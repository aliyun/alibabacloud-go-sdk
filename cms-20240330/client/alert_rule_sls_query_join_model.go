// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlertRuleSlsQueryJoin interface {
	dara.Model
	String() string
	GoString() string
	SetConditions(v []*AlertRuleSlsQueryJoinConditions) *AlertRuleSlsQueryJoin
	GetConditions() []*AlertRuleSlsQueryJoinConditions
	SetType(v string) *AlertRuleSlsQueryJoin
	GetType() *string
}

type AlertRuleSlsQueryJoin struct {
	// The list of join conditions.
	Conditions []*AlertRuleSlsQueryJoinConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	// The type of the collection operation.
	//
	// - CrossJoin: The Cartesian product.
	//
	// - FullJoin: The full join.
	//
	// - InnerJoin: The inner join.
	//
	// - LeftExclude: The left exclusion.
	//
	// - RightExclude: The right exclusion.
	//
	// - LeftJoin: The left join.
	//
	// - RightJoin: The right join.
	//
	// - NoJoin: No merge operation is performed.
	//
	// - Concat: Concatenation.
	//
	// For more information, see https\\://www\\.alibabacloud.com/help/en/sls/user-guide/set-operations.
	//
	// This parameter is required.
	//
	// example:
	//
	// CrossJoin
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AlertRuleSlsQueryJoin) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleSlsQueryJoin) GoString() string {
	return s.String()
}

func (s *AlertRuleSlsQueryJoin) GetConditions() []*AlertRuleSlsQueryJoinConditions {
	return s.Conditions
}

func (s *AlertRuleSlsQueryJoin) GetType() *string {
	return s.Type
}

func (s *AlertRuleSlsQueryJoin) SetConditions(v []*AlertRuleSlsQueryJoinConditions) *AlertRuleSlsQueryJoin {
	s.Conditions = v
	return s
}

func (s *AlertRuleSlsQueryJoin) SetType(v string) *AlertRuleSlsQueryJoin {
	s.Type = &v
	return s
}

func (s *AlertRuleSlsQueryJoin) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AlertRuleSlsQueryJoinConditions struct {
	// The left operand of the condition. The format is $\\<query_idx>.\\<field_name_in_result_set>.
	//
	// example:
	//
	// $0.__topic__
	FirstField *string `json:"firstField,omitempty" xml:"firstField,omitempty"`
	// The comparison operator. Valid values are <, >, ==, !=, <=, and >=.
	//
	// example:
	//
	// ==
	Oper *string `json:"oper,omitempty" xml:"oper,omitempty"`
	// The right operand of the condition. The format is $\\<query_idx>.\\<field_name_in_result_set>.
	//
	// example:
	//
	// $0.__topic__
	SecondField *string `json:"secondField,omitempty" xml:"secondField,omitempty"`
}

func (s AlertRuleSlsQueryJoinConditions) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleSlsQueryJoinConditions) GoString() string {
	return s.String()
}

func (s *AlertRuleSlsQueryJoinConditions) GetFirstField() *string {
	return s.FirstField
}

func (s *AlertRuleSlsQueryJoinConditions) GetOper() *string {
	return s.Oper
}

func (s *AlertRuleSlsQueryJoinConditions) GetSecondField() *string {
	return s.SecondField
}

func (s *AlertRuleSlsQueryJoinConditions) SetFirstField(v string) *AlertRuleSlsQueryJoinConditions {
	s.FirstField = &v
	return s
}

func (s *AlertRuleSlsQueryJoinConditions) SetOper(v string) *AlertRuleSlsQueryJoinConditions {
	s.Oper = &v
	return s
}

func (s *AlertRuleSlsQueryJoinConditions) SetSecondField(v string) *AlertRuleSlsQueryJoinConditions {
	s.SecondField = &v
	return s
}

func (s *AlertRuleSlsQueryJoinConditions) Validate() error {
	return dara.Validate(s)
}
