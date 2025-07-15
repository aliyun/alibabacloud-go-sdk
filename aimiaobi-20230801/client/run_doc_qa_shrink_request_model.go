// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDocQaShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryIdsShrink(v string) *RunDocQaShrinkRequest
	GetCategoryIdsShrink() *string
	SetConversationContextsShrink(v string) *RunDocQaShrinkRequest
	GetConversationContextsShrink() *string
	SetDocIdsShrink(v string) *RunDocQaShrinkRequest
	GetDocIdsShrink() *string
	SetModelName(v string) *RunDocQaShrinkRequest
	GetModelName() *string
	SetQuery(v string) *RunDocQaShrinkRequest
	GetQuery() *string
	SetReferenceContent(v string) *RunDocQaShrinkRequest
	GetReferenceContent() *string
	SetSearchSource(v string) *RunDocQaShrinkRequest
	GetSearchSource() *string
	SetSessionId(v string) *RunDocQaShrinkRequest
	GetSessionId() *string
	SetWorkspaceId(v string) *RunDocQaShrinkRequest
	GetWorkspaceId() *string
}

type RunDocQaShrinkRequest struct {
	CategoryIdsShrink          *string `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty"`
	ConversationContextsShrink *string `json:"ConversationContexts,omitempty" xml:"ConversationContexts,omitempty"`
	DocIdsShrink               *string `json:"DocIds,omitempty" xml:"DocIds,omitempty"`
	ModelName                  *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// This parameter is required.
	Query            *string `json:"Query,omitempty" xml:"Query,omitempty"`
	ReferenceContent *string `json:"ReferenceContent,omitempty" xml:"ReferenceContent,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// fromWeb
	SearchSource *string `json:"SearchSource,omitempty" xml:"SearchSource,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f486c4e2-b773-4d65-88f8-2ba540610456
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-dswd4003ny4gh9rw
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunDocQaShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunDocQaShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunDocQaShrinkRequest) GetCategoryIdsShrink() *string {
	return s.CategoryIdsShrink
}

func (s *RunDocQaShrinkRequest) GetConversationContextsShrink() *string {
	return s.ConversationContextsShrink
}

func (s *RunDocQaShrinkRequest) GetDocIdsShrink() *string {
	return s.DocIdsShrink
}

func (s *RunDocQaShrinkRequest) GetModelName() *string {
	return s.ModelName
}

func (s *RunDocQaShrinkRequest) GetQuery() *string {
	return s.Query
}

func (s *RunDocQaShrinkRequest) GetReferenceContent() *string {
	return s.ReferenceContent
}

func (s *RunDocQaShrinkRequest) GetSearchSource() *string {
	return s.SearchSource
}

func (s *RunDocQaShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RunDocQaShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunDocQaShrinkRequest) SetCategoryIdsShrink(v string) *RunDocQaShrinkRequest {
	s.CategoryIdsShrink = &v
	return s
}

func (s *RunDocQaShrinkRequest) SetConversationContextsShrink(v string) *RunDocQaShrinkRequest {
	s.ConversationContextsShrink = &v
	return s
}

func (s *RunDocQaShrinkRequest) SetDocIdsShrink(v string) *RunDocQaShrinkRequest {
	s.DocIdsShrink = &v
	return s
}

func (s *RunDocQaShrinkRequest) SetModelName(v string) *RunDocQaShrinkRequest {
	s.ModelName = &v
	return s
}

func (s *RunDocQaShrinkRequest) SetQuery(v string) *RunDocQaShrinkRequest {
	s.Query = &v
	return s
}

func (s *RunDocQaShrinkRequest) SetReferenceContent(v string) *RunDocQaShrinkRequest {
	s.ReferenceContent = &v
	return s
}

func (s *RunDocQaShrinkRequest) SetSearchSource(v string) *RunDocQaShrinkRequest {
	s.SearchSource = &v
	return s
}

func (s *RunDocQaShrinkRequest) SetSessionId(v string) *RunDocQaShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *RunDocQaShrinkRequest) SetWorkspaceId(v string) *RunDocQaShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunDocQaShrinkRequest) Validate() error {
	return dara.Validate(s)
}
