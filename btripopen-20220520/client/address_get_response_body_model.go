// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddressGetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddressGetResponseBody
	GetCode() *string
	SetMessage(v string) *AddressGetResponseBody
	GetMessage() *string
	SetModule(v *AddressGetResponseBodyModule) *AddressGetResponseBody
	GetModule() *AddressGetResponseBodyModule
	SetRequestId(v string) *AddressGetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddressGetResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *AddressGetResponseBody
	GetTraceId() *string
}

type AddressGetResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                       `json:"message,omitempty" xml:"message,omitempty"`
	Module  *AddressGetResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// 407543AF-2BD9-5890-BD92-9D1AB7218B27
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 210bcc3a16583004579056128d33d7
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s AddressGetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddressGetResponseBody) GoString() string {
	return s.String()
}

func (s *AddressGetResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddressGetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddressGetResponseBody) GetModule() *AddressGetResponseBodyModule {
	return s.Module
}

func (s *AddressGetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddressGetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddressGetResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *AddressGetResponseBody) SetCode(v string) *AddressGetResponseBody {
	s.Code = &v
	return s
}

func (s *AddressGetResponseBody) SetMessage(v string) *AddressGetResponseBody {
	s.Message = &v
	return s
}

func (s *AddressGetResponseBody) SetModule(v *AddressGetResponseBodyModule) *AddressGetResponseBody {
	s.Module = v
	return s
}

func (s *AddressGetResponseBody) SetRequestId(v string) *AddressGetResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddressGetResponseBody) SetSuccess(v bool) *AddressGetResponseBody {
	s.Success = &v
	return s
}

func (s *AddressGetResponseBody) SetTraceId(v string) *AddressGetResponseBody {
	s.TraceId = &v
	return s
}

func (s *AddressGetResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddressGetResponseBodyModule struct {
	// example:
	//
	// https://trip-hisv.taobao.com/ding/trustLogin.htm?redirectUrl=https%3A%2F%2Fmarket.m.taobao.com%2Fapp%2Fbtrip-fe%2Frx-alitrip-main%2Fhome.html%3Ffpt%3DbIdentify%2528dingtalk.isv.h5.home%2529%26corpId%3Dding3f9797e277423f14a1320dcb25e91351%26dingUserId%3Dmanager9302%26dingAppId%3D1692%26fit%3Dtrue&token=b73e0b9e-d25a-40f0-aff8-2c2e58da659b
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s AddressGetResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s AddressGetResponseBodyModule) GoString() string {
	return s.String()
}

func (s *AddressGetResponseBodyModule) GetUrl() *string {
	return s.Url
}

func (s *AddressGetResponseBodyModule) SetUrl(v string) *AddressGetResponseBodyModule {
	s.Url = &v
	return s
}

func (s *AddressGetResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
