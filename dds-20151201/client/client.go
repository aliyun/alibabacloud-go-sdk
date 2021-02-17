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

type AllocateNodePrivateNetworkAddressRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPassword      *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
}

func (s AllocateNodePrivateNetworkAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s AllocateNodePrivateNetworkAddressRequest) GoString() string {
	return s.String()
}

func (s *AllocateNodePrivateNetworkAddressRequest) SetSecurityToken(v string) *AllocateNodePrivateNetworkAddressRequest {
	s.SecurityToken = &v
	return s
}

func (s *AllocateNodePrivateNetworkAddressRequest) SetOwnerId(v int64) *AllocateNodePrivateNetworkAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocateNodePrivateNetworkAddressRequest) SetResourceOwnerAccount(v string) *AllocateNodePrivateNetworkAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocateNodePrivateNetworkAddressRequest) SetResourceOwnerId(v int64) *AllocateNodePrivateNetworkAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AllocateNodePrivateNetworkAddressRequest) SetOwnerAccount(v string) *AllocateNodePrivateNetworkAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AllocateNodePrivateNetworkAddressRequest) SetZoneId(v string) *AllocateNodePrivateNetworkAddressRequest {
	s.ZoneId = &v
	return s
}

func (s *AllocateNodePrivateNetworkAddressRequest) SetDBInstanceId(v string) *AllocateNodePrivateNetworkAddressRequest {
	s.DBInstanceId = &v
	return s
}

func (s *AllocateNodePrivateNetworkAddressRequest) SetNodeId(v string) *AllocateNodePrivateNetworkAddressRequest {
	s.NodeId = &v
	return s
}

func (s *AllocateNodePrivateNetworkAddressRequest) SetAccountName(v string) *AllocateNodePrivateNetworkAddressRequest {
	s.AccountName = &v
	return s
}

func (s *AllocateNodePrivateNetworkAddressRequest) SetAccountPassword(v string) *AllocateNodePrivateNetworkAddressRequest {
	s.AccountPassword = &v
	return s
}

type AllocateNodePrivateNetworkAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateNodePrivateNetworkAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AllocateNodePrivateNetworkAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateNodePrivateNetworkAddressResponseBody) SetRequestId(v string) *AllocateNodePrivateNetworkAddressResponseBody {
	s.RequestId = &v
	return s
}

type AllocateNodePrivateNetworkAddressResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AllocateNodePrivateNetworkAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AllocateNodePrivateNetworkAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s AllocateNodePrivateNetworkAddressResponse) GoString() string {
	return s.String()
}

func (s *AllocateNodePrivateNetworkAddressResponse) SetHeaders(v map[string]*string) *AllocateNodePrivateNetworkAddressResponse {
	s.Headers = v
	return s
}

func (s *AllocateNodePrivateNetworkAddressResponse) SetBody(v *AllocateNodePrivateNetworkAddressResponseBody) *AllocateNodePrivateNetworkAddressResponse {
	s.Body = v
	return s
}

type AllocatePublicNetworkAddressRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s AllocatePublicNetworkAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s AllocatePublicNetworkAddressRequest) GoString() string {
	return s.String()
}

func (s *AllocatePublicNetworkAddressRequest) SetSecurityToken(v string) *AllocatePublicNetworkAddressRequest {
	s.SecurityToken = &v
	return s
}

func (s *AllocatePublicNetworkAddressRequest) SetOwnerId(v int64) *AllocatePublicNetworkAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocatePublicNetworkAddressRequest) SetResourceOwnerAccount(v string) *AllocatePublicNetworkAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocatePublicNetworkAddressRequest) SetResourceOwnerId(v int64) *AllocatePublicNetworkAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AllocatePublicNetworkAddressRequest) SetOwnerAccount(v string) *AllocatePublicNetworkAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AllocatePublicNetworkAddressRequest) SetDBInstanceId(v string) *AllocatePublicNetworkAddressRequest {
	s.DBInstanceId = &v
	return s
}

func (s *AllocatePublicNetworkAddressRequest) SetNodeId(v string) *AllocatePublicNetworkAddressRequest {
	s.NodeId = &v
	return s
}

type AllocatePublicNetworkAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocatePublicNetworkAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AllocatePublicNetworkAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AllocatePublicNetworkAddressResponseBody) SetRequestId(v string) *AllocatePublicNetworkAddressResponseBody {
	s.RequestId = &v
	return s
}

type AllocatePublicNetworkAddressResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AllocatePublicNetworkAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AllocatePublicNetworkAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s AllocatePublicNetworkAddressResponse) GoString() string {
	return s.String()
}

func (s *AllocatePublicNetworkAddressResponse) SetHeaders(v map[string]*string) *AllocatePublicNetworkAddressResponse {
	s.Headers = v
	return s
}

func (s *AllocatePublicNetworkAddressResponse) SetBody(v *AllocatePublicNetworkAddressResponseBody) *AllocatePublicNetworkAddressResponse {
	s.Body = v
	return s
}

type CheckCloudResourceAuthorizedRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	TargetRegionId       *string `json:"TargetRegionId,omitempty" xml:"TargetRegionId,omitempty"`
}

func (s CheckCloudResourceAuthorizedRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckCloudResourceAuthorizedRequest) GoString() string {
	return s.String()
}

func (s *CheckCloudResourceAuthorizedRequest) SetSecurityToken(v string) *CheckCloudResourceAuthorizedRequest {
	s.SecurityToken = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetOwnerId(v int64) *CheckCloudResourceAuthorizedRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetResourceOwnerAccount(v string) *CheckCloudResourceAuthorizedRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetResourceOwnerId(v int64) *CheckCloudResourceAuthorizedRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetOwnerAccount(v string) *CheckCloudResourceAuthorizedRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetDBInstanceId(v string) *CheckCloudResourceAuthorizedRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetTargetRegionId(v string) *CheckCloudResourceAuthorizedRequest {
	s.TargetRegionId = &v
	return s
}

type CheckCloudResourceAuthorizedResponseBody struct {
	AuthorizationState *int32  `json:"AuthorizationState,omitempty" xml:"AuthorizationState,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RoleArn            *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s CheckCloudResourceAuthorizedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckCloudResourceAuthorizedResponseBody) GoString() string {
	return s.String()
}

func (s *CheckCloudResourceAuthorizedResponseBody) SetAuthorizationState(v int32) *CheckCloudResourceAuthorizedResponseBody {
	s.AuthorizationState = &v
	return s
}

func (s *CheckCloudResourceAuthorizedResponseBody) SetRequestId(v string) *CheckCloudResourceAuthorizedResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckCloudResourceAuthorizedResponseBody) SetRoleArn(v string) *CheckCloudResourceAuthorizedResponseBody {
	s.RoleArn = &v
	return s
}

type CheckCloudResourceAuthorizedResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckCloudResourceAuthorizedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckCloudResourceAuthorizedResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckCloudResourceAuthorizedResponse) GoString() string {
	return s.String()
}

func (s *CheckCloudResourceAuthorizedResponse) SetHeaders(v map[string]*string) *CheckCloudResourceAuthorizedResponse {
	s.Headers = v
	return s
}

func (s *CheckCloudResourceAuthorizedResponse) SetBody(v *CheckCloudResourceAuthorizedResponseBody) *CheckCloudResourceAuthorizedResponse {
	s.Body = v
	return s
}

type CheckRecoveryConditionRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	SourceDBInstance     *string `json:"SourceDBInstance,omitempty" xml:"SourceDBInstance,omitempty"`
	DatabaseNames        *string `json:"DatabaseNames,omitempty" xml:"DatabaseNames,omitempty"`
	RestoreTime          *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	BackupId             *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CheckRecoveryConditionRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckRecoveryConditionRequest) GoString() string {
	return s.String()
}

func (s *CheckRecoveryConditionRequest) SetSecurityToken(v string) *CheckRecoveryConditionRequest {
	s.SecurityToken = &v
	return s
}

func (s *CheckRecoveryConditionRequest) SetOwnerId(v int64) *CheckRecoveryConditionRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckRecoveryConditionRequest) SetResourceOwnerAccount(v string) *CheckRecoveryConditionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckRecoveryConditionRequest) SetResourceOwnerId(v int64) *CheckRecoveryConditionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckRecoveryConditionRequest) SetOwnerAccount(v string) *CheckRecoveryConditionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckRecoveryConditionRequest) SetSourceDBInstance(v string) *CheckRecoveryConditionRequest {
	s.SourceDBInstance = &v
	return s
}

func (s *CheckRecoveryConditionRequest) SetDatabaseNames(v string) *CheckRecoveryConditionRequest {
	s.DatabaseNames = &v
	return s
}

func (s *CheckRecoveryConditionRequest) SetRestoreTime(v string) *CheckRecoveryConditionRequest {
	s.RestoreTime = &v
	return s
}

func (s *CheckRecoveryConditionRequest) SetBackupId(v string) *CheckRecoveryConditionRequest {
	s.BackupId = &v
	return s
}

func (s *CheckRecoveryConditionRequest) SetResourceGroupId(v string) *CheckRecoveryConditionRequest {
	s.ResourceGroupId = &v
	return s
}

type CheckRecoveryConditionResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	IsValid        *bool   `json:"IsValid,omitempty" xml:"IsValid,omitempty"`
}

func (s CheckRecoveryConditionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckRecoveryConditionResponseBody) GoString() string {
	return s.String()
}

func (s *CheckRecoveryConditionResponseBody) SetRequestId(v string) *CheckRecoveryConditionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckRecoveryConditionResponseBody) SetDBInstanceName(v string) *CheckRecoveryConditionResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *CheckRecoveryConditionResponseBody) SetIsValid(v bool) *CheckRecoveryConditionResponseBody {
	s.IsValid = &v
	return s
}

type CheckRecoveryConditionResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckRecoveryConditionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckRecoveryConditionResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckRecoveryConditionResponse) GoString() string {
	return s.String()
}

func (s *CheckRecoveryConditionResponse) SetHeaders(v map[string]*string) *CheckRecoveryConditionResponse {
	s.Headers = v
	return s
}

func (s *CheckRecoveryConditionResponse) SetBody(v *CheckRecoveryConditionResponseBody) *CheckRecoveryConditionResponse {
	s.Body = v
	return s
}

type CreateBackupRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	BackupMethod         *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
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

func (s *CreateBackupRequest) SetDBInstanceId(v string) *CreateBackupRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateBackupRequest) SetBackupMethod(v string) *CreateBackupRequest {
	s.BackupMethod = &v
	return s
}

type CreateBackupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BackupId  *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
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

func (s *CreateBackupResponseBody) SetBackupId(v string) *CreateBackupResponseBody {
	s.BackupId = &v
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

type CreateDBInstanceRequest struct {
	SecurityToken         *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	ClientToken           *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ZoneId                *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Engine                *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion         *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	DBInstanceClass       *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	DBInstanceStorage     *int32  `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	SecurityIPList        *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	AccountPassword       *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	ChargeType            *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Period                *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	NetworkType           *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	VpcId                 *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VSwitchId             *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	SrcDBInstanceId       *string `json:"SrcDBInstanceId,omitempty" xml:"SrcDBInstanceId,omitempty"`
	BackupId              *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	RestoreTime           *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	BusinessInfo          *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	AutoRenew             *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	DatabaseNames         *string `json:"DatabaseNames,omitempty" xml:"DatabaseNames,omitempty"`
	CouponNo              *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	StorageEngine         *string `json:"StorageEngine,omitempty" xml:"StorageEngine,omitempty"`
	ReplicationFactor     *string `json:"ReplicationFactor,omitempty" xml:"ReplicationFactor,omitempty"`
	ReadonlyReplicas      *string `json:"ReadonlyReplicas,omitempty" xml:"ReadonlyReplicas,omitempty"`
	ResourceGroupId       *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ClusterId             *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s CreateDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceRequest) SetSecurityToken(v string) *CreateDBInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateDBInstanceRequest) SetOwnerId(v int64) *CreateDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetResourceOwnerAccount(v string) *CreateDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBInstanceRequest) SetResourceOwnerId(v int64) *CreateDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetOwnerAccount(v string) *CreateDBInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBInstanceRequest) SetClientToken(v string) *CreateDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceRequest) SetZoneId(v string) *CreateDBInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEngine(v string) *CreateDBInstanceRequest {
	s.Engine = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEngineVersion(v string) *CreateDBInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceClass(v string) *CreateDBInstanceRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceStorage(v int32) *CreateDBInstanceRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceDescription(v string) *CreateDBInstanceRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSecurityIPList(v string) *CreateDBInstanceRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateDBInstanceRequest) SetAccountPassword(v string) *CreateDBInstanceRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateDBInstanceRequest) SetChargeType(v string) *CreateDBInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPeriod(v int32) *CreateDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateDBInstanceRequest) SetNetworkType(v string) *CreateDBInstanceRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVpcId(v string) *CreateDBInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVSwitchId(v string) *CreateDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSrcDBInstanceId(v string) *CreateDBInstanceRequest {
	s.SrcDBInstanceId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetBackupId(v string) *CreateDBInstanceRequest {
	s.BackupId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetRestoreTime(v string) *CreateDBInstanceRequest {
	s.RestoreTime = &v
	return s
}

func (s *CreateDBInstanceRequest) SetBusinessInfo(v string) *CreateDBInstanceRequest {
	s.BusinessInfo = &v
	return s
}

func (s *CreateDBInstanceRequest) SetAutoRenew(v string) *CreateDBInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDatabaseNames(v string) *CreateDBInstanceRequest {
	s.DatabaseNames = &v
	return s
}

func (s *CreateDBInstanceRequest) SetCouponNo(v string) *CreateDBInstanceRequest {
	s.CouponNo = &v
	return s
}

func (s *CreateDBInstanceRequest) SetStorageEngine(v string) *CreateDBInstanceRequest {
	s.StorageEngine = &v
	return s
}

func (s *CreateDBInstanceRequest) SetReplicationFactor(v string) *CreateDBInstanceRequest {
	s.ReplicationFactor = &v
	return s
}

func (s *CreateDBInstanceRequest) SetReadonlyReplicas(v string) *CreateDBInstanceRequest {
	s.ReadonlyReplicas = &v
	return s
}

func (s *CreateDBInstanceRequest) SetResourceGroupId(v string) *CreateDBInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetClusterId(v string) *CreateDBInstanceRequest {
	s.ClusterId = &v
	return s
}

type CreateDBInstanceResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OrderId      *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceResponseBody) SetRequestId(v string) *CreateDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetDBInstanceId(v string) *CreateDBInstanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetOrderId(v string) *CreateDBInstanceResponseBody {
	s.OrderId = &v
	return s
}

type CreateDBInstanceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceResponse) SetHeaders(v map[string]*string) *CreateDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateDBInstanceResponse) SetBody(v *CreateDBInstanceResponseBody) *CreateDBInstanceResponse {
	s.Body = v
	return s
}

type CreateNodeRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NodeClass            *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	NodeStorage          *int32  `json:"NodeStorage,omitempty" xml:"NodeStorage,omitempty"`
	NodeType             *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	FromApp              *string `json:"FromApp,omitempty" xml:"FromApp,omitempty"`
	AutoPay              *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	ReadonlyReplicas     *int32  `json:"ReadonlyReplicas,omitempty" xml:"ReadonlyReplicas,omitempty"`
}

func (s CreateNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeRequest) GoString() string {
	return s.String()
}

func (s *CreateNodeRequest) SetSecurityToken(v string) *CreateNodeRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateNodeRequest) SetOwnerId(v int64) *CreateNodeRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateNodeRequest) SetResourceOwnerAccount(v string) *CreateNodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateNodeRequest) SetResourceOwnerId(v int64) *CreateNodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateNodeRequest) SetOwnerAccount(v string) *CreateNodeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateNodeRequest) SetDBInstanceId(v string) *CreateNodeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateNodeRequest) SetNodeClass(v string) *CreateNodeRequest {
	s.NodeClass = &v
	return s
}

func (s *CreateNodeRequest) SetNodeStorage(v int32) *CreateNodeRequest {
	s.NodeStorage = &v
	return s
}

func (s *CreateNodeRequest) SetNodeType(v string) *CreateNodeRequest {
	s.NodeType = &v
	return s
}

func (s *CreateNodeRequest) SetClientToken(v string) *CreateNodeRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateNodeRequest) SetFromApp(v string) *CreateNodeRequest {
	s.FromApp = &v
	return s
}

func (s *CreateNodeRequest) SetAutoPay(v bool) *CreateNodeRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateNodeRequest) SetReadonlyReplicas(v int32) *CreateNodeRequest {
	s.ReadonlyReplicas = &v
	return s
}

type CreateNodeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	NodeId    *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNodeResponseBody) SetRequestId(v string) *CreateNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNodeResponseBody) SetNodeId(v string) *CreateNodeResponseBody {
	s.NodeId = &v
	return s
}

func (s *CreateNodeResponseBody) SetOrderId(v string) *CreateNodeResponseBody {
	s.OrderId = &v
	return s
}

type CreateNodeResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeResponse) GoString() string {
	return s.String()
}

func (s *CreateNodeResponse) SetHeaders(v map[string]*string) *CreateNodeResponse {
	s.Headers = v
	return s
}

func (s *CreateNodeResponse) SetBody(v *CreateNodeResponseBody) *CreateNodeResponse {
	s.Body = v
	return s
}

type CreateRecommendationTaskRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s CreateRecommendationTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRecommendationTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateRecommendationTaskRequest) SetSecurityToken(v string) *CreateRecommendationTaskRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateRecommendationTaskRequest) SetOwnerId(v int64) *CreateRecommendationTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateRecommendationTaskRequest) SetResourceOwnerAccount(v string) *CreateRecommendationTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateRecommendationTaskRequest) SetResourceOwnerId(v int64) *CreateRecommendationTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateRecommendationTaskRequest) SetOwnerAccount(v string) *CreateRecommendationTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateRecommendationTaskRequest) SetInstanceId(v string) *CreateRecommendationTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRecommendationTaskRequest) SetNodeId(v string) *CreateRecommendationTaskRequest {
	s.NodeId = &v
	return s
}

func (s *CreateRecommendationTaskRequest) SetStartTime(v string) *CreateRecommendationTaskRequest {
	s.StartTime = &v
	return s
}

func (s *CreateRecommendationTaskRequest) SetEndTime(v string) *CreateRecommendationTaskRequest {
	s.EndTime = &v
	return s
}

type CreateRecommendationTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRecommendationTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRecommendationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRecommendationTaskResponseBody) SetRequestId(v string) *CreateRecommendationTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateRecommendationTaskResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRecommendationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRecommendationTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRecommendationTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateRecommendationTaskResponse) SetHeaders(v map[string]*string) *CreateRecommendationTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateRecommendationTaskResponse) SetBody(v *CreateRecommendationTaskResponseBody) *CreateRecommendationTaskResponse {
	s.Body = v
	return s
}

type CreateServerlessDBInstanceRequest struct {
	SecurityToken         *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceStorage     *int32  `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	ZoneId                *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Engine                *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion         *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	SecurityIPList        *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	AccountPassword       *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	ChargeType            *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Period                *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	NetworkType           *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	VpcId                 *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VSwitchId             *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ClientToken           *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	StorageEngine         *string `json:"StorageEngine,omitempty" xml:"StorageEngine,omitempty"`
	AutoRenew             *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ResourceGroupId       *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	PeriodPriceType       *string `json:"PeriodPriceType,omitempty" xml:"PeriodPriceType,omitempty"`
}

func (s CreateServerlessDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServerlessDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateServerlessDBInstanceRequest) SetSecurityToken(v string) *CreateServerlessDBInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateServerlessDBInstanceRequest) SetOwnerId(v int64) *CreateServerlessDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateServerlessDBInstanceRequest) SetResourceOwnerAccount(v string) *CreateServerlessDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateServerlessDBInstanceRequest) SetResourceOwnerId(v int64) *CreateServerlessDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateServerlessDBInstanceRequest) SetOwnerAccount(v string) *CreateServerlessDBInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateServerlessDBInstanceRequest) SetRegionId(v string) *CreateServerlessDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServerlessDBInstanceRequest) SetDBInstanceStorage(v int32) *CreateServerlessDBInstanceRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *CreateServerlessDBInstanceRequest) SetZoneId(v string) *CreateServerlessDBInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateServerlessDBInstanceRequest) SetEngine(v string) *CreateServerlessDBInstanceRequest {
	s.Engine = &v
	return s
}

func (s *CreateServerlessDBInstanceRequest) SetEngineVersion(v string) *CreateServerlessDBInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateServerlessDBInstanceRequest) SetDBInstanceDescription(v string) *CreateServerlessDBInstanceRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CreateServerlessDBInstanceRequest) SetSecurityIPList(v string) *CreateServerlessDBInstanceRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateServerlessDBInstanceRequest) SetAccountPassword(v string) *CreateServerlessDBInstanceRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateServerlessDBInstanceRequest) SetChargeType(v string) *CreateServerlessDBInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateServerlessDBInstanceRequest) SetPeriod(v int32) *CreateServerlessDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateServerlessDBInstanceRequest) SetNetworkType(v string) *CreateServerlessDBInstanceRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateServerlessDBInstanceRequest) SetVpcId(v string) *CreateServerlessDBInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateServerlessDBInstanceRequest) SetVSwitchId(v string) *CreateServerlessDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateServerlessDBInstanceRequest) SetClientToken(v string) *CreateServerlessDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServerlessDBInstanceRequest) SetStorageEngine(v string) *CreateServerlessDBInstanceRequest {
	s.StorageEngine = &v
	return s
}

func (s *CreateServerlessDBInstanceRequest) SetAutoRenew(v string) *CreateServerlessDBInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateServerlessDBInstanceRequest) SetResourceGroupId(v string) *CreateServerlessDBInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateServerlessDBInstanceRequest) SetPeriodPriceType(v string) *CreateServerlessDBInstanceRequest {
	s.PeriodPriceType = &v
	return s
}

type CreateServerlessDBInstanceResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OrderId      *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateServerlessDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServerlessDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServerlessDBInstanceResponseBody) SetRequestId(v string) *CreateServerlessDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServerlessDBInstanceResponseBody) SetDBInstanceId(v string) *CreateServerlessDBInstanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *CreateServerlessDBInstanceResponseBody) SetOrderId(v string) *CreateServerlessDBInstanceResponseBody {
	s.OrderId = &v
	return s
}

type CreateServerlessDBInstanceResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateServerlessDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServerlessDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServerlessDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateServerlessDBInstanceResponse) SetHeaders(v map[string]*string) *CreateServerlessDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateServerlessDBInstanceResponse) SetBody(v *CreateServerlessDBInstanceResponseBody) *CreateServerlessDBInstanceResponse {
	s.Body = v
	return s
}

type CreateShardingDBInstanceRequest struct {
	SecurityToken         *string                                        `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId               *int64                                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string                                        `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64                                         `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount          *string                                        `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	ZoneId                *string                                        `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Engine                *string                                        `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion         *string                                        `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	DBInstanceDescription *string                                        `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	SecurityIPList        *string                                        `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	AccountPassword       *string                                        `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	ChargeType            *string                                        `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Period                *int32                                         `json:"Period,omitempty" xml:"Period,omitempty"`
	NetworkType           *string                                        `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	VpcId                 *string                                        `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VSwitchId             *string                                        `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	SrcDBInstanceId       *string                                        `json:"SrcDBInstanceId,omitempty" xml:"SrcDBInstanceId,omitempty"`
	RestoreTime           *string                                        `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	ClientToken           *string                                        `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	StorageEngine         *string                                        `json:"StorageEngine,omitempty" xml:"StorageEngine,omitempty"`
	AutoRenew             *string                                        `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ProtocolType          *string                                        `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	Mongos                []*CreateShardingDBInstanceRequestMongos       `json:"Mongos,omitempty" xml:"Mongos,omitempty" type:"Repeated"`
	ReplicaSet            []*CreateShardingDBInstanceRequestReplicaSet   `json:"ReplicaSet,omitempty" xml:"ReplicaSet,omitempty" type:"Repeated"`
	ConfigServer          []*CreateShardingDBInstanceRequestConfigServer `json:"ConfigServer,omitempty" xml:"ConfigServer,omitempty" type:"Repeated"`
}

func (s CreateShardingDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateShardingDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateShardingDBInstanceRequest) SetSecurityToken(v string) *CreateShardingDBInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetOwnerId(v int64) *CreateShardingDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetResourceOwnerAccount(v string) *CreateShardingDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetResourceOwnerId(v int64) *CreateShardingDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetOwnerAccount(v string) *CreateShardingDBInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetZoneId(v string) *CreateShardingDBInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetEngine(v string) *CreateShardingDBInstanceRequest {
	s.Engine = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetEngineVersion(v string) *CreateShardingDBInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetDBInstanceDescription(v string) *CreateShardingDBInstanceRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetSecurityIPList(v string) *CreateShardingDBInstanceRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetAccountPassword(v string) *CreateShardingDBInstanceRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetChargeType(v string) *CreateShardingDBInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetPeriod(v int32) *CreateShardingDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetNetworkType(v string) *CreateShardingDBInstanceRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetVpcId(v string) *CreateShardingDBInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetVSwitchId(v string) *CreateShardingDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetSrcDBInstanceId(v string) *CreateShardingDBInstanceRequest {
	s.SrcDBInstanceId = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetRestoreTime(v string) *CreateShardingDBInstanceRequest {
	s.RestoreTime = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetClientToken(v string) *CreateShardingDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetStorageEngine(v string) *CreateShardingDBInstanceRequest {
	s.StorageEngine = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetAutoRenew(v string) *CreateShardingDBInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetProtocolType(v string) *CreateShardingDBInstanceRequest {
	s.ProtocolType = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetMongos(v []*CreateShardingDBInstanceRequestMongos) *CreateShardingDBInstanceRequest {
	s.Mongos = v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetReplicaSet(v []*CreateShardingDBInstanceRequestReplicaSet) *CreateShardingDBInstanceRequest {
	s.ReplicaSet = v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetConfigServer(v []*CreateShardingDBInstanceRequestConfigServer) *CreateShardingDBInstanceRequest {
	s.ConfigServer = v
	return s
}

type CreateShardingDBInstanceRequestMongos struct {
	Class *string `json:"Class,omitempty" xml:"Class,omitempty"`
}

func (s CreateShardingDBInstanceRequestMongos) String() string {
	return tea.Prettify(s)
}

func (s CreateShardingDBInstanceRequestMongos) GoString() string {
	return s.String()
}

func (s *CreateShardingDBInstanceRequestMongos) SetClass(v string) *CreateShardingDBInstanceRequestMongos {
	s.Class = &v
	return s
}

type CreateShardingDBInstanceRequestReplicaSet struct {
	Class            *string `json:"Class,omitempty" xml:"Class,omitempty"`
	Storage          *int32  `json:"Storage,omitempty" xml:"Storage,omitempty"`
	ReadonlyReplicas *int32  `json:"ReadonlyReplicas,omitempty" xml:"ReadonlyReplicas,omitempty"`
}

func (s CreateShardingDBInstanceRequestReplicaSet) String() string {
	return tea.Prettify(s)
}

func (s CreateShardingDBInstanceRequestReplicaSet) GoString() string {
	return s.String()
}

func (s *CreateShardingDBInstanceRequestReplicaSet) SetClass(v string) *CreateShardingDBInstanceRequestReplicaSet {
	s.Class = &v
	return s
}

func (s *CreateShardingDBInstanceRequestReplicaSet) SetStorage(v int32) *CreateShardingDBInstanceRequestReplicaSet {
	s.Storage = &v
	return s
}

func (s *CreateShardingDBInstanceRequestReplicaSet) SetReadonlyReplicas(v int32) *CreateShardingDBInstanceRequestReplicaSet {
	s.ReadonlyReplicas = &v
	return s
}

type CreateShardingDBInstanceRequestConfigServer struct {
	Class   *string `json:"Class,omitempty" xml:"Class,omitempty"`
	Storage *int32  `json:"Storage,omitempty" xml:"Storage,omitempty"`
}

func (s CreateShardingDBInstanceRequestConfigServer) String() string {
	return tea.Prettify(s)
}

func (s CreateShardingDBInstanceRequestConfigServer) GoString() string {
	return s.String()
}

func (s *CreateShardingDBInstanceRequestConfigServer) SetClass(v string) *CreateShardingDBInstanceRequestConfigServer {
	s.Class = &v
	return s
}

func (s *CreateShardingDBInstanceRequestConfigServer) SetStorage(v int32) *CreateShardingDBInstanceRequestConfigServer {
	s.Storage = &v
	return s
}

type CreateShardingDBInstanceResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OrderId      *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateShardingDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateShardingDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateShardingDBInstanceResponseBody) SetRequestId(v string) *CreateShardingDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateShardingDBInstanceResponseBody) SetDBInstanceId(v string) *CreateShardingDBInstanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *CreateShardingDBInstanceResponseBody) SetOrderId(v string) *CreateShardingDBInstanceResponseBody {
	s.OrderId = &v
	return s
}

type CreateShardingDBInstanceResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateShardingDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateShardingDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateShardingDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateShardingDBInstanceResponse) SetHeaders(v map[string]*string) *CreateShardingDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateShardingDBInstanceResponse) SetBody(v *CreateShardingDBInstanceResponseBody) *CreateShardingDBInstanceResponse {
	s.Body = v
	return s
}

type DeleteDBInstanceRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DeleteDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceRequest) SetSecurityToken(v string) *DeleteDBInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetOwnerId(v int64) *DeleteDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetResourceOwnerAccount(v string) *DeleteDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetResourceOwnerId(v int64) *DeleteDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetOwnerAccount(v string) *DeleteDBInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetDBInstanceId(v string) *DeleteDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetClientToken(v string) *DeleteDBInstanceRequest {
	s.ClientToken = &v
	return s
}

type DeleteDBInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceResponseBody) SetRequestId(v string) *DeleteDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDBInstanceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceResponse) SetHeaders(v map[string]*string) *DeleteDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBInstanceResponse) SetBody(v *DeleteDBInstanceResponseBody) *DeleteDBInstanceResponse {
	s.Body = v
	return s
}

type DeleteNodeRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DeleteNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteNodeRequest) SetSecurityToken(v string) *DeleteNodeRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteNodeRequest) SetOwnerId(v int64) *DeleteNodeRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteNodeRequest) SetResourceOwnerAccount(v string) *DeleteNodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteNodeRequest) SetResourceOwnerId(v int64) *DeleteNodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteNodeRequest) SetOwnerAccount(v string) *DeleteNodeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteNodeRequest) SetDBInstanceId(v string) *DeleteNodeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteNodeRequest) SetNodeId(v string) *DeleteNodeRequest {
	s.NodeId = &v
	return s
}

func (s *DeleteNodeRequest) SetClientToken(v string) *DeleteNodeRequest {
	s.ClientToken = &v
	return s
}

type DeleteNodeResponseBody struct {
	TaskId    *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s DeleteNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNodeResponseBody) SetTaskId(v int32) *DeleteNodeResponseBody {
	s.TaskId = &v
	return s
}

func (s *DeleteNodeResponseBody) SetRequestId(v string) *DeleteNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNodeResponseBody) SetOrderId(v string) *DeleteNodeResponseBody {
	s.OrderId = &v
	return s
}

type DeleteNodeResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteNodeResponse) SetHeaders(v map[string]*string) *DeleteNodeResponse {
	s.Headers = v
	return s
}

func (s *DeleteNodeResponse) SetBody(v *DeleteNodeResponseBody) *DeleteNodeResponse {
	s.Body = v
	return s
}

type DescribeAccountsRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
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

func (s *DescribeAccountsRequest) SetDBInstanceId(v string) *DescribeAccountsRequest {
	s.DBInstanceId = &v
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
	CharacterType      *string `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	AccountStatus      *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	DBInstanceId       *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	AccountName        *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s DescribeAccountsResponseBodyAccountsAccount) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountsAccount) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountsAccount) SetCharacterType(v string) *DescribeAccountsResponseBodyAccountsAccount {
	s.CharacterType = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsAccount) SetAccountStatus(v string) *DescribeAccountsResponseBodyAccountsAccount {
	s.AccountStatus = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsAccount) SetAccountDescription(v string) *DescribeAccountsResponseBodyAccountsAccount {
	s.AccountDescription = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsAccount) SetDBInstanceId(v string) *DescribeAccountsResponseBodyAccountsAccount {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsAccount) SetAccountName(v string) *DescribeAccountsResponseBodyAccountsAccount {
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

type DescribeActiveOperationTaskCountRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeActiveOperationTaskCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveOperationTaskCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskCountRequest) SetSecurityToken(v string) *DescribeActiveOperationTaskCountRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeActiveOperationTaskCountRequest) SetOwnerId(v int64) *DescribeActiveOperationTaskCountRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeActiveOperationTaskCountRequest) SetResourceOwnerAccount(v string) *DescribeActiveOperationTaskCountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeActiveOperationTaskCountRequest) SetResourceOwnerId(v int64) *DescribeActiveOperationTaskCountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeActiveOperationTaskCountRequest) SetOwnerAccount(v string) *DescribeActiveOperationTaskCountRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeActiveOperationTaskCountRequest) SetResourceGroupId(v string) *DescribeActiveOperationTaskCountRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeActiveOperationTaskCountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	NeedPop   *int32  `json:"NeedPop,omitempty" xml:"NeedPop,omitempty"`
	TaskCount *int32  `json:"TaskCount,omitempty" xml:"TaskCount,omitempty"`
}

func (s DescribeActiveOperationTaskCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveOperationTaskCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskCountResponseBody) SetRequestId(v string) *DescribeActiveOperationTaskCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeActiveOperationTaskCountResponseBody) SetNeedPop(v int32) *DescribeActiveOperationTaskCountResponseBody {
	s.NeedPop = &v
	return s
}

func (s *DescribeActiveOperationTaskCountResponseBody) SetTaskCount(v int32) *DescribeActiveOperationTaskCountResponseBody {
	s.TaskCount = &v
	return s
}

type DescribeActiveOperationTaskCountResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeActiveOperationTaskCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeActiveOperationTaskCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveOperationTaskCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskCountResponse) SetHeaders(v map[string]*string) *DescribeActiveOperationTaskCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeActiveOperationTaskCountResponse) SetBody(v *DescribeActiveOperationTaskCountResponseBody) *DescribeActiveOperationTaskCountResponse {
	s.Body = v
	return s
}

type DescribeActiveOperationTaskTypeRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	IsHistory            *int32  `json:"IsHistory,omitempty" xml:"IsHistory,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeActiveOperationTaskTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveOperationTaskTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskTypeRequest) SetSecurityToken(v string) *DescribeActiveOperationTaskTypeRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeRequest) SetOwnerId(v int64) *DescribeActiveOperationTaskTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeRequest) SetResourceOwnerAccount(v string) *DescribeActiveOperationTaskTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeRequest) SetResourceOwnerId(v int64) *DescribeActiveOperationTaskTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeRequest) SetOwnerAccount(v string) *DescribeActiveOperationTaskTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeRequest) SetIsHistory(v int32) *DescribeActiveOperationTaskTypeRequest {
	s.IsHistory = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeRequest) SetResourceGroupId(v string) *DescribeActiveOperationTaskTypeRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeActiveOperationTaskTypeResponseBody struct {
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TypeList  []*DescribeActiveOperationTaskTypeResponseBodyTypeList `json:"TypeList,omitempty" xml:"TypeList,omitempty" type:"Repeated"`
}

func (s DescribeActiveOperationTaskTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveOperationTaskTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskTypeResponseBody) SetRequestId(v string) *DescribeActiveOperationTaskTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeResponseBody) SetTypeList(v []*DescribeActiveOperationTaskTypeResponseBodyTypeList) *DescribeActiveOperationTaskTypeResponseBody {
	s.TypeList = v
	return s
}

type DescribeActiveOperationTaskTypeResponseBodyTypeList struct {
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	Count    *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeActiveOperationTaskTypeResponseBodyTypeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveOperationTaskTypeResponseBodyTypeList) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskTypeResponseBodyTypeList) SetTaskType(v string) *DescribeActiveOperationTaskTypeResponseBodyTypeList {
	s.TaskType = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeResponseBodyTypeList) SetCount(v int32) *DescribeActiveOperationTaskTypeResponseBodyTypeList {
	s.Count = &v
	return s
}

type DescribeActiveOperationTaskTypeResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeActiveOperationTaskTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeActiveOperationTaskTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveOperationTaskTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskTypeResponse) SetHeaders(v map[string]*string) *DescribeActiveOperationTaskTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeActiveOperationTaskTypeResponse) SetBody(v *DescribeActiveOperationTaskTypeResponseBody) *DescribeActiveOperationTaskTypeResponse {
	s.Body = v
	return s
}

type DescribeAuditFilesRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeAuditFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuditFilesRequest) SetSecurityToken(v string) *DescribeAuditFilesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeAuditFilesRequest) SetOwnerId(v int64) *DescribeAuditFilesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAuditFilesRequest) SetResourceOwnerAccount(v string) *DescribeAuditFilesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAuditFilesRequest) SetResourceOwnerId(v int64) *DescribeAuditFilesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAuditFilesRequest) SetOwnerAccount(v string) *DescribeAuditFilesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAuditFilesRequest) SetDBInstanceId(v string) *DescribeAuditFilesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAuditFilesRequest) SetNodeId(v string) *DescribeAuditFilesRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeAuditFilesRequest) SetPageSize(v int32) *DescribeAuditFilesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAuditFilesRequest) SetPageNumber(v int32) *DescribeAuditFilesRequest {
	s.PageNumber = &v
	return s
}

type DescribeAuditFilesResponseBody struct {
	TotalRecordCount *int32                               `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	PageRecordCount  *int32                               `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber       *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Items            *DescribeAuditFilesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeAuditFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditFilesResponseBody) SetTotalRecordCount(v int32) *DescribeAuditFilesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeAuditFilesResponseBody) SetPageRecordCount(v int32) *DescribeAuditFilesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeAuditFilesResponseBody) SetRequestId(v string) *DescribeAuditFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuditFilesResponseBody) SetPageNumber(v int32) *DescribeAuditFilesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAuditFilesResponseBody) SetItems(v *DescribeAuditFilesResponseBodyItems) *DescribeAuditFilesResponseBody {
	s.Items = v
	return s
}

type DescribeAuditFilesResponseBodyItems struct {
	LogFile []*DescribeAuditFilesResponseBodyItemsLogFile `json:"LogFile,omitempty" xml:"LogFile,omitempty" type:"Repeated"`
}

func (s DescribeAuditFilesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditFilesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAuditFilesResponseBodyItems) SetLogFile(v []*DescribeAuditFilesResponseBodyItemsLogFile) *DescribeAuditFilesResponseBodyItems {
	s.LogFile = v
	return s
}

type DescribeAuditFilesResponseBodyItemsLogFile struct {
	FileID         *int32  `json:"FileID,omitempty" xml:"FileID,omitempty"`
	LogStartTime   *string `json:"LogStartTime,omitempty" xml:"LogStartTime,omitempty"`
	LogSize        *int64  `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	LogDownloadURL *string `json:"LogDownloadURL,omitempty" xml:"LogDownloadURL,omitempty"`
	LogEndTime     *string `json:"LogEndTime,omitempty" xml:"LogEndTime,omitempty"`
	LogStatus      *string `json:"LogStatus,omitempty" xml:"LogStatus,omitempty"`
}

func (s DescribeAuditFilesResponseBodyItemsLogFile) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditFilesResponseBodyItemsLogFile) GoString() string {
	return s.String()
}

func (s *DescribeAuditFilesResponseBodyItemsLogFile) SetFileID(v int32) *DescribeAuditFilesResponseBodyItemsLogFile {
	s.FileID = &v
	return s
}

func (s *DescribeAuditFilesResponseBodyItemsLogFile) SetLogStartTime(v string) *DescribeAuditFilesResponseBodyItemsLogFile {
	s.LogStartTime = &v
	return s
}

func (s *DescribeAuditFilesResponseBodyItemsLogFile) SetLogSize(v int64) *DescribeAuditFilesResponseBodyItemsLogFile {
	s.LogSize = &v
	return s
}

func (s *DescribeAuditFilesResponseBodyItemsLogFile) SetLogDownloadURL(v string) *DescribeAuditFilesResponseBodyItemsLogFile {
	s.LogDownloadURL = &v
	return s
}

func (s *DescribeAuditFilesResponseBodyItemsLogFile) SetLogEndTime(v string) *DescribeAuditFilesResponseBodyItemsLogFile {
	s.LogEndTime = &v
	return s
}

func (s *DescribeAuditFilesResponseBodyItemsLogFile) SetLogStatus(v string) *DescribeAuditFilesResponseBodyItemsLogFile {
	s.LogStatus = &v
	return s
}

type DescribeAuditFilesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAuditFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAuditFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditFilesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuditFilesResponse) SetHeaders(v map[string]*string) *DescribeAuditFilesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuditFilesResponse) SetBody(v *DescribeAuditFilesResponseBody) *DescribeAuditFilesResponse {
	s.Body = v
	return s
}

type DescribeAuditLogFilterRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	RoleType             *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeAuditLogFilterRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogFilterRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogFilterRequest) SetSecurityToken(v string) *DescribeAuditLogFilterRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeAuditLogFilterRequest) SetOwnerId(v int64) *DescribeAuditLogFilterRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAuditLogFilterRequest) SetResourceOwnerAccount(v string) *DescribeAuditLogFilterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAuditLogFilterRequest) SetResourceOwnerId(v int64) *DescribeAuditLogFilterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAuditLogFilterRequest) SetOwnerAccount(v string) *DescribeAuditLogFilterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAuditLogFilterRequest) SetDBInstanceId(v string) *DescribeAuditLogFilterRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAuditLogFilterRequest) SetRoleType(v string) *DescribeAuditLogFilterRequest {
	s.RoleType = &v
	return s
}

type DescribeAuditLogFilterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Filter    *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	RoleType  *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeAuditLogFilterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogFilterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogFilterResponseBody) SetRequestId(v string) *DescribeAuditLogFilterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuditLogFilterResponseBody) SetFilter(v string) *DescribeAuditLogFilterResponseBody {
	s.Filter = &v
	return s
}

func (s *DescribeAuditLogFilterResponseBody) SetRoleType(v string) *DescribeAuditLogFilterResponseBody {
	s.RoleType = &v
	return s
}

type DescribeAuditLogFilterResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAuditLogFilterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAuditLogFilterResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogFilterResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogFilterResponse) SetHeaders(v map[string]*string) *DescribeAuditLogFilterResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuditLogFilterResponse) SetBody(v *DescribeAuditLogFilterResponseBody) *DescribeAuditLogFilterResponse {
	s.Body = v
	return s
}

type DescribeAuditPolicyRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeAuditPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuditPolicyRequest) SetSecurityToken(v string) *DescribeAuditPolicyRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeAuditPolicyRequest) SetOwnerId(v int64) *DescribeAuditPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAuditPolicyRequest) SetResourceOwnerAccount(v string) *DescribeAuditPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAuditPolicyRequest) SetResourceOwnerId(v int64) *DescribeAuditPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAuditPolicyRequest) SetOwnerAccount(v string) *DescribeAuditPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAuditPolicyRequest) SetDBInstanceId(v string) *DescribeAuditPolicyRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeAuditPolicyResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LogAuditStatus *string `json:"LogAuditStatus,omitempty" xml:"LogAuditStatus,omitempty"`
}

func (s DescribeAuditPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditPolicyResponseBody) SetRequestId(v string) *DescribeAuditPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuditPolicyResponseBody) SetLogAuditStatus(v string) *DescribeAuditPolicyResponseBody {
	s.LogAuditStatus = &v
	return s
}

type DescribeAuditPolicyResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAuditPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAuditPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuditPolicyResponse) SetHeaders(v map[string]*string) *DescribeAuditPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuditPolicyResponse) SetBody(v *DescribeAuditPolicyResponseBody) *DescribeAuditPolicyResponse {
	s.Body = v
	return s
}

type DescribeAuditRecordsRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Database             *string `json:"Database,omitempty" xml:"Database,omitempty"`
	User                 *string `json:"User,omitempty" xml:"User,omitempty"`
	Form                 *string `json:"Form,omitempty" xml:"Form,omitempty"`
	QueryKeywords        *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	OrderType            *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
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

func (s *DescribeAuditRecordsRequest) SetDBInstanceId(v string) *DescribeAuditRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetNodeId(v string) *DescribeAuditRecordsRequest {
	s.NodeId = &v
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

func (s *DescribeAuditRecordsRequest) SetDatabase(v string) *DescribeAuditRecordsRequest {
	s.Database = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetUser(v string) *DescribeAuditRecordsRequest {
	s.User = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetForm(v string) *DescribeAuditRecordsRequest {
	s.Form = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetQueryKeywords(v string) *DescribeAuditRecordsRequest {
	s.QueryKeywords = &v
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

func (s *DescribeAuditRecordsRequest) SetOrderType(v string) *DescribeAuditRecordsRequest {
	s.OrderType = &v
	return s
}

type DescribeAuditRecordsResponseBody struct {
	TotalRecordCount *int32                                 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	PageRecordCount  *int32                                 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber       *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Items            *DescribeAuditRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeAuditRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditRecordsResponseBody) SetTotalRecordCount(v int32) *DescribeAuditRecordsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeAuditRecordsResponseBody) SetPageRecordCount(v int32) *DescribeAuditRecordsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeAuditRecordsResponseBody) SetRequestId(v string) *DescribeAuditRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuditRecordsResponseBody) SetPageNumber(v int32) *DescribeAuditRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAuditRecordsResponseBody) SetItems(v *DescribeAuditRecordsResponseBodyItems) *DescribeAuditRecordsResponseBody {
	s.Items = v
	return s
}

type DescribeAuditRecordsResponseBodyItems struct {
	SQLRecord []*DescribeAuditRecordsResponseBodyItemsSQLRecord `json:"SQLRecord,omitempty" xml:"SQLRecord,omitempty" type:"Repeated"`
}

func (s DescribeAuditRecordsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAuditRecordsResponseBodyItems) SetSQLRecord(v []*DescribeAuditRecordsResponseBodyItemsSQLRecord) *DescribeAuditRecordsResponseBodyItems {
	s.SQLRecord = v
	return s
}

type DescribeAuditRecordsResponseBodyItemsSQLRecord struct {
	HostAddress         *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	TableName           *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	ReturnRowCounts     *int64  `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	DBName              *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	ExecuteTime         *string `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	ThreadID            *string `json:"ThreadID,omitempty" xml:"ThreadID,omitempty"`
	TotalExecutionTimes *int64  `json:"TotalExecutionTimes,omitempty" xml:"TotalExecutionTimes,omitempty"`
	Syntax              *string `json:"Syntax,omitempty" xml:"Syntax,omitempty"`
	AccountName         *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s DescribeAuditRecordsResponseBodyItemsSQLRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditRecordsResponseBodyItemsSQLRecord) GoString() string {
	return s.String()
}

func (s *DescribeAuditRecordsResponseBodyItemsSQLRecord) SetHostAddress(v string) *DescribeAuditRecordsResponseBodyItemsSQLRecord {
	s.HostAddress = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQLRecord) SetTableName(v string) *DescribeAuditRecordsResponseBodyItemsSQLRecord {
	s.TableName = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQLRecord) SetReturnRowCounts(v int64) *DescribeAuditRecordsResponseBodyItemsSQLRecord {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQLRecord) SetDBName(v string) *DescribeAuditRecordsResponseBodyItemsSQLRecord {
	s.DBName = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQLRecord) SetExecuteTime(v string) *DescribeAuditRecordsResponseBodyItemsSQLRecord {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQLRecord) SetThreadID(v string) *DescribeAuditRecordsResponseBodyItemsSQLRecord {
	s.ThreadID = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQLRecord) SetTotalExecutionTimes(v int64) *DescribeAuditRecordsResponseBodyItemsSQLRecord {
	s.TotalExecutionTimes = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQLRecord) SetSyntax(v string) *DescribeAuditRecordsResponseBodyItemsSQLRecord {
	s.Syntax = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQLRecord) SetAccountName(v string) *DescribeAuditRecordsResponseBodyItemsSQLRecord {
	s.AccountName = &v
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

type DescribeAvailableResourceRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	InstanceChargeType   *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
}

func (s DescribeAvailableResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceRequest) SetSecurityToken(v string) *DescribeAvailableResourceRequest {
	s.SecurityToken = &v
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

func (s *DescribeAvailableResourceRequest) SetRegionId(v string) *DescribeAvailableResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetZoneId(v string) *DescribeAvailableResourceRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetInstanceChargeType(v string) *DescribeAvailableResourceRequest {
	s.InstanceChargeType = &v
	return s
}

type DescribeAvailableResourceResponseBody struct {
	RequestId        *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SupportedDBTypes *DescribeAvailableResourceResponseBodySupportedDBTypes `json:"SupportedDBTypes,omitempty" xml:"SupportedDBTypes,omitempty" type:"Struct"`
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

func (s *DescribeAvailableResourceResponseBody) SetSupportedDBTypes(v *DescribeAvailableResourceResponseBodySupportedDBTypes) *DescribeAvailableResourceResponseBody {
	s.SupportedDBTypes = v
	return s
}

type DescribeAvailableResourceResponseBodySupportedDBTypes struct {
	SupportedDBType []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBType `json:"SupportedDBType,omitempty" xml:"SupportedDBType,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypes) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypes) SetSupportedDBType(v []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBType) *DescribeAvailableResourceResponseBodySupportedDBTypes {
	s.SupportedDBType = v
	return s
}

type DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBType struct {
	AvailableZones *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZones `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Struct"`
	DbType         *string                                                                             `json:"DbType,omitempty" xml:"DbType,omitempty"`
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBType) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBType) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBType) SetAvailableZones(v *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZones) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBType {
	s.AvailableZones = v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBType) SetDbType(v string) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBType {
	s.DbType = &v
	return s
}

type DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZones struct {
	AvailableZone []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZone `json:"AvailableZone,omitempty" xml:"AvailableZone,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZones) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZones) SetAvailableZone(v []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZone) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZones {
	s.AvailableZone = v
	return s
}

type DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZone struct {
	SupportedEngineVersions *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersions `json:"SupportedEngineVersions,omitempty" xml:"SupportedEngineVersions,omitempty" type:"Struct"`
	ZoneId                  *string                                                                                                                 `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	RegionId                *string                                                                                                                 `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZone) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZone) SetSupportedEngineVersions(v *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersions) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZone {
	s.SupportedEngineVersions = v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZone) SetZoneId(v string) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZone) SetRegionId(v string) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZone {
	s.RegionId = &v
	return s
}

type DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersions struct {
	SupportedEngineVersion []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersion `json:"SupportedEngineVersion,omitempty" xml:"SupportedEngineVersion,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersions) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersions) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersions) SetSupportedEngineVersion(v []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersion) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersions {
	s.SupportedEngineVersion = v
	return s
}

type DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersion struct {
	SupportedEngines *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEngines `json:"SupportedEngines,omitempty" xml:"SupportedEngines,omitempty" type:"Struct"`
	Version          *string                                                                                                                                                       `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersion) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersion) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersion) SetSupportedEngines(v *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEngines) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersion {
	s.SupportedEngines = v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersion) SetVersion(v string) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersion {
	s.Version = &v
	return s
}

type DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEngines struct {
	SupportedEngine []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngine `json:"SupportedEngine,omitempty" xml:"SupportedEngine,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEngines) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEngines) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEngines) SetSupportedEngine(v []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngine) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEngines {
	s.SupportedEngine = v
	return s
}

type DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngine struct {
	SupportedNodeTypes *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypes `json:"SupportedNodeTypes,omitempty" xml:"SupportedNodeTypes,omitempty" type:"Struct"`
	Engine             *string                                                                                                                                                                                        `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngine) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngine) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngine) SetSupportedNodeTypes(v *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypes) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngine {
	s.SupportedNodeTypes = v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngine) SetEngine(v string) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngine {
	s.Engine = &v
	return s
}

type DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypes struct {
	SupportedNodeType []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeType `json:"SupportedNodeType,omitempty" xml:"SupportedNodeType,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypes) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypes) SetSupportedNodeType(v []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeType) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypes {
	s.SupportedNodeType = v
	return s
}

type DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeType struct {
	NodeType           *string                                                                                                                                                                                                                           `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	NetworkTypes       *string                                                                                                                                                                                                                           `json:"NetworkTypes,omitempty" xml:"NetworkTypes,omitempty"`
	AvailableResources *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResources `json:"AvailableResources,omitempty" xml:"AvailableResources,omitempty" type:"Struct"`
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeType) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeType) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeType) SetNodeType(v string) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeType {
	s.NodeType = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeType) SetNetworkTypes(v string) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeType {
	s.NetworkTypes = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeType) SetAvailableResources(v *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResources) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeType {
	s.AvailableResources = v
	return s
}

type DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResources struct {
	AvailableResource []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResource `json:"AvailableResource,omitempty" xml:"AvailableResource,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResources) SetAvailableResource(v []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResource) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResources {
	s.AvailableResource = v
	return s
}

type DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResource struct {
	InstanceClassRemark *string `json:"InstanceClassRemark,omitempty" xml:"InstanceClassRemark,omitempty"`
	InstanceClass       *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResource) SetInstanceClassRemark(v string) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResource {
	s.InstanceClassRemark = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResource) SetInstanceClass(v string) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResource {
	s.InstanceClass = &v
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

type DescribeAvailableTimeRangeRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeAvailableTimeRangeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableTimeRangeRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableTimeRangeRequest) SetSecurityToken(v string) *DescribeAvailableTimeRangeRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeAvailableTimeRangeRequest) SetOwnerId(v int64) *DescribeAvailableTimeRangeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAvailableTimeRangeRequest) SetResourceOwnerAccount(v string) *DescribeAvailableTimeRangeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAvailableTimeRangeRequest) SetResourceOwnerId(v int64) *DescribeAvailableTimeRangeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAvailableTimeRangeRequest) SetOwnerAccount(v string) *DescribeAvailableTimeRangeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAvailableTimeRangeRequest) SetInstanceId(v string) *DescribeAvailableTimeRangeRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAvailableTimeRangeRequest) SetNodeId(v string) *DescribeAvailableTimeRangeRequest {
	s.NodeId = &v
	return s
}

type DescribeAvailableTimeRangeResponseBody struct {
	TimeRange *DescribeAvailableTimeRangeResponseBodyTimeRange `json:"TimeRange,omitempty" xml:"TimeRange,omitempty" type:"Struct"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAvailableTimeRangeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableTimeRangeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableTimeRangeResponseBody) SetTimeRange(v *DescribeAvailableTimeRangeResponseBodyTimeRange) *DescribeAvailableTimeRangeResponseBody {
	s.TimeRange = v
	return s
}

func (s *DescribeAvailableTimeRangeResponseBody) SetRequestId(v string) *DescribeAvailableTimeRangeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAvailableTimeRangeResponseBodyTimeRange struct {
	TimeRange []*DescribeAvailableTimeRangeResponseBodyTimeRangeTimeRange `json:"TimeRange,omitempty" xml:"TimeRange,omitempty" type:"Repeated"`
}

func (s DescribeAvailableTimeRangeResponseBodyTimeRange) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableTimeRangeResponseBodyTimeRange) GoString() string {
	return s.String()
}

func (s *DescribeAvailableTimeRangeResponseBodyTimeRange) SetTimeRange(v []*DescribeAvailableTimeRangeResponseBodyTimeRangeTimeRange) *DescribeAvailableTimeRangeResponseBodyTimeRange {
	s.TimeRange = v
	return s
}

type DescribeAvailableTimeRangeResponseBodyTimeRangeTimeRange struct {
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	NodeId    *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeAvailableTimeRangeResponseBodyTimeRangeTimeRange) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableTimeRangeResponseBodyTimeRangeTimeRange) GoString() string {
	return s.String()
}

func (s *DescribeAvailableTimeRangeResponseBodyTimeRangeTimeRange) SetStatus(v string) *DescribeAvailableTimeRangeResponseBodyTimeRangeTimeRange {
	s.Status = &v
	return s
}

func (s *DescribeAvailableTimeRangeResponseBodyTimeRangeTimeRange) SetEndTime(v string) *DescribeAvailableTimeRangeResponseBodyTimeRangeTimeRange {
	s.EndTime = &v
	return s
}

func (s *DescribeAvailableTimeRangeResponseBodyTimeRangeTimeRange) SetStartTime(v string) *DescribeAvailableTimeRangeResponseBodyTimeRangeTimeRange {
	s.StartTime = &v
	return s
}

func (s *DescribeAvailableTimeRangeResponseBodyTimeRangeTimeRange) SetTaskId(v string) *DescribeAvailableTimeRangeResponseBodyTimeRangeTimeRange {
	s.TaskId = &v
	return s
}

func (s *DescribeAvailableTimeRangeResponseBodyTimeRangeTimeRange) SetNodeId(v string) *DescribeAvailableTimeRangeResponseBodyTimeRangeTimeRange {
	s.NodeId = &v
	return s
}

type DescribeAvailableTimeRangeResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAvailableTimeRangeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAvailableTimeRangeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableTimeRangeResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableTimeRangeResponse) SetHeaders(v map[string]*string) *DescribeAvailableTimeRangeResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableTimeRangeResponse) SetBody(v *DescribeAvailableTimeRangeResponseBody) *DescribeAvailableTimeRangeResponse {
	s.Body = v
	return s
}

type DescribeBackupDBsRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SourceDBInstance     *string `json:"SourceDBInstance,omitempty" xml:"SourceDBInstance,omitempty"`
	RestoreTime          *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	BackupId             *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeBackupDBsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupDBsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupDBsRequest) SetSecurityToken(v string) *DescribeBackupDBsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeBackupDBsRequest) SetOwnerId(v int64) *DescribeBackupDBsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupDBsRequest) SetResourceOwnerAccount(v string) *DescribeBackupDBsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBackupDBsRequest) SetResourceOwnerId(v int64) *DescribeBackupDBsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBackupDBsRequest) SetOwnerAccount(v string) *DescribeBackupDBsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBackupDBsRequest) SetPageNumber(v int32) *DescribeBackupDBsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupDBsRequest) SetPageSize(v int32) *DescribeBackupDBsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupDBsRequest) SetSourceDBInstance(v string) *DescribeBackupDBsRequest {
	s.SourceDBInstance = &v
	return s
}

func (s *DescribeBackupDBsRequest) SetRestoreTime(v string) *DescribeBackupDBsRequest {
	s.RestoreTime = &v
	return s
}

func (s *DescribeBackupDBsRequest) SetBackupId(v string) *DescribeBackupDBsRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupDBsRequest) SetResourceGroupId(v string) *DescribeBackupDBsRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeBackupDBsResponseBody struct {
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Databases  *DescribeBackupDBsResponseBodyDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Struct"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeBackupDBsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupDBsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupDBsResponseBody) SetTotalCount(v int32) *DescribeBackupDBsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBackupDBsResponseBody) SetDatabases(v *DescribeBackupDBsResponseBodyDatabases) *DescribeBackupDBsResponseBody {
	s.Databases = v
	return s
}

func (s *DescribeBackupDBsResponseBody) SetRequestId(v string) *DescribeBackupDBsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupDBsResponseBody) SetPageSize(v int32) *DescribeBackupDBsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupDBsResponseBody) SetPageNumber(v int32) *DescribeBackupDBsResponseBody {
	s.PageNumber = &v
	return s
}

type DescribeBackupDBsResponseBodyDatabases struct {
	Database []*DescribeBackupDBsResponseBodyDatabasesDatabase `json:"Database,omitempty" xml:"Database,omitempty" type:"Repeated"`
}

func (s DescribeBackupDBsResponseBodyDatabases) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupDBsResponseBodyDatabases) GoString() string {
	return s.String()
}

func (s *DescribeBackupDBsResponseBodyDatabases) SetDatabase(v []*DescribeBackupDBsResponseBodyDatabasesDatabase) *DescribeBackupDBsResponseBodyDatabases {
	s.Database = v
	return s
}

type DescribeBackupDBsResponseBodyDatabasesDatabase struct {
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
}

func (s DescribeBackupDBsResponseBodyDatabasesDatabase) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupDBsResponseBodyDatabasesDatabase) GoString() string {
	return s.String()
}

func (s *DescribeBackupDBsResponseBodyDatabasesDatabase) SetDBName(v string) *DescribeBackupDBsResponseBodyDatabasesDatabase {
	s.DBName = &v
	return s
}

type DescribeBackupDBsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupDBsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupDBsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupDBsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupDBsResponse) SetHeaders(v map[string]*string) *DescribeBackupDBsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupDBsResponse) SetBody(v *DescribeBackupDBsResponseBody) *DescribeBackupDBsResponse {
	s.Body = v
	return s
}

type DescribeBackupPolicyRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
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

func (s *DescribeBackupPolicyRequest) SetDBInstanceId(v string) *DescribeBackupPolicyRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeBackupPolicyResponseBody struct {
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PreferredBackupTime   *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	BackupRetentionPeriod *string `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
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

func (s *DescribeBackupPolicyResponseBody) SetBackupRetentionPeriod(v string) *DescribeBackupPolicyResponseBody {
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
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	BackupId             *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
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

func (s *DescribeBackupsRequest) SetDBInstanceId(v string) *DescribeBackupsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeBackupsRequest) SetNodeId(v string) *DescribeBackupsRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeBackupsRequest) SetBackupId(v string) *DescribeBackupsRequest {
	s.BackupId = &v
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

func (s *DescribeBackupsRequest) SetStartTime(v string) *DescribeBackupsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupsRequest) SetEndTime(v string) *DescribeBackupsRequest {
	s.EndTime = &v
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
	BackupIntranetDownloadURL *string `json:"BackupIntranetDownloadURL,omitempty" xml:"BackupIntranetDownloadURL,omitempty"`
	BackupSize                *int64  `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	BackupDownloadURL         *string `json:"BackupDownloadURL,omitempty" xml:"BackupDownloadURL,omitempty"`
	BackupMode                *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	BackupEndTime             *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	BackupId                  *int32  `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	BackupDBNames             *string `json:"BackupDBNames,omitempty" xml:"BackupDBNames,omitempty"`
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

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupIntranetDownloadURL(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupIntranetDownloadURL = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupSize(v int64) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupSize = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupDownloadURL(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupDownloadURL = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupMode(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupEndTime(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupId(v int32) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupDBNames(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupDBNames = &v
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

type DescribeDBInstanceAttributeRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	Engine               *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDBInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeRequest) SetSecurityToken(v string) *DescribeDBInstanceAttributeRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetOwnerId(v int64) *DescribeDBInstanceAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetResourceOwnerAccount(v string) *DescribeDBInstanceAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetOwnerAccount(v string) *DescribeDBInstanceAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetEngine(v string) *DescribeDBInstanceAttributeRequest {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetDBInstanceId(v string) *DescribeDBInstanceAttributeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetResourceGroupId(v string) *DescribeDBInstanceAttributeRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeDBInstanceAttributeResponseBody struct {
	RequestId   *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DBInstances *DescribeDBInstanceAttributeResponseBodyDBInstances `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Struct"`
}

func (s DescribeDBInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBody) SetRequestId(v string) *DescribeDBInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetDBInstances(v *DescribeDBInstanceAttributeResponseBodyDBInstances) *DescribeDBInstanceAttributeResponseBody {
	s.DBInstances = v
	return s
}

type DescribeDBInstanceAttributeResponseBodyDBInstances struct {
	DBInstance []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance `json:"DBInstance,omitempty" xml:"DBInstance,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstances) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstances) SetDBInstance(v []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) *DescribeDBInstanceAttributeResponseBodyDBInstances {
	s.DBInstance = v
	return s
}

type DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance struct {
	CreationTime                *string                                                                       `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ReplicaSets                 *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSets      `json:"ReplicaSets,omitempty" xml:"ReplicaSets,omitempty" type:"Struct"`
	ReplacateId                 *string                                                                       `json:"ReplacateId,omitempty" xml:"ReplacateId,omitempty"`
	ChargeType                  *string                                                                       `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Tags                        *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTags             `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VpcAuthMode                 *string                                                                       `json:"VpcAuthMode,omitempty" xml:"VpcAuthMode,omitempty"`
	NetworkType                 *string                                                                       `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	LockMode                    *string                                                                       `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	EngineVersion               *string                                                                       `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	MaxIOPS                     *int32                                                                        `json:"MaxIOPS,omitempty" xml:"MaxIOPS,omitempty"`
	VPCCloudInstanceIds         *string                                                                       `json:"VPCCloudInstanceIds,omitempty" xml:"VPCCloudInstanceIds,omitempty"`
	MongosList                  *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosList       `json:"MongosList,omitempty" xml:"MongosList,omitempty" type:"Struct"`
	ProtocolType                *string                                                                       `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	DBInstanceDescription       *string                                                                       `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	CurrentKernelVersion        *string                                                                       `json:"CurrentKernelVersion,omitempty" xml:"CurrentKernelVersion,omitempty"`
	DBInstanceReleaseProtection *bool                                                                         `json:"DBInstanceReleaseProtection,omitempty" xml:"DBInstanceReleaseProtection,omitempty"`
	ExpireTime                  *string                                                                       `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	MaintainStartTime           *string                                                                       `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	DBInstanceType              *string                                                                       `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	LastDowngradeTime           *string                                                                       `json:"LastDowngradeTime,omitempty" xml:"LastDowngradeTime,omitempty"`
	ShardList                   *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardList        `json:"ShardList,omitempty" xml:"ShardList,omitempty" type:"Struct"`
	MaintainEndTime             *string                                                                       `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	DBInstanceStatus            *string                                                                       `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	VPCId                       *string                                                                       `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	RegionId                    *string                                                                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceStorage           *int32                                                                        `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	ReplicaSetName              *string                                                                       `json:"ReplicaSetName,omitempty" xml:"ReplicaSetName,omitempty"`
	VSwitchId                   *string                                                                       `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	StorageEngine               *string                                                                       `json:"StorageEngine,omitempty" xml:"StorageEngine,omitempty"`
	ConfigserverList            *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverList `json:"ConfigserverList,omitempty" xml:"ConfigserverList,omitempty" type:"Struct"`
	ResourceGroupId             *string                                                                       `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ZoneId                      *string                                                                       `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	MaxConnections              *int32                                                                        `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	DBInstanceId                *string                                                                       `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceClass             *string                                                                       `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	Engine                      *string                                                                       `json:"Engine,omitempty" xml:"Engine,omitempty"`
	ReadonlyReplicas            *string                                                                       `json:"ReadonlyReplicas,omitempty" xml:"ReadonlyReplicas,omitempty"`
	ReplicationFactor           *string                                                                       `json:"ReplicationFactor,omitempty" xml:"ReplicationFactor,omitempty"`
	KindCode                    *string                                                                       `json:"KindCode,omitempty" xml:"KindCode,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetCreationTime(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.CreationTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetReplicaSets(v *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSets) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.ReplicaSets = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetReplacateId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.ReplacateId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetChargeType(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.ChargeType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetTags(v *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTags) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.Tags = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetVpcAuthMode(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.VpcAuthMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetNetworkType(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetLockMode(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetEngineVersion(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetMaxIOPS(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.MaxIOPS = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetVPCCloudInstanceIds(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.VPCCloudInstanceIds = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetMongosList(v *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosList) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.MongosList = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetProtocolType(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.ProtocolType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetDBInstanceDescription(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetCurrentKernelVersion(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.CurrentKernelVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetDBInstanceReleaseProtection(v bool) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.DBInstanceReleaseProtection = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetExpireTime(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetMaintainStartTime(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetDBInstanceType(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetLastDowngradeTime(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.LastDowngradeTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetShardList(v *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardList) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.ShardList = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetMaintainEndTime(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetDBInstanceStatus(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetVPCId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.VPCId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetRegionId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetDBInstanceStorage(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.DBInstanceStorage = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetReplicaSetName(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.ReplicaSetName = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetVSwitchId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetStorageEngine(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.StorageEngine = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetConfigserverList(v *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverList) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.ConfigserverList = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetResourceGroupId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetZoneId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetMaxConnections(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.MaxConnections = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetDBInstanceId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetDBInstanceClass(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.DBInstanceClass = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetEngine(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetReadonlyReplicas(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.ReadonlyReplicas = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetReplicationFactor(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.ReplicationFactor = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetKindCode(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.KindCode = &v
	return s
}

type DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSets struct {
	ReplicaSet []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet `json:"ReplicaSet,omitempty" xml:"ReplicaSet,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSets) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSets) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSets) SetReplicaSet(v []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSets {
	s.ReplicaSet = v
	return s
}

type DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet struct {
	VSwitchId          *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ConnectionPort     *string `json:"ConnectionPort,omitempty" xml:"ConnectionPort,omitempty"`
	ReplicaSetRole     *string `json:"ReplicaSetRole,omitempty" xml:"ReplicaSetRole,omitempty"`
	ConnectionDomain   *string `json:"ConnectionDomain,omitempty" xml:"ConnectionDomain,omitempty"`
	VPCCloudInstanceId *string `json:"VPCCloudInstanceId,omitempty" xml:"VPCCloudInstanceId,omitempty"`
	NetworkType        *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	VPCId              *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet) SetVSwitchId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet) SetConnectionPort(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet {
	s.ConnectionPort = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet) SetReplicaSetRole(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet {
	s.ReplicaSetRole = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet) SetConnectionDomain(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet {
	s.ConnectionDomain = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet) SetVPCCloudInstanceId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet {
	s.VPCCloudInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet) SetNetworkType(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet {
	s.NetworkType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet) SetVPCId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet {
	s.VPCId = &v
	return s
}

type DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTags struct {
	Tag []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTags) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTags) SetTag(v []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTagsTag) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTags {
	s.Tag = v
	return s
}

type DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTagsTag) SetKey(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTagsTag) SetValue(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTagsTag {
	s.Value = &v
	return s
}

type DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosList struct {
	MongosAttribute []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute `json:"MongosAttribute,omitempty" xml:"MongosAttribute,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosList) SetMongosAttribute(v []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosList {
	s.MongosAttribute = v
	return s
}

type DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute struct {
	VpcCloudInstanceId *string `json:"VpcCloudInstanceId,omitempty" xml:"VpcCloudInstanceId,omitempty"`
	MaxIOPS            *int32  `json:"MaxIOPS,omitempty" xml:"MaxIOPS,omitempty"`
	VSwitchId          *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	NodeClass          *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	MaxConnections     *int32  `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	Port               *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	VPCId              *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	ConnectSting       *string `json:"ConnectSting,omitempty" xml:"ConnectSting,omitempty"`
	NodeDescription    *string `json:"NodeDescription,omitempty" xml:"NodeDescription,omitempty"`
	NodeId             *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetVpcCloudInstanceId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.VpcCloudInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetMaxIOPS(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.MaxIOPS = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetVSwitchId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetNodeClass(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.NodeClass = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetMaxConnections(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.MaxConnections = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetPort(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetVPCId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.VPCId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetConnectSting(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.ConnectSting = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetNodeDescription(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.NodeDescription = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetNodeId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.NodeId = &v
	return s
}

type DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardList struct {
	ShardAttribute []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute `json:"ShardAttribute,omitempty" xml:"ShardAttribute,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardList) SetShardAttribute(v []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardList {
	s.ShardAttribute = v
	return s
}

type DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute struct {
	MaxIOPS          *int32  `json:"MaxIOPS,omitempty" xml:"MaxIOPS,omitempty"`
	ConnectString    *string `json:"ConnectString,omitempty" xml:"ConnectString,omitempty"`
	NodeClass        *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	MaxConnections   *int32  `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	Port             *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	NodeDescription  *string `json:"NodeDescription,omitempty" xml:"NodeDescription,omitempty"`
	NodeId           *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeStorage      *int32  `json:"NodeStorage,omitempty" xml:"NodeStorage,omitempty"`
	ReadonlyReplicas *int32  `json:"ReadonlyReplicas,omitempty" xml:"ReadonlyReplicas,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetMaxIOPS(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.MaxIOPS = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetConnectString(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.ConnectString = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetNodeClass(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.NodeClass = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetMaxConnections(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.MaxConnections = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetPort(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetNodeDescription(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.NodeDescription = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetNodeId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.NodeId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetNodeStorage(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.NodeStorage = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetReadonlyReplicas(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.ReadonlyReplicas = &v
	return s
}

type DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverList struct {
	ConfigserverAttribute []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute `json:"ConfigserverAttribute,omitempty" xml:"ConfigserverAttribute,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverList) SetConfigserverAttribute(v []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverList {
	s.ConfigserverAttribute = v
	return s
}

type DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute struct {
	MaxIOPS         *int32  `json:"MaxIOPS,omitempty" xml:"MaxIOPS,omitempty"`
	ConnectString   *string `json:"ConnectString,omitempty" xml:"ConnectString,omitempty"`
	NodeClass       *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	MaxConnections  *int32  `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	Port            *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	NodeDescription *string `json:"NodeDescription,omitempty" xml:"NodeDescription,omitempty"`
	NodeId          *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeStorage     *int32  `json:"NodeStorage,omitempty" xml:"NodeStorage,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) SetMaxIOPS(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute {
	s.MaxIOPS = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) SetConnectString(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute {
	s.ConnectString = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) SetNodeClass(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute {
	s.NodeClass = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) SetMaxConnections(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute {
	s.MaxConnections = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) SetPort(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) SetNodeDescription(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute {
	s.NodeDescription = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) SetNodeId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute {
	s.NodeId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) SetNodeStorage(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute {
	s.NodeStorage = &v
	return s
}

type DescribeDBInstanceAttributeResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceAttributeResponse) SetBody(v *DescribeDBInstanceAttributeResponseBody) *DescribeDBInstanceAttributeResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceEncryptionKeyRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	EncryptionKey        *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
}

func (s DescribeDBInstanceEncryptionKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceEncryptionKeyRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceEncryptionKeyRequest) SetSecurityToken(v string) *DescribeDBInstanceEncryptionKeyRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyRequest) SetOwnerId(v int64) *DescribeDBInstanceEncryptionKeyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyRequest) SetResourceOwnerAccount(v string) *DescribeDBInstanceEncryptionKeyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceEncryptionKeyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyRequest) SetOwnerAccount(v string) *DescribeDBInstanceEncryptionKeyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyRequest) SetDBInstanceId(v string) *DescribeDBInstanceEncryptionKeyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyRequest) SetEncryptionKey(v string) *DescribeDBInstanceEncryptionKeyRequest {
	s.EncryptionKey = &v
	return s
}

type DescribeDBInstanceEncryptionKeyResponseBody struct {
	Origin              *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EncryptionKeyStatus *string `json:"EncryptionKeyStatus,omitempty" xml:"EncryptionKeyStatus,omitempty"`
	MaterialExpireTime  *string `json:"MaterialExpireTime,omitempty" xml:"MaterialExpireTime,omitempty"`
	KeyUsage            *string `json:"KeyUsage,omitempty" xml:"KeyUsage,omitempty"`
	EncryptionKey       *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	Creator             *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	DeleteDate          *string `json:"DeleteDate,omitempty" xml:"DeleteDate,omitempty"`
}

func (s DescribeDBInstanceEncryptionKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceEncryptionKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceEncryptionKeyResponseBody) SetOrigin(v string) *DescribeDBInstanceEncryptionKeyResponseBody {
	s.Origin = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponseBody) SetDescription(v string) *DescribeDBInstanceEncryptionKeyResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponseBody) SetRequestId(v string) *DescribeDBInstanceEncryptionKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponseBody) SetEncryptionKeyStatus(v string) *DescribeDBInstanceEncryptionKeyResponseBody {
	s.EncryptionKeyStatus = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponseBody) SetMaterialExpireTime(v string) *DescribeDBInstanceEncryptionKeyResponseBody {
	s.MaterialExpireTime = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponseBody) SetKeyUsage(v string) *DescribeDBInstanceEncryptionKeyResponseBody {
	s.KeyUsage = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponseBody) SetEncryptionKey(v string) *DescribeDBInstanceEncryptionKeyResponseBody {
	s.EncryptionKey = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponseBody) SetCreator(v string) *DescribeDBInstanceEncryptionKeyResponseBody {
	s.Creator = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponseBody) SetDeleteDate(v string) *DescribeDBInstanceEncryptionKeyResponseBody {
	s.DeleteDate = &v
	return s
}

type DescribeDBInstanceEncryptionKeyResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBInstanceEncryptionKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceEncryptionKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceEncryptionKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceEncryptionKeyResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceEncryptionKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponse) SetBody(v *DescribeDBInstanceEncryptionKeyResponseBody) *DescribeDBInstanceEncryptionKeyResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceMonitorRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeDBInstanceMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceMonitorRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceMonitorRequest) SetSecurityToken(v string) *DescribeDBInstanceMonitorRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDBInstanceMonitorRequest) SetOwnerId(v int64) *DescribeDBInstanceMonitorRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceMonitorRequest) SetResourceOwnerAccount(v string) *DescribeDBInstanceMonitorRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceMonitorRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceMonitorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceMonitorRequest) SetOwnerAccount(v string) *DescribeDBInstanceMonitorRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceMonitorRequest) SetDBInstanceId(v string) *DescribeDBInstanceMonitorRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeDBInstanceMonitorResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
}

func (s DescribeDBInstanceMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceMonitorResponseBody) SetRequestId(v string) *DescribeDBInstanceMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceMonitorResponseBody) SetGranularity(v string) *DescribeDBInstanceMonitorResponseBody {
	s.Granularity = &v
	return s
}

type DescribeDBInstanceMonitorResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBInstanceMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceMonitorResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceMonitorResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceMonitorResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceMonitorResponse) SetBody(v *DescribeDBInstanceMonitorResponseBody) *DescribeDBInstanceMonitorResponse {
	s.Body = v
	return s
}

type DescribeDBInstancePerformanceRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Key                  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RoleId               *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	ReplicaSetRole       *string `json:"ReplicaSetRole,omitempty" xml:"ReplicaSetRole,omitempty"`
}

func (s DescribeDBInstancePerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancePerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceRequest) SetSecurityToken(v string) *DescribeDBInstancePerformanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetOwnerId(v int64) *DescribeDBInstancePerformanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetResourceOwnerAccount(v string) *DescribeDBInstancePerformanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetResourceOwnerId(v int64) *DescribeDBInstancePerformanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetOwnerAccount(v string) *DescribeDBInstancePerformanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetDBInstanceId(v string) *DescribeDBInstancePerformanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetNodeId(v string) *DescribeDBInstancePerformanceRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetKey(v string) *DescribeDBInstancePerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetStartTime(v string) *DescribeDBInstancePerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetEndTime(v string) *DescribeDBInstancePerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetRoleId(v string) *DescribeDBInstancePerformanceRequest {
	s.RoleId = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetReplicaSetRole(v string) *DescribeDBInstancePerformanceRequest {
	s.ReplicaSetRole = &v
	return s
}

type DescribeDBInstancePerformanceResponseBody struct {
	PerformanceKeys *DescribeDBInstancePerformanceResponseBodyPerformanceKeys `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Struct"`
	EndTime         *string                                                   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId       *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime       *string                                                   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBInstancePerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancePerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceResponseBody) SetPerformanceKeys(v *DescribeDBInstancePerformanceResponseBodyPerformanceKeys) *DescribeDBInstancePerformanceResponseBody {
	s.PerformanceKeys = v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetEndTime(v string) *DescribeDBInstancePerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetRequestId(v string) *DescribeDBInstancePerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetStartTime(v string) *DescribeDBInstancePerformanceResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDBInstancePerformanceResponseBodyPerformanceKeys struct {
	PerformanceKey []*DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey `json:"PerformanceKey,omitempty" xml:"PerformanceKey,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancePerformanceResponseBodyPerformanceKeys) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancePerformanceResponseBodyPerformanceKeys) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeys) SetPerformanceKey(v []*DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey) *DescribeDBInstancePerformanceResponseBodyPerformanceKeys {
	s.PerformanceKey = v
	return s
}

type DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey struct {
	Key               *string                                                                                  `json:"Key,omitempty" xml:"Key,omitempty"`
	Unit              *string                                                                                  `json:"Unit,omitempty" xml:"Unit,omitempty"`
	ValueFormat       *string                                                                                  `json:"ValueFormat,omitempty" xml:"ValueFormat,omitempty"`
	PerformanceValues *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues `json:"PerformanceValues,omitempty" xml:"PerformanceValues,omitempty" type:"Struct"`
}

func (s DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey) SetKey(v string) *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey) SetUnit(v string) *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey {
	s.Unit = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey) SetValueFormat(v string) *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey {
	s.ValueFormat = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey) SetPerformanceValues(v *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues) *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey {
	s.PerformanceValues = v
	return s
}

type DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues struct {
	PerformanceValue []*DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue `json:"PerformanceValue,omitempty" xml:"PerformanceValue,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues) SetPerformanceValue(v []*DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue) *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues {
	s.PerformanceValue = v
	return s
}

type DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Date  *string `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue) SetValue(v string) *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue {
	s.Value = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue) SetDate(v string) *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue {
	s.Date = &v
	return s
}

type DescribeDBInstancePerformanceResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBInstancePerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstancePerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancePerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceResponse) SetHeaders(v map[string]*string) *DescribeDBInstancePerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstancePerformanceResponse) SetBody(v *DescribeDBInstancePerformanceResponseBody) *DescribeDBInstancePerformanceResponse {
	s.Body = v
	return s
}

type DescribeDBInstancesRequest struct {
	SecurityToken         *string                          `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId               *int64                           `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string                          `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64                           `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount          *string                          `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	PageNumber            *int32                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int32                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	DBInstanceId          *string                          `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ReplicationFactor     *string                          `json:"ReplicationFactor,omitempty" xml:"ReplicationFactor,omitempty"`
	DBInstanceDescription *string                          `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	ExpireTime            *string                          `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	DBInstanceStatus      *string                          `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	DBInstanceType        *string                          `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	DBInstanceClass       *string                          `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	Engine                *string                          `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion         *string                          `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	NetworkType           *string                          `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	VpcId                 *string                          `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VSwitchId             *string                          `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ChargeType            *string                          `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ZoneId                *string                          `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Expired               *string                          `json:"Expired,omitempty" xml:"Expired,omitempty"`
	ConnectionDomain      *string                          `json:"ConnectionDomain,omitempty" xml:"ConnectionDomain,omitempty"`
	ResourceGroupId       *string                          `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag                   []*DescribeDBInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesRequest) SetSecurityToken(v string) *DescribeDBInstancesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetOwnerId(v int64) *DescribeDBInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetResourceOwnerAccount(v string) *DescribeDBInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetResourceOwnerId(v int64) *DescribeDBInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetOwnerAccount(v string) *DescribeDBInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetPageNumber(v int32) *DescribeDBInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetPageSize(v int32) *DescribeDBInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceId(v string) *DescribeDBInstancesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetReplicationFactor(v string) *DescribeDBInstancesRequest {
	s.ReplicationFactor = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceDescription(v string) *DescribeDBInstancesRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetExpireTime(v string) *DescribeDBInstancesRequest {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceStatus(v string) *DescribeDBInstancesRequest {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceType(v string) *DescribeDBInstancesRequest {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceClass(v string) *DescribeDBInstancesRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetEngine(v string) *DescribeDBInstancesRequest {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetEngineVersion(v string) *DescribeDBInstancesRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetNetworkType(v string) *DescribeDBInstancesRequest {
	s.NetworkType = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetVpcId(v string) *DescribeDBInstancesRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetVSwitchId(v string) *DescribeDBInstancesRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetChargeType(v string) *DescribeDBInstancesRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetZoneId(v string) *DescribeDBInstancesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetExpired(v string) *DescribeDBInstancesRequest {
	s.Expired = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetConnectionDomain(v string) *DescribeDBInstancesRequest {
	s.ConnectionDomain = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetResourceGroupId(v string) *DescribeDBInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetTag(v []*DescribeDBInstancesRequestTag) *DescribeDBInstancesRequest {
	s.Tag = v
	return s
}

type DescribeDBInstancesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesRequestTag) SetKey(v string) *DescribeDBInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancesRequestTag) SetValue(v string) *DescribeDBInstancesRequestTag {
	s.Value = &v
	return s
}

type DescribeDBInstancesResponseBody struct {
	TotalCount  *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber  *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	DBInstances *DescribeDBInstancesResponseBodyDBInstances `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Struct"`
}

func (s DescribeDBInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBody) SetTotalCount(v int32) *DescribeDBInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetRequestId(v string) *DescribeDBInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetPageSize(v int32) *DescribeDBInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetPageNumber(v int32) *DescribeDBInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetDBInstances(v *DescribeDBInstancesResponseBodyDBInstances) *DescribeDBInstancesResponseBody {
	s.DBInstances = v
	return s
}

type DescribeDBInstancesResponseBodyDBInstances struct {
	DBInstance []*DescribeDBInstancesResponseBodyDBInstancesDBInstance `json:"DBInstance,omitempty" xml:"DBInstance,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancesResponseBodyDBInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyDBInstances) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetDBInstance(v []*DescribeDBInstancesResponseBodyDBInstancesDBInstance) *DescribeDBInstancesResponseBodyDBInstances {
	s.DBInstance = v
	return s
}

type DescribeDBInstancesResponseBodyDBInstancesDBInstance struct {
	CreationTime          *string                                                         `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ChargeType            *string                                                         `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Tags                  *DescribeDBInstancesResponseBodyDBInstancesDBInstanceTags       `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VpcAuthMode           *string                                                         `json:"VpcAuthMode,omitempty" xml:"VpcAuthMode,omitempty"`
	NetworkType           *string                                                         `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	LockMode              *string                                                         `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	EngineVersion         *string                                                         `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	MongosList            *DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosList `json:"MongosList,omitempty" xml:"MongosList,omitempty" type:"Struct"`
	DBInstanceDescription *string                                                         `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	ExpireTime            *string                                                         `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	DBInstanceType        *string                                                         `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	LastDowngradeTime     *string                                                         `json:"LastDowngradeTime,omitempty" xml:"LastDowngradeTime,omitempty"`
	ShardList             *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardList  `json:"ShardList,omitempty" xml:"ShardList,omitempty" type:"Struct"`
	DestroyTime           *string                                                         `json:"DestroyTime,omitempty" xml:"DestroyTime,omitempty"`
	DBInstanceStatus      *string                                                         `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	RegionId              *string                                                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceStorage     *int32                                                          `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	ResourceGroupId       *string                                                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ZoneId                *string                                                         `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	DBInstanceId          *string                                                         `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceClass       *string                                                         `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	Engine                *string                                                         `json:"Engine,omitempty" xml:"Engine,omitempty"`
	ReplicationFactor     *string                                                         `json:"ReplicationFactor,omitempty" xml:"ReplicationFactor,omitempty"`
	KindCode              *string                                                         `json:"KindCode,omitempty" xml:"KindCode,omitempty"`
}

func (s DescribeDBInstancesResponseBodyDBInstancesDBInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyDBInstancesDBInstance) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetCreationTime(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.CreationTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetChargeType(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.ChargeType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetTags(v *DescribeDBInstancesResponseBodyDBInstancesDBInstanceTags) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.Tags = v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetVpcAuthMode(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.VpcAuthMode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetNetworkType(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetLockMode(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetEngineVersion(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetMongosList(v *DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosList) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.MongosList = v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetDBInstanceDescription(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetExpireTime(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetDBInstanceType(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetLastDowngradeTime(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.LastDowngradeTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetShardList(v *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardList) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.ShardList = v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetDestroyTime(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.DestroyTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetDBInstanceStatus(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetRegionId(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetDBInstanceStorage(v int32) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.DBInstanceStorage = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetResourceGroupId(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetZoneId(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetDBInstanceId(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetDBInstanceClass(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.DBInstanceClass = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetEngine(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetReplicationFactor(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.ReplicationFactor = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstance) SetKindCode(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstance {
	s.KindCode = &v
	return s
}

type DescribeDBInstancesResponseBodyDBInstancesDBInstanceTags struct {
	Tag []*DescribeDBInstancesResponseBodyDBInstancesDBInstanceTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancesResponseBodyDBInstancesDBInstanceTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyDBInstancesDBInstanceTags) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceTags) SetTag(v []*DescribeDBInstancesResponseBodyDBInstancesDBInstanceTagsTag) *DescribeDBInstancesResponseBodyDBInstancesDBInstanceTags {
	s.Tag = v
	return s
}

type DescribeDBInstancesResponseBodyDBInstancesDBInstanceTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstancesResponseBodyDBInstancesDBInstanceTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyDBInstancesDBInstanceTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceTagsTag) SetKey(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstanceTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceTagsTag) SetValue(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstanceTagsTag {
	s.Value = &v
	return s
}

type DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosList struct {
	MongosAttribute []*DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosListMongosAttribute `json:"MongosAttribute,omitempty" xml:"MongosAttribute,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosList) SetMongosAttribute(v []*DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) *DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosList {
	s.MongosAttribute = v
	return s
}

type DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosListMongosAttribute struct {
	NodeClass       *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	NodeDescription *string `json:"NodeDescription,omitempty" xml:"NodeDescription,omitempty"`
	NodeId          *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetNodeClass(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.NodeClass = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetNodeDescription(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.NodeDescription = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetNodeId(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.NodeId = &v
	return s
}

type DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardList struct {
	ShardAttribute []*DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute `json:"ShardAttribute,omitempty" xml:"ShardAttribute,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardList) SetShardAttribute(v []*DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute) *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardList {
	s.ShardAttribute = v
	return s
}

type DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute struct {
	NodeClass        *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	NodeDescription  *string `json:"NodeDescription,omitempty" xml:"NodeDescription,omitempty"`
	NodeStorage      *int32  `json:"NodeStorage,omitempty" xml:"NodeStorage,omitempty"`
	NodeId           *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	ReadonlyReplicas *int32  `json:"ReadonlyReplicas,omitempty" xml:"ReadonlyReplicas,omitempty"`
}

func (s DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetNodeClass(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.NodeClass = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetNodeDescription(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.NodeDescription = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetNodeStorage(v int32) *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.NodeStorage = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetNodeId(v string) *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.NodeId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetReadonlyReplicas(v int32) *DescribeDBInstancesResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.ReadonlyReplicas = &v
	return s
}

type DescribeDBInstancesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponse) SetHeaders(v map[string]*string) *DescribeDBInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstancesResponse) SetBody(v *DescribeDBInstancesResponseBody) *DescribeDBInstancesResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceSSLRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeDBInstanceSSLRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceSSLRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSSLRequest) SetSecurityToken(v string) *DescribeDBInstanceSSLRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDBInstanceSSLRequest) SetOwnerId(v int64) *DescribeDBInstanceSSLRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceSSLRequest) SetResourceOwnerAccount(v string) *DescribeDBInstanceSSLRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceSSLRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceSSLRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceSSLRequest) SetOwnerAccount(v string) *DescribeDBInstanceSSLRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceSSLRequest) SetDBInstanceId(v string) *DescribeDBInstanceSSLRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeDBInstanceSSLResponseBody struct {
	SSLExpiredTime *string `json:"SSLExpiredTime,omitempty" xml:"SSLExpiredTime,omitempty"`
	SSLStatus      *string `json:"SSLStatus,omitempty" xml:"SSLStatus,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CertCommonName *string `json:"CertCommonName,omitempty" xml:"CertCommonName,omitempty"`
}

func (s DescribeDBInstanceSSLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceSSLResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSSLResponseBody) SetSSLExpiredTime(v string) *DescribeDBInstanceSSLResponseBody {
	s.SSLExpiredTime = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetSSLStatus(v string) *DescribeDBInstanceSSLResponseBody {
	s.SSLStatus = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetRequestId(v string) *DescribeDBInstanceSSLResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetCertCommonName(v string) *DescribeDBInstanceSSLResponseBody {
	s.CertCommonName = &v
	return s
}

type DescribeDBInstanceSSLResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBInstanceSSLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceSSLResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceSSLResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSSLResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceSSLResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceSSLResponse) SetBody(v *DescribeDBInstanceSSLResponseBody) *DescribeDBInstanceSSLResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceTDEInfoRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeDBInstanceTDEInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceTDEInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTDEInfoRequest) SetSecurityToken(v string) *DescribeDBInstanceTDEInfoRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDBInstanceTDEInfoRequest) SetOwnerId(v int64) *DescribeDBInstanceTDEInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceTDEInfoRequest) SetResourceOwnerAccount(v string) *DescribeDBInstanceTDEInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceTDEInfoRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceTDEInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceTDEInfoRequest) SetOwnerAccount(v string) *DescribeDBInstanceTDEInfoRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceTDEInfoRequest) SetDBInstanceId(v string) *DescribeDBInstanceTDEInfoRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeDBInstanceTDEInfoResponseBody struct {
	TDEStatus *string `json:"TDEStatus,omitempty" xml:"TDEStatus,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceTDEInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceTDEInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTDEInfoResponseBody) SetTDEStatus(v string) *DescribeDBInstanceTDEInfoResponseBody {
	s.TDEStatus = &v
	return s
}

func (s *DescribeDBInstanceTDEInfoResponseBody) SetRequestId(v string) *DescribeDBInstanceTDEInfoResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBInstanceTDEInfoResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBInstanceTDEInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceTDEInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceTDEInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTDEInfoResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceTDEInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceTDEInfoResponse) SetBody(v *DescribeDBInstanceTDEInfoResponseBody) *DescribeDBInstanceTDEInfoResponse {
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
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceStatus       *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
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

func (s *DescribeDedicatedClusterInstanceListRequest) SetRegion(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.Region = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetZoneId(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetInstanceId(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetInstanceStatus(v string) *DescribeDedicatedClusterInstanceListRequest {
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
	Instances  *DescribeDedicatedClusterInstanceListResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	TotalCount *int32                                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeDedicatedClusterInstanceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedClusterInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedClusterInstanceListResponseBody) SetInstances(v *DescribeDedicatedClusterInstanceListResponseBodyInstances) *DescribeDedicatedClusterInstanceListResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBody) SetTotalCount(v int32) *DescribeDedicatedClusterInstanceListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBody) SetRequestId(v string) *DescribeDedicatedClusterInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBody) SetPageSize(v int32) *DescribeDedicatedClusterInstanceListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBody) SetPageNumber(v int32) *DescribeDedicatedClusterInstanceListResponseBody {
	s.PageNumber = &v
	return s
}

type DescribeDedicatedClusterInstanceListResponseBodyInstances struct {
	DbInstance []*DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance `json:"dbInstance,omitempty" xml:"dbInstance,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedClusterInstanceListResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedClusterInstanceListResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetDbInstance(v []*DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.DbInstance = v
	return s
}

type DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance struct {
	VpcId             *string                                                                              `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	CharacterType     *string                                                                              `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	VswitchId         *string                                                                              `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	MaintainStartTime *string                                                                              `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	InstanceClass     *string                                                                              `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	CreateTime        *string                                                                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	MaintainEndTime   *string                                                                              `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	StorageType       *string                                                                              `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	InstanceNodeList  *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstanceInstanceNodeList `json:"InstanceNodeList,omitempty" xml:"InstanceNodeList,omitempty" type:"Struct"`
	InstanceId        *string                                                                              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	EngineVersion     *string                                                                              `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	RegionId          *string                                                                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceName      *string                                                                              `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Region            *string                                                                              `json:"Region,omitempty" xml:"Region,omitempty"`
	ZoneId            *string                                                                              `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ClusterName       *string                                                                              `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	InstanceStatus    *string                                                                              `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	Engine            *string                                                                              `json:"Engine,omitempty" xml:"Engine,omitempty"`
	CustomId          *string                                                                              `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	ClusterId         *string                                                                              `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance) SetVpcId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance {
	s.VpcId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance) SetCharacterType(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance {
	s.CharacterType = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance) SetVswitchId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance {
	s.VswitchId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance) SetMaintainStartTime(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance) SetInstanceClass(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance {
	s.InstanceClass = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance) SetCreateTime(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance) SetMaintainEndTime(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance) SetStorageType(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance {
	s.StorageType = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance) SetInstanceNodeList(v *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstanceInstanceNodeList) *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance {
	s.InstanceNodeList = v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance) SetInstanceId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance) SetEngineVersion(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance) SetRegionId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance) SetInstanceName(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance) SetRegion(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance {
	s.Region = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance) SetZoneId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance) SetClusterName(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance {
	s.ClusterName = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance) SetInstanceStatus(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance) SetEngine(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance) SetCustomId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance {
	s.CustomId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance) SetClusterId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstance {
	s.ClusterId = &v
	return s
}

type DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstanceInstanceNodeList struct {
	InstanceNodes []*DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstanceInstanceNodeListInstanceNodes `json:"InstanceNodes,omitempty" xml:"InstanceNodes,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstanceInstanceNodeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstanceInstanceNodeList) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstanceInstanceNodeList) SetInstanceNodes(v []*DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstanceInstanceNodeListInstanceNodes) *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstanceInstanceNodeList {
	s.InstanceNodes = v
	return s
}

type DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstanceInstanceNodeListInstanceNodes struct {
	NodeIp            *string `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	DedicatedHostName *string `json:"DedicatedHostName,omitempty" xml:"DedicatedHostName,omitempty"`
	InsName           *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	NodeType          *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	ZoneId            *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Role              *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Port              *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	NodeId            *int32  `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstanceInstanceNodeListInstanceNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstanceInstanceNodeListInstanceNodes) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstanceInstanceNodeListInstanceNodes) SetNodeIp(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstanceInstanceNodeListInstanceNodes {
	s.NodeIp = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstanceInstanceNodeListInstanceNodes) SetDedicatedHostName(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstanceInstanceNodeListInstanceNodes {
	s.DedicatedHostName = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstanceInstanceNodeListInstanceNodes) SetInsName(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstanceInstanceNodeListInstanceNodes {
	s.InsName = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstanceInstanceNodeListInstanceNodes) SetNodeType(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstanceInstanceNodeListInstanceNodes {
	s.NodeType = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstanceInstanceNodeListInstanceNodes) SetZoneId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstanceInstanceNodeListInstanceNodes {
	s.ZoneId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstanceInstanceNodeListInstanceNodes) SetRole(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstanceInstanceNodeListInstanceNodes {
	s.Role = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstanceInstanceNodeListInstanceNodes) SetPort(v int32) *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstanceInstanceNodeListInstanceNodes {
	s.Port = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstanceInstanceNodeListInstanceNodes) SetNodeId(v int32) *DescribeDedicatedClusterInstanceListResponseBodyInstancesDbInstanceInstanceNodeListInstanceNodes {
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

type DescribeErrorLogRecordsRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	RoleType             *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeErrorLogRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeErrorLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeErrorLogRecordsRequest) SetSecurityToken(v string) *DescribeErrorLogRecordsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetOwnerId(v int64) *DescribeErrorLogRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetResourceOwnerAccount(v string) *DescribeErrorLogRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetResourceOwnerId(v int64) *DescribeErrorLogRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetOwnerAccount(v string) *DescribeErrorLogRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetDBInstanceId(v string) *DescribeErrorLogRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetNodeId(v string) *DescribeErrorLogRecordsRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetRoleType(v string) *DescribeErrorLogRecordsRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetStartTime(v string) *DescribeErrorLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetEndTime(v string) *DescribeErrorLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetDBName(v string) *DescribeErrorLogRecordsRequest {
	s.DBName = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetPageSize(v int32) *DescribeErrorLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetPageNumber(v int32) *DescribeErrorLogRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetResourceGroupId(v string) *DescribeErrorLogRecordsRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeErrorLogRecordsResponseBody struct {
	TotalRecordCount *int32                                    `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	PageRecordCount  *int32                                    `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber       *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Items            *DescribeErrorLogRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	Engine           *string                                   `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s DescribeErrorLogRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeErrorLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeErrorLogRecordsResponseBody) SetTotalRecordCount(v int32) *DescribeErrorLogRecordsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBody) SetPageRecordCount(v int32) *DescribeErrorLogRecordsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBody) SetRequestId(v string) *DescribeErrorLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBody) SetPageNumber(v int32) *DescribeErrorLogRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBody) SetItems(v *DescribeErrorLogRecordsResponseBodyItems) *DescribeErrorLogRecordsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeErrorLogRecordsResponseBody) SetEngine(v string) *DescribeErrorLogRecordsResponseBody {
	s.Engine = &v
	return s
}

type DescribeErrorLogRecordsResponseBodyItems struct {
	LogRecords []*DescribeErrorLogRecordsResponseBodyItemsLogRecords `json:"LogRecords,omitempty" xml:"LogRecords,omitempty" type:"Repeated"`
}

func (s DescribeErrorLogRecordsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeErrorLogRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeErrorLogRecordsResponseBodyItems) SetLogRecords(v []*DescribeErrorLogRecordsResponseBodyItemsLogRecords) *DescribeErrorLogRecordsResponseBodyItems {
	s.LogRecords = v
	return s
}

type DescribeErrorLogRecordsResponseBodyItemsLogRecords struct {
	ConnInfo   *string `json:"ConnInfo,omitempty" xml:"ConnInfo,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Category   *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Id         *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeErrorLogRecordsResponseBodyItemsLogRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeErrorLogRecordsResponseBodyItemsLogRecords) GoString() string {
	return s.String()
}

func (s *DescribeErrorLogRecordsResponseBodyItemsLogRecords) SetConnInfo(v string) *DescribeErrorLogRecordsResponseBodyItemsLogRecords {
	s.ConnInfo = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBodyItemsLogRecords) SetCreateTime(v string) *DescribeErrorLogRecordsResponseBodyItemsLogRecords {
	s.CreateTime = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBodyItemsLogRecords) SetCategory(v string) *DescribeErrorLogRecordsResponseBodyItemsLogRecords {
	s.Category = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBodyItemsLogRecords) SetContent(v string) *DescribeErrorLogRecordsResponseBodyItemsLogRecords {
	s.Content = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBodyItemsLogRecords) SetId(v int32) *DescribeErrorLogRecordsResponseBodyItemsLogRecords {
	s.Id = &v
	return s
}

type DescribeErrorLogRecordsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeErrorLogRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeErrorLogRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeErrorLogRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeErrorLogRecordsResponse) SetHeaders(v map[string]*string) *DescribeErrorLogRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeErrorLogRecordsResponse) SetBody(v *DescribeErrorLogRecordsResponseBody) *DescribeErrorLogRecordsResponse {
	s.Body = v
	return s
}

type DescribeInstanceAutoRenewalAttributeRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceType       *string `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeInstanceAutoRenewalAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAutoRenewalAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetSecurityToken(v string) *DescribeInstanceAutoRenewalAttributeRequest {
	s.SecurityToken = &v
	return s
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

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetOwnerAccount(v string) *DescribeInstanceAutoRenewalAttributeRequest {
	s.OwnerAccount = &v
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

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetDBInstanceType(v string) *DescribeInstanceAutoRenewalAttributeRequest {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetPageSize(v int64) *DescribeInstanceAutoRenewalAttributeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetPageNumber(v int64) *DescribeInstanceAutoRenewalAttributeRequest {
	s.PageNumber = &v
	return s
}

type DescribeInstanceAutoRenewalAttributeResponseBody struct {
	ItemsNumbers    *int32                                                 `json:"ItemsNumbers,omitempty" xml:"ItemsNumbers,omitempty"`
	PageRecordCount *int32                                                 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId       *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber      *int32                                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Items           *DescribeInstanceAutoRenewalAttributeResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeInstanceAutoRenewalAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAutoRenewalAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) SetItemsNumbers(v int32) *DescribeInstanceAutoRenewalAttributeResponseBody {
	s.ItemsNumbers = &v
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
	DBInstanceType *string `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	AutoRenew      *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	Duration       *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	DbInstanceId   *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem) SetDBInstanceType(v string) *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem) SetAutoRenew(v string) *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem {
	s.AutoRenew = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem) SetDuration(v string) *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem {
	s.Duration = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem) SetDbInstanceId(v string) *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem {
	s.DbInstanceId = &v
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

type DescribeKernelReleaseNotesRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	KernelVersion        *string `json:"KernelVersion,omitempty" xml:"KernelVersion,omitempty"`
}

func (s DescribeKernelReleaseNotesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeKernelReleaseNotesRequest) GoString() string {
	return s.String()
}

func (s *DescribeKernelReleaseNotesRequest) SetSecurityToken(v string) *DescribeKernelReleaseNotesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeKernelReleaseNotesRequest) SetOwnerId(v int64) *DescribeKernelReleaseNotesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeKernelReleaseNotesRequest) SetResourceOwnerAccount(v string) *DescribeKernelReleaseNotesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeKernelReleaseNotesRequest) SetResourceOwnerId(v int64) *DescribeKernelReleaseNotesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeKernelReleaseNotesRequest) SetOwnerAccount(v string) *DescribeKernelReleaseNotesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeKernelReleaseNotesRequest) SetKernelVersion(v string) *DescribeKernelReleaseNotesRequest {
	s.KernelVersion = &v
	return s
}

type DescribeKernelReleaseNotesResponseBody struct {
	RequestId    *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ReleaseNotes *DescribeKernelReleaseNotesResponseBodyReleaseNotes `json:"ReleaseNotes,omitempty" xml:"ReleaseNotes,omitempty" type:"Struct"`
}

func (s DescribeKernelReleaseNotesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeKernelReleaseNotesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKernelReleaseNotesResponseBody) SetRequestId(v string) *DescribeKernelReleaseNotesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKernelReleaseNotesResponseBody) SetReleaseNotes(v *DescribeKernelReleaseNotesResponseBodyReleaseNotes) *DescribeKernelReleaseNotesResponseBody {
	s.ReleaseNotes = v
	return s
}

type DescribeKernelReleaseNotesResponseBodyReleaseNotes struct {
	ReleaseNote []*DescribeKernelReleaseNotesResponseBodyReleaseNotesReleaseNote `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty" type:"Repeated"`
}

func (s DescribeKernelReleaseNotesResponseBodyReleaseNotes) String() string {
	return tea.Prettify(s)
}

func (s DescribeKernelReleaseNotesResponseBodyReleaseNotes) GoString() string {
	return s.String()
}

func (s *DescribeKernelReleaseNotesResponseBodyReleaseNotes) SetReleaseNote(v []*DescribeKernelReleaseNotesResponseBodyReleaseNotesReleaseNote) *DescribeKernelReleaseNotesResponseBodyReleaseNotes {
	s.ReleaseNote = v
	return s
}

type DescribeKernelReleaseNotesResponseBodyReleaseNotesReleaseNote struct {
	ReleaseNote   *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	KernelVersion *string `json:"KernelVersion,omitempty" xml:"KernelVersion,omitempty"`
}

func (s DescribeKernelReleaseNotesResponseBodyReleaseNotesReleaseNote) String() string {
	return tea.Prettify(s)
}

func (s DescribeKernelReleaseNotesResponseBodyReleaseNotesReleaseNote) GoString() string {
	return s.String()
}

func (s *DescribeKernelReleaseNotesResponseBodyReleaseNotesReleaseNote) SetReleaseNote(v string) *DescribeKernelReleaseNotesResponseBodyReleaseNotesReleaseNote {
	s.ReleaseNote = &v
	return s
}

func (s *DescribeKernelReleaseNotesResponseBodyReleaseNotesReleaseNote) SetKernelVersion(v string) *DescribeKernelReleaseNotesResponseBodyReleaseNotesReleaseNote {
	s.KernelVersion = &v
	return s
}

type DescribeKernelReleaseNotesResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeKernelReleaseNotesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeKernelReleaseNotesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeKernelReleaseNotesResponse) GoString() string {
	return s.String()
}

func (s *DescribeKernelReleaseNotesResponse) SetHeaders(v map[string]*string) *DescribeKernelReleaseNotesResponse {
	s.Headers = v
	return s
}

func (s *DescribeKernelReleaseNotesResponse) SetBody(v *DescribeKernelReleaseNotesResponseBody) *DescribeKernelReleaseNotesResponse {
	s.Body = v
	return s
}

type DescribeMongoDBLogConfigRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeMongoDBLogConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMongoDBLogConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeMongoDBLogConfigRequest) SetSecurityToken(v string) *DescribeMongoDBLogConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeMongoDBLogConfigRequest) SetOwnerId(v int64) *DescribeMongoDBLogConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMongoDBLogConfigRequest) SetResourceOwnerAccount(v string) *DescribeMongoDBLogConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeMongoDBLogConfigRequest) SetResourceOwnerId(v int64) *DescribeMongoDBLogConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeMongoDBLogConfigRequest) SetOwnerAccount(v string) *DescribeMongoDBLogConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeMongoDBLogConfigRequest) SetDBInstanceId(v string) *DescribeMongoDBLogConfigRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeMongoDBLogConfigResponseBody struct {
	UserProjectName            *string `json:"UserProjectName,omitempty" xml:"UserProjectName,omitempty"`
	RequestId                  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	IsUserProjectLogstoreExist *int32  `json:"IsUserProjectLogstoreExist,omitempty" xml:"IsUserProjectLogstoreExist,omitempty"`
	IsEtlMetaExist             *int32  `json:"IsEtlMetaExist,omitempty" xml:"IsEtlMetaExist,omitempty"`
}

func (s DescribeMongoDBLogConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMongoDBLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMongoDBLogConfigResponseBody) SetUserProjectName(v string) *DescribeMongoDBLogConfigResponseBody {
	s.UserProjectName = &v
	return s
}

func (s *DescribeMongoDBLogConfigResponseBody) SetRequestId(v string) *DescribeMongoDBLogConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMongoDBLogConfigResponseBody) SetIsUserProjectLogstoreExist(v int32) *DescribeMongoDBLogConfigResponseBody {
	s.IsUserProjectLogstoreExist = &v
	return s
}

func (s *DescribeMongoDBLogConfigResponseBody) SetIsEtlMetaExist(v int32) *DescribeMongoDBLogConfigResponseBody {
	s.IsEtlMetaExist = &v
	return s
}

type DescribeMongoDBLogConfigResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMongoDBLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMongoDBLogConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMongoDBLogConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeMongoDBLogConfigResponse) SetHeaders(v map[string]*string) *DescribeMongoDBLogConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeMongoDBLogConfigResponse) SetBody(v *DescribeMongoDBLogConfigResponseBody) *DescribeMongoDBLogConfigResponse {
	s.Body = v
	return s
}

type DescribeParameterModificationHistoryRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	CharacterType        *string `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
}

func (s DescribeParameterModificationHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterModificationHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeParameterModificationHistoryRequest) SetSecurityToken(v string) *DescribeParameterModificationHistoryRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeParameterModificationHistoryRequest) SetOwnerId(v int64) *DescribeParameterModificationHistoryRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeParameterModificationHistoryRequest) SetResourceOwnerAccount(v string) *DescribeParameterModificationHistoryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeParameterModificationHistoryRequest) SetResourceOwnerId(v int64) *DescribeParameterModificationHistoryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeParameterModificationHistoryRequest) SetOwnerAccount(v string) *DescribeParameterModificationHistoryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeParameterModificationHistoryRequest) SetDBInstanceId(v string) *DescribeParameterModificationHistoryRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeParameterModificationHistoryRequest) SetNodeId(v string) *DescribeParameterModificationHistoryRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeParameterModificationHistoryRequest) SetStartTime(v string) *DescribeParameterModificationHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeParameterModificationHistoryRequest) SetEndTime(v string) *DescribeParameterModificationHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeParameterModificationHistoryRequest) SetCharacterType(v string) *DescribeParameterModificationHistoryRequest {
	s.CharacterType = &v
	return s
}

type DescribeParameterModificationHistoryResponseBody struct {
	RequestId            *string                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HistoricalParameters *DescribeParameterModificationHistoryResponseBodyHistoricalParameters `json:"HistoricalParameters,omitempty" xml:"HistoricalParameters,omitempty" type:"Struct"`
}

func (s DescribeParameterModificationHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterModificationHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParameterModificationHistoryResponseBody) SetRequestId(v string) *DescribeParameterModificationHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParameterModificationHistoryResponseBody) SetHistoricalParameters(v *DescribeParameterModificationHistoryResponseBodyHistoricalParameters) *DescribeParameterModificationHistoryResponseBody {
	s.HistoricalParameters = v
	return s
}

type DescribeParameterModificationHistoryResponseBodyHistoricalParameters struct {
	HistoricalParameter []*DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter `json:"HistoricalParameter,omitempty" xml:"HistoricalParameter,omitempty" type:"Repeated"`
}

func (s DescribeParameterModificationHistoryResponseBodyHistoricalParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterModificationHistoryResponseBodyHistoricalParameters) GoString() string {
	return s.String()
}

func (s *DescribeParameterModificationHistoryResponseBodyHistoricalParameters) SetHistoricalParameter(v []*DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter) *DescribeParameterModificationHistoryResponseBodyHistoricalParameters {
	s.HistoricalParameter = v
	return s
}

type DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter struct {
	ParameterName     *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	OldParameterValue *string `json:"OldParameterValue,omitempty" xml:"OldParameterValue,omitempty"`
	NewParameterValue *string `json:"NewParameterValue,omitempty" xml:"NewParameterValue,omitempty"`
	ModifyTime        *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
}

func (s DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter) GoString() string {
	return s.String()
}

func (s *DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter) SetParameterName(v string) *DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter {
	s.ParameterName = &v
	return s
}

func (s *DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter) SetOldParameterValue(v string) *DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter {
	s.OldParameterValue = &v
	return s
}

func (s *DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter) SetNewParameterValue(v string) *DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter {
	s.NewParameterValue = &v
	return s
}

func (s *DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter) SetModifyTime(v string) *DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter {
	s.ModifyTime = &v
	return s
}

type DescribeParameterModificationHistoryResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeParameterModificationHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeParameterModificationHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterModificationHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeParameterModificationHistoryResponse) SetHeaders(v map[string]*string) *DescribeParameterModificationHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeParameterModificationHistoryResponse) SetBody(v *DescribeParameterModificationHistoryResponseBody) *DescribeParameterModificationHistoryResponse {
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
	CharacterType        *string `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
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

func (s *DescribeParametersRequest) SetCharacterType(v string) *DescribeParametersRequest {
	s.CharacterType = &v
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
	Engine               *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion        *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
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

func (s *DescribeParameterTemplatesRequest) SetEngine(v string) *DescribeParameterTemplatesRequest {
	s.Engine = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetEngineVersion(v string) *DescribeParameterTemplatesRequest {
	s.EngineVersion = &v
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
	OrderType            *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	DBInstances          *string `json:"DBInstances,omitempty" xml:"DBInstances,omitempty"`
	CommodityCode        *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ProductCode          *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	BusinessInfo         *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	CouponNo             *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	OrderParamOut        *string `json:"OrderParamOut,omitempty" xml:"OrderParamOut,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *DescribePriceRequest) SetOrderType(v string) *DescribePriceRequest {
	s.OrderType = &v
	return s
}

func (s *DescribePriceRequest) SetDBInstances(v string) *DescribePriceRequest {
	s.DBInstances = &v
	return s
}

func (s *DescribePriceRequest) SetCommodityCode(v string) *DescribePriceRequest {
	s.CommodityCode = &v
	return s
}

func (s *DescribePriceRequest) SetProductCode(v string) *DescribePriceRequest {
	s.ProductCode = &v
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

func (s *DescribePriceRequest) SetOrderParamOut(v string) *DescribePriceRequest {
	s.OrderParamOut = &v
	return s
}

func (s *DescribePriceRequest) SetResourceGroupId(v string) *DescribePriceRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribePriceResponseBody struct {
	Order       *DescribePriceResponseBodyOrder     `json:"Order,omitempty" xml:"Order,omitempty" type:"Struct"`
	RequestId   *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubOrders   *DescribePriceResponseBodySubOrders `json:"SubOrders,omitempty" xml:"SubOrders,omitempty" type:"Struct"`
	TraceId     *string                             `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
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

func (s *DescribePriceResponseBody) SetTraceId(v string) *DescribePriceResponseBody {
	s.TraceId = &v
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
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsRequest) SetZoneId(v string) *DescribeRegionsRequest {
	s.ZoneId = &v
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
	DdsRegion []*DescribeRegionsResponseBodyRegionsDdsRegion `json:"DdsRegion,omitempty" xml:"DdsRegion,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetDdsRegion(v []*DescribeRegionsResponseBodyRegionsDdsRegion) *DescribeRegionsResponseBodyRegions {
	s.DdsRegion = v
	return s
}

type DescribeRegionsResponseBodyRegionsDdsRegion struct {
	Zones    *DescribeRegionsResponseBodyRegionsDdsRegionZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
	RegionId *string                                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsDdsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsDdsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsDdsRegion) SetZones(v *DescribeRegionsResponseBodyRegionsDdsRegionZones) *DescribeRegionsResponseBodyRegionsDdsRegion {
	s.Zones = v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsDdsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsDdsRegion {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponseBodyRegionsDdsRegionZones struct {
	Zone []*DescribeRegionsResponseBodyRegionsDdsRegionZonesZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegionsDdsRegionZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsDdsRegionZones) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsDdsRegionZones) SetZone(v []*DescribeRegionsResponseBodyRegionsDdsRegionZonesZone) *DescribeRegionsResponseBodyRegionsDdsRegionZones {
	s.Zone = v
	return s
}

type DescribeRegionsResponseBodyRegionsDdsRegionZonesZone struct {
	VpcEnabled *bool   `json:"VpcEnabled,omitempty" xml:"VpcEnabled,omitempty"`
	ZoneName   *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsDdsRegionZonesZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsDdsRegionZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsDdsRegionZonesZone) SetVpcEnabled(v bool) *DescribeRegionsResponseBodyRegionsDdsRegionZonesZone {
	s.VpcEnabled = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsDdsRegionZonesZone) SetZoneName(v string) *DescribeRegionsResponseBodyRegionsDdsRegionZonesZone {
	s.ZoneName = &v
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

type DescribeReplicaSetRoleRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeReplicaSetRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicaSetRoleRequest) GoString() string {
	return s.String()
}

func (s *DescribeReplicaSetRoleRequest) SetSecurityToken(v string) *DescribeReplicaSetRoleRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeReplicaSetRoleRequest) SetOwnerId(v int64) *DescribeReplicaSetRoleRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeReplicaSetRoleRequest) SetResourceOwnerAccount(v string) *DescribeReplicaSetRoleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeReplicaSetRoleRequest) SetResourceOwnerId(v int64) *DescribeReplicaSetRoleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeReplicaSetRoleRequest) SetOwnerAccount(v string) *DescribeReplicaSetRoleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeReplicaSetRoleRequest) SetDBInstanceId(v string) *DescribeReplicaSetRoleRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeReplicaSetRoleResponseBody struct {
	RequestId    *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DBInstanceId *string                                        `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ReplicaSets  *DescribeReplicaSetRoleResponseBodyReplicaSets `json:"ReplicaSets,omitempty" xml:"ReplicaSets,omitempty" type:"Struct"`
}

func (s DescribeReplicaSetRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicaSetRoleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeReplicaSetRoleResponseBody) SetRequestId(v string) *DescribeReplicaSetRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeReplicaSetRoleResponseBody) SetDBInstanceId(v string) *DescribeReplicaSetRoleResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeReplicaSetRoleResponseBody) SetReplicaSets(v *DescribeReplicaSetRoleResponseBodyReplicaSets) *DescribeReplicaSetRoleResponseBody {
	s.ReplicaSets = v
	return s
}

type DescribeReplicaSetRoleResponseBodyReplicaSets struct {
	ReplicaSet []*DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet `json:"ReplicaSet,omitempty" xml:"ReplicaSet,omitempty" type:"Repeated"`
}

func (s DescribeReplicaSetRoleResponseBodyReplicaSets) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicaSetRoleResponseBodyReplicaSets) GoString() string {
	return s.String()
}

func (s *DescribeReplicaSetRoleResponseBodyReplicaSets) SetReplicaSet(v []*DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet) *DescribeReplicaSetRoleResponseBodyReplicaSets {
	s.ReplicaSet = v
	return s
}

type DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet struct {
	ConnectionPort   *string `json:"ConnectionPort,omitempty" xml:"ConnectionPort,omitempty"`
	ReplicaSetRole   *string `json:"ReplicaSetRole,omitempty" xml:"ReplicaSetRole,omitempty"`
	ExpiredTime      *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	ConnectionDomain *string `json:"ConnectionDomain,omitempty" xml:"ConnectionDomain,omitempty"`
	NetworkType      *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	RoleId           *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet) GoString() string {
	return s.String()
}

func (s *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet) SetConnectionPort(v string) *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet {
	s.ConnectionPort = &v
	return s
}

func (s *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet) SetReplicaSetRole(v string) *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet {
	s.ReplicaSetRole = &v
	return s
}

func (s *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet) SetExpiredTime(v string) *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet) SetConnectionDomain(v string) *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet {
	s.ConnectionDomain = &v
	return s
}

func (s *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet) SetNetworkType(v string) *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet {
	s.NetworkType = &v
	return s
}

func (s *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet) SetRoleId(v string) *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet {
	s.RoleId = &v
	return s
}

type DescribeReplicaSetRoleResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeReplicaSetRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeReplicaSetRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicaSetRoleResponse) GoString() string {
	return s.String()
}

func (s *DescribeReplicaSetRoleResponse) SetHeaders(v map[string]*string) *DescribeReplicaSetRoleResponse {
	s.Headers = v
	return s
}

func (s *DescribeReplicaSetRoleResponse) SetBody(v *DescribeReplicaSetRoleResponseBody) *DescribeReplicaSetRoleResponse {
	s.Body = v
	return s
}

type DescribeRoleZoneInfoRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
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

func (s *DescribeRoleZoneInfoRequest) SetDBInstanceId(v string) *DescribeRoleZoneInfoRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeRoleZoneInfoResponseBody struct {
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ZoneInfos *DescribeRoleZoneInfoResponseBodyZoneInfos `json:"ZoneInfos,omitempty" xml:"ZoneInfos,omitempty" type:"Struct"`
}

func (s DescribeRoleZoneInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoleZoneInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRoleZoneInfoResponseBody) SetRequestId(v string) *DescribeRoleZoneInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBody) SetZoneInfos(v *DescribeRoleZoneInfoResponseBodyZoneInfos) *DescribeRoleZoneInfoResponseBody {
	s.ZoneInfos = v
	return s
}

type DescribeRoleZoneInfoResponseBodyZoneInfos struct {
	ZoneInfo []*DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo `json:"ZoneInfo,omitempty" xml:"ZoneInfo,omitempty" type:"Repeated"`
}

func (s DescribeRoleZoneInfoResponseBodyZoneInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoleZoneInfoResponseBodyZoneInfos) GoString() string {
	return s.String()
}

func (s *DescribeRoleZoneInfoResponseBodyZoneInfos) SetZoneInfo(v []*DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo) *DescribeRoleZoneInfoResponseBodyZoneInfos {
	s.ZoneInfo = v
	return s
}

type DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo struct {
	InsName  *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	ZoneId   *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	RoleId   *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo) GoString() string {
	return s.String()
}

func (s *DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo) SetInsName(v string) *DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo {
	s.InsName = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo) SetNodeType(v string) *DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo {
	s.NodeType = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo) SetRoleType(v string) *DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo {
	s.RoleType = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo) SetZoneId(v string) *DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo {
	s.ZoneId = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo) SetRoleId(v string) *DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo {
	s.RoleId = &v
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
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	RoleType             *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	OrderType            *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RoleId               *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
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

func (s *DescribeRunningLogRecordsRequest) SetDBInstanceId(v string) *DescribeRunningLogRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetNodeId(v string) *DescribeRunningLogRecordsRequest {
	s.NodeId = &v
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

func (s *DescribeRunningLogRecordsRequest) SetOrderType(v string) *DescribeRunningLogRecordsRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetResourceGroupId(v string) *DescribeRunningLogRecordsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetRoleId(v string) *DescribeRunningLogRecordsRequest {
	s.RoleId = &v
	return s
}

type DescribeRunningLogRecordsResponseBody struct {
	TotalRecordCount *int32                                      `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	PageRecordCount  *int32                                      `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber       *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
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

func (s *DescribeRunningLogRecordsResponseBody) SetPageNumber(v int32) *DescribeRunningLogRecordsResponseBody {
	s.PageNumber = &v
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
	ConnInfo   *string `json:"ConnInfo,omitempty" xml:"ConnInfo,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Category   *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Id         *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeRunningLogRecordsResponseBodyItemsLogRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeRunningLogRecordsResponseBodyItemsLogRecords) GoString() string {
	return s.String()
}

func (s *DescribeRunningLogRecordsResponseBodyItemsLogRecords) SetConnInfo(v string) *DescribeRunningLogRecordsResponseBodyItemsLogRecords {
	s.ConnInfo = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBodyItemsLogRecords) SetCreateTime(v string) *DescribeRunningLogRecordsResponseBodyItemsLogRecords {
	s.CreateTime = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBodyItemsLogRecords) SetCategory(v string) *DescribeRunningLogRecordsResponseBodyItemsLogRecords {
	s.Category = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBodyItemsLogRecords) SetContent(v string) *DescribeRunningLogRecordsResponseBodyItemsLogRecords {
	s.Content = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBodyItemsLogRecords) SetId(v int32) *DescribeRunningLogRecordsResponseBodyItemsLogRecords {
	s.Id = &v
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
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
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

func (s *DescribeSecurityGroupConfigurationRequest) SetDBInstanceId(v string) *DescribeSecurityGroupConfigurationRequest {
	s.DBInstanceId = &v
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
	RdsEcsSecurityGroupRel []*DescribeSecurityGroupConfigurationResponseBodyItemsRdsEcsSecurityGroupRel `json:"RdsEcsSecurityGroupRel,omitempty" xml:"RdsEcsSecurityGroupRel,omitempty" type:"Repeated"`
}

func (s DescribeSecurityGroupConfigurationResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupConfigurationResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItems) SetRdsEcsSecurityGroupRel(v []*DescribeSecurityGroupConfigurationResponseBodyItemsRdsEcsSecurityGroupRel) *DescribeSecurityGroupConfigurationResponseBodyItems {
	s.RdsEcsSecurityGroupRel = v
	return s
}

type DescribeSecurityGroupConfigurationResponseBodyItemsRdsEcsSecurityGroupRel struct {
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	NetType         *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSecurityGroupConfigurationResponseBodyItemsRdsEcsSecurityGroupRel) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupConfigurationResponseBodyItemsRdsEcsSecurityGroupRel) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItemsRdsEcsSecurityGroupRel) SetSecurityGroupId(v string) *DescribeSecurityGroupConfigurationResponseBodyItemsRdsEcsSecurityGroupRel {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItemsRdsEcsSecurityGroupRel) SetNetType(v string) *DescribeSecurityGroupConfigurationResponseBodyItemsRdsEcsSecurityGroupRel {
	s.NetType = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItemsRdsEcsSecurityGroupRel) SetRegionId(v string) *DescribeSecurityGroupConfigurationResponseBodyItemsRdsEcsSecurityGroupRel {
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
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
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

func (s *DescribeSecurityIpsRequest) SetDBInstanceId(v string) *DescribeSecurityIpsRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeSecurityIpsResponseBody struct {
	SecurityIps      *string                                          `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
	RequestId        *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityIpGroups *DescribeSecurityIpsResponseBodySecurityIpGroups `json:"SecurityIpGroups,omitempty" xml:"SecurityIpGroups,omitempty" type:"Struct"`
}

func (s DescribeSecurityIpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityIpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIpsResponseBody) SetSecurityIps(v string) *DescribeSecurityIpsResponseBody {
	s.SecurityIps = &v
	return s
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

type DescribeShardingNetworkAddressRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeShardingNetworkAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardingNetworkAddressRequest) GoString() string {
	return s.String()
}

func (s *DescribeShardingNetworkAddressRequest) SetSecurityToken(v string) *DescribeShardingNetworkAddressRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeShardingNetworkAddressRequest) SetOwnerId(v int64) *DescribeShardingNetworkAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeShardingNetworkAddressRequest) SetResourceOwnerAccount(v string) *DescribeShardingNetworkAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeShardingNetworkAddressRequest) SetResourceOwnerId(v int64) *DescribeShardingNetworkAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeShardingNetworkAddressRequest) SetOwnerAccount(v string) *DescribeShardingNetworkAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeShardingNetworkAddressRequest) SetDBInstanceId(v string) *DescribeShardingNetworkAddressRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeShardingNetworkAddressRequest) SetNodeId(v string) *DescribeShardingNetworkAddressRequest {
	s.NodeId = &v
	return s
}

type DescribeShardingNetworkAddressResponseBody struct {
	CompatibleConnections *DescribeShardingNetworkAddressResponseBodyCompatibleConnections `json:"CompatibleConnections,omitempty" xml:"CompatibleConnections,omitempty" type:"Struct"`
	RequestId             *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	NetworkAddresses      *DescribeShardingNetworkAddressResponseBodyNetworkAddresses      `json:"NetworkAddresses,omitempty" xml:"NetworkAddresses,omitempty" type:"Struct"`
}

func (s DescribeShardingNetworkAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardingNetworkAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeShardingNetworkAddressResponseBody) SetCompatibleConnections(v *DescribeShardingNetworkAddressResponseBodyCompatibleConnections) *DescribeShardingNetworkAddressResponseBody {
	s.CompatibleConnections = v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBody) SetRequestId(v string) *DescribeShardingNetworkAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBody) SetNetworkAddresses(v *DescribeShardingNetworkAddressResponseBodyNetworkAddresses) *DescribeShardingNetworkAddressResponseBody {
	s.NetworkAddresses = v
	return s
}

type DescribeShardingNetworkAddressResponseBodyCompatibleConnections struct {
	CompatibleConnection []*DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection `json:"CompatibleConnection,omitempty" xml:"CompatibleConnection,omitempty" type:"Repeated"`
}

func (s DescribeShardingNetworkAddressResponseBodyCompatibleConnections) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardingNetworkAddressResponseBodyCompatibleConnections) GoString() string {
	return s.String()
}

func (s *DescribeShardingNetworkAddressResponseBodyCompatibleConnections) SetCompatibleConnection(v []*DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection) *DescribeShardingNetworkAddressResponseBodyCompatibleConnections {
	s.CompatibleConnection = v
	return s
}

type DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection struct {
	VswitchId      *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	ExpiredTime    *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	NetworkType    *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Port           *string `json:"Port,omitempty" xml:"Port,omitempty"`
	NetworkAddress *string `json:"NetworkAddress,omitempty" xml:"NetworkAddress,omitempty"`
	VPCId          *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	IPAddress      *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
}

func (s DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection) GoString() string {
	return s.String()
}

func (s *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection) SetVswitchId(v string) *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection {
	s.VswitchId = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection) SetExpiredTime(v string) *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection) SetNetworkType(v string) *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection {
	s.NetworkType = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection) SetPort(v string) *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection {
	s.Port = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection) SetNetworkAddress(v string) *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection {
	s.NetworkAddress = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection) SetVPCId(v string) *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection {
	s.VPCId = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection) SetIPAddress(v string) *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection {
	s.IPAddress = &v
	return s
}

type DescribeShardingNetworkAddressResponseBodyNetworkAddresses struct {
	NetworkAddress []*DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress `json:"NetworkAddress,omitempty" xml:"NetworkAddress,omitempty" type:"Repeated"`
}

func (s DescribeShardingNetworkAddressResponseBodyNetworkAddresses) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardingNetworkAddressResponseBodyNetworkAddresses) GoString() string {
	return s.String()
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddresses) SetNetworkAddress(v []*DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) *DescribeShardingNetworkAddressResponseBodyNetworkAddresses {
	s.NetworkAddress = v
	return s
}

type DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress struct {
	NodeType       *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	VswitchId      *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	ExpiredTime    *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	NetworkType    *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Role           *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Port           *string `json:"Port,omitempty" xml:"Port,omitempty"`
	VPCId          *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	NetworkAddress *string `json:"NetworkAddress,omitempty" xml:"NetworkAddress,omitempty"`
	NodeId         *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	IPAddress      *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
}

func (s DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) GoString() string {
	return s.String()
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) SetNodeType(v string) *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress {
	s.NodeType = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) SetVswitchId(v string) *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress {
	s.VswitchId = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) SetExpiredTime(v string) *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) SetNetworkType(v string) *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress {
	s.NetworkType = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) SetRole(v string) *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress {
	s.Role = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) SetPort(v string) *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress {
	s.Port = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) SetVPCId(v string) *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress {
	s.VPCId = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) SetNetworkAddress(v string) *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress {
	s.NetworkAddress = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) SetNodeId(v string) *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress {
	s.NodeId = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) SetIPAddress(v string) *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress {
	s.IPAddress = &v
	return s
}

type DescribeShardingNetworkAddressResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeShardingNetworkAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeShardingNetworkAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeShardingNetworkAddressResponse) GoString() string {
	return s.String()
}

func (s *DescribeShardingNetworkAddressResponse) SetHeaders(v map[string]*string) *DescribeShardingNetworkAddressResponse {
	s.Headers = v
	return s
}

func (s *DescribeShardingNetworkAddressResponse) SetBody(v *DescribeShardingNetworkAddressResponseBody) *DescribeShardingNetworkAddressResponse {
	s.Body = v
	return s
}

type DescribeSlowLogRecordsRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	OrderType            *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *DescribeSlowLogRecordsRequest) SetDBInstanceId(v string) *DescribeSlowLogRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetNodeId(v string) *DescribeSlowLogRecordsRequest {
	s.NodeId = &v
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

func (s *DescribeSlowLogRecordsRequest) SetOrderType(v string) *DescribeSlowLogRecordsRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetResourceGroupId(v string) *DescribeSlowLogRecordsRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeSlowLogRecordsResponseBody struct {
	TotalRecordCount *int32                                   `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	PageRecordCount  *int32                                   `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber       *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
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

func (s *DescribeSlowLogRecordsResponseBody) SetPageNumber(v int32) *DescribeSlowLogRecordsResponseBody {
	s.PageNumber = &v
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
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" xml:"ExecutionStartTime,omitempty"`
	HostAddress        *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	QueryTimes         *string `json:"QueryTimes,omitempty" xml:"QueryTimes,omitempty"`
	TableName          *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	SQLText            *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	ReturnRowCounts    *int64  `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	KeysExamined       *int64  `json:"KeysExamined,omitempty" xml:"KeysExamined,omitempty"`
	DBName             *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	DocsExamined       *int64  `json:"DocsExamined,omitempty" xml:"DocsExamined,omitempty"`
	AccountName        *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBodyItemsLogRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyItemsLogRecords) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetExecutionStartTime(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.ExecutionStartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetHostAddress(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.HostAddress = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetQueryTimes(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.QueryTimes = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetTableName(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.TableName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetSQLText(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.SQLText = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetReturnRowCounts(v int64) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetKeysExamined(v int64) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.KeysExamined = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetDBName(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetDocsExamined(v int64) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.DocsExamined = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetAccountName(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.AccountName = &v
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

type DestroyInstanceRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DestroyInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DestroyInstanceRequest) GoString() string {
	return s.String()
}

func (s *DestroyInstanceRequest) SetSecurityToken(v string) *DestroyInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *DestroyInstanceRequest) SetOwnerId(v int64) *DestroyInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DestroyInstanceRequest) SetResourceOwnerAccount(v string) *DestroyInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DestroyInstanceRequest) SetResourceOwnerId(v int64) *DestroyInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DestroyInstanceRequest) SetOwnerAccount(v string) *DestroyInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DestroyInstanceRequest) SetInstanceId(v string) *DestroyInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DestroyInstanceRequest) SetDBInstanceId(v string) *DestroyInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DestroyInstanceRequest) SetClientToken(v string) *DestroyInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DestroyInstanceRequest) SetResourceGroupId(v string) *DestroyInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

type DestroyInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DestroyInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DestroyInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DestroyInstanceResponseBody) SetRequestId(v string) *DestroyInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DestroyInstanceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DestroyInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DestroyInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DestroyInstanceResponse) GoString() string {
	return s.String()
}

func (s *DestroyInstanceResponse) SetHeaders(v map[string]*string) *DestroyInstanceResponse {
	s.Headers = v
	return s
}

func (s *DestroyInstanceResponse) SetBody(v *DestroyInstanceResponseBody) *DestroyInstanceResponse {
	s.Body = v
	return s
}

type EvaluateResourceRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Engine               *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion        *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	DBInstanceClass      *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	ShardsInfo           *string `json:"ShardsInfo,omitempty" xml:"ShardsInfo,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s EvaluateResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s EvaluateResourceRequest) GoString() string {
	return s.String()
}

func (s *EvaluateResourceRequest) SetSecurityToken(v string) *EvaluateResourceRequest {
	s.SecurityToken = &v
	return s
}

func (s *EvaluateResourceRequest) SetOwnerId(v int64) *EvaluateResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *EvaluateResourceRequest) SetResourceOwnerAccount(v string) *EvaluateResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *EvaluateResourceRequest) SetResourceOwnerId(v int64) *EvaluateResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *EvaluateResourceRequest) SetOwnerAccount(v string) *EvaluateResourceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *EvaluateResourceRequest) SetRegionId(v string) *EvaluateResourceRequest {
	s.RegionId = &v
	return s
}

func (s *EvaluateResourceRequest) SetZoneId(v string) *EvaluateResourceRequest {
	s.ZoneId = &v
	return s
}

func (s *EvaluateResourceRequest) SetEngine(v string) *EvaluateResourceRequest {
	s.Engine = &v
	return s
}

func (s *EvaluateResourceRequest) SetEngineVersion(v string) *EvaluateResourceRequest {
	s.EngineVersion = &v
	return s
}

func (s *EvaluateResourceRequest) SetDBInstanceClass(v string) *EvaluateResourceRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *EvaluateResourceRequest) SetShardsInfo(v string) *EvaluateResourceRequest {
	s.ShardsInfo = &v
	return s
}

func (s *EvaluateResourceRequest) SetDBInstanceId(v string) *EvaluateResourceRequest {
	s.DBInstanceId = &v
	return s
}

type EvaluateResourceResponseBody struct {
	DBInstanceAvailable *string `json:"DBInstanceAvailable,omitempty" xml:"DBInstanceAvailable,omitempty"`
	EngineVersion       *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Engine              *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s EvaluateResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EvaluateResourceResponseBody) GoString() string {
	return s.String()
}

func (s *EvaluateResourceResponseBody) SetDBInstanceAvailable(v string) *EvaluateResourceResponseBody {
	s.DBInstanceAvailable = &v
	return s
}

func (s *EvaluateResourceResponseBody) SetEngineVersion(v string) *EvaluateResourceResponseBody {
	s.EngineVersion = &v
	return s
}

func (s *EvaluateResourceResponseBody) SetRequestId(v string) *EvaluateResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *EvaluateResourceResponseBody) SetEngine(v string) *EvaluateResourceResponseBody {
	s.Engine = &v
	return s
}

type EvaluateResourceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EvaluateResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EvaluateResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s EvaluateResourceResponse) GoString() string {
	return s.String()
}

func (s *EvaluateResourceResponse) SetHeaders(v map[string]*string) *EvaluateResourceResponse {
	s.Headers = v
	return s
}

func (s *EvaluateResourceResponse) SetBody(v *EvaluateResourceResponseBody) *EvaluateResourceResponse {
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

type MigrateAvailableZoneRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Vswitch              *string `json:"Vswitch,omitempty" xml:"Vswitch,omitempty"`
	EffectiveTime        *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
}

func (s MigrateAvailableZoneRequest) String() string {
	return tea.Prettify(s)
}

func (s MigrateAvailableZoneRequest) GoString() string {
	return s.String()
}

func (s *MigrateAvailableZoneRequest) SetOwnerId(v int64) *MigrateAvailableZoneRequest {
	s.OwnerId = &v
	return s
}

func (s *MigrateAvailableZoneRequest) SetResourceOwnerAccount(v string) *MigrateAvailableZoneRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *MigrateAvailableZoneRequest) SetResourceOwnerId(v int64) *MigrateAvailableZoneRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *MigrateAvailableZoneRequest) SetOwnerAccount(v string) *MigrateAvailableZoneRequest {
	s.OwnerAccount = &v
	return s
}

func (s *MigrateAvailableZoneRequest) SetDBInstanceId(v string) *MigrateAvailableZoneRequest {
	s.DBInstanceId = &v
	return s
}

func (s *MigrateAvailableZoneRequest) SetZoneId(v string) *MigrateAvailableZoneRequest {
	s.ZoneId = &v
	return s
}

func (s *MigrateAvailableZoneRequest) SetVswitch(v string) *MigrateAvailableZoneRequest {
	s.Vswitch = &v
	return s
}

func (s *MigrateAvailableZoneRequest) SetEffectiveTime(v string) *MigrateAvailableZoneRequest {
	s.EffectiveTime = &v
	return s
}

type MigrateAvailableZoneResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MigrateAvailableZoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MigrateAvailableZoneResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateAvailableZoneResponseBody) SetRequestId(v string) *MigrateAvailableZoneResponseBody {
	s.RequestId = &v
	return s
}

type MigrateAvailableZoneResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MigrateAvailableZoneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MigrateAvailableZoneResponse) String() string {
	return tea.Prettify(s)
}

func (s MigrateAvailableZoneResponse) GoString() string {
	return s.String()
}

func (s *MigrateAvailableZoneResponse) SetHeaders(v map[string]*string) *MigrateAvailableZoneResponse {
	s.Headers = v
	return s
}

func (s *MigrateAvailableZoneResponse) SetBody(v *MigrateAvailableZoneResponseBody) *MigrateAvailableZoneResponse {
	s.Body = v
	return s
}

type MigrateToOtherZoneRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	EffectiveTime        *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s MigrateToOtherZoneRequest) String() string {
	return tea.Prettify(s)
}

func (s MigrateToOtherZoneRequest) GoString() string {
	return s.String()
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

func (s *MigrateToOtherZoneRequest) SetInstanceId(v string) *MigrateToOtherZoneRequest {
	s.InstanceId = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetZoneId(v string) *MigrateToOtherZoneRequest {
	s.ZoneId = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetOwnerAccount(v string) *MigrateToOtherZoneRequest {
	s.OwnerAccount = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetEffectiveTime(v string) *MigrateToOtherZoneRequest {
	s.EffectiveTime = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetVSwitchId(v string) *MigrateToOtherZoneRequest {
	s.VSwitchId = &v
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
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
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

func (s *ModifyAccountDescriptionRequest) SetDBInstanceId(v string) *ModifyAccountDescriptionRequest {
	s.DBInstanceId = &v
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

type ModifyAuditLogFilterRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Filter               *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	RoleType             *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s ModifyAuditLogFilterRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAuditLogFilterRequest) GoString() string {
	return s.String()
}

func (s *ModifyAuditLogFilterRequest) SetSecurityToken(v string) *ModifyAuditLogFilterRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyAuditLogFilterRequest) SetOwnerId(v int64) *ModifyAuditLogFilterRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAuditLogFilterRequest) SetResourceOwnerAccount(v string) *ModifyAuditLogFilterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAuditLogFilterRequest) SetResourceOwnerId(v int64) *ModifyAuditLogFilterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAuditLogFilterRequest) SetOwnerAccount(v string) *ModifyAuditLogFilterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAuditLogFilterRequest) SetDBInstanceId(v string) *ModifyAuditLogFilterRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyAuditLogFilterRequest) SetFilter(v string) *ModifyAuditLogFilterRequest {
	s.Filter = &v
	return s
}

func (s *ModifyAuditLogFilterRequest) SetRoleType(v string) *ModifyAuditLogFilterRequest {
	s.RoleType = &v
	return s
}

type ModifyAuditLogFilterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAuditLogFilterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAuditLogFilterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAuditLogFilterResponseBody) SetRequestId(v string) *ModifyAuditLogFilterResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAuditLogFilterResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyAuditLogFilterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAuditLogFilterResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAuditLogFilterResponse) GoString() string {
	return s.String()
}

func (s *ModifyAuditLogFilterResponse) SetHeaders(v map[string]*string) *ModifyAuditLogFilterResponse {
	s.Headers = v
	return s
}

func (s *ModifyAuditLogFilterResponse) SetBody(v *ModifyAuditLogFilterResponseBody) *ModifyAuditLogFilterResponse {
	s.Body = v
	return s
}

type ModifyAuditPolicyRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	AuditStatus          *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	StoragePeriod        *int32  `json:"StoragePeriod,omitempty" xml:"StoragePeriod,omitempty"`
}

func (s ModifyAuditPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAuditPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyAuditPolicyRequest) SetSecurityToken(v string) *ModifyAuditPolicyRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyAuditPolicyRequest) SetOwnerId(v int64) *ModifyAuditPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAuditPolicyRequest) SetResourceOwnerAccount(v string) *ModifyAuditPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAuditPolicyRequest) SetResourceOwnerId(v int64) *ModifyAuditPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAuditPolicyRequest) SetOwnerAccount(v string) *ModifyAuditPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAuditPolicyRequest) SetDBInstanceId(v string) *ModifyAuditPolicyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyAuditPolicyRequest) SetAuditStatus(v string) *ModifyAuditPolicyRequest {
	s.AuditStatus = &v
	return s
}

func (s *ModifyAuditPolicyRequest) SetStoragePeriod(v int32) *ModifyAuditPolicyRequest {
	s.StoragePeriod = &v
	return s
}

type ModifyAuditPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAuditPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAuditPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAuditPolicyResponseBody) SetRequestId(v string) *ModifyAuditPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAuditPolicyResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyAuditPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAuditPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAuditPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyAuditPolicyResponse) SetHeaders(v map[string]*string) *ModifyAuditPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyAuditPolicyResponse) SetBody(v *ModifyAuditPolicyResponseBody) *ModifyAuditPolicyResponse {
	s.Body = v
	return s
}

type ModifyBackupPolicyRequest struct {
	SecurityToken         *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId          *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	PreferredBackupTime   *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
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

func (s *ModifyBackupPolicyRequest) SetDBInstanceId(v string) *ModifyBackupPolicyRequest {
	s.DBInstanceId = &v
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
	NodeId                  *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
	NewConnectionString     *string `json:"NewConnectionString,omitempty" xml:"NewConnectionString,omitempty"`
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

func (s *ModifyDBInstanceConnectionStringRequest) SetNodeId(v string) *ModifyDBInstanceConnectionStringRequest {
	s.NodeId = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetCurrentConnectionString(v string) *ModifyDBInstanceConnectionStringRequest {
	s.CurrentConnectionString = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetNewConnectionString(v string) *ModifyDBInstanceConnectionStringRequest {
	s.NewConnectionString = &v
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

type ModifyDBInstanceDescriptionRequest struct {
	SecurityToken         *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId          *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NodeId                *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
}

func (s ModifyDBInstanceDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDescriptionRequest) SetSecurityToken(v string) *ModifyDBInstanceDescriptionRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetOwnerId(v int64) *ModifyDBInstanceDescriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceDescriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceDescriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetOwnerAccount(v string) *ModifyDBInstanceDescriptionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetDBInstanceId(v string) *ModifyDBInstanceDescriptionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetNodeId(v string) *ModifyDBInstanceDescriptionRequest {
	s.NodeId = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetDBInstanceDescription(v string) *ModifyDBInstanceDescriptionRequest {
	s.DBInstanceDescription = &v
	return s
}

type ModifyDBInstanceDescriptionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDescriptionResponseBody) SetRequestId(v string) *ModifyDBInstanceDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceDescriptionResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBInstanceDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBInstanceDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDescriptionResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceDescriptionResponse) SetBody(v *ModifyDBInstanceDescriptionResponseBody) *ModifyDBInstanceDescriptionResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceMaintainTimeRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	MaintainStartTime    *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	MaintainEndTime      *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
}

func (s ModifyDBInstanceMaintainTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceMaintainTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetSecurityToken(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetOwnerId(v int64) *ModifyDBInstanceMaintainTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceMaintainTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetOwnerAccount(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetDBInstanceId(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetMaintainStartTime(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.MaintainStartTime = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetMaintainEndTime(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.MaintainEndTime = &v
	return s
}

type ModifyDBInstanceMaintainTimeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceMaintainTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceMaintainTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMaintainTimeResponseBody) SetRequestId(v string) *ModifyDBInstanceMaintainTimeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceMaintainTimeResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBInstanceMaintainTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBInstanceMaintainTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceMaintainTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMaintainTimeResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceMaintainTimeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceMaintainTimeResponse) SetBody(v *ModifyDBInstanceMaintainTimeResponseBody) *ModifyDBInstanceMaintainTimeResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceMonitorRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Granularity          *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
}

func (s ModifyDBInstanceMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceMonitorRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMonitorRequest) SetSecurityToken(v string) *ModifyDBInstanceMonitorRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyDBInstanceMonitorRequest) SetOwnerId(v int64) *ModifyDBInstanceMonitorRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceMonitorRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceMonitorRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceMonitorRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceMonitorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceMonitorRequest) SetOwnerAccount(v string) *ModifyDBInstanceMonitorRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceMonitorRequest) SetDBInstanceId(v string) *ModifyDBInstanceMonitorRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceMonitorRequest) SetGranularity(v string) *ModifyDBInstanceMonitorRequest {
	s.Granularity = &v
	return s
}

type ModifyDBInstanceMonitorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMonitorResponseBody) SetRequestId(v string) *ModifyDBInstanceMonitorResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceMonitorResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBInstanceMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBInstanceMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceMonitorResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMonitorResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceMonitorResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceMonitorResponse) SetBody(v *ModifyDBInstanceMonitorResponseBody) *ModifyDBInstanceMonitorResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceNetExpireTimeRequest struct {
	SecurityToken            *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount     *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId          *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount             *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId             *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ConnectionString         *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	ClassicExpendExpiredDays *int32  `json:"ClassicExpendExpiredDays,omitempty" xml:"ClassicExpendExpiredDays,omitempty"`
}

func (s ModifyDBInstanceNetExpireTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceNetExpireTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceNetExpireTimeRequest) SetSecurityToken(v string) *ModifyDBInstanceNetExpireTimeRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyDBInstanceNetExpireTimeRequest) SetOwnerId(v int64) *ModifyDBInstanceNetExpireTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceNetExpireTimeRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceNetExpireTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceNetExpireTimeRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceNetExpireTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceNetExpireTimeRequest) SetOwnerAccount(v string) *ModifyDBInstanceNetExpireTimeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceNetExpireTimeRequest) SetDBInstanceId(v string) *ModifyDBInstanceNetExpireTimeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceNetExpireTimeRequest) SetConnectionString(v string) *ModifyDBInstanceNetExpireTimeRequest {
	s.ConnectionString = &v
	return s
}

func (s *ModifyDBInstanceNetExpireTimeRequest) SetClassicExpendExpiredDays(v int32) *ModifyDBInstanceNetExpireTimeRequest {
	s.ClassicExpendExpiredDays = &v
	return s
}

type ModifyDBInstanceNetExpireTimeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceNetExpireTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceNetExpireTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceNetExpireTimeResponseBody) SetRequestId(v string) *ModifyDBInstanceNetExpireTimeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceNetExpireTimeResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBInstanceNetExpireTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBInstanceNetExpireTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceNetExpireTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceNetExpireTimeResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceNetExpireTimeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceNetExpireTimeResponse) SetBody(v *ModifyDBInstanceNetExpireTimeResponseBody) *ModifyDBInstanceNetExpireTimeResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceNetworkTypeRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NetworkType          *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	RetainClassic        *string `json:"RetainClassic,omitempty" xml:"RetainClassic,omitempty"`
	ClassicExpiredDays   *int32  `json:"ClassicExpiredDays,omitempty" xml:"ClassicExpiredDays,omitempty"`
}

func (s ModifyDBInstanceNetworkTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceNetworkTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetSecurityToken(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetOwnerId(v int64) *ModifyDBInstanceNetworkTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceNetworkTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetOwnerAccount(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetDBInstanceId(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetNetworkType(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.NetworkType = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetVpcId(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.VpcId = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetVSwitchId(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetRetainClassic(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.RetainClassic = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetClassicExpiredDays(v int32) *ModifyDBInstanceNetworkTypeRequest {
	s.ClassicExpiredDays = &v
	return s
}

type ModifyDBInstanceNetworkTypeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceNetworkTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceNetworkTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceNetworkTypeResponseBody) SetRequestId(v string) *ModifyDBInstanceNetworkTypeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceNetworkTypeResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBInstanceNetworkTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBInstanceNetworkTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceNetworkTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceNetworkTypeResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceNetworkTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceNetworkTypeResponse) SetBody(v *ModifyDBInstanceNetworkTypeResponseBody) *ModifyDBInstanceNetworkTypeResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceSpecRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceClass      *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	DBInstanceStorage    *string `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	OrderType            *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	AutoPay              *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	FromApp              *string `json:"FromApp,omitempty" xml:"FromApp,omitempty"`
	BusinessInfo         *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	ReplicationFactor    *string `json:"ReplicationFactor,omitempty" xml:"ReplicationFactor,omitempty"`
	ReadonlyReplicas     *string `json:"ReadonlyReplicas,omitempty" xml:"ReadonlyReplicas,omitempty"`
	CouponNo             *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	EffectiveTime        *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
}

func (s ModifyDBInstanceSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSpecRequest) SetSecurityToken(v string) *ModifyDBInstanceSpecRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetOwnerId(v int64) *ModifyDBInstanceSpecRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceSpecRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceSpecRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetOwnerAccount(v string) *ModifyDBInstanceSpecRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetDBInstanceId(v string) *ModifyDBInstanceSpecRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetDBInstanceClass(v string) *ModifyDBInstanceSpecRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetDBInstanceStorage(v string) *ModifyDBInstanceSpecRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetOrderType(v string) *ModifyDBInstanceSpecRequest {
	s.OrderType = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetAutoPay(v bool) *ModifyDBInstanceSpecRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetFromApp(v string) *ModifyDBInstanceSpecRequest {
	s.FromApp = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetBusinessInfo(v string) *ModifyDBInstanceSpecRequest {
	s.BusinessInfo = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetReplicationFactor(v string) *ModifyDBInstanceSpecRequest {
	s.ReplicationFactor = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetReadonlyReplicas(v string) *ModifyDBInstanceSpecRequest {
	s.ReadonlyReplicas = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetCouponNo(v string) *ModifyDBInstanceSpecRequest {
	s.CouponNo = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetEffectiveTime(v string) *ModifyDBInstanceSpecRequest {
	s.EffectiveTime = &v
	return s
}

type ModifyDBInstanceSpecResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ModifyDBInstanceSpecResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSpecResponseBody) SetRequestId(v string) *ModifyDBInstanceSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceSpecResponseBody) SetOrderId(v string) *ModifyDBInstanceSpecResponseBody {
	s.OrderId = &v
	return s
}

type ModifyDBInstanceSpecResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBInstanceSpecResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBInstanceSpecResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSpecResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceSpecResponse) SetBody(v *ModifyDBInstanceSpecResponseBody) *ModifyDBInstanceSpecResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceSSLRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	SSLAction            *string `json:"SSLAction,omitempty" xml:"SSLAction,omitempty"`
}

func (s ModifyDBInstanceSSLRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceSSLRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSSLRequest) SetSecurityToken(v string) *ModifyDBInstanceSSLRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetOwnerId(v int64) *ModifyDBInstanceSSLRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceSSLRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceSSLRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetOwnerAccount(v string) *ModifyDBInstanceSSLRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetDBInstanceId(v string) *ModifyDBInstanceSSLRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetSSLAction(v string) *ModifyDBInstanceSSLRequest {
	s.SSLAction = &v
	return s
}

type ModifyDBInstanceSSLResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceSSLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceSSLResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSSLResponseBody) SetRequestId(v string) *ModifyDBInstanceSSLResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceSSLResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBInstanceSSLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBInstanceSSLResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceSSLResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSSLResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceSSLResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceSSLResponse) SetBody(v *ModifyDBInstanceSSLResponseBody) *ModifyDBInstanceSSLResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceTDERequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	TDEStatus            *string `json:"TDEStatus,omitempty" xml:"TDEStatus,omitempty"`
	EncryptorName        *string `json:"EncryptorName,omitempty" xml:"EncryptorName,omitempty"`
	EncryptionKey        *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	RoleARN              *string `json:"RoleARN,omitempty" xml:"RoleARN,omitempty"`
}

func (s ModifyDBInstanceTDERequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceTDERequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceTDERequest) SetSecurityToken(v string) *ModifyDBInstanceTDERequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetOwnerId(v int64) *ModifyDBInstanceTDERequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceTDERequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetResourceOwnerId(v int64) *ModifyDBInstanceTDERequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetOwnerAccount(v string) *ModifyDBInstanceTDERequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetDBInstanceId(v string) *ModifyDBInstanceTDERequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetTDEStatus(v string) *ModifyDBInstanceTDERequest {
	s.TDEStatus = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetEncryptorName(v string) *ModifyDBInstanceTDERequest {
	s.EncryptorName = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetEncryptionKey(v string) *ModifyDBInstanceTDERequest {
	s.EncryptionKey = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetRoleARN(v string) *ModifyDBInstanceTDERequest {
	s.RoleARN = &v
	return s
}

type ModifyDBInstanceTDEResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceTDEResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceTDEResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceTDEResponseBody) SetRequestId(v string) *ModifyDBInstanceTDEResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceTDEResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBInstanceTDEResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBInstanceTDEResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceTDEResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceTDEResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceTDEResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceTDEResponse) SetBody(v *ModifyDBInstanceTDEResponseBody) *ModifyDBInstanceTDEResponse {
	s.Body = v
	return s
}

type ModifyInstanceAutoRenewalAttributeRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *ModifyInstanceAutoRenewalAttributeRequest) SetRegionId(v string) *ModifyInstanceAutoRenewalAttributeRequest {
	s.RegionId = &v
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

type ModifyInstanceVpcAuthModeRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
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

func (s *ModifyInstanceVpcAuthModeRequest) SetDBInstanceId(v string) *ModifyInstanceVpcAuthModeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyInstanceVpcAuthModeRequest) SetNodeId(v string) *ModifyInstanceVpcAuthModeRequest {
	s.NodeId = &v
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

type ModifyNodeSpecRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeClass            *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	NodeStorage          *int32  `json:"NodeStorage,omitempty" xml:"NodeStorage,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	FromApp              *string `json:"FromApp,omitempty" xml:"FromApp,omitempty"`
	AutoPay              *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	EffectiveTime        *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	OrderType            *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	ReadonlyReplicas     *int32  `json:"ReadonlyReplicas,omitempty" xml:"ReadonlyReplicas,omitempty"`
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

func (s *ModifyNodeSpecRequest) SetDBInstanceId(v string) *ModifyNodeSpecRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetNodeId(v string) *ModifyNodeSpecRequest {
	s.NodeId = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetNodeClass(v string) *ModifyNodeSpecRequest {
	s.NodeClass = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetNodeStorage(v int32) *ModifyNodeSpecRequest {
	s.NodeStorage = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetClientToken(v string) *ModifyNodeSpecRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetFromApp(v string) *ModifyNodeSpecRequest {
	s.FromApp = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetAutoPay(v bool) *ModifyNodeSpecRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetEffectiveTime(v string) *ModifyNodeSpecRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetOrderType(v string) *ModifyNodeSpecRequest {
	s.OrderType = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetReadonlyReplicas(v int32) *ModifyNodeSpecRequest {
	s.ReadonlyReplicas = &v
	return s
}

type ModifyNodeSpecResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
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

func (s *ModifyNodeSpecResponseBody) SetOrderId(v string) *ModifyNodeSpecResponseBody {
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

type ModifyParametersRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Parameters           *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	CharacterType        *string `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
}

func (s ModifyParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyParametersRequest) GoString() string {
	return s.String()
}

func (s *ModifyParametersRequest) SetSecurityToken(v string) *ModifyParametersRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyParametersRequest) SetOwnerId(v int64) *ModifyParametersRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyParametersRequest) SetResourceOwnerAccount(v string) *ModifyParametersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyParametersRequest) SetResourceOwnerId(v int64) *ModifyParametersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyParametersRequest) SetOwnerAccount(v string) *ModifyParametersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyParametersRequest) SetDBInstanceId(v string) *ModifyParametersRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyParametersRequest) SetNodeId(v string) *ModifyParametersRequest {
	s.NodeId = &v
	return s
}

func (s *ModifyParametersRequest) SetParameters(v string) *ModifyParametersRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyParametersRequest) SetCharacterType(v string) *ModifyParametersRequest {
	s.CharacterType = &v
	return s
}

type ModifyParametersResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyParametersResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyParametersResponseBody) SetRequestId(v string) *ModifyParametersResponseBody {
	s.RequestId = &v
	return s
}

type ModifyParametersResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyParametersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyParametersResponse) GoString() string {
	return s.String()
}

func (s *ModifyParametersResponse) SetHeaders(v map[string]*string) *ModifyParametersResponse {
	s.Headers = v
	return s
}

func (s *ModifyParametersResponse) SetBody(v *ModifyParametersResponseBody) *ModifyParametersResponse {
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
	DBInstanceId             *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	SecurityIps              *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
	ModifyMode               *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	SecurityIpGroupName      *string `json:"SecurityIpGroupName,omitempty" xml:"SecurityIpGroupName,omitempty"`
	SecurityIpGroupAttribute *string `json:"SecurityIpGroupAttribute,omitempty" xml:"SecurityIpGroupAttribute,omitempty"`
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

func (s *ModifySecurityIpsRequest) SetDBInstanceId(v string) *ModifySecurityIpsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetSecurityIps(v string) *ModifySecurityIpsRequest {
	s.SecurityIps = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetModifyMode(v string) *ModifySecurityIpsRequest {
	s.ModifyMode = &v
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

type ReleaseNodePrivateNetworkAddressRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NetworkType          *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
}

func (s ReleaseNodePrivateNetworkAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseNodePrivateNetworkAddressRequest) GoString() string {
	return s.String()
}

func (s *ReleaseNodePrivateNetworkAddressRequest) SetSecurityToken(v string) *ReleaseNodePrivateNetworkAddressRequest {
	s.SecurityToken = &v
	return s
}

func (s *ReleaseNodePrivateNetworkAddressRequest) SetOwnerId(v int64) *ReleaseNodePrivateNetworkAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseNodePrivateNetworkAddressRequest) SetResourceOwnerAccount(v string) *ReleaseNodePrivateNetworkAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleaseNodePrivateNetworkAddressRequest) SetResourceOwnerId(v int64) *ReleaseNodePrivateNetworkAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleaseNodePrivateNetworkAddressRequest) SetOwnerAccount(v string) *ReleaseNodePrivateNetworkAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReleaseNodePrivateNetworkAddressRequest) SetDBInstanceId(v string) *ReleaseNodePrivateNetworkAddressRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ReleaseNodePrivateNetworkAddressRequest) SetNodeId(v string) *ReleaseNodePrivateNetworkAddressRequest {
	s.NodeId = &v
	return s
}

func (s *ReleaseNodePrivateNetworkAddressRequest) SetNetworkType(v string) *ReleaseNodePrivateNetworkAddressRequest {
	s.NetworkType = &v
	return s
}

type ReleaseNodePrivateNetworkAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseNodePrivateNetworkAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseNodePrivateNetworkAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseNodePrivateNetworkAddressResponseBody) SetRequestId(v string) *ReleaseNodePrivateNetworkAddressResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseNodePrivateNetworkAddressResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleaseNodePrivateNetworkAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseNodePrivateNetworkAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseNodePrivateNetworkAddressResponse) GoString() string {
	return s.String()
}

func (s *ReleaseNodePrivateNetworkAddressResponse) SetHeaders(v map[string]*string) *ReleaseNodePrivateNetworkAddressResponse {
	s.Headers = v
	return s
}

func (s *ReleaseNodePrivateNetworkAddressResponse) SetBody(v *ReleaseNodePrivateNetworkAddressResponseBody) *ReleaseNodePrivateNetworkAddressResponse {
	s.Body = v
	return s
}

type ReleasePublicNetworkAddressRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s ReleasePublicNetworkAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleasePublicNetworkAddressRequest) GoString() string {
	return s.String()
}

func (s *ReleasePublicNetworkAddressRequest) SetSecurityToken(v string) *ReleasePublicNetworkAddressRequest {
	s.SecurityToken = &v
	return s
}

func (s *ReleasePublicNetworkAddressRequest) SetOwnerId(v int64) *ReleasePublicNetworkAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleasePublicNetworkAddressRequest) SetResourceOwnerAccount(v string) *ReleasePublicNetworkAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleasePublicNetworkAddressRequest) SetResourceOwnerId(v int64) *ReleasePublicNetworkAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleasePublicNetworkAddressRequest) SetOwnerAccount(v string) *ReleasePublicNetworkAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReleasePublicNetworkAddressRequest) SetDBInstanceId(v string) *ReleasePublicNetworkAddressRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ReleasePublicNetworkAddressRequest) SetNodeId(v string) *ReleasePublicNetworkAddressRequest {
	s.NodeId = &v
	return s
}

type ReleasePublicNetworkAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleasePublicNetworkAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleasePublicNetworkAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ReleasePublicNetworkAddressResponseBody) SetRequestId(v string) *ReleasePublicNetworkAddressResponseBody {
	s.RequestId = &v
	return s
}

type ReleasePublicNetworkAddressResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleasePublicNetworkAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleasePublicNetworkAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleasePublicNetworkAddressResponse) GoString() string {
	return s.String()
}

func (s *ReleasePublicNetworkAddressResponse) SetHeaders(v map[string]*string) *ReleasePublicNetworkAddressResponse {
	s.Headers = v
	return s
}

func (s *ReleasePublicNetworkAddressResponse) SetBody(v *ReleasePublicNetworkAddressResponseBody) *ReleasePublicNetworkAddressResponse {
	s.Body = v
	return s
}

type RenewDBInstanceRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Period               *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	AutoPay              *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	BusinessInfo         *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	CouponNo             *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
}

func (s RenewDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewDBInstanceRequest) SetSecurityToken(v string) *RenewDBInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *RenewDBInstanceRequest) SetOwnerId(v int64) *RenewDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RenewDBInstanceRequest) SetResourceOwnerAccount(v string) *RenewDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RenewDBInstanceRequest) SetResourceOwnerId(v int64) *RenewDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RenewDBInstanceRequest) SetOwnerAccount(v string) *RenewDBInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RenewDBInstanceRequest) SetClientToken(v string) *RenewDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RenewDBInstanceRequest) SetDBInstanceId(v string) *RenewDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *RenewDBInstanceRequest) SetPeriod(v int32) *RenewDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *RenewDBInstanceRequest) SetAutoPay(v bool) *RenewDBInstanceRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewDBInstanceRequest) SetBusinessInfo(v string) *RenewDBInstanceRequest {
	s.BusinessInfo = &v
	return s
}

func (s *RenewDBInstanceRequest) SetCouponNo(v string) *RenewDBInstanceRequest {
	s.CouponNo = &v
	return s
}

type RenewDBInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s RenewDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewDBInstanceResponseBody) SetRequestId(v string) *RenewDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewDBInstanceResponseBody) SetOrderId(v string) *RenewDBInstanceResponseBody {
	s.OrderId = &v
	return s
}

type RenewDBInstanceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenewDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenewDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *RenewDBInstanceResponse) SetHeaders(v map[string]*string) *RenewDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *RenewDBInstanceResponse) SetBody(v *RenewDBInstanceResponseBody) *RenewDBInstanceResponse {
	s.Body = v
	return s
}

type ResetAccountPasswordRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
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

func (s *ResetAccountPasswordRequest) SetDBInstanceId(v string) *ResetAccountPasswordRequest {
	s.DBInstanceId = &v
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

type RestartDBInstanceRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s RestartDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartDBInstanceRequest) SetSecurityToken(v string) *RestartDBInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *RestartDBInstanceRequest) SetOwnerId(v int64) *RestartDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RestartDBInstanceRequest) SetResourceOwnerAccount(v string) *RestartDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestartDBInstanceRequest) SetResourceOwnerId(v int64) *RestartDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestartDBInstanceRequest) SetOwnerAccount(v string) *RestartDBInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RestartDBInstanceRequest) SetDBInstanceId(v string) *RestartDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *RestartDBInstanceRequest) SetNodeId(v string) *RestartDBInstanceRequest {
	s.NodeId = &v
	return s
}

type RestartDBInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartDBInstanceResponseBody) SetRequestId(v string) *RestartDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RestartDBInstanceResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RestartDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartDBInstanceResponse) SetHeaders(v map[string]*string) *RestartDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestartDBInstanceResponse) SetBody(v *RestartDBInstanceResponseBody) *RestartDBInstanceResponse {
	s.Body = v
	return s
}

type RestoreDBInstanceRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	BackupId             *int32  `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
}

func (s RestoreDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RestoreDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestoreDBInstanceRequest) SetSecurityToken(v string) *RestoreDBInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetOwnerId(v int64) *RestoreDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetResourceOwnerAccount(v string) *RestoreDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetResourceOwnerId(v int64) *RestoreDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetOwnerAccount(v string) *RestoreDBInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetDBInstanceId(v string) *RestoreDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetBackupId(v int32) *RestoreDBInstanceRequest {
	s.BackupId = &v
	return s
}

type RestoreDBInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestoreDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestoreDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreDBInstanceResponseBody) SetRequestId(v string) *RestoreDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RestoreDBInstanceResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RestoreDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestoreDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RestoreDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestoreDBInstanceResponse) SetHeaders(v map[string]*string) *RestoreDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestoreDBInstanceResponse) SetBody(v *RestoreDBInstanceResponseBody) *RestoreDBInstanceResponse {
	s.Body = v
	return s
}

type SwitchDBInstanceHARequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	RoleIds              *string `json:"RoleIds,omitempty" xml:"RoleIds,omitempty"`
	SwitchMode           *int32  `json:"SwitchMode,omitempty" xml:"SwitchMode,omitempty"`
}

func (s SwitchDBInstanceHARequest) String() string {
	return tea.Prettify(s)
}

func (s SwitchDBInstanceHARequest) GoString() string {
	return s.String()
}

func (s *SwitchDBInstanceHARequest) SetSecurityToken(v string) *SwitchDBInstanceHARequest {
	s.SecurityToken = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetOwnerId(v int64) *SwitchDBInstanceHARequest {
	s.OwnerId = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetResourceOwnerAccount(v string) *SwitchDBInstanceHARequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetResourceOwnerId(v int64) *SwitchDBInstanceHARequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetOwnerAccount(v string) *SwitchDBInstanceHARequest {
	s.OwnerAccount = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetDBInstanceId(v string) *SwitchDBInstanceHARequest {
	s.DBInstanceId = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetNodeId(v string) *SwitchDBInstanceHARequest {
	s.NodeId = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetRoleIds(v string) *SwitchDBInstanceHARequest {
	s.RoleIds = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetSwitchMode(v int32) *SwitchDBInstanceHARequest {
	s.SwitchMode = &v
	return s
}

type SwitchDBInstanceHAResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchDBInstanceHAResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SwitchDBInstanceHAResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchDBInstanceHAResponseBody) SetRequestId(v string) *SwitchDBInstanceHAResponseBody {
	s.RequestId = &v
	return s
}

type SwitchDBInstanceHAResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SwitchDBInstanceHAResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SwitchDBInstanceHAResponse) String() string {
	return tea.Prettify(s)
}

func (s SwitchDBInstanceHAResponse) GoString() string {
	return s.String()
}

func (s *SwitchDBInstanceHAResponse) SetHeaders(v map[string]*string) *SwitchDBInstanceHAResponse {
	s.Headers = v
	return s
}

func (s *SwitchDBInstanceHAResponse) SetBody(v *SwitchDBInstanceHAResponseBody) *SwitchDBInstanceHAResponse {
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
	ResourceGroupId      *string                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *TagResourcesRequest) SetResourceGroupId(v string) *TagResourcesRequest {
	s.ResourceGroupId = &v
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
	BusinessInfo         *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	AutoRenew            *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	CouponNo             *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
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

func (s *TransformToPrePaidRequest) SetBusinessInfo(v string) *TransformToPrePaidRequest {
	s.BusinessInfo = &v
	return s
}

func (s *TransformToPrePaidRequest) SetAutoRenew(v string) *TransformToPrePaidRequest {
	s.AutoRenew = &v
	return s
}

func (s *TransformToPrePaidRequest) SetCouponNo(v string) *TransformToPrePaidRequest {
	s.CouponNo = &v
	return s
}

type TransformToPrePaidResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s TransformToPrePaidResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransformToPrePaidResponseBody) GoString() string {
	return s.String()
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
	ResourceGroupId      *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *UntagResourcesRequest) SetResourceGroupId(v string) *UntagResourcesRequest {
	s.ResourceGroupId = &v
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

type UpgradeDBInstanceEngineVersionRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	EngineVersion        *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
}

func (s UpgradeDBInstanceEngineVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBInstanceEngineVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetSecurityToken(v string) *UpgradeDBInstanceEngineVersionRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetOwnerId(v int64) *UpgradeDBInstanceEngineVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetResourceOwnerAccount(v string) *UpgradeDBInstanceEngineVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetResourceOwnerId(v int64) *UpgradeDBInstanceEngineVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetOwnerAccount(v string) *UpgradeDBInstanceEngineVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetDBInstanceId(v string) *UpgradeDBInstanceEngineVersionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetEngineVersion(v string) *UpgradeDBInstanceEngineVersionRequest {
	s.EngineVersion = &v
	return s
}

type UpgradeDBInstanceEngineVersionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeDBInstanceEngineVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBInstanceEngineVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceEngineVersionResponseBody) SetRequestId(v string) *UpgradeDBInstanceEngineVersionResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeDBInstanceEngineVersionResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeDBInstanceEngineVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeDBInstanceEngineVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBInstanceEngineVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceEngineVersionResponse) SetHeaders(v map[string]*string) *UpgradeDBInstanceEngineVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeDBInstanceEngineVersionResponse) SetBody(v *UpgradeDBInstanceEngineVersionResponseBody) *UpgradeDBInstanceEngineVersionResponse {
	s.Body = v
	return s
}

type UpgradeDBInstanceKernelVersionRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s UpgradeDBInstanceKernelVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBInstanceKernelVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetSecurityToken(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetOwnerId(v int64) *UpgradeDBInstanceKernelVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetResourceOwnerAccount(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetResourceOwnerId(v int64) *UpgradeDBInstanceKernelVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetOwnerAccount(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetDBInstanceId(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.DBInstanceId = &v
	return s
}

type UpgradeDBInstanceKernelVersionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeDBInstanceKernelVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBInstanceKernelVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceKernelVersionResponseBody) SetRequestId(v string) *UpgradeDBInstanceKernelVersionResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeDBInstanceKernelVersionResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeDBInstanceKernelVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeDBInstanceKernelVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBInstanceKernelVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceKernelVersionResponse) SetHeaders(v map[string]*string) *UpgradeDBInstanceKernelVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeDBInstanceKernelVersionResponse) SetBody(v *UpgradeDBInstanceKernelVersionResponseBody) *UpgradeDBInstanceKernelVersionResponse {
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
		"cn-qingdao":                  tea.String("mongodb.aliyuncs.com"),
		"cn-beijing":                  tea.String("mongodb.aliyuncs.com"),
		"cn-chengdu":                  tea.String("mongodb.cn-chengdu.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("mongodb.cn-zhangjiakou.aliyuncs.com"),
		"cn-huhehaote":                tea.String("mongodb.cn-huhehaote.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("mongodb.aliyuncs.com"),
		"cn-shanghai":                 tea.String("mongodb.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("mongodb.aliyuncs.com"),
		"cn-heyuan":                   tea.String("mongodb.aliyuncs.com"),
		"cn-hongkong":                 tea.String("mongodb.aliyuncs.com"),
		"ap-southeast-1":              tea.String("mongodb.aliyuncs.com"),
		"ap-southeast-2":              tea.String("mongodb.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-3":              tea.String("mongodb.ap-southeast-3.aliyuncs.com"),
		"ap-southeast-5":              tea.String("mongodb.ap-southeast-5.aliyuncs.com"),
		"ap-northeast-1":              tea.String("mongodb.ap-northeast-1.aliyuncs.com"),
		"eu-west-1":                   tea.String("mongodb.eu-west-1.aliyuncs.com"),
		"us-west-1":                   tea.String("mongodb.aliyuncs.com"),
		"us-east-1":                   tea.String("mongodb.aliyuncs.com"),
		"eu-central-1":                tea.String("mongodb.eu-central-1.aliyuncs.com"),
		"me-east-1":                   tea.String("mongodb.me-east-1.aliyuncs.com"),
		"ap-south-1":                  tea.String("mongodb.ap-south-1.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("mongodb.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("mongodb.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("mongodb.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("mongodb.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("mongodb.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("mongodb.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("mongodb.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("mongodb.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("mongodb.aliyuncs.com"),
		"cn-edge-1":                   tea.String("mongodb.aliyuncs.com"),
		"cn-fujian":                   tea.String("mongodb.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("mongodb.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("mongodb.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("mongodb.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("mongodb.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("mongodb.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("mongodb.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("mongodb.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("mongodb.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("mongodb.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("mongodb.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("mongodb.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("mongodb.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("mongodb.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("mongodb.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("mongodb.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("mongodb.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("mongodb.aliyuncs.com"),
		"cn-wuhan":                    tea.String("mongodb.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("mongodb.aliyuncs.com"),
		"cn-yushanfang":               tea.String("mongodb.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("mongodb.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("mongodb.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("mongodb.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("mongodb.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("mongodb.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("dds"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AllocateNodePrivateNetworkAddressWithOptions(request *AllocateNodePrivateNetworkAddressRequest, runtime *util.RuntimeOptions) (_result *AllocateNodePrivateNetworkAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AllocateNodePrivateNetworkAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AllocateNodePrivateNetworkAddress"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AllocateNodePrivateNetworkAddress(request *AllocateNodePrivateNetworkAddressRequest) (_result *AllocateNodePrivateNetworkAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AllocateNodePrivateNetworkAddressResponse{}
	_body, _err := client.AllocateNodePrivateNetworkAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AllocatePublicNetworkAddressWithOptions(request *AllocatePublicNetworkAddressRequest, runtime *util.RuntimeOptions) (_result *AllocatePublicNetworkAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AllocatePublicNetworkAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AllocatePublicNetworkAddress"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AllocatePublicNetworkAddress(request *AllocatePublicNetworkAddressRequest) (_result *AllocatePublicNetworkAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AllocatePublicNetworkAddressResponse{}
	_body, _err := client.AllocatePublicNetworkAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckCloudResourceAuthorizedWithOptions(request *CheckCloudResourceAuthorizedRequest, runtime *util.RuntimeOptions) (_result *CheckCloudResourceAuthorizedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckCloudResourceAuthorizedResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckCloudResourceAuthorized"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckCloudResourceAuthorized(request *CheckCloudResourceAuthorizedRequest) (_result *CheckCloudResourceAuthorizedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckCloudResourceAuthorizedResponse{}
	_body, _err := client.CheckCloudResourceAuthorizedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckRecoveryConditionWithOptions(request *CheckRecoveryConditionRequest, runtime *util.RuntimeOptions) (_result *CheckRecoveryConditionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckRecoveryConditionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckRecoveryCondition"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckRecoveryCondition(request *CheckRecoveryConditionRequest) (_result *CheckRecoveryConditionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckRecoveryConditionResponse{}
	_body, _err := client.CheckRecoveryConditionWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("CreateBackup"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateDBInstanceWithOptions(request *CreateDBInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDBInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDBInstance"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDBInstance(request *CreateDBInstanceRequest) (_result *CreateDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBInstanceResponse{}
	_body, _err := client.CreateDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateNodeWithOptions(request *CreateNodeRequest, runtime *util.RuntimeOptions) (_result *CreateNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateNodeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateNode"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateNode(request *CreateNodeRequest) (_result *CreateNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNodeResponse{}
	_body, _err := client.CreateNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRecommendationTaskWithOptions(request *CreateRecommendationTaskRequest, runtime *util.RuntimeOptions) (_result *CreateRecommendationTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateRecommendationTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateRecommendationTask"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRecommendationTask(request *CreateRecommendationTaskRequest) (_result *CreateRecommendationTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRecommendationTaskResponse{}
	_body, _err := client.CreateRecommendationTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServerlessDBInstanceWithOptions(request *CreateServerlessDBInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateServerlessDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateServerlessDBInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateServerlessDBInstance"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateServerlessDBInstance(request *CreateServerlessDBInstanceRequest) (_result *CreateServerlessDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServerlessDBInstanceResponse{}
	_body, _err := client.CreateServerlessDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateShardingDBInstanceWithOptions(request *CreateShardingDBInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateShardingDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateShardingDBInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateShardingDBInstance"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateShardingDBInstance(request *CreateShardingDBInstanceRequest) (_result *CreateShardingDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateShardingDBInstanceResponse{}
	_body, _err := client.CreateShardingDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDBInstanceWithOptions(request *DeleteDBInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDBInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDBInstance"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDBInstance(request *DeleteDBInstanceRequest) (_result *DeleteDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDBInstanceResponse{}
	_body, _err := client.DeleteDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteNodeWithOptions(request *DeleteNodeRequest, runtime *util.RuntimeOptions) (_result *DeleteNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteNodeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteNode"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteNode(request *DeleteNodeRequest) (_result *DeleteNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNodeResponse{}
	_body, _err := client.DeleteNodeWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeAccounts"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeActiveOperationTaskCountWithOptions(request *DescribeActiveOperationTaskCountRequest, runtime *util.RuntimeOptions) (_result *DescribeActiveOperationTaskCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeActiveOperationTaskCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeActiveOperationTaskCount"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeActiveOperationTaskCount(request *DescribeActiveOperationTaskCountRequest) (_result *DescribeActiveOperationTaskCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeActiveOperationTaskCountResponse{}
	_body, _err := client.DescribeActiveOperationTaskCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeActiveOperationTaskTypeWithOptions(request *DescribeActiveOperationTaskTypeRequest, runtime *util.RuntimeOptions) (_result *DescribeActiveOperationTaskTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeActiveOperationTaskTypeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeActiveOperationTaskType"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeActiveOperationTaskType(request *DescribeActiveOperationTaskTypeRequest) (_result *DescribeActiveOperationTaskTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeActiveOperationTaskTypeResponse{}
	_body, _err := client.DescribeActiveOperationTaskTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAuditFilesWithOptions(request *DescribeAuditFilesRequest, runtime *util.RuntimeOptions) (_result *DescribeAuditFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAuditFilesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAuditFiles"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAuditFiles(request *DescribeAuditFilesRequest) (_result *DescribeAuditFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAuditFilesResponse{}
	_body, _err := client.DescribeAuditFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAuditLogFilterWithOptions(request *DescribeAuditLogFilterRequest, runtime *util.RuntimeOptions) (_result *DescribeAuditLogFilterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAuditLogFilterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAuditLogFilter"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAuditLogFilter(request *DescribeAuditLogFilterRequest) (_result *DescribeAuditLogFilterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAuditLogFilterResponse{}
	_body, _err := client.DescribeAuditLogFilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAuditPolicyWithOptions(request *DescribeAuditPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeAuditPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAuditPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAuditPolicy"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAuditPolicy(request *DescribeAuditPolicyRequest) (_result *DescribeAuditPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAuditPolicyResponse{}
	_body, _err := client.DescribeAuditPolicyWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeAuditRecords"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeAvailableResourceWithOptions(request *DescribeAvailableResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeAvailableResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAvailableResourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAvailableResource"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeAvailableTimeRangeWithOptions(request *DescribeAvailableTimeRangeRequest, runtime *util.RuntimeOptions) (_result *DescribeAvailableTimeRangeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAvailableTimeRangeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAvailableTimeRange"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAvailableTimeRange(request *DescribeAvailableTimeRangeRequest) (_result *DescribeAvailableTimeRangeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAvailableTimeRangeResponse{}
	_body, _err := client.DescribeAvailableTimeRangeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupDBsWithOptions(request *DescribeBackupDBsRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupDBsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackupDBsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackupDBs"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupDBs(request *DescribeBackupDBsRequest) (_result *DescribeBackupDBsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupDBsResponse{}
	_body, _err := client.DescribeBackupDBsWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackupPolicy"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackups"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeDBInstanceAttributeWithOptions(request *DescribeDBInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBInstanceAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBInstanceAttribute"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceAttribute(request *DescribeDBInstanceAttributeRequest) (_result *DescribeDBInstanceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceAttributeResponse{}
	_body, _err := client.DescribeDBInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstanceEncryptionKeyWithOptions(request *DescribeDBInstanceEncryptionKeyRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceEncryptionKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBInstanceEncryptionKeyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBInstanceEncryptionKey"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceEncryptionKey(request *DescribeDBInstanceEncryptionKeyRequest) (_result *DescribeDBInstanceEncryptionKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceEncryptionKeyResponse{}
	_body, _err := client.DescribeDBInstanceEncryptionKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstanceMonitorWithOptions(request *DescribeDBInstanceMonitorRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBInstanceMonitorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBInstanceMonitor"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceMonitor(request *DescribeDBInstanceMonitorRequest) (_result *DescribeDBInstanceMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceMonitorResponse{}
	_body, _err := client.DescribeDBInstanceMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstancePerformanceWithOptions(request *DescribeDBInstancePerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstancePerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBInstancePerformanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBInstancePerformance"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstancePerformance(request *DescribeDBInstancePerformanceRequest) (_result *DescribeDBInstancePerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstancePerformanceResponse{}
	_body, _err := client.DescribeDBInstancePerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstancesWithOptions(request *DescribeDBInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBInstances"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (_result *DescribeDBInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstancesResponse{}
	_body, _err := client.DescribeDBInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstanceSSLWithOptions(request *DescribeDBInstanceSSLRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceSSLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBInstanceSSLResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBInstanceSSL"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceSSL(request *DescribeDBInstanceSSLRequest) (_result *DescribeDBInstanceSSLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceSSLResponse{}
	_body, _err := client.DescribeDBInstanceSSLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstanceTDEInfoWithOptions(request *DescribeDBInstanceTDEInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceTDEInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBInstanceTDEInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBInstanceTDEInfo"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceTDEInfo(request *DescribeDBInstanceTDEInfoRequest) (_result *DescribeDBInstanceTDEInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceTDEInfoResponse{}
	_body, _err := client.DescribeDBInstanceTDEInfoWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeDedicatedClusterInstanceList"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeErrorLogRecordsWithOptions(request *DescribeErrorLogRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeErrorLogRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeErrorLogRecordsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeErrorLogRecords"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeErrorLogRecords(request *DescribeErrorLogRecordsRequest) (_result *DescribeErrorLogRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeErrorLogRecordsResponse{}
	_body, _err := client.DescribeErrorLogRecordsWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceAutoRenewalAttribute"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeKernelReleaseNotesWithOptions(request *DescribeKernelReleaseNotesRequest, runtime *util.RuntimeOptions) (_result *DescribeKernelReleaseNotesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeKernelReleaseNotesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeKernelReleaseNotes"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeKernelReleaseNotes(request *DescribeKernelReleaseNotesRequest) (_result *DescribeKernelReleaseNotesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeKernelReleaseNotesResponse{}
	_body, _err := client.DescribeKernelReleaseNotesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMongoDBLogConfigWithOptions(request *DescribeMongoDBLogConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeMongoDBLogConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMongoDBLogConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMongoDBLogConfig"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMongoDBLogConfig(request *DescribeMongoDBLogConfigRequest) (_result *DescribeMongoDBLogConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMongoDBLogConfigResponse{}
	_body, _err := client.DescribeMongoDBLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeParameterModificationHistoryWithOptions(request *DescribeParameterModificationHistoryRequest, runtime *util.RuntimeOptions) (_result *DescribeParameterModificationHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeParameterModificationHistoryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeParameterModificationHistory"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeParameterModificationHistory(request *DescribeParameterModificationHistoryRequest) (_result *DescribeParameterModificationHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeParameterModificationHistoryResponse{}
	_body, _err := client.DescribeParameterModificationHistoryWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeParameters"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeParameterTemplates"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribePrice"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeReplicaSetRoleWithOptions(request *DescribeReplicaSetRoleRequest, runtime *util.RuntimeOptions) (_result *DescribeReplicaSetRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeReplicaSetRoleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeReplicaSetRole"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeReplicaSetRole(request *DescribeReplicaSetRoleRequest) (_result *DescribeReplicaSetRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeReplicaSetRoleResponse{}
	_body, _err := client.DescribeReplicaSetRoleWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeRoleZoneInfo"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeRunningLogRecords"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeSecurityGroupConfiguration"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeSecurityIps"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeShardingNetworkAddressWithOptions(request *DescribeShardingNetworkAddressRequest, runtime *util.RuntimeOptions) (_result *DescribeShardingNetworkAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeShardingNetworkAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeShardingNetworkAddress"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeShardingNetworkAddress(request *DescribeShardingNetworkAddressRequest) (_result *DescribeShardingNetworkAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeShardingNetworkAddressResponse{}
	_body, _err := client.DescribeShardingNetworkAddressWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeSlowLogRecords"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DestroyInstanceWithOptions(request *DestroyInstanceRequest, runtime *util.RuntimeOptions) (_result *DestroyInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DestroyInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DestroyInstance"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DestroyInstance(request *DestroyInstanceRequest) (_result *DestroyInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DestroyInstanceResponse{}
	_body, _err := client.DestroyInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EvaluateResourceWithOptions(request *EvaluateResourceRequest, runtime *util.RuntimeOptions) (_result *EvaluateResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EvaluateResourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EvaluateResource"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EvaluateResource(request *EvaluateResourceRequest) (_result *EvaluateResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EvaluateResourceResponse{}
	_body, _err := client.EvaluateResourceWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("ListTagResources"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) MigrateAvailableZoneWithOptions(request *MigrateAvailableZoneRequest, runtime *util.RuntimeOptions) (_result *MigrateAvailableZoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MigrateAvailableZoneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MigrateAvailableZone"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MigrateAvailableZone(request *MigrateAvailableZoneRequest) (_result *MigrateAvailableZoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MigrateAvailableZoneResponse{}
	_body, _err := client.MigrateAvailableZoneWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("MigrateToOtherZone"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("ModifyAccountDescription"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ModifyAuditLogFilterWithOptions(request *ModifyAuditLogFilterRequest, runtime *util.RuntimeOptions) (_result *ModifyAuditLogFilterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyAuditLogFilterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyAuditLogFilter"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAuditLogFilter(request *ModifyAuditLogFilterRequest) (_result *ModifyAuditLogFilterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAuditLogFilterResponse{}
	_body, _err := client.ModifyAuditLogFilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAuditPolicyWithOptions(request *ModifyAuditPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifyAuditPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyAuditPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyAuditPolicy"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAuditPolicy(request *ModifyAuditPolicyRequest) (_result *ModifyAuditPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAuditPolicyResponse{}
	_body, _err := client.ModifyAuditPolicyWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("ModifyBackupPolicy"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBInstanceConnectionString"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ModifyDBInstanceDescriptionWithOptions(request *ModifyDBInstanceDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBInstanceDescriptionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBInstanceDescription"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceDescription(request *ModifyDBInstanceDescriptionRequest) (_result *ModifyDBInstanceDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceDescriptionResponse{}
	_body, _err := client.ModifyDBInstanceDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBInstanceMaintainTimeWithOptions(request *ModifyDBInstanceMaintainTimeRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceMaintainTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBInstanceMaintainTimeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBInstanceMaintainTime"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceMaintainTime(request *ModifyDBInstanceMaintainTimeRequest) (_result *ModifyDBInstanceMaintainTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceMaintainTimeResponse{}
	_body, _err := client.ModifyDBInstanceMaintainTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBInstanceMonitorWithOptions(request *ModifyDBInstanceMonitorRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBInstanceMonitorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBInstanceMonitor"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceMonitor(request *ModifyDBInstanceMonitorRequest) (_result *ModifyDBInstanceMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceMonitorResponse{}
	_body, _err := client.ModifyDBInstanceMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBInstanceNetExpireTimeWithOptions(request *ModifyDBInstanceNetExpireTimeRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceNetExpireTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBInstanceNetExpireTimeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBInstanceNetExpireTime"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceNetExpireTime(request *ModifyDBInstanceNetExpireTimeRequest) (_result *ModifyDBInstanceNetExpireTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceNetExpireTimeResponse{}
	_body, _err := client.ModifyDBInstanceNetExpireTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBInstanceNetworkTypeWithOptions(request *ModifyDBInstanceNetworkTypeRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceNetworkTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBInstanceNetworkTypeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBInstanceNetworkType"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceNetworkType(request *ModifyDBInstanceNetworkTypeRequest) (_result *ModifyDBInstanceNetworkTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceNetworkTypeResponse{}
	_body, _err := client.ModifyDBInstanceNetworkTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBInstanceSpecWithOptions(request *ModifyDBInstanceSpecRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceSpecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBInstanceSpecResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBInstanceSpec"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceSpec(request *ModifyDBInstanceSpecRequest) (_result *ModifyDBInstanceSpecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceSpecResponse{}
	_body, _err := client.ModifyDBInstanceSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBInstanceSSLWithOptions(request *ModifyDBInstanceSSLRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceSSLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBInstanceSSLResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBInstanceSSL"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceSSL(request *ModifyDBInstanceSSLRequest) (_result *ModifyDBInstanceSSLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceSSLResponse{}
	_body, _err := client.ModifyDBInstanceSSLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBInstanceTDEWithOptions(request *ModifyDBInstanceTDERequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceTDEResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBInstanceTDEResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBInstanceTDE"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceTDE(request *ModifyDBInstanceTDERequest) (_result *ModifyDBInstanceTDEResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceTDEResponse{}
	_body, _err := client.ModifyDBInstanceTDEWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstanceAutoRenewalAttribute"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ModifyInstanceVpcAuthModeWithOptions(request *ModifyInstanceVpcAuthModeRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceVpcAuthModeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceVpcAuthModeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstanceVpcAuthMode"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ModifyNodeSpecWithOptions(request *ModifyNodeSpecRequest, runtime *util.RuntimeOptions) (_result *ModifyNodeSpecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyNodeSpecResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyNodeSpec"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ModifyParametersWithOptions(request *ModifyParametersRequest, runtime *util.RuntimeOptions) (_result *ModifyParametersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyParametersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyParameters"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyParameters(request *ModifyParametersRequest) (_result *ModifyParametersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyParametersResponse{}
	_body, _err := client.ModifyParametersWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("ModifySecurityGroupConfiguration"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("ModifySecurityIps"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ReleaseNodePrivateNetworkAddressWithOptions(request *ReleaseNodePrivateNetworkAddressRequest, runtime *util.RuntimeOptions) (_result *ReleaseNodePrivateNetworkAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReleaseNodePrivateNetworkAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReleaseNodePrivateNetworkAddress"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseNodePrivateNetworkAddress(request *ReleaseNodePrivateNetworkAddressRequest) (_result *ReleaseNodePrivateNetworkAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseNodePrivateNetworkAddressResponse{}
	_body, _err := client.ReleaseNodePrivateNetworkAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleasePublicNetworkAddressWithOptions(request *ReleasePublicNetworkAddressRequest, runtime *util.RuntimeOptions) (_result *ReleasePublicNetworkAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReleasePublicNetworkAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReleasePublicNetworkAddress"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleasePublicNetworkAddress(request *ReleasePublicNetworkAddressRequest) (_result *ReleasePublicNetworkAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleasePublicNetworkAddressResponse{}
	_body, _err := client.ReleasePublicNetworkAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenewDBInstanceWithOptions(request *RenewDBInstanceRequest, runtime *util.RuntimeOptions) (_result *RenewDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RenewDBInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RenewDBInstance"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenewDBInstance(request *RenewDBInstanceRequest) (_result *RenewDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewDBInstanceResponse{}
	_body, _err := client.RenewDBInstanceWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("ResetAccountPassword"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RestartDBInstanceWithOptions(request *RestartDBInstanceRequest, runtime *util.RuntimeOptions) (_result *RestartDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RestartDBInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RestartDBInstance"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartDBInstance(request *RestartDBInstanceRequest) (_result *RestartDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartDBInstanceResponse{}
	_body, _err := client.RestartDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestoreDBInstanceWithOptions(request *RestoreDBInstanceRequest, runtime *util.RuntimeOptions) (_result *RestoreDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RestoreDBInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RestoreDBInstance"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestoreDBInstance(request *RestoreDBInstanceRequest) (_result *RestoreDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestoreDBInstanceResponse{}
	_body, _err := client.RestoreDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SwitchDBInstanceHAWithOptions(request *SwitchDBInstanceHARequest, runtime *util.RuntimeOptions) (_result *SwitchDBInstanceHAResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SwitchDBInstanceHAResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SwitchDBInstanceHA"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SwitchDBInstanceHA(request *SwitchDBInstanceHARequest) (_result *SwitchDBInstanceHAResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SwitchDBInstanceHAResponse{}
	_body, _err := client.SwitchDBInstanceHAWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("TransformToPrePaid"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("UntagResources"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UpgradeDBInstanceEngineVersionWithOptions(request *UpgradeDBInstanceEngineVersionRequest, runtime *util.RuntimeOptions) (_result *UpgradeDBInstanceEngineVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpgradeDBInstanceEngineVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpgradeDBInstanceEngineVersion"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeDBInstanceEngineVersion(request *UpgradeDBInstanceEngineVersionRequest) (_result *UpgradeDBInstanceEngineVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeDBInstanceEngineVersionResponse{}
	_body, _err := client.UpgradeDBInstanceEngineVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeDBInstanceKernelVersionWithOptions(request *UpgradeDBInstanceKernelVersionRequest, runtime *util.RuntimeOptions) (_result *UpgradeDBInstanceKernelVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpgradeDBInstanceKernelVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpgradeDBInstanceKernelVersion"), tea.String("2015-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeDBInstanceKernelVersion(request *UpgradeDBInstanceKernelVersionRequest) (_result *UpgradeDBInstanceKernelVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeDBInstanceKernelVersionResponse{}
	_body, _err := client.UpgradeDBInstanceKernelVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
