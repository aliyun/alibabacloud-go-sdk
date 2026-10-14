// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSparkSQLDiagnosisListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeSparkSQLDiagnosisListRequest
	GetDBClusterId() *string
	SetMaxStartTime(v string) *DescribeSparkSQLDiagnosisListRequest
	GetMaxStartTime() *string
	SetMinStartTime(v string) *DescribeSparkSQLDiagnosisListRequest
	GetMinStartTime() *string
	SetOrder(v string) *DescribeSparkSQLDiagnosisListRequest
	GetOrder() *string
	SetPageNumber(v int32) *DescribeSparkSQLDiagnosisListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSparkSQLDiagnosisListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSparkSQLDiagnosisListRequest
	GetRegionId() *string
	SetStatementId(v int64) *DescribeSparkSQLDiagnosisListRequest
	GetStatementId() *int64
}

type DescribeSparkSQLDiagnosisListRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the details of all AnalyticDB for MySQL Data Lakehouse Edition (3.0) clusters in a specific region, including the cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-2zez35ww415xjwk5
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The maximum start time. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2024-11-28 23:00:00
	MaxStartTime *string `json:"MaxStartTime,omitempty" xml:"MaxStartTime,omitempty"`
	// The minimum start time. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2024-11-28 22:00:00
	MinStartTime *string `json:"MinStartTime,omitempty" xml:"MinStartTime,omitempty"`
	// Sorts the query results by a specified field. The value must be a JSON string. Example: `[{"Field":"MaxExclusiveTime","Type":"Asc"}]`. Parameters in the JSON string:
	//
	// - `Field`: the sort field. Valid values:
	//
	//     - `MaxExclusiveTime`: the maximum execution time.
	//
	//     - `PeakMemory`: the peak memory usage.
	//
	//     - `QueryStartTime`: the query start time.
	//
	//     - `QueryWallclockTime`: the query execution time.
	//
	// - `Type`: the sort order. Valid values:
	//
	//     - `Asc`: ascending order.
	//
	//     - `Desc`: descending order.
	//
	// > - If this parameter is not specified, the results are sorted in ascending order by the `MaxExclusiveTime` field by default.
	//
	// example:
	//
	// [{\\"Field\\":\\"QueryStartTime\\",\\"Type\\":\\"Desc\\"}]
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
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
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the regions and zones supported by AnalyticDB for MySQL, including region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The unique ID of the code block in the Spark job.
	//
	// example:
	//
	// 123
	StatementId *int64 `json:"StatementId,omitempty" xml:"StatementId,omitempty"`
}

func (s DescribeSparkSQLDiagnosisListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSparkSQLDiagnosisListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSparkSQLDiagnosisListRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSparkSQLDiagnosisListRequest) GetMaxStartTime() *string {
	return s.MaxStartTime
}

func (s *DescribeSparkSQLDiagnosisListRequest) GetMinStartTime() *string {
	return s.MinStartTime
}

func (s *DescribeSparkSQLDiagnosisListRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeSparkSQLDiagnosisListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSparkSQLDiagnosisListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSparkSQLDiagnosisListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSparkSQLDiagnosisListRequest) GetStatementId() *int64 {
	return s.StatementId
}

func (s *DescribeSparkSQLDiagnosisListRequest) SetDBClusterId(v string) *DescribeSparkSQLDiagnosisListRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisListRequest) SetMaxStartTime(v string) *DescribeSparkSQLDiagnosisListRequest {
	s.MaxStartTime = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisListRequest) SetMinStartTime(v string) *DescribeSparkSQLDiagnosisListRequest {
	s.MinStartTime = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisListRequest) SetOrder(v string) *DescribeSparkSQLDiagnosisListRequest {
	s.Order = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisListRequest) SetPageNumber(v int32) *DescribeSparkSQLDiagnosisListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisListRequest) SetPageSize(v int32) *DescribeSparkSQLDiagnosisListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisListRequest) SetRegionId(v string) *DescribeSparkSQLDiagnosisListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisListRequest) SetStatementId(v int64) *DescribeSparkSQLDiagnosisListRequest {
	s.StatementId = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisListRequest) Validate() error {
	return dara.Validate(s)
}
