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
	// example:
	//
	// 192.168.XX.XX
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1j7******78j8i
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// 2025-09-25T01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// [{\\"Field\\":\\"ExecuteTime\\",\\"Type\\":\\"Desc\\"}]
	Order        *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 999f2439-6b10-xxxx-a5d3-daf3b35c****
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// example:
	//
	// user
	ProxyUser *string `json:"ProxyUser,omitempty" xml:"ProxyUser,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// test_job
	ResourceGroupName    *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// test_table_name
	SQLText *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	// example:
	//
	// 2025-09-25T00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Statement ID。
	//
	// example:
	//
	// fbd22066-1c03-xxxx-aa16-6ae28288****
	StatementId *string `json:"StatementId,omitempty" xml:"StatementId,omitempty"`
	// example:
	//
	// SQL_EDITOR
	StatementSource *string `json:"StatementSource,omitempty" xml:"StatementSource,omitempty"`
	// example:
	//
	// finish
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 3000
	TotalTime *string `json:"TotalTime,omitempty" xml:"TotalTime,omitempty"`
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
