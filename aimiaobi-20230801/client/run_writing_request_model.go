// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunWritingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOriginSessionId(v string) *RunWritingRequest
	GetOriginSessionId() *string
	SetPrompt(v string) *RunWritingRequest
	GetPrompt() *string
	SetReferenceData(v *RunWritingRequestReferenceData) *RunWritingRequest
	GetReferenceData() *RunWritingRequestReferenceData
	SetSessionId(v string) *RunWritingRequest
	GetSessionId() *string
	SetTaskId(v string) *RunWritingRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *RunWritingRequest
	GetWorkspaceId() *string
	SetWritingConfig(v *RunWritingRequestWritingConfig) *RunWritingRequest
	GetWritingConfig() *RunWritingRequestWritingConfig
}

type RunWritingRequest struct {
	// The ID of the original conversation to use for regeneration.
	//
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	OriginSessionId *string `json:"OriginSessionId,omitempty" xml:"OriginSessionId,omitempty"`
	// The prompt.
	//
	// This parameter is required.
	//
	// example:
	//
	// 提示词
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// The reference article data for writing.
	ReferenceData *RunWritingRequestReferenceData `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty" type:"Struct"`
	// The ID of a single-turn conversation.
	//
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The ID of the task. You can reuse the same task ID in a multi-turn conversation.
	//
	// > You do not need to specify TaskId. The system generates one automatically. If you use the same TaskId for multiple tasks, they are grouped into a single conversation.
	//
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The unique ID of the Alibaba Cloud Model Studio workspace. For more information, see [Get a Workspace ID](https://help.aliyun.com/document_detail/2782167.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The writing configuration.
	WritingConfig *RunWritingRequestWritingConfig `json:"WritingConfig,omitempty" xml:"WritingConfig,omitempty" type:"Struct"`
}

func (s RunWritingRequest) String() string {
	return dara.Prettify(s)
}

func (s RunWritingRequest) GoString() string {
	return s.String()
}

func (s *RunWritingRequest) GetOriginSessionId() *string {
	return s.OriginSessionId
}

func (s *RunWritingRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunWritingRequest) GetReferenceData() *RunWritingRequestReferenceData {
	return s.ReferenceData
}

func (s *RunWritingRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RunWritingRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RunWritingRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunWritingRequest) GetWritingConfig() *RunWritingRequestWritingConfig {
	return s.WritingConfig
}

func (s *RunWritingRequest) SetOriginSessionId(v string) *RunWritingRequest {
	s.OriginSessionId = &v
	return s
}

func (s *RunWritingRequest) SetPrompt(v string) *RunWritingRequest {
	s.Prompt = &v
	return s
}

func (s *RunWritingRequest) SetReferenceData(v *RunWritingRequestReferenceData) *RunWritingRequest {
	s.ReferenceData = v
	return s
}

func (s *RunWritingRequest) SetSessionId(v string) *RunWritingRequest {
	s.SessionId = &v
	return s
}

func (s *RunWritingRequest) SetTaskId(v string) *RunWritingRequest {
	s.TaskId = &v
	return s
}

func (s *RunWritingRequest) SetWorkspaceId(v string) *RunWritingRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunWritingRequest) SetWritingConfig(v *RunWritingRequestWritingConfig) *RunWritingRequest {
	s.WritingConfig = v
	return s
}

func (s *RunWritingRequest) Validate() error {
	if s.ReferenceData != nil {
		if err := s.ReferenceData.Validate(); err != nil {
			return err
		}
	}
	if s.WritingConfig != nil {
		if err := s.WritingConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunWritingRequestReferenceData struct {
	// The reference article data for writing.
	Articles []*RunWritingRequestReferenceDataArticles `json:"Articles,omitempty" xml:"Articles,omitempty" type:"Repeated"`
}

func (s RunWritingRequestReferenceData) String() string {
	return dara.Prettify(s)
}

func (s RunWritingRequestReferenceData) GoString() string {
	return s.String()
}

func (s *RunWritingRequestReferenceData) GetArticles() []*RunWritingRequestReferenceDataArticles {
	return s.Articles
}

func (s *RunWritingRequestReferenceData) SetArticles(v []*RunWritingRequestReferenceDataArticles) *RunWritingRequestReferenceData {
	s.Articles = v
	return s
}

func (s *RunWritingRequestReferenceData) Validate() error {
	if s.Articles != nil {
		for _, item := range s.Articles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunWritingRequestReferenceDataArticles struct {
	// The author.
	//
	// example:
	//
	// 作者
	Author *string `json:"Author,omitempty" xml:"Author,omitempty"`
	// The content.
	//
	// example:
	//
	// 文章内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The custom unique ID of the document.
	//
	// example:
	//
	// 文档-自定义的唯一ID
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// The internal unique ID of the document.
	//
	// example:
	//
	// 2124ca4d48a542d788aa86151e1a8c8b
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// The publication time.
	//
	// example:
	//
	// 2024-08-28 11:38:28
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// The source.
	//
	// example:
	//
	// 央视网
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The article summary.
	//
	// example:
	//
	// 文章摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// The tag.
	//
	// example:
	//
	// 文章标签
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The title.
	//
	// example:
	//
	// 文章标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The URL of the article.
	//
	// example:
	//
	// https://www.example.com/aaa.docx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunWritingRequestReferenceDataArticles) String() string {
	return dara.Prettify(s)
}

func (s RunWritingRequestReferenceDataArticles) GoString() string {
	return s.String()
}

func (s *RunWritingRequestReferenceDataArticles) GetAuthor() *string {
	return s.Author
}

func (s *RunWritingRequestReferenceDataArticles) GetContent() *string {
	return s.Content
}

func (s *RunWritingRequestReferenceDataArticles) GetDocId() *string {
	return s.DocId
}

func (s *RunWritingRequestReferenceDataArticles) GetDocUuid() *string {
	return s.DocUuid
}

func (s *RunWritingRequestReferenceDataArticles) GetPubTime() *string {
	return s.PubTime
}

func (s *RunWritingRequestReferenceDataArticles) GetSource() *string {
	return s.Source
}

func (s *RunWritingRequestReferenceDataArticles) GetSummary() *string {
	return s.Summary
}

func (s *RunWritingRequestReferenceDataArticles) GetTag() *string {
	return s.Tag
}

func (s *RunWritingRequestReferenceDataArticles) GetTitle() *string {
	return s.Title
}

func (s *RunWritingRequestReferenceDataArticles) GetUrl() *string {
	return s.Url
}

func (s *RunWritingRequestReferenceDataArticles) SetAuthor(v string) *RunWritingRequestReferenceDataArticles {
	s.Author = &v
	return s
}

func (s *RunWritingRequestReferenceDataArticles) SetContent(v string) *RunWritingRequestReferenceDataArticles {
	s.Content = &v
	return s
}

func (s *RunWritingRequestReferenceDataArticles) SetDocId(v string) *RunWritingRequestReferenceDataArticles {
	s.DocId = &v
	return s
}

func (s *RunWritingRequestReferenceDataArticles) SetDocUuid(v string) *RunWritingRequestReferenceDataArticles {
	s.DocUuid = &v
	return s
}

func (s *RunWritingRequestReferenceDataArticles) SetPubTime(v string) *RunWritingRequestReferenceDataArticles {
	s.PubTime = &v
	return s
}

func (s *RunWritingRequestReferenceDataArticles) SetSource(v string) *RunWritingRequestReferenceDataArticles {
	s.Source = &v
	return s
}

func (s *RunWritingRequestReferenceDataArticles) SetSummary(v string) *RunWritingRequestReferenceDataArticles {
	s.Summary = &v
	return s
}

func (s *RunWritingRequestReferenceDataArticles) SetTag(v string) *RunWritingRequestReferenceDataArticles {
	s.Tag = &v
	return s
}

func (s *RunWritingRequestReferenceDataArticles) SetTitle(v string) *RunWritingRequestReferenceDataArticles {
	s.Title = &v
	return s
}

func (s *RunWritingRequestReferenceDataArticles) SetUrl(v string) *RunWritingRequestReferenceDataArticles {
	s.Url = &v
	return s
}

func (s *RunWritingRequestReferenceDataArticles) Validate() error {
	return dara.Validate(s)
}

type RunWritingRequestWritingConfig struct {
	// The writing domain.
	//
	// - media: Media
	//
	// - government: Government
	//
	// - market: Marketing
	//
	// example:
	//
	// media
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The prompt assistant.
	PromptTag *RunWritingRequestWritingConfigPromptTag `json:"PromptTag,omitempty" xml:"PromptTag,omitempty" type:"Struct"`
	// Control parameters for writing, such as the style, length, and output language.
	Tags []*RunWritingRequestWritingConfigTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Specifies whether to automatically add reference materials.
	//
	// example:
	//
	// true
	UseSearch *bool `json:"UseSearch,omitempty" xml:"UseSearch,omitempty"`
}

func (s RunWritingRequestWritingConfig) String() string {
	return dara.Prettify(s)
}

func (s RunWritingRequestWritingConfig) GoString() string {
	return s.String()
}

func (s *RunWritingRequestWritingConfig) GetDomain() *string {
	return s.Domain
}

func (s *RunWritingRequestWritingConfig) GetPromptTag() *RunWritingRequestWritingConfigPromptTag {
	return s.PromptTag
}

func (s *RunWritingRequestWritingConfig) GetTags() []*RunWritingRequestWritingConfigTags {
	return s.Tags
}

func (s *RunWritingRequestWritingConfig) GetUseSearch() *bool {
	return s.UseSearch
}

func (s *RunWritingRequestWritingConfig) SetDomain(v string) *RunWritingRequestWritingConfig {
	s.Domain = &v
	return s
}

func (s *RunWritingRequestWritingConfig) SetPromptTag(v *RunWritingRequestWritingConfigPromptTag) *RunWritingRequestWritingConfig {
	s.PromptTag = v
	return s
}

func (s *RunWritingRequestWritingConfig) SetTags(v []*RunWritingRequestWritingConfigTags) *RunWritingRequestWritingConfig {
	s.Tags = v
	return s
}

func (s *RunWritingRequestWritingConfig) SetUseSearch(v bool) *RunWritingRequestWritingConfig {
	s.UseSearch = &v
	return s
}

func (s *RunWritingRequestWritingConfig) Validate() error {
	if s.PromptTag != nil {
		if err := s.PromptTag.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunWritingRequestWritingConfigPromptTag struct {
	// Necessary tips.
	//
	// example:
	//
	// 必要提示
	NecessaryTips *string `json:"NecessaryTips,omitempty" xml:"NecessaryTips,omitempty"`
	// The stance.
	//
	// example:
	//
	// 立场
	Position *string `json:"Position,omitempty" xml:"Position,omitempty"`
	// Negative keywords.
	//
	// example:
	//
	// 反向词
	ReverseWords *string `json:"ReverseWords,omitempty" xml:"ReverseWords,omitempty"`
	// The theme.
	//
	// example:
	//
	// 主题
	Theme *string `json:"Theme,omitempty" xml:"Theme,omitempty"`
}

func (s RunWritingRequestWritingConfigPromptTag) String() string {
	return dara.Prettify(s)
}

func (s RunWritingRequestWritingConfigPromptTag) GoString() string {
	return s.String()
}

func (s *RunWritingRequestWritingConfigPromptTag) GetNecessaryTips() *string {
	return s.NecessaryTips
}

func (s *RunWritingRequestWritingConfigPromptTag) GetPosition() *string {
	return s.Position
}

func (s *RunWritingRequestWritingConfigPromptTag) GetReverseWords() *string {
	return s.ReverseWords
}

func (s *RunWritingRequestWritingConfigPromptTag) GetTheme() *string {
	return s.Theme
}

func (s *RunWritingRequestWritingConfigPromptTag) SetNecessaryTips(v string) *RunWritingRequestWritingConfigPromptTag {
	s.NecessaryTips = &v
	return s
}

func (s *RunWritingRequestWritingConfigPromptTag) SetPosition(v string) *RunWritingRequestWritingConfigPromptTag {
	s.Position = &v
	return s
}

func (s *RunWritingRequestWritingConfigPromptTag) SetReverseWords(v string) *RunWritingRequestWritingConfigPromptTag {
	s.ReverseWords = &v
	return s
}

func (s *RunWritingRequestWritingConfigPromptTag) SetTheme(v string) *RunWritingRequestWritingConfigPromptTag {
	s.Theme = &v
	return s
}

func (s *RunWritingRequestWritingConfigPromptTag) Validate() error {
	return dara.Validate(s)
}

type RunWritingRequestWritingConfigTags struct {
	// The value of the option.
	//
	// example:
	//
	// 10
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The tag of the option. Example: gcNumberSizeTag=10.
	//
	// example:
	//
	// gcNumberSizeTag
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s RunWritingRequestWritingConfigTags) String() string {
	return dara.Prettify(s)
}

func (s RunWritingRequestWritingConfigTags) GoString() string {
	return s.String()
}

func (s *RunWritingRequestWritingConfigTags) GetKeyword() *string {
	return s.Keyword
}

func (s *RunWritingRequestWritingConfigTags) GetTag() *string {
	return s.Tag
}

func (s *RunWritingRequestWritingConfigTags) SetKeyword(v string) *RunWritingRequestWritingConfigTags {
	s.Keyword = &v
	return s
}

func (s *RunWritingRequestWritingConfigTags) SetTag(v string) *RunWritingRequestWritingConfigTags {
	s.Tag = &v
	return s
}

func (s *RunWritingRequestWritingConfigTags) Validate() error {
	return dara.Validate(s)
}
