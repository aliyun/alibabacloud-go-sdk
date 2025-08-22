// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafBotAppKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v string) *DescribeDcdnWafBotAppKeyResponseBody
	GetAppKey() *string
	SetRequestId(v string) *DescribeDcdnWafBotAppKeyResponseBody
	GetRequestId() *string
}

type DescribeDcdnWafBotAppKeyResponseBody struct {
	// The SDK authentication key for the Alibaba Cloud account.
	//
	// example:
	//
	// examp1eapp_key_xxxiuMWTX4Gw
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F2542B96-B535-5BF9-8EEE-1CF11B20CCA8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnWafBotAppKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafBotAppKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafBotAppKeyResponseBody) GetAppKey() *string {
	return s.AppKey
}

func (s *DescribeDcdnWafBotAppKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnWafBotAppKeyResponseBody) SetAppKey(v string) *DescribeDcdnWafBotAppKeyResponseBody {
	s.AppKey = &v
	return s
}

func (s *DescribeDcdnWafBotAppKeyResponseBody) SetRequestId(v string) *DescribeDcdnWafBotAppKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnWafBotAppKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
