// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskGraphFlow interface {
	dara.Model
	String() string
	GoString() string
	SetFlowRuleScoreType(v int32) *TaskGraphFlow
	GetFlowRuleScoreType() *int32
	SetId(v int64) *TaskGraphFlow
	GetId() *int64
	SetNodes(v []*GraphFlowNode) *TaskGraphFlow
	GetNodes() []*GraphFlowNode
	SetRid(v int64) *TaskGraphFlow
	GetRid() *int64
	SetRuleName(v string) *TaskGraphFlow
	GetRuleName() *string
	SetShowProperties(v string) *TaskGraphFlow
	GetShowProperties() *string
	SetSkipWhenFirstSessionNodeMiss(v bool) *TaskGraphFlow
	GetSkipWhenFirstSessionNodeMiss() *bool
}

type TaskGraphFlow struct {
	FlowRuleScoreType            *int32           `json:"FlowRuleScoreType,omitempty" xml:"FlowRuleScoreType,omitempty"`
	Id                           *int64           `json:"Id,omitempty" xml:"Id,omitempty"`
	Nodes                        []*GraphFlowNode `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	Rid                          *int64           `json:"Rid,omitempty" xml:"Rid,omitempty"`
	RuleName                     *string          `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	ShowProperties               *string          `json:"ShowProperties,omitempty" xml:"ShowProperties,omitempty"`
	SkipWhenFirstSessionNodeMiss *bool            `json:"SkipWhenFirstSessionNodeMiss,omitempty" xml:"SkipWhenFirstSessionNodeMiss,omitempty"`
}

func (s TaskGraphFlow) String() string {
	return dara.Prettify(s)
}

func (s TaskGraphFlow) GoString() string {
	return s.String()
}

func (s *TaskGraphFlow) GetFlowRuleScoreType() *int32 {
	return s.FlowRuleScoreType
}

func (s *TaskGraphFlow) GetId() *int64 {
	return s.Id
}

func (s *TaskGraphFlow) GetNodes() []*GraphFlowNode {
	return s.Nodes
}

func (s *TaskGraphFlow) GetRid() *int64 {
	return s.Rid
}

func (s *TaskGraphFlow) GetRuleName() *string {
	return s.RuleName
}

func (s *TaskGraphFlow) GetShowProperties() *string {
	return s.ShowProperties
}

func (s *TaskGraphFlow) GetSkipWhenFirstSessionNodeMiss() *bool {
	return s.SkipWhenFirstSessionNodeMiss
}

func (s *TaskGraphFlow) SetFlowRuleScoreType(v int32) *TaskGraphFlow {
	s.FlowRuleScoreType = &v
	return s
}

func (s *TaskGraphFlow) SetId(v int64) *TaskGraphFlow {
	s.Id = &v
	return s
}

func (s *TaskGraphFlow) SetNodes(v []*GraphFlowNode) *TaskGraphFlow {
	s.Nodes = v
	return s
}

func (s *TaskGraphFlow) SetRid(v int64) *TaskGraphFlow {
	s.Rid = &v
	return s
}

func (s *TaskGraphFlow) SetRuleName(v string) *TaskGraphFlow {
	s.RuleName = &v
	return s
}

func (s *TaskGraphFlow) SetShowProperties(v string) *TaskGraphFlow {
	s.ShowProperties = &v
	return s
}

func (s *TaskGraphFlow) SetSkipWhenFirstSessionNodeMiss(v bool) *TaskGraphFlow {
	s.SkipWhenFirstSessionNodeMiss = &v
	return s
}

func (s *TaskGraphFlow) Validate() error {
	return dara.Validate(s)
}
