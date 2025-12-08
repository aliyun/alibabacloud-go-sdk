// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllEndEntityInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListAllEndEntityInstanceRequest
	GetCurrentPage() *int32
	SetMaxResults(v int32) *ListAllEndEntityInstanceRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAllEndEntityInstanceRequest
	GetNextToken() *string
	SetParentId(v int64) *ListAllEndEntityInstanceRequest
	GetParentId() *int64
	SetRecursiveChildren(v int32) *ListAllEndEntityInstanceRequest
	GetRecursiveChildren() *int32
	SetShowSize(v int32) *ListAllEndEntityInstanceRequest
	GetShowSize() *int32
}

type ListAllEndEntityInstanceRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 1d2db86sca4384811e0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 37633
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// example:
	//
	// 9
	RecursiveChildren *int32 `json:"RecursiveChildren,omitempty" xml:"RecursiveChildren,omitempty"`
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
}

func (s ListAllEndEntityInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAllEndEntityInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListAllEndEntityInstanceRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAllEndEntityInstanceRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAllEndEntityInstanceRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAllEndEntityInstanceRequest) GetParentId() *int64 {
	return s.ParentId
}

func (s *ListAllEndEntityInstanceRequest) GetRecursiveChildren() *int32 {
	return s.RecursiveChildren
}

func (s *ListAllEndEntityInstanceRequest) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *ListAllEndEntityInstanceRequest) SetCurrentPage(v int32) *ListAllEndEntityInstanceRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListAllEndEntityInstanceRequest) SetMaxResults(v int32) *ListAllEndEntityInstanceRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAllEndEntityInstanceRequest) SetNextToken(v string) *ListAllEndEntityInstanceRequest {
	s.NextToken = &v
	return s
}

func (s *ListAllEndEntityInstanceRequest) SetParentId(v int64) *ListAllEndEntityInstanceRequest {
	s.ParentId = &v
	return s
}

func (s *ListAllEndEntityInstanceRequest) SetRecursiveChildren(v int32) *ListAllEndEntityInstanceRequest {
	s.RecursiveChildren = &v
	return s
}

func (s *ListAllEndEntityInstanceRequest) SetShowSize(v int32) *ListAllEndEntityInstanceRequest {
	s.ShowSize = &v
	return s
}

func (s *ListAllEndEntityInstanceRequest) Validate() error {
	return dara.Validate(s)
}
