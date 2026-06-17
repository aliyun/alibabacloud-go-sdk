// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceAccessInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCallbackUrl(v string) *GetServiceAccessInfoResponseBody
	GetCallbackUrl() *string
	SetRequestId(v string) *GetServiceAccessInfoResponseBody
	GetRequestId() *string
	SetVerifyCode(v string) *GetServiceAccessInfoResponseBody
	GetVerifyCode() *string
}

type GetServiceAccessInfoResponseBody struct {
	// example:
	//
	// http://xxxxxxxxx
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// dcwedsxcx
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s GetServiceAccessInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceAccessInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceAccessInfoResponseBody) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *GetServiceAccessInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceAccessInfoResponseBody) GetVerifyCode() *string {
	return s.VerifyCode
}

func (s *GetServiceAccessInfoResponseBody) SetCallbackUrl(v string) *GetServiceAccessInfoResponseBody {
	s.CallbackUrl = &v
	return s
}

func (s *GetServiceAccessInfoResponseBody) SetRequestId(v string) *GetServiceAccessInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceAccessInfoResponseBody) SetVerifyCode(v string) *GetServiceAccessInfoResponseBody {
	s.VerifyCode = &v
	return s
}

func (s *GetServiceAccessInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
