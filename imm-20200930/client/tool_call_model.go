// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iToolCall interface {
	dara.Model
	String() string
	GoString() string
	SetFunction(v *FunctionCall) *ToolCall
	GetFunction() *FunctionCall
	SetType(v string) *ToolCall
	GetType() *string
}

type ToolCall struct {
	Function *FunctionCall `json:"Function,omitempty" xml:"Function,omitempty"`
	Type     *string       `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ToolCall) String() string {
	return dara.Prettify(s)
}

func (s ToolCall) GoString() string {
	return s.String()
}

func (s *ToolCall) GetFunction() *FunctionCall {
	return s.Function
}

func (s *ToolCall) GetType() *string {
	return s.Type
}

func (s *ToolCall) SetFunction(v *FunctionCall) *ToolCall {
	s.Function = v
	return s
}

func (s *ToolCall) SetType(v string) *ToolCall {
	s.Type = &v
	return s
}

func (s *ToolCall) Validate() error {
	return dara.Validate(s)
}
