// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadDiagnosisRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientIp(v string) *DownloadDiagnosisRecordsRequest
	GetClientIp() *string
	SetDBClusterId(v string) *DownloadDiagnosisRecordsRequest
	GetDBClusterId() *string
	SetDatabase(v string) *DownloadDiagnosisRecordsRequest
	GetDatabase() *string
	SetEndTime(v string) *DownloadDiagnosisRecordsRequest
	GetEndTime() *string
	SetKeyword(v string) *DownloadDiagnosisRecordsRequest
	GetKeyword() *string
	SetLang(v string) *DownloadDiagnosisRecordsRequest
	GetLang() *string
	SetMaxPeakMemory(v int64) *DownloadDiagnosisRecordsRequest
	GetMaxPeakMemory() *int64
	SetMaxScanSize(v int64) *DownloadDiagnosisRecordsRequest
	GetMaxScanSize() *int64
	SetMinPeakMemory(v int64) *DownloadDiagnosisRecordsRequest
	GetMinPeakMemory() *int64
	SetMinScanSize(v int64) *DownloadDiagnosisRecordsRequest
	GetMinScanSize() *int64
	SetQueryCondition(v string) *DownloadDiagnosisRecordsRequest
	GetQueryCondition() *string
	SetRegionId(v string) *DownloadDiagnosisRecordsRequest
	GetRegionId() *string
	SetResourceGroup(v string) *DownloadDiagnosisRecordsRequest
	GetResourceGroup() *string
	SetStartTime(v string) *DownloadDiagnosisRecordsRequest
	GetStartTime() *string
	SetUserName(v string) *DownloadDiagnosisRecordsRequest
	GetUserName() *string
}

type DownloadDiagnosisRecordsRequest struct {
	// The source IP address.
	//
	// >  You can call the [DescribeDiagnosisDimensions](https://help.aliyun.com/document_detail/308210.html) operation to query the resource groups, database names, usernames, and source IP addresses of the SQL statements that meet a query condition.
	//
	// example:
	//
	// 106.11.XX.XX
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1q8bu9a****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database on which the SQL statements are executed.
	//
	// >  You can call the [DescribeDiagnosisDimensions](https://help.aliyun.com/document_detail/308210.html) operation to query the resource groups, database names, usernames, and source IP addresses of the SQL statements that meet a query condition.
	//
	// example:
	//
	// adb_demo
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The end of the time range to query. Set the time to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >
	//
	// 	- The end time must be later than the start time.
	//
	// 	- The maximum time range that can be specified is 24 hours.
	//
	// example:
	//
	// 1662450730000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The query keyword of the SQL statements.
	//
	// example:
	//
	// select
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language. Valid values:
	//
	// 	- **zh**: simplified Chinese.
	//
	// 	- **en**: English.
	//
	// 	- **ja**: Japanese.
	//
	// 	- **zh-tw**: traditional Chinese.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum peak memory of the SQL statements. Unit: bytes.
	//
	// example:
	//
	// 88000000
	MaxPeakMemory *int64 `json:"MaxPeakMemory,omitempty" xml:"MaxPeakMemory,omitempty"`
	// The maximum scan size of the SQL statements. Unit: bytes.
	//
	// example:
	//
	// 64424509440
	MaxScanSize *int64 `json:"MaxScanSize,omitempty" xml:"MaxScanSize,omitempty"`
	// The minimum peak memory of the SQL statements. Unit: bytes.
	//
	// example:
	//
	// 88000000
	MinPeakMemory *int64 `json:"MinPeakMemory,omitempty" xml:"MinPeakMemory,omitempty"`
	// The minimum scan size of the SQL statements. Unit: bytes.
	//
	// example:
	//
	// 1073741824
	MinScanSize *int64 `json:"MinScanSize,omitempty" xml:"MinScanSize,omitempty"`
	// The query condition for SQL statements, which can contain the `Type`, `Value`, `Min`, and `Max` fields. Specify the condition in the JSON format. `Type` specifies the query dimension. Valid values for Type: `maxCost`, `status`, and `cost`. `Value`, `Min`, or `Max` specifies the query range for the dimension. Valid values:
	//
	// 	- `{"Type":"maxCost","Value":"100"}`: queries the top 100 most time-consuming SQL statements. Set `Value` to 100.
	//
	// 	- `{"Type":"status","Value":"finished"}`: queries the executed SQL statements. You can set `Value` to `running` to query the SQL statements that are being executed. You can also set Value to `failed` to query the SQL statements that failed to be executed.
	//
	// 	- `{"Type":"cost","Min":"10","Max":"200"}`: queries the SQL statements whose execution duration is in the range of 10 to 200 milliseconds. You can also specify custom values for the Min and Max fields.
	//
	// example:
	//
	// {"Type":"status","Value":"finished"}
	QueryCondition *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	// The region ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group to which the SQL statements belong.
	//
	// >  You can call the [DescribeDiagnosisDimensions](https://help.aliyun.com/document_detail/308210.html) operation to query the resource groups, database names, usernames, and source IP addresses of the SQL statements that meet a query condition.
	//
	// example:
	//
	// user_default
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The beginning of the time range to query. Set the time to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  You can query data only within the last 14 days.
	//
	// example:
	//
	// 1662364330000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The username that is used to execute the SQL statements.
	//
	// >  You can call the [DescribeDiagnosisDimensions](~~~~) operation to query the resource groups, database names, usernames, and source IP addresses of the SQL statements that meet a query condition.
	//
	// example:
	//
	// test_user
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DownloadDiagnosisRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DownloadDiagnosisRecordsRequest) GoString() string {
	return s.String()
}

func (s *DownloadDiagnosisRecordsRequest) GetClientIp() *string {
	return s.ClientIp
}

func (s *DownloadDiagnosisRecordsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DownloadDiagnosisRecordsRequest) GetDatabase() *string {
	return s.Database
}

func (s *DownloadDiagnosisRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DownloadDiagnosisRecordsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DownloadDiagnosisRecordsRequest) GetLang() *string {
	return s.Lang
}

func (s *DownloadDiagnosisRecordsRequest) GetMaxPeakMemory() *int64 {
	return s.MaxPeakMemory
}

func (s *DownloadDiagnosisRecordsRequest) GetMaxScanSize() *int64 {
	return s.MaxScanSize
}

func (s *DownloadDiagnosisRecordsRequest) GetMinPeakMemory() *int64 {
	return s.MinPeakMemory
}

func (s *DownloadDiagnosisRecordsRequest) GetMinScanSize() *int64 {
	return s.MinScanSize
}

func (s *DownloadDiagnosisRecordsRequest) GetQueryCondition() *string {
	return s.QueryCondition
}

func (s *DownloadDiagnosisRecordsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DownloadDiagnosisRecordsRequest) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *DownloadDiagnosisRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DownloadDiagnosisRecordsRequest) GetUserName() *string {
	return s.UserName
}

func (s *DownloadDiagnosisRecordsRequest) SetClientIp(v string) *DownloadDiagnosisRecordsRequest {
	s.ClientIp = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetDBClusterId(v string) *DownloadDiagnosisRecordsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetDatabase(v string) *DownloadDiagnosisRecordsRequest {
	s.Database = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetEndTime(v string) *DownloadDiagnosisRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetKeyword(v string) *DownloadDiagnosisRecordsRequest {
	s.Keyword = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetLang(v string) *DownloadDiagnosisRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetMaxPeakMemory(v int64) *DownloadDiagnosisRecordsRequest {
	s.MaxPeakMemory = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetMaxScanSize(v int64) *DownloadDiagnosisRecordsRequest {
	s.MaxScanSize = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetMinPeakMemory(v int64) *DownloadDiagnosisRecordsRequest {
	s.MinPeakMemory = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetMinScanSize(v int64) *DownloadDiagnosisRecordsRequest {
	s.MinScanSize = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetQueryCondition(v string) *DownloadDiagnosisRecordsRequest {
	s.QueryCondition = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetRegionId(v string) *DownloadDiagnosisRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetResourceGroup(v string) *DownloadDiagnosisRecordsRequest {
	s.ResourceGroup = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetStartTime(v string) *DownloadDiagnosisRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetUserName(v string) *DownloadDiagnosisRecordsRequest {
	s.UserName = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) Validate() error {
	return dara.Validate(s)
}
