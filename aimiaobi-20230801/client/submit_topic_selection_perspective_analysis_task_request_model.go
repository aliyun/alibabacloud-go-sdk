// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTopicSelectionPerspectiveAnalysisTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskRequest
	GetAgentKey() *string
	SetDocuments(v []*SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) *SubmitTopicSelectionPerspectiveAnalysisTaskRequest
	GetDocuments() []*SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments
	SetPerspectiveTypes(v []*string) *SubmitTopicSelectionPerspectiveAnalysisTaskRequest
	GetPerspectiveTypes() []*string
	SetTopic(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskRequest
	GetTopic() *string
}

type SubmitTopicSelectionPerspectiveAnalysisTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey  *string                                                        `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Documents []*SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments `json:"Documents,omitempty" xml:"Documents,omitempty" type:"Repeated"`
	// example:
	//
	// TimedViewPoints
	PerspectiveTypes []*string `json:"PerspectiveTypes,omitempty" xml:"PerspectiveTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 待分析的主题名（documents与topic二者至少传一个）
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s SubmitTopicSelectionPerspectiveAnalysisTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitTopicSelectionPerspectiveAnalysisTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequest) GetDocuments() []*SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	return s.Documents
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequest) GetPerspectiveTypes() []*string {
	return s.PerspectiveTypes
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequest) GetTopic() *string {
	return s.Topic
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequest) SetAgentKey(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequest) SetDocuments(v []*SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) *SubmitTopicSelectionPerspectiveAnalysisTaskRequest {
	s.Documents = v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequest) SetPerspectiveTypes(v []*string) *SubmitTopicSelectionPerspectiveAnalysisTaskRequest {
	s.PerspectiveTypes = v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequest) SetTopic(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskRequest {
	s.Topic = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequest) Validate() error {
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

type SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments struct {
	// example:
	//
	// 作者
	Author   *string                                                                `json:"Author,omitempty" xml:"Author,omitempty"`
	Comments []*SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocumentsComments `json:"Comments,omitempty" xml:"Comments,omitempty" type:"Repeated"`
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

func (s SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) String() string {
	return dara.Prettify(s)
}

func (s SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) GoString() string {
	return s.String()
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) GetAuthor() *string {
	return s.Author
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) GetComments() []*SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocumentsComments {
	return s.Comments
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) GetContent() *string {
	return s.Content
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) GetPubTime() *string {
	return s.PubTime
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) GetSource() *string {
	return s.Source
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) GetSummary() *string {
	return s.Summary
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) GetTitle() *string {
	return s.Title
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) GetUrl() *string {
	return s.Url
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) SetAuthor(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	s.Author = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) SetComments(v []*SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocumentsComments) *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	s.Comments = v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) SetContent(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	s.Content = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) SetPubTime(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	s.PubTime = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) SetSource(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	s.Source = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) SetSummary(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	s.Summary = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) SetTitle(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	s.Title = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) SetUrl(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	s.Url = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) Validate() error {
	if s.Comments != nil {
		for _, item := range s.Comments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocumentsComments struct {
	Text     *string `json:"Text,omitempty" xml:"Text,omitempty"`
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocumentsComments) String() string {
	return dara.Prettify(s)
}

func (s SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocumentsComments) GoString() string {
	return s.String()
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocumentsComments) GetText() *string {
	return s.Text
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocumentsComments) GetUsername() *string {
	return s.Username
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocumentsComments) SetText(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocumentsComments {
	s.Text = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocumentsComments) SetUsername(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocumentsComments {
	s.Username = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocumentsComments) Validate() error {
	return dara.Validate(s)
}
