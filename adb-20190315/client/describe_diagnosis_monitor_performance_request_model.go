// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosisMonitorPerformanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDiagnosisMonitorPerformanceRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeDiagnosisMonitorPerformanceRequest
	GetEndTime() *string
	SetLang(v string) *DescribeDiagnosisMonitorPerformanceRequest
	GetLang() *string
	SetQueryCondition(v string) *DescribeDiagnosisMonitorPerformanceRequest
	GetQueryCondition() *string
	SetRegionId(v string) *DescribeDiagnosisMonitorPerformanceRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeDiagnosisMonitorPerformanceRequest
	GetStartTime() *string
}

type DescribeDiagnosisMonitorPerformanceRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the cluster IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp108q1py5r78****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1671687948000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of file titles and error messages. Valid values:
	//
	// 	- **zh*	- (default): simplified Chinese.
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
	// The query conditions for SQL statements, which can be a combination of the `Type` and `Value` fields or a combination of the Type, `Min`, and `Max` fields. Specify the conditions in the JSON format. `Type` specifies the query dimension. Valid values for Type: `maxCost`, `status`, and `cost`. `Value`, `Min`, or `Max` specifies the query range for the dimension. Valid values:
	//
	// 	- `{"Type":"maxCost","Value":"100"}`: queries the top 100 most time-consuming SQL statements. Set `Value` to 100.
	//
	// 	- `{"Type":"status","Value":"finished"}`: queries the executed SQL statements, **including successful and failed SQL statements**. You can set `Value` to `running` to query the SQL statements that are being executed. You can also set Value to `failed` to query the SQL statements that failed to be executed.
	//
	// 	- `{"Type":"cost","Min":"10","Max":"200"}`: queries the SQL statements whose execution duration is in the range of 10 to 200 milliseconds. You can also specify custom values for the Min and Max fields.
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"Type\\":\\"maxCost\\",\\"Value\\":\\"100\\"}
	QueryCondition *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1671684348000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDiagnosisMonitorPerformanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisMonitorPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) GetQueryCondition() *string {
	return s.QueryCondition
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetDBClusterId(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetEndTime(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetLang(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetQueryCondition(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.QueryCondition = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetRegionId(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetStartTime(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) Validate() error {
	return dara.Validate(s)
}
