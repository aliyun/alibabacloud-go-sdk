// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodeGroupStateChangeReason interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *NodeGroupStateChangeReason
	GetCode() *string
	SetMessage(v string) *NodeGroupStateChangeReason
	GetMessage() *string
}

type NodeGroupStateChangeReason struct {
	// 状态码。
	//
	// example:
	//
	// MissingParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 描述信息。
	//
	// example:
	//
	// The instance type is required.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s NodeGroupStateChangeReason) String() string {
	return dara.Prettify(s)
}

func (s NodeGroupStateChangeReason) GoString() string {
	return s.String()
}

func (s *NodeGroupStateChangeReason) GetCode() *string {
	return s.Code
}

func (s *NodeGroupStateChangeReason) GetMessage() *string {
	return s.Message
}

func (s *NodeGroupStateChangeReason) SetCode(v string) *NodeGroupStateChangeReason {
	s.Code = &v
	return s
}

func (s *NodeGroupStateChangeReason) SetMessage(v string) *NodeGroupStateChangeReason {
	s.Message = &v
	return s
}

func (s *NodeGroupStateChangeReason) Validate() error {
	return dara.Validate(s)
}
