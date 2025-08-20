// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBadSqlDetectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeBadSqlDetectionResponseBody
	GetAccessDeniedDetail() *string
	SetDBClusterId(v string) *DescribeBadSqlDetectionResponseBody
	GetDBClusterId() *string
	SetDetectionItems(v []*DescribeBadSqlDetectionResponseBodyDetectionItems) *DescribeBadSqlDetectionResponseBody
	GetDetectionItems() []*DescribeBadSqlDetectionResponseBodyDetectionItems
	SetRequestId(v string) *DescribeBadSqlDetectionResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeBadSqlDetectionResponseBody
	GetTotalCount() *string
}

type DescribeBadSqlDetectionResponseBody struct {
	// The information about the request denial.
	//
	// example:
	//
	// {
	//
	//     "PolicyType": "AccountLevelIdentityBasedPolicy",
	//
	//     "AuthPrincipalOwnerId": "1*****************7",
	//
	//     "EncodedDiagnosticMessage": "AQIBIAAAAOPdwKY2QLOvgMEc7SkkoJfj1kvZwsaRqNYMh10Tv0wTe0fCzaCdrvgazfNb0EnJKETgXyhR+3BIQjx9WAqZryejBsp1Bl4qI5En/D9dEhcXAtKCxCmE2kZCiEzpy8BoEUt+bs0DmlaGWO5xkEpttypLIB4rUhDvZd+zwPg4EXk4KSSWSWsurxtqDkKEMshKlQFBTKvJcKwyhk62IeYly4hQ+5IpXjkh1GQXuDRCQ==",
	//
	//     "AuthPrincipalType": "SubUser",
	//
	//     "AuthPrincipalDisplayName": "2***************9",
	//
	//     "NoPermissionType": "ImplicitDeny",
	//
	//     "AuthAction": "adb:DescribeExcessivePrimaryKeys"
	//
	// }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// amv-xxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The queried detection items and detection results.
	DetectionItems []*DescribeBadSqlDetectionResponseBodyDetectionItems `json:"DetectionItems,omitempty" xml:"DetectionItems,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 584CFCAE-E3C8-5BBB-B46C-724E77A830A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 50
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBadSqlDetectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBadSqlDetectionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBadSqlDetectionResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeBadSqlDetectionResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeBadSqlDetectionResponseBody) GetDetectionItems() []*DescribeBadSqlDetectionResponseBodyDetectionItems {
	return s.DetectionItems
}

func (s *DescribeBadSqlDetectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBadSqlDetectionResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeBadSqlDetectionResponseBody) SetAccessDeniedDetail(v string) *DescribeBadSqlDetectionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeBadSqlDetectionResponseBody) SetDBClusterId(v string) *DescribeBadSqlDetectionResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeBadSqlDetectionResponseBody) SetDetectionItems(v []*DescribeBadSqlDetectionResponseBodyDetectionItems) *DescribeBadSqlDetectionResponseBody {
	s.DetectionItems = v
	return s
}

func (s *DescribeBadSqlDetectionResponseBody) SetRequestId(v string) *DescribeBadSqlDetectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBadSqlDetectionResponseBody) SetTotalCount(v string) *DescribeBadSqlDetectionResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBadSqlDetectionResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeBadSqlDetectionResponseBodyDetectionItems struct {
	// The information about the detection result.
	//
	// example:
	//
	// SQL statements that result in high peak memory are detected.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the detection item.
	//
	// example:
	//
	// Cost
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The detection result items.
	Results []*DescribeBadSqlDetectionResponseBodyDetectionItemsResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// The severity level of the detection result. Valid values:
	//
	// 	- NORMAL
	//
	// 	- WARNING
	//
	// 	- CRITICAL
	//
	// example:
	//
	// WARNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeBadSqlDetectionResponseBodyDetectionItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeBadSqlDetectionResponseBodyDetectionItems) GoString() string {
	return s.String()
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItems) GetMessage() *string {
	return s.Message
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItems) GetName() *string {
	return s.Name
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItems) GetResults() []*DescribeBadSqlDetectionResponseBodyDetectionItemsResults {
	return s.Results
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItems) SetMessage(v string) *DescribeBadSqlDetectionResponseBodyDetectionItems {
	s.Message = &v
	return s
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItems) SetName(v string) *DescribeBadSqlDetectionResponseBodyDetectionItems {
	s.Name = &v
	return s
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItems) SetResults(v []*DescribeBadSqlDetectionResponseBodyDetectionItemsResults) *DescribeBadSqlDetectionResponseBodyDetectionItems {
	s.Results = v
	return s
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItems) SetStatus(v string) *DescribeBadSqlDetectionResponseBodyDetectionItems {
	s.Status = &v
	return s
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItems) Validate() error {
	return dara.Validate(s)
}

type DescribeBadSqlDetectionResponseBodyDetectionItemsResults struct {
	// The total execution duration. Unit: milliseconds.
	//
	// >  This value is the cumulative value of the `QueuedTime`, `TotalPlanningTime`, and `ExecutionTime` parameters.
	//
	// example:
	//
	// 709
	Cost *int64 `json:"Cost,omitempty" xml:"Cost,omitempty"`
	// The diagnostic result items.
	DiagnosisResults []*DescribeBadSqlDetectionResponseBodyDetectionItemsResultsDiagnosisResults `json:"DiagnosisResults,omitempty" xml:"DiagnosisResults,omitempty" type:"Repeated"`
	// The total CPU time consumed by all operators in the stage, which is equivalent to the total CPU time of the stage. You can use this parameter to determine which parts of the stage consume a large amount of computing resources. Unit: milliseconds.
	//
	// example:
	//
	// 2345
	OperatorCost *int64 `json:"OperatorCost,omitempty" xml:"OperatorCost,omitempty"`
	// The amount of returned data. Unit: bytes.
	//
	// example:
	//
	// 235433
	OutputDataSize *int64 `json:"OutputDataSize,omitempty" xml:"OutputDataSize,omitempty"`
	// The SQL pattern ID.
	//
	// example:
	//
	// 3467484070025860498
	PatternId *string `json:"PatternId,omitempty" xml:"PatternId,omitempty"`
	// The peak memory. Unit: bytes.
	//
	// example:
	//
	// 234
	PeakMemory *int64 `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	// The query ID.
	//
	// example:
	//
	// 202410161002191720161451770345363xxxx
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The SQL statement.
	//
	// >  For performance considerations, an SQL statement cannot exceed 5,120 characters in length. Otherwise, the SQL statement is truncated. You can call the [DownloadDiagnosisRecords](https://help.aliyun.com/document_detail/308212.html) operation to download the information about SQL statements that meet a query condition for an AnalyticDB for MySQL cluster, including the complete SQL statements.
	//
	// example:
	//
	// SELECT 	- FROM device WHERE product_key = \\"h66zXfxet2X\\" AND name = \\"device@zntbtfptv5_9237117\\"
	SQL *string `json:"SQL,omitempty" xml:"SQL,omitempty"`
	// The amount of scanned data. Unit: bytes.
	//
	// example:
	//
	// 2342
	ScanSize *int64 `json:"ScanSize,omitempty" xml:"ScanSize,omitempty"`
	// The start time of the query. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-09-06T02:11:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The total number of stages generated.
	//
	// example:
	//
	// 5
	TotalStages *int32 `json:"TotalStages,omitempty" xml:"TotalStages,omitempty"`
}

func (s DescribeBadSqlDetectionResponseBodyDetectionItemsResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeBadSqlDetectionResponseBodyDetectionItemsResults) GoString() string {
	return s.String()
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResults) GetCost() *int64 {
	return s.Cost
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResults) GetDiagnosisResults() []*DescribeBadSqlDetectionResponseBodyDetectionItemsResultsDiagnosisResults {
	return s.DiagnosisResults
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResults) GetOperatorCost() *int64 {
	return s.OperatorCost
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResults) GetOutputDataSize() *int64 {
	return s.OutputDataSize
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResults) GetPatternId() *string {
	return s.PatternId
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResults) GetPeakMemory() *int64 {
	return s.PeakMemory
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResults) GetProcessId() *string {
	return s.ProcessId
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResults) GetSQL() *string {
	return s.SQL
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResults) GetScanSize() *int64 {
	return s.ScanSize
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResults) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResults) GetTotalStages() *int32 {
	return s.TotalStages
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResults) SetCost(v int64) *DescribeBadSqlDetectionResponseBodyDetectionItemsResults {
	s.Cost = &v
	return s
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResults) SetDiagnosisResults(v []*DescribeBadSqlDetectionResponseBodyDetectionItemsResultsDiagnosisResults) *DescribeBadSqlDetectionResponseBodyDetectionItemsResults {
	s.DiagnosisResults = v
	return s
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResults) SetOperatorCost(v int64) *DescribeBadSqlDetectionResponseBodyDetectionItemsResults {
	s.OperatorCost = &v
	return s
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResults) SetOutputDataSize(v int64) *DescribeBadSqlDetectionResponseBodyDetectionItemsResults {
	s.OutputDataSize = &v
	return s
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResults) SetPatternId(v string) *DescribeBadSqlDetectionResponseBodyDetectionItemsResults {
	s.PatternId = &v
	return s
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResults) SetPeakMemory(v int64) *DescribeBadSqlDetectionResponseBodyDetectionItemsResults {
	s.PeakMemory = &v
	return s
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResults) SetProcessId(v string) *DescribeBadSqlDetectionResponseBodyDetectionItemsResults {
	s.ProcessId = &v
	return s
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResults) SetSQL(v string) *DescribeBadSqlDetectionResponseBodyDetectionItemsResults {
	s.SQL = &v
	return s
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResults) SetScanSize(v int64) *DescribeBadSqlDetectionResponseBodyDetectionItemsResults {
	s.ScanSize = &v
	return s
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResults) SetStartTime(v string) *DescribeBadSqlDetectionResponseBodyDetectionItemsResults {
	s.StartTime = &v
	return s
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResults) SetTotalStages(v int32) *DescribeBadSqlDetectionResponseBodyDetectionItemsResults {
	s.TotalStages = &v
	return s
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResults) Validate() error {
	return dara.Validate(s)
}

type DescribeBadSqlDetectionResponseBodyDetectionItemsResultsDiagnosisResults struct {
	// The diagnostic code.
	//
	// example:
	//
	// Large amounts of data are returned to the client.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the diagnostic result.
	//
	// example:
	//
	// Large amounts of data are returned to the client. Import the data to OSS.
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The operator ID.
	//
	// example:
	//
	// TableScan[234]
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// The stage ID.
	//
	// example:
	//
	// Stage[67]
	StageId *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
}

func (s DescribeBadSqlDetectionResponseBodyDetectionItemsResultsDiagnosisResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeBadSqlDetectionResponseBodyDetectionItemsResultsDiagnosisResults) GoString() string {
	return s.String()
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResultsDiagnosisResults) GetCode() *string {
	return s.Code
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResultsDiagnosisResults) GetDetail() *string {
	return s.Detail
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResultsDiagnosisResults) GetOperatorId() *string {
	return s.OperatorId
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResultsDiagnosisResults) GetStageId() *string {
	return s.StageId
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResultsDiagnosisResults) SetCode(v string) *DescribeBadSqlDetectionResponseBodyDetectionItemsResultsDiagnosisResults {
	s.Code = &v
	return s
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResultsDiagnosisResults) SetDetail(v string) *DescribeBadSqlDetectionResponseBodyDetectionItemsResultsDiagnosisResults {
	s.Detail = &v
	return s
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResultsDiagnosisResults) SetOperatorId(v string) *DescribeBadSqlDetectionResponseBodyDetectionItemsResultsDiagnosisResults {
	s.OperatorId = &v
	return s
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResultsDiagnosisResults) SetStageId(v string) *DescribeBadSqlDetectionResponseBodyDetectionItemsResultsDiagnosisResults {
	s.StageId = &v
	return s
}

func (s *DescribeBadSqlDetectionResponseBodyDetectionItemsResultsDiagnosisResults) Validate() error {
	return dara.Validate(s)
}
