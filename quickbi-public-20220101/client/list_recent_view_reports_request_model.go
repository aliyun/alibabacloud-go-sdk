// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecentViewReportsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListRecentViewReportsRequest
	GetKeyword() *string
	SetOffsetDay(v int32) *ListRecentViewReportsRequest
	GetOffsetDay() *int32
	SetPageSize(v int32) *ListRecentViewReportsRequest
	GetPageSize() *int32
	SetQueryMode(v string) *ListRecentViewReportsRequest
	GetQueryMode() *string
	SetTreeType(v string) *ListRecentViewReportsRequest
	GetTreeType() *string
	SetUserId(v string) *ListRecentViewReportsRequest
	GetUserId() *string
}

type ListRecentViewReportsRequest struct {
	// The keyword in the name of the work.
	//
	// example:
	//
	// FinancialReport
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The number of recent days to query data for. Default value: 10.
	//
	// example:
	//
	// 10
	OffsetDay *int32 `json:"OffsetDay,omitempty" xml:"OffsetDay,omitempty"`
	// The number of rows in the work list to query:
	//
	// - Default value: 10.
	//
	// - Maximum value: 9999.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query mode. Valid values:
	//
	// - 1: Sorted by number of visits.
	//
	// - 2: Sorted by most recent visit time.
	//
	// example:
	//
	// 1
	QueryMode *string `json:"QueryMode,omitempty" xml:"QueryMode,omitempty"`
	// The type of work to query. Leave this parameter empty to query all types. Valid values:
	//
	// - DATAPRODUCT: data portal
	//
	// - PAGE: dashboard
	//
	// - REPORT: workbook
	//
	// example:
	//
	// PAGE
	TreeType *string `json:"TreeType,omitempty" xml:"TreeType,omitempty"`
	// The UserID of the user to query in Quick BI.
	//
	// This parameter is required.
	//
	// example:
	//
	// b5d8fd9348cc4327****afb604
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListRecentViewReportsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRecentViewReportsRequest) GoString() string {
	return s.String()
}

func (s *ListRecentViewReportsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListRecentViewReportsRequest) GetOffsetDay() *int32 {
	return s.OffsetDay
}

func (s *ListRecentViewReportsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRecentViewReportsRequest) GetQueryMode() *string {
	return s.QueryMode
}

func (s *ListRecentViewReportsRequest) GetTreeType() *string {
	return s.TreeType
}

func (s *ListRecentViewReportsRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListRecentViewReportsRequest) SetKeyword(v string) *ListRecentViewReportsRequest {
	s.Keyword = &v
	return s
}

func (s *ListRecentViewReportsRequest) SetOffsetDay(v int32) *ListRecentViewReportsRequest {
	s.OffsetDay = &v
	return s
}

func (s *ListRecentViewReportsRequest) SetPageSize(v int32) *ListRecentViewReportsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRecentViewReportsRequest) SetQueryMode(v string) *ListRecentViewReportsRequest {
	s.QueryMode = &v
	return s
}

func (s *ListRecentViewReportsRequest) SetTreeType(v string) *ListRecentViewReportsRequest {
	s.TreeType = &v
	return s
}

func (s *ListRecentViewReportsRequest) SetUserId(v string) *ListRecentViewReportsRequest {
	s.UserId = &v
	return s
}

func (s *ListRecentViewReportsRequest) Validate() error {
	return dara.Validate(s)
}
