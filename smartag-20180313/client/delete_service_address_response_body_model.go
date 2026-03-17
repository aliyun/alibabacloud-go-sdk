// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteServiceAddressResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteServiceAddressResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteServiceAddressResponseBody
	GetRequestId() *string
}

type DeleteServiceAddressResponseBody struct {
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

func (s DeleteServiceAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceAddressResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteServiceAddressResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteServiceAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteServiceAddressResponseBody) SetCode(v string) *DeleteServiceAddressResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteServiceAddressResponseBody) SetMessage(v string) *DeleteServiceAddressResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteServiceAddressResponseBody) SetRequestId(v string) *DeleteServiceAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServiceAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
