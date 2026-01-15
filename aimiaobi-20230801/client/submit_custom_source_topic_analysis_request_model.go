// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCustomSourceTopicAnalysisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnalysisTypes(v []*string) *SubmitCustomSourceTopicAnalysisRequest
	GetAnalysisTypes() []*string
	SetFileType(v string) *SubmitCustomSourceTopicAnalysisRequest
	GetFileType() *string
	SetFileUrl(v string) *SubmitCustomSourceTopicAnalysisRequest
	GetFileUrl() *string
	SetMaxTopicSize(v int32) *SubmitCustomSourceTopicAnalysisRequest
	GetMaxTopicSize() *int32
	SetNews(v []*SubmitCustomSourceTopicAnalysisRequestNews) *SubmitCustomSourceTopicAnalysisRequest
	GetNews() []*SubmitCustomSourceTopicAnalysisRequestNews
	SetTopics(v []*SubmitCustomSourceTopicAnalysisRequestTopics) *SubmitCustomSourceTopicAnalysisRequest
	GetTopics() []*SubmitCustomSourceTopicAnalysisRequestTopics
	SetTopicsFileUrl(v string) *SubmitCustomSourceTopicAnalysisRequest
	GetTopicsFileUrl() *string
	SetWorkspaceId(v string) *SubmitCustomSourceTopicAnalysisRequest
	GetWorkspaceId() *string
}

type SubmitCustomSourceTopicAnalysisRequest struct {
	AnalysisTypes []*string `json:"AnalysisTypes,omitempty" xml:"AnalysisTypes,omitempty" type:"Repeated"`
	// example:
	//
	// json
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// example:
	//
	// http://www.example.com/xxx.json
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// 50
	MaxTopicSize *int32                                          `json:"MaxTopicSize,omitempty" xml:"MaxTopicSize,omitempty"`
	News         []*SubmitCustomSourceTopicAnalysisRequestNews   `json:"News,omitempty" xml:"News,omitempty" type:"Repeated"`
	Topics       []*SubmitCustomSourceTopicAnalysisRequestTopics `json:"Topics,omitempty" xml:"Topics,omitempty" type:"Repeated"`
	// example:
	//
	// http://www.example.com/xxx.jsonline
	TopicsFileUrl *string `json:"TopicsFileUrl,omitempty" xml:"TopicsFileUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SubmitCustomSourceTopicAnalysisRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitCustomSourceTopicAnalysisRequest) GoString() string {
	return s.String()
}

func (s *SubmitCustomSourceTopicAnalysisRequest) GetAnalysisTypes() []*string {
	return s.AnalysisTypes
}

func (s *SubmitCustomSourceTopicAnalysisRequest) GetFileType() *string {
	return s.FileType
}

func (s *SubmitCustomSourceTopicAnalysisRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *SubmitCustomSourceTopicAnalysisRequest) GetMaxTopicSize() *int32 {
	return s.MaxTopicSize
}

func (s *SubmitCustomSourceTopicAnalysisRequest) GetNews() []*SubmitCustomSourceTopicAnalysisRequestNews {
	return s.News
}

func (s *SubmitCustomSourceTopicAnalysisRequest) GetTopics() []*SubmitCustomSourceTopicAnalysisRequestTopics {
	return s.Topics
}

func (s *SubmitCustomSourceTopicAnalysisRequest) GetTopicsFileUrl() *string {
	return s.TopicsFileUrl
}

func (s *SubmitCustomSourceTopicAnalysisRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitCustomSourceTopicAnalysisRequest) SetAnalysisTypes(v []*string) *SubmitCustomSourceTopicAnalysisRequest {
	s.AnalysisTypes = v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisRequest) SetFileType(v string) *SubmitCustomSourceTopicAnalysisRequest {
	s.FileType = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisRequest) SetFileUrl(v string) *SubmitCustomSourceTopicAnalysisRequest {
	s.FileUrl = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisRequest) SetMaxTopicSize(v int32) *SubmitCustomSourceTopicAnalysisRequest {
	s.MaxTopicSize = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisRequest) SetNews(v []*SubmitCustomSourceTopicAnalysisRequestNews) *SubmitCustomSourceTopicAnalysisRequest {
	s.News = v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisRequest) SetTopics(v []*SubmitCustomSourceTopicAnalysisRequestTopics) *SubmitCustomSourceTopicAnalysisRequest {
	s.Topics = v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisRequest) SetTopicsFileUrl(v string) *SubmitCustomSourceTopicAnalysisRequest {
	s.TopicsFileUrl = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisRequest) SetWorkspaceId(v string) *SubmitCustomSourceTopicAnalysisRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisRequest) Validate() error {
	if s.News != nil {
		for _, item := range s.News {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Topics != nil {
		for _, item := range s.Topics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitCustomSourceTopicAnalysisRequestNews struct {
	Comments []*SubmitCustomSourceTopicAnalysisRequestNewsComments `json:"Comments,omitempty" xml:"Comments,omitempty" type:"Repeated"`
	Content  *string                                               `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2024-01-22 10:29:00
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	Source  *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Title   *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// http://www.example.com/xxx.html
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SubmitCustomSourceTopicAnalysisRequestNews) String() string {
	return dara.Prettify(s)
}

func (s SubmitCustomSourceTopicAnalysisRequestNews) GoString() string {
	return s.String()
}

func (s *SubmitCustomSourceTopicAnalysisRequestNews) GetComments() []*SubmitCustomSourceTopicAnalysisRequestNewsComments {
	return s.Comments
}

func (s *SubmitCustomSourceTopicAnalysisRequestNews) GetContent() *string {
	return s.Content
}

func (s *SubmitCustomSourceTopicAnalysisRequestNews) GetPubTime() *string {
	return s.PubTime
}

func (s *SubmitCustomSourceTopicAnalysisRequestNews) GetSource() *string {
	return s.Source
}

func (s *SubmitCustomSourceTopicAnalysisRequestNews) GetTitle() *string {
	return s.Title
}

func (s *SubmitCustomSourceTopicAnalysisRequestNews) GetUrl() *string {
	return s.Url
}

func (s *SubmitCustomSourceTopicAnalysisRequestNews) SetComments(v []*SubmitCustomSourceTopicAnalysisRequestNewsComments) *SubmitCustomSourceTopicAnalysisRequestNews {
	s.Comments = v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisRequestNews) SetContent(v string) *SubmitCustomSourceTopicAnalysisRequestNews {
	s.Content = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisRequestNews) SetPubTime(v string) *SubmitCustomSourceTopicAnalysisRequestNews {
	s.PubTime = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisRequestNews) SetSource(v string) *SubmitCustomSourceTopicAnalysisRequestNews {
	s.Source = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisRequestNews) SetTitle(v string) *SubmitCustomSourceTopicAnalysisRequestNews {
	s.Title = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisRequestNews) SetUrl(v string) *SubmitCustomSourceTopicAnalysisRequestNews {
	s.Url = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisRequestNews) Validate() error {
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

type SubmitCustomSourceTopicAnalysisRequestNewsComments struct {
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s SubmitCustomSourceTopicAnalysisRequestNewsComments) String() string {
	return dara.Prettify(s)
}

func (s SubmitCustomSourceTopicAnalysisRequestNewsComments) GoString() string {
	return s.String()
}

func (s *SubmitCustomSourceTopicAnalysisRequestNewsComments) GetText() *string {
	return s.Text
}

func (s *SubmitCustomSourceTopicAnalysisRequestNewsComments) SetText(v string) *SubmitCustomSourceTopicAnalysisRequestNewsComments {
	s.Text = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisRequestNewsComments) Validate() error {
	return dara.Validate(s)
}

type SubmitCustomSourceTopicAnalysisRequestTopics struct {
	News []*HottopicNews `json:"News,omitempty" xml:"News,omitempty" type:"Repeated"`
	// example:
	//
	// 话题名称
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s SubmitCustomSourceTopicAnalysisRequestTopics) String() string {
	return dara.Prettify(s)
}

func (s SubmitCustomSourceTopicAnalysisRequestTopics) GoString() string {
	return s.String()
}

func (s *SubmitCustomSourceTopicAnalysisRequestTopics) GetNews() []*HottopicNews {
	return s.News
}

func (s *SubmitCustomSourceTopicAnalysisRequestTopics) GetTopic() *string {
	return s.Topic
}

func (s *SubmitCustomSourceTopicAnalysisRequestTopics) SetNews(v []*HottopicNews) *SubmitCustomSourceTopicAnalysisRequestTopics {
	s.News = v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisRequestTopics) SetTopic(v string) *SubmitCustomSourceTopicAnalysisRequestTopics {
	s.Topic = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisRequestTopics) Validate() error {
	if s.News != nil {
		for _, item := range s.News {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
