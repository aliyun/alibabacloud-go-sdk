// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListUsersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUsersRequest
	GetPageSize() *int32
}

type ListUsersRequest struct {
	// The number of annotated members displayed per page in a paged query. Default value: 20.
	//
	// example:
	//
	// 20
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page number of the annotate member list. The starting value is 1. The default value is 1.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUsersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUsersRequest) SetPageNumber(v int32) *ListUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUsersRequest) SetPageSize(v int32) *ListUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListUsersRequest) Validate() error {
	return dara.Validate(s)
}
