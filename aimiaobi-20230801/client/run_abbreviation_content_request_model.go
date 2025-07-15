// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunAbbreviationContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *RunAbbreviationContentRequest
	GetContent() *string
	SetPrompt(v string) *RunAbbreviationContentRequest
	GetPrompt() *string
	SetWorkspaceId(v string) *RunAbbreviationContentRequest
	GetWorkspaceId() *string
}

type RunAbbreviationContentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 创新政务社交媒体功能。鼓励各地区、各部门结合实际，开发政务社交媒体的特色功能，如在线咨询服务、政策解读、互动问答等，增强政务社交媒体的互动性和实用性，提升公众参与度。
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Prompt  *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2setzb9x4ewsd
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunAbbreviationContentRequest) String() string {
	return dara.Prettify(s)
}

func (s RunAbbreviationContentRequest) GoString() string {
	return s.String()
}

func (s *RunAbbreviationContentRequest) GetContent() *string {
	return s.Content
}

func (s *RunAbbreviationContentRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunAbbreviationContentRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunAbbreviationContentRequest) SetContent(v string) *RunAbbreviationContentRequest {
	s.Content = &v
	return s
}

func (s *RunAbbreviationContentRequest) SetPrompt(v string) *RunAbbreviationContentRequest {
	s.Prompt = &v
	return s
}

func (s *RunAbbreviationContentRequest) SetWorkspaceId(v string) *RunAbbreviationContentRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunAbbreviationContentRequest) Validate() error {
	return dara.Validate(s)
}
