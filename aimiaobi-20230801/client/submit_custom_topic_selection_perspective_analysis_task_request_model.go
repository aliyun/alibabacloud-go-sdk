// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest
	GetAgentKey() *string
	SetDocuments(v []*SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest
	GetDocuments() []*SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments
	SetPrompt(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest
	GetPrompt() *string
	SetTopic(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest
	GetTopic() *string
}

type SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey  *string                                                              `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Documents []*SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments `json:"Documents,omitempty" xml:"Documents,omitempty" type:"Repeated"`
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

func (s SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest) GetDocuments() []*SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	return s.Documents
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest) GetTopic() *string {
	return s.Topic
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest) SetAgentKey(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest) SetDocuments(v []*SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest {
	s.Documents = v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest) SetPrompt(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest {
	s.Prompt = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest) SetTopic(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest {
	s.Topic = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest) Validate() error {
	if s.Documents != nil {
		for _, item := range s.Documents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments struct {
	// example:
	//
	// 作者
	Author *string `json:"Author,omitempty" xml:"Author,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 文章内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2024-01-22 10:29:00
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// 新浪
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 文章摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// 文章标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// https://www.example.com/aaa.docx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments) String() string {
	return dara.Prettify(s)
}

func (s SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments) GoString() string {
	return s.String()
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments) GetAuthor() *string {
	return s.Author
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments) GetContent() *string {
	return s.Content
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments) GetPubTime() *string {
	return s.PubTime
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments) GetSource() *string {
	return s.Source
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments) GetSummary() *string {
	return s.Summary
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments) GetTitle() *string {
	return s.Title
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments) GetUrl() *string {
	return s.Url
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments) SetAuthor(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	s.Author = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments) SetContent(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	s.Content = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments) SetPubTime(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	s.PubTime = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments) SetSource(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	s.Source = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments) SetSummary(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	s.Summary = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments) SetTitle(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	s.Title = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments) SetUrl(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	s.Url = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments) Validate() error {
	return dara.Validate(s)
}
