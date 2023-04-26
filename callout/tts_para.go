package callout

import (
	"log"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dyvmsapi"
)

func Callout(region, accessKey, accessSecret, calledShowNumber, ttsCode, param string, phones []string, dryrun bool) {
	var ttsClient, _ = dyvmsapi.NewClientWithAccessKey(region, accessKey, accessSecret)
	request := dyvmsapi.CreateSingleCallByTtsRequest()
	request.Scheme = "https"

	request.TtsCode = ttsCode
	request.CalledShowNumber = calledShowNumber
	request.TtsParam = param

	request.PlayTimes = requests.NewInteger(3)
	request.Volume = requests.NewInteger(100)
	request.Speed = "0"

	for _, phone := range phones {
		log.Printf("callout %s, %s, %s\n", phone, ttsCode, param)
		request.CalledNumber = phone
		if !dryrun {
			response, err := ttsClient.SingleCallByTts(request)
			if err != nil {
				log.Printf("callout request err: %s \n", err.Error())
				return
			}
			log.Printf("callout %s, response is %v\n", phone, response)
		}
		time.Sleep(1 * time.Second)
	}
}
