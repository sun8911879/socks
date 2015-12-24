# socks

[socks](https://github.com/sun8911879/socks) is a simple Go library to toggle on and
off socks for Mac OS X. It will extract
a helper tool and use it to actually chage socks setting.

```go
socks.EnsureHelperToolPresent(fullPath, prompt, iconFullPath)
socks.On(SocksIP,SocksPort string)
socks.Off()
```

See 'example/main.go' for detailed usage.

### Embedding socks-cmd

pac uses binaries from the [socks-cmd](https://github.com/sun8911879/socks-cmd) project.

### Thanks lantern project
[pac-cmd](https://github.com/getlantern/pac-cmd)
[pac](https://github.com/getlantern/pac)