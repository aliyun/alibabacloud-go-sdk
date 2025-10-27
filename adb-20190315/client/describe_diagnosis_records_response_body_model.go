// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosisRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeDiagnosisRecordsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDiagnosisRecordsResponseBody
	GetPageSize() *int32
	SetQuerys(v []*DescribeDiagnosisRecordsResponseBodyQuerys) *DescribeDiagnosisRecordsResponseBody
	GetQuerys() []*DescribeDiagnosisRecordsResponseBodyQuerys
	SetRequestId(v string) *DescribeDiagnosisRecordsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDiagnosisRecordsResponseBody
	GetTotalCount() *int32
}

type DescribeDiagnosisRecordsResponseBody struct {
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
	// The queried SQL statements.
	Querys []*DescribeDiagnosisRecordsResponseBodyQuerys `json:"Querys,omitempty" xml:"Querys,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 109462AF-B5FA-3D5A-9377-B27E5B******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDiagnosisRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisRecordsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDiagnosisRecordsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDiagnosisRecordsResponseBody) GetQuerys() []*DescribeDiagnosisRecordsResponseBodyQuerys {
	return s.Querys
}

func (s *DescribeDiagnosisRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiagnosisRecordsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDiagnosisRecordsResponseBody) SetPageNumber(v int32) *DescribeDiagnosisRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBody) SetPageSize(v int32) *DescribeDiagnosisRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBody) SetQuerys(v []*DescribeDiagnosisRecordsResponseBodyQuerys) *DescribeDiagnosisRecordsResponseBody {
	s.Querys = v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBody) SetRequestId(v string) *DescribeDiagnosisRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBody) SetTotalCount(v int32) *DescribeDiagnosisRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBody) Validate() error {
	if s.Querys != nil {
		for _, item := range s.Querys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDiagnosisRecordsResponseBodyQuerys struct {
	// The source IP address.
	//
	// example:
	//
	// 59.82.xx.xx
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The total execution duration. Unit: milliseconds.
	//
	// >  This value is the cumulative value of the `QueuedTime`, `TotalPlanningTime`, and `ExecutionTime` parameters.
	//
	// example:
	//
	// 10
	Cost *int64 `json:"Cost,omitempty" xml:"Cost,omitempty"`
	// The name of the database on which the SQL statement is executed.
	//
	// example:
	//
	// adb_demo
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The number of rows written to the table by an extract, transform, and load (ETL) task.
	//
	// example:
	//
	// 0
	EtlWriteRows *int64 `json:"EtlWriteRows,omitempty" xml:"EtlWriteRows,omitempty"`
	// The execution duration. Unit: milliseconds.
	//
	// example:
	//
	// 6
	ExecutionTime *int64 `json:"ExecutionTime,omitempty" xml:"ExecutionTime,omitempty"`
	// The amount of returned data. Unit: bytes.
	//
	// example:
	//
	// 9
	OutputDataSize *int64 `json:"OutputDataSize,omitempty" xml:"OutputDataSize,omitempty"`
	// The number of rows returned.
	//
	// example:
	//
	// 1
	OutputRows *int64 `json:"OutputRows,omitempty" xml:"OutputRows,omitempty"`
	// The SQL pattern ID.
	//
	// >  You can call the [DescribePatternPerformance](https://help.aliyun.com/document_detail/612503.html) operation to query the performance metrics of an SQL pattern within a time range.
	//
	// example:
	//
	// -5575924945138******
	PatternId *string `json:"PatternId,omitempty" xml:"PatternId,omitempty"`
	// The peak memory. Unit: bytes.
	//
	// example:
	//
	// 16648
	PeakMemory *int64 `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	// The query ID.
	//
	// example:
	//
	// 2021093000414401000000023503151******
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The query properties.
	QueryProperties []*DescribeDiagnosisRecordsResponseBodyQuerysQueryProperties `json:"QueryProperties,omitempty" xml:"QueryProperties,omitempty" type:"Repeated"`
	// The amount of time that is consumed for queuing. Unit: milliseconds.
	//
	// example:
	//
	// 0
	QueueTime *int64 `json:"QueueTime,omitempty" xml:"QueueTime,omitempty"`
	// The IP address and port number of the AnalyticDB for MySQL frontend node on which the SQL statement is executed.
	//
	// example:
	//
	// 10.0.xx.xx:3004
	RcHost *string `json:"RcHost,omitempty" xml:"RcHost,omitempty"`
	// The execution duration rank of operators that are used in the SQL statement.
	//
	// > This field is returned only for SQL statements that have the `Status` parameter set to `running`.
	//
	// example:
	//
	// 1
	ResourceCostRank *int32 `json:"ResourceCostRank,omitempty" xml:"ResourceCostRank,omitempty"`
	// The resource group to which the SQL statement belongs.
	//
	// example:
	//
	// user_default
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The SQL statement.
	//
	// > For performance considerations, an SQL statement cannot exceed 5,120 characters in length. Otherwise, the SQL statement is truncated. You can call the [DownloadDiagnosisRecords](https://help.aliyun.com/document_detail/308212.html) operation to download the diagnostic information about SQL statements that meet a condition in an AnalyticDB for MySQL cluster, including the complete SQL statements.
	//
	// example:
	//
	// SELECT count(*)\\nFROM nation
	SQL *string `json:"SQL,omitempty" xml:"SQL,omitempty"`
	// Indicates whether the SQL statement is truncated. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	SQLTruncated *bool `json:"SQLTruncated,omitempty" xml:"SQLTruncated,omitempty"`
	// The maximum length of the SQL statement. 5120 is returned. Unit: character. SQL statements that exceed this limit are truncated.
	//
	// example:
	//
	// 5120
	SQLTruncatedThreshold *int64 `json:"SQLTruncatedThreshold,omitempty" xml:"SQLTruncatedThreshold,omitempty"`
	// The number of entries scanned.
	//
	// example:
	//
	// 1
	ScanRows *int64 `json:"ScanRows,omitempty" xml:"ScanRows,omitempty"`
	// The amount of scanned data. Unit: bytes.
	//
	// example:
	//
	// 9
	ScanSize *int64 `json:"ScanSize,omitempty" xml:"ScanSize,omitempty"`
	// The beginning of the time range in which the SQL statement is executed. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1632933704000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the SQL statement. Valid values:
	//
	// 	- **running**
	//
	// 	- **finished**
	//
	// 	- **failed**
	//
	// example:
	//
	// finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The amount of time that is consumed to generate an execution plan. Unit: milliseconds.
	//
	// example:
	//
	// 4
	TotalPlanningTime *int64 `json:"TotalPlanningTime,omitempty" xml:"TotalPlanningTime,omitempty"`
	// The total number of stages generated.
	//
	// example:
	//
	// 2
	TotalStages *int32 `json:"TotalStages,omitempty" xml:"TotalStages,omitempty"`
	// The username that is used to execute the SQL statement.
	//
	// example:
	//
	// test_user
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDiagnosisRecordsResponseBodyQuerys) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisRecordsResponseBodyQuerys) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) GetClientIp() *string {
	return s.ClientIp
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) GetCost() *int64 {
	return s.Cost
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) GetDatabase() *string {
	return s.Database
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) GetEtlWriteRows() *int64 {
	return s.EtlWriteRows
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) GetExecutionTime() *int64 {
	return s.ExecutionTime
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) GetOutputDataSize() *int64 {
	return s.OutputDataSize
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) GetOutputRows() *int64 {
	return s.OutputRows
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) GetPatternId() *string {
	return s.PatternId
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) GetPeakMemory() *int64 {
	return s.PeakMemory
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) GetProcessId() *string {
	return s.ProcessId
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) GetQueryProperties() []*DescribeDiagnosisRecordsResponseBodyQuerysQueryProperties {
	return s.QueryProperties
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) GetQueueTime() *int64 {
	return s.QueueTime
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) GetRcHost() *string {
	return s.RcHost
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) GetResourceCostRank() *int32 {
	return s.ResourceCostRank
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) GetSQL() *string {
	return s.SQL
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) GetSQLTruncated() *bool {
	return s.SQLTruncated
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) GetSQLTruncatedThreshold() *int64 {
	return s.SQLTruncatedThreshold
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) GetScanRows() *int64 {
	return s.ScanRows
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) GetScanSize() *int64 {
	return s.ScanSize
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) GetStatus() *string {
	return s.Status
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) GetTotalPlanningTime() *int64 {
	return s.TotalPlanningTime
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) GetTotalStages() *int32 {
	return s.TotalStages
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) GetUserName() *string {
	return s.UserName
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetClientIp(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.ClientIp = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetCost(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.Cost = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetDatabase(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.Database = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetEtlWriteRows(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.EtlWriteRows = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetExecutionTime(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.ExecutionTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetOutputDataSize(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.OutputDataSize = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetOutputRows(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.OutputRows = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetPatternId(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.PatternId = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetPeakMemory(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.PeakMemory = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetProcessId(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.ProcessId = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetQueryProperties(v []*DescribeDiagnosisRecordsResponseBodyQuerysQueryProperties) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.QueryProperties = v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetQueueTime(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.QueueTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetRcHost(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.RcHost = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetResourceCostRank(v int32) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.ResourceCostRank = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetResourceGroup(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.ResourceGroup = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetSQL(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.SQL = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetSQLTruncated(v bool) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.SQLTruncated = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetSQLTruncatedThreshold(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.SQLTruncatedThreshold = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetScanRows(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.ScanRows = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetScanSize(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.ScanSize = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetStartTime(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetStatus(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.Status = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetTotalPlanningTime(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.TotalPlanningTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetTotalStages(v int32) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.TotalStages = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetUserName(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.UserName = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) Validate() error {
	if s.QueryProperties != nil {
		for _, item := range s.QueryProperties {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDiagnosisRecordsResponseBodyQuerysQueryProperties struct {
	// The name of the query property.
	//
	// example:
	//
	// ff
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the query property.
	//
	// example:
	//
	// 40
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDiagnosisRecordsResponseBodyQuerysQueryProperties) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisRecordsResponseBodyQuerysQueryProperties) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerysQueryProperties) GetName() *string {
	return s.Name
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerysQueryProperties) GetValue() *string {
	return s.Value
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerysQueryProperties) SetName(v string) *DescribeDiagnosisRecordsResponseBodyQuerysQueryProperties {
	s.Name = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerysQueryProperties) SetValue(v string) *DescribeDiagnosisRecordsResponseBodyQuerysQueryProperties {
	s.Value = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerysQueryProperties) Validate() error {
	return dara.Validate(s)
}
