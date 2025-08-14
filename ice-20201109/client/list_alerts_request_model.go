// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ListAlertsRequest
	GetCategory() *string
	SetGmtEnd(v string) *ListAlertsRequest
	GetGmtEnd() *string
	SetGmtStart(v string) *ListAlertsRequest
	GetGmtStart() *string
	SetPageNo(v int32) *ListAlertsRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListAlertsRequest
	GetPageSize() *int32
	SetResourceArn(v string) *ListAlertsRequest
	GetResourceArn() *string
	SetSortBy(v string) *ListAlertsRequest
	GetSortBy() *string
	SetSortByModifiedTime(v string) *ListAlertsRequest
	GetSortByModifiedTime() *string
}

type ListAlertsRequest struct {
	// The alert type.
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The end of the time range to query.
	//
	// example:
	//
	// 2024-11-22T16:10:45Z
	GmtEnd *string `json:"GmtEnd,omitempty" xml:"GmtEnd,omitempty"`
	// The beginning of the time range to query.
	//
	// example:
	//
	// 2024-11-21T16:10:45Z
	GmtStart *string `json:"GmtStart,omitempty" xml:"GmtStart,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ARN of the source or program.
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:ims:mediaweaver:<regionId>:<userId>:vodSource/mySourceLocation/MySource
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	// The sorting order. By default, the query results are sorted by creation time in descending order. Valid values: asc and desc.
	//
	// example:
	//
	// asc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The sorting order by modification time. Valid values: asc and desc.
	//
	// example:
	//
	// asc
	SortByModifiedTime *string `json:"SortByModifiedTime,omitempty" xml:"SortByModifiedTime,omitempty"`
}

func (s ListAlertsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlertsRequest) GoString() string {
	return s.String()
}

func (s *ListAlertsRequest) GetCategory() *string {
	return s.Category
}

func (s *ListAlertsRequest) GetGmtEnd() *string {
	return s.GmtEnd
}

func (s *ListAlertsRequest) GetGmtStart() *string {
	return s.GmtStart
}

func (s *ListAlertsRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListAlertsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAlertsRequest) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *ListAlertsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListAlertsRequest) GetSortByModifiedTime() *string {
	return s.SortByModifiedTime
}

func (s *ListAlertsRequest) SetCategory(v string) *ListAlertsRequest {
	s.Category = &v
	return s
}

func (s *ListAlertsRequest) SetGmtEnd(v string) *ListAlertsRequest {
	s.GmtEnd = &v
	return s
}

func (s *ListAlertsRequest) SetGmtStart(v string) *ListAlertsRequest {
	s.GmtStart = &v
	return s
}

func (s *ListAlertsRequest) SetPageNo(v int32) *ListAlertsRequest {
	s.PageNo = &v
	return s
}

func (s *ListAlertsRequest) SetPageSize(v int32) *ListAlertsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlertsRequest) SetResourceArn(v string) *ListAlertsRequest {
	s.ResourceArn = &v
	return s
}

func (s *ListAlertsRequest) SetSortBy(v string) *ListAlertsRequest {
	s.SortBy = &v
	return s
}

func (s *ListAlertsRequest) SetSortByModifiedTime(v string) *ListAlertsRequest {
	s.SortByModifiedTime = &v
	return s
}

func (s *ListAlertsRequest) Validate() error {
	return dara.Validate(s)
}
