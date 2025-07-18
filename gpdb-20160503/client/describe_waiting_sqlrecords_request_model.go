// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWaitingSQLRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeWaitingSQLRecordsRequest
	GetDBInstanceId() *string
	SetDatabase(v string) *DescribeWaitingSQLRecordsRequest
	GetDatabase() *string
	SetEndTime(v string) *DescribeWaitingSQLRecordsRequest
	GetEndTime() *string
	SetKeyword(v string) *DescribeWaitingSQLRecordsRequest
	GetKeyword() *string
	SetOrder(v string) *DescribeWaitingSQLRecordsRequest
	GetOrder() *string
	SetPageNumber(v int32) *DescribeWaitingSQLRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeWaitingSQLRecordsRequest
	GetPageSize() *int32
	SetQueryCondition(v string) *DescribeWaitingSQLRecordsRequest
	GetQueryCondition() *string
	SetStartTime(v string) *DescribeWaitingSQLRecordsRequest
	GetStartTime() *string
	SetUser(v string) *DescribeWaitingSQLRecordsRequest
	GetUser() *string
}

type DescribeWaitingSQLRecordsRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the instance IDs of all AnalyticDB for PostgreSQL instances in a specific region.
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
	// test
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC. The end time must be later than the start time.
	//
	// If this parameter is not specified, all lock diagnostics records that are generated after the query start time are returned. If the query start time is not specified either, all lock diagnostics records are returned.
	//
	// example:
	//
	// 2022-08-20T07:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The keyword used to filter queries.
	//
	// example:
	//
	// table
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The field used to sort lock diagnostics records and the sorting order.
	//
	// Default value: `{"Field":"StartTime","Type":"Desc"}`, which indicates that lock diagnostics records are sorted by the start time in descending order. No other values are supported.
	//
	// example:
	//
	// {"Field":"StartTime","Type":"Desc"}
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
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
	// The filter condition on queries. Valid values:
	//
	// 	- `{"Type":"maxCost","Value":"10"}`: filters the top 10 longest-waiting queries.
	//
	// 	- `{"Type":"status","Value":"LockWaiting"}`: filters lock-waiting queries.
	//
	// 	- `{"Type":"status","Value":"ResourceWaiting"}`: filters resource-waiting queries.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Type":"maxCost","Value":"10"}
	QueryCondition *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// If this parameter is not specified, all lock diagnostics records that are generated before the query end time are returned. If the query end time is not specified either, all lock diagnostics records are returned.
	//
	// example:
	//
	// 2022-08-15T06:59Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the database account. If this parameter is not specified, the lock diagnostics records of all database accounts are queried.
	//
	// example:
	//
	// testUser
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeWaitingSQLRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWaitingSQLRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWaitingSQLRecordsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeWaitingSQLRecordsRequest) GetDatabase() *string {
	return s.Database
}

func (s *DescribeWaitingSQLRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeWaitingSQLRecordsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeWaitingSQLRecordsRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeWaitingSQLRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeWaitingSQLRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWaitingSQLRecordsRequest) GetQueryCondition() *string {
	return s.QueryCondition
}

func (s *DescribeWaitingSQLRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeWaitingSQLRecordsRequest) GetUser() *string {
	return s.User
}

func (s *DescribeWaitingSQLRecordsRequest) SetDBInstanceId(v string) *DescribeWaitingSQLRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeWaitingSQLRecordsRequest) SetDatabase(v string) *DescribeWaitingSQLRecordsRequest {
	s.Database = &v
	return s
}

func (s *DescribeWaitingSQLRecordsRequest) SetEndTime(v string) *DescribeWaitingSQLRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeWaitingSQLRecordsRequest) SetKeyword(v string) *DescribeWaitingSQLRecordsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeWaitingSQLRecordsRequest) SetOrder(v string) *DescribeWaitingSQLRecordsRequest {
	s.Order = &v
	return s
}

func (s *DescribeWaitingSQLRecordsRequest) SetPageNumber(v int32) *DescribeWaitingSQLRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeWaitingSQLRecordsRequest) SetPageSize(v int32) *DescribeWaitingSQLRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWaitingSQLRecordsRequest) SetQueryCondition(v string) *DescribeWaitingSQLRecordsRequest {
	s.QueryCondition = &v
	return s
}

func (s *DescribeWaitingSQLRecordsRequest) SetStartTime(v string) *DescribeWaitingSQLRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeWaitingSQLRecordsRequest) SetUser(v string) *DescribeWaitingSQLRecordsRequest {
	s.User = &v
	return s
}

func (s *DescribeWaitingSQLRecordsRequest) Validate() error {
	return dara.Validate(s)
}
