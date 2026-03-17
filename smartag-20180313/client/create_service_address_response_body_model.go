// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateServiceAddressResponseBody
	GetCode() *string
	SetMessage(v string) *CreateServiceAddressResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateServiceAddressResponseBody
	GetRequestId() *string
}

type CreateServiceAddressResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 324223F3-93D3-4CE4-B26F-66C0C3809922
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServiceAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceAddressResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceAddressResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateServiceAddressResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateServiceAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceAddressResponseBody) SetCode(v string) *CreateServiceAddressResponseBody {
	s.Code = &v
	return s
}

func (s *CreateServiceAddressResponseBody) SetMessage(v string) *CreateServiceAddressResponseBody {
	s.Message = &v
	return s
}

func (s *CreateServiceAddressResponseBody) SetRequestId(v string) *CreateServiceAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
