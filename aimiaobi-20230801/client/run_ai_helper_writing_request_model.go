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
	// example:
	//
	// false
	DistributeWriting *bool `json:"DistributeWriting,omitempty" xml:"DistributeWriting,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 请帮我写一篇关于人工智能发展趋势的文章
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// Template
	PromptMode *string `json:"PromptMode,omitempty" xml:"PromptMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// example:
	//
	// {"wordCount": "1000", "tone": "formal"}
	WritingParams map[string]*string `json:"WritingParams,omitempty" xml:"WritingParams,omitempty"`
	// 写作场景：government(政务)、media(传媒)、market(营销)、office(办公)、custom(自定义)
	//
	// This parameter is required.
	//
	// example:
	//
	// media
	WritingScene *string `json:"WritingScene,omitempty" xml:"WritingScene,omitempty"`
	// 写作文体唯一标识KEY，可通过ListWritingStyles接口获取对应写作场景下的文体列表
	//
	// This parameter is required.
	//
	// example:
	//
	// news_article
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
