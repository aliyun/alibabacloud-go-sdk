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
	// The set operation type. Valid values:
	//
	// - CrossJoin: Cartesian product
	//
	// - FullJoin: full join
	//
	// - InnerJoin: inner join
	//
	// - LeftExclude: left exclude
	//
	// - RightExclude: right exclude
	//
	// - LeftJoin: left join
	//
	// - RightJoin: right join
	//
	// - NoJoin: no join
	//
	// - Concat: concatenation
	//
	// For more information, see https://www.alibabacloud.com/help/en/sls/user-guide/set-query-statistics-statement.
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
	// The left operand of the condition, in the format of $<query_idx>.<result set field name>.
	//
	// example:
	//
	// $0.__topic__
	FirstField *string `json:"firstField,omitempty" xml:"firstField,omitempty"`
	// The comparison operator. Valid values: <, >, ==, !=, <=, and >=.
	//
	// example:
	//
	// ==
	Oper *string `json:"oper,omitempty" xml:"oper,omitempty"`
	// The right operand of the condition, in the format of $<query_idx>.<result set field name>.
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
