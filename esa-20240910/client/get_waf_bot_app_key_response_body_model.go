// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWafBotAppKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v string) *GetWafBotAppKeyResponseBody
	GetAppKey() *string
	SetRequestId(v string) *GetWafBotAppKeyResponseBody
	GetRequestId() *string
}

type GetWafBotAppKeyResponseBody struct {
	// Application key.
	//
	// example:
	//
	// example_appkey
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWafBotAppKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWafBotAppKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetWafBotAppKeyResponseBody) GetAppKey() *string {
	return s.AppKey
}

func (s *GetWafBotAppKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWafBotAppKeyResponseBody) SetAppKey(v string) *GetWafBotAppKeyResponseBody {
	s.AppKey = &v
	return s
}

func (s *GetWafBotAppKeyResponseBody) SetRequestId(v string) *GetWafBotAppKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWafBotAppKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
