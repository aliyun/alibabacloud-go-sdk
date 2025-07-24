// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStateChangeReason interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StateChangeReason
	GetCode() *string
	SetMessage(v string) *StateChangeReason
	GetMessage() *string
}

type StateChangeReason struct {
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

func (s StateChangeReason) String() string {
	return dara.Prettify(s)
}

func (s StateChangeReason) GoString() string {
	return s.String()
}

func (s *StateChangeReason) GetCode() *string {
	return s.Code
}

func (s *StateChangeReason) GetMessage() *string {
	return s.Message
}

func (s *StateChangeReason) SetCode(v string) *StateChangeReason {
	s.Code = &v
	return s
}

func (s *StateChangeReason) SetMessage(v string) *StateChangeReason {
	s.Message = &v
	return s
}

func (s *StateChangeReason) Validate() error {
	return dara.Validate(s)
}
