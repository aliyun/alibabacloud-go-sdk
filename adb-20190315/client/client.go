// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AllocateClusterPublicConnectionRequest struct {
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	DBClusterId            *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AllocateClusterPublicConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s AllocateClusterPublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *AllocateClusterPublicConnectionRequest) SetConnectionStringPrefix(v string) *AllocateClusterPublicConnectionRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *AllocateClusterPublicConnectionRequest) SetDBClusterId(v string) *AllocateClusterPublicConnectionRequest {
	s.DBClusterId = &v
	return s
}

func (s *AllocateClusterPublicConnectionRequest) SetOwnerAccount(v string) *AllocateClusterPublicConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AllocateClusterPublicConnectionRequest) SetOwnerId(v int64) *AllocateClusterPublicConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocateClusterPublicConnectionRequest) SetResourceOwnerAccount(v string) *AllocateClusterPublicConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocateClusterPublicConnectionRequest) SetResourceOwnerId(v int64) *AllocateClusterPublicConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

type AllocateClusterPublicConnectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateClusterPublicConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AllocateClusterPublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateClusterPublicConnectionResponseBody) SetRequestId(v string) *AllocateClusterPublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

type AllocateClusterPublicConnectionResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AllocateClusterPublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AllocateClusterPublicConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s AllocateClusterPublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *AllocateClusterPublicConnectionResponse) SetHeaders(v map[string]*string) *AllocateClusterPublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *AllocateClusterPublicConnectionResponse) SetBody(v *AllocateClusterPublicConnectionResponseBody) *AllocateClusterPublicConnectionResponse {
	s.Body = v
	return s
}

type BindDBResourcePoolWithUserRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PoolName             *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	PoolUser             *string `json:"PoolUser,omitempty" xml:"PoolUser,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s BindDBResourcePoolWithUserRequest) String() string {
	return tea.Prettify(s)
}

func (s BindDBResourcePoolWithUserRequest) GoString() string {
	return s.String()
}

func (s *BindDBResourcePoolWithUserRequest) SetDBClusterId(v string) *BindDBResourcePoolWithUserRequest {
	s.DBClusterId = &v
	return s
}

func (s *BindDBResourcePoolWithUserRequest) SetOwnerAccount(v string) *BindDBResourcePoolWithUserRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BindDBResourcePoolWithUserRequest) SetOwnerId(v int64) *BindDBResourcePoolWithUserRequest {
	s.OwnerId = &v
	return s
}

func (s *BindDBResourcePoolWithUserRequest) SetPoolName(v string) *BindDBResourcePoolWithUserRequest {
	s.PoolName = &v
	return s
}

func (s *BindDBResourcePoolWithUserRequest) SetPoolUser(v string) *BindDBResourcePoolWithUserRequest {
	s.PoolUser = &v
	return s
}

func (s *BindDBResourcePoolWithUserRequest) SetResourceOwnerAccount(v string) *BindDBResourcePoolWithUserRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindDBResourcePoolWithUserRequest) SetResourceOwnerId(v int64) *BindDBResourcePoolWithUserRequest {
	s.ResourceOwnerId = &v
	return s
}

type BindDBResourcePoolWithUserResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindDBResourcePoolWithUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindDBResourcePoolWithUserResponseBody) GoString() string {
	return s.String()
}

func (s *BindDBResourcePoolWithUserResponseBody) SetRequestId(v string) *BindDBResourcePoolWithUserResponseBody {
	s.RequestId = &v
	return s
}

type BindDBResourcePoolWithUserResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindDBResourcePoolWithUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindDBResourcePoolWithUserResponse) String() string {
	return tea.Prettify(s)
}

func (s BindDBResourcePoolWithUserResponse) GoString() string {
	return s.String()
}

func (s *BindDBResourcePoolWithUserResponse) SetHeaders(v map[string]*string) *BindDBResourcePoolWithUserResponse {
	s.Headers = v
	return s
}

func (s *BindDBResourcePoolWithUserResponse) SetBody(v *BindDBResourcePoolWithUserResponseBody) *BindDBResourcePoolWithUserResponse {
	s.Body = v
	return s
}

type CreateAccountRequest struct {
	AccountDescription   *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPassword      *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	AccountType          *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountRequest) SetAccountDescription(v string) *CreateAccountRequest {
	s.AccountDescription = &v
	return s
}

func (s *CreateAccountRequest) SetAccountName(v string) *CreateAccountRequest {
	s.AccountName = &v
	return s
}

func (s *CreateAccountRequest) SetAccountPassword(v string) *CreateAccountRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateAccountRequest) SetAccountType(v string) *CreateAccountRequest {
	s.AccountType = &v
	return s
}

func (s *CreateAccountRequest) SetDBClusterId(v string) *CreateAccountRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateAccountRequest) SetOwnerAccount(v string) *CreateAccountRequest {
	s.OwnerAccount = &v
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

type CreateAccountResponseBody struct {
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId      *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccountResponseBody) SetDBClusterId(v string) *CreateAccountResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *CreateAccountResponseBody) SetRequestId(v string) *CreateAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccountResponseBody) SetTaskId(v int32) *CreateAccountResponseBody {
	s.TaskId = &v
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

type CreateDBClusterRequest struct {
	BackupSetID          *string `json:"BackupSetID,omitempty" xml:"BackupSetID,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ComputeResource      *string `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	DBClusterCategory    *string `json:"DBClusterCategory,omitempty" xml:"DBClusterCategory,omitempty"`
	DBClusterClass       *string `json:"DBClusterClass,omitempty" xml:"DBClusterClass,omitempty"`
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	DBClusterNetworkType *string `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	DBClusterVersion     *string `json:"DBClusterVersion,omitempty" xml:"DBClusterVersion,omitempty"`
	DBNodeGroupCount     *string `json:"DBNodeGroupCount,omitempty" xml:"DBNodeGroupCount,omitempty"`
	DBNodeStorage        *string `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	ElasticIOResource    *string `json:"ElasticIOResource,omitempty" xml:"ElasticIOResource,omitempty"`
	ExecutorCount        *string `json:"ExecutorCount,omitempty" xml:"ExecutorCount,omitempty"`
	Mode                 *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PayType              *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Period               *string `json:"Period,omitempty" xml:"Period,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RestoreTime          *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	RestoreType          *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	SourceDBInstanceName *string `json:"SourceDBInstanceName,omitempty" xml:"SourceDBInstanceName,omitempty"`
	StorageResource      *string `json:"StorageResource,omitempty" xml:"StorageResource,omitempty"`
	StorageType          *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	UsedTime             *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	VPCId                *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDBClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateDBClusterRequest) SetBackupSetID(v string) *CreateDBClusterRequest {
	s.BackupSetID = &v
	return s
}

func (s *CreateDBClusterRequest) SetClientToken(v string) *CreateDBClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBClusterRequest) SetComputeResource(v string) *CreateDBClusterRequest {
	s.ComputeResource = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBClusterCategory(v string) *CreateDBClusterRequest {
	s.DBClusterCategory = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBClusterClass(v string) *CreateDBClusterRequest {
	s.DBClusterClass = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBClusterDescription(v string) *CreateDBClusterRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBClusterNetworkType(v string) *CreateDBClusterRequest {
	s.DBClusterNetworkType = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBClusterVersion(v string) *CreateDBClusterRequest {
	s.DBClusterVersion = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBNodeGroupCount(v string) *CreateDBClusterRequest {
	s.DBNodeGroupCount = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBNodeStorage(v string) *CreateDBClusterRequest {
	s.DBNodeStorage = &v
	return s
}

func (s *CreateDBClusterRequest) SetElasticIOResource(v string) *CreateDBClusterRequest {
	s.ElasticIOResource = &v
	return s
}

func (s *CreateDBClusterRequest) SetExecutorCount(v string) *CreateDBClusterRequest {
	s.ExecutorCount = &v
	return s
}

func (s *CreateDBClusterRequest) SetMode(v string) *CreateDBClusterRequest {
	s.Mode = &v
	return s
}

func (s *CreateDBClusterRequest) SetOwnerAccount(v string) *CreateDBClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBClusterRequest) SetOwnerId(v int64) *CreateDBClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBClusterRequest) SetPayType(v string) *CreateDBClusterRequest {
	s.PayType = &v
	return s
}

func (s *CreateDBClusterRequest) SetPeriod(v string) *CreateDBClusterRequest {
	s.Period = &v
	return s
}

func (s *CreateDBClusterRequest) SetRegionId(v string) *CreateDBClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBClusterRequest) SetResourceGroupId(v string) *CreateDBClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBClusterRequest) SetResourceOwnerAccount(v string) *CreateDBClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBClusterRequest) SetResourceOwnerId(v int64) *CreateDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBClusterRequest) SetRestoreTime(v string) *CreateDBClusterRequest {
	s.RestoreTime = &v
	return s
}

func (s *CreateDBClusterRequest) SetRestoreType(v string) *CreateDBClusterRequest {
	s.RestoreType = &v
	return s
}

func (s *CreateDBClusterRequest) SetSourceDBInstanceName(v string) *CreateDBClusterRequest {
	s.SourceDBInstanceName = &v
	return s
}

func (s *CreateDBClusterRequest) SetStorageResource(v string) *CreateDBClusterRequest {
	s.StorageResource = &v
	return s
}

func (s *CreateDBClusterRequest) SetStorageType(v string) *CreateDBClusterRequest {
	s.StorageType = &v
	return s
}

func (s *CreateDBClusterRequest) SetUsedTime(v string) *CreateDBClusterRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDBClusterRequest) SetVPCId(v string) *CreateDBClusterRequest {
	s.VPCId = &v
	return s
}

func (s *CreateDBClusterRequest) SetVSwitchId(v string) *CreateDBClusterRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBClusterRequest) SetZoneId(v string) *CreateDBClusterRequest {
	s.ZoneId = &v
	return s
}

type CreateDBClusterResponseBody struct {
	DBClusterId     *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OrderId         *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateDBClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBClusterResponseBody) SetDBClusterId(v string) *CreateDBClusterResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBClusterResponseBody) SetOrderId(v string) *CreateDBClusterResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDBClusterResponseBody) SetRequestId(v string) *CreateDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBClusterResponseBody) SetResourceGroupId(v string) *CreateDBClusterResponseBody {
	s.ResourceGroupId = &v
	return s
}

type CreateDBClusterResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDBClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateDBClusterResponse) SetHeaders(v map[string]*string) *CreateDBClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateDBClusterResponse) SetBody(v *CreateDBClusterResponseBody) *CreateDBClusterResponse {
	s.Body = v
	return s
}

type CreateDBResourcePoolRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	NodeNum              *int32  `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PoolName             *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	QueryType            *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateDBResourcePoolRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBResourcePoolRequest) GoString() string {
	return s.String()
}

func (s *CreateDBResourcePoolRequest) SetDBClusterId(v string) *CreateDBResourcePoolRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBResourcePoolRequest) SetNodeNum(v int32) *CreateDBResourcePoolRequest {
	s.NodeNum = &v
	return s
}

func (s *CreateDBResourcePoolRequest) SetOwnerAccount(v string) *CreateDBResourcePoolRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBResourcePoolRequest) SetOwnerId(v int64) *CreateDBResourcePoolRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBResourcePoolRequest) SetPoolName(v string) *CreateDBResourcePoolRequest {
	s.PoolName = &v
	return s
}

func (s *CreateDBResourcePoolRequest) SetQueryType(v string) *CreateDBResourcePoolRequest {
	s.QueryType = &v
	return s
}

func (s *CreateDBResourcePoolRequest) SetResourceOwnerAccount(v string) *CreateDBResourcePoolRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBResourcePoolRequest) SetResourceOwnerId(v int64) *CreateDBResourcePoolRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateDBResourcePoolResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBResourcePoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBResourcePoolResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBResourcePoolResponseBody) SetRequestId(v string) *CreateDBResourcePoolResponseBody {
	s.RequestId = &v
	return s
}

type CreateDBResourcePoolResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDBResourcePoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDBResourcePoolResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBResourcePoolResponse) GoString() string {
	return s.String()
}

func (s *CreateDBResourcePoolResponse) SetHeaders(v map[string]*string) *CreateDBResourcePoolResponse {
	s.Headers = v
	return s
}

func (s *CreateDBResourcePoolResponse) SetBody(v *CreateDBResourcePoolResponseBody) *CreateDBResourcePoolResponse {
	s.Body = v
	return s
}

type CreateElasticPlanRequest struct {
	DBClusterId             *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ElasticPlanEnable       *bool   `json:"ElasticPlanEnable,omitempty" xml:"ElasticPlanEnable,omitempty"`
	ElasticPlanEndDay       *string `json:"ElasticPlanEndDay,omitempty" xml:"ElasticPlanEndDay,omitempty"`
	ElasticPlanName         *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	ElasticPlanNodeNum      *int32  `json:"ElasticPlanNodeNum,omitempty" xml:"ElasticPlanNodeNum,omitempty"`
	ElasticPlanStartDay     *string `json:"ElasticPlanStartDay,omitempty" xml:"ElasticPlanStartDay,omitempty"`
	ElasticPlanTimeEnd      *string `json:"ElasticPlanTimeEnd,omitempty" xml:"ElasticPlanTimeEnd,omitempty"`
	ElasticPlanTimeStart    *string `json:"ElasticPlanTimeStart,omitempty" xml:"ElasticPlanTimeStart,omitempty"`
	ElasticPlanWeeklyRepeat *string `json:"ElasticPlanWeeklyRepeat,omitempty" xml:"ElasticPlanWeeklyRepeat,omitempty"`
	OwnerAccount            *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                 *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount    *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId         *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourcePoolName        *string `json:"ResourcePoolName,omitempty" xml:"ResourcePoolName,omitempty"`
}

func (s CreateElasticPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateElasticPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateElasticPlanRequest) SetDBClusterId(v string) *CreateElasticPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanEnable(v bool) *CreateElasticPlanRequest {
	s.ElasticPlanEnable = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanEndDay(v string) *CreateElasticPlanRequest {
	s.ElasticPlanEndDay = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanName(v string) *CreateElasticPlanRequest {
	s.ElasticPlanName = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanNodeNum(v int32) *CreateElasticPlanRequest {
	s.ElasticPlanNodeNum = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanStartDay(v string) *CreateElasticPlanRequest {
	s.ElasticPlanStartDay = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanTimeEnd(v string) *CreateElasticPlanRequest {
	s.ElasticPlanTimeEnd = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanTimeStart(v string) *CreateElasticPlanRequest {
	s.ElasticPlanTimeStart = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanWeeklyRepeat(v string) *CreateElasticPlanRequest {
	s.ElasticPlanWeeklyRepeat = &v
	return s
}

func (s *CreateElasticPlanRequest) SetOwnerAccount(v string) *CreateElasticPlanRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateElasticPlanRequest) SetOwnerId(v int64) *CreateElasticPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateElasticPlanRequest) SetResourceOwnerAccount(v string) *CreateElasticPlanRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateElasticPlanRequest) SetResourceOwnerId(v int64) *CreateElasticPlanRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateElasticPlanRequest) SetResourcePoolName(v string) *CreateElasticPlanRequest {
	s.ResourcePoolName = &v
	return s
}

type CreateElasticPlanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateElasticPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateElasticPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateElasticPlanResponseBody) SetRequestId(v string) *CreateElasticPlanResponseBody {
	s.RequestId = &v
	return s
}

type CreateElasticPlanResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateElasticPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateElasticPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateElasticPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateElasticPlanResponse) SetHeaders(v map[string]*string) *CreateElasticPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateElasticPlanResponse) SetBody(v *CreateElasticPlanResponseBody) *CreateElasticPlanResponse {
	s.Body = v
	return s
}

type DeleteAccountRequest struct {
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountType          *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccountRequest) SetAccountName(v string) *DeleteAccountRequest {
	s.AccountName = &v
	return s
}

func (s *DeleteAccountRequest) SetAccountType(v string) *DeleteAccountRequest {
	s.AccountType = &v
	return s
}

func (s *DeleteAccountRequest) SetDBClusterId(v string) *DeleteAccountRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteAccountRequest) SetOwnerAccount(v string) *DeleteAccountRequest {
	s.OwnerAccount = &v
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

type DeleteDBClusterRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteDBClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterRequest) SetDBClusterId(v string) *DeleteDBClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBClusterRequest) SetOwnerAccount(v string) *DeleteDBClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDBClusterRequest) SetOwnerId(v int64) *DeleteDBClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBClusterRequest) SetResourceOwnerAccount(v string) *DeleteDBClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDBClusterRequest) SetResourceOwnerId(v int64) *DeleteDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteDBClusterResponseBody struct {
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId      *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteDBClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterResponseBody) SetDBClusterId(v string) *DeleteDBClusterResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBClusterResponseBody) SetRequestId(v string) *DeleteDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBClusterResponseBody) SetTaskId(v int32) *DeleteDBClusterResponseBody {
	s.TaskId = &v
	return s
}

type DeleteDBClusterResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDBClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterResponse) SetHeaders(v map[string]*string) *DeleteDBClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBClusterResponse) SetBody(v *DeleteDBClusterResponseBody) *DeleteDBClusterResponse {
	s.Body = v
	return s
}

type DeleteDBResourcePoolRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PoolName             *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteDBResourcePoolRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBResourcePoolRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBResourcePoolRequest) SetDBClusterId(v string) *DeleteDBResourcePoolRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBResourcePoolRequest) SetOwnerAccount(v string) *DeleteDBResourcePoolRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDBResourcePoolRequest) SetOwnerId(v int64) *DeleteDBResourcePoolRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBResourcePoolRequest) SetPoolName(v string) *DeleteDBResourcePoolRequest {
	s.PoolName = &v
	return s
}

func (s *DeleteDBResourcePoolRequest) SetResourceOwnerAccount(v string) *DeleteDBResourcePoolRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDBResourcePoolRequest) SetResourceOwnerId(v int64) *DeleteDBResourcePoolRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteDBResourcePoolResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBResourcePoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBResourcePoolResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBResourcePoolResponseBody) SetRequestId(v string) *DeleteDBResourcePoolResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDBResourcePoolResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDBResourcePoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDBResourcePoolResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBResourcePoolResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBResourcePoolResponse) SetHeaders(v map[string]*string) *DeleteDBResourcePoolResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBResourcePoolResponse) SetBody(v *DeleteDBResourcePoolResponseBody) *DeleteDBResourcePoolResponse {
	s.Body = v
	return s
}

type DeleteElasticPlanRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ElasticPlanName      *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteElasticPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteElasticPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteElasticPlanRequest) SetDBClusterId(v string) *DeleteElasticPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteElasticPlanRequest) SetElasticPlanName(v string) *DeleteElasticPlanRequest {
	s.ElasticPlanName = &v
	return s
}

func (s *DeleteElasticPlanRequest) SetOwnerAccount(v string) *DeleteElasticPlanRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteElasticPlanRequest) SetOwnerId(v int64) *DeleteElasticPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteElasticPlanRequest) SetResourceOwnerAccount(v string) *DeleteElasticPlanRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteElasticPlanRequest) SetResourceOwnerId(v int64) *DeleteElasticPlanRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteElasticPlanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteElasticPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteElasticPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteElasticPlanResponseBody) SetRequestId(v string) *DeleteElasticPlanResponseBody {
	s.RequestId = &v
	return s
}

type DeleteElasticPlanResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteElasticPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteElasticPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteElasticPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteElasticPlanResponse) SetHeaders(v map[string]*string) *DeleteElasticPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteElasticPlanResponse) SetBody(v *DeleteElasticPlanResponseBody) *DeleteElasticPlanResponse {
	s.Body = v
	return s
}

type DescribeAccountsRequest struct {
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountType          *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountsRequest) SetAccountName(v string) *DescribeAccountsRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountsRequest) SetAccountType(v string) *DescribeAccountsRequest {
	s.AccountType = &v
	return s
}

func (s *DescribeAccountsRequest) SetDBClusterId(v string) *DescribeAccountsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAccountsRequest) SetOwnerAccount(v string) *DescribeAccountsRequest {
	s.OwnerAccount = &v
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

type DescribeAccountsResponseBody struct {
	AccountList *DescribeAccountsResponseBodyAccountList `json:"AccountList,omitempty" xml:"AccountList,omitempty" type:"Struct"`
	RequestId   *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBody) SetAccountList(v *DescribeAccountsResponseBodyAccountList) *DescribeAccountsResponseBody {
	s.AccountList = v
	return s
}

func (s *DescribeAccountsResponseBody) SetRequestId(v string) *DescribeAccountsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAccountsResponseBodyAccountList struct {
	DBAccount []*DescribeAccountsResponseBodyAccountListDBAccount `json:"DBAccount,omitempty" xml:"DBAccount,omitempty" type:"Repeated"`
}

func (s DescribeAccountsResponseBodyAccountList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountList) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountList) SetDBAccount(v []*DescribeAccountsResponseBodyAccountListDBAccount) *DescribeAccountsResponseBodyAccountList {
	s.DBAccount = v
	return s
}

type DescribeAccountsResponseBodyAccountListDBAccount struct {
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	AccountName        *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountStatus      *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	AccountType        *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s DescribeAccountsResponseBodyAccountListDBAccount) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountListDBAccount) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) SetAccountDescription(v string) *DescribeAccountsResponseBodyAccountListDBAccount {
	s.AccountDescription = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) SetAccountName(v string) *DescribeAccountsResponseBodyAccountListDBAccount {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) SetAccountStatus(v string) *DescribeAccountsResponseBodyAccountListDBAccount {
	s.AccountStatus = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) SetAccountType(v string) *DescribeAccountsResponseBodyAccountListDBAccount {
	s.AccountType = &v
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

type DescribeAllAccountsRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAllAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllAccountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAllAccountsRequest) SetDBClusterId(v string) *DescribeAllAccountsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAllAccountsRequest) SetOwnerAccount(v string) *DescribeAllAccountsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAllAccountsRequest) SetOwnerId(v int64) *DescribeAllAccountsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAllAccountsRequest) SetResourceOwnerAccount(v string) *DescribeAllAccountsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAllAccountsRequest) SetResourceOwnerId(v int64) *DescribeAllAccountsRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeAllAccountsResponseBody struct {
	AccountList []*DescribeAllAccountsResponseBodyAccountList `json:"AccountList,omitempty" xml:"AccountList,omitempty" type:"Repeated"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAllAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAllAccountsResponseBody) SetAccountList(v []*DescribeAllAccountsResponseBodyAccountList) *DescribeAllAccountsResponseBody {
	s.AccountList = v
	return s
}

func (s *DescribeAllAccountsResponseBody) SetRequestId(v string) *DescribeAllAccountsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAllAccountsResponseBodyAccountList struct {
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeAllAccountsResponseBodyAccountList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllAccountsResponseBodyAccountList) GoString() string {
	return s.String()
}

func (s *DescribeAllAccountsResponseBodyAccountList) SetUser(v string) *DescribeAllAccountsResponseBodyAccountList {
	s.User = &v
	return s
}

type DescribeAllAccountsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAllAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAllAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllAccountsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAllAccountsResponse) SetHeaders(v map[string]*string) *DescribeAllAccountsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAllAccountsResponse) SetBody(v *DescribeAllAccountsResponseBody) *DescribeAllAccountsResponse {
	s.Body = v
	return s
}

type DescribeAllDataSourceRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SchemaName           *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName            *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeAllDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceRequest) SetDBClusterId(v string) *DescribeAllDataSourceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAllDataSourceRequest) SetOwnerAccount(v string) *DescribeAllDataSourceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAllDataSourceRequest) SetOwnerId(v int64) *DescribeAllDataSourceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAllDataSourceRequest) SetResourceOwnerAccount(v string) *DescribeAllDataSourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAllDataSourceRequest) SetResourceOwnerId(v int64) *DescribeAllDataSourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAllDataSourceRequest) SetSchemaName(v string) *DescribeAllDataSourceRequest {
	s.SchemaName = &v
	return s
}

func (s *DescribeAllDataSourceRequest) SetTableName(v string) *DescribeAllDataSourceRequest {
	s.TableName = &v
	return s
}

type DescribeAllDataSourceResponseBody struct {
	Columns   *DescribeAllDataSourceResponseBodyColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Schemas   *DescribeAllDataSourceResponseBodySchemas `json:"Schemas,omitempty" xml:"Schemas,omitempty" type:"Struct"`
	Tables    *DescribeAllDataSourceResponseBodyTables  `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Struct"`
}

func (s DescribeAllDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBody) SetColumns(v *DescribeAllDataSourceResponseBodyColumns) *DescribeAllDataSourceResponseBody {
	s.Columns = v
	return s
}

func (s *DescribeAllDataSourceResponseBody) SetRequestId(v string) *DescribeAllDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAllDataSourceResponseBody) SetSchemas(v *DescribeAllDataSourceResponseBodySchemas) *DescribeAllDataSourceResponseBody {
	s.Schemas = v
	return s
}

func (s *DescribeAllDataSourceResponseBody) SetTables(v *DescribeAllDataSourceResponseBodyTables) *DescribeAllDataSourceResponseBody {
	s.Tables = v
	return s
}

type DescribeAllDataSourceResponseBodyColumns struct {
	Column []*DescribeAllDataSourceResponseBodyColumnsColumn `json:"Column,omitempty" xml:"Column,omitempty" type:"Repeated"`
}

func (s DescribeAllDataSourceResponseBodyColumns) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceResponseBodyColumns) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBodyColumns) SetColumn(v []*DescribeAllDataSourceResponseBodyColumnsColumn) *DescribeAllDataSourceResponseBodyColumns {
	s.Column = v
	return s
}

type DescribeAllDataSourceResponseBodyColumnsColumn struct {
	AutoIncrementColumn *bool   `json:"AutoIncrementColumn,omitempty" xml:"AutoIncrementColumn,omitempty"`
	ColumnName          *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	DBClusterId         *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	PrimaryKey          *bool   `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	SchemaName          *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName           *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAllDataSourceResponseBodyColumnsColumn) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceResponseBodyColumnsColumn) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetAutoIncrementColumn(v bool) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.AutoIncrementColumn = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetColumnName(v string) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.ColumnName = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetDBClusterId(v string) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetPrimaryKey(v bool) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.PrimaryKey = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetSchemaName(v string) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.SchemaName = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetTableName(v string) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.TableName = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetType(v string) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.Type = &v
	return s
}

type DescribeAllDataSourceResponseBodySchemas struct {
	Schema []*DescribeAllDataSourceResponseBodySchemasSchema `json:"Schema,omitempty" xml:"Schema,omitempty" type:"Repeated"`
}

func (s DescribeAllDataSourceResponseBodySchemas) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceResponseBodySchemas) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBodySchemas) SetSchema(v []*DescribeAllDataSourceResponseBodySchemasSchema) *DescribeAllDataSourceResponseBodySchemas {
	s.Schema = v
	return s
}

type DescribeAllDataSourceResponseBodySchemasSchema struct {
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	SchemaName  *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
}

func (s DescribeAllDataSourceResponseBodySchemasSchema) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceResponseBodySchemasSchema) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBodySchemasSchema) SetDBClusterId(v string) *DescribeAllDataSourceResponseBodySchemasSchema {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodySchemasSchema) SetSchemaName(v string) *DescribeAllDataSourceResponseBodySchemasSchema {
	s.SchemaName = &v
	return s
}

type DescribeAllDataSourceResponseBodyTables struct {
	Table []*DescribeAllDataSourceResponseBodyTablesTable `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeAllDataSourceResponseBodyTables) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceResponseBodyTables) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBodyTables) SetTable(v []*DescribeAllDataSourceResponseBodyTablesTable) *DescribeAllDataSourceResponseBodyTables {
	s.Table = v
	return s
}

type DescribeAllDataSourceResponseBodyTablesTable struct {
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	SchemaName  *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName   *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeAllDataSourceResponseBodyTablesTable) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceResponseBodyTablesTable) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBodyTablesTable) SetDBClusterId(v string) *DescribeAllDataSourceResponseBodyTablesTable {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyTablesTable) SetSchemaName(v string) *DescribeAllDataSourceResponseBodyTablesTable {
	s.SchemaName = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyTablesTable) SetTableName(v string) *DescribeAllDataSourceResponseBodyTablesTable {
	s.TableName = &v
	return s
}

type DescribeAllDataSourceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAllDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAllDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponse) SetHeaders(v map[string]*string) *DescribeAllDataSourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeAllDataSourceResponse) SetBody(v *DescribeAllDataSourceResponseBody) *DescribeAllDataSourceResponse {
	s.Body = v
	return s
}

type DescribeAuditLogConfigRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAuditLogConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogConfigRequest) SetDBClusterId(v string) *DescribeAuditLogConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAuditLogConfigRequest) SetOwnerAccount(v string) *DescribeAuditLogConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAuditLogConfigRequest) SetOwnerId(v int64) *DescribeAuditLogConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAuditLogConfigRequest) SetRegionId(v string) *DescribeAuditLogConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAuditLogConfigRequest) SetResourceOwnerAccount(v string) *DescribeAuditLogConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAuditLogConfigRequest) SetResourceOwnerId(v int64) *DescribeAuditLogConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeAuditLogConfigResponseBody struct {
	AuditLogStatus *string `json:"AuditLogStatus,omitempty" xml:"AuditLogStatus,omitempty"`
	DBClusterId    *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAuditLogConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogConfigResponseBody) SetAuditLogStatus(v string) *DescribeAuditLogConfigResponseBody {
	s.AuditLogStatus = &v
	return s
}

func (s *DescribeAuditLogConfigResponseBody) SetDBClusterId(v string) *DescribeAuditLogConfigResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAuditLogConfigResponseBody) SetRequestId(v string) *DescribeAuditLogConfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAuditLogConfigResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAuditLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAuditLogConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogConfigResponse) SetHeaders(v map[string]*string) *DescribeAuditLogConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuditLogConfigResponse) SetBody(v *DescribeAuditLogConfigResponseBody) *DescribeAuditLogConfigResponse {
	s.Body = v
	return s
}

type DescribeAuditLogRecordsRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	HostAddress          *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	Order                *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OrderType            *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryKeyword         *string `json:"QueryKeyword,omitempty" xml:"QueryKeyword,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SqlType              *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Succeed              *string `json:"Succeed,omitempty" xml:"Succeed,omitempty"`
	User                 *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeAuditLogRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogRecordsRequest) SetDBClusterId(v string) *DescribeAuditLogRecordsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetDBName(v string) *DescribeAuditLogRecordsRequest {
	s.DBName = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetEndTime(v string) *DescribeAuditLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetHostAddress(v string) *DescribeAuditLogRecordsRequest {
	s.HostAddress = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetOrder(v string) *DescribeAuditLogRecordsRequest {
	s.Order = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetOrderType(v string) *DescribeAuditLogRecordsRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetOwnerAccount(v string) *DescribeAuditLogRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetOwnerId(v int64) *DescribeAuditLogRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetPageNumber(v int32) *DescribeAuditLogRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetPageSize(v int32) *DescribeAuditLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetQueryKeyword(v string) *DescribeAuditLogRecordsRequest {
	s.QueryKeyword = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetRegionId(v string) *DescribeAuditLogRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetResourceOwnerAccount(v string) *DescribeAuditLogRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetResourceOwnerId(v int64) *DescribeAuditLogRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetSqlType(v string) *DescribeAuditLogRecordsRequest {
	s.SqlType = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetStartTime(v string) *DescribeAuditLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetSucceed(v string) *DescribeAuditLogRecordsRequest {
	s.Succeed = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetUser(v string) *DescribeAuditLogRecordsRequest {
	s.User = &v
	return s
}

type DescribeAuditLogRecordsResponseBody struct {
	DBClusterId *string                                     `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Items       []*DescribeAuditLogRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber  *string                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *string                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *string                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAuditLogRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogRecordsResponseBody) SetDBClusterId(v string) *DescribeAuditLogRecordsResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBody) SetItems(v []*DescribeAuditLogRecordsResponseBodyItems) *DescribeAuditLogRecordsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAuditLogRecordsResponseBody) SetPageNumber(v string) *DescribeAuditLogRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBody) SetPageSize(v string) *DescribeAuditLogRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBody) SetRequestId(v string) *DescribeAuditLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBody) SetTotalCount(v string) *DescribeAuditLogRecordsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAuditLogRecordsResponseBodyItems struct {
	ConnId      *string `json:"ConnId,omitempty" xml:"ConnId,omitempty"`
	DBName      *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	ExecuteTime *string `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	ProcessID   *string `json:"ProcessID,omitempty" xml:"ProcessID,omitempty"`
	SQLText     *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	SQLType     *string `json:"SQLType,omitempty" xml:"SQLType,omitempty"`
	Succeed     *string `json:"Succeed,omitempty" xml:"Succeed,omitempty"`
	TotalTime   *string `json:"TotalTime,omitempty" xml:"TotalTime,omitempty"`
}

func (s DescribeAuditLogRecordsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetConnId(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.ConnId = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetDBName(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.DBName = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetExecuteTime(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetHostAddress(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.HostAddress = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetProcessID(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.ProcessID = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetSQLText(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.SQLText = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetSQLType(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.SQLType = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetSucceed(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.Succeed = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetTotalTime(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.TotalTime = &v
	return s
}

type DescribeAuditLogRecordsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAuditLogRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAuditLogRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogRecordsResponse) SetHeaders(v map[string]*string) *DescribeAuditLogRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuditLogRecordsResponse) SetBody(v *DescribeAuditLogRecordsResponseBody) *DescribeAuditLogRecordsResponse {
	s.Body = v
	return s
}

type DescribeAutoRenewAttributeRequest struct {
	DBClusterIds         *string `json:"DBClusterIds,omitempty" xml:"DBClusterIds,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAutoRenewAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewAttributeRequest) SetDBClusterIds(v string) *DescribeAutoRenewAttributeRequest {
	s.DBClusterIds = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetOwnerAccount(v string) *DescribeAutoRenewAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetOwnerId(v int64) *DescribeAutoRenewAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetPageNumber(v int32) *DescribeAutoRenewAttributeRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetPageSize(v int32) *DescribeAutoRenewAttributeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetRegionId(v string) *DescribeAutoRenewAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetResourceGroupId(v string) *DescribeAutoRenewAttributeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetResourceOwnerAccount(v string) *DescribeAutoRenewAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetResourceOwnerId(v int64) *DescribeAutoRenewAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeAutoRenewAttributeResponseBody struct {
	Items            *DescribeAutoRenewAttributeResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	PageNumber       *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount  *int32                                       `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalRecordCount *int32                                       `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeAutoRenewAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoRenewAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewAttributeResponseBody) SetItems(v *DescribeAutoRenewAttributeResponseBodyItems) *DescribeAutoRenewAttributeResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBody) SetPageNumber(v int32) *DescribeAutoRenewAttributeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBody) SetPageRecordCount(v int32) *DescribeAutoRenewAttributeResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBody) SetRequestId(v string) *DescribeAutoRenewAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBody) SetTotalRecordCount(v int32) *DescribeAutoRenewAttributeResponseBody {
	s.TotalRecordCount = &v
	return s
}

type DescribeAutoRenewAttributeResponseBodyItems struct {
	AutoRenewAttribute []*DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute `json:"AutoRenewAttribute,omitempty" xml:"AutoRenewAttribute,omitempty" type:"Repeated"`
}

func (s DescribeAutoRenewAttributeResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoRenewAttributeResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewAttributeResponseBodyItems) SetAutoRenewAttribute(v []*DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) *DescribeAutoRenewAttributeResponseBodyItems {
	s.AutoRenewAttribute = v
	return s
}

type DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute struct {
	AutoRenewEnabled *bool   `json:"AutoRenewEnabled,omitempty" xml:"AutoRenewEnabled,omitempty"`
	DBClusterId      *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Duration         *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	PeriodUnit       *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RenewalStatus    *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
}

func (s DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetAutoRenewEnabled(v bool) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.AutoRenewEnabled = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetDBClusterId(v string) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetDuration(v int32) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.Duration = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetPeriodUnit(v string) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetRegionId(v string) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetRenewalStatus(v string) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.RenewalStatus = &v
	return s
}

type DescribeAutoRenewAttributeResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAutoRenewAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAutoRenewAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoRenewAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewAttributeResponse) SetHeaders(v map[string]*string) *DescribeAutoRenewAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoRenewAttributeResponse) SetBody(v *DescribeAutoRenewAttributeResponseBody) *DescribeAutoRenewAttributeResponse {
	s.Body = v
	return s
}

type DescribeAvailableResourceRequest struct {
	AcceptLanguage       *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	ChargeType           *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceRequest) SetAcceptLanguage(v string) *DescribeAvailableResourceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetChargeType(v string) *DescribeAvailableResourceRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetOwnerAccount(v string) *DescribeAvailableResourceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetOwnerId(v int64) *DescribeAvailableResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetRegionId(v string) *DescribeAvailableResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetResourceOwnerAccount(v string) *DescribeAvailableResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetResourceOwnerId(v int64) *DescribeAvailableResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetZoneId(v string) *DescribeAvailableResourceRequest {
	s.ZoneId = &v
	return s
}

type DescribeAvailableResourceResponseBody struct {
	AvailableZoneList []*DescribeAvailableResourceResponseBodyAvailableZoneList `json:"AvailableZoneList,omitempty" xml:"AvailableZoneList,omitempty" type:"Repeated"`
	RegionId          *string                                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId         *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAvailableResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBody) SetAvailableZoneList(v []*DescribeAvailableResourceResponseBodyAvailableZoneList) *DescribeAvailableResourceResponseBody {
	s.AvailableZoneList = v
	return s
}

func (s *DescribeAvailableResourceResponseBody) SetRegionId(v string) *DescribeAvailableResourceResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBody) SetRequestId(v string) *DescribeAvailableResourceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneList struct {
	SupportedMode []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode `json:"SupportedMode,omitempty" xml:"SupportedMode,omitempty" type:"Repeated"`
	ZoneId        *string                                                                `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneList) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneList) SetSupportedMode(v []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode) *DescribeAvailableResourceResponseBodyAvailableZoneList {
	s.SupportedMode = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneList) SetZoneId(v string) *DescribeAvailableResourceResponseBodyAvailableZoneList {
	s.ZoneId = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode struct {
	Mode                *string                                                                                   `json:"Mode,omitempty" xml:"Mode,omitempty"`
	SupportedSerialList []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList `json:"SupportedSerialList,omitempty" xml:"SupportedSerialList,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode) SetMode(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode {
	s.Mode = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode) SetSupportedSerialList(v []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode {
	s.SupportedSerialList = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList struct {
	Serial                     *string                                                                                                             `json:"Serial,omitempty" xml:"Serial,omitempty"`
	SupportedFlexibleResource  []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource  `json:"SupportedFlexibleResource,omitempty" xml:"SupportedFlexibleResource,omitempty" type:"Repeated"`
	SupportedInstanceClassList []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList `json:"SupportedInstanceClassList,omitempty" xml:"SupportedInstanceClassList,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList) SetSerial(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList {
	s.Serial = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList) SetSupportedFlexibleResource(v []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList {
	s.SupportedFlexibleResource = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList) SetSupportedInstanceClassList(v []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList {
	s.SupportedInstanceClassList = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource struct {
	StorageType                *string                                                                                                                                    `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	SupportedComputeResource   []*string                                                                                                                                  `json:"SupportedComputeResource,omitempty" xml:"SupportedComputeResource,omitempty" type:"Repeated"`
	SupportedElasticIOResource *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource `json:"SupportedElasticIOResource,omitempty" xml:"SupportedElasticIOResource,omitempty" type:"Struct"`
	SupportedStorageResource   []*string                                                                                                                                  `json:"SupportedStorageResource,omitempty" xml:"SupportedStorageResource,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) SetStorageType(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource {
	s.StorageType = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) SetSupportedComputeResource(v []*string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource {
	s.SupportedComputeResource = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) SetSupportedElasticIOResource(v *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource {
	s.SupportedElasticIOResource = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) SetSupportedStorageResource(v []*string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource {
	s.SupportedStorageResource = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource struct {
	MaxCount *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	MinCount *string `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	Step     *string `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource) SetMaxCount(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource {
	s.MaxCount = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource) SetMinCount(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource {
	s.MinCount = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource) SetStep(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource {
	s.Step = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList struct {
	InstanceClass          *string                                                                                                                                   `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	SupportedExecutorList  []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorList  `json:"SupportedExecutorList,omitempty" xml:"SupportedExecutorList,omitempty" type:"Repeated"`
	SupportedNodeCountList []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList `json:"SupportedNodeCountList,omitempty" xml:"SupportedNodeCountList,omitempty" type:"Repeated"`
	Tips                   *string                                                                                                                                   `json:"Tips,omitempty" xml:"Tips,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) SetInstanceClass(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList {
	s.InstanceClass = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) SetSupportedExecutorList(v []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorList) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList {
	s.SupportedExecutorList = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) SetSupportedNodeCountList(v []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList {
	s.SupportedNodeCountList = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) SetTips(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList {
	s.Tips = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorList struct {
	NodeCount *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount `json:"NodeCount,omitempty" xml:"NodeCount,omitempty" type:"Struct"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorList) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorList) SetNodeCount(v *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorList {
	s.NodeCount = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount struct {
	MaxCount *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	MinCount *string `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	Step     *string `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount) SetMaxCount(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount {
	s.MaxCount = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount) SetMinCount(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount {
	s.MinCount = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount) SetStep(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount {
	s.Step = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList struct {
	NodeCount   *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount `json:"NodeCount,omitempty" xml:"NodeCount,omitempty" type:"Struct"`
	StorageSize []*string                                                                                                                                        `json:"StorageSize,omitempty" xml:"StorageSize,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList) SetNodeCount(v *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList {
	s.NodeCount = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList) SetStorageSize(v []*string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList {
	s.StorageSize = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount struct {
	MaxCount *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	MinCount *string `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	Step     *string `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount) SetMaxCount(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount {
	s.MaxCount = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount) SetMinCount(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount {
	s.MinCount = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount) SetStep(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount {
	s.Step = &v
	return s
}

type DescribeAvailableResourceResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAvailableResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAvailableResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponse) SetHeaders(v map[string]*string) *DescribeAvailableResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableResourceResponse) SetBody(v *DescribeAvailableResourceResponseBody) *DescribeAvailableResourceResponse {
	s.Body = v
	return s
}

type DescribeBackupPolicyRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyRequest) SetDBClusterId(v string) *DescribeBackupPolicyRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetOwnerAccount(v string) *DescribeBackupPolicyRequest {
	s.OwnerAccount = &v
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

type DescribeBackupPolicyResponseBody struct {
	BackupRetentionPeriod    *int32  `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	EnableBackupLog          *string `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	LogBackupRetentionPeriod *int32  `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	PreferredBackupPeriod    *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	PreferredBackupTime      *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	RequestId                *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBody) SetBackupRetentionPeriod(v int32) *DescribeBackupPolicyResponseBody {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetEnableBackupLog(v string) *DescribeBackupPolicyResponseBody {
	s.EnableBackupLog = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetLogBackupRetentionPeriod(v int32) *DescribeBackupPolicyResponseBody {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupTime(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupTime = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetRequestId(v string) *DescribeBackupPolicyResponseBody {
	s.RequestId = &v
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
	BackupId             *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeBackupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupsRequest) SetBackupId(v string) *DescribeBackupsRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsRequest) SetDBClusterId(v string) *DescribeBackupsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeBackupsRequest) SetEndTime(v string) *DescribeBackupsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupsRequest) SetOwnerAccount(v string) *DescribeBackupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBackupsRequest) SetOwnerId(v int64) *DescribeBackupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupsRequest) SetPageNumber(v int32) *DescribeBackupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupsRequest) SetPageSize(v int32) *DescribeBackupsRequest {
	s.PageSize = &v
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

func (s *DescribeBackupsRequest) SetStartTime(v string) *DescribeBackupsRequest {
	s.StartTime = &v
	return s
}

type DescribeBackupsResponseBody struct {
	Items      *DescribeBackupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	PageNumber *string                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *string                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBackupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBody) SetItems(v *DescribeBackupsResponseBodyItems) *DescribeBackupsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeBackupsResponseBody) SetPageNumber(v string) *DescribeBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetPageSize(v string) *DescribeBackupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetRequestId(v string) *DescribeBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetTotalCount(v string) *DescribeBackupsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeBackupsResponseBodyItems struct {
	Backup []*DescribeBackupsResponseBodyItemsBackup `json:"Backup,omitempty" xml:"Backup,omitempty" type:"Repeated"`
}

func (s DescribeBackupsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyItems) SetBackup(v []*DescribeBackupsResponseBodyItemsBackup) *DescribeBackupsResponseBodyItems {
	s.Backup = v
	return s
}

type DescribeBackupsResponseBodyItemsBackup struct {
	BackupEndTime   *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	BackupId        *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	BackupMethod    *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	BackupSize      *int32  `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	BackupType      *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	DBClusterId     *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeBackupsResponseBodyItemsBackup) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseBodyItemsBackup) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupEndTime(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupId(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupMethod(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupMethod = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupSize(v int32) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupSize = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupStartTime(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupType(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetDBClusterId(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.DBClusterId = &v
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

type DescribeColumnsRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SchemaName           *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName            *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeColumnsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeColumnsRequest) GoString() string {
	return s.String()
}

func (s *DescribeColumnsRequest) SetDBClusterId(v string) *DescribeColumnsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeColumnsRequest) SetOwnerAccount(v string) *DescribeColumnsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeColumnsRequest) SetOwnerId(v int64) *DescribeColumnsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeColumnsRequest) SetResourceOwnerAccount(v string) *DescribeColumnsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeColumnsRequest) SetResourceOwnerId(v int64) *DescribeColumnsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeColumnsRequest) SetSchemaName(v string) *DescribeColumnsRequest {
	s.SchemaName = &v
	return s
}

func (s *DescribeColumnsRequest) SetTableName(v string) *DescribeColumnsRequest {
	s.TableName = &v
	return s
}

type DescribeColumnsResponseBody struct {
	Items     *DescribeColumnsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeColumnsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeColumnsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeColumnsResponseBody) SetItems(v *DescribeColumnsResponseBodyItems) *DescribeColumnsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeColumnsResponseBody) SetRequestId(v string) *DescribeColumnsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeColumnsResponseBodyItems struct {
	Column []*DescribeColumnsResponseBodyItemsColumn `json:"Column,omitempty" xml:"Column,omitempty" type:"Repeated"`
}

func (s DescribeColumnsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeColumnsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeColumnsResponseBodyItems) SetColumn(v []*DescribeColumnsResponseBodyItemsColumn) *DescribeColumnsResponseBodyItems {
	s.Column = v
	return s
}

type DescribeColumnsResponseBodyItemsColumn struct {
	AutoIncrementColumn *bool   `json:"AutoIncrementColumn,omitempty" xml:"AutoIncrementColumn,omitempty"`
	ColumnName          *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	DBClusterId         *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	PrimaryKey          *bool   `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	SchemaName          *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName           *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeColumnsResponseBodyItemsColumn) String() string {
	return tea.Prettify(s)
}

func (s DescribeColumnsResponseBodyItemsColumn) GoString() string {
	return s.String()
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetAutoIncrementColumn(v bool) *DescribeColumnsResponseBodyItemsColumn {
	s.AutoIncrementColumn = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetColumnName(v string) *DescribeColumnsResponseBodyItemsColumn {
	s.ColumnName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetDBClusterId(v string) *DescribeColumnsResponseBodyItemsColumn {
	s.DBClusterId = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetPrimaryKey(v bool) *DescribeColumnsResponseBodyItemsColumn {
	s.PrimaryKey = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetSchemaName(v string) *DescribeColumnsResponseBodyItemsColumn {
	s.SchemaName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetTableName(v string) *DescribeColumnsResponseBodyItemsColumn {
	s.TableName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetType(v string) *DescribeColumnsResponseBodyItemsColumn {
	s.Type = &v
	return s
}

type DescribeColumnsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeColumnsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeColumnsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeColumnsResponse) GoString() string {
	return s.String()
}

func (s *DescribeColumnsResponse) SetHeaders(v map[string]*string) *DescribeColumnsResponse {
	s.Headers = v
	return s
}

func (s *DescribeColumnsResponse) SetBody(v *DescribeColumnsResponseBody) *DescribeColumnsResponse {
	s.Body = v
	return s
}

type DescribeConnectionCountRecordsRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeConnectionCountRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectionCountRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeConnectionCountRecordsRequest) SetDBClusterId(v string) *DescribeConnectionCountRecordsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeConnectionCountRecordsRequest) SetOwnerAccount(v string) *DescribeConnectionCountRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeConnectionCountRecordsRequest) SetOwnerId(v int64) *DescribeConnectionCountRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeConnectionCountRecordsRequest) SetResourceOwnerAccount(v string) *DescribeConnectionCountRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeConnectionCountRecordsRequest) SetResourceOwnerId(v int64) *DescribeConnectionCountRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeConnectionCountRecordsResponseBody struct {
	AccessIpRecords []*DescribeConnectionCountRecordsResponseBodyAccessIpRecords `json:"AccessIpRecords,omitempty" xml:"AccessIpRecords,omitempty" type:"Repeated"`
	DBClusterId     *string                                                      `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	RequestId       *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserRecords     []*DescribeConnectionCountRecordsResponseBodyUserRecords     `json:"UserRecords,omitempty" xml:"UserRecords,omitempty" type:"Repeated"`
}

func (s DescribeConnectionCountRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectionCountRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConnectionCountRecordsResponseBody) SetAccessIpRecords(v []*DescribeConnectionCountRecordsResponseBodyAccessIpRecords) *DescribeConnectionCountRecordsResponseBody {
	s.AccessIpRecords = v
	return s
}

func (s *DescribeConnectionCountRecordsResponseBody) SetDBClusterId(v string) *DescribeConnectionCountRecordsResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeConnectionCountRecordsResponseBody) SetRequestId(v string) *DescribeConnectionCountRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConnectionCountRecordsResponseBody) SetUserRecords(v []*DescribeConnectionCountRecordsResponseBodyUserRecords) *DescribeConnectionCountRecordsResponseBody {
	s.UserRecords = v
	return s
}

type DescribeConnectionCountRecordsResponseBodyAccessIpRecords struct {
	AccessIp *string `json:"AccessIp,omitempty" xml:"AccessIp,omitempty"`
	Count    *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeConnectionCountRecordsResponseBodyAccessIpRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectionCountRecordsResponseBodyAccessIpRecords) GoString() string {
	return s.String()
}

func (s *DescribeConnectionCountRecordsResponseBodyAccessIpRecords) SetAccessIp(v string) *DescribeConnectionCountRecordsResponseBodyAccessIpRecords {
	s.AccessIp = &v
	return s
}

func (s *DescribeConnectionCountRecordsResponseBodyAccessIpRecords) SetCount(v int64) *DescribeConnectionCountRecordsResponseBodyAccessIpRecords {
	s.Count = &v
	return s
}

type DescribeConnectionCountRecordsResponseBodyUserRecords struct {
	Count *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	User  *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeConnectionCountRecordsResponseBodyUserRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectionCountRecordsResponseBodyUserRecords) GoString() string {
	return s.String()
}

func (s *DescribeConnectionCountRecordsResponseBodyUserRecords) SetCount(v int64) *DescribeConnectionCountRecordsResponseBodyUserRecords {
	s.Count = &v
	return s
}

func (s *DescribeConnectionCountRecordsResponseBodyUserRecords) SetUser(v string) *DescribeConnectionCountRecordsResponseBodyUserRecords {
	s.User = &v
	return s
}

type DescribeConnectionCountRecordsResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeConnectionCountRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeConnectionCountRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectionCountRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeConnectionCountRecordsResponse) SetHeaders(v map[string]*string) *DescribeConnectionCountRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeConnectionCountRecordsResponse) SetBody(v *DescribeConnectionCountRecordsResponseBody) *DescribeConnectionCountRecordsResponse {
	s.Body = v
	return s
}

type DescribeDBClusterAccessWhiteListRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterAccessWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAccessWhiteListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhiteListRequest) SetDBClusterId(v string) *DescribeDBClusterAccessWhiteListRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListRequest) SetOwnerAccount(v string) *DescribeDBClusterAccessWhiteListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListRequest) SetOwnerId(v int64) *DescribeDBClusterAccessWhiteListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterAccessWhiteListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListRequest) SetResourceOwnerId(v int64) *DescribeDBClusterAccessWhiteListRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeDBClusterAccessWhiteListResponseBody struct {
	Items     *DescribeDBClusterAccessWhiteListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterAccessWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAccessWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhiteListResponseBody) SetItems(v *DescribeDBClusterAccessWhiteListResponseBodyItems) *DescribeDBClusterAccessWhiteListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBClusterAccessWhiteListResponseBody) SetRequestId(v string) *DescribeDBClusterAccessWhiteListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBClusterAccessWhiteListResponseBodyItems struct {
	IPArray []*DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray `json:"IPArray,omitempty" xml:"IPArray,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterAccessWhiteListResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAccessWhiteListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyItems) SetIPArray(v []*DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray) *DescribeDBClusterAccessWhiteListResponseBodyItems {
	s.IPArray = v
	return s
}

type DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray struct {
	DBClusterIPArrayAttribute *string `json:"DBClusterIPArrayAttribute,omitempty" xml:"DBClusterIPArrayAttribute,omitempty"`
	DBClusterIPArrayName      *string `json:"DBClusterIPArrayName,omitempty" xml:"DBClusterIPArrayName,omitempty"`
	SecurityIPList            *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray) SetDBClusterIPArrayAttribute(v string) *DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray {
	s.DBClusterIPArrayAttribute = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray) SetDBClusterIPArrayName(v string) *DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray {
	s.DBClusterIPArrayName = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray) SetSecurityIPList(v string) *DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray {
	s.SecurityIPList = &v
	return s
}

type DescribeDBClusterAccessWhiteListResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBClusterAccessWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterAccessWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAccessWhiteListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhiteListResponse) SetHeaders(v map[string]*string) *DescribeDBClusterAccessWhiteListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterAccessWhiteListResponse) SetBody(v *DescribeDBClusterAccessWhiteListResponseBody) *DescribeDBClusterAccessWhiteListResponse {
	s.Body = v
	return s
}

type DescribeDBClusterAttributeRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeRequest) SetDBClusterId(v string) *DescribeDBClusterAttributeRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterAttributeRequest) SetOwnerAccount(v string) *DescribeDBClusterAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAttributeRequest) SetOwnerId(v int64) *DescribeDBClusterAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterAttributeRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAttributeRequest) SetResourceOwnerId(v int64) *DescribeDBClusterAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeDBClusterAttributeResponseBody struct {
	Items     *DescribeDBClusterAttributeResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBody) SetItems(v *DescribeDBClusterAttributeResponseBodyItems) *DescribeDBClusterAttributeResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetRequestId(v string) *DescribeDBClusterAttributeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBClusterAttributeResponseBodyItems struct {
	DBCluster []*DescribeDBClusterAttributeResponseBodyItemsDBCluster `json:"DBCluster,omitempty" xml:"DBCluster,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterAttributeResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyItems) SetDBCluster(v []*DescribeDBClusterAttributeResponseBodyItemsDBCluster) *DescribeDBClusterAttributeResponseBodyItems {
	s.DBCluster = v
	return s
}

type DescribeDBClusterAttributeResponseBodyItemsDBCluster struct {
	Category             *string                                                   `json:"Category,omitempty" xml:"Category,omitempty"`
	CommodityCode        *string                                                   `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ComputeResource      *string                                                   `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	ConnectionString     *string                                                   `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	CreationTime         *string                                                   `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DBClusterDescription *string                                                   `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	DBClusterId          *string                                                   `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBClusterNetworkType *string                                                   `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	DBClusterStatus      *string                                                   `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	DBClusterType        *string                                                   `json:"DBClusterType,omitempty" xml:"DBClusterType,omitempty"`
	DBNodeClass          *string                                                   `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	DBNodeCount          *int64                                                    `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	DBNodeStorage        *int64                                                    `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	DBVersion            *string                                                   `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	DiskType             *string                                                   `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	DtsJobId             *string                                                   `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	ElasticIOResource    *int32                                                    `json:"ElasticIOResource,omitempty" xml:"ElasticIOResource,omitempty"`
	EnableAirflow        *bool                                                     `json:"EnableAirflow,omitempty" xml:"EnableAirflow,omitempty"`
	EnableSpark          *bool                                                     `json:"EnableSpark,omitempty" xml:"EnableSpark,omitempty"`
	Engine               *string                                                   `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion        *string                                                   `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ExecutorCount        *string                                                   `json:"ExecutorCount,omitempty" xml:"ExecutorCount,omitempty"`
	ExpireTime           *string                                                   `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Expired              *string                                                   `json:"Expired,omitempty" xml:"Expired,omitempty"`
	KmsId                *string                                                   `json:"KmsId,omitempty" xml:"KmsId,omitempty"`
	LockMode             *string                                                   `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	LockReason           *string                                                   `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	MaintainTime         *string                                                   `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
	Mode                 *string                                                   `json:"Mode,omitempty" xml:"Mode,omitempty"`
	PayType              *string                                                   `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Port                 *int32                                                    `json:"Port,omitempty" xml:"Port,omitempty"`
	RdsInstanceId        *string                                                   `json:"RdsInstanceId,omitempty" xml:"RdsInstanceId,omitempty"`
	RegionId             *string                                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string                                                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StorageResource      *string                                                   `json:"StorageResource,omitempty" xml:"StorageResource,omitempty"`
	Tags                 *DescribeDBClusterAttributeResponseBodyItemsDBClusterTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	UserENIStatus        *bool                                                     `json:"UserENIStatus,omitempty" xml:"UserENIStatus,omitempty"`
	VPCCloudInstanceId   *string                                                   `json:"VPCCloudInstanceId,omitempty" xml:"VPCCloudInstanceId,omitempty"`
	VPCId                *string                                                   `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId            *string                                                   `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId               *string                                                   `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBCluster) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBCluster) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetCategory(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Category = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetCommodityCode(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetComputeResource(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ComputeResource = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetConnectionString(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetCreationTime(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.CreationTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBClusterDescription(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBClusterId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBClusterNetworkType(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBClusterNetworkType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBClusterStatus(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBClusterType(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBClusterType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBNodeClass(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBNodeCount(v int64) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBNodeCount = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBNodeStorage(v int64) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBNodeStorage = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBVersion(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDiskType(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DiskType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDtsJobId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetElasticIOResource(v int32) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ElasticIOResource = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetEnableAirflow(v bool) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.EnableAirflow = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetEnableSpark(v bool) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.EnableSpark = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetEngine(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Engine = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetEngineVersion(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetExecutorCount(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ExecutorCount = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetExpireTime(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetExpired(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Expired = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetKmsId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.KmsId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetLockMode(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.LockMode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetLockReason(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.LockReason = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetMaintainTime(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.MaintainTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetMode(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Mode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetPayType(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.PayType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetPort(v int32) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Port = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetRdsInstanceId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.RdsInstanceId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetRegionId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetResourceGroupId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetStorageResource(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.StorageResource = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetTags(v *DescribeDBClusterAttributeResponseBodyItemsDBClusterTags) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Tags = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetUserENIStatus(v bool) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.UserENIStatus = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetVPCCloudInstanceId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.VPCCloudInstanceId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetVPCId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.VPCId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetVSwitchId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetZoneId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ZoneId = &v
	return s
}

type DescribeDBClusterAttributeResponseBodyItemsDBClusterTags struct {
	Tag []*DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBClusterTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBClusterTags) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTags) SetTag(v []*DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag) *DescribeDBClusterAttributeResponseBodyItemsDBClusterTags {
	s.Tag = v
	return s
}

type DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag) SetKey(v string) *DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag) SetValue(v string) *DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag {
	s.Value = &v
	return s
}

type DescribeDBClusterAttributeResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBClusterAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponse) SetHeaders(v map[string]*string) *DescribeDBClusterAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetBody(v *DescribeDBClusterAttributeResponseBody) *DescribeDBClusterAttributeResponse {
	s.Body = v
	return s
}

type DescribeDBClusterForecastRequest struct {
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	MetricType  *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBClusterForecastRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterForecastRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterForecastRequest) SetDBClusterId(v string) *DescribeDBClusterForecastRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterForecastRequest) SetMetricType(v string) *DescribeDBClusterForecastRequest {
	s.MetricType = &v
	return s
}

func (s *DescribeDBClusterForecastRequest) SetRegionId(v string) *DescribeDBClusterForecastRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterForecastRequest) SetStartTime(v string) *DescribeDBClusterForecastRequest {
	s.StartTime = &v
	return s
}

type DescribeDBClusterForecastResponseBody struct {
	Performances []*DescribeDBClusterForecastResponseBodyPerformances `json:"Performances,omitempty" xml:"Performances,omitempty" type:"Repeated"`
	RequestId    *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterForecastResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterForecastResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterForecastResponseBody) SetPerformances(v []*DescribeDBClusterForecastResponseBodyPerformances) *DescribeDBClusterForecastResponseBody {
	s.Performances = v
	return s
}

func (s *DescribeDBClusterForecastResponseBody) SetRequestId(v string) *DescribeDBClusterForecastResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBClusterForecastResponseBodyPerformances struct {
	Key    *string                                                    `json:"Key,omitempty" xml:"Key,omitempty"`
	Series []*DescribeDBClusterForecastResponseBodyPerformancesSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
	Unit   *string                                                    `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeDBClusterForecastResponseBodyPerformances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterForecastResponseBodyPerformances) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterForecastResponseBodyPerformances) SetKey(v string) *DescribeDBClusterForecastResponseBodyPerformances {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterForecastResponseBodyPerformances) SetSeries(v []*DescribeDBClusterForecastResponseBodyPerformancesSeries) *DescribeDBClusterForecastResponseBodyPerformances {
	s.Series = v
	return s
}

func (s *DescribeDBClusterForecastResponseBodyPerformances) SetUnit(v string) *DescribeDBClusterForecastResponseBodyPerformances {
	s.Unit = &v
	return s
}

type DescribeDBClusterForecastResponseBodyPerformancesSeries struct {
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DescribeDBClusterForecastResponseBodyPerformancesSeries) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterForecastResponseBodyPerformancesSeries) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterForecastResponseBodyPerformancesSeries) SetName(v string) *DescribeDBClusterForecastResponseBodyPerformancesSeries {
	s.Name = &v
	return s
}

func (s *DescribeDBClusterForecastResponseBodyPerformancesSeries) SetValues(v string) *DescribeDBClusterForecastResponseBodyPerformancesSeries {
	s.Values = &v
	return s
}

type DescribeDBClusterForecastResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBClusterForecastResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterForecastResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterForecastResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterForecastResponse) SetHeaders(v map[string]*string) *DescribeDBClusterForecastResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterForecastResponse) SetBody(v *DescribeDBClusterForecastResponseBody) *DescribeDBClusterForecastResponse {
	s.Body = v
	return s
}

type DescribeDBClusterHealthReportRequest struct {
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBClusterHealthReportRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterHealthReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterHealthReportRequest) SetDBClusterId(v string) *DescribeDBClusterHealthReportRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterHealthReportRequest) SetRegionId(v string) *DescribeDBClusterHealthReportRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterHealthReportRequest) SetStartTime(v string) *DescribeDBClusterHealthReportRequest {
	s.StartTime = &v
	return s
}

type DescribeDBClusterHealthReportResponseBody struct {
	Items     []*DescribeDBClusterHealthReportResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterHealthReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterHealthReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterHealthReportResponseBody) SetItems(v []*DescribeDBClusterHealthReportResponseBodyItems) *DescribeDBClusterHealthReportResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBClusterHealthReportResponseBody) SetRequestId(v string) *DescribeDBClusterHealthReportResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBClusterHealthReportResponseBodyItems struct {
	Avg  *string `json:"Avg,omitempty" xml:"Avg,omitempty"`
	Key  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Max  *string `json:"Max,omitempty" xml:"Max,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDBClusterHealthReportResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterHealthReportResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterHealthReportResponseBodyItems) SetAvg(v string) *DescribeDBClusterHealthReportResponseBodyItems {
	s.Avg = &v
	return s
}

func (s *DescribeDBClusterHealthReportResponseBodyItems) SetKey(v string) *DescribeDBClusterHealthReportResponseBodyItems {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterHealthReportResponseBodyItems) SetMax(v string) *DescribeDBClusterHealthReportResponseBodyItems {
	s.Max = &v
	return s
}

func (s *DescribeDBClusterHealthReportResponseBodyItems) SetName(v string) *DescribeDBClusterHealthReportResponseBodyItems {
	s.Name = &v
	return s
}

type DescribeDBClusterHealthReportResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBClusterHealthReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterHealthReportResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterHealthReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterHealthReportResponse) SetHeaders(v map[string]*string) *DescribeDBClusterHealthReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterHealthReportResponse) SetBody(v *DescribeDBClusterHealthReportResponseBody) *DescribeDBClusterHealthReportResponse {
	s.Body = v
	return s
}

type DescribeDBClusterNetInfoRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterNetInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterNetInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoRequest) SetDBClusterId(v string) *DescribeDBClusterNetInfoRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterNetInfoRequest) SetOwnerAccount(v string) *DescribeDBClusterNetInfoRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterNetInfoRequest) SetOwnerId(v int64) *DescribeDBClusterNetInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterNetInfoRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterNetInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterNetInfoRequest) SetResourceOwnerId(v int64) *DescribeDBClusterNetInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeDBClusterNetInfoResponseBody struct {
	ClusterNetworkType *string                                    `json:"ClusterNetworkType,omitempty" xml:"ClusterNetworkType,omitempty"`
	Items              *DescribeDBClusterNetInfoResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	RequestId          *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterNetInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterNetInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoResponseBody) SetClusterNetworkType(v string) *DescribeDBClusterNetInfoResponseBody {
	s.ClusterNetworkType = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBody) SetItems(v *DescribeDBClusterNetInfoResponseBodyItems) *DescribeDBClusterNetInfoResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBody) SetRequestId(v string) *DescribeDBClusterNetInfoResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBClusterNetInfoResponseBodyItems struct {
	Address []*DescribeDBClusterNetInfoResponseBodyItemsAddress `json:"Address,omitempty" xml:"Address,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterNetInfoResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterNetInfoResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoResponseBodyItems) SetAddress(v []*DescribeDBClusterNetInfoResponseBodyItemsAddress) *DescribeDBClusterNetInfoResponseBodyItems {
	s.Address = v
	return s
}

type DescribeDBClusterNetInfoResponseBodyItemsAddress struct {
	ConnectionString       *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	IPAddress              *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	NetType                *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	Port                   *string `json:"Port,omitempty" xml:"Port,omitempty"`
	VPCId                  *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId              *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeDBClusterNetInfoResponseBodyItemsAddress) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterNetInfoResponseBodyItemsAddress) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) SetConnectionString(v string) *DescribeDBClusterNetInfoResponseBodyItemsAddress {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) SetConnectionStringPrefix(v string) *DescribeDBClusterNetInfoResponseBodyItemsAddress {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) SetIPAddress(v string) *DescribeDBClusterNetInfoResponseBodyItemsAddress {
	s.IPAddress = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) SetNetType(v string) *DescribeDBClusterNetInfoResponseBodyItemsAddress {
	s.NetType = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) SetPort(v string) *DescribeDBClusterNetInfoResponseBodyItemsAddress {
	s.Port = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) SetVPCId(v string) *DescribeDBClusterNetInfoResponseBodyItemsAddress {
	s.VPCId = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) SetVSwitchId(v string) *DescribeDBClusterNetInfoResponseBodyItemsAddress {
	s.VSwitchId = &v
	return s
}

type DescribeDBClusterNetInfoResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBClusterNetInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterNetInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterNetInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoResponse) SetHeaders(v map[string]*string) *DescribeDBClusterNetInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterNetInfoResponse) SetBody(v *DescribeDBClusterNetInfoResponseBody) *DescribeDBClusterNetInfoResponse {
	s.Body = v
	return s
}

type DescribeDBClusterPerformanceRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Key                  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBClusterPerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceRequest) SetDBClusterId(v string) *DescribeDBClusterPerformanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetEndTime(v string) *DescribeDBClusterPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetKey(v string) *DescribeDBClusterPerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetOwnerAccount(v string) *DescribeDBClusterPerformanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetOwnerId(v int64) *DescribeDBClusterPerformanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterPerformanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetResourceOwnerId(v int64) *DescribeDBClusterPerformanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetStartTime(v string) *DescribeDBClusterPerformanceRequest {
	s.StartTime = &v
	return s
}

type DescribeDBClusterPerformanceResponseBody struct {
	DBClusterId  *string                                                 `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	EndTime      *string                                                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Performances []*DescribeDBClusterPerformanceResponseBodyPerformances `json:"Performances,omitempty" xml:"Performances,omitempty" type:"Repeated"`
	RequestId    *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime    *string                                                 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBClusterPerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBody) SetDBClusterId(v string) *DescribeDBClusterPerformanceResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetEndTime(v string) *DescribeDBClusterPerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetPerformances(v []*DescribeDBClusterPerformanceResponseBodyPerformances) *DescribeDBClusterPerformanceResponseBody {
	s.Performances = v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetRequestId(v string) *DescribeDBClusterPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetStartTime(v string) *DescribeDBClusterPerformanceResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDBClusterPerformanceResponseBodyPerformances struct {
	Key    *string                                                       `json:"Key,omitempty" xml:"Key,omitempty"`
	Series []*DescribeDBClusterPerformanceResponseBodyPerformancesSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
	Unit   *string                                                       `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeDBClusterPerformanceResponseBodyPerformances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBodyPerformances) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformances) SetKey(v string) *DescribeDBClusterPerformanceResponseBodyPerformances {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformances) SetSeries(v []*DescribeDBClusterPerformanceResponseBodyPerformancesSeries) *DescribeDBClusterPerformanceResponseBodyPerformances {
	s.Series = v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformances) SetUnit(v string) *DescribeDBClusterPerformanceResponseBodyPerformances {
	s.Unit = &v
	return s
}

type DescribeDBClusterPerformanceResponseBodyPerformancesSeries struct {
	Name   *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterPerformanceResponseBodyPerformancesSeries) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBodyPerformancesSeries) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformancesSeries) SetName(v string) *DescribeDBClusterPerformanceResponseBodyPerformancesSeries {
	s.Name = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformancesSeries) SetValues(v []*string) *DescribeDBClusterPerformanceResponseBodyPerformancesSeries {
	s.Values = v
	return s
}

type DescribeDBClusterPerformanceResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBClusterPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterPerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponse) SetHeaders(v map[string]*string) *DescribeDBClusterPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterPerformanceResponse) SetBody(v *DescribeDBClusterPerformanceResponseBody) *DescribeDBClusterPerformanceResponse {
	s.Body = v
	return s
}

type DescribeDBClusterResourcePoolPerformanceRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Key                  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourcePools        *string `json:"ResourcePools,omitempty" xml:"ResourcePools,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBClusterResourcePoolPerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterResourcePoolPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetDBClusterId(v string) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetEndTime(v string) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetKey(v string) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetOwnerAccount(v string) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetOwnerId(v int64) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetResourceOwnerId(v int64) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetResourcePools(v string) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.ResourcePools = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetStartTime(v string) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.StartTime = &v
	return s
}

type DescribeDBClusterResourcePoolPerformanceResponseBody struct {
	DBClusterId  *string                                                             `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	EndTime      *string                                                             `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Performances []*DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances `json:"Performances,omitempty" xml:"Performances,omitempty" type:"Repeated"`
	RequestId    *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime    *string                                                             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBClusterResourcePoolPerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterResourcePoolPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBody) SetDBClusterId(v string) *DescribeDBClusterResourcePoolPerformanceResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBody) SetEndTime(v string) *DescribeDBClusterResourcePoolPerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBody) SetPerformances(v []*DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances) *DescribeDBClusterResourcePoolPerformanceResponseBody {
	s.Performances = v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBody) SetRequestId(v string) *DescribeDBClusterResourcePoolPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBody) SetStartTime(v string) *DescribeDBClusterResourcePoolPerformanceResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances struct {
	Key                      *string                                                                                     `json:"Key,omitempty" xml:"Key,omitempty"`
	ResourcePoolPerformances []*DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances `json:"ResourcePoolPerformances,omitempty" xml:"ResourcePoolPerformances,omitempty" type:"Repeated"`
	Unit                     *string                                                                                     `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances) SetKey(v string) *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances) SetResourcePoolPerformances(v []*DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances) *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances {
	s.ResourcePoolPerformances = v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances) SetUnit(v string) *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances {
	s.Unit = &v
	return s
}

type DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances struct {
	ResourcePoolName   *string                                                                                                       `json:"ResourcePoolName,omitempty" xml:"ResourcePoolName,omitempty"`
	ResourcePoolSeries []*DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries `json:"ResourcePoolSeries,omitempty" xml:"ResourcePoolSeries,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances) SetResourcePoolName(v string) *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances {
	s.ResourcePoolName = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances) SetResourcePoolSeries(v []*DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries) *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances {
	s.ResourcePoolSeries = v
	return s
}

type DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries struct {
	Name   *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries) SetName(v string) *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries {
	s.Name = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries) SetValues(v []*string) *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries {
	s.Values = v
	return s
}

type DescribeDBClusterResourcePoolPerformanceResponse struct {
	Headers map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBClusterResourcePoolPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterResourcePoolPerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterResourcePoolPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterResourcePoolPerformanceResponse) SetHeaders(v map[string]*string) *DescribeDBClusterResourcePoolPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponse) SetBody(v *DescribeDBClusterResourcePoolPerformanceResponseBody) *DescribeDBClusterResourcePoolPerformanceResponse {
	s.Body = v
	return s
}

type DescribeDBClustersRequest struct {
	DBClusterDescription *string                         `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	DBClusterIds         *string                         `json:"DBClusterIds,omitempty" xml:"DBClusterIds,omitempty"`
	DBClusterStatus      *string                         `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	OwnerAccount         *string                         `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                          `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string                         `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                          `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Tag                  []*DescribeDBClustersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersRequest) SetDBClusterDescription(v string) *DescribeDBClustersRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDBClusterIds(v string) *DescribeDBClustersRequest {
	s.DBClusterIds = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDBClusterStatus(v string) *DescribeDBClustersRequest {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClustersRequest) SetOwnerAccount(v string) *DescribeDBClustersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClustersRequest) SetOwnerId(v int64) *DescribeDBClustersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClustersRequest) SetPageNumber(v int32) *DescribeDBClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClustersRequest) SetPageSize(v int32) *DescribeDBClustersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBClustersRequest) SetRegionId(v string) *DescribeDBClustersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersRequest) SetResourceGroupId(v string) *DescribeDBClustersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClustersRequest) SetResourceOwnerAccount(v string) *DescribeDBClustersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClustersRequest) SetResourceOwnerId(v int64) *DescribeDBClustersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClustersRequest) SetTag(v []*DescribeDBClustersRequestTag) *DescribeDBClustersRequest {
	s.Tag = v
	return s
}

type DescribeDBClustersRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBClustersRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersRequestTag) SetKey(v string) *DescribeDBClustersRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDBClustersRequestTag) SetValue(v string) *DescribeDBClustersRequestTag {
	s.Value = &v
	return s
}

type DescribeDBClustersResponseBody struct {
	Items      *DescribeDBClustersResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	PageNumber *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDBClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBody) SetItems(v *DescribeDBClustersResponseBodyItems) *DescribeDBClustersResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBClustersResponseBody) SetPageNumber(v int32) *DescribeDBClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClustersResponseBody) SetPageSize(v int32) *DescribeDBClustersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDBClustersResponseBody) SetRequestId(v string) *DescribeDBClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClustersResponseBody) SetTotalCount(v int32) *DescribeDBClustersResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDBClustersResponseBodyItems struct {
	DBCluster []*DescribeDBClustersResponseBodyItemsDBCluster `json:"DBCluster,omitempty" xml:"DBCluster,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyItems) SetDBCluster(v []*DescribeDBClustersResponseBodyItemsDBCluster) *DescribeDBClustersResponseBodyItems {
	s.DBCluster = v
	return s
}

type DescribeDBClustersResponseBodyItemsDBCluster struct {
	Category             *string                                           `json:"Category,omitempty" xml:"Category,omitempty"`
	CommodityCode        *string                                           `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ComputeResource      *string                                           `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	ConnectionString     *string                                           `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	CreateTime           *string                                           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DBClusterDescription *string                                           `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	DBClusterId          *string                                           `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBClusterNetworkType *string                                           `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	DBClusterStatus      *string                                           `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	DBClusterType        *string                                           `json:"DBClusterType,omitempty" xml:"DBClusterType,omitempty"`
	DBNodeClass          *string                                           `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	DBNodeCount          *int64                                            `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	DBNodeStorage        *int64                                            `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	DBVersion            *string                                           `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	DiskType             *string                                           `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	DtsJobId             *string                                           `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	ElasticIOResource    *int32                                            `json:"ElasticIOResource,omitempty" xml:"ElasticIOResource,omitempty"`
	Engine               *string                                           `json:"Engine,omitempty" xml:"Engine,omitempty"`
	ExecutorCount        *string                                           `json:"ExecutorCount,omitempty" xml:"ExecutorCount,omitempty"`
	ExpireTime           *string                                           `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Expired              *string                                           `json:"Expired,omitempty" xml:"Expired,omitempty"`
	LockMode             *string                                           `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	LockReason           *string                                           `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	Mode                 *string                                           `json:"Mode,omitempty" xml:"Mode,omitempty"`
	PayType              *string                                           `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Port                 *string                                           `json:"Port,omitempty" xml:"Port,omitempty"`
	RdsInstanceId        *string                                           `json:"RdsInstanceId,omitempty" xml:"RdsInstanceId,omitempty"`
	RegionId             *string                                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string                                           `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StorageResource      *string                                           `json:"StorageResource,omitempty" xml:"StorageResource,omitempty"`
	Tags                 *DescribeDBClustersResponseBodyItemsDBClusterTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VPCCloudInstanceId   *string                                           `json:"VPCCloudInstanceId,omitempty" xml:"VPCCloudInstanceId,omitempty"`
	VPCId                *string                                           `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId            *string                                           `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId               *string                                           `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBClustersResponseBodyItemsDBCluster) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBodyItemsDBCluster) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetCategory(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Category = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetCommodityCode(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetComputeResource(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ComputeResource = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetConnectionString(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetCreateTime(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBClusterDescription(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBClusterId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBClusterNetworkType(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBClusterNetworkType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBClusterStatus(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBClusterType(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBClusterType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBNodeClass(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBNodeCount(v int64) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBNodeCount = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBNodeStorage(v int64) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBNodeStorage = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBVersion(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDiskType(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DiskType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDtsJobId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetElasticIOResource(v int32) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ElasticIOResource = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetEngine(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Engine = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetExecutorCount(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ExecutorCount = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetExpireTime(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetExpired(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Expired = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetLockMode(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.LockMode = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetLockReason(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.LockReason = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetMode(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Mode = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetPayType(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.PayType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetPort(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Port = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetRdsInstanceId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.RdsInstanceId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetRegionId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetResourceGroupId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetStorageResource(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.StorageResource = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetTags(v *DescribeDBClustersResponseBodyItemsDBClusterTags) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Tags = v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetVPCCloudInstanceId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.VPCCloudInstanceId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetVPCId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.VPCId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetVSwitchId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetZoneId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ZoneId = &v
	return s
}

type DescribeDBClustersResponseBodyItemsDBClusterTags struct {
	Tag []*DescribeDBClustersResponseBodyItemsDBClusterTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersResponseBodyItemsDBClusterTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBodyItemsDBClusterTags) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTags) SetTag(v []*DescribeDBClustersResponseBodyItemsDBClusterTagsTag) *DescribeDBClustersResponseBodyItemsDBClusterTags {
	s.Tag = v
	return s
}

type DescribeDBClustersResponseBodyItemsDBClusterTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBClustersResponseBodyItemsDBClusterTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBodyItemsDBClusterTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTagsTag) SetKey(v string) *DescribeDBClustersResponseBodyItemsDBClusterTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTagsTag) SetValue(v string) *DescribeDBClustersResponseBodyItemsDBClusterTagsTag {
	s.Value = &v
	return s
}

type DescribeDBClustersResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponse) SetHeaders(v map[string]*string) *DescribeDBClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClustersResponse) SetBody(v *DescribeDBClustersResponseBody) *DescribeDBClustersResponse {
	s.Body = v
	return s
}

type DescribeDBResourcePoolRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PoolName             *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBResourcePoolRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBResourcePoolRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBResourcePoolRequest) SetDBClusterId(v string) *DescribeDBResourcePoolRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBResourcePoolRequest) SetOwnerAccount(v string) *DescribeDBResourcePoolRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBResourcePoolRequest) SetOwnerId(v int64) *DescribeDBResourcePoolRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBResourcePoolRequest) SetPoolName(v string) *DescribeDBResourcePoolRequest {
	s.PoolName = &v
	return s
}

func (s *DescribeDBResourcePoolRequest) SetResourceOwnerAccount(v string) *DescribeDBResourcePoolRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBResourcePoolRequest) SetResourceOwnerId(v int64) *DescribeDBResourcePoolRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeDBResourcePoolResponseBody struct {
	PoolsInfo []*DescribeDBResourcePoolResponseBodyPoolsInfo `json:"PoolsInfo,omitempty" xml:"PoolsInfo,omitempty" type:"Repeated"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBResourcePoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBResourcePoolResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBResourcePoolResponseBody) SetPoolsInfo(v []*DescribeDBResourcePoolResponseBodyPoolsInfo) *DescribeDBResourcePoolResponseBody {
	s.PoolsInfo = v
	return s
}

func (s *DescribeDBResourcePoolResponseBody) SetRequestId(v string) *DescribeDBResourcePoolResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBResourcePoolResponseBodyPoolsInfo struct {
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	NodeNum    *int32  `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	PoolName   *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	PoolUsers  *string `json:"PoolUsers,omitempty" xml:"PoolUsers,omitempty"`
	QueryType  *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeDBResourcePoolResponseBodyPoolsInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBResourcePoolResponseBodyPoolsInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) SetCreateTime(v string) *DescribeDBResourcePoolResponseBodyPoolsInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) SetNodeNum(v int32) *DescribeDBResourcePoolResponseBodyPoolsInfo {
	s.NodeNum = &v
	return s
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) SetPoolName(v string) *DescribeDBResourcePoolResponseBodyPoolsInfo {
	s.PoolName = &v
	return s
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) SetPoolUsers(v string) *DescribeDBResourcePoolResponseBodyPoolsInfo {
	s.PoolUsers = &v
	return s
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) SetQueryType(v string) *DescribeDBResourcePoolResponseBodyPoolsInfo {
	s.QueryType = &v
	return s
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) SetUpdateTime(v string) *DescribeDBResourcePoolResponseBodyPoolsInfo {
	s.UpdateTime = &v
	return s
}

type DescribeDBResourcePoolResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBResourcePoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBResourcePoolResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBResourcePoolResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBResourcePoolResponse) SetHeaders(v map[string]*string) *DescribeDBResourcePoolResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBResourcePoolResponse) SetBody(v *DescribeDBResourcePoolResponseBody) *DescribeDBResourcePoolResponse {
	s.Body = v
	return s
}

type DescribeDiagnosisDimensionsRequest struct {
	DBClusterId    *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	QueryCondition *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDiagnosisDimensionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisDimensionsRequest) GoString() string {
	return s.String()
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

type DescribeDiagnosisDimensionsResponseBody struct {
	ClientIps      []*string `json:"ClientIps,omitempty" xml:"ClientIps,omitempty" type:"Repeated"`
	Databases      []*string `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroups []*string `json:"ResourceGroups,omitempty" xml:"ResourceGroups,omitempty" type:"Repeated"`
	UserNames      []*string `json:"UserNames,omitempty" xml:"UserNames,omitempty" type:"Repeated"`
}

func (s DescribeDiagnosisDimensionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisDimensionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisDimensionsResponseBody) SetClientIps(v []*string) *DescribeDiagnosisDimensionsResponseBody {
	s.ClientIps = v
	return s
}

func (s *DescribeDiagnosisDimensionsResponseBody) SetDatabases(v []*string) *DescribeDiagnosisDimensionsResponseBody {
	s.Databases = v
	return s
}

func (s *DescribeDiagnosisDimensionsResponseBody) SetRequestId(v string) *DescribeDiagnosisDimensionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosisDimensionsResponseBody) SetResourceGroups(v []*string) *DescribeDiagnosisDimensionsResponseBody {
	s.ResourceGroups = v
	return s
}

func (s *DescribeDiagnosisDimensionsResponseBody) SetUserNames(v []*string) *DescribeDiagnosisDimensionsResponseBody {
	s.UserNames = v
	return s
}

type DescribeDiagnosisDimensionsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDiagnosisDimensionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDiagnosisDimensionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisDimensionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisDimensionsResponse) SetHeaders(v map[string]*string) *DescribeDiagnosisDimensionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosisDimensionsResponse) SetBody(v *DescribeDiagnosisDimensionsResponseBody) *DescribeDiagnosisDimensionsResponse {
	s.Body = v
	return s
}

type DescribeDiagnosisRecordsRequest struct {
	ClientIp       *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	DBClusterId    *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Database       *string `json:"Database,omitempty" xml:"Database,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Keyword        *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	MaxPeakMemory  *int64  `json:"MaxPeakMemory,omitempty" xml:"MaxPeakMemory,omitempty"`
	MaxScanSize    *int64  `json:"MaxScanSize,omitempty" xml:"MaxScanSize,omitempty"`
	MinPeakMemory  *int64  `json:"MinPeakMemory,omitempty" xml:"MinPeakMemory,omitempty"`
	MinScanSize    *int64  `json:"MinScanSize,omitempty" xml:"MinScanSize,omitempty"`
	Order          *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PatternId      *int64  `json:"PatternId,omitempty" xml:"PatternId,omitempty"`
	QueryCondition *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroup  *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UserName       *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDiagnosisRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisRecordsRequest) SetClientIp(v string) *DescribeDiagnosisRecordsRequest {
	s.ClientIp = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetDBClusterId(v string) *DescribeDiagnosisRecordsRequest {
	s.DBClusterId = &v
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

func (s *DescribeDiagnosisRecordsRequest) SetLang(v string) *DescribeDiagnosisRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetMaxPeakMemory(v int64) *DescribeDiagnosisRecordsRequest {
	s.MaxPeakMemory = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetMaxScanSize(v int64) *DescribeDiagnosisRecordsRequest {
	s.MaxScanSize = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetMinPeakMemory(v int64) *DescribeDiagnosisRecordsRequest {
	s.MinPeakMemory = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetMinScanSize(v int64) *DescribeDiagnosisRecordsRequest {
	s.MinScanSize = &v
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

func (s *DescribeDiagnosisRecordsRequest) SetPatternId(v int64) *DescribeDiagnosisRecordsRequest {
	s.PatternId = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetQueryCondition(v string) *DescribeDiagnosisRecordsRequest {
	s.QueryCondition = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetRegionId(v string) *DescribeDiagnosisRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetResourceGroup(v string) *DescribeDiagnosisRecordsRequest {
	s.ResourceGroup = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetStartTime(v string) *DescribeDiagnosisRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetUserName(v string) *DescribeDiagnosisRecordsRequest {
	s.UserName = &v
	return s
}

type DescribeDiagnosisRecordsResponseBody struct {
	PageNumber *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Querys     []*DescribeDiagnosisRecordsResponseBodyQuerys `json:"Querys,omitempty" xml:"Querys,omitempty" type:"Repeated"`
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDiagnosisRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisRecordsResponseBody) SetPageNumber(v int32) *DescribeDiagnosisRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBody) SetPageSize(v int32) *DescribeDiagnosisRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBody) SetQuerys(v []*DescribeDiagnosisRecordsResponseBodyQuerys) *DescribeDiagnosisRecordsResponseBody {
	s.Querys = v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBody) SetRequestId(v string) *DescribeDiagnosisRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBody) SetTotalCount(v int32) *DescribeDiagnosisRecordsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDiagnosisRecordsResponseBodyQuerys struct {
	ClientIp              *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	Cost                  *int64  `json:"Cost,omitempty" xml:"Cost,omitempty"`
	Database              *string `json:"Database,omitempty" xml:"Database,omitempty"`
	EtlWriteRows          *int64  `json:"EtlWriteRows,omitempty" xml:"EtlWriteRows,omitempty"`
	ExecutionTime         *int64  `json:"ExecutionTime,omitempty" xml:"ExecutionTime,omitempty"`
	OutputDataSize        *int64  `json:"OutputDataSize,omitempty" xml:"OutputDataSize,omitempty"`
	OutputRows            *int64  `json:"OutputRows,omitempty" xml:"OutputRows,omitempty"`
	PeakMemory            *int64  `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	ProcessId             *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	QueueTime             *int64  `json:"QueueTime,omitempty" xml:"QueueTime,omitempty"`
	RcHost                *string `json:"RcHost,omitempty" xml:"RcHost,omitempty"`
	ResourceCostRank      *int32  `json:"ResourceCostRank,omitempty" xml:"ResourceCostRank,omitempty"`
	ResourceGroup         *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	SQL                   *string `json:"SQL,omitempty" xml:"SQL,omitempty"`
	SQLTruncated          *bool   `json:"SQLTruncated,omitempty" xml:"SQLTruncated,omitempty"`
	SQLTruncatedThreshold *int64  `json:"SQLTruncatedThreshold,omitempty" xml:"SQLTruncatedThreshold,omitempty"`
	ScanRows              *int64  `json:"ScanRows,omitempty" xml:"ScanRows,omitempty"`
	ScanSize              *int64  `json:"ScanSize,omitempty" xml:"ScanSize,omitempty"`
	StartTime             *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalPlanningTime     *int64  `json:"TotalPlanningTime,omitempty" xml:"TotalPlanningTime,omitempty"`
	TotalStages           *int32  `json:"TotalStages,omitempty" xml:"TotalStages,omitempty"`
	UserName              *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDiagnosisRecordsResponseBodyQuerys) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisRecordsResponseBodyQuerys) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetClientIp(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.ClientIp = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetCost(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.Cost = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetDatabase(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.Database = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetEtlWriteRows(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.EtlWriteRows = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetExecutionTime(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.ExecutionTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetOutputDataSize(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.OutputDataSize = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetOutputRows(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.OutputRows = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetPeakMemory(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.PeakMemory = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetProcessId(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.ProcessId = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetQueueTime(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.QueueTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetRcHost(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.RcHost = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetResourceCostRank(v int32) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.ResourceCostRank = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetResourceGroup(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.ResourceGroup = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetSQL(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.SQL = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetSQLTruncated(v bool) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.SQLTruncated = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetSQLTruncatedThreshold(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.SQLTruncatedThreshold = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetScanRows(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.ScanRows = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetScanSize(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.ScanSize = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetStartTime(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetStatus(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.Status = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetTotalPlanningTime(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.TotalPlanningTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetTotalStages(v int32) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.TotalStages = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetUserName(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.UserName = &v
	return s
}

type DescribeDiagnosisRecordsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDiagnosisRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDiagnosisRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisRecordsResponse) SetHeaders(v map[string]*string) *DescribeDiagnosisRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosisRecordsResponse) SetBody(v *DescribeDiagnosisRecordsResponseBody) *DescribeDiagnosisRecordsResponse {
	s.Body = v
	return s
}

type DescribeDownloadRecordsRequest struct {
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDownloadRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDownloadRecordsRequest) SetDBClusterId(v string) *DescribeDownloadRecordsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDownloadRecordsRequest) SetLang(v string) *DescribeDownloadRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDownloadRecordsRequest) SetRegionId(v string) *DescribeDownloadRecordsRequest {
	s.RegionId = &v
	return s
}

type DescribeDownloadRecordsResponseBody struct {
	Records   []*DescribeDownloadRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDownloadRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDownloadRecordsResponseBody) SetRecords(v []*DescribeDownloadRecordsResponseBodyRecords) *DescribeDownloadRecordsResponseBody {
	s.Records = v
	return s
}

func (s *DescribeDownloadRecordsResponseBody) SetRequestId(v string) *DescribeDownloadRecordsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDownloadRecordsResponseBodyRecords struct {
	DownloadId   *int64  `json:"DownloadId,omitempty" xml:"DownloadId,omitempty"`
	ExceptionMsg *string `json:"ExceptionMsg,omitempty" xml:"ExceptionMsg,omitempty"`
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Url          *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeDownloadRecordsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadRecordsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetDownloadId(v int64) *DescribeDownloadRecordsResponseBodyRecords {
	s.DownloadId = &v
	return s
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetExceptionMsg(v string) *DescribeDownloadRecordsResponseBodyRecords {
	s.ExceptionMsg = &v
	return s
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetFileName(v string) *DescribeDownloadRecordsResponseBodyRecords {
	s.FileName = &v
	return s
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetStatus(v string) *DescribeDownloadRecordsResponseBodyRecords {
	s.Status = &v
	return s
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetUrl(v string) *DescribeDownloadRecordsResponseBodyRecords {
	s.Url = &v
	return s
}

type DescribeDownloadRecordsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDownloadRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDownloadRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDownloadRecordsResponse) SetHeaders(v map[string]*string) *DescribeDownloadRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDownloadRecordsResponse) SetBody(v *DescribeDownloadRecordsResponseBody) *DescribeDownloadRecordsResponse {
	s.Body = v
	return s
}

type DescribeElasticDailyPlanRequest struct {
	DBClusterId                *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ElasticDailyPlanDay        *string `json:"ElasticDailyPlanDay,omitempty" xml:"ElasticDailyPlanDay,omitempty"`
	ElasticDailyPlanStatusList *string `json:"ElasticDailyPlanStatusList,omitempty" xml:"ElasticDailyPlanStatusList,omitempty"`
	ElasticPlanName            *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	OwnerAccount               *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount       *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId            *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourcePoolName           *string `json:"ResourcePoolName,omitempty" xml:"ResourcePoolName,omitempty"`
}

func (s DescribeElasticDailyPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticDailyPlanRequest) GoString() string {
	return s.String()
}

func (s *DescribeElasticDailyPlanRequest) SetDBClusterId(v string) *DescribeElasticDailyPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeElasticDailyPlanRequest) SetElasticDailyPlanDay(v string) *DescribeElasticDailyPlanRequest {
	s.ElasticDailyPlanDay = &v
	return s
}

func (s *DescribeElasticDailyPlanRequest) SetElasticDailyPlanStatusList(v string) *DescribeElasticDailyPlanRequest {
	s.ElasticDailyPlanStatusList = &v
	return s
}

func (s *DescribeElasticDailyPlanRequest) SetElasticPlanName(v string) *DescribeElasticDailyPlanRequest {
	s.ElasticPlanName = &v
	return s
}

func (s *DescribeElasticDailyPlanRequest) SetOwnerAccount(v string) *DescribeElasticDailyPlanRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeElasticDailyPlanRequest) SetOwnerId(v int64) *DescribeElasticDailyPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeElasticDailyPlanRequest) SetResourceOwnerAccount(v string) *DescribeElasticDailyPlanRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeElasticDailyPlanRequest) SetResourceOwnerId(v int64) *DescribeElasticDailyPlanRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeElasticDailyPlanRequest) SetResourcePoolName(v string) *DescribeElasticDailyPlanRequest {
	s.ResourcePoolName = &v
	return s
}

type DescribeElasticDailyPlanResponseBody struct {
	ElasticDailyPlanList []*DescribeElasticDailyPlanResponseBodyElasticDailyPlanList `json:"ElasticDailyPlanList,omitempty" xml:"ElasticDailyPlanList,omitempty" type:"Repeated"`
	RequestId            *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeElasticDailyPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticDailyPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticDailyPlanResponseBody) SetElasticDailyPlanList(v []*DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) *DescribeElasticDailyPlanResponseBody {
	s.ElasticDailyPlanList = v
	return s
}

func (s *DescribeElasticDailyPlanResponseBody) SetRequestId(v string) *DescribeElasticDailyPlanResponseBody {
	s.RequestId = &v
	return s
}

type DescribeElasticDailyPlanResponseBodyElasticDailyPlanList struct {
	Day              *string `json:"Day,omitempty" xml:"Day,omitempty"`
	ElasticNodeNum   *int32  `json:"ElasticNodeNum,omitempty" xml:"ElasticNodeNum,omitempty"`
	EndTs            *string `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	PlanEndTs        *string `json:"PlanEndTs,omitempty" xml:"PlanEndTs,omitempty"`
	PlanName         *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	PlanStartTs      *string `json:"PlanStartTs,omitempty" xml:"PlanStartTs,omitempty"`
	ResourcePoolName *string `json:"ResourcePoolName,omitempty" xml:"ResourcePoolName,omitempty"`
	StartTs          *string `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
	Status           *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) GoString() string {
	return s.String()
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetDay(v string) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.Day = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetElasticNodeNum(v int32) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.ElasticNodeNum = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetEndTs(v string) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.EndTs = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetPlanEndTs(v string) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.PlanEndTs = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetPlanName(v string) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.PlanName = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetPlanStartTs(v string) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.PlanStartTs = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetResourcePoolName(v string) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.ResourcePoolName = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetStartTs(v string) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.StartTs = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetStatus(v int32) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.Status = &v
	return s
}

type DescribeElasticDailyPlanResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeElasticDailyPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeElasticDailyPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticDailyPlanResponse) GoString() string {
	return s.String()
}

func (s *DescribeElasticDailyPlanResponse) SetHeaders(v map[string]*string) *DescribeElasticDailyPlanResponse {
	s.Headers = v
	return s
}

func (s *DescribeElasticDailyPlanResponse) SetBody(v *DescribeElasticDailyPlanResponseBody) *DescribeElasticDailyPlanResponse {
	s.Body = v
	return s
}

type DescribeElasticPlanRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ElasticPlanEnable    *bool   `json:"ElasticPlanEnable,omitempty" xml:"ElasticPlanEnable,omitempty"`
	ElasticPlanName      *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourcePoolName     *string `json:"ResourcePoolName,omitempty" xml:"ResourcePoolName,omitempty"`
}

func (s DescribeElasticPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticPlanRequest) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanRequest) SetDBClusterId(v string) *DescribeElasticPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeElasticPlanRequest) SetElasticPlanEnable(v bool) *DescribeElasticPlanRequest {
	s.ElasticPlanEnable = &v
	return s
}

func (s *DescribeElasticPlanRequest) SetElasticPlanName(v string) *DescribeElasticPlanRequest {
	s.ElasticPlanName = &v
	return s
}

func (s *DescribeElasticPlanRequest) SetOwnerAccount(v string) *DescribeElasticPlanRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeElasticPlanRequest) SetOwnerId(v int64) *DescribeElasticPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeElasticPlanRequest) SetResourceOwnerAccount(v string) *DescribeElasticPlanRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeElasticPlanRequest) SetResourceOwnerId(v int64) *DescribeElasticPlanRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeElasticPlanRequest) SetResourcePoolName(v string) *DescribeElasticPlanRequest {
	s.ResourcePoolName = &v
	return s
}

type DescribeElasticPlanResponseBody struct {
	ElasticPlanList []*DescribeElasticPlanResponseBodyElasticPlanList `json:"ElasticPlanList,omitempty" xml:"ElasticPlanList,omitempty" type:"Repeated"`
	RequestId       *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeElasticPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanResponseBody) SetElasticPlanList(v []*DescribeElasticPlanResponseBodyElasticPlanList) *DescribeElasticPlanResponseBody {
	s.ElasticPlanList = v
	return s
}

func (s *DescribeElasticPlanResponseBody) SetRequestId(v string) *DescribeElasticPlanResponseBody {
	s.RequestId = &v
	return s
}

type DescribeElasticPlanResponseBodyElasticPlanList struct {
	ElasticNodeNum   *int32  `json:"ElasticNodeNum,omitempty" xml:"ElasticNodeNum,omitempty"`
	Enable           *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
	EndDay           *string `json:"EndDay,omitempty" xml:"EndDay,omitempty"`
	EndTime          *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PlanName         *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	ResourcePoolName *string `json:"ResourcePoolName,omitempty" xml:"ResourcePoolName,omitempty"`
	StartDay         *string `json:"StartDay,omitempty" xml:"StartDay,omitempty"`
	StartTime        *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	WeeklyRepeat     *string `json:"WeeklyRepeat,omitempty" xml:"WeeklyRepeat,omitempty"`
}

func (s DescribeElasticPlanResponseBodyElasticPlanList) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticPlanResponseBodyElasticPlanList) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetElasticNodeNum(v int32) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.ElasticNodeNum = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetEnable(v bool) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.Enable = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetEndDay(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.EndDay = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetEndTime(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.EndTime = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetPlanName(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.PlanName = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetResourcePoolName(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.ResourcePoolName = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetStartDay(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.StartDay = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetStartTime(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.StartTime = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetWeeklyRepeat(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.WeeklyRepeat = &v
	return s
}

type DescribeElasticPlanResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeElasticPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeElasticPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticPlanResponse) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanResponse) SetHeaders(v map[string]*string) *DescribeElasticPlanResponse {
	s.Headers = v
	return s
}

func (s *DescribeElasticPlanResponse) SetBody(v *DescribeElasticPlanResponseBody) *DescribeElasticPlanResponse {
	s.Body = v
	return s
}

type DescribeInclinedTablesRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Order                *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TableType            *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
}

func (s DescribeInclinedTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInclinedTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInclinedTablesRequest) SetDBClusterId(v string) *DescribeInclinedTablesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetOrder(v string) *DescribeInclinedTablesRequest {
	s.Order = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetOwnerAccount(v string) *DescribeInclinedTablesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetOwnerId(v int64) *DescribeInclinedTablesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetPageNumber(v int32) *DescribeInclinedTablesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetPageSize(v int32) *DescribeInclinedTablesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetResourceOwnerAccount(v string) *DescribeInclinedTablesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetResourceOwnerId(v int64) *DescribeInclinedTablesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetTableType(v string) *DescribeInclinedTablesRequest {
	s.TableType = &v
	return s
}

type DescribeInclinedTablesResponseBody struct {
	Items      *DescribeInclinedTablesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	PageNumber *string                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *string                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInclinedTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInclinedTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInclinedTablesResponseBody) SetItems(v *DescribeInclinedTablesResponseBodyItems) *DescribeInclinedTablesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeInclinedTablesResponseBody) SetPageNumber(v string) *DescribeInclinedTablesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInclinedTablesResponseBody) SetPageSize(v string) *DescribeInclinedTablesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInclinedTablesResponseBody) SetRequestId(v string) *DescribeInclinedTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInclinedTablesResponseBody) SetTotalCount(v string) *DescribeInclinedTablesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeInclinedTablesResponseBodyItems struct {
	Table []*DescribeInclinedTablesResponseBodyItemsTable `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeInclinedTablesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeInclinedTablesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeInclinedTablesResponseBodyItems) SetTable(v []*DescribeInclinedTablesResponseBodyItemsTable) *DescribeInclinedTablesResponseBodyItems {
	s.Table = v
	return s
}

type DescribeInclinedTablesResponseBodyItemsTable struct {
	IsIncline *bool   `json:"IsIncline,omitempty" xml:"IsIncline,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Schema    *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	Size      *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeInclinedTablesResponseBodyItemsTable) String() string {
	return tea.Prettify(s)
}

func (s DescribeInclinedTablesResponseBodyItemsTable) GoString() string {
	return s.String()
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) SetIsIncline(v bool) *DescribeInclinedTablesResponseBodyItemsTable {
	s.IsIncline = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) SetName(v string) *DescribeInclinedTablesResponseBodyItemsTable {
	s.Name = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) SetSchema(v string) *DescribeInclinedTablesResponseBodyItemsTable {
	s.Schema = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) SetSize(v int64) *DescribeInclinedTablesResponseBodyItemsTable {
	s.Size = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) SetType(v string) *DescribeInclinedTablesResponseBodyItemsTable {
	s.Type = &v
	return s
}

type DescribeInclinedTablesResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInclinedTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInclinedTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInclinedTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInclinedTablesResponse) SetHeaders(v map[string]*string) *DescribeInclinedTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInclinedTablesResponse) SetBody(v *DescribeInclinedTablesResponseBody) *DescribeInclinedTablesResponse {
	s.Body = v
	return s
}

type DescribeLoadTasksRecordsRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Order                *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State                *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeLoadTasksRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadTasksRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoadTasksRecordsRequest) SetDBClusterId(v string) *DescribeLoadTasksRecordsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetDBName(v string) *DescribeLoadTasksRecordsRequest {
	s.DBName = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetEndTime(v string) *DescribeLoadTasksRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetOrder(v string) *DescribeLoadTasksRecordsRequest {
	s.Order = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetOwnerAccount(v string) *DescribeLoadTasksRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetOwnerId(v int64) *DescribeLoadTasksRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetPageNumber(v int32) *DescribeLoadTasksRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetPageSize(v int32) *DescribeLoadTasksRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetResourceOwnerAccount(v string) *DescribeLoadTasksRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetResourceOwnerId(v int64) *DescribeLoadTasksRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetStartTime(v string) *DescribeLoadTasksRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetState(v string) *DescribeLoadTasksRecordsRequest {
	s.State = &v
	return s
}

type DescribeLoadTasksRecordsResponseBody struct {
	DBClusterId      *string                                                 `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	LoadTasksRecords []*DescribeLoadTasksRecordsResponseBodyLoadTasksRecords `json:"LoadTasksRecords,omitempty" xml:"LoadTasksRecords,omitempty" type:"Repeated"`
	PageNumber       *string                                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *string                                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId        *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount       *string                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLoadTasksRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadTasksRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadTasksRecordsResponseBody) SetDBClusterId(v string) *DescribeLoadTasksRecordsResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBody) SetLoadTasksRecords(v []*DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) *DescribeLoadTasksRecordsResponseBody {
	s.LoadTasksRecords = v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBody) SetPageNumber(v string) *DescribeLoadTasksRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBody) SetPageSize(v string) *DescribeLoadTasksRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBody) SetRequestId(v string) *DescribeLoadTasksRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBody) SetTotalCount(v string) *DescribeLoadTasksRecordsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeLoadTasksRecordsResponseBodyLoadTasksRecords struct {
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DBName      *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	JobName     *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	ProcessID   *string `json:"ProcessID,omitempty" xml:"ProcessID,omitempty"`
	ProcessRows *int64  `json:"ProcessRows,omitempty" xml:"ProcessRows,omitempty"`
	Sql         *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
	State       *string `json:"State,omitempty" xml:"State,omitempty"`
	UpdateTime  *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) GoString() string {
	return s.String()
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetCreateTime(v string) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.CreateTime = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetDBName(v string) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.DBName = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetJobName(v string) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.JobName = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetProcessID(v string) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.ProcessID = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetProcessRows(v int64) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.ProcessRows = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetSql(v string) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.Sql = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetState(v string) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.State = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetUpdateTime(v string) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.UpdateTime = &v
	return s
}

type DescribeLoadTasksRecordsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLoadTasksRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLoadTasksRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadTasksRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLoadTasksRecordsResponse) SetHeaders(v map[string]*string) *DescribeLoadTasksRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLoadTasksRecordsResponse) SetBody(v *DescribeLoadTasksRecordsResponseBody) *DescribeLoadTasksRecordsResponse {
	s.Body = v
	return s
}

type DescribeMaintenanceActionRequest struct {
	IsHistory            *int32  `json:"IsHistory,omitempty" xml:"IsHistory,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskType             *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeMaintenanceActionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMaintenanceActionRequest) GoString() string {
	return s.String()
}

func (s *DescribeMaintenanceActionRequest) SetIsHistory(v int32) *DescribeMaintenanceActionRequest {
	s.IsHistory = &v
	return s
}

func (s *DescribeMaintenanceActionRequest) SetOwnerAccount(v string) *DescribeMaintenanceActionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeMaintenanceActionRequest) SetOwnerId(v int64) *DescribeMaintenanceActionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMaintenanceActionRequest) SetPageNumber(v int32) *DescribeMaintenanceActionRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMaintenanceActionRequest) SetPageSize(v int32) *DescribeMaintenanceActionRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMaintenanceActionRequest) SetRegion(v string) *DescribeMaintenanceActionRequest {
	s.Region = &v
	return s
}

func (s *DescribeMaintenanceActionRequest) SetRegionId(v string) *DescribeMaintenanceActionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMaintenanceActionRequest) SetResourceOwnerAccount(v string) *DescribeMaintenanceActionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeMaintenanceActionRequest) SetResourceOwnerId(v int64) *DescribeMaintenanceActionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeMaintenanceActionRequest) SetTaskType(v string) *DescribeMaintenanceActionRequest {
	s.TaskType = &v
	return s
}

type DescribeMaintenanceActionResponseBody struct {
	Items            []*DescribeMaintenanceActionResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber       *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId        *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalRecordCount *int32                                        `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeMaintenanceActionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMaintenanceActionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMaintenanceActionResponseBody) SetItems(v []*DescribeMaintenanceActionResponseBodyItems) *DescribeMaintenanceActionResponseBody {
	s.Items = v
	return s
}

func (s *DescribeMaintenanceActionResponseBody) SetPageNumber(v int32) *DescribeMaintenanceActionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBody) SetPageSize(v int32) *DescribeMaintenanceActionResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBody) SetRequestId(v string) *DescribeMaintenanceActionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBody) SetTotalRecordCount(v int32) *DescribeMaintenanceActionResponseBody {
	s.TotalRecordCount = &v
	return s
}

type DescribeMaintenanceActionResponseBodyItems struct {
	CreatedTime     *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	DBClusterId     *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBType          *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	DBVersion       *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	Deadline        *string `json:"Deadline,omitempty" xml:"Deadline,omitempty"`
	Id              *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifiedTime    *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	PrepareInterval *string `json:"PrepareInterval,omitempty" xml:"PrepareInterval,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResultInfo      *string `json:"ResultInfo,omitempty" xml:"ResultInfo,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SwitchTime      *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	TaskType        *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeMaintenanceActionResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeMaintenanceActionResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetCreatedTime(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.CreatedTime = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetDBClusterId(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.DBClusterId = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetDBType(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.DBType = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetDBVersion(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.DBVersion = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetDeadline(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.Deadline = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetId(v int32) *DescribeMaintenanceActionResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetModifiedTime(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetPrepareInterval(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.PrepareInterval = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetRegion(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.Region = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetResultInfo(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.ResultInfo = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetStartTime(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetStatus(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetSwitchTime(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.SwitchTime = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetTaskType(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.TaskType = &v
	return s
}

type DescribeMaintenanceActionResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMaintenanceActionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMaintenanceActionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMaintenanceActionResponse) GoString() string {
	return s.String()
}

func (s *DescribeMaintenanceActionResponse) SetHeaders(v map[string]*string) *DescribeMaintenanceActionResponse {
	s.Headers = v
	return s
}

func (s *DescribeMaintenanceActionResponse) SetBody(v *DescribeMaintenanceActionResponseBody) *DescribeMaintenanceActionResponse {
	s.Body = v
	return s
}

type DescribeOperatorPermissionRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeOperatorPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOperatorPermissionRequest) GoString() string {
	return s.String()
}

func (s *DescribeOperatorPermissionRequest) SetDBClusterId(v string) *DescribeOperatorPermissionRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeOperatorPermissionRequest) SetOwnerAccount(v string) *DescribeOperatorPermissionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeOperatorPermissionRequest) SetOwnerId(v int64) *DescribeOperatorPermissionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeOperatorPermissionRequest) SetResourceOwnerAccount(v string) *DescribeOperatorPermissionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeOperatorPermissionRequest) SetResourceOwnerId(v int64) *DescribeOperatorPermissionRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeOperatorPermissionResponseBody struct {
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Privileges  *string `json:"Privileges,omitempty" xml:"Privileges,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOperatorPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOperatorPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOperatorPermissionResponseBody) SetCreatedTime(v string) *DescribeOperatorPermissionResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeOperatorPermissionResponseBody) SetDBClusterId(v string) *DescribeOperatorPermissionResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeOperatorPermissionResponseBody) SetExpiredTime(v string) *DescribeOperatorPermissionResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeOperatorPermissionResponseBody) SetPrivileges(v string) *DescribeOperatorPermissionResponseBody {
	s.Privileges = &v
	return s
}

func (s *DescribeOperatorPermissionResponseBody) SetRequestId(v string) *DescribeOperatorPermissionResponseBody {
	s.RequestId = &v
	return s
}

type DescribeOperatorPermissionResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeOperatorPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOperatorPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOperatorPermissionResponse) GoString() string {
	return s.String()
}

func (s *DescribeOperatorPermissionResponse) SetHeaders(v map[string]*string) *DescribeOperatorPermissionResponse {
	s.Headers = v
	return s
}

func (s *DescribeOperatorPermissionResponse) SetBody(v *DescribeOperatorPermissionResponseBody) *DescribeOperatorPermissionResponse {
	s.Body = v
	return s
}

type DescribePatternPerformanceRequest struct {
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PatternId   *string `json:"PatternId,omitempty" xml:"PatternId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePatternPerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePatternPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribePatternPerformanceRequest) SetDBClusterId(v string) *DescribePatternPerformanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribePatternPerformanceRequest) SetEndTime(v string) *DescribePatternPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePatternPerformanceRequest) SetPatternId(v string) *DescribePatternPerformanceRequest {
	s.PatternId = &v
	return s
}

func (s *DescribePatternPerformanceRequest) SetRegionId(v string) *DescribePatternPerformanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePatternPerformanceRequest) SetStartTime(v string) *DescribePatternPerformanceRequest {
	s.StartTime = &v
	return s
}

type DescribePatternPerformanceResponseBody struct {
	EndTime      *string                                               `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Performances []*DescribePatternPerformanceResponseBodyPerformances `json:"Performances,omitempty" xml:"Performances,omitempty" type:"Repeated"`
	RequestId    *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime    *string                                               `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePatternPerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePatternPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePatternPerformanceResponseBody) SetEndTime(v string) *DescribePatternPerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribePatternPerformanceResponseBody) SetPerformances(v []*DescribePatternPerformanceResponseBodyPerformances) *DescribePatternPerformanceResponseBody {
	s.Performances = v
	return s
}

func (s *DescribePatternPerformanceResponseBody) SetRequestId(v string) *DescribePatternPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePatternPerformanceResponseBody) SetStartTime(v string) *DescribePatternPerformanceResponseBody {
	s.StartTime = &v
	return s
}

type DescribePatternPerformanceResponseBodyPerformances struct {
	Key    *string                                                     `json:"Key,omitempty" xml:"Key,omitempty"`
	Series []*DescribePatternPerformanceResponseBodyPerformancesSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
	Unit   *string                                                     `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribePatternPerformanceResponseBodyPerformances) String() string {
	return tea.Prettify(s)
}

func (s DescribePatternPerformanceResponseBodyPerformances) GoString() string {
	return s.String()
}

func (s *DescribePatternPerformanceResponseBodyPerformances) SetKey(v string) *DescribePatternPerformanceResponseBodyPerformances {
	s.Key = &v
	return s
}

func (s *DescribePatternPerformanceResponseBodyPerformances) SetSeries(v []*DescribePatternPerformanceResponseBodyPerformancesSeries) *DescribePatternPerformanceResponseBodyPerformances {
	s.Series = v
	return s
}

func (s *DescribePatternPerformanceResponseBodyPerformances) SetUnit(v string) *DescribePatternPerformanceResponseBodyPerformances {
	s.Unit = &v
	return s
}

type DescribePatternPerformanceResponseBodyPerformancesSeries struct {
	Name   *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribePatternPerformanceResponseBodyPerformancesSeries) String() string {
	return tea.Prettify(s)
}

func (s DescribePatternPerformanceResponseBodyPerformancesSeries) GoString() string {
	return s.String()
}

func (s *DescribePatternPerformanceResponseBodyPerformancesSeries) SetName(v string) *DescribePatternPerformanceResponseBodyPerformancesSeries {
	s.Name = &v
	return s
}

func (s *DescribePatternPerformanceResponseBodyPerformancesSeries) SetValues(v []*string) *DescribePatternPerformanceResponseBodyPerformancesSeries {
	s.Values = v
	return s
}

type DescribePatternPerformanceResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePatternPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePatternPerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePatternPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribePatternPerformanceResponse) SetHeaders(v map[string]*string) *DescribePatternPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribePatternPerformanceResponse) SetBody(v *DescribePatternPerformanceResponseBody) *DescribePatternPerformanceResponse {
	s.Body = v
	return s
}

type DescribeProcessListRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Keyword              *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Order                *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RunningTime          *int32  `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	ShowFull             *bool   `json:"ShowFull,omitempty" xml:"ShowFull,omitempty"`
	User                 *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeProcessListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessListRequest) GoString() string {
	return s.String()
}

func (s *DescribeProcessListRequest) SetDBClusterId(v string) *DescribeProcessListRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeProcessListRequest) SetKeyword(v string) *DescribeProcessListRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeProcessListRequest) SetOrder(v string) *DescribeProcessListRequest {
	s.Order = &v
	return s
}

func (s *DescribeProcessListRequest) SetOwnerAccount(v string) *DescribeProcessListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeProcessListRequest) SetOwnerId(v int64) *DescribeProcessListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeProcessListRequest) SetPageNumber(v int32) *DescribeProcessListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeProcessListRequest) SetPageSize(v int32) *DescribeProcessListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeProcessListRequest) SetResourceOwnerAccount(v string) *DescribeProcessListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeProcessListRequest) SetResourceOwnerId(v int64) *DescribeProcessListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeProcessListRequest) SetRunningTime(v int32) *DescribeProcessListRequest {
	s.RunningTime = &v
	return s
}

func (s *DescribeProcessListRequest) SetShowFull(v bool) *DescribeProcessListRequest {
	s.ShowFull = &v
	return s
}

func (s *DescribeProcessListRequest) SetUser(v string) *DescribeProcessListRequest {
	s.User = &v
	return s
}

type DescribeProcessListResponseBody struct {
	Items      *DescribeProcessListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	PageNumber *string                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *string                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeProcessListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBody) SetItems(v *DescribeProcessListResponseBodyItems) *DescribeProcessListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeProcessListResponseBody) SetPageNumber(v string) *DescribeProcessListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeProcessListResponseBody) SetPageSize(v string) *DescribeProcessListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeProcessListResponseBody) SetRequestId(v string) *DescribeProcessListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProcessListResponseBody) SetTotalCount(v string) *DescribeProcessListResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeProcessListResponseBodyItems struct {
	Process []*DescribeProcessListResponseBodyItemsProcess `json:"Process,omitempty" xml:"Process,omitempty" type:"Repeated"`
}

func (s DescribeProcessListResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBodyItems) SetProcess(v []*DescribeProcessListResponseBodyItemsProcess) *DescribeProcessListResponseBodyItems {
	s.Process = v
	return s
}

type DescribeProcessListResponseBodyItemsProcess struct {
	Command   *string `json:"Command,omitempty" xml:"Command,omitempty"`
	DB        *string `json:"DB,omitempty" xml:"DB,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Id        *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	Info      *string `json:"Info,omitempty" xml:"Info,omitempty"`
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Time      *int32  `json:"Time,omitempty" xml:"Time,omitempty"`
	User      *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeProcessListResponseBodyItemsProcess) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessListResponseBodyItemsProcess) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetCommand(v string) *DescribeProcessListResponseBodyItemsProcess {
	s.Command = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetDB(v string) *DescribeProcessListResponseBodyItemsProcess {
	s.DB = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetHost(v string) *DescribeProcessListResponseBodyItemsProcess {
	s.Host = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetId(v int32) *DescribeProcessListResponseBodyItemsProcess {
	s.Id = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetInfo(v string) *DescribeProcessListResponseBodyItemsProcess {
	s.Info = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetProcessId(v string) *DescribeProcessListResponseBodyItemsProcess {
	s.ProcessId = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetStartTime(v string) *DescribeProcessListResponseBodyItemsProcess {
	s.StartTime = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetTime(v int32) *DescribeProcessListResponseBodyItemsProcess {
	s.Time = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetUser(v string) *DescribeProcessListResponseBodyItemsProcess {
	s.User = &v
	return s
}

type DescribeProcessListResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeProcessListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeProcessListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessListResponse) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponse) SetHeaders(v map[string]*string) *DescribeProcessListResponse {
	s.Headers = v
	return s
}

func (s *DescribeProcessListResponse) SetBody(v *DescribeProcessListResponseBody) *DescribeProcessListResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	AcceptLanguage       *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeRegionsRequest) SetOwnerAccount(v string) *DescribeRegionsRequest {
	s.OwnerAccount = &v
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

type DescribeRegionsResponseBody struct {
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	Region []*DescribeRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegion(v []*DescribeRegionsResponseBodyRegionsRegion) *DescribeRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegion struct {
	LocalName      *string                                        `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string                                        `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Zones          *DescribeRegionsResponseBodyRegionsRegionZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetZones(v *DescribeRegionsResponseBodyRegionsRegionZones) *DescribeRegionsResponseBodyRegionsRegion {
	s.Zones = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegionZones struct {
	Zone []*DescribeRegionsResponseBodyRegionsRegionZonesZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegionsRegionZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegionZones) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegionZones) SetZone(v []*DescribeRegionsResponseBodyRegionsRegionZonesZone) *DescribeRegionsResponseBodyRegionsRegionZones {
	s.Zone = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegionZonesZone struct {
	LocalName  *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	VpcEnabled *bool   `json:"VpcEnabled,omitempty" xml:"VpcEnabled,omitempty"`
	ZoneId     *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegionZonesZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegionZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegionZonesZone) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsRegionZonesZone {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegionZonesZone) SetVpcEnabled(v bool) *DescribeRegionsResponseBodyRegionsRegionZonesZone {
	s.VpcEnabled = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegionZonesZone) SetZoneId(v string) *DescribeRegionsResponseBodyRegionsRegionZonesZone {
	s.ZoneId = &v
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

type DescribeSQLPatternAttributeRequest struct {
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PatternId   *int64  `json:"PatternId,omitempty" xml:"PatternId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSQLPatternAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPatternAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLPatternAttributeRequest) SetDBClusterId(v string) *DescribeSQLPatternAttributeRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSQLPatternAttributeRequest) SetEndTime(v string) *DescribeSQLPatternAttributeRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLPatternAttributeRequest) SetLang(v string) *DescribeSQLPatternAttributeRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSQLPatternAttributeRequest) SetPatternId(v int64) *DescribeSQLPatternAttributeRequest {
	s.PatternId = &v
	return s
}

func (s *DescribeSQLPatternAttributeRequest) SetRegionId(v string) *DescribeSQLPatternAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSQLPatternAttributeRequest) SetStartTime(v string) *DescribeSQLPatternAttributeRequest {
	s.StartTime = &v
	return s
}

type DescribeSQLPatternAttributeResponseBody struct {
	PatternDetail *DescribeSQLPatternAttributeResponseBodyPatternDetail `json:"PatternDetail,omitempty" xml:"PatternDetail,omitempty" type:"Struct"`
	RequestId     *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSQLPatternAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPatternAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLPatternAttributeResponseBody) SetPatternDetail(v *DescribeSQLPatternAttributeResponseBodyPatternDetail) *DescribeSQLPatternAttributeResponseBody {
	s.PatternDetail = v
	return s
}

func (s *DescribeSQLPatternAttributeResponseBody) SetRequestId(v string) *DescribeSQLPatternAttributeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSQLPatternAttributeResponseBodyPatternDetail struct {
	AverageMemory    *string `json:"AverageMemory,omitempty" xml:"AverageMemory,omitempty"`
	AverageQueryTime *string `json:"AverageQueryTime,omitempty" xml:"AverageQueryTime,omitempty"`
	QueryCount       *int64  `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
	SQLPattern       *string `json:"SQLPattern,omitempty" xml:"SQLPattern,omitempty"`
	TotalQueryTime   *string `json:"TotalQueryTime,omitempty" xml:"TotalQueryTime,omitempty"`
}

func (s DescribeSQLPatternAttributeResponseBodyPatternDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPatternAttributeResponseBodyPatternDetail) GoString() string {
	return s.String()
}

func (s *DescribeSQLPatternAttributeResponseBodyPatternDetail) SetAverageMemory(v string) *DescribeSQLPatternAttributeResponseBodyPatternDetail {
	s.AverageMemory = &v
	return s
}

func (s *DescribeSQLPatternAttributeResponseBodyPatternDetail) SetAverageQueryTime(v string) *DescribeSQLPatternAttributeResponseBodyPatternDetail {
	s.AverageQueryTime = &v
	return s
}

func (s *DescribeSQLPatternAttributeResponseBodyPatternDetail) SetQueryCount(v int64) *DescribeSQLPatternAttributeResponseBodyPatternDetail {
	s.QueryCount = &v
	return s
}

func (s *DescribeSQLPatternAttributeResponseBodyPatternDetail) SetSQLPattern(v string) *DescribeSQLPatternAttributeResponseBodyPatternDetail {
	s.SQLPattern = &v
	return s
}

func (s *DescribeSQLPatternAttributeResponseBodyPatternDetail) SetTotalQueryTime(v string) *DescribeSQLPatternAttributeResponseBodyPatternDetail {
	s.TotalQueryTime = &v
	return s
}

type DescribeSQLPatternAttributeResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSQLPatternAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQLPatternAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPatternAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLPatternAttributeResponse) SetHeaders(v map[string]*string) *DescribeSQLPatternAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLPatternAttributeResponse) SetBody(v *DescribeSQLPatternAttributeResponseBody) *DescribeSQLPatternAttributeResponse {
	s.Body = v
	return s
}

type DescribeSQLPatternsRequest struct {
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Keyword     *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Order       *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSQLPatternsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPatternsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLPatternsRequest) SetDBClusterId(v string) *DescribeSQLPatternsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetEndTime(v string) *DescribeSQLPatternsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetKeyword(v string) *DescribeSQLPatternsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetLang(v string) *DescribeSQLPatternsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetOrder(v string) *DescribeSQLPatternsRequest {
	s.Order = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetPageNumber(v int32) *DescribeSQLPatternsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetPageSize(v int32) *DescribeSQLPatternsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetRegionId(v string) *DescribeSQLPatternsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetStartTime(v string) *DescribeSQLPatternsRequest {
	s.StartTime = &v
	return s
}

type DescribeSQLPatternsResponseBody struct {
	PageNumber     *int32                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PatternDetails []*DescribeSQLPatternsResponseBodyPatternDetails `json:"PatternDetails,omitempty" xml:"PatternDetails,omitempty" type:"Repeated"`
	RequestId      *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount     *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSQLPatternsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPatternsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLPatternsResponseBody) SetPageNumber(v int32) *DescribeSQLPatternsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLPatternsResponseBody) SetPageSize(v int32) *DescribeSQLPatternsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSQLPatternsResponseBody) SetPatternDetails(v []*DescribeSQLPatternsResponseBodyPatternDetails) *DescribeSQLPatternsResponseBody {
	s.PatternDetails = v
	return s
}

func (s *DescribeSQLPatternsResponseBody) SetRequestId(v string) *DescribeSQLPatternsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLPatternsResponseBody) SetTotalCount(v int32) *DescribeSQLPatternsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeSQLPatternsResponseBodyPatternDetails struct {
	AccessIp             *string  `json:"AccessIp,omitempty" xml:"AccessIp,omitempty"`
	AverageExecutionTime *float64 `json:"AverageExecutionTime,omitempty" xml:"AverageExecutionTime,omitempty"`
	AveragePeakMemory    *float64 `json:"AveragePeakMemory,omitempty" xml:"AveragePeakMemory,omitempty"`
	AverageQueryTime     *float64 `json:"AverageQueryTime,omitempty" xml:"AverageQueryTime,omitempty"`
	AverageScanSize      *float64 `json:"AverageScanSize,omitempty" xml:"AverageScanSize,omitempty"`
	Blockable            *bool    `json:"Blockable,omitempty" xml:"Blockable,omitempty"`
	FailedCount          *int64   `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	MaxExecutionTime     *int64   `json:"MaxExecutionTime,omitempty" xml:"MaxExecutionTime,omitempty"`
	MaxPeakMemory        *int64   `json:"MaxPeakMemory,omitempty" xml:"MaxPeakMemory,omitempty"`
	MaxQueryTime         *int64   `json:"MaxQueryTime,omitempty" xml:"MaxQueryTime,omitempty"`
	MaxScanSize          *int64   `json:"MaxScanSize,omitempty" xml:"MaxScanSize,omitempty"`
	PatternCreationTime  *string  `json:"PatternCreationTime,omitempty" xml:"PatternCreationTime,omitempty"`
	PatternId            *string  `json:"PatternId,omitempty" xml:"PatternId,omitempty"`
	QueryCount           *int64   `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
	SQLPattern           *string  `json:"SQLPattern,omitempty" xml:"SQLPattern,omitempty"`
	Tables               *string  `json:"Tables,omitempty" xml:"Tables,omitempty"`
	User                 *string  `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeSQLPatternsResponseBodyPatternDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPatternsResponseBodyPatternDetails) GoString() string {
	return s.String()
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetAccessIp(v string) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.AccessIp = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetAverageExecutionTime(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.AverageExecutionTime = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetAveragePeakMemory(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.AveragePeakMemory = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetAverageQueryTime(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.AverageQueryTime = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetAverageScanSize(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.AverageScanSize = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetBlockable(v bool) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.Blockable = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetFailedCount(v int64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.FailedCount = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetMaxExecutionTime(v int64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.MaxExecutionTime = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetMaxPeakMemory(v int64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.MaxPeakMemory = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetMaxQueryTime(v int64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.MaxQueryTime = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetMaxScanSize(v int64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.MaxScanSize = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetPatternCreationTime(v string) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.PatternCreationTime = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetPatternId(v string) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.PatternId = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetQueryCount(v int64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.QueryCount = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetSQLPattern(v string) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.SQLPattern = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetTables(v string) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.Tables = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetUser(v string) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.User = &v
	return s
}

type DescribeSQLPatternsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSQLPatternsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQLPatternsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPatternsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLPatternsResponse) SetHeaders(v map[string]*string) *DescribeSQLPatternsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLPatternsResponse) SetBody(v *DescribeSQLPatternsResponseBody) *DescribeSQLPatternsResponse {
	s.Body = v
	return s
}

type DescribeSQLPlanRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProcessId            *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeSQLPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlanRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanRequest) SetDBClusterId(v string) *DescribeSQLPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSQLPlanRequest) SetOwnerAccount(v string) *DescribeSQLPlanRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSQLPlanRequest) SetOwnerId(v int64) *DescribeSQLPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSQLPlanRequest) SetProcessId(v string) *DescribeSQLPlanRequest {
	s.ProcessId = &v
	return s
}

func (s *DescribeSQLPlanRequest) SetResourceOwnerAccount(v string) *DescribeSQLPlanRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSQLPlanRequest) SetResourceOwnerId(v int64) *DescribeSQLPlanRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeSQLPlanResponseBody struct {
	Detail     *DescribeSQLPlanResponseBodyDetail      `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Struct"`
	OriginInfo *string                                 `json:"OriginInfo,omitempty" xml:"OriginInfo,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StageList  []*DescribeSQLPlanResponseBodyStageList `json:"StageList,omitempty" xml:"StageList,omitempty" type:"Repeated"`
}

func (s DescribeSQLPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanResponseBody) SetDetail(v *DescribeSQLPlanResponseBodyDetail) *DescribeSQLPlanResponseBody {
	s.Detail = v
	return s
}

func (s *DescribeSQLPlanResponseBody) SetOriginInfo(v string) *DescribeSQLPlanResponseBody {
	s.OriginInfo = &v
	return s
}

func (s *DescribeSQLPlanResponseBody) SetRequestId(v string) *DescribeSQLPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLPlanResponseBody) SetStageList(v []*DescribeSQLPlanResponseBodyStageList) *DescribeSQLPlanResponseBody {
	s.StageList = v
	return s
}

type DescribeSQLPlanResponseBodyDetail struct {
	CPUTime      *int64  `json:"CPUTime,omitempty" xml:"CPUTime,omitempty"`
	ClientIP     *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
	Database     *string `json:"Database,omitempty" xml:"Database,omitempty"`
	OutputRows   *int64  `json:"OutputRows,omitempty" xml:"OutputRows,omitempty"`
	OutputSize   *int64  `json:"OutputSize,omitempty" xml:"OutputSize,omitempty"`
	PeakMemory   *int64  `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	PlanningTime *int64  `json:"PlanningTime,omitempty" xml:"PlanningTime,omitempty"`
	QueuedTime   *int64  `json:"QueuedTime,omitempty" xml:"QueuedTime,omitempty"`
	SQL          *string `json:"SQL,omitempty" xml:"SQL,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
	TotalStage   *int64  `json:"TotalStage,omitempty" xml:"TotalStage,omitempty"`
	TotalTask    *int64  `json:"TotalTask,omitempty" xml:"TotalTask,omitempty"`
	TotalTime    *int64  `json:"TotalTime,omitempty" xml:"TotalTime,omitempty"`
	User         *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeSQLPlanResponseBodyDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlanResponseBodyDetail) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanResponseBodyDetail) SetCPUTime(v int64) *DescribeSQLPlanResponseBodyDetail {
	s.CPUTime = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetClientIP(v string) *DescribeSQLPlanResponseBodyDetail {
	s.ClientIP = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetDatabase(v string) *DescribeSQLPlanResponseBodyDetail {
	s.Database = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetOutputRows(v int64) *DescribeSQLPlanResponseBodyDetail {
	s.OutputRows = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetOutputSize(v int64) *DescribeSQLPlanResponseBodyDetail {
	s.OutputSize = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetPeakMemory(v int64) *DescribeSQLPlanResponseBodyDetail {
	s.PeakMemory = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetPlanningTime(v int64) *DescribeSQLPlanResponseBodyDetail {
	s.PlanningTime = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetQueuedTime(v int64) *DescribeSQLPlanResponseBodyDetail {
	s.QueuedTime = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetSQL(v string) *DescribeSQLPlanResponseBodyDetail {
	s.SQL = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetStartTime(v string) *DescribeSQLPlanResponseBodyDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetState(v string) *DescribeSQLPlanResponseBodyDetail {
	s.State = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetTotalStage(v int64) *DescribeSQLPlanResponseBodyDetail {
	s.TotalStage = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetTotalTask(v int64) *DescribeSQLPlanResponseBodyDetail {
	s.TotalTask = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetTotalTime(v int64) *DescribeSQLPlanResponseBodyDetail {
	s.TotalTime = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetUser(v string) *DescribeSQLPlanResponseBodyDetail {
	s.User = &v
	return s
}

type DescribeSQLPlanResponseBodyStageList struct {
	CPUTimeAvg   *int64  `json:"CPUTimeAvg,omitempty" xml:"CPUTimeAvg,omitempty"`
	CPUTimeMax   *int64  `json:"CPUTimeMax,omitempty" xml:"CPUTimeMax,omitempty"`
	CPUTimeMin   *int64  `json:"CPUTimeMin,omitempty" xml:"CPUTimeMin,omitempty"`
	InputSizeAvg *int64  `json:"InputSizeAvg,omitempty" xml:"InputSizeAvg,omitempty"`
	InputSizeMax *int64  `json:"InputSizeMax,omitempty" xml:"InputSizeMax,omitempty"`
	InputSizeMin *int64  `json:"InputSizeMin,omitempty" xml:"InputSizeMin,omitempty"`
	OperatorCost *int64  `json:"OperatorCost,omitempty" xml:"OperatorCost,omitempty"`
	PeakMemory   *int64  `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	ScanSizeAvg  *int64  `json:"ScanSizeAvg,omitempty" xml:"ScanSizeAvg,omitempty"`
	ScanSizeMax  *int64  `json:"ScanSizeMax,omitempty" xml:"ScanSizeMax,omitempty"`
	ScanSizeMin  *int64  `json:"ScanSizeMin,omitempty" xml:"ScanSizeMin,omitempty"`
	ScanTimeAvg  *int64  `json:"ScanTimeAvg,omitempty" xml:"ScanTimeAvg,omitempty"`
	ScanTimeMax  *int64  `json:"ScanTimeMax,omitempty" xml:"ScanTimeMax,omitempty"`
	ScanTimeMin  *int64  `json:"ScanTimeMin,omitempty" xml:"ScanTimeMin,omitempty"`
	StageId      *int32  `json:"StageId,omitempty" xml:"StageId,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeSQLPlanResponseBodyStageList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlanResponseBodyStageList) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanResponseBodyStageList) SetCPUTimeAvg(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.CPUTimeAvg = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetCPUTimeMax(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.CPUTimeMax = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetCPUTimeMin(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.CPUTimeMin = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetInputSizeAvg(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.InputSizeAvg = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetInputSizeMax(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.InputSizeMax = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetInputSizeMin(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.InputSizeMin = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetOperatorCost(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.OperatorCost = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetPeakMemory(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.PeakMemory = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetScanSizeAvg(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.ScanSizeAvg = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetScanSizeMax(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.ScanSizeMax = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetScanSizeMin(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.ScanSizeMin = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetScanTimeAvg(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.ScanTimeAvg = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetScanTimeMax(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.ScanTimeMax = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetScanTimeMin(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.ScanTimeMin = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetStageId(v int32) *DescribeSQLPlanResponseBodyStageList {
	s.StageId = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetState(v string) *DescribeSQLPlanResponseBodyStageList {
	s.State = &v
	return s
}

type DescribeSQLPlanResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSQLPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQLPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlanResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanResponse) SetHeaders(v map[string]*string) *DescribeSQLPlanResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLPlanResponse) SetBody(v *DescribeSQLPlanResponseBody) *DescribeSQLPlanResponse {
	s.Body = v
	return s
}

type DescribeSQLPlanTaskRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProcessId            *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StageId              *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
}

func (s DescribeSQLPlanTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlanTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanTaskRequest) SetDBClusterId(v string) *DescribeSQLPlanTaskRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSQLPlanTaskRequest) SetOwnerAccount(v string) *DescribeSQLPlanTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSQLPlanTaskRequest) SetOwnerId(v int64) *DescribeSQLPlanTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSQLPlanTaskRequest) SetProcessId(v string) *DescribeSQLPlanTaskRequest {
	s.ProcessId = &v
	return s
}

func (s *DescribeSQLPlanTaskRequest) SetResourceOwnerAccount(v string) *DescribeSQLPlanTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSQLPlanTaskRequest) SetResourceOwnerId(v int64) *DescribeSQLPlanTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSQLPlanTaskRequest) SetStageId(v string) *DescribeSQLPlanTaskRequest {
	s.StageId = &v
	return s
}

type DescribeSQLPlanTaskResponseBody struct {
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskList  []*DescribeSQLPlanTaskResponseBodyTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
}

func (s DescribeSQLPlanTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanTaskResponseBody) SetRequestId(v string) *DescribeSQLPlanTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBody) SetTaskList(v []*DescribeSQLPlanTaskResponseBodyTaskList) *DescribeSQLPlanTaskResponseBody {
	s.TaskList = v
	return s
}

type DescribeSQLPlanTaskResponseBodyTaskList struct {
	ElapsedTime  *int64  `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	InputRows    *int64  `json:"InputRows,omitempty" xml:"InputRows,omitempty"`
	InputSize    *int64  `json:"InputSize,omitempty" xml:"InputSize,omitempty"`
	OperatorCost *int64  `json:"OperatorCost,omitempty" xml:"OperatorCost,omitempty"`
	OutputRows   *int64  `json:"OutputRows,omitempty" xml:"OutputRows,omitempty"`
	OutputSize   *int64  `json:"OutputSize,omitempty" xml:"OutputSize,omitempty"`
	PeakMemory   *int64  `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	ScanCost     *int64  `json:"ScanCost,omitempty" xml:"ScanCost,omitempty"`
	ScanRows     *int64  `json:"ScanRows,omitempty" xml:"ScanRows,omitempty"`
	ScanSize     *int64  `json:"ScanSize,omitempty" xml:"ScanSize,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
	TaskId       *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeSQLPlanTaskResponseBodyTaskList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlanTaskResponseBodyTaskList) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetElapsedTime(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.ElapsedTime = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetInputRows(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.InputRows = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetInputSize(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.InputSize = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetOperatorCost(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.OperatorCost = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetOutputRows(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.OutputRows = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetOutputSize(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.OutputSize = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetPeakMemory(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.PeakMemory = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetScanCost(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.ScanCost = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetScanRows(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.ScanRows = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetScanSize(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.ScanSize = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetState(v string) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.State = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetTaskId(v int32) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.TaskId = &v
	return s
}

type DescribeSQLPlanTaskResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSQLPlanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQLPlanTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlanTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanTaskResponse) SetHeaders(v map[string]*string) *DescribeSQLPlanTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLPlanTaskResponse) SetBody(v *DescribeSQLPlanTaskResponseBody) *DescribeSQLPlanTaskResponse {
	s.Body = v
	return s
}

type DescribeSchemasRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeSchemasRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSchemasRequest) GoString() string {
	return s.String()
}

func (s *DescribeSchemasRequest) SetDBClusterId(v string) *DescribeSchemasRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSchemasRequest) SetOwnerAccount(v string) *DescribeSchemasRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSchemasRequest) SetOwnerId(v int64) *DescribeSchemasRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSchemasRequest) SetResourceOwnerAccount(v string) *DescribeSchemasRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSchemasRequest) SetResourceOwnerId(v int64) *DescribeSchemasRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeSchemasResponseBody struct {
	Items     *DescribeSchemasResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSchemasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSchemasResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSchemasResponseBody) SetItems(v *DescribeSchemasResponseBodyItems) *DescribeSchemasResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSchemasResponseBody) SetRequestId(v string) *DescribeSchemasResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSchemasResponseBodyItems struct {
	Schema []*DescribeSchemasResponseBodyItemsSchema `json:"Schema,omitempty" xml:"Schema,omitempty" type:"Repeated"`
}

func (s DescribeSchemasResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSchemasResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSchemasResponseBodyItems) SetSchema(v []*DescribeSchemasResponseBodyItemsSchema) *DescribeSchemasResponseBodyItems {
	s.Schema = v
	return s
}

type DescribeSchemasResponseBodyItemsSchema struct {
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	SchemaName  *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
}

func (s DescribeSchemasResponseBodyItemsSchema) String() string {
	return tea.Prettify(s)
}

func (s DescribeSchemasResponseBodyItemsSchema) GoString() string {
	return s.String()
}

func (s *DescribeSchemasResponseBodyItemsSchema) SetDBClusterId(v string) *DescribeSchemasResponseBodyItemsSchema {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSchemasResponseBodyItemsSchema) SetSchemaName(v string) *DescribeSchemasResponseBodyItemsSchema {
	s.SchemaName = &v
	return s
}

type DescribeSchemasResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSchemasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSchemasResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSchemasResponse) GoString() string {
	return s.String()
}

func (s *DescribeSchemasResponse) SetHeaders(v map[string]*string) *DescribeSchemasResponse {
	s.Headers = v
	return s
}

func (s *DescribeSchemasResponse) SetBody(v *DescribeSchemasResponseBody) *DescribeSchemasResponse {
	s.Body = v
	return s
}

type DescribeSlowLogRecordsRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Order                *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProcessID            *string `json:"ProcessID,omitempty" xml:"ProcessID,omitempty"`
	Range                *string `json:"Range,omitempty" xml:"Range,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State                *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeSlowLogRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsRequest) GoString() string {
	return s.String()
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

type DescribeSlowLogRecordsResponseBody struct {
	DBClusterId *string                                  `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Items       *DescribeSlowLogRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	PageNumber  *string                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *string                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *string                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBody) SetDBClusterId(v string) *DescribeSlowLogRecordsResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetItems(v *DescribeSlowLogRecordsResponseBodyItems) *DescribeSlowLogRecordsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetPageNumber(v string) *DescribeSlowLogRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetPageSize(v string) *DescribeSlowLogRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetRequestId(v string) *DescribeSlowLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetTotalCount(v string) *DescribeSlowLogRecordsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeSlowLogRecordsResponseBodyItems struct {
	SlowLogRecord []*DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord `json:"SlowLogRecord,omitempty" xml:"SlowLogRecord,omitempty" type:"Repeated"`
}

func (s DescribeSlowLogRecordsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetSlowLogRecord(v []*DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) *DescribeSlowLogRecordsResponseBodyItems {
	s.SlowLogRecord = v
	return s
}

type DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord struct {
	DBName             *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" xml:"ExecutionStartTime,omitempty"`
	HostAddress        *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	OutputSize         *string `json:"OutputSize,omitempty" xml:"OutputSize,omitempty"`
	ParseRowCounts     *int64  `json:"ParseRowCounts,omitempty" xml:"ParseRowCounts,omitempty"`
	PeakMemoryUsage    *string `json:"PeakMemoryUsage,omitempty" xml:"PeakMemoryUsage,omitempty"`
	PlanningTime       *int64  `json:"PlanningTime,omitempty" xml:"PlanningTime,omitempty"`
	ProcessID          *string `json:"ProcessID,omitempty" xml:"ProcessID,omitempty"`
	QueryTime          *int64  `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
	QueueTime          *int64  `json:"QueueTime,omitempty" xml:"QueueTime,omitempty"`
	ReturnRowCounts    *int64  `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	SQLText            *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	ScanRows           *int64  `json:"ScanRows,omitempty" xml:"ScanRows,omitempty"`
	ScanSize           *string `json:"ScanSize,omitempty" xml:"ScanSize,omitempty"`
	ScanTime           *int64  `json:"ScanTime,omitempty" xml:"ScanTime,omitempty"`
	State              *string `json:"State,omitempty" xml:"State,omitempty"`
	UserName           *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	WallTime           *int64  `json:"WallTime,omitempty" xml:"WallTime,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetDBName(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetExecutionStartTime(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.ExecutionStartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetHostAddress(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.HostAddress = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetOutputSize(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.OutputSize = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetParseRowCounts(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.ParseRowCounts = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetPeakMemoryUsage(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.PeakMemoryUsage = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetPlanningTime(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.PlanningTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetProcessID(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.ProcessID = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetQueryTime(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.QueryTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetQueueTime(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.QueueTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetReturnRowCounts(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetSQLText(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.SQLText = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetScanRows(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.ScanRows = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetScanSize(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.ScanSize = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetScanTime(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.ScanTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetState(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.State = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetUserName(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.UserName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetWallTime(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.WallTime = &v
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

type DescribeSlowLogTrendRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSlowLogTrendRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendRequest) SetDBClusterId(v string) *DescribeSlowLogTrendRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetDBName(v string) *DescribeSlowLogTrendRequest {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetEndTime(v string) *DescribeSlowLogTrendRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetOwnerAccount(v string) *DescribeSlowLogTrendRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetOwnerId(v int64) *DescribeSlowLogTrendRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetResourceOwnerAccount(v string) *DescribeSlowLogTrendRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetResourceOwnerId(v int64) *DescribeSlowLogTrendRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetStartTime(v string) *DescribeSlowLogTrendRequest {
	s.StartTime = &v
	return s
}

type DescribeSlowLogTrendResponseBody struct {
	DBClusterId *string                                `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	EndTime     *string                                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Items       *DescribeSlowLogTrendResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime   *string                                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSlowLogTrendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBody) SetDBClusterId(v string) *DescribeSlowLogTrendResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBody) SetEndTime(v string) *DescribeSlowLogTrendResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBody) SetItems(v *DescribeSlowLogTrendResponseBodyItems) *DescribeSlowLogTrendResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSlowLogTrendResponseBody) SetRequestId(v string) *DescribeSlowLogTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBody) SetStartTime(v string) *DescribeSlowLogTrendResponseBody {
	s.StartTime = &v
	return s
}

type DescribeSlowLogTrendResponseBodyItems struct {
	SlowLogTrendItem []*DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItem `json:"SlowLogTrendItem,omitempty" xml:"SlowLogTrendItem,omitempty" type:"Repeated"`
}

func (s DescribeSlowLogTrendResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBodyItems) SetSlowLogTrendItem(v []*DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItem) *DescribeSlowLogTrendResponseBodyItems {
	s.SlowLogTrendItem = v
	return s
}

type DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItem struct {
	Key    *string                                                      `json:"Key,omitempty" xml:"Key,omitempty"`
	Series *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Struct"`
	Unit   *string                                                      `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItem) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItem) SetKey(v string) *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItem {
	s.Key = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItem) SetSeries(v *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeries) *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItem {
	s.Series = v
	return s
}

func (s *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItem) SetUnit(v string) *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItem {
	s.Unit = &v
	return s
}

type DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeries struct {
	SeriesItem []*DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeriesSeriesItem `json:"SeriesItem,omitempty" xml:"SeriesItem,omitempty" type:"Repeated"`
}

func (s DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeries) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeries) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeries) SetSeriesItem(v []*DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeriesSeriesItem) *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeries {
	s.SeriesItem = v
	return s
}

type DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeriesSeriesItem struct {
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeriesSeriesItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeriesSeriesItem) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeriesSeriesItem) SetName(v string) *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeriesSeriesItem {
	s.Name = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeriesSeriesItem) SetValues(v string) *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeriesSeriesItem {
	s.Values = &v
	return s
}

type DescribeSlowLogTrendResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSlowLogTrendResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSlowLogTrendResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponse) SetHeaders(v map[string]*string) *DescribeSlowLogTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlowLogTrendResponse) SetBody(v *DescribeSlowLogTrendResponseBody) *DescribeSlowLogTrendResponse {
	s.Body = v
	return s
}

type DescribeSqlPatternRequest struct {
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Order       *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SqlPattern  *string `json:"SqlPattern,omitempty" xml:"SqlPattern,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSqlPatternRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlPatternRequest) GoString() string {
	return s.String()
}

func (s *DescribeSqlPatternRequest) SetDBClusterId(v string) *DescribeSqlPatternRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSqlPatternRequest) SetOrder(v string) *DescribeSqlPatternRequest {
	s.Order = &v
	return s
}

func (s *DescribeSqlPatternRequest) SetPageNumber(v int32) *DescribeSqlPatternRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSqlPatternRequest) SetPageSize(v int32) *DescribeSqlPatternRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSqlPatternRequest) SetRegionId(v string) *DescribeSqlPatternRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSqlPatternRequest) SetSqlPattern(v string) *DescribeSqlPatternRequest {
	s.SqlPattern = &v
	return s
}

func (s *DescribeSqlPatternRequest) SetStartTime(v string) *DescribeSqlPatternRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSqlPatternRequest) SetType(v string) *DescribeSqlPatternRequest {
	s.Type = &v
	return s
}

type DescribeSqlPatternResponseBody struct {
	Items      []*DescribeSqlPatternResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSqlPatternResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlPatternResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSqlPatternResponseBody) SetItems(v []*DescribeSqlPatternResponseBodyItems) *DescribeSqlPatternResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSqlPatternResponseBody) SetPageNumber(v int32) *DescribeSqlPatternResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSqlPatternResponseBody) SetPageSize(v int32) *DescribeSqlPatternResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSqlPatternResponseBody) SetRequestId(v string) *DescribeSqlPatternResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSqlPatternResponseBody) SetTotalCount(v int32) *DescribeSqlPatternResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeSqlPatternResponseBodyItems struct {
	AccessIP      *string `json:"AccessIP,omitempty" xml:"AccessIP,omitempty"`
	AvgCpuTime    *string `json:"AvgCpuTime,omitempty" xml:"AvgCpuTime,omitempty"`
	AvgPeakMemory *string `json:"AvgPeakMemory,omitempty" xml:"AvgPeakMemory,omitempty"`
	AvgScanSize   *string `json:"AvgScanSize,omitempty" xml:"AvgScanSize,omitempty"`
	AvgStageCount *string `json:"AvgStageCount,omitempty" xml:"AvgStageCount,omitempty"`
	AvgTaskCount  *string `json:"AvgTaskCount,omitempty" xml:"AvgTaskCount,omitempty"`
	InstanceName  *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	MaxCpuTime    *string `json:"MaxCpuTime,omitempty" xml:"MaxCpuTime,omitempty"`
	MaxPeakMemory *string `json:"MaxPeakMemory,omitempty" xml:"MaxPeakMemory,omitempty"`
	MaxScanSize   *string `json:"MaxScanSize,omitempty" xml:"MaxScanSize,omitempty"`
	MaxStageCount *string `json:"MaxStageCount,omitempty" xml:"MaxStageCount,omitempty"`
	MaxTaskCount  *string `json:"MaxTaskCount,omitempty" xml:"MaxTaskCount,omitempty"`
	Pattern       *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	QueryCount    *string `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
	ReportDate    *string `json:"ReportDate,omitempty" xml:"ReportDate,omitempty"`
	User          *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeSqlPatternResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlPatternResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSqlPatternResponseBodyItems) SetAccessIP(v string) *DescribeSqlPatternResponseBodyItems {
	s.AccessIP = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetAvgCpuTime(v string) *DescribeSqlPatternResponseBodyItems {
	s.AvgCpuTime = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetAvgPeakMemory(v string) *DescribeSqlPatternResponseBodyItems {
	s.AvgPeakMemory = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetAvgScanSize(v string) *DescribeSqlPatternResponseBodyItems {
	s.AvgScanSize = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetAvgStageCount(v string) *DescribeSqlPatternResponseBodyItems {
	s.AvgStageCount = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetAvgTaskCount(v string) *DescribeSqlPatternResponseBodyItems {
	s.AvgTaskCount = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetInstanceName(v string) *DescribeSqlPatternResponseBodyItems {
	s.InstanceName = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetMaxCpuTime(v string) *DescribeSqlPatternResponseBodyItems {
	s.MaxCpuTime = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetMaxPeakMemory(v string) *DescribeSqlPatternResponseBodyItems {
	s.MaxPeakMemory = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetMaxScanSize(v string) *DescribeSqlPatternResponseBodyItems {
	s.MaxScanSize = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetMaxStageCount(v string) *DescribeSqlPatternResponseBodyItems {
	s.MaxStageCount = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetMaxTaskCount(v string) *DescribeSqlPatternResponseBodyItems {
	s.MaxTaskCount = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetPattern(v string) *DescribeSqlPatternResponseBodyItems {
	s.Pattern = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetQueryCount(v string) *DescribeSqlPatternResponseBodyItems {
	s.QueryCount = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetReportDate(v string) *DescribeSqlPatternResponseBodyItems {
	s.ReportDate = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetUser(v string) *DescribeSqlPatternResponseBodyItems {
	s.User = &v
	return s
}

type DescribeSqlPatternResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSqlPatternResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSqlPatternResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlPatternResponse) GoString() string {
	return s.String()
}

func (s *DescribeSqlPatternResponse) SetHeaders(v map[string]*string) *DescribeSqlPatternResponse {
	s.Headers = v
	return s
}

func (s *DescribeSqlPatternResponse) SetBody(v *DescribeSqlPatternResponseBody) *DescribeSqlPatternResponse {
	s.Body = v
	return s
}

type DescribeTableAccessCountRequest struct {
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Order       *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TableName   *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeTableAccessCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableAccessCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeTableAccessCountRequest) SetDBClusterId(v string) *DescribeTableAccessCountRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTableAccessCountRequest) SetOrder(v string) *DescribeTableAccessCountRequest {
	s.Order = &v
	return s
}

func (s *DescribeTableAccessCountRequest) SetPageNumber(v int32) *DescribeTableAccessCountRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTableAccessCountRequest) SetPageSize(v int32) *DescribeTableAccessCountRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTableAccessCountRequest) SetRegionId(v string) *DescribeTableAccessCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTableAccessCountRequest) SetStartTime(v string) *DescribeTableAccessCountRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTableAccessCountRequest) SetTableName(v string) *DescribeTableAccessCountRequest {
	s.TableName = &v
	return s
}

type DescribeTableAccessCountResponseBody struct {
	Items      []*DescribeTableAccessCountResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTableAccessCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableAccessCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTableAccessCountResponseBody) SetItems(v []*DescribeTableAccessCountResponseBodyItems) *DescribeTableAccessCountResponseBody {
	s.Items = v
	return s
}

func (s *DescribeTableAccessCountResponseBody) SetPageNumber(v int32) *DescribeTableAccessCountResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTableAccessCountResponseBody) SetPageSize(v int32) *DescribeTableAccessCountResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTableAccessCountResponseBody) SetRequestId(v string) *DescribeTableAccessCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTableAccessCountResponseBody) SetTotalCount(v int32) *DescribeTableAccessCountResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeTableAccessCountResponseBodyItems struct {
	AccessCount  *string `json:"AccessCount,omitempty" xml:"AccessCount,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ReportDate   *string `json:"ReportDate,omitempty" xml:"ReportDate,omitempty"`
	TableName    *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TableSchema  *string `json:"TableSchema,omitempty" xml:"TableSchema,omitempty"`
}

func (s DescribeTableAccessCountResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableAccessCountResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeTableAccessCountResponseBodyItems) SetAccessCount(v string) *DescribeTableAccessCountResponseBodyItems {
	s.AccessCount = &v
	return s
}

func (s *DescribeTableAccessCountResponseBodyItems) SetInstanceName(v string) *DescribeTableAccessCountResponseBodyItems {
	s.InstanceName = &v
	return s
}

func (s *DescribeTableAccessCountResponseBodyItems) SetReportDate(v string) *DescribeTableAccessCountResponseBodyItems {
	s.ReportDate = &v
	return s
}

func (s *DescribeTableAccessCountResponseBodyItems) SetTableName(v string) *DescribeTableAccessCountResponseBodyItems {
	s.TableName = &v
	return s
}

func (s *DescribeTableAccessCountResponseBodyItems) SetTableSchema(v string) *DescribeTableAccessCountResponseBodyItems {
	s.TableSchema = &v
	return s
}

type DescribeTableAccessCountResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTableAccessCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTableAccessCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableAccessCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeTableAccessCountResponse) SetHeaders(v map[string]*string) *DescribeTableAccessCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeTableAccessCountResponse) SetBody(v *DescribeTableAccessCountResponseBody) *DescribeTableAccessCountResponse {
	s.Body = v
	return s
}

type DescribeTableDetailRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SchemaName           *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName            *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeTableDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeTableDetailRequest) SetDBClusterId(v string) *DescribeTableDetailRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTableDetailRequest) SetOwnerAccount(v string) *DescribeTableDetailRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTableDetailRequest) SetOwnerId(v int64) *DescribeTableDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTableDetailRequest) SetResourceOwnerAccount(v string) *DescribeTableDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTableDetailRequest) SetResourceOwnerId(v int64) *DescribeTableDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTableDetailRequest) SetSchemaName(v string) *DescribeTableDetailRequest {
	s.SchemaName = &v
	return s
}

func (s *DescribeTableDetailRequest) SetTableName(v string) *DescribeTableDetailRequest {
	s.TableName = &v
	return s
}

type DescribeTableDetailResponseBody struct {
	AvgSize   *int64                                `json:"AvgSize,omitempty" xml:"AvgSize,omitempty"`
	Items     *DescribeTableDetailResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTableDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTableDetailResponseBody) SetAvgSize(v int64) *DescribeTableDetailResponseBody {
	s.AvgSize = &v
	return s
}

func (s *DescribeTableDetailResponseBody) SetItems(v *DescribeTableDetailResponseBodyItems) *DescribeTableDetailResponseBody {
	s.Items = v
	return s
}

func (s *DescribeTableDetailResponseBody) SetRequestId(v string) *DescribeTableDetailResponseBody {
	s.RequestId = &v
	return s
}

type DescribeTableDetailResponseBodyItems struct {
	Shard []*DescribeTableDetailResponseBodyItemsShard `json:"Shard,omitempty" xml:"Shard,omitempty" type:"Repeated"`
}

func (s DescribeTableDetailResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableDetailResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeTableDetailResponseBodyItems) SetShard(v []*DescribeTableDetailResponseBodyItemsShard) *DescribeTableDetailResponseBodyItems {
	s.Shard = v
	return s
}

type DescribeTableDetailResponseBodyItemsShard struct {
	Id   *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeTableDetailResponseBodyItemsShard) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableDetailResponseBodyItemsShard) GoString() string {
	return s.String()
}

func (s *DescribeTableDetailResponseBodyItemsShard) SetId(v int32) *DescribeTableDetailResponseBodyItemsShard {
	s.Id = &v
	return s
}

func (s *DescribeTableDetailResponseBodyItemsShard) SetSize(v int64) *DescribeTableDetailResponseBodyItemsShard {
	s.Size = &v
	return s
}

type DescribeTableDetailResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTableDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTableDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeTableDetailResponse) SetHeaders(v map[string]*string) *DescribeTableDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeTableDetailResponse) SetBody(v *DescribeTableDetailResponseBody) *DescribeTableDetailResponse {
	s.Body = v
	return s
}

type DescribeTablePartitionDiagnoseRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeTablePartitionDiagnoseRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablePartitionDiagnoseRequest) GoString() string {
	return s.String()
}

func (s *DescribeTablePartitionDiagnoseRequest) SetDBClusterId(v string) *DescribeTablePartitionDiagnoseRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseRequest) SetOwnerAccount(v string) *DescribeTablePartitionDiagnoseRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseRequest) SetOwnerId(v int64) *DescribeTablePartitionDiagnoseRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseRequest) SetPageNumber(v int32) *DescribeTablePartitionDiagnoseRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseRequest) SetPageSize(v int32) *DescribeTablePartitionDiagnoseRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseRequest) SetRegionId(v string) *DescribeTablePartitionDiagnoseRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseRequest) SetResourceOwnerAccount(v string) *DescribeTablePartitionDiagnoseRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseRequest) SetResourceOwnerId(v int64) *DescribeTablePartitionDiagnoseRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeTablePartitionDiagnoseResponseBody struct {
	DBClusterId                   *string                                            `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Items                         []*DescribeTablePartitionDiagnoseResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber                    *int32                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                      *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId                     *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuggestMaxRecordsPerPartition *int64                                             `json:"SuggestMaxRecordsPerPartition,omitempty" xml:"SuggestMaxRecordsPerPartition,omitempty"`
	SuggestMinRecordsPerPartition *int64                                             `json:"SuggestMinRecordsPerPartition,omitempty" xml:"SuggestMinRecordsPerPartition,omitempty"`
	TotalCount                    *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTablePartitionDiagnoseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablePartitionDiagnoseResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetDBClusterId(v string) *DescribeTablePartitionDiagnoseResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetItems(v []*DescribeTablePartitionDiagnoseResponseBodyItems) *DescribeTablePartitionDiagnoseResponseBody {
	s.Items = v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetPageNumber(v int32) *DescribeTablePartitionDiagnoseResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetPageSize(v int32) *DescribeTablePartitionDiagnoseResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetRequestId(v string) *DescribeTablePartitionDiagnoseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetSuggestMaxRecordsPerPartition(v int64) *DescribeTablePartitionDiagnoseResponseBody {
	s.SuggestMaxRecordsPerPartition = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetSuggestMinRecordsPerPartition(v int64) *DescribeTablePartitionDiagnoseResponseBody {
	s.SuggestMinRecordsPerPartition = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetTotalCount(v int32) *DescribeTablePartitionDiagnoseResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeTablePartitionDiagnoseResponseBodyItems struct {
	PartitionDetail *string `json:"PartitionDetail,omitempty" xml:"PartitionDetail,omitempty"`
	PartitionNumber *int32  `json:"PartitionNumber,omitempty" xml:"PartitionNumber,omitempty"`
	SchemaName      *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName       *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeTablePartitionDiagnoseResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablePartitionDiagnoseResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeTablePartitionDiagnoseResponseBodyItems) SetPartitionDetail(v string) *DescribeTablePartitionDiagnoseResponseBodyItems {
	s.PartitionDetail = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBodyItems) SetPartitionNumber(v int32) *DescribeTablePartitionDiagnoseResponseBodyItems {
	s.PartitionNumber = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBodyItems) SetSchemaName(v string) *DescribeTablePartitionDiagnoseResponseBodyItems {
	s.SchemaName = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBodyItems) SetTableName(v string) *DescribeTablePartitionDiagnoseResponseBodyItems {
	s.TableName = &v
	return s
}

type DescribeTablePartitionDiagnoseResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTablePartitionDiagnoseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTablePartitionDiagnoseResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablePartitionDiagnoseResponse) GoString() string {
	return s.String()
}

func (s *DescribeTablePartitionDiagnoseResponse) SetHeaders(v map[string]*string) *DescribeTablePartitionDiagnoseResponse {
	s.Headers = v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponse) SetBody(v *DescribeTablePartitionDiagnoseResponseBody) *DescribeTablePartitionDiagnoseResponse {
	s.Body = v
	return s
}

type DescribeTableStatisticsRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Order                *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeTableStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTableStatisticsRequest) SetDBClusterId(v string) *DescribeTableStatisticsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTableStatisticsRequest) SetOrder(v string) *DescribeTableStatisticsRequest {
	s.Order = &v
	return s
}

func (s *DescribeTableStatisticsRequest) SetOwnerAccount(v string) *DescribeTableStatisticsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTableStatisticsRequest) SetOwnerId(v int64) *DescribeTableStatisticsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTableStatisticsRequest) SetPageNumber(v int32) *DescribeTableStatisticsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTableStatisticsRequest) SetPageSize(v int32) *DescribeTableStatisticsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTableStatisticsRequest) SetResourceOwnerAccount(v string) *DescribeTableStatisticsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTableStatisticsRequest) SetResourceOwnerId(v int64) *DescribeTableStatisticsRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeTableStatisticsResponseBody struct {
	DBClusterId *string                                   `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Items       *DescribeTableStatisticsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	PageNumber  *string                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *string                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *string                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTableStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTableStatisticsResponseBody) SetDBClusterId(v string) *DescribeTableStatisticsResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTableStatisticsResponseBody) SetItems(v *DescribeTableStatisticsResponseBodyItems) *DescribeTableStatisticsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeTableStatisticsResponseBody) SetPageNumber(v string) *DescribeTableStatisticsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTableStatisticsResponseBody) SetPageSize(v string) *DescribeTableStatisticsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTableStatisticsResponseBody) SetRequestId(v string) *DescribeTableStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTableStatisticsResponseBody) SetTotalCount(v string) *DescribeTableStatisticsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeTableStatisticsResponseBodyItems struct {
	TableStatisticRecords []*DescribeTableStatisticsResponseBodyItemsTableStatisticRecords `json:"TableStatisticRecords,omitempty" xml:"TableStatisticRecords,omitempty" type:"Repeated"`
}

func (s DescribeTableStatisticsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableStatisticsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeTableStatisticsResponseBodyItems) SetTableStatisticRecords(v []*DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) *DescribeTableStatisticsResponseBodyItems {
	s.TableStatisticRecords = v
	return s
}

type DescribeTableStatisticsResponseBodyItemsTableStatisticRecords struct {
	ColdDataSize        *int64  `json:"ColdDataSize,omitempty" xml:"ColdDataSize,omitempty"`
	DataSize            *int64  `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	IndexSize           *int64  `json:"IndexSize,omitempty" xml:"IndexSize,omitempty"`
	PartitionCount      *int64  `json:"PartitionCount,omitempty" xml:"PartitionCount,omitempty"`
	PrimaryKeyIndexSize *int64  `json:"PrimaryKeyIndexSize,omitempty" xml:"PrimaryKeyIndexSize,omitempty"`
	RowCount            *int64  `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	SchemaName          *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName           *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) GoString() string {
	return s.String()
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetColdDataSize(v int64) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.ColdDataSize = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetDataSize(v int64) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.DataSize = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetIndexSize(v int64) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.IndexSize = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetPartitionCount(v int64) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.PartitionCount = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetPrimaryKeyIndexSize(v int64) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.PrimaryKeyIndexSize = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetRowCount(v int64) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.RowCount = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetSchemaName(v string) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.SchemaName = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetTableName(v string) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.TableName = &v
	return s
}

type DescribeTableStatisticsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTableStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTableStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeTableStatisticsResponse) SetHeaders(v map[string]*string) *DescribeTableStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeTableStatisticsResponse) SetBody(v *DescribeTableStatisticsResponseBody) *DescribeTableStatisticsResponse {
	s.Body = v
	return s
}

type DescribeTablesRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SchemaName           *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
}

func (s DescribeTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTablesRequest) SetDBClusterId(v string) *DescribeTablesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTablesRequest) SetOwnerAccount(v string) *DescribeTablesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTablesRequest) SetOwnerId(v int64) *DescribeTablesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTablesRequest) SetResourceOwnerAccount(v string) *DescribeTablesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTablesRequest) SetResourceOwnerId(v int64) *DescribeTablesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTablesRequest) SetSchemaName(v string) *DescribeTablesRequest {
	s.SchemaName = &v
	return s
}

type DescribeTablesResponseBody struct {
	Items     *DescribeTablesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponseBody) SetItems(v *DescribeTablesResponseBodyItems) *DescribeTablesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeTablesResponseBody) SetRequestId(v string) *DescribeTablesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeTablesResponseBodyItems struct {
	Table []*DescribeTablesResponseBodyItemsTable `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeTablesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponseBodyItems) SetTable(v []*DescribeTablesResponseBodyItemsTable) *DescribeTablesResponseBodyItems {
	s.Table = v
	return s
}

type DescribeTablesResponseBodyItemsTable struct {
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	SchemaName  *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName   *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeTablesResponseBodyItemsTable) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesResponseBodyItemsTable) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponseBodyItemsTable) SetDBClusterId(v string) *DescribeTablesResponseBodyItemsTable {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTablesResponseBodyItemsTable) SetSchemaName(v string) *DescribeTablesResponseBodyItemsTable {
	s.SchemaName = &v
	return s
}

func (s *DescribeTablesResponseBodyItemsTable) SetTableName(v string) *DescribeTablesResponseBodyItemsTable {
	s.TableName = &v
	return s
}

type DescribeTablesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponse) SetHeaders(v map[string]*string) *DescribeTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTablesResponse) SetBody(v *DescribeTablesResponseBody) *DescribeTablesResponse {
	s.Body = v
	return s
}

type DescribeTaskInfoRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskId               *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeTaskInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeTaskInfoRequest) SetDBClusterId(v string) *DescribeTaskInfoRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTaskInfoRequest) SetOwnerAccount(v string) *DescribeTaskInfoRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTaskInfoRequest) SetOwnerId(v int64) *DescribeTaskInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTaskInfoRequest) SetResourceOwnerAccount(v string) *DescribeTaskInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTaskInfoRequest) SetResourceOwnerId(v int64) *DescribeTaskInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTaskInfoRequest) SetTaskId(v int32) *DescribeTaskInfoRequest {
	s.TaskId = &v
	return s
}

type DescribeTaskInfoResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskInfo  *DescribeTaskInfoResponseBodyTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Struct"`
}

func (s DescribeTaskInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTaskInfoResponseBody) SetRequestId(v string) *DescribeTaskInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTaskInfoResponseBody) SetTaskInfo(v *DescribeTaskInfoResponseBodyTaskInfo) *DescribeTaskInfoResponseBody {
	s.TaskInfo = v
	return s
}

type DescribeTaskInfoResponseBodyTaskInfo struct {
	BeginTime  *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	Progress   *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId     *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeTaskInfoResponseBodyTaskInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskInfoResponseBodyTaskInfo) GoString() string {
	return s.String()
}

func (s *DescribeTaskInfoResponseBodyTaskInfo) SetBeginTime(v string) *DescribeTaskInfoResponseBodyTaskInfo {
	s.BeginTime = &v
	return s
}

func (s *DescribeTaskInfoResponseBodyTaskInfo) SetFinishTime(v string) *DescribeTaskInfoResponseBodyTaskInfo {
	s.FinishTime = &v
	return s
}

func (s *DescribeTaskInfoResponseBodyTaskInfo) SetProgress(v string) *DescribeTaskInfoResponseBodyTaskInfo {
	s.Progress = &v
	return s
}

func (s *DescribeTaskInfoResponseBodyTaskInfo) SetStatus(v string) *DescribeTaskInfoResponseBodyTaskInfo {
	s.Status = &v
	return s
}

func (s *DescribeTaskInfoResponseBodyTaskInfo) SetTaskId(v int32) *DescribeTaskInfoResponseBodyTaskInfo {
	s.TaskId = &v
	return s
}

type DescribeTaskInfoResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTaskInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeTaskInfoResponse) SetHeaders(v map[string]*string) *DescribeTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeTaskInfoResponse) SetBody(v *DescribeTaskInfoResponseBody) *DescribeTaskInfoResponse {
	s.Body = v
	return s
}

type DownloadDiagnosisRecordsRequest struct {
	ClientIp       *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	DBClusterId    *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Database       *string `json:"Database,omitempty" xml:"Database,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Keyword        *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	MaxPeakMemory  *int64  `json:"MaxPeakMemory,omitempty" xml:"MaxPeakMemory,omitempty"`
	MaxScanSize    *int64  `json:"MaxScanSize,omitempty" xml:"MaxScanSize,omitempty"`
	MinPeakMemory  *int64  `json:"MinPeakMemory,omitempty" xml:"MinPeakMemory,omitempty"`
	MinScanSize    *int64  `json:"MinScanSize,omitempty" xml:"MinScanSize,omitempty"`
	QueryCondition *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroup  *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UserName       *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DownloadDiagnosisRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DownloadDiagnosisRecordsRequest) GoString() string {
	return s.String()
}

func (s *DownloadDiagnosisRecordsRequest) SetClientIp(v string) *DownloadDiagnosisRecordsRequest {
	s.ClientIp = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetDBClusterId(v string) *DownloadDiagnosisRecordsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetDatabase(v string) *DownloadDiagnosisRecordsRequest {
	s.Database = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetEndTime(v string) *DownloadDiagnosisRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetKeyword(v string) *DownloadDiagnosisRecordsRequest {
	s.Keyword = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetLang(v string) *DownloadDiagnosisRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetMaxPeakMemory(v int64) *DownloadDiagnosisRecordsRequest {
	s.MaxPeakMemory = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetMaxScanSize(v int64) *DownloadDiagnosisRecordsRequest {
	s.MaxScanSize = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetMinPeakMemory(v int64) *DownloadDiagnosisRecordsRequest {
	s.MinPeakMemory = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetMinScanSize(v int64) *DownloadDiagnosisRecordsRequest {
	s.MinScanSize = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetQueryCondition(v string) *DownloadDiagnosisRecordsRequest {
	s.QueryCondition = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetRegionId(v string) *DownloadDiagnosisRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetResourceGroup(v string) *DownloadDiagnosisRecordsRequest {
	s.ResourceGroup = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetStartTime(v string) *DownloadDiagnosisRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetUserName(v string) *DownloadDiagnosisRecordsRequest {
	s.UserName = &v
	return s
}

type DownloadDiagnosisRecordsResponseBody struct {
	DownloadId *int32  `json:"DownloadId,omitempty" xml:"DownloadId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DownloadDiagnosisRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DownloadDiagnosisRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadDiagnosisRecordsResponseBody) SetDownloadId(v int32) *DownloadDiagnosisRecordsResponseBody {
	s.DownloadId = &v
	return s
}

func (s *DownloadDiagnosisRecordsResponseBody) SetRequestId(v string) *DownloadDiagnosisRecordsResponseBody {
	s.RequestId = &v
	return s
}

type DownloadDiagnosisRecordsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DownloadDiagnosisRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DownloadDiagnosisRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DownloadDiagnosisRecordsResponse) GoString() string {
	return s.String()
}

func (s *DownloadDiagnosisRecordsResponse) SetHeaders(v map[string]*string) *DownloadDiagnosisRecordsResponse {
	s.Headers = v
	return s
}

func (s *DownloadDiagnosisRecordsResponse) SetBody(v *DownloadDiagnosisRecordsResponseBody) *DownloadDiagnosisRecordsResponse {
	s.Body = v
	return s
}

type GrantOperatorPermissionRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ExpiredTime          *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Privileges           *string `json:"Privileges,omitempty" xml:"Privileges,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GrantOperatorPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantOperatorPermissionRequest) GoString() string {
	return s.String()
}

func (s *GrantOperatorPermissionRequest) SetDBClusterId(v string) *GrantOperatorPermissionRequest {
	s.DBClusterId = &v
	return s
}

func (s *GrantOperatorPermissionRequest) SetExpiredTime(v string) *GrantOperatorPermissionRequest {
	s.ExpiredTime = &v
	return s
}

func (s *GrantOperatorPermissionRequest) SetOwnerAccount(v string) *GrantOperatorPermissionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GrantOperatorPermissionRequest) SetOwnerId(v int64) *GrantOperatorPermissionRequest {
	s.OwnerId = &v
	return s
}

func (s *GrantOperatorPermissionRequest) SetPrivileges(v string) *GrantOperatorPermissionRequest {
	s.Privileges = &v
	return s
}

func (s *GrantOperatorPermissionRequest) SetResourceOwnerAccount(v string) *GrantOperatorPermissionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GrantOperatorPermissionRequest) SetResourceOwnerId(v int64) *GrantOperatorPermissionRequest {
	s.ResourceOwnerId = &v
	return s
}

type GrantOperatorPermissionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GrantOperatorPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantOperatorPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *GrantOperatorPermissionResponseBody) SetRequestId(v string) *GrantOperatorPermissionResponseBody {
	s.RequestId = &v
	return s
}

type GrantOperatorPermissionResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GrantOperatorPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GrantOperatorPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantOperatorPermissionResponse) GoString() string {
	return s.String()
}

func (s *GrantOperatorPermissionResponse) SetHeaders(v map[string]*string) *GrantOperatorPermissionResponse {
	s.Headers = v
	return s
}

func (s *GrantOperatorPermissionResponse) SetBody(v *GrantOperatorPermissionResponseBody) *GrantOperatorPermissionResponse {
	s.Body = v
	return s
}

type KillProcessRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProcessId            *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s KillProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s KillProcessRequest) GoString() string {
	return s.String()
}

func (s *KillProcessRequest) SetDBClusterId(v string) *KillProcessRequest {
	s.DBClusterId = &v
	return s
}

func (s *KillProcessRequest) SetOwnerAccount(v string) *KillProcessRequest {
	s.OwnerAccount = &v
	return s
}

func (s *KillProcessRequest) SetOwnerId(v int64) *KillProcessRequest {
	s.OwnerId = &v
	return s
}

func (s *KillProcessRequest) SetProcessId(v string) *KillProcessRequest {
	s.ProcessId = &v
	return s
}

func (s *KillProcessRequest) SetResourceOwnerAccount(v string) *KillProcessRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *KillProcessRequest) SetResourceOwnerId(v int64) *KillProcessRequest {
	s.ResourceOwnerId = &v
	return s
}

type KillProcessResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s KillProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s KillProcessResponseBody) GoString() string {
	return s.String()
}

func (s *KillProcessResponseBody) SetRequestId(v string) *KillProcessResponseBody {
	s.RequestId = &v
	return s
}

type KillProcessResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *KillProcessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s KillProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s KillProcessResponse) GoString() string {
	return s.String()
}

func (s *KillProcessResponse) SetHeaders(v map[string]*string) *KillProcessResponse {
	s.Headers = v
	return s
}

func (s *KillProcessResponse) SetBody(v *KillProcessResponseBody) *KillProcessResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	NextToken            *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string                       `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                        `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId           []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string                       `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                        `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag                  []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerAccount(v string) *ListTagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerId(v int64) *ListTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
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

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
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
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
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

type ModifyAccountDescriptionRequest struct {
	AccountDescription   *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyAccountDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionRequest) SetAccountDescription(v string) *ModifyAccountDescriptionRequest {
	s.AccountDescription = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetAccountName(v string) *ModifyAccountDescriptionRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetDBClusterId(v string) *ModifyAccountDescriptionRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetOwnerAccount(v string) *ModifyAccountDescriptionRequest {
	s.OwnerAccount = &v
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

type ModifyAuditLogConfigRequest struct {
	AuditLogStatus       *string `json:"AuditLogStatus,omitempty" xml:"AuditLogStatus,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyAuditLogConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAuditLogConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyAuditLogConfigRequest) SetAuditLogStatus(v string) *ModifyAuditLogConfigRequest {
	s.AuditLogStatus = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetDBClusterId(v string) *ModifyAuditLogConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetOwnerAccount(v string) *ModifyAuditLogConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetOwnerId(v int64) *ModifyAuditLogConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetRegionId(v string) *ModifyAuditLogConfigRequest {
	s.RegionId = &v
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

type ModifyAuditLogConfigResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UpdateSucceed *bool   `json:"UpdateSucceed,omitempty" xml:"UpdateSucceed,omitempty"`
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

func (s *ModifyAuditLogConfigResponseBody) SetUpdateSucceed(v bool) *ModifyAuditLogConfigResponseBody {
	s.UpdateSucceed = &v
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

type ModifyAutoRenewAttributeRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Duration             *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PeriodUnit           *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RenewalStatus        *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyAutoRenewAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyAutoRenewAttributeRequest) SetDBClusterId(v string) *ModifyAutoRenewAttributeRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetDuration(v string) *ModifyAutoRenewAttributeRequest {
	s.Duration = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetOwnerAccount(v string) *ModifyAutoRenewAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetOwnerId(v int64) *ModifyAutoRenewAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetPeriodUnit(v string) *ModifyAutoRenewAttributeRequest {
	s.PeriodUnit = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetRegionId(v string) *ModifyAutoRenewAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetRenewalStatus(v string) *ModifyAutoRenewAttributeRequest {
	s.RenewalStatus = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetResourceOwnerAccount(v string) *ModifyAutoRenewAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetResourceOwnerId(v int64) *ModifyAutoRenewAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyAutoRenewAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAutoRenewAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAutoRenewAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAutoRenewAttributeResponseBody) SetRequestId(v string) *ModifyAutoRenewAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAutoRenewAttributeResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyAutoRenewAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAutoRenewAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAutoRenewAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyAutoRenewAttributeResponse) SetHeaders(v map[string]*string) *ModifyAutoRenewAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyAutoRenewAttributeResponse) SetBody(v *ModifyAutoRenewAttributeResponseBody) *ModifyAutoRenewAttributeResponse {
	s.Body = v
	return s
}

type ModifyBackupPolicyRequest struct {
	BackupRetentionPeriod    *string `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	DBClusterId              *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	EnableBackupLog          *string `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	LogBackupRetentionPeriod *int32  `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	OwnerAccount             *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PreferredBackupPeriod    *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	PreferredBackupTime      *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	ResourceOwnerAccount     *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId          *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequest) SetBackupRetentionPeriod(v string) *ModifyBackupPolicyRequest {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetDBClusterId(v string) *ModifyBackupPolicyRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetEnableBackupLog(v string) *ModifyBackupPolicyRequest {
	s.EnableBackupLog = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetLogBackupRetentionPeriod(v int32) *ModifyBackupPolicyRequest {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetOwnerAccount(v string) *ModifyBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetOwnerId(v int64) *ModifyBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupPeriod(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupTime(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupTime = &v
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

type ModifyClusterConnectionStringRequest struct {
	ConnectionStringPrefix  *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
	DBClusterId             *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount            *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                 *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Port                    *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	ResourceOwnerAccount    *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId         *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyClusterConnectionStringRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterConnectionStringRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterConnectionStringRequest) SetConnectionStringPrefix(v string) *ModifyClusterConnectionStringRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetCurrentConnectionString(v string) *ModifyClusterConnectionStringRequest {
	s.CurrentConnectionString = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetDBClusterId(v string) *ModifyClusterConnectionStringRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetOwnerAccount(v string) *ModifyClusterConnectionStringRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetOwnerId(v int64) *ModifyClusterConnectionStringRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetPort(v int32) *ModifyClusterConnectionStringRequest {
	s.Port = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetResourceOwnerAccount(v string) *ModifyClusterConnectionStringRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetResourceOwnerId(v int64) *ModifyClusterConnectionStringRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyClusterConnectionStringResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClusterConnectionStringResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterConnectionStringResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterConnectionStringResponseBody) SetRequestId(v string) *ModifyClusterConnectionStringResponseBody {
	s.RequestId = &v
	return s
}

type ModifyClusterConnectionStringResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyClusterConnectionStringResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyClusterConnectionStringResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterConnectionStringResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterConnectionStringResponse) SetHeaders(v map[string]*string) *ModifyClusterConnectionStringResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterConnectionStringResponse) SetBody(v *ModifyClusterConnectionStringResponseBody) *ModifyClusterConnectionStringResponse {
	s.Body = v
	return s
}

type ModifyDBClusterRequest struct {
	ComputeResource      *string `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	DBClusterCategory    *string `json:"DBClusterCategory,omitempty" xml:"DBClusterCategory,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBNodeClass          *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	DBNodeGroupCount     *string `json:"DBNodeGroupCount,omitempty" xml:"DBNodeGroupCount,omitempty"`
	DBNodeStorage        *string `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	ElasticIOResource    *int32  `json:"ElasticIOResource,omitempty" xml:"ElasticIOResource,omitempty"`
	ExecutorCount        *string `json:"ExecutorCount,omitempty" xml:"ExecutorCount,omitempty"`
	Mode                 *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	ModifyType           *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StorageResource      *string `json:"StorageResource,omitempty" xml:"StorageResource,omitempty"`
}

func (s ModifyDBClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterRequest) SetComputeResource(v string) *ModifyDBClusterRequest {
	s.ComputeResource = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBClusterCategory(v string) *ModifyDBClusterRequest {
	s.DBClusterCategory = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBClusterId(v string) *ModifyDBClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBNodeClass(v string) *ModifyDBClusterRequest {
	s.DBNodeClass = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBNodeGroupCount(v string) *ModifyDBClusterRequest {
	s.DBNodeGroupCount = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBNodeStorage(v string) *ModifyDBClusterRequest {
	s.DBNodeStorage = &v
	return s
}

func (s *ModifyDBClusterRequest) SetElasticIOResource(v int32) *ModifyDBClusterRequest {
	s.ElasticIOResource = &v
	return s
}

func (s *ModifyDBClusterRequest) SetExecutorCount(v string) *ModifyDBClusterRequest {
	s.ExecutorCount = &v
	return s
}

func (s *ModifyDBClusterRequest) SetMode(v string) *ModifyDBClusterRequest {
	s.Mode = &v
	return s
}

func (s *ModifyDBClusterRequest) SetModifyType(v string) *ModifyDBClusterRequest {
	s.ModifyType = &v
	return s
}

func (s *ModifyDBClusterRequest) SetOwnerAccount(v string) *ModifyDBClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterRequest) SetOwnerId(v int64) *ModifyDBClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterRequest) SetRegionId(v string) *ModifyDBClusterRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBClusterRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterRequest) SetResourceOwnerId(v int64) *ModifyDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterRequest) SetStorageResource(v string) *ModifyDBClusterRequest {
	s.StorageResource = &v
	return s
}

type ModifyDBClusterResponseBody struct {
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OrderId     *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterResponseBody) SetDBClusterId(v string) *ModifyDBClusterResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterResponseBody) SetOrderId(v string) *ModifyDBClusterResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyDBClusterResponseBody) SetRequestId(v string) *ModifyDBClusterResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterResponse) SetHeaders(v map[string]*string) *ModifyDBClusterResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterResponse) SetBody(v *ModifyDBClusterResponseBody) *ModifyDBClusterResponse {
	s.Body = v
	return s
}

type ModifyDBClusterAccessWhiteListRequest struct {
	DBClusterIPArrayAttribute *string `json:"DBClusterIPArrayAttribute,omitempty" xml:"DBClusterIPArrayAttribute,omitempty"`
	DBClusterIPArrayName      *string `json:"DBClusterIPArrayName,omitempty" xml:"DBClusterIPArrayName,omitempty"`
	DBClusterId               *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ModifyMode                *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityIps               *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
}

func (s ModifyDBClusterAccessWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterAccessWhiteListRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetDBClusterIPArrayAttribute(v string) *ModifyDBClusterAccessWhiteListRequest {
	s.DBClusterIPArrayAttribute = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetDBClusterIPArrayName(v string) *ModifyDBClusterAccessWhiteListRequest {
	s.DBClusterIPArrayName = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetDBClusterId(v string) *ModifyDBClusterAccessWhiteListRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetModifyMode(v string) *ModifyDBClusterAccessWhiteListRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetOwnerAccount(v string) *ModifyDBClusterAccessWhiteListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetOwnerId(v int64) *ModifyDBClusterAccessWhiteListRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterAccessWhiteListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetResourceOwnerId(v int64) *ModifyDBClusterAccessWhiteListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetSecurityIps(v string) *ModifyDBClusterAccessWhiteListRequest {
	s.SecurityIps = &v
	return s
}

type ModifyDBClusterAccessWhiteListResponseBody struct {
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId      *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyDBClusterAccessWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterAccessWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAccessWhiteListResponseBody) SetDBClusterId(v string) *ModifyDBClusterAccessWhiteListResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListResponseBody) SetRequestId(v string) *ModifyDBClusterAccessWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListResponseBody) SetTaskId(v int32) *ModifyDBClusterAccessWhiteListResponseBody {
	s.TaskId = &v
	return s
}

type ModifyDBClusterAccessWhiteListResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBClusterAccessWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBClusterAccessWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterAccessWhiteListResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAccessWhiteListResponse) SetHeaders(v map[string]*string) *ModifyDBClusterAccessWhiteListResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterAccessWhiteListResponse) SetBody(v *ModifyDBClusterAccessWhiteListResponseBody) *ModifyDBClusterAccessWhiteListResponse {
	s.Body = v
	return s
}

type ModifyDBClusterDescriptionRequest struct {
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBClusterDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterDescriptionRequest) SetDBClusterDescription(v string) *ModifyDBClusterDescriptionRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetDBClusterId(v string) *ModifyDBClusterDescriptionRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetOwnerAccount(v string) *ModifyDBClusterDescriptionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetOwnerId(v int64) *ModifyDBClusterDescriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterDescriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetResourceOwnerId(v int64) *ModifyDBClusterDescriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyDBClusterDescriptionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterDescriptionResponseBody) SetRequestId(v string) *ModifyDBClusterDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterDescriptionResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBClusterDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBClusterDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterDescriptionResponse) SetHeaders(v map[string]*string) *ModifyDBClusterDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterDescriptionResponse) SetBody(v *ModifyDBClusterDescriptionResponseBody) *ModifyDBClusterDescriptionResponse {
	s.Body = v
	return s
}

type ModifyDBClusterMaintainTimeRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	MaintainTime         *string `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBClusterMaintainTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterMaintainTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMaintainTimeRequest) SetDBClusterId(v string) *ModifyDBClusterMaintainTimeRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetMaintainTime(v string) *ModifyDBClusterMaintainTimeRequest {
	s.MaintainTime = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetOwnerAccount(v string) *ModifyDBClusterMaintainTimeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetOwnerId(v int64) *ModifyDBClusterMaintainTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterMaintainTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetResourceOwnerId(v int64) *ModifyDBClusterMaintainTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyDBClusterMaintainTimeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterMaintainTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterMaintainTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMaintainTimeResponseBody) SetRequestId(v string) *ModifyDBClusterMaintainTimeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterMaintainTimeResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBClusterMaintainTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBClusterMaintainTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterMaintainTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMaintainTimeResponse) SetHeaders(v map[string]*string) *ModifyDBClusterMaintainTimeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterMaintainTimeResponse) SetBody(v *ModifyDBClusterMaintainTimeResponseBody) *ModifyDBClusterMaintainTimeResponse {
	s.Body = v
	return s
}

type ModifyDBClusterResourceGroupRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	NewResourceGroupId   *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBClusterResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterResourceGroupRequest) SetDBClusterId(v string) *ModifyDBClusterResourceGroupRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterResourceGroupRequest) SetNewResourceGroupId(v string) *ModifyDBClusterResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *ModifyDBClusterResourceGroupRequest) SetOwnerAccount(v string) *ModifyDBClusterResourceGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterResourceGroupRequest) SetOwnerId(v int64) *ModifyDBClusterResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterResourceGroupRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterResourceGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterResourceGroupRequest) SetResourceOwnerId(v int64) *ModifyDBClusterResourceGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyDBClusterResourceGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterResourceGroupResponseBody) SetRequestId(v string) *ModifyDBClusterResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterResourceGroupResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBClusterResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBClusterResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterResourceGroupResponse) SetHeaders(v map[string]*string) *ModifyDBClusterResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterResourceGroupResponse) SetBody(v *ModifyDBClusterResourceGroupResponseBody) *ModifyDBClusterResourceGroupResponse {
	s.Body = v
	return s
}

type ModifyDBResourcePoolRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	NodeNum              *int32  `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PoolName             *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	QueryType            *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBResourcePoolRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBResourcePoolRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBResourcePoolRequest) SetDBClusterId(v string) *ModifyDBResourcePoolRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBResourcePoolRequest) SetNodeNum(v int32) *ModifyDBResourcePoolRequest {
	s.NodeNum = &v
	return s
}

func (s *ModifyDBResourcePoolRequest) SetOwnerAccount(v string) *ModifyDBResourcePoolRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBResourcePoolRequest) SetOwnerId(v int64) *ModifyDBResourcePoolRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBResourcePoolRequest) SetPoolName(v string) *ModifyDBResourcePoolRequest {
	s.PoolName = &v
	return s
}

func (s *ModifyDBResourcePoolRequest) SetQueryType(v string) *ModifyDBResourcePoolRequest {
	s.QueryType = &v
	return s
}

func (s *ModifyDBResourcePoolRequest) SetResourceOwnerAccount(v string) *ModifyDBResourcePoolRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBResourcePoolRequest) SetResourceOwnerId(v int64) *ModifyDBResourcePoolRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyDBResourcePoolResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBResourcePoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBResourcePoolResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBResourcePoolResponseBody) SetRequestId(v string) *ModifyDBResourcePoolResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBResourcePoolResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBResourcePoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBResourcePoolResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBResourcePoolResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBResourcePoolResponse) SetHeaders(v map[string]*string) *ModifyDBResourcePoolResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBResourcePoolResponse) SetBody(v *ModifyDBResourcePoolResponseBody) *ModifyDBResourcePoolResponse {
	s.Body = v
	return s
}

type ModifyElasticPlanRequest struct {
	DBClusterId             *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ElasticPlanEnable       *bool   `json:"ElasticPlanEnable,omitempty" xml:"ElasticPlanEnable,omitempty"`
	ElasticPlanEndDay       *string `json:"ElasticPlanEndDay,omitempty" xml:"ElasticPlanEndDay,omitempty"`
	ElasticPlanName         *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	ElasticPlanNodeNum      *int32  `json:"ElasticPlanNodeNum,omitempty" xml:"ElasticPlanNodeNum,omitempty"`
	ElasticPlanStartDay     *string `json:"ElasticPlanStartDay,omitempty" xml:"ElasticPlanStartDay,omitempty"`
	ElasticPlanTimeEnd      *string `json:"ElasticPlanTimeEnd,omitempty" xml:"ElasticPlanTimeEnd,omitempty"`
	ElasticPlanTimeStart    *string `json:"ElasticPlanTimeStart,omitempty" xml:"ElasticPlanTimeStart,omitempty"`
	ElasticPlanWeeklyRepeat *string `json:"ElasticPlanWeeklyRepeat,omitempty" xml:"ElasticPlanWeeklyRepeat,omitempty"`
	OwnerAccount            *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                 *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount    *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId         *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourcePoolName        *string `json:"ResourcePoolName,omitempty" xml:"ResourcePoolName,omitempty"`
}

func (s ModifyElasticPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyElasticPlanRequest) GoString() string {
	return s.String()
}

func (s *ModifyElasticPlanRequest) SetDBClusterId(v string) *ModifyElasticPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanEnable(v bool) *ModifyElasticPlanRequest {
	s.ElasticPlanEnable = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanEndDay(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanEndDay = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanName(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanName = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanNodeNum(v int32) *ModifyElasticPlanRequest {
	s.ElasticPlanNodeNum = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanStartDay(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanStartDay = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanTimeEnd(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanTimeEnd = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanTimeStart(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanTimeStart = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanWeeklyRepeat(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanWeeklyRepeat = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetOwnerAccount(v string) *ModifyElasticPlanRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetOwnerId(v int64) *ModifyElasticPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetResourceOwnerAccount(v string) *ModifyElasticPlanRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetResourceOwnerId(v int64) *ModifyElasticPlanRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetResourcePoolName(v string) *ModifyElasticPlanRequest {
	s.ResourcePoolName = &v
	return s
}

type ModifyElasticPlanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyElasticPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyElasticPlanResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyElasticPlanResponseBody) SetRequestId(v string) *ModifyElasticPlanResponseBody {
	s.RequestId = &v
	return s
}

type ModifyElasticPlanResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyElasticPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyElasticPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyElasticPlanResponse) GoString() string {
	return s.String()
}

func (s *ModifyElasticPlanResponse) SetHeaders(v map[string]*string) *ModifyElasticPlanResponse {
	s.Headers = v
	return s
}

func (s *ModifyElasticPlanResponse) SetBody(v *ModifyElasticPlanResponseBody) *ModifyElasticPlanResponse {
	s.Body = v
	return s
}

type ModifyLogBackupPolicyRequest struct {
	DBClusterId              *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	EnableBackupLog          *string `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	LogBackupRetentionPeriod *string `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	OwnerAccount             *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount     *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId          *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyLogBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyLogBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyLogBackupPolicyRequest) SetDBClusterId(v string) *ModifyLogBackupPolicyRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetEnableBackupLog(v string) *ModifyLogBackupPolicyRequest {
	s.EnableBackupLog = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetLogBackupRetentionPeriod(v string) *ModifyLogBackupPolicyRequest {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetOwnerAccount(v string) *ModifyLogBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetOwnerId(v int64) *ModifyLogBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetResourceOwnerAccount(v string) *ModifyLogBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetResourceOwnerId(v int64) *ModifyLogBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyLogBackupPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLogBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyLogBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLogBackupPolicyResponseBody) SetRequestId(v string) *ModifyLogBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyLogBackupPolicyResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyLogBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyLogBackupPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyLogBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyLogBackupPolicyResponse) SetHeaders(v map[string]*string) *ModifyLogBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyLogBackupPolicyResponse) SetBody(v *ModifyLogBackupPolicyResponseBody) *ModifyLogBackupPolicyResponse {
	s.Body = v
	return s
}

type ModifyMaintenanceActionRequest struct {
	Ids                  *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SwitchTime           *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
}

func (s ModifyMaintenanceActionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyMaintenanceActionRequest) GoString() string {
	return s.String()
}

func (s *ModifyMaintenanceActionRequest) SetIds(v string) *ModifyMaintenanceActionRequest {
	s.Ids = &v
	return s
}

func (s *ModifyMaintenanceActionRequest) SetOwnerAccount(v string) *ModifyMaintenanceActionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyMaintenanceActionRequest) SetOwnerId(v int64) *ModifyMaintenanceActionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyMaintenanceActionRequest) SetRegionId(v string) *ModifyMaintenanceActionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyMaintenanceActionRequest) SetResourceOwnerAccount(v string) *ModifyMaintenanceActionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyMaintenanceActionRequest) SetResourceOwnerId(v int64) *ModifyMaintenanceActionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyMaintenanceActionRequest) SetSwitchTime(v string) *ModifyMaintenanceActionRequest {
	s.SwitchTime = &v
	return s
}

type ModifyMaintenanceActionResponseBody struct {
	Ids       *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyMaintenanceActionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyMaintenanceActionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMaintenanceActionResponseBody) SetIds(v string) *ModifyMaintenanceActionResponseBody {
	s.Ids = &v
	return s
}

func (s *ModifyMaintenanceActionResponseBody) SetRequestId(v string) *ModifyMaintenanceActionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyMaintenanceActionResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyMaintenanceActionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyMaintenanceActionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyMaintenanceActionResponse) GoString() string {
	return s.String()
}

func (s *ModifyMaintenanceActionResponse) SetHeaders(v map[string]*string) *ModifyMaintenanceActionResponse {
	s.Headers = v
	return s
}

func (s *ModifyMaintenanceActionResponse) SetBody(v *ModifyMaintenanceActionResponseBody) *ModifyMaintenanceActionResponse {
	s.Body = v
	return s
}

type ReleaseClusterPublicConnectionRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ReleaseClusterPublicConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseClusterPublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *ReleaseClusterPublicConnectionRequest) SetDBClusterId(v string) *ReleaseClusterPublicConnectionRequest {
	s.DBClusterId = &v
	return s
}

func (s *ReleaseClusterPublicConnectionRequest) SetOwnerAccount(v string) *ReleaseClusterPublicConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReleaseClusterPublicConnectionRequest) SetOwnerId(v int64) *ReleaseClusterPublicConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseClusterPublicConnectionRequest) SetResourceOwnerAccount(v string) *ReleaseClusterPublicConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleaseClusterPublicConnectionRequest) SetResourceOwnerId(v int64) *ReleaseClusterPublicConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

type ReleaseClusterPublicConnectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseClusterPublicConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseClusterPublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseClusterPublicConnectionResponseBody) SetRequestId(v string) *ReleaseClusterPublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseClusterPublicConnectionResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleaseClusterPublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseClusterPublicConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseClusterPublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *ReleaseClusterPublicConnectionResponse) SetHeaders(v map[string]*string) *ReleaseClusterPublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *ReleaseClusterPublicConnectionResponse) SetBody(v *ReleaseClusterPublicConnectionResponseBody) *ReleaseClusterPublicConnectionResponse {
	s.Body = v
	return s
}

type ResetAccountPasswordRequest struct {
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPassword      *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	AccountType          *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ResetAccountPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordRequest) SetAccountName(v string) *ResetAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetAccountPassword(v string) *ResetAccountPasswordRequest {
	s.AccountPassword = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetAccountType(v string) *ResetAccountPasswordRequest {
	s.AccountType = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetDBClusterId(v string) *ResetAccountPasswordRequest {
	s.DBClusterId = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetOwnerAccount(v string) *ResetAccountPasswordRequest {
	s.OwnerAccount = &v
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

type ResetAccountPasswordResponseBody struct {
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId      *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ResetAccountPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordResponseBody) SetDBClusterId(v string) *ResetAccountPasswordResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ResetAccountPasswordResponseBody) SetRequestId(v string) *ResetAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetAccountPasswordResponseBody) SetTaskId(v int32) *ResetAccountPasswordResponseBody {
	s.TaskId = &v
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

type RevokeOperatorPermissionRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RevokeOperatorPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeOperatorPermissionRequest) GoString() string {
	return s.String()
}

func (s *RevokeOperatorPermissionRequest) SetDBClusterId(v string) *RevokeOperatorPermissionRequest {
	s.DBClusterId = &v
	return s
}

func (s *RevokeOperatorPermissionRequest) SetOwnerAccount(v string) *RevokeOperatorPermissionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RevokeOperatorPermissionRequest) SetOwnerId(v int64) *RevokeOperatorPermissionRequest {
	s.OwnerId = &v
	return s
}

func (s *RevokeOperatorPermissionRequest) SetResourceOwnerAccount(v string) *RevokeOperatorPermissionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RevokeOperatorPermissionRequest) SetResourceOwnerId(v int64) *RevokeOperatorPermissionRequest {
	s.ResourceOwnerId = &v
	return s
}

type RevokeOperatorPermissionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeOperatorPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokeOperatorPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeOperatorPermissionResponseBody) SetRequestId(v string) *RevokeOperatorPermissionResponseBody {
	s.RequestId = &v
	return s
}

type RevokeOperatorPermissionResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RevokeOperatorPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RevokeOperatorPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeOperatorPermissionResponse) GoString() string {
	return s.String()
}

func (s *RevokeOperatorPermissionResponse) SetHeaders(v map[string]*string) *RevokeOperatorPermissionResponse {
	s.Headers = v
	return s
}

func (s *RevokeOperatorPermissionResponse) SetBody(v *RevokeOperatorPermissionResponseBody) *RevokeOperatorPermissionResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	OwnerAccount         *string                   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId           []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string                   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag                  []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetOwnerAccount(v string) *TagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetOwnerId(v int64) *TagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
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

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
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

type UnbindDBResourcePoolWithUserRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PoolName             *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	PoolUser             *string `json:"PoolUser,omitempty" xml:"PoolUser,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UnbindDBResourcePoolWithUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindDBResourcePoolWithUserRequest) GoString() string {
	return s.String()
}

func (s *UnbindDBResourcePoolWithUserRequest) SetDBClusterId(v string) *UnbindDBResourcePoolWithUserRequest {
	s.DBClusterId = &v
	return s
}

func (s *UnbindDBResourcePoolWithUserRequest) SetOwnerAccount(v string) *UnbindDBResourcePoolWithUserRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnbindDBResourcePoolWithUserRequest) SetOwnerId(v int64) *UnbindDBResourcePoolWithUserRequest {
	s.OwnerId = &v
	return s
}

func (s *UnbindDBResourcePoolWithUserRequest) SetPoolName(v string) *UnbindDBResourcePoolWithUserRequest {
	s.PoolName = &v
	return s
}

func (s *UnbindDBResourcePoolWithUserRequest) SetPoolUser(v string) *UnbindDBResourcePoolWithUserRequest {
	s.PoolUser = &v
	return s
}

func (s *UnbindDBResourcePoolWithUserRequest) SetResourceOwnerAccount(v string) *UnbindDBResourcePoolWithUserRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnbindDBResourcePoolWithUserRequest) SetResourceOwnerId(v int64) *UnbindDBResourcePoolWithUserRequest {
	s.ResourceOwnerId = &v
	return s
}

type UnbindDBResourcePoolWithUserResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindDBResourcePoolWithUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindDBResourcePoolWithUserResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindDBResourcePoolWithUserResponseBody) SetRequestId(v string) *UnbindDBResourcePoolWithUserResponseBody {
	s.RequestId = &v
	return s
}

type UnbindDBResourcePoolWithUserResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnbindDBResourcePoolWithUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindDBResourcePoolWithUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindDBResourcePoolWithUserResponse) GoString() string {
	return s.String()
}

func (s *UnbindDBResourcePoolWithUserResponse) SetHeaders(v map[string]*string) *UnbindDBResourcePoolWithUserResponse {
	s.Headers = v
	return s
}

func (s *UnbindDBResourcePoolWithUserResponse) SetBody(v *UnbindDBResourcePoolWithUserResponseBody) *UnbindDBResourcePoolWithUserResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	All                  *bool     `json:"All,omitempty" xml:"All,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey               []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerAccount(v string) *UntagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerId(v int64) *UntagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
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

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
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
		"cn-qingdao":                  tea.String("adb.aliyuncs.com"),
		"cn-beijing":                  tea.String("adb.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("adb.aliyuncs.com"),
		"cn-shanghai":                 tea.String("adb.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("adb.aliyuncs.com"),
		"cn-hongkong":                 tea.String("adb.aliyuncs.com"),
		"ap-southeast-1":              tea.String("adb.aliyuncs.com"),
		"us-west-1":                   tea.String("adb.aliyuncs.com"),
		"us-east-1":                   tea.String("adb.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("adb.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("adb.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("adb.ap-northeast-1.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("adb.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("adb.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("adb.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("adb.aliyuncs.com"),
		"cn-edge-1":                   tea.String("adb.aliyuncs.com"),
		"cn-fujian":                   tea.String("adb.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("adb.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("adb.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("adb.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("adb.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("adb.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("adb.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("adb.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("adb.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("adb.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("adb.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("adb.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("adb.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("adb.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("adb.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("adb.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("adb.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("adb.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("adb.aliyuncs.com"),
		"cn-wuhan":                    tea.String("adb.aliyuncs.com"),
		"cn-yushanfang":               tea.String("adb.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("adb.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("adb.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("adb.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("adb.ap-northeast-1.aliyuncs.com"),
		"me-east-1":                   tea.String("adb.ap-northeast-1.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("adb.ap-northeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("adb"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AllocateClusterPublicConnectionWithOptions(request *AllocateClusterPublicConnectionRequest, runtime *util.RuntimeOptions) (_result *AllocateClusterPublicConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionStringPrefix)) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AllocateClusterPublicConnection"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AllocateClusterPublicConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AllocateClusterPublicConnection(request *AllocateClusterPublicConnectionRequest) (_result *AllocateClusterPublicConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AllocateClusterPublicConnectionResponse{}
	_body, _err := client.AllocateClusterPublicConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindDBResourcePoolWithUserWithOptions(request *BindDBResourcePoolWithUserRequest, runtime *util.RuntimeOptions) (_result *BindDBResourcePoolWithUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolName)) {
		query["PoolName"] = request.PoolName
	}

	if !tea.BoolValue(util.IsUnset(request.PoolUser)) {
		query["PoolUser"] = request.PoolUser
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindDBResourcePoolWithUser"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindDBResourcePoolWithUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindDBResourcePoolWithUser(request *BindDBResourcePoolWithUserRequest) (_result *BindDBResourcePoolWithUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindDBResourcePoolWithUserResponse{}
	_body, _err := client.BindDBResourcePoolWithUserWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountDescription)) {
		query["AccountDescription"] = request.AccountDescription
	}

	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AccountPassword)) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.AccountType)) {
		query["AccountType"] = request.AccountType
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccount"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateDBClusterWithOptions(request *CreateDBClusterRequest, runtime *util.RuntimeOptions) (_result *CreateDBClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupSetID)) {
		query["BackupSetID"] = request.BackupSetID
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ComputeResource)) {
		query["ComputeResource"] = request.ComputeResource
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterCategory)) {
		query["DBClusterCategory"] = request.DBClusterCategory
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterClass)) {
		query["DBClusterClass"] = request.DBClusterClass
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterDescription)) {
		query["DBClusterDescription"] = request.DBClusterDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterNetworkType)) {
		query["DBClusterNetworkType"] = request.DBClusterNetworkType
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterVersion)) {
		query["DBClusterVersion"] = request.DBClusterVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DBNodeGroupCount)) {
		query["DBNodeGroupCount"] = request.DBNodeGroupCount
	}

	if !tea.BoolValue(util.IsUnset(request.DBNodeStorage)) {
		query["DBNodeStorage"] = request.DBNodeStorage
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticIOResource)) {
		query["ElasticIOResource"] = request.ElasticIOResource
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorCount)) {
		query["ExecutorCount"] = request.ExecutorCount
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RestoreTime)) {
		query["RestoreTime"] = request.RestoreTime
	}

	if !tea.BoolValue(util.IsUnset(request.RestoreType)) {
		query["RestoreType"] = request.RestoreType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceDBInstanceName)) {
		query["SourceDBInstanceName"] = request.SourceDBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.StorageResource)) {
		query["StorageResource"] = request.StorageResource
	}

	if !tea.BoolValue(util.IsUnset(request.StorageType)) {
		query["StorageType"] = request.StorageType
	}

	if !tea.BoolValue(util.IsUnset(request.UsedTime)) {
		query["UsedTime"] = request.UsedTime
	}

	if !tea.BoolValue(util.IsUnset(request.VPCId)) {
		query["VPCId"] = request.VPCId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDBCluster"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDBClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDBCluster(request *CreateDBClusterRequest) (_result *CreateDBClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBClusterResponse{}
	_body, _err := client.CreateDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDBResourcePoolWithOptions(request *CreateDBResourcePoolRequest, runtime *util.RuntimeOptions) (_result *CreateDBResourcePoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeNum)) {
		query["NodeNum"] = request.NodeNum
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolName)) {
		query["PoolName"] = request.PoolName
	}

	if !tea.BoolValue(util.IsUnset(request.QueryType)) {
		query["QueryType"] = request.QueryType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDBResourcePool"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDBResourcePoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDBResourcePool(request *CreateDBResourcePoolRequest) (_result *CreateDBResourcePoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBResourcePoolResponse{}
	_body, _err := client.CreateDBResourcePoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateElasticPlanWithOptions(request *CreateElasticPlanRequest, runtime *util.RuntimeOptions) (_result *CreateElasticPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanEnable)) {
		query["ElasticPlanEnable"] = request.ElasticPlanEnable
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanEndDay)) {
		query["ElasticPlanEndDay"] = request.ElasticPlanEndDay
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanName)) {
		query["ElasticPlanName"] = request.ElasticPlanName
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanNodeNum)) {
		query["ElasticPlanNodeNum"] = request.ElasticPlanNodeNum
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanStartDay)) {
		query["ElasticPlanStartDay"] = request.ElasticPlanStartDay
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanTimeEnd)) {
		query["ElasticPlanTimeEnd"] = request.ElasticPlanTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanTimeStart)) {
		query["ElasticPlanTimeStart"] = request.ElasticPlanTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanWeeklyRepeat)) {
		query["ElasticPlanWeeklyRepeat"] = request.ElasticPlanWeeklyRepeat
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourcePoolName)) {
		query["ResourcePoolName"] = request.ResourcePoolName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateElasticPlan"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateElasticPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateElasticPlan(request *CreateElasticPlanRequest) (_result *CreateElasticPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateElasticPlanResponse{}
	_body, _err := client.CreateElasticPlanWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AccountType)) {
		query["AccountType"] = request.AccountType
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccount"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DeleteDBClusterWithOptions(request *DeleteDBClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteDBClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDBCluster"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDBClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDBCluster(request *DeleteDBClusterRequest) (_result *DeleteDBClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDBClusterResponse{}
	_body, _err := client.DeleteDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDBResourcePoolWithOptions(request *DeleteDBResourcePoolRequest, runtime *util.RuntimeOptions) (_result *DeleteDBResourcePoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolName)) {
		query["PoolName"] = request.PoolName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDBResourcePool"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDBResourcePoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDBResourcePool(request *DeleteDBResourcePoolRequest) (_result *DeleteDBResourcePoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDBResourcePoolResponse{}
	_body, _err := client.DeleteDBResourcePoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteElasticPlanWithOptions(request *DeleteElasticPlanRequest, runtime *util.RuntimeOptions) (_result *DeleteElasticPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanName)) {
		query["ElasticPlanName"] = request.ElasticPlanName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteElasticPlan"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteElasticPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteElasticPlan(request *DeleteElasticPlanRequest) (_result *DeleteElasticPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteElasticPlanResponse{}
	_body, _err := client.DeleteElasticPlanWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AccountType)) {
		query["AccountType"] = request.AccountType
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccounts"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeAllAccountsWithOptions(request *DescribeAllAccountsRequest, runtime *util.RuntimeOptions) (_result *DescribeAllAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAllAccounts"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAllAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAllAccounts(request *DescribeAllAccountsRequest) (_result *DescribeAllAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAllAccountsResponse{}
	_body, _err := client.DescribeAllAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAllDataSourceWithOptions(request *DescribeAllDataSourceRequest, runtime *util.RuntimeOptions) (_result *DescribeAllDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaName)) {
		query["SchemaName"] = request.SchemaName
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAllDataSource"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAllDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAllDataSource(request *DescribeAllDataSourceRequest) (_result *DescribeAllDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAllDataSourceResponse{}
	_body, _err := client.DescribeAllDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAuditLogConfigWithOptions(request *DescribeAuditLogConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeAuditLogConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAuditLogConfig"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAuditLogConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAuditLogConfig(request *DescribeAuditLogConfigRequest) (_result *DescribeAuditLogConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAuditLogConfigResponse{}
	_body, _err := client.DescribeAuditLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAuditLogRecordsWithOptions(request *DescribeAuditLogRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeAuditLogRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DBName)) {
		query["DBName"] = request.DBName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.HostAddress)) {
		query["HostAddress"] = request.HostAddress
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		query["OrderType"] = request.OrderType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryKeyword)) {
		query["QueryKeyword"] = request.QueryKeyword
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SqlType)) {
		query["SqlType"] = request.SqlType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Succeed)) {
		query["Succeed"] = request.Succeed
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		query["User"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAuditLogRecords"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAuditLogRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAuditLogRecords(request *DescribeAuditLogRecordsRequest) (_result *DescribeAuditLogRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAuditLogRecordsResponse{}
	_body, _err := client.DescribeAuditLogRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAutoRenewAttributeWithOptions(request *DescribeAutoRenewAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeAutoRenewAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterIds)) {
		query["DBClusterIds"] = request.DBClusterIds
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAutoRenewAttribute"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAutoRenewAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAutoRenewAttribute(request *DescribeAutoRenewAttributeRequest) (_result *DescribeAutoRenewAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAutoRenewAttributeResponse{}
	_body, _err := client.DescribeAutoRenewAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAvailableResourceWithOptions(request *DescribeAvailableResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeAvailableResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAvailableResource"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAvailableResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAvailableResource(request *DescribeAvailableResourceRequest) (_result *DescribeAvailableResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAvailableResourceResponse{}
	_body, _err := client.DescribeAvailableResourceWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupPolicy"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		query["BackupId"] = request.BackupId
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackups"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeColumnsWithOptions(request *DescribeColumnsRequest, runtime *util.RuntimeOptions) (_result *DescribeColumnsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaName)) {
		query["SchemaName"] = request.SchemaName
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeColumns"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeColumnsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeColumns(request *DescribeColumnsRequest) (_result *DescribeColumnsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeColumnsResponse{}
	_body, _err := client.DescribeColumnsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeConnectionCountRecordsWithOptions(request *DescribeConnectionCountRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeConnectionCountRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeConnectionCountRecords"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeConnectionCountRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeConnectionCountRecords(request *DescribeConnectionCountRecordsRequest) (_result *DescribeConnectionCountRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeConnectionCountRecordsResponse{}
	_body, _err := client.DescribeConnectionCountRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterAccessWhiteListWithOptions(request *DescribeDBClusterAccessWhiteListRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterAccessWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterAccessWhiteList"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClusterAccessWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterAccessWhiteList(request *DescribeDBClusterAccessWhiteListRequest) (_result *DescribeDBClusterAccessWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterAccessWhiteListResponse{}
	_body, _err := client.DescribeDBClusterAccessWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterAttributeWithOptions(request *DescribeDBClusterAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterAttribute"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClusterAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterAttribute(request *DescribeDBClusterAttributeRequest) (_result *DescribeDBClusterAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterAttributeResponse{}
	_body, _err := client.DescribeDBClusterAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterForecastWithOptions(request *DescribeDBClusterForecastRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterForecastResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.MetricType)) {
		query["MetricType"] = request.MetricType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterForecast"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClusterForecastResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterForecast(request *DescribeDBClusterForecastRequest) (_result *DescribeDBClusterForecastResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterForecastResponse{}
	_body, _err := client.DescribeDBClusterForecastWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterHealthReportWithOptions(request *DescribeDBClusterHealthReportRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterHealthReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterHealthReport"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClusterHealthReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterHealthReport(request *DescribeDBClusterHealthReportRequest) (_result *DescribeDBClusterHealthReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterHealthReportResponse{}
	_body, _err := client.DescribeDBClusterHealthReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterNetInfoWithOptions(request *DescribeDBClusterNetInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterNetInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterNetInfo"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClusterNetInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterNetInfo(request *DescribeDBClusterNetInfoRequest) (_result *DescribeDBClusterNetInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterNetInfoResponse{}
	_body, _err := client.DescribeDBClusterNetInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterPerformanceWithOptions(request *DescribeDBClusterPerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterPerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterPerformance"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClusterPerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterPerformance(request *DescribeDBClusterPerformanceRequest) (_result *DescribeDBClusterPerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterPerformanceResponse{}
	_body, _err := client.DescribeDBClusterPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterResourcePoolPerformanceWithOptions(request *DescribeDBClusterResourcePoolPerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterResourcePoolPerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourcePools)) {
		query["ResourcePools"] = request.ResourcePools
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterResourcePoolPerformance"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClusterResourcePoolPerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterResourcePoolPerformance(request *DescribeDBClusterResourcePoolPerformanceRequest) (_result *DescribeDBClusterResourcePoolPerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterResourcePoolPerformanceResponse{}
	_body, _err := client.DescribeDBClusterResourcePoolPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClustersWithOptions(request *DescribeDBClustersRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterDescription)) {
		query["DBClusterDescription"] = request.DBClusterDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterIds)) {
		query["DBClusterIds"] = request.DBClusterIds
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterStatus)) {
		query["DBClusterStatus"] = request.DBClusterStatus
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusters"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusters(request *DescribeDBClustersRequest) (_result *DescribeDBClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClustersResponse{}
	_body, _err := client.DescribeDBClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBResourcePoolWithOptions(request *DescribeDBResourcePoolRequest, runtime *util.RuntimeOptions) (_result *DescribeDBResourcePoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolName)) {
		query["PoolName"] = request.PoolName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBResourcePool"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBResourcePoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBResourcePool(request *DescribeDBResourcePoolRequest) (_result *DescribeDBResourcePoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBResourcePoolResponse{}
	_body, _err := client.DescribeDBResourcePoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDiagnosisDimensionsWithOptions(request *DescribeDiagnosisDimensionsRequest, runtime *util.RuntimeOptions) (_result *DescribeDiagnosisDimensionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.QueryCondition)) {
		query["QueryCondition"] = request.QueryCondition
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiagnosisDimensions"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiagnosisDimensionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDiagnosisDimensions(request *DescribeDiagnosisDimensionsRequest) (_result *DescribeDiagnosisDimensionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiagnosisDimensionsResponse{}
	_body, _err := client.DescribeDiagnosisDimensionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDiagnosisRecordsWithOptions(request *DescribeDiagnosisRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeDiagnosisRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientIp)) {
		query["ClientIp"] = request.ClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MaxPeakMemory)) {
		query["MaxPeakMemory"] = request.MaxPeakMemory
	}

	if !tea.BoolValue(util.IsUnset(request.MaxScanSize)) {
		query["MaxScanSize"] = request.MaxScanSize
	}

	if !tea.BoolValue(util.IsUnset(request.MinPeakMemory)) {
		query["MinPeakMemory"] = request.MinPeakMemory
	}

	if !tea.BoolValue(util.IsUnset(request.MinScanSize)) {
		query["MinScanSize"] = request.MinScanSize
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PatternId)) {
		query["PatternId"] = request.PatternId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryCondition)) {
		query["QueryCondition"] = request.QueryCondition
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroup)) {
		query["ResourceGroup"] = request.ResourceGroup
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiagnosisRecords"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiagnosisRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDiagnosisRecords(request *DescribeDiagnosisRecordsRequest) (_result *DescribeDiagnosisRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiagnosisRecordsResponse{}
	_body, _err := client.DescribeDiagnosisRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDownloadRecordsWithOptions(request *DescribeDownloadRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeDownloadRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDownloadRecords"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDownloadRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDownloadRecords(request *DescribeDownloadRecordsRequest) (_result *DescribeDownloadRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDownloadRecordsResponse{}
	_body, _err := client.DescribeDownloadRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeElasticDailyPlanWithOptions(request *DescribeElasticDailyPlanRequest, runtime *util.RuntimeOptions) (_result *DescribeElasticDailyPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticDailyPlanDay)) {
		query["ElasticDailyPlanDay"] = request.ElasticDailyPlanDay
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticDailyPlanStatusList)) {
		query["ElasticDailyPlanStatusList"] = request.ElasticDailyPlanStatusList
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanName)) {
		query["ElasticPlanName"] = request.ElasticPlanName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourcePoolName)) {
		query["ResourcePoolName"] = request.ResourcePoolName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeElasticDailyPlan"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeElasticDailyPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeElasticDailyPlan(request *DescribeElasticDailyPlanRequest) (_result *DescribeElasticDailyPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeElasticDailyPlanResponse{}
	_body, _err := client.DescribeElasticDailyPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeElasticPlanWithOptions(request *DescribeElasticPlanRequest, runtime *util.RuntimeOptions) (_result *DescribeElasticPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanEnable)) {
		query["ElasticPlanEnable"] = request.ElasticPlanEnable
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanName)) {
		query["ElasticPlanName"] = request.ElasticPlanName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourcePoolName)) {
		query["ResourcePoolName"] = request.ResourcePoolName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeElasticPlan"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeElasticPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeElasticPlan(request *DescribeElasticPlanRequest) (_result *DescribeElasticPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeElasticPlanResponse{}
	_body, _err := client.DescribeElasticPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInclinedTablesWithOptions(request *DescribeInclinedTablesRequest, runtime *util.RuntimeOptions) (_result *DescribeInclinedTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TableType)) {
		query["TableType"] = request.TableType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInclinedTables"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInclinedTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInclinedTables(request *DescribeInclinedTablesRequest) (_result *DescribeInclinedTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInclinedTablesResponse{}
	_body, _err := client.DescribeInclinedTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLoadTasksRecordsWithOptions(request *DescribeLoadTasksRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeLoadTasksRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DBName)) {
		query["DBName"] = request.DBName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLoadTasksRecords"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLoadTasksRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLoadTasksRecords(request *DescribeLoadTasksRecordsRequest) (_result *DescribeLoadTasksRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLoadTasksRecordsResponse{}
	_body, _err := client.DescribeLoadTasksRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMaintenanceActionWithOptions(request *DescribeMaintenanceActionRequest, runtime *util.RuntimeOptions) (_result *DescribeMaintenanceActionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsHistory)) {
		query["IsHistory"] = request.IsHistory
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMaintenanceAction"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMaintenanceActionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMaintenanceAction(request *DescribeMaintenanceActionRequest) (_result *DescribeMaintenanceActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMaintenanceActionResponse{}
	_body, _err := client.DescribeMaintenanceActionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOperatorPermissionWithOptions(request *DescribeOperatorPermissionRequest, runtime *util.RuntimeOptions) (_result *DescribeOperatorPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOperatorPermission"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOperatorPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOperatorPermission(request *DescribeOperatorPermissionRequest) (_result *DescribeOperatorPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOperatorPermissionResponse{}
	_body, _err := client.DescribeOperatorPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePatternPerformanceWithOptions(request *DescribePatternPerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribePatternPerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PatternId)) {
		query["PatternId"] = request.PatternId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePatternPerformance"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePatternPerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePatternPerformance(request *DescribePatternPerformanceRequest) (_result *DescribePatternPerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePatternPerformanceResponse{}
	_body, _err := client.DescribePatternPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeProcessListWithOptions(request *DescribeProcessListRequest, runtime *util.RuntimeOptions) (_result *DescribeProcessListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RunningTime)) {
		query["RunningTime"] = request.RunningTime
	}

	if !tea.BoolValue(util.IsUnset(request.ShowFull)) {
		query["ShowFull"] = request.ShowFull
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		query["User"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProcessList"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeProcessListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeProcessList(request *DescribeProcessListRequest) (_result *DescribeProcessListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProcessListResponse{}
	_body, _err := client.DescribeProcessListWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeSQLPatternAttributeWithOptions(request *DescribeSQLPatternAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLPatternAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PatternId)) {
		query["PatternId"] = request.PatternId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSQLPatternAttribute"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSQLPatternAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLPatternAttribute(request *DescribeSQLPatternAttributeRequest) (_result *DescribeSQLPatternAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLPatternAttributeResponse{}
	_body, _err := client.DescribeSQLPatternAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLPatternsWithOptions(request *DescribeSQLPatternsRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLPatternsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSQLPatterns"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSQLPatternsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLPatterns(request *DescribeSQLPatternsRequest) (_result *DescribeSQLPatternsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLPatternsResponse{}
	_body, _err := client.DescribeSQLPatternsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLPlanWithOptions(request *DescribeSQLPlanRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessId)) {
		query["ProcessId"] = request.ProcessId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSQLPlan"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSQLPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLPlan(request *DescribeSQLPlanRequest) (_result *DescribeSQLPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLPlanResponse{}
	_body, _err := client.DescribeSQLPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLPlanTaskWithOptions(request *DescribeSQLPlanTaskRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLPlanTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessId)) {
		query["ProcessId"] = request.ProcessId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StageId)) {
		query["StageId"] = request.StageId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSQLPlanTask"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSQLPlanTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLPlanTask(request *DescribeSQLPlanTaskRequest) (_result *DescribeSQLPlanTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLPlanTaskResponse{}
	_body, _err := client.DescribeSQLPlanTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSchemasWithOptions(request *DescribeSchemasRequest, runtime *util.RuntimeOptions) (_result *DescribeSchemasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSchemas"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSchemasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSchemas(request *DescribeSchemasRequest) (_result *DescribeSchemasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSchemasResponse{}
	_body, _err := client.DescribeSchemasWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DBName)) {
		query["DBName"] = request.DBName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessID)) {
		query["ProcessID"] = request.ProcessID
	}

	if !tea.BoolValue(util.IsUnset(request.Range)) {
		query["Range"] = request.Range
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSlowLogRecords"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSlowLogRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeSlowLogTrendWithOptions(request *DescribeSlowLogTrendRequest, runtime *util.RuntimeOptions) (_result *DescribeSlowLogTrendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DBName)) {
		query["DBName"] = request.DBName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSlowLogTrend"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSlowLogTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlowLogTrend(request *DescribeSlowLogTrendRequest) (_result *DescribeSlowLogTrendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlowLogTrendResponse{}
	_body, _err := client.DescribeSlowLogTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSqlPatternWithOptions(request *DescribeSqlPatternRequest, runtime *util.RuntimeOptions) (_result *DescribeSqlPatternResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SqlPattern)) {
		query["SqlPattern"] = request.SqlPattern
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSqlPattern"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSqlPatternResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSqlPattern(request *DescribeSqlPatternRequest) (_result *DescribeSqlPatternResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSqlPatternResponse{}
	_body, _err := client.DescribeSqlPatternWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTableAccessCountWithOptions(request *DescribeTableAccessCountRequest, runtime *util.RuntimeOptions) (_result *DescribeTableAccessCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTableAccessCount"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTableAccessCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTableAccessCount(request *DescribeTableAccessCountRequest) (_result *DescribeTableAccessCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTableAccessCountResponse{}
	_body, _err := client.DescribeTableAccessCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTableDetailWithOptions(request *DescribeTableDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeTableDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaName)) {
		query["SchemaName"] = request.SchemaName
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTableDetail"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTableDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTableDetail(request *DescribeTableDetailRequest) (_result *DescribeTableDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTableDetailResponse{}
	_body, _err := client.DescribeTableDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTablePartitionDiagnoseWithOptions(request *DescribeTablePartitionDiagnoseRequest, runtime *util.RuntimeOptions) (_result *DescribeTablePartitionDiagnoseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTablePartitionDiagnose"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTablePartitionDiagnoseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTablePartitionDiagnose(request *DescribeTablePartitionDiagnoseRequest) (_result *DescribeTablePartitionDiagnoseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTablePartitionDiagnoseResponse{}
	_body, _err := client.DescribeTablePartitionDiagnoseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTableStatisticsWithOptions(request *DescribeTableStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeTableStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTableStatistics"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTableStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTableStatistics(request *DescribeTableStatisticsRequest) (_result *DescribeTableStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTableStatisticsResponse{}
	_body, _err := client.DescribeTableStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTablesWithOptions(request *DescribeTablesRequest, runtime *util.RuntimeOptions) (_result *DescribeTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaName)) {
		query["SchemaName"] = request.SchemaName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTables"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTables(request *DescribeTablesRequest) (_result *DescribeTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTablesResponse{}
	_body, _err := client.DescribeTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTaskInfoWithOptions(request *DescribeTaskInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeTaskInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTaskInfo"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTaskInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTaskInfo(request *DescribeTaskInfoRequest) (_result *DescribeTaskInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTaskInfoResponse{}
	_body, _err := client.DescribeTaskInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DownloadDiagnosisRecordsWithOptions(request *DownloadDiagnosisRecordsRequest, runtime *util.RuntimeOptions) (_result *DownloadDiagnosisRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientIp)) {
		query["ClientIp"] = request.ClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MaxPeakMemory)) {
		query["MaxPeakMemory"] = request.MaxPeakMemory
	}

	if !tea.BoolValue(util.IsUnset(request.MaxScanSize)) {
		query["MaxScanSize"] = request.MaxScanSize
	}

	if !tea.BoolValue(util.IsUnset(request.MinPeakMemory)) {
		query["MinPeakMemory"] = request.MinPeakMemory
	}

	if !tea.BoolValue(util.IsUnset(request.MinScanSize)) {
		query["MinScanSize"] = request.MinScanSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryCondition)) {
		query["QueryCondition"] = request.QueryCondition
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroup)) {
		query["ResourceGroup"] = request.ResourceGroup
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DownloadDiagnosisRecords"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DownloadDiagnosisRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DownloadDiagnosisRecords(request *DownloadDiagnosisRecordsRequest) (_result *DownloadDiagnosisRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DownloadDiagnosisRecordsResponse{}
	_body, _err := client.DownloadDiagnosisRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GrantOperatorPermissionWithOptions(request *GrantOperatorPermissionRequest, runtime *util.RuntimeOptions) (_result *GrantOperatorPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ExpiredTime)) {
		query["ExpiredTime"] = request.ExpiredTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Privileges)) {
		query["Privileges"] = request.Privileges
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GrantOperatorPermission"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GrantOperatorPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GrantOperatorPermission(request *GrantOperatorPermissionRequest) (_result *GrantOperatorPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GrantOperatorPermissionResponse{}
	_body, _err := client.GrantOperatorPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) KillProcessWithOptions(request *KillProcessRequest, runtime *util.RuntimeOptions) (_result *KillProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessId)) {
		query["ProcessId"] = request.ProcessId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("KillProcess"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &KillProcessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) KillProcess(request *KillProcessRequest) (_result *KillProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &KillProcessResponse{}
	_body, _err := client.KillProcessWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ModifyAccountDescriptionWithOptions(request *ModifyAccountDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyAccountDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountDescription)) {
		query["AccountDescription"] = request.AccountDescription
	}

	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAccountDescription"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAccountDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ModifyAuditLogConfigWithOptions(request *ModifyAuditLogConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyAuditLogConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuditLogStatus)) {
		query["AuditLogStatus"] = request.AuditLogStatus
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAuditLogConfig"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAuditLogConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ModifyAutoRenewAttributeWithOptions(request *ModifyAutoRenewAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyAutoRenewAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RenewalStatus)) {
		query["RenewalStatus"] = request.RenewalStatus
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAutoRenewAttribute"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAutoRenewAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAutoRenewAttribute(request *ModifyAutoRenewAttributeRequest) (_result *ModifyAutoRenewAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAutoRenewAttributeResponse{}
	_body, _err := client.ModifyAutoRenewAttributeWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupRetentionPeriod)) {
		query["BackupRetentionPeriod"] = request.BackupRetentionPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EnableBackupLog)) {
		query["EnableBackupLog"] = request.EnableBackupLog
	}

	if !tea.BoolValue(util.IsUnset(request.LogBackupRetentionPeriod)) {
		query["LogBackupRetentionPeriod"] = request.LogBackupRetentionPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredBackupPeriod)) {
		query["PreferredBackupPeriod"] = request.PreferredBackupPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredBackupTime)) {
		query["PreferredBackupTime"] = request.PreferredBackupTime
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyBackupPolicy"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ModifyClusterConnectionStringWithOptions(request *ModifyClusterConnectionStringRequest, runtime *util.RuntimeOptions) (_result *ModifyClusterConnectionStringResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionStringPrefix)) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentConnectionString)) {
		query["CurrentConnectionString"] = request.CurrentConnectionString
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyClusterConnectionString"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyClusterConnectionStringResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyClusterConnectionString(request *ModifyClusterConnectionStringRequest) (_result *ModifyClusterConnectionStringResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyClusterConnectionStringResponse{}
	_body, _err := client.ModifyClusterConnectionStringWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterWithOptions(request *ModifyDBClusterRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ComputeResource)) {
		query["ComputeResource"] = request.ComputeResource
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterCategory)) {
		query["DBClusterCategory"] = request.DBClusterCategory
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DBNodeClass)) {
		query["DBNodeClass"] = request.DBNodeClass
	}

	if !tea.BoolValue(util.IsUnset(request.DBNodeGroupCount)) {
		query["DBNodeGroupCount"] = request.DBNodeGroupCount
	}

	if !tea.BoolValue(util.IsUnset(request.DBNodeStorage)) {
		query["DBNodeStorage"] = request.DBNodeStorage
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticIOResource)) {
		query["ElasticIOResource"] = request.ElasticIOResource
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorCount)) {
		query["ExecutorCount"] = request.ExecutorCount
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyType)) {
		query["ModifyType"] = request.ModifyType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StorageResource)) {
		query["StorageResource"] = request.StorageResource
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBCluster"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBCluster(request *ModifyDBClusterRequest) (_result *ModifyDBClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterResponse{}
	_body, _err := client.ModifyDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterAccessWhiteListWithOptions(request *ModifyDBClusterAccessWhiteListRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterAccessWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterIPArrayAttribute)) {
		query["DBClusterIPArrayAttribute"] = request.DBClusterIPArrayAttribute
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterIPArrayName)) {
		query["DBClusterIPArrayName"] = request.DBClusterIPArrayName
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyMode)) {
		query["ModifyMode"] = request.ModifyMode
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityIps)) {
		query["SecurityIps"] = request.SecurityIps
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBClusterAccessWhiteList"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBClusterAccessWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterAccessWhiteList(request *ModifyDBClusterAccessWhiteListRequest) (_result *ModifyDBClusterAccessWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterAccessWhiteListResponse{}
	_body, _err := client.ModifyDBClusterAccessWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterDescriptionWithOptions(request *ModifyDBClusterDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterDescription)) {
		query["DBClusterDescription"] = request.DBClusterDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBClusterDescription"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBClusterDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterDescription(request *ModifyDBClusterDescriptionRequest) (_result *ModifyDBClusterDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterDescriptionResponse{}
	_body, _err := client.ModifyDBClusterDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterMaintainTimeWithOptions(request *ModifyDBClusterMaintainTimeRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterMaintainTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.MaintainTime)) {
		query["MaintainTime"] = request.MaintainTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBClusterMaintainTime"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBClusterMaintainTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterMaintainTime(request *ModifyDBClusterMaintainTimeRequest) (_result *ModifyDBClusterMaintainTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterMaintainTimeResponse{}
	_body, _err := client.ModifyDBClusterMaintainTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterResourceGroupWithOptions(request *ModifyDBClusterResourceGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.NewResourceGroupId)) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBClusterResourceGroup"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBClusterResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterResourceGroup(request *ModifyDBClusterResourceGroupRequest) (_result *ModifyDBClusterResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterResourceGroupResponse{}
	_body, _err := client.ModifyDBClusterResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBResourcePoolWithOptions(request *ModifyDBResourcePoolRequest, runtime *util.RuntimeOptions) (_result *ModifyDBResourcePoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeNum)) {
		query["NodeNum"] = request.NodeNum
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolName)) {
		query["PoolName"] = request.PoolName
	}

	if !tea.BoolValue(util.IsUnset(request.QueryType)) {
		query["QueryType"] = request.QueryType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBResourcePool"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBResourcePoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBResourcePool(request *ModifyDBResourcePoolRequest) (_result *ModifyDBResourcePoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBResourcePoolResponse{}
	_body, _err := client.ModifyDBResourcePoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyElasticPlanWithOptions(request *ModifyElasticPlanRequest, runtime *util.RuntimeOptions) (_result *ModifyElasticPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanEnable)) {
		query["ElasticPlanEnable"] = request.ElasticPlanEnable
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanEndDay)) {
		query["ElasticPlanEndDay"] = request.ElasticPlanEndDay
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanName)) {
		query["ElasticPlanName"] = request.ElasticPlanName
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanNodeNum)) {
		query["ElasticPlanNodeNum"] = request.ElasticPlanNodeNum
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanStartDay)) {
		query["ElasticPlanStartDay"] = request.ElasticPlanStartDay
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanTimeEnd)) {
		query["ElasticPlanTimeEnd"] = request.ElasticPlanTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanTimeStart)) {
		query["ElasticPlanTimeStart"] = request.ElasticPlanTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanWeeklyRepeat)) {
		query["ElasticPlanWeeklyRepeat"] = request.ElasticPlanWeeklyRepeat
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourcePoolName)) {
		query["ResourcePoolName"] = request.ResourcePoolName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyElasticPlan"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyElasticPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyElasticPlan(request *ModifyElasticPlanRequest) (_result *ModifyElasticPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyElasticPlanResponse{}
	_body, _err := client.ModifyElasticPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyLogBackupPolicyWithOptions(request *ModifyLogBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifyLogBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EnableBackupLog)) {
		query["EnableBackupLog"] = request.EnableBackupLog
	}

	if !tea.BoolValue(util.IsUnset(request.LogBackupRetentionPeriod)) {
		query["LogBackupRetentionPeriod"] = request.LogBackupRetentionPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyLogBackupPolicy"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyLogBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyLogBackupPolicy(request *ModifyLogBackupPolicyRequest) (_result *ModifyLogBackupPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyLogBackupPolicyResponse{}
	_body, _err := client.ModifyLogBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyMaintenanceActionWithOptions(request *ModifyMaintenanceActionRequest, runtime *util.RuntimeOptions) (_result *ModifyMaintenanceActionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ids)) {
		query["Ids"] = request.Ids
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SwitchTime)) {
		query["SwitchTime"] = request.SwitchTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyMaintenanceAction"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyMaintenanceActionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyMaintenanceAction(request *ModifyMaintenanceActionRequest) (_result *ModifyMaintenanceActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyMaintenanceActionResponse{}
	_body, _err := client.ModifyMaintenanceActionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseClusterPublicConnectionWithOptions(request *ReleaseClusterPublicConnectionRequest, runtime *util.RuntimeOptions) (_result *ReleaseClusterPublicConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseClusterPublicConnection"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseClusterPublicConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseClusterPublicConnection(request *ReleaseClusterPublicConnectionRequest) (_result *ReleaseClusterPublicConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseClusterPublicConnectionResponse{}
	_body, _err := client.ReleaseClusterPublicConnectionWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AccountPassword)) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.AccountType)) {
		query["AccountType"] = request.AccountType
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetAccountPassword"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetAccountPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) RevokeOperatorPermissionWithOptions(request *RevokeOperatorPermissionRequest, runtime *util.RuntimeOptions) (_result *RevokeOperatorPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RevokeOperatorPermission"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RevokeOperatorPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RevokeOperatorPermission(request *RevokeOperatorPermissionRequest) (_result *RevokeOperatorPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevokeOperatorPermissionResponse{}
	_body, _err := client.RevokeOperatorPermissionWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) UnbindDBResourcePoolWithUserWithOptions(request *UnbindDBResourcePoolWithUserRequest, runtime *util.RuntimeOptions) (_result *UnbindDBResourcePoolWithUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolName)) {
		query["PoolName"] = request.PoolName
	}

	if !tea.BoolValue(util.IsUnset(request.PoolUser)) {
		query["PoolUser"] = request.PoolUser
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindDBResourcePoolWithUser"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindDBResourcePoolWithUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindDBResourcePoolWithUser(request *UnbindDBResourcePoolWithUserRequest) (_result *UnbindDBResourcePoolWithUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindDBResourcePoolWithUserResponse{}
	_body, _err := client.UnbindDBResourcePoolWithUserWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
