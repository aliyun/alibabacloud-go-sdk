// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchDataTrackResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumnFilter(v *SearchDataTrackResultRequestColumnFilter) *SearchDataTrackResultRequest
	GetColumnFilter() *SearchDataTrackResultRequestColumnFilter
	SetFilterEndTime(v string) *SearchDataTrackResultRequest
	GetFilterEndTime() *string
	SetFilterStartTime(v string) *SearchDataTrackResultRequest
	GetFilterStartTime() *string
	SetFilterTableList(v []*string) *SearchDataTrackResultRequest
	GetFilterTableList() []*string
	SetFilterTypeList(v []*string) *SearchDataTrackResultRequest
	GetFilterTypeList() []*string
	SetOrderId(v int64) *SearchDataTrackResultRequest
	GetOrderId() *int64
	SetTid(v int64) *SearchDataTrackResultRequest
	GetTid() *int64
}

type SearchDataTrackResultRequest struct {
	// The condition to filter columns.
	ColumnFilter *SearchDataTrackResultRequestColumnFilter `json:"ColumnFilter,omitempty" xml:"ColumnFilter,omitempty" type:"Struct"`
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
	FilterTableList []*string `json:"FilterTableList,omitempty" xml:"FilterTableList,omitempty" type:"Repeated"`
	// The types of data operations that you want to track.
	FilterTypeList []*string `json:"FilterTypeList,omitempty" xml:"FilterTypeList,omitempty" type:"Repeated"`
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

func (s SearchDataTrackResultRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchDataTrackResultRequest) GoString() string {
	return s.String()
}

func (s *SearchDataTrackResultRequest) GetColumnFilter() *SearchDataTrackResultRequestColumnFilter {
	return s.ColumnFilter
}

func (s *SearchDataTrackResultRequest) GetFilterEndTime() *string {
	return s.FilterEndTime
}

func (s *SearchDataTrackResultRequest) GetFilterStartTime() *string {
	return s.FilterStartTime
}

func (s *SearchDataTrackResultRequest) GetFilterTableList() []*string {
	return s.FilterTableList
}

func (s *SearchDataTrackResultRequest) GetFilterTypeList() []*string {
	return s.FilterTypeList
}

func (s *SearchDataTrackResultRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *SearchDataTrackResultRequest) GetTid() *int64 {
	return s.Tid
}

func (s *SearchDataTrackResultRequest) SetColumnFilter(v *SearchDataTrackResultRequestColumnFilter) *SearchDataTrackResultRequest {
	s.ColumnFilter = v
	return s
}

func (s *SearchDataTrackResultRequest) SetFilterEndTime(v string) *SearchDataTrackResultRequest {
	s.FilterEndTime = &v
	return s
}

func (s *SearchDataTrackResultRequest) SetFilterStartTime(v string) *SearchDataTrackResultRequest {
	s.FilterStartTime = &v
	return s
}

func (s *SearchDataTrackResultRequest) SetFilterTableList(v []*string) *SearchDataTrackResultRequest {
	s.FilterTableList = v
	return s
}

func (s *SearchDataTrackResultRequest) SetFilterTypeList(v []*string) *SearchDataTrackResultRequest {
	s.FilterTypeList = v
	return s
}

func (s *SearchDataTrackResultRequest) SetOrderId(v int64) *SearchDataTrackResultRequest {
	s.OrderId = &v
	return s
}

func (s *SearchDataTrackResultRequest) SetTid(v int64) *SearchDataTrackResultRequest {
	s.Tid = &v
	return s
}

func (s *SearchDataTrackResultRequest) Validate() error {
	if s.ColumnFilter != nil {
		if err := s.ColumnFilter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchDataTrackResultRequestColumnFilter struct {
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
	// c_payer_name
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The IN list used in the filter condition. This parameter takes effect only when Operator is set to IN or NOT_IN.
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

func (s SearchDataTrackResultRequestColumnFilter) String() string {
	return dara.Prettify(s)
}

func (s SearchDataTrackResultRequestColumnFilter) GoString() string {
	return s.String()
}

func (s *SearchDataTrackResultRequestColumnFilter) GetBetweenEnd() *string {
	return s.BetweenEnd
}

func (s *SearchDataTrackResultRequestColumnFilter) GetBetweenStart() *string {
	return s.BetweenStart
}

func (s *SearchDataTrackResultRequestColumnFilter) GetColumnName() *string {
	return s.ColumnName
}

func (s *SearchDataTrackResultRequestColumnFilter) GetInList() []*string {
	return s.InList
}

func (s *SearchDataTrackResultRequestColumnFilter) GetOperator() *string {
	return s.Operator
}

func (s *SearchDataTrackResultRequestColumnFilter) GetValue() *string {
	return s.Value
}

func (s *SearchDataTrackResultRequestColumnFilter) SetBetweenEnd(v string) *SearchDataTrackResultRequestColumnFilter {
	s.BetweenEnd = &v
	return s
}

func (s *SearchDataTrackResultRequestColumnFilter) SetBetweenStart(v string) *SearchDataTrackResultRequestColumnFilter {
	s.BetweenStart = &v
	return s
}

func (s *SearchDataTrackResultRequestColumnFilter) SetColumnName(v string) *SearchDataTrackResultRequestColumnFilter {
	s.ColumnName = &v
	return s
}

func (s *SearchDataTrackResultRequestColumnFilter) SetInList(v []*string) *SearchDataTrackResultRequestColumnFilter {
	s.InList = v
	return s
}

func (s *SearchDataTrackResultRequestColumnFilter) SetOperator(v string) *SearchDataTrackResultRequestColumnFilter {
	s.Operator = &v
	return s
}

func (s *SearchDataTrackResultRequestColumnFilter) SetValue(v string) *SearchDataTrackResultRequestColumnFilter {
	s.Value = &v
	return s
}

func (s *SearchDataTrackResultRequestColumnFilter) Validate() error {
	return dara.Validate(s)
}
