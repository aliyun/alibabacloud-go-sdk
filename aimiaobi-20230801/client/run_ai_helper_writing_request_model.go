// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunAiHelperWritingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDistributeWriting(v bool) *RunAiHelperWritingRequest
	GetDistributeWriting() *bool
	SetPrompt(v string) *RunAiHelperWritingRequest
	GetPrompt() *string
	SetPromptMode(v string) *RunAiHelperWritingRequest
	GetPromptMode() *string
	SetWorkspaceId(v string) *RunAiHelperWritingRequest
	GetWorkspaceId() *string
	SetWritingParams(v map[string]*string) *RunAiHelperWritingRequest
	GetWritingParams() map[string]*string
	SetWritingScene(v string) *RunAiHelperWritingRequest
	GetWritingScene() *string
	SetWritingStyle(v string) *RunAiHelperWritingRequest
	GetWritingStyle() *string
}

type RunAiHelperWritingRequest struct {
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
	WritingParams map[string]*string `json:"WritingParams,omitempty" xml:"WritingParams,omitempty"`
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

func (s RunAiHelperWritingRequest) String() string {
	return dara.Prettify(s)
}

func (s RunAiHelperWritingRequest) GoString() string {
	return s.String()
}

func (s *RunAiHelperWritingRequest) GetDistributeWriting() *bool {
	return s.DistributeWriting
}

func (s *RunAiHelperWritingRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunAiHelperWritingRequest) GetPromptMode() *string {
	return s.PromptMode
}

func (s *RunAiHelperWritingRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunAiHelperWritingRequest) GetWritingParams() map[string]*string {
	return s.WritingParams
}

func (s *RunAiHelperWritingRequest) GetWritingScene() *string {
	return s.WritingScene
}

func (s *RunAiHelperWritingRequest) GetWritingStyle() *string {
	return s.WritingStyle
}

func (s *RunAiHelperWritingRequest) SetDistributeWriting(v bool) *RunAiHelperWritingRequest {
	s.DistributeWriting = &v
	return s
}

func (s *RunAiHelperWritingRequest) SetPrompt(v string) *RunAiHelperWritingRequest {
	s.Prompt = &v
	return s
}

func (s *RunAiHelperWritingRequest) SetPromptMode(v string) *RunAiHelperWritingRequest {
	s.PromptMode = &v
	return s
}

func (s *RunAiHelperWritingRequest) SetWorkspaceId(v string) *RunAiHelperWritingRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunAiHelperWritingRequest) SetWritingParams(v map[string]*string) *RunAiHelperWritingRequest {
	s.WritingParams = v
	return s
}

func (s *RunAiHelperWritingRequest) SetWritingScene(v string) *RunAiHelperWritingRequest {
	s.WritingScene = &v
	return s
}

func (s *RunAiHelperWritingRequest) SetWritingStyle(v string) *RunAiHelperWritingRequest {
	s.WritingStyle = &v
	return s
}

func (s *RunAiHelperWritingRequest) Validate() error {
	return dara.Validate(s)
}
