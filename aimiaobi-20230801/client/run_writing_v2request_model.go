// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunWritingV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetArticles(v []*RunWritingV2RequestArticles) *RunWritingV2Request
	GetArticles() []*RunWritingV2RequestArticles
	SetDistributeWriting(v bool) *RunWritingV2Request
	GetDistributeWriting() *bool
	SetGcNumberSize(v int32) *RunWritingV2Request
	GetGcNumberSize() *int32
	SetGcNumberSizeTag(v string) *RunWritingV2Request
	GetGcNumberSizeTag() *string
	SetKeywords(v []*string) *RunWritingV2Request
	GetKeywords() []*string
	SetLanguage(v string) *RunWritingV2Request
	GetLanguage() *string
	SetMiniDocs(v []*RunWritingV2RequestMiniDocs) *RunWritingV2Request
	GetMiniDocs() []*RunWritingV2RequestMiniDocs
	SetOutlineList(v []*WritingOutline) *RunWritingV2Request
	GetOutlineList() []*WritingOutline
	SetOutlines(v []*RunWritingV2RequestOutlines) *RunWritingV2Request
	GetOutlines() []*RunWritingV2RequestOutlines
	SetPrompt(v string) *RunWritingV2Request
	GetPrompt() *string
	SetPromptMode(v string) *RunWritingV2Request
	GetPromptMode() *string
	SetSearchSources(v []*RunWritingV2RequestSearchSources) *RunWritingV2Request
	GetSearchSources() []*RunWritingV2RequestSearchSources
	SetSessionId(v string) *RunWritingV2Request
	GetSessionId() *string
	SetSourceTraceMethod(v string) *RunWritingV2Request
	GetSourceTraceMethod() *string
	SetStep(v string) *RunWritingV2Request
	GetStep() *string
	SetSummarization(v []*RunWritingV2RequestSummarization) *RunWritingV2Request
	GetSummarization() []*RunWritingV2RequestSummarization
	SetTaskId(v string) *RunWritingV2Request
	GetTaskId() *string
	SetUseSearch(v bool) *RunWritingV2Request
	GetUseSearch() *bool
	SetWorkspaceId(v string) *RunWritingV2Request
	GetWorkspaceId() *string
	SetWritingParams(v map[string]*string) *RunWritingV2Request
	GetWritingParams() map[string]*string
	SetWritingScene(v string) *RunWritingV2Request
	GetWritingScene() *string
	SetWritingStyle(v string) *RunWritingV2Request
	GetWritingStyle() *string
}

type RunWritingV2Request struct {
	// A list of articles to use as references. **Note:*	- When you provide this parameter, web search is disabled, overriding the `UseSearch` and `SearchSources` parameters.
	Articles []*RunWritingV2RequestArticles `json:"Articles,omitempty" xml:"Articles,omitempty" type:"Repeated"`
	// Specifies whether to enable step-by-step writing. For more information, see the `Step` parameter description.
	//
	// example:
	//
	// false
	DistributeWriting *bool `json:"DistributeWriting,omitempty" xml:"DistributeWriting,omitempty"`
	// The number of articles to write. If you request multiple articles, the system returns them concurrently, each with a unique session ID.
	//
	// example:
	//
	// 2
	GcNumberSize *int32 `json:"GcNumberSize,omitempty" xml:"GcNumberSize,omitempty"`
	// A string that specifies the desired article length. Examples: "about 300 words", "about 600 words", "about 1,000 words", or "about 2,000 words".
	//
	// example:
	//
	// 2000字左右
	GcNumberSizeTag *string `json:"GcNumberSizeTag,omitempty" xml:"GcNumberSizeTag,omitempty"`
	// A list of keywords used for both search and writing.
	Keywords []*string `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Repeated"`
	// The output language for the article.
	//
	// - `en`: English
	//
	// - `zh`: Chinese
	//
	// - Other languages or specific style requirements can also be specified.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// A list of article snippets.
	MiniDocs []*RunWritingV2RequestMiniDocs `json:"MiniDocs,omitempty" xml:"MiniDocs,omitempty" type:"Repeated"`
	// A list of outlines for step-by-step writing.
	OutlineList []*WritingOutline `json:"OutlineList,omitempty" xml:"OutlineList,omitempty" type:"Repeated"`
	// A list of outlines for step-by-step writing. This parameter is deprecated. Use `OutlineList` instead.
	Outlines []*RunWritingV2RequestOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	// The writing prompt. You must provide either `Prompt` or `WritingParams`. For more information, see the description of the `PromptMode` parameter.
	//
	// example:
	//
	// 提示词
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// The prompt mode. Valid values: `Template` (template mode) and `PE` (advanced PE mode).
	//
	// 1. If this parameter is omitted, you must provide the `Prompt` parameter. We recommend that the prompt includes the topic, length, requirements, and prohibitions.
	//
	// 2. If `PromptMode` is set to `Template`, you must provide `WritingParams`, which is a dictionary of string key-value pairs. For the required schema, see the `.Data.TemplateDefine[].Fields` field in the response of the [ListWritingStyles](https://help.aliyun.com/document_detail/2922609.html) operation.
	//
	// 3. If `PromptMode` is set to `PE`, you must pass `WritingParams` with the following two fields:
	//
	//    1. `topic`: Required. The topic to write about.
	//
	//    2. `prompt`: Optional. Any additional custom prompts or writing requirements.
	//
	// example:
	//
	// Template
	PromptMode *string `json:"PromptMode,omitempty" xml:"PromptMode,omitempty"`
	// A list of specified search sources to use.
	SearchSources []*RunWritingV2RequestSearchSources `json:"SearchSources,omitempty" xml:"SearchSources,omitempty" type:"Repeated"`
	// The ID of a single-turn conversation. This parameter is deprecated and its use is discouraged.
	//
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The source tracing method. Currently, only `modelSourceTrace` is supported. If set to `modelSourceTrace`, the model adds citation markers (for example, `[[1]]`) to the end of each cited snippet in the generated text. The citation index starts at 1.
	//
	// example:
	//
	// modelSourceTrace
	SourceTraceMethod *string `json:"SourceTraceMethod,omitempty" xml:"SourceTraceMethod,omitempty"`
	// The step for step-by-step writing. Valid values:
	//
	// - `OutlineGenerate`: Outline generation
	//
	// - `Writing`: Article writing
	//
	// When `DistributeWriting` is `true`, the default flow for step-by-step writing is to first generate an outline and then write the content based on it.
	//
	// example:
	//
	// Writing
	Step *string `json:"Step,omitempty" xml:"Step,omitempty"`
	// A list of summarization objects, used for step-by-step writing.
	Summarization []*RunWritingV2RequestSummarization `json:"Summarization,omitempty" xml:"Summarization,omitempty" type:"Repeated"`
	// The unique ID of the task. You can reuse the same task ID for a multi-turn conversation.
	//
	// > The system automatically generates a `TaskId` if you do not specify one. Reusing the same `TaskId` for subsequent requests groups them into a single conversation.
	//
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Specifies whether to enable web search. If `true`, the system uses its built-in web search feature. Default: `false`.
	//
	// example:
	//
	// true
	UseSearch *bool `json:"UseSearch,omitempty" xml:"UseSearch,omitempty"`
	// The unique ID of the Model Studio workspace. For more information, see [Obtain a Workspace ID](https://help.aliyun.com/document_detail/2782167.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The parameters for template-based writing, provided as a dictionary of string key-value pairs. You must provide either `Prompt` or `WritingParams`. For more information, see the description of the `PromptMode` parameter.
	WritingParams map[string]*string `json:"WritingParams,omitempty" xml:"WritingParams,omitempty"`
	// The writing scene. Valid values: `government` (government affairs), `media`, `market` (marketing), `office`, and `custom`.
	//
	// example:
	//
	// media
	WritingScene *string `json:"WritingScene,omitempty" xml:"WritingScene,omitempty"`
	// The writing style. For a list of supported styles, see [ListWritingStyles](https://help.aliyun.com/document_detail/2922609.html).
	//
	// example:
	//
	// 新闻评论
	WritingStyle *string `json:"WritingStyle,omitempty" xml:"WritingStyle,omitempty"`
}

func (s RunWritingV2Request) String() string {
	return dara.Prettify(s)
}

func (s RunWritingV2Request) GoString() string {
	return s.String()
}

func (s *RunWritingV2Request) GetArticles() []*RunWritingV2RequestArticles {
	return s.Articles
}

func (s *RunWritingV2Request) GetDistributeWriting() *bool {
	return s.DistributeWriting
}

func (s *RunWritingV2Request) GetGcNumberSize() *int32 {
	return s.GcNumberSize
}

func (s *RunWritingV2Request) GetGcNumberSizeTag() *string {
	return s.GcNumberSizeTag
}

func (s *RunWritingV2Request) GetKeywords() []*string {
	return s.Keywords
}

func (s *RunWritingV2Request) GetLanguage() *string {
	return s.Language
}

func (s *RunWritingV2Request) GetMiniDocs() []*RunWritingV2RequestMiniDocs {
	return s.MiniDocs
}

func (s *RunWritingV2Request) GetOutlineList() []*WritingOutline {
	return s.OutlineList
}

func (s *RunWritingV2Request) GetOutlines() []*RunWritingV2RequestOutlines {
	return s.Outlines
}

func (s *RunWritingV2Request) GetPrompt() *string {
	return s.Prompt
}

func (s *RunWritingV2Request) GetPromptMode() *string {
	return s.PromptMode
}

func (s *RunWritingV2Request) GetSearchSources() []*RunWritingV2RequestSearchSources {
	return s.SearchSources
}

func (s *RunWritingV2Request) GetSessionId() *string {
	return s.SessionId
}

func (s *RunWritingV2Request) GetSourceTraceMethod() *string {
	return s.SourceTraceMethod
}

func (s *RunWritingV2Request) GetStep() *string {
	return s.Step
}

func (s *RunWritingV2Request) GetSummarization() []*RunWritingV2RequestSummarization {
	return s.Summarization
}

func (s *RunWritingV2Request) GetTaskId() *string {
	return s.TaskId
}

func (s *RunWritingV2Request) GetUseSearch() *bool {
	return s.UseSearch
}

func (s *RunWritingV2Request) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunWritingV2Request) GetWritingParams() map[string]*string {
	return s.WritingParams
}

func (s *RunWritingV2Request) GetWritingScene() *string {
	return s.WritingScene
}

func (s *RunWritingV2Request) GetWritingStyle() *string {
	return s.WritingStyle
}

func (s *RunWritingV2Request) SetArticles(v []*RunWritingV2RequestArticles) *RunWritingV2Request {
	s.Articles = v
	return s
}

func (s *RunWritingV2Request) SetDistributeWriting(v bool) *RunWritingV2Request {
	s.DistributeWriting = &v
	return s
}

func (s *RunWritingV2Request) SetGcNumberSize(v int32) *RunWritingV2Request {
	s.GcNumberSize = &v
	return s
}

func (s *RunWritingV2Request) SetGcNumberSizeTag(v string) *RunWritingV2Request {
	s.GcNumberSizeTag = &v
	return s
}

func (s *RunWritingV2Request) SetKeywords(v []*string) *RunWritingV2Request {
	s.Keywords = v
	return s
}

func (s *RunWritingV2Request) SetLanguage(v string) *RunWritingV2Request {
	s.Language = &v
	return s
}

func (s *RunWritingV2Request) SetMiniDocs(v []*RunWritingV2RequestMiniDocs) *RunWritingV2Request {
	s.MiniDocs = v
	return s
}

func (s *RunWritingV2Request) SetOutlineList(v []*WritingOutline) *RunWritingV2Request {
	s.OutlineList = v
	return s
}

func (s *RunWritingV2Request) SetOutlines(v []*RunWritingV2RequestOutlines) *RunWritingV2Request {
	s.Outlines = v
	return s
}

func (s *RunWritingV2Request) SetPrompt(v string) *RunWritingV2Request {
	s.Prompt = &v
	return s
}

func (s *RunWritingV2Request) SetPromptMode(v string) *RunWritingV2Request {
	s.PromptMode = &v
	return s
}

func (s *RunWritingV2Request) SetSearchSources(v []*RunWritingV2RequestSearchSources) *RunWritingV2Request {
	s.SearchSources = v
	return s
}

func (s *RunWritingV2Request) SetSessionId(v string) *RunWritingV2Request {
	s.SessionId = &v
	return s
}

func (s *RunWritingV2Request) SetSourceTraceMethod(v string) *RunWritingV2Request {
	s.SourceTraceMethod = &v
	return s
}

func (s *RunWritingV2Request) SetStep(v string) *RunWritingV2Request {
	s.Step = &v
	return s
}

func (s *RunWritingV2Request) SetSummarization(v []*RunWritingV2RequestSummarization) *RunWritingV2Request {
	s.Summarization = v
	return s
}

func (s *RunWritingV2Request) SetTaskId(v string) *RunWritingV2Request {
	s.TaskId = &v
	return s
}

func (s *RunWritingV2Request) SetUseSearch(v bool) *RunWritingV2Request {
	s.UseSearch = &v
	return s
}

func (s *RunWritingV2Request) SetWorkspaceId(v string) *RunWritingV2Request {
	s.WorkspaceId = &v
	return s
}

func (s *RunWritingV2Request) SetWritingParams(v map[string]*string) *RunWritingV2Request {
	s.WritingParams = v
	return s
}

func (s *RunWritingV2Request) SetWritingScene(v string) *RunWritingV2Request {
	s.WritingScene = &v
	return s
}

func (s *RunWritingV2Request) SetWritingStyle(v string) *RunWritingV2Request {
	s.WritingStyle = &v
	return s
}

func (s *RunWritingV2Request) Validate() error {
	if s.Articles != nil {
		for _, item := range s.Articles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MiniDocs != nil {
		for _, item := range s.MiniDocs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OutlineList != nil {
		for _, item := range s.OutlineList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Outlines != nil {
		for _, item := range s.Outlines {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SearchSources != nil {
		for _, item := range s.SearchSources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Summarization != nil {
		for _, item := range s.Summarization {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunWritingV2RequestArticles struct {
	// The content of the article.
	//
	// example:
	//
	// 文章内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The publication time.
	//
	// example:
	//
	// 2024-11-25 14:25:59
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// The name of the search source.
	//
	// example:
	//
	// QuarkCommonNews
	SearchSourceName *string `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
	// The source of the article.
	//
	// example:
	//
	// 新华社
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The title of the article.
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

func (s RunWritingV2RequestArticles) String() string {
	return dara.Prettify(s)
}

func (s RunWritingV2RequestArticles) GoString() string {
	return s.String()
}

func (s *RunWritingV2RequestArticles) GetContent() *string {
	return s.Content
}

func (s *RunWritingV2RequestArticles) GetPubTime() *string {
	return s.PubTime
}

func (s *RunWritingV2RequestArticles) GetSearchSourceName() *string {
	return s.SearchSourceName
}

func (s *RunWritingV2RequestArticles) GetSource() *string {
	return s.Source
}

func (s *RunWritingV2RequestArticles) GetTitle() *string {
	return s.Title
}

func (s *RunWritingV2RequestArticles) GetUrl() *string {
	return s.Url
}

func (s *RunWritingV2RequestArticles) SetContent(v string) *RunWritingV2RequestArticles {
	s.Content = &v
	return s
}

func (s *RunWritingV2RequestArticles) SetPubTime(v string) *RunWritingV2RequestArticles {
	s.PubTime = &v
	return s
}

func (s *RunWritingV2RequestArticles) SetSearchSourceName(v string) *RunWritingV2RequestArticles {
	s.SearchSourceName = &v
	return s
}

func (s *RunWritingV2RequestArticles) SetSource(v string) *RunWritingV2RequestArticles {
	s.Source = &v
	return s
}

func (s *RunWritingV2RequestArticles) SetTitle(v string) *RunWritingV2RequestArticles {
	s.Title = &v
	return s
}

func (s *RunWritingV2RequestArticles) SetUrl(v string) *RunWritingV2RequestArticles {
	s.Url = &v
	return s
}

func (s *RunWritingV2RequestArticles) Validate() error {
	return dara.Validate(s)
}

type RunWritingV2RequestMiniDocs struct {
	// The content of the snippet.
	//
	// example:
	//
	// 片段内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The index of the article snippet in the referenced article.
	//
	// example:
	//
	// 索引
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// Specifies whether to prioritize this snippet.
	//
	// example:
	//
	// true
	Star *bool `json:"Star,omitempty" xml:"Star,omitempty"`
}

func (s RunWritingV2RequestMiniDocs) String() string {
	return dara.Prettify(s)
}

func (s RunWritingV2RequestMiniDocs) GoString() string {
	return s.String()
}

func (s *RunWritingV2RequestMiniDocs) GetContent() *string {
	return s.Content
}

func (s *RunWritingV2RequestMiniDocs) GetIndex() *string {
	return s.Index
}

func (s *RunWritingV2RequestMiniDocs) GetStar() *bool {
	return s.Star
}

func (s *RunWritingV2RequestMiniDocs) SetContent(v string) *RunWritingV2RequestMiniDocs {
	s.Content = &v
	return s
}

func (s *RunWritingV2RequestMiniDocs) SetIndex(v string) *RunWritingV2RequestMiniDocs {
	s.Index = &v
	return s
}

func (s *RunWritingV2RequestMiniDocs) SetStar(v bool) *RunWritingV2RequestMiniDocs {
	s.Star = &v
	return s
}

func (s *RunWritingV2RequestMiniDocs) Validate() error {
	return dara.Validate(s)
}

type RunWritingV2RequestOutlines struct {
	// A list of articles referenced by the outline.
	Articles []*RunWritingV2RequestOutlinesArticles `json:"Articles,omitempty" xml:"Articles,omitempty" type:"Repeated"`
	// The outline.
	//
	// example:
	//
	// 大纲
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
}

func (s RunWritingV2RequestOutlines) String() string {
	return dara.Prettify(s)
}

func (s RunWritingV2RequestOutlines) GoString() string {
	return s.String()
}

func (s *RunWritingV2RequestOutlines) GetArticles() []*RunWritingV2RequestOutlinesArticles {
	return s.Articles
}

func (s *RunWritingV2RequestOutlines) GetOutline() *string {
	return s.Outline
}

func (s *RunWritingV2RequestOutlines) SetArticles(v []*RunWritingV2RequestOutlinesArticles) *RunWritingV2RequestOutlines {
	s.Articles = v
	return s
}

func (s *RunWritingV2RequestOutlines) SetOutline(v string) *RunWritingV2RequestOutlines {
	s.Outline = &v
	return s
}

func (s *RunWritingV2RequestOutlines) Validate() error {
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

type RunWritingV2RequestOutlinesArticles struct {
	// The content of the article.
	//
	// example:
	//
	// 正文内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The title of the article.
	//
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The URL of the article.
	//
	// example:
	//
	// 文章URL
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunWritingV2RequestOutlinesArticles) String() string {
	return dara.Prettify(s)
}

func (s RunWritingV2RequestOutlinesArticles) GoString() string {
	return s.String()
}

func (s *RunWritingV2RequestOutlinesArticles) GetContent() *string {
	return s.Content
}

func (s *RunWritingV2RequestOutlinesArticles) GetTitle() *string {
	return s.Title
}

func (s *RunWritingV2RequestOutlinesArticles) GetUrl() *string {
	return s.Url
}

func (s *RunWritingV2RequestOutlinesArticles) SetContent(v string) *RunWritingV2RequestOutlinesArticles {
	s.Content = &v
	return s
}

func (s *RunWritingV2RequestOutlinesArticles) SetTitle(v string) *RunWritingV2RequestOutlinesArticles {
	s.Title = &v
	return s
}

func (s *RunWritingV2RequestOutlinesArticles) SetUrl(v string) *RunWritingV2RequestOutlinesArticles {
	s.Url = &v
	return s
}

func (s *RunWritingV2RequestOutlinesArticles) Validate() error {
	return dara.Validate(s)
}

type RunWritingV2RequestSearchSources struct {
	// The type of search source. Valid values: `SystemSearch` (built-in system search), `CustomSemanticSearch` (search of a custom semantic index), and `ThirdSearch` (search through a third-party API).
	//
	// example:
	//
	// SystemSearch
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The unique identifier of the data source.
	//
	// example:
	//
	// QuarkCommonNews
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The description of the search source. This parameter is deprecated and has no effect.
	//
	// example:
	//
	// 互联网搜索
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s RunWritingV2RequestSearchSources) String() string {
	return dara.Prettify(s)
}

func (s RunWritingV2RequestSearchSources) GoString() string {
	return s.String()
}

func (s *RunWritingV2RequestSearchSources) GetCode() *string {
	return s.Code
}

func (s *RunWritingV2RequestSearchSources) GetDatasetName() *string {
	return s.DatasetName
}

func (s *RunWritingV2RequestSearchSources) GetName() *string {
	return s.Name
}

func (s *RunWritingV2RequestSearchSources) SetCode(v string) *RunWritingV2RequestSearchSources {
	s.Code = &v
	return s
}

func (s *RunWritingV2RequestSearchSources) SetDatasetName(v string) *RunWritingV2RequestSearchSources {
	s.DatasetName = &v
	return s
}

func (s *RunWritingV2RequestSearchSources) SetName(v string) *RunWritingV2RequestSearchSources {
	s.Name = &v
	return s
}

func (s *RunWritingV2RequestSearchSources) Validate() error {
	return dara.Validate(s)
}

type RunWritingV2RequestSummarization struct {
	// The name of the event.
	//
	// example:
	//
	// 事件名称
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// The summary of the event.
	//
	// example:
	//
	// 事件摘编
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RunWritingV2RequestSummarization) String() string {
	return dara.Prettify(s)
}

func (s RunWritingV2RequestSummarization) GoString() string {
	return s.String()
}

func (s *RunWritingV2RequestSummarization) GetEvent() *string {
	return s.Event
}

func (s *RunWritingV2RequestSummarization) GetMessage() *string {
	return s.Message
}

func (s *RunWritingV2RequestSummarization) SetEvent(v string) *RunWritingV2RequestSummarization {
	s.Event = &v
	return s
}

func (s *RunWritingV2RequestSummarization) SetMessage(v string) *RunWritingV2RequestSummarization {
	s.Message = &v
	return s
}

func (s *RunWritingV2RequestSummarization) Validate() error {
	return dara.Validate(s)
}
