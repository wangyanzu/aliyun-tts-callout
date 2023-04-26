package main

import (
	"aliyun_tts_callout/callout"
	"flag"
	"log"
	"strings"
)

func main() {
	region := flag.String("ar", "cn-hangzhou", "aliyun regionId")
	accessKeyId := flag.String("ak", "", "aliyun accessKeyId")
	accessKeySecret := flag.String("as", "", "aliyun accessKeySecret")
	ttsCode := flag.String("t", "", "tts code, TTS_xxxxxxxxxx")
	showNumber := flag.String("t", "", "CalledShowNumber")
	param := flag.String("b", `{"level":"未知","value":"0","rulename":"xx告警"}`, "tts param")
	phone := flag.String("p", "", "phone list, 135xxx,136xxx,137xxx")
	dryrun := flag.Bool("d", false, "do nothing")

	flag.Parse()
	if *phone == "" {
		panic("phone not found")
	}
	phones := strings.Split(*phone, ",")

	if *dryrun {
		log.Printf("Dryrun")
	}
	callout.Callout(*region, *accessKeyId, *accessKeySecret, *showNumber, *ttsCode, *param, phones, *dryrun)
}
