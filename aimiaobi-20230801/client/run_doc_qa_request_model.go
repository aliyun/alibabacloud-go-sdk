// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDocQaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryIds(v []*string) *RunDocQaRequest
	GetCategoryIds() []*string
	SetConversationContexts(v []*RunDocQaRequestConversationContexts) *RunDocQaRequest
	GetConversationContexts() []*RunDocQaRequestConversationContexts
	SetDocIds(v []*string) *RunDocQaRequest
	GetDocIds() []*string
	SetModelName(v string) *RunDocQaRequest
	GetModelName() *string
	SetQuery(v string) *RunDocQaRequest
	GetQuery() *string
	SetReferenceContent(v string) *RunDocQaRequest
	GetReferenceContent() *string
	SetSearchSource(v string) *RunDocQaRequest
	GetSearchSource() *string
	SetSessionId(v string) *RunDocQaRequest
	GetSessionId() *string
	SetWorkspaceId(v string) *RunDocQaRequest
	GetWorkspaceId() *string
}

type RunDocQaRequest struct {
	CategoryIds          []*string                              `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty" type:"Repeated"`
	ConversationContexts []*RunDocQaRequestConversationContexts `json:"ConversationContexts,omitempty" xml:"ConversationContexts,omitempty" type:"Repeated"`
	DocIds               []*string                              `json:"DocIds,omitempty" xml:"DocIds,omitempty" type:"Repeated"`
	ModelName            *string                                `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
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

func (s RunDocQaRequest) String() string {
	return dara.Prettify(s)
}

func (s RunDocQaRequest) GoString() string {
	return s.String()
}

func (s *RunDocQaRequest) GetCategoryIds() []*string {
	return s.CategoryIds
}

func (s *RunDocQaRequest) GetConversationContexts() []*RunDocQaRequestConversationContexts {
	return s.ConversationContexts
}

func (s *RunDocQaRequest) GetDocIds() []*string {
	return s.DocIds
}

func (s *RunDocQaRequest) GetModelName() *string {
	return s.ModelName
}

func (s *RunDocQaRequest) GetQuery() *string {
	return s.Query
}

func (s *RunDocQaRequest) GetReferenceContent() *string {
	return s.ReferenceContent
}

func (s *RunDocQaRequest) GetSearchSource() *string {
	return s.SearchSource
}

func (s *RunDocQaRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RunDocQaRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunDocQaRequest) SetCategoryIds(v []*string) *RunDocQaRequest {
	s.CategoryIds = v
	return s
}

func (s *RunDocQaRequest) SetConversationContexts(v []*RunDocQaRequestConversationContexts) *RunDocQaRequest {
	s.ConversationContexts = v
	return s
}

func (s *RunDocQaRequest) SetDocIds(v []*string) *RunDocQaRequest {
	s.DocIds = v
	return s
}

func (s *RunDocQaRequest) SetModelName(v string) *RunDocQaRequest {
	s.ModelName = &v
	return s
}

func (s *RunDocQaRequest) SetQuery(v string) *RunDocQaRequest {
	s.Query = &v
	return s
}

func (s *RunDocQaRequest) SetReferenceContent(v string) *RunDocQaRequest {
	s.ReferenceContent = &v
	return s
}

func (s *RunDocQaRequest) SetSearchSource(v string) *RunDocQaRequest {
	s.SearchSource = &v
	return s
}

func (s *RunDocQaRequest) SetSessionId(v string) *RunDocQaRequest {
	s.SessionId = &v
	return s
}

func (s *RunDocQaRequest) SetWorkspaceId(v string) *RunDocQaRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunDocQaRequest) Validate() error {
	return dara.Validate(s)
}

type RunDocQaRequestConversationContexts struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// user
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s RunDocQaRequestConversationContexts) String() string {
	return dara.Prettify(s)
}

func (s RunDocQaRequestConversationContexts) GoString() string {
	return s.String()
}

func (s *RunDocQaRequestConversationContexts) GetContent() *string {
	return s.Content
}

func (s *RunDocQaRequestConversationContexts) GetRole() *string {
	return s.Role
}

func (s *RunDocQaRequestConversationContexts) SetContent(v string) *RunDocQaRequestConversationContexts {
	s.Content = &v
	return s
}

func (s *RunDocQaRequestConversationContexts) SetRole(v string) *RunDocQaRequestConversationContexts {
	s.Role = &v
	return s
}

func (s *RunDocQaRequestConversationContexts) Validate() error {
	return dara.Validate(s)
}
