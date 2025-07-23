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
	Conditions []*AlertRuleSlsQueryJoinConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	// 集合操作类型。
	//
	//   ● CrossJoin： 笛卡尔积
	//
	//   ● FullJoin：全联
	//
	//   ● InnerJoin：内联
	//
	//   ● LeftExclude： 左斥
	//
	//   ● RightExclude：右斥
	//
	//   ● LeftJoin：左联
	//
	//   ● RightJoin：右联
	//
	//   ● NoJoin：不合并
	//
	//   ● Concat： 拼接
	//
	//   https://help.aliyun.com/zh/sls/user-guide/set-query-statistics-statement
	//
	// This parameter is required.
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
	return dara.Validate(s)
}

type AlertRuleSlsQueryJoinConditions struct {
	// 条件的左操作参数，格式为$<query_idx>.<结果集字段名>
	FirstField *string `json:"firstField,omitempty" xml:"firstField,omitempty"`
	// <, >, ==, !=, <=, >=
	Oper *string `json:"oper,omitempty" xml:"oper,omitempty"`
	// 条件的右操作参数，格式为$<query_idx>.<结果集字段名>
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
