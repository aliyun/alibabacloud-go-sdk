// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMigrationTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListMigrationTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMigrationTasksRequest
	GetPageSize() *int32
}

type ListMigrationTasksRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListMigrationTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationTasksRequest) GoString() string {
	return s.String()
}

func (s *ListMigrationTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMigrationTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMigrationTasksRequest) SetPageNumber(v int32) *ListMigrationTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMigrationTasksRequest) SetPageSize(v int32) *ListMigrationTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListMigrationTasksRequest) Validate() error {
	return dara.Validate(s)
}
