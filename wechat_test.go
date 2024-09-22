package wechat

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/markbates/goth"
	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	p := New("wechat_client_id", "wechat_client_secret", "/foo")

	a.Equal(p.ClientKey, "wechat_client_id")
	a.Equal(p.Secret, "wechat_client_secret")
	a.Equal(p.CallbackURL, "/foo")
}

func Test_Implements_Provider(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	a.Implements((*goth.Provider)(nil), New("wechat_client_id", "wechat_client_secret", "/foo"))
}

func Test_BeginAuth(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	p := New("wechat_client_id", "wechat_client_secret", "/foo")
	session, err := p.BeginAuth("test_state")
	s := session.(*Session)
	a.NoError(err)
	a.Contains(s.AuthURL, "open.weixin.qq.com/connect/qrconnect")
	a.Contains(s.AuthURL, "client_id=wechat_client_id")
	a.Contains(s.AuthURL, "state=test_state")
	a.Contains(s.AuthURL, "redirect_uri=/foo")
}

func Test_FetchUser(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	p := New("wechat_client_id", "wechat_client_secret", "/foo")

	// 创建一个模拟的 HTTP 服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"openid":"test_openid","nickname":"Test User","sex":1,"language":"zh_CN","city":"Shenzhen","province":"Guangdong","country":"CN","headimgurl":"http://example.com/avatar.jpg","privilege":["PRIVILEGE1","PRIVILEGE2"]}`))
	}))
	defer server.Close()

	// 创建一个带有 AccessToken 和 OpenID 的会话
	session := &Session{
		AccessToken: "test_access_token",
		OpenID:      "test_openid",
	}

	// 调用 FetchUser 方法
	user, err := p.FetchUser(session)

	// 验证结果
	a.NoError(err)
	a.Equal("test_openid", user.UserID)
	a.Equal("Test User", user.Name)
	a.Equal("http://example.com/avatar.jpg", user.AvatarURL)
	a.Equal("test_access_token", user.AccessToken)
}

func Test_SessionFromJSON(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	p := New("wechat_client_id", "wechat_client_secret", "/foo")
	sessionJSON := `{"AuthURL":"https://open.weixin.qq.com/connect/qrconnect","AccessToken":"1234567890","RefreshToken":"0987654321","ExpiresAt":"2023-01-01T12:00:00Z","OpenID":"test_openid"}`

	session, err := p.UnmarshalSession(sessionJSON)
	a.NoError(err)

	s, ok := session.(*Session)
	a.True(ok)
	a.Equal("1234567890", s.AccessToken)
	a.Equal("0987654321", s.RefreshToken)
	a.Equal("test_openid", s.OpenID)
}
