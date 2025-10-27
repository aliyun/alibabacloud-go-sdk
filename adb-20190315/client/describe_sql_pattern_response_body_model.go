// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlPatternResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeSqlPatternResponseBodyItems) *DescribeSqlPatternResponseBody
	GetItems() []*DescribeSqlPatternResponseBodyItems
	SetPageNumber(v int32) *DescribeSqlPatternResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSqlPatternResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSqlPatternResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeSqlPatternResponseBody
	GetTotalCount() *int32
}

type DescribeSqlPatternResponseBody struct {
	// The queried SQL pattern.
	Items []*DescribeSqlPatternResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6F2D1B4-2C9F-5622-B424-5E7965******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSqlPatternResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlPatternResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSqlPatternResponseBody) GetItems() []*DescribeSqlPatternResponseBodyItems {
	return s.Items
}

func (s *DescribeSqlPatternResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSqlPatternResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSqlPatternResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSqlPatternResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSqlPatternResponseBody) SetItems(v []*DescribeSqlPatternResponseBodyItems) *DescribeSqlPatternResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSqlPatternResponseBody) SetPageNumber(v int32) *DescribeSqlPatternResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSqlPatternResponseBody) SetPageSize(v int32) *DescribeSqlPatternResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSqlPatternResponseBody) SetRequestId(v string) *DescribeSqlPatternResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSqlPatternResponseBody) SetTotalCount(v int32) *DescribeSqlPatternResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSqlPatternResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSqlPatternResponseBodyItems struct {
	// The IP address of the client.
	//
	// > This parameter is returned only when `Type` is set to `accessip`.
	//
	// example:
	//
	// 100.104.***.***
	AccessIP *string `json:"AccessIP,omitempty" xml:"AccessIP,omitempty"`
	// The average execution duration of the SQL pattern within the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1.0625
	AvgCpuTime *string `json:"AvgCpuTime,omitempty" xml:"AvgCpuTime,omitempty"`
	// The average peak memory usage of the SQL pattern within the query time range. Unit: KB.
	//
	// example:
	//
	// 240048
	AvgPeakMemory *string `json:"AvgPeakMemory,omitempty" xml:"AvgPeakMemory,omitempty"`
	// The average amount of data scanned based on the SQL pattern within the query time range. Unit: KB.
	//
	// example:
	//
	// 0
	AvgScanSize *string `json:"AvgScanSize,omitempty" xml:"AvgScanSize,omitempty"`
	// The average number of stages.
	//
	// example:
	//
	// 2
	AvgStageCount *string `json:"AvgStageCount,omitempty" xml:"AvgStageCount,omitempty"`
	// The average number of tasks.
	//
	// example:
	//
	// 2
	AvgTaskCount *string `json:"AvgTaskCount,omitempty" xml:"AvgTaskCount,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// am-bp1r053byu48p****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The maximum execution duration of the SQL pattern within the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 17
	MaxCpuTime *string `json:"MaxCpuTime,omitempty" xml:"MaxCpuTime,omitempty"`
	// The maximum peak memory usage of the SQL pattern within the query time range. Unit: KB.
	//
	// example:
	//
	// 480096
	MaxPeakMemory *string `json:"MaxPeakMemory,omitempty" xml:"MaxPeakMemory,omitempty"`
	// The maximum amount of data scanned based on the SQL pattern within the query time range. Unit: KB.
	//
	// example:
	//
	// 0
	MaxScanSize *string `json:"MaxScanSize,omitempty" xml:"MaxScanSize,omitempty"`
	// The maximum number of stages.
	//
	// example:
	//
	// 2
	MaxStageCount *string `json:"MaxStageCount,omitempty" xml:"MaxStageCount,omitempty"`
	// The maximum number of tasks.
	//
	// example:
	//
	// 2
	MaxTaskCount *string `json:"MaxTaskCount,omitempty" xml:"MaxTaskCount,omitempty"`
	// The SQL pattern.
	//
	// example:
	//
	// SELECT table_name, table_schema AS schema_name, create_time, create_time AS last_ddl_time, table_comment AS description , ceil((data_length + index_length) / ? / ?) AS store_capacity , data_length AS data_bytes, index_length AS index_bytes, table_collation AS collation, auto_increment, table_rows AS num_rows , engine FROM information_schema.tables WHERE table_type != ? AND table_schema = ? AND table_name IN (?) ORDER BY 1
	Pattern *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	// The number of queries performed in association with the SQL pattern within the query time range.
	//
	// example:
	//
	// 16
	QueryCount *string `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
	// The start date of the query.
	//
	// example:
	//
	// 2021-08-30
	ReportDate *string `json:"ReportDate,omitempty" xml:"ReportDate,omitempty"`
	// The username.
	//
	// > This parameter is returned only when `Type` is left empty or set to `user`.
	//
	// example:
	//
	// test_acc
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeSqlPatternResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlPatternResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSqlPatternResponseBodyItems) GetAccessIP() *string {
	return s.AccessIP
}

func (s *DescribeSqlPatternResponseBodyItems) GetAvgCpuTime() *string {
	return s.AvgCpuTime
}

func (s *DescribeSqlPatternResponseBodyItems) GetAvgPeakMemory() *string {
	return s.AvgPeakMemory
}

func (s *DescribeSqlPatternResponseBodyItems) GetAvgScanSize() *string {
	return s.AvgScanSize
}

func (s *DescribeSqlPatternResponseBodyItems) GetAvgStageCount() *string {
	return s.AvgStageCount
}

func (s *DescribeSqlPatternResponseBodyItems) GetAvgTaskCount() *string {
	return s.AvgTaskCount
}

func (s *DescribeSqlPatternResponseBodyItems) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeSqlPatternResponseBodyItems) GetMaxCpuTime() *string {
	return s.MaxCpuTime
}

func (s *DescribeSqlPatternResponseBodyItems) GetMaxPeakMemory() *string {
	return s.MaxPeakMemory
}

func (s *DescribeSqlPatternResponseBodyItems) GetMaxScanSize() *string {
	return s.MaxScanSize
}

func (s *DescribeSqlPatternResponseBodyItems) GetMaxStageCount() *string {
	return s.MaxStageCount
}

func (s *DescribeSqlPatternResponseBodyItems) GetMaxTaskCount() *string {
	return s.MaxTaskCount
}

func (s *DescribeSqlPatternResponseBodyItems) GetPattern() *string {
	return s.Pattern
}

func (s *DescribeSqlPatternResponseBodyItems) GetQueryCount() *string {
	return s.QueryCount
}

func (s *DescribeSqlPatternResponseBodyItems) GetReportDate() *string {
	return s.ReportDate
}

func (s *DescribeSqlPatternResponseBodyItems) GetUser() *string {
	return s.User
}

func (s *DescribeSqlPatternResponseBodyItems) SetAccessIP(v string) *DescribeSqlPatternResponseBodyItems {
	s.AccessIP = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetAvgCpuTime(v string) *DescribeSqlPatternResponseBodyItems {
	s.AvgCpuTime = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetAvgPeakMemory(v string) *DescribeSqlPatternResponseBodyItems {
	s.AvgPeakMemory = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetAvgScanSize(v string) *DescribeSqlPatternResponseBodyItems {
	s.AvgScanSize = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetAvgStageCount(v string) *DescribeSqlPatternResponseBodyItems {
	s.AvgStageCount = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetAvgTaskCount(v string) *DescribeSqlPatternResponseBodyItems {
	s.AvgTaskCount = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetInstanceName(v string) *DescribeSqlPatternResponseBodyItems {
	s.InstanceName = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetMaxCpuTime(v string) *DescribeSqlPatternResponseBodyItems {
	s.MaxCpuTime = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetMaxPeakMemory(v string) *DescribeSqlPatternResponseBodyItems {
	s.MaxPeakMemory = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetMaxScanSize(v string) *DescribeSqlPatternResponseBodyItems {
	s.MaxScanSize = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetMaxStageCount(v string) *DescribeSqlPatternResponseBodyItems {
	s.MaxStageCount = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetMaxTaskCount(v string) *DescribeSqlPatternResponseBodyItems {
	s.MaxTaskCount = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetPattern(v string) *DescribeSqlPatternResponseBodyItems {
	s.Pattern = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetQueryCount(v string) *DescribeSqlPatternResponseBodyItems {
	s.QueryCount = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetReportDate(v string) *DescribeSqlPatternResponseBodyItems {
	s.ReportDate = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetUser(v string) *DescribeSqlPatternResponseBodyItems {
	s.User = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
