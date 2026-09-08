// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentDataSemanticsJoin interface {
	dara.Model
	String() string
	GoString() string
	SetCondition(v *AgentDataSemanticsJoinCondition) *AgentDataSemanticsJoin
	GetCondition() *AgentDataSemanticsJoinCondition
	SetDescription(v string) *AgentDataSemanticsJoin
	GetDescription() *string
	SetLeftTable(v string) *AgentDataSemanticsJoin
	GetLeftTable() *string
	SetRelationshipType(v string) *AgentDataSemanticsJoin
	GetRelationshipType() *string
	SetRightTable(v string) *AgentDataSemanticsJoin
	GetRightTable() *string
}

type AgentDataSemanticsJoin struct {
	// The join condition.
	//
	// This parameter is required.
	Condition *AgentDataSemanticsJoinCondition `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The join usage description.
	//
	// example:
	//
	// Use this join when analyzing customer attributes associated with sales transactions
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The full name of the left table.
	//
	// This parameter is required.
	//
	// example:
	//
	// samples.bakehouse.sales_transactions
	LeftTable *string `json:"LeftTable,omitempty" xml:"LeftTable,omitempty"`
	// The table relationship type.
	//
	// example:
	//
	// many_to_one
	RelationshipType *string `json:"RelationshipType,omitempty" xml:"RelationshipType,omitempty"`
	// The full name of the right table.
	//
	// This parameter is required.
	//
	// example:
	//
	// samples.bakehouse.customers
	RightTable *string `json:"RightTable,omitempty" xml:"RightTable,omitempty"`
}

func (s AgentDataSemanticsJoin) String() string {
	return dara.Prettify(s)
}

func (s AgentDataSemanticsJoin) GoString() string {
	return s.String()
}

func (s *AgentDataSemanticsJoin) GetCondition() *AgentDataSemanticsJoinCondition {
	return s.Condition
}

func (s *AgentDataSemanticsJoin) GetDescription() *string {
	return s.Description
}

func (s *AgentDataSemanticsJoin) GetLeftTable() *string {
	return s.LeftTable
}

func (s *AgentDataSemanticsJoin) GetRelationshipType() *string {
	return s.RelationshipType
}

func (s *AgentDataSemanticsJoin) GetRightTable() *string {
	return s.RightTable
}

func (s *AgentDataSemanticsJoin) SetCondition(v *AgentDataSemanticsJoinCondition) *AgentDataSemanticsJoin {
	s.Condition = v
	return s
}

func (s *AgentDataSemanticsJoin) SetDescription(v string) *AgentDataSemanticsJoin {
	s.Description = &v
	return s
}

func (s *AgentDataSemanticsJoin) SetLeftTable(v string) *AgentDataSemanticsJoin {
	s.LeftTable = &v
	return s
}

func (s *AgentDataSemanticsJoin) SetRelationshipType(v string) *AgentDataSemanticsJoin {
	s.RelationshipType = &v
	return s
}

func (s *AgentDataSemanticsJoin) SetRightTable(v string) *AgentDataSemanticsJoin {
	s.RightTable = &v
	return s
}

func (s *AgentDataSemanticsJoin) Validate() error {
	if s.Condition != nil {
		if err := s.Condition.Validate(); err != nil {
			return err
		}
	}
	return nil
}
