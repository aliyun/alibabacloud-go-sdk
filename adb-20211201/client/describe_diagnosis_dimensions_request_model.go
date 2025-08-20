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
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bt6u59zcmd945****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
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
	// 1625220213000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language. Valid values:
	//
	// 	- **zh-CN*	- (default): simplified Chinese.
	//
	// 	- **en-US**: English.
	//
	// 	- **ja**: Japanese.
	//
	// example:
	//
	// zh-CN
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
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
	// {"Type":"maxCost","Value":"100"}
	QueryCondition *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/454314.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Set the time to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  You can query data only within the last 14 days.
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
