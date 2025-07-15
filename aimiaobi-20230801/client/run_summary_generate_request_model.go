// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSummaryGenerateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *RunSummaryGenerateRequest
	GetContent() *string
	SetPrompt(v string) *RunSummaryGenerateRequest
	GetPrompt() *string
	SetWorkspaceId(v string) *RunSummaryGenerateRequest
	GetWorkspaceId() *string
}

type RunSummaryGenerateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 创新政务社交媒体功能。鼓励各地区、各部门结合实际，开发政务社交媒体的特色功能，如在线咨询服务、政策解读、互动问答等，增强政务社交媒体的互动性和实用性，提升公众参与度。
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 请为上述内容生成一段摘要，字数在100~200字以内。
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2setzb9x4ewsd
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunSummaryGenerateRequest) String() string {
	return dara.Prettify(s)
}

func (s RunSummaryGenerateRequest) GoString() string {
	return s.String()
}

func (s *RunSummaryGenerateRequest) GetContent() *string {
	return s.Content
}

func (s *RunSummaryGenerateRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunSummaryGenerateRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunSummaryGenerateRequest) SetContent(v string) *RunSummaryGenerateRequest {
	s.Content = &v
	return s
}

func (s *RunSummaryGenerateRequest) SetPrompt(v string) *RunSummaryGenerateRequest {
	s.Prompt = &v
	return s
}

func (s *RunSummaryGenerateRequest) SetWorkspaceId(v string) *RunSummaryGenerateRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunSummaryGenerateRequest) Validate() error {
	return dara.Validate(s)
}
