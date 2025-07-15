// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnregisterDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UnregisterDeviceResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UnregisterDeviceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UnregisterDeviceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnregisterDeviceResponseBody
	GetRequestId() *string
}

type UnregisterDeviceResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// BA03159C-E808-4FF1-B27E-A61B6E888D7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnregisterDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnregisterDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UnregisterDeviceResponseBody) GetCode() *string {
	return s.Code
}

func (s *UnregisterDeviceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UnregisterDeviceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnregisterDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnregisterDeviceResponseBody) SetCode(v string) *UnregisterDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *UnregisterDeviceResponseBody) SetHttpStatusCode(v int32) *UnregisterDeviceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UnregisterDeviceResponseBody) SetMessage(v string) *UnregisterDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *UnregisterDeviceResponseBody) SetRequestId(v string) *UnregisterDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnregisterDeviceResponseBody) Validate() error {
	return dara.Validate(s)
}
