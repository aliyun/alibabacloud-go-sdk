// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFavoriteReportsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListFavoriteReportsRequest
	GetKeyword() *string
	SetPageSize(v int32) *ListFavoriteReportsRequest
	GetPageSize() *int32
	SetTreeType(v string) *ListFavoriteReportsRequest
	GetTreeType() *string
	SetUserId(v string) *ListFavoriteReportsRequest
	GetUserId() *string
}

type ListFavoriteReportsRequest struct {
	// Keyword of the work name.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// Number of rows in the work list to be queried:
	//
	// Default value: 10
	//
	// Maximum value: 9999
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Type of the work to be queried (leave blank to query all types). Value range:
	//
	// - DATAPRODUCT: Data Portal
	//
	// - PAGE: Dashboard
	//
	// - REPORT: Spreadsheet
	//
	// example:
	//
	// PAGE
	TreeType *string `json:"TreeType,omitempty" xml:"TreeType,omitempty"`
	// The UserID of the user in Quick BI to be queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// b5d8fd9348cc4327****afb604
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListFavoriteReportsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFavoriteReportsRequest) GoString() string {
	return s.String()
}

func (s *ListFavoriteReportsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListFavoriteReportsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFavoriteReportsRequest) GetTreeType() *string {
	return s.TreeType
}

func (s *ListFavoriteReportsRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListFavoriteReportsRequest) SetKeyword(v string) *ListFavoriteReportsRequest {
	s.Keyword = &v
	return s
}

func (s *ListFavoriteReportsRequest) SetPageSize(v int32) *ListFavoriteReportsRequest {
	s.PageSize = &v
	return s
}

func (s *ListFavoriteReportsRequest) SetTreeType(v string) *ListFavoriteReportsRequest {
	s.TreeType = &v
	return s
}

func (s *ListFavoriteReportsRequest) SetUserId(v string) *ListFavoriteReportsRequest {
	s.UserId = &v
	return s
}

func (s *ListFavoriteReportsRequest) Validate() error {
	return dara.Validate(s)
}
