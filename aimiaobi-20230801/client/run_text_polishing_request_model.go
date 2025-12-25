// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunTextPolishingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *RunTextPolishingRequest
	GetContent() *string
	SetOriginContent(v string) *RunTextPolishingRequest
	GetOriginContent() *string
	SetPrompt(v string) *RunTextPolishingRequest
	GetPrompt() *string
	SetWorkspaceId(v string) *RunTextPolishingRequest
	GetWorkspaceId() *string
}

type RunTextPolishingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 文本内容
	Content       *string `json:"Content,omitempty" xml:"Content,omitempty"`
	OriginContent *string `json:"OriginContent,omitempty" xml:"OriginContent,omitempty"`
	Prompt        *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunTextPolishingRequest) String() string {
	return dara.Prettify(s)
}

func (s RunTextPolishingRequest) GoString() string {
	return s.String()
}

func (s *RunTextPolishingRequest) GetContent() *string {
	return s.Content
}

func (s *RunTextPolishingRequest) GetOriginContent() *string {
	return s.OriginContent
}

func (s *RunTextPolishingRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunTextPolishingRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunTextPolishingRequest) SetContent(v string) *RunTextPolishingRequest {
	s.Content = &v
	return s
}

func (s *RunTextPolishingRequest) SetOriginContent(v string) *RunTextPolishingRequest {
	s.OriginContent = &v
	return s
}

func (s *RunTextPolishingRequest) SetPrompt(v string) *RunTextPolishingRequest {
	s.Prompt = &v
	return s
}

func (s *RunTextPolishingRequest) SetWorkspaceId(v string) *RunTextPolishingRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunTextPolishingRequest) Validate() error {
	return dara.Validate(s)
}
