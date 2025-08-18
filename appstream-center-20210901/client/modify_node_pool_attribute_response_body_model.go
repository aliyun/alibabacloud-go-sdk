// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNodePoolAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyNodePoolAttributeResponseBody
	GetCode() *string
	SetMessage(v string) *ModifyNodePoolAttributeResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyNodePoolAttributeResponseBody
	GetRequestId() *string
}

type ModifyNodePoolAttributeResponseBody struct {
	// example:
	//
	// InvalidParameter.PoolId
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// The parameter PoolId is invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNodePoolAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodePoolAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAttributeResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyNodePoolAttributeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyNodePoolAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyNodePoolAttributeResponseBody) SetCode(v string) *ModifyNodePoolAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyNodePoolAttributeResponseBody) SetMessage(v string) *ModifyNodePoolAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyNodePoolAttributeResponseBody) SetRequestId(v string) *ModifyNodePoolAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNodePoolAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
