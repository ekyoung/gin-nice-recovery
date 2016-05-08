package nice

import (
    "io"
    "log"
    "net/http"
    "net/http/httputil"

    "github.com/go-errors/errors"

    "github.com/gin-gonic/gin"
)

func Recovery(name string, obj interface{}) gin.HandlerFunc {
    return RecoveryWithWriter(name, obj, gin.DefaultErrorWriter)
}

func RecoveryWithWriter(name string, obj interface{}, out io.Writer) gin.HandlerFunc {
    var logger *log.Logger
    if out != nil {
        logger = log.New(out, "\n\n\x1b[31m", log.LstdFlags)
    }

    return func(c *gin.Context) {
        defer func() {
            if err := recover(); err != nil {
                if logger != nil {
                    httprequest, _ := httputil.DumpRequest(c.Request, false)
                    goErr := errors.Wrap(err, 3)
                    reset := string([]byte{27, 91, 48, 109})
                    logger.Printf("[Nice Recovery] panic recovered:\n\n%s%s\n\n%s%s", httprequest, goErr.Error(), goErr.Stack(), reset)
                }

                c.HTML(http.StatusInternalServerError, name, obj)
            }
        }()
        c.Next() // execute all the handlers
    }
}
