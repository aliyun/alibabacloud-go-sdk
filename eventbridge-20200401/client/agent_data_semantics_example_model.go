// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentDataSemanticsExample interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *AgentDataSemanticsExample
	GetDescription() *string
	SetName(v string) *AgentDataSemanticsExample
	GetName() *string
	SetParameters(v []*AgentDataSemanticsExampleParameter) *AgentDataSemanticsExample
	GetParameters() []*AgentDataSemanticsExampleParameter
	SetSQLExpression(v string) *AgentDataSemanticsExample
	GetSQLExpression() *string
}

type AgentDataSemanticsExample struct {
	// The example usage description.
	//
	// example:
	//
	// Use this example when a user asks about high-value customers based on weekly statistics
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The example name.
	//
	// This parameter is required.
	//
	// example:
	//
	// High-value customers with weekly sales exceeding 150
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The SQL example parameter list. A maximum of 20 items are supported.
	Parameters []*AgentDataSemanticsExampleParameter `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The standard SQL example.
	//
	// This parameter is required.
	//
	// example:
	//
	// SELECT customerID FROM sales_transactions
	SQLExpression *string `json:"SQLExpression,omitempty" xml:"SQLExpression,omitempty"`
}

func (s AgentDataSemanticsExample) String() string {
	return dara.Prettify(s)
}

func (s AgentDataSemanticsExample) GoString() string {
	return s.String()
}

func (s *AgentDataSemanticsExample) GetDescription() *string {
	return s.Description
}

func (s *AgentDataSemanticsExample) GetName() *string {
	return s.Name
}

func (s *AgentDataSemanticsExample) GetParameters() []*AgentDataSemanticsExampleParameter {
	return s.Parameters
}

func (s *AgentDataSemanticsExample) GetSQLExpression() *string {
	return s.SQLExpression
}

func (s *AgentDataSemanticsExample) SetDescription(v string) *AgentDataSemanticsExample {
	s.Description = &v
	return s
}

func (s *AgentDataSemanticsExample) SetName(v string) *AgentDataSemanticsExample {
	s.Name = &v
	return s
}

func (s *AgentDataSemanticsExample) SetParameters(v []*AgentDataSemanticsExampleParameter) *AgentDataSemanticsExample {
	s.Parameters = v
	return s
}

func (s *AgentDataSemanticsExample) SetSQLExpression(v string) *AgentDataSemanticsExample {
	s.SQLExpression = &v
	return s
}

func (s *AgentDataSemanticsExample) Validate() error {
	if s.Parameters != nil {
		for _, item := range s.Parameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
