// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeSlowLogRecordsRequest
	GetDBClusterId() *string
	SetDBName(v string) *DescribeSlowLogRecordsRequest
	GetDBName() *string
	SetEndTime(v string) *DescribeSlowLogRecordsRequest
	GetEndTime() *string
	SetOrder(v string) *DescribeSlowLogRecordsRequest
	GetOrder() *string
	SetOwnerAccount(v string) *DescribeSlowLogRecordsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSlowLogRecordsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeSlowLogRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSlowLogRecordsRequest
	GetPageSize() *int32
	SetProcessID(v string) *DescribeSlowLogRecordsRequest
	GetProcessID() *string
	SetRange(v string) *DescribeSlowLogRecordsRequest
	GetRange() *string
	SetResourceOwnerAccount(v string) *DescribeSlowLogRecordsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSlowLogRecordsRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeSlowLogRecordsRequest
	GetStartTime() *string
	SetState(v string) *DescribeSlowLogRecordsRequest
	GetState() *string
}

type DescribeSlowLogRecordsRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the cluster IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a specific region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1rqvm70uh2****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// adb_demo
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time must be in UTC.
	//
	// >  The end time must be later than the start time. The specified time range must be less than seven days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-05-27T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The order in which to sort the retrieved entries by field. Specify this parameter in the JSON format. The value is an ordered array that uses the order of the input array and contains `Field` and `Type`. Example: `[{"Field":"ExecutionStartTime","Type":"Desc"},{"Field":"ScanRows","Type":"Asc"}]`.
	//
	// 	- `Field`: the field that is used to sort the retrieved entries. Valid values:
	//
	//     	- **HostAddress**: the IP address of the client that is used to connect to the database.
	//
	//     	- **UserName**: the username.
	//
	//     	- **ExecutionStartTime**: the start time of the query execution.
	//
	//     	- **QueryTime**: the amount of time consumed to execute the SQL statement.
	//
	//     	- **PeakMemoryUsage**: the maximum memory usage when the SQL statement is executed.
	//
	//     	- **ScanRows**: the number of rows to be scanned from a data source in the task.
	//
	//     	- **ScanSize**: the amount of data to be scanned.
	//
	//     	- **ScanTime**: the total amount of time consumed to scan data.
	//
	//     	- **PlanningTime**: the amount of time consumed to generate execution plans.
	//
	//     	- **WallTime**: the accumulated CPU Time values of all operators in the query on each node.
	//
	//     	- **ProcessID**: the ID of the process.
	//
	// 	- `Type`: the sorting type of the retrieved entries. Valid values:
	//
	//     	- **Desc**: descending order
	//
	//     	- **Asc**: ascending order
	//
	// example:
	//
	// [{"Field":"ExecutionStartTime","Type":"Desc"},{"Field":"ScanRows","Type":"Asc"}]
	Order        *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: **30**, **50**, and **100**. Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the process.
	//
	// example:
	//
	// 2021052716044317201616624903453******
	ProcessID *string `json:"ProcessID,omitempty" xml:"ProcessID,omitempty"`
	// The range conditions used to filter specified fields, including `Max` and `Min`. Specify this parameter in the JSON format. Example: `[{"Field":"ScanSize","Min":"1000000","Max":"10000000"},{"Field":"QueryTime","Min":"1000","Max":"10000"}]`.
	//
	// `Field`: the field to be filtered. Valid values:
	//
	// 	- **ScanSize**: the amount of data to be scanned. Unit: KB.
	//
	// 	- **QueryTime**: the amount of time consumed to execute the statement. Unit: milliseconds.
	//
	// 	- **PeakMemoryUsage**: the maximum memory usage when the SQL statement is executed. Unit: KB.
	//
	// >  `Min` indicates the minimum value of the query range (left operand). `Max` indicates the maximum value of the query range (right operand). Max and Min are both of the String type.
	//
	// example:
	//
	// [{"Field":"ScanSize","Min":"1000000","Max":"10000000"},{"Field":"QueryTime","Min":"1000","Max":"10000"}]
	Range                *string `json:"Range,omitempty" xml:"Range,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-05-20T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the query. Valid values:
	//
	// 	- **Successed**: successful
	//
	// 	- **Failed**: failed
	//
	// example:
	//
	// Successed
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeSlowLogRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSlowLogRecordsRequest) GetDBName() *string {
	return s.DBName
}

func (s *DescribeSlowLogRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSlowLogRecordsRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeSlowLogRecordsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSlowLogRecordsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSlowLogRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSlowLogRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSlowLogRecordsRequest) GetProcessID() *string {
	return s.ProcessID
}

func (s *DescribeSlowLogRecordsRequest) GetRange() *string {
	return s.Range
}

func (s *DescribeSlowLogRecordsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSlowLogRecordsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSlowLogRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSlowLogRecordsRequest) GetState() *string {
	return s.State
}

func (s *DescribeSlowLogRecordsRequest) SetDBClusterId(v string) *DescribeSlowLogRecordsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetDBName(v string) *DescribeSlowLogRecordsRequest {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetEndTime(v string) *DescribeSlowLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetOrder(v string) *DescribeSlowLogRecordsRequest {
	s.Order = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetOwnerAccount(v string) *DescribeSlowLogRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetOwnerId(v int64) *DescribeSlowLogRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetPageNumber(v int32) *DescribeSlowLogRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetPageSize(v int32) *DescribeSlowLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetProcessID(v string) *DescribeSlowLogRecordsRequest {
	s.ProcessID = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetRange(v string) *DescribeSlowLogRecordsRequest {
	s.Range = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetResourceOwnerAccount(v string) *DescribeSlowLogRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetResourceOwnerId(v int64) *DescribeSlowLogRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetStartTime(v string) *DescribeSlowLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetState(v string) *DescribeSlowLogRecordsRequest {
	s.State = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) Validate() error {
	return dara.Validate(s)
}
