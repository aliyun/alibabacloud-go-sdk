// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBotAppKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v string) *DescribeBotAppKeyResponseBody
	GetAppKey() *string
	SetRequestId(v string) *DescribeBotAppKeyResponseBody
	GetRequestId() *string
}

type DescribeBotAppKeyResponseBody struct {
	// AppKey。
	//
	// example:
	//
	// N1Kiv3AGZm******
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// 0C4ADFD4-5B7D-591D-A607-A45C*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBotAppKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBotAppKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBotAppKeyResponseBody) GetAppKey() *string {
	return s.AppKey
}

func (s *DescribeBotAppKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBotAppKeyResponseBody) SetAppKey(v string) *DescribeBotAppKeyResponseBody {
	s.AppKey = &v
	return s
}

func (s *DescribeBotAppKeyResponseBody) SetRequestId(v string) *DescribeBotAppKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBotAppKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
