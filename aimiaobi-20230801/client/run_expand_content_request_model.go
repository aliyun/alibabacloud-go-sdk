// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunExpandContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *RunExpandContentRequest
	GetContent() *string
	SetPrompt(v string) *RunExpandContentRequest
	GetPrompt() *string
	SetWorkspaceId(v string) *RunExpandContentRequest
	GetWorkspaceId() *string
}

type RunExpandContentRequest struct {
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

func (s RunExpandContentRequest) String() string {
	return dara.Prettify(s)
}

func (s RunExpandContentRequest) GoString() string {
	return s.String()
}

func (s *RunExpandContentRequest) GetContent() *string {
	return s.Content
}

func (s *RunExpandContentRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunExpandContentRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunExpandContentRequest) SetContent(v string) *RunExpandContentRequest {
	s.Content = &v
	return s
}

func (s *RunExpandContentRequest) SetPrompt(v string) *RunExpandContentRequest {
	s.Prompt = &v
	return s
}

func (s *RunExpandContentRequest) SetWorkspaceId(v string) *RunExpandContentRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunExpandContentRequest) Validate() error {
	return dara.Validate(s)
}
