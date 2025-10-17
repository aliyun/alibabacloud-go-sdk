// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSparkSQLDiagnosisListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeSparkSQLDiagnosisListResponseBody
	GetAccessDeniedDetail() *string
	SetPageNumber(v int32) *DescribeSparkSQLDiagnosisListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSparkSQLDiagnosisListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSparkSQLDiagnosisListResponseBody
	GetRequestId() *string
	SetSQLDiagnosisList(v []*DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList) *DescribeSparkSQLDiagnosisListResponseBody
	GetSQLDiagnosisList() []*DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList
	SetTotalCount(v int32) *DescribeSparkSQLDiagnosisListResponseBody
	GetTotalCount() *int32
}

type DescribeSparkSQLDiagnosisListResponseBody struct {
	// The information about the request denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
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
	// A91C9D07-7462-5F35-BB47-83629CE6CCAC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried diagnostic information.
	SQLDiagnosisList []*DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList `json:"SQLDiagnosisList,omitempty" xml:"SQLDiagnosisList,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1343
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSparkSQLDiagnosisListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSparkSQLDiagnosisListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSparkSQLDiagnosisListResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeSparkSQLDiagnosisListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSparkSQLDiagnosisListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSparkSQLDiagnosisListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSparkSQLDiagnosisListResponseBody) GetSQLDiagnosisList() []*DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList {
	return s.SQLDiagnosisList
}

func (s *DescribeSparkSQLDiagnosisListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSparkSQLDiagnosisListResponseBody) SetAccessDeniedDetail(v string) *DescribeSparkSQLDiagnosisListResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisListResponseBody) SetPageNumber(v int32) *DescribeSparkSQLDiagnosisListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisListResponseBody) SetPageSize(v int32) *DescribeSparkSQLDiagnosisListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisListResponseBody) SetRequestId(v string) *DescribeSparkSQLDiagnosisListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisListResponseBody) SetSQLDiagnosisList(v []*DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList) *DescribeSparkSQLDiagnosisListResponseBody {
	s.SQLDiagnosisList = v
	return s
}

func (s *DescribeSparkSQLDiagnosisListResponseBody) SetTotalCount(v int32) *DescribeSparkSQLDiagnosisListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisListResponseBody) Validate() error {
	if s.SQLDiagnosisList != nil {
		for _, item := range s.SQLDiagnosisList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList struct {
	// The application ID.
	//
	// >  You can call the [ListSparkApps](https://help.aliyun.com/document_detail/612475.html) operation to query a list of Spark application IDs.
	//
	// example:
	//
	// s202404291020bjd448ad40002122
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The execution duration of the query. Unit: milliseconds.
	//
	// example:
	//
	// 100
	ElapsedTime *int64 `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	// The ID of the query executed within the Spark application.
	//
	// example:
	//
	// 1
	InnerQueryId *int64 `json:"InnerQueryId,omitempty" xml:"InnerQueryId,omitempty"`
	// The maximum operator execution duration. Unit: milliseconds.
	//
	// example:
	//
	// 90
	MaxExclusiveTime *int64 `json:"MaxExclusiveTime,omitempty" xml:"MaxExclusiveTime,omitempty"`
	// The maximum operator memory used. Unit: bytes.
	//
	// example:
	//
	// 1024
	PeakMemory *int64 `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	// The SQL statement.
	//
	// example:
	//
	// select 	- from device where name = \\"105506012112790031\\"
	SQL *string `json:"SQL,omitempty" xml:"SQL,omitempty"`
	// The number of entries scanned.
	//
	// example:
	//
	// 100
	ScanRowCount *int64 `json:"ScanRowCount,omitempty" xml:"ScanRowCount,omitempty"`
	// The start time of the query. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-11-20 09:09:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The execution status of the query. Valid values:
	//
	// 	- COMPLETED
	//
	// 	- CANCELED
	//
	// 	- ABORTED
	//
	// 	- FAILED
	//
	// example:
	//
	// COMPLETED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The unique ID of the code block in the Spark job.
	//
	// example:
	//
	// 1
	StatementId *int64 `json:"StatementId,omitempty" xml:"StatementId,omitempty"`
}

func (s DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList) GoString() string {
	return s.String()
}

func (s *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList) GetAppId() *string {
	return s.AppId
}

func (s *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList) GetElapsedTime() *int64 {
	return s.ElapsedTime
}

func (s *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList) GetInnerQueryId() *int64 {
	return s.InnerQueryId
}

func (s *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList) GetMaxExclusiveTime() *int64 {
	return s.MaxExclusiveTime
}

func (s *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList) GetPeakMemory() *int64 {
	return s.PeakMemory
}

func (s *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList) GetSQL() *string {
	return s.SQL
}

func (s *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList) GetScanRowCount() *int64 {
	return s.ScanRowCount
}

func (s *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList) GetState() *string {
	return s.State
}

func (s *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList) GetStatementId() *int64 {
	return s.StatementId
}

func (s *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList) SetAppId(v string) *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList {
	s.AppId = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList) SetElapsedTime(v int64) *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList {
	s.ElapsedTime = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList) SetInnerQueryId(v int64) *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList {
	s.InnerQueryId = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList) SetMaxExclusiveTime(v int64) *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList {
	s.MaxExclusiveTime = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList) SetPeakMemory(v int64) *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList {
	s.PeakMemory = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList) SetSQL(v string) *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList {
	s.SQL = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList) SetScanRowCount(v int64) *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList {
	s.ScanRowCount = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList) SetStartTime(v string) *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList {
	s.StartTime = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList) SetState(v string) *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList {
	s.State = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList) SetStatementId(v int64) *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList {
	s.StatementId = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisListResponseBodySQLDiagnosisList) Validate() error {
	return dara.Validate(s)
}
