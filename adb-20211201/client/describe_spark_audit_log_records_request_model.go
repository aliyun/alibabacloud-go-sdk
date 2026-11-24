// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSparkAuditLogRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientIp(v string) *DescribeSparkAuditLogRecordsRequest
	GetClientIp() *string
	SetDBClusterId(v string) *DescribeSparkAuditLogRecordsRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeSparkAuditLogRecordsRequest
	GetEndTime() *string
	SetOrder(v string) *DescribeSparkAuditLogRecordsRequest
	GetOrder() *string
	SetOwnerAccount(v string) *DescribeSparkAuditLogRecordsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSparkAuditLogRecordsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeSparkAuditLogRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSparkAuditLogRecordsRequest
	GetPageSize() *int32
	SetProcessId(v string) *DescribeSparkAuditLogRecordsRequest
	GetProcessId() *string
	SetProxyUser(v string) *DescribeSparkAuditLogRecordsRequest
	GetProxyUser() *string
	SetRegionId(v string) *DescribeSparkAuditLogRecordsRequest
	GetRegionId() *string
	SetResourceGroupName(v string) *DescribeSparkAuditLogRecordsRequest
	GetResourceGroupName() *string
	SetResourceOwnerAccount(v string) *DescribeSparkAuditLogRecordsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSparkAuditLogRecordsRequest
	GetResourceOwnerId() *int64
	SetSQLText(v string) *DescribeSparkAuditLogRecordsRequest
	GetSQLText() *string
	SetStartTime(v string) *DescribeSparkAuditLogRecordsRequest
	GetStartTime() *string
	SetStatementId(v string) *DescribeSparkAuditLogRecordsRequest
	GetStatementId() *string
	SetStatementSource(v string) *DescribeSparkAuditLogRecordsRequest
	GetStatementSource() *string
	SetStatus(v string) *DescribeSparkAuditLogRecordsRequest
	GetStatus() *string
	SetTotalTime(v string) *DescribeSparkAuditLogRecordsRequest
	GetTotalTime() *string
	SetUser(v string) *DescribeSparkAuditLogRecordsRequest
	GetUser() *string
}

type DescribeSparkAuditLogRecordsRequest struct {
	// The source IP address.
	//
	// example:
	//
	// 192.168.XX.XX
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The cluster ID.
	//
	// > - You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to query the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1j7******78j8i
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. The end time must be later than the start time, and the interval between the start time and the end time must be less than one day.
	//
	// Format: yyyy-MM-ddTHH:mmZ (UTC).
	//
	// example:
	//
	// 2025-09-25T01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The sorting rule for SQL statements based on specified fields. The value is a JSON array that specifies composite sorting in the order of the input array. The array contains the `Field` and `Type` fields. Example: `[{"Field":"TotalTime", "Type": "desc" }]`.
	//
	// - `Field`: the name of the field to sort by. Valid values:
	//
	//     - `ResourceGroupName`: the name of the resource group.
	//
	//     - `Status`: the execution status of the SQL statement.
	//
	//     - `User`: the username that executed the SQL statement.
	//
	//     - `ExecuteTime`: the start time of the SQL execution.
	//
	//     - `TotalTime`: the execution duration of the SQL statement.
	//
	//     - `ProcessId`: the query ID.
	//
	//     - `ClientIp`: the source IP address.
	//
	//     - `StatementSource`: the source from which the query is initiated.
	//
	// - `Type`: the sorting type. Valid values (case-insensitive):
	//
	//     - `Desc`: descending order.
	//
	//     - `Asc`: ascending order.
	//
	// example:
	//
	// [{\\"Field\\":\\"ExecuteTime\\",\\"Type\\":\\"Desc\\"}]
	Order        *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// The query ID.
	//
	// example:
	//
	// 999f2439-6b10-xxxx-a5d3-daf3b35c****
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated.
	//
	// example:
	//
	// user
	ProxyUser *string `json:"ProxyUser,omitempty" xml:"ProxyUser,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/612293.html) operation to view the regions and zones supported by AnalyticDB for MySQL, including region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the resource group.
	//
	// > You can call the [DescribeDBResourceGroup](https://help.aliyun.com/document_detail/612410.html) operation to view the resource group ID of the cluster.
	//
	// example:
	//
	// test_job
	ResourceGroupName    *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The keyword in the SQL statement.
	//
	// example:
	//
	// test_table_name
	SQLText *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	// The start time of the query. Format: <i>yyyy-MM-ddTHH:mmZ</i> (UTC).
	//
	// > Set the start time to a point in time within the last 30 days.
	//
	// example:
	//
	// 2025-09-25T00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The statement ID.
	//
	// example:
	//
	// fbd22066-1c03-xxxx-aa16-6ae28288****
	StatementId *string `json:"StatementId,omitempty" xml:"StatementId,omitempty"`
	// The source from which the query is initiated.
	//
	// example:
	//
	// SQL_EDITOR
	StatementSource *string `json:"StatementSource,omitempty" xml:"StatementSource,omitempty"`
	// The execution status of the SQL statement.
	//
	// example:
	//
	// finish
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The execution duration of the SQL statement. Unit: milliseconds.
	//
	// example:
	//
	// 3000
	TotalTime *string `json:"TotalTime,omitempty" xml:"TotalTime,omitempty"`
	// The username that executed the SQL statement.
	//
	// example:
	//
	// test_user
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeSparkAuditLogRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSparkAuditLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSparkAuditLogRecordsRequest) GetClientIp() *string {
	return s.ClientIp
}

func (s *DescribeSparkAuditLogRecordsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSparkAuditLogRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSparkAuditLogRecordsRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeSparkAuditLogRecordsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSparkAuditLogRecordsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSparkAuditLogRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSparkAuditLogRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSparkAuditLogRecordsRequest) GetProcessId() *string {
	return s.ProcessId
}

func (s *DescribeSparkAuditLogRecordsRequest) GetProxyUser() *string {
	return s.ProxyUser
}

func (s *DescribeSparkAuditLogRecordsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSparkAuditLogRecordsRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *DescribeSparkAuditLogRecordsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSparkAuditLogRecordsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSparkAuditLogRecordsRequest) GetSQLText() *string {
	return s.SQLText
}

func (s *DescribeSparkAuditLogRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSparkAuditLogRecordsRequest) GetStatementId() *string {
	return s.StatementId
}

func (s *DescribeSparkAuditLogRecordsRequest) GetStatementSource() *string {
	return s.StatementSource
}

func (s *DescribeSparkAuditLogRecordsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeSparkAuditLogRecordsRequest) GetTotalTime() *string {
	return s.TotalTime
}

func (s *DescribeSparkAuditLogRecordsRequest) GetUser() *string {
	return s.User
}

func (s *DescribeSparkAuditLogRecordsRequest) SetClientIp(v string) *DescribeSparkAuditLogRecordsRequest {
	s.ClientIp = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsRequest) SetDBClusterId(v string) *DescribeSparkAuditLogRecordsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsRequest) SetEndTime(v string) *DescribeSparkAuditLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsRequest) SetOrder(v string) *DescribeSparkAuditLogRecordsRequest {
	s.Order = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsRequest) SetOwnerAccount(v string) *DescribeSparkAuditLogRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsRequest) SetOwnerId(v int64) *DescribeSparkAuditLogRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsRequest) SetPageNumber(v int32) *DescribeSparkAuditLogRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsRequest) SetPageSize(v int32) *DescribeSparkAuditLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsRequest) SetProcessId(v string) *DescribeSparkAuditLogRecordsRequest {
	s.ProcessId = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsRequest) SetProxyUser(v string) *DescribeSparkAuditLogRecordsRequest {
	s.ProxyUser = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsRequest) SetRegionId(v string) *DescribeSparkAuditLogRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsRequest) SetResourceGroupName(v string) *DescribeSparkAuditLogRecordsRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsRequest) SetResourceOwnerAccount(v string) *DescribeSparkAuditLogRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsRequest) SetResourceOwnerId(v int64) *DescribeSparkAuditLogRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsRequest) SetSQLText(v string) *DescribeSparkAuditLogRecordsRequest {
	s.SQLText = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsRequest) SetStartTime(v string) *DescribeSparkAuditLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsRequest) SetStatementId(v string) *DescribeSparkAuditLogRecordsRequest {
	s.StatementId = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsRequest) SetStatementSource(v string) *DescribeSparkAuditLogRecordsRequest {
	s.StatementSource = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsRequest) SetStatus(v string) *DescribeSparkAuditLogRecordsRequest {
	s.Status = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsRequest) SetTotalTime(v string) *DescribeSparkAuditLogRecordsRequest {
	s.TotalTime = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsRequest) SetUser(v string) *DescribeSparkAuditLogRecordsRequest {
	s.User = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsRequest) Validate() error {
	return dara.Validate(s)
}
