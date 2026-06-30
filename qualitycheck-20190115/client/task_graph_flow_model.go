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
	// Flow scoring logic settings
	//
	// example:
	//
	// 1
	FlowRuleScoreType *int32 `json:"FlowRuleScoreType,omitempty" xml:"FlowRuleScoreType,omitempty"`
	// Canvas ID of the flow
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// All nodes in the flow
	Nodes []*GraphFlowNode `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// Rule ID
	//
	// example:
	//
	// 1
	Rid *int64 `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// Rule Name
	//
	// example:
	//
	// 违规
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Used for frontend display
	//
	// example:
	//
	// {}
	ShowProperties *string `json:"ShowProperties,omitempty" xml:"ShowProperties,omitempty"`
	// Skip if the first session node is not hit
	//
	// example:
	//
	// true
	SkipWhenFirstSessionNodeMiss *bool `json:"SkipWhenFirstSessionNodeMiss,omitempty" xml:"SkipWhenFirstSessionNodeMiss,omitempty"`
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
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
