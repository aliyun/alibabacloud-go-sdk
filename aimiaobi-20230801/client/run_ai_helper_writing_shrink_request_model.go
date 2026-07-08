// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunAiHelperWritingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDistributeWriting(v bool) *RunAiHelperWritingShrinkRequest
	GetDistributeWriting() *bool
	SetPrompt(v string) *RunAiHelperWritingShrinkRequest
	GetPrompt() *string
	SetPromptMode(v string) *RunAiHelperWritingShrinkRequest
	GetPromptMode() *string
	SetWorkspaceId(v string) *RunAiHelperWritingShrinkRequest
	GetWorkspaceId() *string
	SetWritingParamsShrink(v string) *RunAiHelperWritingShrinkRequest
	GetWritingParamsShrink() *string
	SetWritingScene(v string) *RunAiHelperWritingShrinkRequest
	GetWritingScene() *string
	SetWritingStyle(v string) *RunAiHelperWritingShrinkRequest
	GetWritingStyle() *string
}

type RunAiHelperWritingShrinkRequest struct {
	// Specifies whether to generate the text step by step.
	//
	// example:
	//
	// false
	DistributeWriting *bool `json:"DistributeWriting,omitempty" xml:"DistributeWriting,omitempty"`
	// The prompt, which specifies the subject for the AI to write about.
	//
	// This parameter is required.
	//
	// example:
	//
	// 请帮我写一篇关于人工智能发展趋势的文章
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// The prompt pattern. For example, PE indicates the advanced pattern and Template indicates the template pattern.
	//
	// example:
	//
	// Template
	PromptMode *string `json:"PromptMode,omitempty" xml:"PromptMode,omitempty"`
	// The [workspace](https://help.aliyun.com/document_detail/2782167.html) ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// llm-xxxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The writing parameters from the previous form, specified as key-value pairs.
	//
	// example:
	//
	// {"wordCount": "1000", "tone": "formal"}
	WritingParamsShrink *string `json:"WritingParams,omitempty" xml:"WritingParams,omitempty"`
	// The writing scenario. Valid values: government, media, market, office, and custom.
	//
	// This parameter is required.
	//
	// example:
	//
	// media
	WritingScene *string `json:"WritingScene,omitempty" xml:"WritingScene,omitempty"`
	// The unique key for the writing style. Call the [ListWritingStyles](https://help.aliyun.com/document_detail/2922609.html) operation to get a list of styles for the specified scenario.
	//
	// This parameter is required.
	//
	// example:
	//
	// 通知
	WritingStyle *string `json:"WritingStyle,omitempty" xml:"WritingStyle,omitempty"`
}

func (s RunAiHelperWritingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunAiHelperWritingShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunAiHelperWritingShrinkRequest) GetDistributeWriting() *bool {
	return s.DistributeWriting
}

func (s *RunAiHelperWritingShrinkRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunAiHelperWritingShrinkRequest) GetPromptMode() *string {
	return s.PromptMode
}

func (s *RunAiHelperWritingShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunAiHelperWritingShrinkRequest) GetWritingParamsShrink() *string {
	return s.WritingParamsShrink
}

func (s *RunAiHelperWritingShrinkRequest) GetWritingScene() *string {
	return s.WritingScene
}

func (s *RunAiHelperWritingShrinkRequest) GetWritingStyle() *string {
	return s.WritingStyle
}

func (s *RunAiHelperWritingShrinkRequest) SetDistributeWriting(v bool) *RunAiHelperWritingShrinkRequest {
	s.DistributeWriting = &v
	return s
}

func (s *RunAiHelperWritingShrinkRequest) SetPrompt(v string) *RunAiHelperWritingShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *RunAiHelperWritingShrinkRequest) SetPromptMode(v string) *RunAiHelperWritingShrinkRequest {
	s.PromptMode = &v
	return s
}

func (s *RunAiHelperWritingShrinkRequest) SetWorkspaceId(v string) *RunAiHelperWritingShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunAiHelperWritingShrinkRequest) SetWritingParamsShrink(v string) *RunAiHelperWritingShrinkRequest {
	s.WritingParamsShrink = &v
	return s
}

func (s *RunAiHelperWritingShrinkRequest) SetWritingScene(v string) *RunAiHelperWritingShrinkRequest {
	s.WritingScene = &v
	return s
}

func (s *RunAiHelperWritingShrinkRequest) SetWritingStyle(v string) *RunAiHelperWritingShrinkRequest {
	s.WritingStyle = &v
	return s
}

func (s *RunAiHelperWritingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
