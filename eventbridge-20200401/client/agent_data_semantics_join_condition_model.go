// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentDataSemanticsJoinCondition interface {
	dara.Model
	String() string
	GoString() string
	SetLeftColumn(v string) *AgentDataSemanticsJoinCondition
	GetLeftColumn() *string
	SetMode(v string) *AgentDataSemanticsJoinCondition
	GetMode() *string
	SetRightColumn(v string) *AgentDataSemanticsJoinCondition
	GetRightColumn() *string
	SetSQLExpression(v string) *AgentDataSemanticsJoinCondition
	GetSQLExpression() *string
}

type AgentDataSemanticsJoinCondition struct {
	// The left table field name in form mode.
	//
	// example:
	//
	// customerID
	LeftColumn *string `json:"LeftColumn,omitempty" xml:"LeftColumn,omitempty"`
	// The conditional expression method.
	//
	// This parameter is required.
	//
	// example:
	//
	// form
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The right table field name in form mode.
	//
	// example:
	//
	// customerID
	RightColumn *string `json:"RightColumn,omitempty" xml:"RightColumn,omitempty"`
	// The join SQL expression in SQL mode.
	//
	// example:
	//
	// sales_transactions.customerID = customers.customerID
	SQLExpression *string `json:"SQLExpression,omitempty" xml:"SQLExpression,omitempty"`
}

func (s AgentDataSemanticsJoinCondition) String() string {
	return dara.Prettify(s)
}

func (s AgentDataSemanticsJoinCondition) GoString() string {
	return s.String()
}

func (s *AgentDataSemanticsJoinCondition) GetLeftColumn() *string {
	return s.LeftColumn
}

func (s *AgentDataSemanticsJoinCondition) GetMode() *string {
	return s.Mode
}

func (s *AgentDataSemanticsJoinCondition) GetRightColumn() *string {
	return s.RightColumn
}

func (s *AgentDataSemanticsJoinCondition) GetSQLExpression() *string {
	return s.SQLExpression
}

func (s *AgentDataSemanticsJoinCondition) SetLeftColumn(v string) *AgentDataSemanticsJoinCondition {
	s.LeftColumn = &v
	return s
}

func (s *AgentDataSemanticsJoinCondition) SetMode(v string) *AgentDataSemanticsJoinCondition {
	s.Mode = &v
	return s
}

func (s *AgentDataSemanticsJoinCondition) SetRightColumn(v string) *AgentDataSemanticsJoinCondition {
	s.RightColumn = &v
	return s
}

func (s *AgentDataSemanticsJoinCondition) SetSQLExpression(v string) *AgentDataSemanticsJoinCondition {
	s.SQLExpression = &v
	return s
}

func (s *AgentDataSemanticsJoinCondition) Validate() error {
	return dara.Validate(s)
}
