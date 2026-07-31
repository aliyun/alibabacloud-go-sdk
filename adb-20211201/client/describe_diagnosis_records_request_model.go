// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosisRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientIp(v string) *DescribeDiagnosisRecordsRequest
	GetClientIp() *string
	SetDBClusterId(v string) *DescribeDiagnosisRecordsRequest
	GetDBClusterId() *string
	SetDatabase(v string) *DescribeDiagnosisRecordsRequest
	GetDatabase() *string
	SetEndTime(v string) *DescribeDiagnosisRecordsRequest
	GetEndTime() *string
	SetKeyword(v string) *DescribeDiagnosisRecordsRequest
	GetKeyword() *string
	SetLang(v string) *DescribeDiagnosisRecordsRequest
	GetLang() *string
	SetMaxPeakMemory(v int64) *DescribeDiagnosisRecordsRequest
	GetMaxPeakMemory() *int64
	SetMaxScanSize(v int64) *DescribeDiagnosisRecordsRequest
	GetMaxScanSize() *int64
	SetMinPeakMemory(v int64) *DescribeDiagnosisRecordsRequest
	GetMinPeakMemory() *int64
	SetMinScanSize(v int64) *DescribeDiagnosisRecordsRequest
	GetMinScanSize() *int64
	SetOrder(v string) *DescribeDiagnosisRecordsRequest
	GetOrder() *string
	SetPageNumber(v int32) *DescribeDiagnosisRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDiagnosisRecordsRequest
	GetPageSize() *int32
	SetPatternId(v string) *DescribeDiagnosisRecordsRequest
	GetPatternId() *string
	SetQueryCondition(v string) *DescribeDiagnosisRecordsRequest
	GetQueryCondition() *string
	SetRegionId(v string) *DescribeDiagnosisRecordsRequest
	GetRegionId() *string
	SetResourceGroup(v string) *DescribeDiagnosisRecordsRequest
	GetResourceGroup() *string
	SetStartTime(v string) *DescribeDiagnosisRecordsRequest
	GetStartTime() *string
	SetUserName(v string) *DescribeDiagnosisRecordsRequest
	GetUserName() *string
}

type DescribeDiagnosisRecordsRequest struct {
	// The source IP address.
	//
	// > Call the [DescribeDiagnosisDimensions](https://help.aliyun.com/document_detail/308210.html) operation to view the resource groups, database names, usernames, and source IP addresses for the SQL statements that meet the specified query conditions.
	//
	// example:
	//
	// 59.82.XX.XX
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The ID of the Enterprise Edition, Basic Edition, or Data Lakehouse Edition cluster.
	//
	// > Call the [DescribeDBClusters](https://help.aliyun.com/document_detail/612397.html) operation to view the details of all clusters in your account, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1scs48yc125****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database where the SQL statement is executed.
	//
	// > Call the [DescribeDiagnosisDimensions](https://help.aliyun.com/document_detail/308210.html) operation to view the resource groups, database names, usernames, and source IP addresses for the SQL statements that meet the specified query conditions.
	//
	// example:
	//
	// adb_demo
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The end of the time range to query. Specify the time in the UNIX timestamp format. The time must be in milliseconds.
	//
	// > - The end time must be later than the start time.
	//
	// >
	//
	// > - The interval between the start time and the end time cannot exceed 24 hours.
	//
	// example:
	//
	// 1633017540000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Filters the queries by the keywords contained in the SQL statements.
	//
	// example:
	//
	// select
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language of the file title and some error messages in the downloaded file. Valid values:
	//
	// - **zh**: Simplified Chinese (default).
	//
	// - **en**: English.
	//
	// - **ja**: Japanese.
	//
	// - **zh-tw**: Traditional Chinese.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum peak memory of the SQL statement. Unit: bytes.
	//
	// example:
	//
	// 89000000
	MaxPeakMemory *int64 `json:"MaxPeakMemory,omitempty" xml:"MaxPeakMemory,omitempty"`
	// The maximum scan size of the target SQL statement. Unit: bytes.
	//
	// example:
	//
	// 1024000000
	MaxScanSize *int64 `json:"MaxScanSize,omitempty" xml:"MaxScanSize,omitempty"`
	// The minimum peak memory of the SQL statement. Unit: bytes.
	//
	// example:
	//
	// 0
	MinPeakMemory *int64 `json:"MinPeakMemory,omitempty" xml:"MinPeakMemory,omitempty"`
	// The minimum scan size of the SQL statement. Unit: bytes.
	//
	// example:
	//
	// 0
	MinScanSize *int64 `json:"MinScanSize,omitempty" xml:"MinScanSize,omitempty"`
	// The sorting order of the SQL statements. This parameter is a JSON array that is ordered by the sequence of the input array. It contains the `Field` and `Type` fields. Example: `[{"Field":"StartTime", "Type": "desc" }]`. The fields are described as follows:
	//
	// - `Field` specifies the field by which to sort the SQL statements. Valid values:
	//
	//   - `StartTime`: the start time of the execution.
	//
	//   - `Status`: the execution state.
	//
	//   - `UserName`: the username.
	//
	//   - `Cost`: the execution duration.
	//
	//   - `PeakMemory`: the peak memory.
	//
	//   - `ScanSize`: the amount of scanned data.
	//
	//   - `Database`: the database name.
	//
	//   - `ClientIp`: the source IP address.
	//
	//   - `ResourceGroup`: the resource group.
	//
	//   - `QueueTime`: the amount of time that the query waited in a queue.
	//
	//   - `OutputRows`: the number of output rows.
	//
	//   - `OutputDataSize`: the amount of output data.
	//
	//   - `ResourceCostRank`: the ranking of the execution duration of an operator in the SQL statement. This field is returned only when `QueryCondition` is set to `{"Type":"status","Value":"running"}`.
	//
	// - `Type` specifies the sorting type. Valid values (case-insensitive):
	//
	//   - `Desc`: descending order.
	//
	//   - `Asc`: ascending order.
	//
	// example:
	//
	// [{"Field":"StartTime", "Type": "desc" }]
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. The value must be an integer that is greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// - **30*	- (default)
	//
	// - **50**
	//
	// - **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the SQL pattern.
	//
	// example:
	//
	// 5575924945138******
	PatternId *string `json:"PatternId,omitempty" xml:"PatternId,omitempty"`
	// The conditions for the SQL query. This parameter is a JSON string that contains fields such as Type, `Value`, `Min`, and `Max`. The `Type` field indicates the query dimension. Valid values for `Type`: `maxCost`, `status`, and `cost`. The `Value`, `Min`, and `Max` fields specify the query range for the dimension. Valid values:
	//
	// - `{"Type":"maxCost","Value":"100"}`: queries the details of the top 100 SQL statements that have the longest execution durations. The `Value` field can only be set to 100.
	//
	// - `{"Type":"status","Value":"finished"}`: queries the details of completed SQL statements. You can also set `Value` to `running` or `failed` to query SQL statements that are running or have failed.
	//
	// - `{"Type":"cost","Min":"10","Max":"200"}`: queries the details of SQL statements whose execution durations are between 10 ms and 200 ms. You can customize the minimum and maximum execution durations. Unit: milliseconds.
	//
	// example:
	//
	// {"Type":"status","Value":"finished"}
	QueryCondition *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	// The region ID.
	//
	// > Call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to view the regions and zones supported by AnalyticDB for MySQL, including region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group to which the SQL statement belongs.
	//
	// > Call the [DescribeDiagnosisDimensions](https://help.aliyun.com/document_detail/308210.html) operation to view the resource groups, database names, usernames, and source IP addresses for the SQL statements that meet the specified query conditions.
	//
	// example:
	//
	// user_default
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The start of the time range to query. Specify the time in the UNIX timestamp format. The time must be in milliseconds.
	//
	// > Only data from the last 14 days can be queried.
	//
	// example:
	//
	// 1632931200000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The username used to execute the SQL statement.
	//
	// Call the [DescribeDiagnosisDimensions](https://help.aliyun.com/document_detail/308210.html) operation to view the resource groups, database names, usernames, and source IP addresses for the SQL statements that meet the specified query conditions.
	//
	// example:
	//
	// test_user
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDiagnosisRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisRecordsRequest) GetClientIp() *string {
	return s.ClientIp
}

func (s *DescribeDiagnosisRecordsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDiagnosisRecordsRequest) GetDatabase() *string {
	return s.Database
}

func (s *DescribeDiagnosisRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDiagnosisRecordsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeDiagnosisRecordsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDiagnosisRecordsRequest) GetMaxPeakMemory() *int64 {
	return s.MaxPeakMemory
}

func (s *DescribeDiagnosisRecordsRequest) GetMaxScanSize() *int64 {
	return s.MaxScanSize
}

func (s *DescribeDiagnosisRecordsRequest) GetMinPeakMemory() *int64 {
	return s.MinPeakMemory
}

func (s *DescribeDiagnosisRecordsRequest) GetMinScanSize() *int64 {
	return s.MinScanSize
}

func (s *DescribeDiagnosisRecordsRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeDiagnosisRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDiagnosisRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDiagnosisRecordsRequest) GetPatternId() *string {
	return s.PatternId
}

func (s *DescribeDiagnosisRecordsRequest) GetQueryCondition() *string {
	return s.QueryCondition
}

func (s *DescribeDiagnosisRecordsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDiagnosisRecordsRequest) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *DescribeDiagnosisRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDiagnosisRecordsRequest) GetUserName() *string {
	return s.UserName
}

func (s *DescribeDiagnosisRecordsRequest) SetClientIp(v string) *DescribeDiagnosisRecordsRequest {
	s.ClientIp = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetDBClusterId(v string) *DescribeDiagnosisRecordsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetDatabase(v string) *DescribeDiagnosisRecordsRequest {
	s.Database = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetEndTime(v string) *DescribeDiagnosisRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetKeyword(v string) *DescribeDiagnosisRecordsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetLang(v string) *DescribeDiagnosisRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetMaxPeakMemory(v int64) *DescribeDiagnosisRecordsRequest {
	s.MaxPeakMemory = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetMaxScanSize(v int64) *DescribeDiagnosisRecordsRequest {
	s.MaxScanSize = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetMinPeakMemory(v int64) *DescribeDiagnosisRecordsRequest {
	s.MinPeakMemory = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetMinScanSize(v int64) *DescribeDiagnosisRecordsRequest {
	s.MinScanSize = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetOrder(v string) *DescribeDiagnosisRecordsRequest {
	s.Order = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetPageNumber(v int32) *DescribeDiagnosisRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetPageSize(v int32) *DescribeDiagnosisRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetPatternId(v string) *DescribeDiagnosisRecordsRequest {
	s.PatternId = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetQueryCondition(v string) *DescribeDiagnosisRecordsRequest {
	s.QueryCondition = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetRegionId(v string) *DescribeDiagnosisRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetResourceGroup(v string) *DescribeDiagnosisRecordsRequest {
	s.ResourceGroup = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetStartTime(v string) *DescribeDiagnosisRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetUserName(v string) *DescribeDiagnosisRecordsRequest {
	s.UserName = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) Validate() error {
	return dara.Validate(s)
}
