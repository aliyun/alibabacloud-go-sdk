// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *ListKnowledgeBaseRequest
	GetKnowledgeBaseId() *string
	SetPageNumber(v int32) *ListKnowledgeBaseRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListKnowledgeBaseRequest
	GetPageSize() *int32
}

type ListKnowledgeBaseRequest struct {
	// example:
	//
	// "186432649"
	KnowledgeBaseId *string `json:"knowledgeBaseId,omitempty" xml:"knowledgeBaseId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *ListKnowledgeBaseRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListKnowledgeBaseRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListKnowledgeBaseRequest) SetKnowledgeBaseId(v string) *ListKnowledgeBaseRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *ListKnowledgeBaseRequest) SetPageNumber(v int32) *ListKnowledgeBaseRequest {
	s.PageNumber = &v
	return s
}

func (s *ListKnowledgeBaseRequest) SetPageSize(v int32) *ListKnowledgeBaseRequest {
	s.PageSize = &v
	return s
}

func (s *ListKnowledgeBaseRequest) Validate() error {
	return dara.Validate(s)
}
