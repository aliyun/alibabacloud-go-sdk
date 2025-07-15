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
	SetWorkspaceId(v string) *RunTextPolishingRequest
	GetWorkspaceId() *string
}

type RunTextPolishingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 文本内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
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

func (s *RunTextPolishingRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunTextPolishingRequest) SetContent(v string) *RunTextPolishingRequest {
	s.Content = &v
	return s
}

func (s *RunTextPolishingRequest) SetWorkspaceId(v string) *RunTextPolishingRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunTextPolishingRequest) Validate() error {
	return dara.Validate(s)
}
