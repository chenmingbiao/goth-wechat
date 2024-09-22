# goth-wechat

This package implements the WeChat OAuth2 provider for [Goth](https://github.com/markbates/goth), an authentication library for Go web applications.

## Installation

```bash
go get github.com/chenmingbiao/goth-wechat
```

## Usage

To use this provider with Goth, you need to set it up in your application:

```go
package main

import (
    "github.com/markbates/goth"
    "github.com/chenmingbiao/goth-wechat"
)

func init() {
    goth.UseProviders(
        wechat.New(os.Getenv("WECHAT_KEY"), os.Getenv("WECHAT_SECRET"), "http://localhost:3000/auth/wechat/callback"),
    )
}
```

Make sure to set the `WECHAT_KEY` and `WECHAT_SECRET` environment variables with your WeChat application credentials.

## Authentication

The WeChat authentication flow follows these steps:

1. Redirect the user to the WeChat authorization page
2. WeChat redirects back to your site with a code
3. Your site exchanges the code for an access token
4. Fetch the user information using the access token

For detailed implementation, please refer to the Goth documentation and examples.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

* Thanks to the [Goth](https://github.com/markbates/goth) project for providing the authentication framework.
* This provider is based on the WeChat OAuth 2.0 documentation.