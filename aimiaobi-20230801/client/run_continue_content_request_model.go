// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunContinueContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *RunContinueContentRequest
	GetContent() *string
	SetWorkspaceId(v string) *RunContinueContentRequest
	GetWorkspaceId() *string
}

type RunContinueContentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 创新政务社交媒体功能。鼓励各地区、各部门结合实际，开发政务社交媒体的特色功能，如在线咨询服务、政策解读、互动问答等，增强政务社交媒体的互动性和实用性，提升公众参与度。
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2setzb9x4ewsd
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunContinueContentRequest) String() string {
	return dara.Prettify(s)
}

func (s RunContinueContentRequest) GoString() string {
	return s.String()
}

func (s *RunContinueContentRequest) GetContent() *string {
	return s.Content
}

func (s *RunContinueContentRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunContinueContentRequest) SetContent(v string) *RunContinueContentRequest {
	s.Content = &v
	return s
}

func (s *RunContinueContentRequest) SetWorkspaceId(v string) *RunContinueContentRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunContinueContentRequest) Validate() error {
	return dara.Validate(s)
}
