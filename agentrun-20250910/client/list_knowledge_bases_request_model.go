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

func (s *ListKnowledgeBasesRequest) Validate() error {
	return dara.Validate(s)
}
