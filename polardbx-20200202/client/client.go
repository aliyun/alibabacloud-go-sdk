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

type AllocateInstancePublicConnectionRequest struct {
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DBInstanceName         *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	Port                   *string `json:"Port,omitempty" xml:"Port,omitempty"`
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AllocateInstancePublicConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s AllocateInstancePublicConnectionRequest) GoString() string {
	return s.String()
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

func (s *AllocateInstancePublicConnectionRequest) SetDBInstanceName(v string) *AllocateInstancePublicConnectionRequest {
	s.DBInstanceName = &v
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

func (s *AllocateInstancePublicConnectionRequest) SetOwnerAccount(v string) *AllocateInstancePublicConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetRegionId(v string) *AllocateInstancePublicConnectionRequest {
	s.RegionId = &v
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

type CancelPolarxOrderRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	ScaleOutToken  *string `json:"ScaleOutToken,omitempty" xml:"ScaleOutToken,omitempty"`
}

func (s CancelPolarxOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelPolarxOrderRequest) GoString() string {
	return s.String()
}

func (s *CancelPolarxOrderRequest) SetRegionId(v string) *CancelPolarxOrderRequest {
	s.RegionId = &v
	return s
}

func (s *CancelPolarxOrderRequest) SetDBInstanceName(v string) *CancelPolarxOrderRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CancelPolarxOrderRequest) SetScaleOutToken(v string) *CancelPolarxOrderRequest {
	s.ScaleOutToken = &v
	return s
}

type CancelPolarxOrderResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelPolarxOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelPolarxOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CancelPolarxOrderResponseBody) SetRequestId(v string) *CancelPolarxOrderResponseBody {
	s.RequestId = &v
	return s
}

type CancelPolarxOrderResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelPolarxOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelPolarxOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelPolarxOrderResponse) GoString() string {
	return s.String()
}

func (s *CancelPolarxOrderResponse) SetHeaders(v map[string]*string) *CancelPolarxOrderResponse {
	s.Headers = v
	return s
}

func (s *CancelPolarxOrderResponse) SetBody(v *CancelPolarxOrderResponseBody) *CancelPolarxOrderResponse {
	s.Body = v
	return s
}

type CheckCloudResourceAuthorizedRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RoleArn        *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s CheckCloudResourceAuthorizedRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckCloudResourceAuthorizedRequest) GoString() string {
	return s.String()
}

func (s *CheckCloudResourceAuthorizedRequest) SetRegionId(v string) *CheckCloudResourceAuthorizedRequest {
	s.RegionId = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetDBInstanceName(v string) *CheckCloudResourceAuthorizedRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetRoleArn(v string) *CheckCloudResourceAuthorizedRequest {
	s.RoleArn = &v
	return s
}

type CheckCloudResourceAuthorizedResponseBody struct {
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *CheckCloudResourceAuthorizedResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s CheckCloudResourceAuthorizedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckCloudResourceAuthorizedResponseBody) GoString() string {
	return s.String()
}

func (s *CheckCloudResourceAuthorizedResponseBody) SetRequestId(v string) *CheckCloudResourceAuthorizedResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckCloudResourceAuthorizedResponseBody) SetData(v *CheckCloudResourceAuthorizedResponseBodyData) *CheckCloudResourceAuthorizedResponseBody {
	s.Data = v
	return s
}

type CheckCloudResourceAuthorizedResponseBodyData struct {
	AuthorizationState *string `json:"AuthorizationState,omitempty" xml:"AuthorizationState,omitempty"`
	RoleArn            *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s CheckCloudResourceAuthorizedResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CheckCloudResourceAuthorizedResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckCloudResourceAuthorizedResponseBodyData) SetAuthorizationState(v string) *CheckCloudResourceAuthorizedResponseBodyData {
	s.AuthorizationState = &v
	return s
}

func (s *CheckCloudResourceAuthorizedResponseBodyData) SetRoleArn(v string) *CheckCloudResourceAuthorizedResponseBodyData {
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

type CreateAccountRequest struct {
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName          *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	AccountName             *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPassword         *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	DBName                  *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	AccountPrivilege        *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	AccountDescription      *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	SecurityAccountName     *string `json:"SecurityAccountName,omitempty" xml:"SecurityAccountName,omitempty"`
	SecurityAccountPassword *string `json:"SecurityAccountPassword,omitempty" xml:"SecurityAccountPassword,omitempty"`
}

func (s CreateAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountRequest) SetRegionId(v string) *CreateAccountRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAccountRequest) SetDBInstanceName(v string) *CreateAccountRequest {
	s.DBInstanceName = &v
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

func (s *CreateAccountRequest) SetDBName(v string) *CreateAccountRequest {
	s.DBName = &v
	return s
}

func (s *CreateAccountRequest) SetAccountPrivilege(v string) *CreateAccountRequest {
	s.AccountPrivilege = &v
	return s
}

func (s *CreateAccountRequest) SetAccountDescription(v string) *CreateAccountRequest {
	s.AccountDescription = &v
	return s
}

func (s *CreateAccountRequest) SetSecurityAccountName(v string) *CreateAccountRequest {
	s.SecurityAccountName = &v
	return s
}

func (s *CreateAccountRequest) SetSecurityAccountPassword(v string) *CreateAccountRequest {
	s.SecurityAccountPassword = &v
	return s
}

type CreateAccountResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccountResponseBody) SetMessage(v string) *CreateAccountResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAccountResponseBody) SetRequestId(v string) *CreateAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccountResponseBody) SetSuccess(v bool) *CreateAccountResponseBody {
	s.Success = &v
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
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	BackupType     *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
}

func (s CreateBackupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupRequest) GoString() string {
	return s.String()
}

func (s *CreateBackupRequest) SetRegionId(v string) *CreateBackupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateBackupRequest) SetDBInstanceName(v string) *CreateBackupRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateBackupRequest) SetBackupType(v string) *CreateBackupRequest {
	s.BackupType = &v
	return s
}

type CreateBackupResponseBody struct {
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      []*CreateBackupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s CreateBackupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBackupResponseBody) SetMessage(v string) *CreateBackupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateBackupResponseBody) SetRequestId(v string) *CreateBackupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBackupResponseBody) SetSuccess(v bool) *CreateBackupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateBackupResponseBody) SetData(v []*CreateBackupResponseBodyData) *CreateBackupResponseBody {
	s.Data = v
	return s
}

type CreateBackupResponseBodyData struct {
	BackupSetId *int64 `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
}

func (s CreateBackupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateBackupResponseBodyData) SetBackupSetId(v int64) *CreateBackupResponseBodyData {
	s.BackupSetId = &v
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

type CreateDBRequest struct {
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName          *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	AccountName             *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	Charset                 *string `json:"Charset,omitempty" xml:"Charset,omitempty"`
	DbName                  *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	AccountPrivilege        *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	DbDescription           *string `json:"DbDescription,omitempty" xml:"DbDescription,omitempty"`
	SecurityAccountName     *string `json:"SecurityAccountName,omitempty" xml:"SecurityAccountName,omitempty"`
	SecurityAccountPassword *string `json:"SecurityAccountPassword,omitempty" xml:"SecurityAccountPassword,omitempty"`
}

func (s CreateDBRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBRequest) GoString() string {
	return s.String()
}

func (s *CreateDBRequest) SetRegionId(v string) *CreateDBRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBRequest) SetDBInstanceName(v string) *CreateDBRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateDBRequest) SetAccountName(v string) *CreateDBRequest {
	s.AccountName = &v
	return s
}

func (s *CreateDBRequest) SetCharset(v string) *CreateDBRequest {
	s.Charset = &v
	return s
}

func (s *CreateDBRequest) SetDbName(v string) *CreateDBRequest {
	s.DbName = &v
	return s
}

func (s *CreateDBRequest) SetAccountPrivilege(v string) *CreateDBRequest {
	s.AccountPrivilege = &v
	return s
}

func (s *CreateDBRequest) SetDbDescription(v string) *CreateDBRequest {
	s.DbDescription = &v
	return s
}

func (s *CreateDBRequest) SetSecurityAccountName(v string) *CreateDBRequest {
	s.SecurityAccountName = &v
	return s
}

func (s *CreateDBRequest) SetSecurityAccountPassword(v string) *CreateDBRequest {
	s.SecurityAccountPassword = &v
	return s
}

type CreateDBResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBResponseBody) SetMessage(v string) *CreateDBResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDBResponseBody) SetRequestId(v string) *CreateDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBResponseBody) SetSuccess(v bool) *CreateDBResponseBody {
	s.Success = &v
	return s
}

type CreateDBResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDBResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDBResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBResponse) GoString() string {
	return s.String()
}

func (s *CreateDBResponse) SetHeaders(v map[string]*string) *CreateDBResponse {
	s.Headers = v
	return s
}

func (s *CreateDBResponse) SetBody(v *CreateDBResponseBody) *CreateDBResponse {
	s.Body = v
	return s
}

type CreateDBInstanceRequest struct {
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PayType               *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	DBNodeCount           *int32  `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	DBNodeClass           *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	ZoneId                *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ClientToken           *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	NetworkType           *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	VPCId                 *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId             *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	UsedTime              *int32  `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	Period                *string `json:"Period,omitempty" xml:"Period,omitempty"`
	ResourceGroupId       *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	AutoRenew             *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	EngineVersion         *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	IsReadDBInstance      *bool   `json:"IsReadDBInstance,omitempty" xml:"IsReadDBInstance,omitempty"`
	PrimaryDBInstanceName *string `json:"PrimaryDBInstanceName,omitempty" xml:"PrimaryDBInstanceName,omitempty"`
}

func (s CreateDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceRequest) SetRegionId(v string) *CreateDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPayType(v string) *CreateDBInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBNodeCount(v int32) *CreateDBInstanceRequest {
	s.DBNodeCount = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBNodeClass(v string) *CreateDBInstanceRequest {
	s.DBNodeClass = &v
	return s
}

func (s *CreateDBInstanceRequest) SetZoneId(v string) *CreateDBInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetClientToken(v string) *CreateDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceRequest) SetNetworkType(v string) *CreateDBInstanceRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVPCId(v string) *CreateDBInstanceRequest {
	s.VPCId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVSwitchId(v string) *CreateDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetUsedTime(v int32) *CreateDBInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPeriod(v string) *CreateDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateDBInstanceRequest) SetResourceGroupId(v string) *CreateDBInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetAutoRenew(v bool) *CreateDBInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEngineVersion(v string) *CreateDBInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateDBInstanceRequest) SetIsReadDBInstance(v bool) *CreateDBInstanceRequest {
	s.IsReadDBInstance = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPrimaryDBInstanceName(v string) *CreateDBInstanceRequest {
	s.PrimaryDBInstanceName = &v
	return s
}

type CreateDBInstanceResponseBody struct {
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId        *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceResponseBody) SetDBInstanceName(v string) *CreateDBInstanceResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetRequestId(v string) *CreateDBInstanceResponseBody {
	s.RequestId = &v
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

type CreatePolarxInstanceRequest struct {
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId         *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Quantity       *int32  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	InstanceSeries *string `json:"InstanceSeries,omitempty" xml:"InstanceSeries,omitempty"`
	Specification  *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	PayType        *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	VpcId          *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId      *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	IsHa           *bool   `json:"isHa,omitempty" xml:"isHa,omitempty"`
	PricingCycle   *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	Duration       *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	IsAutoRenew    *bool   `json:"IsAutoRenew,omitempty" xml:"IsAutoRenew,omitempty"`
	MasterInstId   *string `json:"MasterInstId,omitempty" xml:"MasterInstId,omitempty"`
	MySQLVersion   *int32  `json:"MySQLVersion,omitempty" xml:"MySQLVersion,omitempty"`
}

func (s CreatePolarxInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePolarxInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreatePolarxInstanceRequest) SetDescription(v string) *CreatePolarxInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreatePolarxInstanceRequest) SetRegionId(v string) *CreatePolarxInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePolarxInstanceRequest) SetZoneId(v string) *CreatePolarxInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreatePolarxInstanceRequest) SetType(v string) *CreatePolarxInstanceRequest {
	s.Type = &v
	return s
}

func (s *CreatePolarxInstanceRequest) SetQuantity(v int32) *CreatePolarxInstanceRequest {
	s.Quantity = &v
	return s
}

func (s *CreatePolarxInstanceRequest) SetInstanceSeries(v string) *CreatePolarxInstanceRequest {
	s.InstanceSeries = &v
	return s
}

func (s *CreatePolarxInstanceRequest) SetSpecification(v string) *CreatePolarxInstanceRequest {
	s.Specification = &v
	return s
}

func (s *CreatePolarxInstanceRequest) SetClientToken(v string) *CreatePolarxInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreatePolarxInstanceRequest) SetPayType(v string) *CreatePolarxInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreatePolarxInstanceRequest) SetVpcId(v string) *CreatePolarxInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreatePolarxInstanceRequest) SetVswitchId(v string) *CreatePolarxInstanceRequest {
	s.VswitchId = &v
	return s
}

func (s *CreatePolarxInstanceRequest) SetIsHa(v bool) *CreatePolarxInstanceRequest {
	s.IsHa = &v
	return s
}

func (s *CreatePolarxInstanceRequest) SetPricingCycle(v string) *CreatePolarxInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreatePolarxInstanceRequest) SetDuration(v int32) *CreatePolarxInstanceRequest {
	s.Duration = &v
	return s
}

func (s *CreatePolarxInstanceRequest) SetIsAutoRenew(v bool) *CreatePolarxInstanceRequest {
	s.IsAutoRenew = &v
	return s
}

func (s *CreatePolarxInstanceRequest) SetMasterInstId(v string) *CreatePolarxInstanceRequest {
	s.MasterInstId = &v
	return s
}

func (s *CreatePolarxInstanceRequest) SetMySQLVersion(v int32) *CreatePolarxInstanceRequest {
	s.MySQLVersion = &v
	return s
}

type CreatePolarxInstanceResponseBody struct {
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *CreatePolarxInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s CreatePolarxInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePolarxInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolarxInstanceResponseBody) SetSuccess(v bool) *CreatePolarxInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePolarxInstanceResponseBody) SetRequestId(v string) *CreatePolarxInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePolarxInstanceResponseBody) SetData(v *CreatePolarxInstanceResponseBodyData) *CreatePolarxInstanceResponseBody {
	s.Data = v
	return s
}

type CreatePolarxInstanceResponseBodyData struct {
	OrderId            *int64                                                  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	DrdsInstanceIdList *CreatePolarxInstanceResponseBodyDataDrdsInstanceIdList `json:"DrdsInstanceIdList,omitempty" xml:"DrdsInstanceIdList,omitempty" type:"Struct"`
}

func (s CreatePolarxInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreatePolarxInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePolarxInstanceResponseBodyData) SetOrderId(v int64) *CreatePolarxInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreatePolarxInstanceResponseBodyData) SetDrdsInstanceIdList(v *CreatePolarxInstanceResponseBodyDataDrdsInstanceIdList) *CreatePolarxInstanceResponseBodyData {
	s.DrdsInstanceIdList = v
	return s
}

type CreatePolarxInstanceResponseBodyDataDrdsInstanceIdList struct {
	DrdsInstanceIdList []*string `json:"drdsInstanceIdList,omitempty" xml:"drdsInstanceIdList,omitempty" type:"Repeated"`
}

func (s CreatePolarxInstanceResponseBodyDataDrdsInstanceIdList) String() string {
	return tea.Prettify(s)
}

func (s CreatePolarxInstanceResponseBodyDataDrdsInstanceIdList) GoString() string {
	return s.String()
}

func (s *CreatePolarxInstanceResponseBodyDataDrdsInstanceIdList) SetDrdsInstanceIdList(v []*string) *CreatePolarxInstanceResponseBodyDataDrdsInstanceIdList {
	s.DrdsInstanceIdList = v
	return s
}

type CreatePolarxInstanceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreatePolarxInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePolarxInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePolarxInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreatePolarxInstanceResponse) SetHeaders(v map[string]*string) *CreatePolarxInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreatePolarxInstanceResponse) SetBody(v *CreatePolarxInstanceResponseBody) *CreatePolarxInstanceResponse {
	s.Body = v
	return s
}

type CreatePolarxOrderRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	NodeCount      *string `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
}

func (s CreatePolarxOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePolarxOrderRequest) GoString() string {
	return s.String()
}

func (s *CreatePolarxOrderRequest) SetRegionId(v string) *CreatePolarxOrderRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePolarxOrderRequest) SetDBInstanceName(v string) *CreatePolarxOrderRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreatePolarxOrderRequest) SetNodeCount(v string) *CreatePolarxOrderRequest {
	s.NodeCount = &v
	return s
}

type CreatePolarxOrderResponseBody struct {
	RequestId       *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderResultList []*CreatePolarxOrderResponseBodyOrderResultList `json:"OrderResultList,omitempty" xml:"OrderResultList,omitempty" type:"Repeated"`
}

func (s CreatePolarxOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePolarxOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolarxOrderResponseBody) SetRequestId(v string) *CreatePolarxOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePolarxOrderResponseBody) SetOrderResultList(v []*CreatePolarxOrderResponseBodyOrderResultList) *CreatePolarxOrderResponseBody {
	s.OrderResultList = v
	return s
}

type CreatePolarxOrderResponseBodyOrderResultList struct {
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	OrderId        *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreatePolarxOrderResponseBodyOrderResultList) String() string {
	return tea.Prettify(s)
}

func (s CreatePolarxOrderResponseBodyOrderResultList) GoString() string {
	return s.String()
}

func (s *CreatePolarxOrderResponseBodyOrderResultList) SetDBInstanceName(v string) *CreatePolarxOrderResponseBodyOrderResultList {
	s.DBInstanceName = &v
	return s
}

func (s *CreatePolarxOrderResponseBodyOrderResultList) SetOrderId(v int64) *CreatePolarxOrderResponseBodyOrderResultList {
	s.OrderId = &v
	return s
}

type CreatePolarxOrderResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreatePolarxOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePolarxOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePolarxOrderResponse) GoString() string {
	return s.String()
}

func (s *CreatePolarxOrderResponse) SetHeaders(v map[string]*string) *CreatePolarxOrderResponse {
	s.Headers = v
	return s
}

func (s *CreatePolarxOrderResponse) SetBody(v *CreatePolarxOrderResponseBody) *CreatePolarxOrderResponse {
	s.Body = v
	return s
}

type CreateSuperAccountRequest struct {
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName     *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	AccountName        *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPassword    *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
}

func (s CreateSuperAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSuperAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateSuperAccountRequest) SetRegionId(v string) *CreateSuperAccountRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSuperAccountRequest) SetDBInstanceName(v string) *CreateSuperAccountRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateSuperAccountRequest) SetAccountName(v string) *CreateSuperAccountRequest {
	s.AccountName = &v
	return s
}

func (s *CreateSuperAccountRequest) SetAccountPassword(v string) *CreateSuperAccountRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateSuperAccountRequest) SetAccountDescription(v string) *CreateSuperAccountRequest {
	s.AccountDescription = &v
	return s
}

type CreateSuperAccountResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSuperAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSuperAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSuperAccountResponseBody) SetMessage(v string) *CreateSuperAccountResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSuperAccountResponseBody) SetRequestId(v string) *CreateSuperAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSuperAccountResponseBody) SetSuccess(v bool) *CreateSuperAccountResponseBody {
	s.Success = &v
	return s
}

type CreateSuperAccountResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSuperAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSuperAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSuperAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateSuperAccountResponse) SetHeaders(v map[string]*string) *CreateSuperAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateSuperAccountResponse) SetBody(v *CreateSuperAccountResponseBody) *CreateSuperAccountResponse {
	s.Body = v
	return s
}

type DeleteAccountRequest struct {
	DBInstanceName          *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	AccountName             *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityAccountName     *string `json:"SecurityAccountName,omitempty" xml:"SecurityAccountName,omitempty"`
	SecurityAccountPassword *string `json:"SecurityAccountPassword,omitempty" xml:"SecurityAccountPassword,omitempty"`
}

func (s DeleteAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccountRequest) SetDBInstanceName(v string) *DeleteAccountRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DeleteAccountRequest) SetAccountName(v string) *DeleteAccountRequest {
	s.AccountName = &v
	return s
}

func (s *DeleteAccountRequest) SetRegionId(v string) *DeleteAccountRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAccountRequest) SetSecurityAccountName(v string) *DeleteAccountRequest {
	s.SecurityAccountName = &v
	return s
}

func (s *DeleteAccountRequest) SetSecurityAccountPassword(v string) *DeleteAccountRequest {
	s.SecurityAccountPassword = &v
	return s
}

type DeleteAccountResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccountResponseBody) SetMessage(v string) *DeleteAccountResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAccountResponseBody) SetRequestId(v string) *DeleteAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAccountResponseBody) SetSuccess(v bool) *DeleteAccountResponseBody {
	s.Success = &v
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

type DeleteDBRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
}

func (s DeleteDBRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBRequest) SetRegionId(v string) *DeleteDBRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDBRequest) SetDBInstanceName(v string) *DeleteDBRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DeleteDBRequest) SetDbName(v string) *DeleteDBRequest {
	s.DbName = &v
	return s
}

type DeleteDBResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBResponseBody) SetMessage(v string) *DeleteDBResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDBResponseBody) SetRequestId(v string) *DeleteDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBResponseBody) SetSuccess(v bool) *DeleteDBResponseBody {
	s.Success = &v
	return s
}

type DeleteDBResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDBResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDBResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBResponse) SetHeaders(v map[string]*string) *DeleteDBResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBResponse) SetBody(v *DeleteDBResponseBody) *DeleteDBResponse {
	s.Body = v
	return s
}

type DeleteDBInstanceRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
}

func (s DeleteDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceRequest) SetRegionId(v string) *DeleteDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetDBInstanceName(v string) *DeleteDBInstanceRequest {
	s.DBInstanceName = &v
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

type DescribeAccountListRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	AccountName    *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountType    *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s DescribeAccountListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountListRequest) SetRegionId(v string) *DescribeAccountListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAccountListRequest) SetDBInstanceName(v string) *DescribeAccountListRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeAccountListRequest) SetAccountName(v string) *DescribeAccountListRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountListRequest) SetAccountType(v string) *DescribeAccountListRequest {
	s.AccountType = &v
	return s
}

type DescribeAccountListResponseBody struct {
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      []*DescribeAccountListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s DescribeAccountListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountListResponseBody) SetMessage(v string) *DescribeAccountListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAccountListResponseBody) SetRequestId(v string) *DescribeAccountListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountListResponseBody) SetSuccess(v bool) *DescribeAccountListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAccountListResponseBody) SetData(v []*DescribeAccountListResponseBodyData) *DescribeAccountListResponseBody {
	s.Data = v
	return s
}

type DescribeAccountListResponseBodyData struct {
	GmtCreated         *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	DBInstanceName     *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	DBName             *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	AccountPrivilege   *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	AccountType        *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	AccountName        *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s DescribeAccountListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAccountListResponseBodyData) SetGmtCreated(v string) *DescribeAccountListResponseBodyData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeAccountListResponseBodyData) SetDBInstanceName(v string) *DescribeAccountListResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeAccountListResponseBodyData) SetAccountDescription(v string) *DescribeAccountListResponseBodyData {
	s.AccountDescription = &v
	return s
}

func (s *DescribeAccountListResponseBodyData) SetDBName(v string) *DescribeAccountListResponseBodyData {
	s.DBName = &v
	return s
}

func (s *DescribeAccountListResponseBodyData) SetAccountPrivilege(v string) *DescribeAccountListResponseBodyData {
	s.AccountPrivilege = &v
	return s
}

func (s *DescribeAccountListResponseBodyData) SetAccountType(v string) *DescribeAccountListResponseBodyData {
	s.AccountType = &v
	return s
}

func (s *DescribeAccountListResponseBodyData) SetAccountName(v string) *DescribeAccountListResponseBodyData {
	s.AccountName = &v
	return s
}

type DescribeAccountListResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAccountListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAccountListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountListResponse) SetHeaders(v map[string]*string) *DescribeAccountListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountListResponse) SetBody(v *DescribeAccountListResponseBody) *DescribeAccountListResponse {
	s.Body = v
	return s
}

type DescribeBackupPolicyRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
}

func (s DescribeBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyRequest) SetRegionId(v string) *DescribeBackupPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetDBInstanceName(v string) *DescribeBackupPolicyRequest {
	s.DBInstanceName = &v
	return s
}

type DescribeBackupPolicyResponseBody struct {
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      []*DescribeBackupPolicyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s DescribeBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBody) SetMessage(v string) *DescribeBackupPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetRequestId(v string) *DescribeBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetSuccess(v bool) *DescribeBackupPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetData(v []*DescribeBackupPolicyResponseBodyData) *DescribeBackupPolicyResponseBody {
	s.Data = v
	return s
}

type DescribeBackupPolicyResponseBodyData struct {
	LogLocalRetentionSpace     *int32  `json:"LogLocalRetentionSpace,omitempty" xml:"LogLocalRetentionSpace,omitempty"`
	DBInstanceName             *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	BackupWay                  *string `json:"BackupWay,omitempty" xml:"BackupWay,omitempty"`
	BackupPeriod               *string `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	ForceCleanOnHighSpaceUsage *int32  `json:"ForceCleanOnHighSpaceUsage,omitempty" xml:"ForceCleanOnHighSpaceUsage,omitempty"`
	BackupType                 *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	LocalLogRetention          *int32  `json:"LocalLogRetention,omitempty" xml:"LocalLogRetention,omitempty"`
	RemoveLogRetention         *int32  `json:"RemoveLogRetention,omitempty" xml:"RemoveLogRetention,omitempty"`
	BackupPlanBegin            *string `json:"BackupPlanBegin,omitempty" xml:"BackupPlanBegin,omitempty"`
	BackupSetRetention         *int32  `json:"BackupSetRetention,omitempty" xml:"BackupSetRetention,omitempty"`
	IsEnabled                  *int32  `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
}

func (s DescribeBackupPolicyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBodyData) SetLogLocalRetentionSpace(v int32) *DescribeBackupPolicyResponseBodyData {
	s.LogLocalRetentionSpace = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetDBInstanceName(v string) *DescribeBackupPolicyResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetBackupWay(v string) *DescribeBackupPolicyResponseBodyData {
	s.BackupWay = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetBackupPeriod(v string) *DescribeBackupPolicyResponseBodyData {
	s.BackupPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetForceCleanOnHighSpaceUsage(v int32) *DescribeBackupPolicyResponseBodyData {
	s.ForceCleanOnHighSpaceUsage = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetBackupType(v string) *DescribeBackupPolicyResponseBodyData {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetLocalLogRetention(v int32) *DescribeBackupPolicyResponseBodyData {
	s.LocalLogRetention = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetRemoveLogRetention(v int32) *DescribeBackupPolicyResponseBodyData {
	s.RemoveLogRetention = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetBackupPlanBegin(v string) *DescribeBackupPolicyResponseBodyData {
	s.BackupPlanBegin = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetBackupSetRetention(v int32) *DescribeBackupPolicyResponseBodyData {
	s.BackupSetRetention = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetIsEnabled(v int32) *DescribeBackupPolicyResponseBodyData {
	s.IsEnabled = &v
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

type DescribeBackupSetListRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	StartTime      *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeBackupSetListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetListRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetListRequest) SetRegionId(v string) *DescribeBackupSetListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBackupSetListRequest) SetDBInstanceName(v string) *DescribeBackupSetListRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeBackupSetListRequest) SetStartTime(v int64) *DescribeBackupSetListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupSetListRequest) SetEndTime(v int64) *DescribeBackupSetListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupSetListRequest) SetPageSize(v int32) *DescribeBackupSetListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupSetListRequest) SetPageNumber(v int32) *DescribeBackupSetListRequest {
	s.PageNumber = &v
	return s
}

type DescribeBackupSetListResponseBody struct {
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      []*DescribeBackupSetListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s DescribeBackupSetListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetListResponseBody) SetMessage(v string) *DescribeBackupSetListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeBackupSetListResponseBody) SetRequestId(v string) *DescribeBackupSetListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupSetListResponseBody) SetSuccess(v bool) *DescribeBackupSetListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupSetListResponseBody) SetData(v []*DescribeBackupSetListResponseBodyData) *DescribeBackupSetListResponseBody {
	s.Data = v
	return s
}

type DescribeBackupSetListResponseBodyData struct {
	EndTime       *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Status        *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	BackupSetSize *int64 `json:"BackupSetSize,omitempty" xml:"BackupSetSize,omitempty"`
	BackupSetId   *int64 `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	BackupType    *int32 `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	BeginTime     *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	BackupModel   *int32 `json:"BackupModel,omitempty" xml:"BackupModel,omitempty"`
}

func (s DescribeBackupSetListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetListResponseBodyData) SetEndTime(v int64) *DescribeBackupSetListResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupSetListResponseBodyData) SetStatus(v int32) *DescribeBackupSetListResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeBackupSetListResponseBodyData) SetBackupSetSize(v int64) *DescribeBackupSetListResponseBodyData {
	s.BackupSetSize = &v
	return s
}

func (s *DescribeBackupSetListResponseBodyData) SetBackupSetId(v int64) *DescribeBackupSetListResponseBodyData {
	s.BackupSetId = &v
	return s
}

func (s *DescribeBackupSetListResponseBodyData) SetBackupType(v int32) *DescribeBackupSetListResponseBodyData {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupSetListResponseBodyData) SetBeginTime(v int64) *DescribeBackupSetListResponseBodyData {
	s.BeginTime = &v
	return s
}

func (s *DescribeBackupSetListResponseBodyData) SetBackupModel(v int32) *DescribeBackupSetListResponseBodyData {
	s.BackupModel = &v
	return s
}

type DescribeBackupSetListResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupSetListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupSetListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetListResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetListResponse) SetHeaders(v map[string]*string) *DescribeBackupSetListResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupSetListResponse) SetBody(v *DescribeBackupSetListResponseBody) *DescribeBackupSetListResponse {
	s.Body = v
	return s
}

type DescribeBinaryLogListRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeBinaryLogListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBinaryLogListRequest) GoString() string {
	return s.String()
}

func (s *DescribeBinaryLogListRequest) SetRegionId(v string) *DescribeBinaryLogListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBinaryLogListRequest) SetPageNumber(v int32) *DescribeBinaryLogListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBinaryLogListRequest) SetPageSize(v int32) *DescribeBinaryLogListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBinaryLogListRequest) SetDBInstanceName(v string) *DescribeBinaryLogListRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeBinaryLogListRequest) SetStartTime(v string) *DescribeBinaryLogListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBinaryLogListRequest) SetEndTime(v string) *DescribeBinaryLogListRequest {
	s.EndTime = &v
	return s
}

type DescribeBinaryLogListResponseBody struct {
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber  *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalNumber *int32                                      `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
	LogList     []*DescribeBinaryLogListResponseBodyLogList `json:"LogList,omitempty" xml:"LogList,omitempty" type:"Repeated"`
}

func (s DescribeBinaryLogListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBinaryLogListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBinaryLogListResponseBody) SetRequestId(v string) *DescribeBinaryLogListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBinaryLogListResponseBody) SetPageNumber(v int32) *DescribeBinaryLogListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBinaryLogListResponseBody) SetPageSize(v int32) *DescribeBinaryLogListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBinaryLogListResponseBody) SetTotalNumber(v int32) *DescribeBinaryLogListResponseBody {
	s.TotalNumber = &v
	return s
}

func (s *DescribeBinaryLogListResponseBody) SetLogList(v []*DescribeBinaryLogListResponseBodyLogList) *DescribeBinaryLogListResponseBody {
	s.LogList = v
	return s
}

type DescribeBinaryLogListResponseBodyLogList struct {
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	UploadHost   *string `json:"UploadHost,omitempty" xml:"UploadHost,omitempty"`
	UploadStatus *int32  `json:"UploadStatus,omitempty" xml:"UploadStatus,omitempty"`
	DownloadLink *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
	BeginTime    *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	LogSize      *int64  `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	CreatedTime  *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	PurgeStatus  *int32  `json:"PurgeStatus,omitempty" xml:"PurgeStatus,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeBinaryLogListResponseBodyLogList) String() string {
	return tea.Prettify(s)
}

func (s DescribeBinaryLogListResponseBodyLogList) GoString() string {
	return s.String()
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetEndTime(v string) *DescribeBinaryLogListResponseBodyLogList {
	s.EndTime = &v
	return s
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetModifiedTime(v string) *DescribeBinaryLogListResponseBodyLogList {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetUploadHost(v string) *DescribeBinaryLogListResponseBodyLogList {
	s.UploadHost = &v
	return s
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetUploadStatus(v int32) *DescribeBinaryLogListResponseBodyLogList {
	s.UploadStatus = &v
	return s
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetDownloadLink(v string) *DescribeBinaryLogListResponseBodyLogList {
	s.DownloadLink = &v
	return s
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetBeginTime(v string) *DescribeBinaryLogListResponseBodyLogList {
	s.BeginTime = &v
	return s
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetLogSize(v int64) *DescribeBinaryLogListResponseBodyLogList {
	s.LogSize = &v
	return s
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetFileName(v string) *DescribeBinaryLogListResponseBodyLogList {
	s.FileName = &v
	return s
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetCreatedTime(v string) *DescribeBinaryLogListResponseBodyLogList {
	s.CreatedTime = &v
	return s
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetPurgeStatus(v int32) *DescribeBinaryLogListResponseBodyLogList {
	s.PurgeStatus = &v
	return s
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetId(v int64) *DescribeBinaryLogListResponseBodyLogList {
	s.Id = &v
	return s
}

type DescribeBinaryLogListResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBinaryLogListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBinaryLogListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBinaryLogListResponse) GoString() string {
	return s.String()
}

func (s *DescribeBinaryLogListResponse) SetHeaders(v map[string]*string) *DescribeBinaryLogListResponse {
	s.Headers = v
	return s
}

func (s *DescribeBinaryLogListResponse) SetBody(v *DescribeBinaryLogListResponseBody) *DescribeBinaryLogListResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceAttributeRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
}

func (s DescribeDBInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeRequest) SetRegionId(v string) *DescribeDBInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetDBInstanceName(v string) *DescribeDBInstanceAttributeRequest {
	s.DBInstanceName = &v
	return s
}

type DescribeDBInstanceAttributeResponseBody struct {
	RequestId  *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DBInstance *DescribeDBInstanceAttributeResponseBodyDBInstance `json:"DBInstance,omitempty" xml:"DBInstance,omitempty" type:"Struct"`
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

func (s *DescribeDBInstanceAttributeResponseBody) SetDBInstance(v *DescribeDBInstanceAttributeResponseBodyDBInstance) *DescribeDBInstanceAttributeResponseBody {
	s.DBInstance = v
	return s
}

type DescribeDBInstanceAttributeResponseBodyDBInstance struct {
	Type                    *string                                                       `json:"Type,omitempty" xml:"Type,omitempty"`
	Status                  *string                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	RightsSeparationStatus  *string                                                       `json:"RightsSeparationStatus,omitempty" xml:"RightsSeparationStatus,omitempty"`
	DBNodeCount             *int32                                                        `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	Expired                 *string                                                       `json:"Expired,omitempty" xml:"Expired,omitempty"`
	CreateTime              *string                                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	PayType                 *string                                                       `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Port                    *string                                                       `json:"Port,omitempty" xml:"Port,omitempty"`
	LockMode                *string                                                       `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	Description             *string                                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	ConnectionString        *string                                                       `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	StorageUsed             *int64                                                        `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	ExpireDate              *string                                                       `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	CommodityCode           *string                                                       `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	MaintainStartTime       *string                                                       `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	DBInstanceType          *string                                                       `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	DBNodeClass             *string                                                       `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	LatestMinorVersion      *string                                                       `json:"LatestMinorVersion,omitempty" xml:"LatestMinorVersion,omitempty"`
	MaintainEndTime         *string                                                       `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	DBType                  *string                                                       `json:"DBType,omitempty" xml:"DBType,omitempty"`
	RightsSeparationEnabled *bool                                                         `json:"RightsSeparationEnabled,omitempty" xml:"RightsSeparationEnabled,omitempty"`
	VPCId                   *string                                                       `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	MinorVersion            *string                                                       `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	RegionId                *string                                                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Network                 *string                                                       `json:"Network,omitempty" xml:"Network,omitempty"`
	DBVersion               *string                                                       `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	VSwitchId               *string                                                       `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId                  *string                                                       `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Engine                  *string                                                       `json:"Engine,omitempty" xml:"Engine,omitempty"`
	KindCode                *int32                                                        `json:"KindCode,omitempty" xml:"KindCode,omitempty"`
	Id                      *string                                                       `json:"Id,omitempty" xml:"Id,omitempty"`
	DBNodes                 []*DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes   `json:"DBNodes,omitempty" xml:"DBNodes,omitempty" type:"Repeated"`
	ConnAddrs               []*DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs `json:"ConnAddrs,omitempty" xml:"ConnAddrs,omitempty" type:"Repeated"`
	ReadDBInstances         []*string                                                     `json:"ReadDBInstances,omitempty" xml:"ReadDBInstances,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstance) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetType(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.Type = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetStatus(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.Status = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetRightsSeparationStatus(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.RightsSeparationStatus = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetDBNodeCount(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.DBNodeCount = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetExpired(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.Expired = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetCreateTime(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetPayType(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetPort(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetLockMode(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetDescription(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.Description = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetConnectionString(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetStorageUsed(v int64) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.StorageUsed = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetExpireDate(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.ExpireDate = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetCommodityCode(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetMaintainStartTime(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetDBInstanceType(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetDBNodeClass(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetLatestMinorVersion(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.LatestMinorVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetMaintainEndTime(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetDBType(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.DBType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetRightsSeparationEnabled(v bool) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.RightsSeparationEnabled = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetVPCId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.VPCId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetMinorVersion(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.MinorVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetRegionId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetNetwork(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.Network = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetDBVersion(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetVSwitchId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetZoneId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetEngine(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetKindCode(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.KindCode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.Id = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetDBNodes(v []*DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.DBNodes = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetConnAddrs(v []*DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.ConnAddrs = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetReadDBInstances(v []*string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.ReadDBInstances = v
	return s
}

type DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes struct {
	ComputeNodeId *string `json:"ComputeNodeId,omitempty" xml:"ComputeNodeId,omitempty"`
	NodeClass     *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	DataNodeId    *string `json:"DataNodeId,omitempty" xml:"DataNodeId,omitempty"`
	ZoneId        *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) SetComputeNodeId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes {
	s.ComputeNodeId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) SetNodeClass(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes {
	s.NodeClass = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) SetDataNodeId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes {
	s.DataNodeId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) SetZoneId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) SetId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes {
	s.Id = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) SetRegionId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes {
	s.RegionId = &v
	return s
}

type DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs struct {
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VSwitchId        *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	Port             *string `json:"Port,omitempty" xml:"Port,omitempty"`
	VPCId            *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) SetType(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs {
	s.Type = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) SetVSwitchId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) SetPort(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) SetVPCId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs {
	s.VPCId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) SetConnectionString(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs {
	s.ConnectionString = &v
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

type DescribeDBInstancesRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDBInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesRequest) SetRegionId(v string) *DescribeDBInstancesRequest {
	s.RegionId = &v
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

type DescribeDBInstancesResponseBody struct {
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber  *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalNumber *int32                                        `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
	DBInstances []*DescribeDBInstancesResponseBodyDBInstances `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBody) SetRequestId(v string) *DescribeDBInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetPageNumber(v int32) *DescribeDBInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetPageSize(v int32) *DescribeDBInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetTotalNumber(v int32) *DescribeDBInstancesResponseBody {
	s.TotalNumber = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetDBInstances(v []*DescribeDBInstancesResponseBodyDBInstances) *DescribeDBInstancesResponseBody {
	s.DBInstances = v
	return s
}

type DescribeDBInstancesResponseBodyDBInstances struct {
	Type            *string                                            `json:"Type,omitempty" xml:"Type,omitempty"`
	Status          *string                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	CommodityCode   *string                                            `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ExpireTime      *string                                            `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Expired         *bool                                              `json:"Expired,omitempty" xml:"Expired,omitempty"`
	CreateTime      *string                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	PayType         *string                                            `json:"PayType,omitempty" xml:"PayType,omitempty"`
	LockReason      *string                                            `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	DBType          *string                                            `json:"DBType,omitempty" xml:"DBType,omitempty"`
	LockMode        *string                                            `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	VPCId           *string                                            `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	MinorVersion    *string                                            `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	RegionId        *string                                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Network         *string                                            `json:"Network,omitempty" xml:"Network,omitempty"`
	DBVersion       *string                                            `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	Description     *string                                            `json:"Description,omitempty" xml:"Description,omitempty"`
	NodeClass       *string                                            `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	StorageUsed     *int64                                             `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	NodeCount       *int32                                             `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	ZoneId          *string                                            `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Engine          *string                                            `json:"Engine,omitempty" xml:"Engine,omitempty"`
	Id              *string                                            `json:"Id,omitempty" xml:"Id,omitempty"`
	Nodes           []*DescribeDBInstancesResponseBodyDBInstancesNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	ReadDBInstances []*string                                          `json:"ReadDBInstances,omitempty" xml:"ReadDBInstances,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancesResponseBodyDBInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyDBInstances) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetType(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.Type = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetStatus(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.Status = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetCommodityCode(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetExpireTime(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetExpired(v bool) *DescribeDBInstancesResponseBodyDBInstances {
	s.Expired = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetCreateTime(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetPayType(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.PayType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetLockReason(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.LockReason = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetDBType(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.DBType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetLockMode(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetVPCId(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.VPCId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetMinorVersion(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.MinorVersion = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetRegionId(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetNetwork(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.Network = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetDBVersion(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetDescription(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.Description = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetNodeClass(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.NodeClass = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetStorageUsed(v int64) *DescribeDBInstancesResponseBodyDBInstances {
	s.StorageUsed = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetNodeCount(v int32) *DescribeDBInstancesResponseBodyDBInstances {
	s.NodeCount = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetZoneId(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetEngine(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetId(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.Id = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetNodes(v []*DescribeDBInstancesResponseBodyDBInstancesNodes) *DescribeDBInstancesResponseBodyDBInstances {
	s.Nodes = v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetReadDBInstances(v []*string) *DescribeDBInstancesResponseBodyDBInstances {
	s.ReadDBInstances = v
	return s
}

type DescribeDBInstancesResponseBodyDBInstancesNodes struct {
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBInstancesResponseBodyDBInstancesNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyDBInstancesNodes) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyDBInstancesNodes) SetZoneId(v string) *DescribeDBInstancesResponseBodyDBInstancesNodes {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesNodes) SetId(v string) *DescribeDBInstancesResponseBodyDBInstancesNodes {
	s.Id = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesNodes) SetClassCode(v string) *DescribeDBInstancesResponseBodyDBInstancesNodes {
	s.ClassCode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesNodes) SetRegionId(v string) *DescribeDBInstancesResponseBodyDBInstancesNodes {
	s.RegionId = &v
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
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBInstanceSSLRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceSSLRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSSLRequest) SetDBInstanceName(v string) *DescribeDBInstanceSSLRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstanceSSLRequest) SetRegionId(v string) *DescribeDBInstanceSSLRequest {
	s.RegionId = &v
	return s
}

type DescribeDBInstanceSSLResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DescribeDBInstanceSSLResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeDBInstanceSSLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceSSLResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSSLResponseBody) SetRequestId(v string) *DescribeDBInstanceSSLResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetData(v *DescribeDBInstanceSSLResponseBodyData) *DescribeDBInstanceSSLResponseBody {
	s.Data = v
	return s
}

type DescribeDBInstanceSSLResponseBodyData struct {
	SSLEnabled     *bool   `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
	SSLExpiredTime *string `json:"SSLExpiredTime,omitempty" xml:"SSLExpiredTime,omitempty"`
	CertCommonName *string `json:"CertCommonName,omitempty" xml:"CertCommonName,omitempty"`
}

func (s DescribeDBInstanceSSLResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceSSLResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSSLResponseBodyData) SetSSLEnabled(v bool) *DescribeDBInstanceSSLResponseBodyData {
	s.SSLEnabled = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBodyData) SetSSLExpiredTime(v string) *DescribeDBInstanceSSLResponseBodyData {
	s.SSLExpiredTime = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBodyData) SetCertCommonName(v string) *DescribeDBInstanceSSLResponseBodyData {
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

type DescribeDBInstanceTDERequest struct {
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBInstanceTDERequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceTDERequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTDERequest) SetDBInstanceName(v string) *DescribeDBInstanceTDERequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstanceTDERequest) SetRegionId(v string) *DescribeDBInstanceTDERequest {
	s.RegionId = &v
	return s
}

type DescribeDBInstanceTDEResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DescribeDBInstanceTDEResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeDBInstanceTDEResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceTDEResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTDEResponseBody) SetRequestId(v string) *DescribeDBInstanceTDEResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceTDEResponseBody) SetData(v *DescribeDBInstanceTDEResponseBodyData) *DescribeDBInstanceTDEResponseBody {
	s.Data = v
	return s
}

type DescribeDBInstanceTDEResponseBodyData struct {
	TDEStatus *string `json:"TDEStatus,omitempty" xml:"TDEStatus,omitempty"`
}

func (s DescribeDBInstanceTDEResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceTDEResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTDEResponseBodyData) SetTDEStatus(v string) *DescribeDBInstanceTDEResponseBodyData {
	s.TDEStatus = &v
	return s
}

type DescribeDBInstanceTDEResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBInstanceTDEResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceTDEResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceTDEResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTDEResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceTDEResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceTDEResponse) SetBody(v *DescribeDBInstanceTDEResponseBody) *DescribeDBInstanceTDEResponse {
	s.Body = v
	return s
}

type DescribeDbListRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DBName         *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
}

func (s DescribeDbListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDbListRequest) SetRegionId(v string) *DescribeDbListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDbListRequest) SetDBInstanceName(v string) *DescribeDbListRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDbListRequest) SetDBName(v string) *DescribeDbListRequest {
	s.DBName = &v
	return s
}

type DescribeDbListResponseBody struct {
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      []*DescribeDbListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s DescribeDbListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDbListResponseBody) SetMessage(v string) *DescribeDbListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDbListResponseBody) SetRequestId(v string) *DescribeDbListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDbListResponseBody) SetSuccess(v bool) *DescribeDbListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDbListResponseBody) SetData(v []*DescribeDbListResponseBodyData) *DescribeDbListResponseBody {
	s.Data = v
	return s
}

type DescribeDbListResponseBodyData struct {
	DBName           *string                                   `json:"DBName,omitempty" xml:"DBName,omitempty"`
	DBDescription    *string                                   `json:"DBDescription,omitempty" xml:"DBDescription,omitempty"`
	DBInstanceName   *string                                   `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	CharacterSetName *string                                   `json:"CharacterSetName,omitempty" xml:"CharacterSetName,omitempty"`
	Accounts         []*DescribeDbListResponseBodyDataAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
}

func (s DescribeDbListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDbListResponseBodyData) SetDBName(v string) *DescribeDbListResponseBodyData {
	s.DBName = &v
	return s
}

func (s *DescribeDbListResponseBodyData) SetDBDescription(v string) *DescribeDbListResponseBodyData {
	s.DBDescription = &v
	return s
}

func (s *DescribeDbListResponseBodyData) SetDBInstanceName(v string) *DescribeDbListResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDbListResponseBodyData) SetCharacterSetName(v string) *DescribeDbListResponseBodyData {
	s.CharacterSetName = &v
	return s
}

func (s *DescribeDbListResponseBodyData) SetAccounts(v []*DescribeDbListResponseBodyDataAccounts) *DescribeDbListResponseBodyData {
	s.Accounts = v
	return s
}

type DescribeDbListResponseBodyDataAccounts struct {
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	AccountName      *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s DescribeDbListResponseBodyDataAccounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbListResponseBodyDataAccounts) GoString() string {
	return s.String()
}

func (s *DescribeDbListResponseBodyDataAccounts) SetAccountPrivilege(v string) *DescribeDbListResponseBodyDataAccounts {
	s.AccountPrivilege = &v
	return s
}

func (s *DescribeDbListResponseBodyDataAccounts) SetAccountName(v string) *DescribeDbListResponseBodyDataAccounts {
	s.AccountName = &v
	return s
}

type DescribeDbListResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDbListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDbListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDbListResponse) SetHeaders(v map[string]*string) *DescribeDbListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDbListResponse) SetBody(v *DescribeDbListResponseBody) *DescribeDbListResponse {
	s.Body = v
	return s
}

type DescribeDistributeTableListRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
}

func (s DescribeDistributeTableListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDistributeTableListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDistributeTableListRequest) SetRegionId(v string) *DescribeDistributeTableListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDistributeTableListRequest) SetDBInstanceName(v string) *DescribeDistributeTableListRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDistributeTableListRequest) SetDbName(v string) *DescribeDistributeTableListRequest {
	s.DbName = &v
	return s
}

type DescribeDistributeTableListResponseBody struct {
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      *DescribeDistributeTableListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeDistributeTableListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDistributeTableListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDistributeTableListResponseBody) SetMessage(v string) *DescribeDistributeTableListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDistributeTableListResponseBody) SetRequestId(v string) *DescribeDistributeTableListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDistributeTableListResponseBody) SetSuccess(v bool) *DescribeDistributeTableListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDistributeTableListResponseBody) SetData(v *DescribeDistributeTableListResponseBodyData) *DescribeDistributeTableListResponseBody {
	s.Data = v
	return s
}

type DescribeDistributeTableListResponseBodyData struct {
	Tables []*DescribeDistributeTableListResponseBodyDataTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
}

func (s DescribeDistributeTableListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDistributeTableListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDistributeTableListResponseBodyData) SetTables(v []*DescribeDistributeTableListResponseBodyDataTables) *DescribeDistributeTableListResponseBodyData {
	s.Tables = v
	return s
}

type DescribeDistributeTableListResponseBodyDataTables struct {
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TableType *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
	TbKey     *string `json:"TbKey,omitempty" xml:"TbKey,omitempty"`
	DbKey     *string `json:"DbKey,omitempty" xml:"DbKey,omitempty"`
}

func (s DescribeDistributeTableListResponseBodyDataTables) String() string {
	return tea.Prettify(s)
}

func (s DescribeDistributeTableListResponseBodyDataTables) GoString() string {
	return s.String()
}

func (s *DescribeDistributeTableListResponseBodyDataTables) SetTableName(v string) *DescribeDistributeTableListResponseBodyDataTables {
	s.TableName = &v
	return s
}

func (s *DescribeDistributeTableListResponseBodyDataTables) SetTableType(v string) *DescribeDistributeTableListResponseBodyDataTables {
	s.TableType = &v
	return s
}

func (s *DescribeDistributeTableListResponseBodyDataTables) SetTbKey(v string) *DescribeDistributeTableListResponseBodyDataTables {
	s.TbKey = &v
	return s
}

func (s *DescribeDistributeTableListResponseBodyDataTables) SetDbKey(v string) *DescribeDistributeTableListResponseBodyDataTables {
	s.DbKey = &v
	return s
}

type DescribeDistributeTableListResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDistributeTableListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDistributeTableListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDistributeTableListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDistributeTableListResponse) SetHeaders(v map[string]*string) *DescribeDistributeTableListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDistributeTableListResponse) SetBody(v *DescribeDistributeTableListResponseBody) *DescribeDistributeTableListResponse {
	s.Body = v
	return s
}

type DescribeInstanceDbPerformanceRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DbInstanceName *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
	Keys           *string `json:"Keys,omitempty" xml:"Keys,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
}

func (s DescribeInstanceDbPerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceDbPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDbPerformanceRequest) SetRegionId(v string) *DescribeInstanceDbPerformanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceDbPerformanceRequest) SetDbInstanceName(v string) *DescribeInstanceDbPerformanceRequest {
	s.DbInstanceName = &v
	return s
}

func (s *DescribeInstanceDbPerformanceRequest) SetKeys(v string) *DescribeInstanceDbPerformanceRequest {
	s.Keys = &v
	return s
}

func (s *DescribeInstanceDbPerformanceRequest) SetStartTime(v string) *DescribeInstanceDbPerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceDbPerformanceRequest) SetEndTime(v string) *DescribeInstanceDbPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInstanceDbPerformanceRequest) SetDbName(v string) *DescribeInstanceDbPerformanceRequest {
	s.DbName = &v
	return s
}

type DescribeInstanceDbPerformanceResponseBody struct {
	Message   *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      *DescribeInstanceDbPerformanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeInstanceDbPerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceDbPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDbPerformanceResponseBody) SetMessage(v string) *DescribeInstanceDbPerformanceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeInstanceDbPerformanceResponseBody) SetRequestId(v string) *DescribeInstanceDbPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceDbPerformanceResponseBody) SetSuccess(v bool) *DescribeInstanceDbPerformanceResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstanceDbPerformanceResponseBody) SetData(v *DescribeInstanceDbPerformanceResponseBodyData) *DescribeInstanceDbPerformanceResponseBody {
	s.Data = v
	return s
}

type DescribeInstanceDbPerformanceResponseBodyData struct {
	PerformanceItems []*DescribeInstanceDbPerformanceResponseBodyDataPerformanceItems `json:"PerformanceItems,omitempty" xml:"PerformanceItems,omitempty" type:"Repeated"`
}

func (s DescribeInstanceDbPerformanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceDbPerformanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDbPerformanceResponseBodyData) SetPerformanceItems(v []*DescribeInstanceDbPerformanceResponseBodyDataPerformanceItems) *DescribeInstanceDbPerformanceResponseBodyData {
	s.PerformanceItems = v
	return s
}

type DescribeInstanceDbPerformanceResponseBodyDataPerformanceItems struct {
	MetricName  *string                                                                `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Measurement *string                                                                `json:"Measurement,omitempty" xml:"Measurement,omitempty"`
	Points      []*DescribeInstanceDbPerformanceResponseBodyDataPerformanceItemsPoints `json:"Points,omitempty" xml:"Points,omitempty" type:"Repeated"`
}

func (s DescribeInstanceDbPerformanceResponseBodyDataPerformanceItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceDbPerformanceResponseBodyDataPerformanceItems) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDbPerformanceResponseBodyDataPerformanceItems) SetMetricName(v string) *DescribeInstanceDbPerformanceResponseBodyDataPerformanceItems {
	s.MetricName = &v
	return s
}

func (s *DescribeInstanceDbPerformanceResponseBodyDataPerformanceItems) SetMeasurement(v string) *DescribeInstanceDbPerformanceResponseBodyDataPerformanceItems {
	s.Measurement = &v
	return s
}

func (s *DescribeInstanceDbPerformanceResponseBodyDataPerformanceItems) SetPoints(v []*DescribeInstanceDbPerformanceResponseBodyDataPerformanceItemsPoints) *DescribeInstanceDbPerformanceResponseBodyDataPerformanceItems {
	s.Points = v
	return s
}

type DescribeInstanceDbPerformanceResponseBodyDataPerformanceItemsPoints struct {
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Timestamp *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeInstanceDbPerformanceResponseBodyDataPerformanceItemsPoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceDbPerformanceResponseBodyDataPerformanceItemsPoints) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDbPerformanceResponseBodyDataPerformanceItemsPoints) SetValue(v string) *DescribeInstanceDbPerformanceResponseBodyDataPerformanceItemsPoints {
	s.Value = &v
	return s
}

func (s *DescribeInstanceDbPerformanceResponseBodyDataPerformanceItemsPoints) SetTimestamp(v int64) *DescribeInstanceDbPerformanceResponseBodyDataPerformanceItemsPoints {
	s.Timestamp = &v
	return s
}

type DescribeInstanceDbPerformanceResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceDbPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceDbPerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceDbPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDbPerformanceResponse) SetHeaders(v map[string]*string) *DescribeInstanceDbPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceDbPerformanceResponse) SetBody(v *DescribeInstanceDbPerformanceResponseBody) *DescribeInstanceDbPerformanceResponse {
	s.Body = v
	return s
}

type DescribeInstancePerformanceRequest struct {
	DbInstanceName *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
	NodeId         *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Keys           *string `json:"Keys,omitempty" xml:"Keys,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeInstancePerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancePerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancePerformanceRequest) SetDbInstanceName(v string) *DescribeInstancePerformanceRequest {
	s.DbInstanceName = &v
	return s
}

func (s *DescribeInstancePerformanceRequest) SetNodeId(v string) *DescribeInstancePerformanceRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeInstancePerformanceRequest) SetKeys(v string) *DescribeInstancePerformanceRequest {
	s.Keys = &v
	return s
}

func (s *DescribeInstancePerformanceRequest) SetStartTime(v string) *DescribeInstancePerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInstancePerformanceRequest) SetEndTime(v string) *DescribeInstancePerformanceRequest {
	s.EndTime = &v
	return s
}

type DescribeInstancePerformanceResponseBody struct {
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      *DescribeInstancePerformanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeInstancePerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancePerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancePerformanceResponseBody) SetMessage(v string) *DescribeInstancePerformanceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeInstancePerformanceResponseBody) SetRequestId(v string) *DescribeInstancePerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancePerformanceResponseBody) SetSuccess(v bool) *DescribeInstancePerformanceResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstancePerformanceResponseBody) SetData(v *DescribeInstancePerformanceResponseBodyData) *DescribeInstancePerformanceResponseBody {
	s.Data = v
	return s
}

type DescribeInstancePerformanceResponseBodyData struct {
	PerformanceItems []*DescribeInstancePerformanceResponseBodyDataPerformanceItems `json:"PerformanceItems,omitempty" xml:"PerformanceItems,omitempty" type:"Repeated"`
}

func (s DescribeInstancePerformanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancePerformanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeInstancePerformanceResponseBodyData) SetPerformanceItems(v []*DescribeInstancePerformanceResponseBodyDataPerformanceItems) *DescribeInstancePerformanceResponseBodyData {
	s.PerformanceItems = v
	return s
}

type DescribeInstancePerformanceResponseBodyDataPerformanceItems struct {
	MetricName  *string                                                              `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Measurement *string                                                              `json:"Measurement,omitempty" xml:"Measurement,omitempty"`
	Points      []*DescribeInstancePerformanceResponseBodyDataPerformanceItemsPoints `json:"Points,omitempty" xml:"Points,omitempty" type:"Repeated"`
}

func (s DescribeInstancePerformanceResponseBodyDataPerformanceItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancePerformanceResponseBodyDataPerformanceItems) GoString() string {
	return s.String()
}

func (s *DescribeInstancePerformanceResponseBodyDataPerformanceItems) SetMetricName(v string) *DescribeInstancePerformanceResponseBodyDataPerformanceItems {
	s.MetricName = &v
	return s
}

func (s *DescribeInstancePerformanceResponseBodyDataPerformanceItems) SetMeasurement(v string) *DescribeInstancePerformanceResponseBodyDataPerformanceItems {
	s.Measurement = &v
	return s
}

func (s *DescribeInstancePerformanceResponseBodyDataPerformanceItems) SetPoints(v []*DescribeInstancePerformanceResponseBodyDataPerformanceItemsPoints) *DescribeInstancePerformanceResponseBodyDataPerformanceItems {
	s.Points = v
	return s
}

type DescribeInstancePerformanceResponseBodyDataPerformanceItemsPoints struct {
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Timestamp *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeInstancePerformanceResponseBodyDataPerformanceItemsPoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancePerformanceResponseBodyDataPerformanceItemsPoints) GoString() string {
	return s.String()
}

func (s *DescribeInstancePerformanceResponseBodyDataPerformanceItemsPoints) SetValue(v string) *DescribeInstancePerformanceResponseBodyDataPerformanceItemsPoints {
	s.Value = &v
	return s
}

func (s *DescribeInstancePerformanceResponseBodyDataPerformanceItemsPoints) SetTimestamp(v int64) *DescribeInstancePerformanceResponseBodyDataPerformanceItemsPoints {
	s.Timestamp = &v
	return s
}

type DescribeInstancePerformanceResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstancePerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstancePerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancePerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstancePerformanceResponse) SetHeaders(v map[string]*string) *DescribeInstancePerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstancePerformanceResponse) SetBody(v *DescribeInstancePerformanceResponseBody) *DescribeInstancePerformanceResponse {
	s.Body = v
	return s
}

type DescribeInstanceStoragePerformanceRequest struct {
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DbInstanceName    *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
	StorageInstanceId *string `json:"StorageInstanceId,omitempty" xml:"StorageInstanceId,omitempty"`
	Keys              *string `json:"Keys,omitempty" xml:"Keys,omitempty"`
	StartTime         *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime           *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeInstanceStoragePerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStoragePerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStoragePerformanceRequest) SetRegionId(v string) *DescribeInstanceStoragePerformanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceStoragePerformanceRequest) SetDbInstanceName(v string) *DescribeInstanceStoragePerformanceRequest {
	s.DbInstanceName = &v
	return s
}

func (s *DescribeInstanceStoragePerformanceRequest) SetStorageInstanceId(v string) *DescribeInstanceStoragePerformanceRequest {
	s.StorageInstanceId = &v
	return s
}

func (s *DescribeInstanceStoragePerformanceRequest) SetKeys(v string) *DescribeInstanceStoragePerformanceRequest {
	s.Keys = &v
	return s
}

func (s *DescribeInstanceStoragePerformanceRequest) SetStartTime(v string) *DescribeInstanceStoragePerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceStoragePerformanceRequest) SetEndTime(v string) *DescribeInstanceStoragePerformanceRequest {
	s.EndTime = &v
	return s
}

type DescribeInstanceStoragePerformanceResponseBody struct {
	Message   *string                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      *DescribeInstanceStoragePerformanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeInstanceStoragePerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStoragePerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStoragePerformanceResponseBody) SetMessage(v string) *DescribeInstanceStoragePerformanceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeInstanceStoragePerformanceResponseBody) SetRequestId(v string) *DescribeInstanceStoragePerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceStoragePerformanceResponseBody) SetSuccess(v bool) *DescribeInstanceStoragePerformanceResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstanceStoragePerformanceResponseBody) SetData(v *DescribeInstanceStoragePerformanceResponseBodyData) *DescribeInstanceStoragePerformanceResponseBody {
	s.Data = v
	return s
}

type DescribeInstanceStoragePerformanceResponseBodyData struct {
	PerformanceItems []*DescribeInstanceStoragePerformanceResponseBodyDataPerformanceItems `json:"PerformanceItems,omitempty" xml:"PerformanceItems,omitempty" type:"Repeated"`
}

func (s DescribeInstanceStoragePerformanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStoragePerformanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStoragePerformanceResponseBodyData) SetPerformanceItems(v []*DescribeInstanceStoragePerformanceResponseBodyDataPerformanceItems) *DescribeInstanceStoragePerformanceResponseBodyData {
	s.PerformanceItems = v
	return s
}

type DescribeInstanceStoragePerformanceResponseBodyDataPerformanceItems struct {
	MetricName  *string                                                                     `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Measurement *string                                                                     `json:"Measurement,omitempty" xml:"Measurement,omitempty"`
	Points      []*DescribeInstanceStoragePerformanceResponseBodyDataPerformanceItemsPoints `json:"Points,omitempty" xml:"Points,omitempty" type:"Repeated"`
}

func (s DescribeInstanceStoragePerformanceResponseBodyDataPerformanceItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStoragePerformanceResponseBodyDataPerformanceItems) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStoragePerformanceResponseBodyDataPerformanceItems) SetMetricName(v string) *DescribeInstanceStoragePerformanceResponseBodyDataPerformanceItems {
	s.MetricName = &v
	return s
}

func (s *DescribeInstanceStoragePerformanceResponseBodyDataPerformanceItems) SetMeasurement(v string) *DescribeInstanceStoragePerformanceResponseBodyDataPerformanceItems {
	s.Measurement = &v
	return s
}

func (s *DescribeInstanceStoragePerformanceResponseBodyDataPerformanceItems) SetPoints(v []*DescribeInstanceStoragePerformanceResponseBodyDataPerformanceItemsPoints) *DescribeInstanceStoragePerformanceResponseBodyDataPerformanceItems {
	s.Points = v
	return s
}

type DescribeInstanceStoragePerformanceResponseBodyDataPerformanceItemsPoints struct {
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Timestamp *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeInstanceStoragePerformanceResponseBodyDataPerformanceItemsPoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStoragePerformanceResponseBodyDataPerformanceItemsPoints) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStoragePerformanceResponseBodyDataPerformanceItemsPoints) SetValue(v string) *DescribeInstanceStoragePerformanceResponseBodyDataPerformanceItemsPoints {
	s.Value = &v
	return s
}

func (s *DescribeInstanceStoragePerformanceResponseBodyDataPerformanceItemsPoints) SetTimestamp(v int64) *DescribeInstanceStoragePerformanceResponseBodyDataPerformanceItemsPoints {
	s.Timestamp = &v
	return s
}

type DescribeInstanceStoragePerformanceResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceStoragePerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceStoragePerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStoragePerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStoragePerformanceResponse) SetHeaders(v map[string]*string) *DescribeInstanceStoragePerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceStoragePerformanceResponse) SetBody(v *DescribeInstanceStoragePerformanceResponseBody) *DescribeInstanceStoragePerformanceResponse {
	s.Body = v
	return s
}

type DescribeParametersRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ParamLevel   *string `json:"ParamLevel,omitempty" xml:"ParamLevel,omitempty"`
}

func (s DescribeParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersRequest) GoString() string {
	return s.String()
}

func (s *DescribeParametersRequest) SetRegionId(v string) *DescribeParametersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeParametersRequest) SetDBInstanceId(v string) *DescribeParametersRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeParametersRequest) SetParamLevel(v string) *DescribeParametersRequest {
	s.ParamLevel = &v
	return s
}

type DescribeParametersResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DescribeParametersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBody) SetRequestId(v string) *DescribeParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParametersResponseBody) SetData(v *DescribeParametersResponseBodyData) *DescribeParametersResponseBody {
	s.Data = v
	return s
}

type DescribeParametersResponseBodyData struct {
	EngineVersion     *string                                                `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	Engine            *string                                                `json:"Engine,omitempty" xml:"Engine,omitempty"`
	ConfigParameters  []*DescribeParametersResponseBodyDataConfigParameters  `json:"ConfigParameters,omitempty" xml:"ConfigParameters,omitempty" type:"Repeated"`
	RunningParameters []*DescribeParametersResponseBodyDataRunningParameters `json:"RunningParameters,omitempty" xml:"RunningParameters,omitempty" type:"Repeated"`
}

func (s DescribeParametersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBodyData) SetEngineVersion(v string) *DescribeParametersResponseBodyData {
	s.EngineVersion = &v
	return s
}

func (s *DescribeParametersResponseBodyData) SetEngine(v string) *DescribeParametersResponseBodyData {
	s.Engine = &v
	return s
}

func (s *DescribeParametersResponseBodyData) SetConfigParameters(v []*DescribeParametersResponseBodyDataConfigParameters) *DescribeParametersResponseBodyData {
	s.ConfigParameters = v
	return s
}

func (s *DescribeParametersResponseBodyData) SetRunningParameters(v []*DescribeParametersResponseBodyDataRunningParameters) *DescribeParametersResponseBodyData {
	s.RunningParameters = v
	return s
}

type DescribeParametersResponseBodyDataConfigParameters struct {
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	ParameterName        *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ParameterValue       *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s DescribeParametersResponseBodyDataConfigParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponseBodyDataConfigParameters) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBodyDataConfigParameters) SetParameterDescription(v string) *DescribeParametersResponseBodyDataConfigParameters {
	s.ParameterDescription = &v
	return s
}

func (s *DescribeParametersResponseBodyDataConfigParameters) SetParameterName(v string) *DescribeParametersResponseBodyDataConfigParameters {
	s.ParameterName = &v
	return s
}

func (s *DescribeParametersResponseBodyDataConfigParameters) SetParameterValue(v string) *DescribeParametersResponseBodyDataConfigParameters {
	s.ParameterValue = &v
	return s
}

type DescribeParametersResponseBodyDataRunningParameters struct {
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	ParameterName        *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ParameterValue       *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s DescribeParametersResponseBodyDataRunningParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponseBodyDataRunningParameters) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBodyDataRunningParameters) SetParameterDescription(v string) *DescribeParametersResponseBodyDataRunningParameters {
	s.ParameterDescription = &v
	return s
}

func (s *DescribeParametersResponseBodyDataRunningParameters) SetParameterName(v string) *DescribeParametersResponseBodyDataRunningParameters {
	s.ParameterName = &v
	return s
}

func (s *DescribeParametersResponseBodyDataRunningParameters) SetParameterValue(v string) *DescribeParametersResponseBodyDataRunningParameters {
	s.ParameterValue = &v
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

type DescribePolarxDbInstancesRequest struct {
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribePolarxDbInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePolarxDbInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarxDbInstancesRequest) SetDrdsInstanceId(v string) *DescribePolarxDbInstancesRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribePolarxDbInstancesRequest) SetDbName(v string) *DescribePolarxDbInstancesRequest {
	s.DbName = &v
	return s
}

func (s *DescribePolarxDbInstancesRequest) SetPageNumber(v int32) *DescribePolarxDbInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePolarxDbInstancesRequest) SetPageSize(v int32) *DescribePolarxDbInstancesRequest {
	s.PageSize = &v
	return s
}

type DescribePolarxDbInstancesResponseBody struct {
	PageSize    *string                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber  *string                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId   *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total       *string                                           `json:"Total,omitempty" xml:"Total,omitempty"`
	Success     *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
	DbInstances *DescribePolarxDbInstancesResponseBodyDbInstances `json:"DbInstances,omitempty" xml:"DbInstances,omitempty" type:"Struct"`
}

func (s DescribePolarxDbInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePolarxDbInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarxDbInstancesResponseBody) SetPageSize(v string) *DescribePolarxDbInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBody) SetPageNumber(v string) *DescribePolarxDbInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBody) SetRequestId(v string) *DescribePolarxDbInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBody) SetTotal(v string) *DescribePolarxDbInstancesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBody) SetSuccess(v bool) *DescribePolarxDbInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBody) SetDbInstances(v *DescribePolarxDbInstancesResponseBodyDbInstances) *DescribePolarxDbInstancesResponseBody {
	s.DbInstances = v
	return s
}

type DescribePolarxDbInstancesResponseBodyDbInstances struct {
	DbInstance []*DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance `json:"DbInstance,omitempty" xml:"DbInstance,omitempty" type:"Repeated"`
}

func (s DescribePolarxDbInstancesResponseBodyDbInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribePolarxDbInstancesResponseBodyDbInstances) GoString() string {
	return s.String()
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstances) SetDbInstance(v []*DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) *DescribePolarxDbInstancesResponseBodyDbInstances {
	s.DbInstance = v
	return s
}

type DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ExpireTime   *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	PayType      *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	DBType       *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	LockMode     *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	VPCId        *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Network      *string `json:"Network,omitempty" xml:"Network,omitempty"`
	DBVersion    *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	NodeClass    *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	StorageUsed  *int32  `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	NodeCount    *int32  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	ZoneId       *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Engine       *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	LockReason   *string `json:"lockReason,omitempty" xml:"lockReason,omitempty"`
	StatusDesc   *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
}

func (s DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) GoString() string {
	return s.String()
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetStatus(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.Status = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetExpireTime(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetCreateTime(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetPayType(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.PayType = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetDBType(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.DBType = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetLockMode(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.LockMode = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetVPCId(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.VPCId = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetRegionId(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.RegionId = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetNetwork(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.Network = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetDBVersion(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.DBVersion = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetDescription(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.Description = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetNodeClass(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.NodeClass = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetStorageUsed(v int32) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.StorageUsed = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetNodeCount(v int32) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.NodeCount = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetZoneId(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceId(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetEngine(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.Engine = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetLockReason(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.LockReason = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetStatusDesc(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.StatusDesc = &v
	return s
}

type DescribePolarxDbInstancesResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePolarxDbInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePolarxDbInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePolarxDbInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribePolarxDbInstancesResponse) SetHeaders(v map[string]*string) *DescribePolarxDbInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribePolarxDbInstancesResponse) SetBody(v *DescribePolarxDbInstancesResponseBody) *DescribePolarxDbInstancesResponse {
	s.Body = v
	return s
}

type DescribeRegionsResponseBody struct {
	Code      *int32                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorCode *int32                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetCode(v int32) *DescribeRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetMessage(v string) *DescribeRegionsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetSuccess(v bool) *DescribeRegionsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetErrorCode(v int32) *DescribeRegionsResponseBody {
	s.ErrorCode = &v
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
	SupportPolarx10 *bool                                          `json:"SupportPolarx10,omitempty" xml:"SupportPolarx10,omitempty"`
	SupportPolarx20 *bool                                          `json:"SupportPolarx20,omitempty" xml:"SupportPolarx20,omitempty"`
	RegionId        *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Zones           *DescribeRegionsResponseBodyRegionsRegionZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetSupportPolarx10(v bool) *DescribeRegionsResponseBodyRegionsRegion {
	s.SupportPolarx10 = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetSupportPolarx20(v bool) *DescribeRegionsResponseBodyRegionsRegion {
	s.SupportPolarx20 = &v
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
	ZoneId     *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	VpcEnabled *bool   `json:"VpcEnabled,omitempty" xml:"VpcEnabled,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegionZonesZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegionZonesZone) GoString() string {
	return s.String()
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

type DescribeScaleOutMigrateTaskListRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DBInstanceName       *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s DescribeScaleOutMigrateTaskListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScaleOutMigrateTaskListRequest) GoString() string {
	return s.String()
}

func (s *DescribeScaleOutMigrateTaskListRequest) SetOwnerId(v int64) *DescribeScaleOutMigrateTaskListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScaleOutMigrateTaskListRequest) SetResourceOwnerAccount(v string) *DescribeScaleOutMigrateTaskListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeScaleOutMigrateTaskListRequest) SetResourceOwnerId(v int64) *DescribeScaleOutMigrateTaskListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScaleOutMigrateTaskListRequest) SetDBInstanceName(v string) *DescribeScaleOutMigrateTaskListRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeScaleOutMigrateTaskListRequest) SetOwnerAccount(v string) *DescribeScaleOutMigrateTaskListRequest {
	s.OwnerAccount = &v
	return s
}

type DescribeScaleOutMigrateTaskListResponseBody struct {
	Progress  *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeScaleOutMigrateTaskListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScaleOutMigrateTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScaleOutMigrateTaskListResponseBody) SetProgress(v int32) *DescribeScaleOutMigrateTaskListResponseBody {
	s.Progress = &v
	return s
}

func (s *DescribeScaleOutMigrateTaskListResponseBody) SetRequestId(v string) *DescribeScaleOutMigrateTaskListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeScaleOutMigrateTaskListResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeScaleOutMigrateTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScaleOutMigrateTaskListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScaleOutMigrateTaskListResponse) GoString() string {
	return s.String()
}

func (s *DescribeScaleOutMigrateTaskListResponse) SetHeaders(v map[string]*string) *DescribeScaleOutMigrateTaskListResponse {
	s.Headers = v
	return s
}

func (s *DescribeScaleOutMigrateTaskListResponse) SetBody(v *DescribeScaleOutMigrateTaskListResponseBody) *DescribeScaleOutMigrateTaskListResponse {
	s.Body = v
	return s
}

type DescribeSecurityIpsRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
}

func (s DescribeSecurityIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIpsRequest) SetRegionId(v string) *DescribeSecurityIpsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSecurityIpsRequest) SetDBInstanceName(v string) *DescribeSecurityIpsRequest {
	s.DBInstanceName = &v
	return s
}

type DescribeSecurityIpsResponseBody struct {
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      *DescribeSecurityIpsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeSecurityIpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityIpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIpsResponseBody) SetMessage(v string) *DescribeSecurityIpsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSecurityIpsResponseBody) SetRequestId(v string) *DescribeSecurityIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityIpsResponseBody) SetSuccess(v bool) *DescribeSecurityIpsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSecurityIpsResponseBody) SetData(v *DescribeSecurityIpsResponseBodyData) *DescribeSecurityIpsResponseBody {
	s.Data = v
	return s
}

type DescribeSecurityIpsResponseBodyData struct {
	DBInstanceName *string                                          `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	GroupItems     []*DescribeSecurityIpsResponseBodyDataGroupItems `json:"GroupItems,omitempty" xml:"GroupItems,omitempty" type:"Repeated"`
}

func (s DescribeSecurityIpsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityIpsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIpsResponseBodyData) SetDBInstanceName(v string) *DescribeSecurityIpsResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeSecurityIpsResponseBodyData) SetGroupItems(v []*DescribeSecurityIpsResponseBodyDataGroupItems) *DescribeSecurityIpsResponseBodyData {
	s.GroupItems = v
	return s
}

type DescribeSecurityIpsResponseBodyDataGroupItems struct {
	GroupName      *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s DescribeSecurityIpsResponseBodyDataGroupItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityIpsResponseBodyDataGroupItems) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIpsResponseBodyDataGroupItems) SetGroupName(v string) *DescribeSecurityIpsResponseBodyDataGroupItems {
	s.GroupName = &v
	return s
}

func (s *DescribeSecurityIpsResponseBodyDataGroupItems) SetSecurityIPList(v string) *DescribeSecurityIpsResponseBodyDataGroupItems {
	s.SecurityIPList = &v
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

type DescribeSqlAuditInfoRequest struct {
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	AuditAccountName     *string `json:"AuditAccountName,omitempty" xml:"AuditAccountName,omitempty"`
	AuditAccountPassword *string `json:"AuditAccountPassword,omitempty" xml:"AuditAccountPassword,omitempty"`
}

func (s DescribeSqlAuditInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlAuditInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeSqlAuditInfoRequest) SetRegionId(v string) *DescribeSqlAuditInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSqlAuditInfoRequest) SetDBInstanceId(v string) *DescribeSqlAuditInfoRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSqlAuditInfoRequest) SetAuditAccountName(v string) *DescribeSqlAuditInfoRequest {
	s.AuditAccountName = &v
	return s
}

func (s *DescribeSqlAuditInfoRequest) SetAuditAccountPassword(v string) *DescribeSqlAuditInfoRequest {
	s.AuditAccountPassword = &v
	return s
}

type DescribeSqlAuditInfoResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DescribeSqlAuditInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeSqlAuditInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlAuditInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSqlAuditInfoResponseBody) SetRequestId(v string) *DescribeSqlAuditInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSqlAuditInfoResponseBody) SetData(v *DescribeSqlAuditInfoResponseBodyData) *DescribeSqlAuditInfoResponseBody {
	s.Data = v
	return s
}

type DescribeSqlAuditInfoResponseBodyData struct {
	SLSLogStore *string `json:"SLSLogStore,omitempty" xml:"SLSLogStore,omitempty"`
	SLSProject  *string `json:"SLSProject,omitempty" xml:"SLSProject,omitempty"`
	IsEnabled   *bool   `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
}

func (s DescribeSqlAuditInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlAuditInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSqlAuditInfoResponseBodyData) SetSLSLogStore(v string) *DescribeSqlAuditInfoResponseBodyData {
	s.SLSLogStore = &v
	return s
}

func (s *DescribeSqlAuditInfoResponseBodyData) SetSLSProject(v string) *DescribeSqlAuditInfoResponseBodyData {
	s.SLSProject = &v
	return s
}

func (s *DescribeSqlAuditInfoResponseBodyData) SetIsEnabled(v bool) *DescribeSqlAuditInfoResponseBodyData {
	s.IsEnabled = &v
	return s
}

type DescribeSqlAuditInfoResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSqlAuditInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSqlAuditInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlAuditInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeSqlAuditInfoResponse) SetHeaders(v map[string]*string) *DescribeSqlAuditInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeSqlAuditInfoResponse) SetBody(v *DescribeSqlAuditInfoResponseBody) *DescribeSqlAuditInfoResponse {
	s.Body = v
	return s
}

type DescribeTasksRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskAction           *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksRequest) GoString() string {
	return s.String()
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

func (s *DescribeTasksRequest) SetDBInstanceId(v string) *DescribeTasksRequest {
	s.DBInstanceId = &v
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

func (s *DescribeTasksRequest) SetPageSize(v int32) *DescribeTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTasksRequest) SetPageNumber(v int32) *DescribeTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTasksRequest) SetStatus(v string) *DescribeTasksRequest {
	s.Status = &v
	return s
}

func (s *DescribeTasksRequest) SetTaskAction(v string) *DescribeTasksRequest {
	s.TaskAction = &v
	return s
}

func (s *DescribeTasksRequest) SetOwnerAccount(v string) *DescribeTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTasksRequest) SetRegionId(v string) *DescribeTasksRequest {
	s.RegionId = &v
	return s
}

type DescribeTasksResponseBody struct {
	RequestId        *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber       *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount  *int32                            `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	TotalRecordCount *int32                            `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	Items            []*DescribeTasksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBody) SetRequestId(v string) *DescribeTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTasksResponseBody) SetPageNumber(v int32) *DescribeTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTasksResponseBody) SetPageRecordCount(v int32) *DescribeTasksResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeTasksResponseBody) SetTotalRecordCount(v int32) *DescribeTasksResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeTasksResponseBody) SetItems(v []*DescribeTasksResponseBodyItems) *DescribeTasksResponseBody {
	s.Items = v
	return s
}

type DescribeTasksResponseBodyItems struct {
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	FinishTime       *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	Progress         *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	BeginTime        *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	TaskErrorCode    *string `json:"TaskErrorCode,omitempty" xml:"TaskErrorCode,omitempty"`
	DBName           *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	ProgressInfo     *string `json:"ProgressInfo,omitempty" xml:"ProgressInfo,omitempty"`
	ScaleOutToken    *string `json:"ScaleOutToken,omitempty" xml:"ScaleOutToken,omitempty"`
	TaskId           *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskErrorMessage *string `json:"TaskErrorMessage,omitempty" xml:"TaskErrorMessage,omitempty"`
	TaskAction       *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
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

func (s *DescribeTasksResponseBodyItems) SetProgress(v string) *DescribeTasksResponseBodyItems {
	s.Progress = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetBeginTime(v string) *DescribeTasksResponseBodyItems {
	s.BeginTime = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetTaskErrorCode(v string) *DescribeTasksResponseBodyItems {
	s.TaskErrorCode = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetDBName(v string) *DescribeTasksResponseBodyItems {
	s.DBName = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetProgressInfo(v string) *DescribeTasksResponseBodyItems {
	s.ProgressInfo = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetScaleOutToken(v string) *DescribeTasksResponseBodyItems {
	s.ScaleOutToken = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetTaskId(v string) *DescribeTasksResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetTaskErrorMessage(v string) *DescribeTasksResponseBodyItems {
	s.TaskErrorMessage = &v
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

type DescribeUserEncryptionKeyListRequest struct {
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeUserEncryptionKeyListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserEncryptionKeyListRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListRequest) SetDBInstanceName(v string) *DescribeUserEncryptionKeyListRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) SetRegionId(v string) *DescribeUserEncryptionKeyListRequest {
	s.RegionId = &v
	return s
}

type DescribeUserEncryptionKeyListResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DescribeUserEncryptionKeyListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeUserEncryptionKeyListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserEncryptionKeyListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListResponseBody) SetRequestId(v string) *DescribeUserEncryptionKeyListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserEncryptionKeyListResponseBody) SetData(v *DescribeUserEncryptionKeyListResponseBodyData) *DescribeUserEncryptionKeyListResponseBody {
	s.Data = v
	return s
}

type DescribeUserEncryptionKeyListResponseBodyData struct {
	KeyIds []*string `json:"KeyIds,omitempty" xml:"KeyIds,omitempty" type:"Repeated"`
}

func (s DescribeUserEncryptionKeyListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserEncryptionKeyListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListResponseBodyData) SetKeyIds(v []*string) *DescribeUserEncryptionKeyListResponseBodyData {
	s.KeyIds = v
	return s
}

type DescribeUserEncryptionKeyListResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserEncryptionKeyListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserEncryptionKeyListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserEncryptionKeyListResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListResponse) SetHeaders(v map[string]*string) *DescribeUserEncryptionKeyListResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserEncryptionKeyListResponse) SetBody(v *DescribeUserEncryptionKeyListResponseBody) *DescribeUserEncryptionKeyListResponse {
	s.Body = v
	return s
}

type GetPolarxCommodityRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	OrderType      *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
}

func (s GetPolarxCommodityRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPolarxCommodityRequest) GoString() string {
	return s.String()
}

func (s *GetPolarxCommodityRequest) SetRegionId(v string) *GetPolarxCommodityRequest {
	s.RegionId = &v
	return s
}

func (s *GetPolarxCommodityRequest) SetDBInstanceName(v string) *GetPolarxCommodityRequest {
	s.DBInstanceName = &v
	return s
}

func (s *GetPolarxCommodityRequest) SetOrderType(v string) *GetPolarxCommodityRequest {
	s.OrderType = &v
	return s
}

type GetPolarxCommodityResponseBody struct {
	RequestId     *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DBInstance    *GetPolarxCommodityResponseBodyDBInstance      `json:"DBInstance,omitempty" xml:"DBInstance,omitempty" type:"Struct"`
	ComponentList []*GetPolarxCommodityResponseBodyComponentList `json:"ComponentList,omitempty" xml:"ComponentList,omitempty" type:"Repeated"`
}

func (s GetPolarxCommodityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPolarxCommodityResponseBody) GoString() string {
	return s.String()
}

func (s *GetPolarxCommodityResponseBody) SetRequestId(v string) *GetPolarxCommodityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPolarxCommodityResponseBody) SetDBInstance(v *GetPolarxCommodityResponseBodyDBInstance) *GetPolarxCommodityResponseBody {
	s.DBInstance = v
	return s
}

func (s *GetPolarxCommodityResponseBody) SetComponentList(v []*GetPolarxCommodityResponseBodyComponentList) *GetPolarxCommodityResponseBody {
	s.ComponentList = v
	return s
}

type GetPolarxCommodityResponseBodyDBInstance struct {
	Type               *string                                              `json:"Type,omitempty" xml:"Type,omitempty"`
	Status             *string                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	DBNodeCount        *int32                                               `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	Expired            *string                                              `json:"Expired,omitempty" xml:"Expired,omitempty"`
	CreateTime         *string                                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	PayType            *string                                              `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Port               *string                                              `json:"Port,omitempty" xml:"Port,omitempty"`
	LockMode           *string                                              `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	Description        *string                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	ConnectionString   *string                                              `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	StorageUsed        *int64                                               `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	ExpireDate         *string                                              `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	CommodityCode      *string                                              `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	MaintainStartTime  *string                                              `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	DBInstanceType     *string                                              `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	DBNodeClass        *string                                              `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	LatestMinorVersion *string                                              `json:"LatestMinorVersion,omitempty" xml:"LatestMinorVersion,omitempty"`
	MaintainEndTime    *string                                              `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	DBType             *string                                              `json:"DBType,omitempty" xml:"DBType,omitempty"`
	VPCId              *string                                              `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	MinorVersion       *string                                              `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	RegionId           *string                                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Network            *string                                              `json:"Network,omitempty" xml:"Network,omitempty"`
	DBVersion          *string                                              `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	VSwitchId          *string                                              `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId             *string                                              `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Engine             *string                                              `json:"Engine,omitempty" xml:"Engine,omitempty"`
	Id                 *string                                              `json:"Id,omitempty" xml:"Id,omitempty"`
	DBNodes            []*GetPolarxCommodityResponseBodyDBInstanceDBNodes   `json:"DBNodes,omitempty" xml:"DBNodes,omitempty" type:"Repeated"`
	ConnAddrs          []*GetPolarxCommodityResponseBodyDBInstanceConnAddrs `json:"ConnAddrs,omitempty" xml:"ConnAddrs,omitempty" type:"Repeated"`
	ReadDBInstances    []*string                                            `json:"ReadDBInstances,omitempty" xml:"ReadDBInstances,omitempty" type:"Repeated"`
}

func (s GetPolarxCommodityResponseBodyDBInstance) String() string {
	return tea.Prettify(s)
}

func (s GetPolarxCommodityResponseBodyDBInstance) GoString() string {
	return s.String()
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetType(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.Type = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetStatus(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.Status = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetDBNodeCount(v int32) *GetPolarxCommodityResponseBodyDBInstance {
	s.DBNodeCount = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetExpired(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.Expired = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetCreateTime(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.CreateTime = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetPayType(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.PayType = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetPort(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.Port = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetLockMode(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.LockMode = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetDescription(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.Description = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetConnectionString(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.ConnectionString = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetStorageUsed(v int64) *GetPolarxCommodityResponseBodyDBInstance {
	s.StorageUsed = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetExpireDate(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.ExpireDate = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetCommodityCode(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.CommodityCode = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetMaintainStartTime(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.MaintainStartTime = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetDBInstanceType(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.DBInstanceType = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetDBNodeClass(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.DBNodeClass = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetLatestMinorVersion(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.LatestMinorVersion = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetMaintainEndTime(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.MaintainEndTime = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetDBType(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.DBType = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetVPCId(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.VPCId = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetMinorVersion(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.MinorVersion = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetRegionId(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.RegionId = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetNetwork(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.Network = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetDBVersion(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.DBVersion = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetVSwitchId(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.VSwitchId = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetZoneId(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.ZoneId = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetEngine(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.Engine = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetId(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.Id = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetDBNodes(v []*GetPolarxCommodityResponseBodyDBInstanceDBNodes) *GetPolarxCommodityResponseBodyDBInstance {
	s.DBNodes = v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetConnAddrs(v []*GetPolarxCommodityResponseBodyDBInstanceConnAddrs) *GetPolarxCommodityResponseBodyDBInstance {
	s.ConnAddrs = v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetReadDBInstances(v []*string) *GetPolarxCommodityResponseBodyDBInstance {
	s.ReadDBInstances = v
	return s
}

type GetPolarxCommodityResponseBodyDBInstanceDBNodes struct {
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	NodeClass *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetPolarxCommodityResponseBodyDBInstanceDBNodes) String() string {
	return tea.Prettify(s)
}

func (s GetPolarxCommodityResponseBodyDBInstanceDBNodes) GoString() string {
	return s.String()
}

func (s *GetPolarxCommodityResponseBodyDBInstanceDBNodes) SetZoneId(v string) *GetPolarxCommodityResponseBodyDBInstanceDBNodes {
	s.ZoneId = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstanceDBNodes) SetNodeClass(v string) *GetPolarxCommodityResponseBodyDBInstanceDBNodes {
	s.NodeClass = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstanceDBNodes) SetId(v string) *GetPolarxCommodityResponseBodyDBInstanceDBNodes {
	s.Id = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstanceDBNodes) SetRegionId(v string) *GetPolarxCommodityResponseBodyDBInstanceDBNodes {
	s.RegionId = &v
	return s
}

type GetPolarxCommodityResponseBodyDBInstanceConnAddrs struct {
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VSwitchId        *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	Port             *string `json:"Port,omitempty" xml:"Port,omitempty"`
	VPCId            *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
}

func (s GetPolarxCommodityResponseBodyDBInstanceConnAddrs) String() string {
	return tea.Prettify(s)
}

func (s GetPolarxCommodityResponseBodyDBInstanceConnAddrs) GoString() string {
	return s.String()
}

func (s *GetPolarxCommodityResponseBodyDBInstanceConnAddrs) SetType(v string) *GetPolarxCommodityResponseBodyDBInstanceConnAddrs {
	s.Type = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstanceConnAddrs) SetVSwitchId(v string) *GetPolarxCommodityResponseBodyDBInstanceConnAddrs {
	s.VSwitchId = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstanceConnAddrs) SetPort(v string) *GetPolarxCommodityResponseBodyDBInstanceConnAddrs {
	s.Port = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstanceConnAddrs) SetVPCId(v string) *GetPolarxCommodityResponseBodyDBInstanceConnAddrs {
	s.VPCId = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstanceConnAddrs) SetConnectionString(v string) *GetPolarxCommodityResponseBodyDBInstanceConnAddrs {
	s.ConnectionString = &v
	return s
}

type GetPolarxCommodityResponseBodyComponentList struct {
	Type   *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	Name   *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s GetPolarxCommodityResponseBodyComponentList) String() string {
	return tea.Prettify(s)
}

func (s GetPolarxCommodityResponseBodyComponentList) GoString() string {
	return s.String()
}

func (s *GetPolarxCommodityResponseBodyComponentList) SetType(v string) *GetPolarxCommodityResponseBodyComponentList {
	s.Type = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyComponentList) SetName(v string) *GetPolarxCommodityResponseBodyComponentList {
	s.Name = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyComponentList) SetValues(v []*string) *GetPolarxCommodityResponseBodyComponentList {
	s.Values = v
	return s
}

type GetPolarxCommodityResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPolarxCommodityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPolarxCommodityResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPolarxCommodityResponse) GoString() string {
	return s.String()
}

func (s *GetPolarxCommodityResponse) SetHeaders(v map[string]*string) *GetPolarxCommodityResponse {
	s.Headers = v
	return s
}

func (s *GetPolarxCommodityResponse) SetBody(v *GetPolarxCommodityResponseBody) *GetPolarxCommodityResponse {
	s.Body = v
	return s
}

type GetPolarXPriceRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	NodeCount      *string `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
}

func (s GetPolarXPriceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPolarXPriceRequest) GoString() string {
	return s.String()
}

func (s *GetPolarXPriceRequest) SetRegionId(v string) *GetPolarXPriceRequest {
	s.RegionId = &v
	return s
}

func (s *GetPolarXPriceRequest) SetDBInstanceName(v string) *GetPolarXPriceRequest {
	s.DBInstanceName = &v
	return s
}

func (s *GetPolarXPriceRequest) SetNodeCount(v string) *GetPolarXPriceRequest {
	s.NodeCount = &v
	return s
}

type GetPolarXPriceResponseBody struct {
	RequestId      *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderPriceList []*GetPolarXPriceResponseBodyOrderPriceList `json:"OrderPriceList,omitempty" xml:"OrderPriceList,omitempty" type:"Repeated"`
}

func (s GetPolarXPriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPolarXPriceResponseBody) GoString() string {
	return s.String()
}

func (s *GetPolarXPriceResponseBody) SetRequestId(v string) *GetPolarXPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPolarXPriceResponseBody) SetOrderPriceList(v []*GetPolarXPriceResponseBodyOrderPriceList) *GetPolarXPriceResponseBody {
	s.OrderPriceList = v
	return s
}

type GetPolarXPriceResponseBodyOrderPriceList struct {
	DBInstanceName  *string                                          `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	OriginalAmount  *string                                          `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	DiscountAmount  *string                                          `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	TradeAmount     *string                                          `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
	PayType         *string                                          `json:"PayType,omitempty" xml:"PayType,omitempty"`
	TotalCostAmount *string                                          `json:"TotalCostAmount,omitempty" xml:"TotalCostAmount,omitempty"`
	Rules           []*GetPolarXPriceResponseBodyOrderPriceListRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s GetPolarXPriceResponseBodyOrderPriceList) String() string {
	return tea.Prettify(s)
}

func (s GetPolarXPriceResponseBodyOrderPriceList) GoString() string {
	return s.String()
}

func (s *GetPolarXPriceResponseBodyOrderPriceList) SetDBInstanceName(v string) *GetPolarXPriceResponseBodyOrderPriceList {
	s.DBInstanceName = &v
	return s
}

func (s *GetPolarXPriceResponseBodyOrderPriceList) SetOriginalAmount(v string) *GetPolarXPriceResponseBodyOrderPriceList {
	s.OriginalAmount = &v
	return s
}

func (s *GetPolarXPriceResponseBodyOrderPriceList) SetDiscountAmount(v string) *GetPolarXPriceResponseBodyOrderPriceList {
	s.DiscountAmount = &v
	return s
}

func (s *GetPolarXPriceResponseBodyOrderPriceList) SetTradeAmount(v string) *GetPolarXPriceResponseBodyOrderPriceList {
	s.TradeAmount = &v
	return s
}

func (s *GetPolarXPriceResponseBodyOrderPriceList) SetPayType(v string) *GetPolarXPriceResponseBodyOrderPriceList {
	s.PayType = &v
	return s
}

func (s *GetPolarXPriceResponseBodyOrderPriceList) SetTotalCostAmount(v string) *GetPolarXPriceResponseBodyOrderPriceList {
	s.TotalCostAmount = &v
	return s
}

func (s *GetPolarXPriceResponseBodyOrderPriceList) SetRules(v []*GetPolarXPriceResponseBodyOrderPriceListRules) *GetPolarXPriceResponseBodyOrderPriceList {
	s.Rules = v
	return s
}

type GetPolarXPriceResponseBodyOrderPriceListRules struct {
	RuleDescId *int64  `json:"RuleDescId,omitempty" xml:"RuleDescId,omitempty"`
	Title      *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPolarXPriceResponseBodyOrderPriceListRules) String() string {
	return tea.Prettify(s)
}

func (s GetPolarXPriceResponseBodyOrderPriceListRules) GoString() string {
	return s.String()
}

func (s *GetPolarXPriceResponseBodyOrderPriceListRules) SetRuleDescId(v int64) *GetPolarXPriceResponseBodyOrderPriceListRules {
	s.RuleDescId = &v
	return s
}

func (s *GetPolarXPriceResponseBodyOrderPriceListRules) SetTitle(v string) *GetPolarXPriceResponseBodyOrderPriceListRules {
	s.Title = &v
	return s
}

func (s *GetPolarXPriceResponseBodyOrderPriceListRules) SetName(v string) *GetPolarXPriceResponseBodyOrderPriceListRules {
	s.Name = &v
	return s
}

type GetPolarXPriceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPolarXPriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPolarXPriceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPolarXPriceResponse) GoString() string {
	return s.String()
}

func (s *GetPolarXPriceResponse) SetHeaders(v map[string]*string) *GetPolarXPriceResponse {
	s.Headers = v
	return s
}

func (s *GetPolarXPriceResponse) SetBody(v *GetPolarXPriceResponseBody) *GetPolarXPriceResponse {
	s.Body = v
	return s
}

type ModifyAccountDescriptionRequest struct {
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName     *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	AccountName        *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
}

func (s ModifyAccountDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionRequest) SetRegionId(v string) *ModifyAccountDescriptionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetDBInstanceName(v string) *ModifyAccountDescriptionRequest {
	s.DBInstanceName = &v
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
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyAccountDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionResponseBody) SetMessage(v string) *ModifyAccountDescriptionResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyAccountDescriptionResponseBody) SetRequestId(v string) *ModifyAccountDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAccountDescriptionResponseBody) SetSuccess(v bool) *ModifyAccountDescriptionResponseBody {
	s.Success = &v
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

type ModifyDatabaseDescriptionRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DbDescription  *string `json:"DbDescription,omitempty" xml:"DbDescription,omitempty"`
}

func (s ModifyDatabaseDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseDescriptionRequest) SetRegionId(v string) *ModifyDatabaseDescriptionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDatabaseDescriptionRequest) SetDBInstanceName(v string) *ModifyDatabaseDescriptionRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyDatabaseDescriptionRequest) SetDbName(v string) *ModifyDatabaseDescriptionRequest {
	s.DbName = &v
	return s
}

func (s *ModifyDatabaseDescriptionRequest) SetDbDescription(v string) *ModifyDatabaseDescriptionRequest {
	s.DbDescription = &v
	return s
}

type ModifyDatabaseDescriptionResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDatabaseDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseDescriptionResponseBody) SetMessage(v string) *ModifyDatabaseDescriptionResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyDatabaseDescriptionResponseBody) SetRequestId(v string) *ModifyDatabaseDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDatabaseDescriptionResponseBody) SetSuccess(v bool) *ModifyDatabaseDescriptionResponseBody {
	s.Success = &v
	return s
}

type ModifyDatabaseDescriptionResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDatabaseDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDatabaseDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseDescriptionResponse) SetHeaders(v map[string]*string) *ModifyDatabaseDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyDatabaseDescriptionResponse) SetBody(v *ModifyDatabaseDescriptionResponseBody) *ModifyDatabaseDescriptionResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceClassRequest struct {
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName        *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	TargetDBInstanceClass *string `json:"TargetDBInstanceClass,omitempty" xml:"TargetDBInstanceClass,omitempty"`
	ClientToken           *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s ModifyDBInstanceClassRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceClassRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceClassRequest) SetRegionId(v string) *ModifyDBInstanceClassRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetDBInstanceName(v string) *ModifyDBInstanceClassRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetTargetDBInstanceClass(v string) *ModifyDBInstanceClassRequest {
	s.TargetDBInstanceClass = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetClientToken(v string) *ModifyDBInstanceClassRequest {
	s.ClientToken = &v
	return s
}

type ModifyDBInstanceClassResponseBody struct {
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceClassResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceClassResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceClassResponseBody) SetOrderId(v string) *ModifyDBInstanceClassResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyDBInstanceClassResponseBody) SetRequestId(v string) *ModifyDBInstanceClassResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceClassResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBInstanceClassResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBInstanceClassResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceClassResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceClassResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceClassResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceClassResponse) SetBody(v *ModifyDBInstanceClassResponseBody) *ModifyDBInstanceClassResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceConfigRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	ConfigName     *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	ConfigValue    *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
}

func (s ModifyDBInstanceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConfigRequest) SetRegionId(v string) *ModifyDBInstanceConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetDBInstanceName(v string) *ModifyDBInstanceConfigRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetConfigName(v string) *ModifyDBInstanceConfigRequest {
	s.ConfigName = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetConfigValue(v string) *ModifyDBInstanceConfigRequest {
	s.ConfigValue = &v
	return s
}

type ModifyDBInstanceConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConfigResponseBody) SetRequestId(v string) *ModifyDBInstanceConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceConfigResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBInstanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBInstanceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConfigResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceConfigResponse) SetBody(v *ModifyDBInstanceConfigResponseBody) *ModifyDBInstanceConfigResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceDescriptionRequest struct {
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName        *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
}

func (s ModifyDBInstanceDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDescriptionRequest) SetRegionId(v string) *ModifyDBInstanceDescriptionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetDBInstanceName(v string) *ModifyDBInstanceDescriptionRequest {
	s.DBInstanceName = &v
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

type ModifyParameterRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ParamLevel   *string `json:"ParamLevel,omitempty" xml:"ParamLevel,omitempty"`
	Parameters   *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s ModifyParameterRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyParameterRequest) GoString() string {
	return s.String()
}

func (s *ModifyParameterRequest) SetRegionId(v string) *ModifyParameterRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyParameterRequest) SetDBInstanceId(v string) *ModifyParameterRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyParameterRequest) SetParamLevel(v string) *ModifyParameterRequest {
	s.ParamLevel = &v
	return s
}

func (s *ModifyParameterRequest) SetParameters(v string) *ModifyParameterRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyParameterRequest) SetClientToken(v string) *ModifyParameterRequest {
	s.ClientToken = &v
	return s
}

type ModifyParameterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyParameterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyParameterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyParameterResponseBody) SetRequestId(v string) *ModifyParameterResponseBody {
	s.RequestId = &v
	return s
}

type ModifyParameterResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyParameterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyParameterResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyParameterResponse) GoString() string {
	return s.String()
}

func (s *ModifyParameterResponse) SetHeaders(v map[string]*string) *ModifyParameterResponse {
	s.Headers = v
	return s
}

func (s *ModifyParameterResponse) SetBody(v *ModifyParameterResponseBody) *ModifyParameterResponse {
	s.Body = v
	return s
}

type ModifySecurityIpsRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	GroupName      *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	ModifyMode     *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
}

func (s ModifySecurityIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsRequest) SetRegionId(v string) *ModifySecurityIpsRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetDBInstanceName(v string) *ModifySecurityIpsRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetGroupName(v string) *ModifySecurityIpsRequest {
	s.GroupName = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetSecurityIPList(v string) *ModifySecurityIpsRequest {
	s.SecurityIPList = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetModifyMode(v string) *ModifySecurityIpsRequest {
	s.ModifyMode = &v
	return s
}

type ModifySecurityIpsResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifySecurityIpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityIpsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsResponseBody) SetMessage(v string) *ModifySecurityIpsResponseBody {
	s.Message = &v
	return s
}

func (s *ModifySecurityIpsResponseBody) SetRequestId(v string) *ModifySecurityIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySecurityIpsResponseBody) SetSuccess(v bool) *ModifySecurityIpsResponseBody {
	s.Success = &v
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

type ReleaseInstancePublicConnectionRequest struct {
	OwnerId                 *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount    *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId         *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DBInstanceName          *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
	OwnerAccount            *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ReleaseInstancePublicConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstancePublicConnectionRequest) GoString() string {
	return s.String()
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

func (s *ReleaseInstancePublicConnectionRequest) SetDBInstanceName(v string) *ReleaseInstancePublicConnectionRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetCurrentConnectionString(v string) *ReleaseInstancePublicConnectionRequest {
	s.CurrentConnectionString = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetOwnerAccount(v string) *ReleaseInstancePublicConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetRegionId(v string) *ReleaseInstancePublicConnectionRequest {
	s.RegionId = &v
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

type RestartDBInstanceRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
}

func (s RestartDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartDBInstanceRequest) SetRegionId(v string) *RestartDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RestartDBInstanceRequest) SetDBInstanceName(v string) *RestartDBInstanceRequest {
	s.DBInstanceName = &v
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

type RetryPolarxOrderRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	ScaleOutToken  *string `json:"ScaleOutToken,omitempty" xml:"ScaleOutToken,omitempty"`
}

func (s RetryPolarxOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s RetryPolarxOrderRequest) GoString() string {
	return s.String()
}

func (s *RetryPolarxOrderRequest) SetRegionId(v string) *RetryPolarxOrderRequest {
	s.RegionId = &v
	return s
}

func (s *RetryPolarxOrderRequest) SetDBInstanceName(v string) *RetryPolarxOrderRequest {
	s.DBInstanceName = &v
	return s
}

func (s *RetryPolarxOrderRequest) SetScaleOutToken(v string) *RetryPolarxOrderRequest {
	s.ScaleOutToken = &v
	return s
}

type RetryPolarxOrderResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RetryPolarxOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetryPolarxOrderResponseBody) GoString() string {
	return s.String()
}

func (s *RetryPolarxOrderResponseBody) SetRequestId(v string) *RetryPolarxOrderResponseBody {
	s.RequestId = &v
	return s
}

type RetryPolarxOrderResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RetryPolarxOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RetryPolarxOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s RetryPolarxOrderResponse) GoString() string {
	return s.String()
}

func (s *RetryPolarxOrderResponse) SetHeaders(v map[string]*string) *RetryPolarxOrderResponse {
	s.Headers = v
	return s
}

func (s *RetryPolarxOrderResponse) SetBody(v *RetryPolarxOrderResponseBody) *RetryPolarxOrderResponse {
	s.Body = v
	return s
}

type UpdateBackupPolicyRequest struct {
	RegionId                   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName             *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	BackupPeriod               *string `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	BackupPlanBegin            *string `json:"BackupPlanBegin,omitempty" xml:"BackupPlanBegin,omitempty"`
	BackupSetRetention         *int32  `json:"BackupSetRetention,omitempty" xml:"BackupSetRetention,omitempty"`
	BackupType                 *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	BackupWay                  *string `json:"BackupWay,omitempty" xml:"BackupWay,omitempty"`
	ForceCleanOnHighSpaceUsage *int32  `json:"ForceCleanOnHighSpaceUsage,omitempty" xml:"ForceCleanOnHighSpaceUsage,omitempty"`
	IsEnabled                  *int32  `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	LocalLogRetention          *int32  `json:"LocalLogRetention,omitempty" xml:"LocalLogRetention,omitempty"`
	RemoveLogRetention         *int32  `json:"RemoveLogRetention,omitempty" xml:"RemoveLogRetention,omitempty"`
	LogLocalRetentionSpace     *int32  `json:"LogLocalRetentionSpace,omitempty" xml:"LogLocalRetentionSpace,omitempty"`
}

func (s UpdateBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateBackupPolicyRequest) SetRegionId(v string) *UpdateBackupPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetDBInstanceName(v string) *UpdateBackupPolicyRequest {
	s.DBInstanceName = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetBackupPeriod(v string) *UpdateBackupPolicyRequest {
	s.BackupPeriod = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetBackupPlanBegin(v string) *UpdateBackupPolicyRequest {
	s.BackupPlanBegin = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetBackupSetRetention(v int32) *UpdateBackupPolicyRequest {
	s.BackupSetRetention = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetBackupType(v string) *UpdateBackupPolicyRequest {
	s.BackupType = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetBackupWay(v string) *UpdateBackupPolicyRequest {
	s.BackupWay = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetForceCleanOnHighSpaceUsage(v int32) *UpdateBackupPolicyRequest {
	s.ForceCleanOnHighSpaceUsage = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetIsEnabled(v int32) *UpdateBackupPolicyRequest {
	s.IsEnabled = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetLocalLogRetention(v int32) *UpdateBackupPolicyRequest {
	s.LocalLogRetention = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetRemoveLogRetention(v int32) *UpdateBackupPolicyRequest {
	s.RemoveLogRetention = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetLogLocalRetentionSpace(v int32) *UpdateBackupPolicyRequest {
	s.LogLocalRetentionSpace = &v
	return s
}

type UpdateBackupPolicyResponseBody struct {
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      []*UpdateBackupPolicyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s UpdateBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBackupPolicyResponseBody) SetMessage(v string) *UpdateBackupPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateBackupPolicyResponseBody) SetRequestId(v string) *UpdateBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBackupPolicyResponseBody) SetSuccess(v bool) *UpdateBackupPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateBackupPolicyResponseBody) SetData(v []*UpdateBackupPolicyResponseBodyData) *UpdateBackupPolicyResponseBody {
	s.Data = v
	return s
}

type UpdateBackupPolicyResponseBodyData struct {
	LogLocalRetentionSpace     *int32  `json:"LogLocalRetentionSpace,omitempty" xml:"LogLocalRetentionSpace,omitempty"`
	DBInstanceName             *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	BackupWay                  *string `json:"BackupWay,omitempty" xml:"BackupWay,omitempty"`
	BackupPeriod               *string `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	ForceCleanOnHighSpaceUsage *int32  `json:"ForceCleanOnHighSpaceUsage,omitempty" xml:"ForceCleanOnHighSpaceUsage,omitempty"`
	BackupType                 *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	LocalLogRetention          *int32  `json:"LocalLogRetention,omitempty" xml:"LocalLogRetention,omitempty"`
	RemoveLogRetention         *int32  `json:"RemoveLogRetention,omitempty" xml:"RemoveLogRetention,omitempty"`
	BackupPlanBegin            *string `json:"BackupPlanBegin,omitempty" xml:"BackupPlanBegin,omitempty"`
	BackupSetRetention         *int32  `json:"BackupSetRetention,omitempty" xml:"BackupSetRetention,omitempty"`
	IsEnabled                  *int32  `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
}

func (s UpdateBackupPolicyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateBackupPolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateBackupPolicyResponseBodyData) SetLogLocalRetentionSpace(v int32) *UpdateBackupPolicyResponseBodyData {
	s.LogLocalRetentionSpace = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetDBInstanceName(v string) *UpdateBackupPolicyResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetBackupWay(v string) *UpdateBackupPolicyResponseBodyData {
	s.BackupWay = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetBackupPeriod(v string) *UpdateBackupPolicyResponseBodyData {
	s.BackupPeriod = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetForceCleanOnHighSpaceUsage(v int32) *UpdateBackupPolicyResponseBodyData {
	s.ForceCleanOnHighSpaceUsage = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetBackupType(v string) *UpdateBackupPolicyResponseBodyData {
	s.BackupType = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetLocalLogRetention(v int32) *UpdateBackupPolicyResponseBodyData {
	s.LocalLogRetention = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetRemoveLogRetention(v int32) *UpdateBackupPolicyResponseBodyData {
	s.RemoveLogRetention = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetBackupPlanBegin(v string) *UpdateBackupPolicyResponseBodyData {
	s.BackupPlanBegin = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetBackupSetRetention(v int32) *UpdateBackupPolicyResponseBodyData {
	s.BackupSetRetention = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetIsEnabled(v int32) *UpdateBackupPolicyResponseBodyData {
	s.IsEnabled = &v
	return s
}

type UpdateBackupPolicyResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateBackupPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateBackupPolicyResponse) SetHeaders(v map[string]*string) *UpdateBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateBackupPolicyResponse) SetBody(v *UpdateBackupPolicyResponseBody) *UpdateBackupPolicyResponse {
	s.Body = v
	return s
}

type UpdateDBInstanceSSLRequest struct {
	EnableSSL      *bool   `json:"EnableSSL,omitempty" xml:"EnableSSL,omitempty"`
	CertCommonName *string `json:"CertCommonName,omitempty" xml:"CertCommonName,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateDBInstanceSSLRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDBInstanceSSLRequest) GoString() string {
	return s.String()
}

func (s *UpdateDBInstanceSSLRequest) SetEnableSSL(v bool) *UpdateDBInstanceSSLRequest {
	s.EnableSSL = &v
	return s
}

func (s *UpdateDBInstanceSSLRequest) SetCertCommonName(v string) *UpdateDBInstanceSSLRequest {
	s.CertCommonName = &v
	return s
}

func (s *UpdateDBInstanceSSLRequest) SetDBInstanceName(v string) *UpdateDBInstanceSSLRequest {
	s.DBInstanceName = &v
	return s
}

func (s *UpdateDBInstanceSSLRequest) SetRegionId(v string) *UpdateDBInstanceSSLRequest {
	s.RegionId = &v
	return s
}

type UpdateDBInstanceSSLResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *UpdateDBInstanceSSLResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s UpdateDBInstanceSSLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDBInstanceSSLResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDBInstanceSSLResponseBody) SetRequestId(v string) *UpdateDBInstanceSSLResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDBInstanceSSLResponseBody) SetData(v *UpdateDBInstanceSSLResponseBodyData) *UpdateDBInstanceSSLResponseBody {
	s.Data = v
	return s
}

type UpdateDBInstanceSSLResponseBodyData struct {
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateDBInstanceSSLResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateDBInstanceSSLResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateDBInstanceSSLResponseBodyData) SetTaskId(v int64) *UpdateDBInstanceSSLResponseBodyData {
	s.TaskId = &v
	return s
}

type UpdateDBInstanceSSLResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDBInstanceSSLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDBInstanceSSLResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDBInstanceSSLResponse) GoString() string {
	return s.String()
}

func (s *UpdateDBInstanceSSLResponse) SetHeaders(v map[string]*string) *UpdateDBInstanceSSLResponse {
	s.Headers = v
	return s
}

func (s *UpdateDBInstanceSSLResponse) SetBody(v *UpdateDBInstanceSSLResponseBody) *UpdateDBInstanceSSLResponse {
	s.Body = v
	return s
}

type UpdateDBInstanceTDERequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	TDEStatus      *int32  `json:"TDEStatus,omitempty" xml:"TDEStatus,omitempty"`
	EncryptionKey  *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	RoleArn        *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s UpdateDBInstanceTDERequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDBInstanceTDERequest) GoString() string {
	return s.String()
}

func (s *UpdateDBInstanceTDERequest) SetRegionId(v string) *UpdateDBInstanceTDERequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDBInstanceTDERequest) SetDBInstanceName(v string) *UpdateDBInstanceTDERequest {
	s.DBInstanceName = &v
	return s
}

func (s *UpdateDBInstanceTDERequest) SetTDEStatus(v int32) *UpdateDBInstanceTDERequest {
	s.TDEStatus = &v
	return s
}

func (s *UpdateDBInstanceTDERequest) SetEncryptionKey(v string) *UpdateDBInstanceTDERequest {
	s.EncryptionKey = &v
	return s
}

func (s *UpdateDBInstanceTDERequest) SetRoleArn(v string) *UpdateDBInstanceTDERequest {
	s.RoleArn = &v
	return s
}

type UpdateDBInstanceTDEResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *UpdateDBInstanceTDEResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s UpdateDBInstanceTDEResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDBInstanceTDEResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDBInstanceTDEResponseBody) SetRequestId(v string) *UpdateDBInstanceTDEResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDBInstanceTDEResponseBody) SetData(v *UpdateDBInstanceTDEResponseBodyData) *UpdateDBInstanceTDEResponseBody {
	s.Data = v
	return s
}

type UpdateDBInstanceTDEResponseBodyData struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateDBInstanceTDEResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateDBInstanceTDEResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateDBInstanceTDEResponseBodyData) SetTaskId(v string) *UpdateDBInstanceTDEResponseBodyData {
	s.TaskId = &v
	return s
}

type UpdateDBInstanceTDEResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDBInstanceTDEResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDBInstanceTDEResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDBInstanceTDEResponse) GoString() string {
	return s.String()
}

func (s *UpdateDBInstanceTDEResponse) SetHeaders(v map[string]*string) *UpdateDBInstanceTDEResponse {
	s.Headers = v
	return s
}

func (s *UpdateDBInstanceTDEResponse) SetBody(v *UpdateDBInstanceTDEResponseBody) *UpdateDBInstanceTDEResponse {
	s.Body = v
	return s
}

type UpdatePolarDBXInstanceNodeRequest struct {
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName      *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DbInstanceNodeCount *string `json:"DbInstanceNodeCount,omitempty" xml:"DbInstanceNodeCount,omitempty"`
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s UpdatePolarDBXInstanceNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePolarDBXInstanceNodeRequest) GoString() string {
	return s.String()
}

func (s *UpdatePolarDBXInstanceNodeRequest) SetRegionId(v string) *UpdatePolarDBXInstanceNodeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdatePolarDBXInstanceNodeRequest) SetDBInstanceName(v string) *UpdatePolarDBXInstanceNodeRequest {
	s.DBInstanceName = &v
	return s
}

func (s *UpdatePolarDBXInstanceNodeRequest) SetDbInstanceNodeCount(v string) *UpdatePolarDBXInstanceNodeRequest {
	s.DbInstanceNodeCount = &v
	return s
}

func (s *UpdatePolarDBXInstanceNodeRequest) SetClientToken(v string) *UpdatePolarDBXInstanceNodeRequest {
	s.ClientToken = &v
	return s
}

type UpdatePolarDBXInstanceNodeResponseBody struct {
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePolarDBXInstanceNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePolarDBXInstanceNodeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePolarDBXInstanceNodeResponseBody) SetOrderId(v string) *UpdatePolarDBXInstanceNodeResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpdatePolarDBXInstanceNodeResponseBody) SetRequestId(v string) *UpdatePolarDBXInstanceNodeResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePolarDBXInstanceNodeResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdatePolarDBXInstanceNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePolarDBXInstanceNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePolarDBXInstanceNodeResponse) GoString() string {
	return s.String()
}

func (s *UpdatePolarDBXInstanceNodeResponse) SetHeaders(v map[string]*string) *UpdatePolarDBXInstanceNodeResponse {
	s.Headers = v
	return s
}

func (s *UpdatePolarDBXInstanceNodeResponse) SetBody(v *UpdatePolarDBXInstanceNodeResponseBody) *UpdatePolarDBXInstanceNodeResponse {
	s.Body = v
	return s
}

type UpgradeDBInstanceKernelVersionRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	UpgradeTime    *string `json:"UpgradeTime,omitempty" xml:"UpgradeTime,omitempty"`
	SwitchTime     *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
}

func (s UpgradeDBInstanceKernelVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBInstanceKernelVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetRegionId(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetDBInstanceName(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.DBInstanceName = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetUpgradeTime(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.UpgradeTime = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetSwitchTime(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.SwitchTime = &v
	return s
}

type UpgradeDBInstanceKernelVersionResponseBody struct {
	DBInstanceName     *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	TargetMinorVersion *string `json:"TargetMinorVersion,omitempty" xml:"TargetMinorVersion,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId             *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpgradeDBInstanceKernelVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBInstanceKernelVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceKernelVersionResponseBody) SetDBInstanceName(v string) *UpgradeDBInstanceKernelVersionResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionResponseBody) SetTargetMinorVersion(v string) *UpgradeDBInstanceKernelVersionResponseBody {
	s.TargetMinorVersion = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionResponseBody) SetRequestId(v string) *UpgradeDBInstanceKernelVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionResponseBody) SetTaskId(v string) *UpgradeDBInstanceKernelVersionResponseBody {
	s.TaskId = &v
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
		"ap-northeast-1":              tea.String("polardbx.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("polardbx.aliyuncs.com"),
		"ap-south-1":                  tea.String("polardbx.aliyuncs.com"),
		"ap-southeast-2":              tea.String("polardbx.aliyuncs.com"),
		"ap-southeast-3":              tea.String("polardbx.aliyuncs.com"),
		"ap-southeast-5":              tea.String("polardbx.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("polardbx.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("polardbx.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("polardbx.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("polardbx.aliyuncs.com"),
		"cn-edge-1":                   tea.String("polardbx.aliyuncs.com"),
		"cn-fujian":                   tea.String("polardbx.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("polardbx.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("polardbx.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("polardbx.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("polardbx.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("polardbx.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("polardbx.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("polardbx.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("polardbx.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("polardbx.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("polardbx.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("polardbx.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("polardbx.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("polardbx.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("polardbx.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("polardbx.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("polardbx.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("polardbx.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("polardbx.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("polardbx.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("polardbx.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("polardbx.aliyuncs.com"),
		"cn-wuhan":                    tea.String("polardbx.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("polardbx.aliyuncs.com"),
		"cn-yushanfang":               tea.String("polardbx.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("polardbx.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("polardbx.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("polardbx.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("polardbx.aliyuncs.com"),
		"eu-central-1":                tea.String("polardbx.aliyuncs.com"),
		"eu-west-1":                   tea.String("polardbx.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("polardbx.aliyuncs.com"),
		"me-east-1":                   tea.String("polardbx.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("polardbx.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("polardbx"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AllocateInstancePublicConnectionWithOptions(request *AllocateInstancePublicConnectionRequest, runtime *util.RuntimeOptions) (_result *AllocateInstancePublicConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AllocateInstancePublicConnectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AllocateInstancePublicConnection"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CancelPolarxOrderWithOptions(request *CancelPolarxOrderRequest, runtime *util.RuntimeOptions) (_result *CancelPolarxOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CancelPolarxOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CancelPolarxOrder"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelPolarxOrder(request *CancelPolarxOrderRequest) (_result *CancelPolarxOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelPolarxOrderResponse{}
	_body, _err := client.CancelPolarxOrderWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("CheckCloudResourceAuthorized"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateAccountWithOptions(request *CreateAccountRequest, runtime *util.RuntimeOptions) (_result *CreateAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAccount"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("CreateBackup"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateDBWithOptions(request *CreateDBRequest, runtime *util.RuntimeOptions) (_result *CreateDBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDBResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDB"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDB(request *CreateDBRequest) (_result *CreateDBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBResponse{}
	_body, _err := client.CreateDBWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("CreateDBInstance"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreatePolarxInstanceWithOptions(request *CreatePolarxInstanceRequest, runtime *util.RuntimeOptions) (_result *CreatePolarxInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreatePolarxInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreatePolarxInstance"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePolarxInstance(request *CreatePolarxInstanceRequest) (_result *CreatePolarxInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePolarxInstanceResponse{}
	_body, _err := client.CreatePolarxInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePolarxOrderWithOptions(request *CreatePolarxOrderRequest, runtime *util.RuntimeOptions) (_result *CreatePolarxOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreatePolarxOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreatePolarxOrder"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePolarxOrder(request *CreatePolarxOrderRequest) (_result *CreatePolarxOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePolarxOrderResponse{}
	_body, _err := client.CreatePolarxOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSuperAccountWithOptions(request *CreateSuperAccountRequest, runtime *util.RuntimeOptions) (_result *CreateSuperAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSuperAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSuperAccount"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSuperAccount(request *CreateSuperAccountRequest) (_result *CreateSuperAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSuperAccountResponse{}
	_body, _err := client.CreateSuperAccountWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DeleteAccount"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DeleteDBWithOptions(request *DeleteDBRequest, runtime *util.RuntimeOptions) (_result *DeleteDBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDBResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDB"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDB(request *DeleteDBRequest) (_result *DeleteDBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDBResponse{}
	_body, _err := client.DeleteDBWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DeleteDBInstance"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeAccountListWithOptions(request *DescribeAccountListRequest, runtime *util.RuntimeOptions) (_result *DescribeAccountListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAccountListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAccountList"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccountList(request *DescribeAccountListRequest) (_result *DescribeAccountListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccountListResponse{}
	_body, _err := client.DescribeAccountListWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackupPolicy"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeBackupSetListWithOptions(request *DescribeBackupSetListRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupSetListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeBackupSetListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackupSetList"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupSetList(request *DescribeBackupSetListRequest) (_result *DescribeBackupSetListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupSetListResponse{}
	_body, _err := client.DescribeBackupSetListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBinaryLogListWithOptions(request *DescribeBinaryLogListRequest, runtime *util.RuntimeOptions) (_result *DescribeBinaryLogListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBinaryLogListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBinaryLogList"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBinaryLogList(request *DescribeBinaryLogListRequest) (_result *DescribeBinaryLogListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBinaryLogListResponse{}
	_body, _err := client.DescribeBinaryLogListWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBInstanceAttribute"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeDBInstancesWithOptions(request *DescribeDBInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBInstances"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBInstanceSSL"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeDBInstanceTDEWithOptions(request *DescribeDBInstanceTDERequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceTDEResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBInstanceTDEResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBInstanceTDE"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceTDE(request *DescribeDBInstanceTDERequest) (_result *DescribeDBInstanceTDEResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceTDEResponse{}
	_body, _err := client.DescribeDBInstanceTDEWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDbListWithOptions(request *DescribeDbListRequest, runtime *util.RuntimeOptions) (_result *DescribeDbListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDbListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDbList"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDbList(request *DescribeDbListRequest) (_result *DescribeDbListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDbListResponse{}
	_body, _err := client.DescribeDbListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDistributeTableListWithOptions(request *DescribeDistributeTableListRequest, runtime *util.RuntimeOptions) (_result *DescribeDistributeTableListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDistributeTableListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDistributeTableList"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDistributeTableList(request *DescribeDistributeTableListRequest) (_result *DescribeDistributeTableListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDistributeTableListResponse{}
	_body, _err := client.DescribeDistributeTableListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceDbPerformanceWithOptions(request *DescribeInstanceDbPerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceDbPerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceDbPerformanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceDbPerformance"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceDbPerformance(request *DescribeInstanceDbPerformanceRequest) (_result *DescribeInstanceDbPerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceDbPerformanceResponse{}
	_body, _err := client.DescribeInstanceDbPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstancePerformanceWithOptions(request *DescribeInstancePerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeInstancePerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstancePerformanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstancePerformance"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstancePerformance(request *DescribeInstancePerformanceRequest) (_result *DescribeInstancePerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstancePerformanceResponse{}
	_body, _err := client.DescribeInstancePerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceStoragePerformanceWithOptions(request *DescribeInstanceStoragePerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceStoragePerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceStoragePerformanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceStoragePerformance"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceStoragePerformance(request *DescribeInstanceStoragePerformanceRequest) (_result *DescribeInstanceStoragePerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceStoragePerformanceResponse{}
	_body, _err := client.DescribeInstanceStoragePerformanceWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeParameters"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribePolarxDbInstancesWithOptions(request *DescribePolarxDbInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribePolarxDbInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePolarxDbInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePolarxDbInstances"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePolarxDbInstances(request *DescribePolarxDbInstancesRequest) (_result *DescribePolarxDbInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePolarxDbInstancesResponse{}
	_body, _err := client.DescribePolarxDbInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions() (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScaleOutMigrateTaskListWithOptions(request *DescribeScaleOutMigrateTaskListRequest, runtime *util.RuntimeOptions) (_result *DescribeScaleOutMigrateTaskListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeScaleOutMigrateTaskListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeScaleOutMigrateTaskList"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScaleOutMigrateTaskList(request *DescribeScaleOutMigrateTaskListRequest) (_result *DescribeScaleOutMigrateTaskListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScaleOutMigrateTaskListResponse{}
	_body, _err := client.DescribeScaleOutMigrateTaskListWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeSecurityIps"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeSqlAuditInfoWithOptions(request *DescribeSqlAuditInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeSqlAuditInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSqlAuditInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSqlAuditInfo"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSqlAuditInfo(request *DescribeSqlAuditInfoRequest) (_result *DescribeSqlAuditInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSqlAuditInfoResponse{}
	_body, _err := client.DescribeSqlAuditInfoWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeTasks"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeUserEncryptionKeyListWithOptions(request *DescribeUserEncryptionKeyListRequest, runtime *util.RuntimeOptions) (_result *DescribeUserEncryptionKeyListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUserEncryptionKeyListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUserEncryptionKeyList"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserEncryptionKeyList(request *DescribeUserEncryptionKeyListRequest) (_result *DescribeUserEncryptionKeyListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserEncryptionKeyListResponse{}
	_body, _err := client.DescribeUserEncryptionKeyListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPolarxCommodityWithOptions(request *GetPolarxCommodityRequest, runtime *util.RuntimeOptions) (_result *GetPolarxCommodityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPolarxCommodityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPolarxCommodity"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPolarxCommodity(request *GetPolarxCommodityRequest) (_result *GetPolarxCommodityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPolarxCommodityResponse{}
	_body, _err := client.GetPolarxCommodityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPolarXPriceWithOptions(request *GetPolarXPriceRequest, runtime *util.RuntimeOptions) (_result *GetPolarXPriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPolarXPriceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPolarXPrice"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPolarXPrice(request *GetPolarXPriceRequest) (_result *GetPolarXPriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPolarXPriceResponse{}
	_body, _err := client.GetPolarXPriceWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("ModifyAccountDescription"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ModifyDatabaseDescriptionWithOptions(request *ModifyDatabaseDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyDatabaseDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDatabaseDescriptionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDatabaseDescription"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDatabaseDescription(request *ModifyDatabaseDescriptionRequest) (_result *ModifyDatabaseDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDatabaseDescriptionResponse{}
	_body, _err := client.ModifyDatabaseDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBInstanceClassWithOptions(request *ModifyDBInstanceClassRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceClassResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBInstanceClassResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBInstanceClass"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceClass(request *ModifyDBInstanceClassRequest) (_result *ModifyDBInstanceClassResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceClassResponse{}
	_body, _err := client.ModifyDBInstanceClassWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBInstanceConfigWithOptions(request *ModifyDBInstanceConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBInstanceConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBInstanceConfig"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceConfig(request *ModifyDBInstanceConfigRequest) (_result *ModifyDBInstanceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceConfigResponse{}
	_body, _err := client.ModifyDBInstanceConfigWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBInstanceDescription"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ModifyParameterWithOptions(request *ModifyParameterRequest, runtime *util.RuntimeOptions) (_result *ModifyParameterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyParameterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyParameter"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyParameter(request *ModifyParameterRequest) (_result *ModifyParameterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyParameterResponse{}
	_body, _err := client.ModifyParameterWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("ModifySecurityIps"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ReleaseInstancePublicConnectionWithOptions(request *ReleaseInstancePublicConnectionRequest, runtime *util.RuntimeOptions) (_result *ReleaseInstancePublicConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReleaseInstancePublicConnectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReleaseInstancePublicConnection"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RestartDBInstanceWithOptions(request *RestartDBInstanceRequest, runtime *util.RuntimeOptions) (_result *RestartDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RestartDBInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RestartDBInstance"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RetryPolarxOrderWithOptions(request *RetryPolarxOrderRequest, runtime *util.RuntimeOptions) (_result *RetryPolarxOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RetryPolarxOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RetryPolarxOrder"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RetryPolarxOrder(request *RetryPolarxOrderRequest) (_result *RetryPolarxOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RetryPolarxOrderResponse{}
	_body, _err := client.RetryPolarxOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateBackupPolicyWithOptions(request *UpdateBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *UpdateBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateBackupPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateBackupPolicy"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateBackupPolicy(request *UpdateBackupPolicyRequest) (_result *UpdateBackupPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateBackupPolicyResponse{}
	_body, _err := client.UpdateBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDBInstanceSSLWithOptions(request *UpdateDBInstanceSSLRequest, runtime *util.RuntimeOptions) (_result *UpdateDBInstanceSSLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDBInstanceSSLResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDBInstanceSSL"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDBInstanceSSL(request *UpdateDBInstanceSSLRequest) (_result *UpdateDBInstanceSSLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDBInstanceSSLResponse{}
	_body, _err := client.UpdateDBInstanceSSLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDBInstanceTDEWithOptions(request *UpdateDBInstanceTDERequest, runtime *util.RuntimeOptions) (_result *UpdateDBInstanceTDEResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDBInstanceTDEResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDBInstanceTDE"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDBInstanceTDE(request *UpdateDBInstanceTDERequest) (_result *UpdateDBInstanceTDEResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDBInstanceTDEResponse{}
	_body, _err := client.UpdateDBInstanceTDEWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePolarDBXInstanceNodeWithOptions(request *UpdatePolarDBXInstanceNodeRequest, runtime *util.RuntimeOptions) (_result *UpdatePolarDBXInstanceNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdatePolarDBXInstanceNodeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdatePolarDBXInstanceNode"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePolarDBXInstanceNode(request *UpdatePolarDBXInstanceNodeRequest) (_result *UpdatePolarDBXInstanceNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePolarDBXInstanceNodeResponse{}
	_body, _err := client.UpdatePolarDBXInstanceNodeWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("UpgradeDBInstanceKernelVersion"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
