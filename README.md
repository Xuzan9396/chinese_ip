## golang 随机获取国内ip
** 需要1.1.6以上的版本才能运行

```go
func Test_getIp(t *testing.T)  {
    ip := GetChineseIp()
    t.Log("ip",ip)
}
```
