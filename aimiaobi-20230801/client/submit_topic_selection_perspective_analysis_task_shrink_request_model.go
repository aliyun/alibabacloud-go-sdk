// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest
	GetAgentKey() *string
	SetDocumentsShrink(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest
	GetDocumentsShrink() *string
	SetPerspectiveTypesShrink(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest
	GetPerspectiveTypesShrink() *string
	SetTopic(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest
	GetTopic() *string
}

type SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest struct {
	// The unique identifier of the workspace: [AgentKey](https://help.aliyun.com/document_detail/2587494.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The list of documents to be analyzed. (Provide at least one of documents or topic)
	DocumentsShrink *string `json:"Documents,omitempty" xml:"Documents,omitempty"`
	// The topic selection perspective tasks to be analyzed. By default, this parameter is empty, which means all tasks are analyzed. (TopicSummary: Topic event summary, HotViewPoints: Hot topic selection perspectives, TimedViewPoints: Timeliness topic selection perspectives, WebReviewPoints: Online review topic selection perspectives, FreshViewPoints: Novel topic selection perspectives)
	//
	// example:
	//
	// TimedViewPoints
	PerspectiveTypesShrink *string `json:"PerspectiveTypes,omitempty" xml:"PerspectiveTypes,omitempty"`
	// The name of the topic to be analyzed. (Provide at least one of documents or topic)
	//
	// example:
	//
	// 待分析的主题名（documents与topic二者至少传一个）
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest) GetDocumentsShrink() *string {
	return s.DocumentsShrink
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest) GetPerspectiveTypesShrink() *string {
	return s.PerspectiveTypesShrink
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest) GetTopic() *string {
	return s.Topic
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest) SetAgentKey(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest) SetDocumentsShrink(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest {
	s.DocumentsShrink = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest) SetPerspectiveTypesShrink(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest {
	s.PerspectiveTypesShrink = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest) SetTopic(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest {
	s.Topic = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
