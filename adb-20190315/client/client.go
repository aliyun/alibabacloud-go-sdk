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

type AllocateClusterPublicConnectionRequest struct {
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId            *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
}

func (s AllocateClusterPublicConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s AllocateClusterPublicConnectionRequest) GoString() string {
	return s.String()
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

func (s *AllocateClusterPublicConnectionRequest) SetOwnerAccount(v string) *AllocateClusterPublicConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AllocateClusterPublicConnectionRequest) SetDBClusterId(v string) *AllocateClusterPublicConnectionRequest {
	s.DBClusterId = &v
	return s
}

func (s *AllocateClusterPublicConnectionRequest) SetConnectionStringPrefix(v string) *AllocateClusterPublicConnectionRequest {
	s.ConnectionStringPrefix = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	PoolName             *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	PoolUser             *string `json:"PoolUser,omitempty" xml:"PoolUser,omitempty"`
}

func (s BindDBResourcePoolWithUserRequest) String() string {
	return tea.Prettify(s)
}

func (s BindDBResourcePoolWithUserRequest) GoString() string {
	return s.String()
}

func (s *BindDBResourcePoolWithUserRequest) SetOwnerId(v int64) *BindDBResourcePoolWithUserRequest {
	s.OwnerId = &v
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

func (s *BindDBResourcePoolWithUserRequest) SetOwnerAccount(v string) *BindDBResourcePoolWithUserRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BindDBResourcePoolWithUserRequest) SetDBClusterId(v string) *BindDBResourcePoolWithUserRequest {
	s.DBClusterId = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPassword      *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	AccountDescription   *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	AccountType          *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s CreateAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountRequest) GoString() string {
	return s.String()
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

func (s *CreateAccountRequest) SetDBClusterId(v string) *CreateAccountRequest {
	s.DBClusterId = &v
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

func (s *CreateAccountRequest) SetAccountDescription(v string) *CreateAccountRequest {
	s.AccountDescription = &v
	return s
}

func (s *CreateAccountRequest) SetAccountType(v string) *CreateAccountRequest {
	s.AccountType = &v
	return s
}

type CreateAccountResponseBody struct {
	TaskId      *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s CreateAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccountResponseBody) SetTaskId(v int32) *CreateAccountResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateAccountResponseBody) SetRequestId(v string) *CreateAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccountResponseBody) SetDBClusterId(v string) *CreateAccountResponseBody {
	s.DBClusterId = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	DBClusterVersion     *string `json:"DBClusterVersion,omitempty" xml:"DBClusterVersion,omitempty"`
	DBClusterCategory    *string `json:"DBClusterCategory,omitempty" xml:"DBClusterCategory,omitempty"`
	DBClusterClass       *string `json:"DBClusterClass,omitempty" xml:"DBClusterClass,omitempty"`
	DBNodeGroupCount     *string `json:"DBNodeGroupCount,omitempty" xml:"DBNodeGroupCount,omitempty"`
	DBNodeStorage        *string `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	DBClusterNetworkType *string `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	PayType              *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Period               *string `json:"Period,omitempty" xml:"Period,omitempty"`
	UsedTime             *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	VPCId                *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ExecutorCount        *string `json:"ExecutorCount,omitempty" xml:"ExecutorCount,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Mode                 *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	StorageResource      *string `json:"StorageResource,omitempty" xml:"StorageResource,omitempty"`
	StorageType          *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	ComputeResource      *string `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	RestoreType          *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	SourceDBInstanceName *string `json:"SourceDBInstanceName,omitempty" xml:"SourceDBInstanceName,omitempty"`
	BackupSetID          *string `json:"BackupSetID,omitempty" xml:"BackupSetID,omitempty"`
	RestoreTime          *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	ElasticIOResource    *string `json:"ElasticIOResource,omitempty" xml:"ElasticIOResource,omitempty"`
}

func (s CreateDBClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateDBClusterRequest) SetOwnerId(v int64) *CreateDBClusterRequest {
	s.OwnerId = &v
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

func (s *CreateDBClusterRequest) SetOwnerAccount(v string) *CreateDBClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBClusterRequest) SetRegionId(v string) *CreateDBClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBClusterRequest) SetZoneId(v string) *CreateDBClusterRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBClusterVersion(v string) *CreateDBClusterRequest {
	s.DBClusterVersion = &v
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

func (s *CreateDBClusterRequest) SetDBNodeGroupCount(v string) *CreateDBClusterRequest {
	s.DBNodeGroupCount = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBNodeStorage(v string) *CreateDBClusterRequest {
	s.DBNodeStorage = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBClusterNetworkType(v string) *CreateDBClusterRequest {
	s.DBClusterNetworkType = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBClusterDescription(v string) *CreateDBClusterRequest {
	s.DBClusterDescription = &v
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

func (s *CreateDBClusterRequest) SetClientToken(v string) *CreateDBClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBClusterRequest) SetExecutorCount(v string) *CreateDBClusterRequest {
	s.ExecutorCount = &v
	return s
}

func (s *CreateDBClusterRequest) SetResourceGroupId(v string) *CreateDBClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBClusterRequest) SetMode(v string) *CreateDBClusterRequest {
	s.Mode = &v
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

func (s *CreateDBClusterRequest) SetComputeResource(v string) *CreateDBClusterRequest {
	s.ComputeResource = &v
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

func (s *CreateDBClusterRequest) SetBackupSetID(v string) *CreateDBClusterRequest {
	s.BackupSetID = &v
	return s
}

func (s *CreateDBClusterRequest) SetRestoreTime(v string) *CreateDBClusterRequest {
	s.RestoreTime = &v
	return s
}

func (s *CreateDBClusterRequest) SetElasticIOResource(v string) *CreateDBClusterRequest {
	s.ElasticIOResource = &v
	return s
}

type CreateDBClusterResponseBody struct {
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	DBClusterId     *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OrderId         *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateDBClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBClusterResponseBody) SetRequestId(v string) *CreateDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBClusterResponseBody) SetResourceGroupId(v string) *CreateDBClusterResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBClusterResponseBody) SetDBClusterId(v string) *CreateDBClusterResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBClusterResponseBody) SetOrderId(v string) *CreateDBClusterResponseBody {
	s.OrderId = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	PoolName             *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	QueryType            *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	NodeNum              *int32  `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
}

func (s CreateDBResourcePoolRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBResourcePoolRequest) GoString() string {
	return s.String()
}

func (s *CreateDBResourcePoolRequest) SetOwnerId(v int64) *CreateDBResourcePoolRequest {
	s.OwnerId = &v
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

func (s *CreateDBResourcePoolRequest) SetOwnerAccount(v string) *CreateDBResourcePoolRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBResourcePoolRequest) SetDBClusterId(v string) *CreateDBResourcePoolRequest {
	s.DBClusterId = &v
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

func (s *CreateDBResourcePoolRequest) SetNodeNum(v int32) *CreateDBResourcePoolRequest {
	s.NodeNum = &v
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
	OwnerId                 *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount    *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId         *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount            *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId             *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ElasticPlanName         *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	ResourcePoolName        *string `json:"ResourcePoolName,omitempty" xml:"ResourcePoolName,omitempty"`
	ElasticPlanNodeNum      *int32  `json:"ElasticPlanNodeNum,omitempty" xml:"ElasticPlanNodeNum,omitempty"`
	ElasticPlanTimeStart    *string `json:"ElasticPlanTimeStart,omitempty" xml:"ElasticPlanTimeStart,omitempty"`
	ElasticPlanTimeEnd      *string `json:"ElasticPlanTimeEnd,omitempty" xml:"ElasticPlanTimeEnd,omitempty"`
	ElasticPlanWeeklyRepeat *string `json:"ElasticPlanWeeklyRepeat,omitempty" xml:"ElasticPlanWeeklyRepeat,omitempty"`
	ElasticPlanStartDay     *string `json:"ElasticPlanStartDay,omitempty" xml:"ElasticPlanStartDay,omitempty"`
	ElasticPlanEndDay       *string `json:"ElasticPlanEndDay,omitempty" xml:"ElasticPlanEndDay,omitempty"`
	ElasticPlanEnable       *bool   `json:"ElasticPlanEnable,omitempty" xml:"ElasticPlanEnable,omitempty"`
}

func (s CreateElasticPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateElasticPlanRequest) GoString() string {
	return s.String()
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

func (s *CreateElasticPlanRequest) SetOwnerAccount(v string) *CreateElasticPlanRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateElasticPlanRequest) SetDBClusterId(v string) *CreateElasticPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanName(v string) *CreateElasticPlanRequest {
	s.ElasticPlanName = &v
	return s
}

func (s *CreateElasticPlanRequest) SetResourcePoolName(v string) *CreateElasticPlanRequest {
	s.ResourcePoolName = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanNodeNum(v int32) *CreateElasticPlanRequest {
	s.ElasticPlanNodeNum = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanTimeStart(v string) *CreateElasticPlanRequest {
	s.ElasticPlanTimeStart = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanTimeEnd(v string) *CreateElasticPlanRequest {
	s.ElasticPlanTimeEnd = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanWeeklyRepeat(v string) *CreateElasticPlanRequest {
	s.ElasticPlanWeeklyRepeat = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanStartDay(v string) *CreateElasticPlanRequest {
	s.ElasticPlanStartDay = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanEndDay(v string) *CreateElasticPlanRequest {
	s.ElasticPlanEndDay = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanEnable(v bool) *CreateElasticPlanRequest {
	s.ElasticPlanEnable = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountType          *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s DeleteAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountRequest) GoString() string {
	return s.String()
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

func (s *DeleteAccountRequest) SetDBClusterId(v string) *DeleteAccountRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteAccountRequest) SetAccountName(v string) *DeleteAccountRequest {
	s.AccountName = &v
	return s
}

func (s *DeleteAccountRequest) SetAccountType(v string) *DeleteAccountRequest {
	s.AccountType = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DeleteDBClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBClusterRequest) GoString() string {
	return s.String()
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

func (s *DeleteDBClusterRequest) SetOwnerAccount(v string) *DeleteDBClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDBClusterRequest) SetDBClusterId(v string) *DeleteDBClusterRequest {
	s.DBClusterId = &v
	return s
}

type DeleteDBClusterResponseBody struct {
	TaskId      *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DeleteDBClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterResponseBody) SetTaskId(v int32) *DeleteDBClusterResponseBody {
	s.TaskId = &v
	return s
}

func (s *DeleteDBClusterResponseBody) SetRequestId(v string) *DeleteDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBClusterResponseBody) SetDBClusterId(v string) *DeleteDBClusterResponseBody {
	s.DBClusterId = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	PoolName             *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
}

func (s DeleteDBResourcePoolRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBResourcePoolRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBResourcePoolRequest) SetOwnerId(v int64) *DeleteDBResourcePoolRequest {
	s.OwnerId = &v
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

func (s *DeleteDBResourcePoolRequest) SetOwnerAccount(v string) *DeleteDBResourcePoolRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDBResourcePoolRequest) SetDBClusterId(v string) *DeleteDBResourcePoolRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBResourcePoolRequest) SetPoolName(v string) *DeleteDBResourcePoolRequest {
	s.PoolName = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ElasticPlanName      *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
}

func (s DeleteElasticPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteElasticPlanRequest) GoString() string {
	return s.String()
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

func (s *DeleteElasticPlanRequest) SetOwnerAccount(v string) *DeleteElasticPlanRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteElasticPlanRequest) SetDBClusterId(v string) *DeleteElasticPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteElasticPlanRequest) SetElasticPlanName(v string) *DeleteElasticPlanRequest {
	s.ElasticPlanName = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountType          *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s DescribeAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsRequest) GoString() string {
	return s.String()
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

func (s *DescribeAccountsRequest) SetDBClusterId(v string) *DescribeAccountsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAccountsRequest) SetAccountName(v string) *DescribeAccountsRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountsRequest) SetAccountType(v string) *DescribeAccountsRequest {
	s.AccountType = &v
	return s
}

type DescribeAccountsResponseBody struct {
	RequestId   *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AccountList *DescribeAccountsResponseBodyAccountList `json:"AccountList,omitempty" xml:"AccountList,omitempty" type:"Struct"`
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

func (s *DescribeAccountsResponseBody) SetAccountList(v *DescribeAccountsResponseBodyAccountList) *DescribeAccountsResponseBody {
	s.AccountList = v
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
	AccountStatus      *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	AccountType        *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	AccountName        *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s DescribeAccountsResponseBodyAccountListDBAccount) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountListDBAccount) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) SetAccountStatus(v string) *DescribeAccountsResponseBodyAccountListDBAccount {
	s.AccountStatus = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) SetAccountDescription(v string) *DescribeAccountsResponseBodyAccountListDBAccount {
	s.AccountDescription = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) SetAccountType(v string) *DescribeAccountsResponseBodyAccountListDBAccount {
	s.AccountType = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) SetAccountName(v string) *DescribeAccountsResponseBodyAccountListDBAccount {
	s.AccountName = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeAllAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllAccountsRequest) GoString() string {
	return s.String()
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

func (s *DescribeAllAccountsRequest) SetOwnerAccount(v string) *DescribeAllAccountsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAllAccountsRequest) SetDBClusterId(v string) *DescribeAllAccountsRequest {
	s.DBClusterId = &v
	return s
}

type DescribeAllAccountsResponseBody struct {
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AccountList []*DescribeAllAccountsResponseBodyAccountList `json:"AccountList,omitempty" xml:"AccountList,omitempty" type:"Repeated"`
}

func (s DescribeAllAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAllAccountsResponseBody) SetRequestId(v string) *DescribeAllAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAllAccountsResponseBody) SetAccountList(v []*DescribeAllAccountsResponseBodyAccountList) *DescribeAllAccountsResponseBody {
	s.AccountList = v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	SchemaName           *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName            *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeAllDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceRequest) GoString() string {
	return s.String()
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

func (s *DescribeAllDataSourceRequest) SetOwnerAccount(v string) *DescribeAllDataSourceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAllDataSourceRequest) SetDBClusterId(v string) *DescribeAllDataSourceRequest {
	s.DBClusterId = &v
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
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tables    *DescribeAllDataSourceResponseBodyTables  `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Struct"`
	Columns   *DescribeAllDataSourceResponseBodyColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Struct"`
	Schemas   *DescribeAllDataSourceResponseBodySchemas `json:"Schemas,omitempty" xml:"Schemas,omitempty" type:"Struct"`
}

func (s DescribeAllDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBody) SetRequestId(v string) *DescribeAllDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAllDataSourceResponseBody) SetTables(v *DescribeAllDataSourceResponseBodyTables) *DescribeAllDataSourceResponseBody {
	s.Tables = v
	return s
}

func (s *DescribeAllDataSourceResponseBody) SetColumns(v *DescribeAllDataSourceResponseBodyColumns) *DescribeAllDataSourceResponseBody {
	s.Columns = v
	return s
}

func (s *DescribeAllDataSourceResponseBody) SetSchemas(v *DescribeAllDataSourceResponseBodySchemas) *DescribeAllDataSourceResponseBody {
	s.Schemas = v
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
	TableName   *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	SchemaName  *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
}

func (s DescribeAllDataSourceResponseBodyTablesTable) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceResponseBodyTablesTable) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBodyTablesTable) SetTableName(v string) *DescribeAllDataSourceResponseBodyTablesTable {
	s.TableName = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyTablesTable) SetDBClusterId(v string) *DescribeAllDataSourceResponseBodyTablesTable {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyTablesTable) SetSchemaName(v string) *DescribeAllDataSourceResponseBodyTablesTable {
	s.SchemaName = &v
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
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	ColumnName          *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	TableName           *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	AutoIncrementColumn *bool   `json:"AutoIncrementColumn,omitempty" xml:"AutoIncrementColumn,omitempty"`
	DBClusterId         *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	PrimaryKey          *bool   `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	SchemaName          *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
}

func (s DescribeAllDataSourceResponseBodyColumnsColumn) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceResponseBodyColumnsColumn) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetType(v string) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.Type = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetColumnName(v string) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.ColumnName = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetTableName(v string) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.TableName = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetAutoIncrementColumn(v bool) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.AutoIncrementColumn = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAuditLogConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogConfigRequest) SetOwnerId(v int64) *DescribeAuditLogConfigRequest {
	s.OwnerId = &v
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

func (s *DescribeAuditLogConfigRequest) SetOwnerAccount(v string) *DescribeAuditLogConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAuditLogConfigRequest) SetDBClusterId(v string) *DescribeAuditLogConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAuditLogConfigRequest) SetRegionId(v string) *DescribeAuditLogConfigRequest {
	s.RegionId = &v
	return s
}

type DescribeAuditLogConfigResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AuditLogStatus *string `json:"AuditLogStatus,omitempty" xml:"AuditLogStatus,omitempty"`
	DBClusterId    *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeAuditLogConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogConfigResponseBody) SetRequestId(v string) *DescribeAuditLogConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuditLogConfigResponseBody) SetAuditLogStatus(v string) *DescribeAuditLogConfigResponseBody {
	s.AuditLogStatus = &v
	return s
}

func (s *DescribeAuditLogConfigResponseBody) SetDBClusterId(v string) *DescribeAuditLogConfigResponseBody {
	s.DBClusterId = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	QueryKeyword         *string `json:"QueryKeyword,omitempty" xml:"QueryKeyword,omitempty"`
	SqlType              *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	Succeed              *string `json:"Succeed,omitempty" xml:"Succeed,omitempty"`
	HostAddress          *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	OrderType            *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	User                 *string `json:"User,omitempty" xml:"User,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeAuditLogRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogRecordsRequest) SetOwnerId(v int64) *DescribeAuditLogRecordsRequest {
	s.OwnerId = &v
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

func (s *DescribeAuditLogRecordsRequest) SetOwnerAccount(v string) *DescribeAuditLogRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetDBClusterId(v string) *DescribeAuditLogRecordsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetRegionId(v string) *DescribeAuditLogRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetStartTime(v string) *DescribeAuditLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetEndTime(v string) *DescribeAuditLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetDBName(v string) *DescribeAuditLogRecordsRequest {
	s.DBName = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetQueryKeyword(v string) *DescribeAuditLogRecordsRequest {
	s.QueryKeyword = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetSqlType(v string) *DescribeAuditLogRecordsRequest {
	s.SqlType = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetSucceed(v string) *DescribeAuditLogRecordsRequest {
	s.Succeed = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetHostAddress(v string) *DescribeAuditLogRecordsRequest {
	s.HostAddress = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetOrderType(v string) *DescribeAuditLogRecordsRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetUser(v string) *DescribeAuditLogRecordsRequest {
	s.User = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetPageSize(v int32) *DescribeAuditLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetPageNumber(v int32) *DescribeAuditLogRecordsRequest {
	s.PageNumber = &v
	return s
}

type DescribeAuditLogRecordsResponseBody struct {
	TotalCount  *string                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize    *string                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber  *string                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	DBClusterId *string                                     `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Items       []*DescribeAuditLogRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeAuditLogRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogRecordsResponseBody) SetTotalCount(v string) *DescribeAuditLogRecordsResponseBody {
	s.TotalCount = &v
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

func (s *DescribeAuditLogRecordsResponseBody) SetPageNumber(v string) *DescribeAuditLogRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBody) SetDBClusterId(v string) *DescribeAuditLogRecordsResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBody) SetItems(v []*DescribeAuditLogRecordsResponseBodyItems) *DescribeAuditLogRecordsResponseBody {
	s.Items = v
	return s
}

type DescribeAuditLogRecordsResponseBodyItems struct {
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	Succeed     *string `json:"Succeed,omitempty" xml:"Succeed,omitempty"`
	SQLText     *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	TotalTime   *string `json:"TotalTime,omitempty" xml:"TotalTime,omitempty"`
	ConnId      *string `json:"ConnId,omitempty" xml:"ConnId,omitempty"`
	DBName      *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	SQLType     *string `json:"SQLType,omitempty" xml:"SQLType,omitempty"`
	ExecuteTime *string `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	ProcessID   *string `json:"ProcessID,omitempty" xml:"ProcessID,omitempty"`
}

func (s DescribeAuditLogRecordsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetHostAddress(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.HostAddress = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetSucceed(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.Succeed = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetSQLText(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.SQLText = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetTotalTime(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.TotalTime = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetConnId(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.ConnId = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetDBName(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.DBName = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetSQLType(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.SQLType = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetExecuteTime(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetProcessID(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.ProcessID = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBClusterIds         *string `json:"DBClusterIds,omitempty" xml:"DBClusterIds,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeAutoRenewAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewAttributeRequest) SetOwnerId(v int64) *DescribeAutoRenewAttributeRequest {
	s.OwnerId = &v
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

func (s *DescribeAutoRenewAttributeRequest) SetOwnerAccount(v string) *DescribeAutoRenewAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetRegionId(v string) *DescribeAutoRenewAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetDBClusterIds(v string) *DescribeAutoRenewAttributeRequest {
	s.DBClusterIds = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetPageSize(v int32) *DescribeAutoRenewAttributeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetPageNumber(v int32) *DescribeAutoRenewAttributeRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetResourceGroupId(v string) *DescribeAutoRenewAttributeRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeAutoRenewAttributeResponseBody struct {
	TotalRecordCount *int32                                       `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	PageRecordCount  *int32                                       `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber       *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Items            *DescribeAutoRenewAttributeResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeAutoRenewAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoRenewAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewAttributeResponseBody) SetTotalRecordCount(v int32) *DescribeAutoRenewAttributeResponseBody {
	s.TotalRecordCount = &v
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

func (s *DescribeAutoRenewAttributeResponseBody) SetPageNumber(v int32) *DescribeAutoRenewAttributeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBody) SetItems(v *DescribeAutoRenewAttributeResponseBodyItems) *DescribeAutoRenewAttributeResponseBody {
	s.Items = v
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
	DBClusterId      *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	PeriodUnit       *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	Duration         *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	RenewalStatus    *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
	AutoRenewEnabled *bool   `json:"AutoRenewEnabled,omitempty" xml:"AutoRenewEnabled,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetDBClusterId(v string) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetPeriodUnit(v string) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetDuration(v int32) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.Duration = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetRenewalStatus(v string) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.RenewalStatus = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetAutoRenewEnabled(v bool) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.AutoRenewEnabled = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetRegionId(v string) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.RegionId = &v
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
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ChargeType           *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	AcceptLanguage       *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeAvailableResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceRequest) SetRegionId(v string) *DescribeAvailableResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetZoneId(v string) *DescribeAvailableResourceRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetChargeType(v string) *DescribeAvailableResourceRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetOwnerId(v int64) *DescribeAvailableResourceRequest {
	s.OwnerId = &v
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

func (s *DescribeAvailableResourceRequest) SetOwnerAccount(v string) *DescribeAvailableResourceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetAcceptLanguage(v string) *DescribeAvailableResourceRequest {
	s.AcceptLanguage = &v
	return s
}

type DescribeAvailableResourceResponseBody struct {
	RequestId         *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RegionId          *string                                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AvailableZoneList []*DescribeAvailableResourceResponseBodyAvailableZoneList `json:"AvailableZoneList,omitempty" xml:"AvailableZoneList,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBody) SetRequestId(v string) *DescribeAvailableResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBody) SetRegionId(v string) *DescribeAvailableResourceResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBody) SetAvailableZoneList(v []*DescribeAvailableResourceResponseBodyAvailableZoneList) *DescribeAvailableResourceResponseBody {
	s.AvailableZoneList = v
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
	SupportedSerialList []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList `json:"SupportedSerialList,omitempty" xml:"SupportedSerialList,omitempty" type:"Repeated"`
	Mode                *string                                                                                   `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode) SetSupportedSerialList(v []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode {
	s.SupportedSerialList = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode) SetMode(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode {
	s.Mode = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList struct {
	Serial                     *string                                                                                                             `json:"Serial,omitempty" xml:"Serial,omitempty"`
	SupportedInstanceClassList []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList `json:"SupportedInstanceClassList,omitempty" xml:"SupportedInstanceClassList,omitempty" type:"Repeated"`
	SupportedFlexibleResource  []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource  `json:"SupportedFlexibleResource,omitempty" xml:"SupportedFlexibleResource,omitempty" type:"Repeated"`
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

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList) SetSupportedInstanceClassList(v []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList {
	s.SupportedInstanceClassList = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList) SetSupportedFlexibleResource(v []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList {
	s.SupportedFlexibleResource = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList struct {
	SupportedExecutorList  []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorList  `json:"SupportedExecutorList,omitempty" xml:"SupportedExecutorList,omitempty" type:"Repeated"`
	InstanceClass          *string                                                                                                                                   `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	SupportedNodeCountList []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList `json:"SupportedNodeCountList,omitempty" xml:"SupportedNodeCountList,omitempty" type:"Repeated"`
	Tips                   *string                                                                                                                                   `json:"Tips,omitempty" xml:"Tips,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) SetSupportedExecutorList(v []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorList) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList {
	s.SupportedExecutorList = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) SetInstanceClass(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList {
	s.InstanceClass = &v
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
	Step     *string `json:"Step,omitempty" xml:"Step,omitempty"`
	MinCount *string `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	MaxCount *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount) SetStep(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount {
	s.Step = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount) SetMinCount(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount {
	s.MinCount = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount) SetMaxCount(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount {
	s.MaxCount = &v
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
	Step     *string `json:"Step,omitempty" xml:"Step,omitempty"`
	MinCount *string `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	MaxCount *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount) SetStep(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount {
	s.Step = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount) SetMinCount(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount {
	s.MinCount = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount) SetMaxCount(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount {
	s.MaxCount = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource struct {
	SupportedElasticIOResource *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource `json:"SupportedElasticIOResource,omitempty" xml:"SupportedElasticIOResource,omitempty" type:"Struct"`
	SupportedStorageResource   []*string                                                                                                                                  `json:"SupportedStorageResource,omitempty" xml:"SupportedStorageResource,omitempty" type:"Repeated"`
	StorageType                *string                                                                                                                                    `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	SupportedComputeResource   []*string                                                                                                                                  `json:"SupportedComputeResource,omitempty" xml:"SupportedComputeResource,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) SetSupportedElasticIOResource(v *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource {
	s.SupportedElasticIOResource = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) SetSupportedStorageResource(v []*string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource {
	s.SupportedStorageResource = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) SetStorageType(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource {
	s.StorageType = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) SetSupportedComputeResource(v []*string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource {
	s.SupportedComputeResource = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource struct {
	Step     *string `json:"Step,omitempty" xml:"Step,omitempty"`
	MinCount *string `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	MaxCount *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource) SetStep(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource {
	s.Step = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource) SetMinCount(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource {
	s.MinCount = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource) SetMaxCount(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource {
	s.MaxCount = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyRequest) GoString() string {
	return s.String()
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

func (s *DescribeBackupPolicyRequest) SetDBClusterId(v string) *DescribeBackupPolicyRequest {
	s.DBClusterId = &v
	return s
}

type DescribeBackupPolicyResponseBody struct {
	PreferredBackupPeriod    *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	LogBackupRetentionPeriod *int32  `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	RequestId                *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PreferredBackupTime      *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	EnableBackupLog          *string `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	BackupRetentionPeriod    *int32  `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
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

func (s *DescribeBackupPolicyResponseBody) SetLogBackupRetentionPeriod(v int32) *DescribeBackupPolicyResponseBody {
	s.LogBackupRetentionPeriod = &v
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

func (s *DescribeBackupPolicyResponseBody) SetEnableBackupLog(v string) *DescribeBackupPolicyResponseBody {
	s.EnableBackupLog = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetBackupRetentionPeriod(v int32) *DescribeBackupPolicyResponseBody {
	s.BackupRetentionPeriod = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	BackupId             *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeBackupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsRequest) GoString() string {
	return s.String()
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

func (s *DescribeBackupsRequest) SetDBClusterId(v string) *DescribeBackupsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeBackupsRequest) SetBackupId(v string) *DescribeBackupsRequest {
	s.BackupId = &v
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

func (s *DescribeBackupsRequest) SetPageSize(v int32) *DescribeBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupsRequest) SetPageNumber(v int32) *DescribeBackupsRequest {
	s.PageNumber = &v
	return s
}

type DescribeBackupsResponseBody struct {
	TotalCount *string                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *string                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *string                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Items      *DescribeBackupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeBackupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBody) SetTotalCount(v string) *DescribeBackupsResponseBody {
	s.TotalCount = &v
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

func (s *DescribeBackupsResponseBody) SetPageNumber(v string) *DescribeBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetItems(v *DescribeBackupsResponseBodyItems) *DescribeBackupsResponseBody {
	s.Items = v
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
	DBClusterId     *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	BackupType      *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	BackupSize      *int32  `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	BackupEndTime   *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	BackupId        *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	BackupMethod    *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
}

func (s DescribeBackupsResponseBodyItemsBackup) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseBodyItemsBackup) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetDBClusterId(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.DBClusterId = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupType(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupStartTime(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupSize(v int32) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupSize = &v
	return s
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	SchemaName           *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName            *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeColumnsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeColumnsRequest) GoString() string {
	return s.String()
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

func (s *DescribeColumnsRequest) SetOwnerAccount(v string) *DescribeColumnsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeColumnsRequest) SetDBClusterId(v string) *DescribeColumnsRequest {
	s.DBClusterId = &v
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
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Items     *DescribeColumnsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeColumnsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeColumnsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeColumnsResponseBody) SetRequestId(v string) *DescribeColumnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeColumnsResponseBody) SetItems(v *DescribeColumnsResponseBodyItems) *DescribeColumnsResponseBody {
	s.Items = v
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
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	ColumnName          *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	TableName           *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	AutoIncrementColumn *bool   `json:"AutoIncrementColumn,omitempty" xml:"AutoIncrementColumn,omitempty"`
	DBClusterId         *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	PrimaryKey          *bool   `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	SchemaName          *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
}

func (s DescribeColumnsResponseBodyItemsColumn) String() string {
	return tea.Prettify(s)
}

func (s DescribeColumnsResponseBodyItemsColumn) GoString() string {
	return s.String()
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetType(v string) *DescribeColumnsResponseBodyItemsColumn {
	s.Type = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetColumnName(v string) *DescribeColumnsResponseBodyItemsColumn {
	s.ColumnName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetTableName(v string) *DescribeColumnsResponseBodyItemsColumn {
	s.TableName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetAutoIncrementColumn(v bool) *DescribeColumnsResponseBodyItemsColumn {
	s.AutoIncrementColumn = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeConnectionCountRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectionCountRecordsRequest) GoString() string {
	return s.String()
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

func (s *DescribeConnectionCountRecordsRequest) SetOwnerAccount(v string) *DescribeConnectionCountRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeConnectionCountRecordsRequest) SetDBClusterId(v string) *DescribeConnectionCountRecordsRequest {
	s.DBClusterId = &v
	return s
}

type DescribeConnectionCountRecordsResponseBody struct {
	TotalCount      *string                                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId       *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DBClusterId     *string                                                      `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	AccessIpRecords []*DescribeConnectionCountRecordsResponseBodyAccessIpRecords `json:"AccessIpRecords,omitempty" xml:"AccessIpRecords,omitempty" type:"Repeated"`
	UserRecords     []*DescribeConnectionCountRecordsResponseBodyUserRecords     `json:"UserRecords,omitempty" xml:"UserRecords,omitempty" type:"Repeated"`
}

func (s DescribeConnectionCountRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectionCountRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConnectionCountRecordsResponseBody) SetTotalCount(v string) *DescribeConnectionCountRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeConnectionCountRecordsResponseBody) SetRequestId(v string) *DescribeConnectionCountRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConnectionCountRecordsResponseBody) SetDBClusterId(v string) *DescribeConnectionCountRecordsResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeConnectionCountRecordsResponseBody) SetAccessIpRecords(v []*DescribeConnectionCountRecordsResponseBodyAccessIpRecords) *DescribeConnectionCountRecordsResponseBody {
	s.AccessIpRecords = v
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
	User  *string `json:"User,omitempty" xml:"User,omitempty"`
	Count *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeConnectionCountRecordsResponseBodyUserRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectionCountRecordsResponseBodyUserRecords) GoString() string {
	return s.String()
}

func (s *DescribeConnectionCountRecordsResponseBodyUserRecords) SetUser(v string) *DescribeConnectionCountRecordsResponseBodyUserRecords {
	s.User = &v
	return s
}

func (s *DescribeConnectionCountRecordsResponseBodyUserRecords) SetCount(v int64) *DescribeConnectionCountRecordsResponseBodyUserRecords {
	s.Count = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeDBClusterAccessWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAccessWhiteListRequest) GoString() string {
	return s.String()
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

func (s *DescribeDBClusterAccessWhiteListRequest) SetOwnerAccount(v string) *DescribeDBClusterAccessWhiteListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListRequest) SetDBClusterId(v string) *DescribeDBClusterAccessWhiteListRequest {
	s.DBClusterId = &v
	return s
}

type DescribeDBClusterAccessWhiteListResponseBody struct {
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Items     *DescribeDBClusterAccessWhiteListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeDBClusterAccessWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAccessWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhiteListResponseBody) SetRequestId(v string) *DescribeDBClusterAccessWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListResponseBody) SetItems(v *DescribeDBClusterAccessWhiteListResponseBodyItems) *DescribeDBClusterAccessWhiteListResponseBody {
	s.Items = v
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
	SecurityIPList            *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	DBClusterIPArrayName      *string `json:"DBClusterIPArrayName,omitempty" xml:"DBClusterIPArrayName,omitempty"`
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

func (s *DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray) SetSecurityIPList(v string) *DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray {
	s.SecurityIPList = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray) SetDBClusterIPArrayName(v string) *DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray {
	s.DBClusterIPArrayName = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeDBClusterAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeRequest) GoString() string {
	return s.String()
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

func (s *DescribeDBClusterAttributeRequest) SetOwnerAccount(v string) *DescribeDBClusterAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAttributeRequest) SetDBClusterId(v string) *DescribeDBClusterAttributeRequest {
	s.DBClusterId = &v
	return s
}

type DescribeDBClusterAttributeResponseBody struct {
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Items     *DescribeDBClusterAttributeResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeDBClusterAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBody) SetRequestId(v string) *DescribeDBClusterAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetItems(v *DescribeDBClusterAttributeResponseBodyItems) *DescribeDBClusterAttributeResponseBody {
	s.Items = v
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
	CreationTime         *string                                                   `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	EnableSpark          *bool                                                     `json:"EnableSpark,omitempty" xml:"EnableSpark,omitempty"`
	DtsJobId             *string                                                   `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	DBNodeCount          *int64                                                    `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	Expired              *string                                                   `json:"Expired,omitempty" xml:"Expired,omitempty"`
	MaintainTime         *string                                                   `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
	PayType              *string                                                   `json:"PayType,omitempty" xml:"PayType,omitempty"`
	DiskType             *string                                                   `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	Tags                 *DescribeDBClusterAttributeResponseBodyItemsDBClusterTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	Mode                 *string                                                   `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Port                 *int32                                                    `json:"Port,omitempty" xml:"Port,omitempty"`
	LockMode             *string                                                   `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	EngineVersion        *string                                                   `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	EnableAirflow        *bool                                                     `json:"EnableAirflow,omitempty" xml:"EnableAirflow,omitempty"`
	ExecutorCount        *string                                                   `json:"ExecutorCount,omitempty" xml:"ExecutorCount,omitempty"`
	StorageResource      *string                                                   `json:"StorageResource,omitempty" xml:"StorageResource,omitempty"`
	DBClusterId          *string                                                   `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ConnectionString     *string                                                   `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	RdsInstanceId        *string                                                   `json:"RdsInstanceId,omitempty" xml:"RdsInstanceId,omitempty"`
	DBClusterType        *string                                                   `json:"DBClusterType,omitempty" xml:"DBClusterType,omitempty"`
	CommodityCode        *string                                                   `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ExpireTime           *string                                                   `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	DBNodeStorage        *int64                                                    `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	DBNodeClass          *string                                                   `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	LockReason           *string                                                   `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	VPCId                *string                                                   `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	ComputeResource      *string                                                   `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	RegionId             *string                                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ElasticIOResource    *int32                                                    `json:"ElasticIOResource,omitempty" xml:"ElasticIOResource,omitempty"`
	VSwitchId            *string                                                   `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	DBVersion            *string                                                   `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	VPCCloudInstanceId   *string                                                   `json:"VPCCloudInstanceId,omitempty" xml:"VPCCloudInstanceId,omitempty"`
	DBClusterStatus      *string                                                   `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	ResourceGroupId      *string                                                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	DBClusterNetworkType *string                                                   `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	DBClusterDescription *string                                                   `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	UserENIStatus        *bool                                                     `json:"UserENIStatus,omitempty" xml:"UserENIStatus,omitempty"`
	ZoneId               *string                                                   `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Category             *string                                                   `json:"Category,omitempty" xml:"Category,omitempty"`
	Engine               *string                                                   `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBCluster) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBCluster) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetCreationTime(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.CreationTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetEnableSpark(v bool) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.EnableSpark = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDtsJobId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBNodeCount(v int64) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBNodeCount = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetExpired(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Expired = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetMaintainTime(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.MaintainTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetPayType(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.PayType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDiskType(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DiskType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetTags(v *DescribeDBClusterAttributeResponseBodyItemsDBClusterTags) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Tags = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetMode(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Mode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetPort(v int32) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Port = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetLockMode(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.LockMode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetEngineVersion(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetEnableAirflow(v bool) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.EnableAirflow = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetExecutorCount(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ExecutorCount = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetStorageResource(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.StorageResource = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBClusterId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetConnectionString(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetRdsInstanceId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.RdsInstanceId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBClusterType(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBClusterType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetCommodityCode(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetExpireTime(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBNodeStorage(v int64) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBNodeStorage = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBNodeClass(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetLockReason(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.LockReason = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetVPCId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.VPCId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetComputeResource(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ComputeResource = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetRegionId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetElasticIOResource(v int32) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ElasticIOResource = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetVSwitchId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBVersion(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetVPCCloudInstanceId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.VPCCloudInstanceId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBClusterStatus(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetResourceGroupId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBClusterNetworkType(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBClusterNetworkType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBClusterDescription(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetUserENIStatus(v bool) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.UserENIStatus = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetZoneId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetCategory(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Category = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetEngine(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Engine = &v
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

type DescribeDBClusterNetInfoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeDBClusterNetInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterNetInfoRequest) GoString() string {
	return s.String()
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

func (s *DescribeDBClusterNetInfoRequest) SetOwnerAccount(v string) *DescribeDBClusterNetInfoRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterNetInfoRequest) SetDBClusterId(v string) *DescribeDBClusterNetInfoRequest {
	s.DBClusterId = &v
	return s
}

type DescribeDBClusterNetInfoResponseBody struct {
	ClusterNetworkType *string                                    `json:"ClusterNetworkType,omitempty" xml:"ClusterNetworkType,omitempty"`
	RequestId          *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Items              *DescribeDBClusterNetInfoResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
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

func (s *DescribeDBClusterNetInfoResponseBody) SetRequestId(v string) *DescribeDBClusterNetInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBody) SetItems(v *DescribeDBClusterNetInfoResponseBodyItems) *DescribeDBClusterNetInfoResponseBody {
	s.Items = v
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
	VSwitchId              *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ConnectionString       *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	NetType                *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	Port                   *string `json:"Port,omitempty" xml:"Port,omitempty"`
	VPCId                  *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	IPAddress              *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
}

func (s DescribeDBClusterNetInfoResponseBodyItemsAddress) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterNetInfoResponseBodyItemsAddress) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) SetVSwitchId(v string) *DescribeDBClusterNetInfoResponseBodyItemsAddress {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) SetConnectionString(v string) *DescribeDBClusterNetInfoResponseBodyItemsAddress {
	s.ConnectionString = &v
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

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) SetIPAddress(v string) *DescribeDBClusterNetInfoResponseBodyItemsAddress {
	s.IPAddress = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) SetConnectionStringPrefix(v string) *DescribeDBClusterNetInfoResponseBodyItemsAddress {
	s.ConnectionStringPrefix = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Key                  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeDBClusterPerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceRequest) GoString() string {
	return s.String()
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

func (s *DescribeDBClusterPerformanceRequest) SetOwnerAccount(v string) *DescribeDBClusterPerformanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetDBClusterId(v string) *DescribeDBClusterPerformanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetKey(v string) *DescribeDBClusterPerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetStartTime(v string) *DescribeDBClusterPerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetEndTime(v string) *DescribeDBClusterPerformanceRequest {
	s.EndTime = &v
	return s
}

type DescribeDBClusterPerformanceResponseBody struct {
	EndTime      *string                                                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId    *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime    *string                                                 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	DBClusterId  *string                                                 `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Performances []*DescribeDBClusterPerformanceResponseBodyPerformances `json:"Performances,omitempty" xml:"Performances,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterPerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBody) SetEndTime(v string) *DescribeDBClusterPerformanceResponseBody {
	s.EndTime = &v
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

func (s *DescribeDBClusterPerformanceResponseBody) SetDBClusterId(v string) *DescribeDBClusterPerformanceResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetPerformances(v []*DescribeDBClusterPerformanceResponseBodyPerformances) *DescribeDBClusterPerformanceResponseBody {
	s.Performances = v
	return s
}

type DescribeDBClusterPerformanceResponseBodyPerformances struct {
	Key    *string                                                       `json:"Key,omitempty" xml:"Key,omitempty"`
	Unit   *string                                                       `json:"Unit,omitempty" xml:"Unit,omitempty"`
	Series []*DescribeDBClusterPerformanceResponseBodyPerformancesSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
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

func (s *DescribeDBClusterPerformanceResponseBodyPerformances) SetUnit(v string) *DescribeDBClusterPerformanceResponseBodyPerformances {
	s.Unit = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformances) SetSeries(v []*DescribeDBClusterPerformanceResponseBodyPerformancesSeries) *DescribeDBClusterPerformanceResponseBodyPerformances {
	s.Series = v
	return s
}

type DescribeDBClusterPerformanceResponseBodyPerformancesSeries struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
	Name   *string   `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDBClusterPerformanceResponseBodyPerformancesSeries) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBodyPerformancesSeries) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformancesSeries) SetValues(v []*string) *DescribeDBClusterPerformanceResponseBodyPerformancesSeries {
	s.Values = v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformancesSeries) SetName(v string) *DescribeDBClusterPerformanceResponseBodyPerformancesSeries {
	s.Name = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Key                  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	ResourcePools        *string `json:"ResourcePools,omitempty" xml:"ResourcePools,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeDBClusterResourcePoolPerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterResourcePoolPerformanceRequest) GoString() string {
	return s.String()
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

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetOwnerAccount(v string) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetDBClusterId(v string) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetKey(v string) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.Key = &v
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

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetEndTime(v string) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.EndTime = &v
	return s
}

type DescribeDBClusterResourcePoolPerformanceResponseBody struct {
	EndTime      *string                                                             `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId    *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime    *string                                                             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	DBClusterId  *string                                                             `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Performances []*DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances `json:"Performances,omitempty" xml:"Performances,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterResourcePoolPerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterResourcePoolPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBody) SetEndTime(v string) *DescribeDBClusterResourcePoolPerformanceResponseBody {
	s.EndTime = &v
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

func (s *DescribeDBClusterResourcePoolPerformanceResponseBody) SetDBClusterId(v string) *DescribeDBClusterResourcePoolPerformanceResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBody) SetPerformances(v []*DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances) *DescribeDBClusterResourcePoolPerformanceResponseBody {
	s.Performances = v
	return s
}

type DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances struct {
	Key                      *string                                                                                     `json:"Key,omitempty" xml:"Key,omitempty"`
	Unit                     *string                                                                                     `json:"Unit,omitempty" xml:"Unit,omitempty"`
	ResourcePoolPerformances []*DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances `json:"ResourcePoolPerformances,omitempty" xml:"ResourcePoolPerformances,omitempty" type:"Repeated"`
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

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances) SetUnit(v string) *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances {
	s.Unit = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances) SetResourcePoolPerformances(v []*DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances) *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances {
	s.ResourcePoolPerformances = v
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
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
	Name   *string   `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries) SetValues(v []*string) *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries {
	s.Values = v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries) SetName(v string) *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries {
	s.Name = &v
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
	OwnerId              *int64                          `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                         `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                          `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string                         `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBClusterIds         *string                         `json:"DBClusterIds,omitempty" xml:"DBClusterIds,omitempty"`
	DBClusterDescription *string                         `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	DBClusterStatus      *string                         `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	PageSize             *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ResourceGroupId      *string                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag                  []*DescribeDBClustersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersRequest) SetOwnerId(v int64) *DescribeDBClustersRequest {
	s.OwnerId = &v
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

func (s *DescribeDBClustersRequest) SetOwnerAccount(v string) *DescribeDBClustersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClustersRequest) SetRegionId(v string) *DescribeDBClustersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDBClusterIds(v string) *DescribeDBClustersRequest {
	s.DBClusterIds = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDBClusterDescription(v string) *DescribeDBClustersRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDBClusterStatus(v string) *DescribeDBClustersRequest {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClustersRequest) SetPageSize(v int32) *DescribeDBClustersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBClustersRequest) SetPageNumber(v int32) *DescribeDBClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClustersRequest) SetResourceGroupId(v string) *DescribeDBClustersRequest {
	s.ResourceGroupId = &v
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
	TotalCount *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Items      *DescribeDBClustersResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeDBClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBody) SetTotalCount(v int32) *DescribeDBClustersResponseBody {
	s.TotalCount = &v
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

func (s *DescribeDBClustersResponseBody) SetPageNumber(v int32) *DescribeDBClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClustersResponseBody) SetItems(v *DescribeDBClustersResponseBodyItems) *DescribeDBClustersResponseBody {
	s.Items = v
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
	DtsJobId             *string                                           `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	DBNodeCount          *int64                                            `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	Expired              *string                                           `json:"Expired,omitempty" xml:"Expired,omitempty"`
	CreateTime           *string                                           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	PayType              *string                                           `json:"PayType,omitempty" xml:"PayType,omitempty"`
	DiskType             *string                                           `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	Tags                 *DescribeDBClustersResponseBodyItemsDBClusterTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	Mode                 *string                                           `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Port                 *string                                           `json:"Port,omitempty" xml:"Port,omitempty"`
	LockMode             *string                                           `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	StorageResource      *string                                           `json:"StorageResource,omitempty" xml:"StorageResource,omitempty"`
	ExecutorCount        *string                                           `json:"ExecutorCount,omitempty" xml:"ExecutorCount,omitempty"`
	DBClusterId          *string                                           `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ConnectionString     *string                                           `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	RdsInstanceId        *string                                           `json:"RdsInstanceId,omitempty" xml:"RdsInstanceId,omitempty"`
	DBClusterType        *string                                           `json:"DBClusterType,omitempty" xml:"DBClusterType,omitempty"`
	CommodityCode        *string                                           `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ExpireTime           *string                                           `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	DBNodeStorage        *int64                                            `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	DBNodeClass          *string                                           `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	LockReason           *string                                           `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	VPCId                *string                                           `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	RegionId             *string                                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ComputeResource      *string                                           `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	ElasticIOResource    *int32                                            `json:"ElasticIOResource,omitempty" xml:"ElasticIOResource,omitempty"`
	VSwitchId            *string                                           `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	DBVersion            *string                                           `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	VPCCloudInstanceId   *string                                           `json:"VPCCloudInstanceId,omitempty" xml:"VPCCloudInstanceId,omitempty"`
	DBClusterStatus      *string                                           `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	ResourceGroupId      *string                                           `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	DBClusterNetworkType *string                                           `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	DBClusterDescription *string                                           `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	ZoneId               *string                                           `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Category             *string                                           `json:"Category,omitempty" xml:"Category,omitempty"`
	Engine               *string                                           `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s DescribeDBClustersResponseBodyItemsDBCluster) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBodyItemsDBCluster) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDtsJobId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBNodeCount(v int64) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBNodeCount = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetExpired(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Expired = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetCreateTime(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetPayType(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.PayType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDiskType(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DiskType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetTags(v *DescribeDBClustersResponseBodyItemsDBClusterTags) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Tags = v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetMode(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Mode = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetPort(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Port = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetLockMode(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.LockMode = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetStorageResource(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.StorageResource = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetExecutorCount(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ExecutorCount = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBClusterId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetConnectionString(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetRdsInstanceId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.RdsInstanceId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBClusterType(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBClusterType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetCommodityCode(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetExpireTime(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBNodeStorage(v int64) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBNodeStorage = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBNodeClass(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetLockReason(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.LockReason = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetVPCId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.VPCId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetRegionId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetComputeResource(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ComputeResource = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetElasticIOResource(v int32) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ElasticIOResource = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetVSwitchId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBVersion(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetVPCCloudInstanceId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.VPCCloudInstanceId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBClusterStatus(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetResourceGroupId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBClusterNetworkType(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBClusterNetworkType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBClusterDescription(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetZoneId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetCategory(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Category = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetEngine(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Engine = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	PoolName             *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
}

func (s DescribeDBResourcePoolRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBResourcePoolRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBResourcePoolRequest) SetOwnerId(v int64) *DescribeDBResourcePoolRequest {
	s.OwnerId = &v
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

func (s *DescribeDBResourcePoolRequest) SetOwnerAccount(v string) *DescribeDBResourcePoolRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBResourcePoolRequest) SetDBClusterId(v string) *DescribeDBResourcePoolRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBResourcePoolRequest) SetPoolName(v string) *DescribeDBResourcePoolRequest {
	s.PoolName = &v
	return s
}

type DescribeDBResourcePoolResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PoolsInfo []*DescribeDBResourcePoolResponseBodyPoolsInfo `json:"PoolsInfo,omitempty" xml:"PoolsInfo,omitempty" type:"Repeated"`
}

func (s DescribeDBResourcePoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBResourcePoolResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBResourcePoolResponseBody) SetRequestId(v string) *DescribeDBResourcePoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBResourcePoolResponseBody) SetPoolsInfo(v []*DescribeDBResourcePoolResponseBodyPoolsInfo) *DescribeDBResourcePoolResponseBody {
	s.PoolsInfo = v
	return s
}

type DescribeDBResourcePoolResponseBodyPoolsInfo struct {
	QueryType  *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	PoolName   *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	PoolUsers  *string `json:"PoolUsers,omitempty" xml:"PoolUsers,omitempty"`
	NodeNum    *int32  `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
}

func (s DescribeDBResourcePoolResponseBodyPoolsInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBResourcePoolResponseBodyPoolsInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) SetQueryType(v string) *DescribeDBResourcePoolResponseBodyPoolsInfo {
	s.QueryType = &v
	return s
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) SetUpdateTime(v string) *DescribeDBResourcePoolResponseBodyPoolsInfo {
	s.UpdateTime = &v
	return s
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) SetPoolName(v string) *DescribeDBResourcePoolResponseBodyPoolsInfo {
	s.PoolName = &v
	return s
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) SetCreateTime(v string) *DescribeDBResourcePoolResponseBodyPoolsInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) SetPoolUsers(v string) *DescribeDBResourcePoolResponseBodyPoolsInfo {
	s.PoolUsers = &v
	return s
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) SetNodeNum(v int32) *DescribeDBResourcePoolResponseBodyPoolsInfo {
	s.NodeNum = &v
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

type DescribeElasticDailyPlanRequest struct {
	OwnerId                    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount       *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId            *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount               *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId                *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ElasticPlanName            *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	ResourcePoolName           *string `json:"ResourcePoolName,omitempty" xml:"ResourcePoolName,omitempty"`
	ElasticDailyPlanDay        *string `json:"ElasticDailyPlanDay,omitempty" xml:"ElasticDailyPlanDay,omitempty"`
	ElasticDailyPlanStatusList *string `json:"ElasticDailyPlanStatusList,omitempty" xml:"ElasticDailyPlanStatusList,omitempty"`
}

func (s DescribeElasticDailyPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticDailyPlanRequest) GoString() string {
	return s.String()
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

func (s *DescribeElasticDailyPlanRequest) SetOwnerAccount(v string) *DescribeElasticDailyPlanRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeElasticDailyPlanRequest) SetDBClusterId(v string) *DescribeElasticDailyPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeElasticDailyPlanRequest) SetElasticPlanName(v string) *DescribeElasticDailyPlanRequest {
	s.ElasticPlanName = &v
	return s
}

func (s *DescribeElasticDailyPlanRequest) SetResourcePoolName(v string) *DescribeElasticDailyPlanRequest {
	s.ResourcePoolName = &v
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

type DescribeElasticDailyPlanResponseBody struct {
	RequestId            *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ElasticDailyPlanList []*DescribeElasticDailyPlanResponseBodyElasticDailyPlanList `json:"ElasticDailyPlanList,omitempty" xml:"ElasticDailyPlanList,omitempty" type:"Repeated"`
}

func (s DescribeElasticDailyPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticDailyPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticDailyPlanResponseBody) SetRequestId(v string) *DescribeElasticDailyPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBody) SetElasticDailyPlanList(v []*DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) *DescribeElasticDailyPlanResponseBody {
	s.ElasticDailyPlanList = v
	return s
}

type DescribeElasticDailyPlanResponseBodyElasticDailyPlanList struct {
	Status           *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Day              *string `json:"Day,omitempty" xml:"Day,omitempty"`
	ResourcePoolName *string `json:"ResourcePoolName,omitempty" xml:"ResourcePoolName,omitempty"`
	StartTs          *string `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
	PlanEndTs        *string `json:"PlanEndTs,omitempty" xml:"PlanEndTs,omitempty"`
	PlanStartTs      *string `json:"PlanStartTs,omitempty" xml:"PlanStartTs,omitempty"`
	ElasticNodeNum   *int32  `json:"ElasticNodeNum,omitempty" xml:"ElasticNodeNum,omitempty"`
	EndTs            *string `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	PlanName         *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
}

func (s DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) GoString() string {
	return s.String()
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetStatus(v int32) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.Status = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetDay(v string) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.Day = &v
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

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetPlanEndTs(v string) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.PlanEndTs = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetPlanStartTs(v string) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.PlanStartTs = &v
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

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetPlanName(v string) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.PlanName = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ElasticPlanName      *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	ResourcePoolName     *string `json:"ResourcePoolName,omitempty" xml:"ResourcePoolName,omitempty"`
	ElasticPlanEnable    *bool   `json:"ElasticPlanEnable,omitempty" xml:"ElasticPlanEnable,omitempty"`
}

func (s DescribeElasticPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticPlanRequest) GoString() string {
	return s.String()
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

func (s *DescribeElasticPlanRequest) SetOwnerAccount(v string) *DescribeElasticPlanRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeElasticPlanRequest) SetDBClusterId(v string) *DescribeElasticPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeElasticPlanRequest) SetElasticPlanName(v string) *DescribeElasticPlanRequest {
	s.ElasticPlanName = &v
	return s
}

func (s *DescribeElasticPlanRequest) SetResourcePoolName(v string) *DescribeElasticPlanRequest {
	s.ResourcePoolName = &v
	return s
}

func (s *DescribeElasticPlanRequest) SetElasticPlanEnable(v bool) *DescribeElasticPlanRequest {
	s.ElasticPlanEnable = &v
	return s
}

type DescribeElasticPlanResponseBody struct {
	RequestId       *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ElasticPlanList []*DescribeElasticPlanResponseBodyElasticPlanList `json:"ElasticPlanList,omitempty" xml:"ElasticPlanList,omitempty" type:"Repeated"`
}

func (s DescribeElasticPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanResponseBody) SetRequestId(v string) *DescribeElasticPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeElasticPlanResponseBody) SetElasticPlanList(v []*DescribeElasticPlanResponseBodyElasticPlanList) *DescribeElasticPlanResponseBody {
	s.ElasticPlanList = v
	return s
}

type DescribeElasticPlanResponseBodyElasticPlanList struct {
	EndTime          *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	WeeklyRepeat     *string `json:"WeeklyRepeat,omitempty" xml:"WeeklyRepeat,omitempty"`
	StartTime        *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ResourcePoolName *string `json:"ResourcePoolName,omitempty" xml:"ResourcePoolName,omitempty"`
	StartDay         *string `json:"StartDay,omitempty" xml:"StartDay,omitempty"`
	ElasticNodeNum   *int32  `json:"ElasticNodeNum,omitempty" xml:"ElasticNodeNum,omitempty"`
	Enable           *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
	EndDay           *string `json:"EndDay,omitempty" xml:"EndDay,omitempty"`
	PlanName         *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
}

func (s DescribeElasticPlanResponseBodyElasticPlanList) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticPlanResponseBodyElasticPlanList) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetEndTime(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.EndTime = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetWeeklyRepeat(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.WeeklyRepeat = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetStartTime(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.StartTime = &v
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

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetPlanName(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.PlanName = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	TableType            *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
	Order                *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeInclinedTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInclinedTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInclinedTablesRequest) SetOwnerId(v int64) *DescribeInclinedTablesRequest {
	s.OwnerId = &v
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

func (s *DescribeInclinedTablesRequest) SetOwnerAccount(v string) *DescribeInclinedTablesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetDBClusterId(v string) *DescribeInclinedTablesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetTableType(v string) *DescribeInclinedTablesRequest {
	s.TableType = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetOrder(v string) *DescribeInclinedTablesRequest {
	s.Order = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetPageSize(v int32) *DescribeInclinedTablesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetPageNumber(v int32) *DescribeInclinedTablesRequest {
	s.PageNumber = &v
	return s
}

type DescribeInclinedTablesResponseBody struct {
	TotalCount *string                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *string                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *string                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Items      *DescribeInclinedTablesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeInclinedTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInclinedTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInclinedTablesResponseBody) SetTotalCount(v string) *DescribeInclinedTablesResponseBody {
	s.TotalCount = &v
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

func (s *DescribeInclinedTablesResponseBody) SetPageNumber(v string) *DescribeInclinedTablesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInclinedTablesResponseBody) SetItems(v *DescribeInclinedTablesResponseBodyItems) *DescribeInclinedTablesResponseBody {
	s.Items = v
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
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Schema    *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	Size      *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	IsIncline *bool   `json:"IsIncline,omitempty" xml:"IsIncline,omitempty"`
}

func (s DescribeInclinedTablesResponseBodyItemsTable) String() string {
	return tea.Prettify(s)
}

func (s DescribeInclinedTablesResponseBodyItemsTable) GoString() string {
	return s.String()
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) SetType(v string) *DescribeInclinedTablesResponseBodyItemsTable {
	s.Type = &v
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

func (s *DescribeInclinedTablesResponseBodyItemsTable) SetName(v string) *DescribeInclinedTablesResponseBodyItemsTable {
	s.Name = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) SetIsIncline(v bool) *DescribeInclinedTablesResponseBodyItemsTable {
	s.IsIncline = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Order                *string `json:"Order,omitempty" xml:"Order,omitempty"`
	Range                *string `json:"Range,omitempty" xml:"Range,omitempty"`
	State                *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeLoadTasksRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadTasksRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoadTasksRecordsRequest) SetOwnerId(v int64) *DescribeLoadTasksRecordsRequest {
	s.OwnerId = &v
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

func (s *DescribeLoadTasksRecordsRequest) SetOwnerAccount(v string) *DescribeLoadTasksRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetDBClusterId(v string) *DescribeLoadTasksRecordsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetStartTime(v string) *DescribeLoadTasksRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetEndTime(v string) *DescribeLoadTasksRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetDBName(v string) *DescribeLoadTasksRecordsRequest {
	s.DBName = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetPageSize(v int32) *DescribeLoadTasksRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetPageNumber(v int32) *DescribeLoadTasksRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetOrder(v string) *DescribeLoadTasksRecordsRequest {
	s.Order = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetRange(v string) *DescribeLoadTasksRecordsRequest {
	s.Range = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetState(v string) *DescribeLoadTasksRecordsRequest {
	s.State = &v
	return s
}

type DescribeLoadTasksRecordsResponseBody struct {
	TotalCount       *string                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize         *string                                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId        *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber       *string                                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	DBClusterId      *string                                                 `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	LoadTasksRecords []*DescribeLoadTasksRecordsResponseBodyLoadTasksRecords `json:"LoadTasksRecords,omitempty" xml:"LoadTasksRecords,omitempty" type:"Repeated"`
}

func (s DescribeLoadTasksRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadTasksRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadTasksRecordsResponseBody) SetTotalCount(v string) *DescribeLoadTasksRecordsResponseBody {
	s.TotalCount = &v
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

func (s *DescribeLoadTasksRecordsResponseBody) SetPageNumber(v string) *DescribeLoadTasksRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBody) SetDBClusterId(v string) *DescribeLoadTasksRecordsResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBody) SetLoadTasksRecords(v []*DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) *DescribeLoadTasksRecordsResponseBody {
	s.LoadTasksRecords = v
	return s
}

type DescribeLoadTasksRecordsResponseBodyLoadTasksRecords struct {
	Sql         *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
	State       *string `json:"State,omitempty" xml:"State,omitempty"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DBName      *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	ProcessID   *string `json:"ProcessID,omitempty" xml:"ProcessID,omitempty"`
	UpdateTime  *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	JobName     *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	ProcessRows *int64  `json:"ProcessRows,omitempty" xml:"ProcessRows,omitempty"`
}

func (s DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) GoString() string {
	return s.String()
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetSql(v string) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.Sql = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetState(v string) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.State = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetCreateTime(v string) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.CreateTime = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetDBName(v string) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.DBName = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetProcessID(v string) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.ProcessID = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetUpdateTime(v string) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.UpdateTime = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetJobName(v string) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.JobName = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetProcessRows(v int64) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.ProcessRows = &v
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

type DescribeOperatorPermissionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeOperatorPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOperatorPermissionRequest) GoString() string {
	return s.String()
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

func (s *DescribeOperatorPermissionRequest) SetOwnerAccount(v string) *DescribeOperatorPermissionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeOperatorPermissionRequest) SetDBClusterId(v string) *DescribeOperatorPermissionRequest {
	s.DBClusterId = &v
	return s
}

type DescribeOperatorPermissionResponseBody struct {
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Privileges  *string `json:"Privileges,omitempty" xml:"Privileges,omitempty"`
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
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

func (s *DescribeOperatorPermissionResponseBody) SetRequestId(v string) *DescribeOperatorPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOperatorPermissionResponseBody) SetPrivileges(v string) *DescribeOperatorPermissionResponseBody {
	s.Privileges = &v
	return s
}

func (s *DescribeOperatorPermissionResponseBody) SetExpiredTime(v string) *DescribeOperatorPermissionResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeOperatorPermissionResponseBody) SetDBClusterId(v string) *DescribeOperatorPermissionResponseBody {
	s.DBClusterId = &v
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

type DescribeProcessListRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ShowFull             *bool   `json:"ShowFull,omitempty" xml:"ShowFull,omitempty"`
	RunningTime          *int32  `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	User                 *string `json:"User,omitempty" xml:"User,omitempty"`
	Keyword              *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Order                *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeProcessListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessListRequest) GoString() string {
	return s.String()
}

func (s *DescribeProcessListRequest) SetOwnerId(v int64) *DescribeProcessListRequest {
	s.OwnerId = &v
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

func (s *DescribeProcessListRequest) SetOwnerAccount(v string) *DescribeProcessListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeProcessListRequest) SetDBClusterId(v string) *DescribeProcessListRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeProcessListRequest) SetShowFull(v bool) *DescribeProcessListRequest {
	s.ShowFull = &v
	return s
}

func (s *DescribeProcessListRequest) SetRunningTime(v int32) *DescribeProcessListRequest {
	s.RunningTime = &v
	return s
}

func (s *DescribeProcessListRequest) SetUser(v string) *DescribeProcessListRequest {
	s.User = &v
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

func (s *DescribeProcessListRequest) SetPageSize(v int32) *DescribeProcessListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeProcessListRequest) SetPageNumber(v int32) *DescribeProcessListRequest {
	s.PageNumber = &v
	return s
}

type DescribeProcessListResponseBody struct {
	TotalCount *string                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *string                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *string                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Items      *DescribeProcessListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeProcessListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBody) SetTotalCount(v string) *DescribeProcessListResponseBody {
	s.TotalCount = &v
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

func (s *DescribeProcessListResponseBody) SetPageNumber(v string) *DescribeProcessListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeProcessListResponseBody) SetItems(v *DescribeProcessListResponseBodyItems) *DescribeProcessListResponseBody {
	s.Items = v
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
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Time      *int32  `json:"Time,omitempty" xml:"Time,omitempty"`
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	DB        *string `json:"DB,omitempty" xml:"DB,omitempty"`
	Command   *string `json:"Command,omitempty" xml:"Command,omitempty"`
	User      *string `json:"User,omitempty" xml:"User,omitempty"`
	Id        *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	Info      *string `json:"Info,omitempty" xml:"Info,omitempty"`
}

func (s DescribeProcessListResponseBodyItemsProcess) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessListResponseBodyItemsProcess) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetStartTime(v string) *DescribeProcessListResponseBodyItemsProcess {
	s.StartTime = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetTime(v int32) *DescribeProcessListResponseBodyItemsProcess {
	s.Time = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetProcessId(v string) *DescribeProcessListResponseBodyItemsProcess {
	s.ProcessId = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetHost(v string) *DescribeProcessListResponseBodyItemsProcess {
	s.Host = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetDB(v string) *DescribeProcessListResponseBodyItemsProcess {
	s.DB = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetCommand(v string) *DescribeProcessListResponseBodyItemsProcess {
	s.Command = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetUser(v string) *DescribeProcessListResponseBodyItemsProcess {
	s.User = &v
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
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
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

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
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
	Zones          *DescribeRegionsResponseBodyRegionsRegionZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
	LocalName      *string                                        `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string                                        `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetZones(v *DescribeRegionsResponseBodyRegionsRegionZones) *DescribeRegionsResponseBodyRegionsRegion {
	s.Zones = v
	return s
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
	ZoneId     *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	VpcEnabled *bool   `json:"VpcEnabled,omitempty" xml:"VpcEnabled,omitempty"`
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

func (s *DescribeRegionsResponseBodyRegionsRegionZonesZone) SetZoneId(v string) *DescribeRegionsResponseBodyRegionsRegionZonesZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegionZonesZone) SetVpcEnabled(v bool) *DescribeRegionsResponseBodyRegionsRegionZonesZone {
	s.VpcEnabled = &v
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

type DescribeSchemasRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeSchemasRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSchemasRequest) GoString() string {
	return s.String()
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

func (s *DescribeSchemasRequest) SetOwnerAccount(v string) *DescribeSchemasRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSchemasRequest) SetDBClusterId(v string) *DescribeSchemasRequest {
	s.DBClusterId = &v
	return s
}

type DescribeSchemasResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Items     *DescribeSchemasResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeSchemasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSchemasResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSchemasResponseBody) SetRequestId(v string) *DescribeSchemasResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSchemasResponseBody) SetItems(v *DescribeSchemasResponseBodyItems) *DescribeSchemasResponseBody {
	s.Items = v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ProcessID            *string `json:"ProcessID,omitempty" xml:"ProcessID,omitempty"`
	Order                *string `json:"Order,omitempty" xml:"Order,omitempty"`
	Range                *string `json:"Range,omitempty" xml:"Range,omitempty"`
	State                *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeSlowLogRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsRequest) GoString() string {
	return s.String()
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

func (s *DescribeSlowLogRecordsRequest) SetDBClusterId(v string) *DescribeSlowLogRecordsRequest {
	s.DBClusterId = &v
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

func (s *DescribeSlowLogRecordsRequest) SetProcessID(v string) *DescribeSlowLogRecordsRequest {
	s.ProcessID = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetOrder(v string) *DescribeSlowLogRecordsRequest {
	s.Order = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetRange(v string) *DescribeSlowLogRecordsRequest {
	s.Range = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetState(v string) *DescribeSlowLogRecordsRequest {
	s.State = &v
	return s
}

type DescribeSlowLogRecordsResponseBody struct {
	TotalCount  *string                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize    *string                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber  *string                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	DBClusterId *string                                  `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Items       *DescribeSlowLogRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeSlowLogRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBody) SetTotalCount(v string) *DescribeSlowLogRecordsResponseBody {
	s.TotalCount = &v
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

func (s *DescribeSlowLogRecordsResponseBody) SetPageNumber(v string) *DescribeSlowLogRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetDBClusterId(v string) *DescribeSlowLogRecordsResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetItems(v *DescribeSlowLogRecordsResponseBodyItems) *DescribeSlowLogRecordsResponseBody {
	s.Items = v
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
	HostAddress        *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	ScanTime           *int64  `json:"ScanTime,omitempty" xml:"ScanTime,omitempty"`
	SQLText            *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	OutputSize         *string `json:"OutputSize,omitempty" xml:"OutputSize,omitempty"`
	PeakMemoryUsage    *string `json:"PeakMemoryUsage,omitempty" xml:"PeakMemoryUsage,omitempty"`
	State              *string `json:"State,omitempty" xml:"State,omitempty"`
	WallTime           *int64  `json:"WallTime,omitempty" xml:"WallTime,omitempty"`
	ScanSize           *string `json:"ScanSize,omitempty" xml:"ScanSize,omitempty"`
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" xml:"ExecutionStartTime,omitempty"`
	QueryTime          *int64  `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
	ReturnRowCounts    *int64  `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	ScanRows           *int64  `json:"ScanRows,omitempty" xml:"ScanRows,omitempty"`
	ParseRowCounts     *int64  `json:"ParseRowCounts,omitempty" xml:"ParseRowCounts,omitempty"`
	DBName             *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	PlanningTime       *int64  `json:"PlanningTime,omitempty" xml:"PlanningTime,omitempty"`
	QueueTime          *int64  `json:"QueueTime,omitempty" xml:"QueueTime,omitempty"`
	UserName           *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	ProcessID          *string `json:"ProcessID,omitempty" xml:"ProcessID,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetHostAddress(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.HostAddress = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetScanTime(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.ScanTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetSQLText(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.SQLText = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetOutputSize(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.OutputSize = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetPeakMemoryUsage(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.PeakMemoryUsage = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetState(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.State = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetWallTime(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.WallTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetScanSize(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.ScanSize = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetExecutionStartTime(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.ExecutionStartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetQueryTime(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.QueryTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetReturnRowCounts(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetScanRows(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.ScanRows = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetParseRowCounts(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.ParseRowCounts = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetDBName(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetPlanningTime(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.PlanningTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetQueueTime(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.QueueTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetUserName(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.UserName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetProcessID(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.ProcessID = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
}

func (s DescribeSlowLogTrendRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogTrendRequest) GoString() string {
	return s.String()
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

func (s *DescribeSlowLogTrendRequest) SetOwnerAccount(v string) *DescribeSlowLogTrendRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetDBClusterId(v string) *DescribeSlowLogTrendRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetStartTime(v string) *DescribeSlowLogTrendRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetEndTime(v string) *DescribeSlowLogTrendRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetDBName(v string) *DescribeSlowLogTrendRequest {
	s.DBName = &v
	return s
}

type DescribeSlowLogTrendResponseBody struct {
	EndTime     *string                                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime   *string                                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	DBClusterId *string                                `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Items       *DescribeSlowLogTrendResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeSlowLogTrendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBody) SetEndTime(v string) *DescribeSlowLogTrendResponseBody {
	s.EndTime = &v
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

func (s *DescribeSlowLogTrendResponseBody) SetDBClusterId(v string) *DescribeSlowLogTrendResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBody) SetItems(v *DescribeSlowLogTrendResponseBodyItems) *DescribeSlowLogTrendResponseBody {
	s.Items = v
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
	Unit   *string                                                      `json:"Unit,omitempty" xml:"Unit,omitempty"`
	Series *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Struct"`
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

func (s *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItem) SetUnit(v string) *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItem {
	s.Unit = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItem) SetSeries(v *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeries) *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItem {
	s.Series = v
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
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeriesSeriesItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeriesSeriesItem) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeriesSeriesItem) SetValues(v string) *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeriesSeriesItem {
	s.Values = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeriesSeriesItem) SetName(v string) *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeriesSeriesItem {
	s.Name = &v
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

type DescribeSQLPlanRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ProcessId            *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
}

func (s DescribeSQLPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlanRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanRequest) SetOwnerId(v int64) *DescribeSQLPlanRequest {
	s.OwnerId = &v
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

func (s *DescribeSQLPlanRequest) SetOwnerAccount(v string) *DescribeSQLPlanRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSQLPlanRequest) SetDBClusterId(v string) *DescribeSQLPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSQLPlanRequest) SetProcessId(v string) *DescribeSQLPlanRequest {
	s.ProcessId = &v
	return s
}

type DescribeSQLPlanResponseBody struct {
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StageList  []*DescribeSQLPlanResponseBodyStageList `json:"StageList,omitempty" xml:"StageList,omitempty" type:"Repeated"`
	OriginInfo *string                                 `json:"OriginInfo,omitempty" xml:"OriginInfo,omitempty"`
	Detail     *DescribeSQLPlanResponseBodyDetail      `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Struct"`
}

func (s DescribeSQLPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanResponseBody) SetRequestId(v string) *DescribeSQLPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLPlanResponseBody) SetStageList(v []*DescribeSQLPlanResponseBodyStageList) *DescribeSQLPlanResponseBody {
	s.StageList = v
	return s
}

func (s *DescribeSQLPlanResponseBody) SetOriginInfo(v string) *DescribeSQLPlanResponseBody {
	s.OriginInfo = &v
	return s
}

func (s *DescribeSQLPlanResponseBody) SetDetail(v *DescribeSQLPlanResponseBodyDetail) *DescribeSQLPlanResponseBody {
	s.Detail = v
	return s
}

type DescribeSQLPlanResponseBodyStageList struct {
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
	CPUTimeAvg   *int64  `json:"CPUTimeAvg,omitempty" xml:"CPUTimeAvg,omitempty"`
	CPUTimeMax   *int64  `json:"CPUTimeMax,omitempty" xml:"CPUTimeMax,omitempty"`
	OperatorCost *int64  `json:"OperatorCost,omitempty" xml:"OperatorCost,omitempty"`
	ScanTimeMax  *int64  `json:"ScanTimeMax,omitempty" xml:"ScanTimeMax,omitempty"`
	InputSizeMax *int64  `json:"InputSizeMax,omitempty" xml:"InputSizeMax,omitempty"`
	StageId      *int32  `json:"StageId,omitempty" xml:"StageId,omitempty"`
	ScanSizeMax  *int64  `json:"ScanSizeMax,omitempty" xml:"ScanSizeMax,omitempty"`
	CPUTimeMin   *int64  `json:"CPUTimeMin,omitempty" xml:"CPUTimeMin,omitempty"`
	ScanTimeMin  *int64  `json:"ScanTimeMin,omitempty" xml:"ScanTimeMin,omitempty"`
	ScanSizeMin  *int64  `json:"ScanSizeMin,omitempty" xml:"ScanSizeMin,omitempty"`
	InputSizeMin *int64  `json:"InputSizeMin,omitempty" xml:"InputSizeMin,omitempty"`
	PeakMemory   *int64  `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	ScanTimeAvg  *int64  `json:"ScanTimeAvg,omitempty" xml:"ScanTimeAvg,omitempty"`
	ScanSizeAvg  *int64  `json:"ScanSizeAvg,omitempty" xml:"ScanSizeAvg,omitempty"`
	InputSizeAvg *int64  `json:"InputSizeAvg,omitempty" xml:"InputSizeAvg,omitempty"`
}

func (s DescribeSQLPlanResponseBodyStageList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlanResponseBodyStageList) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanResponseBodyStageList) SetState(v string) *DescribeSQLPlanResponseBodyStageList {
	s.State = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetCPUTimeAvg(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.CPUTimeAvg = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetCPUTimeMax(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.CPUTimeMax = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetOperatorCost(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.OperatorCost = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetScanTimeMax(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.ScanTimeMax = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetInputSizeMax(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.InputSizeMax = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetStageId(v int32) *DescribeSQLPlanResponseBodyStageList {
	s.StageId = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetScanSizeMax(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.ScanSizeMax = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetCPUTimeMin(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.CPUTimeMin = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetScanTimeMin(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.ScanTimeMin = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetScanSizeMin(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.ScanSizeMin = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetInputSizeMin(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.InputSizeMin = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetPeakMemory(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.PeakMemory = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetScanTimeAvg(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.ScanTimeAvg = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetScanSizeAvg(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.ScanSizeAvg = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetInputSizeAvg(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.InputSizeAvg = &v
	return s
}

type DescribeSQLPlanResponseBodyDetail struct {
	SQL          *string `json:"SQL,omitempty" xml:"SQL,omitempty"`
	OutputSize   *int64  `json:"OutputSize,omitempty" xml:"OutputSize,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
	OutputRows   *int64  `json:"OutputRows,omitempty" xml:"OutputRows,omitempty"`
	User         *string `json:"User,omitempty" xml:"User,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TotalStage   *int64  `json:"TotalStage,omitempty" xml:"TotalStage,omitempty"`
	QueuedTime   *int64  `json:"QueuedTime,omitempty" xml:"QueuedTime,omitempty"`
	TotalTime    *int64  `json:"TotalTime,omitempty" xml:"TotalTime,omitempty"`
	TotalTask    *int64  `json:"TotalTask,omitempty" xml:"TotalTask,omitempty"`
	Database     *string `json:"Database,omitempty" xml:"Database,omitempty"`
	PeakMemory   *int64  `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	ClientIP     *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
	PlanningTime *int64  `json:"PlanningTime,omitempty" xml:"PlanningTime,omitempty"`
	CPUTime      *int64  `json:"CPUTime,omitempty" xml:"CPUTime,omitempty"`
}

func (s DescribeSQLPlanResponseBodyDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlanResponseBodyDetail) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanResponseBodyDetail) SetSQL(v string) *DescribeSQLPlanResponseBodyDetail {
	s.SQL = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetOutputSize(v int64) *DescribeSQLPlanResponseBodyDetail {
	s.OutputSize = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetState(v string) *DescribeSQLPlanResponseBodyDetail {
	s.State = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetOutputRows(v int64) *DescribeSQLPlanResponseBodyDetail {
	s.OutputRows = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetUser(v string) *DescribeSQLPlanResponseBodyDetail {
	s.User = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetStartTime(v string) *DescribeSQLPlanResponseBodyDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetTotalStage(v int64) *DescribeSQLPlanResponseBodyDetail {
	s.TotalStage = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetQueuedTime(v int64) *DescribeSQLPlanResponseBodyDetail {
	s.QueuedTime = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetTotalTime(v int64) *DescribeSQLPlanResponseBodyDetail {
	s.TotalTime = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetTotalTask(v int64) *DescribeSQLPlanResponseBodyDetail {
	s.TotalTask = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetDatabase(v string) *DescribeSQLPlanResponseBodyDetail {
	s.Database = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetPeakMemory(v int64) *DescribeSQLPlanResponseBodyDetail {
	s.PeakMemory = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetClientIP(v string) *DescribeSQLPlanResponseBodyDetail {
	s.ClientIP = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetPlanningTime(v int64) *DescribeSQLPlanResponseBodyDetail {
	s.PlanningTime = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetCPUTime(v int64) *DescribeSQLPlanResponseBodyDetail {
	s.CPUTime = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ProcessId            *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	StageId              *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
}

func (s DescribeSQLPlanTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlanTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanTaskRequest) SetOwnerId(v int64) *DescribeSQLPlanTaskRequest {
	s.OwnerId = &v
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

func (s *DescribeSQLPlanTaskRequest) SetOwnerAccount(v string) *DescribeSQLPlanTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSQLPlanTaskRequest) SetDBClusterId(v string) *DescribeSQLPlanTaskRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSQLPlanTaskRequest) SetProcessId(v string) *DescribeSQLPlanTaskRequest {
	s.ProcessId = &v
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
	ScanCost     *int64  `json:"ScanCost,omitempty" xml:"ScanCost,omitempty"`
	OutputSize   *int64  `json:"OutputSize,omitempty" xml:"OutputSize,omitempty"`
	InputSize    *int64  `json:"InputSize,omitempty" xml:"InputSize,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
	OperatorCost *int64  `json:"OperatorCost,omitempty" xml:"OperatorCost,omitempty"`
	OutputRows   *int64  `json:"OutputRows,omitempty" xml:"OutputRows,omitempty"`
	ScanSize     *int64  `json:"ScanSize,omitempty" xml:"ScanSize,omitempty"`
	ElapsedTime  *int64  `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	ScanRows     *int64  `json:"ScanRows,omitempty" xml:"ScanRows,omitempty"`
	PeakMemory   *int64  `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	TaskId       *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	InputRows    *int64  `json:"InputRows,omitempty" xml:"InputRows,omitempty"`
}

func (s DescribeSQLPlanTaskResponseBodyTaskList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlanTaskResponseBodyTaskList) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetScanCost(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.ScanCost = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetOutputSize(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.OutputSize = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetInputSize(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.InputSize = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetState(v string) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.State = &v
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

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetScanSize(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.ScanSize = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetElapsedTime(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.ElapsedTime = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetScanRows(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.ScanRows = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetPeakMemory(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.PeakMemory = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetTaskId(v int32) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.TaskId = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetInputRows(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.InputRows = &v
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

type DescribeTableDetailRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	SchemaName           *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName            *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeTableDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableDetailRequest) GoString() string {
	return s.String()
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

func (s *DescribeTableDetailRequest) SetOwnerAccount(v string) *DescribeTableDetailRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTableDetailRequest) SetDBClusterId(v string) *DescribeTableDetailRequest {
	s.DBClusterId = &v
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
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Items     *DescribeTableDetailResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
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

func (s *DescribeTableDetailResponseBody) SetRequestId(v string) *DescribeTableDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTableDetailResponseBody) SetItems(v *DescribeTableDetailResponseBodyItems) *DescribeTableDetailResponseBody {
	s.Items = v
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
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	Id   *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeTableDetailResponseBodyItemsShard) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableDetailResponseBodyItemsShard) GoString() string {
	return s.String()
}

func (s *DescribeTableDetailResponseBodyItemsShard) SetSize(v int64) *DescribeTableDetailResponseBodyItemsShard {
	s.Size = &v
	return s
}

func (s *DescribeTableDetailResponseBodyItemsShard) SetId(v int32) *DescribeTableDetailResponseBodyItemsShard {
	s.Id = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeTablePartitionDiagnoseRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablePartitionDiagnoseRequest) GoString() string {
	return s.String()
}

func (s *DescribeTablePartitionDiagnoseRequest) SetOwnerId(v int64) *DescribeTablePartitionDiagnoseRequest {
	s.OwnerId = &v
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

func (s *DescribeTablePartitionDiagnoseRequest) SetOwnerAccount(v string) *DescribeTablePartitionDiagnoseRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseRequest) SetRegionId(v string) *DescribeTablePartitionDiagnoseRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseRequest) SetDBClusterId(v string) *DescribeTablePartitionDiagnoseRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseRequest) SetPageSize(v int32) *DescribeTablePartitionDiagnoseRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseRequest) SetPageNumber(v int32) *DescribeTablePartitionDiagnoseRequest {
	s.PageNumber = &v
	return s
}

type DescribeTablePartitionDiagnoseResponseBody struct {
	TotalCount                    *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId                     *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize                      *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber                    *int32                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	DBClusterId                   *string                                            `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	SuggestMaxRecordsPerPartition *int64                                             `json:"SuggestMaxRecordsPerPartition,omitempty" xml:"SuggestMaxRecordsPerPartition,omitempty"`
	SuggestMinRecordsPerPartition *int64                                             `json:"SuggestMinRecordsPerPartition,omitempty" xml:"SuggestMinRecordsPerPartition,omitempty"`
	Items                         []*DescribeTablePartitionDiagnoseResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeTablePartitionDiagnoseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablePartitionDiagnoseResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetTotalCount(v int32) *DescribeTablePartitionDiagnoseResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetRequestId(v string) *DescribeTablePartitionDiagnoseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetPageSize(v int32) *DescribeTablePartitionDiagnoseResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetPageNumber(v int32) *DescribeTablePartitionDiagnoseResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetDBClusterId(v string) *DescribeTablePartitionDiagnoseResponseBody {
	s.DBClusterId = &v
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

func (s *DescribeTablePartitionDiagnoseResponseBody) SetItems(v []*DescribeTablePartitionDiagnoseResponseBodyItems) *DescribeTablePartitionDiagnoseResponseBody {
	s.Items = v
	return s
}

type DescribeTablePartitionDiagnoseResponseBodyItems struct {
	TableName       *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	PartitionDetail *string `json:"PartitionDetail,omitempty" xml:"PartitionDetail,omitempty"`
	SchemaName      *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	PartitionNumber *int32  `json:"PartitionNumber,omitempty" xml:"PartitionNumber,omitempty"`
}

func (s DescribeTablePartitionDiagnoseResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablePartitionDiagnoseResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeTablePartitionDiagnoseResponseBodyItems) SetTableName(v string) *DescribeTablePartitionDiagnoseResponseBodyItems {
	s.TableName = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBodyItems) SetPartitionDetail(v string) *DescribeTablePartitionDiagnoseResponseBodyItems {
	s.PartitionDetail = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBodyItems) SetSchemaName(v string) *DescribeTablePartitionDiagnoseResponseBodyItems {
	s.SchemaName = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBodyItems) SetPartitionNumber(v int32) *DescribeTablePartitionDiagnoseResponseBodyItems {
	s.PartitionNumber = &v
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

type DescribeTablesRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	SchemaName           *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
}

func (s DescribeTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesRequest) GoString() string {
	return s.String()
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

func (s *DescribeTablesRequest) SetOwnerAccount(v string) *DescribeTablesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTablesRequest) SetDBClusterId(v string) *DescribeTablesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTablesRequest) SetSchemaName(v string) *DescribeTablesRequest {
	s.SchemaName = &v
	return s
}

type DescribeTablesResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Items     *DescribeTablesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponseBody) SetRequestId(v string) *DescribeTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTablesResponseBody) SetItems(v *DescribeTablesResponseBodyItems) *DescribeTablesResponseBody {
	s.Items = v
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
	TableName   *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	SchemaName  *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
}

func (s DescribeTablesResponseBodyItemsTable) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesResponseBodyItemsTable) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponseBodyItemsTable) SetTableName(v string) *DescribeTablesResponseBodyItemsTable {
	s.TableName = &v
	return s
}

func (s *DescribeTablesResponseBodyItemsTable) SetDBClusterId(v string) *DescribeTablesResponseBodyItemsTable {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTablesResponseBodyItemsTable) SetSchemaName(v string) *DescribeTablesResponseBodyItemsTable {
	s.SchemaName = &v
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

type DescribeTableStatisticsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Order                *string `json:"Order,omitempty" xml:"Order,omitempty"`
}

func (s DescribeTableStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTableStatisticsRequest) SetOwnerId(v int64) *DescribeTableStatisticsRequest {
	s.OwnerId = &v
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

func (s *DescribeTableStatisticsRequest) SetOwnerAccount(v string) *DescribeTableStatisticsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTableStatisticsRequest) SetDBClusterId(v string) *DescribeTableStatisticsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTableStatisticsRequest) SetPageSize(v int32) *DescribeTableStatisticsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTableStatisticsRequest) SetPageNumber(v int32) *DescribeTableStatisticsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTableStatisticsRequest) SetOrder(v string) *DescribeTableStatisticsRequest {
	s.Order = &v
	return s
}

type DescribeTableStatisticsResponseBody struct {
	TotalCount  *string                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize    *string                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber  *string                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	DBClusterId *string                                   `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Items       *DescribeTableStatisticsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeTableStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTableStatisticsResponseBody) SetTotalCount(v string) *DescribeTableStatisticsResponseBody {
	s.TotalCount = &v
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

func (s *DescribeTableStatisticsResponseBody) SetPageNumber(v string) *DescribeTableStatisticsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTableStatisticsResponseBody) SetDBClusterId(v string) *DescribeTableStatisticsResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTableStatisticsResponseBody) SetItems(v *DescribeTableStatisticsResponseBodyItems) *DescribeTableStatisticsResponseBody {
	s.Items = v
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
	SchemaName          *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName           *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	RowCount            *int64  `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	DataSize            *int64  `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	IndexSize           *int64  `json:"IndexSize,omitempty" xml:"IndexSize,omitempty"`
	PrimaryKeyIndexSize *int64  `json:"PrimaryKeyIndexSize,omitempty" xml:"PrimaryKeyIndexSize,omitempty"`
	PartitionCount      *int64  `json:"PartitionCount,omitempty" xml:"PartitionCount,omitempty"`
	ColdDataSize        *int64  `json:"ColdDataSize,omitempty" xml:"ColdDataSize,omitempty"`
}

func (s DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) GoString() string {
	return s.String()
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetSchemaName(v string) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.SchemaName = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetTableName(v string) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.TableName = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetRowCount(v int64) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.RowCount = &v
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

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetPrimaryKeyIndexSize(v int64) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.PrimaryKeyIndexSize = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetPartitionCount(v int64) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.PartitionCount = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetColdDataSize(v int64) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.ColdDataSize = &v
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

type DescribeTaskInfoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	TaskId               *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeTaskInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskInfoRequest) GoString() string {
	return s.String()
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

func (s *DescribeTaskInfoRequest) SetOwnerAccount(v string) *DescribeTaskInfoRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTaskInfoRequest) SetDBClusterId(v string) *DescribeTaskInfoRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTaskInfoRequest) SetTaskId(v int32) *DescribeTaskInfoRequest {
	s.TaskId = &v
	return s
}

type DescribeTaskInfoResponseBody struct {
	TaskInfo  *DescribeTaskInfoResponseBodyTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTaskInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTaskInfoResponseBody) SetTaskInfo(v *DescribeTaskInfoResponseBodyTaskInfo) *DescribeTaskInfoResponseBody {
	s.TaskInfo = v
	return s
}

func (s *DescribeTaskInfoResponseBody) SetRequestId(v string) *DescribeTaskInfoResponseBody {
	s.RequestId = &v
	return s
}

type DescribeTaskInfoResponseBodyTaskInfo struct {
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	Progress   *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	BeginTime  *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	TaskId     *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeTaskInfoResponseBodyTaskInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskInfoResponseBodyTaskInfo) GoString() string {
	return s.String()
}

func (s *DescribeTaskInfoResponseBodyTaskInfo) SetStatus(v string) *DescribeTaskInfoResponseBodyTaskInfo {
	s.Status = &v
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

func (s *DescribeTaskInfoResponseBodyTaskInfo) SetBeginTime(v string) *DescribeTaskInfoResponseBodyTaskInfo {
	s.BeginTime = &v
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

type GrantOperatorPermissionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ExpiredTime          *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Privileges           *string `json:"Privileges,omitempty" xml:"Privileges,omitempty"`
}

func (s GrantOperatorPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantOperatorPermissionRequest) GoString() string {
	return s.String()
}

func (s *GrantOperatorPermissionRequest) SetOwnerId(v int64) *GrantOperatorPermissionRequest {
	s.OwnerId = &v
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

func (s *GrantOperatorPermissionRequest) SetOwnerAccount(v string) *GrantOperatorPermissionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GrantOperatorPermissionRequest) SetDBClusterId(v string) *GrantOperatorPermissionRequest {
	s.DBClusterId = &v
	return s
}

func (s *GrantOperatorPermissionRequest) SetExpiredTime(v string) *GrantOperatorPermissionRequest {
	s.ExpiredTime = &v
	return s
}

func (s *GrantOperatorPermissionRequest) SetPrivileges(v string) *GrantOperatorPermissionRequest {
	s.Privileges = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ProcessId            *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
}

func (s KillProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s KillProcessRequest) GoString() string {
	return s.String()
}

func (s *KillProcessRequest) SetOwnerId(v int64) *KillProcessRequest {
	s.OwnerId = &v
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

func (s *KillProcessRequest) SetOwnerAccount(v string) *KillProcessRequest {
	s.OwnerAccount = &v
	return s
}

func (s *KillProcessRequest) SetDBClusterId(v string) *KillProcessRequest {
	s.DBClusterId = &v
	return s
}

func (s *KillProcessRequest) SetProcessId(v string) *KillProcessRequest {
	s.ProcessId = &v
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

type ModifyAccountDescriptionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountDescription   *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
}

func (s ModifyAccountDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionRequest) GoString() string {
	return s.String()
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

func (s *ModifyAccountDescriptionRequest) SetDBClusterId(v string) *ModifyAccountDescriptionRequest {
	s.DBClusterId = &v
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

type ModifyAuditLogConfigRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AuditLogStatus       *string `json:"AuditLogStatus,omitempty" xml:"AuditLogStatus,omitempty"`
}

func (s ModifyAuditLogConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAuditLogConfigRequest) GoString() string {
	return s.String()
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

func (s *ModifyAuditLogConfigRequest) SetDBClusterId(v string) *ModifyAuditLogConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetRegionId(v string) *ModifyAuditLogConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetAuditLogStatus(v string) *ModifyAuditLogConfigRequest {
	s.AuditLogStatus = &v
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

type ModifyAutoRenewAttributeRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RenewalStatus        *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
	Duration             *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	PeriodUnit           *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
}

func (s ModifyAutoRenewAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyAutoRenewAttributeRequest) SetOwnerId(v int64) *ModifyAutoRenewAttributeRequest {
	s.OwnerId = &v
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

func (s *ModifyAutoRenewAttributeRequest) SetOwnerAccount(v string) *ModifyAutoRenewAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetDBClusterId(v string) *ModifyAutoRenewAttributeRequest {
	s.DBClusterId = &v
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

func (s *ModifyAutoRenewAttributeRequest) SetDuration(v string) *ModifyAutoRenewAttributeRequest {
	s.Duration = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetPeriodUnit(v string) *ModifyAutoRenewAttributeRequest {
	s.PeriodUnit = &v
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
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount     *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId          *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount             *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId              *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	PreferredBackupTime      *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	PreferredBackupPeriod    *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	BackupRetentionPeriod    *string `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	EnableBackupLog          *string `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	LogBackupRetentionPeriod *int32  `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
}

func (s ModifyBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPolicyRequest) GoString() string {
	return s.String()
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

func (s *ModifyBackupPolicyRequest) SetDBClusterId(v string) *ModifyBackupPolicyRequest {
	s.DBClusterId = &v
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

func (s *ModifyBackupPolicyRequest) SetBackupRetentionPeriod(v string) *ModifyBackupPolicyRequest {
	s.BackupRetentionPeriod = &v
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
	OwnerId                 *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount    *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId         *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount            *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId             *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ConnectionStringPrefix  *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
	Port                    *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s ModifyClusterConnectionStringRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterConnectionStringRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterConnectionStringRequest) SetOwnerId(v int64) *ModifyClusterConnectionStringRequest {
	s.OwnerId = &v
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

func (s *ModifyClusterConnectionStringRequest) SetOwnerAccount(v string) *ModifyClusterConnectionStringRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetDBClusterId(v string) *ModifyClusterConnectionStringRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetConnectionStringPrefix(v string) *ModifyClusterConnectionStringRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetCurrentConnectionString(v string) *ModifyClusterConnectionStringRequest {
	s.CurrentConnectionString = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetPort(v int32) *ModifyClusterConnectionStringRequest {
	s.Port = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBNodeGroupCount     *string `json:"DBNodeGroupCount,omitempty" xml:"DBNodeGroupCount,omitempty"`
	DBNodeStorage        *string `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	DBNodeClass          *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	ModifyType           *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	ExecutorCount        *string `json:"ExecutorCount,omitempty" xml:"ExecutorCount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StorageResource      *string `json:"StorageResource,omitempty" xml:"StorageResource,omitempty"`
	ComputeResource      *string `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	ElasticIOResource    *int32  `json:"ElasticIOResource,omitempty" xml:"ElasticIOResource,omitempty"`
	DBClusterCategory    *string `json:"DBClusterCategory,omitempty" xml:"DBClusterCategory,omitempty"`
	Mode                 *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s ModifyDBClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterRequest) SetOwnerId(v int64) *ModifyDBClusterRequest {
	s.OwnerId = &v
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

func (s *ModifyDBClusterRequest) SetOwnerAccount(v string) *ModifyDBClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBClusterId(v string) *ModifyDBClusterRequest {
	s.DBClusterId = &v
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

func (s *ModifyDBClusterRequest) SetDBNodeClass(v string) *ModifyDBClusterRequest {
	s.DBNodeClass = &v
	return s
}

func (s *ModifyDBClusterRequest) SetModifyType(v string) *ModifyDBClusterRequest {
	s.ModifyType = &v
	return s
}

func (s *ModifyDBClusterRequest) SetExecutorCount(v string) *ModifyDBClusterRequest {
	s.ExecutorCount = &v
	return s
}

func (s *ModifyDBClusterRequest) SetRegionId(v string) *ModifyDBClusterRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBClusterRequest) SetStorageResource(v string) *ModifyDBClusterRequest {
	s.StorageResource = &v
	return s
}

func (s *ModifyDBClusterRequest) SetComputeResource(v string) *ModifyDBClusterRequest {
	s.ComputeResource = &v
	return s
}

func (s *ModifyDBClusterRequest) SetElasticIOResource(v int32) *ModifyDBClusterRequest {
	s.ElasticIOResource = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBClusterCategory(v string) *ModifyDBClusterRequest {
	s.DBClusterCategory = &v
	return s
}

func (s *ModifyDBClusterRequest) SetMode(v string) *ModifyDBClusterRequest {
	s.Mode = &v
	return s
}

type ModifyDBClusterResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OrderId     *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ModifyDBClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterResponseBody) SetRequestId(v string) *ModifyDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterResponseBody) SetDBClusterId(v string) *ModifyDBClusterResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterResponseBody) SetOrderId(v string) *ModifyDBClusterResponseBody {
	s.OrderId = &v
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
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId               *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	SecurityIps               *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
	DBClusterIPArrayName      *string `json:"DBClusterIPArrayName,omitempty" xml:"DBClusterIPArrayName,omitempty"`
	DBClusterIPArrayAttribute *string `json:"DBClusterIPArrayAttribute,omitempty" xml:"DBClusterIPArrayAttribute,omitempty"`
	ModifyMode                *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
}

func (s ModifyDBClusterAccessWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterAccessWhiteListRequest) GoString() string {
	return s.String()
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

func (s *ModifyDBClusterAccessWhiteListRequest) SetOwnerAccount(v string) *ModifyDBClusterAccessWhiteListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetDBClusterId(v string) *ModifyDBClusterAccessWhiteListRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetSecurityIps(v string) *ModifyDBClusterAccessWhiteListRequest {
	s.SecurityIps = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetDBClusterIPArrayName(v string) *ModifyDBClusterAccessWhiteListRequest {
	s.DBClusterIPArrayName = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetDBClusterIPArrayAttribute(v string) *ModifyDBClusterAccessWhiteListRequest {
	s.DBClusterIPArrayAttribute = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetModifyMode(v string) *ModifyDBClusterAccessWhiteListRequest {
	s.ModifyMode = &v
	return s
}

type ModifyDBClusterAccessWhiteListResponseBody struct {
	TaskId      *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s ModifyDBClusterAccessWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterAccessWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAccessWhiteListResponseBody) SetTaskId(v int32) *ModifyDBClusterAccessWhiteListResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListResponseBody) SetRequestId(v string) *ModifyDBClusterAccessWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListResponseBody) SetDBClusterId(v string) *ModifyDBClusterAccessWhiteListResponseBody {
	s.DBClusterId = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
}

func (s ModifyDBClusterDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterDescriptionRequest) GoString() string {
	return s.String()
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

func (s *ModifyDBClusterDescriptionRequest) SetOwnerAccount(v string) *ModifyDBClusterDescriptionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetDBClusterId(v string) *ModifyDBClusterDescriptionRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetDBClusterDescription(v string) *ModifyDBClusterDescriptionRequest {
	s.DBClusterDescription = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	MaintainTime         *string `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
}

func (s ModifyDBClusterMaintainTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterMaintainTimeRequest) GoString() string {
	return s.String()
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

func (s *ModifyDBClusterMaintainTimeRequest) SetOwnerAccount(v string) *ModifyDBClusterMaintainTimeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetDBClusterId(v string) *ModifyDBClusterMaintainTimeRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetMaintainTime(v string) *ModifyDBClusterMaintainTimeRequest {
	s.MaintainTime = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	NewResourceGroupId   *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
}

func (s ModifyDBClusterResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterResourceGroupRequest) GoString() string {
	return s.String()
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

func (s *ModifyDBClusterResourceGroupRequest) SetOwnerAccount(v string) *ModifyDBClusterResourceGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterResourceGroupRequest) SetDBClusterId(v string) *ModifyDBClusterResourceGroupRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterResourceGroupRequest) SetNewResourceGroupId(v string) *ModifyDBClusterResourceGroupRequest {
	s.NewResourceGroupId = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	PoolName             *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	QueryType            *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	NodeNum              *int32  `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
}

func (s ModifyDBResourcePoolRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBResourcePoolRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBResourcePoolRequest) SetOwnerId(v int64) *ModifyDBResourcePoolRequest {
	s.OwnerId = &v
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

func (s *ModifyDBResourcePoolRequest) SetOwnerAccount(v string) *ModifyDBResourcePoolRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBResourcePoolRequest) SetDBClusterId(v string) *ModifyDBResourcePoolRequest {
	s.DBClusterId = &v
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

func (s *ModifyDBResourcePoolRequest) SetNodeNum(v int32) *ModifyDBResourcePoolRequest {
	s.NodeNum = &v
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
	OwnerId                 *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount    *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId         *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount            *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId             *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ElasticPlanName         *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	ResourcePoolName        *string `json:"ResourcePoolName,omitempty" xml:"ResourcePoolName,omitempty"`
	ElasticPlanNodeNum      *int32  `json:"ElasticPlanNodeNum,omitempty" xml:"ElasticPlanNodeNum,omitempty"`
	ElasticPlanTimeStart    *string `json:"ElasticPlanTimeStart,omitempty" xml:"ElasticPlanTimeStart,omitempty"`
	ElasticPlanTimeEnd      *string `json:"ElasticPlanTimeEnd,omitempty" xml:"ElasticPlanTimeEnd,omitempty"`
	ElasticPlanWeeklyRepeat *string `json:"ElasticPlanWeeklyRepeat,omitempty" xml:"ElasticPlanWeeklyRepeat,omitempty"`
	ElasticPlanStartDay     *string `json:"ElasticPlanStartDay,omitempty" xml:"ElasticPlanStartDay,omitempty"`
	ElasticPlanEndDay       *string `json:"ElasticPlanEndDay,omitempty" xml:"ElasticPlanEndDay,omitempty"`
	ElasticPlanEnable       *bool   `json:"ElasticPlanEnable,omitempty" xml:"ElasticPlanEnable,omitempty"`
}

func (s ModifyElasticPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyElasticPlanRequest) GoString() string {
	return s.String()
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

func (s *ModifyElasticPlanRequest) SetOwnerAccount(v string) *ModifyElasticPlanRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetDBClusterId(v string) *ModifyElasticPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanName(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanName = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetResourcePoolName(v string) *ModifyElasticPlanRequest {
	s.ResourcePoolName = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanNodeNum(v int32) *ModifyElasticPlanRequest {
	s.ElasticPlanNodeNum = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanTimeStart(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanTimeStart = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanTimeEnd(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanTimeEnd = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanWeeklyRepeat(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanWeeklyRepeat = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanStartDay(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanStartDay = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanEndDay(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanEndDay = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanEnable(v bool) *ModifyElasticPlanRequest {
	s.ElasticPlanEnable = &v
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
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount     *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId          *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount             *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId              *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	EnableBackupLog          *string `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	LogBackupRetentionPeriod *string `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
}

func (s ModifyLogBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyLogBackupPolicyRequest) GoString() string {
	return s.String()
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

func (s *ModifyLogBackupPolicyRequest) SetOwnerAccount(v string) *ModifyLogBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
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

type ReleaseClusterPublicConnectionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s ReleaseClusterPublicConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseClusterPublicConnectionRequest) GoString() string {
	return s.String()
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

func (s *ReleaseClusterPublicConnectionRequest) SetOwnerAccount(v string) *ReleaseClusterPublicConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReleaseClusterPublicConnectionRequest) SetDBClusterId(v string) *ReleaseClusterPublicConnectionRequest {
	s.DBClusterId = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPassword      *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	AccountType          *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s ResetAccountPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountPasswordRequest) GoString() string {
	return s.String()
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

func (s *ResetAccountPasswordRequest) SetDBClusterId(v string) *ResetAccountPasswordRequest {
	s.DBClusterId = &v
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

func (s *ResetAccountPasswordRequest) SetAccountType(v string) *ResetAccountPasswordRequest {
	s.AccountType = &v
	return s
}

type ResetAccountPasswordResponseBody struct {
	TaskId      *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s ResetAccountPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordResponseBody) SetTaskId(v int32) *ResetAccountPasswordResponseBody {
	s.TaskId = &v
	return s
}

func (s *ResetAccountPasswordResponseBody) SetRequestId(v string) *ResetAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetAccountPasswordResponseBody) SetDBClusterId(v string) *ResetAccountPasswordResponseBody {
	s.DBClusterId = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s RevokeOperatorPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeOperatorPermissionRequest) GoString() string {
	return s.String()
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

func (s *RevokeOperatorPermissionRequest) SetOwnerAccount(v string) *RevokeOperatorPermissionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RevokeOperatorPermissionRequest) SetDBClusterId(v string) *RevokeOperatorPermissionRequest {
	s.DBClusterId = &v
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

type UnbindDBResourcePoolWithUserRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	PoolName             *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	PoolUser             *string `json:"PoolUser,omitempty" xml:"PoolUser,omitempty"`
}

func (s UnbindDBResourcePoolWithUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindDBResourcePoolWithUserRequest) GoString() string {
	return s.String()
}

func (s *UnbindDBResourcePoolWithUserRequest) SetOwnerId(v int64) *UnbindDBResourcePoolWithUserRequest {
	s.OwnerId = &v
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

func (s *UnbindDBResourcePoolWithUserRequest) SetOwnerAccount(v string) *UnbindDBResourcePoolWithUserRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnbindDBResourcePoolWithUserRequest) SetDBClusterId(v string) *UnbindDBResourcePoolWithUserRequest {
	s.DBClusterId = &v
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AllocateClusterPublicConnectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AllocateClusterPublicConnection"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BindDBResourcePoolWithUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BindDBResourcePoolWithUser"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAccount"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDBClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDBCluster"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDBResourcePoolResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDBResourcePool"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateElasticPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateElasticPlan"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAccount"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDBClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDBCluster"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDBResourcePoolResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDBResourcePool"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteElasticPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteElasticPlan"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAccountsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAccounts"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAllAccountsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAllAccounts"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAllDataSourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAllDataSource"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAuditLogConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAuditLogConfig"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAuditLogRecordsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAuditLogRecords"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAutoRenewAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAutoRenewAttribute"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAvailableResourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAvailableResource"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackupPolicy"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackups"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeColumnsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeColumns"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeConnectionCountRecordsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeConnectionCountRecords"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBClusterAccessWhiteListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBClusterAccessWhiteList"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBClusterAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBClusterAttribute"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeDBClusterNetInfoWithOptions(request *DescribeDBClusterNetInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterNetInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBClusterNetInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBClusterNetInfo"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBClusterPerformanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBClusterPerformance"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBClusterResourcePoolPerformanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBClusterResourcePoolPerformance"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBClustersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBClusters"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBResourcePoolResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBResourcePool"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeElasticDailyPlanWithOptions(request *DescribeElasticDailyPlanRequest, runtime *util.RuntimeOptions) (_result *DescribeElasticDailyPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeElasticDailyPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeElasticDailyPlan"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeElasticPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeElasticPlan"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInclinedTablesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInclinedTables"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeLoadTasksRecordsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeLoadTasksRecords"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeOperatorPermissionWithOptions(request *DescribeOperatorPermissionRequest, runtime *util.RuntimeOptions) (_result *DescribeOperatorPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeOperatorPermissionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeOperatorPermission"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeProcessListWithOptions(request *DescribeProcessListRequest, runtime *util.RuntimeOptions) (_result *DescribeProcessListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeProcessListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeProcessList"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeSchemasWithOptions(request *DescribeSchemasRequest, runtime *util.RuntimeOptions) (_result *DescribeSchemasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSchemasResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSchemas"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSlowLogRecordsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSlowLogRecords"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSlowLogTrendResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSlowLogTrend"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeSQLPlanWithOptions(request *DescribeSQLPlanRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSQLPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSQLPlan"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSQLPlanTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSQLPlanTask"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeTableDetailWithOptions(request *DescribeTableDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeTableDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeTableDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTableDetail"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeTablePartitionDiagnoseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTablePartitionDiagnose"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeTablesWithOptions(request *DescribeTablesRequest, runtime *util.RuntimeOptions) (_result *DescribeTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeTablesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTables"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeTableStatisticsWithOptions(request *DescribeTableStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeTableStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeTableStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTableStatistics"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeTaskInfoWithOptions(request *DescribeTaskInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeTaskInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeTaskInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTaskInfo"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GrantOperatorPermissionWithOptions(request *GrantOperatorPermissionRequest, runtime *util.RuntimeOptions) (_result *GrantOperatorPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GrantOperatorPermissionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GrantOperatorPermission"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &KillProcessResponse{}
	_body, _err := client.DoRPCRequest(tea.String("KillProcess"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagResources"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyAccountDescriptionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyAccountDescription"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyAuditLogConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyAuditLogConfig"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyAutoRenewAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyAutoRenewAttribute"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyBackupPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyBackupPolicy"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyClusterConnectionStringResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyClusterConnectionString"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBCluster"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBClusterAccessWhiteListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBClusterAccessWhiteList"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBClusterDescriptionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBClusterDescription"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBClusterMaintainTimeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBClusterMaintainTime"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBClusterResourceGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBClusterResourceGroup"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBResourcePoolResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBResourcePool"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyElasticPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyElasticPlan"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyLogBackupPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyLogBackupPolicy"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ReleaseClusterPublicConnectionWithOptions(request *ReleaseClusterPublicConnectionRequest, runtime *util.RuntimeOptions) (_result *ReleaseClusterPublicConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReleaseClusterPublicConnectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReleaseClusterPublicConnection"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResetAccountPasswordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResetAccountPassword"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RevokeOperatorPermissionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RevokeOperatorPermission"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnbindDBResourcePoolWithUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnbindDBResourcePoolWithUser"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UntagResources"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
