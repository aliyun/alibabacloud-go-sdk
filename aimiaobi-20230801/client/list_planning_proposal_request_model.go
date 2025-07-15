// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPlanningProposalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListPlanningProposalRequest
	GetAgentKey() *string
	SetCustomViewPointId(v string) *ListPlanningProposalRequest
	GetCustomViewPointId() *string
	SetCustomViewPointIds(v []*string) *ListPlanningProposalRequest
	GetCustomViewPointIds() []*string
	SetMaxResults(v int32) *ListPlanningProposalRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListPlanningProposalRequest
	GetNextToken() *string
	SetTitles(v []*string) *ListPlanningProposalRequest
	GetTitles() []*string
	SetTopic(v string) *ListPlanningProposalRequest
	GetTopic() *string
	SetTopicSource(v string) *ListPlanningProposalRequest
	GetTopicSource() *string
	SetTopicVersion(v string) *ListPlanningProposalRequest
	GetTopicVersion() *string
	SetViewPointType(v string) *ListPlanningProposalRequest
	GetViewPointType() *string
}

type ListPlanningProposalRequest struct {
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
	CustomViewPointIds []*string `json:"CustomViewPointIds,omitempty" xml:"CustomViewPointIds,omitempty" type:"Repeated"`
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
	Titles []*string `json:"Titles,omitempty" xml:"Titles,omitempty" type:"Repeated"`
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

func (s ListPlanningProposalRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPlanningProposalRequest) GoString() string {
	return s.String()
}

func (s *ListPlanningProposalRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListPlanningProposalRequest) GetCustomViewPointId() *string {
	return s.CustomViewPointId
}

func (s *ListPlanningProposalRequest) GetCustomViewPointIds() []*string {
	return s.CustomViewPointIds
}

func (s *ListPlanningProposalRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPlanningProposalRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPlanningProposalRequest) GetTitles() []*string {
	return s.Titles
}

func (s *ListPlanningProposalRequest) GetTopic() *string {
	return s.Topic
}

func (s *ListPlanningProposalRequest) GetTopicSource() *string {
	return s.TopicSource
}

func (s *ListPlanningProposalRequest) GetTopicVersion() *string {
	return s.TopicVersion
}

func (s *ListPlanningProposalRequest) GetViewPointType() *string {
	return s.ViewPointType
}

func (s *ListPlanningProposalRequest) SetAgentKey(v string) *ListPlanningProposalRequest {
	s.AgentKey = &v
	return s
}

func (s *ListPlanningProposalRequest) SetCustomViewPointId(v string) *ListPlanningProposalRequest {
	s.CustomViewPointId = &v
	return s
}

func (s *ListPlanningProposalRequest) SetCustomViewPointIds(v []*string) *ListPlanningProposalRequest {
	s.CustomViewPointIds = v
	return s
}

func (s *ListPlanningProposalRequest) SetMaxResults(v int32) *ListPlanningProposalRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPlanningProposalRequest) SetNextToken(v string) *ListPlanningProposalRequest {
	s.NextToken = &v
	return s
}

func (s *ListPlanningProposalRequest) SetTitles(v []*string) *ListPlanningProposalRequest {
	s.Titles = v
	return s
}

func (s *ListPlanningProposalRequest) SetTopic(v string) *ListPlanningProposalRequest {
	s.Topic = &v
	return s
}

func (s *ListPlanningProposalRequest) SetTopicSource(v string) *ListPlanningProposalRequest {
	s.TopicSource = &v
	return s
}

func (s *ListPlanningProposalRequest) SetTopicVersion(v string) *ListPlanningProposalRequest {
	s.TopicVersion = &v
	return s
}

func (s *ListPlanningProposalRequest) SetViewPointType(v string) *ListPlanningProposalRequest {
	s.ViewPointType = &v
	return s
}

func (s *ListPlanningProposalRequest) Validate() error {
	return dara.Validate(s)
}
