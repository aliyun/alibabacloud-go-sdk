// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFunctionCall interface {
	dara.Model
	String() string
	GoString() string
	SetArguments(v string) *FunctionCall
	GetArguments() *string
	SetName(v string) *FunctionCall
	GetName() *string
}

type FunctionCall struct {
	// The parameters detected by the large language model.
	Arguments *string `json:"Arguments,omitempty" xml:"Arguments,omitempty"`
	// The function name.
	//
	// This parameter is required.
	//
	// example:
	//
	// search_file
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s FunctionCall) String() string {
	return dara.Prettify(s)
}

func (s FunctionCall) GoString() string {
	return s.String()
}

func (s *FunctionCall) GetArguments() *string {
	return s.Arguments
}

func (s *FunctionCall) GetName() *string {
	return s.Name
}

func (s *FunctionCall) SetArguments(v string) *FunctionCall {
	s.Arguments = &v
	return s
}

func (s *FunctionCall) SetName(v string) *FunctionCall {
	s.Name = &v
	return s
}

func (s *FunctionCall) Validate() error {
	return dara.Validate(s)
}
