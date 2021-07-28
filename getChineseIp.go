package chinese_ip

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"net"
	_ "embed"
	"time"
)

func InetNtoA(ip int64) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

func InetAtoN(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}

type CityInfo struct {
	Min string `json:"min"`
	Max string `json:"max"`
}

type CityInfoResult struct {
	MinInt int64
	MaxInt int64
	RandInt int64
	Name string
}

//go:embed chineseIp.json
var g_file_chineseIp []byte

var g_rand_ip []*CityInfoResult

func init()  {
	g_city_info := make(map[string][]CityInfo)
	err := json.Unmarshal(g_file_chineseIp,&g_city_info)
	if err != nil {
		log.Println(err)
		return
	}
	g_rand_ip = make([]*CityInfoResult,0)
	for city, infos := range g_city_info {
		for _, info := range infos {
			minInt,maxInt := InetAtoN(info.Min),InetAtoN(info.Max)
			g_rand_ip = append(g_rand_ip,&CityInfoResult{
				MinInt: minInt,
				MaxInt:maxInt,
				RandInt: maxInt-minInt+1,
				Name: city,
			})
			//log.Println(city,info.Min,info.Max,minInt,maxInt,maxInt-minInt+1)
		}
	}
}
func GetChineseIp() string  {

	rand.Seed(time.Now().UnixNano())
	lens := len(g_rand_ip)
	randInt := rand.Intn(lens)
	ipLong := g_rand_ip[randInt].MinInt + int64(rand.Intn(int(g_rand_ip[randInt].RandInt)))

	return InetNtoA(ipLong)

}