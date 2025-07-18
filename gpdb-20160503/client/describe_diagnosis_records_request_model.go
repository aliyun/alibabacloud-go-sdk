// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosisRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDiagnosisRecordsRequest
	GetDBInstanceId() *string
	SetDatabase(v string) *DescribeDiagnosisRecordsRequest
	GetDatabase() *string
	SetEndTime(v string) *DescribeDiagnosisRecordsRequest
	GetEndTime() *string
	SetKeyword(v string) *DescribeDiagnosisRecordsRequest
	GetKeyword() *string
	SetOrder(v string) *DescribeDiagnosisRecordsRequest
	GetOrder() *string
	SetPageNumber(v int32) *DescribeDiagnosisRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDiagnosisRecordsRequest
	GetPageSize() *int32
	SetQueryCondition(v string) *DescribeDiagnosisRecordsRequest
	GetQueryCondition() *string
	SetStartTime(v string) *DescribeDiagnosisRecordsRequest
	GetStartTime() *string
	SetUser(v string) *DescribeDiagnosisRecordsRequest
	GetUser() *string
}

type DescribeDiagnosisRecordsRequest struct {
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
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
	// The keyword of the SQL statement.
	//
	// example:
	//
	// SELECT
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The order of fields in the console. You do not need to specify this parameter.
	//
	// example:
	//
	// null
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Pages start from page 1. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The filter condition on queries. Specify the value in the JSON format. Valid values:
	//
	// 	- `{"Type":"maxCost", "Value":"100"}`: filters the top 100 queries that are the most time-consuming.
	//
	// 	- `{"Type":"status","Value":"finished"}`: filters completed queries.
	//
	// 	- `{"Type":"status","Value":"running"}`: filters running queries.
	//
	// 	- `{"Type":"cost","Min":"30","Max":"50"}`: filters the queries that consume a period of 30 milliseconds to less than 50 milliseconds. You can customize a filter condition by setting **Min*	- and **Max**.
	//
	//     	- If only **Min*	- is specified, the queries that consume a period of time that is greater than the Min value are filtered.
	//
	//     	- If only **Max*	- is specified, the queries that consume a period of time that is less than the Max value are filtered.
	//
	//     	- If both **Min*	- and **Max*	- are specified, the queries that consume a period of time that is greater than or equal to the **Min*	- value and less than or equal to the **Max*	- value are filtered.
	//
	// example:
	//
	// { "Type":"maxCost", "Value":"100" }
	QueryCondition *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
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

func (s DescribeDiagnosisRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisRecordsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
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

func (s *DescribeDiagnosisRecordsRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeDiagnosisRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDiagnosisRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDiagnosisRecordsRequest) GetQueryCondition() *string {
	return s.QueryCondition
}

func (s *DescribeDiagnosisRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDiagnosisRecordsRequest) GetUser() *string {
	return s.User
}

func (s *DescribeDiagnosisRecordsRequest) SetDBInstanceId(v string) *DescribeDiagnosisRecordsRequest {
	s.DBInstanceId = &v
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

func (s *DescribeDiagnosisRecordsRequest) SetQueryCondition(v string) *DescribeDiagnosisRecordsRequest {
	s.QueryCondition = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetStartTime(v string) *DescribeDiagnosisRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetUser(v string) *DescribeDiagnosisRecordsRequest {
	s.User = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) Validate() error {
	return dara.Validate(s)
}
