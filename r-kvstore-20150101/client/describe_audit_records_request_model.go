// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *DescribeAuditRecordsRequest
	GetAccountName() *string
	SetDatabaseName(v string) *DescribeAuditRecordsRequest
	GetDatabaseName() *string
	SetEndTime(v string) *DescribeAuditRecordsRequest
	GetEndTime() *string
	SetHostAddress(v string) *DescribeAuditRecordsRequest
	GetHostAddress() *string
	SetInstanceId(v string) *DescribeAuditRecordsRequest
	GetInstanceId() *string
	SetNodeId(v string) *DescribeAuditRecordsRequest
	GetNodeId() *string
	SetOwnerAccount(v string) *DescribeAuditRecordsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAuditRecordsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeAuditRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAuditRecordsRequest
	GetPageSize() *int32
	SetQueryKeywords(v string) *DescribeAuditRecordsRequest
	GetQueryKeywords() *string
	SetResourceOwnerAccount(v string) *DescribeAuditRecordsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAuditRecordsRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeAuditRecordsRequest
	GetSecurityToken() *string
	SetStartTime(v string) *DescribeAuditRecordsRequest
	GetStartTime() *string
}

type DescribeAuditRecordsRequest struct {
	// The name of the account. If you do not specify this parameter, the audit logs of all accounts are queried.
	//
	// example:
	//
	// demo
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The database in the instance. If you do not specify this parameter, the audit logs of all databases are queried. Valid values: 0 to 255. For example, 0 indicates database 0.
	//
	// example:
	//
	// 0
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The end time of the query. The end time must be later than the start time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time must be in UTC.
	//
	// > Because audit logs contain a large number of entries, specify a time range of up to 10 minutes and a maximum of 1 day.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-03-25T12:10:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IP address of the client. If you do not specify this parameter, the audit logs of all IP addresses are queried.
	//
	// example:
	//
	// 127.0.0.1
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the node in the instance. Specify this parameter to query the monitoring information of a specific node.
	//
	// > 	- This parameter is available only when the instance uses the read/write splitting or cluster architecture.
	//
	// > 	- You can call the [DescribeLogicInstanceTopology](https://help.aliyun.com/document_detail/473786.html) operation to query the node ID.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****-db-0
	NodeId       *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keyword used to query audit logs based on a specific command. By default, the audit logs of all executed commands are queried.
	//
	// > You can specify only one keyword.
	//
	// example:
	//
	// maxclients
	QueryKeywords        *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The start time of the query. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-03-24T12:10:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAuditRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuditRecordsRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeAuditRecordsRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeAuditRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeAuditRecordsRequest) GetHostAddress() *string {
	return s.HostAddress
}

func (s *DescribeAuditRecordsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAuditRecordsRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeAuditRecordsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAuditRecordsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAuditRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAuditRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAuditRecordsRequest) GetQueryKeywords() *string {
	return s.QueryKeywords
}

func (s *DescribeAuditRecordsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAuditRecordsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAuditRecordsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeAuditRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeAuditRecordsRequest) SetAccountName(v string) *DescribeAuditRecordsRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetDatabaseName(v string) *DescribeAuditRecordsRequest {
	s.DatabaseName = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetEndTime(v string) *DescribeAuditRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetHostAddress(v string) *DescribeAuditRecordsRequest {
	s.HostAddress = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetInstanceId(v string) *DescribeAuditRecordsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetNodeId(v string) *DescribeAuditRecordsRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetOwnerAccount(v string) *DescribeAuditRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetOwnerId(v int64) *DescribeAuditRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetPageNumber(v int32) *DescribeAuditRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetPageSize(v int32) *DescribeAuditRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetQueryKeywords(v string) *DescribeAuditRecordsRequest {
	s.QueryKeywords = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetResourceOwnerAccount(v string) *DescribeAuditRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetResourceOwnerId(v int64) *DescribeAuditRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetSecurityToken(v string) *DescribeAuditRecordsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetStartTime(v string) *DescribeAuditRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAuditRecordsRequest) Validate() error {
	return dara.Validate(s)
}
