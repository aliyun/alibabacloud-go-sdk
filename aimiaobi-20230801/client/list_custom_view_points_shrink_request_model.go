// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomViewPointsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListCustomViewPointsShrinkRequest
	GetAgentKey() *string
	SetAttitude(v string) *ListCustomViewPointsShrinkRequest
	GetAttitude() *string
	SetAttitudesShrink(v string) *ListCustomViewPointsShrinkRequest
	GetAttitudesShrink() *string
	SetCustomViewPointId(v string) *ListCustomViewPointsShrinkRequest
	GetCustomViewPointId() *string
	SetCustomViewPointIdsShrink(v string) *ListCustomViewPointsShrinkRequest
	GetCustomViewPointIdsShrink() *string
	SetMaxResults(v int32) *ListCustomViewPointsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListCustomViewPointsShrinkRequest
	GetNextToken() *string
	SetTopic(v string) *ListCustomViewPointsShrinkRequest
	GetTopic() *string
	SetTopicId(v string) *ListCustomViewPointsShrinkRequest
	GetTopicId() *string
}

type ListCustomViewPointsShrinkRequest struct {
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
	AttitudesShrink *string `json:"Attitudes,omitempty" xml:"Attitudes,omitempty"`
	// example:
	//
	// 461591f4880747f890702c1b90494d1a
	CustomViewPointId *string `json:"CustomViewPointId,omitempty" xml:"CustomViewPointId,omitempty"`
	// example:
	//
	// 7ece3d1212e04c9ca716ae2486228f3f
	CustomViewPointIdsShrink *string `json:"CustomViewPointIds,omitempty" xml:"CustomViewPointIds,omitempty"`
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

func (s ListCustomViewPointsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomViewPointsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListCustomViewPointsShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListCustomViewPointsShrinkRequest) GetAttitude() *string {
	return s.Attitude
}

func (s *ListCustomViewPointsShrinkRequest) GetAttitudesShrink() *string {
	return s.AttitudesShrink
}

func (s *ListCustomViewPointsShrinkRequest) GetCustomViewPointId() *string {
	return s.CustomViewPointId
}

func (s *ListCustomViewPointsShrinkRequest) GetCustomViewPointIdsShrink() *string {
	return s.CustomViewPointIdsShrink
}

func (s *ListCustomViewPointsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCustomViewPointsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCustomViewPointsShrinkRequest) GetTopic() *string {
	return s.Topic
}

func (s *ListCustomViewPointsShrinkRequest) GetTopicId() *string {
	return s.TopicId
}

func (s *ListCustomViewPointsShrinkRequest) SetAgentKey(v string) *ListCustomViewPointsShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *ListCustomViewPointsShrinkRequest) SetAttitude(v string) *ListCustomViewPointsShrinkRequest {
	s.Attitude = &v
	return s
}

func (s *ListCustomViewPointsShrinkRequest) SetAttitudesShrink(v string) *ListCustomViewPointsShrinkRequest {
	s.AttitudesShrink = &v
	return s
}

func (s *ListCustomViewPointsShrinkRequest) SetCustomViewPointId(v string) *ListCustomViewPointsShrinkRequest {
	s.CustomViewPointId = &v
	return s
}

func (s *ListCustomViewPointsShrinkRequest) SetCustomViewPointIdsShrink(v string) *ListCustomViewPointsShrinkRequest {
	s.CustomViewPointIdsShrink = &v
	return s
}

func (s *ListCustomViewPointsShrinkRequest) SetMaxResults(v int32) *ListCustomViewPointsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCustomViewPointsShrinkRequest) SetNextToken(v string) *ListCustomViewPointsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListCustomViewPointsShrinkRequest) SetTopic(v string) *ListCustomViewPointsShrinkRequest {
	s.Topic = &v
	return s
}

func (s *ListCustomViewPointsShrinkRequest) SetTopicId(v string) *ListCustomViewPointsShrinkRequest {
	s.TopicId = &v
	return s
}

func (s *ListCustomViewPointsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
