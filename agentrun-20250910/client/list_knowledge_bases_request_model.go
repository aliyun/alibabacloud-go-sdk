// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListKnowledgeBasesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListKnowledgeBasesRequest
	GetPageSize() *int32
	SetProvider(v string) *ListKnowledgeBasesRequest
	GetProvider() *string
	SetWorkspaceId(v string) *ListKnowledgeBasesRequest
	GetWorkspaceId() *string
	SetWorkspaceIds(v string) *ListKnowledgeBasesRequest
	GetWorkspaceIds() *string
}

type ListKnowledgeBasesRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// bailian
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// example:
	//
	// 1e546db8-bd93-5a52-9be1-5a1351cdac95
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// example:
	//
	// 1e546db8-bd93-5a52-9be1-5a1351cdac95,1e546db8-bd93-5a52-9be1-5a1351caass4
	WorkspaceIds *string `json:"workspaceIds,omitempty" xml:"workspaceIds,omitempty"`
}

func (s ListKnowledgeBasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBasesRequest) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBasesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListKnowledgeBasesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListKnowledgeBasesRequest) GetProvider() *string {
	return s.Provider
}

func (s *ListKnowledgeBasesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListKnowledgeBasesRequest) GetWorkspaceIds() *string {
	return s.WorkspaceIds
}

func (s *ListKnowledgeBasesRequest) SetPageNumber(v int32) *ListKnowledgeBasesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListKnowledgeBasesRequest) SetPageSize(v int32) *ListKnowledgeBasesRequest {
	s.PageSize = &v
	return s
}

func (s *ListKnowledgeBasesRequest) SetProvider(v string) *ListKnowledgeBasesRequest {
	s.Provider = &v
	return s
}

func (s *ListKnowledgeBasesRequest) SetWorkspaceId(v string) *ListKnowledgeBasesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListKnowledgeBasesRequest) SetWorkspaceIds(v string) *ListKnowledgeBasesRequest {
	s.WorkspaceIds = &v
	return s
}

func (s *ListKnowledgeBasesRequest) Validate() error {
	return dara.Validate(s)
}
