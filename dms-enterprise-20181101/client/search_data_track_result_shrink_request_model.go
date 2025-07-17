// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchDataTrackResultShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumnFilterShrink(v string) *SearchDataTrackResultShrinkRequest
	GetColumnFilterShrink() *string
	SetFilterEndTime(v string) *SearchDataTrackResultShrinkRequest
	GetFilterEndTime() *string
	SetFilterStartTime(v string) *SearchDataTrackResultShrinkRequest
	GetFilterStartTime() *string
	SetFilterTableListShrink(v string) *SearchDataTrackResultShrinkRequest
	GetFilterTableListShrink() *string
	SetFilterTypeListShrink(v string) *SearchDataTrackResultShrinkRequest
	GetFilterTypeListShrink() *string
	SetOrderId(v int64) *SearchDataTrackResultShrinkRequest
	GetOrderId() *int64
	SetTid(v int64) *SearchDataTrackResultShrinkRequest
	GetTid() *int64
}

type SearchDataTrackResultShrinkRequest struct {
	// The condition to filter columns.
	ColumnFilterShrink *string `json:"ColumnFilter,omitempty" xml:"ColumnFilter,omitempty"`
	// The end time of the time range in which you want to track data operations. The time must be in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2023-04-23 10:00:00
	FilterEndTime *string `json:"FilterEndTime,omitempty" xml:"FilterEndTime,omitempty"`
	// The start time of the time range in which you want to track data operations. The time must be in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2023-04-23 00:00:00
	FilterStartTime *string `json:"FilterStartTime,omitempty" xml:"FilterStartTime,omitempty"`
	// The names of the tables for which you want to track data operations.
	FilterTableListShrink *string `json:"FilterTableList,omitempty" xml:"FilterTableList,omitempty"`
	// The types of data operations that you want to track.
	FilterTypeListShrink *string `json:"FilterTypeList,omitempty" xml:"FilterTypeList,omitempty"`
	// The ID of the ticket. You can call the [ListOrders](https://help.aliyun.com/document_detail/144643.html) operation to query the ticket ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 420****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to query the tenant ID.
	//
	// example:
	//
	// 62***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s SearchDataTrackResultShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchDataTrackResultShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchDataTrackResultShrinkRequest) GetColumnFilterShrink() *string {
	return s.ColumnFilterShrink
}

func (s *SearchDataTrackResultShrinkRequest) GetFilterEndTime() *string {
	return s.FilterEndTime
}

func (s *SearchDataTrackResultShrinkRequest) GetFilterStartTime() *string {
	return s.FilterStartTime
}

func (s *SearchDataTrackResultShrinkRequest) GetFilterTableListShrink() *string {
	return s.FilterTableListShrink
}

func (s *SearchDataTrackResultShrinkRequest) GetFilterTypeListShrink() *string {
	return s.FilterTypeListShrink
}

func (s *SearchDataTrackResultShrinkRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *SearchDataTrackResultShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *SearchDataTrackResultShrinkRequest) SetColumnFilterShrink(v string) *SearchDataTrackResultShrinkRequest {
	s.ColumnFilterShrink = &v
	return s
}

func (s *SearchDataTrackResultShrinkRequest) SetFilterEndTime(v string) *SearchDataTrackResultShrinkRequest {
	s.FilterEndTime = &v
	return s
}

func (s *SearchDataTrackResultShrinkRequest) SetFilterStartTime(v string) *SearchDataTrackResultShrinkRequest {
	s.FilterStartTime = &v
	return s
}

func (s *SearchDataTrackResultShrinkRequest) SetFilterTableListShrink(v string) *SearchDataTrackResultShrinkRequest {
	s.FilterTableListShrink = &v
	return s
}

func (s *SearchDataTrackResultShrinkRequest) SetFilterTypeListShrink(v string) *SearchDataTrackResultShrinkRequest {
	s.FilterTypeListShrink = &v
	return s
}

func (s *SearchDataTrackResultShrinkRequest) SetOrderId(v int64) *SearchDataTrackResultShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *SearchDataTrackResultShrinkRequest) SetTid(v int64) *SearchDataTrackResultShrinkRequest {
	s.Tid = &v
	return s
}

func (s *SearchDataTrackResultShrinkRequest) Validate() error {
	return dara.Validate(s)
}
