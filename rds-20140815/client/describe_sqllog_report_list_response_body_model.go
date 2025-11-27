// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLLogReportListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeSQLLogReportListResponseBodyItems) *DescribeSQLLogReportListResponseBody
	GetItems() *DescribeSQLLogReportListResponseBodyItems
	SetPageNumber(v int32) *DescribeSQLLogReportListResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeSQLLogReportListResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeSQLLogReportListResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeSQLLogReportListResponseBody
	GetTotalRecordCount() *int32
}

type DescribeSQLLogReportListResponseBody struct {
	// An array that consists of SQL log reports.
	Items *DescribeSQLLogReportListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of SQL log reports on the current page.
	//
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 08A3B71B-FE08-4B03-974F-CC7EA6DB1828
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 60
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeSQLLogReportListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogReportListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogReportListResponseBody) GetItems() *DescribeSQLLogReportListResponseBodyItems {
	return s.Items
}

func (s *DescribeSQLLogReportListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSQLLogReportListResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeSQLLogReportListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSQLLogReportListResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeSQLLogReportListResponseBody) SetItems(v *DescribeSQLLogReportListResponseBodyItems) *DescribeSQLLogReportListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSQLLogReportListResponseBody) SetPageNumber(v int32) *DescribeSQLLogReportListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogReportListResponseBody) SetPageRecordCount(v int32) *DescribeSQLLogReportListResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSQLLogReportListResponseBody) SetRequestId(v string) *DescribeSQLLogReportListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLLogReportListResponseBody) SetTotalRecordCount(v int32) *DescribeSQLLogReportListResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeSQLLogReportListResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSQLLogReportListResponseBodyItems struct {
	Item []*DescribeSQLLogReportListResponseBodyItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s DescribeSQLLogReportListResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogReportListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogReportListResponseBodyItems) GetItem() []*DescribeSQLLogReportListResponseBodyItemsItem {
	return s.Item
}

func (s *DescribeSQLLogReportListResponseBodyItems) SetItem(v []*DescribeSQLLogReportListResponseBodyItemsItem) *DescribeSQLLogReportListResponseBodyItems {
	s.Item = v
	return s
}

func (s *DescribeSQLLogReportListResponseBodyItems) Validate() error {
	if s.Item != nil {
		for _, item := range s.Item {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSQLLogReportListResponseBodyItemsItem struct {
	// An array that consists of SQL statements executed with the highest latency.
	LatencyTopNItems *DescribeSQLLogReportListResponseBodyItemsItemLatencyTopNItems `json:"LatencyTopNItems,omitempty" xml:"LatencyTopNItems,omitempty" type:"Struct"`
	// An array that consists of SQL statements executed the most frequently.
	QPSTopNItems *DescribeSQLLogReportListResponseBodyItemsItemQPSTopNItems `json:"QPSTopNItems,omitempty" xml:"QPSTopNItems,omitempty" type:"Struct"`
	// The time when the report was generated. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-03-27T16:00:00Z
	ReportTime *string `json:"ReportTime,omitempty" xml:"ReportTime,omitempty"`
}

func (s DescribeSQLLogReportListResponseBodyItemsItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogReportListResponseBodyItemsItem) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogReportListResponseBodyItemsItem) GetLatencyTopNItems() *DescribeSQLLogReportListResponseBodyItemsItemLatencyTopNItems {
	return s.LatencyTopNItems
}

func (s *DescribeSQLLogReportListResponseBodyItemsItem) GetQPSTopNItems() *DescribeSQLLogReportListResponseBodyItemsItemQPSTopNItems {
	return s.QPSTopNItems
}

func (s *DescribeSQLLogReportListResponseBodyItemsItem) GetReportTime() *string {
	return s.ReportTime
}

func (s *DescribeSQLLogReportListResponseBodyItemsItem) SetLatencyTopNItems(v *DescribeSQLLogReportListResponseBodyItemsItemLatencyTopNItems) *DescribeSQLLogReportListResponseBodyItemsItem {
	s.LatencyTopNItems = v
	return s
}

func (s *DescribeSQLLogReportListResponseBodyItemsItem) SetQPSTopNItems(v *DescribeSQLLogReportListResponseBodyItemsItemQPSTopNItems) *DescribeSQLLogReportListResponseBodyItemsItem {
	s.QPSTopNItems = v
	return s
}

func (s *DescribeSQLLogReportListResponseBodyItemsItem) SetReportTime(v string) *DescribeSQLLogReportListResponseBodyItemsItem {
	s.ReportTime = &v
	return s
}

func (s *DescribeSQLLogReportListResponseBodyItemsItem) Validate() error {
	if s.LatencyTopNItems != nil {
		if err := s.LatencyTopNItems.Validate(); err != nil {
			return err
		}
	}
	if s.QPSTopNItems != nil {
		if err := s.QPSTopNItems.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSQLLogReportListResponseBodyItemsItemLatencyTopNItems struct {
	LatencyTopNItem []*DescribeSQLLogReportListResponseBodyItemsItemLatencyTopNItemsLatencyTopNItem `json:"LatencyTopNItem,omitempty" xml:"LatencyTopNItem,omitempty" type:"Repeated"`
}

func (s DescribeSQLLogReportListResponseBodyItemsItemLatencyTopNItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogReportListResponseBodyItemsItemLatencyTopNItems) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogReportListResponseBodyItemsItemLatencyTopNItems) GetLatencyTopNItem() []*DescribeSQLLogReportListResponseBodyItemsItemLatencyTopNItemsLatencyTopNItem {
	return s.LatencyTopNItem
}

func (s *DescribeSQLLogReportListResponseBodyItemsItemLatencyTopNItems) SetLatencyTopNItem(v []*DescribeSQLLogReportListResponseBodyItemsItemLatencyTopNItemsLatencyTopNItem) *DescribeSQLLogReportListResponseBodyItemsItemLatencyTopNItems {
	s.LatencyTopNItem = v
	return s
}

func (s *DescribeSQLLogReportListResponseBodyItemsItemLatencyTopNItems) Validate() error {
	if s.LatencyTopNItem != nil {
		for _, item := range s.LatencyTopNItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSQLLogReportListResponseBodyItemsItemLatencyTopNItemsLatencyTopNItem struct {
	// The average time that is required to execute the SQL statement. Unit: milliseconds.
	//
	// example:
	//
	// 500
	AvgLatency *int64 `json:"AvgLatency,omitempty" xml:"AvgLatency,omitempty"`
	// The number of times that the SQL statement is executed.
	//
	// example:
	//
	// 5
	SQLExecuteTimes *int64 `json:"SQLExecuteTimes,omitempty" xml:"SQLExecuteTimes,omitempty"`
	// The SQL statement.
	//
	// >  Only the first 128 characters of the SQL statement are returned. In addition, only the SQL statements that take more than 100 ms to execute are returned.
	//
	// example:
	//
	// SELECT 	- FROM table_name;
	SQLText *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
}

func (s DescribeSQLLogReportListResponseBodyItemsItemLatencyTopNItemsLatencyTopNItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogReportListResponseBodyItemsItemLatencyTopNItemsLatencyTopNItem) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogReportListResponseBodyItemsItemLatencyTopNItemsLatencyTopNItem) GetAvgLatency() *int64 {
	return s.AvgLatency
}

func (s *DescribeSQLLogReportListResponseBodyItemsItemLatencyTopNItemsLatencyTopNItem) GetSQLExecuteTimes() *int64 {
	return s.SQLExecuteTimes
}

func (s *DescribeSQLLogReportListResponseBodyItemsItemLatencyTopNItemsLatencyTopNItem) GetSQLText() *string {
	return s.SQLText
}

func (s *DescribeSQLLogReportListResponseBodyItemsItemLatencyTopNItemsLatencyTopNItem) SetAvgLatency(v int64) *DescribeSQLLogReportListResponseBodyItemsItemLatencyTopNItemsLatencyTopNItem {
	s.AvgLatency = &v
	return s
}

func (s *DescribeSQLLogReportListResponseBodyItemsItemLatencyTopNItemsLatencyTopNItem) SetSQLExecuteTimes(v int64) *DescribeSQLLogReportListResponseBodyItemsItemLatencyTopNItemsLatencyTopNItem {
	s.SQLExecuteTimes = &v
	return s
}

func (s *DescribeSQLLogReportListResponseBodyItemsItemLatencyTopNItemsLatencyTopNItem) SetSQLText(v string) *DescribeSQLLogReportListResponseBodyItemsItemLatencyTopNItemsLatencyTopNItem {
	s.SQLText = &v
	return s
}

func (s *DescribeSQLLogReportListResponseBodyItemsItemLatencyTopNItemsLatencyTopNItem) Validate() error {
	return dara.Validate(s)
}

type DescribeSQLLogReportListResponseBodyItemsItemQPSTopNItems struct {
	QPSTopNItem []*DescribeSQLLogReportListResponseBodyItemsItemQPSTopNItemsQPSTopNItem `json:"QPSTopNItem,omitempty" xml:"QPSTopNItem,omitempty" type:"Repeated"`
}

func (s DescribeSQLLogReportListResponseBodyItemsItemQPSTopNItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogReportListResponseBodyItemsItemQPSTopNItems) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogReportListResponseBodyItemsItemQPSTopNItems) GetQPSTopNItem() []*DescribeSQLLogReportListResponseBodyItemsItemQPSTopNItemsQPSTopNItem {
	return s.QPSTopNItem
}

func (s *DescribeSQLLogReportListResponseBodyItemsItemQPSTopNItems) SetQPSTopNItem(v []*DescribeSQLLogReportListResponseBodyItemsItemQPSTopNItemsQPSTopNItem) *DescribeSQLLogReportListResponseBodyItemsItemQPSTopNItems {
	s.QPSTopNItem = v
	return s
}

func (s *DescribeSQLLogReportListResponseBodyItemsItemQPSTopNItems) Validate() error {
	if s.QPSTopNItem != nil {
		for _, item := range s.QPSTopNItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSQLLogReportListResponseBodyItemsItemQPSTopNItemsQPSTopNItem struct {
	// The number of times that the SQL statement is executed.
	//
	// example:
	//
	// 5
	SQLExecuteTimes *int64 `json:"SQLExecuteTimes,omitempty" xml:"SQLExecuteTimes,omitempty"`
	// The SQL statement.
	//
	// >  Only the first 128 characters of the SQL statement are returned. In addition, only the SQL statements that take more than 5 ms to execute are returned.
	//
	// example:
	//
	// SELECT 	- FROM table_name;
	SQLText *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
}

func (s DescribeSQLLogReportListResponseBodyItemsItemQPSTopNItemsQPSTopNItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogReportListResponseBodyItemsItemQPSTopNItemsQPSTopNItem) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogReportListResponseBodyItemsItemQPSTopNItemsQPSTopNItem) GetSQLExecuteTimes() *int64 {
	return s.SQLExecuteTimes
}

func (s *DescribeSQLLogReportListResponseBodyItemsItemQPSTopNItemsQPSTopNItem) GetSQLText() *string {
	return s.SQLText
}

func (s *DescribeSQLLogReportListResponseBodyItemsItemQPSTopNItemsQPSTopNItem) SetSQLExecuteTimes(v int64) *DescribeSQLLogReportListResponseBodyItemsItemQPSTopNItemsQPSTopNItem {
	s.SQLExecuteTimes = &v
	return s
}

func (s *DescribeSQLLogReportListResponseBodyItemsItemQPSTopNItemsQPSTopNItem) SetSQLText(v string) *DescribeSQLLogReportListResponseBodyItemsItemQPSTopNItemsQPSTopNItem {
	s.SQLText = &v
	return s
}

func (s *DescribeSQLLogReportListResponseBodyItemsItemQPSTopNItemsQPSTopNItem) Validate() error {
	return dara.Validate(s)
}
