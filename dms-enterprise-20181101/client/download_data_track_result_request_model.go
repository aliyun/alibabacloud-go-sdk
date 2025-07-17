// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadDataTrackResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumnFilter(v *DownloadDataTrackResultRequestColumnFilter) *DownloadDataTrackResultRequest
	GetColumnFilter() *DownloadDataTrackResultRequestColumnFilter
	SetEventIdList(v []*int64) *DownloadDataTrackResultRequest
	GetEventIdList() []*int64
	SetFilterEndTime(v string) *DownloadDataTrackResultRequest
	GetFilterEndTime() *string
	SetFilterStartTime(v string) *DownloadDataTrackResultRequest
	GetFilterStartTime() *string
	SetFilterTableList(v []*string) *DownloadDataTrackResultRequest
	GetFilterTableList() []*string
	SetFilterTypeList(v []*string) *DownloadDataTrackResultRequest
	GetFilterTypeList() []*string
	SetOrderId(v int64) *DownloadDataTrackResultRequest
	GetOrderId() *int64
	SetRollbackSQLType(v string) *DownloadDataTrackResultRequest
	GetRollbackSQLType() *string
	SetTid(v int64) *DownloadDataTrackResultRequest
	GetTid() *int64
}

type DownloadDataTrackResultRequest struct {
	// The condition to filter columns.
	ColumnFilter *DownloadDataTrackResultRequestColumnFilter `json:"ColumnFilter,omitempty" xml:"ColumnFilter,omitempty" type:"Struct"`
	// The IDs of the events.
	EventIdList []*int64 `json:"EventIdList,omitempty" xml:"EventIdList,omitempty" type:"Repeated"`
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
	FilterTableList []*string `json:"FilterTableList,omitempty" xml:"FilterTableList,omitempty" type:"Repeated"`
	// The types of data operations that you want to track.
	FilterTypeList []*string `json:"FilterTypeList,omitempty" xml:"FilterTypeList,omitempty" type:"Repeated"`
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

func (s DownloadDataTrackResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DownloadDataTrackResultRequest) GoString() string {
	return s.String()
}

func (s *DownloadDataTrackResultRequest) GetColumnFilter() *DownloadDataTrackResultRequestColumnFilter {
	return s.ColumnFilter
}

func (s *DownloadDataTrackResultRequest) GetEventIdList() []*int64 {
	return s.EventIdList
}

func (s *DownloadDataTrackResultRequest) GetFilterEndTime() *string {
	return s.FilterEndTime
}

func (s *DownloadDataTrackResultRequest) GetFilterStartTime() *string {
	return s.FilterStartTime
}

func (s *DownloadDataTrackResultRequest) GetFilterTableList() []*string {
	return s.FilterTableList
}

func (s *DownloadDataTrackResultRequest) GetFilterTypeList() []*string {
	return s.FilterTypeList
}

func (s *DownloadDataTrackResultRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *DownloadDataTrackResultRequest) GetRollbackSQLType() *string {
	return s.RollbackSQLType
}

func (s *DownloadDataTrackResultRequest) GetTid() *int64 {
	return s.Tid
}

func (s *DownloadDataTrackResultRequest) SetColumnFilter(v *DownloadDataTrackResultRequestColumnFilter) *DownloadDataTrackResultRequest {
	s.ColumnFilter = v
	return s
}

func (s *DownloadDataTrackResultRequest) SetEventIdList(v []*int64) *DownloadDataTrackResultRequest {
	s.EventIdList = v
	return s
}

func (s *DownloadDataTrackResultRequest) SetFilterEndTime(v string) *DownloadDataTrackResultRequest {
	s.FilterEndTime = &v
	return s
}

func (s *DownloadDataTrackResultRequest) SetFilterStartTime(v string) *DownloadDataTrackResultRequest {
	s.FilterStartTime = &v
	return s
}

func (s *DownloadDataTrackResultRequest) SetFilterTableList(v []*string) *DownloadDataTrackResultRequest {
	s.FilterTableList = v
	return s
}

func (s *DownloadDataTrackResultRequest) SetFilterTypeList(v []*string) *DownloadDataTrackResultRequest {
	s.FilterTypeList = v
	return s
}

func (s *DownloadDataTrackResultRequest) SetOrderId(v int64) *DownloadDataTrackResultRequest {
	s.OrderId = &v
	return s
}

func (s *DownloadDataTrackResultRequest) SetRollbackSQLType(v string) *DownloadDataTrackResultRequest {
	s.RollbackSQLType = &v
	return s
}

func (s *DownloadDataTrackResultRequest) SetTid(v int64) *DownloadDataTrackResultRequest {
	s.Tid = &v
	return s
}

func (s *DownloadDataTrackResultRequest) Validate() error {
	return dara.Validate(s)
}

type DownloadDataTrackResultRequestColumnFilter struct {
	// The end value of the range used in the filter condition. This parameter takes effect only when Operator is set to BETWEEN.
	//
	// example:
	//
	// 10
	BetweenEnd *string `json:"BetweenEnd,omitempty" xml:"BetweenEnd,omitempty"`
	// The start value of the range used in the filter condition. This parameter takes effect only when Operator is set to BETWEEN.
	//
	// example:
	//
	// 1
	BetweenStart *string `json:"BetweenStart,omitempty" xml:"BetweenStart,omitempty"`
	// The name of the column.
	//
	// example:
	//
	// account_name
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The IN list used in the filter condition.
	InList []*string `json:"InList,omitempty" xml:"InList,omitempty" type:"Repeated"`
	// The type of the operator used to configure the filter condition. Valid values:
	//
	// 	- **EQUAL**: retrieves the column whose value is equal to the specified value.
	//
	// 	- **NOT_EQUAL**: retrieves the column whose value is not equal to the specified value.
	//
	// 	- **IN**: retrieves the column whose value is in the IN list.
	//
	// 	- **BETWEEN**: retrieves the column whose value is in the specified range.
	//
	// 	- **LESS**: retrieves the column whose value is less than the specified value.
	//
	// 	- **MORE**: retrieves the column whose value is greater than the specified value.
	//
	// 	- **NOT_IN**: retrieves the column whose value is not in the IN list.
	//
	// example:
	//
	// EQUAL
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The value used in the filter condition.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DownloadDataTrackResultRequestColumnFilter) String() string {
	return dara.Prettify(s)
}

func (s DownloadDataTrackResultRequestColumnFilter) GoString() string {
	return s.String()
}

func (s *DownloadDataTrackResultRequestColumnFilter) GetBetweenEnd() *string {
	return s.BetweenEnd
}

func (s *DownloadDataTrackResultRequestColumnFilter) GetBetweenStart() *string {
	return s.BetweenStart
}

func (s *DownloadDataTrackResultRequestColumnFilter) GetColumnName() *string {
	return s.ColumnName
}

func (s *DownloadDataTrackResultRequestColumnFilter) GetInList() []*string {
	return s.InList
}

func (s *DownloadDataTrackResultRequestColumnFilter) GetOperator() *string {
	return s.Operator
}

func (s *DownloadDataTrackResultRequestColumnFilter) GetValue() *string {
	return s.Value
}

func (s *DownloadDataTrackResultRequestColumnFilter) SetBetweenEnd(v string) *DownloadDataTrackResultRequestColumnFilter {
	s.BetweenEnd = &v
	return s
}

func (s *DownloadDataTrackResultRequestColumnFilter) SetBetweenStart(v string) *DownloadDataTrackResultRequestColumnFilter {
	s.BetweenStart = &v
	return s
}

func (s *DownloadDataTrackResultRequestColumnFilter) SetColumnName(v string) *DownloadDataTrackResultRequestColumnFilter {
	s.ColumnName = &v
	return s
}

func (s *DownloadDataTrackResultRequestColumnFilter) SetInList(v []*string) *DownloadDataTrackResultRequestColumnFilter {
	s.InList = v
	return s
}

func (s *DownloadDataTrackResultRequestColumnFilter) SetOperator(v string) *DownloadDataTrackResultRequestColumnFilter {
	s.Operator = &v
	return s
}

func (s *DownloadDataTrackResultRequestColumnFilter) SetValue(v string) *DownloadDataTrackResultRequestColumnFilter {
	s.Value = &v
	return s
}

func (s *DownloadDataTrackResultRequestColumnFilter) Validate() error {
	return dara.Validate(s)
}
