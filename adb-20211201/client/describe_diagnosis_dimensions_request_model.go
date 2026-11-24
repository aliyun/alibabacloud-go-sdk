// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosisDimensionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDiagnosisDimensionsRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeDiagnosisDimensionsRequest
	GetEndTime() *string
	SetLang(v string) *DescribeDiagnosisDimensionsRequest
	GetLang() *string
	SetQueryCondition(v string) *DescribeDiagnosisDimensionsRequest
	GetQueryCondition() *string
	SetRegionId(v string) *DescribeDiagnosisDimensionsRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeDiagnosisDimensionsRequest
	GetStartTime() *string
}

type DescribeDiagnosisDimensionsRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (3.0) clusters in the specified region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bt6u59zcmd945****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. Specify the time as a UNIX timestamp in milliseconds.
	//
	// > 	- The end time must be later than the start time.
	//
	// > 	- The interval between the start time and the end time cannot exceed 24 hours.
	//
	// example:
	//
	// 1625220213000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language configuration. Valid values:
	//
	// - **zh-CN**: Simplified Chinese (default).
	//
	// - **en-US**: English.
	//
	// - **ja**: Japanese.
	//
	// example:
	//
	// zh-CN
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The search conditions for SQL queries. The value is a JSON string that contains fields such as `Type`, `Value`, `Min`, and `Max`. The `Type` field specifies the query dimension. Valid values: `maxCost`, `status`, and `cost`. The `Value`, `Min`, or `Max` field specifies the query range for the dimension. Valid values:
	//
	// - `{"Type":"maxCost","Value":"100"}`: queries the details of the top 100 SQL statements with the longest execution duration. Currently, you can only set `Value` to 100.
	//
	// - `{"Type":"status","Value":"finished"}`: queries the details of completed SQL statements. You can also set `Value` to `running` or `failed` to query the details of running or failed SQL statements.
	//
	// - `{"Type":"cost","Min":"10","Max":"200"}`: queries the details of SQL statements whose execution duration is between 10 milliseconds and 200 milliseconds. You can also customize the minimum and maximum execution duration in milliseconds.
	//
	// example:
	//
	// {"Type":"maxCost","Value":"100"}
	QueryCondition *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/454314.html) operation to query the region ID of the specified AnalyticDB for MySQL Data Lakehouse Edition (3.0) cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time of the query. Specify the time as a UNIX timestamp in milliseconds.
	//
	// > You can only query data from the last 14 days. If you query data from more than 14 days ago, an empty result is returned.
	//
	// example:
	//
	// 1625220210000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDiagnosisDimensionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisDimensionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisDimensionsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDiagnosisDimensionsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDiagnosisDimensionsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDiagnosisDimensionsRequest) GetQueryCondition() *string {
	return s.QueryCondition
}

func (s *DescribeDiagnosisDimensionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDiagnosisDimensionsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDiagnosisDimensionsRequest) SetDBClusterId(v string) *DescribeDiagnosisDimensionsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDiagnosisDimensionsRequest) SetEndTime(v string) *DescribeDiagnosisDimensionsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDiagnosisDimensionsRequest) SetLang(v string) *DescribeDiagnosisDimensionsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDiagnosisDimensionsRequest) SetQueryCondition(v string) *DescribeDiagnosisDimensionsRequest {
	s.QueryCondition = &v
	return s
}

func (s *DescribeDiagnosisDimensionsRequest) SetRegionId(v string) *DescribeDiagnosisDimensionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiagnosisDimensionsRequest) SetStartTime(v string) *DescribeDiagnosisDimensionsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosisDimensionsRequest) Validate() error {
	return dara.Validate(s)
}
