// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RegisterDevicesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *RegisterDevicesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RegisterDevicesResponseBody
	GetMessage() *string
	SetParams(v []*string) *RegisterDevicesResponseBody
	GetParams() []*string
	SetRequestId(v string) *RegisterDevicesResponseBody
	GetRequestId() *string
}

type RegisterDevicesResponseBody struct {
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
	// BA03159C-E808-4FF1-B27E-A61B6E888D7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterDevicesResponseBody) GetCode() *string {
	return s.Code
}

func (s *RegisterDevicesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RegisterDevicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RegisterDevicesResponseBody) GetParams() []*string {
	return s.Params
}

func (s *RegisterDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterDevicesResponseBody) SetCode(v string) *RegisterDevicesResponseBody {
	s.Code = &v
	return s
}

func (s *RegisterDevicesResponseBody) SetHttpStatusCode(v int32) *RegisterDevicesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RegisterDevicesResponseBody) SetMessage(v string) *RegisterDevicesResponseBody {
	s.Message = &v
	return s
}

func (s *RegisterDevicesResponseBody) SetParams(v []*string) *RegisterDevicesResponseBody {
	s.Params = v
	return s
}

func (s *RegisterDevicesResponseBody) SetRequestId(v string) *RegisterDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterDevicesResponseBody) Validate() error {
	return dara.Validate(s)
}
