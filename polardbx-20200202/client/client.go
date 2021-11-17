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
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	DBInstanceName         *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Port                   *string `json:"Port,omitempty" xml:"Port,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AllocateInstancePublicConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s AllocateInstancePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *AllocateInstancePublicConnectionRequest) SetConnectionStringPrefix(v string) *AllocateInstancePublicConnectionRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetDBInstanceName(v string) *AllocateInstancePublicConnectionRequest {
	s.DBInstanceName = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetOwnerAccount(v string) *AllocateInstancePublicConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetOwnerId(v int64) *AllocateInstancePublicConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetPort(v string) *AllocateInstancePublicConnectionRequest {
	s.Port = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetRegionId(v string) *AllocateInstancePublicConnectionRequest {
	s.RegionId = &v
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
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ScaleOutToken  *string `json:"ScaleOutToken,omitempty" xml:"ScaleOutToken,omitempty"`
}

func (s CancelPolarxOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelPolarxOrderRequest) GoString() string {
	return s.String()
}

func (s *CancelPolarxOrderRequest) SetDBInstanceName(v string) *CancelPolarxOrderRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CancelPolarxOrderRequest) SetRegionId(v string) *CancelPolarxOrderRequest {
	s.RegionId = &v
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
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RoleArn        *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s CheckCloudResourceAuthorizedRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckCloudResourceAuthorizedRequest) GoString() string {
	return s.String()
}

func (s *CheckCloudResourceAuthorizedRequest) SetDBInstanceName(v string) *CheckCloudResourceAuthorizedRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetRegionId(v string) *CheckCloudResourceAuthorizedRequest {
	s.RegionId = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetRoleArn(v string) *CheckCloudResourceAuthorizedRequest {
	s.RoleArn = &v
	return s
}

type CheckCloudResourceAuthorizedResponseBody struct {
	Data      *CheckCloudResourceAuthorizedResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckCloudResourceAuthorizedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckCloudResourceAuthorizedResponseBody) GoString() string {
	return s.String()
}

func (s *CheckCloudResourceAuthorizedResponseBody) SetData(v *CheckCloudResourceAuthorizedResponseBodyData) *CheckCloudResourceAuthorizedResponseBody {
	s.Data = v
	return s
}

func (s *CheckCloudResourceAuthorizedResponseBody) SetRequestId(v string) *CheckCloudResourceAuthorizedResponseBody {
	s.RequestId = &v
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
	AccountDescription      *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	AccountName             *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPassword         *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	AccountPrivilege        *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	DBInstanceName          *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DBName                  *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityAccountName     *string `json:"SecurityAccountName,omitempty" xml:"SecurityAccountName,omitempty"`
	SecurityAccountPassword *string `json:"SecurityAccountPassword,omitempty" xml:"SecurityAccountPassword,omitempty"`
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

func (s *CreateAccountRequest) SetAccountPrivilege(v string) *CreateAccountRequest {
	s.AccountPrivilege = &v
	return s
}

func (s *CreateAccountRequest) SetDBInstanceName(v string) *CreateAccountRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateAccountRequest) SetDBName(v string) *CreateAccountRequest {
	s.DBName = &v
	return s
}

func (s *CreateAccountRequest) SetRegionId(v string) *CreateAccountRequest {
	s.RegionId = &v
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
	BackupType     *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateBackupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupRequest) GoString() string {
	return s.String()
}

func (s *CreateBackupRequest) SetBackupType(v string) *CreateBackupRequest {
	s.BackupType = &v
	return s
}

func (s *CreateBackupRequest) SetDBInstanceName(v string) *CreateBackupRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateBackupRequest) SetRegionId(v string) *CreateBackupRequest {
	s.RegionId = &v
	return s
}

type CreateBackupResponseBody struct {
	Data      []*CreateBackupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateBackupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBackupResponseBody) SetData(v []*CreateBackupResponseBodyData) *CreateBackupResponseBody {
	s.Data = v
	return s
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
	AccountName             *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPrivilege        *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	Charset                 *string `json:"Charset,omitempty" xml:"Charset,omitempty"`
	DBInstanceName          *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DbDescription           *string `json:"DbDescription,omitempty" xml:"DbDescription,omitempty"`
	DbName                  *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityAccountName     *string `json:"SecurityAccountName,omitempty" xml:"SecurityAccountName,omitempty"`
	SecurityAccountPassword *string `json:"SecurityAccountPassword,omitempty" xml:"SecurityAccountPassword,omitempty"`
}

func (s CreateDBRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBRequest) GoString() string {
	return s.String()
}

func (s *CreateDBRequest) SetAccountName(v string) *CreateDBRequest {
	s.AccountName = &v
	return s
}

func (s *CreateDBRequest) SetAccountPrivilege(v string) *CreateDBRequest {
	s.AccountPrivilege = &v
	return s
}

func (s *CreateDBRequest) SetCharset(v string) *CreateDBRequest {
	s.Charset = &v
	return s
}

func (s *CreateDBRequest) SetDBInstanceName(v string) *CreateDBRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateDBRequest) SetDbDescription(v string) *CreateDBRequest {
	s.DbDescription = &v
	return s
}

func (s *CreateDBRequest) SetDbName(v string) *CreateDBRequest {
	s.DbName = &v
	return s
}

func (s *CreateDBRequest) SetRegionId(v string) *CreateDBRequest {
	s.RegionId = &v
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
	AutoRenew             *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ClientToken           *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DBNodeClass           *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	DBNodeCount           *int32  `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	EngineVersion         *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	IsReadDBInstance      *bool   `json:"IsReadDBInstance,omitempty" xml:"IsReadDBInstance,omitempty"`
	NetworkType           *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	PayType               *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Period                *string `json:"Period,omitempty" xml:"Period,omitempty"`
	PrimaryDBInstanceName *string `json:"PrimaryDBInstanceName,omitempty" xml:"PrimaryDBInstanceName,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId       *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	UsedTime              *int32  `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	VPCId                 *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId             *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId                *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceRequest) SetAutoRenew(v bool) *CreateDBInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateDBInstanceRequest) SetClientToken(v string) *CreateDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBNodeClass(v string) *CreateDBInstanceRequest {
	s.DBNodeClass = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBNodeCount(v int32) *CreateDBInstanceRequest {
	s.DBNodeCount = &v
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

func (s *CreateDBInstanceRequest) SetNetworkType(v string) *CreateDBInstanceRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPayType(v string) *CreateDBInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPeriod(v string) *CreateDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPrimaryDBInstanceName(v string) *CreateDBInstanceRequest {
	s.PrimaryDBInstanceName = &v
	return s
}

func (s *CreateDBInstanceRequest) SetRegionId(v string) *CreateDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetResourceGroupId(v string) *CreateDBInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetUsedTime(v int32) *CreateDBInstanceRequest {
	s.UsedTime = &v
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

func (s *CreateDBInstanceRequest) SetZoneId(v string) *CreateDBInstanceRequest {
	s.ZoneId = &v
	return s
}

type CreateDBInstanceResponseBody struct {
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	OrderId        *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *CreateDBInstanceResponseBody) SetOrderId(v string) *CreateDBInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetRequestId(v string) *CreateDBInstanceResponseBody {
	s.RequestId = &v
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

type CreatePolarxOrderRequest struct {
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	NodeCount      *string `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreatePolarxOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePolarxOrderRequest) GoString() string {
	return s.String()
}

func (s *CreatePolarxOrderRequest) SetDBInstanceName(v string) *CreatePolarxOrderRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreatePolarxOrderRequest) SetNodeCount(v string) *CreatePolarxOrderRequest {
	s.NodeCount = &v
	return s
}

func (s *CreatePolarxOrderRequest) SetRegionId(v string) *CreatePolarxOrderRequest {
	s.RegionId = &v
	return s
}

type CreatePolarxOrderResponseBody struct {
	OrderResultList []*CreatePolarxOrderResponseBodyOrderResultList `json:"OrderResultList,omitempty" xml:"OrderResultList,omitempty" type:"Repeated"`
	RequestId       *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePolarxOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePolarxOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolarxOrderResponseBody) SetOrderResultList(v []*CreatePolarxOrderResponseBodyOrderResultList) *CreatePolarxOrderResponseBody {
	s.OrderResultList = v
	return s
}

func (s *CreatePolarxOrderResponseBody) SetRequestId(v string) *CreatePolarxOrderResponseBody {
	s.RequestId = &v
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
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	AccountName        *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPassword    *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	DBInstanceName     *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateSuperAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSuperAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateSuperAccountRequest) SetAccountDescription(v string) *CreateSuperAccountRequest {
	s.AccountDescription = &v
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

func (s *CreateSuperAccountRequest) SetDBInstanceName(v string) *CreateSuperAccountRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateSuperAccountRequest) SetRegionId(v string) *CreateSuperAccountRequest {
	s.RegionId = &v
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
	AccountName             *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DBInstanceName          *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
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

func (s *DeleteAccountRequest) SetAccountName(v string) *DeleteAccountRequest {
	s.AccountName = &v
	return s
}

func (s *DeleteAccountRequest) SetDBInstanceName(v string) *DeleteAccountRequest {
	s.DBInstanceName = &v
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
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDBRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBRequest) SetDBInstanceName(v string) *DeleteDBRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DeleteDBRequest) SetDbName(v string) *DeleteDBRequest {
	s.DbName = &v
	return s
}

func (s *DeleteDBRequest) SetRegionId(v string) *DeleteDBRequest {
	s.RegionId = &v
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
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceRequest) SetDBInstanceName(v string) *DeleteDBInstanceRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetRegionId(v string) *DeleteDBInstanceRequest {
	s.RegionId = &v
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
	AccountName    *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountType    *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAccountListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountListRequest) SetAccountName(v string) *DescribeAccountListRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountListRequest) SetAccountType(v string) *DescribeAccountListRequest {
	s.AccountType = &v
	return s
}

func (s *DescribeAccountListRequest) SetDBInstanceName(v string) *DescribeAccountListRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeAccountListRequest) SetRegionId(v string) *DescribeAccountListRequest {
	s.RegionId = &v
	return s
}

type DescribeAccountListResponseBody struct {
	Data      []*DescribeAccountListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAccountListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountListResponseBody) SetData(v []*DescribeAccountListResponseBodyData) *DescribeAccountListResponseBody {
	s.Data = v
	return s
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

type DescribeAccountListResponseBodyData struct {
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	AccountName        *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPrivilege   *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	AccountType        *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	DBInstanceName     *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DBName             *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	GmtCreated         *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
}

func (s DescribeAccountListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAccountListResponseBodyData) SetAccountDescription(v string) *DescribeAccountListResponseBodyData {
	s.AccountDescription = &v
	return s
}

func (s *DescribeAccountListResponseBodyData) SetAccountName(v string) *DescribeAccountListResponseBodyData {
	s.AccountName = &v
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

func (s *DescribeAccountListResponseBodyData) SetDBInstanceName(v string) *DescribeAccountListResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeAccountListResponseBodyData) SetDBName(v string) *DescribeAccountListResponseBodyData {
	s.DBName = &v
	return s
}

func (s *DescribeAccountListResponseBodyData) SetGmtCreated(v string) *DescribeAccountListResponseBodyData {
	s.GmtCreated = &v
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
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyRequest) SetDBInstanceName(v string) *DescribeBackupPolicyRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetRegionId(v string) *DescribeBackupPolicyRequest {
	s.RegionId = &v
	return s
}

type DescribeBackupPolicyResponseBody struct {
	Data      []*DescribeBackupPolicyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBody) SetData(v []*DescribeBackupPolicyResponseBodyData) *DescribeBackupPolicyResponseBody {
	s.Data = v
	return s
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

type DescribeBackupPolicyResponseBodyData struct {
	BackupPeriod               *string `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	BackupPlanBegin            *string `json:"BackupPlanBegin,omitempty" xml:"BackupPlanBegin,omitempty"`
	BackupSetRetention         *int32  `json:"BackupSetRetention,omitempty" xml:"BackupSetRetention,omitempty"`
	BackupType                 *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	BackupWay                  *string `json:"BackupWay,omitempty" xml:"BackupWay,omitempty"`
	DBInstanceName             *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	ForceCleanOnHighSpaceUsage *int32  `json:"ForceCleanOnHighSpaceUsage,omitempty" xml:"ForceCleanOnHighSpaceUsage,omitempty"`
	IsEnabled                  *int32  `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	LocalLogRetention          *int32  `json:"LocalLogRetention,omitempty" xml:"LocalLogRetention,omitempty"`
	LogLocalRetentionSpace     *int32  `json:"LogLocalRetentionSpace,omitempty" xml:"LogLocalRetentionSpace,omitempty"`
	RemoveLogRetention         *int32  `json:"RemoveLogRetention,omitempty" xml:"RemoveLogRetention,omitempty"`
}

func (s DescribeBackupPolicyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBodyData) SetBackupPeriod(v string) *DescribeBackupPolicyResponseBodyData {
	s.BackupPeriod = &v
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

func (s *DescribeBackupPolicyResponseBodyData) SetBackupType(v string) *DescribeBackupPolicyResponseBodyData {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetBackupWay(v string) *DescribeBackupPolicyResponseBodyData {
	s.BackupWay = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetDBInstanceName(v string) *DescribeBackupPolicyResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetForceCleanOnHighSpaceUsage(v int32) *DescribeBackupPolicyResponseBodyData {
	s.ForceCleanOnHighSpaceUsage = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetIsEnabled(v int32) *DescribeBackupPolicyResponseBodyData {
	s.IsEnabled = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetLocalLogRetention(v int32) *DescribeBackupPolicyResponseBodyData {
	s.LocalLogRetention = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetLogLocalRetentionSpace(v int32) *DescribeBackupPolicyResponseBodyData {
	s.LogLocalRetentionSpace = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetRemoveLogRetention(v int32) *DescribeBackupPolicyResponseBodyData {
	s.RemoveLogRetention = &v
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
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	EndTime        *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime      *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeBackupSetListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetListRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetListRequest) SetDBInstanceName(v string) *DescribeBackupSetListRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeBackupSetListRequest) SetEndTime(v int64) *DescribeBackupSetListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupSetListRequest) SetPageNumber(v int32) *DescribeBackupSetListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupSetListRequest) SetPageSize(v int32) *DescribeBackupSetListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupSetListRequest) SetRegionId(v string) *DescribeBackupSetListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBackupSetListRequest) SetStartTime(v int64) *DescribeBackupSetListRequest {
	s.StartTime = &v
	return s
}

type DescribeBackupSetListResponseBody struct {
	Data      []*DescribeBackupSetListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackupSetListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetListResponseBody) SetData(v []*DescribeBackupSetListResponseBodyData) *DescribeBackupSetListResponseBody {
	s.Data = v
	return s
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

type DescribeBackupSetListResponseBodyData struct {
	BackupModel   *int32 `json:"BackupModel,omitempty" xml:"BackupModel,omitempty"`
	BackupSetId   *int64 `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	BackupSetSize *int64 `json:"BackupSetSize,omitempty" xml:"BackupSetSize,omitempty"`
	BackupType    *int32 `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	BeginTime     *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime       *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Status        *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeBackupSetListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetListResponseBodyData) SetBackupModel(v int32) *DescribeBackupSetListResponseBodyData {
	s.BackupModel = &v
	return s
}

func (s *DescribeBackupSetListResponseBodyData) SetBackupSetId(v int64) *DescribeBackupSetListResponseBodyData {
	s.BackupSetId = &v
	return s
}

func (s *DescribeBackupSetListResponseBodyData) SetBackupSetSize(v int64) *DescribeBackupSetListResponseBodyData {
	s.BackupSetSize = &v
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

func (s *DescribeBackupSetListResponseBodyData) SetEndTime(v int64) *DescribeBackupSetListResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupSetListResponseBodyData) SetStatus(v int32) *DescribeBackupSetListResponseBodyData {
	s.Status = &v
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
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeBinaryLogListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBinaryLogListRequest) GoString() string {
	return s.String()
}

func (s *DescribeBinaryLogListRequest) SetDBInstanceName(v string) *DescribeBinaryLogListRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeBinaryLogListRequest) SetEndTime(v string) *DescribeBinaryLogListRequest {
	s.EndTime = &v
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

func (s *DescribeBinaryLogListRequest) SetRegionId(v string) *DescribeBinaryLogListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBinaryLogListRequest) SetStartTime(v string) *DescribeBinaryLogListRequest {
	s.StartTime = &v
	return s
}

type DescribeBinaryLogListResponseBody struct {
	LogList     []*DescribeBinaryLogListResponseBodyLogList `json:"LogList,omitempty" xml:"LogList,omitempty" type:"Repeated"`
	PageNumber  *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalNumber *int32                                      `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
}

func (s DescribeBinaryLogListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBinaryLogListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBinaryLogListResponseBody) SetLogList(v []*DescribeBinaryLogListResponseBodyLogList) *DescribeBinaryLogListResponseBody {
	s.LogList = v
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

func (s *DescribeBinaryLogListResponseBody) SetRequestId(v string) *DescribeBinaryLogListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBinaryLogListResponseBody) SetTotalNumber(v int32) *DescribeBinaryLogListResponseBody {
	s.TotalNumber = &v
	return s
}

type DescribeBinaryLogListResponseBodyLogList struct {
	BeginTime    *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	CreatedTime  *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	DownloadLink *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LogSize      *int64  `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	PurgeStatus  *int32  `json:"PurgeStatus,omitempty" xml:"PurgeStatus,omitempty"`
	UploadHost   *string `json:"UploadHost,omitempty" xml:"UploadHost,omitempty"`
	UploadStatus *int32  `json:"UploadStatus,omitempty" xml:"UploadStatus,omitempty"`
}

func (s DescribeBinaryLogListResponseBodyLogList) String() string {
	return tea.Prettify(s)
}

func (s DescribeBinaryLogListResponseBodyLogList) GoString() string {
	return s.String()
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetBeginTime(v string) *DescribeBinaryLogListResponseBodyLogList {
	s.BeginTime = &v
	return s
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetCreatedTime(v string) *DescribeBinaryLogListResponseBodyLogList {
	s.CreatedTime = &v
	return s
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetDownloadLink(v string) *DescribeBinaryLogListResponseBodyLogList {
	s.DownloadLink = &v
	return s
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetEndTime(v string) *DescribeBinaryLogListResponseBodyLogList {
	s.EndTime = &v
	return s
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetFileName(v string) *DescribeBinaryLogListResponseBodyLogList {
	s.FileName = &v
	return s
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetId(v int64) *DescribeBinaryLogListResponseBodyLogList {
	s.Id = &v
	return s
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetLogSize(v int64) *DescribeBinaryLogListResponseBodyLogList {
	s.LogSize = &v
	return s
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetModifiedTime(v string) *DescribeBinaryLogListResponseBodyLogList {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetPurgeStatus(v int32) *DescribeBinaryLogListResponseBodyLogList {
	s.PurgeStatus = &v
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

type DescribeCharacterSetRequest struct {
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCharacterSetRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCharacterSetRequest) GoString() string {
	return s.String()
}

func (s *DescribeCharacterSetRequest) SetDBInstanceName(v string) *DescribeCharacterSetRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeCharacterSetRequest) SetRegionId(v string) *DescribeCharacterSetRequest {
	s.RegionId = &v
	return s
}

type DescribeCharacterSetResponseBody struct {
	Data      *DescribeCharacterSetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCharacterSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCharacterSetResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCharacterSetResponseBody) SetData(v *DescribeCharacterSetResponseBodyData) *DescribeCharacterSetResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCharacterSetResponseBody) SetMessage(v string) *DescribeCharacterSetResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCharacterSetResponseBody) SetRequestId(v string) *DescribeCharacterSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCharacterSetResponseBody) SetSuccess(v bool) *DescribeCharacterSetResponseBody {
	s.Success = &v
	return s
}

type DescribeCharacterSetResponseBodyData struct {
	CharacterSet []*string `json:"CharacterSet,omitempty" xml:"CharacterSet,omitempty" type:"Repeated"`
	Engine       *string   `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s DescribeCharacterSetResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCharacterSetResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCharacterSetResponseBodyData) SetCharacterSet(v []*string) *DescribeCharacterSetResponseBodyData {
	s.CharacterSet = v
	return s
}

func (s *DescribeCharacterSetResponseBodyData) SetEngine(v string) *DescribeCharacterSetResponseBodyData {
	s.Engine = &v
	return s
}

type DescribeCharacterSetResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCharacterSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCharacterSetResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCharacterSetResponse) GoString() string {
	return s.String()
}

func (s *DescribeCharacterSetResponse) SetHeaders(v map[string]*string) *DescribeCharacterSetResponse {
	s.Headers = v
	return s
}

func (s *DescribeCharacterSetResponse) SetBody(v *DescribeCharacterSetResponseBody) *DescribeCharacterSetResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceAttributeRequest struct {
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeRequest) SetDBInstanceName(v string) *DescribeDBInstanceAttributeRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetRegionId(v string) *DescribeDBInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

type DescribeDBInstanceAttributeResponseBody struct {
	DBInstance *DescribeDBInstanceAttributeResponseBodyDBInstance `json:"DBInstance,omitempty" xml:"DBInstance,omitempty" type:"Struct"`
	RequestId  *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBody) SetDBInstance(v *DescribeDBInstanceAttributeResponseBodyDBInstance) *DescribeDBInstanceAttributeResponseBody {
	s.DBInstance = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetRequestId(v string) *DescribeDBInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBInstanceAttributeResponseBodyDBInstance struct {
	CommodityCode           *string                                                       `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ConnAddrs               []*DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs `json:"ConnAddrs,omitempty" xml:"ConnAddrs,omitempty" type:"Repeated"`
	ConnectionString        *string                                                       `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	CreateTime              *string                                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DBInstanceType          *string                                                       `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	DBNodeClass             *string                                                       `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	DBNodeCount             *int32                                                        `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	DBNodes                 []*DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes   `json:"DBNodes,omitempty" xml:"DBNodes,omitempty" type:"Repeated"`
	DBType                  *string                                                       `json:"DBType,omitempty" xml:"DBType,omitempty"`
	DBVersion               *string                                                       `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	Description             *string                                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	Engine                  *string                                                       `json:"Engine,omitempty" xml:"Engine,omitempty"`
	ExpireDate              *string                                                       `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	Expired                 *string                                                       `json:"Expired,omitempty" xml:"Expired,omitempty"`
	Id                      *string                                                       `json:"Id,omitempty" xml:"Id,omitempty"`
	KindCode                *int32                                                        `json:"KindCode,omitempty" xml:"KindCode,omitempty"`
	LatestMinorVersion      *string                                                       `json:"LatestMinorVersion,omitempty" xml:"LatestMinorVersion,omitempty"`
	LockMode                *string                                                       `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	MaintainEndTime         *string                                                       `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	MaintainStartTime       *string                                                       `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	MinorVersion            *string                                                       `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	Network                 *string                                                       `json:"Network,omitempty" xml:"Network,omitempty"`
	PayType                 *string                                                       `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Port                    *string                                                       `json:"Port,omitempty" xml:"Port,omitempty"`
	ReadDBInstances         []*string                                                     `json:"ReadDBInstances,omitempty" xml:"ReadDBInstances,omitempty" type:"Repeated"`
	RegionId                *string                                                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RightsSeparationEnabled *bool                                                         `json:"RightsSeparationEnabled,omitempty" xml:"RightsSeparationEnabled,omitempty"`
	RightsSeparationStatus  *string                                                       `json:"RightsSeparationStatus,omitempty" xml:"RightsSeparationStatus,omitempty"`
	Status                  *string                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageUsed             *int64                                                        `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	Type                    *string                                                       `json:"Type,omitempty" xml:"Type,omitempty"`
	VPCId                   *string                                                       `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId               *string                                                       `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId                  *string                                                       `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstance) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetCommodityCode(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetConnAddrs(v []*DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.ConnAddrs = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetConnectionString(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetCreateTime(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.CreateTime = &v
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

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetDBNodeCount(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.DBNodeCount = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetDBNodes(v []*DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.DBNodes = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetDBType(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.DBType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetDBVersion(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetDescription(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.Description = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetEngine(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetExpireDate(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.ExpireDate = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetExpired(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.Expired = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.Id = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetKindCode(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.KindCode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetLatestMinorVersion(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.LatestMinorVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetLockMode(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetMaintainEndTime(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetMaintainStartTime(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetMinorVersion(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.MinorVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetNetwork(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.Network = &v
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

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetReadDBInstances(v []*string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.ReadDBInstances = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetRegionId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetRightsSeparationEnabled(v bool) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.RightsSeparationEnabled = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetRightsSeparationStatus(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.RightsSeparationStatus = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetStatus(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.Status = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetStorageUsed(v int64) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.StorageUsed = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetType(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.Type = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetVPCId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.VPCId = &v
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

type DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs struct {
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	Port             *int64  `json:"Port,omitempty" xml:"Port,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VPCId            *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId        *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) SetConnectionString(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) SetPort(v int64) *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) SetType(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs {
	s.Type = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) SetVPCId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs {
	s.VPCId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) SetVSwitchId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs {
	s.VSwitchId = &v
	return s
}

type DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes struct {
	ComputeNodeId *string `json:"ComputeNodeId,omitempty" xml:"ComputeNodeId,omitempty"`
	DataNodeId    *string `json:"DataNodeId,omitempty" xml:"DataNodeId,omitempty"`
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	NodeClass     *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId        *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) SetDataNodeId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes {
	s.DataNodeId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) SetId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes {
	s.Id = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) SetNodeClass(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes {
	s.NodeClass = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) SetRegionId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) SetZoneId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes {
	s.ZoneId = &v
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

type DescribeDBInstanceConfigRequest struct {
	ConfigName     *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBInstanceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceConfigRequest) SetConfigName(v string) *DescribeDBInstanceConfigRequest {
	s.ConfigName = &v
	return s
}

func (s *DescribeDBInstanceConfigRequest) SetDBInstanceName(v string) *DescribeDBInstanceConfigRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstanceConfigRequest) SetRegionId(v string) *DescribeDBInstanceConfigRequest {
	s.RegionId = &v
	return s
}

type DescribeDBInstanceConfigResponseBody struct {
	Data      *DescribeDBInstanceConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceConfigResponseBody) SetData(v *DescribeDBInstanceConfigResponseBodyData) *DescribeDBInstanceConfigResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDBInstanceConfigResponseBody) SetRequestId(v string) *DescribeDBInstanceConfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBInstanceConfigResponseBodyData struct {
	ConfigName     *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	ConfigValue    *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	DbInstanceName *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
}

func (s DescribeDBInstanceConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceConfigResponseBodyData) SetConfigName(v string) *DescribeDBInstanceConfigResponseBodyData {
	s.ConfigName = &v
	return s
}

func (s *DescribeDBInstanceConfigResponseBodyData) SetConfigValue(v string) *DescribeDBInstanceConfigResponseBodyData {
	s.ConfigValue = &v
	return s
}

func (s *DescribeDBInstanceConfigResponseBodyData) SetDbInstanceName(v string) *DescribeDBInstanceConfigResponseBodyData {
	s.DbInstanceName = &v
	return s
}

type DescribeDBInstanceConfigResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBInstanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceConfigResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceConfigResponse) SetBody(v *DescribeDBInstanceConfigResponseBody) *DescribeDBInstanceConfigResponse {
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
	Data      *DescribeDBInstanceSSLResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceSSLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceSSLResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSSLResponseBody) SetData(v *DescribeDBInstanceSSLResponseBodyData) *DescribeDBInstanceSSLResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetRequestId(v string) *DescribeDBInstanceSSLResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBInstanceSSLResponseBodyData struct {
	CertCommonName *string `json:"CertCommonName,omitempty" xml:"CertCommonName,omitempty"`
	SSLEnabled     *bool   `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
	SSLExpiredTime *string `json:"SSLExpiredTime,omitempty" xml:"SSLExpiredTime,omitempty"`
}

func (s DescribeDBInstanceSSLResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceSSLResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSSLResponseBodyData) SetCertCommonName(v string) *DescribeDBInstanceSSLResponseBodyData {
	s.CertCommonName = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBodyData) SetSSLEnabled(v bool) *DescribeDBInstanceSSLResponseBodyData {
	s.SSLEnabled = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBodyData) SetSSLExpiredTime(v string) *DescribeDBInstanceSSLResponseBodyData {
	s.SSLExpiredTime = &v
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
	Data      *DescribeDBInstanceTDEResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceTDEResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceTDEResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTDEResponseBody) SetData(v *DescribeDBInstanceTDEResponseBodyData) *DescribeDBInstanceTDEResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDBInstanceTDEResponseBody) SetRequestId(v string) *DescribeDBInstanceTDEResponseBody {
	s.RequestId = &v
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

type DescribeDBInstanceTopologyRequest struct {
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBInstanceTopologyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceTopologyRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTopologyRequest) SetDBInstanceName(v string) *DescribeDBInstanceTopologyRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstanceTopologyRequest) SetRegionId(v string) *DescribeDBInstanceTopologyRequest {
	s.RegionId = &v
	return s
}

type DescribeDBInstanceTopologyResponseBody struct {
	Data      *DescribeDBInstanceTopologyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceTopologyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceTopologyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTopologyResponseBody) SetData(v *DescribeDBInstanceTopologyResponseBodyData) *DescribeDBInstanceTopologyResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBody) SetRequestId(v string) *DescribeDBInstanceTopologyResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBInstanceTopologyResponseBodyData struct {
	LogicInstanceTopology *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology `json:"LogicInstanceTopology,omitempty" xml:"LogicInstanceTopology,omitempty" type:"Struct"`
}

func (s DescribeDBInstanceTopologyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceTopologyResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTopologyResponseBodyData) SetLogicInstanceTopology(v *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) *DescribeDBInstanceTopologyResponseBodyData {
	s.LogicInstanceTopology = v
	return s
}

type DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology struct {
	DBInstanceConnType          *string                                                                 `json:"DBInstanceConnType,omitempty" xml:"DBInstanceConnType,omitempty"`
	DBInstanceCreateTime        *string                                                                 `json:"DBInstanceCreateTime,omitempty" xml:"DBInstanceCreateTime,omitempty"`
	DBInstanceDescription       *string                                                                 `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	DBInstanceId                *string                                                                 `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceName              *string                                                                 `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DBInstanceStatus            *int32                                                                  `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	DBInstanceStatusDescription *string                                                                 `json:"DBInstanceStatusDescription,omitempty" xml:"DBInstanceStatusDescription,omitempty"`
	DBInstanceStorage           *int32                                                                  `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	Engine                      *string                                                                 `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion               *string                                                                 `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	Items                       []*DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	LockMode                    *int32                                                                  `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	LockReason                  *string                                                                 `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	MaintainEndTime             *string                                                                 `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	MaintainStartTime           *string                                                                 `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
}

func (s DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetDBInstanceConnType(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.DBInstanceConnType = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetDBInstanceCreateTime(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.DBInstanceCreateTime = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetDBInstanceDescription(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetDBInstanceId(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetDBInstanceName(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetDBInstanceStatus(v int32) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetDBInstanceStatusDescription(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.DBInstanceStatusDescription = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetDBInstanceStorage(v int32) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.DBInstanceStorage = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetEngine(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetEngineVersion(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetItems(v []*DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.Items = v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetLockMode(v int32) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetLockReason(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.LockReason = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetMaintainEndTime(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetMaintainStartTime(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.MaintainStartTime = &v
	return s
}

type DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems struct {
	CharacterType               *string                                                                             `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	ConnectionIp                []*DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsConnectionIp `json:"ConnectionIp,omitempty" xml:"ConnectionIp,omitempty" type:"Repeated"`
	DBInstanceConnType          *int32                                                                              `json:"DBInstanceConnType,omitempty" xml:"DBInstanceConnType,omitempty"`
	DBInstanceCreateTime        *string                                                                             `json:"DBInstanceCreateTime,omitempty" xml:"DBInstanceCreateTime,omitempty"`
	DBInstanceDescription       *string                                                                             `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	DBInstanceId                *string                                                                             `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceName              *string                                                                             `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DBInstanceStatus            *int32                                                                              `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	DBInstanceStatusDescription *string                                                                             `json:"DBInstanceStatusDescription,omitempty" xml:"DBInstanceStatusDescription,omitempty"`
	DiskSize                    *int64                                                                              `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	Engine                      *string                                                                             `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion               *string                                                                             `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	LockMode                    *int32                                                                              `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	LockReason                  *string                                                                             `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	MaintainEndTime             *string                                                                             `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	MaintainStartTime           *string                                                                             `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	MaxConnections              *int32                                                                              `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	MaxIops                     *int32                                                                              `json:"MaxIops,omitempty" xml:"MaxIops,omitempty"`
}

func (s DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetCharacterType(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.CharacterType = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetConnectionIp(v []*DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsConnectionIp) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.ConnectionIp = v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetDBInstanceConnType(v int32) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.DBInstanceConnType = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetDBInstanceCreateTime(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.DBInstanceCreateTime = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetDBInstanceDescription(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetDBInstanceId(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetDBInstanceName(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetDBInstanceStatus(v int32) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetDBInstanceStatusDescription(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.DBInstanceStatusDescription = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetDiskSize(v int64) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.DiskSize = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetEngine(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetEngineVersion(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetLockMode(v int32) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetLockReason(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.LockReason = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetMaintainEndTime(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetMaintainStartTime(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetMaxConnections(v int32) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.MaxConnections = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetMaxIops(v int32) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.MaxIops = &v
	return s
}

type DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsConnectionIp struct {
	ConnectionString  *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	DBInstanceNetType *int32  `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	Port              *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsConnectionIp) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsConnectionIp) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsConnectionIp) SetConnectionString(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsConnectionIp {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsConnectionIp) SetDBInstanceNetType(v int32) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsConnectionIp {
	s.DBInstanceNetType = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsConnectionIp) SetPort(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsConnectionIp {
	s.Port = &v
	return s
}

type DescribeDBInstanceTopologyResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBInstanceTopologyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceTopologyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceTopologyResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTopologyResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceTopologyResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceTopologyResponse) SetBody(v *DescribeDBInstanceTopologyResponseBody) *DescribeDBInstanceTopologyResponse {
	s.Body = v
	return s
}

type DescribeDBInstancesRequest struct {
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesRequest) SetPageNumber(v int32) *DescribeDBInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetPageSize(v int32) *DescribeDBInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetRegionId(v string) *DescribeDBInstancesRequest {
	s.RegionId = &v
	return s
}

type DescribeDBInstancesResponseBody struct {
	DBInstances []*DescribeDBInstancesResponseBodyDBInstances `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Repeated"`
	PageNumber  *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalNumber *int32                                        `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
}

func (s DescribeDBInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBody) SetDBInstances(v []*DescribeDBInstancesResponseBodyDBInstances) *DescribeDBInstancesResponseBody {
	s.DBInstances = v
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

func (s *DescribeDBInstancesResponseBody) SetRequestId(v string) *DescribeDBInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetTotalNumber(v int32) *DescribeDBInstancesResponseBody {
	s.TotalNumber = &v
	return s
}

type DescribeDBInstancesResponseBodyDBInstances struct {
	CommodityCode   *string                                            `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CreateTime      *string                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DBType          *string                                            `json:"DBType,omitempty" xml:"DBType,omitempty"`
	DBVersion       *string                                            `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	Description     *string                                            `json:"Description,omitempty" xml:"Description,omitempty"`
	Engine          *string                                            `json:"Engine,omitempty" xml:"Engine,omitempty"`
	ExpireTime      *string                                            `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Expired         *bool                                              `json:"Expired,omitempty" xml:"Expired,omitempty"`
	Id              *string                                            `json:"Id,omitempty" xml:"Id,omitempty"`
	LockMode        *string                                            `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	LockReason      *string                                            `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	MinorVersion    *string                                            `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	Network         *string                                            `json:"Network,omitempty" xml:"Network,omitempty"`
	NodeClass       *string                                            `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	NodeCount       *int32                                             `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	Nodes           []*DescribeDBInstancesResponseBodyDBInstancesNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	PayType         *string                                            `json:"PayType,omitempty" xml:"PayType,omitempty"`
	ReadDBInstances []*string                                          `json:"ReadDBInstances,omitempty" xml:"ReadDBInstances,omitempty" type:"Repeated"`
	RegionId        *string                                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status          *string                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageUsed     *int64                                             `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	Type            *string                                            `json:"Type,omitempty" xml:"Type,omitempty"`
	VPCId           *string                                            `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	ZoneId          *string                                            `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstancesResponseBodyDBInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyDBInstances) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetCommodityCode(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetCreateTime(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetDBType(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.DBType = &v
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

func (s *DescribeDBInstancesResponseBodyDBInstances) SetEngine(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.Engine = &v
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

func (s *DescribeDBInstancesResponseBodyDBInstances) SetId(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.Id = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetLockMode(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetLockReason(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.LockReason = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetMinorVersion(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.MinorVersion = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetNetwork(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.Network = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetNodeClass(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.NodeClass = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetNodeCount(v int32) *DescribeDBInstancesResponseBodyDBInstances {
	s.NodeCount = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetNodes(v []*DescribeDBInstancesResponseBodyDBInstancesNodes) *DescribeDBInstancesResponseBodyDBInstances {
	s.Nodes = v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetPayType(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.PayType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetReadDBInstances(v []*string) *DescribeDBInstancesResponseBodyDBInstances {
	s.ReadDBInstances = v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetRegionId(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetStatus(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.Status = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetStorageUsed(v int64) *DescribeDBInstancesResponseBodyDBInstances {
	s.StorageUsed = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetType(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.Type = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetVPCId(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.VPCId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetZoneId(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.ZoneId = &v
	return s
}

type DescribeDBInstancesResponseBodyDBInstancesNodes struct {
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstancesResponseBodyDBInstancesNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyDBInstancesNodes) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyDBInstancesNodes) SetClassCode(v string) *DescribeDBInstancesResponseBodyDBInstancesNodes {
	s.ClassCode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesNodes) SetId(v string) *DescribeDBInstancesResponseBodyDBInstancesNodes {
	s.Id = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesNodes) SetRegionId(v string) *DescribeDBInstancesResponseBodyDBInstancesNodes {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesNodes) SetZoneId(v string) *DescribeDBInstancesResponseBodyDBInstancesNodes {
	s.ZoneId = &v
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

type DescribeDBNodePerformanceRequest struct {
	CharacterType  *string `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DBNodeIds      *string `json:"DBNodeIds,omitempty" xml:"DBNodeIds,omitempty"`
	DBNodeRole     *string `json:"DBNodeRole,omitempty" xml:"DBNodeRole,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Key            *string `json:"Key,omitempty" xml:"Key,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBNodePerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBNodePerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceRequest) SetCharacterType(v string) *DescribeDBNodePerformanceRequest {
	s.CharacterType = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) SetDBInstanceName(v string) *DescribeDBNodePerformanceRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) SetDBNodeIds(v string) *DescribeDBNodePerformanceRequest {
	s.DBNodeIds = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) SetDBNodeRole(v string) *DescribeDBNodePerformanceRequest {
	s.DBNodeRole = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) SetEndTime(v string) *DescribeDBNodePerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) SetKey(v string) *DescribeDBNodePerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) SetRegionId(v string) *DescribeDBNodePerformanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) SetStartTime(v string) *DescribeDBNodePerformanceRequest {
	s.StartTime = &v
	return s
}

type DescribeDBNodePerformanceResponseBody struct {
	DBInstanceName  *string                                               `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	EndTime         *string                                               `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PerformanceKeys *DescribeDBNodePerformanceResponseBodyPerformanceKeys `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Struct"`
	RequestId       *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime       *string                                               `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBNodePerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBNodePerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceResponseBody) SetDBInstanceName(v string) *DescribeDBNodePerformanceResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBNodePerformanceResponseBody) SetEndTime(v string) *DescribeDBNodePerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDBNodePerformanceResponseBody) SetPerformanceKeys(v *DescribeDBNodePerformanceResponseBodyPerformanceKeys) *DescribeDBNodePerformanceResponseBody {
	s.PerformanceKeys = v
	return s
}

func (s *DescribeDBNodePerformanceResponseBody) SetRequestId(v string) *DescribeDBNodePerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBNodePerformanceResponseBody) SetStartTime(v string) *DescribeDBNodePerformanceResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDBNodePerformanceResponseBodyPerformanceKeys struct {
	PerformanceItem []*DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem `json:"PerformanceItem,omitempty" xml:"PerformanceItem,omitempty" type:"Repeated"`
}

func (s DescribeDBNodePerformanceResponseBodyPerformanceKeys) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBNodePerformanceResponseBodyPerformanceKeys) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeys) SetPerformanceItem(v []*DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem) *DescribeDBNodePerformanceResponseBodyPerformanceKeys {
	s.PerformanceItem = v
	return s
}

type DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem struct {
	DBNodeId    *string                                                                    `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	Measurement *string                                                                    `json:"Measurement,omitempty" xml:"Measurement,omitempty"`
	MetricName  *string                                                                    `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Points      *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPoints `json:"Points,omitempty" xml:"Points,omitempty" type:"Struct"`
}

func (s DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem) SetDBNodeId(v string) *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.DBNodeId = &v
	return s
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem) SetMeasurement(v string) *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.Measurement = &v
	return s
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem) SetMetricName(v string) *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.MetricName = &v
	return s
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem) SetPoints(v *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPoints) *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.Points = v
	return s
}

type DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPoints struct {
	PerformanceItemValue []*DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue `json:"PerformanceItemValue,omitempty" xml:"PerformanceItemValue,omitempty" type:"Repeated"`
}

func (s DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPoints) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPoints) SetPerformanceItemValue(v []*DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPoints {
	s.PerformanceItemValue = v
	return s
}

type DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue struct {
	Timestamp *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) SetTimestamp(v int64) *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue {
	s.Timestamp = &v
	return s
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) SetValue(v string) *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue {
	s.Value = &v
	return s
}

type DescribeDBNodePerformanceResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBNodePerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBNodePerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBNodePerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceResponse) SetHeaders(v map[string]*string) *DescribeDBNodePerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBNodePerformanceResponse) SetBody(v *DescribeDBNodePerformanceResponseBody) *DescribeDBNodePerformanceResponse {
	s.Body = v
	return s
}

type DescribeDbListRequest struct {
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DBName         *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDbListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDbListRequest) SetDBInstanceName(v string) *DescribeDbListRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDbListRequest) SetDBName(v string) *DescribeDbListRequest {
	s.DBName = &v
	return s
}

func (s *DescribeDbListRequest) SetRegionId(v string) *DescribeDbListRequest {
	s.RegionId = &v
	return s
}

type DescribeDbListResponseBody struct {
	Data      []*DescribeDbListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDbListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDbListResponseBody) SetData(v []*DescribeDbListResponseBodyData) *DescribeDbListResponseBody {
	s.Data = v
	return s
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

type DescribeDbListResponseBodyData struct {
	Accounts         []*DescribeDbListResponseBodyDataAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	CharacterSetName *string                                   `json:"CharacterSetName,omitempty" xml:"CharacterSetName,omitempty"`
	DBDescription    *string                                   `json:"DBDescription,omitempty" xml:"DBDescription,omitempty"`
	DBInstanceName   *string                                   `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DBName           *string                                   `json:"DBName,omitempty" xml:"DBName,omitempty"`
}

func (s DescribeDbListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDbListResponseBodyData) SetAccounts(v []*DescribeDbListResponseBodyDataAccounts) *DescribeDbListResponseBodyData {
	s.Accounts = v
	return s
}

func (s *DescribeDbListResponseBodyData) SetCharacterSetName(v string) *DescribeDbListResponseBodyData {
	s.CharacterSetName = &v
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

func (s *DescribeDbListResponseBodyData) SetDBName(v string) *DescribeDbListResponseBodyData {
	s.DBName = &v
	return s
}

type DescribeDbListResponseBodyDataAccounts struct {
	AccountName      *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
}

func (s DescribeDbListResponseBodyDataAccounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbListResponseBodyDataAccounts) GoString() string {
	return s.String()
}

func (s *DescribeDbListResponseBodyDataAccounts) SetAccountName(v string) *DescribeDbListResponseBodyDataAccounts {
	s.AccountName = &v
	return s
}

func (s *DescribeDbListResponseBodyDataAccounts) SetAccountPrivilege(v string) *DescribeDbListResponseBodyDataAccounts {
	s.AccountPrivilege = &v
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
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDistributeTableListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDistributeTableListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDistributeTableListRequest) SetDBInstanceName(v string) *DescribeDistributeTableListRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDistributeTableListRequest) SetDbName(v string) *DescribeDistributeTableListRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDistributeTableListRequest) SetRegionId(v string) *DescribeDistributeTableListRequest {
	s.RegionId = &v
	return s
}

type DescribeDistributeTableListResponseBody struct {
	Data      *DescribeDistributeTableListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDistributeTableListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDistributeTableListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDistributeTableListResponseBody) SetData(v *DescribeDistributeTableListResponseBodyData) *DescribeDistributeTableListResponseBody {
	s.Data = v
	return s
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
	DbKey     *string `json:"DbKey,omitempty" xml:"DbKey,omitempty"`
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TableType *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
	TbKey     *string `json:"TbKey,omitempty" xml:"TbKey,omitempty"`
}

func (s DescribeDistributeTableListResponseBodyDataTables) String() string {
	return tea.Prettify(s)
}

func (s DescribeDistributeTableListResponseBodyDataTables) GoString() string {
	return s.String()
}

func (s *DescribeDistributeTableListResponseBodyDataTables) SetDbKey(v string) *DescribeDistributeTableListResponseBodyDataTables {
	s.DbKey = &v
	return s
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

type DescribeParameterTemplatesRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ParamLevel   *string `json:"ParamLevel,omitempty" xml:"ParamLevel,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeParameterTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeParameterTemplatesRequest) SetDBInstanceId(v string) *DescribeParameterTemplatesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetParamLevel(v string) *DescribeParameterTemplatesRequest {
	s.ParamLevel = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetRegionId(v string) *DescribeParameterTemplatesRequest {
	s.RegionId = &v
	return s
}

type DescribeParameterTemplatesResponseBody struct {
	Data      *DescribeParameterTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeParameterTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParameterTemplatesResponseBody) SetData(v *DescribeParameterTemplatesResponseBodyData) *DescribeParameterTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeParameterTemplatesResponseBody) SetRequestId(v string) *DescribeParameterTemplatesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeParameterTemplatesResponseBodyData struct {
	Engine         *string                                                 `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion  *string                                                 `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ParameterCount *int32                                                  `json:"ParameterCount,omitempty" xml:"ParameterCount,omitempty"`
	Parameters     []*DescribeParameterTemplatesResponseBodyDataParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s DescribeParameterTemplatesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeParameterTemplatesResponseBodyData) SetEngine(v string) *DescribeParameterTemplatesResponseBodyData {
	s.Engine = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyData) SetEngineVersion(v string) *DescribeParameterTemplatesResponseBodyData {
	s.EngineVersion = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyData) SetParameterCount(v int32) *DescribeParameterTemplatesResponseBodyData {
	s.ParameterCount = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyData) SetParameters(v []*DescribeParameterTemplatesResponseBodyDataParameters) *DescribeParameterTemplatesResponseBodyData {
	s.Parameters = v
	return s
}

type DescribeParameterTemplatesResponseBodyDataParameters struct {
	CheckingCode         *string `json:"CheckingCode,omitempty" xml:"CheckingCode,omitempty"`
	Dynamic              *int32  `json:"Dynamic,omitempty" xml:"Dynamic,omitempty"`
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	ParameterName        *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ParameterValue       *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
	Revisable            *int32  `json:"Revisable,omitempty" xml:"Revisable,omitempty"`
}

func (s DescribeParameterTemplatesResponseBodyDataParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterTemplatesResponseBodyDataParameters) GoString() string {
	return s.String()
}

func (s *DescribeParameterTemplatesResponseBodyDataParameters) SetCheckingCode(v string) *DescribeParameterTemplatesResponseBodyDataParameters {
	s.CheckingCode = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyDataParameters) SetDynamic(v int32) *DescribeParameterTemplatesResponseBodyDataParameters {
	s.Dynamic = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyDataParameters) SetParameterDescription(v string) *DescribeParameterTemplatesResponseBodyDataParameters {
	s.ParameterDescription = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyDataParameters) SetParameterName(v string) *DescribeParameterTemplatesResponseBodyDataParameters {
	s.ParameterName = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyDataParameters) SetParameterValue(v string) *DescribeParameterTemplatesResponseBodyDataParameters {
	s.ParameterValue = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyDataParameters) SetRevisable(v int32) *DescribeParameterTemplatesResponseBodyDataParameters {
	s.Revisable = &v
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

type DescribeParametersRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ParamLevel   *string `json:"ParamLevel,omitempty" xml:"ParamLevel,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersRequest) GoString() string {
	return s.String()
}

func (s *DescribeParametersRequest) SetDBInstanceId(v string) *DescribeParametersRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeParametersRequest) SetParamLevel(v string) *DescribeParametersRequest {
	s.ParamLevel = &v
	return s
}

func (s *DescribeParametersRequest) SetRegionId(v string) *DescribeParametersRequest {
	s.RegionId = &v
	return s
}

type DescribeParametersResponseBody struct {
	Data      *DescribeParametersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBody) SetData(v *DescribeParametersResponseBodyData) *DescribeParametersResponseBody {
	s.Data = v
	return s
}

func (s *DescribeParametersResponseBody) SetRequestId(v string) *DescribeParametersResponseBody {
	s.RequestId = &v
	return s
}

type DescribeParametersResponseBodyData struct {
	ConfigParameters  []*DescribeParametersResponseBodyDataConfigParameters  `json:"ConfigParameters,omitempty" xml:"ConfigParameters,omitempty" type:"Repeated"`
	Engine            *string                                                `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion     *string                                                `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	RunningParameters []*DescribeParametersResponseBodyDataRunningParameters `json:"RunningParameters,omitempty" xml:"RunningParameters,omitempty" type:"Repeated"`
}

func (s DescribeParametersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBodyData) SetConfigParameters(v []*DescribeParametersResponseBodyDataConfigParameters) *DescribeParametersResponseBodyData {
	s.ConfigParameters = v
	return s
}

func (s *DescribeParametersResponseBodyData) SetEngine(v string) *DescribeParametersResponseBodyData {
	s.Engine = &v
	return s
}

func (s *DescribeParametersResponseBodyData) SetEngineVersion(v string) *DescribeParametersResponseBodyData {
	s.EngineVersion = &v
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

type DescribePolarxDataNodesRequest struct {
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribePolarxDataNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePolarxDataNodesRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarxDataNodesRequest) SetPageNumber(v int32) *DescribePolarxDataNodesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePolarxDataNodesRequest) SetPageSize(v int32) *DescribePolarxDataNodesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePolarxDataNodesRequest) SetRegionId(v string) *DescribePolarxDataNodesRequest {
	s.RegionId = &v
	return s
}

type DescribePolarxDataNodesResponseBody struct {
	DBInstanceDataNodes []*DescribePolarxDataNodesResponseBodyDBInstanceDataNodes `json:"DBInstanceDataNodes,omitempty" xml:"DBInstanceDataNodes,omitempty" type:"Repeated"`
	PageNumber          *int32                                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize            *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId           *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalNumber         *int32                                                    `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
}

func (s DescribePolarxDataNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePolarxDataNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarxDataNodesResponseBody) SetDBInstanceDataNodes(v []*DescribePolarxDataNodesResponseBodyDBInstanceDataNodes) *DescribePolarxDataNodesResponseBody {
	s.DBInstanceDataNodes = v
	return s
}

func (s *DescribePolarxDataNodesResponseBody) SetPageNumber(v int32) *DescribePolarxDataNodesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePolarxDataNodesResponseBody) SetPageSize(v int32) *DescribePolarxDataNodesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePolarxDataNodesResponseBody) SetRequestId(v string) *DescribePolarxDataNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarxDataNodesResponseBody) SetTotalNumber(v int32) *DescribePolarxDataNodesResponseBody {
	s.TotalNumber = &v
	return s
}

type DescribePolarxDataNodesResponseBodyDBInstanceDataNodes struct {
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	DBInstanceId          *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceName        *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
}

func (s DescribePolarxDataNodesResponseBodyDBInstanceDataNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribePolarxDataNodesResponseBodyDBInstanceDataNodes) GoString() string {
	return s.String()
}

func (s *DescribePolarxDataNodesResponseBodyDBInstanceDataNodes) SetDBInstanceDescription(v string) *DescribePolarxDataNodesResponseBodyDBInstanceDataNodes {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribePolarxDataNodesResponseBodyDBInstanceDataNodes) SetDBInstanceId(v string) *DescribePolarxDataNodesResponseBodyDBInstanceDataNodes {
	s.DBInstanceId = &v
	return s
}

func (s *DescribePolarxDataNodesResponseBodyDBInstanceDataNodes) SetDBInstanceName(v string) *DescribePolarxDataNodesResponseBodyDBInstanceDataNodes {
	s.DBInstanceName = &v
	return s
}

type DescribePolarxDataNodesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePolarxDataNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePolarxDataNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePolarxDataNodesResponse) GoString() string {
	return s.String()
}

func (s *DescribePolarxDataNodesResponse) SetHeaders(v map[string]*string) *DescribePolarxDataNodesResponse {
	s.Headers = v
	return s
}

func (s *DescribePolarxDataNodesResponse) SetBody(v *DescribePolarxDataNodesResponseBody) *DescribePolarxDataNodesResponse {
	s.Body = v
	return s
}

type DescribePolarxDbInstancesRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribePolarxDbInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePolarxDbInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarxDbInstancesRequest) SetDbName(v string) *DescribePolarxDbInstancesRequest {
	s.DbName = &v
	return s
}

func (s *DescribePolarxDbInstancesRequest) SetDrdsInstanceId(v string) *DescribePolarxDbInstancesRequest {
	s.DrdsInstanceId = &v
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
	DbInstances *DescribePolarxDbInstancesResponseBodyDbInstances `json:"DbInstances,omitempty" xml:"DbInstances,omitempty" type:"Struct"`
	PageNumber  *string                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *string                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
	Total       *string                                           `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribePolarxDbInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePolarxDbInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarxDbInstancesResponseBody) SetDbInstances(v *DescribePolarxDbInstancesResponseBodyDbInstances) *DescribePolarxDbInstancesResponseBody {
	s.DbInstances = v
	return s
}

func (s *DescribePolarxDbInstancesResponseBody) SetPageNumber(v string) *DescribePolarxDbInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBody) SetPageSize(v string) *DescribePolarxDbInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBody) SetRequestId(v string) *DescribePolarxDbInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBody) SetSuccess(v bool) *DescribePolarxDbInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBody) SetTotal(v string) *DescribePolarxDbInstancesResponseBody {
	s.Total = &v
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
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBType       *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	DBVersion    *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Engine       *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	ExpireTime   *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	LockMode     *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	Network      *string `json:"Network,omitempty" xml:"Network,omitempty"`
	NodeClass    *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	NodeCount    *int32  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	PayType      *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusDesc   *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	StorageUsed  *int32  `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	VPCId        *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	ZoneId       *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	LockReason   *string `json:"lockReason,omitempty" xml:"lockReason,omitempty"`
}

func (s DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) GoString() string {
	return s.String()
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetCreateTime(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceId(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetDBType(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.DBType = &v
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

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetEngine(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.Engine = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetExpireTime(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetLockMode(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.LockMode = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetNetwork(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.Network = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetNodeClass(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.NodeClass = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetNodeCount(v int32) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.NodeCount = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetPayType(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.PayType = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetRegionId(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.RegionId = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetStatus(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.Status = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetStatusDesc(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.StatusDesc = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetStorageUsed(v int32) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.StorageUsed = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetVPCId(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.VPCId = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetZoneId(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance) SetLockReason(v string) *DescribePolarxDbInstancesResponseBodyDbInstancesDbInstance {
	s.LockReason = &v
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
	ErrorCode *int32                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *DescribeRegionsResponseBody) SetErrorCode(v int32) *DescribeRegionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetMessage(v string) *DescribeRegionsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
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
	RegionId        *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SupportPolarx10 *bool                                          `json:"SupportPolarx10,omitempty" xml:"SupportPolarx10,omitempty"`
	SupportPolarx20 *bool                                          `json:"SupportPolarx20,omitempty" xml:"SupportPolarx20,omitempty"`
	Zones           *DescribeRegionsResponseBodyRegionsRegionZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetSupportPolarx10(v bool) *DescribeRegionsResponseBodyRegionsRegion {
	s.SupportPolarx10 = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetSupportPolarx20(v bool) *DescribeRegionsResponseBodyRegionsRegion {
	s.SupportPolarx20 = &v
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
	VpcEnabled *bool   `json:"VpcEnabled,omitempty" xml:"VpcEnabled,omitempty"`
	ZoneId     *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegionZonesZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegionZonesZone) GoString() string {
	return s.String()
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

type DescribeScaleOutMigrateTaskListRequest struct {
	DBInstanceName       *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeScaleOutMigrateTaskListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScaleOutMigrateTaskListRequest) GoString() string {
	return s.String()
}

func (s *DescribeScaleOutMigrateTaskListRequest) SetDBInstanceName(v string) *DescribeScaleOutMigrateTaskListRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeScaleOutMigrateTaskListRequest) SetOwnerAccount(v string) *DescribeScaleOutMigrateTaskListRequest {
	s.OwnerAccount = &v
	return s
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
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSecurityIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIpsRequest) SetDBInstanceName(v string) *DescribeSecurityIpsRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeSecurityIpsRequest) SetRegionId(v string) *DescribeSecurityIpsRequest {
	s.RegionId = &v
	return s
}

type DescribeSecurityIpsResponseBody struct {
	Data      *DescribeSecurityIpsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSecurityIpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityIpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIpsResponseBody) SetData(v *DescribeSecurityIpsResponseBodyData) *DescribeSecurityIpsResponseBody {
	s.Data = v
	return s
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

type DescribeTableDetailRequest struct {
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TableName      *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeTableDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeTableDetailRequest) SetDBInstanceName(v string) *DescribeTableDetailRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeTableDetailRequest) SetDbName(v string) *DescribeTableDetailRequest {
	s.DbName = &v
	return s
}

func (s *DescribeTableDetailRequest) SetRegionId(v string) *DescribeTableDetailRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTableDetailRequest) SetTableName(v string) *DescribeTableDetailRequest {
	s.TableName = &v
	return s
}

type DescribeTableDetailResponseBody struct {
	Data      *DescribeTableDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeTableDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTableDetailResponseBody) SetData(v *DescribeTableDetailResponseBodyData) *DescribeTableDetailResponseBody {
	s.Data = v
	return s
}

func (s *DescribeTableDetailResponseBody) SetMessage(v string) *DescribeTableDetailResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTableDetailResponseBody) SetRequestId(v string) *DescribeTableDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTableDetailResponseBody) SetSuccess(v bool) *DescribeTableDetailResponseBody {
	s.Success = &v
	return s
}

type DescribeTableDetailResponseBodyData struct {
	Fields []*DescribeTableDetailResponseBodyDataFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
}

func (s DescribeTableDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeTableDetailResponseBodyData) SetFields(v []*DescribeTableDetailResponseBodyDataFields) *DescribeTableDetailResponseBodyData {
	s.Fields = v
	return s
}

type DescribeTableDetailResponseBodyDataFields struct {
	Column   *string `json:"Column,omitempty" xml:"Column,omitempty"`
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Extra    *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s DescribeTableDetailResponseBodyDataFields) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableDetailResponseBodyDataFields) GoString() string {
	return s.String()
}

func (s *DescribeTableDetailResponseBodyDataFields) SetColumn(v string) *DescribeTableDetailResponseBodyDataFields {
	s.Column = &v
	return s
}

func (s *DescribeTableDetailResponseBodyDataFields) SetDataType(v string) *DescribeTableDetailResponseBodyDataFields {
	s.DataType = &v
	return s
}

func (s *DescribeTableDetailResponseBodyDataFields) SetExtra(v string) *DescribeTableDetailResponseBodyDataFields {
	s.Extra = &v
	return s
}

func (s *DescribeTableDetailResponseBodyDataFields) SetKey(v string) *DescribeTableDetailResponseBodyDataFields {
	s.Key = &v
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

type DescribeTasksRequest struct {
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeTasksRequest) SetDBInstanceId(v string) *DescribeTasksRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeTasksRequest) SetEndTime(v string) *DescribeTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTasksRequest) SetOwnerAccount(v string) *DescribeTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTasksRequest) SetOwnerId(v int64) *DescribeTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTasksRequest) SetPageNumber(v int32) *DescribeTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTasksRequest) SetPageSize(v int32) *DescribeTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTasksRequest) SetRegionId(v string) *DescribeTasksRequest {
	s.RegionId = &v
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

func (s *DescribeTasksRequest) SetStartTime(v string) *DescribeTasksRequest {
	s.StartTime = &v
	return s
}

type DescribeTasksResponseBody struct {
	Items            []*DescribeTasksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber       *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount  *int32                            `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalRecordCount *int32                            `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBody) SetItems(v []*DescribeTasksResponseBodyItems) *DescribeTasksResponseBody {
	s.Items = v
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

func (s *DescribeTasksResponseBody) SetRequestId(v string) *DescribeTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTasksResponseBody) SetTotalRecordCount(v int32) *DescribeTasksResponseBody {
	s.TotalRecordCount = &v
	return s
}

type DescribeTasksResponseBodyItems struct {
	BeginTime        *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	DBName           *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	FinishTime       *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	Progress         *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	ProgressInfo     *string `json:"ProgressInfo,omitempty" xml:"ProgressInfo,omitempty"`
	ScaleOutToken    *string `json:"ScaleOutToken,omitempty" xml:"ScaleOutToken,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskAction       *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	TaskErrorCode    *string `json:"TaskErrorCode,omitempty" xml:"TaskErrorCode,omitempty"`
	TaskErrorMessage *string `json:"TaskErrorMessage,omitempty" xml:"TaskErrorMessage,omitempty"`
	TaskId           *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeTasksResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBodyItems) SetBeginTime(v string) *DescribeTasksResponseBodyItems {
	s.BeginTime = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetDBName(v string) *DescribeTasksResponseBodyItems {
	s.DBName = &v
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

func (s *DescribeTasksResponseBodyItems) SetProgressInfo(v string) *DescribeTasksResponseBodyItems {
	s.ProgressInfo = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetScaleOutToken(v string) *DescribeTasksResponseBodyItems {
	s.ScaleOutToken = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetStatus(v string) *DescribeTasksResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetTaskAction(v string) *DescribeTasksResponseBodyItems {
	s.TaskAction = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetTaskErrorCode(v string) *DescribeTasksResponseBodyItems {
	s.TaskErrorCode = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetTaskErrorMessage(v string) *DescribeTasksResponseBodyItems {
	s.TaskErrorMessage = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetTaskId(v string) *DescribeTasksResponseBodyItems {
	s.TaskId = &v
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
	Data      *DescribeUserEncryptionKeyListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserEncryptionKeyListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserEncryptionKeyListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListResponseBody) SetData(v *DescribeUserEncryptionKeyListResponseBodyData) *DescribeUserEncryptionKeyListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeUserEncryptionKeyListResponseBody) SetRequestId(v string) *DescribeUserEncryptionKeyListResponseBody {
	s.RequestId = &v
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

type DisableRightsSeparationRequest struct {
	DBInstanceName     *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DbaAccountName     *string `json:"DbaAccountName,omitempty" xml:"DbaAccountName,omitempty"`
	DbaAccountPassword *string `json:"DbaAccountPassword,omitempty" xml:"DbaAccountPassword,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisableRightsSeparationRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableRightsSeparationRequest) GoString() string {
	return s.String()
}

func (s *DisableRightsSeparationRequest) SetDBInstanceName(v string) *DisableRightsSeparationRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DisableRightsSeparationRequest) SetDbaAccountName(v string) *DisableRightsSeparationRequest {
	s.DbaAccountName = &v
	return s
}

func (s *DisableRightsSeparationRequest) SetDbaAccountPassword(v string) *DisableRightsSeparationRequest {
	s.DbaAccountPassword = &v
	return s
}

func (s *DisableRightsSeparationRequest) SetRegionId(v string) *DisableRightsSeparationRequest {
	s.RegionId = &v
	return s
}

type DisableRightsSeparationResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableRightsSeparationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableRightsSeparationResponseBody) GoString() string {
	return s.String()
}

func (s *DisableRightsSeparationResponseBody) SetMessage(v string) *DisableRightsSeparationResponseBody {
	s.Message = &v
	return s
}

func (s *DisableRightsSeparationResponseBody) SetRequestId(v string) *DisableRightsSeparationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableRightsSeparationResponseBody) SetSuccess(v bool) *DisableRightsSeparationResponseBody {
	s.Success = &v
	return s
}

type DisableRightsSeparationResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableRightsSeparationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableRightsSeparationResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableRightsSeparationResponse) GoString() string {
	return s.String()
}

func (s *DisableRightsSeparationResponse) SetHeaders(v map[string]*string) *DisableRightsSeparationResponse {
	s.Headers = v
	return s
}

func (s *DisableRightsSeparationResponse) SetBody(v *DisableRightsSeparationResponseBody) *DisableRightsSeparationResponse {
	s.Body = v
	return s
}

type EnableRightsSeparationRequest struct {
	AuditAccountDescription    *string `json:"AuditAccountDescription,omitempty" xml:"AuditAccountDescription,omitempty"`
	AuditAccountName           *string `json:"AuditAccountName,omitempty" xml:"AuditAccountName,omitempty"`
	AuditAccountPassword       *string `json:"AuditAccountPassword,omitempty" xml:"AuditAccountPassword,omitempty"`
	DBInstanceName             *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RegionId                   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityAccountDescription *string `json:"SecurityAccountDescription,omitempty" xml:"SecurityAccountDescription,omitempty"`
	SecurityAccountName        *string `json:"SecurityAccountName,omitempty" xml:"SecurityAccountName,omitempty"`
	SecurityAccountPassword    *string `json:"SecurityAccountPassword,omitempty" xml:"SecurityAccountPassword,omitempty"`
}

func (s EnableRightsSeparationRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableRightsSeparationRequest) GoString() string {
	return s.String()
}

func (s *EnableRightsSeparationRequest) SetAuditAccountDescription(v string) *EnableRightsSeparationRequest {
	s.AuditAccountDescription = &v
	return s
}

func (s *EnableRightsSeparationRequest) SetAuditAccountName(v string) *EnableRightsSeparationRequest {
	s.AuditAccountName = &v
	return s
}

func (s *EnableRightsSeparationRequest) SetAuditAccountPassword(v string) *EnableRightsSeparationRequest {
	s.AuditAccountPassword = &v
	return s
}

func (s *EnableRightsSeparationRequest) SetDBInstanceName(v string) *EnableRightsSeparationRequest {
	s.DBInstanceName = &v
	return s
}

func (s *EnableRightsSeparationRequest) SetRegionId(v string) *EnableRightsSeparationRequest {
	s.RegionId = &v
	return s
}

func (s *EnableRightsSeparationRequest) SetSecurityAccountDescription(v string) *EnableRightsSeparationRequest {
	s.SecurityAccountDescription = &v
	return s
}

func (s *EnableRightsSeparationRequest) SetSecurityAccountName(v string) *EnableRightsSeparationRequest {
	s.SecurityAccountName = &v
	return s
}

func (s *EnableRightsSeparationRequest) SetSecurityAccountPassword(v string) *EnableRightsSeparationRequest {
	s.SecurityAccountPassword = &v
	return s
}

type EnableRightsSeparationResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableRightsSeparationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableRightsSeparationResponseBody) GoString() string {
	return s.String()
}

func (s *EnableRightsSeparationResponseBody) SetMessage(v string) *EnableRightsSeparationResponseBody {
	s.Message = &v
	return s
}

func (s *EnableRightsSeparationResponseBody) SetRequestId(v string) *EnableRightsSeparationResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableRightsSeparationResponseBody) SetSuccess(v bool) *EnableRightsSeparationResponseBody {
	s.Success = &v
	return s
}

type EnableRightsSeparationResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableRightsSeparationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableRightsSeparationResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableRightsSeparationResponse) GoString() string {
	return s.String()
}

func (s *EnableRightsSeparationResponse) SetHeaders(v map[string]*string) *EnableRightsSeparationResponse {
	s.Headers = v
	return s
}

func (s *EnableRightsSeparationResponse) SetBody(v *EnableRightsSeparationResponseBody) *EnableRightsSeparationResponse {
	s.Body = v
	return s
}

type GetPolarxCommodityRequest struct {
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	OrderType      *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetPolarxCommodityRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPolarxCommodityRequest) GoString() string {
	return s.String()
}

func (s *GetPolarxCommodityRequest) SetDBInstanceName(v string) *GetPolarxCommodityRequest {
	s.DBInstanceName = &v
	return s
}

func (s *GetPolarxCommodityRequest) SetOrderType(v string) *GetPolarxCommodityRequest {
	s.OrderType = &v
	return s
}

func (s *GetPolarxCommodityRequest) SetRegionId(v string) *GetPolarxCommodityRequest {
	s.RegionId = &v
	return s
}

type GetPolarxCommodityResponseBody struct {
	ComponentList []*GetPolarxCommodityResponseBodyComponentList `json:"ComponentList,omitempty" xml:"ComponentList,omitempty" type:"Repeated"`
	DBInstance    *GetPolarxCommodityResponseBodyDBInstance      `json:"DBInstance,omitempty" xml:"DBInstance,omitempty" type:"Struct"`
	RequestId     *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPolarxCommodityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPolarxCommodityResponseBody) GoString() string {
	return s.String()
}

func (s *GetPolarxCommodityResponseBody) SetComponentList(v []*GetPolarxCommodityResponseBodyComponentList) *GetPolarxCommodityResponseBody {
	s.ComponentList = v
	return s
}

func (s *GetPolarxCommodityResponseBody) SetDBInstance(v *GetPolarxCommodityResponseBodyDBInstance) *GetPolarxCommodityResponseBody {
	s.DBInstance = v
	return s
}

func (s *GetPolarxCommodityResponseBody) SetRequestId(v string) *GetPolarxCommodityResponseBody {
	s.RequestId = &v
	return s
}

type GetPolarxCommodityResponseBodyComponentList struct {
	Name   *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Type   *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s GetPolarxCommodityResponseBodyComponentList) String() string {
	return tea.Prettify(s)
}

func (s GetPolarxCommodityResponseBodyComponentList) GoString() string {
	return s.String()
}

func (s *GetPolarxCommodityResponseBodyComponentList) SetName(v string) *GetPolarxCommodityResponseBodyComponentList {
	s.Name = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyComponentList) SetType(v string) *GetPolarxCommodityResponseBodyComponentList {
	s.Type = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyComponentList) SetValues(v []*string) *GetPolarxCommodityResponseBodyComponentList {
	s.Values = v
	return s
}

type GetPolarxCommodityResponseBodyDBInstance struct {
	CommodityCode      *string                                              `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ConnAddrs          []*GetPolarxCommodityResponseBodyDBInstanceConnAddrs `json:"ConnAddrs,omitempty" xml:"ConnAddrs,omitempty" type:"Repeated"`
	ConnectionString   *string                                              `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	CreateTime         *string                                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DBInstanceType     *string                                              `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	DBNodeClass        *string                                              `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	DBNodeCount        *int32                                               `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	DBNodes            []*GetPolarxCommodityResponseBodyDBInstanceDBNodes   `json:"DBNodes,omitempty" xml:"DBNodes,omitempty" type:"Repeated"`
	DBType             *string                                              `json:"DBType,omitempty" xml:"DBType,omitempty"`
	DBVersion          *string                                              `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	Description        *string                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	Engine             *string                                              `json:"Engine,omitempty" xml:"Engine,omitempty"`
	ExpireDate         *string                                              `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	Expired            *string                                              `json:"Expired,omitempty" xml:"Expired,omitempty"`
	Id                 *string                                              `json:"Id,omitempty" xml:"Id,omitempty"`
	LatestMinorVersion *string                                              `json:"LatestMinorVersion,omitempty" xml:"LatestMinorVersion,omitempty"`
	LockMode           *string                                              `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	MaintainEndTime    *string                                              `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	MaintainStartTime  *string                                              `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	MinorVersion       *string                                              `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	Network            *string                                              `json:"Network,omitempty" xml:"Network,omitempty"`
	PayType            *string                                              `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Port               *string                                              `json:"Port,omitempty" xml:"Port,omitempty"`
	ReadDBInstances    []*string                                            `json:"ReadDBInstances,omitempty" xml:"ReadDBInstances,omitempty" type:"Repeated"`
	RegionId           *string                                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status             *string                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageUsed        *int64                                               `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	Type               *string                                              `json:"Type,omitempty" xml:"Type,omitempty"`
	VPCId              *string                                              `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId          *string                                              `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId             *string                                              `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetPolarxCommodityResponseBodyDBInstance) String() string {
	return tea.Prettify(s)
}

func (s GetPolarxCommodityResponseBodyDBInstance) GoString() string {
	return s.String()
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetCommodityCode(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.CommodityCode = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetConnAddrs(v []*GetPolarxCommodityResponseBodyDBInstanceConnAddrs) *GetPolarxCommodityResponseBodyDBInstance {
	s.ConnAddrs = v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetConnectionString(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.ConnectionString = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetCreateTime(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.CreateTime = &v
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

func (s *GetPolarxCommodityResponseBodyDBInstance) SetDBNodeCount(v int32) *GetPolarxCommodityResponseBodyDBInstance {
	s.DBNodeCount = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetDBNodes(v []*GetPolarxCommodityResponseBodyDBInstanceDBNodes) *GetPolarxCommodityResponseBodyDBInstance {
	s.DBNodes = v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetDBType(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.DBType = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetDBVersion(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.DBVersion = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetDescription(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.Description = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetEngine(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.Engine = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetExpireDate(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.ExpireDate = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetExpired(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.Expired = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetId(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.Id = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetLatestMinorVersion(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.LatestMinorVersion = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetLockMode(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.LockMode = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetMaintainEndTime(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.MaintainEndTime = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetMaintainStartTime(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.MaintainStartTime = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetMinorVersion(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.MinorVersion = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetNetwork(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.Network = &v
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

func (s *GetPolarxCommodityResponseBodyDBInstance) SetReadDBInstances(v []*string) *GetPolarxCommodityResponseBodyDBInstance {
	s.ReadDBInstances = v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetRegionId(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.RegionId = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetStatus(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.Status = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetStorageUsed(v int64) *GetPolarxCommodityResponseBodyDBInstance {
	s.StorageUsed = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetType(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.Type = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstance) SetVPCId(v string) *GetPolarxCommodityResponseBodyDBInstance {
	s.VPCId = &v
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

type GetPolarxCommodityResponseBodyDBInstanceConnAddrs struct {
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	Port             *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VPCId            *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId        *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s GetPolarxCommodityResponseBodyDBInstanceConnAddrs) String() string {
	return tea.Prettify(s)
}

func (s GetPolarxCommodityResponseBodyDBInstanceConnAddrs) GoString() string {
	return s.String()
}

func (s *GetPolarxCommodityResponseBodyDBInstanceConnAddrs) SetConnectionString(v string) *GetPolarxCommodityResponseBodyDBInstanceConnAddrs {
	s.ConnectionString = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstanceConnAddrs) SetPort(v string) *GetPolarxCommodityResponseBodyDBInstanceConnAddrs {
	s.Port = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstanceConnAddrs) SetType(v string) *GetPolarxCommodityResponseBodyDBInstanceConnAddrs {
	s.Type = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstanceConnAddrs) SetVPCId(v string) *GetPolarxCommodityResponseBodyDBInstanceConnAddrs {
	s.VPCId = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstanceConnAddrs) SetVSwitchId(v string) *GetPolarxCommodityResponseBodyDBInstanceConnAddrs {
	s.VSwitchId = &v
	return s
}

type GetPolarxCommodityResponseBodyDBInstanceDBNodes struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	NodeClass *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetPolarxCommodityResponseBodyDBInstanceDBNodes) String() string {
	return tea.Prettify(s)
}

func (s GetPolarxCommodityResponseBodyDBInstanceDBNodes) GoString() string {
	return s.String()
}

func (s *GetPolarxCommodityResponseBodyDBInstanceDBNodes) SetId(v string) *GetPolarxCommodityResponseBodyDBInstanceDBNodes {
	s.Id = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstanceDBNodes) SetNodeClass(v string) *GetPolarxCommodityResponseBodyDBInstanceDBNodes {
	s.NodeClass = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstanceDBNodes) SetRegionId(v string) *GetPolarxCommodityResponseBodyDBInstanceDBNodes {
	s.RegionId = &v
	return s
}

func (s *GetPolarxCommodityResponseBodyDBInstanceDBNodes) SetZoneId(v string) *GetPolarxCommodityResponseBodyDBInstanceDBNodes {
	s.ZoneId = &v
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

type ModifyAccountDescriptionRequest struct {
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	AccountName        *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DBInstanceName     *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *ModifyAccountDescriptionRequest) SetDBInstanceName(v string) *ModifyAccountDescriptionRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetRegionId(v string) *ModifyAccountDescriptionRequest {
	s.RegionId = &v
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

type ModifyAccountPrivilegeRequest struct {
	AccountName             *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPrivilege        *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	DBInstanceName          *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DbName                  *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityAccountName     *string `json:"SecurityAccountName,omitempty" xml:"SecurityAccountName,omitempty"`
	SecurityAccountPassword *string `json:"SecurityAccountPassword,omitempty" xml:"SecurityAccountPassword,omitempty"`
}

func (s ModifyAccountPrivilegeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountPrivilegeRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegeRequest) SetAccountName(v string) *ModifyAccountPrivilegeRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountPrivilegeRequest) SetAccountPrivilege(v string) *ModifyAccountPrivilegeRequest {
	s.AccountPrivilege = &v
	return s
}

func (s *ModifyAccountPrivilegeRequest) SetDBInstanceName(v string) *ModifyAccountPrivilegeRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyAccountPrivilegeRequest) SetDbName(v string) *ModifyAccountPrivilegeRequest {
	s.DbName = &v
	return s
}

func (s *ModifyAccountPrivilegeRequest) SetRegionId(v string) *ModifyAccountPrivilegeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAccountPrivilegeRequest) SetSecurityAccountName(v string) *ModifyAccountPrivilegeRequest {
	s.SecurityAccountName = &v
	return s
}

func (s *ModifyAccountPrivilegeRequest) SetSecurityAccountPassword(v string) *ModifyAccountPrivilegeRequest {
	s.SecurityAccountPassword = &v
	return s
}

type ModifyAccountPrivilegeResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyAccountPrivilegeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountPrivilegeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegeResponseBody) SetMessage(v string) *ModifyAccountPrivilegeResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyAccountPrivilegeResponseBody) SetRequestId(v string) *ModifyAccountPrivilegeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAccountPrivilegeResponseBody) SetSuccess(v bool) *ModifyAccountPrivilegeResponseBody {
	s.Success = &v
	return s
}

type ModifyAccountPrivilegeResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyAccountPrivilegeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAccountPrivilegeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountPrivilegeResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegeResponse) SetHeaders(v map[string]*string) *ModifyAccountPrivilegeResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccountPrivilegeResponse) SetBody(v *ModifyAccountPrivilegeResponseBody) *ModifyAccountPrivilegeResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceClassRequest struct {
	ClientToken           *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DBInstanceName        *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TargetDBInstanceClass *string `json:"TargetDBInstanceClass,omitempty" xml:"TargetDBInstanceClass,omitempty"`
}

func (s ModifyDBInstanceClassRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceClassRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceClassRequest) SetClientToken(v string) *ModifyDBInstanceClassRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetDBInstanceName(v string) *ModifyDBInstanceClassRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetRegionId(v string) *ModifyDBInstanceClassRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetTargetDBInstanceClass(v string) *ModifyDBInstanceClassRequest {
	s.TargetDBInstanceClass = &v
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
	ConfigName     *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	ConfigValue    *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDBInstanceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConfigRequest) SetConfigName(v string) *ModifyDBInstanceConfigRequest {
	s.ConfigName = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetConfigValue(v string) *ModifyDBInstanceConfigRequest {
	s.ConfigValue = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetDBInstanceName(v string) *ModifyDBInstanceConfigRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetRegionId(v string) *ModifyDBInstanceConfigRequest {
	s.RegionId = &v
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
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	DBInstanceName        *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDBInstanceDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDescriptionRequest) SetDBInstanceDescription(v string) *ModifyDBInstanceDescriptionRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetDBInstanceName(v string) *ModifyDBInstanceDescriptionRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetRegionId(v string) *ModifyDBInstanceDescriptionRequest {
	s.RegionId = &v
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
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	MaintainTime   *string `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDBInstanceMaintainTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceMaintainTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetClientToken(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetDBInstanceName(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetMaintainTime(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.MaintainTime = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetRegionId(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.RegionId = &v
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

type ModifyDatabaseDescriptionRequest struct {
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DbDescription  *string `json:"DbDescription,omitempty" xml:"DbDescription,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDatabaseDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseDescriptionRequest) SetDBInstanceName(v string) *ModifyDatabaseDescriptionRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyDatabaseDescriptionRequest) SetDbDescription(v string) *ModifyDatabaseDescriptionRequest {
	s.DbDescription = &v
	return s
}

func (s *ModifyDatabaseDescriptionRequest) SetDbName(v string) *ModifyDatabaseDescriptionRequest {
	s.DbName = &v
	return s
}

func (s *ModifyDatabaseDescriptionRequest) SetRegionId(v string) *ModifyDatabaseDescriptionRequest {
	s.RegionId = &v
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

type ModifyParameterRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ParamLevel   *string `json:"ParamLevel,omitempty" xml:"ParamLevel,omitempty"`
	Parameters   *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyParameterRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyParameterRequest) GoString() string {
	return s.String()
}

func (s *ModifyParameterRequest) SetClientToken(v string) *ModifyParameterRequest {
	s.ClientToken = &v
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

func (s *ModifyParameterRequest) SetRegionId(v string) *ModifyParameterRequest {
	s.RegionId = &v
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
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	GroupName      *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ModifyMode     *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s ModifySecurityIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsRequest) SetDBInstanceName(v string) *ModifySecurityIpsRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetGroupName(v string) *ModifySecurityIpsRequest {
	s.GroupName = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetModifyMode(v string) *ModifySecurityIpsRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetRegionId(v string) *ModifySecurityIpsRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetSecurityIPList(v string) *ModifySecurityIpsRequest {
	s.SecurityIPList = &v
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
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
	DBInstanceName          *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	OwnerAccount            *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                 *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount    *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId         *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ReleaseInstancePublicConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstancePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstancePublicConnectionRequest) SetCurrentConnectionString(v string) *ReleaseInstancePublicConnectionRequest {
	s.CurrentConnectionString = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetDBInstanceName(v string) *ReleaseInstancePublicConnectionRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetOwnerAccount(v string) *ReleaseInstancePublicConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetOwnerId(v int64) *ReleaseInstancePublicConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetRegionId(v string) *ReleaseInstancePublicConnectionRequest {
	s.RegionId = &v
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

type ResetAccountPasswordRequest struct {
	AccountName             *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPassword         *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	DBInstanceName          *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityAccountName     *string `json:"SecurityAccountName,omitempty" xml:"SecurityAccountName,omitempty"`
	SecurityAccountPassword *string `json:"SecurityAccountPassword,omitempty" xml:"SecurityAccountPassword,omitempty"`
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

func (s *ResetAccountPasswordRequest) SetDBInstanceName(v string) *ResetAccountPasswordRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetRegionId(v string) *ResetAccountPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetSecurityAccountName(v string) *ResetAccountPasswordRequest {
	s.SecurityAccountName = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetSecurityAccountPassword(v string) *ResetAccountPasswordRequest {
	s.SecurityAccountPassword = &v
	return s
}

type ResetAccountPasswordResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResetAccountPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordResponseBody) SetMessage(v string) *ResetAccountPasswordResponseBody {
	s.Message = &v
	return s
}

func (s *ResetAccountPasswordResponseBody) SetRequestId(v string) *ResetAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetAccountPasswordResponseBody) SetSuccess(v bool) *ResetAccountPasswordResponseBody {
	s.Success = &v
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
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RestartDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartDBInstanceRequest) SetDBInstanceName(v string) *RestartDBInstanceRequest {
	s.DBInstanceName = &v
	return s
}

func (s *RestartDBInstanceRequest) SetRegionId(v string) *RestartDBInstanceRequest {
	s.RegionId = &v
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

type UpdateBackupPolicyRequest struct {
	BackupPeriod               *string `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	BackupPlanBegin            *string `json:"BackupPlanBegin,omitempty" xml:"BackupPlanBegin,omitempty"`
	BackupSetRetention         *int32  `json:"BackupSetRetention,omitempty" xml:"BackupSetRetention,omitempty"`
	BackupType                 *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	BackupWay                  *string `json:"BackupWay,omitempty" xml:"BackupWay,omitempty"`
	DBInstanceName             *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	ForceCleanOnHighSpaceUsage *int32  `json:"ForceCleanOnHighSpaceUsage,omitempty" xml:"ForceCleanOnHighSpaceUsage,omitempty"`
	IsEnabled                  *int32  `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	LocalLogRetention          *int32  `json:"LocalLogRetention,omitempty" xml:"LocalLogRetention,omitempty"`
	LogLocalRetentionSpace     *int32  `json:"LogLocalRetentionSpace,omitempty" xml:"LogLocalRetentionSpace,omitempty"`
	RegionId                   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RemoveLogRetention         *int32  `json:"RemoveLogRetention,omitempty" xml:"RemoveLogRetention,omitempty"`
}

func (s UpdateBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBackupPolicyRequest) GoString() string {
	return s.String()
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

func (s *UpdateBackupPolicyRequest) SetDBInstanceName(v string) *UpdateBackupPolicyRequest {
	s.DBInstanceName = &v
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

func (s *UpdateBackupPolicyRequest) SetLogLocalRetentionSpace(v int32) *UpdateBackupPolicyRequest {
	s.LogLocalRetentionSpace = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetRegionId(v string) *UpdateBackupPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetRemoveLogRetention(v int32) *UpdateBackupPolicyRequest {
	s.RemoveLogRetention = &v
	return s
}

type UpdateBackupPolicyResponseBody struct {
	Data      []*UpdateBackupPolicyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBackupPolicyResponseBody) SetData(v []*UpdateBackupPolicyResponseBodyData) *UpdateBackupPolicyResponseBody {
	s.Data = v
	return s
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

type UpdateBackupPolicyResponseBodyData struct {
	BackupPeriod               *string `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	BackupPlanBegin            *string `json:"BackupPlanBegin,omitempty" xml:"BackupPlanBegin,omitempty"`
	BackupSetRetention         *int32  `json:"BackupSetRetention,omitempty" xml:"BackupSetRetention,omitempty"`
	BackupType                 *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	BackupWay                  *string `json:"BackupWay,omitempty" xml:"BackupWay,omitempty"`
	DBInstanceName             *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	ForceCleanOnHighSpaceUsage *int32  `json:"ForceCleanOnHighSpaceUsage,omitempty" xml:"ForceCleanOnHighSpaceUsage,omitempty"`
	IsEnabled                  *int32  `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	LocalLogRetention          *int32  `json:"LocalLogRetention,omitempty" xml:"LocalLogRetention,omitempty"`
	LogLocalRetentionSpace     *int32  `json:"LogLocalRetentionSpace,omitempty" xml:"LogLocalRetentionSpace,omitempty"`
	RemoveLogRetention         *int32  `json:"RemoveLogRetention,omitempty" xml:"RemoveLogRetention,omitempty"`
}

func (s UpdateBackupPolicyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateBackupPolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateBackupPolicyResponseBodyData) SetBackupPeriod(v string) *UpdateBackupPolicyResponseBodyData {
	s.BackupPeriod = &v
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

func (s *UpdateBackupPolicyResponseBodyData) SetBackupType(v string) *UpdateBackupPolicyResponseBodyData {
	s.BackupType = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetBackupWay(v string) *UpdateBackupPolicyResponseBodyData {
	s.BackupWay = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetDBInstanceName(v string) *UpdateBackupPolicyResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetForceCleanOnHighSpaceUsage(v int32) *UpdateBackupPolicyResponseBodyData {
	s.ForceCleanOnHighSpaceUsage = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetIsEnabled(v int32) *UpdateBackupPolicyResponseBodyData {
	s.IsEnabled = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetLocalLogRetention(v int32) *UpdateBackupPolicyResponseBodyData {
	s.LocalLogRetention = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetLogLocalRetentionSpace(v int32) *UpdateBackupPolicyResponseBodyData {
	s.LogLocalRetentionSpace = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetRemoveLogRetention(v int32) *UpdateBackupPolicyResponseBodyData {
	s.RemoveLogRetention = &v
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
	CertCommonName *string `json:"CertCommonName,omitempty" xml:"CertCommonName,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	EnableSSL      *bool   `json:"EnableSSL,omitempty" xml:"EnableSSL,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateDBInstanceSSLRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDBInstanceSSLRequest) GoString() string {
	return s.String()
}

func (s *UpdateDBInstanceSSLRequest) SetCertCommonName(v string) *UpdateDBInstanceSSLRequest {
	s.CertCommonName = &v
	return s
}

func (s *UpdateDBInstanceSSLRequest) SetDBInstanceName(v string) *UpdateDBInstanceSSLRequest {
	s.DBInstanceName = &v
	return s
}

func (s *UpdateDBInstanceSSLRequest) SetEnableSSL(v bool) *UpdateDBInstanceSSLRequest {
	s.EnableSSL = &v
	return s
}

func (s *UpdateDBInstanceSSLRequest) SetRegionId(v string) *UpdateDBInstanceSSLRequest {
	s.RegionId = &v
	return s
}

type UpdateDBInstanceSSLResponseBody struct {
	Data      *UpdateDBInstanceSSLResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDBInstanceSSLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDBInstanceSSLResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDBInstanceSSLResponseBody) SetData(v *UpdateDBInstanceSSLResponseBodyData) *UpdateDBInstanceSSLResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDBInstanceSSLResponseBody) SetRequestId(v string) *UpdateDBInstanceSSLResponseBody {
	s.RequestId = &v
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
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	EncryptionKey  *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RoleArn        *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	TDEStatus      *int32  `json:"TDEStatus,omitempty" xml:"TDEStatus,omitempty"`
}

func (s UpdateDBInstanceTDERequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDBInstanceTDERequest) GoString() string {
	return s.String()
}

func (s *UpdateDBInstanceTDERequest) SetDBInstanceName(v string) *UpdateDBInstanceTDERequest {
	s.DBInstanceName = &v
	return s
}

func (s *UpdateDBInstanceTDERequest) SetEncryptionKey(v string) *UpdateDBInstanceTDERequest {
	s.EncryptionKey = &v
	return s
}

func (s *UpdateDBInstanceTDERequest) SetRegionId(v string) *UpdateDBInstanceTDERequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDBInstanceTDERequest) SetRoleArn(v string) *UpdateDBInstanceTDERequest {
	s.RoleArn = &v
	return s
}

func (s *UpdateDBInstanceTDERequest) SetTDEStatus(v int32) *UpdateDBInstanceTDERequest {
	s.TDEStatus = &v
	return s
}

type UpdateDBInstanceTDEResponseBody struct {
	Data      *UpdateDBInstanceTDEResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDBInstanceTDEResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDBInstanceTDEResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDBInstanceTDEResponseBody) SetData(v *UpdateDBInstanceTDEResponseBodyData) *UpdateDBInstanceTDEResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDBInstanceTDEResponseBody) SetRequestId(v string) *UpdateDBInstanceTDEResponseBody {
	s.RequestId = &v
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
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DBInstanceName      *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DbInstanceNodeCount *string `json:"DbInstanceNodeCount,omitempty" xml:"DbInstanceNodeCount,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdatePolarDBXInstanceNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePolarDBXInstanceNodeRequest) GoString() string {
	return s.String()
}

func (s *UpdatePolarDBXInstanceNodeRequest) SetClientToken(v string) *UpdatePolarDBXInstanceNodeRequest {
	s.ClientToken = &v
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

func (s *UpdatePolarDBXInstanceNodeRequest) SetRegionId(v string) *UpdatePolarDBXInstanceNodeRequest {
	s.RegionId = &v
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
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpgradeDBInstanceKernelVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBInstanceKernelVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetDBInstanceName(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.DBInstanceName = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetRegionId(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.RegionId = &v
	return s
}

type UpgradeDBInstanceKernelVersionResponseBody struct {
	DBInstanceName     *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TargetMinorVersion *string `json:"TargetMinorVersion,omitempty" xml:"TargetMinorVersion,omitempty"`
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

func (s *UpgradeDBInstanceKernelVersionResponseBody) SetRequestId(v string) *UpgradeDBInstanceKernelVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionResponseBody) SetTargetMinorVersion(v string) *UpgradeDBInstanceKernelVersionResponseBody {
	s.TargetMinorVersion = &v
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

func (client *Client) DescribeCharacterSetWithOptions(request *DescribeCharacterSetRequest, runtime *util.RuntimeOptions) (_result *DescribeCharacterSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCharacterSetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCharacterSet"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCharacterSet(request *DescribeCharacterSetRequest) (_result *DescribeCharacterSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCharacterSetResponse{}
	_body, _err := client.DescribeCharacterSetWithOptions(request, runtime)
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

func (client *Client) DescribeDBInstanceConfigWithOptions(request *DescribeDBInstanceConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBInstanceConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBInstanceConfig"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceConfig(request *DescribeDBInstanceConfigRequest) (_result *DescribeDBInstanceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceConfigResponse{}
	_body, _err := client.DescribeDBInstanceConfigWithOptions(request, runtime)
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

func (client *Client) DescribeDBInstanceTopologyWithOptions(request *DescribeDBInstanceTopologyRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceTopologyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBInstanceTopologyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBInstanceTopology"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceTopology(request *DescribeDBInstanceTopologyRequest) (_result *DescribeDBInstanceTopologyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceTopologyResponse{}
	_body, _err := client.DescribeDBInstanceTopologyWithOptions(request, runtime)
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

func (client *Client) DescribeDBNodePerformanceWithOptions(request *DescribeDBNodePerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDBNodePerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBNodePerformanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBNodePerformance"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBNodePerformance(request *DescribeDBNodePerformanceRequest) (_result *DescribeDBNodePerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBNodePerformanceResponse{}
	_body, _err := client.DescribeDBNodePerformanceWithOptions(request, runtime)
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

func (client *Client) DescribeParameterTemplatesWithOptions(request *DescribeParameterTemplatesRequest, runtime *util.RuntimeOptions) (_result *DescribeParameterTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeParameterTemplatesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeParameterTemplates"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribePolarxDataNodesWithOptions(request *DescribePolarxDataNodesRequest, runtime *util.RuntimeOptions) (_result *DescribePolarxDataNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePolarxDataNodesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePolarxDataNodes"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePolarxDataNodes(request *DescribePolarxDataNodesRequest) (_result *DescribePolarxDataNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePolarxDataNodesResponse{}
	_body, _err := client.DescribePolarxDataNodesWithOptions(request, runtime)
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

func (client *Client) DescribeTableDetailWithOptions(request *DescribeTableDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeTableDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeTableDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTableDetail"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DisableRightsSeparationWithOptions(request *DisableRightsSeparationRequest, runtime *util.RuntimeOptions) (_result *DisableRightsSeparationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableRightsSeparationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableRightsSeparation"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableRightsSeparation(request *DisableRightsSeparationRequest) (_result *DisableRightsSeparationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableRightsSeparationResponse{}
	_body, _err := client.DisableRightsSeparationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableRightsSeparationWithOptions(request *EnableRightsSeparationRequest, runtime *util.RuntimeOptions) (_result *EnableRightsSeparationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableRightsSeparationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableRightsSeparation"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableRightsSeparation(request *EnableRightsSeparationRequest) (_result *EnableRightsSeparationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableRightsSeparationResponse{}
	_body, _err := client.EnableRightsSeparationWithOptions(request, runtime)
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

func (client *Client) ModifyAccountPrivilegeWithOptions(request *ModifyAccountPrivilegeRequest, runtime *util.RuntimeOptions) (_result *ModifyAccountPrivilegeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyAccountPrivilegeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyAccountPrivilege"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAccountPrivilege(request *ModifyAccountPrivilegeRequest) (_result *ModifyAccountPrivilegeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAccountPrivilegeResponse{}
	_body, _err := client.ModifyAccountPrivilegeWithOptions(request, runtime)
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

func (client *Client) ModifyDBInstanceMaintainTimeWithOptions(request *ModifyDBInstanceMaintainTimeRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceMaintainTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBInstanceMaintainTimeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBInstanceMaintainTime"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ResetAccountPasswordWithOptions(request *ResetAccountPasswordRequest, runtime *util.RuntimeOptions) (_result *ResetAccountPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResetAccountPasswordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResetAccountPassword"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
