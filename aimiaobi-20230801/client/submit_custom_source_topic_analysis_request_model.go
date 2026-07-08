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
	// The types of analysis for hot topic selection. Multiple values are supported. If you omit this parameter, the service analyzes all types by default. If you pass an empty array, the service performs only clustering and skips the analysis of hot topics for selection.
	//
	// `HotViewPoints`: Analyzes perspectives on hot topics.
	//
	// `WebReviewPoints`: Analyzes user viewpoints. This requires comments.
	//
	// `TimedViewPoints`: Analyzes perspectives on timeliness.
	//
	// `FreshViewPoints`: Analyzes novel perspectives.
	//
	// `TopicSummary`: Summarizes news content.
	AnalysisTypes []*string `json:"AnalysisTypes,omitempty" xml:"AnalysisTypes,omitempty" type:"Repeated"`
	// The file type. Valid values: `json` (JSON array) and `jsonLine` (JSON Lines).
	//
	// example:
	//
	// json
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The file URL. You must specify either `FileUrl` or `News`. For details on the file structure, see the description of the `News` parameter.
	//
	// example:
	//
	// http://www.example.com/xxx.json
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The maximum number of topics to analyze. By default, the service sorts clustered news by count in descending order and analyzes the top 50 topics. The maximum value is 200.
	//
	// example:
	//
	// 50
	MaxTopicSize *int32 `json:"MaxTopicSize,omitempty" xml:"MaxTopicSize,omitempty"`
	// A list of news articles. You must specify either `News` or `FileUrl`.
	News []*SubmitCustomSourceTopicAnalysisRequestNews `json:"News,omitempty" xml:"News,omitempty" type:"Repeated"`
	// A list of topics.
	Topics []*SubmitCustomSourceTopicAnalysisRequestTopics `json:"Topics,omitempty" xml:"Topics,omitempty" type:"Repeated"`
	// The URL of the file that contains the topic list. The file must be in JSON Lines format, with each line representing a single JSON object.
	//
	// example:
	//
	// http://www.example.com/xxx.jsonline
	TopicsFileUrl *string `json:"TopicsFileUrl,omitempty" xml:"TopicsFileUrl,omitempty"`
	// [The Model Studio workspace ID.](https://help.aliyun.com/document_detail/2782167.html)
	//
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
	// A list of comments.
	Comments []*SubmitCustomSourceTopicAnalysisRequestNewsComments `json:"Comments,omitempty" xml:"Comments,omitempty" type:"Repeated"`
	// The content of the news article.
	//
	// example:
	//
	// 新闻正文
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The publication time. The format must be `YYYY-MM-dd HH:mm:ss`.
	//
	// example:
	//
	// 2024-01-22 10:29:00
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// The source of the news article.
	//
	// example:
	//
	// 百度
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The title of the news article.
	//
	// example:
	//
	// 新闻标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The URL of the news article.
	//
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
	// The comment text.
	//
	// example:
	//
	// 评论内容
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
	// A custom field. You can use this field to filter results when you call the `ListHotTopics` operation.
	//
	// example:
	//
	// xxx
	CustomField *string `json:"CustomField,omitempty" xml:"CustomField,omitempty"`
	// A list of news articles.
	News []*HottopicNews `json:"News,omitempty" xml:"News,omitempty" type:"Repeated"`
	// The topic name.
	//
	// example:
	//
	// 话题名称
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The URL of the topic. This value is passed through to the `ListHotTopics` response without being processed.
	//
	// example:
	//
	// https://www.example.com/topic/123
	TopicUrl *string `json:"TopicUrl,omitempty" xml:"TopicUrl,omitempty"`
}

func (s SubmitCustomSourceTopicAnalysisRequestTopics) String() string {
	return dara.Prettify(s)
}

func (s SubmitCustomSourceTopicAnalysisRequestTopics) GoString() string {
	return s.String()
}

func (s *SubmitCustomSourceTopicAnalysisRequestTopics) GetCustomField() *string {
	return s.CustomField
}

func (s *SubmitCustomSourceTopicAnalysisRequestTopics) GetNews() []*HottopicNews {
	return s.News
}

func (s *SubmitCustomSourceTopicAnalysisRequestTopics) GetTopic() *string {
	return s.Topic
}

func (s *SubmitCustomSourceTopicAnalysisRequestTopics) GetTopicUrl() *string {
	return s.TopicUrl
}

func (s *SubmitCustomSourceTopicAnalysisRequestTopics) SetCustomField(v string) *SubmitCustomSourceTopicAnalysisRequestTopics {
	s.CustomField = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisRequestTopics) SetNews(v []*HottopicNews) *SubmitCustomSourceTopicAnalysisRequestTopics {
	s.News = v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisRequestTopics) SetTopic(v string) *SubmitCustomSourceTopicAnalysisRequestTopics {
	s.Topic = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisRequestTopics) SetTopicUrl(v string) *SubmitCustomSourceTopicAnalysisRequestTopics {
	s.TopicUrl = &v
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
