// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCodeInterpreterSessionInput interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *StartCodeInterpreterSessionInput
	GetName() *string
}

type StartCodeInterpreterSessionInput struct {
	// 代码解释器会话的名称，用于标识和区分不同的会话实例
	//
	// if can be null:
	// true
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s StartCodeInterpreterSessionInput) String() string {
	return dara.Prettify(s)
}

func (s StartCodeInterpreterSessionInput) GoString() string {
	return s.String()
}

func (s *StartCodeInterpreterSessionInput) GetName() *string {
	return s.Name
}

func (s *StartCodeInterpreterSessionInput) SetName(v string) *StartCodeInterpreterSessionInput {
	s.Name = &v
	return s
}

func (s *StartCodeInterpreterSessionInput) Validate() error {
	return dara.Validate(s)
}
