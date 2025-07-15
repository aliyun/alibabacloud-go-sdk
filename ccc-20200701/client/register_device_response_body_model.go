// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RegisterDeviceResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *RegisterDeviceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RegisterDeviceResponseBody
	GetMessage() *string
	SetParams(v []*string) *RegisterDeviceResponseBody
	GetParams() []*string
	SetRequestId(v string) *RegisterDeviceResponseBody
	GetRequestId() *string
}

type RegisterDeviceResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterDeviceResponseBody) GetCode() *string {
	return s.Code
}

func (s *RegisterDeviceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RegisterDeviceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RegisterDeviceResponseBody) GetParams() []*string {
	return s.Params
}

func (s *RegisterDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterDeviceResponseBody) SetCode(v string) *RegisterDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *RegisterDeviceResponseBody) SetHttpStatusCode(v int32) *RegisterDeviceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RegisterDeviceResponseBody) SetMessage(v string) *RegisterDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *RegisterDeviceResponseBody) SetParams(v []*string) *RegisterDeviceResponseBody {
	s.Params = v
	return s
}

func (s *RegisterDeviceResponseBody) SetRequestId(v string) *RegisterDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterDeviceResponseBody) Validate() error {
	return dara.Validate(s)
}
