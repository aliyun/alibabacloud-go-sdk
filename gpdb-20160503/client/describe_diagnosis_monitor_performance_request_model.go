// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosisMonitorPerformanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDiagnosisMonitorPerformanceRequest
	GetDBInstanceId() *string
	SetDatabase(v string) *DescribeDiagnosisMonitorPerformanceRequest
	GetDatabase() *string
	SetEndTime(v string) *DescribeDiagnosisMonitorPerformanceRequest
	GetEndTime() *string
	SetQueryCondition(v string) *DescribeDiagnosisMonitorPerformanceRequest
	GetQueryCondition() *string
	SetStartTime(v string) *DescribeDiagnosisMonitorPerformanceRequest
	GetStartTime() *string
	SetUser(v string) *DescribeDiagnosisMonitorPerformanceRequest
	GetUser() *string
}

type DescribeDiagnosisMonitorPerformanceRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// adbtest
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The end of the time range to query. Specify the time in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC. The end time must be later than the start time.
	//
	// example:
	//
	// 2022-05-07T07:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The filter condition on queries. Specify the value in the JSON format. Valid values:
	//
	// 	- `{"Type":"maxCost", "Value":"100"}`: filters the top 100 queries that are the most time-consuming.
	//
	// 	- `{"Type":"status","Value":"finished"}`: filters completed queries.
	//
	// 	- `{"Type":"status","Value":"running"}`: filters running queries.
	//
	// 	- `{"Type":"cost","Min":"30","Max":"50"}`: filters the queries that consume 30 milliseconds or more and less than 50 milliseconds. You can customize a filter condition by setting **Min*	- and **Max**.
	//
	//     	- If only **Min*	- is specified, the queries that consume a period of time that is greater than or equal to the Min value are filtered.
	//
	//     	- If only **Max*	- is specified, the queries that consume a period of time that is less than the Max value are filtered.
	//
	//     	- If both **Min*	- and **Max*	- are specified, the queries that consume a period of time that is greater than or equal to the **Min*	- value and less than the **Max*	- value are filtered.
	//
	// example:
	//
	// {"Type":"maxCost", "Value":"100"}
	QueryCondition *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	// The beginning of the time range to query. Specify the time in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2022-05-07T06:59Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the database account.
	//
	// example:
	//
	// adbpguser
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeDiagnosisMonitorPerformanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisMonitorPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) GetDatabase() *string {
	return s.Database
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) GetQueryCondition() *string {
	return s.QueryCondition
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) GetUser() *string {
	return s.User
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetDBInstanceId(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetDatabase(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.Database = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetEndTime(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetQueryCondition(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.QueryCondition = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetStartTime(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetUser(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.User = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) Validate() error {
	return dara.Validate(s)
}
