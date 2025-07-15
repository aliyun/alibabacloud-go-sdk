// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddDevicesResponseBody
	GetCode() *string
	SetMessage(v string) *AddDevicesResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddDevicesResponseBody
	GetRequestId() *string
}

type AddDevicesResponseBody struct {
	// The execution result. If the request was successful, `success` is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned error message. This parameter is not returned if the value of Code is `success`.
	//
	// example:
	//
	// The parameter is not specified.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A87DBB05-653A-5E4B-B72B-5F4A1E07****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *AddDevicesResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddDevicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDevicesResponseBody) SetCode(v string) *AddDevicesResponseBody {
	s.Code = &v
	return s
}

func (s *AddDevicesResponseBody) SetMessage(v string) *AddDevicesResponseBody {
	s.Message = &v
	return s
}

func (s *AddDevicesResponseBody) SetRequestId(v string) *AddDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDevicesResponseBody) Validate() error {
	return dara.Validate(s)
}
