// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemoryCollectionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemoryCollectionName(v string) *ListMemoryCollectionsRequest
	GetMemoryCollectionName() *string
	SetPageNumber(v int32) *ListMemoryCollectionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMemoryCollectionsRequest
	GetPageSize() *int32
	SetStatus(v string) *ListMemoryCollectionsRequest
	GetStatus() *string
	SetType(v string) *ListMemoryCollectionsRequest
	GetType() *string
}

type ListMemoryCollectionsRequest struct {
	// example:
	//
	// my-memory-1
	MemoryCollectionName *string `json:"memoryCollectionName,omitempty" xml:"memoryCollectionName,omitempty"`
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
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// mem0
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListMemoryCollectionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryCollectionsRequest) GoString() string {
	return s.String()
}

func (s *ListMemoryCollectionsRequest) GetMemoryCollectionName() *string {
	return s.MemoryCollectionName
}

func (s *ListMemoryCollectionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMemoryCollectionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMemoryCollectionsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListMemoryCollectionsRequest) GetType() *string {
	return s.Type
}

func (s *ListMemoryCollectionsRequest) SetMemoryCollectionName(v string) *ListMemoryCollectionsRequest {
	s.MemoryCollectionName = &v
	return s
}

func (s *ListMemoryCollectionsRequest) SetPageNumber(v int32) *ListMemoryCollectionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMemoryCollectionsRequest) SetPageSize(v int32) *ListMemoryCollectionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMemoryCollectionsRequest) SetStatus(v string) *ListMemoryCollectionsRequest {
	s.Status = &v
	return s
}

func (s *ListMemoryCollectionsRequest) SetType(v string) *ListMemoryCollectionsRequest {
	s.Type = &v
	return s
}

func (s *ListMemoryCollectionsRequest) Validate() error {
	return dara.Validate(s)
}
