# zerologForEcho

middleware that provides logging and panic recovery for [Echo](https://github.com/labstack/echo) using [Zerolog](https://github.com/rs/zerolog).
same as [https://github.com/tomruk/zap4echo](https://github.com/tomruk/zap4echo) but using zerolog as logger

<a name="section-1"></a> [Usage](#section-1)

```
logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

e.Use(
  zerologForEcho.RecoverWithConfig(zerologForEcho.RecoverConfig{}, &logger)
)
```
