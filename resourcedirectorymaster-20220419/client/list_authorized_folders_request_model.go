// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedFoldersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListAuthorizedFoldersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAuthorizedFoldersRequest
	GetPageSize() *int32
}

type ListAuthorizedFoldersRequest struct {
	// The page number.
	//
	// Page starts from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAuthorizedFoldersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedFoldersRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizedFoldersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAuthorizedFoldersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAuthorizedFoldersRequest) SetPageNumber(v int32) *ListAuthorizedFoldersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAuthorizedFoldersRequest) SetPageSize(v int32) *ListAuthorizedFoldersRequest {
	s.PageSize = &v
	return s
}

func (s *ListAuthorizedFoldersRequest) Validate() error {
	return dara.Validate(s)
}
