// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest
	GetAgentKey() *string
	SetDocumentsShrink(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest
	GetDocumentsShrink() *string
	SetPrompt(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest
	GetPrompt() *string
	SetTopic(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest
	GetTopic() *string
}

type SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey        *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DocumentsShrink *string `json:"Documents,omitempty" xml:"Documents,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 自定义观点的输入Prompt
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// 待分析的主题名（documents与topic二者至少传一个）
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest) GetDocumentsShrink() *string {
	return s.DocumentsShrink
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest) GetTopic() *string {
	return s.Topic
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest) SetAgentKey(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest) SetDocumentsShrink(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest {
	s.DocumentsShrink = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest) SetPrompt(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest) SetTopic(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest {
	s.Topic = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
