// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentDataSemanticsMetric interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *AgentDataSemanticsMetric
	GetDescription() *string
	SetName(v string) *AgentDataSemanticsMetric
	GetName() *string
	SetSQLExpression(v string) *AgentDataSemanticsMetric
	GetSQLExpression() *string
	SetSynonyms(v []*string) *AgentDataSemanticsMetric
	GetSynonyms() []*string
	SetType(v string) *AgentDataSemanticsMetric
	GetType() *string
}

type AgentDataSemanticsMetric struct {
	// The usage description.
	//
	// example:
	//
	// Use totalPrice to calculate total sales for sales performance comparison
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the SQL expression.
	//
	// This parameter is required.
	//
	// example:
	//
	// TotalSales
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The SQL expression.
	//
	// This parameter is required.
	//
	// example:
	//
	// SUM(sales_transactions.totalPrice)
	SQLExpression *string `json:"SQLExpression,omitempty" xml:"SQLExpression,omitempty"`
	// The list of synonyms. A maximum of 10 items are supported. Each item can contain up to 64 characters.
	//
	// example:
	//
	// ["Sales","Revenue","GMV"]
	Synonyms []*string `json:"Synonyms,omitempty" xml:"Synonyms,omitempty" type:"Repeated"`
	// The type of the SQL expression.
	//
	// This parameter is required.
	//
	// example:
	//
	// measure
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AgentDataSemanticsMetric) String() string {
	return dara.Prettify(s)
}

func (s AgentDataSemanticsMetric) GoString() string {
	return s.String()
}

func (s *AgentDataSemanticsMetric) GetDescription() *string {
	return s.Description
}

func (s *AgentDataSemanticsMetric) GetName() *string {
	return s.Name
}

func (s *AgentDataSemanticsMetric) GetSQLExpression() *string {
	return s.SQLExpression
}

func (s *AgentDataSemanticsMetric) GetSynonyms() []*string {
	return s.Synonyms
}

func (s *AgentDataSemanticsMetric) GetType() *string {
	return s.Type
}

func (s *AgentDataSemanticsMetric) SetDescription(v string) *AgentDataSemanticsMetric {
	s.Description = &v
	return s
}

func (s *AgentDataSemanticsMetric) SetName(v string) *AgentDataSemanticsMetric {
	s.Name = &v
	return s
}

func (s *AgentDataSemanticsMetric) SetSQLExpression(v string) *AgentDataSemanticsMetric {
	s.SQLExpression = &v
	return s
}

func (s *AgentDataSemanticsMetric) SetSynonyms(v []*string) *AgentDataSemanticsMetric {
	s.Synonyms = v
	return s
}

func (s *AgentDataSemanticsMetric) SetType(v string) *AgentDataSemanticsMetric {
	s.Type = &v
	return s
}

func (s *AgentDataSemanticsMetric) Validate() error {
	return dara.Validate(s)
}
