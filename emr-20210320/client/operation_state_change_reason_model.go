// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperationStateChangeReason interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *OperationStateChangeReason
	GetCode() *string
	SetMessage(v string) *OperationStateChangeReason
	GetMessage() *string
}

type OperationStateChangeReason struct {
	// 状态码。
	//
	// example:
	//
	// OutOfStock
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 状态变化信息。
	//
	// example:
	//
	// The requested resource is sold out in the specified zone, try other types of resources or other regions and zones.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s OperationStateChangeReason) String() string {
	return dara.Prettify(s)
}

func (s OperationStateChangeReason) GoString() string {
	return s.String()
}

func (s *OperationStateChangeReason) GetCode() *string {
	return s.Code
}

func (s *OperationStateChangeReason) GetMessage() *string {
	return s.Message
}

func (s *OperationStateChangeReason) SetCode(v string) *OperationStateChangeReason {
	s.Code = &v
	return s
}

func (s *OperationStateChangeReason) SetMessage(v string) *OperationStateChangeReason {
	s.Message = &v
	return s
}

func (s *OperationStateChangeReason) Validate() error {
	return dara.Validate(s)
}
