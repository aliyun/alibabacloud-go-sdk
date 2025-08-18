// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppInstanceGroupAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyAppInstanceGroupAttributeResponseBody
	GetCode() *string
	SetMessage(v string) *ModifyAppInstanceGroupAttributeResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyAppInstanceGroupAttributeResponseBody
	GetRequestId() *string
}

type ModifyAppInstanceGroupAttributeResponseBody struct {
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

func (s ModifyAppInstanceGroupAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyAppInstanceGroupAttributeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyAppInstanceGroupAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAppInstanceGroupAttributeResponseBody) SetCode(v string) *ModifyAppInstanceGroupAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeResponseBody) SetMessage(v string) *ModifyAppInstanceGroupAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeResponseBody) SetRequestId(v string) *ModifyAppInstanceGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
