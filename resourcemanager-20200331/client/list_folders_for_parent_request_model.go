// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFoldersForParentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListFoldersForParentRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListFoldersForParentRequest
	GetPageSize() *int32
	SetParentFolderId(v string) *ListFoldersForParentRequest
	GetParentFolderId() *string
	SetQueryKeyword(v string) *ListFoldersForParentRequest
	GetQueryKeyword() *string
}

type ListFoldersForParentRequest struct {
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the parent folder.
	//
	// If you leave this parameter empty, the information of the first-level subfolders of the Root folder is queried.
	//
	// example:
	//
	// r-b1****
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	// The keyword used for the query, such as a folder name.
	//
	// Fuzzy match is supported.
	//
	// example:
	//
	// rdFolder
	QueryKeyword *string `json:"QueryKeyword,omitempty" xml:"QueryKeyword,omitempty"`
}

func (s ListFoldersForParentRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFoldersForParentRequest) GoString() string {
	return s.String()
}

func (s *ListFoldersForParentRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFoldersForParentRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFoldersForParentRequest) GetParentFolderId() *string {
	return s.ParentFolderId
}

func (s *ListFoldersForParentRequest) GetQueryKeyword() *string {
	return s.QueryKeyword
}

func (s *ListFoldersForParentRequest) SetPageNumber(v int32) *ListFoldersForParentRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFoldersForParentRequest) SetPageSize(v int32) *ListFoldersForParentRequest {
	s.PageSize = &v
	return s
}

func (s *ListFoldersForParentRequest) SetParentFolderId(v string) *ListFoldersForParentRequest {
	s.ParentFolderId = &v
	return s
}

func (s *ListFoldersForParentRequest) SetQueryKeyword(v string) *ListFoldersForParentRequest {
	s.QueryKeyword = &v
	return s
}

func (s *ListFoldersForParentRequest) Validate() error {
	return dara.Validate(s)
}
