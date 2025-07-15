// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomViewPointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListCustomViewPointsRequest
	GetAgentKey() *string
	SetAttitude(v string) *ListCustomViewPointsRequest
	GetAttitude() *string
	SetAttitudes(v []*string) *ListCustomViewPointsRequest
	GetAttitudes() []*string
	SetCustomViewPointId(v string) *ListCustomViewPointsRequest
	GetCustomViewPointId() *string
	SetCustomViewPointIds(v []*string) *ListCustomViewPointsRequest
	GetCustomViewPointIds() []*string
	SetMaxResults(v int32) *ListCustomViewPointsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListCustomViewPointsRequest
	GetNextToken() *string
	SetTopic(v string) *ListCustomViewPointsRequest
	GetTopic() *string
	SetTopicId(v string) *ListCustomViewPointsRequest
	GetTopicId() *string
}

type ListCustomViewPointsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 观点
	Attitude *string `json:"Attitude,omitempty" xml:"Attitude,omitempty"`
	// example:
	//
	// 观点
	//
	//      *
	Attitudes []*string `json:"Attitudes,omitempty" xml:"Attitudes,omitempty" type:"Repeated"`
	// example:
	//
	// 461591f4880747f890702c1b90494d1a
	CustomViewPointId *string `json:"CustomViewPointId,omitempty" xml:"CustomViewPointId,omitempty"`
	// example:
	//
	// 7ece3d1212e04c9ca716ae2486228f3f
	CustomViewPointIds []*string `json:"CustomViewPointIds,omitempty" xml:"CustomViewPointIds,omitempty" type:"Repeated"`
	// example:
	//
	// 52
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 下一页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 热榜主题
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// example:
	//
	// 1d20ed14db0840efb1c7eaaf4d46352b
	TopicId *string `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
}

func (s ListCustomViewPointsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomViewPointsRequest) GoString() string {
	return s.String()
}

func (s *ListCustomViewPointsRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListCustomViewPointsRequest) GetAttitude() *string {
	return s.Attitude
}

func (s *ListCustomViewPointsRequest) GetAttitudes() []*string {
	return s.Attitudes
}

func (s *ListCustomViewPointsRequest) GetCustomViewPointId() *string {
	return s.CustomViewPointId
}

func (s *ListCustomViewPointsRequest) GetCustomViewPointIds() []*string {
	return s.CustomViewPointIds
}

func (s *ListCustomViewPointsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCustomViewPointsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCustomViewPointsRequest) GetTopic() *string {
	return s.Topic
}

func (s *ListCustomViewPointsRequest) GetTopicId() *string {
	return s.TopicId
}

func (s *ListCustomViewPointsRequest) SetAgentKey(v string) *ListCustomViewPointsRequest {
	s.AgentKey = &v
	return s
}

func (s *ListCustomViewPointsRequest) SetAttitude(v string) *ListCustomViewPointsRequest {
	s.Attitude = &v
	return s
}

func (s *ListCustomViewPointsRequest) SetAttitudes(v []*string) *ListCustomViewPointsRequest {
	s.Attitudes = v
	return s
}

func (s *ListCustomViewPointsRequest) SetCustomViewPointId(v string) *ListCustomViewPointsRequest {
	s.CustomViewPointId = &v
	return s
}

func (s *ListCustomViewPointsRequest) SetCustomViewPointIds(v []*string) *ListCustomViewPointsRequest {
	s.CustomViewPointIds = v
	return s
}

func (s *ListCustomViewPointsRequest) SetMaxResults(v int32) *ListCustomViewPointsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCustomViewPointsRequest) SetNextToken(v string) *ListCustomViewPointsRequest {
	s.NextToken = &v
	return s
}

func (s *ListCustomViewPointsRequest) SetTopic(v string) *ListCustomViewPointsRequest {
	s.Topic = &v
	return s
}

func (s *ListCustomViewPointsRequest) SetTopicId(v string) *ListCustomViewPointsRequest {
	s.TopicId = &v
	return s
}

func (s *ListCustomViewPointsRequest) Validate() error {
	return dara.Validate(s)
}
