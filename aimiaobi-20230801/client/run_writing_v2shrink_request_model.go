// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunWritingV2ShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArticlesShrink(v string) *RunWritingV2ShrinkRequest
	GetArticlesShrink() *string
	SetDistributeWriting(v bool) *RunWritingV2ShrinkRequest
	GetDistributeWriting() *bool
	SetGcNumberSize(v int32) *RunWritingV2ShrinkRequest
	GetGcNumberSize() *int32
	SetGcNumberSizeTag(v string) *RunWritingV2ShrinkRequest
	GetGcNumberSizeTag() *string
	SetKeywordsShrink(v string) *RunWritingV2ShrinkRequest
	GetKeywordsShrink() *string
	SetLanguage(v string) *RunWritingV2ShrinkRequest
	GetLanguage() *string
	SetMiniDocsShrink(v string) *RunWritingV2ShrinkRequest
	GetMiniDocsShrink() *string
	SetOutlineListShrink(v string) *RunWritingV2ShrinkRequest
	GetOutlineListShrink() *string
	SetOutlinesShrink(v string) *RunWritingV2ShrinkRequest
	GetOutlinesShrink() *string
	SetPrompt(v string) *RunWritingV2ShrinkRequest
	GetPrompt() *string
	SetPromptMode(v string) *RunWritingV2ShrinkRequest
	GetPromptMode() *string
	SetSearchSourcesShrink(v string) *RunWritingV2ShrinkRequest
	GetSearchSourcesShrink() *string
	SetSessionId(v string) *RunWritingV2ShrinkRequest
	GetSessionId() *string
	SetSourceTraceMethod(v string) *RunWritingV2ShrinkRequest
	GetSourceTraceMethod() *string
	SetStep(v string) *RunWritingV2ShrinkRequest
	GetStep() *string
	SetSummarizationShrink(v string) *RunWritingV2ShrinkRequest
	GetSummarizationShrink() *string
	SetTaskId(v string) *RunWritingV2ShrinkRequest
	GetTaskId() *string
	SetUseSearch(v bool) *RunWritingV2ShrinkRequest
	GetUseSearch() *bool
	SetWorkspaceId(v string) *RunWritingV2ShrinkRequest
	GetWorkspaceId() *string
	SetWritingParamsShrink(v string) *RunWritingV2ShrinkRequest
	GetWritingParamsShrink() *string
	SetWritingScene(v string) *RunWritingV2ShrinkRequest
	GetWritingScene() *string
	SetWritingStyle(v string) *RunWritingV2ShrinkRequest
	GetWritingStyle() *string
}

type RunWritingV2ShrinkRequest struct {
	// A list of articles to use as references. **Note:*	- When you provide this parameter, web search is disabled, overriding the `UseSearch` and `SearchSources` parameters.
	ArticlesShrink *string `json:"Articles,omitempty" xml:"Articles,omitempty"`
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
	KeywordsShrink *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
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
	MiniDocsShrink *string `json:"MiniDocs,omitempty" xml:"MiniDocs,omitempty"`
	// A list of outlines for step-by-step writing.
	OutlineListShrink *string `json:"OutlineList,omitempty" xml:"OutlineList,omitempty"`
	// A list of outlines for step-by-step writing. This parameter is deprecated. Use `OutlineList` instead.
	OutlinesShrink *string `json:"Outlines,omitempty" xml:"Outlines,omitempty"`
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
	SearchSourcesShrink *string `json:"SearchSources,omitempty" xml:"SearchSources,omitempty"`
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
	SummarizationShrink *string `json:"Summarization,omitempty" xml:"Summarization,omitempty"`
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
	WritingParamsShrink *string `json:"WritingParams,omitempty" xml:"WritingParams,omitempty"`
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

func (s RunWritingV2ShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunWritingV2ShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunWritingV2ShrinkRequest) GetArticlesShrink() *string {
	return s.ArticlesShrink
}

func (s *RunWritingV2ShrinkRequest) GetDistributeWriting() *bool {
	return s.DistributeWriting
}

func (s *RunWritingV2ShrinkRequest) GetGcNumberSize() *int32 {
	return s.GcNumberSize
}

func (s *RunWritingV2ShrinkRequest) GetGcNumberSizeTag() *string {
	return s.GcNumberSizeTag
}

func (s *RunWritingV2ShrinkRequest) GetKeywordsShrink() *string {
	return s.KeywordsShrink
}

func (s *RunWritingV2ShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *RunWritingV2ShrinkRequest) GetMiniDocsShrink() *string {
	return s.MiniDocsShrink
}

func (s *RunWritingV2ShrinkRequest) GetOutlineListShrink() *string {
	return s.OutlineListShrink
}

func (s *RunWritingV2ShrinkRequest) GetOutlinesShrink() *string {
	return s.OutlinesShrink
}

func (s *RunWritingV2ShrinkRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunWritingV2ShrinkRequest) GetPromptMode() *string {
	return s.PromptMode
}

func (s *RunWritingV2ShrinkRequest) GetSearchSourcesShrink() *string {
	return s.SearchSourcesShrink
}

func (s *RunWritingV2ShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RunWritingV2ShrinkRequest) GetSourceTraceMethod() *string {
	return s.SourceTraceMethod
}

func (s *RunWritingV2ShrinkRequest) GetStep() *string {
	return s.Step
}

func (s *RunWritingV2ShrinkRequest) GetSummarizationShrink() *string {
	return s.SummarizationShrink
}

func (s *RunWritingV2ShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RunWritingV2ShrinkRequest) GetUseSearch() *bool {
	return s.UseSearch
}

func (s *RunWritingV2ShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunWritingV2ShrinkRequest) GetWritingParamsShrink() *string {
	return s.WritingParamsShrink
}

func (s *RunWritingV2ShrinkRequest) GetWritingScene() *string {
	return s.WritingScene
}

func (s *RunWritingV2ShrinkRequest) GetWritingStyle() *string {
	return s.WritingStyle
}

func (s *RunWritingV2ShrinkRequest) SetArticlesShrink(v string) *RunWritingV2ShrinkRequest {
	s.ArticlesShrink = &v
	return s
}

func (s *RunWritingV2ShrinkRequest) SetDistributeWriting(v bool) *RunWritingV2ShrinkRequest {
	s.DistributeWriting = &v
	return s
}

func (s *RunWritingV2ShrinkRequest) SetGcNumberSize(v int32) *RunWritingV2ShrinkRequest {
	s.GcNumberSize = &v
	return s
}

func (s *RunWritingV2ShrinkRequest) SetGcNumberSizeTag(v string) *RunWritingV2ShrinkRequest {
	s.GcNumberSizeTag = &v
	return s
}

func (s *RunWritingV2ShrinkRequest) SetKeywordsShrink(v string) *RunWritingV2ShrinkRequest {
	s.KeywordsShrink = &v
	return s
}

func (s *RunWritingV2ShrinkRequest) SetLanguage(v string) *RunWritingV2ShrinkRequest {
	s.Language = &v
	return s
}

func (s *RunWritingV2ShrinkRequest) SetMiniDocsShrink(v string) *RunWritingV2ShrinkRequest {
	s.MiniDocsShrink = &v
	return s
}

func (s *RunWritingV2ShrinkRequest) SetOutlineListShrink(v string) *RunWritingV2ShrinkRequest {
	s.OutlineListShrink = &v
	return s
}

func (s *RunWritingV2ShrinkRequest) SetOutlinesShrink(v string) *RunWritingV2ShrinkRequest {
	s.OutlinesShrink = &v
	return s
}

func (s *RunWritingV2ShrinkRequest) SetPrompt(v string) *RunWritingV2ShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *RunWritingV2ShrinkRequest) SetPromptMode(v string) *RunWritingV2ShrinkRequest {
	s.PromptMode = &v
	return s
}

func (s *RunWritingV2ShrinkRequest) SetSearchSourcesShrink(v string) *RunWritingV2ShrinkRequest {
	s.SearchSourcesShrink = &v
	return s
}

func (s *RunWritingV2ShrinkRequest) SetSessionId(v string) *RunWritingV2ShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *RunWritingV2ShrinkRequest) SetSourceTraceMethod(v string) *RunWritingV2ShrinkRequest {
	s.SourceTraceMethod = &v
	return s
}

func (s *RunWritingV2ShrinkRequest) SetStep(v string) *RunWritingV2ShrinkRequest {
	s.Step = &v
	return s
}

func (s *RunWritingV2ShrinkRequest) SetSummarizationShrink(v string) *RunWritingV2ShrinkRequest {
	s.SummarizationShrink = &v
	return s
}

func (s *RunWritingV2ShrinkRequest) SetTaskId(v string) *RunWritingV2ShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *RunWritingV2ShrinkRequest) SetUseSearch(v bool) *RunWritingV2ShrinkRequest {
	s.UseSearch = &v
	return s
}

func (s *RunWritingV2ShrinkRequest) SetWorkspaceId(v string) *RunWritingV2ShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunWritingV2ShrinkRequest) SetWritingParamsShrink(v string) *RunWritingV2ShrinkRequest {
	s.WritingParamsShrink = &v
	return s
}

func (s *RunWritingV2ShrinkRequest) SetWritingScene(v string) *RunWritingV2ShrinkRequest {
	s.WritingScene = &v
	return s
}

func (s *RunWritingV2ShrinkRequest) SetWritingStyle(v string) *RunWritingV2ShrinkRequest {
	s.WritingStyle = &v
	return s
}

func (s *RunWritingV2ShrinkRequest) Validate() error {
	return dara.Validate(s)
}
