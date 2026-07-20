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
	// The status code.
	//
	// example:
	//
	// 0
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The description.
	//
	// example:
	//
	// 成功
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The response data.
	Module *AddressGetResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// The unique identifier of this request.
	//
	// example:
	//
	// 407543AF-****-****-****-9D1AB7218B27
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. Use this parameter to determine the result of the call.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The global trace identifier of the request, typically used for troubleshooting.
	//
	// example:
	//
	// 210bcc3a********d33d7
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
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddressGetResponseBodyModule struct {
	// The redirect URL.
	//
	// example:
	//
	// https://trip-hisv.taobao.com/ding/trustLogin.htm?********aff8-2c2e58da659b
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
