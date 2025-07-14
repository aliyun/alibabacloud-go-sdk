// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSharedReportsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListSharedReportsRequest
	GetKeyword() *string
	SetPageSize(v int32) *ListSharedReportsRequest
	GetPageSize() *int32
	SetTreeType(v string) *ListSharedReportsRequest
	GetTreeType() *string
	SetUserId(v string) *ListSharedReportsRequest
	GetUserId() *string
}

type ListSharedReportsRequest struct {
	// Keyword of the name of the work.
	//
	// example:
	//
	// Test report
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// Query the number of rows in the work list:
	//
	// 	- Default value: 10.
	//
	// 	- Maximum value: 9999
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Query the type of the work (fill in the blank to query all types). Valid values:
	//
	// 	- DATAPRODUCT: BI portal
	//
	// 	- PAGE: Dashboard
	//
	// 	- REPORT: workbook
	//
	// example:
	//
	// PAGE
	TreeType *string `json:"TreeType,omitempty" xml:"TreeType,omitempty"`
	// The UserID of the user to be queried in the Quick BI.
	//
	// This parameter is required.
	//
	// example:
	//
	// b5d8fd9348cc4327****afb604
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListSharedReportsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSharedReportsRequest) GoString() string {
	return s.String()
}

func (s *ListSharedReportsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListSharedReportsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSharedReportsRequest) GetTreeType() *string {
	return s.TreeType
}

func (s *ListSharedReportsRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListSharedReportsRequest) SetKeyword(v string) *ListSharedReportsRequest {
	s.Keyword = &v
	return s
}

func (s *ListSharedReportsRequest) SetPageSize(v int32) *ListSharedReportsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSharedReportsRequest) SetTreeType(v string) *ListSharedReportsRequest {
	s.TreeType = &v
	return s
}

func (s *ListSharedReportsRequest) SetUserId(v string) *ListSharedReportsRequest {
	s.UserId = &v
	return s
}

func (s *ListSharedReportsRequest) Validate() error {
	return dara.Validate(s)
}
