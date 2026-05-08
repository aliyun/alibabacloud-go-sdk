// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeSlowLogRecordsResponseBody
	GetDBClusterId() *string
	SetItems(v *DescribeSlowLogRecordsResponseBodyItems) *DescribeSlowLogRecordsResponseBody
	GetItems() *DescribeSlowLogRecordsResponseBodyItems
	SetPageNumber(v string) *DescribeSlowLogRecordsResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeSlowLogRecordsResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeSlowLogRecordsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeSlowLogRecordsResponseBody
	GetTotalCount() *string
}

type DescribeSlowLogRecordsResponseBody struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// example:
	//
	// am-bp1rqvm70uh2****
	DBClusterId *string                                  `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Items       *DescribeSlowLogRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D7559209-7EC3-41E1-8F78-156990******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSlowLogRecordsResponseBody) GetItems() *DescribeSlowLogRecordsResponseBodyItems {
	return s.Items
}

func (s *DescribeSlowLogRecordsResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeSlowLogRecordsResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeSlowLogRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSlowLogRecordsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeSlowLogRecordsResponseBody) SetDBClusterId(v string) *DescribeSlowLogRecordsResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetItems(v *DescribeSlowLogRecordsResponseBodyItems) *DescribeSlowLogRecordsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetPageNumber(v string) *DescribeSlowLogRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetPageSize(v string) *DescribeSlowLogRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetRequestId(v string) *DescribeSlowLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetTotalCount(v string) *DescribeSlowLogRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSlowLogRecordsResponseBodyItems struct {
	SlowLogRecord []*DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord `json:"SlowLogRecord,omitempty" xml:"SlowLogRecord,omitempty" type:"Repeated"`
}

func (s DescribeSlowLogRecordsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetSlowLogRecord() []*DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	return s.SlowLogRecord
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetSlowLogRecord(v []*DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) *DescribeSlowLogRecordsResponseBodyItems {
	s.SlowLogRecord = v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) Validate() error {
	if s.SlowLogRecord != nil {
		for _, item := range s.SlowLogRecord {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord struct {
	DBName             *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" xml:"ExecutionStartTime,omitempty"`
	HostAddress        *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	OutputSize         *string `json:"OutputSize,omitempty" xml:"OutputSize,omitempty"`
	ParseRowCounts     *int64  `json:"ParseRowCounts,omitempty" xml:"ParseRowCounts,omitempty"`
	PeakMemoryUsage    *string `json:"PeakMemoryUsage,omitempty" xml:"PeakMemoryUsage,omitempty"`
	PlanningTime       *int64  `json:"PlanningTime,omitempty" xml:"PlanningTime,omitempty"`
	ProcessID          *string `json:"ProcessID,omitempty" xml:"ProcessID,omitempty"`
	QueryTime          *int64  `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
	QueueTime          *int64  `json:"QueueTime,omitempty" xml:"QueueTime,omitempty"`
	ReturnRowCounts    *int64  `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	SQLText            *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	ScanRows           *int64  `json:"ScanRows,omitempty" xml:"ScanRows,omitempty"`
	ScanSize           *string `json:"ScanSize,omitempty" xml:"ScanSize,omitempty"`
	ScanTime           *int64  `json:"ScanTime,omitempty" xml:"ScanTime,omitempty"`
	State              *string `json:"State,omitempty" xml:"State,omitempty"`
	UserName           *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	WallTime           *int64  `json:"WallTime,omitempty" xml:"WallTime,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) GetDBName() *string {
	return s.DBName
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) GetExecutionStartTime() *string {
	return s.ExecutionStartTime
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) GetHostAddress() *string {
	return s.HostAddress
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) GetOutputSize() *string {
	return s.OutputSize
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) GetParseRowCounts() *int64 {
	return s.ParseRowCounts
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) GetPeakMemoryUsage() *string {
	return s.PeakMemoryUsage
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) GetPlanningTime() *int64 {
	return s.PlanningTime
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) GetProcessID() *string {
	return s.ProcessID
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) GetQueryTime() *int64 {
	return s.QueryTime
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) GetQueueTime() *int64 {
	return s.QueueTime
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) GetReturnRowCounts() *int64 {
	return s.ReturnRowCounts
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) GetSQLText() *string {
	return s.SQLText
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) GetScanRows() *int64 {
	return s.ScanRows
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) GetScanSize() *string {
	return s.ScanSize
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) GetScanTime() *int64 {
	return s.ScanTime
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) GetState() *string {
	return s.State
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) GetUserName() *string {
	return s.UserName
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) GetWallTime() *int64 {
	return s.WallTime
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetDBName(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetExecutionStartTime(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.ExecutionStartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetHostAddress(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.HostAddress = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetOutputSize(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.OutputSize = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetParseRowCounts(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.ParseRowCounts = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetPeakMemoryUsage(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.PeakMemoryUsage = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetPlanningTime(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.PlanningTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetProcessID(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.ProcessID = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetQueryTime(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.QueryTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetQueueTime(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.QueueTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetReturnRowCounts(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetSQLText(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.SQLText = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetScanRows(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.ScanRows = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetScanSize(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.ScanSize = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetScanTime(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.ScanTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetState(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.State = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetUserName(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.UserName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetWallTime(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.WallTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) Validate() error {
	return dara.Validate(s)
}
