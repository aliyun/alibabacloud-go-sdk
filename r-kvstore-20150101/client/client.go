// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddShardingNodeRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ShardCount           *int32  `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
	ShardClass           *string `json:"ShardClass,omitempty" xml:"ShardClass,omitempty"`
	AutoPay              *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	CouponNo             *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	BusinessInfo         *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s AddShardingNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s AddShardingNodeRequest) GoString() string {
	return s.String()
}

func (s *AddShardingNodeRequest) SetSecurityToken(v string) *AddShardingNodeRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddShardingNodeRequest) SetOwnerId(v int64) *AddShardingNodeRequest {
	s.OwnerId = &v
	return s
}

func (s *AddShardingNodeRequest) SetResourceOwnerAccount(v string) *AddShardingNodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddShardingNodeRequest) SetResourceOwnerId(v int64) *AddShardingNodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddShardingNodeRequest) SetOwnerAccount(v string) *AddShardingNodeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddShardingNodeRequest) SetInstanceId(v string) *AddShardingNodeRequest {
	s.InstanceId = &v
	return s
}

func (s *AddShardingNodeRequest) SetShardCount(v int32) *AddShardingNodeRequest {
	s.ShardCount = &v
	return s
}

func (s *AddShardingNodeRequest) SetShardClass(v string) *AddShardingNodeRequest {
	s.ShardClass = &v
	return s
}

func (s *AddShardingNodeRequest) SetAutoPay(v bool) *AddShardingNodeRequest {
	s.AutoPay = &v
	return s
}

func (s *AddShardingNodeRequest) SetCouponNo(v string) *AddShardingNodeRequest {
	s.CouponNo = &v
	return s
}

func (s *AddShardingNodeRequest) SetBusinessInfo(v string) *AddShardingNodeRequest {
	s.BusinessInfo = &v
	return s
}

func (s *AddShardingNodeRequest) SetClientToken(v string) *AddShardingNodeRequest {
	s.ClientToken = &v
	return s
}

type AddShardingNodeResponseBody struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	NodeIds   []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	OrderId   *int64    `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s AddShardingNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddShardingNodeResponseBody) GoString() string {
	return s.String()
}

func (s *AddShardingNodeResponseBody) SetRequestId(v string) *AddShardingNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddShardingNodeResponseBody) SetNodeIds(v []*string) *AddShardingNodeResponseBody {
	s.NodeIds = v
	return s
}

func (s *AddShardingNodeResponseBody) SetOrderId(v int64) *AddShardingNodeResponseBody {
	s.OrderId = &v
	return s
}

type AddShardingNodeResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddShardingNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddShardingNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s AddShardingNodeResponse) GoString() string {
	return s.String()
}

func (s *AddShardingNodeResponse) SetHeaders(v map[string]*string) *AddShardingNodeResponse {
	s.Headers = v
	return s
}

func (s *AddShardingNodeResponse) SetBody(v *AddShardingNodeResponseBody) *AddShardingNodeResponse {
	s.Body = v
	return s
}

type AllocateDirectConnectionRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ConnectionString     *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	Port                 *string `json:"Port,omitempty" xml:"Port,omitempty"`
	OnlyAllocateIp       *bool   `json:"OnlyAllocateIp,omitempty" xml:"OnlyAllocateIp,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s AllocateDirectConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s AllocateDirectConnectionRequest) GoString() string {
	return s.String()
}

func (s *AllocateDirectConnectionRequest) SetSecurityToken(v string) *AllocateDirectConnectionRequest {
	s.SecurityToken = &v
	return s
}

func (s *AllocateDirectConnectionRequest) SetOwnerId(v int64) *AllocateDirectConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocateDirectConnectionRequest) SetResourceOwnerAccount(v string) *AllocateDirectConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocateDirectConnectionRequest) SetResourceOwnerId(v int64) *AllocateDirectConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AllocateDirectConnectionRequest) SetOwnerAccount(v string) *AllocateDirectConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AllocateDirectConnectionRequest) SetInstanceId(v string) *AllocateDirectConnectionRequest {
	s.InstanceId = &v
	return s
}

func (s *AllocateDirectConnectionRequest) SetConnectionString(v string) *AllocateDirectConnectionRequest {
	s.ConnectionString = &v
	return s
}

func (s *AllocateDirectConnectionRequest) SetPort(v string) *AllocateDirectConnectionRequest {
	s.Port = &v
	return s
}

func (s *AllocateDirectConnectionRequest) SetOnlyAllocateIp(v bool) *AllocateDirectConnectionRequest {
	s.OnlyAllocateIp = &v
	return s
}

func (s *AllocateDirectConnectionRequest) SetClientToken(v string) *AllocateDirectConnectionRequest {
	s.ClientToken = &v
	return s
}

type AllocateDirectConnectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateDirectConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AllocateDirectConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateDirectConnectionResponseBody) SetRequestId(v string) *AllocateDirectConnectionResponseBody {
	s.RequestId = &v
	return s
}

type AllocateDirectConnectionResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AllocateDirectConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AllocateDirectConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s AllocateDirectConnectionResponse) GoString() string {
	return s.String()
}

func (s *AllocateDirectConnectionResponse) SetHeaders(v map[string]*string) *AllocateDirectConnectionResponse {
	s.Headers = v
	return s
}

func (s *AllocateDirectConnectionResponse) SetBody(v *AllocateDirectConnectionResponseBody) *AllocateDirectConnectionResponse {
	s.Body = v
	return s
}

type AllocateInstancePublicConnectionRequest struct {
	SecurityToken          *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	Port                   *string `json:"Port,omitempty" xml:"Port,omitempty"`
	ClientToken            *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s AllocateInstancePublicConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s AllocateInstancePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *AllocateInstancePublicConnectionRequest) SetSecurityToken(v string) *AllocateInstancePublicConnectionRequest {
	s.SecurityToken = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetOwnerId(v int64) *AllocateInstancePublicConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetResourceOwnerAccount(v string) *AllocateInstancePublicConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetResourceOwnerId(v int64) *AllocateInstancePublicConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetOwnerAccount(v string) *AllocateInstancePublicConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetInstanceId(v string) *AllocateInstancePublicConnectionRequest {
	s.InstanceId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetConnectionStringPrefix(v string) *AllocateInstancePublicConnectionRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetPort(v string) *AllocateInstancePublicConnectionRequest {
	s.Port = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetClientToken(v string) *AllocateInstancePublicConnectionRequest {
	s.ClientToken = &v
	return s
}

type AllocateInstancePublicConnectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateInstancePublicConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AllocateInstancePublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateInstancePublicConnectionResponseBody) SetRequestId(v string) *AllocateInstancePublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

type AllocateInstancePublicConnectionResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AllocateInstancePublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AllocateInstancePublicConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s AllocateInstancePublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *AllocateInstancePublicConnectionResponse) SetHeaders(v map[string]*string) *AllocateInstancePublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *AllocateInstancePublicConnectionResponse) SetBody(v *AllocateInstancePublicConnectionResponseBody) *AllocateInstancePublicConnectionResponse {
	s.Body = v
	return s
}

type CreateAccountRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPrivilege     *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	AccountPassword      *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	AccountDescription   *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	AccountType          *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountRequest) SetSecurityToken(v string) *CreateAccountRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateAccountRequest) SetOwnerId(v int64) *CreateAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAccountRequest) SetResourceOwnerAccount(v string) *CreateAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAccountRequest) SetResourceOwnerId(v int64) *CreateAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAccountRequest) SetOwnerAccount(v string) *CreateAccountRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateAccountRequest) SetInstanceId(v string) *CreateAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAccountRequest) SetAccountName(v string) *CreateAccountRequest {
	s.AccountName = &v
	return s
}

func (s *CreateAccountRequest) SetAccountPrivilege(v string) *CreateAccountRequest {
	s.AccountPrivilege = &v
	return s
}

func (s *CreateAccountRequest) SetAccountPassword(v string) *CreateAccountRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateAccountRequest) SetAccountDescription(v string) *CreateAccountRequest {
	s.AccountDescription = &v
	return s
}

func (s *CreateAccountRequest) SetAccountType(v string) *CreateAccountRequest {
	s.AccountType = &v
	return s
}

func (s *CreateAccountRequest) SetClientToken(v string) *CreateAccountRequest {
	s.ClientToken = &v
	return s
}

type CreateAccountResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AcountName *string `json:"AcountName,omitempty" xml:"AcountName,omitempty"`
}

func (s CreateAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccountResponseBody) SetRequestId(v string) *CreateAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccountResponseBody) SetInstanceId(v string) *CreateAccountResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateAccountResponseBody) SetAcountName(v string) *CreateAccountResponseBody {
	s.AcountName = &v
	return s
}

type CreateAccountResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateAccountResponse) SetHeaders(v map[string]*string) *CreateAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateAccountResponse) SetBody(v *CreateAccountResponseBody) *CreateAccountResponse {
	s.Body = v
	return s
}

type CreateBackupRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateBackupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupRequest) GoString() string {
	return s.String()
}

func (s *CreateBackupRequest) SetSecurityToken(v string) *CreateBackupRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateBackupRequest) SetOwnerId(v int64) *CreateBackupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateBackupRequest) SetResourceOwnerAccount(v string) *CreateBackupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateBackupRequest) SetResourceOwnerId(v int64) *CreateBackupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateBackupRequest) SetOwnerAccount(v string) *CreateBackupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateBackupRequest) SetInstanceId(v string) *CreateBackupRequest {
	s.InstanceId = &v
	return s
}

type CreateBackupResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BackupJobID *string `json:"BackupJobID,omitempty" xml:"BackupJobID,omitempty"`
}

func (s CreateBackupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBackupResponseBody) SetRequestId(v string) *CreateBackupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBackupResponseBody) SetBackupJobID(v string) *CreateBackupResponseBody {
	s.BackupJobID = &v
	return s
}

type CreateBackupResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateBackupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBackupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupResponse) GoString() string {
	return s.String()
}

func (s *CreateBackupResponse) SetHeaders(v map[string]*string) *CreateBackupResponse {
	s.Headers = v
	return s
}

func (s *CreateBackupResponse) SetBody(v *CreateBackupResponseBody) *CreateBackupResponse {
	s.Body = v
	return s
}

type CreateCacheAnalysisTaskRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateCacheAnalysisTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCacheAnalysisTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateCacheAnalysisTaskRequest) SetSecurityToken(v string) *CreateCacheAnalysisTaskRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateCacheAnalysisTaskRequest) SetOwnerId(v int64) *CreateCacheAnalysisTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCacheAnalysisTaskRequest) SetResourceOwnerAccount(v string) *CreateCacheAnalysisTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCacheAnalysisTaskRequest) SetResourceOwnerId(v int64) *CreateCacheAnalysisTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCacheAnalysisTaskRequest) SetOwnerAccount(v string) *CreateCacheAnalysisTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateCacheAnalysisTaskRequest) SetInstanceId(v string) *CreateCacheAnalysisTaskRequest {
	s.InstanceId = &v
	return s
}

type CreateCacheAnalysisTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCacheAnalysisTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCacheAnalysisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCacheAnalysisTaskResponseBody) SetRequestId(v string) *CreateCacheAnalysisTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateCacheAnalysisTaskResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCacheAnalysisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCacheAnalysisTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCacheAnalysisTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateCacheAnalysisTaskResponse) SetHeaders(v map[string]*string) *CreateCacheAnalysisTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateCacheAnalysisTaskResponse) SetBody(v *CreateCacheAnalysisTaskResponseBody) *CreateCacheAnalysisTaskResponse {
	s.Body = v
	return s
}

type CreateGlobalDistributeCacheRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	SeedSubInstanceId    *string `json:"SeedSubInstanceId,omitempty" xml:"SeedSubInstanceId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateGlobalDistributeCacheRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGlobalDistributeCacheRequest) GoString() string {
	return s.String()
}

func (s *CreateGlobalDistributeCacheRequest) SetSecurityToken(v string) *CreateGlobalDistributeCacheRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateGlobalDistributeCacheRequest) SetOwnerId(v int64) *CreateGlobalDistributeCacheRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateGlobalDistributeCacheRequest) SetResourceOwnerAccount(v string) *CreateGlobalDistributeCacheRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateGlobalDistributeCacheRequest) SetResourceOwnerId(v int64) *CreateGlobalDistributeCacheRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateGlobalDistributeCacheRequest) SetOwnerAccount(v string) *CreateGlobalDistributeCacheRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateGlobalDistributeCacheRequest) SetSeedSubInstanceId(v string) *CreateGlobalDistributeCacheRequest {
	s.SeedSubInstanceId = &v
	return s
}

func (s *CreateGlobalDistributeCacheRequest) SetClientToken(v string) *CreateGlobalDistributeCacheRequest {
	s.ClientToken = &v
	return s
}

type CreateGlobalDistributeCacheResponseBody struct {
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	GlobalInstanceId *string `json:"GlobalInstanceId,omitempty" xml:"GlobalInstanceId,omitempty"`
}

func (s CreateGlobalDistributeCacheResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGlobalDistributeCacheResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGlobalDistributeCacheResponseBody) SetRequestId(v string) *CreateGlobalDistributeCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGlobalDistributeCacheResponseBody) SetGlobalInstanceId(v string) *CreateGlobalDistributeCacheResponseBody {
	s.GlobalInstanceId = &v
	return s
}

type CreateGlobalDistributeCacheResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGlobalDistributeCacheResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGlobalDistributeCacheResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGlobalDistributeCacheResponse) GoString() string {
	return s.String()
}

func (s *CreateGlobalDistributeCacheResponse) SetHeaders(v map[string]*string) *CreateGlobalDistributeCacheResponse {
	s.Headers = v
	return s
}

func (s *CreateGlobalDistributeCacheResponse) SetBody(v *CreateGlobalDistributeCacheResponseBody) *CreateGlobalDistributeCacheResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Token                *string `json:"Token,omitempty" xml:"Token,omitempty"`
	InstanceName         *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Password             *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Capacity             *int64  `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	InstanceClass        *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Config               *string `json:"Config,omitempty" xml:"Config,omitempty"`
	ChargeType           *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	NodeType             *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	NetworkType          *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	Period               *string `json:"Period,omitempty" xml:"Period,omitempty"`
	BusinessInfo         *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	CouponNo             *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	SrcDBInstanceId      *string `json:"SrcDBInstanceId,omitempty" xml:"SrcDBInstanceId,omitempty"`
	BackupId             *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	InstanceType         *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	EngineVersion        *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	PrivateIpAddress     *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	AutoUseCoupon        *string `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	AutoRenew            *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewPeriod      *string `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RestoreTime          *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	ShardCount           *int32  `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
	GlobalInstanceId     *string `json:"GlobalInstanceId,omitempty" xml:"GlobalInstanceId,omitempty"`
	GlobalInstance       *bool   `json:"GlobalInstance,omitempty" xml:"GlobalInstance,omitempty"`
	SecondaryZoneId      *string `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetSecurityToken(v string) *CreateInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateInstanceRequest) SetOwnerId(v int64) *CreateInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceOwnerAccount(v string) *CreateInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceOwnerId(v int64) *CreateInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateInstanceRequest) SetOwnerAccount(v string) *CreateInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateInstanceRequest) SetRegionId(v string) *CreateInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateInstanceRequest) SetToken(v string) *CreateInstanceRequest {
	s.Token = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetPassword(v string) *CreateInstanceRequest {
	s.Password = &v
	return s
}

func (s *CreateInstanceRequest) SetCapacity(v int64) *CreateInstanceRequest {
	s.Capacity = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceClass(v string) *CreateInstanceRequest {
	s.InstanceClass = &v
	return s
}

func (s *CreateInstanceRequest) SetZoneId(v string) *CreateInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateInstanceRequest) SetConfig(v string) *CreateInstanceRequest {
	s.Config = &v
	return s
}

func (s *CreateInstanceRequest) SetChargeType(v string) *CreateInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetNodeType(v string) *CreateInstanceRequest {
	s.NodeType = &v
	return s
}

func (s *CreateInstanceRequest) SetNetworkType(v string) *CreateInstanceRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateInstanceRequest) SetVpcId(v string) *CreateInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateInstanceRequest) SetVSwitchId(v string) *CreateInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateInstanceRequest) SetPeriod(v string) *CreateInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateInstanceRequest) SetBusinessInfo(v string) *CreateInstanceRequest {
	s.BusinessInfo = &v
	return s
}

func (s *CreateInstanceRequest) SetCouponNo(v string) *CreateInstanceRequest {
	s.CouponNo = &v
	return s
}

func (s *CreateInstanceRequest) SetSrcDBInstanceId(v string) *CreateInstanceRequest {
	s.SrcDBInstanceId = &v
	return s
}

func (s *CreateInstanceRequest) SetBackupId(v string) *CreateInstanceRequest {
	s.BackupId = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceType(v string) *CreateInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateInstanceRequest) SetEngineVersion(v string) *CreateInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateInstanceRequest) SetPrivateIpAddress(v string) *CreateInstanceRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoUseCoupon(v string) *CreateInstanceRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenew(v string) *CreateInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenewPeriod(v string) *CreateInstanceRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceGroupId(v string) *CreateInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstanceRequest) SetRestoreTime(v string) *CreateInstanceRequest {
	s.RestoreTime = &v
	return s
}

func (s *CreateInstanceRequest) SetDedicatedHostGroupId(v string) *CreateInstanceRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *CreateInstanceRequest) SetShardCount(v int32) *CreateInstanceRequest {
	s.ShardCount = &v
	return s
}

func (s *CreateInstanceRequest) SetGlobalInstanceId(v string) *CreateInstanceRequest {
	s.GlobalInstanceId = &v
	return s
}

func (s *CreateInstanceRequest) SetGlobalInstance(v bool) *CreateInstanceRequest {
	s.GlobalInstance = &v
	return s
}

func (s *CreateInstanceRequest) SetSecondaryZoneId(v string) *CreateInstanceRequest {
	s.SecondaryZoneId = &v
	return s
}

type CreateInstanceResponseBody struct {
	Connections      *int64  `json:"Connections,omitempty" xml:"Connections,omitempty"`
	UserName         *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	EndTime          *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ZoneId           *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Config           *string `json:"Config,omitempty" xml:"Config,omitempty"`
	Port             *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	VSwitchId        *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	PrivateIpAddr    *string `json:"PrivateIpAddr,omitempty" xml:"PrivateIpAddr,omitempty"`
	ConnectionDomain *string `json:"ConnectionDomain,omitempty" xml:"ConnectionDomain,omitempty"`
	InstanceName     *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	VpcId            *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	QPS              *int64  `json:"QPS,omitempty" xml:"QPS,omitempty"`
	Capacity         *int64  `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	ChargeType       *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	NetworkType      *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	InstanceStatus   *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	NodeType         *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	Bandwidth        *int64  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetConnections(v int64) *CreateInstanceResponseBody {
	s.Connections = &v
	return s
}

func (s *CreateInstanceResponseBody) SetUserName(v string) *CreateInstanceResponseBody {
	s.UserName = &v
	return s
}

func (s *CreateInstanceResponseBody) SetEndTime(v string) *CreateInstanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetZoneId(v string) *CreateInstanceResponseBody {
	s.ZoneId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetInstanceId(v string) *CreateInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetConfig(v string) *CreateInstanceResponseBody {
	s.Config = &v
	return s
}

func (s *CreateInstanceResponseBody) SetPort(v int32) *CreateInstanceResponseBody {
	s.Port = &v
	return s
}

func (s *CreateInstanceResponseBody) SetVSwitchId(v string) *CreateInstanceResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetPrivateIpAddr(v string) *CreateInstanceResponseBody {
	s.PrivateIpAddr = &v
	return s
}

func (s *CreateInstanceResponseBody) SetConnectionDomain(v string) *CreateInstanceResponseBody {
	s.ConnectionDomain = &v
	return s
}

func (s *CreateInstanceResponseBody) SetInstanceName(v string) *CreateInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceResponseBody) SetVpcId(v string) *CreateInstanceResponseBody {
	s.VpcId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetQPS(v int64) *CreateInstanceResponseBody {
	s.QPS = &v
	return s
}

func (s *CreateInstanceResponseBody) SetCapacity(v int64) *CreateInstanceResponseBody {
	s.Capacity = &v
	return s
}

func (s *CreateInstanceResponseBody) SetChargeType(v string) *CreateInstanceResponseBody {
	s.ChargeType = &v
	return s
}

func (s *CreateInstanceResponseBody) SetNetworkType(v string) *CreateInstanceResponseBody {
	s.NetworkType = &v
	return s
}

func (s *CreateInstanceResponseBody) SetInstanceStatus(v string) *CreateInstanceResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *CreateInstanceResponseBody) SetNodeType(v string) *CreateInstanceResponseBody {
	s.NodeType = &v
	return s
}

func (s *CreateInstanceResponseBody) SetBandwidth(v int64) *CreateInstanceResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *CreateInstanceResponseBody) SetRegionId(v string) *CreateInstanceResponseBody {
	s.RegionId = &v
	return s
}

type CreateInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponse) SetHeaders(v map[string]*string) *CreateInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceResponse) SetBody(v *CreateInstanceResponseBody) *CreateInstanceResponse {
	s.Body = v
	return s
}

type CreateTairInstanceRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceName         *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Password             *string `json:"Password,omitempty" xml:"Password,omitempty"`
	InstanceClass        *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ChargeType           *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	Period               *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	BusinessInfo         *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	CouponNo             *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	SrcDBInstanceId      *string `json:"SrcDBInstanceId,omitempty" xml:"SrcDBInstanceId,omitempty"`
	BackupId             *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	PrivateIpAddress     *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	AutoUseCoupon        *string `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	AutoRenew            *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewPeriod      *string `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	ResourceGroupId      *int32  `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	AutoPay              *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	StorageType          *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	Storage              *int32  `json:"Storage,omitempty" xml:"Storage,omitempty"`
	ShardType            *string `json:"ShardType,omitempty" xml:"ShardType,omitempty"`
	ShardCount           *int32  `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
	EngineVersion        *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	InstanceType         *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s CreateTairInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTairInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateTairInstanceRequest) SetSecurityToken(v string) *CreateTairInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateTairInstanceRequest) SetOwnerId(v int64) *CreateTairInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTairInstanceRequest) SetResourceOwnerAccount(v string) *CreateTairInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTairInstanceRequest) SetResourceOwnerId(v int64) *CreateTairInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTairInstanceRequest) SetOwnerAccount(v string) *CreateTairInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTairInstanceRequest) SetRegionId(v string) *CreateTairInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTairInstanceRequest) SetInstanceName(v string) *CreateTairInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateTairInstanceRequest) SetPassword(v string) *CreateTairInstanceRequest {
	s.Password = &v
	return s
}

func (s *CreateTairInstanceRequest) SetInstanceClass(v string) *CreateTairInstanceRequest {
	s.InstanceClass = &v
	return s
}

func (s *CreateTairInstanceRequest) SetZoneId(v string) *CreateTairInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateTairInstanceRequest) SetChargeType(v string) *CreateTairInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateTairInstanceRequest) SetVpcId(v string) *CreateTairInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateTairInstanceRequest) SetVSwitchId(v string) *CreateTairInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateTairInstanceRequest) SetPeriod(v int32) *CreateTairInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateTairInstanceRequest) SetBusinessInfo(v string) *CreateTairInstanceRequest {
	s.BusinessInfo = &v
	return s
}

func (s *CreateTairInstanceRequest) SetCouponNo(v string) *CreateTairInstanceRequest {
	s.CouponNo = &v
	return s
}

func (s *CreateTairInstanceRequest) SetSrcDBInstanceId(v string) *CreateTairInstanceRequest {
	s.SrcDBInstanceId = &v
	return s
}

func (s *CreateTairInstanceRequest) SetBackupId(v string) *CreateTairInstanceRequest {
	s.BackupId = &v
	return s
}

func (s *CreateTairInstanceRequest) SetPrivateIpAddress(v string) *CreateTairInstanceRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CreateTairInstanceRequest) SetAutoUseCoupon(v string) *CreateTairInstanceRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateTairInstanceRequest) SetAutoRenew(v string) *CreateTairInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateTairInstanceRequest) SetAutoRenewPeriod(v string) *CreateTairInstanceRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateTairInstanceRequest) SetResourceGroupId(v int32) *CreateTairInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTairInstanceRequest) SetAutoPay(v bool) *CreateTairInstanceRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateTairInstanceRequest) SetClientToken(v string) *CreateTairInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTairInstanceRequest) SetStorageType(v string) *CreateTairInstanceRequest {
	s.StorageType = &v
	return s
}

func (s *CreateTairInstanceRequest) SetStorage(v int32) *CreateTairInstanceRequest {
	s.Storage = &v
	return s
}

func (s *CreateTairInstanceRequest) SetShardType(v string) *CreateTairInstanceRequest {
	s.ShardType = &v
	return s
}

func (s *CreateTairInstanceRequest) SetShardCount(v int32) *CreateTairInstanceRequest {
	s.ShardCount = &v
	return s
}

func (s *CreateTairInstanceRequest) SetEngineVersion(v string) *CreateTairInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateTairInstanceRequest) SetInstanceType(v string) *CreateTairInstanceRequest {
	s.InstanceType = &v
	return s
}

type CreateTairInstanceResponseBody struct {
	Connections      *int64  `json:"Connections,omitempty" xml:"Connections,omitempty"`
	TaskId           *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ZoneId           *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Config           *string `json:"Config,omitempty" xml:"Config,omitempty"`
	Port             *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	ConnectionDomain *string `json:"ConnectionDomain,omitempty" xml:"ConnectionDomain,omitempty"`
	InstanceName     *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	QPS              *int64  `json:"QPS,omitempty" xml:"QPS,omitempty"`
	ChargeType       *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	InstanceStatus   *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	Bandwidth        *int64  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateTairInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTairInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTairInstanceResponseBody) SetConnections(v int64) *CreateTairInstanceResponseBody {
	s.Connections = &v
	return s
}

func (s *CreateTairInstanceResponseBody) SetTaskId(v string) *CreateTairInstanceResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateTairInstanceResponseBody) SetRequestId(v string) *CreateTairInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTairInstanceResponseBody) SetZoneId(v string) *CreateTairInstanceResponseBody {
	s.ZoneId = &v
	return s
}

func (s *CreateTairInstanceResponseBody) SetInstanceId(v string) *CreateTairInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateTairInstanceResponseBody) SetConfig(v string) *CreateTairInstanceResponseBody {
	s.Config = &v
	return s
}

func (s *CreateTairInstanceResponseBody) SetPort(v int32) *CreateTairInstanceResponseBody {
	s.Port = &v
	return s
}

func (s *CreateTairInstanceResponseBody) SetConnectionDomain(v string) *CreateTairInstanceResponseBody {
	s.ConnectionDomain = &v
	return s
}

func (s *CreateTairInstanceResponseBody) SetInstanceName(v string) *CreateTairInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *CreateTairInstanceResponseBody) SetQPS(v int64) *CreateTairInstanceResponseBody {
	s.QPS = &v
	return s
}

func (s *CreateTairInstanceResponseBody) SetChargeType(v string) *CreateTairInstanceResponseBody {
	s.ChargeType = &v
	return s
}

func (s *CreateTairInstanceResponseBody) SetInstanceStatus(v string) *CreateTairInstanceResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *CreateTairInstanceResponseBody) SetBandwidth(v int64) *CreateTairInstanceResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *CreateTairInstanceResponseBody) SetRegionId(v string) *CreateTairInstanceResponseBody {
	s.RegionId = &v
	return s
}

type CreateTairInstanceResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTairInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTairInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTairInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateTairInstanceResponse) SetHeaders(v map[string]*string) *CreateTairInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateTairInstanceResponse) SetBody(v *CreateTairInstanceResponseBody) *CreateTairInstanceResponse {
	s.Body = v
	return s
}

type CreateUserClusterHostRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ChargeType           *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	HostClass            *string `json:"HostClass,omitempty" xml:"HostClass,omitempty"`
	OrderNum             *int32  `json:"OrderNum,omitempty" xml:"OrderNum,omitempty"`
	OrderPeriod          *int32  `json:"OrderPeriod,omitempty" xml:"OrderPeriod,omitempty"`
	AutoPay              *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoRenew            *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AgentId              *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	BusinessInfo         *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	CouponNo             *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	Engine               *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateUserClusterHostRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserClusterHostRequest) GoString() string {
	return s.String()
}

func (s *CreateUserClusterHostRequest) SetSecurityToken(v string) *CreateUserClusterHostRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateUserClusterHostRequest) SetOwnerId(v int64) *CreateUserClusterHostRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateUserClusterHostRequest) SetResourceOwnerAccount(v string) *CreateUserClusterHostRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateUserClusterHostRequest) SetResourceOwnerId(v int64) *CreateUserClusterHostRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateUserClusterHostRequest) SetOwnerAccount(v string) *CreateUserClusterHostRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateUserClusterHostRequest) SetRegionId(v string) *CreateUserClusterHostRequest {
	s.RegionId = &v
	return s
}

func (s *CreateUserClusterHostRequest) SetClusterId(v string) *CreateUserClusterHostRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateUserClusterHostRequest) SetChargeType(v string) *CreateUserClusterHostRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateUserClusterHostRequest) SetHostClass(v string) *CreateUserClusterHostRequest {
	s.HostClass = &v
	return s
}

func (s *CreateUserClusterHostRequest) SetOrderNum(v int32) *CreateUserClusterHostRequest {
	s.OrderNum = &v
	return s
}

func (s *CreateUserClusterHostRequest) SetOrderPeriod(v int32) *CreateUserClusterHostRequest {
	s.OrderPeriod = &v
	return s
}

func (s *CreateUserClusterHostRequest) SetAutoPay(v bool) *CreateUserClusterHostRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateUserClusterHostRequest) SetAutoRenew(v bool) *CreateUserClusterHostRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateUserClusterHostRequest) SetAgentId(v string) *CreateUserClusterHostRequest {
	s.AgentId = &v
	return s
}

func (s *CreateUserClusterHostRequest) SetBusinessInfo(v string) *CreateUserClusterHostRequest {
	s.BusinessInfo = &v
	return s
}

func (s *CreateUserClusterHostRequest) SetCouponNo(v string) *CreateUserClusterHostRequest {
	s.CouponNo = &v
	return s
}

func (s *CreateUserClusterHostRequest) SetEngine(v string) *CreateUserClusterHostRequest {
	s.Engine = &v
	return s
}

func (s *CreateUserClusterHostRequest) SetZoneId(v string) *CreateUserClusterHostRequest {
	s.ZoneId = &v
	return s
}

type CreateUserClusterHostResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	HostId    *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s CreateUserClusterHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserClusterHostResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserClusterHostResponseBody) SetRequestId(v string) *CreateUserClusterHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserClusterHostResponseBody) SetClusterId(v string) *CreateUserClusterHostResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateUserClusterHostResponseBody) SetHostId(v string) *CreateUserClusterHostResponseBody {
	s.HostId = &v
	return s
}

type CreateUserClusterHostResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateUserClusterHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUserClusterHostResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserClusterHostResponse) GoString() string {
	return s.String()
}

func (s *CreateUserClusterHostResponse) SetHeaders(v map[string]*string) *CreateUserClusterHostResponse {
	s.Headers = v
	return s
}

func (s *CreateUserClusterHostResponse) SetBody(v *CreateUserClusterHostResponseBody) *CreateUserClusterHostResponse {
	s.Body = v
	return s
}

type DeleteAccountRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s DeleteAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccountRequest) SetSecurityToken(v string) *DeleteAccountRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteAccountRequest) SetOwnerId(v int64) *DeleteAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAccountRequest) SetResourceOwnerAccount(v string) *DeleteAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAccountRequest) SetResourceOwnerId(v int64) *DeleteAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteAccountRequest) SetOwnerAccount(v string) *DeleteAccountRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteAccountRequest) SetInstanceId(v string) *DeleteAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteAccountRequest) SetAccountName(v string) *DeleteAccountRequest {
	s.AccountName = &v
	return s
}

type DeleteAccountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccountResponseBody) SetRequestId(v string) *DeleteAccountResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAccountResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccountResponse) SetHeaders(v map[string]*string) *DeleteAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccountResponse) SetBody(v *DeleteAccountResponseBody) *DeleteAccountResponse {
	s.Body = v
	return s
}

type DeleteInstanceRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	GlobalInstanceId     *string `json:"GlobalInstanceId,omitempty" xml:"GlobalInstanceId,omitempty"`
	ReleaseSubInstance   *bool   `json:"ReleaseSubInstance,omitempty" xml:"ReleaseSubInstance,omitempty"`
}

func (s DeleteInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) SetSecurityToken(v string) *DeleteInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteInstanceRequest) SetOwnerId(v int64) *DeleteInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteInstanceRequest) SetResourceOwnerAccount(v string) *DeleteInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteInstanceRequest) SetResourceOwnerId(v int64) *DeleteInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteInstanceRequest) SetOwnerAccount(v string) *DeleteInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteInstanceRequest) SetInstanceId(v string) *DeleteInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceRequest) SetGlobalInstanceId(v string) *DeleteInstanceRequest {
	s.GlobalInstanceId = &v
	return s
}

func (s *DeleteInstanceRequest) SetReleaseSubInstance(v bool) *DeleteInstanceRequest {
	s.ReleaseSubInstance = &v
	return s
}

type DeleteInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) SetRequestId(v string) *DeleteInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponse) SetHeaders(v map[string]*string) *DeleteInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceResponse) SetBody(v *DeleteInstanceResponseBody) *DeleteInstanceResponse {
	s.Body = v
	return s
}

type DeleteShardingNodeRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DeleteShardingNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteShardingNodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteShardingNodeRequest) SetSecurityToken(v string) *DeleteShardingNodeRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteShardingNodeRequest) SetOwnerId(v int64) *DeleteShardingNodeRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteShardingNodeRequest) SetResourceOwnerAccount(v string) *DeleteShardingNodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteShardingNodeRequest) SetResourceOwnerId(v int64) *DeleteShardingNodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteShardingNodeRequest) SetOwnerAccount(v string) *DeleteShardingNodeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteShardingNodeRequest) SetInstanceId(v string) *DeleteShardingNodeRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteShardingNodeRequest) SetNodeId(v string) *DeleteShardingNodeRequest {
	s.NodeId = &v
	return s
}

type DeleteShardingNodeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteShardingNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteShardingNodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteShardingNodeResponseBody) SetRequestId(v string) *DeleteShardingNodeResponseBody {
	s.RequestId = &v
	return s
}

type DeleteShardingNodeResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteShardingNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteShardingNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteShardingNodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteShardingNodeResponse) SetHeaders(v map[string]*string) *DeleteShardingNodeResponse {
	s.Headers = v
	return s
}

func (s *DeleteShardingNodeResponse) SetBody(v *DeleteShardingNodeResponseBody) *DeleteShardingNodeResponse {
	s.Body = v
	return s
}

type DeleteUserClusterHostRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	HostId               *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	Engine               *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DeleteUserClusterHostRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserClusterHostRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserClusterHostRequest) SetSecurityToken(v string) *DeleteUserClusterHostRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteUserClusterHostRequest) SetOwnerId(v int64) *DeleteUserClusterHostRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteUserClusterHostRequest) SetResourceOwnerAccount(v string) *DeleteUserClusterHostRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteUserClusterHostRequest) SetResourceOwnerId(v int64) *DeleteUserClusterHostRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteUserClusterHostRequest) SetOwnerAccount(v string) *DeleteUserClusterHostRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteUserClusterHostRequest) SetRegionId(v string) *DeleteUserClusterHostRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteUserClusterHostRequest) SetHostId(v string) *DeleteUserClusterHostRequest {
	s.HostId = &v
	return s
}

func (s *DeleteUserClusterHostRequest) SetEngine(v string) *DeleteUserClusterHostRequest {
	s.Engine = &v
	return s
}

func (s *DeleteUserClusterHostRequest) SetZoneId(v string) *DeleteUserClusterHostRequest {
	s.ZoneId = &v
	return s
}

func (s *DeleteUserClusterHostRequest) SetClusterId(v string) *DeleteUserClusterHostRequest {
	s.ClusterId = &v
	return s
}

type DeleteUserClusterHostResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserClusterHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserClusterHostResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserClusterHostResponseBody) SetRequestId(v string) *DeleteUserClusterHostResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserClusterHostResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteUserClusterHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUserClusterHostResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserClusterHostResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserClusterHostResponse) SetHeaders(v map[string]*string) *DeleteUserClusterHostResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserClusterHostResponse) SetBody(v *DeleteUserClusterHostResponseBody) *DeleteUserClusterHostResponse {
	s.Body = v
	return s
}

type DescribeAccountsRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s DescribeAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountsRequest) SetSecurityToken(v string) *DescribeAccountsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeAccountsRequest) SetOwnerId(v int64) *DescribeAccountsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAccountsRequest) SetResourceOwnerAccount(v string) *DescribeAccountsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAccountsRequest) SetResourceOwnerId(v int64) *DescribeAccountsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAccountsRequest) SetOwnerAccount(v string) *DescribeAccountsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAccountsRequest) SetInstanceId(v string) *DescribeAccountsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAccountsRequest) SetAccountName(v string) *DescribeAccountsRequest {
	s.AccountName = &v
	return s
}

type DescribeAccountsResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Accounts  *DescribeAccountsResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
}

func (s DescribeAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBody) SetRequestId(v string) *DescribeAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountsResponseBody) SetAccounts(v *DescribeAccountsResponseBodyAccounts) *DescribeAccountsResponseBody {
	s.Accounts = v
	return s
}

type DescribeAccountsResponseBodyAccounts struct {
	Account []*DescribeAccountsResponseBodyAccountsAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Repeated"`
}

func (s DescribeAccountsResponseBodyAccounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccounts) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccounts) SetAccount(v []*DescribeAccountsResponseBodyAccountsAccount) *DescribeAccountsResponseBodyAccounts {
	s.Account = v
	return s
}

type DescribeAccountsResponseBodyAccountsAccount struct {
	AccountStatus      *string                                                        `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	DatabasePrivileges *DescribeAccountsResponseBodyAccountsAccountDatabasePrivileges `json:"DatabasePrivileges,omitempty" xml:"DatabasePrivileges,omitempty" type:"Struct"`
	AccountDescription *string                                                        `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	InstanceId         *string                                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountType        *string                                                        `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	AccountName        *string                                                        `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s DescribeAccountsResponseBodyAccountsAccount) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountsAccount) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountsAccount) SetAccountStatus(v string) *DescribeAccountsResponseBodyAccountsAccount {
	s.AccountStatus = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsAccount) SetDatabasePrivileges(v *DescribeAccountsResponseBodyAccountsAccountDatabasePrivileges) *DescribeAccountsResponseBodyAccountsAccount {
	s.DatabasePrivileges = v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsAccount) SetAccountDescription(v string) *DescribeAccountsResponseBodyAccountsAccount {
	s.AccountDescription = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsAccount) SetInstanceId(v string) *DescribeAccountsResponseBodyAccountsAccount {
	s.InstanceId = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsAccount) SetAccountType(v string) *DescribeAccountsResponseBodyAccountsAccount {
	s.AccountType = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsAccount) SetAccountName(v string) *DescribeAccountsResponseBodyAccountsAccount {
	s.AccountName = &v
	return s
}

type DescribeAccountsResponseBodyAccountsAccountDatabasePrivileges struct {
	DatabasePrivilege []*DescribeAccountsResponseBodyAccountsAccountDatabasePrivilegesDatabasePrivilege `json:"DatabasePrivilege,omitempty" xml:"DatabasePrivilege,omitempty" type:"Repeated"`
}

func (s DescribeAccountsResponseBodyAccountsAccountDatabasePrivileges) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountsAccountDatabasePrivileges) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountsAccountDatabasePrivileges) SetDatabasePrivilege(v []*DescribeAccountsResponseBodyAccountsAccountDatabasePrivilegesDatabasePrivilege) *DescribeAccountsResponseBodyAccountsAccountDatabasePrivileges {
	s.DatabasePrivilege = v
	return s
}

type DescribeAccountsResponseBodyAccountsAccountDatabasePrivilegesDatabasePrivilege struct {
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
}

func (s DescribeAccountsResponseBodyAccountsAccountDatabasePrivilegesDatabasePrivilege) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountsAccountDatabasePrivilegesDatabasePrivilege) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountsAccountDatabasePrivilegesDatabasePrivilege) SetAccountPrivilege(v string) *DescribeAccountsResponseBodyAccountsAccountDatabasePrivilegesDatabasePrivilege {
	s.AccountPrivilege = &v
	return s
}

type DescribeAccountsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponse) SetHeaders(v map[string]*string) *DescribeAccountsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountsResponse) SetBody(v *DescribeAccountsResponseBody) *DescribeAccountsResponse {
	s.Body = v
	return s
}

type DescribeActiveOperationTaskRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty"`
	TaskType             *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	IsHistory            *int32  `json:"IsHistory,omitempty" xml:"IsHistory,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ProductId            *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s DescribeActiveOperationTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveOperationTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskRequest) SetSecurityToken(v string) *DescribeActiveOperationTaskRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeActiveOperationTaskRequest) SetOwnerId(v int64) *DescribeActiveOperationTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeActiveOperationTaskRequest) SetResourceOwnerAccount(v string) *DescribeActiveOperationTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeActiveOperationTaskRequest) SetResourceOwnerId(v int64) *DescribeActiveOperationTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeActiveOperationTaskRequest) SetOwnerAccount(v string) *DescribeActiveOperationTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeActiveOperationTaskRequest) SetRegion(v string) *DescribeActiveOperationTaskRequest {
	s.Region = &v
	return s
}

func (s *DescribeActiveOperationTaskRequest) SetTaskType(v string) *DescribeActiveOperationTaskRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeActiveOperationTaskRequest) SetIsHistory(v int32) *DescribeActiveOperationTaskRequest {
	s.IsHistory = &v
	return s
}

func (s *DescribeActiveOperationTaskRequest) SetPageSize(v int32) *DescribeActiveOperationTaskRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeActiveOperationTaskRequest) SetPageNumber(v int32) *DescribeActiveOperationTaskRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeActiveOperationTaskRequest) SetProductId(v string) *DescribeActiveOperationTaskRequest {
	s.ProductId = &v
	return s
}

type DescribeActiveOperationTaskResponseBody struct {
	TotalRecordCount *int32                                          `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	RequestId        *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize         *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber       *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Items            []*DescribeActiveOperationTaskResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeActiveOperationTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveOperationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskResponseBody) SetTotalRecordCount(v int32) *DescribeActiveOperationTaskResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBody) SetRequestId(v string) *DescribeActiveOperationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBody) SetPageSize(v int32) *DescribeActiveOperationTaskResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBody) SetPageNumber(v int32) *DescribeActiveOperationTaskResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBody) SetItems(v []*DescribeActiveOperationTaskResponseBodyItems) *DescribeActiveOperationTaskResponseBody {
	s.Items = v
	return s
}

type DescribeActiveOperationTaskResponseBodyItems struct {
	Status          *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	PrepareInterval *string `json:"PrepareInterval,omitempty" xml:"PrepareInterval,omitempty"`
	Deadline        *string `json:"Deadline,omitempty" xml:"Deadline,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TaskType        *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	ModifiedTime    *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	InsName         *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	DbType          *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	CreatedTime     *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Id              *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	SwitchTime      *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
}

func (s DescribeActiveOperationTaskResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveOperationTaskResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskResponseBodyItems) SetStatus(v int32) *DescribeActiveOperationTaskResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBodyItems) SetPrepareInterval(v string) *DescribeActiveOperationTaskResponseBodyItems {
	s.PrepareInterval = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBodyItems) SetDeadline(v string) *DescribeActiveOperationTaskResponseBodyItems {
	s.Deadline = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBodyItems) SetStartTime(v string) *DescribeActiveOperationTaskResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBodyItems) SetTaskType(v string) *DescribeActiveOperationTaskResponseBodyItems {
	s.TaskType = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBodyItems) SetModifiedTime(v string) *DescribeActiveOperationTaskResponseBodyItems {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBodyItems) SetInsName(v string) *DescribeActiveOperationTaskResponseBodyItems {
	s.InsName = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBodyItems) SetDbType(v string) *DescribeActiveOperationTaskResponseBodyItems {
	s.DbType = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBodyItems) SetRegion(v string) *DescribeActiveOperationTaskResponseBodyItems {
	s.Region = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBodyItems) SetCreatedTime(v string) *DescribeActiveOperationTaskResponseBodyItems {
	s.CreatedTime = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBodyItems) SetId(v int32) *DescribeActiveOperationTaskResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBodyItems) SetSwitchTime(v string) *DescribeActiveOperationTaskResponseBodyItems {
	s.SwitchTime = &v
	return s
}

type DescribeActiveOperationTaskResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeActiveOperationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeActiveOperationTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveOperationTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskResponse) SetHeaders(v map[string]*string) *DescribeActiveOperationTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeActiveOperationTaskResponse) SetBody(v *DescribeActiveOperationTaskResponseBody) *DescribeActiveOperationTaskResponse {
	s.Body = v
	return s
}

type DescribeAuditRecordsRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DatabaseName         *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	QueryKeywords        *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	HostAddress          *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeAuditRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuditRecordsRequest) SetSecurityToken(v string) *DescribeAuditRecordsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetOwnerId(v int64) *DescribeAuditRecordsRequest {
	s.OwnerId = &v
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

func (s *DescribeAuditRecordsRequest) SetOwnerAccount(v string) *DescribeAuditRecordsRequest {
	s.OwnerAccount = &v
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

func (s *DescribeAuditRecordsRequest) SetAccountName(v string) *DescribeAuditRecordsRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetDatabaseName(v string) *DescribeAuditRecordsRequest {
	s.DatabaseName = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetQueryKeywords(v string) *DescribeAuditRecordsRequest {
	s.QueryKeywords = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetHostAddress(v string) *DescribeAuditRecordsRequest {
	s.HostAddress = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetPageSize(v int32) *DescribeAuditRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetPageNumber(v int32) *DescribeAuditRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetStartTime(v string) *DescribeAuditRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetEndTime(v string) *DescribeAuditRecordsRequest {
	s.EndTime = &v
	return s
}

type DescribeAuditRecordsResponseBody struct {
	InstanceName     *string                                `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	TotalRecordCount *int32                                 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	EndTime          *string                                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId        *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize         *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber       *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	StartTime        *string                                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Items            *DescribeAuditRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeAuditRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditRecordsResponseBody) SetInstanceName(v string) *DescribeAuditRecordsResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeAuditRecordsResponseBody) SetTotalRecordCount(v int32) *DescribeAuditRecordsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeAuditRecordsResponseBody) SetEndTime(v string) *DescribeAuditRecordsResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeAuditRecordsResponseBody) SetRequestId(v string) *DescribeAuditRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuditRecordsResponseBody) SetPageSize(v int32) *DescribeAuditRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAuditRecordsResponseBody) SetPageNumber(v int32) *DescribeAuditRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAuditRecordsResponseBody) SetStartTime(v string) *DescribeAuditRecordsResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeAuditRecordsResponseBody) SetItems(v *DescribeAuditRecordsResponseBodyItems) *DescribeAuditRecordsResponseBody {
	s.Items = v
	return s
}

type DescribeAuditRecordsResponseBodyItems struct {
	SQL []*DescribeAuditRecordsResponseBodyItemsSQL `json:"SQL,omitempty" xml:"SQL,omitempty" type:"Repeated"`
}

func (s DescribeAuditRecordsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAuditRecordsResponseBodyItems) SetSQL(v []*DescribeAuditRecordsResponseBodyItemsSQL) *DescribeAuditRecordsResponseBodyItems {
	s.SQL = v
	return s
}

type DescribeAuditRecordsResponseBodyItemsSQL struct {
	HostAddress         *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	SQLText             *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	DatabaseName        *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	SQLType             *string `json:"SQLType,omitempty" xml:"SQLType,omitempty"`
	ExecuteTime         *string `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	TotalExecutionTimes *string `json:"TotalExecutionTimes,omitempty" xml:"TotalExecutionTimes,omitempty"`
	AccountName         *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	IPAddress           *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
}

func (s DescribeAuditRecordsResponseBodyItemsSQL) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditRecordsResponseBodyItemsSQL) GoString() string {
	return s.String()
}

func (s *DescribeAuditRecordsResponseBodyItemsSQL) SetHostAddress(v string) *DescribeAuditRecordsResponseBodyItemsSQL {
	s.HostAddress = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQL) SetSQLText(v string) *DescribeAuditRecordsResponseBodyItemsSQL {
	s.SQLText = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQL) SetDatabaseName(v string) *DescribeAuditRecordsResponseBodyItemsSQL {
	s.DatabaseName = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQL) SetSQLType(v string) *DescribeAuditRecordsResponseBodyItemsSQL {
	s.SQLType = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQL) SetExecuteTime(v string) *DescribeAuditRecordsResponseBodyItemsSQL {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQL) SetTotalExecutionTimes(v string) *DescribeAuditRecordsResponseBodyItemsSQL {
	s.TotalExecutionTimes = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQL) SetAccountName(v string) *DescribeAuditRecordsResponseBodyItemsSQL {
	s.AccountName = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQL) SetIPAddress(v string) *DescribeAuditRecordsResponseBodyItemsSQL {
	s.IPAddress = &v
	return s
}

type DescribeAuditRecordsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAuditRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAuditRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuditRecordsResponse) SetHeaders(v map[string]*string) *DescribeAuditRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuditRecordsResponse) SetBody(v *DescribeAuditRecordsResponseBody) *DescribeAuditRecordsResponse {
	s.Body = v
	return s
}

type DescribeBackupPolicyRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyRequest) SetSecurityToken(v string) *DescribeBackupPolicyRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetOwnerId(v int64) *DescribeBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetResourceOwnerAccount(v string) *DescribeBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetResourceOwnerId(v int64) *DescribeBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetOwnerAccount(v string) *DescribeBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetInstanceId(v string) *DescribeBackupPolicyRequest {
	s.InstanceId = &v
	return s
}

type DescribeBackupPolicyResponseBody struct {
	PreferredBackupPeriod   *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PreferredBackupTime     *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	EnableBackupLog         *int32  `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	BackupRetentionPeriod   *string `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	PreferredNextBackupTime *string `json:"PreferredNextBackupTime,omitempty" xml:"PreferredNextBackupTime,omitempty"`
}

func (s DescribeBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetRequestId(v string) *DescribeBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupTime(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupTime = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetEnableBackupLog(v int32) *DescribeBackupPolicyResponseBody {
	s.EnableBackupLog = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetBackupRetentionPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredNextBackupTime(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredNextBackupTime = &v
	return s
}

type DescribeBackupPolicyResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponse) SetHeaders(v map[string]*string) *DescribeBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupPolicyResponse) SetBody(v *DescribeBackupPolicyResponseBody) *DescribeBackupPolicyResponse {
	s.Body = v
	return s
}

type DescribeBackupsRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	BackupId             *int32  `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	NeedAof              *string `json:"NeedAof,omitempty" xml:"NeedAof,omitempty"`
}

func (s DescribeBackupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupsRequest) SetSecurityToken(v string) *DescribeBackupsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeBackupsRequest) SetOwnerId(v int64) *DescribeBackupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupsRequest) SetResourceOwnerAccount(v string) *DescribeBackupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBackupsRequest) SetResourceOwnerId(v int64) *DescribeBackupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBackupsRequest) SetOwnerAccount(v string) *DescribeBackupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBackupsRequest) SetInstanceId(v string) *DescribeBackupsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeBackupsRequest) SetBackupId(v int32) *DescribeBackupsRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsRequest) SetPageSize(v int32) *DescribeBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupsRequest) SetPageNumber(v int32) *DescribeBackupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupsRequest) SetStartTime(v string) *DescribeBackupsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupsRequest) SetEndTime(v string) *DescribeBackupsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupsRequest) SetNeedAof(v string) *DescribeBackupsRequest {
	s.NeedAof = &v
	return s
}

type DescribeBackupsResponseBody struct {
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Backups    *DescribeBackupsResponseBodyBackups `json:"Backups,omitempty" xml:"Backups,omitempty" type:"Struct"`
}

func (s DescribeBackupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBody) SetTotalCount(v int32) *DescribeBackupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetRequestId(v string) *DescribeBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetPageSize(v int32) *DescribeBackupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetPageNumber(v int32) *DescribeBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetBackups(v *DescribeBackupsResponseBodyBackups) *DescribeBackupsResponseBody {
	s.Backups = v
	return s
}

type DescribeBackupsResponseBodyBackups struct {
	Backup []*DescribeBackupsResponseBodyBackupsBackup `json:"Backup,omitempty" xml:"Backup,omitempty" type:"Repeated"`
}

func (s DescribeBackupsResponseBodyBackups) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseBodyBackups) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyBackups) SetBackup(v []*DescribeBackupsResponseBodyBackupsBackup) *DescribeBackupsResponseBodyBackups {
	s.Backup = v
	return s
}

type DescribeBackupsResponseBodyBackupsBackup struct {
	BackupStatus              *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	BackupType                *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	BackupStartTime           *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	NodeInstanceId            *string `json:"NodeInstanceId,omitempty" xml:"NodeInstanceId,omitempty"`
	BackupDownloadURL         *string `json:"BackupDownloadURL,omitempty" xml:"BackupDownloadURL,omitempty"`
	BackupEndTime             *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	EngineVersion             *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	BackupDBNames             *string `json:"BackupDBNames,omitempty" xml:"BackupDBNames,omitempty"`
	BackupId                  *int32  `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	BackupIntranetDownloadURL *string `json:"BackupIntranetDownloadURL,omitempty" xml:"BackupIntranetDownloadURL,omitempty"`
	BackupSize                *int64  `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	BackupMode                *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	BackupMethod              *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
}

func (s DescribeBackupsResponseBodyBackupsBackup) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseBodyBackupsBackup) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupStatus(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupStatus = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupType(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupStartTime(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetNodeInstanceId(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.NodeInstanceId = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupDownloadURL(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupDownloadURL = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupEndTime(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetEngineVersion(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.EngineVersion = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupDBNames(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupDBNames = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupId(v int32) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupIntranetDownloadURL(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupIntranetDownloadURL = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupSize(v int64) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupSize = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupMode(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupMethod(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupMethod = &v
	return s
}

type DescribeBackupsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponse) SetHeaders(v map[string]*string) *DescribeBackupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupsResponse) SetBody(v *DescribeBackupsResponseBody) *DescribeBackupsResponse {
	s.Body = v
	return s
}

type DescribeBackupTasksRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	BackupJobId          *string `json:"BackupJobId,omitempty" xml:"BackupJobId,omitempty"`
	JobMode              *string `json:"JobMode,omitempty" xml:"JobMode,omitempty"`
}

func (s DescribeBackupTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupTasksRequest) SetSecurityToken(v string) *DescribeBackupTasksRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeBackupTasksRequest) SetOwnerId(v int64) *DescribeBackupTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupTasksRequest) SetResourceOwnerAccount(v string) *DescribeBackupTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBackupTasksRequest) SetResourceOwnerId(v int64) *DescribeBackupTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBackupTasksRequest) SetOwnerAccount(v string) *DescribeBackupTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBackupTasksRequest) SetInstanceId(v string) *DescribeBackupTasksRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeBackupTasksRequest) SetBackupJobId(v string) *DescribeBackupTasksRequest {
	s.BackupJobId = &v
	return s
}

func (s *DescribeBackupTasksRequest) SetJobMode(v string) *DescribeBackupTasksRequest {
	s.JobMode = &v
	return s
}

type DescribeBackupTasksResponseBody struct {
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceId *string                                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	BackupJobs []*DescribeBackupTasksResponseBodyBackupJobs `json:"BackupJobs,omitempty" xml:"BackupJobs,omitempty" type:"Repeated"`
}

func (s DescribeBackupTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupTasksResponseBody) SetRequestId(v string) *DescribeBackupTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupTasksResponseBody) SetInstanceId(v string) *DescribeBackupTasksResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeBackupTasksResponseBody) SetBackupJobs(v []*DescribeBackupTasksResponseBodyBackupJobs) *DescribeBackupTasksResponseBody {
	s.BackupJobs = v
	return s
}

type DescribeBackupTasksResponseBodyBackupJobs struct {
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Process              *string `json:"Process,omitempty" xml:"Process,omitempty"`
	JobMode              *string `json:"JobMode,omitempty" xml:"JobMode,omitempty"`
	BackupJobID          *int32  `json:"BackupJobID,omitempty" xml:"BackupJobID,omitempty"`
	BackupProgressStatus *string `json:"BackupProgressStatus,omitempty" xml:"BackupProgressStatus,omitempty"`
	TaskAction           *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
}

func (s DescribeBackupTasksResponseBodyBackupJobs) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTasksResponseBodyBackupJobs) GoString() string {
	return s.String()
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) SetStartTime(v string) *DescribeBackupTasksResponseBodyBackupJobs {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) SetProcess(v string) *DescribeBackupTasksResponseBodyBackupJobs {
	s.Process = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) SetJobMode(v string) *DescribeBackupTasksResponseBodyBackupJobs {
	s.JobMode = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) SetBackupJobID(v int32) *DescribeBackupTasksResponseBodyBackupJobs {
	s.BackupJobID = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) SetBackupProgressStatus(v string) *DescribeBackupTasksResponseBodyBackupJobs {
	s.BackupProgressStatus = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) SetTaskAction(v string) *DescribeBackupTasksResponseBodyBackupJobs {
	s.TaskAction = &v
	return s
}

type DescribeBackupTasksResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupTasksResponse) SetHeaders(v map[string]*string) *DescribeBackupTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupTasksResponse) SetBody(v *DescribeBackupTasksResponseBody) *DescribeBackupTasksResponse {
	s.Body = v
	return s
}

type DescribeCacheAnalysisReportRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Date                 *string `json:"Date,omitempty" xml:"Date,omitempty"`
	AnalysisType         *string `json:"AnalysisType,omitempty" xml:"AnalysisType,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumbers          *int32  `json:"PageNumbers,omitempty" xml:"PageNumbers,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeCacheAnalysisReportRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCacheAnalysisReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisReportRequest) SetSecurityToken(v string) *DescribeCacheAnalysisReportRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeCacheAnalysisReportRequest) SetOwnerId(v int64) *DescribeCacheAnalysisReportRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCacheAnalysisReportRequest) SetResourceOwnerAccount(v string) *DescribeCacheAnalysisReportRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCacheAnalysisReportRequest) SetResourceOwnerId(v int64) *DescribeCacheAnalysisReportRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCacheAnalysisReportRequest) SetOwnerAccount(v string) *DescribeCacheAnalysisReportRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCacheAnalysisReportRequest) SetInstanceId(v string) *DescribeCacheAnalysisReportRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCacheAnalysisReportRequest) SetDate(v string) *DescribeCacheAnalysisReportRequest {
	s.Date = &v
	return s
}

func (s *DescribeCacheAnalysisReportRequest) SetAnalysisType(v string) *DescribeCacheAnalysisReportRequest {
	s.AnalysisType = &v
	return s
}

func (s *DescribeCacheAnalysisReportRequest) SetPageSize(v int32) *DescribeCacheAnalysisReportRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCacheAnalysisReportRequest) SetPageNumbers(v int32) *DescribeCacheAnalysisReportRequest {
	s.PageNumbers = &v
	return s
}

func (s *DescribeCacheAnalysisReportRequest) SetNodeId(v string) *DescribeCacheAnalysisReportRequest {
	s.NodeId = &v
	return s
}

type DescribeCacheAnalysisReportResponseBody struct {
	TotalRecordCount *int32                   `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	HotKeys          []map[string]interface{} `json:"HotKeys,omitempty" xml:"HotKeys,omitempty" type:"Repeated"`
	PageRecordCount  *int32                   `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	BigKeys          []map[string]interface{} `json:"BigKeys,omitempty" xml:"BigKeys,omitempty" type:"Repeated"`
	PageSize         *int32                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId        *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber       *int32                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeCacheAnalysisReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCacheAnalysisReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisReportResponseBody) SetTotalRecordCount(v int32) *DescribeCacheAnalysisReportResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeCacheAnalysisReportResponseBody) SetHotKeys(v []map[string]interface{}) *DescribeCacheAnalysisReportResponseBody {
	s.HotKeys = v
	return s
}

func (s *DescribeCacheAnalysisReportResponseBody) SetPageRecordCount(v int32) *DescribeCacheAnalysisReportResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeCacheAnalysisReportResponseBody) SetBigKeys(v []map[string]interface{}) *DescribeCacheAnalysisReportResponseBody {
	s.BigKeys = v
	return s
}

func (s *DescribeCacheAnalysisReportResponseBody) SetPageSize(v int32) *DescribeCacheAnalysisReportResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCacheAnalysisReportResponseBody) SetRequestId(v string) *DescribeCacheAnalysisReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCacheAnalysisReportResponseBody) SetPageNumber(v int32) *DescribeCacheAnalysisReportResponseBody {
	s.PageNumber = &v
	return s
}

type DescribeCacheAnalysisReportResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCacheAnalysisReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCacheAnalysisReportResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCacheAnalysisReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisReportResponse) SetHeaders(v map[string]*string) *DescribeCacheAnalysisReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeCacheAnalysisReportResponse) SetBody(v *DescribeCacheAnalysisReportResponseBody) *DescribeCacheAnalysisReportResponse {
	s.Body = v
	return s
}

type DescribeCacheAnalysisReportListRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Days                 *int32  `json:"Days,omitempty" xml:"Days,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumbers          *int32  `json:"PageNumbers,omitempty" xml:"PageNumbers,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Date                 *string `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s DescribeCacheAnalysisReportListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCacheAnalysisReportListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisReportListRequest) SetSecurityToken(v string) *DescribeCacheAnalysisReportListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeCacheAnalysisReportListRequest) SetOwnerId(v int64) *DescribeCacheAnalysisReportListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCacheAnalysisReportListRequest) SetResourceOwnerAccount(v string) *DescribeCacheAnalysisReportListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCacheAnalysisReportListRequest) SetResourceOwnerId(v int64) *DescribeCacheAnalysisReportListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCacheAnalysisReportListRequest) SetOwnerAccount(v string) *DescribeCacheAnalysisReportListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCacheAnalysisReportListRequest) SetInstanceId(v string) *DescribeCacheAnalysisReportListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCacheAnalysisReportListRequest) SetDays(v int32) *DescribeCacheAnalysisReportListRequest {
	s.Days = &v
	return s
}

func (s *DescribeCacheAnalysisReportListRequest) SetPageSize(v int32) *DescribeCacheAnalysisReportListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCacheAnalysisReportListRequest) SetPageNumbers(v int32) *DescribeCacheAnalysisReportListRequest {
	s.PageNumbers = &v
	return s
}

func (s *DescribeCacheAnalysisReportListRequest) SetNodeId(v string) *DescribeCacheAnalysisReportListRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeCacheAnalysisReportListRequest) SetDate(v string) *DescribeCacheAnalysisReportListRequest {
	s.Date = &v
	return s
}

type DescribeCacheAnalysisReportListResponseBody struct {
	DailyTasks *DescribeCacheAnalysisReportListResponseBodyDailyTasks `json:"DailyTasks,omitempty" xml:"DailyTasks,omitempty" type:"Struct"`
	RequestId  *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceId *string                                                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeCacheAnalysisReportListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCacheAnalysisReportListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisReportListResponseBody) SetDailyTasks(v *DescribeCacheAnalysisReportListResponseBodyDailyTasks) *DescribeCacheAnalysisReportListResponseBody {
	s.DailyTasks = v
	return s
}

func (s *DescribeCacheAnalysisReportListResponseBody) SetRequestId(v string) *DescribeCacheAnalysisReportListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCacheAnalysisReportListResponseBody) SetInstanceId(v string) *DescribeCacheAnalysisReportListResponseBody {
	s.InstanceId = &v
	return s
}

type DescribeCacheAnalysisReportListResponseBodyDailyTasks struct {
	DailyTask []*DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTask `json:"DailyTask,omitempty" xml:"DailyTask,omitempty" type:"Repeated"`
}

func (s DescribeCacheAnalysisReportListResponseBodyDailyTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeCacheAnalysisReportListResponseBodyDailyTasks) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisReportListResponseBodyDailyTasks) SetDailyTask(v []*DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTask) *DescribeCacheAnalysisReportListResponseBodyDailyTasks {
	s.DailyTask = v
	return s
}

type DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTask struct {
	Tasks *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
	Date  *string                                                              `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTask) String() string {
	return tea.Prettify(s)
}

func (s DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTask) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTask) SetTasks(v *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasks) *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTask {
	s.Tasks = v
	return s
}

func (s *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTask) SetDate(v string) *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTask {
	s.Date = &v
	return s
}

type DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasks struct {
	Task []*DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasks) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasks) SetTask(v []*DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask) *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasks {
	s.Task = v
	return s
}

type DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask struct {
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	NodeId    *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask) String() string {
	return tea.Prettify(s)
}

func (s DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask) SetStatus(v string) *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask {
	s.Status = &v
	return s
}

func (s *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask) SetStartTime(v string) *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask {
	s.StartTime = &v
	return s
}

func (s *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask) SetTaskId(v string) *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask {
	s.TaskId = &v
	return s
}

func (s *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask) SetNodeId(v string) *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask {
	s.NodeId = &v
	return s
}

type DescribeCacheAnalysisReportListResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCacheAnalysisReportListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCacheAnalysisReportListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCacheAnalysisReportListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisReportListResponse) SetHeaders(v map[string]*string) *DescribeCacheAnalysisReportListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCacheAnalysisReportListResponse) SetBody(v *DescribeCacheAnalysisReportListResponseBody) *DescribeCacheAnalysisReportListResponse {
	s.Body = v
	return s
}

type DescribeClusterMemberInfoRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	FilterService        *string `json:"FilterService,omitempty" xml:"FilterService,omitempty"`
	NeedReplica          *string `json:"NeedReplica,omitempty" xml:"NeedReplica,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeClusterMemberInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterMemberInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterMemberInfoRequest) SetSecurityToken(v string) *DescribeClusterMemberInfoRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeClusterMemberInfoRequest) SetOwnerId(v int64) *DescribeClusterMemberInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeClusterMemberInfoRequest) SetResourceOwnerAccount(v string) *DescribeClusterMemberInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeClusterMemberInfoRequest) SetResourceOwnerId(v int64) *DescribeClusterMemberInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeClusterMemberInfoRequest) SetOwnerAccount(v string) *DescribeClusterMemberInfoRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeClusterMemberInfoRequest) SetInstanceId(v string) *DescribeClusterMemberInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeClusterMemberInfoRequest) SetFilterService(v string) *DescribeClusterMemberInfoRequest {
	s.FilterService = &v
	return s
}

func (s *DescribeClusterMemberInfoRequest) SetNeedReplica(v string) *DescribeClusterMemberInfoRequest {
	s.NeedReplica = &v
	return s
}

func (s *DescribeClusterMemberInfoRequest) SetPageNumber(v int32) *DescribeClusterMemberInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeClusterMemberInfoRequest) SetPageSize(v int32) *DescribeClusterMemberInfoRequest {
	s.PageSize = &v
	return s
}

type DescribeClusterMemberInfoResponseBody struct {
	RequestId       *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ClusterChildren []*DescribeClusterMemberInfoResponseBodyClusterChildren `json:"ClusterChildren,omitempty" xml:"ClusterChildren,omitempty" type:"Repeated"`
}

func (s DescribeClusterMemberInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterMemberInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterMemberInfoResponseBody) SetRequestId(v string) *DescribeClusterMemberInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBody) SetClusterChildren(v []*DescribeClusterMemberInfoResponseBodyClusterChildren) *DescribeClusterMemberInfoResponseBody {
	s.ClusterChildren = v
	return s
}

type DescribeClusterMemberInfoResponseBodyClusterChildren struct {
	Capacity            *int64  `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	UserId              *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DiskSizeMB          *int32  `json:"DiskSizeMB,omitempty" xml:"DiskSizeMB,omitempty"`
	BandWidth           *int64  `json:"BandWidth,omitempty" xml:"BandWidth,omitempty"`
	CurrentBandWidth    *int64  `json:"CurrentBandWidth,omitempty" xml:"CurrentBandWidth,omitempty"`
	ClassCode           *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	BizType             *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Service             *string `json:"Service,omitempty" xml:"Service,omitempty"`
	BinlogRetentionDays *int32  `json:"BinlogRetentionDays,omitempty" xml:"BinlogRetentionDays,omitempty"`
	Connections         *int64  `json:"Connections,omitempty" xml:"Connections,omitempty"`
	ResourceGroupName   *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	ServiceVersion      *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	ReplicaSize         *int32  `json:"ReplicaSize,omitempty" xml:"ReplicaSize,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id                  *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeClusterMemberInfoResponseBodyClusterChildren) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterMemberInfoResponseBodyClusterChildren) GoString() string {
	return s.String()
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetCapacity(v int64) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.Capacity = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetUserId(v string) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.UserId = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetDiskSizeMB(v int32) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.DiskSizeMB = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetBandWidth(v int64) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.BandWidth = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetCurrentBandWidth(v int64) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.CurrentBandWidth = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetClassCode(v string) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.ClassCode = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetBizType(v string) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.BizType = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetService(v string) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.Service = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetBinlogRetentionDays(v int32) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.BinlogRetentionDays = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetConnections(v int64) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.Connections = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetResourceGroupName(v string) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetServiceVersion(v string) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.ServiceVersion = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetReplicaSize(v int32) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.ReplicaSize = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetName(v string) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.Name = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetId(v int64) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.Id = &v
	return s
}

type DescribeClusterMemberInfoResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterMemberInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterMemberInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterMemberInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterMemberInfoResponse) SetHeaders(v map[string]*string) *DescribeClusterMemberInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterMemberInfoResponse) SetBody(v *DescribeClusterMemberInfoResponseBody) *DescribeClusterMemberInfoResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceNetInfoRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeDBInstanceNetInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceNetInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoRequest) SetSecurityToken(v string) *DescribeDBInstanceNetInfoRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDBInstanceNetInfoRequest) SetOwnerId(v int64) *DescribeDBInstanceNetInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoRequest) SetResourceOwnerAccount(v string) *DescribeDBInstanceNetInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceNetInfoRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceNetInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoRequest) SetOwnerAccount(v string) *DescribeDBInstanceNetInfoRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceNetInfoRequest) SetInstanceId(v string) *DescribeDBInstanceNetInfoRequest {
	s.InstanceId = &v
	return s
}

type DescribeDBInstanceNetInfoResponseBody struct {
	NetInfoItems        *DescribeDBInstanceNetInfoResponseBodyNetInfoItems `json:"NetInfoItems,omitempty" xml:"NetInfoItems,omitempty" type:"Struct"`
	RequestId           *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceNetworkType *string                                            `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetNetInfoItems(v *DescribeDBInstanceNetInfoResponseBodyNetInfoItems) *DescribeDBInstanceNetInfoResponseBody {
	s.NetInfoItems = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetRequestId(v string) *DescribeDBInstanceNetInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetInstanceNetworkType(v string) *DescribeDBInstanceNetInfoResponseBody {
	s.InstanceNetworkType = &v
	return s
}

type DescribeDBInstanceNetInfoResponseBodyNetInfoItems struct {
	InstanceNetInfo []*DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo `json:"InstanceNetInfo,omitempty" xml:"InstanceNetInfo,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceNetInfoResponseBodyNetInfoItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBodyNetInfoItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItems) SetInstanceNetInfo(v []*DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) *DescribeDBInstanceNetInfoResponseBodyNetInfoItems {
	s.InstanceNetInfo = v
	return s
}

type DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo struct {
	DirectConnection  *int32  `json:"DirectConnection,omitempty" xml:"DirectConnection,omitempty"`
	VSwitchId         *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	DBInstanceNetType *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	Upgradeable       *string `json:"Upgradeable,omitempty" xml:"Upgradeable,omitempty"`
	ExpiredTime       *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	ConnectionString  *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	IPType            *string `json:"IPType,omitempty" xml:"IPType,omitempty"`
	VPCInstanceId     *string `json:"VPCInstanceId,omitempty" xml:"VPCInstanceId,omitempty"`
	Port              *string `json:"Port,omitempty" xml:"Port,omitempty"`
	VPCId             *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	IPAddress         *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetDirectConnection(v int32) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.DirectConnection = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetVSwitchId(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetDBInstanceNetType(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.DBInstanceNetType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetUpgradeable(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.Upgradeable = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetExpiredTime(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetConnectionString(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetIPType(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.IPType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetVPCInstanceId(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.VPCInstanceId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetPort(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetVPCId(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.VPCId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetIPAddress(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.IPAddress = &v
	return s
}

type DescribeDBInstanceNetInfoResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBInstanceNetInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceNetInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceNetInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponse) SetBody(v *DescribeDBInstanceNetInfoResponseBody) *DescribeDBInstanceNetInfoResponse {
	s.Body = v
	return s
}

type DescribeDedicatedClusterInstanceListRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceStatus       *int32  `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	InstanceNetType      *string `json:"InstanceNetType,omitempty" xml:"InstanceNetType,omitempty"`
	Engine               *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion        *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DedicatedHostName    *string `json:"DedicatedHostName,omitempty" xml:"DedicatedHostName,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDedicatedClusterInstanceListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedClusterInstanceListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetSecurityToken(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetOwnerId(v int64) *DescribeDedicatedClusterInstanceListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetResourceOwnerAccount(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetResourceOwnerId(v int64) *DescribeDedicatedClusterInstanceListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetOwnerAccount(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetRegionId(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetInstanceId(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetInstanceStatus(v int32) *DescribeDedicatedClusterInstanceListRequest {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetInstanceNetType(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.InstanceNetType = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetEngine(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.Engine = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetEngineVersion(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetClusterId(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetDedicatedHostName(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.DedicatedHostName = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetPageNumber(v int32) *DescribeDedicatedClusterInstanceListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetPageSize(v int32) *DescribeDedicatedClusterInstanceListRequest {
	s.PageSize = &v
	return s
}

type DescribeDedicatedClusterInstanceListResponseBody struct {
	Instances  []*DescribeDedicatedClusterInstanceListResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	TotalCount *int32                                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeDedicatedClusterInstanceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedClusterInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedClusterInstanceListResponseBody) SetInstances(v []*DescribeDedicatedClusterInstanceListResponseBodyInstances) *DescribeDedicatedClusterInstanceListResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBody) SetTotalCount(v int32) *DescribeDedicatedClusterInstanceListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBody) SetPageSize(v int32) *DescribeDedicatedClusterInstanceListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBody) SetRequestId(v string) *DescribeDedicatedClusterInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBody) SetPageNumber(v int32) *DescribeDedicatedClusterInstanceListResponseBody {
	s.PageNumber = &v
	return s
}

type DescribeDedicatedClusterInstanceListResponseBodyInstances struct {
	VpcId             *string                                                                      `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	CharacterType     *int32                                                                       `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	VswitchId         *string                                                                      `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	MaintainStartTime *string                                                                      `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	InstanceClass     *string                                                                      `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	ConnectionDomain  *string                                                                      `json:"ConnectionDomain,omitempty" xml:"ConnectionDomain,omitempty"`
	CreateTime        *string                                                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	MaintainEndTime   *string                                                                      `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	StorageType       *string                                                                      `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	InstanceNodeList  []*DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList `json:"InstanceNodeList,omitempty" xml:"InstanceNodeList,omitempty" type:"Repeated"`
	InstanceId        *string                                                                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	BandWidth         *int64                                                                       `json:"BandWidth,omitempty" xml:"BandWidth,omitempty"`
	CurrentBandWidth  *int64                                                                       `json:"CurrentBandWidth,omitempty" xml:"CurrentBandWidth,omitempty"`
	EngineVersion     *string                                                                      `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	RegionId          *string                                                                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceName      *string                                                                      `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ZoneId            *string                                                                      `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ClusterName       *string                                                                      `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	InstanceStatus    *string                                                                      `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	Engine            *string                                                                      `json:"Engine,omitempty" xml:"Engine,omitempty"`
	ShardCount        *int32                                                                       `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
	CustomId          *string                                                                      `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	ClusterId         *string                                                                      `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeDedicatedClusterInstanceListResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedClusterInstanceListResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetVpcId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.VpcId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetCharacterType(v int32) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.CharacterType = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetVswitchId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.VswitchId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetMaintainStartTime(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetInstanceClass(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.InstanceClass = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetConnectionDomain(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.ConnectionDomain = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetCreateTime(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.CreateTime = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetMaintainEndTime(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetStorageType(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.StorageType = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetInstanceNodeList(v []*DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.InstanceNodeList = v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetInstanceId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetBandWidth(v int64) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.BandWidth = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetCurrentBandWidth(v int64) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.CurrentBandWidth = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetEngineVersion(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetRegionId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetInstanceName(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetZoneId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.ZoneId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetClusterName(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.ClusterName = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetInstanceStatus(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetEngine(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.Engine = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetShardCount(v int32) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.ShardCount = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetCustomId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.CustomId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetClusterId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.ClusterId = &v
	return s
}

type DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList struct {
	NodeIp            *string `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	DedicatedHostName *string `json:"DedicatedHostName,omitempty" xml:"DedicatedHostName,omitempty"`
	NodeType          *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	ZoneId            *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Port              *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Role              *string `json:"Role,omitempty" xml:"Role,omitempty"`
	NodeId            *int32  `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) SetNodeIp(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList {
	s.NodeIp = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) SetDedicatedHostName(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList {
	s.DedicatedHostName = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) SetNodeType(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList {
	s.NodeType = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) SetZoneId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList {
	s.ZoneId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) SetInstanceId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList {
	s.InstanceId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) SetPort(v int32) *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList {
	s.Port = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) SetRole(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList {
	s.Role = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) SetNodeId(v int32) *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList {
	s.NodeId = &v
	return s
}

type DescribeDedicatedClusterInstanceListResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDedicatedClusterInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDedicatedClusterInstanceListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedClusterInstanceListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedClusterInstanceListResponse) SetHeaders(v map[string]*string) *DescribeDedicatedClusterInstanceListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponse) SetBody(v *DescribeDedicatedClusterInstanceListResponseBody) *DescribeDedicatedClusterInstanceListResponse {
	s.Body = v
	return s
}

type DescribeEngineVersionRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Parameters           *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s DescribeEngineVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEngineVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeEngineVersionRequest) SetSecurityToken(v string) *DescribeEngineVersionRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeEngineVersionRequest) SetOwnerId(v int64) *DescribeEngineVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeEngineVersionRequest) SetResourceOwnerAccount(v string) *DescribeEngineVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeEngineVersionRequest) SetResourceOwnerId(v int64) *DescribeEngineVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeEngineVersionRequest) SetOwnerAccount(v string) *DescribeEngineVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeEngineVersionRequest) SetInstanceId(v string) *DescribeEngineVersionRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeEngineVersionRequest) SetParameters(v string) *DescribeEngineVersionRequest {
	s.Parameters = &v
	return s
}

type DescribeEngineVersionResponseBody struct {
	MajorVersion              *string `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	IsLatestVersion           *bool   `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
	RequestId                 *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MinorVersion              *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	EnableUpgradeMinorVersion *bool   `json:"EnableUpgradeMinorVersion,omitempty" xml:"EnableUpgradeMinorVersion,omitempty"`
	Engine                    *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EnableUpgradeMajorVersion *bool   `json:"EnableUpgradeMajorVersion,omitempty" xml:"EnableUpgradeMajorVersion,omitempty"`
}

func (s DescribeEngineVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEngineVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEngineVersionResponseBody) SetMajorVersion(v string) *DescribeEngineVersionResponseBody {
	s.MajorVersion = &v
	return s
}

func (s *DescribeEngineVersionResponseBody) SetIsLatestVersion(v bool) *DescribeEngineVersionResponseBody {
	s.IsLatestVersion = &v
	return s
}

func (s *DescribeEngineVersionResponseBody) SetRequestId(v string) *DescribeEngineVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEngineVersionResponseBody) SetMinorVersion(v string) *DescribeEngineVersionResponseBody {
	s.MinorVersion = &v
	return s
}

func (s *DescribeEngineVersionResponseBody) SetEnableUpgradeMinorVersion(v bool) *DescribeEngineVersionResponseBody {
	s.EnableUpgradeMinorVersion = &v
	return s
}

func (s *DescribeEngineVersionResponseBody) SetEngine(v string) *DescribeEngineVersionResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeEngineVersionResponseBody) SetEnableUpgradeMajorVersion(v bool) *DescribeEngineVersionResponseBody {
	s.EnableUpgradeMajorVersion = &v
	return s
}

type DescribeEngineVersionResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEngineVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEngineVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEngineVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeEngineVersionResponse) SetHeaders(v map[string]*string) *DescribeEngineVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeEngineVersionResponse) SetBody(v *DescribeEngineVersionResponseBody) *DescribeEngineVersionResponse {
	s.Body = v
	return s
}

type DescribeGlobalDistributeCacheRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	GlobalInstanceId     *string `json:"GlobalInstanceId,omitempty" xml:"GlobalInstanceId,omitempty"`
	PageNumber           *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SubInstanceId        *string `json:"SubInstanceId,omitempty" xml:"SubInstanceId,omitempty"`
}

func (s DescribeGlobalDistributeCacheRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDistributeCacheRequest) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDistributeCacheRequest) SetSecurityToken(v string) *DescribeGlobalDistributeCacheRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeGlobalDistributeCacheRequest) SetOwnerId(v int64) *DescribeGlobalDistributeCacheRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeGlobalDistributeCacheRequest) SetResourceOwnerAccount(v string) *DescribeGlobalDistributeCacheRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeGlobalDistributeCacheRequest) SetResourceOwnerId(v int64) *DescribeGlobalDistributeCacheRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeGlobalDistributeCacheRequest) SetOwnerAccount(v string) *DescribeGlobalDistributeCacheRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeGlobalDistributeCacheRequest) SetGlobalInstanceId(v string) *DescribeGlobalDistributeCacheRequest {
	s.GlobalInstanceId = &v
	return s
}

func (s *DescribeGlobalDistributeCacheRequest) SetPageNumber(v string) *DescribeGlobalDistributeCacheRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGlobalDistributeCacheRequest) SetPageSize(v string) *DescribeGlobalDistributeCacheRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGlobalDistributeCacheRequest) SetSubInstanceId(v string) *DescribeGlobalDistributeCacheRequest {
	s.SubInstanceId = &v
	return s
}

type DescribeGlobalDistributeCacheResponseBody struct {
	TotalRecordCount       *int32                                                             `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	GlobalDistributeCaches []*DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCaches `json:"GlobalDistributeCaches,omitempty" xml:"GlobalDistributeCaches,omitempty" type:"Repeated"`
	PageSize               *int32                                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId              *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber             *int32                                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeGlobalDistributeCacheResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDistributeCacheResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDistributeCacheResponseBody) SetTotalRecordCount(v int32) *DescribeGlobalDistributeCacheResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeGlobalDistributeCacheResponseBody) SetGlobalDistributeCaches(v []*DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCaches) *DescribeGlobalDistributeCacheResponseBody {
	s.GlobalDistributeCaches = v
	return s
}

func (s *DescribeGlobalDistributeCacheResponseBody) SetPageSize(v int32) *DescribeGlobalDistributeCacheResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGlobalDistributeCacheResponseBody) SetRequestId(v string) *DescribeGlobalDistributeCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGlobalDistributeCacheResponseBody) SetPageNumber(v int32) *DescribeGlobalDistributeCacheResponseBody {
	s.PageNumber = &v
	return s
}

type DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCaches struct {
	Status           *string                                                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	SubInstances     []*DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances `json:"SubInstances,omitempty" xml:"SubInstances,omitempty" type:"Repeated"`
	GlobalInstanceId *string                                                                        `json:"GlobalInstanceId,omitempty" xml:"GlobalInstanceId,omitempty"`
}

func (s DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCaches) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCaches) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCaches) SetStatus(v string) *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCaches {
	s.Status = &v
	return s
}

func (s *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCaches) SetSubInstances(v []*DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances) *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCaches {
	s.SubInstances = v
	return s
}

func (s *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCaches) SetGlobalInstanceId(v string) *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCaches {
	s.GlobalInstanceId = &v
	return s
}

type DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances struct {
	InstanceClass    *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	InstanceStatus   *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	InstanceID       *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	GlobalInstanceId *string `json:"GlobalInstanceId,omitempty" xml:"GlobalInstanceId,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances) SetInstanceClass(v string) *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances {
	s.InstanceClass = &v
	return s
}

func (s *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances) SetInstanceStatus(v string) *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances) SetInstanceID(v string) *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances {
	s.InstanceID = &v
	return s
}

func (s *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances) SetGlobalInstanceId(v string) *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances {
	s.GlobalInstanceId = &v
	return s
}

func (s *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances) SetRegionId(v string) *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances {
	s.RegionId = &v
	return s
}

type DescribeGlobalDistributeCacheResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGlobalDistributeCacheResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGlobalDistributeCacheResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDistributeCacheResponse) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDistributeCacheResponse) SetHeaders(v map[string]*string) *DescribeGlobalDistributeCacheResponse {
	s.Headers = v
	return s
}

func (s *DescribeGlobalDistributeCacheResponse) SetBody(v *DescribeGlobalDistributeCacheResponseBody) *DescribeGlobalDistributeCacheResponse {
	s.Body = v
	return s
}

type DescribeHistoryMonitorValuesRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IntervalForHistory   *string `json:"IntervalForHistory,omitempty" xml:"IntervalForHistory,omitempty"`
	MonitorKeys          *string `json:"MonitorKeys,omitempty" xml:"MonitorKeys,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	AccessType           *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	Product              *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Category             *string `json:"Category,omitempty" xml:"Category,omitempty"`
	ReplicatorJobId      *string `json:"ReplicatorJobId,omitempty" xml:"ReplicatorJobId,omitempty"`
}

func (s DescribeHistoryMonitorValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryMonitorValuesRequest) GoString() string {
	return s.String()
}

func (s *DescribeHistoryMonitorValuesRequest) SetSecurityToken(v string) *DescribeHistoryMonitorValuesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeHistoryMonitorValuesRequest) SetOwnerId(v int64) *DescribeHistoryMonitorValuesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeHistoryMonitorValuesRequest) SetResourceOwnerAccount(v string) *DescribeHistoryMonitorValuesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeHistoryMonitorValuesRequest) SetResourceOwnerId(v int64) *DescribeHistoryMonitorValuesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeHistoryMonitorValuesRequest) SetOwnerAccount(v string) *DescribeHistoryMonitorValuesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeHistoryMonitorValuesRequest) SetInstanceId(v string) *DescribeHistoryMonitorValuesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHistoryMonitorValuesRequest) SetStartTime(v string) *DescribeHistoryMonitorValuesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeHistoryMonitorValuesRequest) SetEndTime(v string) *DescribeHistoryMonitorValuesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeHistoryMonitorValuesRequest) SetIntervalForHistory(v string) *DescribeHistoryMonitorValuesRequest {
	s.IntervalForHistory = &v
	return s
}

func (s *DescribeHistoryMonitorValuesRequest) SetMonitorKeys(v string) *DescribeHistoryMonitorValuesRequest {
	s.MonitorKeys = &v
	return s
}

func (s *DescribeHistoryMonitorValuesRequest) SetNodeId(v string) *DescribeHistoryMonitorValuesRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeHistoryMonitorValuesRequest) SetAccessType(v string) *DescribeHistoryMonitorValuesRequest {
	s.AccessType = &v
	return s
}

func (s *DescribeHistoryMonitorValuesRequest) SetProduct(v string) *DescribeHistoryMonitorValuesRequest {
	s.Product = &v
	return s
}

func (s *DescribeHistoryMonitorValuesRequest) SetCategory(v string) *DescribeHistoryMonitorValuesRequest {
	s.Category = &v
	return s
}

func (s *DescribeHistoryMonitorValuesRequest) SetReplicatorJobId(v string) *DescribeHistoryMonitorValuesRequest {
	s.ReplicatorJobId = &v
	return s
}

type DescribeHistoryMonitorValuesResponseBody struct {
	MonitorHistory *string `json:"MonitorHistory,omitempty" xml:"MonitorHistory,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHistoryMonitorValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryMonitorValuesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHistoryMonitorValuesResponseBody) SetMonitorHistory(v string) *DescribeHistoryMonitorValuesResponseBody {
	s.MonitorHistory = &v
	return s
}

func (s *DescribeHistoryMonitorValuesResponseBody) SetRequestId(v string) *DescribeHistoryMonitorValuesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeHistoryMonitorValuesResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHistoryMonitorValuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHistoryMonitorValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryMonitorValuesResponse) GoString() string {
	return s.String()
}

func (s *DescribeHistoryMonitorValuesResponse) SetHeaders(v map[string]*string) *DescribeHistoryMonitorValuesResponse {
	s.Headers = v
	return s
}

func (s *DescribeHistoryMonitorValuesResponse) SetBody(v *DescribeHistoryMonitorValuesResponseBody) *DescribeHistoryMonitorValuesResponse {
	s.Body = v
	return s
}

type DescribeInstanceAttributeRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeRequest) SetSecurityToken(v string) *DescribeInstanceAttributeRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeInstanceAttributeRequest) SetOwnerId(v int64) *DescribeInstanceAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceAttributeRequest) SetResourceOwnerAccount(v string) *DescribeInstanceAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceAttributeRequest) SetResourceOwnerId(v int64) *DescribeInstanceAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceAttributeRequest) SetOwnerAccount(v string) *DescribeInstanceAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstanceAttributeRequest) SetInstanceId(v string) *DescribeInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

type DescribeInstanceAttributeResponseBody struct {
	Instances *DescribeInstanceAttributeResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBody) SetInstances(v *DescribeInstanceAttributeResponseBodyInstances) *DescribeInstanceAttributeResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetRequestId(v string) *DescribeInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceAttributeResponseBodyInstances struct {
	DBInstanceAttribute []*DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute `json:"DBInstanceAttribute,omitempty" xml:"DBInstanceAttribute,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAttributeResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyInstances) SetDBInstanceAttribute(v []*DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) *DescribeInstanceAttributeResponseBodyInstances {
	s.DBInstanceAttribute = v
	return s
}

type DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute struct {
	VpcId                     *string                                                                `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	PrivateIp                 *string                                                                `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	Capacity                  *int64                                                                 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	CreateTime                *string                                                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ConnectionDomain          *string                                                                `json:"ConnectionDomain,omitempty" xml:"ConnectionDomain,omitempty"`
	IsRds                     *bool                                                                  `json:"IsRds,omitempty" xml:"IsRds,omitempty"`
	ChargeType                *string                                                                `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Tags                      *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VpcAuthMode               *string                                                                `json:"VpcAuthMode,omitempty" xml:"VpcAuthMode,omitempty"`
	ArchitectureType          *string                                                                `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	NetworkType               *string                                                                `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	AvailabilityValue         *string                                                                `json:"AvailabilityValue,omitempty" xml:"AvailabilityValue,omitempty"`
	Port                      *int64                                                                 `json:"Port,omitempty" xml:"Port,omitempty"`
	Config                    *string                                                                `json:"Config,omitempty" xml:"Config,omitempty"`
	EngineVersion             *string                                                                `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	PackageType               *string                                                                `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	VpcCloudInstanceId        *string                                                                `json:"VpcCloudInstanceId,omitempty" xml:"VpcCloudInstanceId,omitempty"`
	Bandwidth                 *int64                                                                 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	InstanceName              *string                                                                `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	SecurityIPList            *string                                                                `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	ShardCount                *int32                                                                 `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
	GlobalInstanceId          *string                                                                `json:"GlobalInstanceId,omitempty" xml:"GlobalInstanceId,omitempty"`
	QPS                       *int64                                                                 `json:"QPS,omitempty" xml:"QPS,omitempty"`
	AuditLogRetention         *string                                                                `json:"AuditLogRetention,omitempty" xml:"AuditLogRetention,omitempty"`
	MaintainStartTime         *string                                                                `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	InstanceClass             *string                                                                `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	MaintainEndTime           *string                                                                `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	InstanceId                *string                                                                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType              *string                                                                `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	HasRenewChangeOrder       *string                                                                `json:"HasRenewChangeOrder,omitempty" xml:"HasRenewChangeOrder,omitempty"`
	InstanceReleaseProtection *bool                                                                  `json:"InstanceReleaseProtection,omitempty" xml:"InstanceReleaseProtection,omitempty"`
	ReplicationMode           *string                                                                `json:"ReplicationMode,omitempty" xml:"ReplicationMode,omitempty"`
	RegionId                  *string                                                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	EndTime                   *string                                                                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	VSwitchId                 *string                                                                `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ReplicaId                 *string                                                                `json:"ReplicaId,omitempty" xml:"ReplicaId,omitempty"`
	NodeType                  *string                                                                `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	Connections               *int64                                                                 `json:"Connections,omitempty" xml:"Connections,omitempty"`
	ResourceGroupId           *string                                                                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ZoneId                    *string                                                                `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	InstanceStatus            *string                                                                `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	Engine                    *string                                                                `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetVpcId(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.VpcId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetPrivateIp(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.PrivateIp = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetCapacity(v int64) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.Capacity = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetCreateTime(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.CreateTime = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetConnectionDomain(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.ConnectionDomain = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetIsRds(v bool) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.IsRds = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetChargeType(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.ChargeType = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetTags(v *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTags) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.Tags = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetVpcAuthMode(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.VpcAuthMode = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetArchitectureType(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.ArchitectureType = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetNetworkType(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.NetworkType = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetAvailabilityValue(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.AvailabilityValue = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetPort(v int64) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.Port = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetConfig(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.Config = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetEngineVersion(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.EngineVersion = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetPackageType(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.PackageType = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetVpcCloudInstanceId(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.VpcCloudInstanceId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetBandwidth(v int64) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.Bandwidth = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetInstanceName(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetSecurityIPList(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.SecurityIPList = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetShardCount(v int32) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.ShardCount = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetGlobalInstanceId(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.GlobalInstanceId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetQPS(v int64) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.QPS = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetAuditLogRetention(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.AuditLogRetention = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetMaintainStartTime(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetInstanceClass(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.InstanceClass = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetMaintainEndTime(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetInstanceId(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetInstanceType(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetHasRenewChangeOrder(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.HasRenewChangeOrder = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetInstanceReleaseProtection(v bool) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.InstanceReleaseProtection = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetReplicationMode(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.ReplicationMode = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetRegionId(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetEndTime(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.EndTime = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetVSwitchId(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.VSwitchId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetReplicaId(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.ReplicaId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetNodeType(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.NodeType = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetConnections(v int64) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.Connections = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetResourceGroupId(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetZoneId(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.ZoneId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetInstanceStatus(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetEngine(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.Engine = &v
	return s
}

type DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTags struct {
	Tag []*DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTags) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTags) SetTag(v []*DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTags {
	s.Tag = v
	return s
}

type DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag) SetKey(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag) SetValue(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag {
	s.Value = &v
	return s
}

type DescribeInstanceAttributeResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponse) SetHeaders(v map[string]*string) *DescribeInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceAttributeResponse) SetBody(v *DescribeInstanceAttributeResponseBody) *DescribeInstanceAttributeResponse {
	s.Body = v
	return s
}

type DescribeInstanceAutoRenewalAttributeRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ProxyId              *string `json:"proxyId,omitempty" xml:"proxyId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	Category             *string `json:"Category,omitempty" xml:"Category,omitempty"`
}

func (s DescribeInstanceAutoRenewalAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAutoRenewalAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetOwnerId(v int64) *DescribeInstanceAutoRenewalAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetResourceOwnerAccount(v string) *DescribeInstanceAutoRenewalAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetResourceOwnerId(v int64) *DescribeInstanceAutoRenewalAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetClientToken(v string) *DescribeInstanceAutoRenewalAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetProxyId(v string) *DescribeInstanceAutoRenewalAttributeRequest {
	s.ProxyId = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetRegionId(v string) *DescribeInstanceAutoRenewalAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetDBInstanceId(v string) *DescribeInstanceAutoRenewalAttributeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetPageSize(v int32) *DescribeInstanceAutoRenewalAttributeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetPageNumber(v int32) *DescribeInstanceAutoRenewalAttributeRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetOwnerAccount(v string) *DescribeInstanceAutoRenewalAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetCategory(v string) *DescribeInstanceAutoRenewalAttributeRequest {
	s.Category = &v
	return s
}

type DescribeInstanceAutoRenewalAttributeResponseBody struct {
	TotalRecordCount *int32                                                 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	PageRecordCount  *int32                                                 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber       *int32                                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Items            *DescribeInstanceAutoRenewalAttributeResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeInstanceAutoRenewalAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAutoRenewalAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) SetTotalRecordCount(v int32) *DescribeInstanceAutoRenewalAttributeResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) SetPageRecordCount(v int32) *DescribeInstanceAutoRenewalAttributeResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) SetRequestId(v string) *DescribeInstanceAutoRenewalAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) SetPageNumber(v int32) *DescribeInstanceAutoRenewalAttributeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) SetItems(v *DescribeInstanceAutoRenewalAttributeResponseBodyItems) *DescribeInstanceAutoRenewalAttributeResponseBody {
	s.Items = v
	return s
}

type DescribeInstanceAutoRenewalAttributeResponseBodyItems struct {
	Item []*DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAutoRenewalAttributeResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAutoRenewalAttributeResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyItems) SetItem(v []*DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem) *DescribeInstanceAutoRenewalAttributeResponseBodyItems {
	s.Item = v
	return s
}

type DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem struct {
	AutoRenew    *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	Duration     *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem) SetAutoRenew(v string) *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem {
	s.AutoRenew = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem) SetDuration(v int32) *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem {
	s.Duration = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem) SetDBInstanceId(v string) *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem) SetRegionId(v string) *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem {
	s.RegionId = &v
	return s
}

type DescribeInstanceAutoRenewalAttributeResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceAutoRenewalAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceAutoRenewalAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAutoRenewalAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewalAttributeResponse) SetHeaders(v map[string]*string) *DescribeInstanceAutoRenewalAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponse) SetBody(v *DescribeInstanceAutoRenewalAttributeResponseBody) *DescribeInstanceAutoRenewalAttributeResponse {
	s.Body = v
	return s
}

type DescribeInstanceConfigRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceConfigRequest) SetSecurityToken(v string) *DescribeInstanceConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeInstanceConfigRequest) SetOwnerId(v int64) *DescribeInstanceConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceConfigRequest) SetResourceOwnerAccount(v string) *DescribeInstanceConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceConfigRequest) SetResourceOwnerId(v int64) *DescribeInstanceConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceConfigRequest) SetOwnerAccount(v string) *DescribeInstanceConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstanceConfigRequest) SetInstanceId(v string) *DescribeInstanceConfigRequest {
	s.InstanceId = &v
	return s
}

type DescribeInstanceConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Config    *string `json:"Config,omitempty" xml:"Config,omitempty"`
}

func (s DescribeInstanceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceConfigResponseBody) SetRequestId(v string) *DescribeInstanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceConfigResponseBody) SetConfig(v string) *DescribeInstanceConfigResponseBody {
	s.Config = &v
	return s
}

type DescribeInstanceConfigResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceConfigResponse) SetHeaders(v map[string]*string) *DescribeInstanceConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceConfigResponse) SetBody(v *DescribeInstanceConfigResponseBody) *DescribeInstanceConfigResponse {
	s.Body = v
	return s
}

type DescribeInstancesRequest struct {
	SecurityToken        *string                        `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                        `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                         `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string                        `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceIds          *string                        `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	InstanceStatus       *string                        `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	ChargeType           *string                        `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	NetworkType          *string                        `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	EngineVersion        *string                        `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	InstanceClass        *string                        `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	VpcId                *string                        `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VSwitchId            *string                        `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	PageNumber           *int32                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	InstanceType         *string                        `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	SearchKey            *string                        `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	ArchitectureType     *string                        `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	Expired              *string                        `json:"Expired,omitempty" xml:"Expired,omitempty"`
	ZoneId               *string                        `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	VpcCloudInsInfo      *int32                         `json:"VpcCloudInsInfo,omitempty" xml:"VpcCloudInsInfo,omitempty"`
	ResourceGroupId      *string                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	GlobalInstance       *bool                          `json:"GlobalInstance,omitempty" xml:"GlobalInstance,omitempty"`
	EditionType          *string                        `json:"EditionType,omitempty" xml:"EditionType,omitempty"`
	Tag                  []*DescribeInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) SetSecurityToken(v string) *DescribeInstancesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeInstancesRequest) SetOwnerId(v int64) *DescribeInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstancesRequest) SetResourceOwnerAccount(v string) *DescribeInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstancesRequest) SetResourceOwnerId(v int64) *DescribeInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstancesRequest) SetOwnerAccount(v string) *DescribeInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstancesRequest) SetRegionId(v string) *DescribeInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceIds(v string) *DescribeInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceStatus(v string) *DescribeInstancesRequest {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstancesRequest) SetChargeType(v string) *DescribeInstancesRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeInstancesRequest) SetNetworkType(v string) *DescribeInstancesRequest {
	s.NetworkType = &v
	return s
}

func (s *DescribeInstancesRequest) SetEngineVersion(v string) *DescribeInstancesRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceClass(v string) *DescribeInstancesRequest {
	s.InstanceClass = &v
	return s
}

func (s *DescribeInstancesRequest) SetVpcId(v string) *DescribeInstancesRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeInstancesRequest) SetVSwitchId(v string) *DescribeInstancesRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageNumber(v int32) *DescribeInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageSize(v int32) *DescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceType(v string) *DescribeInstancesRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstancesRequest) SetSearchKey(v string) *DescribeInstancesRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeInstancesRequest) SetArchitectureType(v string) *DescribeInstancesRequest {
	s.ArchitectureType = &v
	return s
}

func (s *DescribeInstancesRequest) SetExpired(v string) *DescribeInstancesRequest {
	s.Expired = &v
	return s
}

func (s *DescribeInstancesRequest) SetZoneId(v string) *DescribeInstancesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeInstancesRequest) SetVpcCloudInsInfo(v int32) *DescribeInstancesRequest {
	s.VpcCloudInsInfo = &v
	return s
}

func (s *DescribeInstancesRequest) SetResourceGroupId(v string) *DescribeInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesRequest) SetGlobalInstance(v bool) *DescribeInstancesRequest {
	s.GlobalInstance = &v
	return s
}

func (s *DescribeInstancesRequest) SetEditionType(v string) *DescribeInstancesRequest {
	s.EditionType = &v
	return s
}

func (s *DescribeInstancesRequest) SetTag(v []*DescribeInstancesRequestTag) *DescribeInstancesRequest {
	s.Tag = v
	return s
}

type DescribeInstancesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequestTag) SetKey(v string) *DescribeInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeInstancesRequestTag) SetValue(v string) *DescribeInstancesRequestTag {
	s.Value = &v
	return s
}

type DescribeInstancesResponseBody struct {
	Instances  *DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) SetInstances(v *DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotalCount(v int32) *DescribeInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageSize(v int32) *DescribeInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageNumber(v int32) *DescribeInstancesResponseBody {
	s.PageNumber = &v
	return s
}

type DescribeInstancesResponseBodyInstances struct {
	KVStoreInstance []*DescribeInstancesResponseBodyInstancesKVStoreInstance `json:"KVStoreInstance,omitempty" xml:"KVStoreInstance,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) SetKVStoreInstance(v []*DescribeInstancesResponseBodyInstancesKVStoreInstance) *DescribeInstancesResponseBodyInstances {
	s.KVStoreInstance = v
	return s
}

type DescribeInstancesResponseBodyInstancesKVStoreInstance struct {
	VpcId               *string                                                    `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	PrivateIp           *string                                                    `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	Capacity            *int64                                                     `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	ReplacateId         *string                                                    `json:"ReplacateId,omitempty" xml:"ReplacateId,omitempty"`
	CreateTime          *string                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ConnectionDomain    *string                                                    `json:"ConnectionDomain,omitempty" xml:"ConnectionDomain,omitempty"`
	IsRds               *bool                                                      `json:"IsRds,omitempty" xml:"IsRds,omitempty"`
	ChargeType          *string                                                    `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Tags                *DescribeInstancesResponseBodyInstancesKVStoreInstanceTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	ArchitectureType    *string                                                    `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	NetworkType         *string                                                    `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Port                *int64                                                     `json:"Port,omitempty" xml:"Port,omitempty"`
	ConnectionMode      *string                                                    `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
	Config              *string                                                    `json:"Config,omitempty" xml:"Config,omitempty"`
	PackageType         *string                                                    `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	EngineVersion       *string                                                    `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	Bandwidth           *int64                                                     `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	InstanceName        *string                                                    `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ShardCount          *int32                                                     `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
	UserName            *string                                                    `json:"UserName,omitempty" xml:"UserName,omitempty"`
	QPS                 *int64                                                     `json:"QPS,omitempty" xml:"QPS,omitempty"`
	InstanceClass       *string                                                    `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	DestroyTime         *string                                                    `json:"DestroyTime,omitempty" xml:"DestroyTime,omitempty"`
	InstanceId          *string                                                    `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType        *string                                                    `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	HasRenewChangeOrder *bool                                                      `json:"HasRenewChangeOrder,omitempty" xml:"HasRenewChangeOrder,omitempty"`
	RegionId            *string                                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SearchKey           *string                                                    `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	EndTime             *string                                                    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	VSwitchId           *string                                                    `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	NodeType            *string                                                    `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	Connections         *int64                                                     `json:"Connections,omitempty" xml:"Connections,omitempty"`
	ResourceGroupId     *string                                                    `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ZoneId              *string                                                    `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	InstanceStatus      *string                                                    `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesKVStoreInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesKVStoreInstance) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetVpcId(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.VpcId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetPrivateIp(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.PrivateIp = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetCapacity(v int64) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.Capacity = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetReplacateId(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.ReplacateId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetCreateTime(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetConnectionDomain(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.ConnectionDomain = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetIsRds(v bool) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.IsRds = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetChargeType(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.ChargeType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetTags(v *DescribeInstancesResponseBodyInstancesKVStoreInstanceTags) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.Tags = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetArchitectureType(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.ArchitectureType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetNetworkType(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetPort(v int64) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.Port = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetConnectionMode(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.ConnectionMode = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetConfig(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.Config = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetPackageType(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.PackageType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetEngineVersion(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetBandwidth(v int64) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.Bandwidth = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetInstanceName(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetShardCount(v int32) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.ShardCount = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetUserName(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.UserName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetQPS(v int64) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.QPS = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetInstanceClass(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.InstanceClass = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetDestroyTime(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.DestroyTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetInstanceId(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetInstanceType(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetHasRenewChangeOrder(v bool) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.HasRenewChangeOrder = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetRegionId(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetSearchKey(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.SearchKey = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetEndTime(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.EndTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetVSwitchId(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.VSwitchId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetNodeType(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.NodeType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetConnections(v int64) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.Connections = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetResourceGroupId(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetZoneId(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetInstanceStatus(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.InstanceStatus = &v
	return s
}

type DescribeInstancesResponseBodyInstancesKVStoreInstanceTags struct {
	Tag []*DescribeInstancesResponseBodyInstancesKVStoreInstanceTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesKVStoreInstanceTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesKVStoreInstanceTags) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstanceTags) SetTag(v []*DescribeInstancesResponseBodyInstancesKVStoreInstanceTagsTag) *DescribeInstancesResponseBodyInstancesKVStoreInstanceTags {
	s.Tag = v
	return s
}

type DescribeInstancesResponseBodyInstancesKVStoreInstanceTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesKVStoreInstanceTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesKVStoreInstanceTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstanceTagsTag) SetKey(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstanceTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstanceTagsTag) SetValue(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstanceTagsTag {
	s.Value = &v
	return s
}

type DescribeInstancesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponse) SetHeaders(v map[string]*string) *DescribeInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstancesResponse) SetBody(v *DescribeInstancesResponseBody) *DescribeInstancesResponse {
	s.Body = v
	return s
}

type DescribeInstanceSSLRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceSSLRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSSLRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSSLRequest) SetSecurityToken(v string) *DescribeInstanceSSLRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeInstanceSSLRequest) SetOwnerId(v int64) *DescribeInstanceSSLRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceSSLRequest) SetResourceOwnerAccount(v string) *DescribeInstanceSSLRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceSSLRequest) SetResourceOwnerId(v int64) *DescribeInstanceSSLRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceSSLRequest) SetOwnerAccount(v string) *DescribeInstanceSSLRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstanceSSLRequest) SetInstanceId(v string) *DescribeInstanceSSLRequest {
	s.InstanceId = &v
	return s
}

type DescribeInstanceSSLResponseBody struct {
	SSLExpiredTime *string `json:"SSLExpiredTime,omitempty" xml:"SSLExpiredTime,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SSLEnabled     *string `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CertCommonName *string `json:"CertCommonName,omitempty" xml:"CertCommonName,omitempty"`
}

func (s DescribeInstanceSSLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSSLResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSSLResponseBody) SetSSLExpiredTime(v string) *DescribeInstanceSSLResponseBody {
	s.SSLExpiredTime = &v
	return s
}

func (s *DescribeInstanceSSLResponseBody) SetRequestId(v string) *DescribeInstanceSSLResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceSSLResponseBody) SetSSLEnabled(v string) *DescribeInstanceSSLResponseBody {
	s.SSLEnabled = &v
	return s
}

func (s *DescribeInstanceSSLResponseBody) SetInstanceId(v string) *DescribeInstanceSSLResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceSSLResponseBody) SetCertCommonName(v string) *DescribeInstanceSSLResponseBody {
	s.CertCommonName = &v
	return s
}

type DescribeInstanceSSLResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceSSLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceSSLResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSSLResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSSLResponse) SetHeaders(v map[string]*string) *DescribeInstanceSSLResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceSSLResponse) SetBody(v *DescribeInstanceSSLResponseBody) *DescribeInstanceSSLResponse {
	s.Body = v
	return s
}

type DescribeIntranetAttributeRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeIntranetAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIntranetAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeIntranetAttributeRequest) SetSecurityToken(v string) *DescribeIntranetAttributeRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeIntranetAttributeRequest) SetOwnerId(v int64) *DescribeIntranetAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeIntranetAttributeRequest) SetResourceOwnerAccount(v string) *DescribeIntranetAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeIntranetAttributeRequest) SetResourceOwnerId(v int64) *DescribeIntranetAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeIntranetAttributeRequest) SetOwnerAccount(v string) *DescribeIntranetAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeIntranetAttributeRequest) SetInstanceId(v string) *DescribeIntranetAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeIntranetAttributeRequest) SetResourceGroupId(v string) *DescribeIntranetAttributeRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeIntranetAttributeResponseBody struct {
	IntranetBandwidth   *int32  `json:"IntranetBandwidth,omitempty" xml:"IntranetBandwidth,omitempty"`
	BandwidthExpireTime *string `json:"BandwidthExpireTime,omitempty" xml:"BandwidthExpireTime,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ExpireTime          *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
}

func (s DescribeIntranetAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIntranetAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIntranetAttributeResponseBody) SetIntranetBandwidth(v int32) *DescribeIntranetAttributeResponseBody {
	s.IntranetBandwidth = &v
	return s
}

func (s *DescribeIntranetAttributeResponseBody) SetBandwidthExpireTime(v string) *DescribeIntranetAttributeResponseBody {
	s.BandwidthExpireTime = &v
	return s
}

func (s *DescribeIntranetAttributeResponseBody) SetRequestId(v string) *DescribeIntranetAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIntranetAttributeResponseBody) SetExpireTime(v string) *DescribeIntranetAttributeResponseBody {
	s.ExpireTime = &v
	return s
}

type DescribeIntranetAttributeResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeIntranetAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeIntranetAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIntranetAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeIntranetAttributeResponse) SetHeaders(v map[string]*string) *DescribeIntranetAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeIntranetAttributeResponse) SetBody(v *DescribeIntranetAttributeResponseBody) *DescribeIntranetAttributeResponse {
	s.Body = v
	return s
}

type DescribeLogicInstanceTopologyRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeLogicInstanceTopologyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogicInstanceTopologyRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogicInstanceTopologyRequest) SetSecurityToken(v string) *DescribeLogicInstanceTopologyRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLogicInstanceTopologyRequest) SetOwnerId(v int64) *DescribeLogicInstanceTopologyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLogicInstanceTopologyRequest) SetResourceOwnerAccount(v string) *DescribeLogicInstanceTopologyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLogicInstanceTopologyRequest) SetResourceOwnerId(v int64) *DescribeLogicInstanceTopologyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLogicInstanceTopologyRequest) SetOwnerAccount(v string) *DescribeLogicInstanceTopologyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLogicInstanceTopologyRequest) SetInstanceId(v string) *DescribeLogicInstanceTopologyRequest {
	s.InstanceId = &v
	return s
}

type DescribeLogicInstanceTopologyResponseBody struct {
	RedisShardList *DescribeLogicInstanceTopologyResponseBodyRedisShardList `json:"RedisShardList,omitempty" xml:"RedisShardList,omitempty" type:"Struct"`
	RequestId      *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceId     *string                                                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RedisProxyList *DescribeLogicInstanceTopologyResponseBodyRedisProxyList `json:"RedisProxyList,omitempty" xml:"RedisProxyList,omitempty" type:"Struct"`
}

func (s DescribeLogicInstanceTopologyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogicInstanceTopologyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogicInstanceTopologyResponseBody) SetRedisShardList(v *DescribeLogicInstanceTopologyResponseBodyRedisShardList) *DescribeLogicInstanceTopologyResponseBody {
	s.RedisShardList = v
	return s
}

func (s *DescribeLogicInstanceTopologyResponseBody) SetRequestId(v string) *DescribeLogicInstanceTopologyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogicInstanceTopologyResponseBody) SetInstanceId(v string) *DescribeLogicInstanceTopologyResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeLogicInstanceTopologyResponseBody) SetRedisProxyList(v *DescribeLogicInstanceTopologyResponseBodyRedisProxyList) *DescribeLogicInstanceTopologyResponseBody {
	s.RedisProxyList = v
	return s
}

type DescribeLogicInstanceTopologyResponseBodyRedisShardList struct {
	NodeInfo []*DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Repeated"`
}

func (s DescribeLogicInstanceTopologyResponseBodyRedisShardList) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogicInstanceTopologyResponseBodyRedisShardList) GoString() string {
	return s.String()
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisShardList) SetNodeInfo(v []*DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo) *DescribeLogicInstanceTopologyResponseBodyRedisShardList {
	s.NodeInfo = v
	return s
}

type DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo struct {
	Capacity   *string `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	Connection *string `json:"Connection,omitempty" xml:"Connection,omitempty"`
	NodeType   *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	Bandwidth  *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	NodeId     *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo) GoString() string {
	return s.String()
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo) SetCapacity(v string) *DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo {
	s.Capacity = &v
	return s
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo) SetConnection(v string) *DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo {
	s.Connection = &v
	return s
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo) SetNodeType(v string) *DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo {
	s.NodeType = &v
	return s
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo) SetBandwidth(v string) *DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo {
	s.Bandwidth = &v
	return s
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo) SetNodeId(v string) *DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo {
	s.NodeId = &v
	return s
}

type DescribeLogicInstanceTopologyResponseBodyRedisProxyList struct {
	NodeInfo []*DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Repeated"`
}

func (s DescribeLogicInstanceTopologyResponseBodyRedisProxyList) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogicInstanceTopologyResponseBodyRedisProxyList) GoString() string {
	return s.String()
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisProxyList) SetNodeInfo(v []*DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo) *DescribeLogicInstanceTopologyResponseBodyRedisProxyList {
	s.NodeInfo = v
	return s
}

type DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo struct {
	Capacity   *string `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	Connection *string `json:"Connection,omitempty" xml:"Connection,omitempty"`
	NodeType   *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	Bandwidth  *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	NodeId     *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo) GoString() string {
	return s.String()
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo) SetCapacity(v string) *DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo {
	s.Capacity = &v
	return s
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo) SetConnection(v string) *DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo {
	s.Connection = &v
	return s
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo) SetNodeType(v string) *DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo {
	s.NodeType = &v
	return s
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo) SetBandwidth(v string) *DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo {
	s.Bandwidth = &v
	return s
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo) SetNodeId(v string) *DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo {
	s.NodeId = &v
	return s
}

type DescribeLogicInstanceTopologyResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLogicInstanceTopologyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLogicInstanceTopologyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogicInstanceTopologyResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogicInstanceTopologyResponse) SetHeaders(v map[string]*string) *DescribeLogicInstanceTopologyResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogicInstanceTopologyResponse) SetBody(v *DescribeLogicInstanceTopologyResponseBody) *DescribeLogicInstanceTopologyResponse {
	s.Body = v
	return s
}

type DescribeMonitorItemsRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s DescribeMonitorItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorItemsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitorItemsRequest) SetSecurityToken(v string) *DescribeMonitorItemsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeMonitorItemsRequest) SetOwnerId(v int64) *DescribeMonitorItemsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMonitorItemsRequest) SetResourceOwnerAccount(v string) *DescribeMonitorItemsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeMonitorItemsRequest) SetResourceOwnerId(v int64) *DescribeMonitorItemsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeMonitorItemsRequest) SetOwnerAccount(v string) *DescribeMonitorItemsRequest {
	s.OwnerAccount = &v
	return s
}

type DescribeMonitorItemsResponseBody struct {
	MonitorItems *DescribeMonitorItemsResponseBodyMonitorItems `json:"MonitorItems,omitempty" xml:"MonitorItems,omitempty" type:"Struct"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMonitorItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorItemsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitorItemsResponseBody) SetMonitorItems(v *DescribeMonitorItemsResponseBodyMonitorItems) *DescribeMonitorItemsResponseBody {
	s.MonitorItems = v
	return s
}

func (s *DescribeMonitorItemsResponseBody) SetRequestId(v string) *DescribeMonitorItemsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMonitorItemsResponseBodyMonitorItems struct {
	KVStoreMonitorItem []*DescribeMonitorItemsResponseBodyMonitorItemsKVStoreMonitorItem `json:"KVStoreMonitorItem,omitempty" xml:"KVStoreMonitorItem,omitempty" type:"Repeated"`
}

func (s DescribeMonitorItemsResponseBodyMonitorItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorItemsResponseBodyMonitorItems) GoString() string {
	return s.String()
}

func (s *DescribeMonitorItemsResponseBodyMonitorItems) SetKVStoreMonitorItem(v []*DescribeMonitorItemsResponseBodyMonitorItemsKVStoreMonitorItem) *DescribeMonitorItemsResponseBodyMonitorItems {
	s.KVStoreMonitorItem = v
	return s
}

type DescribeMonitorItemsResponseBodyMonitorItemsKVStoreMonitorItem struct {
	MonitorKey *string `json:"MonitorKey,omitempty" xml:"MonitorKey,omitempty"`
	Unit       *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeMonitorItemsResponseBodyMonitorItemsKVStoreMonitorItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorItemsResponseBodyMonitorItemsKVStoreMonitorItem) GoString() string {
	return s.String()
}

func (s *DescribeMonitorItemsResponseBodyMonitorItemsKVStoreMonitorItem) SetMonitorKey(v string) *DescribeMonitorItemsResponseBodyMonitorItemsKVStoreMonitorItem {
	s.MonitorKey = &v
	return s
}

func (s *DescribeMonitorItemsResponseBodyMonitorItemsKVStoreMonitorItem) SetUnit(v string) *DescribeMonitorItemsResponseBodyMonitorItemsKVStoreMonitorItem {
	s.Unit = &v
	return s
}

type DescribeMonitorItemsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMonitorItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMonitorItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorItemsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitorItemsResponse) SetHeaders(v map[string]*string) *DescribeMonitorItemsResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitorItemsResponse) SetBody(v *DescribeMonitorItemsResponseBody) *DescribeMonitorItemsResponse {
	s.Body = v
	return s
}

type DescribeParametersRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersRequest) GoString() string {
	return s.String()
}

func (s *DescribeParametersRequest) SetSecurityToken(v string) *DescribeParametersRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeParametersRequest) SetOwnerId(v int64) *DescribeParametersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeParametersRequest) SetResourceOwnerAccount(v string) *DescribeParametersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeParametersRequest) SetResourceOwnerId(v int64) *DescribeParametersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeParametersRequest) SetOwnerAccount(v string) *DescribeParametersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeParametersRequest) SetDBInstanceId(v string) *DescribeParametersRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeParametersRequest) SetNodeId(v string) *DescribeParametersRequest {
	s.NodeId = &v
	return s
}

type DescribeParametersResponseBody struct {
	RunningParameters *DescribeParametersResponseBodyRunningParameters `json:"RunningParameters,omitempty" xml:"RunningParameters,omitempty" type:"Struct"`
	EngineVersion     *string                                          `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	RequestId         *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ConfigParameters  *DescribeParametersResponseBodyConfigParameters  `json:"ConfigParameters,omitempty" xml:"ConfigParameters,omitempty" type:"Struct"`
	Engine            *string                                          `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s DescribeParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBody) SetRunningParameters(v *DescribeParametersResponseBodyRunningParameters) *DescribeParametersResponseBody {
	s.RunningParameters = v
	return s
}

func (s *DescribeParametersResponseBody) SetEngineVersion(v string) *DescribeParametersResponseBody {
	s.EngineVersion = &v
	return s
}

func (s *DescribeParametersResponseBody) SetRequestId(v string) *DescribeParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParametersResponseBody) SetConfigParameters(v *DescribeParametersResponseBodyConfigParameters) *DescribeParametersResponseBody {
	s.ConfigParameters = v
	return s
}

func (s *DescribeParametersResponseBody) SetEngine(v string) *DescribeParametersResponseBody {
	s.Engine = &v
	return s
}

type DescribeParametersResponseBodyRunningParameters struct {
	Parameter []*DescribeParametersResponseBodyRunningParametersParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Repeated"`
}

func (s DescribeParametersResponseBodyRunningParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponseBodyRunningParameters) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBodyRunningParameters) SetParameter(v []*DescribeParametersResponseBodyRunningParametersParameter) *DescribeParametersResponseBodyRunningParameters {
	s.Parameter = v
	return s
}

type DescribeParametersResponseBodyRunningParametersParameter struct {
	CheckingCode         *string `json:"CheckingCode,omitempty" xml:"CheckingCode,omitempty"`
	ParameterName        *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ParameterValue       *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
	ForceRestart         *string `json:"ForceRestart,omitempty" xml:"ForceRestart,omitempty"`
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	ModifiableStatus     *string `json:"ModifiableStatus,omitempty" xml:"ModifiableStatus,omitempty"`
}

func (s DescribeParametersResponseBodyRunningParametersParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponseBodyRunningParametersParameter) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBodyRunningParametersParameter) SetCheckingCode(v string) *DescribeParametersResponseBodyRunningParametersParameter {
	s.CheckingCode = &v
	return s
}

func (s *DescribeParametersResponseBodyRunningParametersParameter) SetParameterName(v string) *DescribeParametersResponseBodyRunningParametersParameter {
	s.ParameterName = &v
	return s
}

func (s *DescribeParametersResponseBodyRunningParametersParameter) SetParameterValue(v string) *DescribeParametersResponseBodyRunningParametersParameter {
	s.ParameterValue = &v
	return s
}

func (s *DescribeParametersResponseBodyRunningParametersParameter) SetForceRestart(v string) *DescribeParametersResponseBodyRunningParametersParameter {
	s.ForceRestart = &v
	return s
}

func (s *DescribeParametersResponseBodyRunningParametersParameter) SetParameterDescription(v string) *DescribeParametersResponseBodyRunningParametersParameter {
	s.ParameterDescription = &v
	return s
}

func (s *DescribeParametersResponseBodyRunningParametersParameter) SetModifiableStatus(v string) *DescribeParametersResponseBodyRunningParametersParameter {
	s.ModifiableStatus = &v
	return s
}

type DescribeParametersResponseBodyConfigParameters struct {
	Parameter []*DescribeParametersResponseBodyConfigParametersParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Repeated"`
}

func (s DescribeParametersResponseBodyConfigParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponseBodyConfigParameters) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBodyConfigParameters) SetParameter(v []*DescribeParametersResponseBodyConfigParametersParameter) *DescribeParametersResponseBodyConfigParameters {
	s.Parameter = v
	return s
}

type DescribeParametersResponseBodyConfigParametersParameter struct {
	CheckingCode         *string `json:"CheckingCode,omitempty" xml:"CheckingCode,omitempty"`
	ParameterName        *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ParameterValue       *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
	ForceRestart         *bool   `json:"ForceRestart,omitempty" xml:"ForceRestart,omitempty"`
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	ModifiableStatus     *bool   `json:"ModifiableStatus,omitempty" xml:"ModifiableStatus,omitempty"`
}

func (s DescribeParametersResponseBodyConfigParametersParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponseBodyConfigParametersParameter) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBodyConfigParametersParameter) SetCheckingCode(v string) *DescribeParametersResponseBodyConfigParametersParameter {
	s.CheckingCode = &v
	return s
}

func (s *DescribeParametersResponseBodyConfigParametersParameter) SetParameterName(v string) *DescribeParametersResponseBodyConfigParametersParameter {
	s.ParameterName = &v
	return s
}

func (s *DescribeParametersResponseBodyConfigParametersParameter) SetParameterValue(v string) *DescribeParametersResponseBodyConfigParametersParameter {
	s.ParameterValue = &v
	return s
}

func (s *DescribeParametersResponseBodyConfigParametersParameter) SetForceRestart(v bool) *DescribeParametersResponseBodyConfigParametersParameter {
	s.ForceRestart = &v
	return s
}

func (s *DescribeParametersResponseBodyConfigParametersParameter) SetParameterDescription(v string) *DescribeParametersResponseBodyConfigParametersParameter {
	s.ParameterDescription = &v
	return s
}

func (s *DescribeParametersResponseBodyConfigParametersParameter) SetModifiableStatus(v bool) *DescribeParametersResponseBodyConfigParametersParameter {
	s.ModifiableStatus = &v
	return s
}

type DescribeParametersResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeParametersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponse) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponse) SetHeaders(v map[string]*string) *DescribeParametersResponse {
	s.Headers = v
	return s
}

func (s *DescribeParametersResponse) SetBody(v *DescribeParametersResponseBody) *DescribeParametersResponse {
	s.Body = v
	return s
}

type DescribeParameterTemplatesRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	CharacterType        *string `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	Engine               *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion        *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeParameterTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeParameterTemplatesRequest) SetSecurityToken(v string) *DescribeParameterTemplatesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetOwnerId(v int64) *DescribeParameterTemplatesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetResourceOwnerAccount(v string) *DescribeParameterTemplatesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetResourceOwnerId(v int64) *DescribeParameterTemplatesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetOwnerAccount(v string) *DescribeParameterTemplatesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetCharacterType(v string) *DescribeParameterTemplatesRequest {
	s.CharacterType = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetEngine(v string) *DescribeParameterTemplatesRequest {
	s.Engine = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetEngineVersion(v string) *DescribeParameterTemplatesRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetResourceGroupId(v string) *DescribeParameterTemplatesRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeParameterTemplatesResponseBody struct {
	ParameterCount *string                                           `json:"ParameterCount,omitempty" xml:"ParameterCount,omitempty"`
	EngineVersion  *string                                           `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	Parameters     *DescribeParameterTemplatesResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Struct"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Engine         *string                                           `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s DescribeParameterTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParameterTemplatesResponseBody) SetParameterCount(v string) *DescribeParameterTemplatesResponseBody {
	s.ParameterCount = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBody) SetEngineVersion(v string) *DescribeParameterTemplatesResponseBody {
	s.EngineVersion = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBody) SetParameters(v *DescribeParameterTemplatesResponseBodyParameters) *DescribeParameterTemplatesResponseBody {
	s.Parameters = v
	return s
}

func (s *DescribeParameterTemplatesResponseBody) SetRequestId(v string) *DescribeParameterTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBody) SetEngine(v string) *DescribeParameterTemplatesResponseBody {
	s.Engine = &v
	return s
}

type DescribeParameterTemplatesResponseBodyParameters struct {
	TemplateRecord []*DescribeParameterTemplatesResponseBodyParametersTemplateRecord `json:"TemplateRecord,omitempty" xml:"TemplateRecord,omitempty" type:"Repeated"`
}

func (s DescribeParameterTemplatesResponseBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterTemplatesResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *DescribeParameterTemplatesResponseBodyParameters) SetTemplateRecord(v []*DescribeParameterTemplatesResponseBodyParametersTemplateRecord) *DescribeParameterTemplatesResponseBodyParameters {
	s.TemplateRecord = v
	return s
}

type DescribeParameterTemplatesResponseBodyParametersTemplateRecord struct {
	CheckingCode         *string `json:"CheckingCode,omitempty" xml:"CheckingCode,omitempty"`
	ParameterName        *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ParameterValue       *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
	ForceModify          *bool   `json:"ForceModify,omitempty" xml:"ForceModify,omitempty"`
	ForceRestart         *bool   `json:"ForceRestart,omitempty" xml:"ForceRestart,omitempty"`
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
}

func (s DescribeParameterTemplatesResponseBodyParametersTemplateRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterTemplatesResponseBodyParametersTemplateRecord) GoString() string {
	return s.String()
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) SetCheckingCode(v string) *DescribeParameterTemplatesResponseBodyParametersTemplateRecord {
	s.CheckingCode = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) SetParameterName(v string) *DescribeParameterTemplatesResponseBodyParametersTemplateRecord {
	s.ParameterName = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) SetParameterValue(v string) *DescribeParameterTemplatesResponseBodyParametersTemplateRecord {
	s.ParameterValue = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) SetForceModify(v bool) *DescribeParameterTemplatesResponseBodyParametersTemplateRecord {
	s.ForceModify = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) SetForceRestart(v bool) *DescribeParameterTemplatesResponseBodyParametersTemplateRecord {
	s.ForceRestart = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) SetParameterDescription(v string) *DescribeParameterTemplatesResponseBodyParametersTemplateRecord {
	s.ParameterDescription = &v
	return s
}

type DescribeParameterTemplatesResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeParameterTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeParameterTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeParameterTemplatesResponse) SetHeaders(v map[string]*string) *DescribeParameterTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeParameterTemplatesResponse) SetBody(v *DescribeParameterTemplatesResponseBody) *DescribeParameterTemplatesResponse {
	s.Body = v
	return s
}

type DescribePriceRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Capacity             *int64  `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	InstanceClass        *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	OrderType            *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ChargeType           *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	NodeType             *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	Period               *int64  `json:"Period,omitempty" xml:"Period,omitempty"`
	Quantity             *int64  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Instances            *string `json:"Instances,omitempty" xml:"Instances,omitempty"`
	BusinessInfo         *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	CouponNo             *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	ForceUpgrade         *bool   `json:"ForceUpgrade,omitempty" xml:"ForceUpgrade,omitempty"`
	OrderParamOut        *string `json:"OrderParamOut,omitempty" xml:"OrderParamOut,omitempty"`
}

func (s DescribePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceRequest) GoString() string {
	return s.String()
}

func (s *DescribePriceRequest) SetSecurityToken(v string) *DescribePriceRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribePriceRequest) SetOwnerId(v int64) *DescribePriceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePriceRequest) SetResourceOwnerAccount(v string) *DescribePriceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePriceRequest) SetResourceOwnerId(v int64) *DescribePriceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePriceRequest) SetOwnerAccount(v string) *DescribePriceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribePriceRequest) SetRegionId(v string) *DescribePriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePriceRequest) SetCapacity(v int64) *DescribePriceRequest {
	s.Capacity = &v
	return s
}

func (s *DescribePriceRequest) SetInstanceClass(v string) *DescribePriceRequest {
	s.InstanceClass = &v
	return s
}

func (s *DescribePriceRequest) SetOrderType(v string) *DescribePriceRequest {
	s.OrderType = &v
	return s
}

func (s *DescribePriceRequest) SetZoneId(v string) *DescribePriceRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribePriceRequest) SetChargeType(v string) *DescribePriceRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribePriceRequest) SetNodeType(v string) *DescribePriceRequest {
	s.NodeType = &v
	return s
}

func (s *DescribePriceRequest) SetPeriod(v int64) *DescribePriceRequest {
	s.Period = &v
	return s
}

func (s *DescribePriceRequest) SetQuantity(v int64) *DescribePriceRequest {
	s.Quantity = &v
	return s
}

func (s *DescribePriceRequest) SetInstanceId(v string) *DescribePriceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePriceRequest) SetInstances(v string) *DescribePriceRequest {
	s.Instances = &v
	return s
}

func (s *DescribePriceRequest) SetBusinessInfo(v string) *DescribePriceRequest {
	s.BusinessInfo = &v
	return s
}

func (s *DescribePriceRequest) SetCouponNo(v string) *DescribePriceRequest {
	s.CouponNo = &v
	return s
}

func (s *DescribePriceRequest) SetForceUpgrade(v bool) *DescribePriceRequest {
	s.ForceUpgrade = &v
	return s
}

func (s *DescribePriceRequest) SetOrderParamOut(v string) *DescribePriceRequest {
	s.OrderParamOut = &v
	return s
}

type DescribePriceResponseBody struct {
	Order       *DescribePriceResponseBodyOrder     `json:"Order,omitempty" xml:"Order,omitempty" type:"Struct"`
	RequestId   *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubOrders   *DescribePriceResponseBodySubOrders `json:"SubOrders,omitempty" xml:"SubOrders,omitempty" type:"Struct"`
	OrderParams *string                             `json:"OrderParams,omitempty" xml:"OrderParams,omitempty"`
	Rules       *DescribePriceResponseBodyRules     `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
}

func (s DescribePriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBody) SetOrder(v *DescribePriceResponseBodyOrder) *DescribePriceResponseBody {
	s.Order = v
	return s
}

func (s *DescribePriceResponseBody) SetRequestId(v string) *DescribePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePriceResponseBody) SetSubOrders(v *DescribePriceResponseBodySubOrders) *DescribePriceResponseBody {
	s.SubOrders = v
	return s
}

func (s *DescribePriceResponseBody) SetOrderParams(v string) *DescribePriceResponseBody {
	s.OrderParams = &v
	return s
}

func (s *DescribePriceResponseBody) SetRules(v *DescribePriceResponseBodyRules) *DescribePriceResponseBody {
	s.Rules = v
	return s
}

type DescribePriceResponseBodyOrder struct {
	Coupons        *DescribePriceResponseBodyOrderCoupons `json:"Coupons,omitempty" xml:"Coupons,omitempty" type:"Struct"`
	OriginalAmount *string                                `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	DiscountAmount *string                                `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	RuleIds        *DescribePriceResponseBodyOrderRuleIds `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Struct"`
	TradeAmount    *string                                `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
	Currency       *string                                `json:"Currency,omitempty" xml:"Currency,omitempty"`
}

func (s DescribePriceResponseBodyOrder) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBodyOrder) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyOrder) SetCoupons(v *DescribePriceResponseBodyOrderCoupons) *DescribePriceResponseBodyOrder {
	s.Coupons = v
	return s
}

func (s *DescribePriceResponseBodyOrder) SetOriginalAmount(v string) *DescribePriceResponseBodyOrder {
	s.OriginalAmount = &v
	return s
}

func (s *DescribePriceResponseBodyOrder) SetDiscountAmount(v string) *DescribePriceResponseBodyOrder {
	s.DiscountAmount = &v
	return s
}

func (s *DescribePriceResponseBodyOrder) SetRuleIds(v *DescribePriceResponseBodyOrderRuleIds) *DescribePriceResponseBodyOrder {
	s.RuleIds = v
	return s
}

func (s *DescribePriceResponseBodyOrder) SetTradeAmount(v string) *DescribePriceResponseBodyOrder {
	s.TradeAmount = &v
	return s
}

func (s *DescribePriceResponseBodyOrder) SetCurrency(v string) *DescribePriceResponseBodyOrder {
	s.Currency = &v
	return s
}

type DescribePriceResponseBodyOrderCoupons struct {
	Coupon []*DescribePriceResponseBodyOrderCouponsCoupon `json:"Coupon,omitempty" xml:"Coupon,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodyOrderCoupons) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBodyOrderCoupons) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyOrderCoupons) SetCoupon(v []*DescribePriceResponseBodyOrderCouponsCoupon) *DescribePriceResponseBodyOrderCoupons {
	s.Coupon = v
	return s
}

type DescribePriceResponseBodyOrderCouponsCoupon struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IsSelected  *string `json:"IsSelected,omitempty" xml:"IsSelected,omitempty"`
	CouponNo    *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribePriceResponseBodyOrderCouponsCoupon) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBodyOrderCouponsCoupon) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyOrderCouponsCoupon) SetDescription(v string) *DescribePriceResponseBodyOrderCouponsCoupon {
	s.Description = &v
	return s
}

func (s *DescribePriceResponseBodyOrderCouponsCoupon) SetIsSelected(v string) *DescribePriceResponseBodyOrderCouponsCoupon {
	s.IsSelected = &v
	return s
}

func (s *DescribePriceResponseBodyOrderCouponsCoupon) SetCouponNo(v string) *DescribePriceResponseBodyOrderCouponsCoupon {
	s.CouponNo = &v
	return s
}

func (s *DescribePriceResponseBodyOrderCouponsCoupon) SetName(v string) *DescribePriceResponseBodyOrderCouponsCoupon {
	s.Name = &v
	return s
}

type DescribePriceResponseBodyOrderRuleIds struct {
	RuleId []*string `json:"RuleId,omitempty" xml:"RuleId,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodyOrderRuleIds) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBodyOrderRuleIds) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyOrderRuleIds) SetRuleId(v []*string) *DescribePriceResponseBodyOrderRuleIds {
	s.RuleId = v
	return s
}

type DescribePriceResponseBodySubOrders struct {
	SubOrder []*DescribePriceResponseBodySubOrdersSubOrder `json:"SubOrder,omitempty" xml:"SubOrder,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodySubOrders) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBodySubOrders) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodySubOrders) SetSubOrder(v []*DescribePriceResponseBodySubOrdersSubOrder) *DescribePriceResponseBodySubOrders {
	s.SubOrder = v
	return s
}

type DescribePriceResponseBodySubOrdersSubOrder struct {
	OriginalAmount *string                                            `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	DiscountAmount *string                                            `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	RuleIds        *DescribePriceResponseBodySubOrdersSubOrderRuleIds `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Struct"`
	TradeAmount    *string                                            `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
	InstanceId     *string                                            `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribePriceResponseBodySubOrdersSubOrder) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBodySubOrdersSubOrder) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) SetOriginalAmount(v string) *DescribePriceResponseBodySubOrdersSubOrder {
	s.OriginalAmount = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) SetDiscountAmount(v string) *DescribePriceResponseBodySubOrdersSubOrder {
	s.DiscountAmount = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) SetRuleIds(v *DescribePriceResponseBodySubOrdersSubOrderRuleIds) *DescribePriceResponseBodySubOrdersSubOrder {
	s.RuleIds = v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) SetTradeAmount(v string) *DescribePriceResponseBodySubOrdersSubOrder {
	s.TradeAmount = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) SetInstanceId(v string) *DescribePriceResponseBodySubOrdersSubOrder {
	s.InstanceId = &v
	return s
}

type DescribePriceResponseBodySubOrdersSubOrderRuleIds struct {
	RuleId []*string `json:"RuleId,omitempty" xml:"RuleId,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodySubOrdersSubOrderRuleIds) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBodySubOrdersSubOrderRuleIds) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodySubOrdersSubOrderRuleIds) SetRuleId(v []*string) *DescribePriceResponseBodySubOrdersSubOrderRuleIds {
	s.RuleId = v
	return s
}

type DescribePriceResponseBodyRules struct {
	Rule []*DescribePriceResponseBodyRulesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodyRules) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyRules) SetRule(v []*DescribePriceResponseBodyRulesRule) *DescribePriceResponseBodyRules {
	s.Rule = v
	return s
}

type DescribePriceResponseBodyRulesRule struct {
	RuleDescId *int64  `json:"RuleDescId,omitempty" xml:"RuleDescId,omitempty"`
	Title      *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribePriceResponseBodyRulesRule) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBodyRulesRule) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyRulesRule) SetRuleDescId(v int64) *DescribePriceResponseBodyRulesRule {
	s.RuleDescId = &v
	return s
}

func (s *DescribePriceResponseBodyRulesRule) SetTitle(v string) *DescribePriceResponseBodyRulesRule {
	s.Title = &v
	return s
}

func (s *DescribePriceResponseBodyRulesRule) SetName(v string) *DescribePriceResponseBodyRulesRule {
	s.Name = &v
	return s
}

type DescribePriceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePriceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponse) GoString() string {
	return s.String()
}

func (s *DescribePriceResponse) SetHeaders(v map[string]*string) *DescribePriceResponse {
	s.Headers = v
	return s
}

func (s *DescribePriceResponse) SetBody(v *DescribePriceResponseBody) *DescribePriceResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	AcceptLanguage       *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetSecurityToken(v string) *DescribeRegionsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeRegionsRequest) SetOwnerId(v int64) *DescribeRegionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerAccount(v string) *DescribeRegionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerId(v int64) *DescribeRegionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRegionsRequest) SetOwnerAccount(v string) *DescribeRegionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

type DescribeRegionsResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RegionIds *DescribeRegionsResponseBodyRegionIds `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Struct"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegionIds(v *DescribeRegionsResponseBodyRegionIds) *DescribeRegionsResponseBody {
	s.RegionIds = v
	return s
}

type DescribeRegionsResponseBodyRegionIds struct {
	KVStoreRegion []*DescribeRegionsResponseBodyRegionIdsKVStoreRegion `json:"KVStoreRegion,omitempty" xml:"KVStoreRegion,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegionIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionIds) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionIds) SetKVStoreRegion(v []*DescribeRegionsResponseBodyRegionIdsKVStoreRegion) *DescribeRegionsResponseBodyRegionIds {
	s.KVStoreRegion = v
	return s
}

type DescribeRegionsResponseBodyRegionIdsKVStoreRegion struct {
	LocalName      *string                                                      `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	ZoneIdList     *DescribeRegionsResponseBodyRegionIdsKVStoreRegionZoneIdList `json:"ZoneIdList,omitempty" xml:"ZoneIdList,omitempty" type:"Struct"`
	RegionEndpoint *string                                                      `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string                                                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneIds        *string                                                      `json:"ZoneIds,omitempty" xml:"ZoneIds,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionIdsKVStoreRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionIdsKVStoreRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionIdsKVStoreRegion) SetLocalName(v string) *DescribeRegionsResponseBodyRegionIdsKVStoreRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionIdsKVStoreRegion) SetZoneIdList(v *DescribeRegionsResponseBodyRegionIdsKVStoreRegionZoneIdList) *DescribeRegionsResponseBodyRegionIdsKVStoreRegion {
	s.ZoneIdList = v
	return s
}

func (s *DescribeRegionsResponseBodyRegionIdsKVStoreRegion) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegionIdsKVStoreRegion {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionIdsKVStoreRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionIdsKVStoreRegion {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionIdsKVStoreRegion) SetZoneIds(v string) *DescribeRegionsResponseBodyRegionIdsKVStoreRegion {
	s.ZoneIds = &v
	return s
}

type DescribeRegionsResponseBodyRegionIdsKVStoreRegionZoneIdList struct {
	ZoneId []*string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegionIdsKVStoreRegionZoneIdList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionIdsKVStoreRegionZoneIdList) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionIdsKVStoreRegionZoneIdList) SetZoneId(v []*string) *DescribeRegionsResponseBodyRegionIdsKVStoreRegionZoneIdList {
	s.ZoneId = v
	return s
}

type DescribeRegionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeRoleZoneInfoRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Product              *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Category             *string `json:"Category,omitempty" xml:"Category,omitempty"`
	QueryType            *int32  `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Role                 *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s DescribeRoleZoneInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoleZoneInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeRoleZoneInfoRequest) SetSecurityToken(v string) *DescribeRoleZoneInfoRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetOwnerId(v int64) *DescribeRoleZoneInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetResourceOwnerAccount(v string) *DescribeRoleZoneInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetResourceOwnerId(v int64) *DescribeRoleZoneInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetOwnerAccount(v string) *DescribeRoleZoneInfoRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetInstanceId(v string) *DescribeRoleZoneInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetProduct(v string) *DescribeRoleZoneInfoRequest {
	s.Product = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetCategory(v string) *DescribeRoleZoneInfoRequest {
	s.Category = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetQueryType(v int32) *DescribeRoleZoneInfoRequest {
	s.QueryType = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetPageNumber(v int32) *DescribeRoleZoneInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetPageSize(v int32) *DescribeRoleZoneInfoRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetRole(v string) *DescribeRoleZoneInfoRequest {
	s.Role = &v
	return s
}

type DescribeRoleZoneInfoResponseBody struct {
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Node       *DescribeRoleZoneInfoResponseBodyNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Struct"`
}

func (s DescribeRoleZoneInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoleZoneInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRoleZoneInfoResponseBody) SetTotalCount(v int32) *DescribeRoleZoneInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBody) SetRequestId(v string) *DescribeRoleZoneInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBody) SetPageSize(v int32) *DescribeRoleZoneInfoResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBody) SetPageNumber(v int32) *DescribeRoleZoneInfoResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBody) SetNode(v *DescribeRoleZoneInfoResponseBodyNode) *DescribeRoleZoneInfoResponseBody {
	s.Node = v
	return s
}

type DescribeRoleZoneInfoResponseBodyNode struct {
	NodeInfo []*DescribeRoleZoneInfoResponseBodyNodeNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Repeated"`
}

func (s DescribeRoleZoneInfoResponseBodyNode) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoleZoneInfoResponseBodyNode) GoString() string {
	return s.String()
}

func (s *DescribeRoleZoneInfoResponseBodyNode) SetNodeInfo(v []*DescribeRoleZoneInfoResponseBodyNodeNodeInfo) *DescribeRoleZoneInfoResponseBodyNode {
	s.NodeInfo = v
	return s
}

type DescribeRoleZoneInfoResponseBodyNodeNodeInfo struct {
	DefaultBandWidth       *int64  `json:"DefaultBandWidth,omitempty" xml:"DefaultBandWidth,omitempty"`
	CurrentMinorVersion    *string `json:"CurrentMinorVersion,omitempty" xml:"CurrentMinorVersion,omitempty"`
	CurrentBandWidth       *int64  `json:"CurrentBandWidth,omitempty" xml:"CurrentBandWidth,omitempty"`
	InsType                *int32  `json:"InsType,omitempty" xml:"InsType,omitempty"`
	IsLatestVersion        *int32  `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
	NodeType               *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	InsName                *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	ZoneId                 *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	IsOpenBandWidthService *bool   `json:"IsOpenBandWidthService,omitempty" xml:"IsOpenBandWidthService,omitempty"`
	CustinsId              *string `json:"CustinsId,omitempty" xml:"CustinsId,omitempty"`
	Role                   *string `json:"Role,omitempty" xml:"Role,omitempty"`
	NodeId                 *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeRoleZoneInfoResponseBodyNodeNodeInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoleZoneInfoResponseBodyNodeNodeInfo) GoString() string {
	return s.String()
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetDefaultBandWidth(v int64) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.DefaultBandWidth = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetCurrentMinorVersion(v string) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.CurrentMinorVersion = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetCurrentBandWidth(v int64) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.CurrentBandWidth = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetInsType(v int32) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.InsType = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetIsLatestVersion(v int32) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.IsLatestVersion = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetNodeType(v string) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.NodeType = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetInsName(v string) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.InsName = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetZoneId(v string) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.ZoneId = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetIsOpenBandWidthService(v bool) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.IsOpenBandWidthService = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetCustinsId(v string) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.CustinsId = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetRole(v string) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.Role = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetNodeId(v string) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.NodeId = &v
	return s
}

type DescribeRoleZoneInfoResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRoleZoneInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRoleZoneInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoleZoneInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeRoleZoneInfoResponse) SetHeaders(v map[string]*string) *DescribeRoleZoneInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeRoleZoneInfoResponse) SetBody(v *DescribeRoleZoneInfoResponseBody) *DescribeRoleZoneInfoResponse {
	s.Body = v
	return s
}

type DescribeRunningLogRecordsRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	SQLId                *int64  `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	RoleType             *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	CharacterType        *string `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	QueryKeyword         *string `json:"QueryKeyword,omitempty" xml:"QueryKeyword,omitempty"`
	OrderType            *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
}

func (s DescribeRunningLogRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRunningLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRunningLogRecordsRequest) SetSecurityToken(v string) *DescribeRunningLogRecordsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetOwnerId(v int64) *DescribeRunningLogRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetResourceOwnerAccount(v string) *DescribeRunningLogRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetResourceOwnerId(v int64) *DescribeRunningLogRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetOwnerAccount(v string) *DescribeRunningLogRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetInstanceId(v string) *DescribeRunningLogRecordsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetNodeId(v string) *DescribeRunningLogRecordsRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetSQLId(v int64) *DescribeRunningLogRecordsRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetStartTime(v string) *DescribeRunningLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetEndTime(v string) *DescribeRunningLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetDBName(v string) *DescribeRunningLogRecordsRequest {
	s.DBName = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetRoleType(v string) *DescribeRunningLogRecordsRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetPageSize(v int32) *DescribeRunningLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetPageNumber(v int32) *DescribeRunningLogRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetResourceGroupId(v string) *DescribeRunningLogRecordsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetCharacterType(v string) *DescribeRunningLogRecordsRequest {
	s.CharacterType = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetQueryKeyword(v string) *DescribeRunningLogRecordsRequest {
	s.QueryKeyword = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetOrderType(v string) *DescribeRunningLogRecordsRequest {
	s.OrderType = &v
	return s
}

type DescribeRunningLogRecordsResponseBody struct {
	TotalRecordCount *int32                                      `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	PageRecordCount  *int32                                      `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize         *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	InstanceId       *string                                     `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber       *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	StartTime        *string                                     `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Items            *DescribeRunningLogRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	Engine           *string                                     `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s DescribeRunningLogRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRunningLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRunningLogRecordsResponseBody) SetTotalRecordCount(v int32) *DescribeRunningLogRecordsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBody) SetPageRecordCount(v int32) *DescribeRunningLogRecordsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBody) SetRequestId(v string) *DescribeRunningLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBody) SetPageSize(v int32) *DescribeRunningLogRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBody) SetInstanceId(v string) *DescribeRunningLogRecordsResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBody) SetPageNumber(v int32) *DescribeRunningLogRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBody) SetStartTime(v string) *DescribeRunningLogRecordsResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBody) SetItems(v *DescribeRunningLogRecordsResponseBodyItems) *DescribeRunningLogRecordsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeRunningLogRecordsResponseBody) SetEngine(v string) *DescribeRunningLogRecordsResponseBody {
	s.Engine = &v
	return s
}

type DescribeRunningLogRecordsResponseBodyItems struct {
	LogRecords []*DescribeRunningLogRecordsResponseBodyItemsLogRecords `json:"LogRecords,omitempty" xml:"LogRecords,omitempty" type:"Repeated"`
}

func (s DescribeRunningLogRecordsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeRunningLogRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeRunningLogRecordsResponseBodyItems) SetLogRecords(v []*DescribeRunningLogRecordsResponseBodyItemsLogRecords) *DescribeRunningLogRecordsResponseBodyItems {
	s.LogRecords = v
	return s
}

type DescribeRunningLogRecordsResponseBodyItemsLogRecords struct {
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s DescribeRunningLogRecordsResponseBodyItemsLogRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeRunningLogRecordsResponseBodyItemsLogRecords) GoString() string {
	return s.String()
}

func (s *DescribeRunningLogRecordsResponseBodyItemsLogRecords) SetCreateTime(v string) *DescribeRunningLogRecordsResponseBodyItemsLogRecords {
	s.CreateTime = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBodyItemsLogRecords) SetInstanceId(v string) *DescribeRunningLogRecordsResponseBodyItemsLogRecords {
	s.InstanceId = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBodyItemsLogRecords) SetContent(v string) *DescribeRunningLogRecordsResponseBodyItemsLogRecords {
	s.Content = &v
	return s
}

type DescribeRunningLogRecordsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRunningLogRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRunningLogRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRunningLogRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRunningLogRecordsResponse) SetHeaders(v map[string]*string) *DescribeRunningLogRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRunningLogRecordsResponse) SetBody(v *DescribeRunningLogRecordsResponseBody) *DescribeRunningLogRecordsResponse {
	s.Body = v
	return s
}

type DescribeSecurityGroupConfigurationRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeSecurityGroupConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupConfigurationRequest) SetSecurityToken(v string) *DescribeSecurityGroupConfigurationRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationRequest) SetOwnerId(v int64) *DescribeSecurityGroupConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationRequest) SetResourceOwnerAccount(v string) *DescribeSecurityGroupConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationRequest) SetResourceOwnerId(v int64) *DescribeSecurityGroupConfigurationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationRequest) SetOwnerAccount(v string) *DescribeSecurityGroupConfigurationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationRequest) SetInstanceId(v string) *DescribeSecurityGroupConfigurationRequest {
	s.InstanceId = &v
	return s
}

type DescribeSecurityGroupConfigurationResponseBody struct {
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Items     *DescribeSecurityGroupConfigurationResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeSecurityGroupConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupConfigurationResponseBody) SetRequestId(v string) *DescribeSecurityGroupConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationResponseBody) SetItems(v *DescribeSecurityGroupConfigurationResponseBodyItems) *DescribeSecurityGroupConfigurationResponseBody {
	s.Items = v
	return s
}

type DescribeSecurityGroupConfigurationResponseBodyItems struct {
	EcsSecurityGroupRelation []*DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation `json:"EcsSecurityGroupRelation,omitempty" xml:"EcsSecurityGroupRelation,omitempty" type:"Repeated"`
}

func (s DescribeSecurityGroupConfigurationResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupConfigurationResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItems) SetEcsSecurityGroupRelation(v []*DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) *DescribeSecurityGroupConfigurationResponseBodyItems {
	s.EcsSecurityGroupRelation = v
	return s
}

type DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation struct {
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	NetType         *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) SetSecurityGroupId(v string) *DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) SetNetType(v string) *DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation {
	s.NetType = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) SetRegionId(v string) *DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation {
	s.RegionId = &v
	return s
}

type DescribeSecurityGroupConfigurationResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSecurityGroupConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSecurityGroupConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupConfigurationResponse) SetHeaders(v map[string]*string) *DescribeSecurityGroupConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityGroupConfigurationResponse) SetBody(v *DescribeSecurityGroupConfigurationResponseBody) *DescribeSecurityGroupConfigurationResponse {
	s.Body = v
	return s
}

type DescribeSecurityIpsRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeSecurityIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIpsRequest) SetSecurityToken(v string) *DescribeSecurityIpsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeSecurityIpsRequest) SetOwnerId(v int64) *DescribeSecurityIpsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSecurityIpsRequest) SetResourceOwnerAccount(v string) *DescribeSecurityIpsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSecurityIpsRequest) SetResourceOwnerId(v int64) *DescribeSecurityIpsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSecurityIpsRequest) SetOwnerAccount(v string) *DescribeSecurityIpsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSecurityIpsRequest) SetInstanceId(v string) *DescribeSecurityIpsRequest {
	s.InstanceId = &v
	return s
}

type DescribeSecurityIpsResponseBody struct {
	RequestId        *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityIpGroups *DescribeSecurityIpsResponseBodySecurityIpGroups `json:"SecurityIpGroups,omitempty" xml:"SecurityIpGroups,omitempty" type:"Struct"`
}

func (s DescribeSecurityIpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityIpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIpsResponseBody) SetRequestId(v string) *DescribeSecurityIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityIpsResponseBody) SetSecurityIpGroups(v *DescribeSecurityIpsResponseBodySecurityIpGroups) *DescribeSecurityIpsResponseBody {
	s.SecurityIpGroups = v
	return s
}

type DescribeSecurityIpsResponseBodySecurityIpGroups struct {
	SecurityIpGroup []*DescribeSecurityIpsResponseBodySecurityIpGroupsSecurityIpGroup `json:"SecurityIpGroup,omitempty" xml:"SecurityIpGroup,omitempty" type:"Repeated"`
}

func (s DescribeSecurityIpsResponseBodySecurityIpGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityIpsResponseBodySecurityIpGroups) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIpsResponseBodySecurityIpGroups) SetSecurityIpGroup(v []*DescribeSecurityIpsResponseBodySecurityIpGroupsSecurityIpGroup) *DescribeSecurityIpsResponseBodySecurityIpGroups {
	s.SecurityIpGroup = v
	return s
}

type DescribeSecurityIpsResponseBodySecurityIpGroupsSecurityIpGroup struct {
	SecurityIpGroupName      *string `json:"SecurityIpGroupName,omitempty" xml:"SecurityIpGroupName,omitempty"`
	SecurityIpGroupAttribute *string `json:"SecurityIpGroupAttribute,omitempty" xml:"SecurityIpGroupAttribute,omitempty"`
	SecurityIpList           *string `json:"SecurityIpList,omitempty" xml:"SecurityIpList,omitempty"`
}

func (s DescribeSecurityIpsResponseBodySecurityIpGroupsSecurityIpGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityIpsResponseBodySecurityIpGroupsSecurityIpGroup) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIpsResponseBodySecurityIpGroupsSecurityIpGroup) SetSecurityIpGroupName(v string) *DescribeSecurityIpsResponseBodySecurityIpGroupsSecurityIpGroup {
	s.SecurityIpGroupName = &v
	return s
}

func (s *DescribeSecurityIpsResponseBodySecurityIpGroupsSecurityIpGroup) SetSecurityIpGroupAttribute(v string) *DescribeSecurityIpsResponseBodySecurityIpGroupsSecurityIpGroup {
	s.SecurityIpGroupAttribute = &v
	return s
}

func (s *DescribeSecurityIpsResponseBodySecurityIpGroupsSecurityIpGroup) SetSecurityIpList(v string) *DescribeSecurityIpsResponseBodySecurityIpGroupsSecurityIpGroup {
	s.SecurityIpList = &v
	return s
}

type DescribeSecurityIpsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSecurityIpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSecurityIpsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityIpsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIpsResponse) SetHeaders(v map[string]*string) *DescribeSecurityIpsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityIpsResponse) SetBody(v *DescribeSecurityIpsResponseBody) *DescribeSecurityIpsResponse {
	s.Body = v
	return s
}

type DescribeSlowLogRecordsRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	SQLId                *int64  `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	SlowLogRecordType    *string `json:"SlowLogRecordType,omitempty" xml:"SlowLogRecordType,omitempty"`
	QueryKeyword         *string `json:"QueryKeyword,omitempty" xml:"QueryKeyword,omitempty"`
	OrderType            *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OrderBy              *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
}

func (s DescribeSlowLogRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsRequest) SetSecurityToken(v string) *DescribeSlowLogRecordsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetOwnerId(v int64) *DescribeSlowLogRecordsRequest {
	s.OwnerId = &v
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

func (s *DescribeSlowLogRecordsRequest) SetOwnerAccount(v string) *DescribeSlowLogRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetInstanceId(v string) *DescribeSlowLogRecordsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetNodeId(v string) *DescribeSlowLogRecordsRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetSQLId(v int64) *DescribeSlowLogRecordsRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetStartTime(v string) *DescribeSlowLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetEndTime(v string) *DescribeSlowLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetDBName(v string) *DescribeSlowLogRecordsRequest {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetPageSize(v int32) *DescribeSlowLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetPageNumber(v int32) *DescribeSlowLogRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetSlowLogRecordType(v string) *DescribeSlowLogRecordsRequest {
	s.SlowLogRecordType = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetQueryKeyword(v string) *DescribeSlowLogRecordsRequest {
	s.QueryKeyword = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetOrderType(v string) *DescribeSlowLogRecordsRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetOrderBy(v string) *DescribeSlowLogRecordsRequest {
	s.OrderBy = &v
	return s
}

type DescribeSlowLogRecordsResponseBody struct {
	TotalRecordCount *int32                                   `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	PageRecordCount  *int32                                   `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize         *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	InstanceId       *string                                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber       *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	StartTime        *string                                  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Items            *DescribeSlowLogRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	Engine           *string                                  `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBody) SetTotalRecordCount(v int32) *DescribeSlowLogRecordsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetPageRecordCount(v int32) *DescribeSlowLogRecordsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetRequestId(v string) *DescribeSlowLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetPageSize(v int32) *DescribeSlowLogRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetInstanceId(v string) *DescribeSlowLogRecordsResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetPageNumber(v int32) *DescribeSlowLogRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetStartTime(v string) *DescribeSlowLogRecordsResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetItems(v *DescribeSlowLogRecordsResponseBodyItems) *DescribeSlowLogRecordsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetEngine(v string) *DescribeSlowLogRecordsResponseBody {
	s.Engine = &v
	return s
}

type DescribeSlowLogRecordsResponseBodyItems struct {
	LogRecords []*DescribeSlowLogRecordsResponseBodyItemsLogRecords `json:"LogRecords,omitempty" xml:"LogRecords,omitempty" type:"Repeated"`
}

func (s DescribeSlowLogRecordsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetLogRecords(v []*DescribeSlowLogRecordsResponseBodyItemsLogRecords) *DescribeSlowLogRecordsResponseBodyItems {
	s.LogRecords = v
	return s
}

type DescribeSlowLogRecordsResponseBodyItemsLogRecords struct {
	Account      *string `json:"Account,omitempty" xml:"Account,omitempty"`
	ElapsedTime  *int64  `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	Command      *string `json:"Command,omitempty" xml:"Command,omitempty"`
	DBName       *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	ExecuteTime  *string `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	DataBaseName *string `json:"DataBaseName,omitempty" xml:"DataBaseName,omitempty"`
	NodeId       *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	AccountName  *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	IPAddress    *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBodyItemsLogRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyItemsLogRecords) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetAccount(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.Account = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetElapsedTime(v int64) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.ElapsedTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetCommand(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.Command = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetDBName(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetExecuteTime(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetDataBaseName(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.DataBaseName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetNodeId(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.NodeId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetAccountName(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.AccountName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetIPAddress(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.IPAddress = &v
	return s
}

type DescribeSlowLogRecordsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSlowLogRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSlowLogRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponse) SetHeaders(v map[string]*string) *DescribeSlowLogRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlowLogRecordsResponse) SetBody(v *DescribeSlowLogRecordsResponseBody) *DescribeSlowLogRecordsResponse {
	s.Body = v
	return s
}

type DescribeTasksRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeTasksRequest) SetSecurityToken(v string) *DescribeTasksRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeTasksRequest) SetOwnerId(v int64) *DescribeTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTasksRequest) SetResourceOwnerAccount(v string) *DescribeTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTasksRequest) SetResourceOwnerId(v int64) *DescribeTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTasksRequest) SetOwnerAccount(v string) *DescribeTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTasksRequest) SetInstanceId(v string) *DescribeTasksRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTasksRequest) SetPageSize(v int32) *DescribeTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTasksRequest) SetPageNumber(v int32) *DescribeTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTasksRequest) SetStartTime(v string) *DescribeTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTasksRequest) SetEndTime(v string) *DescribeTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTasksRequest) SetStatus(v string) *DescribeTasksRequest {
	s.Status = &v
	return s
}

type DescribeTasksResponseBody struct {
	TotalRecordCount *int32                            `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	RequestId        *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize         *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber       *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Items            []*DescribeTasksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBody) SetTotalRecordCount(v int32) *DescribeTasksResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeTasksResponseBody) SetRequestId(v string) *DescribeTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTasksResponseBody) SetPageSize(v int32) *DescribeTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTasksResponseBody) SetPageNumber(v int32) *DescribeTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTasksResponseBody) SetItems(v []*DescribeTasksResponseBodyItems) *DescribeTasksResponseBody {
	s.Items = v
	return s
}

type DescribeTasksResponseBodyItems struct {
	Status           *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	FinishTime       *string  `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	StepsInfo        *string  `json:"StepsInfo,omitempty" xml:"StepsInfo,omitempty"`
	Progress         *float32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	BeginTime        *string  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	Remain           *int32   `json:"Remain,omitempty" xml:"Remain,omitempty"`
	CurrentStepName  *string  `json:"CurrentStepName,omitempty" xml:"CurrentStepName,omitempty"`
	StepProgressInfo *string  `json:"StepProgressInfo,omitempty" xml:"StepProgressInfo,omitempty"`
	TaskId           *string  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskAction       *string  `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
}

func (s DescribeTasksResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBodyItems) SetStatus(v string) *DescribeTasksResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetFinishTime(v string) *DescribeTasksResponseBodyItems {
	s.FinishTime = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetStepsInfo(v string) *DescribeTasksResponseBodyItems {
	s.StepsInfo = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetProgress(v float32) *DescribeTasksResponseBodyItems {
	s.Progress = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetBeginTime(v string) *DescribeTasksResponseBodyItems {
	s.BeginTime = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetRemain(v int32) *DescribeTasksResponseBodyItems {
	s.Remain = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetCurrentStepName(v string) *DescribeTasksResponseBodyItems {
	s.CurrentStepName = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetStepProgressInfo(v string) *DescribeTasksResponseBodyItems {
	s.StepProgressInfo = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetTaskId(v string) *DescribeTasksResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetTaskAction(v string) *DescribeTasksResponseBodyItems {
	s.TaskAction = &v
	return s
}

type DescribeTasksResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponse) SetHeaders(v map[string]*string) *DescribeTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeTasksResponse) SetBody(v *DescribeTasksResponseBody) *DescribeTasksResponse {
	s.Body = v
	return s
}

type DescribeUserClusterHostRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MaxRecordsPerPage    *int32  `json:"MaxRecordsPerPage,omitempty" xml:"MaxRecordsPerPage,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Engine               *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeUserClusterHostRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserClusterHostRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserClusterHostRequest) SetSecurityToken(v string) *DescribeUserClusterHostRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeUserClusterHostRequest) SetOwnerId(v int64) *DescribeUserClusterHostRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserClusterHostRequest) SetResourceOwnerAccount(v string) *DescribeUserClusterHostRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeUserClusterHostRequest) SetResourceOwnerId(v int64) *DescribeUserClusterHostRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeUserClusterHostRequest) SetOwnerAccount(v string) *DescribeUserClusterHostRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeUserClusterHostRequest) SetRegionId(v string) *DescribeUserClusterHostRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserClusterHostRequest) SetMaxRecordsPerPage(v int32) *DescribeUserClusterHostRequest {
	s.MaxRecordsPerPage = &v
	return s
}

func (s *DescribeUserClusterHostRequest) SetPageNumber(v int32) *DescribeUserClusterHostRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserClusterHostRequest) SetZoneId(v string) *DescribeUserClusterHostRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeUserClusterHostRequest) SetEngine(v string) *DescribeUserClusterHostRequest {
	s.Engine = &v
	return s
}

func (s *DescribeUserClusterHostRequest) SetClusterId(v string) *DescribeUserClusterHostRequest {
	s.ClusterId = &v
	return s
}

type DescribeUserClusterHostResponseBody struct {
	MaxRecordsPerPage *string                                         `json:"MaxRecordsPerPage,omitempty" xml:"MaxRecordsPerPage,omitempty"`
	HostItems         []*DescribeUserClusterHostResponseBodyHostItems `json:"HostItems,omitempty" xml:"HostItems,omitempty" type:"Repeated"`
	RequestId         *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber        *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalRecords      *int32                                          `json:"TotalRecords,omitempty" xml:"TotalRecords,omitempty"`
	ItemNumbers       *int32                                          `json:"ItemNumbers,omitempty" xml:"ItemNumbers,omitempty"`
}

func (s DescribeUserClusterHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserClusterHostResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserClusterHostResponseBody) SetMaxRecordsPerPage(v string) *DescribeUserClusterHostResponseBody {
	s.MaxRecordsPerPage = &v
	return s
}

func (s *DescribeUserClusterHostResponseBody) SetHostItems(v []*DescribeUserClusterHostResponseBodyHostItems) *DescribeUserClusterHostResponseBody {
	s.HostItems = v
	return s
}

func (s *DescribeUserClusterHostResponseBody) SetRequestId(v string) *DescribeUserClusterHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserClusterHostResponseBody) SetPageNumber(v int32) *DescribeUserClusterHostResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserClusterHostResponseBody) SetTotalRecords(v int32) *DescribeUserClusterHostResponseBody {
	s.TotalRecords = &v
	return s
}

func (s *DescribeUserClusterHostResponseBody) SetItemNumbers(v int32) *DescribeUserClusterHostResponseBody {
	s.ItemNumbers = &v
	return s
}

type DescribeUserClusterHostResponseBodyHostItems struct {
	HostIP           *string `json:"HostIP,omitempty" xml:"HostIP,omitempty"`
	ExpireTime       *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	HostStatus       *string `json:"HostStatus,omitempty" xml:"HostStatus,omitempty"`
	ChargeType       *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	HostName         *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	HostStorage      *string `json:"HostStorage,omitempty" xml:"HostStorage,omitempty"`
	InstanceNumber   *string `json:"InstanceNumber,omitempty" xml:"InstanceNumber,omitempty"`
	HostId           *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	HostClass        *string `json:"HostClass,omitempty" xml:"HostClass,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AllocationStatus *string `json:"AllocationStatus,omitempty" xml:"AllocationStatus,omitempty"`
	ZoneId           *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	HostCpu          *string `json:"HostCpu,omitempty" xml:"HostCpu,omitempty"`
	Engine           *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	HostMem          *string `json:"HostMem,omitempty" xml:"HostMem,omitempty"`
	Id               *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	ClusterId        *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeUserClusterHostResponseBodyHostItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserClusterHostResponseBodyHostItems) GoString() string {
	return s.String()
}

func (s *DescribeUserClusterHostResponseBodyHostItems) SetHostIP(v string) *DescribeUserClusterHostResponseBodyHostItems {
	s.HostIP = &v
	return s
}

func (s *DescribeUserClusterHostResponseBodyHostItems) SetExpireTime(v string) *DescribeUserClusterHostResponseBodyHostItems {
	s.ExpireTime = &v
	return s
}

func (s *DescribeUserClusterHostResponseBodyHostItems) SetCreateTime(v string) *DescribeUserClusterHostResponseBodyHostItems {
	s.CreateTime = &v
	return s
}

func (s *DescribeUserClusterHostResponseBodyHostItems) SetHostStatus(v string) *DescribeUserClusterHostResponseBodyHostItems {
	s.HostStatus = &v
	return s
}

func (s *DescribeUserClusterHostResponseBodyHostItems) SetChargeType(v string) *DescribeUserClusterHostResponseBodyHostItems {
	s.ChargeType = &v
	return s
}

func (s *DescribeUserClusterHostResponseBodyHostItems) SetHostName(v string) *DescribeUserClusterHostResponseBodyHostItems {
	s.HostName = &v
	return s
}

func (s *DescribeUserClusterHostResponseBodyHostItems) SetHostStorage(v string) *DescribeUserClusterHostResponseBodyHostItems {
	s.HostStorage = &v
	return s
}

func (s *DescribeUserClusterHostResponseBodyHostItems) SetInstanceNumber(v string) *DescribeUserClusterHostResponseBodyHostItems {
	s.InstanceNumber = &v
	return s
}

func (s *DescribeUserClusterHostResponseBodyHostItems) SetHostId(v string) *DescribeUserClusterHostResponseBodyHostItems {
	s.HostId = &v
	return s
}

func (s *DescribeUserClusterHostResponseBodyHostItems) SetHostClass(v string) *DescribeUserClusterHostResponseBodyHostItems {
	s.HostClass = &v
	return s
}

func (s *DescribeUserClusterHostResponseBodyHostItems) SetRegionId(v string) *DescribeUserClusterHostResponseBodyHostItems {
	s.RegionId = &v
	return s
}

func (s *DescribeUserClusterHostResponseBodyHostItems) SetAllocationStatus(v string) *DescribeUserClusterHostResponseBodyHostItems {
	s.AllocationStatus = &v
	return s
}

func (s *DescribeUserClusterHostResponseBodyHostItems) SetZoneId(v string) *DescribeUserClusterHostResponseBodyHostItems {
	s.ZoneId = &v
	return s
}

func (s *DescribeUserClusterHostResponseBodyHostItems) SetHostCpu(v string) *DescribeUserClusterHostResponseBodyHostItems {
	s.HostCpu = &v
	return s
}

func (s *DescribeUserClusterHostResponseBodyHostItems) SetEngine(v string) *DescribeUserClusterHostResponseBodyHostItems {
	s.Engine = &v
	return s
}

func (s *DescribeUserClusterHostResponseBodyHostItems) SetHostMem(v string) *DescribeUserClusterHostResponseBodyHostItems {
	s.HostMem = &v
	return s
}

func (s *DescribeUserClusterHostResponseBodyHostItems) SetId(v int32) *DescribeUserClusterHostResponseBodyHostItems {
	s.Id = &v
	return s
}

func (s *DescribeUserClusterHostResponseBodyHostItems) SetClusterId(v string) *DescribeUserClusterHostResponseBodyHostItems {
	s.ClusterId = &v
	return s
}

type DescribeUserClusterHostResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserClusterHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserClusterHostResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserClusterHostResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserClusterHostResponse) SetHeaders(v map[string]*string) *DescribeUserClusterHostResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserClusterHostResponse) SetBody(v *DescribeUserClusterHostResponseBody) *DescribeUserClusterHostResponse {
	s.Body = v
	return s
}

type DescribeUserClusterHostInstanceRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MaxRecordsPerPage    *int32  `json:"MaxRecordsPerPage,omitempty" xml:"MaxRecordsPerPage,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Engine               *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	InstanceIds          *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	InstanceStatus       *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
}

func (s DescribeUserClusterHostInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserClusterHostInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserClusterHostInstanceRequest) SetSecurityToken(v string) *DescribeUserClusterHostInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeUserClusterHostInstanceRequest) SetOwnerId(v int64) *DescribeUserClusterHostInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserClusterHostInstanceRequest) SetResourceOwnerAccount(v string) *DescribeUserClusterHostInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeUserClusterHostInstanceRequest) SetResourceOwnerId(v int64) *DescribeUserClusterHostInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeUserClusterHostInstanceRequest) SetOwnerAccount(v string) *DescribeUserClusterHostInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeUserClusterHostInstanceRequest) SetRegionId(v string) *DescribeUserClusterHostInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserClusterHostInstanceRequest) SetMaxRecordsPerPage(v int32) *DescribeUserClusterHostInstanceRequest {
	s.MaxRecordsPerPage = &v
	return s
}

func (s *DescribeUserClusterHostInstanceRequest) SetPageNumber(v int32) *DescribeUserClusterHostInstanceRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserClusterHostInstanceRequest) SetZoneId(v string) *DescribeUserClusterHostInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeUserClusterHostInstanceRequest) SetEngine(v string) *DescribeUserClusterHostInstanceRequest {
	s.Engine = &v
	return s
}

func (s *DescribeUserClusterHostInstanceRequest) SetClusterId(v string) *DescribeUserClusterHostInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeUserClusterHostInstanceRequest) SetInstanceIds(v string) *DescribeUserClusterHostInstanceRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeUserClusterHostInstanceRequest) SetInstanceStatus(v string) *DescribeUserClusterHostInstanceRequest {
	s.InstanceStatus = &v
	return s
}

type DescribeUserClusterHostInstanceResponseBody struct {
	MaxRecordsPerPage *int32                                                     `json:"MaxRecordsPerPage,omitempty" xml:"MaxRecordsPerPage,omitempty"`
	RequestId         *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber        *int32                                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	InstancesItems    *DescribeUserClusterHostInstanceResponseBodyInstancesItems `json:"InstancesItems,omitempty" xml:"InstancesItems,omitempty" type:"Struct"`
	TotalRecords      *int32                                                     `json:"TotalRecords,omitempty" xml:"TotalRecords,omitempty"`
	ItemNumbers       *int32                                                     `json:"ItemNumbers,omitempty" xml:"ItemNumbers,omitempty"`
}

func (s DescribeUserClusterHostInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserClusterHostInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserClusterHostInstanceResponseBody) SetMaxRecordsPerPage(v int32) *DescribeUserClusterHostInstanceResponseBody {
	s.MaxRecordsPerPage = &v
	return s
}

func (s *DescribeUserClusterHostInstanceResponseBody) SetRequestId(v string) *DescribeUserClusterHostInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserClusterHostInstanceResponseBody) SetPageNumber(v int32) *DescribeUserClusterHostInstanceResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserClusterHostInstanceResponseBody) SetInstancesItems(v *DescribeUserClusterHostInstanceResponseBodyInstancesItems) *DescribeUserClusterHostInstanceResponseBody {
	s.InstancesItems = v
	return s
}

func (s *DescribeUserClusterHostInstanceResponseBody) SetTotalRecords(v int32) *DescribeUserClusterHostInstanceResponseBody {
	s.TotalRecords = &v
	return s
}

func (s *DescribeUserClusterHostInstanceResponseBody) SetItemNumbers(v int32) *DescribeUserClusterHostInstanceResponseBody {
	s.ItemNumbers = &v
	return s
}

type DescribeUserClusterHostInstanceResponseBodyInstancesItems struct {
	InstanceInfo []*DescribeUserClusterHostInstanceResponseBodyInstancesItemsInstanceInfo `json:"InstanceInfo,omitempty" xml:"InstanceInfo,omitempty" type:"Repeated"`
}

func (s DescribeUserClusterHostInstanceResponseBodyInstancesItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserClusterHostInstanceResponseBodyInstancesItems) GoString() string {
	return s.String()
}

func (s *DescribeUserClusterHostInstanceResponseBodyInstancesItems) SetInstanceInfo(v []*DescribeUserClusterHostInstanceResponseBodyInstancesItemsInstanceInfo) *DescribeUserClusterHostInstanceResponseBodyInstancesItems {
	s.InstanceInfo = v
	return s
}

type DescribeUserClusterHostInstanceResponseBodyInstancesItemsInstanceInfo struct {
	InstanceClass  *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ZoneId         *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	Engine         *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType   *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	EngineVersion  *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId      *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeUserClusterHostInstanceResponseBodyInstancesItemsInstanceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserClusterHostInstanceResponseBodyInstancesItemsInstanceInfo) GoString() string {
	return s.String()
}

func (s *DescribeUserClusterHostInstanceResponseBodyInstancesItemsInstanceInfo) SetInstanceClass(v string) *DescribeUserClusterHostInstanceResponseBodyInstancesItemsInstanceInfo {
	s.InstanceClass = &v
	return s
}

func (s *DescribeUserClusterHostInstanceResponseBodyInstancesItemsInstanceInfo) SetCreateTime(v string) *DescribeUserClusterHostInstanceResponseBodyInstancesItemsInstanceInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeUserClusterHostInstanceResponseBodyInstancesItemsInstanceInfo) SetZoneId(v string) *DescribeUserClusterHostInstanceResponseBodyInstancesItemsInstanceInfo {
	s.ZoneId = &v
	return s
}

func (s *DescribeUserClusterHostInstanceResponseBodyInstancesItemsInstanceInfo) SetInstanceStatus(v string) *DescribeUserClusterHostInstanceResponseBodyInstancesItemsInstanceInfo {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeUserClusterHostInstanceResponseBodyInstancesItemsInstanceInfo) SetEngine(v string) *DescribeUserClusterHostInstanceResponseBodyInstancesItemsInstanceInfo {
	s.Engine = &v
	return s
}

func (s *DescribeUserClusterHostInstanceResponseBodyInstancesItemsInstanceInfo) SetInstanceId(v string) *DescribeUserClusterHostInstanceResponseBodyInstancesItemsInstanceInfo {
	s.InstanceId = &v
	return s
}

func (s *DescribeUserClusterHostInstanceResponseBodyInstancesItemsInstanceInfo) SetInstanceType(v string) *DescribeUserClusterHostInstanceResponseBodyInstancesItemsInstanceInfo {
	s.InstanceType = &v
	return s
}

func (s *DescribeUserClusterHostInstanceResponseBodyInstancesItemsInstanceInfo) SetEngineVersion(v string) *DescribeUserClusterHostInstanceResponseBodyInstancesItemsInstanceInfo {
	s.EngineVersion = &v
	return s
}

func (s *DescribeUserClusterHostInstanceResponseBodyInstancesItemsInstanceInfo) SetRegionId(v string) *DescribeUserClusterHostInstanceResponseBodyInstancesItemsInstanceInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeUserClusterHostInstanceResponseBodyInstancesItemsInstanceInfo) SetClusterId(v string) *DescribeUserClusterHostInstanceResponseBodyInstancesItemsInstanceInfo {
	s.ClusterId = &v
	return s
}

type DescribeUserClusterHostInstanceResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserClusterHostInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserClusterHostInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserClusterHostInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserClusterHostInstanceResponse) SetHeaders(v map[string]*string) *DescribeUserClusterHostInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserClusterHostInstanceResponse) SetBody(v *DescribeUserClusterHostInstanceResponseBody) *DescribeUserClusterHostInstanceResponse {
	s.Body = v
	return s
}

type DescribeZonesRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AcceptLanguage       *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) SetSecurityToken(v string) *DescribeZonesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeZonesRequest) SetOwnerId(v int64) *DescribeZonesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeZonesRequest) SetResourceOwnerAccount(v string) *DescribeZonesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeZonesRequest) SetResourceOwnerId(v int64) *DescribeZonesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeZonesRequest) SetOwnerAccount(v string) *DescribeZonesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeZonesRequest) SetRegionId(v string) *DescribeZonesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeZonesRequest) SetAcceptLanguage(v string) *DescribeZonesRequest {
	s.AcceptLanguage = &v
	return s
}

type DescribeZonesResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Zones     *DescribeZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBody) SetRequestId(v string) *DescribeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponseBody) SetZones(v *DescribeZonesResponseBodyZones) *DescribeZonesResponseBody {
	s.Zones = v
	return s
}

type DescribeZonesResponseBodyZones struct {
	KVStoreZone []*DescribeZonesResponseBodyZonesKVStoreZone `json:"KVStoreZone,omitempty" xml:"KVStoreZone,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZones) SetKVStoreZone(v []*DescribeZonesResponseBodyZonesKVStoreZone) *DescribeZonesResponseBodyZones {
	s.KVStoreZone = v
	return s
}

type DescribeZonesResponseBodyZonesKVStoreZone struct {
	IsRds         *bool   `json:"IsRds,omitempty" xml:"IsRds,omitempty"`
	ZoneId        *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Disabled      *bool   `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	SwitchNetwork *bool   `json:"SwitchNetwork,omitempty" xml:"SwitchNetwork,omitempty"`
	ZoneName      *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeZonesResponseBodyZonesKVStoreZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesKVStoreZone) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesKVStoreZone) SetIsRds(v bool) *DescribeZonesResponseBodyZonesKVStoreZone {
	s.IsRds = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesKVStoreZone) SetZoneId(v string) *DescribeZonesResponseBodyZonesKVStoreZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesKVStoreZone) SetDisabled(v bool) *DescribeZonesResponseBodyZonesKVStoreZone {
	s.Disabled = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesKVStoreZone) SetSwitchNetwork(v bool) *DescribeZonesResponseBodyZonesKVStoreZone {
	s.SwitchNetwork = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesKVStoreZone) SetZoneName(v string) *DescribeZonesResponseBodyZonesKVStoreZone {
	s.ZoneName = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesKVStoreZone) SetRegionId(v string) *DescribeZonesResponseBodyZonesKVStoreZone {
	s.RegionId = &v
	return s
}

type DescribeZonesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponse) SetHeaders(v map[string]*string) *DescribeZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeZonesResponse) SetBody(v *DescribeZonesResponseBody) *DescribeZonesResponse {
	s.Body = v
	return s
}

type EnableAdditionalBandwidthRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CouponNo             *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	AutoPay              *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Bandwidth            *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	OrderTimeLength      *string `json:"OrderTimeLength,omitempty" xml:"OrderTimeLength,omitempty"`
}

func (s EnableAdditionalBandwidthRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableAdditionalBandwidthRequest) GoString() string {
	return s.String()
}

func (s *EnableAdditionalBandwidthRequest) SetSecurityToken(v string) *EnableAdditionalBandwidthRequest {
	s.SecurityToken = &v
	return s
}

func (s *EnableAdditionalBandwidthRequest) SetOwnerId(v int64) *EnableAdditionalBandwidthRequest {
	s.OwnerId = &v
	return s
}

func (s *EnableAdditionalBandwidthRequest) SetResourceOwnerAccount(v string) *EnableAdditionalBandwidthRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *EnableAdditionalBandwidthRequest) SetResourceOwnerId(v int64) *EnableAdditionalBandwidthRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *EnableAdditionalBandwidthRequest) SetOwnerAccount(v string) *EnableAdditionalBandwidthRequest {
	s.OwnerAccount = &v
	return s
}

func (s *EnableAdditionalBandwidthRequest) SetInstanceId(v string) *EnableAdditionalBandwidthRequest {
	s.InstanceId = &v
	return s
}

func (s *EnableAdditionalBandwidthRequest) SetCouponNo(v string) *EnableAdditionalBandwidthRequest {
	s.CouponNo = &v
	return s
}

func (s *EnableAdditionalBandwidthRequest) SetAutoPay(v bool) *EnableAdditionalBandwidthRequest {
	s.AutoPay = &v
	return s
}

func (s *EnableAdditionalBandwidthRequest) SetNodeId(v string) *EnableAdditionalBandwidthRequest {
	s.NodeId = &v
	return s
}

func (s *EnableAdditionalBandwidthRequest) SetBandwidth(v string) *EnableAdditionalBandwidthRequest {
	s.Bandwidth = &v
	return s
}

func (s *EnableAdditionalBandwidthRequest) SetOrderTimeLength(v string) *EnableAdditionalBandwidthRequest {
	s.OrderTimeLength = &v
	return s
}

type EnableAdditionalBandwidthResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s EnableAdditionalBandwidthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableAdditionalBandwidthResponseBody) GoString() string {
	return s.String()
}

func (s *EnableAdditionalBandwidthResponseBody) SetRequestId(v string) *EnableAdditionalBandwidthResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableAdditionalBandwidthResponseBody) SetOrderId(v string) *EnableAdditionalBandwidthResponseBody {
	s.OrderId = &v
	return s
}

type EnableAdditionalBandwidthResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableAdditionalBandwidthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableAdditionalBandwidthResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableAdditionalBandwidthResponse) GoString() string {
	return s.String()
}

func (s *EnableAdditionalBandwidthResponse) SetHeaders(v map[string]*string) *EnableAdditionalBandwidthResponse {
	s.Headers = v
	return s
}

func (s *EnableAdditionalBandwidthResponse) SetBody(v *EnableAdditionalBandwidthResponseBody) *EnableAdditionalBandwidthResponse {
	s.Body = v
	return s
}

type FlushExpireKeysRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	EffectiveTime        *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
}

func (s FlushExpireKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s FlushExpireKeysRequest) GoString() string {
	return s.String()
}

func (s *FlushExpireKeysRequest) SetSecurityToken(v string) *FlushExpireKeysRequest {
	s.SecurityToken = &v
	return s
}

func (s *FlushExpireKeysRequest) SetOwnerId(v int64) *FlushExpireKeysRequest {
	s.OwnerId = &v
	return s
}

func (s *FlushExpireKeysRequest) SetResourceOwnerAccount(v string) *FlushExpireKeysRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *FlushExpireKeysRequest) SetResourceOwnerId(v int64) *FlushExpireKeysRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *FlushExpireKeysRequest) SetOwnerAccount(v string) *FlushExpireKeysRequest {
	s.OwnerAccount = &v
	return s
}

func (s *FlushExpireKeysRequest) SetInstanceId(v string) *FlushExpireKeysRequest {
	s.InstanceId = &v
	return s
}

func (s *FlushExpireKeysRequest) SetEffectiveTime(v string) *FlushExpireKeysRequest {
	s.EffectiveTime = &v
	return s
}

type FlushExpireKeysResponseBody struct {
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s FlushExpireKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FlushExpireKeysResponseBody) GoString() string {
	return s.String()
}

func (s *FlushExpireKeysResponseBody) SetTaskId(v string) *FlushExpireKeysResponseBody {
	s.TaskId = &v
	return s
}

func (s *FlushExpireKeysResponseBody) SetRequestId(v string) *FlushExpireKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlushExpireKeysResponseBody) SetInstanceId(v string) *FlushExpireKeysResponseBody {
	s.InstanceId = &v
	return s
}

type FlushExpireKeysResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FlushExpireKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FlushExpireKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s FlushExpireKeysResponse) GoString() string {
	return s.String()
}

func (s *FlushExpireKeysResponse) SetHeaders(v map[string]*string) *FlushExpireKeysResponse {
	s.Headers = v
	return s
}

func (s *FlushExpireKeysResponse) SetBody(v *FlushExpireKeysResponseBody) *FlushExpireKeysResponse {
	s.Body = v
	return s
}

type FlushInstanceRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s FlushInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s FlushInstanceRequest) GoString() string {
	return s.String()
}

func (s *FlushInstanceRequest) SetSecurityToken(v string) *FlushInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *FlushInstanceRequest) SetOwnerId(v int64) *FlushInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *FlushInstanceRequest) SetResourceOwnerAccount(v string) *FlushInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *FlushInstanceRequest) SetResourceOwnerId(v int64) *FlushInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *FlushInstanceRequest) SetOwnerAccount(v string) *FlushInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *FlushInstanceRequest) SetInstanceId(v string) *FlushInstanceRequest {
	s.InstanceId = &v
	return s
}

type FlushInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FlushInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FlushInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *FlushInstanceResponseBody) SetRequestId(v string) *FlushInstanceResponseBody {
	s.RequestId = &v
	return s
}

type FlushInstanceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FlushInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FlushInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s FlushInstanceResponse) GoString() string {
	return s.String()
}

func (s *FlushInstanceResponse) SetHeaders(v map[string]*string) *FlushInstanceResponse {
	s.Headers = v
	return s
}

func (s *FlushInstanceResponse) SetBody(v *FlushInstanceResponseBody) *FlushInstanceResponse {
	s.Body = v
	return s
}

type GrantAccountPrivilegeRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPrivilege     *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
}

func (s GrantAccountPrivilegeRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantAccountPrivilegeRequest) GoString() string {
	return s.String()
}

func (s *GrantAccountPrivilegeRequest) SetSecurityToken(v string) *GrantAccountPrivilegeRequest {
	s.SecurityToken = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetOwnerId(v int64) *GrantAccountPrivilegeRequest {
	s.OwnerId = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetResourceOwnerAccount(v string) *GrantAccountPrivilegeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetResourceOwnerId(v int64) *GrantAccountPrivilegeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetOwnerAccount(v string) *GrantAccountPrivilegeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetInstanceId(v string) *GrantAccountPrivilegeRequest {
	s.InstanceId = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetAccountName(v string) *GrantAccountPrivilegeRequest {
	s.AccountName = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetAccountPrivilege(v string) *GrantAccountPrivilegeRequest {
	s.AccountPrivilege = &v
	return s
}

type GrantAccountPrivilegeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GrantAccountPrivilegeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantAccountPrivilegeResponseBody) GoString() string {
	return s.String()
}

func (s *GrantAccountPrivilegeResponseBody) SetRequestId(v string) *GrantAccountPrivilegeResponseBody {
	s.RequestId = &v
	return s
}

type GrantAccountPrivilegeResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GrantAccountPrivilegeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GrantAccountPrivilegeResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantAccountPrivilegeResponse) GoString() string {
	return s.String()
}

func (s *GrantAccountPrivilegeResponse) SetHeaders(v map[string]*string) *GrantAccountPrivilegeResponse {
	s.Headers = v
	return s
}

func (s *GrantAccountPrivilegeResponse) SetBody(v *GrantAccountPrivilegeResponseBody) *GrantAccountPrivilegeResponse {
	s.Body = v
	return s
}

type InitializeKvstorePermissionRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceName          *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s InitializeKvstorePermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s InitializeKvstorePermissionRequest) GoString() string {
	return s.String()
}

func (s *InitializeKvstorePermissionRequest) SetSecurityToken(v string) *InitializeKvstorePermissionRequest {
	s.SecurityToken = &v
	return s
}

func (s *InitializeKvstorePermissionRequest) SetOwnerId(v int64) *InitializeKvstorePermissionRequest {
	s.OwnerId = &v
	return s
}

func (s *InitializeKvstorePermissionRequest) SetResourceOwnerAccount(v string) *InitializeKvstorePermissionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *InitializeKvstorePermissionRequest) SetResourceOwnerId(v int64) *InitializeKvstorePermissionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *InitializeKvstorePermissionRequest) SetOwnerAccount(v string) *InitializeKvstorePermissionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *InitializeKvstorePermissionRequest) SetRegionId(v string) *InitializeKvstorePermissionRequest {
	s.RegionId = &v
	return s
}

func (s *InitializeKvstorePermissionRequest) SetServiceName(v string) *InitializeKvstorePermissionRequest {
	s.ServiceName = &v
	return s
}

type InitializeKvstorePermissionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InitializeKvstorePermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitializeKvstorePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *InitializeKvstorePermissionResponseBody) SetRequestId(v string) *InitializeKvstorePermissionResponseBody {
	s.RequestId = &v
	return s
}

type InitializeKvstorePermissionResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InitializeKvstorePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InitializeKvstorePermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s InitializeKvstorePermissionResponse) GoString() string {
	return s.String()
}

func (s *InitializeKvstorePermissionResponse) SetHeaders(v map[string]*string) *InitializeKvstorePermissionResponse {
	s.Headers = v
	return s
}

func (s *InitializeKvstorePermissionResponse) SetBody(v *InitializeKvstorePermissionResponseBody) *InitializeKvstorePermissionResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	OwnerId              *int64                        `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                       `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                        `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string                       `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType         *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	NextToken            *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceId           []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag                  []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetOwnerId(v int64) *ListTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerAccount(v string) *ListTagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerId(v int64) *ListTagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerAccount(v string) *ListTagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	NextToken    *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type MigrateToOtherZoneRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	EffectiveTime        *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
}

func (s MigrateToOtherZoneRequest) String() string {
	return tea.Prettify(s)
}

func (s MigrateToOtherZoneRequest) GoString() string {
	return s.String()
}

func (s *MigrateToOtherZoneRequest) SetSecurityToken(v string) *MigrateToOtherZoneRequest {
	s.SecurityToken = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetOwnerId(v int64) *MigrateToOtherZoneRequest {
	s.OwnerId = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetResourceOwnerAccount(v string) *MigrateToOtherZoneRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetResourceOwnerId(v int64) *MigrateToOtherZoneRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetOwnerAccount(v string) *MigrateToOtherZoneRequest {
	s.OwnerAccount = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetDBInstanceId(v string) *MigrateToOtherZoneRequest {
	s.DBInstanceId = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetZoneId(v string) *MigrateToOtherZoneRequest {
	s.ZoneId = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetVSwitchId(v string) *MigrateToOtherZoneRequest {
	s.VSwitchId = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetEffectiveTime(v string) *MigrateToOtherZoneRequest {
	s.EffectiveTime = &v
	return s
}

type MigrateToOtherZoneResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MigrateToOtherZoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MigrateToOtherZoneResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateToOtherZoneResponseBody) SetRequestId(v string) *MigrateToOtherZoneResponseBody {
	s.RequestId = &v
	return s
}

type MigrateToOtherZoneResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MigrateToOtherZoneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MigrateToOtherZoneResponse) String() string {
	return tea.Prettify(s)
}

func (s MigrateToOtherZoneResponse) GoString() string {
	return s.String()
}

func (s *MigrateToOtherZoneResponse) SetHeaders(v map[string]*string) *MigrateToOtherZoneResponse {
	s.Headers = v
	return s
}

func (s *MigrateToOtherZoneResponse) SetBody(v *MigrateToOtherZoneResponseBody) *MigrateToOtherZoneResponse {
	s.Body = v
	return s
}

type ModifyAccountDescriptionRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountDescription   *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
}

func (s ModifyAccountDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionRequest) SetSecurityToken(v string) *ModifyAccountDescriptionRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetOwnerId(v int64) *ModifyAccountDescriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetResourceOwnerAccount(v string) *ModifyAccountDescriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetResourceOwnerId(v int64) *ModifyAccountDescriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetOwnerAccount(v string) *ModifyAccountDescriptionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetInstanceId(v string) *ModifyAccountDescriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetAccountName(v string) *ModifyAccountDescriptionRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetAccountDescription(v string) *ModifyAccountDescriptionRequest {
	s.AccountDescription = &v
	return s
}

type ModifyAccountDescriptionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccountDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionResponseBody) SetRequestId(v string) *ModifyAccountDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAccountDescriptionResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyAccountDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAccountDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionResponse) SetHeaders(v map[string]*string) *ModifyAccountDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccountDescriptionResponse) SetBody(v *ModifyAccountDescriptionResponseBody) *ModifyAccountDescriptionResponse {
	s.Body = v
	return s
}

type ModifyAccountPasswordRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	OldAccountPassword   *string `json:"OldAccountPassword,omitempty" xml:"OldAccountPassword,omitempty"`
	NewAccountPassword   *string `json:"NewAccountPassword,omitempty" xml:"NewAccountPassword,omitempty"`
}

func (s ModifyAccountPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountPasswordRequest) SetSecurityToken(v string) *ModifyAccountPasswordRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetOwnerId(v int64) *ModifyAccountPasswordRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetResourceOwnerAccount(v string) *ModifyAccountPasswordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetResourceOwnerId(v int64) *ModifyAccountPasswordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetOwnerAccount(v string) *ModifyAccountPasswordRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetInstanceId(v string) *ModifyAccountPasswordRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetAccountName(v string) *ModifyAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetOldAccountPassword(v string) *ModifyAccountPasswordRequest {
	s.OldAccountPassword = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetNewAccountPassword(v string) *ModifyAccountPasswordRequest {
	s.NewAccountPassword = &v
	return s
}

type ModifyAccountPasswordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccountPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountPasswordResponseBody) SetRequestId(v string) *ModifyAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAccountPasswordResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAccountPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountPasswordResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountPasswordResponse) SetHeaders(v map[string]*string) *ModifyAccountPasswordResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccountPasswordResponse) SetBody(v *ModifyAccountPasswordResponseBody) *ModifyAccountPasswordResponse {
	s.Body = v
	return s
}

type ModifyActiveOperationTaskRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	Ids                  *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	SwitchTime           *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
}

func (s ModifyActiveOperationTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyActiveOperationTaskRequest) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationTaskRequest) SetSecurityToken(v string) *ModifyActiveOperationTaskRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyActiveOperationTaskRequest) SetOwnerId(v int64) *ModifyActiveOperationTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyActiveOperationTaskRequest) SetResourceOwnerAccount(v string) *ModifyActiveOperationTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyActiveOperationTaskRequest) SetResourceOwnerId(v int64) *ModifyActiveOperationTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyActiveOperationTaskRequest) SetOwnerAccount(v string) *ModifyActiveOperationTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyActiveOperationTaskRequest) SetIds(v string) *ModifyActiveOperationTaskRequest {
	s.Ids = &v
	return s
}

func (s *ModifyActiveOperationTaskRequest) SetSwitchTime(v string) *ModifyActiveOperationTaskRequest {
	s.SwitchTime = &v
	return s
}

type ModifyActiveOperationTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Ids       *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
}

func (s ModifyActiveOperationTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyActiveOperationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationTaskResponseBody) SetRequestId(v string) *ModifyActiveOperationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyActiveOperationTaskResponseBody) SetIds(v string) *ModifyActiveOperationTaskResponseBody {
	s.Ids = &v
	return s
}

type ModifyActiveOperationTaskResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyActiveOperationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyActiveOperationTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyActiveOperationTaskResponse) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationTaskResponse) SetHeaders(v map[string]*string) *ModifyActiveOperationTaskResponse {
	s.Headers = v
	return s
}

func (s *ModifyActiveOperationTaskResponse) SetBody(v *ModifyActiveOperationTaskResponseBody) *ModifyActiveOperationTaskResponse {
	s.Body = v
	return s
}

type ModifyAuditLogConfigRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AuditLogSwitchSource *string `json:"AuditLogSwitchSource,omitempty" xml:"AuditLogSwitchSource,omitempty"`
	ServiceType          *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Retention            *int32  `json:"Retention,omitempty" xml:"Retention,omitempty"`
	ProxyAudit           *string `json:"ProxyAudit,omitempty" xml:"ProxyAudit,omitempty"`
	DbAudit              *bool   `json:"DbAudit,omitempty" xml:"DbAudit,omitempty"`
	AuditCommand         *string `json:"AuditCommand,omitempty" xml:"AuditCommand,omitempty"`
}

func (s ModifyAuditLogConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAuditLogConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyAuditLogConfigRequest) SetSecurityToken(v string) *ModifyAuditLogConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetOwnerId(v int64) *ModifyAuditLogConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetResourceOwnerAccount(v string) *ModifyAuditLogConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetResourceOwnerId(v int64) *ModifyAuditLogConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetOwnerAccount(v string) *ModifyAuditLogConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetInstanceId(v string) *ModifyAuditLogConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetAuditLogSwitchSource(v string) *ModifyAuditLogConfigRequest {
	s.AuditLogSwitchSource = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetServiceType(v string) *ModifyAuditLogConfigRequest {
	s.ServiceType = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetRetention(v int32) *ModifyAuditLogConfigRequest {
	s.Retention = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetProxyAudit(v string) *ModifyAuditLogConfigRequest {
	s.ProxyAudit = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetDbAudit(v bool) *ModifyAuditLogConfigRequest {
	s.DbAudit = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetAuditCommand(v string) *ModifyAuditLogConfigRequest {
	s.AuditCommand = &v
	return s
}

type ModifyAuditLogConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAuditLogConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAuditLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAuditLogConfigResponseBody) SetRequestId(v string) *ModifyAuditLogConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAuditLogConfigResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyAuditLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAuditLogConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAuditLogConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyAuditLogConfigResponse) SetHeaders(v map[string]*string) *ModifyAuditLogConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyAuditLogConfigResponse) SetBody(v *ModifyAuditLogConfigResponseBody) *ModifyAuditLogConfigResponse {
	s.Body = v
	return s
}

type ModifyBackupPolicyRequest struct {
	SecurityToken         *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId            *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PreferredBackupTime   *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	EnableBackupLog       *int32  `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
}

func (s ModifyBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequest) SetSecurityToken(v string) *ModifyBackupPolicyRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetOwnerId(v int64) *ModifyBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetResourceOwnerAccount(v string) *ModifyBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetResourceOwnerId(v int64) *ModifyBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetOwnerAccount(v string) *ModifyBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetInstanceId(v string) *ModifyBackupPolicyRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupTime(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupTime = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupPeriod(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetEnableBackupLog(v int32) *ModifyBackupPolicyRequest {
	s.EnableBackupLog = &v
	return s
}

type ModifyBackupPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyResponseBody) SetRequestId(v string) *ModifyBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyBackupPolicyResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyBackupPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyResponse) SetHeaders(v map[string]*string) *ModifyBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyBackupPolicyResponse) SetBody(v *ModifyBackupPolicyResponseBody) *ModifyBackupPolicyResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceConnectionStringRequest struct {
	SecurityToken           *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId                 *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount    *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId         *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount            *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId            *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NewConnectionString     *string `json:"NewConnectionString,omitempty" xml:"NewConnectionString,omitempty"`
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
	Port                    *string `json:"Port,omitempty" xml:"Port,omitempty"`
	IPType                  *string `json:"IPType,omitempty" xml:"IPType,omitempty"`
}

func (s ModifyDBInstanceConnectionStringRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConnectionStringRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionStringRequest) SetSecurityToken(v string) *ModifyDBInstanceConnectionStringRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetOwnerId(v int64) *ModifyDBInstanceConnectionStringRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceConnectionStringRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceConnectionStringRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetOwnerAccount(v string) *ModifyDBInstanceConnectionStringRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetDBInstanceId(v string) *ModifyDBInstanceConnectionStringRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetNewConnectionString(v string) *ModifyDBInstanceConnectionStringRequest {
	s.NewConnectionString = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetCurrentConnectionString(v string) *ModifyDBInstanceConnectionStringRequest {
	s.CurrentConnectionString = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetPort(v string) *ModifyDBInstanceConnectionStringRequest {
	s.Port = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetIPType(v string) *ModifyDBInstanceConnectionStringRequest {
	s.IPType = &v
	return s
}

type ModifyDBInstanceConnectionStringResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceConnectionStringResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConnectionStringResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionStringResponseBody) SetRequestId(v string) *ModifyDBInstanceConnectionStringResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceConnectionStringResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBInstanceConnectionStringResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBInstanceConnectionStringResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConnectionStringResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionStringResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceConnectionStringResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponse) SetBody(v *ModifyDBInstanceConnectionStringResponseBody) *ModifyDBInstanceConnectionStringResponse {
	s.Body = v
	return s
}

type ModifyInstanceAttributeRequest struct {
	SecurityToken             *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId                *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName              *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	NewPassword               *string `json:"NewPassword,omitempty" xml:"NewPassword,omitempty"`
	InstanceReleaseProtection *bool   `json:"InstanceReleaseProtection,omitempty" xml:"InstanceReleaseProtection,omitempty"`
	Product                   *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Category                  *string `json:"Category,omitempty" xml:"Category,omitempty"`
}

func (s ModifyInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeRequest) SetSecurityToken(v string) *ModifyInstanceAttributeRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetOwnerId(v int64) *ModifyInstanceAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetResourceOwnerAccount(v string) *ModifyInstanceAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetResourceOwnerId(v int64) *ModifyInstanceAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetOwnerAccount(v string) *ModifyInstanceAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetInstanceId(v string) *ModifyInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetInstanceName(v string) *ModifyInstanceAttributeRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetNewPassword(v string) *ModifyInstanceAttributeRequest {
	s.NewPassword = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetInstanceReleaseProtection(v bool) *ModifyInstanceAttributeRequest {
	s.InstanceReleaseProtection = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetProduct(v string) *ModifyInstanceAttributeRequest {
	s.Product = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetCategory(v string) *ModifyInstanceAttributeRequest {
	s.Category = &v
	return s
}

type ModifyInstanceAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeResponseBody) SetRequestId(v string) *ModifyInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceAttributeResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeResponse) SetHeaders(v map[string]*string) *ModifyInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceAttributeResponse) SetBody(v *ModifyInstanceAttributeResponseBody) *ModifyInstanceAttributeResponse {
	s.Body = v
	return s
}

type ModifyInstanceAutoRenewalAttributeRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Duration             *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	AutoRenew            *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
}

func (s ModifyInstanceAutoRenewalAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceAutoRenewalAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAutoRenewalAttributeRequest) SetSecurityToken(v string) *ModifyInstanceAutoRenewalAttributeRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyInstanceAutoRenewalAttributeRequest) SetOwnerId(v int64) *ModifyInstanceAutoRenewalAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceAutoRenewalAttributeRequest) SetResourceOwnerAccount(v string) *ModifyInstanceAutoRenewalAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceAutoRenewalAttributeRequest) SetResourceOwnerId(v int64) *ModifyInstanceAutoRenewalAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceAutoRenewalAttributeRequest) SetOwnerAccount(v string) *ModifyInstanceAutoRenewalAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceAutoRenewalAttributeRequest) SetDBInstanceId(v string) *ModifyInstanceAutoRenewalAttributeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyInstanceAutoRenewalAttributeRequest) SetDuration(v string) *ModifyInstanceAutoRenewalAttributeRequest {
	s.Duration = &v
	return s
}

func (s *ModifyInstanceAutoRenewalAttributeRequest) SetAutoRenew(v string) *ModifyInstanceAutoRenewalAttributeRequest {
	s.AutoRenew = &v
	return s
}

type ModifyInstanceAutoRenewalAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceAutoRenewalAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceAutoRenewalAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAutoRenewalAttributeResponseBody) SetRequestId(v string) *ModifyInstanceAutoRenewalAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceAutoRenewalAttributeResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceAutoRenewalAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceAutoRenewalAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceAutoRenewalAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAutoRenewalAttributeResponse) SetHeaders(v map[string]*string) *ModifyInstanceAutoRenewalAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceAutoRenewalAttributeResponse) SetBody(v *ModifyInstanceAutoRenewalAttributeResponseBody) *ModifyInstanceAutoRenewalAttributeResponse {
	s.Body = v
	return s
}

type ModifyInstanceConfigRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Config               *string `json:"Config,omitempty" xml:"Config,omitempty"`
}

func (s ModifyInstanceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceConfigRequest) SetSecurityToken(v string) *ModifyInstanceConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetOwnerId(v int64) *ModifyInstanceConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetResourceOwnerAccount(v string) *ModifyInstanceConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetResourceOwnerId(v int64) *ModifyInstanceConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetOwnerAccount(v string) *ModifyInstanceConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetInstanceId(v string) *ModifyInstanceConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetConfig(v string) *ModifyInstanceConfigRequest {
	s.Config = &v
	return s
}

type ModifyInstanceConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceConfigResponseBody) SetRequestId(v string) *ModifyInstanceConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceConfigResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceConfigResponse) SetHeaders(v map[string]*string) *ModifyInstanceConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceConfigResponse) SetBody(v *ModifyInstanceConfigResponseBody) *ModifyInstanceConfigResponse {
	s.Body = v
	return s
}

type ModifyInstanceMaintainTimeRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaintainStartTime    *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	MaintainEndTime      *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
}

func (s ModifyInstanceMaintainTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceMaintainTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMaintainTimeRequest) SetSecurityToken(v string) *ModifyInstanceMaintainTimeRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyInstanceMaintainTimeRequest) SetOwnerId(v int64) *ModifyInstanceMaintainTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceMaintainTimeRequest) SetResourceOwnerAccount(v string) *ModifyInstanceMaintainTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceMaintainTimeRequest) SetResourceOwnerId(v int64) *ModifyInstanceMaintainTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceMaintainTimeRequest) SetOwnerAccount(v string) *ModifyInstanceMaintainTimeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceMaintainTimeRequest) SetInstanceId(v string) *ModifyInstanceMaintainTimeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceMaintainTimeRequest) SetMaintainStartTime(v string) *ModifyInstanceMaintainTimeRequest {
	s.MaintainStartTime = &v
	return s
}

func (s *ModifyInstanceMaintainTimeRequest) SetMaintainEndTime(v string) *ModifyInstanceMaintainTimeRequest {
	s.MaintainEndTime = &v
	return s
}

type ModifyInstanceMaintainTimeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceMaintainTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceMaintainTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMaintainTimeResponseBody) SetRequestId(v string) *ModifyInstanceMaintainTimeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceMaintainTimeResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceMaintainTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceMaintainTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceMaintainTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMaintainTimeResponse) SetHeaders(v map[string]*string) *ModifyInstanceMaintainTimeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceMaintainTimeResponse) SetBody(v *ModifyInstanceMaintainTimeResponseBody) *ModifyInstanceMaintainTimeResponse {
	s.Body = v
	return s
}

type ModifyInstanceMajorVersionRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MajorVersion         *string `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	EffectTime           *string `json:"EffectTime,omitempty" xml:"EffectTime,omitempty"`
	EffectiveTime        *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
}

func (s ModifyInstanceMajorVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceMajorVersionRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMajorVersionRequest) SetSecurityToken(v string) *ModifyInstanceMajorVersionRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyInstanceMajorVersionRequest) SetOwnerId(v int64) *ModifyInstanceMajorVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceMajorVersionRequest) SetResourceOwnerAccount(v string) *ModifyInstanceMajorVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceMajorVersionRequest) SetResourceOwnerId(v int64) *ModifyInstanceMajorVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceMajorVersionRequest) SetOwnerAccount(v string) *ModifyInstanceMajorVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceMajorVersionRequest) SetInstanceId(v string) *ModifyInstanceMajorVersionRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceMajorVersionRequest) SetMajorVersion(v string) *ModifyInstanceMajorVersionRequest {
	s.MajorVersion = &v
	return s
}

func (s *ModifyInstanceMajorVersionRequest) SetEffectTime(v string) *ModifyInstanceMajorVersionRequest {
	s.EffectTime = &v
	return s
}

func (s *ModifyInstanceMajorVersionRequest) SetEffectiveTime(v string) *ModifyInstanceMajorVersionRequest {
	s.EffectiveTime = &v
	return s
}

type ModifyInstanceMajorVersionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceMajorVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceMajorVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMajorVersionResponseBody) SetRequestId(v string) *ModifyInstanceMajorVersionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceMajorVersionResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceMajorVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceMajorVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceMajorVersionResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMajorVersionResponse) SetHeaders(v map[string]*string) *ModifyInstanceMajorVersionResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceMajorVersionResponse) SetBody(v *ModifyInstanceMajorVersionResponseBody) *ModifyInstanceMajorVersionResponse {
	s.Body = v
	return s
}

type ModifyInstanceMinorVersionRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Minorversion         *string `json:"Minorversion,omitempty" xml:"Minorversion,omitempty"`
	EffectTime           *string `json:"EffectTime,omitempty" xml:"EffectTime,omitempty"`
	EffectiveTime        *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
}

func (s ModifyInstanceMinorVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceMinorVersionRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMinorVersionRequest) SetSecurityToken(v string) *ModifyInstanceMinorVersionRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyInstanceMinorVersionRequest) SetOwnerId(v int64) *ModifyInstanceMinorVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceMinorVersionRequest) SetResourceOwnerAccount(v string) *ModifyInstanceMinorVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceMinorVersionRequest) SetResourceOwnerId(v int64) *ModifyInstanceMinorVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceMinorVersionRequest) SetOwnerAccount(v string) *ModifyInstanceMinorVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceMinorVersionRequest) SetInstanceId(v string) *ModifyInstanceMinorVersionRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceMinorVersionRequest) SetMinorversion(v string) *ModifyInstanceMinorVersionRequest {
	s.Minorversion = &v
	return s
}

func (s *ModifyInstanceMinorVersionRequest) SetEffectTime(v string) *ModifyInstanceMinorVersionRequest {
	s.EffectTime = &v
	return s
}

func (s *ModifyInstanceMinorVersionRequest) SetEffectiveTime(v string) *ModifyInstanceMinorVersionRequest {
	s.EffectiveTime = &v
	return s
}

type ModifyInstanceMinorVersionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceMinorVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceMinorVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMinorVersionResponseBody) SetRequestId(v string) *ModifyInstanceMinorVersionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceMinorVersionResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceMinorVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceMinorVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceMinorVersionResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMinorVersionResponse) SetHeaders(v map[string]*string) *ModifyInstanceMinorVersionResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceMinorVersionResponse) SetBody(v *ModifyInstanceMinorVersionResponseBody) *ModifyInstanceMinorVersionResponse {
	s.Body = v
	return s
}

type ModifyInstanceNetExpireTimeRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ConnectionString     *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	ClassicExpiredDays   *int32  `json:"ClassicExpiredDays,omitempty" xml:"ClassicExpiredDays,omitempty"`
}

func (s ModifyInstanceNetExpireTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceNetExpireTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNetExpireTimeRequest) SetSecurityToken(v string) *ModifyInstanceNetExpireTimeRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyInstanceNetExpireTimeRequest) SetOwnerId(v int64) *ModifyInstanceNetExpireTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceNetExpireTimeRequest) SetResourceOwnerAccount(v string) *ModifyInstanceNetExpireTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceNetExpireTimeRequest) SetResourceOwnerId(v int64) *ModifyInstanceNetExpireTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceNetExpireTimeRequest) SetOwnerAccount(v string) *ModifyInstanceNetExpireTimeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceNetExpireTimeRequest) SetInstanceId(v string) *ModifyInstanceNetExpireTimeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceNetExpireTimeRequest) SetConnectionString(v string) *ModifyInstanceNetExpireTimeRequest {
	s.ConnectionString = &v
	return s
}

func (s *ModifyInstanceNetExpireTimeRequest) SetClassicExpiredDays(v int32) *ModifyInstanceNetExpireTimeRequest {
	s.ClassicExpiredDays = &v
	return s
}

type ModifyInstanceNetExpireTimeResponseBody struct {
	NetInfoItems *ModifyInstanceNetExpireTimeResponseBodyNetInfoItems `json:"NetInfoItems,omitempty" xml:"NetInfoItems,omitempty" type:"Struct"`
	RequestId    *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceId   *string                                              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ModifyInstanceNetExpireTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceNetExpireTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNetExpireTimeResponseBody) SetNetInfoItems(v *ModifyInstanceNetExpireTimeResponseBodyNetInfoItems) *ModifyInstanceNetExpireTimeResponseBody {
	s.NetInfoItems = v
	return s
}

func (s *ModifyInstanceNetExpireTimeResponseBody) SetRequestId(v string) *ModifyInstanceNetExpireTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceNetExpireTimeResponseBody) SetInstanceId(v string) *ModifyInstanceNetExpireTimeResponseBody {
	s.InstanceId = &v
	return s
}

type ModifyInstanceNetExpireTimeResponseBodyNetInfoItems struct {
	NetInfoItem []*ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem `json:"NetInfoItem,omitempty" xml:"NetInfoItem,omitempty" type:"Repeated"`
}

func (s ModifyInstanceNetExpireTimeResponseBodyNetInfoItems) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceNetExpireTimeResponseBodyNetInfoItems) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNetExpireTimeResponseBodyNetInfoItems) SetNetInfoItem(v []*ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem) *ModifyInstanceNetExpireTimeResponseBodyNetInfoItems {
	s.NetInfoItem = v
	return s
}

type ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem struct {
	DBInstanceNetType *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	ConnectionString  *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	ExpiredTime       *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Port              *string `json:"Port,omitempty" xml:"Port,omitempty"`
	IPAddress         *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
}

func (s ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem) SetDBInstanceNetType(v string) *ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem {
	s.DBInstanceNetType = &v
	return s
}

func (s *ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem) SetConnectionString(v string) *ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem {
	s.ConnectionString = &v
	return s
}

func (s *ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem) SetExpiredTime(v string) *ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem {
	s.ExpiredTime = &v
	return s
}

func (s *ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem) SetPort(v string) *ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem {
	s.Port = &v
	return s
}

func (s *ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem) SetIPAddress(v string) *ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem {
	s.IPAddress = &v
	return s
}

type ModifyInstanceNetExpireTimeResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceNetExpireTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceNetExpireTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceNetExpireTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNetExpireTimeResponse) SetHeaders(v map[string]*string) *ModifyInstanceNetExpireTimeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceNetExpireTimeResponse) SetBody(v *ModifyInstanceNetExpireTimeResponseBody) *ModifyInstanceNetExpireTimeResponse {
	s.Body = v
	return s
}

type ModifyInstanceSpecRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceClass        *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	FromApp              *string `json:"FromApp,omitempty" xml:"FromApp,omitempty"`
	BusinessInfo         *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	CouponNo             *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	ForceUpgrade         *bool   `json:"ForceUpgrade,omitempty" xml:"ForceUpgrade,omitempty"`
	EffectiveTime        *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	AutoPay              *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	OrderType            *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	MajorVersion         *string `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	SourceBiz            *string `json:"SourceBiz,omitempty" xml:"SourceBiz,omitempty"`
}

func (s ModifyInstanceSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSpecRequest) SetSecurityToken(v string) *ModifyInstanceSpecRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetOwnerId(v int64) *ModifyInstanceSpecRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetResourceOwnerAccount(v string) *ModifyInstanceSpecRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetResourceOwnerId(v int64) *ModifyInstanceSpecRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetOwnerAccount(v string) *ModifyInstanceSpecRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetInstanceId(v string) *ModifyInstanceSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetInstanceClass(v string) *ModifyInstanceSpecRequest {
	s.InstanceClass = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetFromApp(v string) *ModifyInstanceSpecRequest {
	s.FromApp = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetBusinessInfo(v string) *ModifyInstanceSpecRequest {
	s.BusinessInfo = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetCouponNo(v string) *ModifyInstanceSpecRequest {
	s.CouponNo = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetForceUpgrade(v bool) *ModifyInstanceSpecRequest {
	s.ForceUpgrade = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetEffectiveTime(v string) *ModifyInstanceSpecRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetAutoPay(v bool) *ModifyInstanceSpecRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetOrderType(v string) *ModifyInstanceSpecRequest {
	s.OrderType = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetMajorVersion(v string) *ModifyInstanceSpecRequest {
	s.MajorVersion = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetClientToken(v string) *ModifyInstanceSpecRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetSourceBiz(v string) *ModifyInstanceSpecRequest {
	s.SourceBiz = &v
	return s
}

type ModifyInstanceSpecResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ModifyInstanceSpecResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSpecResponseBody) SetRequestId(v string) *ModifyInstanceSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceSpecResponseBody) SetOrderId(v string) *ModifyInstanceSpecResponseBody {
	s.OrderId = &v
	return s
}

type ModifyInstanceSpecResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceSpecResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceSpecResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSpecResponse) SetHeaders(v map[string]*string) *ModifyInstanceSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceSpecResponse) SetBody(v *ModifyInstanceSpecResponseBody) *ModifyInstanceSpecResponse {
	s.Body = v
	return s
}

type ModifyInstanceSSLRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SSLEnabled           *string `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
}

func (s ModifyInstanceSSLRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceSSLRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSSLRequest) SetSecurityToken(v string) *ModifyInstanceSSLRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyInstanceSSLRequest) SetOwnerId(v int64) *ModifyInstanceSSLRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceSSLRequest) SetResourceOwnerAccount(v string) *ModifyInstanceSSLRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceSSLRequest) SetResourceOwnerId(v int64) *ModifyInstanceSSLRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceSSLRequest) SetOwnerAccount(v string) *ModifyInstanceSSLRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceSSLRequest) SetInstanceId(v string) *ModifyInstanceSSLRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceSSLRequest) SetSSLEnabled(v string) *ModifyInstanceSSLRequest {
	s.SSLEnabled = &v
	return s
}

type ModifyInstanceSSLResponseBody struct {
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ModifyInstanceSSLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceSSLResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSSLResponseBody) SetTaskId(v string) *ModifyInstanceSSLResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyInstanceSSLResponseBody) SetRequestId(v string) *ModifyInstanceSSLResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceSSLResponseBody) SetInstanceId(v string) *ModifyInstanceSSLResponseBody {
	s.InstanceId = &v
	return s
}

type ModifyInstanceSSLResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceSSLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceSSLResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceSSLResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSSLResponse) SetHeaders(v map[string]*string) *ModifyInstanceSSLResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceSSLResponse) SetBody(v *ModifyInstanceSSLResponseBody) *ModifyInstanceSSLResponse {
	s.Body = v
	return s
}

type ModifyInstanceVpcAuthModeRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VpcAuthMode          *string `json:"VpcAuthMode,omitempty" xml:"VpcAuthMode,omitempty"`
}

func (s ModifyInstanceVpcAuthModeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceVpcAuthModeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceVpcAuthModeRequest) SetSecurityToken(v string) *ModifyInstanceVpcAuthModeRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyInstanceVpcAuthModeRequest) SetOwnerId(v int64) *ModifyInstanceVpcAuthModeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceVpcAuthModeRequest) SetResourceOwnerAccount(v string) *ModifyInstanceVpcAuthModeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceVpcAuthModeRequest) SetResourceOwnerId(v int64) *ModifyInstanceVpcAuthModeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceVpcAuthModeRequest) SetOwnerAccount(v string) *ModifyInstanceVpcAuthModeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceVpcAuthModeRequest) SetInstanceId(v string) *ModifyInstanceVpcAuthModeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceVpcAuthModeRequest) SetVpcAuthMode(v string) *ModifyInstanceVpcAuthModeRequest {
	s.VpcAuthMode = &v
	return s
}

type ModifyInstanceVpcAuthModeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceVpcAuthModeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceVpcAuthModeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceVpcAuthModeResponseBody) SetRequestId(v string) *ModifyInstanceVpcAuthModeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceVpcAuthModeResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceVpcAuthModeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceVpcAuthModeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceVpcAuthModeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceVpcAuthModeResponse) SetHeaders(v map[string]*string) *ModifyInstanceVpcAuthModeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceVpcAuthModeResponse) SetBody(v *ModifyInstanceVpcAuthModeResponseBody) *ModifyInstanceVpcAuthModeResponse {
	s.Body = v
	return s
}

type ModifyIntranetAttributeRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	BandWidth            *int64  `json:"BandWidth,omitempty" xml:"BandWidth,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s ModifyIntranetAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyIntranetAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyIntranetAttributeRequest) SetSecurityToken(v string) *ModifyIntranetAttributeRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyIntranetAttributeRequest) SetOwnerId(v int64) *ModifyIntranetAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyIntranetAttributeRequest) SetResourceOwnerAccount(v string) *ModifyIntranetAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyIntranetAttributeRequest) SetResourceOwnerId(v int64) *ModifyIntranetAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyIntranetAttributeRequest) SetOwnerAccount(v string) *ModifyIntranetAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyIntranetAttributeRequest) SetInstanceId(v string) *ModifyIntranetAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyIntranetAttributeRequest) SetBandWidth(v int64) *ModifyIntranetAttributeRequest {
	s.BandWidth = &v
	return s
}

func (s *ModifyIntranetAttributeRequest) SetNodeId(v string) *ModifyIntranetAttributeRequest {
	s.NodeId = &v
	return s
}

type ModifyIntranetAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyIntranetAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyIntranetAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIntranetAttributeResponseBody) SetRequestId(v string) *ModifyIntranetAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyIntranetAttributeResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyIntranetAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyIntranetAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyIntranetAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyIntranetAttributeResponse) SetHeaders(v map[string]*string) *ModifyIntranetAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyIntranetAttributeResponse) SetBody(v *ModifyIntranetAttributeResponseBody) *ModifyIntranetAttributeResponse {
	s.Body = v
	return s
}

type ModifyNodeSpecRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceClass        *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	BusinessInfo         *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	CouponNo             *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	AutoPay              *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	OrderType            *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	SwitchTimeMode       *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
}

func (s ModifyNodeSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodeSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyNodeSpecRequest) SetSecurityToken(v string) *ModifyNodeSpecRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetOwnerId(v int64) *ModifyNodeSpecRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetResourceOwnerAccount(v string) *ModifyNodeSpecRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetResourceOwnerId(v int64) *ModifyNodeSpecRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetOwnerAccount(v string) *ModifyNodeSpecRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetInstanceId(v string) *ModifyNodeSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetInstanceClass(v string) *ModifyNodeSpecRequest {
	s.InstanceClass = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetBusinessInfo(v string) *ModifyNodeSpecRequest {
	s.BusinessInfo = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetCouponNo(v string) *ModifyNodeSpecRequest {
	s.CouponNo = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetAutoPay(v bool) *ModifyNodeSpecRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetOrderType(v string) *ModifyNodeSpecRequest {
	s.OrderType = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetNodeId(v string) *ModifyNodeSpecRequest {
	s.NodeId = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetSwitchTimeMode(v string) *ModifyNodeSpecRequest {
	s.SwitchTimeMode = &v
	return s
}

type ModifyNodeSpecResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ModifyNodeSpecResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodeSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNodeSpecResponseBody) SetRequestId(v string) *ModifyNodeSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNodeSpecResponseBody) SetOrderId(v int64) *ModifyNodeSpecResponseBody {
	s.OrderId = &v
	return s
}

type ModifyNodeSpecResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyNodeSpecResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyNodeSpecResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodeSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyNodeSpecResponse) SetHeaders(v map[string]*string) *ModifyNodeSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyNodeSpecResponse) SetBody(v *ModifyNodeSpecResponseBody) *ModifyNodeSpecResponse {
	s.Body = v
	return s
}

type ModifyResourceGroupRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Product              *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Category             *string `json:"Category,omitempty" xml:"Category,omitempty"`
}

func (s ModifyResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyResourceGroupRequest) SetOwnerId(v int64) *ModifyResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyResourceGroupRequest) SetResourceOwnerAccount(v string) *ModifyResourceGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyResourceGroupRequest) SetResourceOwnerId(v int64) *ModifyResourceGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyResourceGroupRequest) SetClientToken(v string) *ModifyResourceGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyResourceGroupRequest) SetOwnerAccount(v string) *ModifyResourceGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyResourceGroupRequest) SetInstanceId(v string) *ModifyResourceGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyResourceGroupRequest) SetResourceGroupId(v string) *ModifyResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyResourceGroupRequest) SetRegionId(v string) *ModifyResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyResourceGroupRequest) SetProduct(v string) *ModifyResourceGroupRequest {
	s.Product = &v
	return s
}

func (s *ModifyResourceGroupRequest) SetCategory(v string) *ModifyResourceGroupRequest {
	s.Category = &v
	return s
}

type ModifyResourceGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyResourceGroupResponseBody) SetRequestId(v string) *ModifyResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyResourceGroupResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyResourceGroupResponse) SetHeaders(v map[string]*string) *ModifyResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyResourceGroupResponse) SetBody(v *ModifyResourceGroupResponseBody) *ModifyResourceGroupResponse {
	s.Body = v
	return s
}

type ModifySecurityGroupConfigurationRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	SecurityGroupId      *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s ModifySecurityGroupConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityGroupConfigurationRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupConfigurationRequest) SetSecurityToken(v string) *ModifySecurityGroupConfigurationRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifySecurityGroupConfigurationRequest) SetOwnerId(v int64) *ModifySecurityGroupConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySecurityGroupConfigurationRequest) SetResourceOwnerAccount(v string) *ModifySecurityGroupConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySecurityGroupConfigurationRequest) SetResourceOwnerId(v int64) *ModifySecurityGroupConfigurationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySecurityGroupConfigurationRequest) SetOwnerAccount(v string) *ModifySecurityGroupConfigurationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySecurityGroupConfigurationRequest) SetDBInstanceId(v string) *ModifySecurityGroupConfigurationRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifySecurityGroupConfigurationRequest) SetSecurityGroupId(v string) *ModifySecurityGroupConfigurationRequest {
	s.SecurityGroupId = &v
	return s
}

type ModifySecurityGroupConfigurationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySecurityGroupConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityGroupConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupConfigurationResponseBody) SetRequestId(v string) *ModifySecurityGroupConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type ModifySecurityGroupConfigurationResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifySecurityGroupConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySecurityGroupConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityGroupConfigurationResponse) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupConfigurationResponse) SetHeaders(v map[string]*string) *ModifySecurityGroupConfigurationResponse {
	s.Headers = v
	return s
}

func (s *ModifySecurityGroupConfigurationResponse) SetBody(v *ModifySecurityGroupConfigurationResponseBody) *ModifySecurityGroupConfigurationResponse {
	s.Body = v
	return s
}

type ModifySecurityIpsRequest struct {
	SecurityToken            *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount     *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId          *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount             *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId               *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SecurityIps              *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
	SecurityIpGroupName      *string `json:"SecurityIpGroupName,omitempty" xml:"SecurityIpGroupName,omitempty"`
	SecurityIpGroupAttribute *string `json:"SecurityIpGroupAttribute,omitempty" xml:"SecurityIpGroupAttribute,omitempty"`
	ModifyMode               *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
}

func (s ModifySecurityIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsRequest) SetSecurityToken(v string) *ModifySecurityIpsRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetOwnerId(v int64) *ModifySecurityIpsRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetResourceOwnerAccount(v string) *ModifySecurityIpsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetResourceOwnerId(v int64) *ModifySecurityIpsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetOwnerAccount(v string) *ModifySecurityIpsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetRegionId(v string) *ModifySecurityIpsRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetInstanceId(v string) *ModifySecurityIpsRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetSecurityIps(v string) *ModifySecurityIpsRequest {
	s.SecurityIps = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetSecurityIpGroupName(v string) *ModifySecurityIpsRequest {
	s.SecurityIpGroupName = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetSecurityIpGroupAttribute(v string) *ModifySecurityIpsRequest {
	s.SecurityIpGroupAttribute = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetModifyMode(v string) *ModifySecurityIpsRequest {
	s.ModifyMode = &v
	return s
}

type ModifySecurityIpsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySecurityIpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityIpsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsResponseBody) SetRequestId(v string) *ModifySecurityIpsResponseBody {
	s.RequestId = &v
	return s
}

type ModifySecurityIpsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifySecurityIpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySecurityIpsResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityIpsResponse) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsResponse) SetHeaders(v map[string]*string) *ModifySecurityIpsResponse {
	s.Headers = v
	return s
}

func (s *ModifySecurityIpsResponse) SetBody(v *ModifySecurityIpsResponseBody) *ModifySecurityIpsResponse {
	s.Body = v
	return s
}

type ModifyUserClusterHostRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	HostId               *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	AllocationStatus     *int32  `json:"AllocationStatus,omitempty" xml:"AllocationStatus,omitempty"`
	Engine               *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyUserClusterHostRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserClusterHostRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserClusterHostRequest) SetSecurityToken(v string) *ModifyUserClusterHostRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyUserClusterHostRequest) SetOwnerId(v int64) *ModifyUserClusterHostRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyUserClusterHostRequest) SetResourceOwnerAccount(v string) *ModifyUserClusterHostRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyUserClusterHostRequest) SetResourceOwnerId(v int64) *ModifyUserClusterHostRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyUserClusterHostRequest) SetOwnerAccount(v string) *ModifyUserClusterHostRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyUserClusterHostRequest) SetRegionId(v string) *ModifyUserClusterHostRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyUserClusterHostRequest) SetHostId(v string) *ModifyUserClusterHostRequest {
	s.HostId = &v
	return s
}

func (s *ModifyUserClusterHostRequest) SetClusterId(v string) *ModifyUserClusterHostRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyUserClusterHostRequest) SetAllocationStatus(v int32) *ModifyUserClusterHostRequest {
	s.AllocationStatus = &v
	return s
}

func (s *ModifyUserClusterHostRequest) SetEngine(v string) *ModifyUserClusterHostRequest {
	s.Engine = &v
	return s
}

func (s *ModifyUserClusterHostRequest) SetZoneId(v string) *ModifyUserClusterHostRequest {
	s.ZoneId = &v
	return s
}

type ModifyUserClusterHostResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserClusterHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserClusterHostResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserClusterHostResponseBody) SetRequestId(v string) *ModifyUserClusterHostResponseBody {
	s.RequestId = &v
	return s
}

type ModifyUserClusterHostResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyUserClusterHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyUserClusterHostResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserClusterHostResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserClusterHostResponse) SetHeaders(v map[string]*string) *ModifyUserClusterHostResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserClusterHostResponse) SetBody(v *ModifyUserClusterHostResponseBody) *ModifyUserClusterHostResponse {
	s.Body = v
	return s
}

type ReleaseDirectConnectionRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ReleaseDirectConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseDirectConnectionRequest) GoString() string {
	return s.String()
}

func (s *ReleaseDirectConnectionRequest) SetSecurityToken(v string) *ReleaseDirectConnectionRequest {
	s.SecurityToken = &v
	return s
}

func (s *ReleaseDirectConnectionRequest) SetOwnerId(v int64) *ReleaseDirectConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseDirectConnectionRequest) SetResourceOwnerAccount(v string) *ReleaseDirectConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleaseDirectConnectionRequest) SetResourceOwnerId(v int64) *ReleaseDirectConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleaseDirectConnectionRequest) SetOwnerAccount(v string) *ReleaseDirectConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReleaseDirectConnectionRequest) SetInstanceId(v string) *ReleaseDirectConnectionRequest {
	s.InstanceId = &v
	return s
}

type ReleaseDirectConnectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseDirectConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseDirectConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseDirectConnectionResponseBody) SetRequestId(v string) *ReleaseDirectConnectionResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseDirectConnectionResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleaseDirectConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseDirectConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseDirectConnectionResponse) GoString() string {
	return s.String()
}

func (s *ReleaseDirectConnectionResponse) SetHeaders(v map[string]*string) *ReleaseDirectConnectionResponse {
	s.Headers = v
	return s
}

func (s *ReleaseDirectConnectionResponse) SetBody(v *ReleaseDirectConnectionResponseBody) *ReleaseDirectConnectionResponse {
	s.Body = v
	return s
}

type ReleaseInstancePublicConnectionRequest struct {
	SecurityToken           *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId                 *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount    *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId         *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount            *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId              *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
}

func (s ReleaseInstancePublicConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstancePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstancePublicConnectionRequest) SetSecurityToken(v string) *ReleaseInstancePublicConnectionRequest {
	s.SecurityToken = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetOwnerId(v int64) *ReleaseInstancePublicConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetResourceOwnerAccount(v string) *ReleaseInstancePublicConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetResourceOwnerId(v int64) *ReleaseInstancePublicConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetOwnerAccount(v string) *ReleaseInstancePublicConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetInstanceId(v string) *ReleaseInstancePublicConnectionRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetCurrentConnectionString(v string) *ReleaseInstancePublicConnectionRequest {
	s.CurrentConnectionString = &v
	return s
}

type ReleaseInstancePublicConnectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseInstancePublicConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstancePublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseInstancePublicConnectionResponseBody) SetRequestId(v string) *ReleaseInstancePublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseInstancePublicConnectionResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleaseInstancePublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseInstancePublicConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstancePublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *ReleaseInstancePublicConnectionResponse) SetHeaders(v map[string]*string) *ReleaseInstancePublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *ReleaseInstancePublicConnectionResponse) SetBody(v *ReleaseInstancePublicConnectionResponseBody) *ReleaseInstancePublicConnectionResponse {
	s.Body = v
	return s
}

type RenewInstanceRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Capacity             *string `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	InstanceClass        *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	Period               *int64  `json:"Period,omitempty" xml:"Period,omitempty"`
	AutoPay              *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	FromApp              *string `json:"FromApp,omitempty" xml:"FromApp,omitempty"`
	BusinessInfo         *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	CouponNo             *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	ForceUpgrade         *bool   `json:"ForceUpgrade,omitempty" xml:"ForceUpgrade,omitempty"`
}

func (s RenewInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewInstanceRequest) SetSecurityToken(v string) *RenewInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *RenewInstanceRequest) SetOwnerId(v int64) *RenewInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RenewInstanceRequest) SetResourceOwnerAccount(v string) *RenewInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RenewInstanceRequest) SetResourceOwnerId(v int64) *RenewInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RenewInstanceRequest) SetOwnerAccount(v string) *RenewInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RenewInstanceRequest) SetInstanceId(v string) *RenewInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RenewInstanceRequest) SetCapacity(v string) *RenewInstanceRequest {
	s.Capacity = &v
	return s
}

func (s *RenewInstanceRequest) SetInstanceClass(v string) *RenewInstanceRequest {
	s.InstanceClass = &v
	return s
}

func (s *RenewInstanceRequest) SetPeriod(v int64) *RenewInstanceRequest {
	s.Period = &v
	return s
}

func (s *RenewInstanceRequest) SetAutoPay(v bool) *RenewInstanceRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewInstanceRequest) SetFromApp(v string) *RenewInstanceRequest {
	s.FromApp = &v
	return s
}

func (s *RenewInstanceRequest) SetBusinessInfo(v string) *RenewInstanceRequest {
	s.BusinessInfo = &v
	return s
}

func (s *RenewInstanceRequest) SetCouponNo(v string) *RenewInstanceRequest {
	s.CouponNo = &v
	return s
}

func (s *RenewInstanceRequest) SetForceUpgrade(v bool) *RenewInstanceRequest {
	s.ForceUpgrade = &v
	return s
}

type RenewInstanceResponseBody struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s RenewInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponseBody) SetEndTime(v string) *RenewInstanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *RenewInstanceResponseBody) SetRequestId(v string) *RenewInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewInstanceResponseBody) SetOrderId(v string) *RenewInstanceResponseBody {
	s.OrderId = &v
	return s
}

type RenewInstanceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenewInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenewInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceResponse) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponse) SetHeaders(v map[string]*string) *RenewInstanceResponse {
	s.Headers = v
	return s
}

func (s *RenewInstanceResponse) SetBody(v *RenewInstanceResponseBody) *RenewInstanceResponse {
	s.Body = v
	return s
}

type ReplaceUserClusterHostRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	HostId               *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	Engine               *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ReplaceUserClusterHostRequest) String() string {
	return tea.Prettify(s)
}

func (s ReplaceUserClusterHostRequest) GoString() string {
	return s.String()
}

func (s *ReplaceUserClusterHostRequest) SetSecurityToken(v string) *ReplaceUserClusterHostRequest {
	s.SecurityToken = &v
	return s
}

func (s *ReplaceUserClusterHostRequest) SetOwnerId(v int64) *ReplaceUserClusterHostRequest {
	s.OwnerId = &v
	return s
}

func (s *ReplaceUserClusterHostRequest) SetResourceOwnerAccount(v string) *ReplaceUserClusterHostRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReplaceUserClusterHostRequest) SetResourceOwnerId(v int64) *ReplaceUserClusterHostRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReplaceUserClusterHostRequest) SetOwnerAccount(v string) *ReplaceUserClusterHostRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReplaceUserClusterHostRequest) SetRegionId(v string) *ReplaceUserClusterHostRequest {
	s.RegionId = &v
	return s
}

func (s *ReplaceUserClusterHostRequest) SetClusterId(v string) *ReplaceUserClusterHostRequest {
	s.ClusterId = &v
	return s
}

func (s *ReplaceUserClusterHostRequest) SetHostId(v string) *ReplaceUserClusterHostRequest {
	s.HostId = &v
	return s
}

func (s *ReplaceUserClusterHostRequest) SetEngine(v string) *ReplaceUserClusterHostRequest {
	s.Engine = &v
	return s
}

func (s *ReplaceUserClusterHostRequest) SetZoneId(v string) *ReplaceUserClusterHostRequest {
	s.ZoneId = &v
	return s
}

type ReplaceUserClusterHostResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	NewHostId *string `json:"NewHostId,omitempty" xml:"NewHostId,omitempty"`
}

func (s ReplaceUserClusterHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReplaceUserClusterHostResponseBody) GoString() string {
	return s.String()
}

func (s *ReplaceUserClusterHostResponseBody) SetRequestId(v string) *ReplaceUserClusterHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplaceUserClusterHostResponseBody) SetNewHostId(v string) *ReplaceUserClusterHostResponseBody {
	s.NewHostId = &v
	return s
}

type ReplaceUserClusterHostResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReplaceUserClusterHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReplaceUserClusterHostResponse) String() string {
	return tea.Prettify(s)
}

func (s ReplaceUserClusterHostResponse) GoString() string {
	return s.String()
}

func (s *ReplaceUserClusterHostResponse) SetHeaders(v map[string]*string) *ReplaceUserClusterHostResponse {
	s.Headers = v
	return s
}

func (s *ReplaceUserClusterHostResponse) SetBody(v *ReplaceUserClusterHostResponseBody) *ReplaceUserClusterHostResponse {
	s.Body = v
	return s
}

type ResetAccountPasswordRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPassword      *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
}

func (s ResetAccountPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordRequest) SetSecurityToken(v string) *ResetAccountPasswordRequest {
	s.SecurityToken = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetOwnerId(v int64) *ResetAccountPasswordRequest {
	s.OwnerId = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetResourceOwnerAccount(v string) *ResetAccountPasswordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetResourceOwnerId(v int64) *ResetAccountPasswordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetOwnerAccount(v string) *ResetAccountPasswordRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetInstanceId(v string) *ResetAccountPasswordRequest {
	s.InstanceId = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetAccountName(v string) *ResetAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetAccountPassword(v string) *ResetAccountPasswordRequest {
	s.AccountPassword = &v
	return s
}

type ResetAccountPasswordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetAccountPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordResponseBody) SetRequestId(v string) *ResetAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

type ResetAccountPasswordResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetAccountPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordResponse) SetHeaders(v map[string]*string) *ResetAccountPasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetAccountPasswordResponse) SetBody(v *ResetAccountPasswordResponseBody) *ResetAccountPasswordResponse {
	s.Body = v
	return s
}

type RestartInstanceRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	EffectiveTime        *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	UpgradeMinorVersion  *bool   `json:"UpgradeMinorVersion,omitempty" xml:"UpgradeMinorVersion,omitempty"`
}

func (s RestartInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartInstanceRequest) SetSecurityToken(v string) *RestartInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *RestartInstanceRequest) SetOwnerId(v int64) *RestartInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RestartInstanceRequest) SetResourceOwnerAccount(v string) *RestartInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestartInstanceRequest) SetResourceOwnerId(v int64) *RestartInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestartInstanceRequest) SetOwnerAccount(v string) *RestartInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RestartInstanceRequest) SetInstanceId(v string) *RestartInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RestartInstanceRequest) SetEffectiveTime(v string) *RestartInstanceRequest {
	s.EffectiveTime = &v
	return s
}

func (s *RestartInstanceRequest) SetUpgradeMinorVersion(v bool) *RestartInstanceRequest {
	s.UpgradeMinorVersion = &v
	return s
}

type RestartInstanceResponseBody struct {
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s RestartInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponseBody) SetTaskId(v string) *RestartInstanceResponseBody {
	s.TaskId = &v
	return s
}

func (s *RestartInstanceResponseBody) SetRequestId(v string) *RestartInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartInstanceResponseBody) SetInstanceId(v string) *RestartInstanceResponseBody {
	s.InstanceId = &v
	return s
}

type RestartInstanceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RestartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponse) SetHeaders(v map[string]*string) *RestartInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestartInstanceResponse) SetBody(v *RestartInstanceResponseBody) *RestartInstanceResponse {
	s.Body = v
	return s
}

type RestoreInstanceRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	BackupId             *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	RestoreType          *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	RestoreTime          *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
}

func (s RestoreInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RestoreInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestoreInstanceRequest) SetSecurityToken(v string) *RestoreInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *RestoreInstanceRequest) SetOwnerId(v int64) *RestoreInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RestoreInstanceRequest) SetResourceOwnerAccount(v string) *RestoreInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestoreInstanceRequest) SetResourceOwnerId(v int64) *RestoreInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestoreInstanceRequest) SetOwnerAccount(v string) *RestoreInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RestoreInstanceRequest) SetInstanceId(v string) *RestoreInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RestoreInstanceRequest) SetBackupId(v string) *RestoreInstanceRequest {
	s.BackupId = &v
	return s
}

func (s *RestoreInstanceRequest) SetRestoreType(v string) *RestoreInstanceRequest {
	s.RestoreType = &v
	return s
}

func (s *RestoreInstanceRequest) SetRestoreTime(v string) *RestoreInstanceRequest {
	s.RestoreTime = &v
	return s
}

type RestoreInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestoreInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestoreInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreInstanceResponseBody) SetRequestId(v string) *RestoreInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RestoreInstanceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RestoreInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestoreInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RestoreInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestoreInstanceResponse) SetHeaders(v map[string]*string) *RestoreInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestoreInstanceResponse) SetBody(v *RestoreInstanceResponseBody) *RestoreInstanceResponse {
	s.Body = v
	return s
}

type SwitchInstanceHARequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	SwitchMode           *int32  `json:"SwitchMode,omitempty" xml:"SwitchMode,omitempty"`
	Product              *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Category             *string `json:"Category,omitempty" xml:"Category,omitempty"`
	SwitchType           *string `json:"SwitchType,omitempty" xml:"SwitchType,omitempty"`
}

func (s SwitchInstanceHARequest) String() string {
	return tea.Prettify(s)
}

func (s SwitchInstanceHARequest) GoString() string {
	return s.String()
}

func (s *SwitchInstanceHARequest) SetSecurityToken(v string) *SwitchInstanceHARequest {
	s.SecurityToken = &v
	return s
}

func (s *SwitchInstanceHARequest) SetOwnerId(v int64) *SwitchInstanceHARequest {
	s.OwnerId = &v
	return s
}

func (s *SwitchInstanceHARequest) SetResourceOwnerAccount(v string) *SwitchInstanceHARequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SwitchInstanceHARequest) SetResourceOwnerId(v int64) *SwitchInstanceHARequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SwitchInstanceHARequest) SetOwnerAccount(v string) *SwitchInstanceHARequest {
	s.OwnerAccount = &v
	return s
}

func (s *SwitchInstanceHARequest) SetInstanceId(v string) *SwitchInstanceHARequest {
	s.InstanceId = &v
	return s
}

func (s *SwitchInstanceHARequest) SetNodeId(v string) *SwitchInstanceHARequest {
	s.NodeId = &v
	return s
}

func (s *SwitchInstanceHARequest) SetSwitchMode(v int32) *SwitchInstanceHARequest {
	s.SwitchMode = &v
	return s
}

func (s *SwitchInstanceHARequest) SetProduct(v string) *SwitchInstanceHARequest {
	s.Product = &v
	return s
}

func (s *SwitchInstanceHARequest) SetCategory(v string) *SwitchInstanceHARequest {
	s.Category = &v
	return s
}

func (s *SwitchInstanceHARequest) SetSwitchType(v string) *SwitchInstanceHARequest {
	s.SwitchType = &v
	return s
}

type SwitchInstanceHAResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchInstanceHAResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SwitchInstanceHAResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchInstanceHAResponseBody) SetRequestId(v string) *SwitchInstanceHAResponseBody {
	s.RequestId = &v
	return s
}

type SwitchInstanceHAResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SwitchInstanceHAResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SwitchInstanceHAResponse) String() string {
	return tea.Prettify(s)
}

func (s SwitchInstanceHAResponse) GoString() string {
	return s.String()
}

func (s *SwitchInstanceHAResponse) SetHeaders(v map[string]*string) *SwitchInstanceHAResponse {
	s.Headers = v
	return s
}

func (s *SwitchInstanceHAResponse) SetBody(v *SwitchInstanceHAResponseBody) *SwitchInstanceHAResponse {
	s.Body = v
	return s
}

type SwitchNetworkRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	TargetNetworkType    *string `json:"TargetNetworkType,omitempty" xml:"TargetNetworkType,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RetainClassic        *string `json:"RetainClassic,omitempty" xml:"RetainClassic,omitempty"`
	ClassicExpiredDays   *string `json:"ClassicExpiredDays,omitempty" xml:"ClassicExpiredDays,omitempty"`
}

func (s SwitchNetworkRequest) String() string {
	return tea.Prettify(s)
}

func (s SwitchNetworkRequest) GoString() string {
	return s.String()
}

func (s *SwitchNetworkRequest) SetSecurityToken(v string) *SwitchNetworkRequest {
	s.SecurityToken = &v
	return s
}

func (s *SwitchNetworkRequest) SetOwnerId(v int64) *SwitchNetworkRequest {
	s.OwnerId = &v
	return s
}

func (s *SwitchNetworkRequest) SetResourceOwnerAccount(v string) *SwitchNetworkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SwitchNetworkRequest) SetResourceOwnerId(v int64) *SwitchNetworkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SwitchNetworkRequest) SetOwnerAccount(v string) *SwitchNetworkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SwitchNetworkRequest) SetTargetNetworkType(v string) *SwitchNetworkRequest {
	s.TargetNetworkType = &v
	return s
}

func (s *SwitchNetworkRequest) SetVpcId(v string) *SwitchNetworkRequest {
	s.VpcId = &v
	return s
}

func (s *SwitchNetworkRequest) SetVSwitchId(v string) *SwitchNetworkRequest {
	s.VSwitchId = &v
	return s
}

func (s *SwitchNetworkRequest) SetInstanceId(v string) *SwitchNetworkRequest {
	s.InstanceId = &v
	return s
}

func (s *SwitchNetworkRequest) SetRetainClassic(v string) *SwitchNetworkRequest {
	s.RetainClassic = &v
	return s
}

func (s *SwitchNetworkRequest) SetClassicExpiredDays(v string) *SwitchNetworkRequest {
	s.ClassicExpiredDays = &v
	return s
}

type SwitchNetworkResponseBody struct {
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchNetworkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SwitchNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchNetworkResponseBody) SetTaskId(v string) *SwitchNetworkResponseBody {
	s.TaskId = &v
	return s
}

func (s *SwitchNetworkResponseBody) SetRequestId(v string) *SwitchNetworkResponseBody {
	s.RequestId = &v
	return s
}

type SwitchNetworkResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SwitchNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SwitchNetworkResponse) String() string {
	return tea.Prettify(s)
}

func (s SwitchNetworkResponse) GoString() string {
	return s.String()
}

func (s *SwitchNetworkResponse) SetHeaders(v map[string]*string) *SwitchNetworkResponse {
	s.Headers = v
	return s
}

func (s *SwitchNetworkResponse) SetBody(v *SwitchNetworkResponseBody) *SwitchNetworkResponse {
	s.Body = v
	return s
}

type SyncDtsStatusRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SyncDtsStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncDtsStatusRequest) GoString() string {
	return s.String()
}

func (s *SyncDtsStatusRequest) SetSecurityToken(v string) *SyncDtsStatusRequest {
	s.SecurityToken = &v
	return s
}

func (s *SyncDtsStatusRequest) SetOwnerId(v int64) *SyncDtsStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *SyncDtsStatusRequest) SetResourceOwnerAccount(v string) *SyncDtsStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SyncDtsStatusRequest) SetResourceOwnerId(v int64) *SyncDtsStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SyncDtsStatusRequest) SetOwnerAccount(v string) *SyncDtsStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SyncDtsStatusRequest) SetRegionId(v string) *SyncDtsStatusRequest {
	s.RegionId = &v
	return s
}

func (s *SyncDtsStatusRequest) SetInstanceId(v string) *SyncDtsStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *SyncDtsStatusRequest) SetStatus(v string) *SyncDtsStatusRequest {
	s.Status = &v
	return s
}

func (s *SyncDtsStatusRequest) SetTaskId(v string) *SyncDtsStatusRequest {
	s.TaskId = &v
	return s
}

type SyncDtsStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SyncDtsStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncDtsStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SyncDtsStatusResponseBody) SetRequestId(v string) *SyncDtsStatusResponseBody {
	s.RequestId = &v
	return s
}

type SyncDtsStatusResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SyncDtsStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SyncDtsStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncDtsStatusResponse) GoString() string {
	return s.String()
}

func (s *SyncDtsStatusResponse) SetHeaders(v map[string]*string) *SyncDtsStatusResponse {
	s.Headers = v
	return s
}

func (s *SyncDtsStatusResponse) SetBody(v *SyncDtsStatusResponseBody) *SyncDtsStatusResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	OwnerId              *int64                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string                   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType         *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId           []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag                  []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetOwnerId(v int64) *TagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerAccount(v string) *TagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerId(v int64) *TagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetOwnerAccount(v string) *TagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type TransformToPrePaidRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Period               *int64  `json:"Period,omitempty" xml:"Period,omitempty"`
	AutoPay              *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	FromApp              *string `json:"FromApp,omitempty" xml:"FromApp,omitempty"`
}

func (s TransformToPrePaidRequest) String() string {
	return tea.Prettify(s)
}

func (s TransformToPrePaidRequest) GoString() string {
	return s.String()
}

func (s *TransformToPrePaidRequest) SetSecurityToken(v string) *TransformToPrePaidRequest {
	s.SecurityToken = &v
	return s
}

func (s *TransformToPrePaidRequest) SetOwnerId(v int64) *TransformToPrePaidRequest {
	s.OwnerId = &v
	return s
}

func (s *TransformToPrePaidRequest) SetResourceOwnerAccount(v string) *TransformToPrePaidRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TransformToPrePaidRequest) SetResourceOwnerId(v int64) *TransformToPrePaidRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TransformToPrePaidRequest) SetOwnerAccount(v string) *TransformToPrePaidRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TransformToPrePaidRequest) SetInstanceId(v string) *TransformToPrePaidRequest {
	s.InstanceId = &v
	return s
}

func (s *TransformToPrePaidRequest) SetPeriod(v int64) *TransformToPrePaidRequest {
	s.Period = &v
	return s
}

func (s *TransformToPrePaidRequest) SetAutoPay(v bool) *TransformToPrePaidRequest {
	s.AutoPay = &v
	return s
}

func (s *TransformToPrePaidRequest) SetFromApp(v string) *TransformToPrePaidRequest {
	s.FromApp = &v
	return s
}

type TransformToPrePaidResponseBody struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s TransformToPrePaidResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransformToPrePaidResponseBody) GoString() string {
	return s.String()
}

func (s *TransformToPrePaidResponseBody) SetEndTime(v string) *TransformToPrePaidResponseBody {
	s.EndTime = &v
	return s
}

func (s *TransformToPrePaidResponseBody) SetRequestId(v string) *TransformToPrePaidResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransformToPrePaidResponseBody) SetOrderId(v string) *TransformToPrePaidResponseBody {
	s.OrderId = &v
	return s
}

type TransformToPrePaidResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TransformToPrePaidResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TransformToPrePaidResponse) String() string {
	return tea.Prettify(s)
}

func (s TransformToPrePaidResponse) GoString() string {
	return s.String()
}

func (s *TransformToPrePaidResponse) SetHeaders(v map[string]*string) *TransformToPrePaidResponse {
	s.Headers = v
	return s
}

func (s *TransformToPrePaidResponse) SetBody(v *TransformToPrePaidResponseBody) *TransformToPrePaidResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType         *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	All                  *bool     `json:"All,omitempty" xml:"All,omitempty"`
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	TagKey               []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetOwnerId(v int64) *UntagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerAccount(v string) *UntagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerId(v int64) *UntagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerAccount(v string) *UntagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-qingdao":                  tea.String("r-kvstore.aliyuncs.com"),
		"cn-beijing":                  tea.String("r-kvstore.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("r-kvstore.aliyuncs.com"),
		"cn-shanghai":                 tea.String("r-kvstore.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("r-kvstore.aliyuncs.com"),
		"cn-heyuan":                   tea.String("r-kvstore.aliyuncs.com"),
		"ap-southeast-1":              tea.String("r-kvstore.aliyuncs.com"),
		"us-west-1":                   tea.String("r-kvstore.aliyuncs.com"),
		"us-east-1":                   tea.String("r-kvstore.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("r-kvstore.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("r-kvstore.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("r-kvstore.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("r-kvstore.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("r-kvstore.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("r-kvstore.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("r-kvstore.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("r-kvstore.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("r-kvstore.aliyuncs.com"),
		"cn-edge-1":                   tea.String("r-kvstore.aliyuncs.com"),
		"cn-fujian":                   tea.String("r-kvstore.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("r-kvstore.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("r-kvstore.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("r-kvstore.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("r-kvstore.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("r-kvstore.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("r-kvstore.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("r-kvstore.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("r-kvstore.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("r-kvstore.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("r-kvstore.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("r-kvstore.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("r-kvstore.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("r-kvstore.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("r-kvstore.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("r-kvstore.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("r-kvstore.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("r-kvstore.aliyuncs.com"),
		"cn-wuhan":                    tea.String("r-kvstore.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("r-kvstore.aliyuncs.com"),
		"cn-yushanfang":               tea.String("r-kvstore.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("r-kvstore.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("r-kvstore.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("r-kvstore.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("r-kvstore.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("r-kvstore.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("r-kvstore"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddShardingNodeWithOptions(request *AddShardingNodeRequest, runtime *util.RuntimeOptions) (_result *AddShardingNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddShardingNodeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddShardingNode"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddShardingNode(request *AddShardingNodeRequest) (_result *AddShardingNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddShardingNodeResponse{}
	_body, _err := client.AddShardingNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AllocateDirectConnectionWithOptions(request *AllocateDirectConnectionRequest, runtime *util.RuntimeOptions) (_result *AllocateDirectConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AllocateDirectConnectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AllocateDirectConnection"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AllocateDirectConnection(request *AllocateDirectConnectionRequest) (_result *AllocateDirectConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AllocateDirectConnectionResponse{}
	_body, _err := client.AllocateDirectConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AllocateInstancePublicConnectionWithOptions(request *AllocateInstancePublicConnectionRequest, runtime *util.RuntimeOptions) (_result *AllocateInstancePublicConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AllocateInstancePublicConnectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AllocateInstancePublicConnection"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AllocateInstancePublicConnection(request *AllocateInstancePublicConnectionRequest) (_result *AllocateInstancePublicConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AllocateInstancePublicConnectionResponse{}
	_body, _err := client.AllocateInstancePublicConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAccountWithOptions(request *CreateAccountRequest, runtime *util.RuntimeOptions) (_result *CreateAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAccount"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAccount(request *CreateAccountRequest) (_result *CreateAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAccountResponse{}
	_body, _err := client.CreateAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBackupWithOptions(request *CreateBackupRequest, runtime *util.RuntimeOptions) (_result *CreateBackupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateBackupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateBackup"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBackup(request *CreateBackupRequest) (_result *CreateBackupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBackupResponse{}
	_body, _err := client.CreateBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCacheAnalysisTaskWithOptions(request *CreateCacheAnalysisTaskRequest, runtime *util.RuntimeOptions) (_result *CreateCacheAnalysisTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCacheAnalysisTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCacheAnalysisTask"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCacheAnalysisTask(request *CreateCacheAnalysisTaskRequest) (_result *CreateCacheAnalysisTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCacheAnalysisTaskResponse{}
	_body, _err := client.CreateCacheAnalysisTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGlobalDistributeCacheWithOptions(request *CreateGlobalDistributeCacheRequest, runtime *util.RuntimeOptions) (_result *CreateGlobalDistributeCacheResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateGlobalDistributeCacheResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateGlobalDistributeCache"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGlobalDistributeCache(request *CreateGlobalDistributeCacheRequest) (_result *CreateGlobalDistributeCacheResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGlobalDistributeCacheResponse{}
	_body, _err := client.CreateGlobalDistributeCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateInstance"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTairInstanceWithOptions(request *CreateTairInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateTairInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateTairInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateTairInstance"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTairInstance(request *CreateTairInstanceRequest) (_result *CreateTairInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTairInstanceResponse{}
	_body, _err := client.CreateTairInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUserClusterHostWithOptions(request *CreateUserClusterHostRequest, runtime *util.RuntimeOptions) (_result *CreateUserClusterHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateUserClusterHostResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateUserClusterHost"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUserClusterHost(request *CreateUserClusterHostRequest) (_result *CreateUserClusterHostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUserClusterHostResponse{}
	_body, _err := client.CreateUserClusterHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAccountWithOptions(request *DeleteAccountRequest, runtime *util.RuntimeOptions) (_result *DeleteAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAccount"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAccount(request *DeleteAccountRequest) (_result *DeleteAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAccountResponse{}
	_body, _err := client.DeleteAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteInstance"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteInstance(request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteShardingNodeWithOptions(request *DeleteShardingNodeRequest, runtime *util.RuntimeOptions) (_result *DeleteShardingNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteShardingNodeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteShardingNode"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteShardingNode(request *DeleteShardingNodeRequest) (_result *DeleteShardingNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteShardingNodeResponse{}
	_body, _err := client.DeleteShardingNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserClusterHostWithOptions(request *DeleteUserClusterHostRequest, runtime *util.RuntimeOptions) (_result *DeleteUserClusterHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteUserClusterHostResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteUserClusterHost"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUserClusterHost(request *DeleteUserClusterHostRequest) (_result *DeleteUserClusterHostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserClusterHostResponse{}
	_body, _err := client.DeleteUserClusterHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAccountsWithOptions(request *DescribeAccountsRequest, runtime *util.RuntimeOptions) (_result *DescribeAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAccountsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAccounts"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccounts(request *DescribeAccountsRequest) (_result *DescribeAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccountsResponse{}
	_body, _err := client.DescribeAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeActiveOperationTaskWithOptions(request *DescribeActiveOperationTaskRequest, runtime *util.RuntimeOptions) (_result *DescribeActiveOperationTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeActiveOperationTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeActiveOperationTask"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeActiveOperationTask(request *DescribeActiveOperationTaskRequest) (_result *DescribeActiveOperationTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeActiveOperationTaskResponse{}
	_body, _err := client.DescribeActiveOperationTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAuditRecordsWithOptions(request *DescribeAuditRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeAuditRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAuditRecordsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAuditRecords"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAuditRecords(request *DescribeAuditRecordsRequest) (_result *DescribeAuditRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAuditRecordsResponse{}
	_body, _err := client.DescribeAuditRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupPolicyWithOptions(request *DescribeBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackupPolicy"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupPolicy(request *DescribeBackupPolicyRequest) (_result *DescribeBackupPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.DescribeBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupsWithOptions(request *DescribeBackupsRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackups"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackups(request *DescribeBackupsRequest) (_result *DescribeBackupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupsResponse{}
	_body, _err := client.DescribeBackupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupTasksWithOptions(request *DescribeBackupTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackupTasksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackupTasks"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupTasks(request *DescribeBackupTasksRequest) (_result *DescribeBackupTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupTasksResponse{}
	_body, _err := client.DescribeBackupTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCacheAnalysisReportWithOptions(request *DescribeCacheAnalysisReportRequest, runtime *util.RuntimeOptions) (_result *DescribeCacheAnalysisReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCacheAnalysisReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCacheAnalysisReport"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCacheAnalysisReport(request *DescribeCacheAnalysisReportRequest) (_result *DescribeCacheAnalysisReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCacheAnalysisReportResponse{}
	_body, _err := client.DescribeCacheAnalysisReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCacheAnalysisReportListWithOptions(request *DescribeCacheAnalysisReportListRequest, runtime *util.RuntimeOptions) (_result *DescribeCacheAnalysisReportListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCacheAnalysisReportListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCacheAnalysisReportList"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCacheAnalysisReportList(request *DescribeCacheAnalysisReportListRequest) (_result *DescribeCacheAnalysisReportListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCacheAnalysisReportListResponse{}
	_body, _err := client.DescribeCacheAnalysisReportListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterMemberInfoWithOptions(request *DescribeClusterMemberInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterMemberInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeClusterMemberInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeClusterMemberInfo"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterMemberInfo(request *DescribeClusterMemberInfoRequest) (_result *DescribeClusterMemberInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterMemberInfoResponse{}
	_body, _err := client.DescribeClusterMemberInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstanceNetInfoWithOptions(request *DescribeDBInstanceNetInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceNetInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBInstanceNetInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBInstanceNetInfo"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceNetInfo(request *DescribeDBInstanceNetInfoRequest) (_result *DescribeDBInstanceNetInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceNetInfoResponse{}
	_body, _err := client.DescribeDBInstanceNetInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDedicatedClusterInstanceListWithOptions(request *DescribeDedicatedClusterInstanceListRequest, runtime *util.RuntimeOptions) (_result *DescribeDedicatedClusterInstanceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDedicatedClusterInstanceListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDedicatedClusterInstanceList"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDedicatedClusterInstanceList(request *DescribeDedicatedClusterInstanceListRequest) (_result *DescribeDedicatedClusterInstanceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDedicatedClusterInstanceListResponse{}
	_body, _err := client.DescribeDedicatedClusterInstanceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEngineVersionWithOptions(request *DescribeEngineVersionRequest, runtime *util.RuntimeOptions) (_result *DescribeEngineVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEngineVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEngineVersion"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEngineVersion(request *DescribeEngineVersionRequest) (_result *DescribeEngineVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEngineVersionResponse{}
	_body, _err := client.DescribeEngineVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGlobalDistributeCacheWithOptions(request *DescribeGlobalDistributeCacheRequest, runtime *util.RuntimeOptions) (_result *DescribeGlobalDistributeCacheResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGlobalDistributeCacheResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGlobalDistributeCache"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGlobalDistributeCache(request *DescribeGlobalDistributeCacheRequest) (_result *DescribeGlobalDistributeCacheResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGlobalDistributeCacheResponse{}
	_body, _err := client.DescribeGlobalDistributeCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHistoryMonitorValuesWithOptions(request *DescribeHistoryMonitorValuesRequest, runtime *util.RuntimeOptions) (_result *DescribeHistoryMonitorValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeHistoryMonitorValuesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeHistoryMonitorValues"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHistoryMonitorValues(request *DescribeHistoryMonitorValuesRequest) (_result *DescribeHistoryMonitorValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHistoryMonitorValuesResponse{}
	_body, _err := client.DescribeHistoryMonitorValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceAttributeWithOptions(request *DescribeInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceAttribute"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceAttribute(request *DescribeInstanceAttributeRequest) (_result *DescribeInstanceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceAttributeResponse{}
	_body, _err := client.DescribeInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceAutoRenewalAttributeWithOptions(request *DescribeInstanceAutoRenewalAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceAutoRenewalAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceAutoRenewalAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceAutoRenewalAttribute"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceAutoRenewalAttribute(request *DescribeInstanceAutoRenewalAttributeRequest) (_result *DescribeInstanceAutoRenewalAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceAutoRenewalAttributeResponse{}
	_body, _err := client.DescribeInstanceAutoRenewalAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceConfigWithOptions(request *DescribeInstanceConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceConfig"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceConfig(request *DescribeInstanceConfigRequest) (_result *DescribeInstanceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceConfigResponse{}
	_body, _err := client.DescribeInstanceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstancesWithOptions(request *DescribeInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstances"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstances(request *DescribeInstancesRequest) (_result *DescribeInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DescribeInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceSSLWithOptions(request *DescribeInstanceSSLRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceSSLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceSSLResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceSSL"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceSSL(request *DescribeInstanceSSLRequest) (_result *DescribeInstanceSSLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceSSLResponse{}
	_body, _err := client.DescribeInstanceSSLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeIntranetAttributeWithOptions(request *DescribeIntranetAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeIntranetAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeIntranetAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeIntranetAttribute"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIntranetAttribute(request *DescribeIntranetAttributeRequest) (_result *DescribeIntranetAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIntranetAttributeResponse{}
	_body, _err := client.DescribeIntranetAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLogicInstanceTopologyWithOptions(request *DescribeLogicInstanceTopologyRequest, runtime *util.RuntimeOptions) (_result *DescribeLogicInstanceTopologyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeLogicInstanceTopologyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeLogicInstanceTopology"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLogicInstanceTopology(request *DescribeLogicInstanceTopologyRequest) (_result *DescribeLogicInstanceTopologyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLogicInstanceTopologyResponse{}
	_body, _err := client.DescribeLogicInstanceTopologyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMonitorItemsWithOptions(request *DescribeMonitorItemsRequest, runtime *util.RuntimeOptions) (_result *DescribeMonitorItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMonitorItemsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMonitorItems"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMonitorItems(request *DescribeMonitorItemsRequest) (_result *DescribeMonitorItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMonitorItemsResponse{}
	_body, _err := client.DescribeMonitorItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeParametersWithOptions(request *DescribeParametersRequest, runtime *util.RuntimeOptions) (_result *DescribeParametersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeParametersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeParameters"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeParameters(request *DescribeParametersRequest) (_result *DescribeParametersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeParametersResponse{}
	_body, _err := client.DescribeParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeParameterTemplatesWithOptions(request *DescribeParameterTemplatesRequest, runtime *util.RuntimeOptions) (_result *DescribeParameterTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeParameterTemplatesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeParameterTemplates"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeParameterTemplates(request *DescribeParameterTemplatesRequest) (_result *DescribeParameterTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeParameterTemplatesResponse{}
	_body, _err := client.DescribeParameterTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePriceWithOptions(request *DescribePriceRequest, runtime *util.RuntimeOptions) (_result *DescribePriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePriceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePrice"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePrice(request *DescribePriceRequest) (_result *DescribePriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePriceResponse{}
	_body, _err := client.DescribePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRoleZoneInfoWithOptions(request *DescribeRoleZoneInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeRoleZoneInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRoleZoneInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRoleZoneInfo"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRoleZoneInfo(request *DescribeRoleZoneInfoRequest) (_result *DescribeRoleZoneInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRoleZoneInfoResponse{}
	_body, _err := client.DescribeRoleZoneInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRunningLogRecordsWithOptions(request *DescribeRunningLogRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeRunningLogRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRunningLogRecordsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRunningLogRecords"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRunningLogRecords(request *DescribeRunningLogRecordsRequest) (_result *DescribeRunningLogRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRunningLogRecordsResponse{}
	_body, _err := client.DescribeRunningLogRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSecurityGroupConfigurationWithOptions(request *DescribeSecurityGroupConfigurationRequest, runtime *util.RuntimeOptions) (_result *DescribeSecurityGroupConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSecurityGroupConfigurationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSecurityGroupConfiguration"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSecurityGroupConfiguration(request *DescribeSecurityGroupConfigurationRequest) (_result *DescribeSecurityGroupConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSecurityGroupConfigurationResponse{}
	_body, _err := client.DescribeSecurityGroupConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSecurityIpsWithOptions(request *DescribeSecurityIpsRequest, runtime *util.RuntimeOptions) (_result *DescribeSecurityIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSecurityIpsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSecurityIps"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSecurityIps(request *DescribeSecurityIpsRequest) (_result *DescribeSecurityIpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSecurityIpsResponse{}
	_body, _err := client.DescribeSecurityIpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlowLogRecordsWithOptions(request *DescribeSlowLogRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeSlowLogRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSlowLogRecordsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSlowLogRecords"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlowLogRecords(request *DescribeSlowLogRecordsRequest) (_result *DescribeSlowLogRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlowLogRecordsResponse{}
	_body, _err := client.DescribeSlowLogRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTasksWithOptions(request *DescribeTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeTasksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTasks"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTasks(request *DescribeTasksRequest) (_result *DescribeTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTasksResponse{}
	_body, _err := client.DescribeTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserClusterHostWithOptions(request *DescribeUserClusterHostRequest, runtime *util.RuntimeOptions) (_result *DescribeUserClusterHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUserClusterHostResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUserClusterHost"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserClusterHost(request *DescribeUserClusterHostRequest) (_result *DescribeUserClusterHostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserClusterHostResponse{}
	_body, _err := client.DescribeUserClusterHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserClusterHostInstanceWithOptions(request *DescribeUserClusterHostInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeUserClusterHostInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUserClusterHostInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUserClusterHostInstance"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserClusterHostInstance(request *DescribeUserClusterHostInstanceRequest) (_result *DescribeUserClusterHostInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserClusterHostInstanceResponse{}
	_body, _err := client.DescribeUserClusterHostInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *util.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeZones"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeZones(request *DescribeZonesRequest) (_result *DescribeZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DescribeZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableAdditionalBandwidthWithOptions(request *EnableAdditionalBandwidthRequest, runtime *util.RuntimeOptions) (_result *EnableAdditionalBandwidthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableAdditionalBandwidthResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableAdditionalBandwidth"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableAdditionalBandwidth(request *EnableAdditionalBandwidthRequest) (_result *EnableAdditionalBandwidthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableAdditionalBandwidthResponse{}
	_body, _err := client.EnableAdditionalBandwidthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FlushExpireKeysWithOptions(request *FlushExpireKeysRequest, runtime *util.RuntimeOptions) (_result *FlushExpireKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FlushExpireKeysResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FlushExpireKeys"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FlushExpireKeys(request *FlushExpireKeysRequest) (_result *FlushExpireKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FlushExpireKeysResponse{}
	_body, _err := client.FlushExpireKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FlushInstanceWithOptions(request *FlushInstanceRequest, runtime *util.RuntimeOptions) (_result *FlushInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FlushInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FlushInstance"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FlushInstance(request *FlushInstanceRequest) (_result *FlushInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FlushInstanceResponse{}
	_body, _err := client.FlushInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GrantAccountPrivilegeWithOptions(request *GrantAccountPrivilegeRequest, runtime *util.RuntimeOptions) (_result *GrantAccountPrivilegeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GrantAccountPrivilegeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GrantAccountPrivilege"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GrantAccountPrivilege(request *GrantAccountPrivilegeRequest) (_result *GrantAccountPrivilegeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GrantAccountPrivilegeResponse{}
	_body, _err := client.GrantAccountPrivilegeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitializeKvstorePermissionWithOptions(request *InitializeKvstorePermissionRequest, runtime *util.RuntimeOptions) (_result *InitializeKvstorePermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InitializeKvstorePermissionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InitializeKvstorePermission"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitializeKvstorePermission(request *InitializeKvstorePermissionRequest) (_result *InitializeKvstorePermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitializeKvstorePermissionResponse{}
	_body, _err := client.InitializeKvstorePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagResources"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MigrateToOtherZoneWithOptions(request *MigrateToOtherZoneRequest, runtime *util.RuntimeOptions) (_result *MigrateToOtherZoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MigrateToOtherZoneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MigrateToOtherZone"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MigrateToOtherZone(request *MigrateToOtherZoneRequest) (_result *MigrateToOtherZoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MigrateToOtherZoneResponse{}
	_body, _err := client.MigrateToOtherZoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAccountDescriptionWithOptions(request *ModifyAccountDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyAccountDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyAccountDescriptionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyAccountDescription"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAccountDescription(request *ModifyAccountDescriptionRequest) (_result *ModifyAccountDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAccountDescriptionResponse{}
	_body, _err := client.ModifyAccountDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAccountPasswordWithOptions(request *ModifyAccountPasswordRequest, runtime *util.RuntimeOptions) (_result *ModifyAccountPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyAccountPasswordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyAccountPassword"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAccountPassword(request *ModifyAccountPasswordRequest) (_result *ModifyAccountPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAccountPasswordResponse{}
	_body, _err := client.ModifyAccountPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyActiveOperationTaskWithOptions(request *ModifyActiveOperationTaskRequest, runtime *util.RuntimeOptions) (_result *ModifyActiveOperationTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyActiveOperationTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyActiveOperationTask"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyActiveOperationTask(request *ModifyActiveOperationTaskRequest) (_result *ModifyActiveOperationTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyActiveOperationTaskResponse{}
	_body, _err := client.ModifyActiveOperationTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAuditLogConfigWithOptions(request *ModifyAuditLogConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyAuditLogConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyAuditLogConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyAuditLogConfig"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAuditLogConfig(request *ModifyAuditLogConfigRequest) (_result *ModifyAuditLogConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAuditLogConfigResponse{}
	_body, _err := client.ModifyAuditLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyBackupPolicyWithOptions(request *ModifyBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifyBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyBackupPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyBackupPolicy"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyBackupPolicy(request *ModifyBackupPolicyRequest) (_result *ModifyBackupPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyBackupPolicyResponse{}
	_body, _err := client.ModifyBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBInstanceConnectionStringWithOptions(request *ModifyDBInstanceConnectionStringRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceConnectionStringResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBInstanceConnectionStringResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBInstanceConnectionString"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceConnectionString(request *ModifyDBInstanceConnectionStringRequest) (_result *ModifyDBInstanceConnectionStringResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceConnectionStringResponse{}
	_body, _err := client.ModifyDBInstanceConnectionStringWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceAttributeWithOptions(request *ModifyInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstanceAttribute"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceAttribute(request *ModifyInstanceAttributeRequest) (_result *ModifyInstanceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceAttributeResponse{}
	_body, _err := client.ModifyInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceAutoRenewalAttributeWithOptions(request *ModifyInstanceAutoRenewalAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceAutoRenewalAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceAutoRenewalAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstanceAutoRenewalAttribute"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceAutoRenewalAttribute(request *ModifyInstanceAutoRenewalAttributeRequest) (_result *ModifyInstanceAutoRenewalAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceAutoRenewalAttributeResponse{}
	_body, _err := client.ModifyInstanceAutoRenewalAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceConfigWithOptions(request *ModifyInstanceConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstanceConfig"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceConfig(request *ModifyInstanceConfigRequest) (_result *ModifyInstanceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceConfigResponse{}
	_body, _err := client.ModifyInstanceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceMaintainTimeWithOptions(request *ModifyInstanceMaintainTimeRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceMaintainTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceMaintainTimeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstanceMaintainTime"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceMaintainTime(request *ModifyInstanceMaintainTimeRequest) (_result *ModifyInstanceMaintainTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceMaintainTimeResponse{}
	_body, _err := client.ModifyInstanceMaintainTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceMajorVersionWithOptions(request *ModifyInstanceMajorVersionRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceMajorVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceMajorVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstanceMajorVersion"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceMajorVersion(request *ModifyInstanceMajorVersionRequest) (_result *ModifyInstanceMajorVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceMajorVersionResponse{}
	_body, _err := client.ModifyInstanceMajorVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceMinorVersionWithOptions(request *ModifyInstanceMinorVersionRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceMinorVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceMinorVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstanceMinorVersion"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceMinorVersion(request *ModifyInstanceMinorVersionRequest) (_result *ModifyInstanceMinorVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceMinorVersionResponse{}
	_body, _err := client.ModifyInstanceMinorVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceNetExpireTimeWithOptions(request *ModifyInstanceNetExpireTimeRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceNetExpireTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceNetExpireTimeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstanceNetExpireTime"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceNetExpireTime(request *ModifyInstanceNetExpireTimeRequest) (_result *ModifyInstanceNetExpireTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceNetExpireTimeResponse{}
	_body, _err := client.ModifyInstanceNetExpireTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceSpecWithOptions(request *ModifyInstanceSpecRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceSpecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceSpecResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstanceSpec"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceSpec(request *ModifyInstanceSpecRequest) (_result *ModifyInstanceSpecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceSpecResponse{}
	_body, _err := client.ModifyInstanceSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceSSLWithOptions(request *ModifyInstanceSSLRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceSSLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceSSLResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstanceSSL"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceSSL(request *ModifyInstanceSSLRequest) (_result *ModifyInstanceSSLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceSSLResponse{}
	_body, _err := client.ModifyInstanceSSLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceVpcAuthModeWithOptions(request *ModifyInstanceVpcAuthModeRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceVpcAuthModeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceVpcAuthModeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstanceVpcAuthMode"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceVpcAuthMode(request *ModifyInstanceVpcAuthModeRequest) (_result *ModifyInstanceVpcAuthModeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceVpcAuthModeResponse{}
	_body, _err := client.ModifyInstanceVpcAuthModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyIntranetAttributeWithOptions(request *ModifyIntranetAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyIntranetAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyIntranetAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyIntranetAttribute"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyIntranetAttribute(request *ModifyIntranetAttributeRequest) (_result *ModifyIntranetAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyIntranetAttributeResponse{}
	_body, _err := client.ModifyIntranetAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyNodeSpecWithOptions(request *ModifyNodeSpecRequest, runtime *util.RuntimeOptions) (_result *ModifyNodeSpecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyNodeSpecResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyNodeSpec"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyNodeSpec(request *ModifyNodeSpecRequest) (_result *ModifyNodeSpecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyNodeSpecResponse{}
	_body, _err := client.ModifyNodeSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyResourceGroupWithOptions(request *ModifyResourceGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyResourceGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyResourceGroup"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyResourceGroup(request *ModifyResourceGroupRequest) (_result *ModifyResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyResourceGroupResponse{}
	_body, _err := client.ModifyResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySecurityGroupConfigurationWithOptions(request *ModifySecurityGroupConfigurationRequest, runtime *util.RuntimeOptions) (_result *ModifySecurityGroupConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifySecurityGroupConfigurationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifySecurityGroupConfiguration"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySecurityGroupConfiguration(request *ModifySecurityGroupConfigurationRequest) (_result *ModifySecurityGroupConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySecurityGroupConfigurationResponse{}
	_body, _err := client.ModifySecurityGroupConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySecurityIpsWithOptions(request *ModifySecurityIpsRequest, runtime *util.RuntimeOptions) (_result *ModifySecurityIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifySecurityIpsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifySecurityIps"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySecurityIps(request *ModifySecurityIpsRequest) (_result *ModifySecurityIpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySecurityIpsResponse{}
	_body, _err := client.ModifySecurityIpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyUserClusterHostWithOptions(request *ModifyUserClusterHostRequest, runtime *util.RuntimeOptions) (_result *ModifyUserClusterHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyUserClusterHostResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyUserClusterHost"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyUserClusterHost(request *ModifyUserClusterHostRequest) (_result *ModifyUserClusterHostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyUserClusterHostResponse{}
	_body, _err := client.ModifyUserClusterHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseDirectConnectionWithOptions(request *ReleaseDirectConnectionRequest, runtime *util.RuntimeOptions) (_result *ReleaseDirectConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReleaseDirectConnectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReleaseDirectConnection"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseDirectConnection(request *ReleaseDirectConnectionRequest) (_result *ReleaseDirectConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseDirectConnectionResponse{}
	_body, _err := client.ReleaseDirectConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseInstancePublicConnectionWithOptions(request *ReleaseInstancePublicConnectionRequest, runtime *util.RuntimeOptions) (_result *ReleaseInstancePublicConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReleaseInstancePublicConnectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReleaseInstancePublicConnection"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseInstancePublicConnection(request *ReleaseInstancePublicConnectionRequest) (_result *ReleaseInstancePublicConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseInstancePublicConnectionResponse{}
	_body, _err := client.ReleaseInstancePublicConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenewInstanceWithOptions(request *RenewInstanceRequest, runtime *util.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RenewInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RenewInstance"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenewInstance(request *RenewInstanceRequest) (_result *RenewInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewInstanceResponse{}
	_body, _err := client.RenewInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReplaceUserClusterHostWithOptions(request *ReplaceUserClusterHostRequest, runtime *util.RuntimeOptions) (_result *ReplaceUserClusterHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReplaceUserClusterHostResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReplaceUserClusterHost"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReplaceUserClusterHost(request *ReplaceUserClusterHostRequest) (_result *ReplaceUserClusterHostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReplaceUserClusterHostResponse{}
	_body, _err := client.ReplaceUserClusterHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetAccountPasswordWithOptions(request *ResetAccountPasswordRequest, runtime *util.RuntimeOptions) (_result *ResetAccountPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResetAccountPasswordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResetAccountPassword"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) (_result *ResetAccountPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetAccountPasswordResponse{}
	_body, _err := client.ResetAccountPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestartInstanceWithOptions(request *RestartInstanceRequest, runtime *util.RuntimeOptions) (_result *RestartInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RestartInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RestartInstance"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartInstance(request *RestartInstanceRequest) (_result *RestartInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartInstanceResponse{}
	_body, _err := client.RestartInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestoreInstanceWithOptions(request *RestoreInstanceRequest, runtime *util.RuntimeOptions) (_result *RestoreInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RestoreInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RestoreInstance"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestoreInstance(request *RestoreInstanceRequest) (_result *RestoreInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestoreInstanceResponse{}
	_body, _err := client.RestoreInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SwitchInstanceHAWithOptions(request *SwitchInstanceHARequest, runtime *util.RuntimeOptions) (_result *SwitchInstanceHAResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SwitchInstanceHAResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SwitchInstanceHA"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SwitchInstanceHA(request *SwitchInstanceHARequest) (_result *SwitchInstanceHAResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SwitchInstanceHAResponse{}
	_body, _err := client.SwitchInstanceHAWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SwitchNetworkWithOptions(request *SwitchNetworkRequest, runtime *util.RuntimeOptions) (_result *SwitchNetworkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SwitchNetworkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SwitchNetwork"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SwitchNetwork(request *SwitchNetworkRequest) (_result *SwitchNetworkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SwitchNetworkResponse{}
	_body, _err := client.SwitchNetworkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncDtsStatusWithOptions(request *SyncDtsStatusRequest, runtime *util.RuntimeOptions) (_result *SyncDtsStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SyncDtsStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SyncDtsStatus"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncDtsStatus(request *SyncDtsStatusRequest) (_result *SyncDtsStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SyncDtsStatusResponse{}
	_body, _err := client.SyncDtsStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransformToPrePaidWithOptions(request *TransformToPrePaidRequest, runtime *util.RuntimeOptions) (_result *TransformToPrePaidResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TransformToPrePaidResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TransformToPrePaid"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransformToPrePaid(request *TransformToPrePaidRequest) (_result *TransformToPrePaidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransformToPrePaidResponse{}
	_body, _err := client.TransformToPrePaidWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UntagResources"), tea.String("2015-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
