// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentDataSemanticsExampleParameter interface {
	dara.Model
	String() string
	GoString() string
	SetDataType(v string) *AgentDataSemanticsExampleParameter
	GetDataType() *string
	SetDescription(v string) *AgentDataSemanticsExampleParameter
	GetDescription() *string
	SetName(v string) *AgentDataSemanticsExampleParameter
	GetName() *string
	SetValue(v string) *AgentDataSemanticsExampleParameter
	GetValue() *string
}

type AgentDataSemanticsExampleParameter struct {
	// The data type of the parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// date
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The parameter description.
	//
	// example:
	//
	// The lower bound of the transaction date
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameter name.
	//
	// This parameter is required.
	//
	// example:
	//
	// start_date
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The example value of the parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-01-01
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AgentDataSemanticsExampleParameter) String() string {
	return dara.Prettify(s)
}

func (s AgentDataSemanticsExampleParameter) GoString() string {
	return s.String()
}

func (s *AgentDataSemanticsExampleParameter) GetDataType() *string {
	return s.DataType
}

func (s *AgentDataSemanticsExampleParameter) GetDescription() *string {
	return s.Description
}

func (s *AgentDataSemanticsExampleParameter) GetName() *string {
	return s.Name
}

func (s *AgentDataSemanticsExampleParameter) GetValue() *string {
	return s.Value
}

func (s *AgentDataSemanticsExampleParameter) SetDataType(v string) *AgentDataSemanticsExampleParameter {
	s.DataType = &v
	return s
}

func (s *AgentDataSemanticsExampleParameter) SetDescription(v string) *AgentDataSemanticsExampleParameter {
	s.Description = &v
	return s
}

func (s *AgentDataSemanticsExampleParameter) SetName(v string) *AgentDataSemanticsExampleParameter {
	s.Name = &v
	return s
}

func (s *AgentDataSemanticsExampleParameter) SetValue(v string) *AgentDataSemanticsExampleParameter {
	s.Value = &v
	return s
}

func (s *AgentDataSemanticsExampleParameter) Validate() error {
	return dara.Validate(s)
}
