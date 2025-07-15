// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCustomHotTopicViewPointAnalysisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAskUser(v string) *RunCustomHotTopicViewPointAnalysisRequest
	GetAskUser() *string
	SetPrompt(v string) *RunCustomHotTopicViewPointAnalysisRequest
	GetPrompt() *string
	SetSearchQuery(v string) *RunCustomHotTopicViewPointAnalysisRequest
	GetSearchQuery() *string
	SetSkipAskUser(v bool) *RunCustomHotTopicViewPointAnalysisRequest
	GetSkipAskUser() *bool
	SetTopic(v string) *RunCustomHotTopicViewPointAnalysisRequest
	GetTopic() *string
	SetTopicId(v string) *RunCustomHotTopicViewPointAnalysisRequest
	GetTopicId() *string
	SetTopicSource(v string) *RunCustomHotTopicViewPointAnalysisRequest
	GetTopicSource() *string
	SetTopicVersion(v string) *RunCustomHotTopicViewPointAnalysisRequest
	GetTopicVersion() *string
	SetUserBack(v string) *RunCustomHotTopicViewPointAnalysisRequest
	GetUserBack() *string
	SetWorkspaceId(v string) *RunCustomHotTopicViewPointAnalysisRequest
	GetWorkspaceId() *string
}

type RunCustomHotTopicViewPointAnalysisRequest struct {
	// example:
	//
	// 模型反问
	AskUser *string `json:"AskUser,omitempty" xml:"AskUser,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 自定义选题视角的Prompt
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// 改写后的Query
	SearchQuery *string `json:"SearchQuery,omitempty" xml:"SearchQuery,omitempty"`
	// example:
	//
	// true
	SkipAskUser *bool `json:"SkipAskUser,omitempty" xml:"SkipAskUser,omitempty"`
	// example:
	//
	// 热点主题
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// example:
	//
	// 热点主题ID
	TopicId *string `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
	// example:
	//
	// 热点主题来源
	TopicSource *string `json:"TopicSource,omitempty" xml:"TopicSource,omitempty"`
	// example:
	//
	// 热点主题版本
	TopicVersion *string `json:"TopicVersion,omitempty" xml:"TopicVersion,omitempty"`
	// example:
	//
	// 用户反馈
	UserBack *string `json:"UserBack,omitempty" xml:"UserBack,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunCustomHotTopicViewPointAnalysisRequest) String() string {
	return dara.Prettify(s)
}

func (s RunCustomHotTopicViewPointAnalysisRequest) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) GetAskUser() *string {
	return s.AskUser
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) GetSearchQuery() *string {
	return s.SearchQuery
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) GetSkipAskUser() *bool {
	return s.SkipAskUser
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) GetTopic() *string {
	return s.Topic
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) GetTopicId() *string {
	return s.TopicId
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) GetTopicSource() *string {
	return s.TopicSource
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) GetTopicVersion() *string {
	return s.TopicVersion
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) GetUserBack() *string {
	return s.UserBack
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) SetAskUser(v string) *RunCustomHotTopicViewPointAnalysisRequest {
	s.AskUser = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) SetPrompt(v string) *RunCustomHotTopicViewPointAnalysisRequest {
	s.Prompt = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) SetSearchQuery(v string) *RunCustomHotTopicViewPointAnalysisRequest {
	s.SearchQuery = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) SetSkipAskUser(v bool) *RunCustomHotTopicViewPointAnalysisRequest {
	s.SkipAskUser = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) SetTopic(v string) *RunCustomHotTopicViewPointAnalysisRequest {
	s.Topic = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) SetTopicId(v string) *RunCustomHotTopicViewPointAnalysisRequest {
	s.TopicId = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) SetTopicSource(v string) *RunCustomHotTopicViewPointAnalysisRequest {
	s.TopicSource = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) SetTopicVersion(v string) *RunCustomHotTopicViewPointAnalysisRequest {
	s.TopicVersion = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) SetUserBack(v string) *RunCustomHotTopicViewPointAnalysisRequest {
	s.UserBack = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) SetWorkspaceId(v string) *RunCustomHotTopicViewPointAnalysisRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) Validate() error {
	return dara.Validate(s)
}
