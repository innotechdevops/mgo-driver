# mgo-driver

## Install

```
$ go get github.com/innotechdevops/mgo-driver
```

## How to use

- Wtih env

```golang
driver := mgodriver.New(mgodriver.ConfigEnv())
```

- With config

```golang
driver := mgodriver.New(mgodriver.Config{
    User:         os.Getenv("MARIA_USER"),
    Pass:         os.Getenv("MARIA_PASS"),
    Host:         os.Getenv("MARIA_HOST"),
    DatabaseName: os.Getenv("MARIA_DATABASE"),
    Port:         mgodriver.DefaultPort,
})
```