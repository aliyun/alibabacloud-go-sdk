// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadDataTrackResultShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumnFilterShrink(v string) *DownloadDataTrackResultShrinkRequest
	GetColumnFilterShrink() *string
	SetEventIdListShrink(v string) *DownloadDataTrackResultShrinkRequest
	GetEventIdListShrink() *string
	SetFilterEndTime(v string) *DownloadDataTrackResultShrinkRequest
	GetFilterEndTime() *string
	SetFilterStartTime(v string) *DownloadDataTrackResultShrinkRequest
	GetFilterStartTime() *string
	SetFilterTableListShrink(v string) *DownloadDataTrackResultShrinkRequest
	GetFilterTableListShrink() *string
	SetFilterTypeListShrink(v string) *DownloadDataTrackResultShrinkRequest
	GetFilterTypeListShrink() *string
	SetOrderId(v int64) *DownloadDataTrackResultShrinkRequest
	GetOrderId() *int64
	SetRollbackSQLType(v string) *DownloadDataTrackResultShrinkRequest
	GetRollbackSQLType() *string
	SetTid(v int64) *DownloadDataTrackResultShrinkRequest
	GetTid() *int64
}

type DownloadDataTrackResultShrinkRequest struct {
	// The condition to filter columns.
	ColumnFilterShrink *string `json:"ColumnFilter,omitempty" xml:"ColumnFilter,omitempty"`
	// The IDs of the events.
	EventIdListShrink *string `json:"EventIdList,omitempty" xml:"EventIdList,omitempty"`
	// The end time of the time range in which you want to track data operations. The time must be specified in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2023-04-23 10:00:00
	FilterEndTime *string `json:"FilterEndTime,omitempty" xml:"FilterEndTime,omitempty"`
	// The start time of the time range in which you want to track data operations. The time must be specified in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2023-04-23 00:00:00
	FilterStartTime *string `json:"FilterStartTime,omitempty" xml:"FilterStartTime,omitempty"`
	// The names of the tables for which you want to track data operations.
	FilterTableListShrink *string `json:"FilterTableList,omitempty" xml:"FilterTableList,omitempty"`
	// The types of data operations that you want to track.
	FilterTypeListShrink *string `json:"FilterTypeList,omitempty" xml:"FilterTypeList,omitempty"`
	// The ID of the ticket. You can call the [ListOrders](https://help.aliyun.com/document_detail/144643.html) operation to obtain the ticket ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 406****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The type of the SQL statement.
	//
	// 	- **REVERSE**: undoes or rolls back an executed SQL statement, which is equivalent to the UNDO SQL statement.
	//
	// 	- **FORWARD**: redoes or re-executes an SQL statement that failed to be executed, which is equivalent to the REDO SQL statement.
	//
	// This parameter is required.
	//
	// example:
	//
	// REVERSE
	RollbackSQLType *string `json:"RollbackSQLType,omitempty" xml:"RollbackSQLType,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s DownloadDataTrackResultShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DownloadDataTrackResultShrinkRequest) GoString() string {
	return s.String()
}

func (s *DownloadDataTrackResultShrinkRequest) GetColumnFilterShrink() *string {
	return s.ColumnFilterShrink
}

func (s *DownloadDataTrackResultShrinkRequest) GetEventIdListShrink() *string {
	return s.EventIdListShrink
}

func (s *DownloadDataTrackResultShrinkRequest) GetFilterEndTime() *string {
	return s.FilterEndTime
}

func (s *DownloadDataTrackResultShrinkRequest) GetFilterStartTime() *string {
	return s.FilterStartTime
}

func (s *DownloadDataTrackResultShrinkRequest) GetFilterTableListShrink() *string {
	return s.FilterTableListShrink
}

func (s *DownloadDataTrackResultShrinkRequest) GetFilterTypeListShrink() *string {
	return s.FilterTypeListShrink
}

func (s *DownloadDataTrackResultShrinkRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *DownloadDataTrackResultShrinkRequest) GetRollbackSQLType() *string {
	return s.RollbackSQLType
}

func (s *DownloadDataTrackResultShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *DownloadDataTrackResultShrinkRequest) SetColumnFilterShrink(v string) *DownloadDataTrackResultShrinkRequest {
	s.ColumnFilterShrink = &v
	return s
}

func (s *DownloadDataTrackResultShrinkRequest) SetEventIdListShrink(v string) *DownloadDataTrackResultShrinkRequest {
	s.EventIdListShrink = &v
	return s
}

func (s *DownloadDataTrackResultShrinkRequest) SetFilterEndTime(v string) *DownloadDataTrackResultShrinkRequest {
	s.FilterEndTime = &v
	return s
}

func (s *DownloadDataTrackResultShrinkRequest) SetFilterStartTime(v string) *DownloadDataTrackResultShrinkRequest {
	s.FilterStartTime = &v
	return s
}

func (s *DownloadDataTrackResultShrinkRequest) SetFilterTableListShrink(v string) *DownloadDataTrackResultShrinkRequest {
	s.FilterTableListShrink = &v
	return s
}

func (s *DownloadDataTrackResultShrinkRequest) SetFilterTypeListShrink(v string) *DownloadDataTrackResultShrinkRequest {
	s.FilterTypeListShrink = &v
	return s
}

func (s *DownloadDataTrackResultShrinkRequest) SetOrderId(v int64) *DownloadDataTrackResultShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *DownloadDataTrackResultShrinkRequest) SetRollbackSQLType(v string) *DownloadDataTrackResultShrinkRequest {
	s.RollbackSQLType = &v
	return s
}

func (s *DownloadDataTrackResultShrinkRequest) SetTid(v int64) *DownloadDataTrackResultShrinkRequest {
	s.Tid = &v
	return s
}

func (s *DownloadDataTrackResultShrinkRequest) Validate() error {
	return dara.Validate(s)
}
