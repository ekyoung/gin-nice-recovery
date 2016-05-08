# gin-nice-recovery

[Gin](https://gin-gonic.github.io/gin/) middleware to render a nice looking error page when recovering from a panic.

## Why?

The default `gin.Recovery()` middleware leaves the user looking a blank white page. This middleware renders the
specified html template. It logs the same HTTP request information and stack trace as the default middleware.

## Installation

```bash
$ go get github.com/ekyoung/gin-nice-recovery
```

## Usage

```go
import (
    "github.com/ekyoung/gin-nice-recovery"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.New()       // gin.Default() installs gin.Recovery() so use gin.New() instead
    router.Use(gin.Logger())  // Install the default logger, not required
    
    // Install nice.Recovery, passing the name of the html template to render, and data to use
    router.Use(nice.Recovery("error.tmpl", gin.H{
        "title": "Error",
    }))    
    
    // Load templates as usual
    router.LoadHTMLFiles("error.tmpl")

    router.GET("/", func(c *gin.Context) {
        panic("Doh!")
    })

    router.Run(":8080")
}
```