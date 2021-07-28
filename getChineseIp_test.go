package chinese_ip

import "testing"

func Test_getIp(t *testing.T)  {
	ip := GetChineseIp()
	t.Log("ip",ip)
}