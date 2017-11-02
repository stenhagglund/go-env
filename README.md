# GO-ENV - simple environment to struct parser

## Install

```bash
go get github.com/stenhagglund/go-env
```

## Usage and Examples
Define a struct and call env.Parse with a pointer to it.
```go
type Config struct {
    Host 	 string   `env:"HOST,required,default=localhost"`
    Secret 	 []byte   `env:"SECRET,required,type=byte"`
    Versions []string `env:"VALUES,default=v1,v2"`
    Names 	 []string `env:"VALUES,default=n1:n2:n3,separator=:"`
}

config := &Config{}
if err := env.Parse(&config); err != nil {
    panic(err) // or however you want to handle errors
}
```

Usage is also supported for nested structs
```go
type Config struct {
    Connection struct {
        Host 	 string   `env:"HOST,required,default=localhost"`
    }
    Secret 	 []byte   `env:"SECRET,required,type=byte"`
    Versions []string `env:"VALUES,default=v1,v2"`
    Names 	 []string `env:"VALUES,default=n1:n2:n3,separator=:"`
}

config := &Config{}
if err := env.Parse(&config); err != nil {
    panic(err) // or however you want to handle errors
}
```

## Supported types
- [Boolean types](https://golang.org/ref/spec#Boolean_types)
- [Numeric types](https://golang.org/ref/spec#Numeric_types)
- [String types](https://golang.org/ref/spec#String_types)
- [Slice types](https://golang.org/ref/spec#Slice_types)
- [Struct types](https://golang.org/ref/spec#Struct_types)
- [time.Duration](https://golang.org/pkg/time/#Duration)

## License
The MIT License (MIT) - see LICENSE for more details