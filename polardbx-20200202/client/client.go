// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AlignStoragePrimaryAzoneRequest struct {
	// This parameter is required.
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StorageInstanceName *string `json:"StorageInstanceName,omitempty" xml:"StorageInstanceName,omitempty"`
	SwitchTime          *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	SwitchTimeMode      *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
}

func (s AlignStoragePrimaryAzoneRequest) String() string {
	return tea.Prettify(s)
}

func (s AlignStoragePrimaryAzoneRequest) GoString() string {
	return s.String()
}

func (s *AlignStoragePrimaryAzoneRequest) SetDBInstanceName(v string) *AlignStoragePrimaryAzoneRequest {
	s.DBInstanceName = &v
	return s
}

func (s *AlignStoragePrimaryAzoneRequest) SetRegionId(v string) *AlignStoragePrimaryAzoneRequest {
	s.RegionId = &v
	return s
}

func (s *AlignStoragePrimaryAzoneRequest) SetStorageInstanceName(v string) *AlignStoragePrimaryAzoneRequest {
	s.StorageInstanceName = &v
	return s
}

func (s *AlignStoragePrimaryAzoneRequest) SetSwitchTime(v string) *AlignStoragePrimaryAzoneRequest {
	s.SwitchTime = &v
	return s
}

func (s *AlignStoragePrimaryAzoneRequest) SetSwitchTimeMode(v string) *AlignStoragePrimaryAzoneRequest {
	s.SwitchTimeMode = &v
	return s
}

type AlignStoragePrimaryAzoneResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AlignStoragePrimaryAzoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AlignStoragePrimaryAzoneResponseBody) GoString() string {
	return s.String()
}

func (s *AlignStoragePrimaryAzoneResponseBody) SetMessage(v string) *AlignStoragePrimaryAzoneResponseBody {
	s.Message = &v
	return s
}

func (s *AlignStoragePrimaryAzoneResponseBody) SetRequestId(v string) *AlignStoragePrimaryAzoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *AlignStoragePrimaryAzoneResponseBody) SetSuccess(v bool) *AlignStoragePrimaryAzoneResponseBody {
	s.Success = &v
	return s
}

type AlignStoragePrimaryAzoneResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AlignStoragePrimaryAzoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AlignStoragePrimaryAzoneResponse) String() string {
	return tea.Prettify(s)
}

func (s AlignStoragePrimaryAzoneResponse) GoString() string {
	return s.String()
}

func (s *AlignStoragePrimaryAzoneResponse) SetHeaders(v map[string]*string) *AlignStoragePrimaryAzoneResponse {
	s.Headers = v
	return s
}

func (s *AlignStoragePrimaryAzoneResponse) SetStatusCode(v int32) *AlignStoragePrimaryAzoneResponse {
	s.StatusCode = &v
	return s
}

func (s *AlignStoragePrimaryAzoneResponse) SetBody(v *AlignStoragePrimaryAzoneResponseBody) *AlignStoragePrimaryAzoneResponse {
	s.Body = v
	return s
}

type AllocateColdDataVolumeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzravgpt8q****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AllocateColdDataVolumeRequest) String() string {
	return tea.Prettify(s)
}

func (s AllocateColdDataVolumeRequest) GoString() string {
	return s.String()
}

func (s *AllocateColdDataVolumeRequest) SetDBInstanceName(v string) *AllocateColdDataVolumeRequest {
	s.DBInstanceName = &v
	return s
}

func (s *AllocateColdDataVolumeRequest) SetRegionId(v string) *AllocateColdDataVolumeRequest {
	s.RegionId = &v
	return s
}

type AllocateColdDataVolumeResponseBody struct {
	// example:
	//
	// D6A4256F-7B83-5BD7-9AC0-72E1FAC05330
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateColdDataVolumeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AllocateColdDataVolumeResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateColdDataVolumeResponseBody) SetRequestId(v string) *AllocateColdDataVolumeResponseBody {
	s.RequestId = &v
	return s
}

type AllocateColdDataVolumeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocateColdDataVolumeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocateColdDataVolumeResponse) String() string {
	return tea.Prettify(s)
}

func (s AllocateColdDataVolumeResponse) GoString() string {
	return s.String()
}

func (s *AllocateColdDataVolumeResponse) SetHeaders(v map[string]*string) *AllocateColdDataVolumeResponse {
	s.Headers = v
	return s
}

func (s *AllocateColdDataVolumeResponse) SetStatusCode(v int32) *AllocateColdDataVolumeResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocateColdDataVolumeResponse) SetBody(v *AllocateColdDataVolumeResponseBody) *AllocateColdDataVolumeResponse {
	s.Body = v
	return s
}

type AllocateInstancePublicConnectionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-**************
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-**************
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
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
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocateInstancePublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *AllocateInstancePublicConnectionResponse) SetStatusCode(v int32) *AllocateInstancePublicConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocateInstancePublicConnectionResponse) SetBody(v *AllocateInstancePublicConnectionResponseBody) *AllocateInstancePublicConnectionResponse {
	s.Body = v
	return s
}

type CancelActiveOperationTasksRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 111
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CancelActiveOperationTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelActiveOperationTasksRequest) GoString() string {
	return s.String()
}

func (s *CancelActiveOperationTasksRequest) SetIds(v string) *CancelActiveOperationTasksRequest {
	s.Ids = &v
	return s
}

func (s *CancelActiveOperationTasksRequest) SetRegionId(v string) *CancelActiveOperationTasksRequest {
	s.RegionId = &v
	return s
}

type CancelActiveOperationTasksResponseBody struct {
	// example:
	//
	// 111,1223
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// example:
	//
	// AE4F6C34-065F-45AA-B5DC-4B8D816F6305
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelActiveOperationTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelActiveOperationTasksResponseBody) GoString() string {
	return s.String()
}

func (s *CancelActiveOperationTasksResponseBody) SetIds(v string) *CancelActiveOperationTasksResponseBody {
	s.Ids = &v
	return s
}

func (s *CancelActiveOperationTasksResponseBody) SetRequestId(v string) *CancelActiveOperationTasksResponseBody {
	s.RequestId = &v
	return s
}

type CancelActiveOperationTasksResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelActiveOperationTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelActiveOperationTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelActiveOperationTasksResponse) GoString() string {
	return s.String()
}

func (s *CancelActiveOperationTasksResponse) SetHeaders(v map[string]*string) *CancelActiveOperationTasksResponse {
	s.Headers = v
	return s
}

func (s *CancelActiveOperationTasksResponse) SetStatusCode(v int32) *CancelActiveOperationTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelActiveOperationTasksResponse) SetBody(v *CancelActiveOperationTasksResponseBody) *CancelActiveOperationTasksResponse {
	s.Body = v
	return s
}

type ChangeResourceGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rg-acfmwolx3j4****
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-szrwrbp693****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PolarDBXInstance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) SetNewResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetRegionId(v string) *ChangeResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceId(v string) *ChangeResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceType(v string) *ChangeResourceGroupRequest {
	s.ResourceType = &v
	return s
}

type ChangeResourceGroupResponseBody struct {
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3WE34
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) SetRequestId(v string) *ChangeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type ChangeResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponse) SetHeaders(v map[string]*string) *ChangeResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ChangeResourceGroupResponse) SetStatusCode(v int32) *ChangeResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeResourceGroupResponse) SetBody(v *ChangeResourceGroupResponseBody) *ChangeResourceGroupResponse {
	s.Body = v
	return s
}

type CheckCloudResourceAuthorizedRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// acs:ram::123456789012****:role/AliyunRdsInstanceEncryptionDefaultRole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
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
	Data *CheckCloudResourceAuthorizedResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// A501A191-BD70-5E50-98A9-C2A486A82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 0
	AuthorizationState *string `json:"AuthorizationState,omitempty" xml:"AuthorizationState,omitempty"`
	// example:
	//
	// acs:ram::123456789012****:role/AliyunRdsInstanceEncryptionDefaultRole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckCloudResourceAuthorizedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CheckCloudResourceAuthorizedResponse) SetStatusCode(v int32) *CheckCloudResourceAuthorizedResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckCloudResourceAuthorizedResponse) SetBody(v *CheckCloudResourceAuthorizedResponseBody) *CheckCloudResourceAuthorizedResponse {
	s.Body = v
	return s
}

type CreateAccountRequest struct {
	// example:
	//
	// test
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testAccount
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Test@1111
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// example:
	//
	// ReadWrite
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-**************
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// testdb
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// securityAccount
	SecurityAccountName *string `json:"SecurityAccountName,omitempty" xml:"SecurityAccountName,omitempty"`
	// example:
	//
	// securityPassword
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
	// example:
	//
	// ****
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateAccountResponse) SetStatusCode(v int32) *CreateAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccountResponse) SetBody(v *CreateAccountResponseBody) *CreateAccountResponse {
	s.Body = v
	return s
}

type CreateBackupRequest struct {
	// example:
	//
	// 0
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-**************
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Data *CreateBackupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// *****
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateBackupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBackupResponseBody) SetData(v *CreateBackupResponseBodyData) *CreateBackupResponseBody {
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
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
}

func (s CreateBackupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateBackupResponseBodyData) SetBackupSetId(v string) *CreateBackupResponseBodyData {
	s.BackupSetId = &v
	return s
}

type CreateBackupResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBackupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateBackupResponse) SetStatusCode(v int32) *CreateBackupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBackupResponse) SetBody(v *CreateBackupResponseBody) *CreateBackupResponse {
	s.Body = v
	return s
}

type CreateDBRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// testaccount
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// ReadWrite
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// utf8mb4
	Charset *string `json:"Charset,omitempty" xml:"Charset,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-**************
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// db for test
	DbDescription *string `json:"DbDescription,omitempty" xml:"DbDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testdb
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// example:
	//
	// auto
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// securityAccount
	SecurityAccountName *string `json:"SecurityAccountName,omitempty" xml:"SecurityAccountName,omitempty"`
	// example:
	//
	// securityPassword
	SecurityAccountPassword *string `json:"SecurityAccountPassword,omitempty" xml:"SecurityAccountPassword,omitempty"`
	StoragePoolName         *string `json:"StoragePoolName,omitempty" xml:"StoragePoolName,omitempty"`
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

func (s *CreateDBRequest) SetMode(v string) *CreateDBRequest {
	s.Mode = &v
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

func (s *CreateDBRequest) SetStoragePoolName(v string) *CreateDBRequest {
	s.StoragePoolName = &v
	return s
}

type CreateDBResponseBody struct {
	// example:
	//
	// *****
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateDBResponse) SetStatusCode(v int32) *CreateDBResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBResponse) SetBody(v *CreateDBResponseBody) *CreateDBResponse {
	s.Body = v
	return s
}

type CreateDBInstanceRequest struct {
	// example:
	//
	// true
	AutoRenew   *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	CNNodeCount *string `json:"CNNodeCount,omitempty" xml:"CNNodeCount,omitempty"`
	// example:
	//
	// xxxxxx-xxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CnClass     *string `json:"CnClass,omitempty" xml:"CnClass,omitempty"`
	// example:
	//
	// polarx.x4.2xlarge.2d
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// example:
	//
	// 2
	DBNodeCount    *int32  `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	DNNodeCount    *string `json:"DNNodeCount,omitempty" xml:"DNNodeCount,omitempty"`
	DnClass        *string `json:"DnClass,omitempty" xml:"DnClass,omitempty"`
	DnStorageSpace *string `json:"DnStorageSpace,omitempty" xml:"DnStorageSpace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2.0
	EngineVersion            *string            `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ExtraParams              map[string]*string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	IsColumnarReadDBInstance *bool              `json:"IsColumnarReadDBInstance,omitempty" xml:"IsColumnarReadDBInstance,omitempty"`
	// example:
	//
	// false
	IsReadDBInstance *bool `json:"IsReadDBInstance,omitempty" xml:"IsReadDBInstance,omitempty"`
	// example:
	//
	// vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PREPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// pxc-*********
	PrimaryDBInstanceName *string `json:"PrimaryDBInstanceName,omitempty" xml:"PrimaryDBInstanceName,omitempty"`
	// example:
	//
	// cn-shenzhen-e
	PrimaryZone *string `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// null
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// cn-shenzhen-a
	SecondaryZone *string `json:"SecondaryZone,omitempty" xml:"SecondaryZone,omitempty"`
	Series        *string `json:"Series,omitempty" xml:"Series,omitempty"`
	// example:
	//
	// cn-shenzhen-e
	TertiaryZone *string `json:"TertiaryZone,omitempty" xml:"TertiaryZone,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3azones
	TopologyType *string `json:"TopologyType,omitempty" xml:"TopologyType,omitempty"`
	// example:
	//
	// 1
	UsedTime *int32 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-*****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// example:
	//
	// vsw-*********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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

func (s *CreateDBInstanceRequest) SetCNNodeCount(v string) *CreateDBInstanceRequest {
	s.CNNodeCount = &v
	return s
}

func (s *CreateDBInstanceRequest) SetClientToken(v string) *CreateDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceRequest) SetCnClass(v string) *CreateDBInstanceRequest {
	s.CnClass = &v
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

func (s *CreateDBInstanceRequest) SetDNNodeCount(v string) *CreateDBInstanceRequest {
	s.DNNodeCount = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDnClass(v string) *CreateDBInstanceRequest {
	s.DnClass = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDnStorageSpace(v string) *CreateDBInstanceRequest {
	s.DnStorageSpace = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEngineVersion(v string) *CreateDBInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateDBInstanceRequest) SetExtraParams(v map[string]*string) *CreateDBInstanceRequest {
	s.ExtraParams = v
	return s
}

func (s *CreateDBInstanceRequest) SetIsColumnarReadDBInstance(v bool) *CreateDBInstanceRequest {
	s.IsColumnarReadDBInstance = &v
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

func (s *CreateDBInstanceRequest) SetPrimaryZone(v string) *CreateDBInstanceRequest {
	s.PrimaryZone = &v
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

func (s *CreateDBInstanceRequest) SetSecondaryZone(v string) *CreateDBInstanceRequest {
	s.SecondaryZone = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSeries(v string) *CreateDBInstanceRequest {
	s.Series = &v
	return s
}

func (s *CreateDBInstanceRequest) SetTertiaryZone(v string) *CreateDBInstanceRequest {
	s.TertiaryZone = &v
	return s
}

func (s *CreateDBInstanceRequest) SetTopologyType(v string) *CreateDBInstanceRequest {
	s.TopologyType = &v
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

type CreateDBInstanceShrinkRequest struct {
	// example:
	//
	// true
	AutoRenew   *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	CNNodeCount *string `json:"CNNodeCount,omitempty" xml:"CNNodeCount,omitempty"`
	// example:
	//
	// xxxxxx-xxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CnClass     *string `json:"CnClass,omitempty" xml:"CnClass,omitempty"`
	// example:
	//
	// polarx.x4.2xlarge.2d
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// example:
	//
	// 2
	DBNodeCount    *int32  `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	DNNodeCount    *string `json:"DNNodeCount,omitempty" xml:"DNNodeCount,omitempty"`
	DnClass        *string `json:"DnClass,omitempty" xml:"DnClass,omitempty"`
	DnStorageSpace *string `json:"DnStorageSpace,omitempty" xml:"DnStorageSpace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2.0
	EngineVersion            *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ExtraParamsShrink        *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	IsColumnarReadDBInstance *bool   `json:"IsColumnarReadDBInstance,omitempty" xml:"IsColumnarReadDBInstance,omitempty"`
	// example:
	//
	// false
	IsReadDBInstance *bool `json:"IsReadDBInstance,omitempty" xml:"IsReadDBInstance,omitempty"`
	// example:
	//
	// vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PREPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// pxc-*********
	PrimaryDBInstanceName *string `json:"PrimaryDBInstanceName,omitempty" xml:"PrimaryDBInstanceName,omitempty"`
	// example:
	//
	// cn-shenzhen-e
	PrimaryZone *string `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// null
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// cn-shenzhen-a
	SecondaryZone *string `json:"SecondaryZone,omitempty" xml:"SecondaryZone,omitempty"`
	Series        *string `json:"Series,omitempty" xml:"Series,omitempty"`
	// example:
	//
	// cn-shenzhen-e
	TertiaryZone *string `json:"TertiaryZone,omitempty" xml:"TertiaryZone,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3azones
	TopologyType *string `json:"TopologyType,omitempty" xml:"TopologyType,omitempty"`
	// example:
	//
	// 1
	UsedTime *int32 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-*****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// example:
	//
	// vsw-*********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDBInstanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceShrinkRequest) SetAutoRenew(v bool) *CreateDBInstanceShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetCNNodeCount(v string) *CreateDBInstanceShrinkRequest {
	s.CNNodeCount = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetClientToken(v string) *CreateDBInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetCnClass(v string) *CreateDBInstanceShrinkRequest {
	s.CnClass = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDBNodeClass(v string) *CreateDBInstanceShrinkRequest {
	s.DBNodeClass = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDBNodeCount(v int32) *CreateDBInstanceShrinkRequest {
	s.DBNodeCount = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDNNodeCount(v string) *CreateDBInstanceShrinkRequest {
	s.DNNodeCount = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDnClass(v string) *CreateDBInstanceShrinkRequest {
	s.DnClass = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDnStorageSpace(v string) *CreateDBInstanceShrinkRequest {
	s.DnStorageSpace = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetEngineVersion(v string) *CreateDBInstanceShrinkRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetExtraParamsShrink(v string) *CreateDBInstanceShrinkRequest {
	s.ExtraParamsShrink = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetIsColumnarReadDBInstance(v bool) *CreateDBInstanceShrinkRequest {
	s.IsColumnarReadDBInstance = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetIsReadDBInstance(v bool) *CreateDBInstanceShrinkRequest {
	s.IsReadDBInstance = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetNetworkType(v string) *CreateDBInstanceShrinkRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetPayType(v string) *CreateDBInstanceShrinkRequest {
	s.PayType = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetPeriod(v string) *CreateDBInstanceShrinkRequest {
	s.Period = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetPrimaryDBInstanceName(v string) *CreateDBInstanceShrinkRequest {
	s.PrimaryDBInstanceName = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetPrimaryZone(v string) *CreateDBInstanceShrinkRequest {
	s.PrimaryZone = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetRegionId(v string) *CreateDBInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetResourceGroupId(v string) *CreateDBInstanceShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetSecondaryZone(v string) *CreateDBInstanceShrinkRequest {
	s.SecondaryZone = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetSeries(v string) *CreateDBInstanceShrinkRequest {
	s.Series = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetTertiaryZone(v string) *CreateDBInstanceShrinkRequest {
	s.TertiaryZone = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetTopologyType(v string) *CreateDBInstanceShrinkRequest {
	s.TopologyType = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetUsedTime(v int32) *CreateDBInstanceShrinkRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetVPCId(v string) *CreateDBInstanceShrinkRequest {
	s.VPCId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetVSwitchId(v string) *CreateDBInstanceShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetZoneId(v string) *CreateDBInstanceShrinkRequest {
	s.ZoneId = &v
	return s
}

type CreateDBInstanceResponseBody struct {
	// example:
	//
	// pxc-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// 12345
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateDBInstanceResponse) SetStatusCode(v int32) *CreateDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBInstanceResponse) SetBody(v *CreateDBInstanceResponseBody) *CreateDBInstanceResponse {
	s.Body = v
	return s
}

type CreateSuperAccountRequest struct {
	// example:
	//
	// testdbadescription
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dba
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testdbapassword
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-************
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// example:
	//
	// *****
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSuperAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateSuperAccountResponse) SetStatusCode(v int32) *CreateSuperAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSuperAccountResponse) SetBody(v *CreateSuperAccountResponseBody) *CreateSuperAccountResponse {
	s.Body = v
	return s
}

type DeleteAccountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// testaccount
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// securityAccount
	SecurityAccountName *string `json:"SecurityAccountName,omitempty" xml:"SecurityAccountName,omitempty"`
	// example:
	//
	// securityPassword
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
	// example:
	//
	// ****
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteAccountResponse) SetStatusCode(v int32) *DeleteAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccountResponse) SetBody(v *DeleteAccountResponseBody) *DeleteAccountResponse {
	s.Body = v
	return s
}

type DeleteDBRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testdb
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// example:
	//
	// *****
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteDBResponse) SetStatusCode(v int32) *DeleteDBResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBResponse) SetBody(v *DeleteDBResponseBody) *DeleteDBResponse {
	s.Body = v
	return s
}

type DeleteDBInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteDBInstanceResponse) SetStatusCode(v int32) *DeleteDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBInstanceResponse) SetBody(v *DeleteDBInstanceResponseBody) *DeleteDBInstanceResponse {
	s.Body = v
	return s
}

type DescribeAccountListRequest struct {
	// example:
	//
	// testaccount
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 0
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Data []*DescribeAccountListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// testaccount desc
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// example:
	//
	// testAccount
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// ReadWrite
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	// example:
	//
	// 0
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// example:
	//
	// pxc-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// testdb
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// example:
	//
	// 2012-06-08T15:00Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccountListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeAccountListResponse) SetStatusCode(v int32) *DescribeAccountListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccountListResponse) SetBody(v *DescribeAccountListResponseBody) *DescribeAccountListResponse {
	s.Body = v
	return s
}

type DescribeActiveOperationMaintainConfRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeActiveOperationMaintainConfRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveOperationMaintainConfRequest) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationMaintainConfRequest) SetRegionId(v string) *DescribeActiveOperationMaintainConfRequest {
	s.RegionId = &v
	return s
}

type DescribeActiveOperationMaintainConfResponseBody struct {
	Config *DescribeActiveOperationMaintainConfResponseBodyConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// example:
	//
	// 1
	HasConfig *int64 `json:"HasConfig,omitempty" xml:"HasConfig,omitempty"`
	// example:
	//
	// 1A586DCB-39A6-4050-81CC-C7BD4CCDB49F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeActiveOperationMaintainConfResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveOperationMaintainConfResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationMaintainConfResponseBody) SetConfig(v *DescribeActiveOperationMaintainConfResponseBodyConfig) *DescribeActiveOperationMaintainConfResponseBody {
	s.Config = v
	return s
}

func (s *DescribeActiveOperationMaintainConfResponseBody) SetHasConfig(v int64) *DescribeActiveOperationMaintainConfResponseBody {
	s.HasConfig = &v
	return s
}

func (s *DescribeActiveOperationMaintainConfResponseBody) SetRequestId(v string) *DescribeActiveOperationMaintainConfResponseBody {
	s.RequestId = &v
	return s
}

type DescribeActiveOperationMaintainConfResponseBodyConfig struct {
	// example:
	//
	// 2021-08-11T10:08:27Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// 1,2,3,4,5,6,7
	CycleTime *string `json:"CycleTime,omitempty" xml:"CycleTime,omitempty"`
	// example:
	//
	// Week
	CycleType *string `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	// example:
	//
	// 04:00:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// example:
	//
	// 04:00:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// example:
	//
	// 2021-08-11T10:08:27Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeActiveOperationMaintainConfResponseBodyConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveOperationMaintainConfResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationMaintainConfResponseBodyConfig) SetCreatedTime(v string) *DescribeActiveOperationMaintainConfResponseBodyConfig {
	s.CreatedTime = &v
	return s
}

func (s *DescribeActiveOperationMaintainConfResponseBodyConfig) SetCycleTime(v string) *DescribeActiveOperationMaintainConfResponseBodyConfig {
	s.CycleTime = &v
	return s
}

func (s *DescribeActiveOperationMaintainConfResponseBodyConfig) SetCycleType(v string) *DescribeActiveOperationMaintainConfResponseBodyConfig {
	s.CycleType = &v
	return s
}

func (s *DescribeActiveOperationMaintainConfResponseBodyConfig) SetMaintainEndTime(v string) *DescribeActiveOperationMaintainConfResponseBodyConfig {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeActiveOperationMaintainConfResponseBodyConfig) SetMaintainStartTime(v string) *DescribeActiveOperationMaintainConfResponseBodyConfig {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeActiveOperationMaintainConfResponseBodyConfig) SetModifiedTime(v string) *DescribeActiveOperationMaintainConfResponseBodyConfig {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeActiveOperationMaintainConfResponseBodyConfig) SetStatus(v int64) *DescribeActiveOperationMaintainConfResponseBodyConfig {
	s.Status = &v
	return s
}

type DescribeActiveOperationMaintainConfResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeActiveOperationMaintainConfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeActiveOperationMaintainConfResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveOperationMaintainConfResponse) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationMaintainConfResponse) SetHeaders(v map[string]*string) *DescribeActiveOperationMaintainConfResponse {
	s.Headers = v
	return s
}

func (s *DescribeActiveOperationMaintainConfResponse) SetStatusCode(v int32) *DescribeActiveOperationMaintainConfResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeActiveOperationMaintainConfResponse) SetBody(v *DescribeActiveOperationMaintainConfResponseBody) *DescribeActiveOperationMaintainConfResponse {
	s.Body = v
	return s
}

type DescribeActiveOperationTaskCountRequest struct {
	// example:
	//
	// Category
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// polarx
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeActiveOperationTaskCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveOperationTaskCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskCountRequest) SetCategory(v string) *DescribeActiveOperationTaskCountRequest {
	s.Category = &v
	return s
}

func (s *DescribeActiveOperationTaskCountRequest) SetProduct(v string) *DescribeActiveOperationTaskCountRequest {
	s.Product = &v
	return s
}

func (s *DescribeActiveOperationTaskCountRequest) SetRegionId(v string) *DescribeActiveOperationTaskCountRequest {
	s.RegionId = &v
	return s
}

type DescribeActiveOperationTaskCountResponseBody struct {
	// example:
	//
	// 1
	NeedPop *int64 `json:"NeedPop,omitempty" xml:"NeedPop,omitempty"`
	// example:
	//
	// EC7E27FC-58F8-4722-89BB-D1B6D0971956
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	TaskCount *int64 `json:"TaskCount,omitempty" xml:"TaskCount,omitempty"`
}

func (s DescribeActiveOperationTaskCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveOperationTaskCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskCountResponseBody) SetNeedPop(v int64) *DescribeActiveOperationTaskCountResponseBody {
	s.NeedPop = &v
	return s
}

func (s *DescribeActiveOperationTaskCountResponseBody) SetRequestId(v string) *DescribeActiveOperationTaskCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeActiveOperationTaskCountResponseBody) SetTaskCount(v int64) *DescribeActiveOperationTaskCountResponseBody {
	s.TaskCount = &v
	return s
}

type DescribeActiveOperationTaskCountResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeActiveOperationTaskCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeActiveOperationTaskCountResponse) SetStatusCode(v int32) *DescribeActiveOperationTaskCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeActiveOperationTaskCountResponse) SetBody(v *DescribeActiveOperationTaskCountResponseBody) *DescribeActiveOperationTaskCountResponse {
	s.Body = v
	return s
}

type DescribeActiveOperationTasksRequest struct {
	// example:
	//
	// -1
	AllowCancel *int64 `json:"AllowCancel,omitempty" xml:"AllowCancel,omitempty"`
	// example:
	//
	// -1
	AllowChange *int64 `json:"AllowChange,omitempty" xml:"AllowChange,omitempty"`
	// example:
	//
	// all
	ChangeLevel *string `json:"ChangeLevel,omitempty" xml:"ChangeLevel,omitempty"`
	// example:
	//
	// polarx
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// example:
	//
	// pxc-xxxxx
	InsName *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 25
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// polarx
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// -1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// all
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeActiveOperationTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveOperationTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTasksRequest) SetAllowCancel(v int64) *DescribeActiveOperationTasksRequest {
	s.AllowCancel = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetAllowChange(v int64) *DescribeActiveOperationTasksRequest {
	s.AllowChange = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetChangeLevel(v string) *DescribeActiveOperationTasksRequest {
	s.ChangeLevel = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetDbType(v string) *DescribeActiveOperationTasksRequest {
	s.DbType = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetInsName(v string) *DescribeActiveOperationTasksRequest {
	s.InsName = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetPageNumber(v int64) *DescribeActiveOperationTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetPageSize(v int64) *DescribeActiveOperationTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetProductId(v string) *DescribeActiveOperationTasksRequest {
	s.ProductId = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetRegion(v string) *DescribeActiveOperationTasksRequest {
	s.Region = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetRegionId(v string) *DescribeActiveOperationTasksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetStatus(v int64) *DescribeActiveOperationTasksRequest {
	s.Status = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetTaskType(v string) *DescribeActiveOperationTasksRequest {
	s.TaskType = &v
	return s
}

type DescribeActiveOperationTasksResponseBody struct {
	Items []*DescribeActiveOperationTasksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 12
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 23
	TotalRecordCount *int64 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeActiveOperationTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveOperationTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTasksResponseBody) SetItems(v []*DescribeActiveOperationTasksResponseBodyItems) *DescribeActiveOperationTasksResponseBody {
	s.Items = v
	return s
}

func (s *DescribeActiveOperationTasksResponseBody) SetPageNumber(v int64) *DescribeActiveOperationTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBody) SetPageSize(v int64) *DescribeActiveOperationTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBody) SetRequestId(v string) *DescribeActiveOperationTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBody) SetTotalRecordCount(v int64) *DescribeActiveOperationTasksResponseBody {
	s.TotalRecordCount = &v
	return s
}

type DescribeActiveOperationTasksResponseBodyItems struct {
	// example:
	//
	// 0
	AllowCancel *string `json:"AllowCancel,omitempty" xml:"AllowCancel,omitempty"`
	// example:
	//
	// 0
	AllowChange *string `json:"AllowChange,omitempty" xml:"AllowChange,omitempty"`
	// example:
	//
	// Risk
	ChangeLevel *string `json:"ChangeLevel,omitempty" xml:"ChangeLevel,omitempty"`
	// example:
	//
	// Risk repairment
	ChangeLevelEn *string `json:"ChangeLevelEn,omitempty" xml:"ChangeLevelEn,omitempty"`
	ChangeLevelZh *string `json:"ChangeLevelZh,omitempty" xml:"ChangeLevelZh,omitempty"`
	// example:
	//
	// 2021-06-15T16:00:00Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// cn-shanghai-et-a
	CurrentAVZ *string `json:"CurrentAVZ,omitempty" xml:"CurrentAVZ,omitempty"`
	// example:
	//
	// polarx
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// example:
	//
	// 2.0
	DbVersion *string `json:"DbVersion,omitempty" xml:"DbVersion,omitempty"`
	// example:
	//
	// 2021-06-15T16:00:00Z
	Deadline *string `json:"Deadline,omitempty" xml:"Deadline,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// --
	Impact *string `json:"Impact,omitempty" xml:"Impact,omitempty"`
	// example:
	//
	// Transient instance disconnection
	ImpactEn *string `json:"ImpactEn,omitempty" xml:"ImpactEn,omitempty"`
	ImpactZh *string `json:"ImpactZh,omitempty" xml:"ImpactZh,omitempty"`
	// example:
	//
	// xxx
	InsComment *string `json:"InsComment,omitempty" xml:"InsComment,omitempty"`
	// example:
	//
	// pxc-dd
	InsName *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	// example:
	//
	// 2021-06-15T16:00:00Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// 03:00:00
	PrepareInterval *string `json:"PrepareInterval,omitempty" xml:"PrepareInterval,omitempty"`
	// example:
	//
	// cn-shanghai-et15-b01
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// --
	ResultInfo *string `json:"ResultInfo,omitempty" xml:"ResultInfo,omitempty"`
	// example:
	//
	// 2021-06-15T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 3
	Status      *int64    `json:"Status,omitempty" xml:"Status,omitempty"`
	SubInsNames []*string `json:"SubInsNames,omitempty" xml:"SubInsNames,omitempty" type:"Repeated"`
	// example:
	//
	// 2021-06-15T16:00:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// example:
	//
	// rds_apsaradb_transfer
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// example:
	//
	// --
	TaskTypeEn *string `json:"TaskTypeEn,omitempty" xml:"TaskTypeEn,omitempty"`
	// example:
	//
	// --
	TaskTypeZh *string `json:"TaskTypeZh,omitempty" xml:"TaskTypeZh,omitempty"`
}

func (s DescribeActiveOperationTasksResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveOperationTasksResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetAllowCancel(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.AllowCancel = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetAllowChange(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.AllowChange = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetChangeLevel(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.ChangeLevel = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetChangeLevelEn(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.ChangeLevelEn = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetChangeLevelZh(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.ChangeLevelZh = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetCreatedTime(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.CreatedTime = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetCurrentAVZ(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.CurrentAVZ = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetDbType(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.DbType = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetDbVersion(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.DbVersion = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetDeadline(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.Deadline = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetId(v int64) *DescribeActiveOperationTasksResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetImpact(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.Impact = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetImpactEn(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.ImpactEn = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetImpactZh(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.ImpactZh = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetInsComment(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.InsComment = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetInsName(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.InsName = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetModifiedTime(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetPrepareInterval(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.PrepareInterval = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetRegion(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.Region = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetResultInfo(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.ResultInfo = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetStartTime(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetStatus(v int64) *DescribeActiveOperationTasksResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetSubInsNames(v []*string) *DescribeActiveOperationTasksResponseBodyItems {
	s.SubInsNames = v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetSwitchTime(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.SwitchTime = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetTaskType(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.TaskType = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetTaskTypeEn(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.TaskTypeEn = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetTaskTypeZh(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.TaskTypeZh = &v
	return s
}

type DescribeActiveOperationTasksResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeActiveOperationTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeActiveOperationTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveOperationTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTasksResponse) SetHeaders(v map[string]*string) *DescribeActiveOperationTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeActiveOperationTasksResponse) SetStatusCode(v int32) *DescribeActiveOperationTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeActiveOperationTasksResponse) SetBody(v *DescribeActiveOperationTasksResponseBody) *DescribeActiveOperationTasksResponse {
	s.Body = v
	return s
}

type DescribeArchiveTableListRequest struct {
	// This parameter is required.
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	PageIndex *int64 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// This parameter is required.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TableName  *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeArchiveTableListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeArchiveTableListRequest) GoString() string {
	return s.String()
}

func (s *DescribeArchiveTableListRequest) SetDBInstanceName(v string) *DescribeArchiveTableListRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeArchiveTableListRequest) SetPageIndex(v int64) *DescribeArchiveTableListRequest {
	s.PageIndex = &v
	return s
}

func (s *DescribeArchiveTableListRequest) SetPageSize(v int64) *DescribeArchiveTableListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeArchiveTableListRequest) SetRegionId(v string) *DescribeArchiveTableListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeArchiveTableListRequest) SetSchemaName(v string) *DescribeArchiveTableListRequest {
	s.SchemaName = &v
	return s
}

func (s *DescribeArchiveTableListRequest) SetStatus(v string) *DescribeArchiveTableListRequest {
	s.Status = &v
	return s
}

func (s *DescribeArchiveTableListRequest) SetTableName(v string) *DescribeArchiveTableListRequest {
	s.TableName = &v
	return s
}

type DescribeArchiveTableListResponseBody struct {
	Data      *DescribeArchiveTableListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeArchiveTableListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeArchiveTableListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeArchiveTableListResponseBody) SetData(v *DescribeArchiveTableListResponseBodyData) *DescribeArchiveTableListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeArchiveTableListResponseBody) SetRequestId(v string) *DescribeArchiveTableListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeArchiveTableListResponseBodyData struct {
	PageIndex         *int64                                            `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize          *int64                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PausedCount       *int32                                            `json:"PausedCount,omitempty" xml:"PausedCount,omitempty"`
	RunningCount      *int32                                            `json:"RunningCount,omitempty" xml:"RunningCount,omitempty"`
	SuccessCount      *int32                                            `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	Tables            []*DescribeArchiveTableListResponseBodyDataTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	TobeArchivedConut *int32                                            `json:"TobeArchivedConut,omitempty" xml:"TobeArchivedConut,omitempty"`
	Total             *int64                                            `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeArchiveTableListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeArchiveTableListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeArchiveTableListResponseBodyData) SetPageIndex(v int64) *DescribeArchiveTableListResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *DescribeArchiveTableListResponseBodyData) SetPageSize(v int64) *DescribeArchiveTableListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeArchiveTableListResponseBodyData) SetPausedCount(v int32) *DescribeArchiveTableListResponseBodyData {
	s.PausedCount = &v
	return s
}

func (s *DescribeArchiveTableListResponseBodyData) SetRunningCount(v int32) *DescribeArchiveTableListResponseBodyData {
	s.RunningCount = &v
	return s
}

func (s *DescribeArchiveTableListResponseBodyData) SetSuccessCount(v int32) *DescribeArchiveTableListResponseBodyData {
	s.SuccessCount = &v
	return s
}

func (s *DescribeArchiveTableListResponseBodyData) SetTables(v []*DescribeArchiveTableListResponseBodyDataTables) *DescribeArchiveTableListResponseBodyData {
	s.Tables = v
	return s
}

func (s *DescribeArchiveTableListResponseBodyData) SetTobeArchivedConut(v int32) *DescribeArchiveTableListResponseBodyData {
	s.TobeArchivedConut = &v
	return s
}

func (s *DescribeArchiveTableListResponseBodyData) SetTotal(v int64) *DescribeArchiveTableListResponseBodyData {
	s.Total = &v
	return s
}

type DescribeArchiveTableListResponseBodyDataTables struct {
	ArchiveStatus          *string  `json:"ArchiveStatus,omitempty" xml:"ArchiveStatus,omitempty"`
	CreatedDate            *int64   `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	FileCount              *int32   `json:"FileCount,omitempty" xml:"FileCount,omitempty"`
	LastSuccessArchiveTime *int64   `json:"LastSuccessArchiveTime,omitempty" xml:"LastSuccessArchiveTime,omitempty"`
	SchemaName             *string  `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	SpaceSize              *float64 `json:"SpaceSize,omitempty" xml:"SpaceSize,omitempty"`
	TableName              *string  `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeArchiveTableListResponseBodyDataTables) String() string {
	return tea.Prettify(s)
}

func (s DescribeArchiveTableListResponseBodyDataTables) GoString() string {
	return s.String()
}

func (s *DescribeArchiveTableListResponseBodyDataTables) SetArchiveStatus(v string) *DescribeArchiveTableListResponseBodyDataTables {
	s.ArchiveStatus = &v
	return s
}

func (s *DescribeArchiveTableListResponseBodyDataTables) SetCreatedDate(v int64) *DescribeArchiveTableListResponseBodyDataTables {
	s.CreatedDate = &v
	return s
}

func (s *DescribeArchiveTableListResponseBodyDataTables) SetFileCount(v int32) *DescribeArchiveTableListResponseBodyDataTables {
	s.FileCount = &v
	return s
}

func (s *DescribeArchiveTableListResponseBodyDataTables) SetLastSuccessArchiveTime(v int64) *DescribeArchiveTableListResponseBodyDataTables {
	s.LastSuccessArchiveTime = &v
	return s
}

func (s *DescribeArchiveTableListResponseBodyDataTables) SetSchemaName(v string) *DescribeArchiveTableListResponseBodyDataTables {
	s.SchemaName = &v
	return s
}

func (s *DescribeArchiveTableListResponseBodyDataTables) SetSpaceSize(v float64) *DescribeArchiveTableListResponseBodyDataTables {
	s.SpaceSize = &v
	return s
}

func (s *DescribeArchiveTableListResponseBodyDataTables) SetTableName(v string) *DescribeArchiveTableListResponseBodyDataTables {
	s.TableName = &v
	return s
}

type DescribeArchiveTableListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeArchiveTableListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeArchiveTableListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeArchiveTableListResponse) GoString() string {
	return s.String()
}

func (s *DescribeArchiveTableListResponse) SetHeaders(v map[string]*string) *DescribeArchiveTableListResponse {
	s.Headers = v
	return s
}

func (s *DescribeArchiveTableListResponse) SetStatusCode(v int32) *DescribeArchiveTableListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeArchiveTableListResponse) SetBody(v *DescribeArchiveTableListResponseBody) *DescribeArchiveTableListResponse {
	s.Body = v
	return s
}

type DescribeBackupPolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Data *DescribeBackupPolicyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// B87E2AB3-B7C9-4394-9160-7F639F732031
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBody) SetData(v *DescribeBackupPolicyResponseBodyData) *DescribeBackupPolicyResponseBody {
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
	BackupPeriod                   *string `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	BackupPlanBegin                *string `json:"BackupPlanBegin,omitempty" xml:"BackupPlanBegin,omitempty"`
	BackupSetRetention             *int32  `json:"BackupSetRetention,omitempty" xml:"BackupSetRetention,omitempty"`
	BackupType                     *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	BackupWay                      *string `json:"BackupWay,omitempty" xml:"BackupWay,omitempty"`
	ColdDataBackupInterval         *int32  `json:"ColdDataBackupInterval,omitempty" xml:"ColdDataBackupInterval,omitempty"`
	ColdDataBackupRetention        *int32  `json:"ColdDataBackupRetention,omitempty" xml:"ColdDataBackupRetention,omitempty"`
	CrossRegionDataBackupRetention *int32  `json:"CrossRegionDataBackupRetention,omitempty" xml:"CrossRegionDataBackupRetention,omitempty"`
	CrossRegionLogBackupRetention  *int32  `json:"CrossRegionLogBackupRetention,omitempty" xml:"CrossRegionLogBackupRetention,omitempty"`
	DBInstanceName                 *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DestCrossRegion                *string `json:"DestCrossRegion,omitempty" xml:"DestCrossRegion,omitempty"`
	ForceCleanOnHighSpaceUsage     *int32  `json:"ForceCleanOnHighSpaceUsage,omitempty" xml:"ForceCleanOnHighSpaceUsage,omitempty"`
	IsCrossRegionDataBackupEnabled *bool   `json:"IsCrossRegionDataBackupEnabled,omitempty" xml:"IsCrossRegionDataBackupEnabled,omitempty"`
	IsCrossRegionLogBackupEnabled  *bool   `json:"IsCrossRegionLogBackupEnabled,omitempty" xml:"IsCrossRegionLogBackupEnabled,omitempty"`
	IsEnabled                      *int32  `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	LocalLogRetention              *int32  `json:"LocalLogRetention,omitempty" xml:"LocalLogRetention,omitempty"`
	LocalLogRetentionNumber        *int32  `json:"LocalLogRetentionNumber,omitempty" xml:"LocalLogRetentionNumber,omitempty"`
	LogLocalRetentionSpace         *int32  `json:"LogLocalRetentionSpace,omitempty" xml:"LogLocalRetentionSpace,omitempty"`
	RemoveLogRetention             *int32  `json:"RemoveLogRetention,omitempty" xml:"RemoveLogRetention,omitempty"`
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

func (s *DescribeBackupPolicyResponseBodyData) SetColdDataBackupInterval(v int32) *DescribeBackupPolicyResponseBodyData {
	s.ColdDataBackupInterval = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetColdDataBackupRetention(v int32) *DescribeBackupPolicyResponseBodyData {
	s.ColdDataBackupRetention = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetCrossRegionDataBackupRetention(v int32) *DescribeBackupPolicyResponseBodyData {
	s.CrossRegionDataBackupRetention = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetCrossRegionLogBackupRetention(v int32) *DescribeBackupPolicyResponseBodyData {
	s.CrossRegionLogBackupRetention = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetDBInstanceName(v string) *DescribeBackupPolicyResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetDestCrossRegion(v string) *DescribeBackupPolicyResponseBodyData {
	s.DestCrossRegion = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetForceCleanOnHighSpaceUsage(v int32) *DescribeBackupPolicyResponseBodyData {
	s.ForceCleanOnHighSpaceUsage = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetIsCrossRegionDataBackupEnabled(v bool) *DescribeBackupPolicyResponseBodyData {
	s.IsCrossRegionDataBackupEnabled = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetIsCrossRegionLogBackupEnabled(v bool) *DescribeBackupPolicyResponseBodyData {
	s.IsCrossRegionLogBackupEnabled = &v
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

func (s *DescribeBackupPolicyResponseBodyData) SetLocalLogRetentionNumber(v int32) *DescribeBackupPolicyResponseBodyData {
	s.LocalLogRetentionNumber = &v
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeBackupPolicyResponse) SetStatusCode(v int32) *DescribeBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupPolicyResponse) SetBody(v *DescribeBackupPolicyResponseBody) *DescribeBackupPolicyResponse {
	s.Body = v
	return s
}

type DescribeBackupSetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 111
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-htri0ori2r****
	DBInstanceName  *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DestCrossRegion *string `json:"DestCrossRegion,omitempty" xml:"DestCrossRegion,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeBackupSetRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetRequest) SetBackupSetId(v string) *DescribeBackupSetRequest {
	s.BackupSetId = &v
	return s
}

func (s *DescribeBackupSetRequest) SetDBInstanceName(v string) *DescribeBackupSetRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeBackupSetRequest) SetDestCrossRegion(v string) *DescribeBackupSetRequest {
	s.DestCrossRegion = &v
	return s
}

func (s *DescribeBackupSetRequest) SetRegionId(v string) *DescribeBackupSetRequest {
	s.RegionId = &v
	return s
}

type DescribeBackupSetResponseBody struct {
	Data []*DescribeBackupSetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// successs
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1A6D328C-84B8-40DC-BF49-6C73984D7494
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackupSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetResponseBody) SetData(v []*DescribeBackupSetResponseBodyData) *DescribeBackupSetResponseBody {
	s.Data = v
	return s
}

func (s *DescribeBackupSetResponseBody) SetMessage(v string) *DescribeBackupSetResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeBackupSetResponseBody) SetRequestId(v string) *DescribeBackupSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupSetResponseBody) SetSuccess(v bool) *DescribeBackupSetResponseBody {
	s.Success = &v
	return s
}

type DescribeBackupSetResponseBodyData struct {
	// example:
	//
	// 0
	BackupModel *int32 `json:"BackupModel,omitempty" xml:"BackupModel,omitempty"`
	// example:
	//
	// 111
	BackupSetId *int64 `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// example:
	//
	// 88803195
	BackupSetSize *int64 `json:"BackupSetSize,omitempty" xml:"BackupSetSize,omitempty"`
	// example:
	//
	// 1
	BackupType *int32 `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// example:
	//
	// 1650250861754
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 1650251308000
	EndTime *int64                                      `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OSSList []*DescribeBackupSetResponseBodyDataOSSList `json:"OSSList,omitempty" xml:"OSSList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeBackupSetResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetResponseBodyData) SetBackupModel(v int32) *DescribeBackupSetResponseBodyData {
	s.BackupModel = &v
	return s
}

func (s *DescribeBackupSetResponseBodyData) SetBackupSetId(v int64) *DescribeBackupSetResponseBodyData {
	s.BackupSetId = &v
	return s
}

func (s *DescribeBackupSetResponseBodyData) SetBackupSetSize(v int64) *DescribeBackupSetResponseBodyData {
	s.BackupSetSize = &v
	return s
}

func (s *DescribeBackupSetResponseBodyData) SetBackupType(v int32) *DescribeBackupSetResponseBodyData {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupSetResponseBodyData) SetBeginTime(v int64) *DescribeBackupSetResponseBodyData {
	s.BeginTime = &v
	return s
}

func (s *DescribeBackupSetResponseBodyData) SetEndTime(v int64) *DescribeBackupSetResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupSetResponseBodyData) SetOSSList(v []*DescribeBackupSetResponseBodyDataOSSList) *DescribeBackupSetResponseBodyData {
	s.OSSList = v
	return s
}

func (s *DescribeBackupSetResponseBodyData) SetStatus(v int32) *DescribeBackupSetResponseBodyData {
	s.Status = &v
	return s
}

type DescribeBackupSetResponseBodyDataOSSList struct {
	// example:
	//
	// hins3084_data_20220418110623_qp.xb
	BackupSetFile *string `json:"BackupSetFile,omitempty" xml:"BackupSetFile,omitempty"`
	// example:
	//
	// https://pre-rdsbak-cn-xxx.oss-cn-beijing.aliyuncs.com/custins2255/hins3084_data_20220418110623_qp.xb?OSSAccessKeyId=LTAI5tJEmRFdJ8aBPDR7****&Expires=1650441697&dd=7KJzkUSbXf6dwy
	DownloadLink *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
	// example:
	//
	// http://pre-rdsbak-cn-beijing.oss-cn-beijing-internal.aliyuncs.com/custins2255/hins3084_data_20220418110623_qp.xb?
	IntranetDownloadLink *string `json:"IntranetDownloadLink,omitempty" xml:"IntranetDownloadLink,omitempty"`
	// example:
	//
	// 2022-04-20T08:01:37Z
	LinkExpiredTime *string `json:"LinkExpiredTime,omitempty" xml:"LinkExpiredTime,omitempty"`
}

func (s DescribeBackupSetResponseBodyDataOSSList) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetResponseBodyDataOSSList) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetResponseBodyDataOSSList) SetBackupSetFile(v string) *DescribeBackupSetResponseBodyDataOSSList {
	s.BackupSetFile = &v
	return s
}

func (s *DescribeBackupSetResponseBodyDataOSSList) SetDownloadLink(v string) *DescribeBackupSetResponseBodyDataOSSList {
	s.DownloadLink = &v
	return s
}

func (s *DescribeBackupSetResponseBodyDataOSSList) SetIntranetDownloadLink(v string) *DescribeBackupSetResponseBodyDataOSSList {
	s.IntranetDownloadLink = &v
	return s
}

func (s *DescribeBackupSetResponseBodyDataOSSList) SetLinkExpiredTime(v string) *DescribeBackupSetResponseBodyDataOSSList {
	s.LinkExpiredTime = &v
	return s
}

type DescribeBackupSetResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupSetResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetResponse) SetHeaders(v map[string]*string) *DescribeBackupSetResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupSetResponse) SetStatusCode(v int32) *DescribeBackupSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupSetResponse) SetBody(v *DescribeBackupSetResponseBody) *DescribeBackupSetResponse {
	s.Body = v
	return s
}

type DescribeBackupSetListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-xxxxxx
	DBInstanceName  *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DestCrossRegion *string `json:"DestCrossRegion,omitempty" xml:"DestCrossRegion,omitempty"`
	// example:
	//
	// 1635707845000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 1635707845000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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

func (s *DescribeBackupSetListRequest) SetDestCrossRegion(v string) *DescribeBackupSetListRequest {
	s.DestCrossRegion = &v
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
	Data []*DescribeBackupSetListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1A6D328C-84B8-40DC-BF49-6C73984D7494
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// 0
	BackupModel *int32 `json:"BackupModel,omitempty" xml:"BackupModel,omitempty"`
	// example:
	//
	// 111
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// example:
	//
	// 88803195
	BackupSetSize *int64 `json:"BackupSetSize,omitempty" xml:"BackupSetSize,omitempty"`
	// example:
	//
	// 1
	BackupType *int32 `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// example:
	//
	// 1635706960956
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 1635706960956
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *DescribeBackupSetListResponseBodyData) SetBackupSetId(v string) *DescribeBackupSetListResponseBodyData {
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupSetListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeBackupSetListResponse) SetStatusCode(v int32) *DescribeBackupSetListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupSetListResponse) SetBody(v *DescribeBackupSetListResponseBody) *DescribeBackupSetListResponse {
	s.Body = v
	return s
}

type DescribeBinaryLogListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-hz1fds
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-09 10:27:46
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-09-09 10:27:46
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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

func (s *DescribeBinaryLogListRequest) SetInstanceName(v string) *DescribeBinaryLogListRequest {
	s.InstanceName = &v
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
	LogList []*DescribeBinaryLogListResponseBodyLogList `json:"LogList,omitempty" xml:"LogList,omitempty" type:"Repeated"`
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
	// 2DFF784E-DC31-5BBC-9B25-9261CD9E0AA9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalNumber *int32 `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
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
	// example:
	//
	// 2021-09-09 10:27:46
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 2021-09-09 10:27:46
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// http://polarx-cdc-binlog-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/polardbx_cdc/pxc-hzfd132143sfds1/binlog.000001?Expires=1636469502&OSSAccessKeyId=LT13fds12dsafddsf&Signature=fdpm%bdsfadsa%2F%bdsafdsaf%3D
	DownloadLink *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
	// example:
	//
	// 2021-11-09 10:27:46
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// binlog.000001
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 100
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 536870912
	LogSize *int64 `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	// example:
	//
	// 2021-11-09 10:27:46
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// 0
	PurgeStatus *int32 `json:"PurgeStatus,omitempty" xml:"PurgeStatus,omitempty"`
	// example:
	//
	// 10.110.88.9
	UploadHost *string `json:"UploadHost,omitempty" xml:"UploadHost,omitempty"`
	// example:
	//
	// 2
	UploadStatus *int32 `json:"UploadStatus,omitempty" xml:"UploadStatus,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBinaryLogListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeBinaryLogListResponse) SetStatusCode(v int32) *DescribeBinaryLogListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBinaryLogListResponse) SetBody(v *DescribeBinaryLogListResponseBody) *DescribeBinaryLogListResponse {
	s.Body = v
	return s
}

type DescribeCharacterSetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Data *DescribeCharacterSetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 709C1E40-092D-4A3D-9958-6D7438******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// polarx
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCharacterSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeCharacterSetResponse) SetStatusCode(v int32) *DescribeCharacterSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCharacterSetResponse) SetBody(v *DescribeCharacterSetResponseBody) *DescribeCharacterSetResponse {
	s.Body = v
	return s
}

type DescribeColdDataBasicInfoRequest struct {
	// This parameter is required.
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeColdDataBasicInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeColdDataBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeColdDataBasicInfoRequest) SetDBInstanceName(v string) *DescribeColdDataBasicInfoRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeColdDataBasicInfoRequest) SetRegionId(v string) *DescribeColdDataBasicInfoRequest {
	s.RegionId = &v
	return s
}

type DescribeColdDataBasicInfoResponseBody struct {
	Data      *DescribeColdDataBasicInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeColdDataBasicInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeColdDataBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeColdDataBasicInfoResponseBody) SetData(v *DescribeColdDataBasicInfoResponseBodyData) *DescribeColdDataBasicInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeColdDataBasicInfoResponseBody) SetRequestId(v string) *DescribeColdDataBasicInfoResponseBody {
	s.RequestId = &v
	return s
}

type DescribeColdDataBasicInfoResponseBodyData struct {
	BackupSetCount     *int32   `json:"BackupSetCount,omitempty" xml:"BackupSetCount,omitempty"`
	BackupSetSpaceSize *float64 `json:"BackupSetSpaceSize,omitempty" xml:"BackupSetSpaceSize,omitempty"`
	CloudProduct       *string  `json:"CloudProduct,omitempty" xml:"CloudProduct,omitempty"`
	CurrentSpaceSize   *float64 `json:"CurrentSpaceSize,omitempty" xml:"CurrentSpaceSize,omitempty"`
	DataRedundancyType *string  `json:"DataRedundancyType,omitempty" xml:"DataRedundancyType,omitempty"`
	EnableStatus       *bool    `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	ReadAccessNum      *int64   `json:"ReadAccessNum,omitempty" xml:"ReadAccessNum,omitempty"`
	RegionId           *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VolumeName         *string  `json:"VolumeName,omitempty" xml:"VolumeName,omitempty"`
	WriteAccessNum     *float64 `json:"WriteAccessNum,omitempty" xml:"WriteAccessNum,omitempty"`
}

func (s DescribeColdDataBasicInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeColdDataBasicInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeColdDataBasicInfoResponseBodyData) SetBackupSetCount(v int32) *DescribeColdDataBasicInfoResponseBodyData {
	s.BackupSetCount = &v
	return s
}

func (s *DescribeColdDataBasicInfoResponseBodyData) SetBackupSetSpaceSize(v float64) *DescribeColdDataBasicInfoResponseBodyData {
	s.BackupSetSpaceSize = &v
	return s
}

func (s *DescribeColdDataBasicInfoResponseBodyData) SetCloudProduct(v string) *DescribeColdDataBasicInfoResponseBodyData {
	s.CloudProduct = &v
	return s
}

func (s *DescribeColdDataBasicInfoResponseBodyData) SetCurrentSpaceSize(v float64) *DescribeColdDataBasicInfoResponseBodyData {
	s.CurrentSpaceSize = &v
	return s
}

func (s *DescribeColdDataBasicInfoResponseBodyData) SetDataRedundancyType(v string) *DescribeColdDataBasicInfoResponseBodyData {
	s.DataRedundancyType = &v
	return s
}

func (s *DescribeColdDataBasicInfoResponseBodyData) SetEnableStatus(v bool) *DescribeColdDataBasicInfoResponseBodyData {
	s.EnableStatus = &v
	return s
}

func (s *DescribeColdDataBasicInfoResponseBodyData) SetReadAccessNum(v int64) *DescribeColdDataBasicInfoResponseBodyData {
	s.ReadAccessNum = &v
	return s
}

func (s *DescribeColdDataBasicInfoResponseBodyData) SetRegionId(v string) *DescribeColdDataBasicInfoResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeColdDataBasicInfoResponseBodyData) SetVolumeName(v string) *DescribeColdDataBasicInfoResponseBodyData {
	s.VolumeName = &v
	return s
}

func (s *DescribeColdDataBasicInfoResponseBodyData) SetWriteAccessNum(v float64) *DescribeColdDataBasicInfoResponseBodyData {
	s.WriteAccessNum = &v
	return s
}

type DescribeColdDataBasicInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeColdDataBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeColdDataBasicInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeColdDataBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeColdDataBasicInfoResponse) SetHeaders(v map[string]*string) *DescribeColdDataBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeColdDataBasicInfoResponse) SetStatusCode(v int32) *DescribeColdDataBasicInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeColdDataBasicInfoResponse) SetBody(v *DescribeColdDataBasicInfoResponseBody) *DescribeColdDataBasicInfoResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceAttributeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-aekzaxhezhs5***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *DescribeDBInstanceAttributeRequest) SetResourceGroupId(v string) *DescribeDBInstanceAttributeRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeDBInstanceAttributeResponseBody struct {
	DBInstance *DescribeDBInstanceAttributeResponseBodyDBInstance `json:"DBInstance,omitempty" xml:"DBInstance,omitempty" type:"Struct"`
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CanNotCreateColumnar *bool `json:"CanNotCreateColumnar,omitempty" xml:"CanNotCreateColumnar,omitempty"`
	// example:
	//
	// polarx.x4.xlarge.2e
	CnNodeClassCode *string `json:"CnNodeClassCode,omitempty" xml:"CnNodeClassCode,omitempty"`
	// example:
	//
	// 2
	CnNodeCount             *int32    `json:"CnNodeCount,omitempty" xml:"CnNodeCount,omitempty"`
	ColumnarInstanceName    *string   `json:"ColumnarInstanceName,omitempty" xml:"ColumnarInstanceName,omitempty"`
	ColumnarReadDBInstances []*string `json:"ColumnarReadDBInstances,omitempty" xml:"ColumnarReadDBInstances,omitempty" type:"Repeated"`
	// example:
	//
	// drds_polarxpost_public_cn
	CommodityCode    *string                                                       `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ConnAddrs        []*DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs `json:"ConnAddrs,omitempty" xml:"ConnAddrs,omitempty" type:"Repeated"`
	ConnectionString *string                                                       `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// example:
	//
	// 2021-08-31T08:56:25.000+0000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// ReadWrite
	DBInstanceType *string `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	// example:
	//
	// polarx.x4.large.2e
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// example:
	//
	// 2
	DBNodeCount *int32                                                      `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	DBNodes     []*DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes `json:"DBNodes,omitempty" xml:"DBNodes,omitempty" type:"Repeated"`
	// example:
	//
	// polarx
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// example:
	//
	// 5.5
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// example:
	//
	// test instance
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DifferentDNSpec *bool   `json:"DifferentDNSpec,omitempty" xml:"DifferentDNSpec,omitempty"`
	// example:
	//
	// mysql.x8.large.25
	DnNodeClassCode *string `json:"DnNodeClassCode,omitempty" xml:"DnNodeClassCode,omitempty"`
	// example:
	//
	// 2
	DnNodeCount *int32 `json:"DnNodeCount,omitempty" xml:"DnNodeCount,omitempty"`
	// example:
	//
	// polarx
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 2022-08-31T16:00:00.000+0000
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// example:
	//
	// false
	Expired         *string                                                           `json:"Expired,omitempty" xml:"Expired,omitempty"`
	GdnInstanceName *string                                                           `json:"GdnInstanceName,omitempty" xml:"GdnInstanceName,omitempty"`
	GdnMemberList   []*DescribeDBInstanceAttributeResponseBodyDBInstanceGdnMemberList `json:"GdnMemberList,omitempty" xml:"GdnMemberList,omitempty" type:"Repeated"`
	GdnRole         *string                                                           `json:"GdnRole,omitempty" xml:"GdnRole,omitempty"`
	// example:
	//
	// pxc-zkralxpc5d****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 18
	KindCode *int32 `json:"KindCode,omitempty" xml:"KindCode,omitempty"`
	// This parameter is required.
	LTSVersions []*string `json:"LTSVersions,omitempty" xml:"LTSVersions,omitempty" type:"Repeated"`
	// example:
	//
	// polarx-kernel_5.4.11-16301083_xcluster-20210805
	LatestMinorVersion *string `json:"LatestMinorVersion,omitempty" xml:"LatestMinorVersion,omitempty"`
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// example:
	//
	// 06:00
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// example:
	//
	// 06:00
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// example:
	//
	// polarx-kernel_5.4.11-16301083_xcluster-20210805
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	// example:
	//
	// VPC
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// 3306
	Port              *string `json:"Port,omitempty" xml:"Port,omitempty"`
	PrimaryInstanceId *string `json:"PrimaryInstanceId,omitempty" xml:"PrimaryInstanceId,omitempty"`
	// 主可用区。
	//
	// This parameter is required.
	PrimaryZone     *string   `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	ReadDBInstances []*string `json:"ReadDBInstances,omitempty" xml:"ReadDBInstances,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-xxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// false
	RightsSeparationEnabled *bool `json:"RightsSeparationEnabled,omitempty" xml:"RightsSeparationEnabled,omitempty"`
	// example:
	//
	// disabled
	RightsSeparationStatus *string `json:"RightsSeparationStatus,omitempty" xml:"RightsSeparationStatus,omitempty"`
	// 次可用区。
	SecondaryZone *string `json:"SecondaryZone,omitempty" xml:"SecondaryZone,omitempty"`
	// example:
	//
	// enterprise
	Series       *string `json:"Series,omitempty" xml:"Series,omitempty"`
	SpecCategory *string `json:"SpecCategory,omitempty" xml:"SpecCategory,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 17042505728
	StorageUsed *int64                                                     `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	TagSet      []*DescribeDBInstanceAttributeResponseBodyDBInstanceTagSet `json:"TagSet,omitempty" xml:"TagSet,omitempty" type:"Repeated"`
	// 第三可用区。
	TertiaryZone *string `json:"TertiaryZone,omitempty" xml:"TertiaryZone,omitempty"`
	// 拓扑类型：
	//
	// - **3azones**：三可用区；
	//
	// - **1azone**：单可用区。
	//
	// This parameter is required.
	TopologyType *string `json:"TopologyType,omitempty" xml:"TopologyType,omitempty"`
	// example:
	//
	// ReadWrite
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-xxxxx
	VPCId     *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstance) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetCanNotCreateColumnar(v bool) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.CanNotCreateColumnar = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetCnNodeClassCode(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.CnNodeClassCode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetCnNodeCount(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.CnNodeCount = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetColumnarInstanceName(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.ColumnarInstanceName = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetColumnarReadDBInstances(v []*string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.ColumnarReadDBInstances = v
	return s
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

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetDifferentDNSpec(v bool) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.DifferentDNSpec = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetDnNodeClassCode(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.DnNodeClassCode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetDnNodeCount(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.DnNodeCount = &v
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

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetGdnInstanceName(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.GdnInstanceName = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetGdnMemberList(v []*DescribeDBInstanceAttributeResponseBodyDBInstanceGdnMemberList) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.GdnMemberList = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetGdnRole(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.GdnRole = &v
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

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetLTSVersions(v []*string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.LTSVersions = v
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

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetPrimaryInstanceId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.PrimaryInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetPrimaryZone(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.PrimaryZone = &v
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

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetResourceGroupId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.ResourceGroupId = &v
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

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetSecondaryZone(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.SecondaryZone = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetSeries(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.Series = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetSpecCategory(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.SpecCategory = &v
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

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetTagSet(v []*DescribeDBInstanceAttributeResponseBodyDBInstanceTagSet) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.TagSet = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetTertiaryZone(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.TertiaryZone = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetTopologyType(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.TopologyType = &v
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
	// example:
	//
	// polardbx-xxx.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// example:
	//
	// 3306
	Port *int64 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// VPC
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-xxxxxx
	VPCId     *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// pxc-zkralxpc5d****
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
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

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) SetVpcInstanceId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs {
	s.VpcInstanceId = &v
	return s
}

type DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes struct {
	ComputeNodeId *string `json:"ComputeNodeId,omitempty" xml:"ComputeNodeId,omitempty"`
	// example:
	//
	// pxc-xdb-xxxxxx
	DataNodeId *string `json:"DataNodeId,omitempty" xml:"DataNodeId,omitempty"`
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// polarx.x4.large.2e
	NodeClass *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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

type DescribeDBInstanceAttributeResponseBodyDBInstanceGdnMemberList struct {
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	Role       *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstanceGdnMemberList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstanceGdnMemberList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceGdnMemberList) SetMemberName(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceGdnMemberList {
	s.MemberName = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceGdnMemberList) SetRole(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceGdnMemberList {
	s.Role = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceGdnMemberList) SetStatus(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceGdnMemberList {
	s.Status = &v
	return s
}

type DescribeDBInstanceAttributeResponseBodyDBInstanceTagSet struct {
	// example:
	//
	// key2
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value2
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstanceTagSet) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstanceTagSet) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceTagSet) SetKey(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceTagSet {
	s.Key = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceTagSet) SetValue(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceTagSet {
	s.Value = &v
	return s
}

type DescribeDBInstanceAttributeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeDBInstanceAttributeResponse) SetStatusCode(v int32) *DescribeDBInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponse) SetBody(v *DescribeDBInstanceAttributeResponseBody) *DescribeDBInstanceAttributeResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// htap
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Data *DescribeDBInstanceConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// htap
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeDBInstanceConfigResponse) SetStatusCode(v int32) *DescribeDBInstanceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceConfigResponse) SetBody(v *DescribeDBInstanceConfigResponseBody) *DescribeDBInstanceConfigResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceHARequest struct {
	// This parameter is required.
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBInstanceHARequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceHARequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceHARequest) SetDBInstanceName(v string) *DescribeDBInstanceHARequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstanceHARequest) SetRegionId(v string) *DescribeDBInstanceHARequest {
	s.RegionId = &v
	return s
}

type DescribeDBInstanceHAResponseBody struct {
	Data      *DescribeDBInstanceHAResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDBInstanceHAResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceHAResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceHAResponseBody) SetData(v *DescribeDBInstanceHAResponseBodyData) *DescribeDBInstanceHAResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDBInstanceHAResponseBody) SetMessage(v string) *DescribeDBInstanceHAResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDBInstanceHAResponseBody) SetRequestId(v string) *DescribeDBInstanceHAResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceHAResponseBody) SetSuccess(v bool) *DescribeDBInstanceHAResponseBody {
	s.Success = &v
	return s
}

type DescribeDBInstanceHAResponseBodyData struct {
	PrimaryAzoneId    *string `json:"PrimaryAzoneId,omitempty" xml:"PrimaryAzoneId,omitempty"`
	PrimaryRegionId   *string `json:"PrimaryRegionId,omitempty" xml:"PrimaryRegionId,omitempty"`
	SecondaryAzoneId  *string `json:"SecondaryAzoneId,omitempty" xml:"SecondaryAzoneId,omitempty"`
	SecondaryRegionId *string `json:"SecondaryRegionId,omitempty" xml:"SecondaryRegionId,omitempty"`
	TopologyType      *string `json:"TopologyType,omitempty" xml:"TopologyType,omitempty"`
}

func (s DescribeDBInstanceHAResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceHAResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceHAResponseBodyData) SetPrimaryAzoneId(v string) *DescribeDBInstanceHAResponseBodyData {
	s.PrimaryAzoneId = &v
	return s
}

func (s *DescribeDBInstanceHAResponseBodyData) SetPrimaryRegionId(v string) *DescribeDBInstanceHAResponseBodyData {
	s.PrimaryRegionId = &v
	return s
}

func (s *DescribeDBInstanceHAResponseBodyData) SetSecondaryAzoneId(v string) *DescribeDBInstanceHAResponseBodyData {
	s.SecondaryAzoneId = &v
	return s
}

func (s *DescribeDBInstanceHAResponseBodyData) SetSecondaryRegionId(v string) *DescribeDBInstanceHAResponseBodyData {
	s.SecondaryRegionId = &v
	return s
}

func (s *DescribeDBInstanceHAResponseBodyData) SetTopologyType(v string) *DescribeDBInstanceHAResponseBodyData {
	s.TopologyType = &v
	return s
}

type DescribeDBInstanceHAResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceHAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceHAResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceHAResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceHAResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceHAResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceHAResponse) SetStatusCode(v int32) *DescribeDBInstanceHAResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceHAResponse) SetBody(v *DescribeDBInstanceHAResponseBody) *DescribeDBInstanceHAResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceSSLRequest struct {
	// example:
	//
	// pxc-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Data *DescribeDBInstanceSSLResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// pxc-sddddddcym7g7w****.polarx.singapore.rds.aliyuncs.com
	CertCommonName *string `json:"CertCommonName,omitempty" xml:"CertCommonName,omitempty"`
	// example:
	//
	// false
	SSLEnabled *bool `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
	// example:
	//
	// 2022-11-04T09:39:07Z
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceSSLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeDBInstanceSSLResponse) SetStatusCode(v int32) *DescribeDBInstanceSSLResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceSSLResponse) SetBody(v *DescribeDBInstanceSSLResponseBody) *DescribeDBInstanceSSLResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceTDERequest struct {
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Data *DescribeDBInstanceTDEResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 0
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceTDEResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeDBInstanceTDEResponse) SetStatusCode(v int32) *DescribeDBInstanceTDEResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceTDEResponse) SetBody(v *DescribeDBInstanceTDEResponseBody) *DescribeDBInstanceTDEResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceTopologyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MinuteSimple   *bool   `json:"MinuteSimple,omitempty" xml:"MinuteSimple,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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

func (s *DescribeDBInstanceTopologyRequest) SetEndTime(v string) *DescribeDBInstanceTopologyRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBInstanceTopologyRequest) SetMinuteSimple(v bool) *DescribeDBInstanceTopologyRequest {
	s.MinuteSimple = &v
	return s
}

func (s *DescribeDBInstanceTopologyRequest) SetRegionId(v string) *DescribeDBInstanceTopologyRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceTopologyRequest) SetStartTime(v string) *DescribeDBInstanceTopologyRequest {
	s.StartTime = &v
	return s
}

type DescribeDBInstanceTopologyResponseBody struct {
	Data *DescribeDBInstanceTopologyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// lvs
	DBInstanceConnType *string `json:"DBInstanceConnType,omitempty" xml:"DBInstanceConnType,omitempty"`
	// example:
	//
	// 2021-10-21T10:30:45Z 04:00:00
	DBInstanceCreateTime *string `json:"DBInstanceCreateTime,omitempty" xml:"DBInstanceCreateTime,omitempty"`
	// example:
	//
	// pxc-sprcym7g7wj7k
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// example:
	//
	// 304726047
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// pxc-sprcym7g7w****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// 8
	DBInstanceStatus *int32 `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// example:
	//
	// TDE_MODIFYING
	DBInstanceStatusDescription *string `json:"DBInstanceStatusDescription,omitempty" xml:"DBInstanceStatusDescription,omitempty"`
	// example:
	//
	// 1
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// example:
	//
	// polarx
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 2.0
	EngineVersion *string                                                                        `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	HistoryItems  []*DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems `json:"HistoryItems,omitempty" xml:"HistoryItems,omitempty" type:"Repeated"`
	Items         []*DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems        `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	LockMode   *int32  `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// example:
	//
	// 05:00:00
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// example:
	//
	// 04:00:00
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
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

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetHistoryItems(v []*DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.HistoryItems = v
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

type DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems struct {
	Activated       *bool   `json:"Activated,omitempty" xml:"Activated,omitempty"`
	Azone           *string `json:"Azone,omitempty" xml:"Azone,omitempty"`
	CharacterType   *string `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceName  *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	PhyInstanceName *string `json:"PhyInstanceName,omitempty" xml:"PhyInstanceName,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Role            *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) SetActivated(v bool) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems {
	s.Activated = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) SetAzone(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems {
	s.Azone = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) SetCharacterType(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems {
	s.CharacterType = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) SetDBInstanceId(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) SetDBInstanceName(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) SetPhyInstanceName(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems {
	s.PhyInstanceName = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) SetRegion(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems {
	s.Region = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) SetRole(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems {
	s.Role = &v
	return s
}

type DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems struct {
	// example:
	//
	// true
	Activated *bool `json:"Activated,omitempty" xml:"Activated,omitempty"`
	// example:
	//
	// cn-hangzhou-a
	Azone              *string                                                                              `json:"Azone,omitempty" xml:"Azone,omitempty"`
	AzoneRoleList      []*DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsAzoneRoleList `json:"AzoneRoleList,omitempty" xml:"AzoneRoleList,omitempty" type:"Repeated"`
	CharacterType      *string                                                                              `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	ConnectionIp       []*DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsConnectionIp  `json:"ConnectionIp,omitempty" xml:"ConnectionIp,omitempty" type:"Repeated"`
	DBInstanceConnType *int32                                                                               `json:"DBInstanceConnType,omitempty" xml:"DBInstanceConnType,omitempty"`
	// example:
	//
	// 2021-10-21T10:30:45Z
	DBInstanceCreateTime  *string `json:"DBInstanceCreateTime,omitempty" xml:"DBInstanceCreateTime,omitempty"`
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// example:
	//
	// 304726049
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// pxc-i-tk6t4z****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// 8
	DBInstanceStatus            *int32  `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	DBInstanceStatusDescription *string `json:"DBInstanceStatusDescription,omitempty" xml:"DBInstanceStatusDescription,omitempty"`
	// example:
	//
	// 3145728
	DiskSize *int64 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// example:
	//
	// mysql
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// example:
	//
	// 0
	LockMode          *int32  `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	LockReason        *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	MaintainEndTime   *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// example:
	//
	// 4000
	MaxConnections *int32 `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	// example:
	//
	// 7000
	MaxIops *int32 `json:"MaxIops,omitempty" xml:"MaxIops,omitempty"`
	// example:
	//
	// polarx.x4.large.2e
	NodeClass       *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	PhyInstanceName *string `json:"PhyInstanceName,omitempty" xml:"PhyInstanceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// master
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// 0
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageUsed *string `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	// example:
	//
	// polarx-cdc-kernel-2.0.0-3985896
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetActivated(v bool) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.Activated = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetAzone(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.Azone = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetAzoneRoleList(v []*DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsAzoneRoleList) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.AzoneRoleList = v
	return s
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

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetNodeClass(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.NodeClass = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetPhyInstanceName(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.PhyInstanceName = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetRegion(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.Region = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetRole(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.Role = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetStatus(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.Status = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetStorageUsed(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.StorageUsed = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetVersion(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.Version = &v
	return s
}

type DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsAzoneRoleList struct {
	// example:
	//
	// cn-hangzhou-a
	Azone *string `json:"Azone,omitempty" xml:"Azone,omitempty"`
	// example:
	//
	// leader
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsAzoneRoleList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsAzoneRoleList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsAzoneRoleList) SetAzone(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsAzoneRoleList {
	s.Azone = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsAzoneRoleList) SetRole(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsAzoneRoleList {
	s.Role = &v
	return s
}

type DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsConnectionIp struct {
	// example:
	//
	// pxc-xdb-m-pxcdym7g7w********.mysql.singapore.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// example:
	//
	// 1
	DBInstanceNetType *int32 `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceTopologyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeDBInstanceTopologyResponse) SetStatusCode(v int32) *DescribeDBInstanceTopologyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponse) SetBody(v *DescribeDBInstanceTopologyResponseBody) *DescribeDBInstanceTopologyResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceViaEndpointRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-hz*******.polarx.rds.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBInstanceViaEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceViaEndpointRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceViaEndpointRequest) SetEndpoint(v string) *DescribeDBInstanceViaEndpointRequest {
	s.Endpoint = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointRequest) SetRegionId(v string) *DescribeDBInstanceViaEndpointRequest {
	s.RegionId = &v
	return s
}

type DescribeDBInstanceViaEndpointResponseBody struct {
	DBInstance *DescribeDBInstanceViaEndpointResponseBodyDBInstance `json:"DBInstance,omitempty" xml:"DBInstance,omitempty" type:"Struct"`
	// example:
	//
	// 14036EBE-CF2E-44DB-ACE9-AC6157D3A6FC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceViaEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceViaEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceViaEndpointResponseBody) SetDBInstance(v *DescribeDBInstanceViaEndpointResponseBodyDBInstance) *DescribeDBInstanceViaEndpointResponseBody {
	s.DBInstance = v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBody) SetRequestId(v string) *DescribeDBInstanceViaEndpointResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBInstanceViaEndpointResponseBodyDBInstance struct {
	// example:
	//
	// polarx.x4.xlarge.2e
	CnNodeClassCode *string `json:"CnNodeClassCode,omitempty" xml:"CnNodeClassCode,omitempty"`
	// example:
	//
	// 2
	CnNodeCount *int32 `json:"CnNodeCount,omitempty" xml:"CnNodeCount,omitempty"`
	// example:
	//
	// drds_polarxpost_public_cn
	CommodityCode *string                                                         `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ConnAddrs     []*DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs `json:"ConnAddrs,omitempty" xml:"ConnAddrs,omitempty" type:"Repeated"`
	// example:
	//
	// pxc-sprpx766vo****.polarx.singapore.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// example:
	//
	// 2021-08-31T08:56:25.000+0000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// ReadWrite
	DBInstanceType *string `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	// example:
	//
	// polarx.x4.large.2e
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// example:
	//
	// 2
	DBNodeCount *int32                                                        `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	DBNodes     []*DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes `json:"DBNodes,omitempty" xml:"DBNodes,omitempty" type:"Repeated"`
	// example:
	//
	// polarx
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// example:
	//
	// 5.5
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// example:
	//
	// test instance
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// mysql.x8.large.25
	DnNodeClassCode *string `json:"DnNodeClassCode,omitempty" xml:"DnNodeClassCode,omitempty"`
	// example:
	//
	// 2
	DnNodeCount *int32 `json:"DnNodeCount,omitempty" xml:"DnNodeCount,omitempty"`
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 2022-08-31T16:00:00.000+0000
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// example:
	//
	// false
	Expired *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// example:
	//
	// pxc-zkralxpc5d****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 18
	KindCode *int32 `json:"KindCode,omitempty" xml:"KindCode,omitempty"`
	// This parameter is required.
	LTSVersions []*string `json:"LTSVersions,omitempty" xml:"LTSVersions,omitempty" type:"Repeated"`
	// example:
	//
	// polarx-kernel_5.4.11-16301083_xcluster-20210805
	LatestMinorVersion *string `json:"LatestMinorVersion,omitempty" xml:"LatestMinorVersion,omitempty"`
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// example:
	//
	// 06:00
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// example:
	//
	// 06:00
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// example:
	//
	// polarx-kernel_5.4.11-16301083_xcluster-20210805
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	// example:
	//
	// VPC
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// 3306
	Port            *string   `json:"Port,omitempty" xml:"Port,omitempty"`
	ReadDBInstances []*string `json:"ReadDBInstances,omitempty" xml:"ReadDBInstances,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-xxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// false
	RightsSeparationEnabled *bool `json:"RightsSeparationEnabled,omitempty" xml:"RightsSeparationEnabled,omitempty"`
	// example:
	//
	// disabled
	RightsSeparationStatus *string `json:"RightsSeparationStatus,omitempty" xml:"RightsSeparationStatus,omitempty"`
	// example:
	//
	// enterprise
	Series *string `json:"Series,omitempty" xml:"Series,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 17042505728
	StorageUsed *int64                                                       `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	TagSet      []*DescribeDBInstanceViaEndpointResponseBodyDBInstanceTagSet `json:"TagSet,omitempty" xml:"TagSet,omitempty" type:"Repeated"`
	// example:
	//
	// ReadWrite
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-xxxxx
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// example:
	//
	// vsw-********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstanceViaEndpointResponseBodyDBInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceViaEndpointResponseBodyDBInstance) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetCnNodeClassCode(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.CnNodeClassCode = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetCnNodeCount(v int32) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.CnNodeCount = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetCommodityCode(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetConnAddrs(v []*DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.ConnAddrs = v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetConnectionString(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetCreateTime(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetDBInstanceType(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetDBNodeClass(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetDBNodeCount(v int32) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.DBNodeCount = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetDBNodes(v []*DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.DBNodes = v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetDBType(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.DBType = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetDBVersion(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetDescription(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.Description = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetDnNodeClassCode(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.DnNodeClassCode = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetDnNodeCount(v int32) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.DnNodeCount = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetEngine(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetExpireDate(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.ExpireDate = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetExpired(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.Expired = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetId(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.Id = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetKindCode(v int32) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.KindCode = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetLTSVersions(v []*string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.LTSVersions = v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetLatestMinorVersion(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.LatestMinorVersion = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetLockMode(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetMaintainEndTime(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetMaintainStartTime(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetMinorVersion(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.MinorVersion = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetNetwork(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.Network = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetPayType(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetPort(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetReadDBInstances(v []*string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.ReadDBInstances = v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetRegionId(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetResourceGroupId(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetRightsSeparationEnabled(v bool) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.RightsSeparationEnabled = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetRightsSeparationStatus(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.RightsSeparationStatus = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetSeries(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.Series = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetStatus(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.Status = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetStorageUsed(v int64) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.StorageUsed = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetTagSet(v []*DescribeDBInstanceViaEndpointResponseBodyDBInstanceTagSet) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.TagSet = v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetType(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.Type = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetVPCId(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.VPCId = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetVSwitchId(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetZoneId(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.ZoneId = &v
	return s
}

type DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs struct {
	// example:
	//
	// polardbx-xxx.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// example:
	//
	// 3306
	Port *int64 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// VPC
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-xxxxxx
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// example:
	//
	// vsw-*********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// pxc-zkralxpc5d****
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs) SetConnectionString(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs) SetPort(v int64) *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs) SetType(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs {
	s.Type = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs) SetVPCId(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs {
	s.VPCId = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs) SetVSwitchId(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs) SetVpcInstanceId(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs {
	s.VpcInstanceId = &v
	return s
}

type DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes struct {
	// example:
	//
	// pxc-i-********
	ComputeNodeId *string `json:"ComputeNodeId,omitempty" xml:"ComputeNodeId,omitempty"`
	// example:
	//
	// pxc-xdb-xxxxxx
	DataNodeId *string `json:"DataNodeId,omitempty" xml:"DataNodeId,omitempty"`
	// example:
	//
	// pxi-*********
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// polarx.x4.large.2e
	NodeClass *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes) SetComputeNodeId(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes {
	s.ComputeNodeId = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes) SetDataNodeId(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes {
	s.DataNodeId = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes) SetId(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes {
	s.Id = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes) SetNodeClass(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes {
	s.NodeClass = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes) SetRegionId(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes) SetZoneId(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes {
	s.ZoneId = &v
	return s
}

type DescribeDBInstanceViaEndpointResponseBodyDBInstanceTagSet struct {
	// example:
	//
	// key2
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value2
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstanceViaEndpointResponseBodyDBInstanceTagSet) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceViaEndpointResponseBodyDBInstanceTagSet) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceTagSet) SetKey(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstanceTagSet {
	s.Key = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceTagSet) SetValue(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstanceTagSet {
	s.Value = &v
	return s
}

type DescribeDBInstanceViaEndpointResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceViaEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceViaEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceViaEndpointResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceViaEndpointResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceViaEndpointResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponse) SetStatusCode(v int32) *DescribeDBInstanceViaEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponse) SetBody(v *DescribeDBInstanceViaEndpointResponseBody) *DescribeDBInstanceViaEndpointResponse {
	s.Body = v
	return s
}

type DescribeDBInstancesRequest struct {
	DbVersion *string `json:"DbVersion,omitempty" xml:"DbVersion,omitempty"`
	// example:
	//
	// dinga93c84f4d***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MustHasCdc *bool   `json:"MustHasCdc,omitempty" xml:"MustHasCdc,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfmyst47hjw***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// enterprise
	Series *string `json:"Series,omitempty" xml:"Series,omitempty"`
	// example:
	//
	// [{\\"TagKey\\":\\"test\\",\\"TagValue\\":\\"test-value\\"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s DescribeDBInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesRequest) SetDbVersion(v string) *DescribeDBInstancesRequest {
	s.DbVersion = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetInstanceId(v string) *DescribeDBInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetMustHasCdc(v bool) *DescribeDBInstancesRequest {
	s.MustHasCdc = &v
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

func (s *DescribeDBInstancesRequest) SetRegionId(v string) *DescribeDBInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetResourceGroupId(v string) *DescribeDBInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetSeries(v string) *DescribeDBInstancesRequest {
	s.Series = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetTags(v string) *DescribeDBInstancesRequest {
	s.Tags = &v
	return s
}

type DescribeDBInstancesResponseBody struct {
	DBInstances []*DescribeDBInstancesResponseBodyDBInstances `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Repeated"`
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
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 30
	TotalNumber *int32 `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
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
	// example:
	//
	// pxc-c-dmlgit****
	CdcInstanceName *string `json:"CdcInstanceName,omitempty" xml:"CdcInstanceName,omitempty"`
	// example:
	//
	// polarx.x4.large.2e
	CnNodeClassCode *string `json:"CnNodeClassCode,omitempty" xml:"CnNodeClassCode,omitempty"`
	// example:
	//
	// 2
	CnNodeCount             *int32    `json:"CnNodeCount,omitempty" xml:"CnNodeCount,omitempty"`
	ColumnarInstanceName    *string   `json:"ColumnarInstanceName,omitempty" xml:"ColumnarInstanceName,omitempty"`
	ColumnarReadDBInstances []*string `json:"ColumnarReadDBInstances,omitempty" xml:"ColumnarReadDBInstances,omitempty" type:"Repeated"`
	// example:
	//
	// drds_polarxpre_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// example:
	//
	// true
	ContainBinlogX *bool   `json:"ContainBinlogX,omitempty" xml:"ContainBinlogX,omitempty"`
	CpuType        *string `json:"CpuType,omitempty" xml:"CpuType,omitempty"`
	// example:
	//
	// 2021-11-01T03:49:50.000+0000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// pxc-xxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// polarx
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// example:
	//
	// 5.7
	DBVersion   *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// mysql.n4.medium.25
	DnNodeClassCode *string `json:"DnNodeClassCode,omitempty" xml:"DnNodeClassCode,omitempty"`
	// example:
	//
	// 2
	DnNodeCount *int32 `json:"DnNodeCount,omitempty" xml:"DnNodeCount,omitempty"`
	// example:
	//
	// polarx
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 2021-12-01T16:00:00.000+0000
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// false
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// example:
	//
	// pxc-hzr2yeov9jmg3z
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// Unlock
	LockMode   *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// example:
	//
	// polarx-kernel_5.4.12-16349923_xcluster-20210926
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	// example:
	//
	// VPC
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// example:
	//
	// polarx.x4.large.2e
	NodeClass *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	// example:
	//
	// 5
	NodeCount *int32                                             `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	Nodes     []*DescribeDBInstancesResponseBodyDBInstancesNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// Prepaid
	PayType           *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PrimaryInstanceId *string `json:"PrimaryInstanceId,omitempty" xml:"PrimaryInstanceId,omitempty"`
	// 主可用区。
	//
	// This parameter is required.
	PrimaryZone     *string   `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	ReadDBInstances []*string `json:"ReadDBInstances,omitempty" xml:"ReadDBInstances,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-xxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 次可用区。
	SecondaryZone *string `json:"SecondaryZone,omitempty" xml:"SecondaryZone,omitempty"`
	// example:
	//
	// enterprise
	Series *string `json:"Series,omitempty" xml:"Series,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 40658534400
	StorageUsed *int64 `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	// example:
	//
	// true
	SupportBinlogX *bool                                               `json:"SupportBinlogX,omitempty" xml:"SupportBinlogX,omitempty"`
	TagSet         []*DescribeDBInstancesResponseBodyDBInstancesTagSet `json:"TagSet,omitempty" xml:"TagSet,omitempty" type:"Repeated"`
	// 第三可用区。
	TertiaryZone *string `json:"TertiaryZone,omitempty" xml:"TertiaryZone,omitempty"`
	// 拓扑类型：
	//
	// - **3azones**：三可用区；
	//
	// - **1azone**：单可用区。
	//
	// This parameter is required.
	TopologyType *string `json:"TopologyType,omitempty" xml:"TopologyType,omitempty"`
	// example:
	//
	// ReadWrite
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// VPCID
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// example:
	//
	// cn-hangzhou-g
	ZoneId  *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	GdnRole *string `json:"gdnRole,omitempty" xml:"gdnRole,omitempty"`
	IsInGdn *bool   `json:"isInGdn,omitempty" xml:"isInGdn,omitempty"`
}

func (s DescribeDBInstancesResponseBodyDBInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyDBInstances) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetCdcInstanceName(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.CdcInstanceName = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetCnNodeClassCode(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.CnNodeClassCode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetCnNodeCount(v int32) *DescribeDBInstancesResponseBodyDBInstances {
	s.CnNodeCount = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetColumnarInstanceName(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.ColumnarInstanceName = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetColumnarReadDBInstances(v []*string) *DescribeDBInstancesResponseBodyDBInstances {
	s.ColumnarReadDBInstances = v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetCommodityCode(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetContainBinlogX(v bool) *DescribeDBInstancesResponseBodyDBInstances {
	s.ContainBinlogX = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetCpuType(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.CpuType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetCreateTime(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetDBInstanceName(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.DBInstanceName = &v
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

func (s *DescribeDBInstancesResponseBodyDBInstances) SetDnNodeClassCode(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.DnNodeClassCode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetDnNodeCount(v int32) *DescribeDBInstancesResponseBodyDBInstances {
	s.DnNodeCount = &v
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

func (s *DescribeDBInstancesResponseBodyDBInstances) SetPrimaryInstanceId(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.PrimaryInstanceId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetPrimaryZone(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.PrimaryZone = &v
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

func (s *DescribeDBInstancesResponseBodyDBInstances) SetResourceGroupId(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetSecondaryZone(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.SecondaryZone = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetSeries(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.Series = &v
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

func (s *DescribeDBInstancesResponseBodyDBInstances) SetSupportBinlogX(v bool) *DescribeDBInstancesResponseBodyDBInstances {
	s.SupportBinlogX = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetTagSet(v []*DescribeDBInstancesResponseBodyDBInstancesTagSet) *DescribeDBInstancesResponseBodyDBInstances {
	s.TagSet = v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetTertiaryZone(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.TertiaryZone = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetTopologyType(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.TopologyType = &v
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

func (s *DescribeDBInstancesResponseBodyDBInstances) SetGdnRole(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.GdnRole = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetIsInGdn(v bool) *DescribeDBInstancesResponseBodyDBInstances {
	s.IsInGdn = &v
	return s
}

type DescribeDBInstancesResponseBodyDBInstancesNodes struct {
	// example:
	//
	// polarx.x4.large.2e
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// cn-hangzhou-g-aliyun
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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

type DescribeDBInstancesResponseBodyDBInstancesTagSet struct {
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstancesResponseBodyDBInstancesTagSet) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyDBInstancesTagSet) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyDBInstancesTagSet) SetKey(v string) *DescribeDBInstancesResponseBodyDBInstancesTagSet {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesTagSet) SetValue(v string) *DescribeDBInstancesResponseBodyDBInstancesTagSet {
	s.Value = &v
	return s
}

type DescribeDBInstancesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeDBInstancesResponse) SetStatusCode(v int32) *DescribeDBInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstancesResponse) SetBody(v *DescribeDBInstancesResponseBody) *DescribeDBInstancesResponse {
	s.Body = v
	return s
}

type DescribeDBNodePerformanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// polarx_cn
	CharacterType *string `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-*******
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-i-******,pxc-i-*******
	DBNodeIds *string `json:"DBNodeIds,omitempty" xml:"DBNodeIds,omitempty"`
	// example:
	//
	// master
	DBNodeRole *string `json:"DBNodeRole,omitempty" xml:"DBNodeRole,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2012-06-18T15:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Cpu_Usage,Mem_Usage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2012-06-08T15:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// 2021-10-20T02:00Z
	EndTime         *string                                               `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PerformanceKeys *DescribeDBNodePerformanceResponseBodyPerformanceKeys `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Struct"`
	// example:
	//
	// EFB5E10B-5268-170F-A378-9AF86CCEACC8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2021-10-20T00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
	// example:
	//
	// pxc-i-********
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// example:
	//
	// Logic_TPS
	Measurement *string `json:"Measurement,omitempty" xml:"Measurement,omitempty"`
	// example:
	//
	// logic_tps
	MetricName *string                                                                    `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Points     *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPoints `json:"Points,omitempty" xml:"Points,omitempty" type:"Struct"`
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
	// example:
	//
	// 1600822800000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// 42.38
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBNodePerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeDBNodePerformanceResponse) SetStatusCode(v int32) *DescribeDBNodePerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBNodePerformanceResponse) SetBody(v *DescribeDBNodePerformanceResponseBody) *DescribeDBNodePerformanceResponse {
	s.Body = v
	return s
}

type DescribeDbListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// db_name
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Data []*DescribeDbListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Accounts []*DescribeDbListResponseBodyDataAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	// example:
	//
	// utf8mb4
	CharacterSetName *string `json:"CharacterSetName,omitempty" xml:"CharacterSetName,omitempty"`
	// example:
	//
	// test
	DBDescription *string `json:"DBDescription,omitempty" xml:"DBDescription,omitempty"`
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// test
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
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
	// example:
	//
	// root4test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// ReadWrite
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDbListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeDbListResponse) SetStatusCode(v int32) *DescribeDbListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDbListResponse) SetBody(v *DescribeDbListResponseBody) *DescribeDbListResponse {
	s.Body = v
	return s
}

type DescribeDistributeTableListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sbtest1
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Data *DescribeDistributeTableListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// id
	DbKey *string `json:"DbKey,omitempty" xml:"DbKey,omitempty"`
	// example:
	//
	// sbtest1
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// multi
	TableType *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
	// example:
	//
	// “”
	TbKey *string `json:"TbKey,omitempty" xml:"TbKey,omitempty"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDistributeTableListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeDistributeTableListResponse) SetStatusCode(v int32) *DescribeDistributeTableListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDistributeTableListResponse) SetBody(v *DescribeDistributeTableListResponseBody) *DescribeDistributeTableListResponse {
	s.Body = v
	return s
}

type DescribeEventsRequest struct {
	// example:
	//
	// 2021-10-18T03:07:25Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 2021-10-18T03:07:25Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventsRequest) SetEndTime(v string) *DescribeEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEventsRequest) SetPageNumber(v int32) *DescribeEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEventsRequest) SetPageSize(v int32) *DescribeEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEventsRequest) SetRegionId(v string) *DescribeEventsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEventsRequest) SetStartTime(v string) *DescribeEventsRequest {
	s.StartTime = &v
	return s
}

type DescribeEventsResponseBody struct {
	EventItems []*DescribeEventsResponseBodyEventItems `json:"EventItems,omitempty" xml:"EventItems,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 4748127A-6D50-432C-B635-433467074C27
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	TotalRecordCount *int64 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBody) SetEventItems(v []*DescribeEventsResponseBodyEventItems) *DescribeEventsResponseBody {
	s.EventItems = v
	return s
}

func (s *DescribeEventsResponseBody) SetPageNumber(v int64) *DescribeEventsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEventsResponseBody) SetPageSize(v int64) *DescribeEventsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEventsResponseBody) SetRequestId(v string) *DescribeEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventsResponseBody) SetTotalRecordCount(v int64) *DescribeEventsResponseBody {
	s.TotalRecordCount = &v
	return s
}

type DescribeEventsResponseBodyEventItems struct {
	// example:
	//
	// 50421290
	EventId *int64 `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// example:
	//
	// ModifySecurityIps
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// example:
	//
	// {\\"Domain\\": \\"rds-cn-hangzhou.aliyuncs.com\\"}
	EventPayload *string `json:"EventPayload,omitempty" xml:"EventPayload,omitempty"`
	// example:
	//
	// FROM_USER
	EventReason *string `json:"EventReason,omitempty" xml:"EventReason,omitempty"`
	// example:
	//
	// 2021-10-15T06:39:49Z
	EventRecordTime *string `json:"EventRecordTime,omitempty" xml:"EventRecordTime,omitempty"`
	// example:
	//
	// 2021-10-15T06:35:00Z
	EventTime *string `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	// example:
	//
	// SecurityManagement
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// example:
	//
	// USRE
	EventUserType *string `json:"EventUserType,omitempty" xml:"EventUserType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// px-bp1v8udesc89g156g
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeEventsResponseBodyEventItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsResponseBodyEventItems) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBodyEventItems) SetEventId(v int64) *DescribeEventsResponseBodyEventItems {
	s.EventId = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItems) SetEventName(v string) *DescribeEventsResponseBodyEventItems {
	s.EventName = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItems) SetEventPayload(v string) *DescribeEventsResponseBodyEventItems {
	s.EventPayload = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItems) SetEventReason(v string) *DescribeEventsResponseBodyEventItems {
	s.EventReason = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItems) SetEventRecordTime(v string) *DescribeEventsResponseBodyEventItems {
	s.EventRecordTime = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItems) SetEventTime(v string) *DescribeEventsResponseBodyEventItems {
	s.EventTime = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItems) SetEventType(v string) *DescribeEventsResponseBodyEventItems {
	s.EventType = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItems) SetEventUserType(v string) *DescribeEventsResponseBodyEventItems {
	s.EventUserType = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItems) SetRegionId(v string) *DescribeEventsResponseBodyEventItems {
	s.RegionId = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItems) SetResourceName(v string) *DescribeEventsResponseBodyEventItems {
	s.ResourceName = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItems) SetResourceType(v string) *DescribeEventsResponseBodyEventItems {
	s.ResourceType = &v
	return s
}

type DescribeEventsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponse) SetHeaders(v map[string]*string) *DescribeEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventsResponse) SetStatusCode(v int32) *DescribeEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventsResponse) SetBody(v *DescribeEventsResponseBody) *DescribeEventsResponse {
	s.Body = v
	return s
}

type DescribeGdnInstancesRequest struct {
	// example:
	//
	// gdn_id、
	//
	// polarx_id
	FilterType *string `json:"FilterType,omitempty" xml:"FilterType,omitempty"`
	// example:
	//
	// gdn-***、
	//
	// pxc-***
	FilterValue *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	// GDN ID。
	//
	// example:
	//
	// gdn-***
	GDNId *string `json:"GDNId,omitempty" xml:"GDNId,omitempty"`
	// example:
	//
	// 50
	PageNum *string `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeGdnInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGdnInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGdnInstancesRequest) SetFilterType(v string) *DescribeGdnInstancesRequest {
	s.FilterType = &v
	return s
}

func (s *DescribeGdnInstancesRequest) SetFilterValue(v string) *DescribeGdnInstancesRequest {
	s.FilterValue = &v
	return s
}

func (s *DescribeGdnInstancesRequest) SetGDNId(v string) *DescribeGdnInstancesRequest {
	s.GDNId = &v
	return s
}

func (s *DescribeGdnInstancesRequest) SetPageNum(v string) *DescribeGdnInstancesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeGdnInstancesRequest) SetPageSize(v string) *DescribeGdnInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGdnInstancesRequest) SetRegionId(v string) *DescribeGdnInstancesRequest {
	s.RegionId = &v
	return s
}

type DescribeGdnInstancesResponseBody struct {
	Data *DescribeGdnInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7B044BD1-6402-5DE9-9AED-63D15A994E37
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeGdnInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGdnInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGdnInstancesResponseBody) SetData(v *DescribeGdnInstancesResponseBodyData) *DescribeGdnInstancesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeGdnInstancesResponseBody) SetMessage(v string) *DescribeGdnInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGdnInstancesResponseBody) SetRequestId(v string) *DescribeGdnInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGdnInstancesResponseBody) SetSuccess(v bool) *DescribeGdnInstancesResponseBody {
	s.Success = &v
	return s
}

type DescribeGdnInstancesResponseBodyData struct {
	GdnInstanceList []*DescribeGdnInstancesResponseBodyDataGdnInstanceList `json:"GdnInstanceList,omitempty" xml:"GdnInstanceList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 130
	TotalNumber *string `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
}

func (s DescribeGdnInstancesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeGdnInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeGdnInstancesResponseBodyData) SetGdnInstanceList(v []*DescribeGdnInstancesResponseBodyDataGdnInstanceList) *DescribeGdnInstancesResponseBodyData {
	s.GdnInstanceList = v
	return s
}

func (s *DescribeGdnInstancesResponseBodyData) SetPageNumber(v string) *DescribeGdnInstancesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyData) SetPageSize(v string) *DescribeGdnInstancesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyData) SetTotalNumber(v string) *DescribeGdnInstancesResponseBodyData {
	s.TotalNumber = &v
	return s
}

type DescribeGdnInstancesResponseBodyDataGdnInstanceList struct {
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// gdn-***
	GdnInstanceName *string `json:"GdnInstanceName,omitempty" xml:"GdnInstanceName,omitempty"`
	// example:
	//
	// 2025-01-02T13:11:10.000+0000
	GmtCreated *string                                                          `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	MemberList []*DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
	// example:
	//
	// 5.7
	MysqlVersion *string `json:"MysqlVersion,omitempty" xml:"MysqlVersion,omitempty"`
	// example:
	//
	// Creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// ""
	SwitchHistory *string `json:"SwitchHistory,omitempty" xml:"SwitchHistory,omitempty"`
}

func (s DescribeGdnInstancesResponseBodyDataGdnInstanceList) String() string {
	return tea.Prettify(s)
}

func (s DescribeGdnInstancesResponseBodyDataGdnInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) SetDescription(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceList {
	s.Description = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) SetGdnInstanceName(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceList {
	s.GdnInstanceName = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) SetGmtCreated(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceList {
	s.GmtCreated = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) SetMemberList(v []*DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) *DescribeGdnInstancesResponseBodyDataGdnInstanceList {
	s.MemberList = v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) SetMysqlVersion(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceList {
	s.MysqlVersion = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) SetStatus(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceList {
	s.Status = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) SetSwitchHistory(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceList {
	s.SwitchHistory = &v
	return s
}

type DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList struct {
	// example:
	//
	// polarx.x4.medium.2e
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	// example:
	//
	// polarx.x4.medium.2e
	CnNodeClassCode *string `json:"CnNodeClassCode,omitempty" xml:"CnNodeClassCode,omitempty"`
	// example:
	//
	// 2
	CnNodeCount *string `json:"CnNodeCount,omitempty" xml:"CnNodeCount,omitempty"`
	// example:
	//
	// drds_polarxpre_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// example:
	//
	// mysql.n4.medium.25
	DnNodeClassCode *string `json:"DnNodeClassCode,omitempty" xml:"DnNodeClassCode,omitempty"`
	// example:
	//
	// 2
	DnNodeCount *string `json:"DnNodeCount,omitempty" xml:"DnNodeCount,omitempty"`
	// example:
	//
	// 2025-01-02T13:11:10.000+0000
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 2025-01-02T13:11:10.000+0000
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// pxc-***
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// cn-zhangjiakou-a
	PrimaryZone *string `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// primary、
	//
	// standby
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// cn-zhangjiakou-a
	SecondaryZone *string `json:"SecondaryZone,omitempty" xml:"SecondaryZone,omitempty"`
	// example:
	//
	// 1s
	SecondsBehindMaster *string `json:"SecondsBehindMaster,omitempty" xml:"SecondsBehindMaster,omitempty"`
	// example:
	//
	// Creating
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// cn-zhangjiakou-a
	TertiaryZone *string `json:"TertiaryZone,omitempty" xml:"TertiaryZone,omitempty"`
	// example:
	//
	// cn-zhangjiakou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) String() string {
	return tea.Prettify(s)
}

func (s DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) GoString() string {
	return s.String()
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetClassCode(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.ClassCode = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetCnNodeClassCode(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.CnNodeClassCode = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetCnNodeCount(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.CnNodeCount = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetCommodityCode(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.CommodityCode = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetDnNodeClassCode(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.DnNodeClassCode = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetDnNodeCount(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.DnNodeCount = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetExpireTime(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.ExpireTime = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetGmtCreated(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.GmtCreated = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetMemberName(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.MemberName = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetPayType(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.PayType = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetPrimaryZone(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.PrimaryZone = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetRegionId(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.RegionId = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetRole(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.Role = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetSecondaryZone(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.SecondaryZone = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetSecondsBehindMaster(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.SecondsBehindMaster = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetStatus(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.Status = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetTaskStatus(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.TaskStatus = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetTertiaryZone(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.TertiaryZone = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetZoneId(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.ZoneId = &v
	return s
}

type DescribeGdnInstancesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGdnInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGdnInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGdnInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGdnInstancesResponse) SetHeaders(v map[string]*string) *DescribeGdnInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGdnInstancesResponse) SetStatusCode(v int32) *DescribeGdnInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGdnInstancesResponse) SetBody(v *DescribeGdnInstancesResponseBody) *DescribeGdnInstancesResponse {
	s.Body = v
	return s
}

type DescribeOpenBackupSetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 2024-10-14T00:00:00Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
}

func (s DescribeOpenBackupSetRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpenBackupSetRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpenBackupSetRequest) SetDBInstanceName(v string) *DescribeOpenBackupSetRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeOpenBackupSetRequest) SetRegionId(v string) *DescribeOpenBackupSetRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOpenBackupSetRequest) SetRestoreTime(v string) *DescribeOpenBackupSetRequest {
	s.RestoreTime = &v
	return s
}

type DescribeOpenBackupSetResponseBody struct {
	// example:
	//
	// {"gmsBackupSet": {"pubFullDownloadUrl": "https://xxxxx","dnName": "pxc-xdb-m-xxxxxx","hostInstanceId": 0001,"binlogs": [],"backupEndTime": "2024-10-21T10:11:56Z","backupLinkExpiredTime": "2024-10-23T06:13:54Z","dnBackupSetId": "00088","notCompletedBinlogs": [],"commitIndex": "15249275","innerFullDownloadUrl": "http://xxxxx","backupStartTime": "2024-10-21T10:09:20Z","backupSetSize": 526118912},"dnBackupSets": [],"insName": "pxc-xxxxx","backupSetId": "cb-xxxxx","canBinlogRecoverToTime": 1729567925000,"backupEndTime": "2024-10-21T10:12:16Z","canBinlogRecoverToTimeUTC": "2024-10-22T03:32:05Z","canBackupMinRecoverToTimeUTC": "2024-10-21T10:11:56Z","pitrInvalid": false,"backupStartTime": "2024-10-21T10:09:16Z","canBackupMinRecoverToTime": 1729505516000}
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOpenBackupSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpenBackupSetResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOpenBackupSetResponseBody) SetData(v interface{}) *DescribeOpenBackupSetResponseBody {
	s.Data = v
	return s
}

func (s *DescribeOpenBackupSetResponseBody) SetRequestId(v string) *DescribeOpenBackupSetResponseBody {
	s.RequestId = &v
	return s
}

type DescribeOpenBackupSetResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOpenBackupSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOpenBackupSetResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpenBackupSetResponse) GoString() string {
	return s.String()
}

func (s *DescribeOpenBackupSetResponse) SetHeaders(v map[string]*string) *DescribeOpenBackupSetResponse {
	s.Headers = v
	return s
}

func (s *DescribeOpenBackupSetResponse) SetStatusCode(v int32) *DescribeOpenBackupSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOpenBackupSetResponse) SetBody(v *DescribeOpenBackupSetResponseBody) *DescribeOpenBackupSetResponse {
	s.Body = v
	return s
}

type DescribeParameterTemplatesRequest struct {
	// example:
	//
	// pxc-********
	DBInstanceId  *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// example:
	//
	// compute
	ParamLevel *string `json:"ParamLevel,omitempty" xml:"ParamLevel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *DescribeParameterTemplatesRequest) SetEngineVersion(v string) *DescribeParameterTemplatesRequest {
	s.EngineVersion = &v
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
	Data *DescribeParameterTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// polarx
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 2.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// example:
	//
	// 218
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
	// example:
	//
	// [0|1]
	CheckingCode *string `json:"CheckingCode,omitempty" xml:"CheckingCode,omitempty"`
	// example:
	//
	// 0
	Dynamic *int32 `json:"Dynamic,omitempty" xml:"Dynamic,omitempty"`
	// example:
	//
	// polarx hatp addition param
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// example:
	//
	// loose_enable_gts
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// example:
	//
	// 1
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
	// example:
	//
	// 0
	Revisable *int32 `json:"Revisable,omitempty" xml:"Revisable,omitempty"`
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeParameterTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeParameterTemplatesResponse) SetStatusCode(v int32) *DescribeParameterTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeParameterTemplatesResponse) SetBody(v *DescribeParameterTemplatesResponseBody) *DescribeParameterTemplatesResponse {
	s.Body = v
	return s
}

type DescribeParametersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// compute
	ParamLevel *string `json:"ParamLevel,omitempty" xml:"ParamLevel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Data *DescribeParametersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ConfigParameters []*DescribeParametersResponseBodyDataConfigParameters `json:"ConfigParameters,omitempty" xml:"ConfigParameters,omitempty" type:"Repeated"`
	DBInstanceId     *string                                               `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// polarx
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 2.0
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

func (s *DescribeParametersResponseBodyData) SetDBInstanceId(v string) *DescribeParametersResponseBodyData {
	s.DBInstanceId = &v
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
	// example:
	//
	// CONN_POOL_XPROTO_STORAGE_DB_PORT
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// example:
	//
	// -1
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
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
	// example:
	//
	// CONN_POOL_XPROTO_STORAGE_DB_PORT
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// example:
	//
	// -1
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeParametersResponse) SetStatusCode(v int32) *DescribeParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeParametersResponse) SetBody(v *DescribeParametersResponseBody) *DescribeParametersResponse {
	s.Body = v
	return s
}

type DescribeRegionsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// success
	Message *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	Regions *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// ch-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// true
	SupportPolarx10 *bool `json:"SupportPolarx10,omitempty" xml:"SupportPolarx10,omitempty"`
	// example:
	//
	// true
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
	// example:
	//
	// true
	VpcEnabled *bool `json:"VpcEnabled,omitempty" xml:"VpcEnabled,omitempty"`
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeScaleOutMigrateTaskListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *DescribeScaleOutMigrateTaskListRequest) SetRegionId(v string) *DescribeScaleOutMigrateTaskListRequest {
	s.RegionId = &v
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
	// example:
	//
	// 32
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
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
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScaleOutMigrateTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeScaleOutMigrateTaskListResponse) SetStatusCode(v int32) *DescribeScaleOutMigrateTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScaleOutMigrateTaskListResponse) SetBody(v *DescribeScaleOutMigrateTaskListResponseBody) *DescribeScaleOutMigrateTaskListResponse {
	s.Body = v
	return s
}

type DescribeSecurityIpsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Data *DescribeSecurityIpsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 14036EBE-CF2E-44DB-ACE9-AC6157D3A6FC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// pxc-hzjasd****
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
	// example:
	//
	// defaultGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// 127.0.0.1,172.168.0.0
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecurityIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeSecurityIpsResponse) SetStatusCode(v int32) *DescribeSecurityIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecurityIpsResponse) SetBody(v *DescribeSecurityIpsResponseBody) *DescribeSecurityIpsResponse {
	s.Body = v
	return s
}

type DescribeSlowLogRecordsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// polarx_cn
	CharacterType *string `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-bjxxxxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// testdb
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// example:
	//
	// pxc-i-mezcj4ejdz
	DBNodeIds *string `json:"DBNodeIds,omitempty" xml:"DBNodeIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-11-22T02:22Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-10-09T02:26
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSlowLogRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsRequest) SetCharacterType(v string) *DescribeSlowLogRecordsRequest {
	s.CharacterType = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetDBInstanceName(v string) *DescribeSlowLogRecordsRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetDBName(v string) *DescribeSlowLogRecordsRequest {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetDBNodeIds(v string) *DescribeSlowLogRecordsRequest {
	s.DBNodeIds = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetEndTime(v string) *DescribeSlowLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetPage(v int32) *DescribeSlowLogRecordsRequest {
	s.Page = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetPageSize(v int32) *DescribeSlowLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetRegionId(v string) *DescribeSlowLogRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetStartTime(v string) *DescribeSlowLogRecordsRequest {
	s.StartTime = &v
	return s
}

type DescribeSlowLogRecordsResponseBody struct {
	// example:
	//
	// pxc-********
	DBInstanceId *string                                    `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Items        []*DescribeSlowLogRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageRecordCount *string `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBody) SetDBInstanceId(v string) *DescribeSlowLogRecordsResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetItems(v []*DescribeSlowLogRecordsResponseBodyItems) *DescribeSlowLogRecordsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetPageNumber(v string) *DescribeSlowLogRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetPageRecordCount(v string) *DescribeSlowLogRecordsResponseBody {
	s.PageRecordCount = &v
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
	// example:
	//
	// pxc-i-xxxx
	CNname *string `json:"CNname,omitempty" xml:"CNname,omitempty"`
	// example:
	//
	// dcdev
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// example:
	//
	// tddl:5.4.19-20240927
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// example:
	//
	// 0
	Fail *string `json:"Fail,omitempty" xml:"Fail,omitempty"`
	// example:
	//
	// 1
	Frows *string `json:"Frows,omitempty" xml:"Frows,omitempty"`
	// example:
	//
	// ****[****] @ [1XX.XX.XX.XX]
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// example:
	//
	// pxc-xdb-s-xxxx
	InsName *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	// example:
	//
	// 0
	IsBind *string `json:"IsBind,omitempty" xml:"IsBind,omitempty"`
	// example:
	//
	// 1
	LockTimeMS *string `json:"LockTimeMS,omitempty" xml:"LockTimeMS,omitempty"`
	// example:
	//
	// ["1"]
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// example:
	//
	// 10
	ParseRowCounts *string `json:"ParseRowCounts,omitempty" xml:"ParseRowCounts,omitempty"`
	// example:
	//
	// 2024-11-22T02:22:22.444Z
	QueryStartTime *string `json:"QueryStartTime,omitempty" xml:"QueryStartTime,omitempty"`
	// example:
	//
	// 3.000
	QueryTime *string `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
	// example:
	//
	// 3000.000
	QueryTimeMS *string `json:"QueryTimeMS,omitempty" xml:"QueryTimeMS,omitempty"`
	// example:
	//
	// 20
	ReturnRowCounts *string `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	// example:
	//
	// 1
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// example:
	//
	// 1
	SCNT *string `json:"SCNT,omitempty" xml:"SCNT,omitempty"`
	// example:
	//
	// c8df07e5d45cd68da8b4771c2016e20b
	SQLHash *string `json:"SQLHash,omitempty" xml:"SQLHash,omitempty"`
	// example:
	//
	// select 	- from test
	SQLText *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	// example:
	//
	// select
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// example:
	//
	// 0
	TooLong *string `json:"TooLong,omitempty" xml:"TooLong,omitempty"`
	// example:
	//
	// 17a5c5c840006000
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	// example:
	//
	// XA
	TransactionPolicy *string `json:"TransactionPolicy,omitempty" xml:"TransactionPolicy,omitempty"`
	// example:
	//
	// 17a5c5c840006000
	TrxId *string `json:"TrxId,omitempty" xml:"TrxId,omitempty"`
	// example:
	//
	// TP
	WT *string `json:"WT,omitempty" xml:"WT,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetCNname(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.CNname = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetDBName(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetExtension(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.Extension = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetFail(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.Fail = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetFrows(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.Frows = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetHostAddress(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.HostAddress = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetInsName(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.InsName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetIsBind(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.IsBind = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetLockTimeMS(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.LockTimeMS = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetParams(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.Params = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetParseRowCounts(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.ParseRowCounts = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetQueryStartTime(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.QueryStartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetQueryTime(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.QueryTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetQueryTimeMS(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.QueryTimeMS = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetReturnRowCounts(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetRows(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.Rows = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetSCNT(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.SCNT = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetSQLHash(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.SQLHash = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetSQLText(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.SQLText = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetSqlType(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.SqlType = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetTooLong(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.TooLong = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetTraceId(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.TraceId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetTransactionPolicy(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.TransactionPolicy = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetTrxId(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.TrxId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetWT(v string) *DescribeSlowLogRecordsResponseBodyItems {
	s.WT = &v
	return s
}

type DescribeSlowLogRecordsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlowLogRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeSlowLogRecordsResponse) SetStatusCode(v int32) *DescribeSlowLogRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlowLogRecordsResponse) SetBody(v *DescribeSlowLogRecordsResponseBody) *DescribeSlowLogRecordsResponse {
	s.Body = v
	return s
}

type DescribeTagsRequest struct {
	// example:
	//
	// pxc-xxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// test
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s DescribeTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagsRequest) SetDBInstanceName(v string) *DescribeTagsRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeTagsRequest) SetRegionId(v string) *DescribeTagsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTagsRequest) SetTagKey(v string) *DescribeTagsRequest {
	s.TagKey = &v
	return s
}

type DescribeTagsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// requestid
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagInfos  []*DescribeTagsResponseBodyTagInfos `json:"TagInfos,omitempty" xml:"TagInfos,omitempty" type:"Repeated"`
}

func (s DescribeTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBody) SetRequestId(v string) *DescribeTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagsResponseBody) SetTagInfos(v []*DescribeTagsResponseBodyTagInfos) *DescribeTagsResponseBody {
	s.TagInfos = v
	return s
}

type DescribeTagsResponseBodyTagInfos struct {
	DBInstanceIds []*string `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// testKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// testValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeTagsResponseBodyTagInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponseBodyTagInfos) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyTagInfos) SetDBInstanceIds(v []*string) *DescribeTagsResponseBodyTagInfos {
	s.DBInstanceIds = v
	return s
}

func (s *DescribeTagsResponseBodyTagInfos) SetTagKey(v string) *DescribeTagsResponseBodyTagInfos {
	s.TagKey = &v
	return s
}

func (s *DescribeTagsResponseBodyTagInfos) SetTagValue(v string) *DescribeTagsResponseBodyTagInfos {
	s.TagValue = &v
	return s
}

type DescribeTagsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponse) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponse) SetHeaders(v map[string]*string) *DescribeTagsResponse {
	s.Headers = v
	return s
}

func (s *DescribeTagsResponse) SetStatusCode(v int32) *DescribeTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTagsResponse) SetBody(v *DescribeTagsResponseBody) *DescribeTagsResponse {
	s.Body = v
	return s
}

type DescribeTasksRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-04
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 30
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-01
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
	Items []*DescribeTasksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// D6045D78-C119-5A17-9DEA-0F810394E008
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
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
	// example:
	//
	// 2021-10-20T19:40:00Z
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// DBName
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// example:
	//
	// 2021-10-20T20:00:00Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// example:
	//
	// 80
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// ProgressInfo
	ProgressInfo *string `json:"ProgressInfo,omitempty" xml:"ProgressInfo,omitempty"`
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	ScaleOutToken *string `json:"ScaleOutToken,omitempty" xml:"ScaleOutToken,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// multi_scale_out
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	// example:
	//
	// TaskErrorCode
	TaskErrorCode *string `json:"TaskErrorCode,omitempty" xml:"TaskErrorCode,omitempty"`
	// example:
	//
	// TaskErrorMessage
	TaskErrorMessage *string `json:"TaskErrorMessage,omitempty" xml:"TaskErrorMessage,omitempty"`
	// example:
	//
	// 20089398
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeTasksResponse) SetStatusCode(v int32) *DescribeTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTasksResponse) SetBody(v *DescribeTasksResponseBody) *DescribeTasksResponse {
	s.Body = v
	return s
}

type DescribeUserEncryptionKeyListRequest struct {
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Data *DescribeUserEncryptionKeyListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserEncryptionKeyListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeUserEncryptionKeyListResponse) SetStatusCode(v int32) *DescribeUserEncryptionKeyListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserEncryptionKeyListResponse) SetBody(v *DescribeUserEncryptionKeyListResponseBody) *DescribeUserEncryptionKeyListResponse {
	s.Body = v
	return s
}

type DisableRightsSeparationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-sprcym7g7w****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// account_1
	DbaAccountName *string `json:"DbaAccountName,omitempty" xml:"DbaAccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// *****
	DbaAccountPassword *string `json:"DbaAccountPassword,omitempty" xml:"DbaAccountPassword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hanghzou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// FE5D94E3-3C93-3594-95D9-AAED2A980915
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableRightsSeparationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DisableRightsSeparationResponse) SetStatusCode(v int32) *DisableRightsSeparationResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableRightsSeparationResponse) SetBody(v *DisableRightsSeparationResponseBody) *DisableRightsSeparationResponse {
	s.Body = v
	return s
}

type EnableRightsSeparationRequest struct {
	// example:
	//
	// test123
	AuditAccountDescription *string `json:"AuditAccountDescription,omitempty" xml:"AuditAccountDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// account_audit
	AuditAccountName *string `json:"AuditAccountName,omitempty" xml:"AuditAccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ******
	AuditAccountPassword *string `json:"AuditAccountPassword,omitempty" xml:"AuditAccountPassword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-htri0ori2r4k9p
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// test123
	SecurityAccountDescription *string `json:"SecurityAccountDescription,omitempty" xml:"SecurityAccountDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// account_sec
	SecurityAccountName *string `json:"SecurityAccountName,omitempty" xml:"SecurityAccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// *****
	SecurityAccountPassword *string `json:"SecurityAccountPassword,omitempty" xml:"SecurityAccountPassword,omitempty"`
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
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73559800-3c8c-11ec-bd40-99cfcff3fe1e
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableRightsSeparationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *EnableRightsSeparationResponse) SetStatusCode(v int32) *EnableRightsSeparationResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableRightsSeparationResponse) SetBody(v *EnableRightsSeparationResponseBody) *EnableRightsSeparationResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// example:
	//
	// xxdds
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// PolarDBXInstance
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
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
	// example:
	//
	// 1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 2
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
	// example:
	//
	// xxdd
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// xxxx-xxxxxx
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
	// example:
	//
	// pxc-xxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// PolarDBXInstance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// 1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ModifyAccountDescriptionRequest struct {
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// account
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// example:
	//
	// msg
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAccountDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyAccountDescriptionResponse) SetStatusCode(v int32) *ModifyAccountDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccountDescriptionResponse) SetBody(v *ModifyAccountDescriptionResponseBody) *ModifyAccountDescriptionResponse {
	s.Body = v
	return s
}

type ModifyAccountPrivilegeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// account_sec
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// ReadWrite
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-htri0ori2r4k9p
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// sbtest
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// account_audit
	SecurityAccountName *string `json:"SecurityAccountName,omitempty" xml:"SecurityAccountName,omitempty"`
	// example:
	//
	// *****
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
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73559800-3c8c-11ec-bd40-99cfcff3fe1e
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAccountPrivilegeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyAccountPrivilegeResponse) SetStatusCode(v int32) *ModifyAccountPrivilegeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccountPrivilegeResponse) SetBody(v *ModifyAccountPrivilegeResponseBody) *ModifyAccountPrivilegeResponse {
	s.Body = v
	return s
}

type ModifyActiveOperationMaintainConfRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1,2,3,4,5,6,7
	CycleTime *string `json:"CycleTime,omitempty" xml:"CycleTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Week
	CycleType *string `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 02:00:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 02:00:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyActiveOperationMaintainConfRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyActiveOperationMaintainConfRequest) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationMaintainConfRequest) SetCycleTime(v string) *ModifyActiveOperationMaintainConfRequest {
	s.CycleTime = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfRequest) SetCycleType(v string) *ModifyActiveOperationMaintainConfRequest {
	s.CycleType = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfRequest) SetMaintainEndTime(v string) *ModifyActiveOperationMaintainConfRequest {
	s.MaintainEndTime = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfRequest) SetMaintainStartTime(v string) *ModifyActiveOperationMaintainConfRequest {
	s.MaintainStartTime = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfRequest) SetRegionId(v string) *ModifyActiveOperationMaintainConfRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfRequest) SetStatus(v int32) *ModifyActiveOperationMaintainConfRequest {
	s.Status = &v
	return s
}

type ModifyActiveOperationMaintainConfResponseBody struct {
	// example:
	//
	// 4035409E-7297-4115-ABC9-C1C3942BC069
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyActiveOperationMaintainConfResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyActiveOperationMaintainConfResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationMaintainConfResponseBody) SetRequestId(v string) *ModifyActiveOperationMaintainConfResponseBody {
	s.RequestId = &v
	return s
}

type ModifyActiveOperationMaintainConfResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyActiveOperationMaintainConfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyActiveOperationMaintainConfResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyActiveOperationMaintainConfResponse) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationMaintainConfResponse) SetHeaders(v map[string]*string) *ModifyActiveOperationMaintainConfResponse {
	s.Headers = v
	return s
}

func (s *ModifyActiveOperationMaintainConfResponse) SetStatusCode(v int32) *ModifyActiveOperationMaintainConfResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfResponse) SetBody(v *ModifyActiveOperationMaintainConfResponseBody) *ModifyActiveOperationMaintainConfResponse {
	s.Body = v
	return s
}

type ModifyActiveOperationTasksRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// example:
	//
	// 1
	ImmediateStart *int64 `json:"ImmediateStart,omitempty" xml:"ImmediateStart,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 2021-08-15T12:00:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
}

func (s ModifyActiveOperationTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyActiveOperationTasksRequest) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationTasksRequest) SetIds(v string) *ModifyActiveOperationTasksRequest {
	s.Ids = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) SetImmediateStart(v int64) *ModifyActiveOperationTasksRequest {
	s.ImmediateStart = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) SetRegionId(v string) *ModifyActiveOperationTasksRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) SetSwitchTime(v string) *ModifyActiveOperationTasksRequest {
	s.SwitchTime = &v
	return s
}

type ModifyActiveOperationTasksResponseBody struct {
	// example:
	//
	// 1
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// example:
	//
	// 8C9CC90A-9532-4752-B59F-580112C5A45B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyActiveOperationTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyActiveOperationTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationTasksResponseBody) SetIds(v string) *ModifyActiveOperationTasksResponseBody {
	s.Ids = &v
	return s
}

func (s *ModifyActiveOperationTasksResponseBody) SetRequestId(v string) *ModifyActiveOperationTasksResponseBody {
	s.RequestId = &v
	return s
}

type ModifyActiveOperationTasksResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyActiveOperationTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyActiveOperationTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyActiveOperationTasksResponse) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationTasksResponse) SetHeaders(v map[string]*string) *ModifyActiveOperationTasksResponse {
	s.Headers = v
	return s
}

func (s *ModifyActiveOperationTasksResponse) SetStatusCode(v int32) *ModifyActiveOperationTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyActiveOperationTasksResponse) SetBody(v *ModifyActiveOperationTasksResponseBody) *ModifyActiveOperationTasksResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceClassRequest struct {
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CnClass     *string `json:"CnClass,omitempty" xml:"CnClass,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DnClass        *string `json:"DnClass,omitempty" xml:"DnClass,omitempty"`
	DnStorageSpace *string `json:"DnStorageSpace,omitempty" xml:"DnStorageSpace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SpecifiedDNScale       *bool   `json:"SpecifiedDNScale,omitempty" xml:"SpecifiedDNScale,omitempty"`
	SpecifiedDNSpecMapJson *string `json:"SpecifiedDNSpecMapJson,omitempty" xml:"SpecifiedDNSpecMapJson,omitempty"`
	SwitchTime             *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	SwitchTimeMode         *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
	// example:
	//
	// polarx.x4.xlarge.2e
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

func (s *ModifyDBInstanceClassRequest) SetCnClass(v string) *ModifyDBInstanceClassRequest {
	s.CnClass = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetDBInstanceName(v string) *ModifyDBInstanceClassRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetDnClass(v string) *ModifyDBInstanceClassRequest {
	s.DnClass = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetDnStorageSpace(v string) *ModifyDBInstanceClassRequest {
	s.DnStorageSpace = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetRegionId(v string) *ModifyDBInstanceClassRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetSpecifiedDNScale(v bool) *ModifyDBInstanceClassRequest {
	s.SpecifiedDNScale = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetSpecifiedDNSpecMapJson(v string) *ModifyDBInstanceClassRequest {
	s.SpecifiedDNSpecMapJson = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetSwitchTime(v string) *ModifyDBInstanceClassRequest {
	s.SwitchTime = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetSwitchTimeMode(v string) *ModifyDBInstanceClassRequest {
	s.SwitchTimeMode = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetTargetDBInstanceClass(v string) *ModifyDBInstanceClassRequest {
	s.TargetDBInstanceClass = &v
	return s
}

type ModifyDBInstanceClassResponseBody struct {
	// example:
	//
	// 20211103105558
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceClassResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyDBInstanceClassResponse) SetStatusCode(v int32) *ModifyDBInstanceClassResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceClassResponse) SetBody(v *ModifyDBInstanceClassResponseBody) *ModifyDBInstanceClassResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ENABLE_CONSISTENT_REPLICA_READ
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyDBInstanceConfigResponse) SetStatusCode(v int32) *ModifyDBInstanceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceConfigResponse) SetBody(v *ModifyDBInstanceConfigResponseBody) *ModifyDBInstanceConfigResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceConnectionStringRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-unrf5ssig0ecg8.polarx.huhehaote.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-unrf5ssig0ecg8
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3300
	NewPort *string `json:"NewPort,omitempty" xml:"NewPort,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test2
	NewPrefix *string `json:"NewPrefix,omitempty" xml:"NewPrefix,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDBInstanceConnectionStringRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConnectionStringRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionStringRequest) SetConnectionString(v string) *ModifyDBInstanceConnectionStringRequest {
	s.ConnectionString = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetDBInstanceName(v string) *ModifyDBInstanceConnectionStringRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetNewPort(v string) *ModifyDBInstanceConnectionStringRequest {
	s.NewPort = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetNewPrefix(v string) *ModifyDBInstanceConnectionStringRequest {
	s.NewPrefix = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetRegionId(v string) *ModifyDBInstanceConnectionStringRequest {
	s.RegionId = &v
	return s
}

type ModifyDBInstanceConnectionStringResponseBody struct {
	// example:
	//
	// 200
	Code *int64                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ModifyDBInstanceConnectionStringResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// c3cf535c-a585-11ea-8263-00163e04d3a7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceConnectionStringResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConnectionStringResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionStringResponseBody) SetCode(v int64) *ModifyDBInstanceConnectionStringResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponseBody) SetData(v *ModifyDBInstanceConnectionStringResponseBodyData) *ModifyDBInstanceConnectionStringResponseBody {
	s.Data = v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponseBody) SetMessage(v string) *ModifyDBInstanceConnectionStringResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponseBody) SetRequestId(v string) *ModifyDBInstanceConnectionStringResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceConnectionStringResponseBodyData struct {
	// example:
	//
	// test2.polarx.huhehaote.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// example:
	//
	// pxc-unrf5ssig0ecg8
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// 1
	DBInstanceNetType *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	// example:
	//
	// 3300
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s ModifyDBInstanceConnectionStringResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConnectionStringResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionStringResponseBodyData) SetConnectionString(v string) *ModifyDBInstanceConnectionStringResponseBodyData {
	s.ConnectionString = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponseBodyData) SetDBInstanceName(v string) *ModifyDBInstanceConnectionStringResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponseBodyData) SetDBInstanceNetType(v string) *ModifyDBInstanceConnectionStringResponseBodyData {
	s.DBInstanceNetType = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponseBodyData) SetPort(v string) *ModifyDBInstanceConnectionStringResponseBodyData {
	s.Port = &v
	return s
}

type ModifyDBInstanceConnectionStringResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceConnectionStringResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyDBInstanceConnectionStringResponse) SetStatusCode(v int32) *ModifyDBInstanceConnectionStringResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponse) SetBody(v *ModifyDBInstanceConnectionStringResponseBody) *ModifyDBInstanceConnectionStringResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceDescriptionRequest struct {
	// This parameter is required.
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyDBInstanceDescriptionResponse) SetStatusCode(v int32) *ModifyDBInstanceDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceDescriptionResponse) SetBody(v *ModifyDBInstanceDescriptionResponseBody) *ModifyDBInstanceDescriptionResponse {
	s.Body = v
	return s
}

type ModifyDatabaseDescriptionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	DbDescription *string `json:"DbDescription,omitempty" xml:"DbDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testDB
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDatabaseDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyDatabaseDescriptionResponse) SetStatusCode(v int32) *ModifyDatabaseDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDatabaseDescriptionResponse) SetBody(v *ModifyDatabaseDescriptionResponseBody) *ModifyDatabaseDescriptionResponse {
	s.Body = v
	return s
}

type ModifyParameterRequest struct {
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasdyuoo
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// compute
	ParamLevel       *string `json:"ParamLevel,omitempty" xml:"ParamLevel,omitempty"`
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// example:
	//
	// {"CONN_POOL_BLOCK_TIMEOUT":6000}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *ModifyParameterRequest) SetParameterGroupId(v string) *ModifyParameterRequest {
	s.ParameterGroupId = &v
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
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyParameterResponse) SetStatusCode(v int32) *ModifyParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyParameterResponse) SetBody(v *ModifyParameterResponseBody) *ModifyParameterResponse {
	s.Body = v
	return s
}

type ModifySecurityIpsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// default
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// 1
	ModifyMode *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 127.0.0.1,192.168.0.0
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
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySecurityIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifySecurityIpsResponse) SetStatusCode(v int32) *ModifySecurityIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySecurityIpsResponse) SetBody(v *ModifySecurityIpsResponseBody) *ModifySecurityIpsResponse {
	s.Body = v
	return s
}

type ReleaseColdDataVolumeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzravgpt8q****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ReleaseColdDataVolumeRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseColdDataVolumeRequest) GoString() string {
	return s.String()
}

func (s *ReleaseColdDataVolumeRequest) SetDBInstanceName(v string) *ReleaseColdDataVolumeRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ReleaseColdDataVolumeRequest) SetRegionId(v string) *ReleaseColdDataVolumeRequest {
	s.RegionId = &v
	return s
}

type ReleaseColdDataVolumeResponseBody struct {
	// example:
	//
	// EA330983-C895-57C0-AE82-5A63106EBB10
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseColdDataVolumeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseColdDataVolumeResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseColdDataVolumeResponseBody) SetRequestId(v string) *ReleaseColdDataVolumeResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseColdDataVolumeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseColdDataVolumeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseColdDataVolumeResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseColdDataVolumeResponse) GoString() string {
	return s.String()
}

func (s *ReleaseColdDataVolumeResponse) SetHeaders(v map[string]*string) *ReleaseColdDataVolumeResponse {
	s.Headers = v
	return s
}

func (s *ReleaseColdDataVolumeResponse) SetStatusCode(v int32) *ReleaseColdDataVolumeResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseColdDataVolumeResponse) SetBody(v *ReleaseColdDataVolumeResponseBody) *ReleaseColdDataVolumeResponse {
	s.Body = v
	return s
}

type ReleaseInstancePublicConnectionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasdyuoo.polarx.rds.aliyuncs.com
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *ReleaseInstancePublicConnectionRequest) SetRegionId(v string) *ReleaseInstancePublicConnectionRequest {
	s.RegionId = &v
	return s
}

type ReleaseInstancePublicConnectionResponseBody struct {
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
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
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseInstancePublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ReleaseInstancePublicConnectionResponse) SetStatusCode(v int32) *ReleaseInstancePublicConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseInstancePublicConnectionResponse) SetBody(v *ReleaseInstancePublicConnectionResponseBody) *ReleaseInstancePublicConnectionResponse {
	s.Body = v
	return s
}

type ResetAccountPasswordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// account
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// *****
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-htri0ori2r4k9p
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// account_sec
	SecurityAccountName *string `json:"SecurityAccountName,omitempty" xml:"SecurityAccountName,omitempty"`
	// example:
	//
	// *****
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
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73559800-3c8c-11ec-bd40-99cfcff3fe1e
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ResetAccountPasswordResponse) SetStatusCode(v int32) *ResetAccountPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetAccountPasswordResponse) SetBody(v *ResetAccountPasswordResponseBody) *ResetAccountPasswordResponse {
	s.Body = v
	return s
}

type RestartDBInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *RestartDBInstanceResponse) SetStatusCode(v int32) *RestartDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartDBInstanceResponse) SetBody(v *RestartDBInstanceResponseBody) *RestartDBInstanceResponse {
	s.Body = v
	return s
}

type SwitchDBInstanceHARequest struct {
	// This parameter is required.
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SwitchTime            *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	SwitchTimeMode        *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
	TargetPrimaryAzoneId  *string `json:"TargetPrimaryAzoneId,omitempty" xml:"TargetPrimaryAzoneId,omitempty"`
	TargetPrimaryRegionId *string `json:"TargetPrimaryRegionId,omitempty" xml:"TargetPrimaryRegionId,omitempty"`
}

func (s SwitchDBInstanceHARequest) String() string {
	return tea.Prettify(s)
}

func (s SwitchDBInstanceHARequest) GoString() string {
	return s.String()
}

func (s *SwitchDBInstanceHARequest) SetDBInstanceName(v string) *SwitchDBInstanceHARequest {
	s.DBInstanceName = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetRegionId(v string) *SwitchDBInstanceHARequest {
	s.RegionId = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetSwitchTime(v string) *SwitchDBInstanceHARequest {
	s.SwitchTime = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetSwitchTimeMode(v string) *SwitchDBInstanceHARequest {
	s.SwitchTimeMode = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetTargetPrimaryAzoneId(v string) *SwitchDBInstanceHARequest {
	s.TargetPrimaryAzoneId = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetTargetPrimaryRegionId(v string) *SwitchDBInstanceHARequest {
	s.TargetPrimaryRegionId = &v
	return s
}

type SwitchDBInstanceHAResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SwitchDBInstanceHAResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SwitchDBInstanceHAResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchDBInstanceHAResponseBody) SetMessage(v string) *SwitchDBInstanceHAResponseBody {
	s.Message = &v
	return s
}

func (s *SwitchDBInstanceHAResponseBody) SetRequestId(v string) *SwitchDBInstanceHAResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchDBInstanceHAResponseBody) SetSuccess(v bool) *SwitchDBInstanceHAResponseBody {
	s.Success = &v
	return s
}

type SwitchDBInstanceHAResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchDBInstanceHAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *SwitchDBInstanceHAResponse) SetStatusCode(v int32) *SwitchDBInstanceHAResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchDBInstanceHAResponse) SetBody(v *SwitchDBInstanceHAResponseBody) *SwitchDBInstanceHAResponse {
	s.Body = v
	return s
}

type SwitchGdnMemberRoleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	SwitchMode  *string `json:"SwitchMode,omitempty" xml:"SwitchMode,omitempty"`
	TaskTimeout *int64  `json:"TaskTimeout,omitempty" xml:"TaskTimeout,omitempty"`
}

func (s SwitchGdnMemberRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s SwitchGdnMemberRoleRequest) GoString() string {
	return s.String()
}

func (s *SwitchGdnMemberRoleRequest) SetDBInstanceName(v string) *SwitchGdnMemberRoleRequest {
	s.DBInstanceName = &v
	return s
}

func (s *SwitchGdnMemberRoleRequest) SetRegionId(v string) *SwitchGdnMemberRoleRequest {
	s.RegionId = &v
	return s
}

func (s *SwitchGdnMemberRoleRequest) SetSwitchMode(v string) *SwitchGdnMemberRoleRequest {
	s.SwitchMode = &v
	return s
}

func (s *SwitchGdnMemberRoleRequest) SetTaskTimeout(v int64) *SwitchGdnMemberRoleRequest {
	s.TaskTimeout = &v
	return s
}

type SwitchGdnMemberRoleResponseBody struct {
	Data *SwitchGdnMemberRoleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SwitchGdnMemberRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SwitchGdnMemberRoleResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchGdnMemberRoleResponseBody) SetData(v *SwitchGdnMemberRoleResponseBodyData) *SwitchGdnMemberRoleResponseBody {
	s.Data = v
	return s
}

func (s *SwitchGdnMemberRoleResponseBody) SetMessage(v string) *SwitchGdnMemberRoleResponseBody {
	s.Message = &v
	return s
}

func (s *SwitchGdnMemberRoleResponseBody) SetRequestId(v string) *SwitchGdnMemberRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchGdnMemberRoleResponseBody) SetSuccess(v bool) *SwitchGdnMemberRoleResponseBody {
	s.Success = &v
	return s
}

type SwitchGdnMemberRoleResponseBodyData struct {
	// example:
	//
	// 2209883
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SwitchGdnMemberRoleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SwitchGdnMemberRoleResponseBodyData) GoString() string {
	return s.String()
}

func (s *SwitchGdnMemberRoleResponseBodyData) SetTaskId(v int32) *SwitchGdnMemberRoleResponseBodyData {
	s.TaskId = &v
	return s
}

type SwitchGdnMemberRoleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchGdnMemberRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchGdnMemberRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s SwitchGdnMemberRoleResponse) GoString() string {
	return s.String()
}

func (s *SwitchGdnMemberRoleResponse) SetHeaders(v map[string]*string) *SwitchGdnMemberRoleResponse {
	s.Headers = v
	return s
}

func (s *SwitchGdnMemberRoleResponse) SetStatusCode(v int32) *SwitchGdnMemberRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchGdnMemberRoleResponse) SetBody(v *SwitchGdnMemberRoleResponseBody) *SwitchGdnMemberRoleResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// PolarDBXInstance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// This parameter is required.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
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
	// example:
	//
	// 12
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 22
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
	// requestId
	//
	// example:
	//
	// xxxx-xxxx
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// PolarDBXInstance
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
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

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
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
	// RequestId
	//
	// example:
	//
	// xxxxxx
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateBackupPolicyRequest struct {
	// example:
	//
	// 1001000
	BackupPeriod *string `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	// example:
	//
	// 03:00Z
	BackupPlanBegin *string `json:"BackupPlanBegin,omitempty" xml:"BackupPlanBegin,omitempty"`
	// example:
	//
	// 30
	BackupSetRetention *int32 `json:"BackupSetRetention,omitempty" xml:"BackupSetRetention,omitempty"`
	// example:
	//
	// 0
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// example:
	//
	// P
	BackupWay *string `json:"BackupWay,omitempty" xml:"BackupWay,omitempty"`
	// example:
	//
	// 30
	ColdDataBackupInterval *int32 `json:"ColdDataBackupInterval,omitempty" xml:"ColdDataBackupInterval,omitempty"`
	// example:
	//
	// 30
	ColdDataBackupRetention        *int32 `json:"ColdDataBackupRetention,omitempty" xml:"ColdDataBackupRetention,omitempty"`
	CrossRegionDataBackupRetention *int32 `json:"CrossRegionDataBackupRetention,omitempty" xml:"CrossRegionDataBackupRetention,omitempty"`
	CrossRegionLogBackupRetention  *int32 `json:"CrossRegionLogBackupRetention,omitempty" xml:"CrossRegionLogBackupRetention,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasdyuoo
	DBInstanceName  *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DestCrossRegion *string `json:"DestCrossRegion,omitempty" xml:"DestCrossRegion,omitempty"`
	// example:
	//
	// 1
	ForceCleanOnHighSpaceUsage     *int32 `json:"ForceCleanOnHighSpaceUsage,omitempty" xml:"ForceCleanOnHighSpaceUsage,omitempty"`
	IsCrossRegionDataBackupEnabled *bool  `json:"IsCrossRegionDataBackupEnabled,omitempty" xml:"IsCrossRegionDataBackupEnabled,omitempty"`
	IsCrossRegionLogBackupEnabled  *bool  `json:"IsCrossRegionLogBackupEnabled,omitempty" xml:"IsCrossRegionLogBackupEnabled,omitempty"`
	// example:
	//
	// 1
	IsEnabled *int32 `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	// example:
	//
	// 7
	LocalLogRetention       *int32 `json:"LocalLogRetention,omitempty" xml:"LocalLogRetention,omitempty"`
	LocalLogRetentionNumber *int32 `json:"LocalLogRetentionNumber,omitempty" xml:"LocalLogRetentionNumber,omitempty"`
	// example:
	//
	// 30
	LogLocalRetentionSpace *int32 `json:"LogLocalRetentionSpace,omitempty" xml:"LogLocalRetentionSpace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 7
	RemoveLogRetention *int32 `json:"RemoveLogRetention,omitempty" xml:"RemoveLogRetention,omitempty"`
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

func (s *UpdateBackupPolicyRequest) SetColdDataBackupInterval(v int32) *UpdateBackupPolicyRequest {
	s.ColdDataBackupInterval = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetColdDataBackupRetention(v int32) *UpdateBackupPolicyRequest {
	s.ColdDataBackupRetention = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetCrossRegionDataBackupRetention(v int32) *UpdateBackupPolicyRequest {
	s.CrossRegionDataBackupRetention = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetCrossRegionLogBackupRetention(v int32) *UpdateBackupPolicyRequest {
	s.CrossRegionLogBackupRetention = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetDBInstanceName(v string) *UpdateBackupPolicyRequest {
	s.DBInstanceName = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetDestCrossRegion(v string) *UpdateBackupPolicyRequest {
	s.DestCrossRegion = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetForceCleanOnHighSpaceUsage(v int32) *UpdateBackupPolicyRequest {
	s.ForceCleanOnHighSpaceUsage = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetIsCrossRegionDataBackupEnabled(v bool) *UpdateBackupPolicyRequest {
	s.IsCrossRegionDataBackupEnabled = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetIsCrossRegionLogBackupEnabled(v bool) *UpdateBackupPolicyRequest {
	s.IsCrossRegionLogBackupEnabled = &v
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

func (s *UpdateBackupPolicyRequest) SetLocalLogRetentionNumber(v int32) *UpdateBackupPolicyRequest {
	s.LocalLogRetentionNumber = &v
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
	Data *UpdateBackupPolicyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBackupPolicyResponseBody) SetData(v *UpdateBackupPolicyResponseBodyData) *UpdateBackupPolicyResponseBody {
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
	BackupPeriod                   *string `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	BackupPlanBegin                *string `json:"BackupPlanBegin,omitempty" xml:"BackupPlanBegin,omitempty"`
	BackupSetRetention             *int32  `json:"BackupSetRetention,omitempty" xml:"BackupSetRetention,omitempty"`
	BackupType                     *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	BackupWay                      *string `json:"BackupWay,omitempty" xml:"BackupWay,omitempty"`
	ColdDataBackupInterval         *int32  `json:"ColdDataBackupInterval,omitempty" xml:"ColdDataBackupInterval,omitempty"`
	ColdDataBackupRetention        *int32  `json:"ColdDataBackupRetention,omitempty" xml:"ColdDataBackupRetention,omitempty"`
	CrossRegionDataBackupRetention *int32  `json:"CrossRegionDataBackupRetention,omitempty" xml:"CrossRegionDataBackupRetention,omitempty"`
	CrossRegionLogBackupRetention  *int32  `json:"CrossRegionLogBackupRetention,omitempty" xml:"CrossRegionLogBackupRetention,omitempty"`
	DBInstanceName                 *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DestCrossRegion                *string `json:"DestCrossRegion,omitempty" xml:"DestCrossRegion,omitempty"`
	ForceCleanOnHighSpaceUsage     *int32  `json:"ForceCleanOnHighSpaceUsage,omitempty" xml:"ForceCleanOnHighSpaceUsage,omitempty"`
	IsCrossRegionDataBackupEnabled *bool   `json:"IsCrossRegionDataBackupEnabled,omitempty" xml:"IsCrossRegionDataBackupEnabled,omitempty"`
	IsCrossRegionLogBackupEnabled  *bool   `json:"IsCrossRegionLogBackupEnabled,omitempty" xml:"IsCrossRegionLogBackupEnabled,omitempty"`
	IsEnabled                      *int32  `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	LocalLogRetention              *int32  `json:"LocalLogRetention,omitempty" xml:"LocalLogRetention,omitempty"`
	LocalLogRetentionNumber        *int32  `json:"LocalLogRetentionNumber,omitempty" xml:"LocalLogRetentionNumber,omitempty"`
	LogLocalRetentionSpace         *int32  `json:"LogLocalRetentionSpace,omitempty" xml:"LogLocalRetentionSpace,omitempty"`
	RemoveLogRetention             *int32  `json:"RemoveLogRetention,omitempty" xml:"RemoveLogRetention,omitempty"`
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

func (s *UpdateBackupPolicyResponseBodyData) SetColdDataBackupInterval(v int32) *UpdateBackupPolicyResponseBodyData {
	s.ColdDataBackupInterval = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetColdDataBackupRetention(v int32) *UpdateBackupPolicyResponseBodyData {
	s.ColdDataBackupRetention = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetCrossRegionDataBackupRetention(v int32) *UpdateBackupPolicyResponseBodyData {
	s.CrossRegionDataBackupRetention = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetCrossRegionLogBackupRetention(v int32) *UpdateBackupPolicyResponseBodyData {
	s.CrossRegionLogBackupRetention = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetDBInstanceName(v string) *UpdateBackupPolicyResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetDestCrossRegion(v string) *UpdateBackupPolicyResponseBodyData {
	s.DestCrossRegion = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetForceCleanOnHighSpaceUsage(v int32) *UpdateBackupPolicyResponseBodyData {
	s.ForceCleanOnHighSpaceUsage = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetIsCrossRegionDataBackupEnabled(v bool) *UpdateBackupPolicyResponseBodyData {
	s.IsCrossRegionDataBackupEnabled = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetIsCrossRegionLogBackupEnabled(v bool) *UpdateBackupPolicyResponseBodyData {
	s.IsCrossRegionLogBackupEnabled = &v
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

func (s *UpdateBackupPolicyResponseBodyData) SetLocalLogRetentionNumber(v int32) *UpdateBackupPolicyResponseBodyData {
	s.LocalLogRetentionNumber = &v
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdateBackupPolicyResponse) SetStatusCode(v int32) *UpdateBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBackupPolicyResponse) SetBody(v *UpdateBackupPolicyResponseBody) *UpdateBackupPolicyResponse {
	s.Body = v
	return s
}

type UpdateDBInstanceSSLRequest struct {
	// example:
	//
	// pxc-hzrqjarxdocd4t.polarx.rds.aliyuncs.com
	CertCommonName *string `json:"CertCommonName,omitempty" xml:"CertCommonName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableSSL *bool `json:"EnableSSL,omitempty" xml:"EnableSSL,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Data *UpdateDBInstanceSSLResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 2209883
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDBInstanceSSLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdateDBInstanceSSLResponse) SetStatusCode(v int32) *UpdateDBInstanceSSLResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDBInstanceSSLResponse) SetBody(v *UpdateDBInstanceSSLResponseBody) *UpdateDBInstanceSSLResponse {
	s.Body = v
	return s
}

type UpdateDBInstanceTDERequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// RkVBNURDMjAtNkQ4QS01OTc5LTk3QUEtRkM1NzU0Nk******
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// acs:ram::1406926****:role/aliyunrdsinstanceencryptiondefaultrole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TDEStatus *int32 `json:"TDEStatus,omitempty" xml:"TDEStatus,omitempty"`
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
	Data *UpdateDBInstanceTDEResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 42292****
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDBInstanceTDEResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdateDBInstanceTDEResponse) SetStatusCode(v int32) *UpdateDBInstanceTDEResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDBInstanceTDEResponse) SetBody(v *UpdateDBInstanceTDEResponseBody) *UpdateDBInstanceTDEResponse {
	s.Body = v
	return s
}

type UpdatePolarDBXInstanceNodeRequest struct {
	AddDNSpec *string `json:"AddDNSpec,omitempty" xml:"AddDNSpec,omitempty"`
	// example:
	//
	// 2
	CNNodeCount *string `json:"CNNodeCount,omitempty" xml:"CNNodeCount,omitempty"`
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasdyuoo
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// 2
	DNNodeCount *string `json:"DNNodeCount,omitempty" xml:"DNNodeCount,omitempty"`
	// example:
	//
	// 3
	DbInstanceNodeCount *string `json:"DbInstanceNodeCount,omitempty" xml:"DbInstanceNodeCount,omitempty"`
	DeleteDNIds         *string `json:"DeleteDNIds,omitempty" xml:"DeleteDNIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StoragePoolName *string `json:"StoragePoolName,omitempty" xml:"StoragePoolName,omitempty"`
}

func (s UpdatePolarDBXInstanceNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePolarDBXInstanceNodeRequest) GoString() string {
	return s.String()
}

func (s *UpdatePolarDBXInstanceNodeRequest) SetAddDNSpec(v string) *UpdatePolarDBXInstanceNodeRequest {
	s.AddDNSpec = &v
	return s
}

func (s *UpdatePolarDBXInstanceNodeRequest) SetCNNodeCount(v string) *UpdatePolarDBXInstanceNodeRequest {
	s.CNNodeCount = &v
	return s
}

func (s *UpdatePolarDBXInstanceNodeRequest) SetClientToken(v string) *UpdatePolarDBXInstanceNodeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdatePolarDBXInstanceNodeRequest) SetDBInstanceName(v string) *UpdatePolarDBXInstanceNodeRequest {
	s.DBInstanceName = &v
	return s
}

func (s *UpdatePolarDBXInstanceNodeRequest) SetDNNodeCount(v string) *UpdatePolarDBXInstanceNodeRequest {
	s.DNNodeCount = &v
	return s
}

func (s *UpdatePolarDBXInstanceNodeRequest) SetDbInstanceNodeCount(v string) *UpdatePolarDBXInstanceNodeRequest {
	s.DbInstanceNodeCount = &v
	return s
}

func (s *UpdatePolarDBXInstanceNodeRequest) SetDeleteDNIds(v string) *UpdatePolarDBXInstanceNodeRequest {
	s.DeleteDNIds = &v
	return s
}

func (s *UpdatePolarDBXInstanceNodeRequest) SetRegionId(v string) *UpdatePolarDBXInstanceNodeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdatePolarDBXInstanceNodeRequest) SetStoragePoolName(v string) *UpdatePolarDBXInstanceNodeRequest {
	s.StoragePoolName = &v
	return s
}

type UpdatePolarDBXInstanceNodeResponseBody struct {
	// example:
	//
	// 20211103105558
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePolarDBXInstanceNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdatePolarDBXInstanceNodeResponse) SetStatusCode(v int32) *UpdatePolarDBXInstanceNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePolarDBXInstanceNodeResponse) SetBody(v *UpdatePolarDBXInstanceNodeResponseBody) *UpdatePolarDBXInstanceNodeResponse {
	s.Body = v
	return s
}

type UpgradeDBInstanceKernelVersionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// polarx-kernel_5.4.12-16349923_xcluster-20210926
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 0
	SwitchMode *string `json:"SwitchMode,omitempty" xml:"SwitchMode,omitempty"`
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

func (s *UpgradeDBInstanceKernelVersionRequest) SetMinorVersion(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.MinorVersion = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetRegionId(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetSwitchMode(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.SwitchMode = &v
	return s
}

type UpgradeDBInstanceKernelVersionResponseBody struct {
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// polarx-kernel_5.4.7-16001481_xcluster-20200910
	TargetMinorVersion *string `json:"TargetMinorVersion,omitempty" xml:"TargetMinorVersion,omitempty"`
	// example:
	//
	// 422922413
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeDBInstanceKernelVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpgradeDBInstanceKernelVersionResponse) SetStatusCode(v int32) *UpgradeDBInstanceKernelVersionResponse {
	s.StatusCode = &v
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

// @param request - AlignStoragePrimaryAzoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AlignStoragePrimaryAzoneResponse
func (client *Client) AlignStoragePrimaryAzoneWithOptions(request *AlignStoragePrimaryAzoneRequest, runtime *util.RuntimeOptions) (_result *AlignStoragePrimaryAzoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StorageInstanceName)) {
		query["StorageInstanceName"] = request.StorageInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.SwitchTime)) {
		query["SwitchTime"] = request.SwitchTime
	}

	if !tea.BoolValue(util.IsUnset(request.SwitchTimeMode)) {
		query["SwitchTimeMode"] = request.SwitchTimeMode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AlignStoragePrimaryAzone"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &AlignStoragePrimaryAzoneResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &AlignStoragePrimaryAzoneResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - AlignStoragePrimaryAzoneRequest
//
// @return AlignStoragePrimaryAzoneResponse
func (client *Client) AlignStoragePrimaryAzone(request *AlignStoragePrimaryAzoneRequest) (_result *AlignStoragePrimaryAzoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AlignStoragePrimaryAzoneResponse{}
	_body, _err := client.AlignStoragePrimaryAzoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开通冷存储
//
// @param request - AllocateColdDataVolumeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllocateColdDataVolumeResponse
func (client *Client) AllocateColdDataVolumeWithOptions(request *AllocateColdDataVolumeRequest, runtime *util.RuntimeOptions) (_result *AllocateColdDataVolumeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AllocateColdDataVolume"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &AllocateColdDataVolumeResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &AllocateColdDataVolumeResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 开通冷存储
//
// @param request - AllocateColdDataVolumeRequest
//
// @return AllocateColdDataVolumeResponse
func (client *Client) AllocateColdDataVolume(request *AllocateColdDataVolumeRequest) (_result *AllocateColdDataVolumeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AllocateColdDataVolumeResponse{}
	_body, _err := client.AllocateColdDataVolumeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AllocateInstancePublicConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllocateInstancePublicConnectionResponse
func (client *Client) AllocateInstancePublicConnectionWithOptions(request *AllocateInstancePublicConnectionRequest, runtime *util.RuntimeOptions) (_result *AllocateInstancePublicConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionStringPrefix)) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
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
		Action:      tea.String("AllocateInstancePublicConnection"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &AllocateInstancePublicConnectionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &AllocateInstancePublicConnectionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - AllocateInstancePublicConnectionRequest
//
// @return AllocateInstancePublicConnectionResponse
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

// Summary:
//
// 取消主动运维任务
//
// @param request - CancelActiveOperationTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelActiveOperationTasksResponse
func (client *Client) CancelActiveOperationTasksWithOptions(request *CancelActiveOperationTasksRequest, runtime *util.RuntimeOptions) (_result *CancelActiveOperationTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelActiveOperationTasks"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CancelActiveOperationTasksResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CancelActiveOperationTasksResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 取消主动运维任务
//
// @param request - CancelActiveOperationTasksRequest
//
// @return CancelActiveOperationTasksResponse
func (client *Client) CancelActiveOperationTasks(request *CancelActiveOperationTasksRequest) (_result *CancelActiveOperationTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelActiveOperationTasksResponse{}
	_body, _err := client.CancelActiveOperationTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改实例所在资源组.
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, runtime *util.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewResourceGroupId)) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeResourceGroup"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ChangeResourceGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ChangeResourceGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 修改实例所在资源组.
//
// @param request - ChangeResourceGroupRequest
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CheckCloudResourceAuthorizedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckCloudResourceAuthorizedResponse
func (client *Client) CheckCloudResourceAuthorizedWithOptions(request *CheckCloudResourceAuthorizedRequest, runtime *util.RuntimeOptions) (_result *CheckCloudResourceAuthorizedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleArn)) {
		query["RoleArn"] = request.RoleArn
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckCloudResourceAuthorized"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CheckCloudResourceAuthorizedResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CheckCloudResourceAuthorizedResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - CheckCloudResourceAuthorizedRequest
//
// @return CheckCloudResourceAuthorizedResponse
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

// @param request - CreateAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccountResponse
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

	if !tea.BoolValue(util.IsUnset(request.AccountPrivilege)) {
		query["AccountPrivilege"] = request.AccountPrivilege
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.DBName)) {
		query["DBName"] = request.DBName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityAccountName)) {
		query["SecurityAccountName"] = request.SecurityAccountName
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityAccountPassword)) {
		query["SecurityAccountPassword"] = request.SecurityAccountPassword
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccount"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateAccountResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateAccountResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - CreateAccountRequest
//
// @return CreateAccountResponse
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

// @param request - CreateBackupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBackupResponse
func (client *Client) CreateBackupWithOptions(request *CreateBackupRequest, runtime *util.RuntimeOptions) (_result *CreateBackupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupType)) {
		query["BackupType"] = request.BackupType
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBackup"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateBackupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateBackupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - CreateBackupRequest
//
// @return CreateBackupResponse
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

// @param request - CreateDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBResponse
func (client *Client) CreateDBWithOptions(request *CreateDBRequest, runtime *util.RuntimeOptions) (_result *CreateDBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AccountPrivilege)) {
		query["AccountPrivilege"] = request.AccountPrivilege
	}

	if !tea.BoolValue(util.IsUnset(request.Charset)) {
		query["Charset"] = request.Charset
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.DbDescription)) {
		query["DbDescription"] = request.DbDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityAccountName)) {
		query["SecurityAccountName"] = request.SecurityAccountName
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityAccountPassword)) {
		query["SecurityAccountPassword"] = request.SecurityAccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.StoragePoolName)) {
		query["StoragePoolName"] = request.StoragePoolName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDB"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateDBResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateDBResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - CreateDBRequest
//
// @return CreateDBResponse
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

// @param tmpReq - CreateDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBInstanceResponse
func (client *Client) CreateDBInstanceWithOptions(tmpReq *CreateDBInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateDBInstanceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateDBInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ExtraParams)) {
		request.ExtraParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExtraParams, tea.String("ExtraParams"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.CNNodeCount)) {
		query["CNNodeCount"] = request.CNNodeCount
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CnClass)) {
		query["CnClass"] = request.CnClass
	}

	if !tea.BoolValue(util.IsUnset(request.DBNodeClass)) {
		query["DBNodeClass"] = request.DBNodeClass
	}

	if !tea.BoolValue(util.IsUnset(request.DBNodeCount)) {
		query["DBNodeCount"] = request.DBNodeCount
	}

	if !tea.BoolValue(util.IsUnset(request.DNNodeCount)) {
		query["DNNodeCount"] = request.DNNodeCount
	}

	if !tea.BoolValue(util.IsUnset(request.DnClass)) {
		query["DnClass"] = request.DnClass
	}

	if !tea.BoolValue(util.IsUnset(request.DnStorageSpace)) {
		query["DnStorageSpace"] = request.DnStorageSpace
	}

	if !tea.BoolValue(util.IsUnset(request.EngineVersion)) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraParamsShrink)) {
		query["ExtraParams"] = request.ExtraParamsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.IsColumnarReadDBInstance)) {
		query["IsColumnarReadDBInstance"] = request.IsColumnarReadDBInstance
	}

	if !tea.BoolValue(util.IsUnset(request.IsReadDBInstance)) {
		query["IsReadDBInstance"] = request.IsReadDBInstance
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkType)) {
		query["NetworkType"] = request.NetworkType
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PrimaryDBInstanceName)) {
		query["PrimaryDBInstanceName"] = request.PrimaryDBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.PrimaryZone)) {
		query["PrimaryZone"] = request.PrimaryZone
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecondaryZone)) {
		query["SecondaryZone"] = request.SecondaryZone
	}

	if !tea.BoolValue(util.IsUnset(request.Series)) {
		query["Series"] = request.Series
	}

	if !tea.BoolValue(util.IsUnset(request.TertiaryZone)) {
		query["TertiaryZone"] = request.TertiaryZone
	}

	if !tea.BoolValue(util.IsUnset(request.TopologyType)) {
		query["TopologyType"] = request.TopologyType
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
		Action:      tea.String("CreateDBInstance"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateDBInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateDBInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - CreateDBInstanceRequest
//
// @return CreateDBInstanceResponse
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

// @param request - CreateSuperAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSuperAccountResponse
func (client *Client) CreateSuperAccountWithOptions(request *CreateSuperAccountRequest, runtime *util.RuntimeOptions) (_result *CreateSuperAccountResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSuperAccount"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateSuperAccountResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateSuperAccountResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - CreateSuperAccountRequest
//
// @return CreateSuperAccountResponse
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

// @param request - DeleteAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccountResponse
func (client *Client) DeleteAccountWithOptions(request *DeleteAccountRequest, runtime *util.RuntimeOptions) (_result *DeleteAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityAccountName)) {
		query["SecurityAccountName"] = request.SecurityAccountName
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityAccountPassword)) {
		query["SecurityAccountPassword"] = request.SecurityAccountPassword
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccount"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteAccountResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteAccountResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DeleteAccountRequest
//
// @return DeleteAccountResponse
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

// @param request - DeleteDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBResponse
func (client *Client) DeleteDBWithOptions(request *DeleteDBRequest, runtime *util.RuntimeOptions) (_result *DeleteDBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDB"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteDBResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteDBResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DeleteDBRequest
//
// @return DeleteDBResponse
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

// @param request - DeleteDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBInstanceResponse
func (client *Client) DeleteDBInstanceWithOptions(request *DeleteDBInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDBInstance"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteDBInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteDBInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DeleteDBInstanceRequest
//
// @return DeleteDBInstanceResponse
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

// @param request - DescribeAccountListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccountListResponse
func (client *Client) DescribeAccountListWithOptions(request *DescribeAccountListRequest, runtime *util.RuntimeOptions) (_result *DescribeAccountListResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccountList"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeAccountListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeAccountListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeAccountListRequest
//
// @return DescribeAccountListResponse
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

// Summary:
//
// 展示全局运维窗口配置
//
// @param request - DescribeActiveOperationMaintainConfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeActiveOperationMaintainConfResponse
func (client *Client) DescribeActiveOperationMaintainConfWithOptions(request *DescribeActiveOperationMaintainConfRequest, runtime *util.RuntimeOptions) (_result *DescribeActiveOperationMaintainConfResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeActiveOperationMaintainConf"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeActiveOperationMaintainConfResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeActiveOperationMaintainConfResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 展示全局运维窗口配置
//
// @param request - DescribeActiveOperationMaintainConfRequest
//
// @return DescribeActiveOperationMaintainConfResponse
func (client *Client) DescribeActiveOperationMaintainConf(request *DescribeActiveOperationMaintainConfRequest) (_result *DescribeActiveOperationMaintainConfResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeActiveOperationMaintainConfResponse{}
	_body, _err := client.DescribeActiveOperationMaintainConfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取主动运维任务数量
//
// @param request - DescribeActiveOperationTaskCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeActiveOperationTaskCountResponse
func (client *Client) DescribeActiveOperationTaskCountWithOptions(request *DescribeActiveOperationTaskCountRequest, runtime *util.RuntimeOptions) (_result *DescribeActiveOperationTaskCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeActiveOperationTaskCount"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeActiveOperationTaskCountResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeActiveOperationTaskCountResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取主动运维任务数量
//
// @param request - DescribeActiveOperationTaskCountRequest
//
// @return DescribeActiveOperationTaskCountResponse
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

// Summary:
//
// 获取待执行自动运维任务列表
//
// @param request - DescribeActiveOperationTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeActiveOperationTasksResponse
func (client *Client) DescribeActiveOperationTasksWithOptions(request *DescribeActiveOperationTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeActiveOperationTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeActiveOperationTasks"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeActiveOperationTasksResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeActiveOperationTasksResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取待执行自动运维任务列表
//
// @param request - DescribeActiveOperationTasksRequest
//
// @return DescribeActiveOperationTasksResponse
func (client *Client) DescribeActiveOperationTasks(request *DescribeActiveOperationTasksRequest) (_result *DescribeActiveOperationTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeActiveOperationTasksResponse{}
	_body, _err := client.DescribeActiveOperationTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 冷存储表列表
//
// @param request - DescribeArchiveTableListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeArchiveTableListResponse
func (client *Client) DescribeArchiveTableListWithOptions(request *DescribeArchiveTableListRequest, runtime *util.RuntimeOptions) (_result *DescribeArchiveTableListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaName)) {
		query["SchemaName"] = request.SchemaName
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeArchiveTableList"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeArchiveTableListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeArchiveTableListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 冷存储表列表
//
// @param request - DescribeArchiveTableListRequest
//
// @return DescribeArchiveTableListResponse
func (client *Client) DescribeArchiveTableList(request *DescribeArchiveTableListRequest) (_result *DescribeArchiveTableListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeArchiveTableListResponse{}
	_body, _err := client.DescribeArchiveTableListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeBackupPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupPolicyResponse
func (client *Client) DescribeBackupPolicyWithOptions(request *DescribeBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupPolicy"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeBackupPolicyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeBackupPolicyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeBackupPolicyRequest
//
// @return DescribeBackupPolicyResponse
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

// Summary:
//
// 备份集详情
//
// @param request - DescribeBackupSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupSetResponse
func (client *Client) DescribeBackupSetWithOptions(request *DescribeBackupSetRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupSetId)) {
		query["BackupSetId"] = request.BackupSetId
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.DestCrossRegion)) {
		query["DestCrossRegion"] = request.DestCrossRegion
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupSet"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeBackupSetResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeBackupSetResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 备份集详情
//
// @param request - DescribeBackupSetRequest
//
// @return DescribeBackupSetResponse
func (client *Client) DescribeBackupSet(request *DescribeBackupSetRequest) (_result *DescribeBackupSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupSetResponse{}
	_body, _err := client.DescribeBackupSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeBackupSetListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupSetListResponse
func (client *Client) DescribeBackupSetListWithOptions(request *DescribeBackupSetListRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupSetListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupSetList"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeBackupSetListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeBackupSetListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeBackupSetListRequest
//
// @return DescribeBackupSetListResponse
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

// @param request - DescribeBinaryLogListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBinaryLogListResponse
func (client *Client) DescribeBinaryLogListWithOptions(request *DescribeBinaryLogListRequest, runtime *util.RuntimeOptions) (_result *DescribeBinaryLogListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
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
		Action:      tea.String("DescribeBinaryLogList"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeBinaryLogListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeBinaryLogListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeBinaryLogListRequest
//
// @return DescribeBinaryLogListResponse
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

// @param request - DescribeCharacterSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCharacterSetResponse
func (client *Client) DescribeCharacterSetWithOptions(request *DescribeCharacterSetRequest, runtime *util.RuntimeOptions) (_result *DescribeCharacterSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCharacterSet"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeCharacterSetResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeCharacterSetResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeCharacterSetRequest
//
// @return DescribeCharacterSetResponse
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

// Summary:
//
// 冷存储基础信息
//
// @param request - DescribeColdDataBasicInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeColdDataBasicInfoResponse
func (client *Client) DescribeColdDataBasicInfoWithOptions(request *DescribeColdDataBasicInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeColdDataBasicInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeColdDataBasicInfo"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeColdDataBasicInfoResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeColdDataBasicInfoResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 冷存储基础信息
//
// @param request - DescribeColdDataBasicInfoRequest
//
// @return DescribeColdDataBasicInfoResponse
func (client *Client) DescribeColdDataBasicInfo(request *DescribeColdDataBasicInfoRequest) (_result *DescribeColdDataBasicInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeColdDataBasicInfoResponse{}
	_body, _err := client.DescribeColdDataBasicInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeDBInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceAttributeResponse
func (client *Client) DescribeDBInstanceAttributeWithOptions(request *DescribeDBInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceAttribute"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDBInstanceAttributeResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDBInstanceAttributeResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeDBInstanceAttributeRequest
//
// @return DescribeDBInstanceAttributeResponse
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

// @param request - DescribeDBInstanceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceConfigResponse
func (client *Client) DescribeDBInstanceConfigWithOptions(request *DescribeDBInstanceConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigName)) {
		query["ConfigName"] = request.ConfigName
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceConfig"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDBInstanceConfigResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDBInstanceConfigResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeDBInstanceConfigRequest
//
// @return DescribeDBInstanceConfigResponse
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

// @param request - DescribeDBInstanceHARequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceHAResponse
func (client *Client) DescribeDBInstanceHAWithOptions(request *DescribeDBInstanceHARequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceHAResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceHA"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDBInstanceHAResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDBInstanceHAResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeDBInstanceHARequest
//
// @return DescribeDBInstanceHAResponse
func (client *Client) DescribeDBInstanceHA(request *DescribeDBInstanceHARequest) (_result *DescribeDBInstanceHAResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceHAResponse{}
	_body, _err := client.DescribeDBInstanceHAWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeDBInstanceSSLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceSSLResponse
func (client *Client) DescribeDBInstanceSSLWithOptions(request *DescribeDBInstanceSSLRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceSSLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceSSL"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDBInstanceSSLResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDBInstanceSSLResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeDBInstanceSSLRequest
//
// @return DescribeDBInstanceSSLResponse
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

// @param request - DescribeDBInstanceTDERequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceTDEResponse
func (client *Client) DescribeDBInstanceTDEWithOptions(request *DescribeDBInstanceTDERequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceTDEResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceTDE"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDBInstanceTDEResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDBInstanceTDEResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeDBInstanceTDERequest
//
// @return DescribeDBInstanceTDEResponse
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

// @param request - DescribeDBInstanceTopologyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceTopologyResponse
func (client *Client) DescribeDBInstanceTopologyWithOptions(request *DescribeDBInstanceTopologyRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceTopologyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MinuteSimple)) {
		query["MinuteSimple"] = request.MinuteSimple
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
		Action:      tea.String("DescribeDBInstanceTopology"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDBInstanceTopologyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDBInstanceTopologyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeDBInstanceTopologyRequest
//
// @return DescribeDBInstanceTopologyResponse
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

// Summary:
//
// 通过Endpoint查询实例
//
// @param request - DescribeDBInstanceViaEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceViaEndpointResponse
func (client *Client) DescribeDBInstanceViaEndpointWithOptions(request *DescribeDBInstanceViaEndpointRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceViaEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Endpoint)) {
		query["Endpoint"] = request.Endpoint
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceViaEndpoint"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDBInstanceViaEndpointResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDBInstanceViaEndpointResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 通过Endpoint查询实例
//
// @param request - DescribeDBInstanceViaEndpointRequest
//
// @return DescribeDBInstanceViaEndpointResponse
func (client *Client) DescribeDBInstanceViaEndpoint(request *DescribeDBInstanceViaEndpointRequest) (_result *DescribeDBInstanceViaEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceViaEndpointResponse{}
	_body, _err := client.DescribeDBInstanceViaEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeDBInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstancesResponse
func (client *Client) DescribeDBInstancesWithOptions(request *DescribeDBInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbVersion)) {
		query["DbVersion"] = request.DbVersion
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MustHasCdc)) {
		query["MustHasCdc"] = request.MustHasCdc
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

	if !tea.BoolValue(util.IsUnset(request.Series)) {
		query["Series"] = request.Series
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstances"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDBInstancesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDBInstancesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeDBInstancesRequest
//
// @return DescribeDBInstancesResponse
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

// @param request - DescribeDBNodePerformanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBNodePerformanceResponse
func (client *Client) DescribeDBNodePerformanceWithOptions(request *DescribeDBNodePerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDBNodePerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CharacterType)) {
		query["CharacterType"] = request.CharacterType
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.DBNodeIds)) {
		query["DBNodeIds"] = request.DBNodeIds
	}

	if !tea.BoolValue(util.IsUnset(request.DBNodeRole)) {
		query["DBNodeRole"] = request.DBNodeRole
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
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
		Action:      tea.String("DescribeDBNodePerformance"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDBNodePerformanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDBNodePerformanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeDBNodePerformanceRequest
//
// @return DescribeDBNodePerformanceResponse
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

// @param request - DescribeDbListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDbListResponse
func (client *Client) DescribeDbListWithOptions(request *DescribeDbListRequest, runtime *util.RuntimeOptions) (_result *DescribeDbListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.DBName)) {
		query["DBName"] = request.DBName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDbList"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDbListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDbListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeDbListRequest
//
// @return DescribeDbListResponse
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

// @param request - DescribeDistributeTableListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDistributeTableListResponse
func (client *Client) DescribeDistributeTableListWithOptions(request *DescribeDistributeTableListRequest, runtime *util.RuntimeOptions) (_result *DescribeDistributeTableListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDistributeTableList"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDistributeTableListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDistributeTableListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeDistributeTableListRequest
//
// @return DescribeDistributeTableListResponse
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

// Summary:
//
// 历史事件
//
// @param request - DescribeEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventsResponse
func (client *Client) DescribeEventsWithOptions(request *DescribeEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEvents"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeEventsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeEventsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 历史事件
//
// @param request - DescribeEventsRequest
//
// @return DescribeEventsResponse
func (client *Client) DescribeEvents(request *DescribeEventsRequest) (_result *DescribeEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEventsResponse{}
	_body, _err := client.DescribeEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取GDN实例列表
//
// @param request - DescribeGdnInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGdnInstancesResponse
func (client *Client) DescribeGdnInstancesWithOptions(request *DescribeGdnInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeGdnInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterType)) {
		query["FilterType"] = request.FilterType
	}

	if !tea.BoolValue(util.IsUnset(request.FilterValue)) {
		query["FilterValue"] = request.FilterValue
	}

	if !tea.BoolValue(util.IsUnset(request.GDNId)) {
		query["GDNId"] = request.GDNId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGdnInstances"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeGdnInstancesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeGdnInstancesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取GDN实例列表
//
// @param request - DescribeGdnInstancesRequest
//
// @return DescribeGdnInstancesResponse
func (client *Client) DescribeGdnInstances(request *DescribeGdnInstancesRequest) (_result *DescribeGdnInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGdnInstancesResponse{}
	_body, _err := client.DescribeGdnInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开放商业备份集
//
// @param request - DescribeOpenBackupSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOpenBackupSetResponse
func (client *Client) DescribeOpenBackupSetWithOptions(request *DescribeOpenBackupSetRequest, runtime *util.RuntimeOptions) (_result *DescribeOpenBackupSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RestoreTime)) {
		query["RestoreTime"] = request.RestoreTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOpenBackupSet"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeOpenBackupSetResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeOpenBackupSetResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 开放商业备份集
//
// @param request - DescribeOpenBackupSetRequest
//
// @return DescribeOpenBackupSetResponse
func (client *Client) DescribeOpenBackupSet(request *DescribeOpenBackupSetRequest) (_result *DescribeOpenBackupSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOpenBackupSetResponse{}
	_body, _err := client.DescribeOpenBackupSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeParameterTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeParameterTemplatesResponse
func (client *Client) DescribeParameterTemplatesWithOptions(request *DescribeParameterTemplatesRequest, runtime *util.RuntimeOptions) (_result *DescribeParameterTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EngineVersion)) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ParamLevel)) {
		query["ParamLevel"] = request.ParamLevel
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeParameterTemplates"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeParameterTemplatesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeParameterTemplatesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeParameterTemplatesRequest
//
// @return DescribeParameterTemplatesResponse
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

// @param request - DescribeParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeParametersResponse
func (client *Client) DescribeParametersWithOptions(request *DescribeParametersRequest, runtime *util.RuntimeOptions) (_result *DescribeParametersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ParamLevel)) {
		query["ParamLevel"] = request.ParamLevel
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeParameters"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeParametersResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeParametersResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeParametersRequest
//
// @return DescribeParametersResponse
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

// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeRegionsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeRegionsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @return DescribeRegionsResponse
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

// @param request - DescribeScaleOutMigrateTaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScaleOutMigrateTaskListResponse
func (client *Client) DescribeScaleOutMigrateTaskListWithOptions(request *DescribeScaleOutMigrateTaskListRequest, runtime *util.RuntimeOptions) (_result *DescribeScaleOutMigrateTaskListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
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
		Action:      tea.String("DescribeScaleOutMigrateTaskList"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeScaleOutMigrateTaskListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeScaleOutMigrateTaskListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeScaleOutMigrateTaskListRequest
//
// @return DescribeScaleOutMigrateTaskListResponse
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

// @param request - DescribeSecurityIpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecurityIpsResponse
func (client *Client) DescribeSecurityIpsWithOptions(request *DescribeSecurityIpsRequest, runtime *util.RuntimeOptions) (_result *DescribeSecurityIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSecurityIps"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeSecurityIpsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeSecurityIpsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeSecurityIpsRequest
//
// @return DescribeSecurityIpsResponse
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

// Summary:
//
// 慢SQL明细
//
// @param request - DescribeSlowLogRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlowLogRecordsResponse
func (client *Client) DescribeSlowLogRecordsWithOptions(request *DescribeSlowLogRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeSlowLogRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CharacterType)) {
		query["CharacterType"] = request.CharacterType
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.DBName)) {
		query["DBName"] = request.DBName
	}

	if !tea.BoolValue(util.IsUnset(request.DBNodeIds)) {
		query["DBNodeIds"] = request.DBNodeIds
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
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
		Action:      tea.String("DescribeSlowLogRecords"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeSlowLogRecordsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeSlowLogRecordsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 慢SQL明细
//
// @param request - DescribeSlowLogRecordsRequest
//
// @return DescribeSlowLogRecordsResponse
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

// Summary:
//
// 标签列表查询
//
// @param request - DescribeTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTagsResponse
func (client *Client) DescribeTagsWithOptions(request *DescribeTagsRequest, runtime *util.RuntimeOptions) (_result *DescribeTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTags"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeTagsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeTagsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 标签列表查询
//
// @param request - DescribeTagsRequest
//
// @return DescribeTagsResponse
func (client *Client) DescribeTags(request *DescribeTagsRequest) (_result *DescribeTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTagsResponse{}
	_body, _err := client.DescribeTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTasksResponse
func (client *Client) DescribeTasksWithOptions(request *DescribeTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
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
		Action:      tea.String("DescribeTasks"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeTasksResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeTasksResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeTasksRequest
//
// @return DescribeTasksResponse
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

// @param request - DescribeUserEncryptionKeyListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserEncryptionKeyListResponse
func (client *Client) DescribeUserEncryptionKeyListWithOptions(request *DescribeUserEncryptionKeyListRequest, runtime *util.RuntimeOptions) (_result *DescribeUserEncryptionKeyListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserEncryptionKeyList"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeUserEncryptionKeyListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeUserEncryptionKeyListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeUserEncryptionKeyListRequest
//
// @return DescribeUserEncryptionKeyListResponse
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

// @param request - DisableRightsSeparationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableRightsSeparationResponse
func (client *Client) DisableRightsSeparationWithOptions(request *DisableRightsSeparationRequest, runtime *util.RuntimeOptions) (_result *DisableRightsSeparationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.DbaAccountName)) {
		query["DbaAccountName"] = request.DbaAccountName
	}

	if !tea.BoolValue(util.IsUnset(request.DbaAccountPassword)) {
		query["DbaAccountPassword"] = request.DbaAccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableRightsSeparation"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DisableRightsSeparationResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DisableRightsSeparationResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DisableRightsSeparationRequest
//
// @return DisableRightsSeparationResponse
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

// Summary:
//
// 开启三权分立
//
// @param request - EnableRightsSeparationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableRightsSeparationResponse
func (client *Client) EnableRightsSeparationWithOptions(request *EnableRightsSeparationRequest, runtime *util.RuntimeOptions) (_result *EnableRightsSeparationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuditAccountDescription)) {
		query["AuditAccountDescription"] = request.AuditAccountDescription
	}

	if !tea.BoolValue(util.IsUnset(request.AuditAccountName)) {
		query["AuditAccountName"] = request.AuditAccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AuditAccountPassword)) {
		query["AuditAccountPassword"] = request.AuditAccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityAccountDescription)) {
		query["SecurityAccountDescription"] = request.SecurityAccountDescription
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityAccountName)) {
		query["SecurityAccountName"] = request.SecurityAccountName
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityAccountPassword)) {
		query["SecurityAccountPassword"] = request.SecurityAccountPassword
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableRightsSeparation"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &EnableRightsSeparationResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &EnableRightsSeparationResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 开启三权分立
//
// @param request - EnableRightsSeparationRequest
//
// @return EnableRightsSeparationResponse
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

// Summary:
//
// 查标签接口
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
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
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListTagResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListTagResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查标签接口
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
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

// @param request - ModifyAccountDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccountDescriptionResponse
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

	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAccountDescription"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyAccountDescriptionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyAccountDescriptionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ModifyAccountDescriptionRequest
//
// @return ModifyAccountDescriptionResponse
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

// @param request - ModifyAccountPrivilegeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccountPrivilegeResponse
func (client *Client) ModifyAccountPrivilegeWithOptions(request *ModifyAccountPrivilegeRequest, runtime *util.RuntimeOptions) (_result *ModifyAccountPrivilegeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AccountPrivilege)) {
		query["AccountPrivilege"] = request.AccountPrivilege
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityAccountName)) {
		query["SecurityAccountName"] = request.SecurityAccountName
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityAccountPassword)) {
		query["SecurityAccountPassword"] = request.SecurityAccountPassword
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAccountPrivilege"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyAccountPrivilegeResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyAccountPrivilegeResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ModifyAccountPrivilegeRequest
//
// @return ModifyAccountPrivilegeResponse
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

// Summary:
//
// 修改全局运维窗口信息
//
// @param request - ModifyActiveOperationMaintainConfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyActiveOperationMaintainConfResponse
func (client *Client) ModifyActiveOperationMaintainConfWithOptions(request *ModifyActiveOperationMaintainConfRequest, runtime *util.RuntimeOptions) (_result *ModifyActiveOperationMaintainConfResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyActiveOperationMaintainConf"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyActiveOperationMaintainConfResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyActiveOperationMaintainConfResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 修改全局运维窗口信息
//
// @param request - ModifyActiveOperationMaintainConfRequest
//
// @return ModifyActiveOperationMaintainConfResponse
func (client *Client) ModifyActiveOperationMaintainConf(request *ModifyActiveOperationMaintainConfRequest) (_result *ModifyActiveOperationMaintainConfResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyActiveOperationMaintainConfResponse{}
	_body, _err := client.ModifyActiveOperationMaintainConfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改主动运维任务
//
// @param request - ModifyActiveOperationTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyActiveOperationTasksResponse
func (client *Client) ModifyActiveOperationTasksWithOptions(request *ModifyActiveOperationTasksRequest, runtime *util.RuntimeOptions) (_result *ModifyActiveOperationTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ids)) {
		query["Ids"] = request.Ids
	}

	if !tea.BoolValue(util.IsUnset(request.ImmediateStart)) {
		query["ImmediateStart"] = request.ImmediateStart
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SwitchTime)) {
		query["SwitchTime"] = request.SwitchTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyActiveOperationTasks"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyActiveOperationTasksResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyActiveOperationTasksResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 修改主动运维任务
//
// @param request - ModifyActiveOperationTasksRequest
//
// @return ModifyActiveOperationTasksResponse
func (client *Client) ModifyActiveOperationTasks(request *ModifyActiveOperationTasksRequest) (_result *ModifyActiveOperationTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyActiveOperationTasksResponse{}
	_body, _err := client.ModifyActiveOperationTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyDBInstanceClassRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceClassResponse
func (client *Client) ModifyDBInstanceClassWithOptions(request *ModifyDBInstanceClassRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceClassResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CnClass)) {
		query["CnClass"] = request.CnClass
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.DnClass)) {
		query["DnClass"] = request.DnClass
	}

	if !tea.BoolValue(util.IsUnset(request.DnStorageSpace)) {
		query["DnStorageSpace"] = request.DnStorageSpace
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SpecifiedDNScale)) {
		query["SpecifiedDNScale"] = request.SpecifiedDNScale
	}

	if !tea.BoolValue(util.IsUnset(request.SpecifiedDNSpecMapJson)) {
		query["SpecifiedDNSpecMapJson"] = request.SpecifiedDNSpecMapJson
	}

	if !tea.BoolValue(util.IsUnset(request.SwitchTime)) {
		query["SwitchTime"] = request.SwitchTime
	}

	if !tea.BoolValue(util.IsUnset(request.SwitchTimeMode)) {
		query["SwitchTimeMode"] = request.SwitchTimeMode
	}

	if !tea.BoolValue(util.IsUnset(request.TargetDBInstanceClass)) {
		query["TargetDBInstanceClass"] = request.TargetDBInstanceClass
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBInstanceClass"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyDBInstanceClassResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyDBInstanceClassResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ModifyDBInstanceClassRequest
//
// @return ModifyDBInstanceClassResponse
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

// Summary:
//
// 修改实例配置
//
// @param request - ModifyDBInstanceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceConfigResponse
func (client *Client) ModifyDBInstanceConfigWithOptions(request *ModifyDBInstanceConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigName)) {
		query["ConfigName"] = request.ConfigName
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigValue)) {
		query["ConfigValue"] = request.ConfigValue
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBInstanceConfig"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyDBInstanceConfigResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyDBInstanceConfigResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 修改实例配置
//
// @param request - ModifyDBInstanceConfigRequest
//
// @return ModifyDBInstanceConfigResponse
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

// Summary:
//
// 修改实例链接串
//
// @param request - ModifyDBInstanceConnectionStringRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceConnectionStringResponse
func (client *Client) ModifyDBInstanceConnectionStringWithOptions(request *ModifyDBInstanceConnectionStringRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceConnectionStringResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionString)) {
		query["ConnectionString"] = request.ConnectionString
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.NewPort)) {
		query["NewPort"] = request.NewPort
	}

	if !tea.BoolValue(util.IsUnset(request.NewPrefix)) {
		query["NewPrefix"] = request.NewPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBInstanceConnectionString"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyDBInstanceConnectionStringResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyDBInstanceConnectionStringResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 修改实例链接串
//
// @param request - ModifyDBInstanceConnectionStringRequest
//
// @return ModifyDBInstanceConnectionStringResponse
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

// @param request - ModifyDBInstanceDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceDescriptionResponse
func (client *Client) ModifyDBInstanceDescriptionWithOptions(request *ModifyDBInstanceDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceDescription)) {
		query["DBInstanceDescription"] = request.DBInstanceDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBInstanceDescription"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyDBInstanceDescriptionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyDBInstanceDescriptionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ModifyDBInstanceDescriptionRequest
//
// @return ModifyDBInstanceDescriptionResponse
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

// @param request - ModifyDatabaseDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDatabaseDescriptionResponse
func (client *Client) ModifyDatabaseDescriptionWithOptions(request *ModifyDatabaseDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyDatabaseDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.DbDescription)) {
		query["DbDescription"] = request.DbDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDatabaseDescription"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyDatabaseDescriptionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyDatabaseDescriptionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ModifyDatabaseDescriptionRequest
//
// @return ModifyDatabaseDescriptionResponse
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

// @param request - ModifyParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyParameterResponse
func (client *Client) ModifyParameterWithOptions(request *ModifyParameterRequest, runtime *util.RuntimeOptions) (_result *ModifyParameterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ParamLevel)) {
		query["ParamLevel"] = request.ParamLevel
	}

	if !tea.BoolValue(util.IsUnset(request.ParameterGroupId)) {
		query["ParameterGroupId"] = request.ParameterGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyParameter"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyParameterResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyParameterResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ModifyParameterRequest
//
// @return ModifyParameterResponse
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

// @param request - ModifySecurityIpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySecurityIpsResponse
func (client *Client) ModifySecurityIpsWithOptions(request *ModifySecurityIpsRequest, runtime *util.RuntimeOptions) (_result *ModifySecurityIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyMode)) {
		query["ModifyMode"] = request.ModifyMode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityIPList)) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySecurityIps"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifySecurityIpsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifySecurityIpsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ModifySecurityIpsRequest
//
// @return ModifySecurityIpsResponse
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

// Summary:
//
// 关闭冷存储
//
// @param request - ReleaseColdDataVolumeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseColdDataVolumeResponse
func (client *Client) ReleaseColdDataVolumeWithOptions(request *ReleaseColdDataVolumeRequest, runtime *util.RuntimeOptions) (_result *ReleaseColdDataVolumeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseColdDataVolume"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ReleaseColdDataVolumeResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ReleaseColdDataVolumeResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 关闭冷存储
//
// @param request - ReleaseColdDataVolumeRequest
//
// @return ReleaseColdDataVolumeResponse
func (client *Client) ReleaseColdDataVolume(request *ReleaseColdDataVolumeRequest) (_result *ReleaseColdDataVolumeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseColdDataVolumeResponse{}
	_body, _err := client.ReleaseColdDataVolumeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ReleaseInstancePublicConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseInstancePublicConnectionResponse
func (client *Client) ReleaseInstancePublicConnectionWithOptions(request *ReleaseInstancePublicConnectionRequest, runtime *util.RuntimeOptions) (_result *ReleaseInstancePublicConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentConnectionString)) {
		query["CurrentConnectionString"] = request.CurrentConnectionString
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseInstancePublicConnection"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ReleaseInstancePublicConnectionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ReleaseInstancePublicConnectionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ReleaseInstancePublicConnectionRequest
//
// @return ReleaseInstancePublicConnectionResponse
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

// @param request - ResetAccountPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetAccountPasswordResponse
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

	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityAccountName)) {
		query["SecurityAccountName"] = request.SecurityAccountName
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityAccountPassword)) {
		query["SecurityAccountPassword"] = request.SecurityAccountPassword
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetAccountPassword"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ResetAccountPasswordResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ResetAccountPasswordResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ResetAccountPasswordRequest
//
// @return ResetAccountPasswordResponse
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

// @param request - RestartDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartDBInstanceResponse
func (client *Client) RestartDBInstanceWithOptions(request *RestartDBInstanceRequest, runtime *util.RuntimeOptions) (_result *RestartDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartDBInstance"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RestartDBInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RestartDBInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - RestartDBInstanceRequest
//
// @return RestartDBInstanceResponse
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

// @param request - SwitchDBInstanceHARequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchDBInstanceHAResponse
func (client *Client) SwitchDBInstanceHAWithOptions(request *SwitchDBInstanceHARequest, runtime *util.RuntimeOptions) (_result *SwitchDBInstanceHAResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SwitchTime)) {
		query["SwitchTime"] = request.SwitchTime
	}

	if !tea.BoolValue(util.IsUnset(request.SwitchTimeMode)) {
		query["SwitchTimeMode"] = request.SwitchTimeMode
	}

	if !tea.BoolValue(util.IsUnset(request.TargetPrimaryAzoneId)) {
		query["TargetPrimaryAzoneId"] = request.TargetPrimaryAzoneId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetPrimaryRegionId)) {
		query["TargetPrimaryRegionId"] = request.TargetPrimaryRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SwitchDBInstanceHA"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SwitchDBInstanceHAResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SwitchDBInstanceHAResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - SwitchDBInstanceHARequest
//
// @return SwitchDBInstanceHAResponse
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

// Summary:
//
// GDN主备切换
//
// @param request - SwitchGdnMemberRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchGdnMemberRoleResponse
func (client *Client) SwitchGdnMemberRoleWithOptions(request *SwitchGdnMemberRoleRequest, runtime *util.RuntimeOptions) (_result *SwitchGdnMemberRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SwitchMode)) {
		query["SwitchMode"] = request.SwitchMode
	}

	if !tea.BoolValue(util.IsUnset(request.TaskTimeout)) {
		query["TaskTimeout"] = request.TaskTimeout
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SwitchGdnMemberRole"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SwitchGdnMemberRoleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SwitchGdnMemberRoleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// GDN主备切换
//
// @param request - SwitchGdnMemberRoleRequest
//
// @return SwitchGdnMemberRoleResponse
func (client *Client) SwitchGdnMemberRole(request *SwitchGdnMemberRoleRequest) (_result *SwitchGdnMemberRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SwitchGdnMemberRoleResponse{}
	_body, _err := client.SwitchGdnMemberRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 打标签接口
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
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
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &TagResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &TagResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 打标签接口
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
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

// Summary:
//
// 删标签接口
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
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
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UntagResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UntagResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删标签接口
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
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

// @param request - UpdateBackupPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBackupPolicyResponse
func (client *Client) UpdateBackupPolicyWithOptions(request *UpdateBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *UpdateBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupPeriod)) {
		query["BackupPeriod"] = request.BackupPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.BackupPlanBegin)) {
		query["BackupPlanBegin"] = request.BackupPlanBegin
	}

	if !tea.BoolValue(util.IsUnset(request.BackupSetRetention)) {
		query["BackupSetRetention"] = request.BackupSetRetention
	}

	if !tea.BoolValue(util.IsUnset(request.BackupType)) {
		query["BackupType"] = request.BackupType
	}

	if !tea.BoolValue(util.IsUnset(request.BackupWay)) {
		query["BackupWay"] = request.BackupWay
	}

	if !tea.BoolValue(util.IsUnset(request.ColdDataBackupInterval)) {
		query["ColdDataBackupInterval"] = request.ColdDataBackupInterval
	}

	if !tea.BoolValue(util.IsUnset(request.ColdDataBackupRetention)) {
		query["ColdDataBackupRetention"] = request.ColdDataBackupRetention
	}

	if !tea.BoolValue(util.IsUnset(request.CrossRegionDataBackupRetention)) {
		query["CrossRegionDataBackupRetention"] = request.CrossRegionDataBackupRetention
	}

	if !tea.BoolValue(util.IsUnset(request.CrossRegionLogBackupRetention)) {
		query["CrossRegionLogBackupRetention"] = request.CrossRegionLogBackupRetention
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.DestCrossRegion)) {
		query["DestCrossRegion"] = request.DestCrossRegion
	}

	if !tea.BoolValue(util.IsUnset(request.ForceCleanOnHighSpaceUsage)) {
		query["ForceCleanOnHighSpaceUsage"] = request.ForceCleanOnHighSpaceUsage
	}

	if !tea.BoolValue(util.IsUnset(request.IsCrossRegionDataBackupEnabled)) {
		query["IsCrossRegionDataBackupEnabled"] = request.IsCrossRegionDataBackupEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.IsCrossRegionLogBackupEnabled)) {
		query["IsCrossRegionLogBackupEnabled"] = request.IsCrossRegionLogBackupEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.IsEnabled)) {
		query["IsEnabled"] = request.IsEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.LocalLogRetention)) {
		query["LocalLogRetention"] = request.LocalLogRetention
	}

	if !tea.BoolValue(util.IsUnset(request.LocalLogRetentionNumber)) {
		query["LocalLogRetentionNumber"] = request.LocalLogRetentionNumber
	}

	if !tea.BoolValue(util.IsUnset(request.LogLocalRetentionSpace)) {
		query["LogLocalRetentionSpace"] = request.LogLocalRetentionSpace
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RemoveLogRetention)) {
		query["RemoveLogRetention"] = request.RemoveLogRetention
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateBackupPolicy"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateBackupPolicyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateBackupPolicyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - UpdateBackupPolicyRequest
//
// @return UpdateBackupPolicyResponse
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

// @param request - UpdateDBInstanceSSLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDBInstanceSSLResponse
func (client *Client) UpdateDBInstanceSSLWithOptions(request *UpdateDBInstanceSSLRequest, runtime *util.RuntimeOptions) (_result *UpdateDBInstanceSSLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertCommonName)) {
		query["CertCommonName"] = request.CertCommonName
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.EnableSSL)) {
		query["EnableSSL"] = request.EnableSSL
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDBInstanceSSL"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateDBInstanceSSLResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateDBInstanceSSLResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - UpdateDBInstanceSSLRequest
//
// @return UpdateDBInstanceSSLResponse
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

// @param request - UpdateDBInstanceTDERequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDBInstanceTDEResponse
func (client *Client) UpdateDBInstanceTDEWithOptions(request *UpdateDBInstanceTDERequest, runtime *util.RuntimeOptions) (_result *UpdateDBInstanceTDEResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptionKey)) {
		query["EncryptionKey"] = request.EncryptionKey
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleArn)) {
		query["RoleArn"] = request.RoleArn
	}

	if !tea.BoolValue(util.IsUnset(request.TDEStatus)) {
		query["TDEStatus"] = request.TDEStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDBInstanceTDE"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateDBInstanceTDEResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateDBInstanceTDEResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - UpdateDBInstanceTDERequest
//
// @return UpdateDBInstanceTDEResponse
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

// @param request - UpdatePolarDBXInstanceNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePolarDBXInstanceNodeResponse
func (client *Client) UpdatePolarDBXInstanceNodeWithOptions(request *UpdatePolarDBXInstanceNodeRequest, runtime *util.RuntimeOptions) (_result *UpdatePolarDBXInstanceNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddDNSpec)) {
		query["AddDNSpec"] = request.AddDNSpec
	}

	if !tea.BoolValue(util.IsUnset(request.CNNodeCount)) {
		query["CNNodeCount"] = request.CNNodeCount
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.DNNodeCount)) {
		query["DNNodeCount"] = request.DNNodeCount
	}

	if !tea.BoolValue(util.IsUnset(request.DbInstanceNodeCount)) {
		query["DbInstanceNodeCount"] = request.DbInstanceNodeCount
	}

	if !tea.BoolValue(util.IsUnset(request.DeleteDNIds)) {
		query["DeleteDNIds"] = request.DeleteDNIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StoragePoolName)) {
		query["StoragePoolName"] = request.StoragePoolName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePolarDBXInstanceNode"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdatePolarDBXInstanceNodeResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdatePolarDBXInstanceNodeResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - UpdatePolarDBXInstanceNodeRequest
//
// @return UpdatePolarDBXInstanceNodeResponse
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

// @param request - UpgradeDBInstanceKernelVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeDBInstanceKernelVersionResponse
func (client *Client) UpgradeDBInstanceKernelVersionWithOptions(request *UpgradeDBInstanceKernelVersionRequest, runtime *util.RuntimeOptions) (_result *UpgradeDBInstanceKernelVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceName)) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.MinorVersion)) {
		query["MinorVersion"] = request.MinorVersion
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SwitchMode)) {
		query["SwitchMode"] = request.SwitchMode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeDBInstanceKernelVersion"),
		Version:     tea.String("2020-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpgradeDBInstanceKernelVersionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpgradeDBInstanceKernelVersionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - UpgradeDBInstanceKernelVersionRequest
//
// @return UpgradeDBInstanceKernelVersionResponse
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
