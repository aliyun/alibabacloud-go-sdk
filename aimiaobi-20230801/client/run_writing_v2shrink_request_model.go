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
	ArticlesShrink *string `json:"Articles,omitempty" xml:"Articles,omitempty"`
	// example:
	//
	// false
	DistributeWriting *bool `json:"DistributeWriting,omitempty" xml:"DistributeWriting,omitempty"`
	// example:
	//
	// 2
	GcNumberSize    *int32  `json:"GcNumberSize,omitempty" xml:"GcNumberSize,omitempty"`
	GcNumberSizeTag *string `json:"GcNumberSizeTag,omitempty" xml:"GcNumberSizeTag,omitempty"`
	KeywordsShrink  *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// example:
	//
	// en
	Language       *string `json:"Language,omitempty" xml:"Language,omitempty"`
	MiniDocsShrink *string `json:"MiniDocs,omitempty" xml:"MiniDocs,omitempty"`
	OutlinesShrink *string `json:"Outlines,omitempty" xml:"Outlines,omitempty"`
	Prompt         *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// Template
	PromptMode          *string `json:"PromptMode,omitempty" xml:"PromptMode,omitempty"`
	SearchSourcesShrink *string `json:"SearchSources,omitempty" xml:"SearchSources,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// Writing
	Step                *string `json:"Step,omitempty" xml:"Step,omitempty"`
	SummarizationShrink *string `json:"Summarization,omitempty" xml:"Summarization,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// true
	UseSearch *bool `json:"UseSearch,omitempty" xml:"UseSearch,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId         *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WritingParamsShrink *string `json:"WritingParams,omitempty" xml:"WritingParams,omitempty"`
	// example:
	//
	// media
	WritingScene *string `json:"WritingScene,omitempty" xml:"WritingScene,omitempty"`
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
