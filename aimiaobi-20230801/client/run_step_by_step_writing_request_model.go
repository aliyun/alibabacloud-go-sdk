// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunStepByStepWritingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOriginSessionId(v string) *RunStepByStepWritingRequest
	GetOriginSessionId() *string
	SetPrompt(v string) *RunStepByStepWritingRequest
	GetPrompt() *string
	SetReferenceData(v *RunStepByStepWritingRequestReferenceData) *RunStepByStepWritingRequest
	GetReferenceData() *RunStepByStepWritingRequestReferenceData
	SetSessionId(v string) *RunStepByStepWritingRequest
	GetSessionId() *string
	SetTaskId(v string) *RunStepByStepWritingRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *RunStepByStepWritingRequest
	GetWorkspaceId() *string
	SetWritingConfig(v *RunStepByStepWritingRequestWritingConfig) *RunStepByStepWritingRequest
	GetWritingConfig() *RunStepByStepWritingRequestWritingConfig
}

type RunStepByStepWritingRequest struct {
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	OriginSessionId *string `json:"OriginSessionId,omitempty" xml:"OriginSessionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 提示词
	Prompt        *string                                   `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	ReferenceData *RunStepByStepWritingRequestReferenceData `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty" type:"Struct"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId   *string                                   `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WritingConfig *RunStepByStepWritingRequestWritingConfig `json:"WritingConfig,omitempty" xml:"WritingConfig,omitempty" type:"Struct"`
}

func (s RunStepByStepWritingRequest) String() string {
	return dara.Prettify(s)
}

func (s RunStepByStepWritingRequest) GoString() string {
	return s.String()
}

func (s *RunStepByStepWritingRequest) GetOriginSessionId() *string {
	return s.OriginSessionId
}

func (s *RunStepByStepWritingRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunStepByStepWritingRequest) GetReferenceData() *RunStepByStepWritingRequestReferenceData {
	return s.ReferenceData
}

func (s *RunStepByStepWritingRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RunStepByStepWritingRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RunStepByStepWritingRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunStepByStepWritingRequest) GetWritingConfig() *RunStepByStepWritingRequestWritingConfig {
	return s.WritingConfig
}

func (s *RunStepByStepWritingRequest) SetOriginSessionId(v string) *RunStepByStepWritingRequest {
	s.OriginSessionId = &v
	return s
}

func (s *RunStepByStepWritingRequest) SetPrompt(v string) *RunStepByStepWritingRequest {
	s.Prompt = &v
	return s
}

func (s *RunStepByStepWritingRequest) SetReferenceData(v *RunStepByStepWritingRequestReferenceData) *RunStepByStepWritingRequest {
	s.ReferenceData = v
	return s
}

func (s *RunStepByStepWritingRequest) SetSessionId(v string) *RunStepByStepWritingRequest {
	s.SessionId = &v
	return s
}

func (s *RunStepByStepWritingRequest) SetTaskId(v string) *RunStepByStepWritingRequest {
	s.TaskId = &v
	return s
}

func (s *RunStepByStepWritingRequest) SetWorkspaceId(v string) *RunStepByStepWritingRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunStepByStepWritingRequest) SetWritingConfig(v *RunStepByStepWritingRequestWritingConfig) *RunStepByStepWritingRequest {
	s.WritingConfig = v
	return s
}

func (s *RunStepByStepWritingRequest) Validate() error {
	return dara.Validate(s)
}

type RunStepByStepWritingRequestReferenceData struct {
	Articles      []*RunStepByStepWritingRequestReferenceDataArticles `json:"Articles,omitempty" xml:"Articles,omitempty" type:"Repeated"`
	MiniDoc       []*string                                           `json:"MiniDoc,omitempty" xml:"MiniDoc,omitempty" type:"Repeated"`
	Outlines      []*RunStepByStepWritingRequestReferenceDataOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	Summarization []*string                                           `json:"Summarization,omitempty" xml:"Summarization,omitempty" type:"Repeated"`
}

func (s RunStepByStepWritingRequestReferenceData) String() string {
	return dara.Prettify(s)
}

func (s RunStepByStepWritingRequestReferenceData) GoString() string {
	return s.String()
}

func (s *RunStepByStepWritingRequestReferenceData) GetArticles() []*RunStepByStepWritingRequestReferenceDataArticles {
	return s.Articles
}

func (s *RunStepByStepWritingRequestReferenceData) GetMiniDoc() []*string {
	return s.MiniDoc
}

func (s *RunStepByStepWritingRequestReferenceData) GetOutlines() []*RunStepByStepWritingRequestReferenceDataOutlines {
	return s.Outlines
}

func (s *RunStepByStepWritingRequestReferenceData) GetSummarization() []*string {
	return s.Summarization
}

func (s *RunStepByStepWritingRequestReferenceData) SetArticles(v []*RunStepByStepWritingRequestReferenceDataArticles) *RunStepByStepWritingRequestReferenceData {
	s.Articles = v
	return s
}

func (s *RunStepByStepWritingRequestReferenceData) SetMiniDoc(v []*string) *RunStepByStepWritingRequestReferenceData {
	s.MiniDoc = v
	return s
}

func (s *RunStepByStepWritingRequestReferenceData) SetOutlines(v []*RunStepByStepWritingRequestReferenceDataOutlines) *RunStepByStepWritingRequestReferenceData {
	s.Outlines = v
	return s
}

func (s *RunStepByStepWritingRequestReferenceData) SetSummarization(v []*string) *RunStepByStepWritingRequestReferenceData {
	s.Summarization = v
	return s
}

func (s *RunStepByStepWritingRequestReferenceData) Validate() error {
	return dara.Validate(s)
}

type RunStepByStepWritingRequestReferenceDataArticles struct {
	// example:
	//
	// 作者
	Author *string `json:"Author,omitempty" xml:"Author,omitempty"`
	// example:
	//
	// 文章内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 文档-自定义的唯一ID
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// 8a20e007a6174522af4d6a2657d5526f
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// http://www.example.com
	MediaUrl *string `json:"MediaUrl,omitempty" xml:"MediaUrl,omitempty"`
	// example:
	//
	// 2024-09-10 14:17:54
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// 央视网
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 文章摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// 文章标签
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// example:
	//
	// 文章标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// https://www.example.com/aaa.docx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunStepByStepWritingRequestReferenceDataArticles) String() string {
	return dara.Prettify(s)
}

func (s RunStepByStepWritingRequestReferenceDataArticles) GoString() string {
	return s.String()
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) GetAuthor() *string {
	return s.Author
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) GetContent() *string {
	return s.Content
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) GetDocId() *string {
	return s.DocId
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) GetDocUuid() *string {
	return s.DocUuid
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) GetMediaUrl() *string {
	return s.MediaUrl
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) GetPubTime() *string {
	return s.PubTime
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) GetSource() *string {
	return s.Source
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) GetSummary() *string {
	return s.Summary
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) GetTag() *string {
	return s.Tag
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) GetTitle() *string {
	return s.Title
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) GetUrl() *string {
	return s.Url
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) SetAuthor(v string) *RunStepByStepWritingRequestReferenceDataArticles {
	s.Author = &v
	return s
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) SetContent(v string) *RunStepByStepWritingRequestReferenceDataArticles {
	s.Content = &v
	return s
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) SetDocId(v string) *RunStepByStepWritingRequestReferenceDataArticles {
	s.DocId = &v
	return s
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) SetDocUuid(v string) *RunStepByStepWritingRequestReferenceDataArticles {
	s.DocUuid = &v
	return s
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) SetMediaUrl(v string) *RunStepByStepWritingRequestReferenceDataArticles {
	s.MediaUrl = &v
	return s
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) SetPubTime(v string) *RunStepByStepWritingRequestReferenceDataArticles {
	s.PubTime = &v
	return s
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) SetSource(v string) *RunStepByStepWritingRequestReferenceDataArticles {
	s.Source = &v
	return s
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) SetSummary(v string) *RunStepByStepWritingRequestReferenceDataArticles {
	s.Summary = &v
	return s
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) SetTag(v string) *RunStepByStepWritingRequestReferenceDataArticles {
	s.Tag = &v
	return s
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) SetTitle(v string) *RunStepByStepWritingRequestReferenceDataArticles {
	s.Title = &v
	return s
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) SetUrl(v string) *RunStepByStepWritingRequestReferenceDataArticles {
	s.Url = &v
	return s
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) Validate() error {
	return dara.Validate(s)
}

type RunStepByStepWritingRequestReferenceDataOutlines struct {
	Articles []*RunStepByStepWritingRequestReferenceDataOutlinesArticles `json:"Articles,omitempty" xml:"Articles,omitempty" type:"Repeated"`
	// example:
	//
	// 大纲
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
}

func (s RunStepByStepWritingRequestReferenceDataOutlines) String() string {
	return dara.Prettify(s)
}

func (s RunStepByStepWritingRequestReferenceDataOutlines) GoString() string {
	return s.String()
}

func (s *RunStepByStepWritingRequestReferenceDataOutlines) GetArticles() []*RunStepByStepWritingRequestReferenceDataOutlinesArticles {
	return s.Articles
}

func (s *RunStepByStepWritingRequestReferenceDataOutlines) GetOutline() *string {
	return s.Outline
}

func (s *RunStepByStepWritingRequestReferenceDataOutlines) SetArticles(v []*RunStepByStepWritingRequestReferenceDataOutlinesArticles) *RunStepByStepWritingRequestReferenceDataOutlines {
	s.Articles = v
	return s
}

func (s *RunStepByStepWritingRequestReferenceDataOutlines) SetOutline(v string) *RunStepByStepWritingRequestReferenceDataOutlines {
	s.Outline = &v
	return s
}

func (s *RunStepByStepWritingRequestReferenceDataOutlines) Validate() error {
	return dara.Validate(s)
}

type RunStepByStepWritingRequestReferenceDataOutlinesArticles struct {
	// example:
	//
	// 文章内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 文章标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 文章链接
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunStepByStepWritingRequestReferenceDataOutlinesArticles) String() string {
	return dara.Prettify(s)
}

func (s RunStepByStepWritingRequestReferenceDataOutlinesArticles) GoString() string {
	return s.String()
}

func (s *RunStepByStepWritingRequestReferenceDataOutlinesArticles) GetContent() *string {
	return s.Content
}

func (s *RunStepByStepWritingRequestReferenceDataOutlinesArticles) GetTitle() *string {
	return s.Title
}

func (s *RunStepByStepWritingRequestReferenceDataOutlinesArticles) GetUrl() *string {
	return s.Url
}

func (s *RunStepByStepWritingRequestReferenceDataOutlinesArticles) SetContent(v string) *RunStepByStepWritingRequestReferenceDataOutlinesArticles {
	s.Content = &v
	return s
}

func (s *RunStepByStepWritingRequestReferenceDataOutlinesArticles) SetTitle(v string) *RunStepByStepWritingRequestReferenceDataOutlinesArticles {
	s.Title = &v
	return s
}

func (s *RunStepByStepWritingRequestReferenceDataOutlinesArticles) SetUrl(v string) *RunStepByStepWritingRequestReferenceDataOutlinesArticles {
	s.Url = &v
	return s
}

func (s *RunStepByStepWritingRequestReferenceDataOutlinesArticles) Validate() error {
	return dara.Validate(s)
}

type RunStepByStepWritingRequestWritingConfig struct {
	// example:
	//
	// media
	Domain    *string                                            `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Keywords  []*string                                          `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Repeated"`
	PromptTag *RunStepByStepWritingRequestWritingConfigPromptTag `json:"PromptTag,omitempty" xml:"PromptTag,omitempty" type:"Struct"`
	// example:
	//
	// 分步骤写作场景，传媒写作支持的写作场景:新闻写作(默认),新闻评论,通用文体，公文写作支持的写作场景:通知(默认),通告,通报,请示,决定,函,通用文体
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// example:
	//
	// Writing
	Step              *string                                         `json:"Step,omitempty" xml:"Step,omitempty"`
	SummaryReturnType *string                                         `json:"SummaryReturnType,omitempty" xml:"SummaryReturnType,omitempty"`
	Tags              []*RunStepByStepWritingRequestWritingConfigTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// true
	UseSearch *bool `json:"UseSearch,omitempty" xml:"UseSearch,omitempty"`
}

func (s RunStepByStepWritingRequestWritingConfig) String() string {
	return dara.Prettify(s)
}

func (s RunStepByStepWritingRequestWritingConfig) GoString() string {
	return s.String()
}

func (s *RunStepByStepWritingRequestWritingConfig) GetDomain() *string {
	return s.Domain
}

func (s *RunStepByStepWritingRequestWritingConfig) GetKeywords() []*string {
	return s.Keywords
}

func (s *RunStepByStepWritingRequestWritingConfig) GetPromptTag() *RunStepByStepWritingRequestWritingConfigPromptTag {
	return s.PromptTag
}

func (s *RunStepByStepWritingRequestWritingConfig) GetScene() *string {
	return s.Scene
}

func (s *RunStepByStepWritingRequestWritingConfig) GetStep() *string {
	return s.Step
}

func (s *RunStepByStepWritingRequestWritingConfig) GetSummaryReturnType() *string {
	return s.SummaryReturnType
}

func (s *RunStepByStepWritingRequestWritingConfig) GetTags() []*RunStepByStepWritingRequestWritingConfigTags {
	return s.Tags
}

func (s *RunStepByStepWritingRequestWritingConfig) GetUseSearch() *bool {
	return s.UseSearch
}

func (s *RunStepByStepWritingRequestWritingConfig) SetDomain(v string) *RunStepByStepWritingRequestWritingConfig {
	s.Domain = &v
	return s
}

func (s *RunStepByStepWritingRequestWritingConfig) SetKeywords(v []*string) *RunStepByStepWritingRequestWritingConfig {
	s.Keywords = v
	return s
}

func (s *RunStepByStepWritingRequestWritingConfig) SetPromptTag(v *RunStepByStepWritingRequestWritingConfigPromptTag) *RunStepByStepWritingRequestWritingConfig {
	s.PromptTag = v
	return s
}

func (s *RunStepByStepWritingRequestWritingConfig) SetScene(v string) *RunStepByStepWritingRequestWritingConfig {
	s.Scene = &v
	return s
}

func (s *RunStepByStepWritingRequestWritingConfig) SetStep(v string) *RunStepByStepWritingRequestWritingConfig {
	s.Step = &v
	return s
}

func (s *RunStepByStepWritingRequestWritingConfig) SetSummaryReturnType(v string) *RunStepByStepWritingRequestWritingConfig {
	s.SummaryReturnType = &v
	return s
}

func (s *RunStepByStepWritingRequestWritingConfig) SetTags(v []*RunStepByStepWritingRequestWritingConfigTags) *RunStepByStepWritingRequestWritingConfig {
	s.Tags = v
	return s
}

func (s *RunStepByStepWritingRequestWritingConfig) SetUseSearch(v bool) *RunStepByStepWritingRequestWritingConfig {
	s.UseSearch = &v
	return s
}

func (s *RunStepByStepWritingRequestWritingConfig) Validate() error {
	return dara.Validate(s)
}

type RunStepByStepWritingRequestWritingConfigPromptTag struct {
	// example:
	//
	// 必要提示
	NecessaryTips *string `json:"NecessaryTips,omitempty" xml:"NecessaryTips,omitempty"`
	// example:
	//
	// 立场
	Position *string `json:"Position,omitempty" xml:"Position,omitempty"`
	// example:
	//
	// 反向词
	ReverseWords *string `json:"ReverseWords,omitempty" xml:"ReverseWords,omitempty"`
	// example:
	//
	// 主题
	Theme *string `json:"Theme,omitempty" xml:"Theme,omitempty"`
}

func (s RunStepByStepWritingRequestWritingConfigPromptTag) String() string {
	return dara.Prettify(s)
}

func (s RunStepByStepWritingRequestWritingConfigPromptTag) GoString() string {
	return s.String()
}

func (s *RunStepByStepWritingRequestWritingConfigPromptTag) GetNecessaryTips() *string {
	return s.NecessaryTips
}

func (s *RunStepByStepWritingRequestWritingConfigPromptTag) GetPosition() *string {
	return s.Position
}

func (s *RunStepByStepWritingRequestWritingConfigPromptTag) GetReverseWords() *string {
	return s.ReverseWords
}

func (s *RunStepByStepWritingRequestWritingConfigPromptTag) GetTheme() *string {
	return s.Theme
}

func (s *RunStepByStepWritingRequestWritingConfigPromptTag) SetNecessaryTips(v string) *RunStepByStepWritingRequestWritingConfigPromptTag {
	s.NecessaryTips = &v
	return s
}

func (s *RunStepByStepWritingRequestWritingConfigPromptTag) SetPosition(v string) *RunStepByStepWritingRequestWritingConfigPromptTag {
	s.Position = &v
	return s
}

func (s *RunStepByStepWritingRequestWritingConfigPromptTag) SetReverseWords(v string) *RunStepByStepWritingRequestWritingConfigPromptTag {
	s.ReverseWords = &v
	return s
}

func (s *RunStepByStepWritingRequestWritingConfigPromptTag) SetTheme(v string) *RunStepByStepWritingRequestWritingConfigPromptTag {
	s.Theme = &v
	return s
}

func (s *RunStepByStepWritingRequestWritingConfigPromptTag) Validate() error {
	return dara.Validate(s)
}

type RunStepByStepWritingRequestWritingConfigTags struct {
	// example:
	//
	// 10
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// gcNumberSizeTag
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s RunStepByStepWritingRequestWritingConfigTags) String() string {
	return dara.Prettify(s)
}

func (s RunStepByStepWritingRequestWritingConfigTags) GoString() string {
	return s.String()
}

func (s *RunStepByStepWritingRequestWritingConfigTags) GetKeyword() *string {
	return s.Keyword
}

func (s *RunStepByStepWritingRequestWritingConfigTags) GetTag() *string {
	return s.Tag
}

func (s *RunStepByStepWritingRequestWritingConfigTags) SetKeyword(v string) *RunStepByStepWritingRequestWritingConfigTags {
	s.Keyword = &v
	return s
}

func (s *RunStepByStepWritingRequestWritingConfigTags) SetTag(v string) *RunStepByStepWritingRequestWritingConfigTags {
	s.Tag = &v
	return s
}

func (s *RunStepByStepWritingRequestWritingConfigTags) Validate() error {
	return dara.Validate(s)
}
