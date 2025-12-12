// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGraphFlowNode interface {
	dara.Model
	String() string
	GoString() string
	SetConditions(v []*ConditionBasicInfo) *GraphFlowNode
	GetConditions() []*ConditionBasicInfo
	SetContent(v string) *GraphFlowNode
	GetContent() *string
	SetId(v int64) *GraphFlowNode
	GetId() *int64
	SetIndex(v int32) *GraphFlowNode
	GetIndex() *int32
	SetName(v string) *GraphFlowNode
	GetName() *string
	SetNextNodes(v []*GraphFlowNodeNextNodes) *GraphFlowNode
	GetNextNodes() []*GraphFlowNodeNextNodes
	SetNodeType(v string) *GraphFlowNode
	GetNodeType() *string
	SetProperties(v *GraphFlowNodeProperties) *GraphFlowNode
	GetProperties() *GraphFlowNodeProperties
	SetRid(v int64) *GraphFlowNode
	GetRid() *int64
	SetUseConditions(v bool) *GraphFlowNode
	GetUseConditions() *bool
}

type GraphFlowNode struct {
	Conditions    []*ConditionBasicInfo     `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	Content       *string                   `json:"Content,omitempty" xml:"Content,omitempty"`
	Id            *int64                    `json:"Id,omitempty" xml:"Id,omitempty"`
	Index         *int32                    `json:"Index,omitempty" xml:"Index,omitempty"`
	Name          *string                   `json:"Name,omitempty" xml:"Name,omitempty"`
	NextNodes     []*GraphFlowNodeNextNodes `json:"NextNodes,omitempty" xml:"NextNodes,omitempty" type:"Repeated"`
	NodeType      *string                   `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	Properties    *GraphFlowNodeProperties  `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	Rid           *int64                    `json:"Rid,omitempty" xml:"Rid,omitempty"`
	UseConditions *bool                     `json:"UseConditions,omitempty" xml:"UseConditions,omitempty"`
}

func (s GraphFlowNode) String() string {
	return dara.Prettify(s)
}

func (s GraphFlowNode) GoString() string {
	return s.String()
}

func (s *GraphFlowNode) GetConditions() []*ConditionBasicInfo {
	return s.Conditions
}

func (s *GraphFlowNode) GetContent() *string {
	return s.Content
}

func (s *GraphFlowNode) GetId() *int64 {
	return s.Id
}

func (s *GraphFlowNode) GetIndex() *int32 {
	return s.Index
}

func (s *GraphFlowNode) GetName() *string {
	return s.Name
}

func (s *GraphFlowNode) GetNextNodes() []*GraphFlowNodeNextNodes {
	return s.NextNodes
}

func (s *GraphFlowNode) GetNodeType() *string {
	return s.NodeType
}

func (s *GraphFlowNode) GetProperties() *GraphFlowNodeProperties {
	return s.Properties
}

func (s *GraphFlowNode) GetRid() *int64 {
	return s.Rid
}

func (s *GraphFlowNode) GetUseConditions() *bool {
	return s.UseConditions
}

func (s *GraphFlowNode) SetConditions(v []*ConditionBasicInfo) *GraphFlowNode {
	s.Conditions = v
	return s
}

func (s *GraphFlowNode) SetContent(v string) *GraphFlowNode {
	s.Content = &v
	return s
}

func (s *GraphFlowNode) SetId(v int64) *GraphFlowNode {
	s.Id = &v
	return s
}

func (s *GraphFlowNode) SetIndex(v int32) *GraphFlowNode {
	s.Index = &v
	return s
}

func (s *GraphFlowNode) SetName(v string) *GraphFlowNode {
	s.Name = &v
	return s
}

func (s *GraphFlowNode) SetNextNodes(v []*GraphFlowNodeNextNodes) *GraphFlowNode {
	s.NextNodes = v
	return s
}

func (s *GraphFlowNode) SetNodeType(v string) *GraphFlowNode {
	s.NodeType = &v
	return s
}

func (s *GraphFlowNode) SetProperties(v *GraphFlowNodeProperties) *GraphFlowNode {
	s.Properties = v
	return s
}

func (s *GraphFlowNode) SetRid(v int64) *GraphFlowNode {
	s.Rid = &v
	return s
}

func (s *GraphFlowNode) SetUseConditions(v bool) *GraphFlowNode {
	s.UseConditions = &v
	return s
}

func (s *GraphFlowNode) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NextNodes != nil {
		for _, item := range s.NextNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Properties != nil {
		if err := s.Properties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GraphFlowNodeNextNodes struct {
	CheckType  *int32    `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	Index      *int32    `json:"Index,omitempty" xml:"Index,omitempty"`
	Lambda     *string   `json:"Lambda,omitempty" xml:"Lambda,omitempty"`
	Name       *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	NextNodeId *int64    `json:"NextNodeId,omitempty" xml:"NextNodeId,omitempty"`
	Triggers   []*string `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
}

func (s GraphFlowNodeNextNodes) String() string {
	return dara.Prettify(s)
}

func (s GraphFlowNodeNextNodes) GoString() string {
	return s.String()
}

func (s *GraphFlowNodeNextNodes) GetCheckType() *int32 {
	return s.CheckType
}

func (s *GraphFlowNodeNextNodes) GetIndex() *int32 {
	return s.Index
}

func (s *GraphFlowNodeNextNodes) GetLambda() *string {
	return s.Lambda
}

func (s *GraphFlowNodeNextNodes) GetName() *string {
	return s.Name
}

func (s *GraphFlowNodeNextNodes) GetNextNodeId() *int64 {
	return s.NextNodeId
}

func (s *GraphFlowNodeNextNodes) GetTriggers() []*string {
	return s.Triggers
}

func (s *GraphFlowNodeNextNodes) SetCheckType(v int32) *GraphFlowNodeNextNodes {
	s.CheckType = &v
	return s
}

func (s *GraphFlowNodeNextNodes) SetIndex(v int32) *GraphFlowNodeNextNodes {
	s.Index = &v
	return s
}

func (s *GraphFlowNodeNextNodes) SetLambda(v string) *GraphFlowNodeNextNodes {
	s.Lambda = &v
	return s
}

func (s *GraphFlowNodeNextNodes) SetName(v string) *GraphFlowNodeNextNodes {
	s.Name = &v
	return s
}

func (s *GraphFlowNodeNextNodes) SetNextNodeId(v int64) *GraphFlowNodeNextNodes {
	s.NextNodeId = &v
	return s
}

func (s *GraphFlowNodeNextNodes) SetTriggers(v []*string) *GraphFlowNodeNextNodes {
	s.Triggers = v
	return s
}

func (s *GraphFlowNodeNextNodes) Validate() error {
	return dara.Validate(s)
}

type GraphFlowNodeProperties struct {
	AutoReview       *int32    `json:"AutoReview,omitempty" xml:"AutoReview,omitempty"`
	BranchJudge      *bool     `json:"BranchJudge,omitempty" xml:"BranchJudge,omitempty"`
	CheckMoreSize    *int32    `json:"CheckMoreSize,omitempty" xml:"CheckMoreSize,omitempty"`
	CheckType        *int32    `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	Lambda           *string   `json:"Lambda,omitempty" xml:"Lambda,omitempty"`
	Role             *string   `json:"Role,omitempty" xml:"Role,omitempty"`
	RuleScoreType    *int32    `json:"RuleScoreType,omitempty" xml:"RuleScoreType,omitempty"`
	SayType          *string   `json:"SayType,omitempty" xml:"SayType,omitempty"`
	ScoreNum         *int32    `json:"ScoreNum,omitempty" xml:"ScoreNum,omitempty"`
	ScoreNumType     *int32    `json:"ScoreNumType,omitempty" xml:"ScoreNumType,omitempty"`
	ScoreRuleHitType *int32    `json:"ScoreRuleHitType,omitempty" xml:"ScoreRuleHitType,omitempty"`
	ScoreType        *int32    `json:"ScoreType,omitempty" xml:"ScoreType,omitempty"`
	Triggers         []*string `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
	Type             *string   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GraphFlowNodeProperties) String() string {
	return dara.Prettify(s)
}

func (s GraphFlowNodeProperties) GoString() string {
	return s.String()
}

func (s *GraphFlowNodeProperties) GetAutoReview() *int32 {
	return s.AutoReview
}

func (s *GraphFlowNodeProperties) GetBranchJudge() *bool {
	return s.BranchJudge
}

func (s *GraphFlowNodeProperties) GetCheckMoreSize() *int32 {
	return s.CheckMoreSize
}

func (s *GraphFlowNodeProperties) GetCheckType() *int32 {
	return s.CheckType
}

func (s *GraphFlowNodeProperties) GetLambda() *string {
	return s.Lambda
}

func (s *GraphFlowNodeProperties) GetRole() *string {
	return s.Role
}

func (s *GraphFlowNodeProperties) GetRuleScoreType() *int32 {
	return s.RuleScoreType
}

func (s *GraphFlowNodeProperties) GetSayType() *string {
	return s.SayType
}

func (s *GraphFlowNodeProperties) GetScoreNum() *int32 {
	return s.ScoreNum
}

func (s *GraphFlowNodeProperties) GetScoreNumType() *int32 {
	return s.ScoreNumType
}

func (s *GraphFlowNodeProperties) GetScoreRuleHitType() *int32 {
	return s.ScoreRuleHitType
}

func (s *GraphFlowNodeProperties) GetScoreType() *int32 {
	return s.ScoreType
}

func (s *GraphFlowNodeProperties) GetTriggers() []*string {
	return s.Triggers
}

func (s *GraphFlowNodeProperties) GetType() *string {
	return s.Type
}

func (s *GraphFlowNodeProperties) SetAutoReview(v int32) *GraphFlowNodeProperties {
	s.AutoReview = &v
	return s
}

func (s *GraphFlowNodeProperties) SetBranchJudge(v bool) *GraphFlowNodeProperties {
	s.BranchJudge = &v
	return s
}

func (s *GraphFlowNodeProperties) SetCheckMoreSize(v int32) *GraphFlowNodeProperties {
	s.CheckMoreSize = &v
	return s
}

func (s *GraphFlowNodeProperties) SetCheckType(v int32) *GraphFlowNodeProperties {
	s.CheckType = &v
	return s
}

func (s *GraphFlowNodeProperties) SetLambda(v string) *GraphFlowNodeProperties {
	s.Lambda = &v
	return s
}

func (s *GraphFlowNodeProperties) SetRole(v string) *GraphFlowNodeProperties {
	s.Role = &v
	return s
}

func (s *GraphFlowNodeProperties) SetRuleScoreType(v int32) *GraphFlowNodeProperties {
	s.RuleScoreType = &v
	return s
}

func (s *GraphFlowNodeProperties) SetSayType(v string) *GraphFlowNodeProperties {
	s.SayType = &v
	return s
}

func (s *GraphFlowNodeProperties) SetScoreNum(v int32) *GraphFlowNodeProperties {
	s.ScoreNum = &v
	return s
}

func (s *GraphFlowNodeProperties) SetScoreNumType(v int32) *GraphFlowNodeProperties {
	s.ScoreNumType = &v
	return s
}

func (s *GraphFlowNodeProperties) SetScoreRuleHitType(v int32) *GraphFlowNodeProperties {
	s.ScoreRuleHitType = &v
	return s
}

func (s *GraphFlowNodeProperties) SetScoreType(v int32) *GraphFlowNodeProperties {
	s.ScoreType = &v
	return s
}

func (s *GraphFlowNodeProperties) SetTriggers(v []*string) *GraphFlowNodeProperties {
	s.Triggers = v
	return s
}

func (s *GraphFlowNodeProperties) SetType(v string) *GraphFlowNodeProperties {
	s.Type = &v
	return s
}

func (s *GraphFlowNodeProperties) Validate() error {
	return dara.Validate(s)
}
