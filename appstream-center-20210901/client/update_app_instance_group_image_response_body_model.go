// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppInstanceGroupImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAppInstanceGroupImageResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateAppInstanceGroupImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAppInstanceGroupImageResponseBody
	GetRequestId() *string
}

type UpdateAppInstanceGroupImageResponseBody struct {
	// The error code.
	//
	// example:
	//
	// InvalidParameter.ProductType
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The parameter ProductType is invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAppInstanceGroupImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppInstanceGroupImageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppInstanceGroupImageResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAppInstanceGroupImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAppInstanceGroupImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAppInstanceGroupImageResponseBody) SetCode(v string) *UpdateAppInstanceGroupImageResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAppInstanceGroupImageResponseBody) SetMessage(v string) *UpdateAppInstanceGroupImageResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAppInstanceGroupImageResponseBody) SetRequestId(v string) *UpdateAppInstanceGroupImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppInstanceGroupImageResponseBody) Validate() error {
	return dara.Validate(s)
}
