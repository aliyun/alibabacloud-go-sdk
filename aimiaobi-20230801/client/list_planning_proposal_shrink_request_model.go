// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPlanningProposalShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListPlanningProposalShrinkRequest
	GetAgentKey() *string
	SetCustomViewPointId(v string) *ListPlanningProposalShrinkRequest
	GetCustomViewPointId() *string
	SetCustomViewPointIdsShrink(v string) *ListPlanningProposalShrinkRequest
	GetCustomViewPointIdsShrink() *string
	SetMaxResults(v int32) *ListPlanningProposalShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListPlanningProposalShrinkRequest
	GetNextToken() *string
	SetTitlesShrink(v string) *ListPlanningProposalShrinkRequest
	GetTitlesShrink() *string
	SetTopic(v string) *ListPlanningProposalShrinkRequest
	GetTopic() *string
	SetTopicSource(v string) *ListPlanningProposalShrinkRequest
	GetTopicSource() *string
	SetTopicVersion(v string) *ListPlanningProposalShrinkRequest
	GetTopicVersion() *string
	SetViewPointType(v string) *ListPlanningProposalShrinkRequest
	GetViewPointType() *string
}

type ListPlanningProposalShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// e7b26a9e1211444db8f0a984361a5e0f
	CustomViewPointId *string `json:"CustomViewPointId,omitempty" xml:"CustomViewPointId,omitempty"`
	// example:
	//
	// 27971fc8f3ce4ed58c7e7fc4b503e432
	CustomViewPointIdsShrink *string `json:"CustomViewPointIds,omitempty" xml:"CustomViewPointIds,omitempty"`
	// example:
	//
	// 73
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 下一页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 标题
	//
	//      *
	TitlesShrink *string `json:"Titles,omitempty" xml:"Titles,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 热榜主题
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 热榜源
	TopicSource *string `json:"TopicSource,omitempty" xml:"TopicSource,omitempty"`
	// example:
	//
	// 2024-09-10_08
	TopicVersion *string `json:"TopicVersion,omitempty" xml:"TopicVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CustomViewPoints
	ViewPointType *string `json:"ViewPointType,omitempty" xml:"ViewPointType,omitempty"`
}

func (s ListPlanningProposalShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPlanningProposalShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPlanningProposalShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListPlanningProposalShrinkRequest) GetCustomViewPointId() *string {
	return s.CustomViewPointId
}

func (s *ListPlanningProposalShrinkRequest) GetCustomViewPointIdsShrink() *string {
	return s.CustomViewPointIdsShrink
}

func (s *ListPlanningProposalShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPlanningProposalShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPlanningProposalShrinkRequest) GetTitlesShrink() *string {
	return s.TitlesShrink
}

func (s *ListPlanningProposalShrinkRequest) GetTopic() *string {
	return s.Topic
}

func (s *ListPlanningProposalShrinkRequest) GetTopicSource() *string {
	return s.TopicSource
}

func (s *ListPlanningProposalShrinkRequest) GetTopicVersion() *string {
	return s.TopicVersion
}

func (s *ListPlanningProposalShrinkRequest) GetViewPointType() *string {
	return s.ViewPointType
}

func (s *ListPlanningProposalShrinkRequest) SetAgentKey(v string) *ListPlanningProposalShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *ListPlanningProposalShrinkRequest) SetCustomViewPointId(v string) *ListPlanningProposalShrinkRequest {
	s.CustomViewPointId = &v
	return s
}

func (s *ListPlanningProposalShrinkRequest) SetCustomViewPointIdsShrink(v string) *ListPlanningProposalShrinkRequest {
	s.CustomViewPointIdsShrink = &v
	return s
}

func (s *ListPlanningProposalShrinkRequest) SetMaxResults(v int32) *ListPlanningProposalShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPlanningProposalShrinkRequest) SetNextToken(v string) *ListPlanningProposalShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListPlanningProposalShrinkRequest) SetTitlesShrink(v string) *ListPlanningProposalShrinkRequest {
	s.TitlesShrink = &v
	return s
}

func (s *ListPlanningProposalShrinkRequest) SetTopic(v string) *ListPlanningProposalShrinkRequest {
	s.Topic = &v
	return s
}

func (s *ListPlanningProposalShrinkRequest) SetTopicSource(v string) *ListPlanningProposalShrinkRequest {
	s.TopicSource = &v
	return s
}

func (s *ListPlanningProposalShrinkRequest) SetTopicVersion(v string) *ListPlanningProposalShrinkRequest {
	s.TopicVersion = &v
	return s
}

func (s *ListPlanningProposalShrinkRequest) SetViewPointType(v string) *ListPlanningProposalShrinkRequest {
	s.ViewPointType = &v
	return s
}

func (s *ListPlanningProposalShrinkRequest) Validate() error {
	return dara.Validate(s)
}
