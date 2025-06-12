// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddUserToGroupRequest struct {
	// The name of the RAM user group.
	//
	// example:
	//
	// Dev-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The name of the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s AddUserToGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUserToGroupRequest) GoString() string {
	return s.String()
}

func (s *AddUserToGroupRequest) SetGroupName(v string) *AddUserToGroupRequest {
	s.GroupName = &v
	return s
}

func (s *AddUserToGroupRequest) SetUserName(v string) *AddUserToGroupRequest {
	s.UserName = &v
	return s
}

type AddUserToGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1B968853-B423-63A6-FE1F-45E81BC2AD61
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUserToGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUserToGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserToGroupResponseBody) SetRequestId(v string) *AddUserToGroupResponseBody {
	s.RequestId = &v
	return s
}

type AddUserToGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUserToGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUserToGroupResponse) GoString() string {
	return s.String()
}

func (s *AddUserToGroupResponse) SetHeaders(v map[string]*string) *AddUserToGroupResponse {
	s.Headers = v
	return s
}

func (s *AddUserToGroupResponse) SetStatusCode(v int32) *AddUserToGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserToGroupResponse) SetBody(v *AddUserToGroupResponseBody) *AddUserToGroupResponse {
	s.Body = v
	return s
}

type AttachPolicyToGroupRequest struct {
	// The name of the RAM user group.
	//
	// example:
	//
	// dev
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy. Valid values: `System` and `Custom`.
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s AttachPolicyToGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachPolicyToGroupRequest) GoString() string {
	return s.String()
}

func (s *AttachPolicyToGroupRequest) SetGroupName(v string) *AttachPolicyToGroupRequest {
	s.GroupName = &v
	return s
}

func (s *AttachPolicyToGroupRequest) SetPolicyName(v string) *AttachPolicyToGroupRequest {
	s.PolicyName = &v
	return s
}

func (s *AttachPolicyToGroupRequest) SetPolicyType(v string) *AttachPolicyToGroupRequest {
	s.PolicyType = &v
	return s
}

type AttachPolicyToGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 697852FB-50D7-44D9-9774-530C31EAC572
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachPolicyToGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachPolicyToGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AttachPolicyToGroupResponseBody) SetRequestId(v string) *AttachPolicyToGroupResponseBody {
	s.RequestId = &v
	return s
}

type AttachPolicyToGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachPolicyToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachPolicyToGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachPolicyToGroupResponse) GoString() string {
	return s.String()
}

func (s *AttachPolicyToGroupResponse) SetHeaders(v map[string]*string) *AttachPolicyToGroupResponse {
	s.Headers = v
	return s
}

func (s *AttachPolicyToGroupResponse) SetStatusCode(v int32) *AttachPolicyToGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachPolicyToGroupResponse) SetBody(v *AttachPolicyToGroupResponseBody) *AttachPolicyToGroupResponse {
	s.Body = v
	return s
}

type AttachPolicyToRoleRequest struct {
	// The name of the policy.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy. Valid values: `System` and `Custom`.
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The name of the RAM role.
	//
	// example:
	//
	// OSSAdminRole
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s AttachPolicyToRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachPolicyToRoleRequest) GoString() string {
	return s.String()
}

func (s *AttachPolicyToRoleRequest) SetPolicyName(v string) *AttachPolicyToRoleRequest {
	s.PolicyName = &v
	return s
}

func (s *AttachPolicyToRoleRequest) SetPolicyType(v string) *AttachPolicyToRoleRequest {
	s.PolicyType = &v
	return s
}

func (s *AttachPolicyToRoleRequest) SetRoleName(v string) *AttachPolicyToRoleRequest {
	s.RoleName = &v
	return s
}

type AttachPolicyToRoleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 697852FB-50D7-44D9-9774-530C31EAC572
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachPolicyToRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachPolicyToRoleResponseBody) GoString() string {
	return s.String()
}

func (s *AttachPolicyToRoleResponseBody) SetRequestId(v string) *AttachPolicyToRoleResponseBody {
	s.RequestId = &v
	return s
}

type AttachPolicyToRoleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachPolicyToRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachPolicyToRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachPolicyToRoleResponse) GoString() string {
	return s.String()
}

func (s *AttachPolicyToRoleResponse) SetHeaders(v map[string]*string) *AttachPolicyToRoleResponse {
	s.Headers = v
	return s
}

func (s *AttachPolicyToRoleResponse) SetStatusCode(v int32) *AttachPolicyToRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachPolicyToRoleResponse) SetBody(v *AttachPolicyToRoleResponseBody) *AttachPolicyToRoleResponse {
	s.Body = v
	return s
}

type AttachPolicyToUserRequest struct {
	// The name of the policy.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy. Valid values: `System` and `Custom`.
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The name of the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s AttachPolicyToUserRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachPolicyToUserRequest) GoString() string {
	return s.String()
}

func (s *AttachPolicyToUserRequest) SetPolicyName(v string) *AttachPolicyToUserRequest {
	s.PolicyName = &v
	return s
}

func (s *AttachPolicyToUserRequest) SetPolicyType(v string) *AttachPolicyToUserRequest {
	s.PolicyType = &v
	return s
}

func (s *AttachPolicyToUserRequest) SetUserName(v string) *AttachPolicyToUserRequest {
	s.UserName = &v
	return s
}

type AttachPolicyToUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 697852FB-50D7-44D9-9774-530C31EAC572
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachPolicyToUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachPolicyToUserResponseBody) GoString() string {
	return s.String()
}

func (s *AttachPolicyToUserResponseBody) SetRequestId(v string) *AttachPolicyToUserResponseBody {
	s.RequestId = &v
	return s
}

type AttachPolicyToUserResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachPolicyToUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachPolicyToUserResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachPolicyToUserResponse) GoString() string {
	return s.String()
}

func (s *AttachPolicyToUserResponse) SetHeaders(v map[string]*string) *AttachPolicyToUserResponse {
	s.Headers = v
	return s
}

func (s *AttachPolicyToUserResponse) SetStatusCode(v int32) *AttachPolicyToUserResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachPolicyToUserResponse) SetBody(v *AttachPolicyToUserResponseBody) *AttachPolicyToUserResponse {
	s.Body = v
	return s
}

type BindMFADeviceRequest struct {
	// The first authentication code.
	//
	// example:
	//
	// 11****
	AuthenticationCode1 *string `json:"AuthenticationCode1,omitempty" xml:"AuthenticationCode1,omitempty"`
	// The second authentication code.
	//
	// example:
	//
	// 33****
	AuthenticationCode2 *string `json:"AuthenticationCode2,omitempty" xml:"AuthenticationCode2,omitempty"`
	// The serial number of the MFA device.
	//
	// example:
	//
	// acs:ram::123456789012****:mfa/device002
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The name of the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s BindMFADeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s BindMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *BindMFADeviceRequest) SetAuthenticationCode1(v string) *BindMFADeviceRequest {
	s.AuthenticationCode1 = &v
	return s
}

func (s *BindMFADeviceRequest) SetAuthenticationCode2(v string) *BindMFADeviceRequest {
	s.AuthenticationCode2 = &v
	return s
}

func (s *BindMFADeviceRequest) SetSerialNumber(v string) *BindMFADeviceRequest {
	s.SerialNumber = &v
	return s
}

func (s *BindMFADeviceRequest) SetUserName(v string) *BindMFADeviceRequest {
	s.UserName = &v
	return s
}

type BindMFADeviceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindMFADeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindMFADeviceResponseBody) GoString() string {
	return s.String()
}

func (s *BindMFADeviceResponseBody) SetRequestId(v string) *BindMFADeviceResponseBody {
	s.RequestId = &v
	return s
}

type BindMFADeviceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindMFADeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindMFADeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s BindMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *BindMFADeviceResponse) SetHeaders(v map[string]*string) *BindMFADeviceResponse {
	s.Headers = v
	return s
}

func (s *BindMFADeviceResponse) SetStatusCode(v int32) *BindMFADeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *BindMFADeviceResponse) SetBody(v *BindMFADeviceResponseBody) *BindMFADeviceResponse {
	s.Body = v
	return s
}

type ChangePasswordRequest struct {
	// The new password that is used to log on to the console.
	//
	// The password must meet the complexity requirements. For more information, see [SetPasswordPolicy](https://help.aliyun.com/document_detail/28739.html).
	//
	// example:
	//
	// aw$2****
	NewPassword *string `json:"NewPassword,omitempty" xml:"NewPassword,omitempty"`
	// The old password that is used to log on to the console.
	//
	// example:
	//
	// 12****
	OldPassword *string `json:"OldPassword,omitempty" xml:"OldPassword,omitempty"`
}

func (s ChangePasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangePasswordRequest) GoString() string {
	return s.String()
}

func (s *ChangePasswordRequest) SetNewPassword(v string) *ChangePasswordRequest {
	s.NewPassword = &v
	return s
}

func (s *ChangePasswordRequest) SetOldPassword(v string) *ChangePasswordRequest {
	s.OldPassword = &v
	return s
}

type ChangePasswordResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangePasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangePasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ChangePasswordResponseBody) SetRequestId(v string) *ChangePasswordResponseBody {
	s.RequestId = &v
	return s
}

type ChangePasswordResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangePasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangePasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangePasswordResponse) GoString() string {
	return s.String()
}

func (s *ChangePasswordResponse) SetHeaders(v map[string]*string) *ChangePasswordResponse {
	s.Headers = v
	return s
}

func (s *ChangePasswordResponse) SetStatusCode(v int32) *ChangePasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangePasswordResponse) SetBody(v *ChangePasswordResponseBody) *ChangePasswordResponse {
	s.Body = v
	return s
}

type ClearAccountAliasResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClearAccountAliasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClearAccountAliasResponseBody) GoString() string {
	return s.String()
}

func (s *ClearAccountAliasResponseBody) SetRequestId(v string) *ClearAccountAliasResponseBody {
	s.RequestId = &v
	return s
}

type ClearAccountAliasResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearAccountAliasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClearAccountAliasResponse) String() string {
	return tea.Prettify(s)
}

func (s ClearAccountAliasResponse) GoString() string {
	return s.String()
}

func (s *ClearAccountAliasResponse) SetHeaders(v map[string]*string) *ClearAccountAliasResponse {
	s.Headers = v
	return s
}

func (s *ClearAccountAliasResponse) SetStatusCode(v int32) *ClearAccountAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearAccountAliasResponse) SetBody(v *ClearAccountAliasResponseBody) *ClearAccountAliasResponse {
	s.Body = v
	return s
}

type CreateAccessKeyRequest struct {
	// The name of the RAM user. If a RAM user calls this operation and does not specify this parameter, an AccessKey pair is created for the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateAccessKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessKeyRequest) SetUserName(v string) *CreateAccessKeyRequest {
	s.UserName = &v
	return s
}

type CreateAccessKeyResponseBody struct {
	// The information about the AccessKey pair.
	AccessKey *CreateAccessKeyResponseBodyAccessKey `json:"AccessKey,omitempty" xml:"AccessKey,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccessKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccessKeyResponseBody) SetAccessKey(v *CreateAccessKeyResponseBodyAccessKey) *CreateAccessKeyResponseBody {
	s.AccessKey = v
	return s
}

func (s *CreateAccessKeyResponseBody) SetRequestId(v string) *CreateAccessKeyResponseBody {
	s.RequestId = &v
	return s
}

type CreateAccessKeyResponseBodyAccessKey struct {
	// The AccessKey ID.
	//
	// example:
	//
	// 0wNEpMMlzy7s****
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The AccessKey secret.
	//
	// example:
	//
	// PupkTg8jdmau1cXxYacgE736PJ****
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// The time when the AccessKey pair was created.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The status of the AccessKey pair. Valid values: Active and Inactive.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateAccessKeyResponseBodyAccessKey) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessKeyResponseBodyAccessKey) GoString() string {
	return s.String()
}

func (s *CreateAccessKeyResponseBodyAccessKey) SetAccessKeyId(v string) *CreateAccessKeyResponseBodyAccessKey {
	s.AccessKeyId = &v
	return s
}

func (s *CreateAccessKeyResponseBodyAccessKey) SetAccessKeySecret(v string) *CreateAccessKeyResponseBodyAccessKey {
	s.AccessKeySecret = &v
	return s
}

func (s *CreateAccessKeyResponseBodyAccessKey) SetCreateDate(v string) *CreateAccessKeyResponseBodyAccessKey {
	s.CreateDate = &v
	return s
}

func (s *CreateAccessKeyResponseBodyAccessKey) SetStatus(v string) *CreateAccessKeyResponseBodyAccessKey {
	s.Status = &v
	return s
}

type CreateAccessKeyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAccessKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAccessKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateAccessKeyResponse) SetHeaders(v map[string]*string) *CreateAccessKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateAccessKeyResponse) SetStatusCode(v int32) *CreateAccessKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccessKeyResponse) SetBody(v *CreateAccessKeyResponseBody) *CreateAccessKeyResponse {
	s.Body = v
	return s
}

type CreateGroupRequest struct {
	// The description.
	//
	// The value can be up to 128 characters in length.
	//
	// example:
	//
	// Dev-Team
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The name of the user group.
	//
	// The name must be 1 to 64 characters in length and can contain letters, digits, periods (.), hyphens (-), and underscores (_).
	//
	// example:
	//
	// Dev-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s CreateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupRequest) SetComments(v string) *CreateGroupRequest {
	s.Comments = &v
	return s
}

func (s *CreateGroupRequest) SetGroupName(v string) *CreateGroupRequest {
	s.GroupName = &v
	return s
}

type CreateGroupResponseBody struct {
	// The information about the group.
	Group *CreateGroupResponseBodyGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D3F0679E-9757-95DB-AF2D-04D5188C69C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupResponseBody) SetGroup(v *CreateGroupResponseBodyGroup) *CreateGroupResponseBody {
	s.Group = v
	return s
}

func (s *CreateGroupResponseBody) SetRequestId(v string) *CreateGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateGroupResponseBodyGroup struct {
	// The description.
	//
	// example:
	//
	// Dev-Team
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The ID of the user group.
	//
	// example:
	//
	// g-FpMEHiMysofp****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the user group.
	//
	// example:
	//
	// Dev-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s CreateGroupResponseBodyGroup) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponseBodyGroup) GoString() string {
	return s.String()
}

func (s *CreateGroupResponseBodyGroup) SetComments(v string) *CreateGroupResponseBodyGroup {
	s.Comments = &v
	return s
}

func (s *CreateGroupResponseBodyGroup) SetCreateDate(v string) *CreateGroupResponseBodyGroup {
	s.CreateDate = &v
	return s
}

func (s *CreateGroupResponseBodyGroup) SetGroupId(v string) *CreateGroupResponseBodyGroup {
	s.GroupId = &v
	return s
}

func (s *CreateGroupResponseBodyGroup) SetGroupName(v string) *CreateGroupResponseBodyGroup {
	s.GroupName = &v
	return s
}

type CreateGroupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupResponse) SetHeaders(v map[string]*string) *CreateGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupResponse) SetStatusCode(v int32) *CreateGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupResponse) SetBody(v *CreateGroupResponseBody) *CreateGroupResponse {
	s.Body = v
	return s
}

type CreateLoginProfileRequest struct {
	// Specifies whether the RAM user must bind a multi-factor authentication (MFA) device upon the next logon. Default value: `false`.
	//
	// example:
	//
	// false
	MFABindRequired *bool `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty"`
	// The logon password of the RAM user. The password must meet the password strength requirements. For more information, see [GetPasswordPolicy](https://help.aliyun.com/document_detail/2337691.html).
	//
	// example:
	//
	// mypassword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether the RAM user has to change the password upon logon. Default value: `false`.
	//
	// example:
	//
	// false
	PasswordResetRequired *bool `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty"`
	// The name of the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateLoginProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLoginProfileRequest) GoString() string {
	return s.String()
}

func (s *CreateLoginProfileRequest) SetMFABindRequired(v bool) *CreateLoginProfileRequest {
	s.MFABindRequired = &v
	return s
}

func (s *CreateLoginProfileRequest) SetPassword(v string) *CreateLoginProfileRequest {
	s.Password = &v
	return s
}

func (s *CreateLoginProfileRequest) SetPasswordResetRequired(v bool) *CreateLoginProfileRequest {
	s.PasswordResetRequired = &v
	return s
}

func (s *CreateLoginProfileRequest) SetUserName(v string) *CreateLoginProfileRequest {
	s.UserName = &v
	return s
}

type CreateLoginProfileResponseBody struct {
	// The logon configurations of the RAM user.
	LoginProfile *CreateLoginProfileResponseBodyLoginProfile `json:"LoginProfile,omitempty" xml:"LoginProfile,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLoginProfileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLoginProfileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLoginProfileResponseBody) SetLoginProfile(v *CreateLoginProfileResponseBodyLoginProfile) *CreateLoginProfileResponseBody {
	s.LoginProfile = v
	return s
}

func (s *CreateLoginProfileResponseBody) SetRequestId(v string) *CreateLoginProfileResponseBody {
	s.RequestId = &v
	return s
}

type CreateLoginProfileResponseBodyLoginProfile struct {
	// The creation time.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// Indicates whether an MFA device must be bound to the RAM user.
	//
	// example:
	//
	// false
	MFABindRequired *bool `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty"`
	// Indicates whether the RAM user must change the password upon logon.
	//
	// example:
	//
	// false
	PasswordResetRequired *bool `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty"`
	// The name of the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateLoginProfileResponseBodyLoginProfile) String() string {
	return tea.Prettify(s)
}

func (s CreateLoginProfileResponseBodyLoginProfile) GoString() string {
	return s.String()
}

func (s *CreateLoginProfileResponseBodyLoginProfile) SetCreateDate(v string) *CreateLoginProfileResponseBodyLoginProfile {
	s.CreateDate = &v
	return s
}

func (s *CreateLoginProfileResponseBodyLoginProfile) SetMFABindRequired(v bool) *CreateLoginProfileResponseBodyLoginProfile {
	s.MFABindRequired = &v
	return s
}

func (s *CreateLoginProfileResponseBodyLoginProfile) SetPasswordResetRequired(v bool) *CreateLoginProfileResponseBodyLoginProfile {
	s.PasswordResetRequired = &v
	return s
}

func (s *CreateLoginProfileResponseBodyLoginProfile) SetUserName(v string) *CreateLoginProfileResponseBodyLoginProfile {
	s.UserName = &v
	return s
}

type CreateLoginProfileResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLoginProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLoginProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLoginProfileResponse) GoString() string {
	return s.String()
}

func (s *CreateLoginProfileResponse) SetHeaders(v map[string]*string) *CreateLoginProfileResponse {
	s.Headers = v
	return s
}

func (s *CreateLoginProfileResponse) SetStatusCode(v int32) *CreateLoginProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLoginProfileResponse) SetBody(v *CreateLoginProfileResponseBody) *CreateLoginProfileResponse {
	s.Body = v
	return s
}

type CreatePolicyRequest struct {
	// The description of the policy.
	//
	// The description must be 1 to 1,024 characters in length.
	//
	// example:
	//
	// Query ECS instances in a specific region
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The document of the policy.
	//
	// The document must be 1 to 6,144 characters in length.
	//
	// For more information about policy elements and sample policies, see [Policy elements](https://help.aliyun.com/document_detail/93738.html) and [Overview of sample policies](https://help.aliyun.com/document_detail/210969.html).
	//
	// example:
	//
	// {"Statement": [{"Effect": "Allow","Action": "ecs:Describe*","Resource": "acs:ecs:cn-qingdao:*:instance/*"}],"Version": "1"}
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The name of the policy.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	//
	// example:
	//
	// View-ECS-instances-in-a-specific-region
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The tags.
	Tag []*CreatePolicyRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreatePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyRequest) SetDescription(v string) *CreatePolicyRequest {
	s.Description = &v
	return s
}

func (s *CreatePolicyRequest) SetPolicyDocument(v string) *CreatePolicyRequest {
	s.PolicyDocument = &v
	return s
}

func (s *CreatePolicyRequest) SetPolicyName(v string) *CreatePolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *CreatePolicyRequest) SetTag(v []*CreatePolicyRequestTag) *CreatePolicyRequest {
	s.Tag = v
	return s
}

type CreatePolicyRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// owner
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// alice
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePolicyRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyRequestTag) GoString() string {
	return s.String()
}

func (s *CreatePolicyRequestTag) SetKey(v string) *CreatePolicyRequestTag {
	s.Key = &v
	return s
}

func (s *CreatePolicyRequestTag) SetValue(v string) *CreatePolicyRequestTag {
	s.Value = &v
	return s
}

type CreatePolicyShrinkRequest struct {
	// The description of the policy.
	//
	// The description must be 1 to 1,024 characters in length.
	//
	// example:
	//
	// Query ECS instances in a specific region
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The document of the policy.
	//
	// The document must be 1 to 6,144 characters in length.
	//
	// For more information about policy elements and sample policies, see [Policy elements](https://help.aliyun.com/document_detail/93738.html) and [Overview of sample policies](https://help.aliyun.com/document_detail/210969.html).
	//
	// example:
	//
	// {"Statement": [{"Effect": "Allow","Action": "ecs:Describe*","Resource": "acs:ecs:cn-qingdao:*:instance/*"}],"Version": "1"}
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The name of the policy.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	//
	// example:
	//
	// View-ECS-instances-in-a-specific-region
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The tags.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s CreatePolicyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyShrinkRequest) SetDescription(v string) *CreatePolicyShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetPolicyDocument(v string) *CreatePolicyShrinkRequest {
	s.PolicyDocument = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetPolicyName(v string) *CreatePolicyShrinkRequest {
	s.PolicyName = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetTagShrink(v string) *CreatePolicyShrinkRequest {
	s.TagShrink = &v
	return s
}

type CreatePolicyResponseBody struct {
	// The information about the policy.
	Policy *CreatePolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// BA34C54A-C2B1-5A65-B6B0-B5842C1DB4DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolicyResponseBody) SetPolicy(v *CreatePolicyResponseBodyPolicy) *CreatePolicyResponseBody {
	s.Policy = v
	return s
}

func (s *CreatePolicyResponseBody) SetRequestId(v string) *CreatePolicyResponseBody {
	s.RequestId = &v
	return s
}

type CreatePolicyResponseBodyPolicy struct {
	// The time when the policy was created.
	//
	// example:
	//
	// 2021-10-13T02:46:57Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The version of the policy. Default value: v1.
	//
	// example:
	//
	// v1
	DefaultVersion *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// The description of the policy.
	//
	// example:
	//
	// Query ECS instances in a specific region
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// View-ECS-instances-in-a-specific-region
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy. Valid values:
	//
	// 	- Custom
	//
	// 	- System
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s CreatePolicyResponseBodyPolicy) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *CreatePolicyResponseBodyPolicy) SetCreateDate(v string) *CreatePolicyResponseBodyPolicy {
	s.CreateDate = &v
	return s
}

func (s *CreatePolicyResponseBodyPolicy) SetDefaultVersion(v string) *CreatePolicyResponseBodyPolicy {
	s.DefaultVersion = &v
	return s
}

func (s *CreatePolicyResponseBodyPolicy) SetDescription(v string) *CreatePolicyResponseBodyPolicy {
	s.Description = &v
	return s
}

func (s *CreatePolicyResponseBodyPolicy) SetPolicyName(v string) *CreatePolicyResponseBodyPolicy {
	s.PolicyName = &v
	return s
}

func (s *CreatePolicyResponseBodyPolicy) SetPolicyType(v string) *CreatePolicyResponseBodyPolicy {
	s.PolicyType = &v
	return s
}

type CreatePolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyResponse) GoString() string {
	return s.String()
}

func (s *CreatePolicyResponse) SetHeaders(v map[string]*string) *CreatePolicyResponse {
	s.Headers = v
	return s
}

func (s *CreatePolicyResponse) SetStatusCode(v int32) *CreatePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePolicyResponse) SetBody(v *CreatePolicyResponseBody) *CreatePolicyResponse {
	s.Body = v
	return s
}

type CreatePolicyVersionRequest struct {
	// The document of the policy. The document can be up to 6,144 bytes in length.
	//
	// example:
	//
	// {"Statement":[{"Action":["oss:*"],"Effect":"Allow","Resource":["acs:oss:*:*:*"]}],"Version":"1"}
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The rotation strategy of the policy. The rotation strategy can be used to delete an early policy version.
	//
	// Valid values:
	//
	// 	- `None`: disables the rotation strategy.
	//
	// 	- `DeleteOldestNonDefaultVersionWhenLimitExceeded`: deletes the earliest non-active version if the number of versions exceeds the limit.
	//
	// Default value: `None`.
	//
	// example:
	//
	// None
	RotateStrategy *string `json:"RotateStrategy,omitempty" xml:"RotateStrategy,omitempty"`
	// Specifies whether to set this policy as the default policy. Default value: `false`.
	//
	// example:
	//
	// false
	SetAsDefault *bool `json:"SetAsDefault,omitempty" xml:"SetAsDefault,omitempty"`
}

func (s CreatePolicyVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyVersionRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyVersionRequest) SetPolicyDocument(v string) *CreatePolicyVersionRequest {
	s.PolicyDocument = &v
	return s
}

func (s *CreatePolicyVersionRequest) SetPolicyName(v string) *CreatePolicyVersionRequest {
	s.PolicyName = &v
	return s
}

func (s *CreatePolicyVersionRequest) SetRotateStrategy(v string) *CreatePolicyVersionRequest {
	s.RotateStrategy = &v
	return s
}

func (s *CreatePolicyVersionRequest) SetSetAsDefault(v bool) *CreatePolicyVersionRequest {
	s.SetAsDefault = &v
	return s
}

type CreatePolicyVersionResponseBody struct {
	// The information about the policy version.
	PolicyVersion *CreatePolicyVersionResponseBodyPolicyVersion `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePolicyVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolicyVersionResponseBody) SetPolicyVersion(v *CreatePolicyVersionResponseBodyPolicyVersion) *CreatePolicyVersionResponseBody {
	s.PolicyVersion = v
	return s
}

func (s *CreatePolicyVersionResponseBody) SetRequestId(v string) *CreatePolicyVersionResponseBody {
	s.RequestId = &v
	return s
}

type CreatePolicyVersionResponseBodyPolicyVersion struct {
	// The time when the policy version was created.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// Indicates whether the policy version is the default version.
	//
	// example:
	//
	// false
	IsDefaultVersion *bool `json:"IsDefaultVersion,omitempty" xml:"IsDefaultVersion,omitempty"`
	// The document of the policy.
	//
	// example:
	//
	// { "Statement": [{ "Action": ["oss:*"], "Effect": "Allow", "Resource": ["acs:oss:*:*:*"]}], "Version": "1"}
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The ID of the policy version.
	//
	// example:
	//
	// v3
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s CreatePolicyVersionResponseBodyPolicyVersion) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyVersionResponseBodyPolicyVersion) GoString() string {
	return s.String()
}

func (s *CreatePolicyVersionResponseBodyPolicyVersion) SetCreateDate(v string) *CreatePolicyVersionResponseBodyPolicyVersion {
	s.CreateDate = &v
	return s
}

func (s *CreatePolicyVersionResponseBodyPolicyVersion) SetIsDefaultVersion(v bool) *CreatePolicyVersionResponseBodyPolicyVersion {
	s.IsDefaultVersion = &v
	return s
}

func (s *CreatePolicyVersionResponseBodyPolicyVersion) SetPolicyDocument(v string) *CreatePolicyVersionResponseBodyPolicyVersion {
	s.PolicyDocument = &v
	return s
}

func (s *CreatePolicyVersionResponseBodyPolicyVersion) SetVersionId(v string) *CreatePolicyVersionResponseBodyPolicyVersion {
	s.VersionId = &v
	return s
}

type CreatePolicyVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePolicyVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePolicyVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyVersionResponse) GoString() string {
	return s.String()
}

func (s *CreatePolicyVersionResponse) SetHeaders(v map[string]*string) *CreatePolicyVersionResponse {
	s.Headers = v
	return s
}

func (s *CreatePolicyVersionResponse) SetStatusCode(v int32) *CreatePolicyVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePolicyVersionResponse) SetBody(v *CreatePolicyVersionResponseBody) *CreatePolicyVersionResponse {
	s.Body = v
	return s
}

type CreateRoleRequest struct {
	// The trust policy that specifies one or more trusted entities to assume the RAM role. The trusted entities can be Alibaba Cloud accounts, Alibaba Cloud services, or identity providers (IdPs).
	//
	// >  RAM users cannot assume the RAM roles of trusted Alibaba Cloud services.
	//
	// example:
	//
	// {"Statement":[{"Action":"sts:AssumeRole","Effect":"Allow","Principal":{"RAM":"acs:ram::123456789012****:root"}}],"Version":"1"}
	AssumeRolePolicyDocument *string `json:"AssumeRolePolicyDocument,omitempty" xml:"AssumeRolePolicyDocument,omitempty"`
	// The description of the RAM role.
	//
	// The description must be 1 to 1,024 characters in length.
	//
	// example:
	//
	// ECS administrator
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum session time of the RAM role.
	//
	// Valid values: 3600 to 43200. Unit: seconds. Default value: 3600.
	//
	// If you do not specify this parameter, the default value is used.
	//
	// example:
	//
	// 3600
	MaxSessionDuration *int64 `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty"`
	// The name of the RAM role.
	//
	// The name must be 1 to 64 characters in length, and can contain letters, digits, periods (.), and hyphens (-).
	//
	// example:
	//
	// ECSAdmin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The tags.
	Tag []*CreateRoleRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateRoleRequest) SetAssumeRolePolicyDocument(v string) *CreateRoleRequest {
	s.AssumeRolePolicyDocument = &v
	return s
}

func (s *CreateRoleRequest) SetDescription(v string) *CreateRoleRequest {
	s.Description = &v
	return s
}

func (s *CreateRoleRequest) SetMaxSessionDuration(v int64) *CreateRoleRequest {
	s.MaxSessionDuration = &v
	return s
}

func (s *CreateRoleRequest) SetRoleName(v string) *CreateRoleRequest {
	s.RoleName = &v
	return s
}

func (s *CreateRoleRequest) SetTag(v []*CreateRoleRequestTag) *CreateRoleRequest {
	s.Tag = v
	return s
}

type CreateRoleRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateRoleRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateRoleRequestTag) GoString() string {
	return s.String()
}

func (s *CreateRoleRequestTag) SetKey(v string) *CreateRoleRequestTag {
	s.Key = &v
	return s
}

func (s *CreateRoleRequestTag) SetValue(v string) *CreateRoleRequestTag {
	s.Value = &v
	return s
}

type CreateRoleShrinkRequest struct {
	// The trust policy that specifies one or more trusted entities to assume the RAM role. The trusted entities can be Alibaba Cloud accounts, Alibaba Cloud services, or identity providers (IdPs).
	//
	// >  RAM users cannot assume the RAM roles of trusted Alibaba Cloud services.
	//
	// example:
	//
	// {"Statement":[{"Action":"sts:AssumeRole","Effect":"Allow","Principal":{"RAM":"acs:ram::123456789012****:root"}}],"Version":"1"}
	AssumeRolePolicyDocument *string `json:"AssumeRolePolicyDocument,omitempty" xml:"AssumeRolePolicyDocument,omitempty"`
	// The description of the RAM role.
	//
	// The description must be 1 to 1,024 characters in length.
	//
	// example:
	//
	// ECS administrator
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum session time of the RAM role.
	//
	// Valid values: 3600 to 43200. Unit: seconds. Default value: 3600.
	//
	// If you do not specify this parameter, the default value is used.
	//
	// example:
	//
	// 3600
	MaxSessionDuration *int64 `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty"`
	// The name of the RAM role.
	//
	// The name must be 1 to 64 characters in length, and can contain letters, digits, periods (.), and hyphens (-).
	//
	// example:
	//
	// ECSAdmin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The tags.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s CreateRoleShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRoleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateRoleShrinkRequest) SetAssumeRolePolicyDocument(v string) *CreateRoleShrinkRequest {
	s.AssumeRolePolicyDocument = &v
	return s
}

func (s *CreateRoleShrinkRequest) SetDescription(v string) *CreateRoleShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateRoleShrinkRequest) SetMaxSessionDuration(v int64) *CreateRoleShrinkRequest {
	s.MaxSessionDuration = &v
	return s
}

func (s *CreateRoleShrinkRequest) SetRoleName(v string) *CreateRoleShrinkRequest {
	s.RoleName = &v
	return s
}

func (s *CreateRoleShrinkRequest) SetTagShrink(v string) *CreateRoleShrinkRequest {
	s.TagShrink = &v
	return s
}

type CreateRoleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the RAM role.
	Role *CreateRoleResponseBodyRole `json:"Role,omitempty" xml:"Role,omitempty" type:"Struct"`
}

func (s CreateRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRoleResponseBody) SetRequestId(v string) *CreateRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRoleResponseBody) SetRole(v *CreateRoleResponseBodyRole) *CreateRoleResponseBody {
	s.Role = v
	return s
}

type CreateRoleResponseBodyRole struct {
	// The Alibaba Cloud Resource Name (ARN) of the RAM role.
	//
	// example:
	//
	// acs:ram::123456789012****:role/ECSAdmin
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The trust policy that specifies the trusted entity to assume the RAM role.
	//
	// example:
	//
	// { "Statement": [ { "Action": "sts:AssumeRole", "Effect": "Allow", "Principal": { "RAM": "acs:ram::123456789012****:root" } } ], "Version": "1" }
	AssumeRolePolicyDocument *string `json:"AssumeRolePolicyDocument,omitempty" xml:"AssumeRolePolicyDocument,omitempty"`
	// The time when the RAM role was created.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the RAM role.
	//
	// example:
	//
	// ECS administrator
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum session time of the RAM role.
	//
	// example:
	//
	// 3600
	MaxSessionDuration *int64 `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty"`
	// The ID of the RAM role.
	//
	// example:
	//
	// 901234567890****
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// The name of the RAM role.
	//
	// example:
	//
	// ECSAdmin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s CreateRoleResponseBodyRole) String() string {
	return tea.Prettify(s)
}

func (s CreateRoleResponseBodyRole) GoString() string {
	return s.String()
}

func (s *CreateRoleResponseBodyRole) SetArn(v string) *CreateRoleResponseBodyRole {
	s.Arn = &v
	return s
}

func (s *CreateRoleResponseBodyRole) SetAssumeRolePolicyDocument(v string) *CreateRoleResponseBodyRole {
	s.AssumeRolePolicyDocument = &v
	return s
}

func (s *CreateRoleResponseBodyRole) SetCreateDate(v string) *CreateRoleResponseBodyRole {
	s.CreateDate = &v
	return s
}

func (s *CreateRoleResponseBodyRole) SetDescription(v string) *CreateRoleResponseBodyRole {
	s.Description = &v
	return s
}

func (s *CreateRoleResponseBodyRole) SetMaxSessionDuration(v int64) *CreateRoleResponseBodyRole {
	s.MaxSessionDuration = &v
	return s
}

func (s *CreateRoleResponseBodyRole) SetRoleId(v string) *CreateRoleResponseBodyRole {
	s.RoleId = &v
	return s
}

func (s *CreateRoleResponseBodyRole) SetRoleName(v string) *CreateRoleResponseBodyRole {
	s.RoleName = &v
	return s
}

type CreateRoleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateRoleResponse) SetHeaders(v map[string]*string) *CreateRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateRoleResponse) SetStatusCode(v int32) *CreateRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRoleResponse) SetBody(v *CreateRoleResponseBody) *CreateRoleResponse {
	s.Body = v
	return s
}

type CreateUserRequest struct {
	// The description of the RAM user.
	//
	// The description must be 1 to 128 characters in length.
	//
	// example:
	//
	// This is a cloud computing engineer.
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The display name of the RAM user.
	//
	// The name must be 1 to 128 characters in length.
	//
	// example:
	//
	// alice
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the RAM user.
	//
	// >  This parameter applies only to the China site (aliyun.com).
	//
	// example:
	//
	// alice@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The mobile phone number of the RAM user.
	//
	// Format: \\<Country code>-\\<Mobile phone number>.
	//
	// >  This parameter applies only to the China site (aliyun.com).
	//
	// example:
	//
	// 86-1868888****
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	// The name of the RAM user.
	//
	// The name must be 1 to 64 characters in length, and can contain letters, digits, periods (.), hyphens (-), and underscores (_).
	//
	// example:
	//
	// alice
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) SetComments(v string) *CreateUserRequest {
	s.Comments = &v
	return s
}

func (s *CreateUserRequest) SetDisplayName(v string) *CreateUserRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateUserRequest) SetEmail(v string) *CreateUserRequest {
	s.Email = &v
	return s
}

func (s *CreateUserRequest) SetMobilePhone(v string) *CreateUserRequest {
	s.MobilePhone = &v
	return s
}

func (s *CreateUserRequest) SetUserName(v string) *CreateUserRequest {
	s.UserName = &v
	return s
}

type CreateUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the RAM user.
	User *CreateUserResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s CreateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBody) SetRequestId(v string) *CreateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserResponseBody) SetUser(v *CreateUserResponseBodyUser) *CreateUserResponseBody {
	s.User = v
	return s
}

type CreateUserResponseBodyUser struct {
	// The description of the RAM user.
	//
	// example:
	//
	// This is a cloud computing engineer.
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The point in time when the RAM user was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The display name of the RAM user.
	//
	// example:
	//
	// alice
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the RAM user.
	//
	// >  This parameter applies only to the China site (aliyun.com).
	//
	// example:
	//
	// alice@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The mobile phone number of the RAM user.
	//
	// >  This parameter applies only to the China site (aliyun.com).
	//
	// example:
	//
	// 86-1868888****
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	// The ID of the RAM user.
	//
	// example:
	//
	// 122748924538****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The name of the RAM user.
	//
	// example:
	//
	// alice
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateUserResponseBodyUser) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponseBodyUser) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBodyUser) SetComments(v string) *CreateUserResponseBodyUser {
	s.Comments = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetCreateDate(v string) *CreateUserResponseBodyUser {
	s.CreateDate = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetDisplayName(v string) *CreateUserResponseBodyUser {
	s.DisplayName = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetEmail(v string) *CreateUserResponseBodyUser {
	s.Email = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetMobilePhone(v string) *CreateUserResponseBodyUser {
	s.MobilePhone = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetUserId(v string) *CreateUserResponseBodyUser {
	s.UserId = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetUserName(v string) *CreateUserResponseBodyUser {
	s.UserName = &v
	return s
}

type CreateUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponse) GoString() string {
	return s.String()
}

func (s *CreateUserResponse) SetHeaders(v map[string]*string) *CreateUserResponse {
	s.Headers = v
	return s
}

func (s *CreateUserResponse) SetStatusCode(v int32) *CreateUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserResponse) SetBody(v *CreateUserResponseBody) *CreateUserResponse {
	s.Body = v
	return s
}

type CreateVirtualMFADeviceRequest struct {
	// The name of the MFA device.
	//
	// The name must be 1 to 64 characters in length, and can contain letters, digits, periods (.), and hyphens (-).
	//
	// example:
	//
	// device001
	VirtualMFADeviceName *string `json:"VirtualMFADeviceName,omitempty" xml:"VirtualMFADeviceName,omitempty"`
}

func (s CreateVirtualMFADeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *CreateVirtualMFADeviceRequest) SetVirtualMFADeviceName(v string) *CreateVirtualMFADeviceRequest {
	s.VirtualMFADeviceName = &v
	return s
}

type CreateVirtualMFADeviceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the MFA device.
	VirtualMFADevice *CreateVirtualMFADeviceResponseBodyVirtualMFADevice `json:"VirtualMFADevice,omitempty" xml:"VirtualMFADevice,omitempty" type:"Struct"`
}

func (s CreateVirtualMFADeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualMFADeviceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVirtualMFADeviceResponseBody) SetRequestId(v string) *CreateVirtualMFADeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVirtualMFADeviceResponseBody) SetVirtualMFADevice(v *CreateVirtualMFADeviceResponseBodyVirtualMFADevice) *CreateVirtualMFADeviceResponseBody {
	s.VirtualMFADevice = v
	return s
}

type CreateVirtualMFADeviceResponseBodyVirtualMFADevice struct {
	// The key of the MFA device.
	//
	// example:
	//
	// DSF98HAD982KJA9SDFNAS9D8FU839B8ADHBGS****
	Base32StringSeed *string `json:"Base32StringSeed,omitempty" xml:"Base32StringSeed,omitempty"`
	// The Base64-encoded QR code, in the PNG format.
	//
	// example:
	//
	// YXNkZmFzZDlmeW5hc2Q5OGZoODd4bXJmcThhaGU5aSBmYXNkZiBzYWRmIGFGIDRxd2VjIGEgdHEz****
	QRCodePNG *string `json:"QRCodePNG,omitempty" xml:"QRCodePNG,omitempty"`
	// The serial number of the MFA device.
	//
	// example:
	//
	// acs:ram::123456789012****:mfa/device001
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s CreateVirtualMFADeviceResponseBodyVirtualMFADevice) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualMFADeviceResponseBodyVirtualMFADevice) GoString() string {
	return s.String()
}

func (s *CreateVirtualMFADeviceResponseBodyVirtualMFADevice) SetBase32StringSeed(v string) *CreateVirtualMFADeviceResponseBodyVirtualMFADevice {
	s.Base32StringSeed = &v
	return s
}

func (s *CreateVirtualMFADeviceResponseBodyVirtualMFADevice) SetQRCodePNG(v string) *CreateVirtualMFADeviceResponseBodyVirtualMFADevice {
	s.QRCodePNG = &v
	return s
}

func (s *CreateVirtualMFADeviceResponseBodyVirtualMFADevice) SetSerialNumber(v string) *CreateVirtualMFADeviceResponseBodyVirtualMFADevice {
	s.SerialNumber = &v
	return s
}

type CreateVirtualMFADeviceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVirtualMFADeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVirtualMFADeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *CreateVirtualMFADeviceResponse) SetHeaders(v map[string]*string) *CreateVirtualMFADeviceResponse {
	s.Headers = v
	return s
}

func (s *CreateVirtualMFADeviceResponse) SetStatusCode(v int32) *CreateVirtualMFADeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVirtualMFADeviceResponse) SetBody(v *CreateVirtualMFADeviceResponseBody) *CreateVirtualMFADeviceResponse {
	s.Body = v
	return s
}

type DecodeDiagnosticMessageRequest struct {
	// The encoded diagnostic information in the response that contains an access denied error. The error is caused by no RAM permissions.
	//
	// example:
	//
	// AQEAAAAAZBgxr0U1MjA1NTM1LUM4BBktMzE5RS1CODgxLUU1QTI0RDNFQTM1****
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
}

func (s DecodeDiagnosticMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s DecodeDiagnosticMessageRequest) GoString() string {
	return s.String()
}

func (s *DecodeDiagnosticMessageRequest) SetEncodedDiagnosticMessage(v string) *DecodeDiagnosticMessageRequest {
	s.EncodedDiagnosticMessage = &v
	return s
}

type DecodeDiagnosticMessageResponseBody struct {
	// The decoded diagnostic information.
	DecodedDiagnosticMessage *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage `json:"DecodedDiagnosticMessage,omitempty" xml:"DecodedDiagnosticMessage,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D2331703-AADF-5564-BA9B-26CD51A33BA0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DecodeDiagnosticMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DecodeDiagnosticMessageResponseBody) GoString() string {
	return s.String()
}

func (s *DecodeDiagnosticMessageResponseBody) SetDecodedDiagnosticMessage(v *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage) *DecodeDiagnosticMessageResponseBody {
	s.DecodedDiagnosticMessage = v
	return s
}

func (s *DecodeDiagnosticMessageResponseBody) SetRequestId(v string) *DecodeDiagnosticMessageResponseBody {
	s.RequestId = &v
	return s
}

type DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage struct {
	// The operation that is used for authentication in the request.
	//
	// example:
	//
	// ram:DecodeDiagnosticMessage
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The conditions that are used for authentication in the request.
	AuthConditions []*DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthConditions `json:"AuthConditions,omitempty" xml:"AuthConditions,omitempty" type:"Repeated"`
	// The operator that is used for authentication in the request.
	AuthPrincipal *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthPrincipal `json:"AuthPrincipal,omitempty" xml:"AuthPrincipal,omitempty" type:"Struct"`
	// The resource that is used for authentication in the request.
	//
	// example:
	//
	// *
	AuthResource *string `json:"AuthResource,omitempty" xml:"AuthResource,omitempty"`
	// Indicates whether the access denied error is caused by an explicit deny.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	ExplicitDeny *bool `json:"ExplicitDeny,omitempty" xml:"ExplicitDeny,omitempty"`
	// The policies that are matched.
	MatchedPolicies []*DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies `json:"MatchedPolicies,omitempty" xml:"MatchedPolicies,omitempty" type:"Repeated"`
	// The type of the policy that causes the access denied error.
	//
	// Valid values:
	//
	// 	- AssumeRolePolicy: role-specific trust policy
	//
	// 	- ControlPolicy: control policy
	//
	// 	- AccountLevelIdentityBasedPolicy: identity-based policy at the account level
	//
	// 	- ResourceGroupLevelIdentityBasedPolicy: identity-based policy at the resource group level
	//
	// 	- SessionPolicy: session policy
	//
	// example:
	//
	// AccountLevelIdentityBasedPolicy
	NoPermissionPolicyType *string `json:"NoPermissionPolicyType,omitempty" xml:"NoPermissionPolicyType,omitempty"`
}

func (s DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage) String() string {
	return tea.Prettify(s)
}

func (s DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage) GoString() string {
	return s.String()
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage) SetAuthAction(v string) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage {
	s.AuthAction = &v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage) SetAuthConditions(v []*DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthConditions) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage {
	s.AuthConditions = v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage) SetAuthPrincipal(v *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthPrincipal) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage {
	s.AuthPrincipal = v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage) SetAuthResource(v string) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage {
	s.AuthResource = &v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage) SetExplicitDeny(v bool) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage {
	s.ExplicitDeny = &v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage) SetMatchedPolicies(v []*DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage {
	s.MatchedPolicies = v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage) SetNoPermissionPolicyType(v string) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage {
	s.NoPermissionPolicyType = &v
	return s
}

type DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthConditions struct {
	// The key of the condition.
	//
	// example:
	//
	// acs:SourceIp
	ConditionKey *string `json:"ConditionKey,omitempty" xml:"ConditionKey,omitempty"`
	// The values that correspond to the key.
	ConditionValues []*string `json:"ConditionValues,omitempty" xml:"ConditionValues,omitempty" type:"Repeated"`
}

func (s DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthConditions) String() string {
	return tea.Prettify(s)
}

func (s DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthConditions) GoString() string {
	return s.String()
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthConditions) SetConditionKey(v string) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthConditions {
	s.ConditionKey = &v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthConditions) SetConditionValues(v []*string) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthConditions {
	s.ConditionValues = v
	return s
}

type DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthPrincipal struct {
	// The identity.
	//
	// 	- If the operator is a RAM user, the ID of the user is displayed.
	//
	// 	- If the operator is a RAM role, the name and session name of the role are displayed. Example: RoleName:RoleSessionName.
	//
	// 	- If the operator is an SSO federated identity, the type and name of the identity provider (IdP) are displayed. Example: saml-provider/AzureAD.
	//
	// example:
	//
	// 28877424437521****
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The ID of the Alibaba Cloud account to which the identity belongs.
	//
	// example:
	//
	// 196813200012****
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The identity type that is used for authentication in the request.
	//
	// Valid values:
	//
	// 	- SubUser: RAM user
	//
	// 	- AssumedRoleUser: RAM role
	//
	// 	- Federated: SSO federated identity
	//
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
}

func (s DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthPrincipal) String() string {
	return tea.Prettify(s)
}

func (s DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthPrincipal) GoString() string {
	return s.String()
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthPrincipal) SetAuthPrincipalDisplayName(v string) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthPrincipal {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthPrincipal) SetAuthPrincipalOwnerId(v string) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthPrincipal {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthPrincipal) SetAuthPrincipalType(v string) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthPrincipal {
	s.AuthPrincipalType = &v
	return s
}

type DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies struct {
	// The type of the entity to which the policy is attached.
	//
	// Valid values:
	//
	// 	- RamUser: RAM user
	//
	// 	- RamRole: RAM role
	//
	// 	- ResourceDirectoryTarget: entity in a resource directory
	//
	// 	- RamGroup: RAM user group
	//
	// example:
	//
	// RamUser
	AttachedEntityType *string `json:"AttachedEntityType,omitempty" xml:"AttachedEntityType,omitempty"`
	// The authorization scope of the policy.
	//
	// Valid values:
	//
	// 	- Account: Alibaba Cloud account
	//
	// 	- Folder: folder in the resource directory
	//
	// 	- ResourceGroup: resource group
	//
	// example:
	//
	// Account
	AttachedScope *string `json:"AttachedScope,omitempty" xml:"AttachedScope,omitempty"`
	// The effect of the policy.
	//
	// Valid values:
	//
	// 	- Deny
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Allow
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Deny
	Effect *string `json:"Effect,omitempty" xml:"Effect,omitempty"`
	// The identifier of the policy.
	//
	// 	- Control policy: the ID of the control policy
	//
	// 	- RAM policy: the name of the policy
	//
	// example:
	//
	// MyPolicyName
	PolicyIdentifier *string `json:"PolicyIdentifier,omitempty" xml:"PolicyIdentifier,omitempty"`
	// The type of the policy.
	//
	// Valid values:
	//
	// 	- Custom: custom policy
	//
	// 	- System: system policy
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The version number of the policy.
	//
	// > Only custom policies have version numbers.
	//
	// example:
	//
	// v1
	PolicyVersion *string `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
}

func (s DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies) String() string {
	return tea.Prettify(s)
}

func (s DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies) GoString() string {
	return s.String()
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies) SetAttachedEntityType(v string) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies {
	s.AttachedEntityType = &v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies) SetAttachedScope(v string) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies {
	s.AttachedScope = &v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies) SetEffect(v string) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies {
	s.Effect = &v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies) SetPolicyIdentifier(v string) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies {
	s.PolicyIdentifier = &v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies) SetPolicyType(v string) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies {
	s.PolicyType = &v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies) SetPolicyVersion(v string) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies {
	s.PolicyVersion = &v
	return s
}

type DecodeDiagnosticMessageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DecodeDiagnosticMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DecodeDiagnosticMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s DecodeDiagnosticMessageResponse) GoString() string {
	return s.String()
}

func (s *DecodeDiagnosticMessageResponse) SetHeaders(v map[string]*string) *DecodeDiagnosticMessageResponse {
	s.Headers = v
	return s
}

func (s *DecodeDiagnosticMessageResponse) SetStatusCode(v int32) *DecodeDiagnosticMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *DecodeDiagnosticMessageResponse) SetBody(v *DecodeDiagnosticMessageResponseBody) *DecodeDiagnosticMessageResponse {
	s.Body = v
	return s
}

type DeleteAccessKeyRequest struct {
	// The AccessKey ID in the AccessKey pair that you want to delete.``
	//
	// example:
	//
	// 0wNEpMMlzy7s****
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// The name of the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DeleteAccessKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessKeyRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessKeyRequest) SetUserAccessKeyId(v string) *DeleteAccessKeyRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *DeleteAccessKeyRequest) SetUserName(v string) *DeleteAccessKeyRequest {
	s.UserName = &v
	return s
}

type DeleteAccessKeyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccessKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccessKeyResponseBody) SetRequestId(v string) *DeleteAccessKeyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAccessKeyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAccessKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAccessKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessKeyResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessKeyResponse) SetHeaders(v map[string]*string) *DeleteAccessKeyResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessKeyResponse) SetStatusCode(v int32) *DeleteAccessKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessKeyResponse) SetBody(v *DeleteAccessKeyResponseBody) *DeleteAccessKeyResponse {
	s.Body = v
	return s
}

type DeleteGroupRequest struct {
	// The name of the RAM user group.
	//
	// If you want to query the name of a RAM user group, you can call the [ListGroups](https://help.aliyun.com/document_detail/28703.html) operation.
	//
	// example:
	//
	// Dev-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s DeleteGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupRequest) SetGroupName(v string) *DeleteGroupRequest {
	s.GroupName = &v
	return s
}

type DeleteGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FCF40AB5-881C-A0F9-334C-B0AD423AA69D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGroupResponseBody) SetRequestId(v string) *DeleteGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGroupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteGroupResponse) SetHeaders(v map[string]*string) *DeleteGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteGroupResponse) SetStatusCode(v int32) *DeleteGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGroupResponse) SetBody(v *DeleteGroupResponseBody) *DeleteGroupResponse {
	s.Body = v
	return s
}

type DeleteLoginProfileRequest struct {
	// The name of the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DeleteLoginProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoginProfileRequest) GoString() string {
	return s.String()
}

func (s *DeleteLoginProfileRequest) SetUserName(v string) *DeleteLoginProfileRequest {
	s.UserName = &v
	return s
}

type DeleteLoginProfileResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1C488B66-B819-4D14-8711-C4EAAA13AC01
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLoginProfileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoginProfileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLoginProfileResponseBody) SetRequestId(v string) *DeleteLoginProfileResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLoginProfileResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLoginProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLoginProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoginProfileResponse) GoString() string {
	return s.String()
}

func (s *DeleteLoginProfileResponse) SetHeaders(v map[string]*string) *DeleteLoginProfileResponse {
	s.Headers = v
	return s
}

func (s *DeleteLoginProfileResponse) SetStatusCode(v int32) *DeleteLoginProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLoginProfileResponse) SetBody(v *DeleteLoginProfileResponseBody) *DeleteLoginProfileResponse {
	s.Body = v
	return s
}

type DeletePolicyRequest struct {
	// Specifies whether to delete all versions of the policy. Valid values:
	//
	// 	- true: deletes all versions of the policy.
	//
	// 	- false: does not delete all versions of the policy. If you set the parameter to false, the non-default versions of the policy are not deleted. Before you delete the policy, you must manually delete all non-default versions of the policy.
	//
	// example:
	//
	// true
	CascadingDelete *bool `json:"CascadingDelete,omitempty" xml:"CascadingDelete,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s DeletePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyRequest) GoString() string {
	return s.String()
}

func (s *DeletePolicyRequest) SetCascadingDelete(v bool) *DeletePolicyRequest {
	s.CascadingDelete = &v
	return s
}

func (s *DeletePolicyRequest) SetPolicyName(v string) *DeletePolicyRequest {
	s.PolicyName = &v
	return s
}

type DeletePolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 898FAB24-7509-43EE-A287-086FE4C44394
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolicyResponseBody) SetRequestId(v string) *DeletePolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeletePolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyResponse) GoString() string {
	return s.String()
}

func (s *DeletePolicyResponse) SetHeaders(v map[string]*string) *DeletePolicyResponse {
	s.Headers = v
	return s
}

func (s *DeletePolicyResponse) SetStatusCode(v int32) *DeletePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePolicyResponse) SetBody(v *DeletePolicyResponseBody) *DeletePolicyResponse {
	s.Body = v
	return s
}

type DeletePolicyVersionRequest struct {
	// The name of the policy.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The ID of the policy version that you want to delete.
	//
	// example:
	//
	// v3
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s DeletePolicyVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyVersionRequest) GoString() string {
	return s.String()
}

func (s *DeletePolicyVersionRequest) SetPolicyName(v string) *DeletePolicyVersionRequest {
	s.PolicyName = &v
	return s
}

func (s *DeletePolicyVersionRequest) SetVersionId(v string) *DeletePolicyVersionRequest {
	s.VersionId = &v
	return s
}

type DeletePolicyVersionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePolicyVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolicyVersionResponseBody) SetRequestId(v string) *DeletePolicyVersionResponseBody {
	s.RequestId = &v
	return s
}

type DeletePolicyVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePolicyVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePolicyVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyVersionResponse) GoString() string {
	return s.String()
}

func (s *DeletePolicyVersionResponse) SetHeaders(v map[string]*string) *DeletePolicyVersionResponse {
	s.Headers = v
	return s
}

func (s *DeletePolicyVersionResponse) SetStatusCode(v int32) *DeletePolicyVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePolicyVersionResponse) SetBody(v *DeletePolicyVersionResponseBody) *DeletePolicyVersionResponse {
	s.Body = v
	return s
}

type DeleteRoleRequest struct {
	// The name of the RAM role.
	//
	// example:
	//
	// ECSAdmin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s DeleteRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRoleRequest) SetRoleName(v string) *DeleteRoleRequest {
	s.RoleName = &v
	return s
}

type DeleteRoleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 898FAB24-7509-43EE-A287-086FE4C44394
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRoleResponseBody) SetRequestId(v string) *DeleteRoleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRoleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoleResponse) GoString() string {
	return s.String()
}

func (s *DeleteRoleResponse) SetHeaders(v map[string]*string) *DeleteRoleResponse {
	s.Headers = v
	return s
}

func (s *DeleteRoleResponse) SetStatusCode(v int32) *DeleteRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRoleResponse) SetBody(v *DeleteRoleResponseBody) *DeleteRoleResponse {
	s.Body = v
	return s
}

type DeleteUserRequest struct {
	// The name of the RAM user.
	//
	// The name must be 1 to 64 characters in length, and can contain letters, digits, periods (.), hyphens (-), and underscores (_).
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DeleteUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserRequest) SetUserName(v string) *DeleteUserRequest {
	s.UserName = &v
	return s
}

type DeleteUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1C488B66-B819-4D14-8711-C4EAAA13AC01
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserResponseBody) SetRequestId(v string) *DeleteUserResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserResponse) SetHeaders(v map[string]*string) *DeleteUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserResponse) SetStatusCode(v int32) *DeleteUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserResponse) SetBody(v *DeleteUserResponseBody) *DeleteUserResponse {
	s.Body = v
	return s
}

type DeleteVirtualMFADeviceRequest struct {
	// The serial number of the MFA device.
	//
	// example:
	//
	// acs:ram::123456789012****:mfa/device002
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s DeleteVirtualMFADeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteVirtualMFADeviceRequest) SetSerialNumber(v string) *DeleteVirtualMFADeviceRequest {
	s.SerialNumber = &v
	return s
}

type DeleteVirtualMFADeviceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVirtualMFADeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualMFADeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVirtualMFADeviceResponseBody) SetRequestId(v string) *DeleteVirtualMFADeviceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVirtualMFADeviceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVirtualMFADeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVirtualMFADeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *DeleteVirtualMFADeviceResponse) SetHeaders(v map[string]*string) *DeleteVirtualMFADeviceResponse {
	s.Headers = v
	return s
}

func (s *DeleteVirtualMFADeviceResponse) SetStatusCode(v int32) *DeleteVirtualMFADeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVirtualMFADeviceResponse) SetBody(v *DeleteVirtualMFADeviceResponseBody) *DeleteVirtualMFADeviceResponse {
	s.Body = v
	return s
}

type DetachPolicyFromGroupRequest struct {
	// The name of the RAM user group.
	//
	// example:
	//
	// dev
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy. Valid values: `System` and `Custom`.
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DetachPolicyFromGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachPolicyFromGroupRequest) GoString() string {
	return s.String()
}

func (s *DetachPolicyFromGroupRequest) SetGroupName(v string) *DetachPolicyFromGroupRequest {
	s.GroupName = &v
	return s
}

func (s *DetachPolicyFromGroupRequest) SetPolicyName(v string) *DetachPolicyFromGroupRequest {
	s.PolicyName = &v
	return s
}

func (s *DetachPolicyFromGroupRequest) SetPolicyType(v string) *DetachPolicyFromGroupRequest {
	s.PolicyType = &v
	return s
}

type DetachPolicyFromGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 697852FB-50D7-44D9-9774-530C31EAC572
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachPolicyFromGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachPolicyFromGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DetachPolicyFromGroupResponseBody) SetRequestId(v string) *DetachPolicyFromGroupResponseBody {
	s.RequestId = &v
	return s
}

type DetachPolicyFromGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachPolicyFromGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachPolicyFromGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachPolicyFromGroupResponse) GoString() string {
	return s.String()
}

func (s *DetachPolicyFromGroupResponse) SetHeaders(v map[string]*string) *DetachPolicyFromGroupResponse {
	s.Headers = v
	return s
}

func (s *DetachPolicyFromGroupResponse) SetStatusCode(v int32) *DetachPolicyFromGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachPolicyFromGroupResponse) SetBody(v *DetachPolicyFromGroupResponseBody) *DetachPolicyFromGroupResponse {
	s.Body = v
	return s
}

type DetachPolicyFromRoleRequest struct {
	// The name of the policy.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy. Valid values: `System` and `Custom`.
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The name of the RAM role.
	//
	// example:
	//
	// OSSAdminRole
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s DetachPolicyFromRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachPolicyFromRoleRequest) GoString() string {
	return s.String()
}

func (s *DetachPolicyFromRoleRequest) SetPolicyName(v string) *DetachPolicyFromRoleRequest {
	s.PolicyName = &v
	return s
}

func (s *DetachPolicyFromRoleRequest) SetPolicyType(v string) *DetachPolicyFromRoleRequest {
	s.PolicyType = &v
	return s
}

func (s *DetachPolicyFromRoleRequest) SetRoleName(v string) *DetachPolicyFromRoleRequest {
	s.RoleName = &v
	return s
}

type DetachPolicyFromRoleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 697852FB-50D7-44D9-9774-530C31EAC572
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachPolicyFromRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachPolicyFromRoleResponseBody) GoString() string {
	return s.String()
}

func (s *DetachPolicyFromRoleResponseBody) SetRequestId(v string) *DetachPolicyFromRoleResponseBody {
	s.RequestId = &v
	return s
}

type DetachPolicyFromRoleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachPolicyFromRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachPolicyFromRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachPolicyFromRoleResponse) GoString() string {
	return s.String()
}

func (s *DetachPolicyFromRoleResponse) SetHeaders(v map[string]*string) *DetachPolicyFromRoleResponse {
	s.Headers = v
	return s
}

func (s *DetachPolicyFromRoleResponse) SetStatusCode(v int32) *DetachPolicyFromRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachPolicyFromRoleResponse) SetBody(v *DetachPolicyFromRoleResponseBody) *DetachPolicyFromRoleResponse {
	s.Body = v
	return s
}

type DetachPolicyFromUserRequest struct {
	// The name of the policy.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy. Valid values: `System` and `Custom`.
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The name of the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DetachPolicyFromUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachPolicyFromUserRequest) GoString() string {
	return s.String()
}

func (s *DetachPolicyFromUserRequest) SetPolicyName(v string) *DetachPolicyFromUserRequest {
	s.PolicyName = &v
	return s
}

func (s *DetachPolicyFromUserRequest) SetPolicyType(v string) *DetachPolicyFromUserRequest {
	s.PolicyType = &v
	return s
}

func (s *DetachPolicyFromUserRequest) SetUserName(v string) *DetachPolicyFromUserRequest {
	s.UserName = &v
	return s
}

type DetachPolicyFromUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 697852FB-50D7-44D9-9774-530C31EAC572
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachPolicyFromUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachPolicyFromUserResponseBody) GoString() string {
	return s.String()
}

func (s *DetachPolicyFromUserResponseBody) SetRequestId(v string) *DetachPolicyFromUserResponseBody {
	s.RequestId = &v
	return s
}

type DetachPolicyFromUserResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachPolicyFromUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachPolicyFromUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachPolicyFromUserResponse) GoString() string {
	return s.String()
}

func (s *DetachPolicyFromUserResponse) SetHeaders(v map[string]*string) *DetachPolicyFromUserResponse {
	s.Headers = v
	return s
}

func (s *DetachPolicyFromUserResponse) SetStatusCode(v int32) *DetachPolicyFromUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachPolicyFromUserResponse) SetBody(v *DetachPolicyFromUserResponseBody) *DetachPolicyFromUserResponse {
	s.Body = v
	return s
}

type GetAccessKeyLastUsedRequest struct {
	// example:
	//
	// LTAI4GFTgcR8m8cZQDTH****
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// example:
	//
	// test
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetAccessKeyLastUsedRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedRequest) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedRequest) SetUserAccessKeyId(v string) *GetAccessKeyLastUsedRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *GetAccessKeyLastUsedRequest) SetUserName(v string) *GetAccessKeyLastUsedRequest {
	s.UserName = &v
	return s
}

type GetAccessKeyLastUsedResponseBody struct {
	// The details of the time when the AccessKey pair was used for the last time.
	AccessKeyLastUsed *GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed `json:"AccessKeyLastUsed,omitempty" xml:"AccessKeyLastUsed,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 5CCE804C-6450-49A7-B1DB-2460F7A97416
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccessKeyLastUsedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedResponseBody) SetAccessKeyLastUsed(v *GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed) *GetAccessKeyLastUsedResponseBody {
	s.AccessKeyLastUsed = v
	return s
}

func (s *GetAccessKeyLastUsedResponseBody) SetRequestId(v string) *GetAccessKeyLastUsedResponseBody {
	s.RequestId = &v
	return s
}

type GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed struct {
	// The time when the AccessKey pair was used for the last time.
	//
	// example:
	//
	// 2020-10-21T06:37:40Z
	LastUsedDate *string `json:"LastUsedDate,omitempty" xml:"LastUsedDate,omitempty"`
}

func (s GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed) SetLastUsedDate(v string) *GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed {
	s.LastUsedDate = &v
	return s
}

type GetAccessKeyLastUsedResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccessKeyLastUsedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccessKeyLastUsedResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedResponse) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedResponse) SetHeaders(v map[string]*string) *GetAccessKeyLastUsedResponse {
	s.Headers = v
	return s
}

func (s *GetAccessKeyLastUsedResponse) SetStatusCode(v int32) *GetAccessKeyLastUsedResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessKeyLastUsedResponse) SetBody(v *GetAccessKeyLastUsedResponseBody) *GetAccessKeyLastUsedResponse {
	s.Body = v
	return s
}

type GetAccountAliasResponseBody struct {
	// The alias of the Alibaba Cloud account.
	//
	// example:
	//
	// myalias
	AccountAlias *string `json:"AccountAlias,omitempty" xml:"AccountAlias,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccountAliasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccountAliasResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountAliasResponseBody) SetAccountAlias(v string) *GetAccountAliasResponseBody {
	s.AccountAlias = &v
	return s
}

func (s *GetAccountAliasResponseBody) SetRequestId(v string) *GetAccountAliasResponseBody {
	s.RequestId = &v
	return s
}

type GetAccountAliasResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountAliasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountAliasResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccountAliasResponse) GoString() string {
	return s.String()
}

func (s *GetAccountAliasResponse) SetHeaders(v map[string]*string) *GetAccountAliasResponse {
	s.Headers = v
	return s
}

func (s *GetAccountAliasResponse) SetStatusCode(v int32) *GetAccountAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountAliasResponse) SetBody(v *GetAccountAliasResponseBody) *GetAccountAliasResponse {
	s.Body = v
	return s
}

type GetGroupRequest struct {
	// The name of the RAM user group.
	//
	// example:
	//
	// Dev-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s GetGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGroupRequest) GoString() string {
	return s.String()
}

func (s *GetGroupRequest) SetGroupName(v string) *GetGroupRequest {
	s.GroupName = &v
	return s
}

type GetGroupResponseBody struct {
	// The information about the RAM user group.
	Group *GetGroupResponseBodyGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D4065824-E422-3ED6-68B1-1AF7D5C7804C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupResponseBody) SetGroup(v *GetGroupResponseBodyGroup) *GetGroupResponseBody {
	s.Group = v
	return s
}

func (s *GetGroupResponseBody) SetRequestId(v string) *GetGroupResponseBody {
	s.RequestId = &v
	return s
}

type GetGroupResponseBodyGroup struct {
	// The description of the RAM user group.
	//
	// example:
	//
	// Dev-Team
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The time when the RAM user group was created.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The ID of the RAM user group.
	//
	// example:
	//
	// g-FpMEHiMysofp****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the RAM user group.
	//
	// example:
	//
	// Dev-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2015-02-11T03:15:21Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s GetGroupResponseBodyGroup) String() string {
	return tea.Prettify(s)
}

func (s GetGroupResponseBodyGroup) GoString() string {
	return s.String()
}

func (s *GetGroupResponseBodyGroup) SetComments(v string) *GetGroupResponseBodyGroup {
	s.Comments = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetCreateDate(v string) *GetGroupResponseBodyGroup {
	s.CreateDate = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetGroupId(v string) *GetGroupResponseBodyGroup {
	s.GroupId = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetGroupName(v string) *GetGroupResponseBodyGroup {
	s.GroupName = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetUpdateDate(v string) *GetGroupResponseBodyGroup {
	s.UpdateDate = &v
	return s
}

type GetGroupResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGroupResponse) GoString() string {
	return s.String()
}

func (s *GetGroupResponse) SetHeaders(v map[string]*string) *GetGroupResponse {
	s.Headers = v
	return s
}

func (s *GetGroupResponse) SetStatusCode(v int32) *GetGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGroupResponse) SetBody(v *GetGroupResponseBody) *GetGroupResponse {
	s.Body = v
	return s
}

type GetLoginProfileRequest struct {
	// The name of the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetLoginProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLoginProfileRequest) GoString() string {
	return s.String()
}

func (s *GetLoginProfileRequest) SetUserName(v string) *GetLoginProfileRequest {
	s.UserName = &v
	return s
}

type GetLoginProfileResponseBody struct {
	// The logon configurations of the RAM user.
	LoginProfile *GetLoginProfileResponseBodyLoginProfile `json:"LoginProfile,omitempty" xml:"LoginProfile,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLoginProfileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLoginProfileResponseBody) GoString() string {
	return s.String()
}

func (s *GetLoginProfileResponseBody) SetLoginProfile(v *GetLoginProfileResponseBodyLoginProfile) *GetLoginProfileResponseBody {
	s.LoginProfile = v
	return s
}

func (s *GetLoginProfileResponseBody) SetRequestId(v string) *GetLoginProfileResponseBody {
	s.RequestId = &v
	return s
}

type GetLoginProfileResponseBodyLoginProfile struct {
	// The creation time.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// Indicates whether a multi-factor authentication (MFA) device must be bound to the RAM user.
	//
	// example:
	//
	// true
	MFABindRequired *bool `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty"`
	// Indicates whether the RAM user must change the password upon logon.
	//
	// example:
	//
	// true
	PasswordResetRequired *bool `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty"`
	// The name of the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetLoginProfileResponseBodyLoginProfile) String() string {
	return tea.Prettify(s)
}

func (s GetLoginProfileResponseBodyLoginProfile) GoString() string {
	return s.String()
}

func (s *GetLoginProfileResponseBodyLoginProfile) SetCreateDate(v string) *GetLoginProfileResponseBodyLoginProfile {
	s.CreateDate = &v
	return s
}

func (s *GetLoginProfileResponseBodyLoginProfile) SetMFABindRequired(v bool) *GetLoginProfileResponseBodyLoginProfile {
	s.MFABindRequired = &v
	return s
}

func (s *GetLoginProfileResponseBodyLoginProfile) SetPasswordResetRequired(v bool) *GetLoginProfileResponseBodyLoginProfile {
	s.PasswordResetRequired = &v
	return s
}

func (s *GetLoginProfileResponseBodyLoginProfile) SetUserName(v string) *GetLoginProfileResponseBodyLoginProfile {
	s.UserName = &v
	return s
}

type GetLoginProfileResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLoginProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLoginProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLoginProfileResponse) GoString() string {
	return s.String()
}

func (s *GetLoginProfileResponse) SetHeaders(v map[string]*string) *GetLoginProfileResponse {
	s.Headers = v
	return s
}

func (s *GetLoginProfileResponse) SetStatusCode(v int32) *GetLoginProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLoginProfileResponse) SetBody(v *GetLoginProfileResponseBody) *GetLoginProfileResponse {
	s.Body = v
	return s
}

type GetPasswordPolicyResponseBody struct {
	// The password policy.
	PasswordPolicy *GetPasswordPolicyResponseBodyPasswordPolicy `json:"PasswordPolicy,omitempty" xml:"PasswordPolicy,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPasswordPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPasswordPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetPasswordPolicyResponseBody) SetPasswordPolicy(v *GetPasswordPolicyResponseBodyPasswordPolicy) *GetPasswordPolicyResponseBody {
	s.PasswordPolicy = v
	return s
}

func (s *GetPasswordPolicyResponseBody) SetRequestId(v string) *GetPasswordPolicyResponseBody {
	s.RequestId = &v
	return s
}

type GetPasswordPolicyResponseBodyPasswordPolicy struct {
	// Indicates whether a password expires. Valid values: `true` and `false`. Default value: `false`. If the parameter is unspecified, the default value false is returned.
	//
	// 	- If this parameter is set to `true`, the Alibaba Cloud account to which the RAM users belong must reset the password before the RAM users can log on to the Alibaba Cloud Management Console.
	//
	// 	- If this parameter is set to `false`, the RAM users can change the passwords after the passwords expire and then log on to the Alibaba Cloud Management Console.
	//
	// example:
	//
	// false
	HardExpiry *bool `json:"HardExpiry,omitempty" xml:"HardExpiry,omitempty"`
	// The maximum number of permitted logon attempts within one hour. The number of logon attempts is reset to zero if a RAM user changes the password.
	//
	// example:
	//
	// 5
	MaxLoginAttemps *int32 `json:"MaxLoginAttemps,omitempty" xml:"MaxLoginAttemps,omitempty"`
	// The number of days for which a password is valid. If you reset a password, the password validity period restarts. Default value: 0. The default value indicates that the password never expires.
	//
	// example:
	//
	// 0
	MaxPasswordAge *int32 `json:"MaxPasswordAge,omitempty" xml:"MaxPasswordAge,omitempty"`
	// The minimum number of characters in a password.
	//
	// example:
	//
	// 12
	MinimumPasswordLength *int32 `json:"MinimumPasswordLength,omitempty" xml:"MinimumPasswordLength,omitempty"`
	// The number of previous passwords that a RAM user is prevented from reusing. Default value: 0. The default value indicates that the RAM user can reuse previous passwords.
	//
	// example:
	//
	// 0
	PasswordReusePrevention *int32 `json:"PasswordReusePrevention,omitempty" xml:"PasswordReusePrevention,omitempty"`
	// Indicates whether a password must contain one or more lowercase letters.
	//
	// example:
	//
	// true
	RequireLowercaseCharacters *bool `json:"RequireLowercaseCharacters,omitempty" xml:"RequireLowercaseCharacters,omitempty"`
	// Indicates whether a password must contain one or more digits.
	//
	// example:
	//
	// true
	RequireNumbers *bool `json:"RequireNumbers,omitempty" xml:"RequireNumbers,omitempty"`
	// Indicates whether a password must contain one or more special characters.
	//
	// example:
	//
	// true
	RequireSymbols *bool `json:"RequireSymbols,omitempty" xml:"RequireSymbols,omitempty"`
	// Indicates whether a password must contain one or more uppercase letters.
	//
	// example:
	//
	// true
	RequireUppercaseCharacters *bool `json:"RequireUppercaseCharacters,omitempty" xml:"RequireUppercaseCharacters,omitempty"`
}

func (s GetPasswordPolicyResponseBodyPasswordPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetPasswordPolicyResponseBodyPasswordPolicy) GoString() string {
	return s.String()
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetHardExpiry(v bool) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.HardExpiry = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetMaxLoginAttemps(v int32) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.MaxLoginAttemps = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetMaxPasswordAge(v int32) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.MaxPasswordAge = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetMinimumPasswordLength(v int32) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.MinimumPasswordLength = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetPasswordReusePrevention(v int32) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.PasswordReusePrevention = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetRequireLowercaseCharacters(v bool) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.RequireLowercaseCharacters = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetRequireNumbers(v bool) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.RequireNumbers = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetRequireSymbols(v bool) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.RequireSymbols = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetRequireUppercaseCharacters(v bool) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.RequireUppercaseCharacters = &v
	return s
}

type GetPasswordPolicyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPasswordPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPasswordPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPasswordPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetPasswordPolicyResponse) SetHeaders(v map[string]*string) *GetPasswordPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetPasswordPolicyResponse) SetStatusCode(v int32) *GetPasswordPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPasswordPolicyResponse) SetBody(v *GetPasswordPolicyResponseBody) *GetPasswordPolicyResponse {
	s.Body = v
	return s
}

type GetPolicyRequest struct {
	// The name of the policy.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy. Valid values: `System` and `Custom`.
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s GetPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetPolicyRequest) SetPolicyName(v string) *GetPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *GetPolicyRequest) SetPolicyType(v string) *GetPolicyRequest {
	s.PolicyType = &v
	return s
}

type GetPolicyResponseBody struct {
	// The information about the default policy version.
	DefaultPolicyVersion *GetPolicyResponseBodyDefaultPolicyVersion `json:"DefaultPolicyVersion,omitempty" xml:"DefaultPolicyVersion,omitempty" type:"Struct"`
	// The basic information about the policy.
	Policy *GetPolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 697852FB-50D7-44D9-9774-530C31EAC572
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetPolicyResponseBody) SetDefaultPolicyVersion(v *GetPolicyResponseBodyDefaultPolicyVersion) *GetPolicyResponseBody {
	s.DefaultPolicyVersion = v
	return s
}

func (s *GetPolicyResponseBody) SetPolicy(v *GetPolicyResponseBodyPolicy) *GetPolicyResponseBody {
	s.Policy = v
	return s
}

func (s *GetPolicyResponseBody) SetRequestId(v string) *GetPolicyResponseBody {
	s.RequestId = &v
	return s
}

type GetPolicyResponseBodyDefaultPolicyVersion struct {
	// The time when the default policy version was created.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// An attribute in the `DefaultPolicyVersion` parameter. The value of the `IsDefaultVersion` parameter is `true`.
	//
	// example:
	//
	// true
	IsDefaultVersion *bool `json:"IsDefaultVersion,omitempty" xml:"IsDefaultVersion,omitempty"`
	// The document of the policy.
	//
	// example:
	//
	// { "Statement": [{ "Action": ["oss:*"], "Effect": "Allow", "Resource": ["acs:oss:*:*:*"]}], "Version": "1"}
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The ID of the default policy version.
	//
	// example:
	//
	// v1
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetPolicyResponseBodyDefaultPolicyVersion) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyResponseBodyDefaultPolicyVersion) GoString() string {
	return s.String()
}

func (s *GetPolicyResponseBodyDefaultPolicyVersion) SetCreateDate(v string) *GetPolicyResponseBodyDefaultPolicyVersion {
	s.CreateDate = &v
	return s
}

func (s *GetPolicyResponseBodyDefaultPolicyVersion) SetIsDefaultVersion(v bool) *GetPolicyResponseBodyDefaultPolicyVersion {
	s.IsDefaultVersion = &v
	return s
}

func (s *GetPolicyResponseBodyDefaultPolicyVersion) SetPolicyDocument(v string) *GetPolicyResponseBodyDefaultPolicyVersion {
	s.PolicyDocument = &v
	return s
}

func (s *GetPolicyResponseBodyDefaultPolicyVersion) SetVersionId(v string) *GetPolicyResponseBodyDefaultPolicyVersion {
	s.VersionId = &v
	return s
}

type GetPolicyResponseBodyPolicy struct {
	// The number of references to the policy.
	//
	// example:
	//
	// 0
	AttachmentCount *int32 `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty"`
	// The time when the policy was created.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The default version of the policy.
	//
	// example:
	//
	// v1
	DefaultVersion *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// The description of the policy.
	//
	// example:
	//
	// OSS administrator
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// N/A
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy.
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The time when the policy was modified.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s GetPolicyResponseBodyPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *GetPolicyResponseBodyPolicy) SetAttachmentCount(v int32) *GetPolicyResponseBodyPolicy {
	s.AttachmentCount = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetCreateDate(v string) *GetPolicyResponseBodyPolicy {
	s.CreateDate = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetDefaultVersion(v string) *GetPolicyResponseBodyPolicy {
	s.DefaultVersion = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetDescription(v string) *GetPolicyResponseBodyPolicy {
	s.Description = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetPolicyDocument(v string) *GetPolicyResponseBodyPolicy {
	s.PolicyDocument = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetPolicyName(v string) *GetPolicyResponseBodyPolicy {
	s.PolicyName = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetPolicyType(v string) *GetPolicyResponseBodyPolicy {
	s.PolicyType = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetUpdateDate(v string) *GetPolicyResponseBodyPolicy {
	s.UpdateDate = &v
	return s
}

type GetPolicyResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetPolicyResponse) SetHeaders(v map[string]*string) *GetPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetPolicyResponse) SetStatusCode(v int32) *GetPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPolicyResponse) SetBody(v *GetPolicyResponseBody) *GetPolicyResponse {
	s.Body = v
	return s
}

type GetPolicyVersionRequest struct {
	// The name of the policy.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy. Valid values: `System` and `Custom`.
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The ID of the policy version.
	//
	// example:
	//
	// v3
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetPolicyVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyVersionRequest) GoString() string {
	return s.String()
}

func (s *GetPolicyVersionRequest) SetPolicyName(v string) *GetPolicyVersionRequest {
	s.PolicyName = &v
	return s
}

func (s *GetPolicyVersionRequest) SetPolicyType(v string) *GetPolicyVersionRequest {
	s.PolicyType = &v
	return s
}

func (s *GetPolicyVersionRequest) SetVersionId(v string) *GetPolicyVersionRequest {
	s.VersionId = &v
	return s
}

type GetPolicyVersionResponseBody struct {
	// The information about the policy version.
	PolicyVersion *GetPolicyVersionResponseBodyPolicyVersion `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPolicyVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetPolicyVersionResponseBody) SetPolicyVersion(v *GetPolicyVersionResponseBodyPolicyVersion) *GetPolicyVersionResponseBody {
	s.PolicyVersion = v
	return s
}

func (s *GetPolicyVersionResponseBody) SetRequestId(v string) *GetPolicyVersionResponseBody {
	s.RequestId = &v
	return s
}

type GetPolicyVersionResponseBodyPolicyVersion struct {
	// The creation time.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// Indicates whether the policy version is the default version.
	//
	// example:
	//
	// false
	IsDefaultVersion *bool `json:"IsDefaultVersion,omitempty" xml:"IsDefaultVersion,omitempty"`
	// The document of the policy.
	//
	// example:
	//
	// { "Statement": [{ "Action": ["oss:*"], "Effect": "Allow", "Resource": ["acs:oss:*:*:*"]}], "Version": "1"}
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The ID of the policy version.
	//
	// example:
	//
	// v3
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetPolicyVersionResponseBodyPolicyVersion) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyVersionResponseBodyPolicyVersion) GoString() string {
	return s.String()
}

func (s *GetPolicyVersionResponseBodyPolicyVersion) SetCreateDate(v string) *GetPolicyVersionResponseBodyPolicyVersion {
	s.CreateDate = &v
	return s
}

func (s *GetPolicyVersionResponseBodyPolicyVersion) SetIsDefaultVersion(v bool) *GetPolicyVersionResponseBodyPolicyVersion {
	s.IsDefaultVersion = &v
	return s
}

func (s *GetPolicyVersionResponseBodyPolicyVersion) SetPolicyDocument(v string) *GetPolicyVersionResponseBodyPolicyVersion {
	s.PolicyDocument = &v
	return s
}

func (s *GetPolicyVersionResponseBodyPolicyVersion) SetVersionId(v string) *GetPolicyVersionResponseBodyPolicyVersion {
	s.VersionId = &v
	return s
}

type GetPolicyVersionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPolicyVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPolicyVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyVersionResponse) GoString() string {
	return s.String()
}

func (s *GetPolicyVersionResponse) SetHeaders(v map[string]*string) *GetPolicyVersionResponse {
	s.Headers = v
	return s
}

func (s *GetPolicyVersionResponse) SetStatusCode(v int32) *GetPolicyVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPolicyVersionResponse) SetBody(v *GetPolicyVersionResponseBody) *GetPolicyVersionResponse {
	s.Body = v
	return s
}

type GetRoleRequest struct {
	// The name of the RAM role.
	//
	// The name must be 1 to 64 characters in length, and can contain letters, digits, periods (.), and hyphens (-).
	//
	// example:
	//
	// ECSAdmin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s GetRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRoleRequest) GoString() string {
	return s.String()
}

func (s *GetRoleRequest) SetRoleName(v string) *GetRoleRequest {
	s.RoleName = &v
	return s
}

type GetRoleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the RAM role.
	Role *GetRoleResponseBodyRole `json:"Role,omitempty" xml:"Role,omitempty" type:"Struct"`
}

func (s GetRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRoleResponseBody) GoString() string {
	return s.String()
}

func (s *GetRoleResponseBody) SetRequestId(v string) *GetRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRoleResponseBody) SetRole(v *GetRoleResponseBodyRole) *GetRoleResponseBody {
	s.Role = v
	return s
}

type GetRoleResponseBodyRole struct {
	// The Alibaba Cloud Resource Name (ARN) of the RAM role.
	//
	// example:
	//
	// acs:ram::123456789012****:role/ECSAdmin
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The policy that specifies the trusted entity to assume the RAM role.
	//
	// example:
	//
	// { "Statement": [ { "Action": "sts:AssumeRole", "Effect": "Allow", "Principal": { "RAM": "acs:ram::123456789012****:root" } } ], "Version": "1" }
	AssumeRolePolicyDocument *string `json:"AssumeRolePolicyDocument,omitempty" xml:"AssumeRolePolicyDocument,omitempty"`
	// The time when the RAM role was created.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the RAM role.
	//
	// example:
	//
	// ECS administrator
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum session duration of the RAM role.
	//
	// example:
	//
	// 3600
	MaxSessionDuration *int64 `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty"`
	// The ID of the RAM role.
	//
	// example:
	//
	// 901234567890****
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// The name of the RAM role.
	//
	// example:
	//
	// ECSAdmin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The time when the RAM role was modified.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s GetRoleResponseBodyRole) String() string {
	return tea.Prettify(s)
}

func (s GetRoleResponseBodyRole) GoString() string {
	return s.String()
}

func (s *GetRoleResponseBodyRole) SetArn(v string) *GetRoleResponseBodyRole {
	s.Arn = &v
	return s
}

func (s *GetRoleResponseBodyRole) SetAssumeRolePolicyDocument(v string) *GetRoleResponseBodyRole {
	s.AssumeRolePolicyDocument = &v
	return s
}

func (s *GetRoleResponseBodyRole) SetCreateDate(v string) *GetRoleResponseBodyRole {
	s.CreateDate = &v
	return s
}

func (s *GetRoleResponseBodyRole) SetDescription(v string) *GetRoleResponseBodyRole {
	s.Description = &v
	return s
}

func (s *GetRoleResponseBodyRole) SetMaxSessionDuration(v int64) *GetRoleResponseBodyRole {
	s.MaxSessionDuration = &v
	return s
}

func (s *GetRoleResponseBodyRole) SetRoleId(v string) *GetRoleResponseBodyRole {
	s.RoleId = &v
	return s
}

func (s *GetRoleResponseBodyRole) SetRoleName(v string) *GetRoleResponseBodyRole {
	s.RoleName = &v
	return s
}

func (s *GetRoleResponseBodyRole) SetUpdateDate(v string) *GetRoleResponseBodyRole {
	s.UpdateDate = &v
	return s
}

type GetRoleResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRoleResponse) GoString() string {
	return s.String()
}

func (s *GetRoleResponse) SetHeaders(v map[string]*string) *GetRoleResponse {
	s.Headers = v
	return s
}

func (s *GetRoleResponse) SetStatusCode(v int32) *GetRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRoleResponse) SetBody(v *GetRoleResponseBody) *GetRoleResponse {
	s.Body = v
	return s
}

type GetSecurityPreferenceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DC1213F1-A9D5-4A01-A996-44983689126C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The security preferences.
	SecurityPreference *GetSecurityPreferenceResponseBodySecurityPreference `json:"SecurityPreference,omitempty" xml:"SecurityPreference,omitempty" type:"Struct"`
}

func (s GetSecurityPreferenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityPreferenceResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseBody) SetRequestId(v string) *GetSecurityPreferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecurityPreferenceResponseBody) SetSecurityPreference(v *GetSecurityPreferenceResponseBodySecurityPreference) *GetSecurityPreferenceResponseBody {
	s.SecurityPreference = v
	return s
}

type GetSecurityPreferenceResponseBodySecurityPreference struct {
	// The AccessKey pair preference.
	AccessKeyPreference *GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference `json:"AccessKeyPreference,omitempty" xml:"AccessKeyPreference,omitempty" type:"Struct"`
	// The logon preference.
	LoginProfilePreference *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference `json:"LoginProfilePreference,omitempty" xml:"LoginProfilePreference,omitempty" type:"Struct"`
	// The multi-factor authentication (MFA) preference.
	MFAPreference *GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference `json:"MFAPreference,omitempty" xml:"MFAPreference,omitempty" type:"Struct"`
	// The public key preference.
	//
	// >  This parameter is valid only for the Japan site.
	PublicKeyPreference *GetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference `json:"PublicKeyPreference,omitempty" xml:"PublicKeyPreference,omitempty" type:"Struct"`
}

func (s GetSecurityPreferenceResponseBodySecurityPreference) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityPreferenceResponseBodySecurityPreference) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) SetAccessKeyPreference(v *GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) *GetSecurityPreferenceResponseBodySecurityPreference {
	s.AccessKeyPreference = v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) SetLoginProfilePreference(v *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) *GetSecurityPreferenceResponseBodySecurityPreference {
	s.LoginProfilePreference = v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) SetMFAPreference(v *GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) *GetSecurityPreferenceResponseBodySecurityPreference {
	s.MFAPreference = v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) SetPublicKeyPreference(v *GetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference) *GetSecurityPreferenceResponseBodySecurityPreference {
	s.PublicKeyPreference = v
	return s
}

type GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference struct {
	// Indicates whether Resource Access Management (RAM) users can manage their AccessKey pairs. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	AllowUserToManageAccessKeys *bool `json:"AllowUserToManageAccessKeys,omitempty" xml:"AllowUserToManageAccessKeys,omitempty"`
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) SetAllowUserToManageAccessKeys(v bool) *GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference {
	s.AllowUserToManageAccessKeys = &v
	return s
}

type GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference struct {
	// Indicates whether RAM users can change their passwords. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AllowUserToChangePassword *bool `json:"AllowUserToChangePassword,omitempty" xml:"AllowUserToChangePassword,omitempty"`
	// Indicates whether RAM users can save security codes for MFA during logon. Each security code is valid for seven days. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	EnableSaveMFATicket *bool `json:"EnableSaveMFATicket,omitempty" xml:"EnableSaveMFATicket,omitempty"`
	// The subnet mask that indicates the IP addresses from which logon to the Alibaba Cloud Management Console is allowed. This parameter applies to password-based logon and single sign-on (SSO). However, this parameter does not apply to API calls that are authenticated based on AccessKey pairs.
	//
	// 	- If you specify a subnet mask, RAM users can use only the IP addresses in the subnet mask to log on to the Alibaba Cloud Management Console.
	//
	// 	- If you do not specify a subnet mask, RAM users can use all IP addresses to log on to the Alibaba Cloud Management Console.
	//
	// If you want to specify more than one subnet mask, separate the masks with semicolons (;). Example: 192.168.0.0/16;10.0.0.0/8.
	//
	// example:
	//
	// 10.0.0.0/8
	LoginNetworkMasks *string `json:"LoginNetworkMasks,omitempty" xml:"LoginNetworkMasks,omitempty"`
	// The validity period of the logon session of RAM users. Unit: hours.
	//
	// example:
	//
	// 6
	LoginSessionDuration *int32 `json:"LoginSessionDuration,omitempty" xml:"LoginSessionDuration,omitempty"`
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetAllowUserToChangePassword(v bool) *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.AllowUserToChangePassword = &v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetEnableSaveMFATicket(v bool) *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.EnableSaveMFATicket = &v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetLoginNetworkMasks(v string) *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.LoginNetworkMasks = &v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetLoginSessionDuration(v int32) *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.LoginSessionDuration = &v
	return s
}

type GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference struct {
	// Indicates whether RAM users can manage their MFA devices. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AllowUserToManageMFADevices *bool `json:"AllowUserToManageMFADevices,omitempty" xml:"AllowUserToManageMFADevices,omitempty"`
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) SetAllowUserToManageMFADevices(v bool) *GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference {
	s.AllowUserToManageMFADevices = &v
	return s
}

type GetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference struct {
	// Indicates whether RAM users can manage their public keys. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	AllowUserToManagePublicKeys *bool `json:"AllowUserToManagePublicKeys,omitempty" xml:"AllowUserToManagePublicKeys,omitempty"`
}

func (s GetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference) SetAllowUserToManagePublicKeys(v bool) *GetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference {
	s.AllowUserToManagePublicKeys = &v
	return s
}

type GetSecurityPreferenceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSecurityPreferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSecurityPreferenceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityPreferenceResponse) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponse) SetHeaders(v map[string]*string) *GetSecurityPreferenceResponse {
	s.Headers = v
	return s
}

func (s *GetSecurityPreferenceResponse) SetStatusCode(v int32) *GetSecurityPreferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecurityPreferenceResponse) SetBody(v *GetSecurityPreferenceResponseBody) *GetSecurityPreferenceResponse {
	s.Body = v
	return s
}

type GetUserRequest struct {
	// The username of the RAM user.
	//
	// The username must be 1 to 64 characters in length, and can contain letters, digits, periods (.), hyphens (-), and underscores (_).
	//
	// example:
	//
	// alice
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) SetUserName(v string) *GetUserRequest {
	s.UserName = &v
	return s
}

type GetUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2D69A58F-345C-4FDE-88E4-BF5189484043
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the RAM user.
	User *GetUserResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s GetUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserResponseBody) SetUser(v *GetUserResponseBodyUser) *GetUserResponseBody {
	s.User = v
	return s
}

type GetUserResponseBodyUser struct {
	// The description of the RAM user.
	//
	// example:
	//
	// Cloud computing engineer
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The point in time when the RAM user was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The display name of the RAM user.
	//
	// example:
	//
	// alice
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the RAM user.
	//
	// >  This parameter can be returned only on the China site (aliyun.com).
	//
	// example:
	//
	// alice@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The point in time when the RAM user last logged on to the Alibaba Cloud Management Console by using the password. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	LastLoginDate *string `json:"LastLoginDate,omitempty" xml:"LastLoginDate,omitempty"`
	// The mobile phone number of the RAM user.
	//
	// >  This parameter can be returned only on the China site (aliyun.com).
	//
	// example:
	//
	// 86-1860000****
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	// The point in time when the information about the RAM user was last modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-02-11T03:15:21Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	// The ID of the RAM user.
	//
	// example:
	//
	// 222748924538****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username of the RAM user.
	//
	// example:
	//
	// alice
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetUserResponseBodyUser) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyUser) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUser) SetComments(v string) *GetUserResponseBodyUser {
	s.Comments = &v
	return s
}

func (s *GetUserResponseBodyUser) SetCreateDate(v string) *GetUserResponseBodyUser {
	s.CreateDate = &v
	return s
}

func (s *GetUserResponseBodyUser) SetDisplayName(v string) *GetUserResponseBodyUser {
	s.DisplayName = &v
	return s
}

func (s *GetUserResponseBodyUser) SetEmail(v string) *GetUserResponseBodyUser {
	s.Email = &v
	return s
}

func (s *GetUserResponseBodyUser) SetLastLoginDate(v string) *GetUserResponseBodyUser {
	s.LastLoginDate = &v
	return s
}

func (s *GetUserResponseBodyUser) SetMobilePhone(v string) *GetUserResponseBodyUser {
	s.MobilePhone = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUpdateDate(v string) *GetUserResponseBodyUser {
	s.UpdateDate = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserId(v string) *GetUserResponseBodyUser {
	s.UserId = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserName(v string) *GetUserResponseBodyUser {
	s.UserName = &v
	return s
}

type GetUserResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponse) GoString() string {
	return s.String()
}

func (s *GetUserResponse) SetHeaders(v map[string]*string) *GetUserResponse {
	s.Headers = v
	return s
}

func (s *GetUserResponse) SetStatusCode(v int32) *GetUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserResponse) SetBody(v *GetUserResponseBody) *GetUserResponse {
	s.Body = v
	return s
}

type GetUserMFAInfoRequest struct {
	// The name of the RAM user.
	//
	// example:
	//
	// test
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetUserMFAInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserMFAInfoRequest) GoString() string {
	return s.String()
}

func (s *GetUserMFAInfoRequest) SetUserName(v string) *GetUserMFAInfoRequest {
	s.UserName = &v
	return s
}

type GetUserMFAInfoResponseBody struct {
	// The information about the MFA device that is bound to the RAM user.
	MFADevice *GetUserMFAInfoResponseBodyMFADevice `json:"MFADevice,omitempty" xml:"MFADevice,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUserMFAInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserMFAInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserMFAInfoResponseBody) SetMFADevice(v *GetUserMFAInfoResponseBodyMFADevice) *GetUserMFAInfoResponseBody {
	s.MFADevice = v
	return s
}

func (s *GetUserMFAInfoResponseBody) SetRequestId(v string) *GetUserMFAInfoResponseBody {
	s.RequestId = &v
	return s
}

type GetUserMFAInfoResponseBodyMFADevice struct {
	// The serial number of the MFA device.
	//
	// example:
	//
	// acs:ram::177242285274****:mfa/test
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The type of the MFA device. Valid values:
	//
	// 	- VMFA: virtual MFA device.
	//
	// 	- U2F: Universal 2nd Factor (U2F) security key.
	//
	// example:
	//
	// VMFA
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetUserMFAInfoResponseBodyMFADevice) String() string {
	return tea.Prettify(s)
}

func (s GetUserMFAInfoResponseBodyMFADevice) GoString() string {
	return s.String()
}

func (s *GetUserMFAInfoResponseBodyMFADevice) SetSerialNumber(v string) *GetUserMFAInfoResponseBodyMFADevice {
	s.SerialNumber = &v
	return s
}

func (s *GetUserMFAInfoResponseBodyMFADevice) SetType(v string) *GetUserMFAInfoResponseBodyMFADevice {
	s.Type = &v
	return s
}

type GetUserMFAInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserMFAInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserMFAInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserMFAInfoResponse) GoString() string {
	return s.String()
}

func (s *GetUserMFAInfoResponse) SetHeaders(v map[string]*string) *GetUserMFAInfoResponse {
	s.Headers = v
	return s
}

func (s *GetUserMFAInfoResponse) SetStatusCode(v int32) *GetUserMFAInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserMFAInfoResponse) SetBody(v *GetUserMFAInfoResponseBody) *GetUserMFAInfoResponse {
	s.Body = v
	return s
}

type ListAccessKeysRequest struct {
	// The name of the RAM user. If a RAM user calls this operation and does not specify this parameter, the AccessKey pairs of the RAM user are returned.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListAccessKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccessKeysRequest) GoString() string {
	return s.String()
}

func (s *ListAccessKeysRequest) SetUserName(v string) *ListAccessKeysRequest {
	s.UserName = &v
	return s
}

type ListAccessKeysResponseBody struct {
	// The AccessKey pairs that belong to the RAM user.
	AccessKeys *ListAccessKeysResponseBodyAccessKeys `json:"AccessKeys,omitempty" xml:"AccessKeys,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4B450CA1-36E8-4AA2-8461-86B42BF4CC4E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAccessKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAccessKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccessKeysResponseBody) SetAccessKeys(v *ListAccessKeysResponseBodyAccessKeys) *ListAccessKeysResponseBody {
	s.AccessKeys = v
	return s
}

func (s *ListAccessKeysResponseBody) SetRequestId(v string) *ListAccessKeysResponseBody {
	s.RequestId = &v
	return s
}

type ListAccessKeysResponseBodyAccessKeys struct {
	AccessKey []*ListAccessKeysResponseBodyAccessKeysAccessKey `json:"AccessKey,omitempty" xml:"AccessKey,omitempty" type:"Repeated"`
}

func (s ListAccessKeysResponseBodyAccessKeys) String() string {
	return tea.Prettify(s)
}

func (s ListAccessKeysResponseBodyAccessKeys) GoString() string {
	return s.String()
}

func (s *ListAccessKeysResponseBodyAccessKeys) SetAccessKey(v []*ListAccessKeysResponseBodyAccessKeysAccessKey) *ListAccessKeysResponseBodyAccessKeys {
	s.AccessKey = v
	return s
}

type ListAccessKeysResponseBodyAccessKeysAccessKey struct {
	// The AccessKey ID.
	//
	// example:
	//
	// 0wNEpMMlzy7s****
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The time when the AccessKey pair was created.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The status of the AccessKey pair. Valid values: Active and Inactive.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAccessKeysResponseBodyAccessKeysAccessKey) String() string {
	return tea.Prettify(s)
}

func (s ListAccessKeysResponseBodyAccessKeysAccessKey) GoString() string {
	return s.String()
}

func (s *ListAccessKeysResponseBodyAccessKeysAccessKey) SetAccessKeyId(v string) *ListAccessKeysResponseBodyAccessKeysAccessKey {
	s.AccessKeyId = &v
	return s
}

func (s *ListAccessKeysResponseBodyAccessKeysAccessKey) SetCreateDate(v string) *ListAccessKeysResponseBodyAccessKeysAccessKey {
	s.CreateDate = &v
	return s
}

func (s *ListAccessKeysResponseBodyAccessKeysAccessKey) SetStatus(v string) *ListAccessKeysResponseBodyAccessKeysAccessKey {
	s.Status = &v
	return s
}

type ListAccessKeysResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccessKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccessKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccessKeysResponse) GoString() string {
	return s.String()
}

func (s *ListAccessKeysResponse) SetHeaders(v map[string]*string) *ListAccessKeysResponse {
	s.Headers = v
	return s
}

func (s *ListAccessKeysResponse) SetStatusCode(v int32) *ListAccessKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccessKeysResponse) SetBody(v *ListAccessKeysResponseBody) *ListAccessKeysResponse {
	s.Body = v
	return s
}

type ListEntitiesForPolicyRequest struct {
	// The name of the policy.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy. Valid values: `System` and `Custom`.
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListEntitiesForPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEntitiesForPolicyRequest) GoString() string {
	return s.String()
}

func (s *ListEntitiesForPolicyRequest) SetPolicyName(v string) *ListEntitiesForPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *ListEntitiesForPolicyRequest) SetPolicyType(v string) *ListEntitiesForPolicyRequest {
	s.PolicyType = &v
	return s
}

type ListEntitiesForPolicyResponseBody struct {
	// The information about the Resource Access Management (RAM) user groups.
	Groups *ListEntitiesForPolicyResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7B8A4E7D-6CFF-471D-84DF-195A7A241ECB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the RAM roles.
	Roles *ListEntitiesForPolicyResponseBodyRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Struct"`
	// The information about the RAM users.
	Users *ListEntitiesForPolicyResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s ListEntitiesForPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEntitiesForPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ListEntitiesForPolicyResponseBody) SetGroups(v *ListEntitiesForPolicyResponseBodyGroups) *ListEntitiesForPolicyResponseBody {
	s.Groups = v
	return s
}

func (s *ListEntitiesForPolicyResponseBody) SetRequestId(v string) *ListEntitiesForPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEntitiesForPolicyResponseBody) SetRoles(v *ListEntitiesForPolicyResponseBodyRoles) *ListEntitiesForPolicyResponseBody {
	s.Roles = v
	return s
}

func (s *ListEntitiesForPolicyResponseBody) SetUsers(v *ListEntitiesForPolicyResponseBodyUsers) *ListEntitiesForPolicyResponseBody {
	s.Users = v
	return s
}

type ListEntitiesForPolicyResponseBodyGroups struct {
	Group []*ListEntitiesForPolicyResponseBodyGroupsGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Repeated"`
}

func (s ListEntitiesForPolicyResponseBodyGroups) String() string {
	return tea.Prettify(s)
}

func (s ListEntitiesForPolicyResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *ListEntitiesForPolicyResponseBodyGroups) SetGroup(v []*ListEntitiesForPolicyResponseBodyGroupsGroup) *ListEntitiesForPolicyResponseBodyGroups {
	s.Group = v
	return s
}

type ListEntitiesForPolicyResponseBodyGroupsGroup struct {
	// The time when the policy was attached to the RAM user group.
	//
	// example:
	//
	// 2015-02-18T17:22:08Z
	AttachDate *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty"`
	// The description of the RAM user group.
	//
	// example:
	//
	// Test team
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The name of the RAM user group.
	//
	// example:
	//
	// QA-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s ListEntitiesForPolicyResponseBodyGroupsGroup) String() string {
	return tea.Prettify(s)
}

func (s ListEntitiesForPolicyResponseBodyGroupsGroup) GoString() string {
	return s.String()
}

func (s *ListEntitiesForPolicyResponseBodyGroupsGroup) SetAttachDate(v string) *ListEntitiesForPolicyResponseBodyGroupsGroup {
	s.AttachDate = &v
	return s
}

func (s *ListEntitiesForPolicyResponseBodyGroupsGroup) SetComments(v string) *ListEntitiesForPolicyResponseBodyGroupsGroup {
	s.Comments = &v
	return s
}

func (s *ListEntitiesForPolicyResponseBodyGroupsGroup) SetGroupName(v string) *ListEntitiesForPolicyResponseBodyGroupsGroup {
	s.GroupName = &v
	return s
}

type ListEntitiesForPolicyResponseBodyRoles struct {
	Role []*ListEntitiesForPolicyResponseBodyRolesRole `json:"Role,omitempty" xml:"Role,omitempty" type:"Repeated"`
}

func (s ListEntitiesForPolicyResponseBodyRoles) String() string {
	return tea.Prettify(s)
}

func (s ListEntitiesForPolicyResponseBodyRoles) GoString() string {
	return s.String()
}

func (s *ListEntitiesForPolicyResponseBodyRoles) SetRole(v []*ListEntitiesForPolicyResponseBodyRolesRole) *ListEntitiesForPolicyResponseBodyRoles {
	s.Role = v
	return s
}

type ListEntitiesForPolicyResponseBodyRolesRole struct {
	// The Alibaba Cloud Resource Name (ARN) of the RAM role.
	//
	// example:
	//
	// acs:ram::123456789012****:role/ECSAdmin
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The time when the policy was attached to the RAM role.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	AttachDate *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty"`
	// The description of the RAM role.
	//
	// example:
	//
	// ECS administrator
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the RAM role.
	//
	// example:
	//
	// 122748924538****
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// The name of the RAM role.
	//
	// example:
	//
	// ECSAdmin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s ListEntitiesForPolicyResponseBodyRolesRole) String() string {
	return tea.Prettify(s)
}

func (s ListEntitiesForPolicyResponseBodyRolesRole) GoString() string {
	return s.String()
}

func (s *ListEntitiesForPolicyResponseBodyRolesRole) SetArn(v string) *ListEntitiesForPolicyResponseBodyRolesRole {
	s.Arn = &v
	return s
}

func (s *ListEntitiesForPolicyResponseBodyRolesRole) SetAttachDate(v string) *ListEntitiesForPolicyResponseBodyRolesRole {
	s.AttachDate = &v
	return s
}

func (s *ListEntitiesForPolicyResponseBodyRolesRole) SetDescription(v string) *ListEntitiesForPolicyResponseBodyRolesRole {
	s.Description = &v
	return s
}

func (s *ListEntitiesForPolicyResponseBodyRolesRole) SetRoleId(v string) *ListEntitiesForPolicyResponseBodyRolesRole {
	s.RoleId = &v
	return s
}

func (s *ListEntitiesForPolicyResponseBodyRolesRole) SetRoleName(v string) *ListEntitiesForPolicyResponseBodyRolesRole {
	s.RoleName = &v
	return s
}

type ListEntitiesForPolicyResponseBodyUsers struct {
	User []*ListEntitiesForPolicyResponseBodyUsersUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ListEntitiesForPolicyResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s ListEntitiesForPolicyResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListEntitiesForPolicyResponseBodyUsers) SetUser(v []*ListEntitiesForPolicyResponseBodyUsersUser) *ListEntitiesForPolicyResponseBodyUsers {
	s.User = v
	return s
}

type ListEntitiesForPolicyResponseBodyUsersUser struct {
	// The time when the policy was attached to the RAM user.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	AttachDate *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty"`
	// The display name of the RAM user.
	//
	// example:
	//
	// Zhang*
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The unique ID of the RAM user.
	//
	// example:
	//
	// 122748924538****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The name of the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListEntitiesForPolicyResponseBodyUsersUser) String() string {
	return tea.Prettify(s)
}

func (s ListEntitiesForPolicyResponseBodyUsersUser) GoString() string {
	return s.String()
}

func (s *ListEntitiesForPolicyResponseBodyUsersUser) SetAttachDate(v string) *ListEntitiesForPolicyResponseBodyUsersUser {
	s.AttachDate = &v
	return s
}

func (s *ListEntitiesForPolicyResponseBodyUsersUser) SetDisplayName(v string) *ListEntitiesForPolicyResponseBodyUsersUser {
	s.DisplayName = &v
	return s
}

func (s *ListEntitiesForPolicyResponseBodyUsersUser) SetUserId(v string) *ListEntitiesForPolicyResponseBodyUsersUser {
	s.UserId = &v
	return s
}

func (s *ListEntitiesForPolicyResponseBodyUsersUser) SetUserName(v string) *ListEntitiesForPolicyResponseBodyUsersUser {
	s.UserName = &v
	return s
}

type ListEntitiesForPolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEntitiesForPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEntitiesForPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEntitiesForPolicyResponse) GoString() string {
	return s.String()
}

func (s *ListEntitiesForPolicyResponse) SetHeaders(v map[string]*string) *ListEntitiesForPolicyResponse {
	s.Headers = v
	return s
}

func (s *ListEntitiesForPolicyResponse) SetStatusCode(v int32) *ListEntitiesForPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEntitiesForPolicyResponse) SetBody(v *ListEntitiesForPolicyResponseBody) *ListEntitiesForPolicyResponse {
	s.Body = v
	return s
}

type ListGroupsRequest struct {
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request.``
	//
	// When you call the operation for the first time, if the total number of returned entries exceeds the value of `MaxItems`, the entries are truncated. The system returns entries based on the value of `MaxItems` and does not return the excess entries. In this case, the value of the response parameter `IsTruncated` is `true`, and `Marker` is returned. In the next call, you can use the value of `Marker` and maintain the settings of the other request parameters to query the excess entries. You can repeat the call until the value of `IsTruncated` becomes `false`. This way, all entries are returned.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 100.
	//
	// example:
	//
	// 100
	MaxItems *int32 `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
}

func (s ListGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsRequest) SetMarker(v string) *ListGroupsRequest {
	s.Marker = &v
	return s
}

func (s *ListGroupsRequest) SetMaxItems(v int32) *ListGroupsRequest {
	s.MaxItems = &v
	return s
}

type ListGroupsResponseBody struct {
	// The information about the RAM user groups.
	Groups *ListGroupsResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Struct"`
	// Indicates whether the response is truncated. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// >  This parameter is returned only when `IsTruncated` is `true`.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 065527AA-2F2E-AD7C-7484-F2626CFE4934
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBody) SetGroups(v *ListGroupsResponseBodyGroups) *ListGroupsResponseBody {
	s.Groups = v
	return s
}

func (s *ListGroupsResponseBody) SetIsTruncated(v bool) *ListGroupsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListGroupsResponseBody) SetMarker(v string) *ListGroupsResponseBody {
	s.Marker = &v
	return s
}

func (s *ListGroupsResponseBody) SetRequestId(v string) *ListGroupsResponseBody {
	s.RequestId = &v
	return s
}

type ListGroupsResponseBodyGroups struct {
	Group []*ListGroupsResponseBodyGroupsGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Repeated"`
}

func (s ListGroupsResponseBodyGroups) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBodyGroups) SetGroup(v []*ListGroupsResponseBodyGroupsGroup) *ListGroupsResponseBodyGroups {
	s.Group = v
	return s
}

type ListGroupsResponseBodyGroupsGroup struct {
	// The description.
	//
	// example:
	//
	// Dev-Team
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The ID of the RAM user group.
	//
	// example:
	//
	// g-FpMEHiMysofp****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the RAM user group.
	//
	// example:
	//
	// Dev-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListGroupsResponseBodyGroupsGroup) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponseBodyGroupsGroup) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBodyGroupsGroup) SetComments(v string) *ListGroupsResponseBodyGroupsGroup {
	s.Comments = &v
	return s
}

func (s *ListGroupsResponseBodyGroupsGroup) SetCreateDate(v string) *ListGroupsResponseBodyGroupsGroup {
	s.CreateDate = &v
	return s
}

func (s *ListGroupsResponseBodyGroupsGroup) SetGroupId(v string) *ListGroupsResponseBodyGroupsGroup {
	s.GroupId = &v
	return s
}

func (s *ListGroupsResponseBodyGroupsGroup) SetGroupName(v string) *ListGroupsResponseBodyGroupsGroup {
	s.GroupName = &v
	return s
}

func (s *ListGroupsResponseBodyGroupsGroup) SetUpdateDate(v string) *ListGroupsResponseBodyGroupsGroup {
	s.UpdateDate = &v
	return s
}

type ListGroupsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListGroupsResponse) SetHeaders(v map[string]*string) *ListGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListGroupsResponse) SetStatusCode(v int32) *ListGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupsResponse) SetBody(v *ListGroupsResponseBody) *ListGroupsResponse {
	s.Body = v
	return s
}

type ListGroupsForUserRequest struct {
	// The name of the RAM user.
	//
	// example:
	//
	// Alice
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListGroupsForUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsForUserRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserRequest) SetUserName(v string) *ListGroupsForUserRequest {
	s.UserName = &v
	return s
}

type ListGroupsForUserResponseBody struct {
	// The information about the RAM user groups.
	Groups *ListGroupsForUserResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// DA772B52-BF9F-54CA-AC77-AA7A2DA89D46
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListGroupsForUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserResponseBody) SetGroups(v *ListGroupsForUserResponseBodyGroups) *ListGroupsForUserResponseBody {
	s.Groups = v
	return s
}

func (s *ListGroupsForUserResponseBody) SetRequestId(v string) *ListGroupsForUserResponseBody {
	s.RequestId = &v
	return s
}

type ListGroupsForUserResponseBodyGroups struct {
	Group []*ListGroupsForUserResponseBodyGroupsGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Repeated"`
}

func (s ListGroupsForUserResponseBodyGroups) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsForUserResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserResponseBodyGroups) SetGroup(v []*ListGroupsForUserResponseBodyGroupsGroup) *ListGroupsForUserResponseBodyGroups {
	s.Group = v
	return s
}

type ListGroupsForUserResponseBodyGroupsGroup struct {
	// The description.
	//
	// example:
	//
	// Dev-Team
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The ID of the RAM user group.
	//
	// example:
	//
	// g-zYtroLrgbZR1****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the RAM user group.
	//
	// example:
	//
	// Dev-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The time when the RAM user was added to the RAM user group.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	JoinDate *string `json:"JoinDate,omitempty" xml:"JoinDate,omitempty"`
}

func (s ListGroupsForUserResponseBodyGroupsGroup) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsForUserResponseBodyGroupsGroup) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserResponseBodyGroupsGroup) SetComments(v string) *ListGroupsForUserResponseBodyGroupsGroup {
	s.Comments = &v
	return s
}

func (s *ListGroupsForUserResponseBodyGroupsGroup) SetGroupId(v string) *ListGroupsForUserResponseBodyGroupsGroup {
	s.GroupId = &v
	return s
}

func (s *ListGroupsForUserResponseBodyGroupsGroup) SetGroupName(v string) *ListGroupsForUserResponseBodyGroupsGroup {
	s.GroupName = &v
	return s
}

func (s *ListGroupsForUserResponseBodyGroupsGroup) SetJoinDate(v string) *ListGroupsForUserResponseBodyGroupsGroup {
	s.JoinDate = &v
	return s
}

type ListGroupsForUserResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupsForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupsForUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsForUserResponse) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserResponse) SetHeaders(v map[string]*string) *ListGroupsForUserResponse {
	s.Headers = v
	return s
}

func (s *ListGroupsForUserResponse) SetStatusCode(v int32) *ListGroupsForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupsForUserResponse) SetBody(v *ListGroupsForUserResponseBody) *ListGroupsForUserResponse {
	s.Body = v
	return s
}

type ListPoliciesRequest struct {
	// The `marker`. If part of a previous response is truncated, you can use this parameter to obtain the truncated part.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The number of entries to return. If a response is truncated because it reaches the value of `MaxItems`, the value of `IsTruncated` will be `true`.
	//
	// Valid values: 1 to 1000. Default value: 100.
	//
	// example:
	//
	// 100
	MaxItems *int32 `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
	// The type of the policies. Valid values: `System` and `Custom`. If you do not specify the parameter, all policies are returned.``
	//
	// example:
	//
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The tags.
	Tag []*ListPoliciesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListPoliciesRequest) SetMarker(v string) *ListPoliciesRequest {
	s.Marker = &v
	return s
}

func (s *ListPoliciesRequest) SetMaxItems(v int32) *ListPoliciesRequest {
	s.MaxItems = &v
	return s
}

func (s *ListPoliciesRequest) SetPolicyType(v string) *ListPoliciesRequest {
	s.PolicyType = &v
	return s
}

func (s *ListPoliciesRequest) SetTag(v []*ListPoliciesRequestTag) *ListPoliciesRequest {
	s.Tag = v
	return s
}

type ListPoliciesRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// owner
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// alice
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPoliciesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesRequestTag) GoString() string {
	return s.String()
}

func (s *ListPoliciesRequestTag) SetKey(v string) *ListPoliciesRequestTag {
	s.Key = &v
	return s
}

func (s *ListPoliciesRequestTag) SetValue(v string) *ListPoliciesRequestTag {
	s.Value = &v
	return s
}

type ListPoliciesShrinkRequest struct {
	// The `marker`. If part of a previous response is truncated, you can use this parameter to obtain the truncated part.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The number of entries to return. If a response is truncated because it reaches the value of `MaxItems`, the value of `IsTruncated` will be `true`.
	//
	// Valid values: 1 to 1000. Default value: 100.
	//
	// example:
	//
	// 100
	MaxItems *int32 `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
	// The type of the policies. Valid values: `System` and `Custom`. If you do not specify the parameter, all policies are returned.``
	//
	// example:
	//
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The tags.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListPoliciesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPoliciesShrinkRequest) SetMarker(v string) *ListPoliciesShrinkRequest {
	s.Marker = &v
	return s
}

func (s *ListPoliciesShrinkRequest) SetMaxItems(v int32) *ListPoliciesShrinkRequest {
	s.MaxItems = &v
	return s
}

func (s *ListPoliciesShrinkRequest) SetPolicyType(v string) *ListPoliciesShrinkRequest {
	s.PolicyType = &v
	return s
}

func (s *ListPoliciesShrinkRequest) SetTagShrink(v string) *ListPoliciesShrinkRequest {
	s.TagShrink = &v
	return s
}

type ListPoliciesResponseBody struct {
	// Indicates whether the response is truncated.
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The marker. This parameter is returned only if the value of `IsTruncated` is `true`. If the parameter is returned, you can call this operation again and set `Marker` to obtain the truncated part.``
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The information about the policies.
	Policies *ListPoliciesResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7B8A4E7D-6CFF-471D-84DF-195A7A241ECB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBody) SetIsTruncated(v bool) *ListPoliciesResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListPoliciesResponseBody) SetMarker(v string) *ListPoliciesResponseBody {
	s.Marker = &v
	return s
}

func (s *ListPoliciesResponseBody) SetPolicies(v *ListPoliciesResponseBodyPolicies) *ListPoliciesResponseBody {
	s.Policies = v
	return s
}

func (s *ListPoliciesResponseBody) SetRequestId(v string) *ListPoliciesResponseBody {
	s.RequestId = &v
	return s
}

type ListPoliciesResponseBodyPolicies struct {
	Policy []*ListPoliciesResponseBodyPoliciesPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Repeated"`
}

func (s ListPoliciesResponseBodyPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBodyPolicies) SetPolicy(v []*ListPoliciesResponseBodyPoliciesPolicy) *ListPoliciesResponseBodyPolicies {
	s.Policy = v
	return s
}

type ListPoliciesResponseBodyPoliciesPolicy struct {
	// The number of references to the policy.
	//
	// example:
	//
	// 3
	AttachmentCount *int32 `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty"`
	// The time when the policy was created.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The default version of the policy.
	//
	// example:
	//
	// v1
	DefaultVersion *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// The description of the policy.
	//
	// example:
	//
	// OSS administrator
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy.
	//
	// example:
	//
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The tags.
	Tags *ListPoliciesResponseBodyPoliciesPolicyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The time when the policy was modified.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListPoliciesResponseBodyPoliciesPolicy) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesResponseBodyPoliciesPolicy) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) SetAttachmentCount(v int32) *ListPoliciesResponseBodyPoliciesPolicy {
	s.AttachmentCount = &v
	return s
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) SetCreateDate(v string) *ListPoliciesResponseBodyPoliciesPolicy {
	s.CreateDate = &v
	return s
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) SetDefaultVersion(v string) *ListPoliciesResponseBodyPoliciesPolicy {
	s.DefaultVersion = &v
	return s
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) SetDescription(v string) *ListPoliciesResponseBodyPoliciesPolicy {
	s.Description = &v
	return s
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) SetPolicyName(v string) *ListPoliciesResponseBodyPoliciesPolicy {
	s.PolicyName = &v
	return s
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) SetPolicyType(v string) *ListPoliciesResponseBodyPoliciesPolicy {
	s.PolicyType = &v
	return s
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) SetTags(v *ListPoliciesResponseBodyPoliciesPolicyTags) *ListPoliciesResponseBodyPoliciesPolicy {
	s.Tags = v
	return s
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) SetUpdateDate(v string) *ListPoliciesResponseBodyPoliciesPolicy {
	s.UpdateDate = &v
	return s
}

type ListPoliciesResponseBodyPoliciesPolicyTags struct {
	Tag []*ListPoliciesResponseBodyPoliciesPolicyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListPoliciesResponseBodyPoliciesPolicyTags) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesResponseBodyPoliciesPolicyTags) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBodyPoliciesPolicyTags) SetTag(v []*ListPoliciesResponseBodyPoliciesPolicyTagsTag) *ListPoliciesResponseBodyPoliciesPolicyTags {
	s.Tag = v
	return s
}

type ListPoliciesResponseBodyPoliciesPolicyTagsTag struct {
	// The key of the tag.
	//
	// example:
	//
	// owner
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// alice
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListPoliciesResponseBodyPoliciesPolicyTagsTag) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesResponseBodyPoliciesPolicyTagsTag) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBodyPoliciesPolicyTagsTag) SetTagKey(v string) *ListPoliciesResponseBodyPoliciesPolicyTagsTag {
	s.TagKey = &v
	return s
}

func (s *ListPoliciesResponseBodyPoliciesPolicyTagsTag) SetTagValue(v string) *ListPoliciesResponseBodyPoliciesPolicyTagsTag {
	s.TagValue = &v
	return s
}

type ListPoliciesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponse) SetHeaders(v map[string]*string) *ListPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListPoliciesResponse) SetStatusCode(v int32) *ListPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPoliciesResponse) SetBody(v *ListPoliciesResponseBody) *ListPoliciesResponse {
	s.Body = v
	return s
}

type ListPoliciesForGroupRequest struct {
	// The name of the RAM user group.
	//
	// example:
	//
	// dev
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s ListPoliciesForGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesForGroupRequest) GoString() string {
	return s.String()
}

func (s *ListPoliciesForGroupRequest) SetGroupName(v string) *ListPoliciesForGroupRequest {
	s.GroupName = &v
	return s
}

type ListPoliciesForGroupResponseBody struct {
	// The information about the policies.
	Policies *ListPoliciesForGroupResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7B8A4E7D-6CFF-471D-84DF-195A7A241ECB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPoliciesForGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesForGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListPoliciesForGroupResponseBody) SetPolicies(v *ListPoliciesForGroupResponseBodyPolicies) *ListPoliciesForGroupResponseBody {
	s.Policies = v
	return s
}

func (s *ListPoliciesForGroupResponseBody) SetRequestId(v string) *ListPoliciesForGroupResponseBody {
	s.RequestId = &v
	return s
}

type ListPoliciesForGroupResponseBodyPolicies struct {
	Policy []*ListPoliciesForGroupResponseBodyPoliciesPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Repeated"`
}

func (s ListPoliciesForGroupResponseBodyPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesForGroupResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *ListPoliciesForGroupResponseBodyPolicies) SetPolicy(v []*ListPoliciesForGroupResponseBodyPoliciesPolicy) *ListPoliciesForGroupResponseBodyPolicies {
	s.Policy = v
	return s
}

type ListPoliciesForGroupResponseBodyPoliciesPolicy struct {
	// The time when the policy was attached to the RAM user group.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	AttachDate *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty"`
	// The default version of the policy.
	//
	// example:
	//
	// v1
	DefaultVersion *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// The description of the policy.
	//
	// example:
	//
	// OSS administrator
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy.
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListPoliciesForGroupResponseBodyPoliciesPolicy) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesForGroupResponseBodyPoliciesPolicy) GoString() string {
	return s.String()
}

func (s *ListPoliciesForGroupResponseBodyPoliciesPolicy) SetAttachDate(v string) *ListPoliciesForGroupResponseBodyPoliciesPolicy {
	s.AttachDate = &v
	return s
}

func (s *ListPoliciesForGroupResponseBodyPoliciesPolicy) SetDefaultVersion(v string) *ListPoliciesForGroupResponseBodyPoliciesPolicy {
	s.DefaultVersion = &v
	return s
}

func (s *ListPoliciesForGroupResponseBodyPoliciesPolicy) SetDescription(v string) *ListPoliciesForGroupResponseBodyPoliciesPolicy {
	s.Description = &v
	return s
}

func (s *ListPoliciesForGroupResponseBodyPoliciesPolicy) SetPolicyName(v string) *ListPoliciesForGroupResponseBodyPoliciesPolicy {
	s.PolicyName = &v
	return s
}

func (s *ListPoliciesForGroupResponseBodyPoliciesPolicy) SetPolicyType(v string) *ListPoliciesForGroupResponseBodyPoliciesPolicy {
	s.PolicyType = &v
	return s
}

type ListPoliciesForGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPoliciesForGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPoliciesForGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesForGroupResponse) GoString() string {
	return s.String()
}

func (s *ListPoliciesForGroupResponse) SetHeaders(v map[string]*string) *ListPoliciesForGroupResponse {
	s.Headers = v
	return s
}

func (s *ListPoliciesForGroupResponse) SetStatusCode(v int32) *ListPoliciesForGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPoliciesForGroupResponse) SetBody(v *ListPoliciesForGroupResponseBody) *ListPoliciesForGroupResponse {
	s.Body = v
	return s
}

type ListPoliciesForRoleRequest struct {
	// The name of the RAM role.
	//
	// example:
	//
	// AdminRole
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s ListPoliciesForRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesForRoleRequest) GoString() string {
	return s.String()
}

func (s *ListPoliciesForRoleRequest) SetRoleName(v string) *ListPoliciesForRoleRequest {
	s.RoleName = &v
	return s
}

type ListPoliciesForRoleResponseBody struct {
	// The information about the policies.
	Policies *ListPoliciesForRoleResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7B8A4E7D-6CFF-471D-84DF-195A7A241ECB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPoliciesForRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesForRoleResponseBody) GoString() string {
	return s.String()
}

func (s *ListPoliciesForRoleResponseBody) SetPolicies(v *ListPoliciesForRoleResponseBodyPolicies) *ListPoliciesForRoleResponseBody {
	s.Policies = v
	return s
}

func (s *ListPoliciesForRoleResponseBody) SetRequestId(v string) *ListPoliciesForRoleResponseBody {
	s.RequestId = &v
	return s
}

type ListPoliciesForRoleResponseBodyPolicies struct {
	Policy []*ListPoliciesForRoleResponseBodyPoliciesPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Repeated"`
}

func (s ListPoliciesForRoleResponseBodyPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesForRoleResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *ListPoliciesForRoleResponseBodyPolicies) SetPolicy(v []*ListPoliciesForRoleResponseBodyPoliciesPolicy) *ListPoliciesForRoleResponseBodyPolicies {
	s.Policy = v
	return s
}

type ListPoliciesForRoleResponseBodyPoliciesPolicy struct {
	// The time when the policy was attached to the RAM role.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	AttachDate *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty"`
	// The default version of the policy.
	//
	// example:
	//
	// v1
	DefaultVersion *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// The description of the policy.
	//
	// example:
	//
	// OSS administrator
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy.
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListPoliciesForRoleResponseBodyPoliciesPolicy) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesForRoleResponseBodyPoliciesPolicy) GoString() string {
	return s.String()
}

func (s *ListPoliciesForRoleResponseBodyPoliciesPolicy) SetAttachDate(v string) *ListPoliciesForRoleResponseBodyPoliciesPolicy {
	s.AttachDate = &v
	return s
}

func (s *ListPoliciesForRoleResponseBodyPoliciesPolicy) SetDefaultVersion(v string) *ListPoliciesForRoleResponseBodyPoliciesPolicy {
	s.DefaultVersion = &v
	return s
}

func (s *ListPoliciesForRoleResponseBodyPoliciesPolicy) SetDescription(v string) *ListPoliciesForRoleResponseBodyPoliciesPolicy {
	s.Description = &v
	return s
}

func (s *ListPoliciesForRoleResponseBodyPoliciesPolicy) SetPolicyName(v string) *ListPoliciesForRoleResponseBodyPoliciesPolicy {
	s.PolicyName = &v
	return s
}

func (s *ListPoliciesForRoleResponseBodyPoliciesPolicy) SetPolicyType(v string) *ListPoliciesForRoleResponseBodyPoliciesPolicy {
	s.PolicyType = &v
	return s
}

type ListPoliciesForRoleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPoliciesForRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPoliciesForRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesForRoleResponse) GoString() string {
	return s.String()
}

func (s *ListPoliciesForRoleResponse) SetHeaders(v map[string]*string) *ListPoliciesForRoleResponse {
	s.Headers = v
	return s
}

func (s *ListPoliciesForRoleResponse) SetStatusCode(v int32) *ListPoliciesForRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPoliciesForRoleResponse) SetBody(v *ListPoliciesForRoleResponseBody) *ListPoliciesForRoleResponse {
	s.Body = v
	return s
}

type ListPoliciesForUserRequest struct {
	// The name of the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListPoliciesForUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesForUserRequest) GoString() string {
	return s.String()
}

func (s *ListPoliciesForUserRequest) SetUserName(v string) *ListPoliciesForUserRequest {
	s.UserName = &v
	return s
}

type ListPoliciesForUserResponseBody struct {
	// The information about the policy.
	Policies *ListPoliciesForUserResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7B8A4E7D-6CFF-471D-84DF-195A7A241ECB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPoliciesForUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListPoliciesForUserResponseBody) SetPolicies(v *ListPoliciesForUserResponseBodyPolicies) *ListPoliciesForUserResponseBody {
	s.Policies = v
	return s
}

func (s *ListPoliciesForUserResponseBody) SetRequestId(v string) *ListPoliciesForUserResponseBody {
	s.RequestId = &v
	return s
}

type ListPoliciesForUserResponseBodyPolicies struct {
	Policy []*ListPoliciesForUserResponseBodyPoliciesPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Repeated"`
}

func (s ListPoliciesForUserResponseBodyPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesForUserResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *ListPoliciesForUserResponseBodyPolicies) SetPolicy(v []*ListPoliciesForUserResponseBodyPoliciesPolicy) *ListPoliciesForUserResponseBodyPolicies {
	s.Policy = v
	return s
}

type ListPoliciesForUserResponseBodyPoliciesPolicy struct {
	// The time at which the policy is attached to the RAM user. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	AttachDate *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty"`
	// The current version.
	//
	// example:
	//
	// v1
	DefaultVersion *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// The description of the policy.
	//
	// example:
	//
	// OSS administrator
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy. Valid values:
	//
	// 	- System: system policy
	//
	// 	- Custom: custom policy
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListPoliciesForUserResponseBodyPoliciesPolicy) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesForUserResponseBodyPoliciesPolicy) GoString() string {
	return s.String()
}

func (s *ListPoliciesForUserResponseBodyPoliciesPolicy) SetAttachDate(v string) *ListPoliciesForUserResponseBodyPoliciesPolicy {
	s.AttachDate = &v
	return s
}

func (s *ListPoliciesForUserResponseBodyPoliciesPolicy) SetDefaultVersion(v string) *ListPoliciesForUserResponseBodyPoliciesPolicy {
	s.DefaultVersion = &v
	return s
}

func (s *ListPoliciesForUserResponseBodyPoliciesPolicy) SetDescription(v string) *ListPoliciesForUserResponseBodyPoliciesPolicy {
	s.Description = &v
	return s
}

func (s *ListPoliciesForUserResponseBodyPoliciesPolicy) SetPolicyName(v string) *ListPoliciesForUserResponseBodyPoliciesPolicy {
	s.PolicyName = &v
	return s
}

func (s *ListPoliciesForUserResponseBodyPoliciesPolicy) SetPolicyType(v string) *ListPoliciesForUserResponseBodyPoliciesPolicy {
	s.PolicyType = &v
	return s
}

type ListPoliciesForUserResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPoliciesForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPoliciesForUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesForUserResponse) GoString() string {
	return s.String()
}

func (s *ListPoliciesForUserResponse) SetHeaders(v map[string]*string) *ListPoliciesForUserResponse {
	s.Headers = v
	return s
}

func (s *ListPoliciesForUserResponse) SetStatusCode(v int32) *ListPoliciesForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPoliciesForUserResponse) SetBody(v *ListPoliciesForUserResponseBody) *ListPoliciesForUserResponse {
	s.Body = v
	return s
}

type ListPolicyVersionsRequest struct {
	// The name of the policy.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy. Valid values: `System` and `Custom`.
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListPolicyVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListPolicyVersionsRequest) SetPolicyName(v string) *ListPolicyVersionsRequest {
	s.PolicyName = &v
	return s
}

func (s *ListPolicyVersionsRequest) SetPolicyType(v string) *ListPolicyVersionsRequest {
	s.PolicyType = &v
	return s
}

type ListPolicyVersionsResponseBody struct {
	// The information about the policy versions.
	PolicyVersions *ListPolicyVersionsResponseBodyPolicyVersions `json:"PolicyVersions,omitempty" xml:"PolicyVersions,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7B8A4E7D-6CFF-471D-84DF-195A7A241ECB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPolicyVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPolicyVersionsResponseBody) SetPolicyVersions(v *ListPolicyVersionsResponseBodyPolicyVersions) *ListPolicyVersionsResponseBody {
	s.PolicyVersions = v
	return s
}

func (s *ListPolicyVersionsResponseBody) SetRequestId(v string) *ListPolicyVersionsResponseBody {
	s.RequestId = &v
	return s
}

type ListPolicyVersionsResponseBodyPolicyVersions struct {
	PolicyVersion []*ListPolicyVersionsResponseBodyPolicyVersionsPolicyVersion `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty" type:"Repeated"`
}

func (s ListPolicyVersionsResponseBodyPolicyVersions) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyVersionsResponseBodyPolicyVersions) GoString() string {
	return s.String()
}

func (s *ListPolicyVersionsResponseBodyPolicyVersions) SetPolicyVersion(v []*ListPolicyVersionsResponseBodyPolicyVersionsPolicyVersion) *ListPolicyVersionsResponseBodyPolicyVersions {
	s.PolicyVersion = v
	return s
}

type ListPolicyVersionsResponseBodyPolicyVersionsPolicyVersion struct {
	// The time when the version was created.
	//
	// example:
	//
	// 2015-02-26T01:25:52Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// Indicates whether the policy version is the default version.
	//
	// example:
	//
	// false
	IsDefaultVersion *bool `json:"IsDefaultVersion,omitempty" xml:"IsDefaultVersion,omitempty"`
	// The document of the policy.
	//
	// example:
	//
	// { "Statement": [{ "Action": ["oss:*"], "Effect": "Allow", "Resource": ["acs:oss:*:*:*"]}], "Version": "1"}
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The ID of the policy version.
	//
	// example:
	//
	// v3
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s ListPolicyVersionsResponseBodyPolicyVersionsPolicyVersion) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyVersionsResponseBodyPolicyVersionsPolicyVersion) GoString() string {
	return s.String()
}

func (s *ListPolicyVersionsResponseBodyPolicyVersionsPolicyVersion) SetCreateDate(v string) *ListPolicyVersionsResponseBodyPolicyVersionsPolicyVersion {
	s.CreateDate = &v
	return s
}

func (s *ListPolicyVersionsResponseBodyPolicyVersionsPolicyVersion) SetIsDefaultVersion(v bool) *ListPolicyVersionsResponseBodyPolicyVersionsPolicyVersion {
	s.IsDefaultVersion = &v
	return s
}

func (s *ListPolicyVersionsResponseBodyPolicyVersionsPolicyVersion) SetPolicyDocument(v string) *ListPolicyVersionsResponseBodyPolicyVersionsPolicyVersion {
	s.PolicyDocument = &v
	return s
}

func (s *ListPolicyVersionsResponseBodyPolicyVersionsPolicyVersion) SetVersionId(v string) *ListPolicyVersionsResponseBodyPolicyVersionsPolicyVersion {
	s.VersionId = &v
	return s
}

type ListPolicyVersionsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPolicyVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPolicyVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListPolicyVersionsResponse) SetHeaders(v map[string]*string) *ListPolicyVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListPolicyVersionsResponse) SetStatusCode(v int32) *ListPolicyVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPolicyVersionsResponse) SetBody(v *ListPolicyVersionsResponseBody) *ListPolicyVersionsResponse {
	s.Body = v
	return s
}

type ListRolesRequest struct {
	// The `marker`. If part of a previous response is truncated, you can use this parameter to obtain the truncated part.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The number of entries to return. If a response is truncated because it reaches the value of `MaxItems`, the value of `IsTruncated` will be `true`.
	//
	// Valid values: 1 to 1000. Default value: 100.
	//
	// example:
	//
	// 100
	MaxItems *int32 `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
	// The tags.
	Tag []*ListRolesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListRolesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRolesRequest) GoString() string {
	return s.String()
}

func (s *ListRolesRequest) SetMarker(v string) *ListRolesRequest {
	s.Marker = &v
	return s
}

func (s *ListRolesRequest) SetMaxItems(v int32) *ListRolesRequest {
	s.MaxItems = &v
	return s
}

func (s *ListRolesRequest) SetTag(v []*ListRolesRequestTag) *ListRolesRequest {
	s.Tag = v
	return s
}

type ListRolesRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// owner
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// alice
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListRolesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListRolesRequestTag) GoString() string {
	return s.String()
}

func (s *ListRolesRequestTag) SetKey(v string) *ListRolesRequestTag {
	s.Key = &v
	return s
}

func (s *ListRolesRequestTag) SetValue(v string) *ListRolesRequestTag {
	s.Value = &v
	return s
}

type ListRolesShrinkRequest struct {
	// The `marker`. If part of a previous response is truncated, you can use this parameter to obtain the truncated part.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The number of entries to return. If a response is truncated because it reaches the value of `MaxItems`, the value of `IsTruncated` will be `true`.
	//
	// Valid values: 1 to 1000. Default value: 100.
	//
	// example:
	//
	// 100
	MaxItems *int32 `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
	// The tags.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListRolesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRolesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListRolesShrinkRequest) SetMarker(v string) *ListRolesShrinkRequest {
	s.Marker = &v
	return s
}

func (s *ListRolesShrinkRequest) SetMaxItems(v int32) *ListRolesShrinkRequest {
	s.MaxItems = &v
	return s
}

func (s *ListRolesShrinkRequest) SetTagShrink(v string) *ListRolesShrinkRequest {
	s.TagShrink = &v
	return s
}

type ListRolesResponseBody struct {
	// Indicates whether the response is truncated.
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The marker. This parameter is returned only if the value of `IsTruncated` is `true`. If the parameter is returned, you can call this operation again and set this parameter to obtain the truncated part.````
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7B8A4E7D-6CFF-471D-84DF-195A7A241ECB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the RAM roles.
	Roles *ListRolesResponseBodyRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Struct"`
}

func (s ListRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBody) SetIsTruncated(v bool) *ListRolesResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListRolesResponseBody) SetMarker(v string) *ListRolesResponseBody {
	s.Marker = &v
	return s
}

func (s *ListRolesResponseBody) SetRequestId(v string) *ListRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRolesResponseBody) SetRoles(v *ListRolesResponseBodyRoles) *ListRolesResponseBody {
	s.Roles = v
	return s
}

type ListRolesResponseBodyRoles struct {
	Role []*ListRolesResponseBodyRolesRole `json:"Role,omitempty" xml:"Role,omitempty" type:"Repeated"`
}

func (s ListRolesResponseBodyRoles) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBodyRoles) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyRoles) SetRole(v []*ListRolesResponseBodyRolesRole) *ListRolesResponseBodyRoles {
	s.Role = v
	return s
}

type ListRolesResponseBodyRolesRole struct {
	// The Alibaba Cloud Resource Name (ARN) of the RAM role.
	//
	// example:
	//
	// acs:ram::123456789012****:role/ECSAdmin
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the RAM role.
	//
	// example:
	//
	// ECS administrator
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum session duration of the RAM role.
	//
	// example:
	//
	// 3600
	MaxSessionDuration *int64 `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty"`
	// The ID of the RAM role.
	//
	// example:
	//
	// 901234567890****
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// The name of the RAM role.
	//
	// example:
	//
	// ECSAdmin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The tags.
	Tags *ListRolesResponseBodyRolesRoleTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The update time.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListRolesResponseBodyRolesRole) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBodyRolesRole) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyRolesRole) SetArn(v string) *ListRolesResponseBodyRolesRole {
	s.Arn = &v
	return s
}

func (s *ListRolesResponseBodyRolesRole) SetCreateDate(v string) *ListRolesResponseBodyRolesRole {
	s.CreateDate = &v
	return s
}

func (s *ListRolesResponseBodyRolesRole) SetDescription(v string) *ListRolesResponseBodyRolesRole {
	s.Description = &v
	return s
}

func (s *ListRolesResponseBodyRolesRole) SetMaxSessionDuration(v int64) *ListRolesResponseBodyRolesRole {
	s.MaxSessionDuration = &v
	return s
}

func (s *ListRolesResponseBodyRolesRole) SetRoleId(v string) *ListRolesResponseBodyRolesRole {
	s.RoleId = &v
	return s
}

func (s *ListRolesResponseBodyRolesRole) SetRoleName(v string) *ListRolesResponseBodyRolesRole {
	s.RoleName = &v
	return s
}

func (s *ListRolesResponseBodyRolesRole) SetTags(v *ListRolesResponseBodyRolesRoleTags) *ListRolesResponseBodyRolesRole {
	s.Tags = v
	return s
}

func (s *ListRolesResponseBodyRolesRole) SetUpdateDate(v string) *ListRolesResponseBodyRolesRole {
	s.UpdateDate = &v
	return s
}

type ListRolesResponseBodyRolesRoleTags struct {
	Tag []*ListRolesResponseBodyRolesRoleTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListRolesResponseBodyRolesRoleTags) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBodyRolesRoleTags) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyRolesRoleTags) SetTag(v []*ListRolesResponseBodyRolesRoleTagsTag) *ListRolesResponseBodyRolesRoleTags {
	s.Tag = v
	return s
}

type ListRolesResponseBodyRolesRoleTagsTag struct {
	// The key of the tag.
	//
	// example:
	//
	// owner
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// alice
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListRolesResponseBodyRolesRoleTagsTag) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBodyRolesRoleTagsTag) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyRolesRoleTagsTag) SetTagKey(v string) *ListRolesResponseBodyRolesRoleTagsTag {
	s.TagKey = &v
	return s
}

func (s *ListRolesResponseBodyRolesRoleTagsTag) SetTagValue(v string) *ListRolesResponseBodyRolesRoleTagsTag {
	s.TagValue = &v
	return s
}

type ListRolesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponse) GoString() string {
	return s.String()
}

func (s *ListRolesResponse) SetHeaders(v map[string]*string) *ListRolesResponse {
	s.Headers = v
	return s
}

func (s *ListRolesResponse) SetStatusCode(v int32) *ListRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRolesResponse) SetBody(v *ListRolesResponseBody) *ListRolesResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// GY/oWREsOP1bPHGcHGrXfYX7UG1k9KqWFYThNDPx1UX26PbWOIu2CMqqiMr68H/K
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The names of the resources. You can specify up to 50 resource names.
	ResourceNames []*string `json:"ResourceNames,omitempty" xml:"ResourceNames,omitempty" type:"Repeated"`
	// The resource type.
	//
	// Enumerated values:
	//
	// 	- role
	//
	// 	- policy
	//
	// example:
	//
	// role
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags. You can specify up to 20 tags.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *ListTagResourcesRequest) SetPageSize(v int32) *ListTagResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceNames(v []*string) *ListTagResourcesRequest {
	s.ResourceNames = v
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
	// The key of the tag. The tag key can be up to 128 characters in length.
	//
	// example:
	//
	// t1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. The tag value can be up to 256 characters in length.
	//
	// example:
	//
	// v1
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

type ListTagResourcesShrinkRequest struct {
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// GY/oWREsOP1bPHGcHGrXfYX7UG1k9KqWFYThNDPx1UX26PbWOIu2CMqqiMr68H/K
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The names of the resources. You can specify up to 50 resource names.
	ResourceNamesShrink *string `json:"ResourceNames,omitempty" xml:"ResourceNames,omitempty"`
	// The resource type.
	//
	// Enumerated values:
	//
	// 	- role
	//
	// 	- policy
	//
	// example:
	//
	// role
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags. You can specify up to 20 tags.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListTagResourcesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesShrinkRequest) SetNextToken(v string) *ListTagResourcesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetPageSize(v int32) *ListTagResourcesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetResourceNamesShrink(v string) *ListTagResourcesShrinkRequest {
	s.ResourceNamesShrink = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetResourceType(v string) *ListTagResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetTagShrink(v string) *ListTagResourcesShrinkRequest {
	s.TagShrink = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// GY/oWREsOP1bPHGcHGrXfYX7UG1k9KqWFYThNDPx1UX26PbWOIu2CMqqiMr68H/K
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5CCE804C-6450-49A7-B1DB-2460F7A97416
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tags that are added to the resources.
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
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

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	// The name of the resource.
	//
	// example:
	//
	// role1
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource type.
	//
	// example:
	//
	// role
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of the tag.
	//
	// example:
	//
	// t1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// v1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceName(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceName = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
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

type ListUsersRequest struct {
	// The `marker`. If part of a previous response is truncated, you can use this parameter to obtain the truncated part.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The number of entries per page. If a response is truncated because it reaches the value of MaxItems, the value of `IsTruncatedg` will be `true`.
	//
	// Valid values: 1 to 1000. Default value: 100.
	//
	// example:
	//
	// 100
	MaxItems *int32 `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetMarker(v string) *ListUsersRequest {
	s.Marker = &v
	return s
}

func (s *ListUsersRequest) SetMaxItems(v int32) *ListUsersRequest {
	s.MaxItems = &v
	return s
}

type ListUsersResponseBody struct {
	// Indicates whether the response is truncated.
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The marker. This parameter is returned only if the value of `IsTruncated` is `true`. If the parameter is returned, you can call this operation again and set `Marker` to obtain the truncated part.``
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4B450CA1-36E8-4AA2-8461-86B42BF4CC4E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The RAM users.
	Users *ListUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s ListUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) SetIsTruncated(v bool) *ListUsersResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListUsersResponseBody) SetMarker(v string) *ListUsersResponseBody {
	s.Marker = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetUsers(v *ListUsersResponseBodyUsers) *ListUsersResponseBody {
	s.Users = v
	return s
}

type ListUsersResponseBodyUsers struct {
	User []*ListUsersResponseBodyUsersUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsers) SetUser(v []*ListUsersResponseBodyUsersUser) *ListUsersResponseBodyUsers {
	s.User = v
	return s
}

type ListUsersResponseBodyUsersUser struct {
	// The description.
	//
	// example:
	//
	// Cloud computing engineer
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The time when the RAM user was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The display name of the RAM user.
	//
	// example:
	//
	// Zhangq****
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the RAM user.
	//
	// > This parameter is unavailable.
	//
	// example:
	//
	// zhangq****@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The mobile phone number of the RAM user.
	//
	// > This parameter is unavailable.
	//
	// example:
	//
	// 86-1860000****
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	// The point in time when the information about the RAM user was last modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	// The ID of the RAM user.
	//
	// example:
	//
	// 122748924538****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The logon name of the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListUsersResponseBodyUsersUser) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyUsersUser) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsersUser) SetComments(v string) *ListUsersResponseBodyUsersUser {
	s.Comments = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetCreateDate(v string) *ListUsersResponseBodyUsersUser {
	s.CreateDate = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetDisplayName(v string) *ListUsersResponseBodyUsersUser {
	s.DisplayName = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetEmail(v string) *ListUsersResponseBodyUsersUser {
	s.Email = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetMobilePhone(v string) *ListUsersResponseBodyUsersUser {
	s.MobilePhone = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetUpdateDate(v string) *ListUsersResponseBodyUsersUser {
	s.UpdateDate = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetUserId(v string) *ListUsersResponseBodyUsersUser {
	s.UserId = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetUserName(v string) *ListUsersResponseBodyUsersUser {
	s.UserName = &v
	return s
}

type ListUsersResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponse) GoString() string {
	return s.String()
}

func (s *ListUsersResponse) SetHeaders(v map[string]*string) *ListUsersResponse {
	s.Headers = v
	return s
}

func (s *ListUsersResponse) SetStatusCode(v int32) *ListUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersResponse) SetBody(v *ListUsersResponseBody) *ListUsersResponse {
	s.Body = v
	return s
}

type ListUsersForGroupRequest struct {
	// The name of the RAM user group.
	//
	// example:
	//
	// Dev-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The `marker`. If part of a previous response is truncated, you can use this parameter to obtain the truncated part.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The number of entries to return. If a response is truncated because it reaches the value of `MaxItems`, the value of `IsTruncated` will be `true`.
	//
	// Valid values: 1 to 1000. Default value: 100.
	//
	// example:
	//
	// 100
	MaxItems *int32 `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
}

func (s ListUsersForGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersForGroupRequest) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupRequest) SetGroupName(v string) *ListUsersForGroupRequest {
	s.GroupName = &v
	return s
}

func (s *ListUsersForGroupRequest) SetMarker(v string) *ListUsersForGroupRequest {
	s.Marker = &v
	return s
}

func (s *ListUsersForGroupRequest) SetMaxItems(v int32) *ListUsersForGroupRequest {
	s.MaxItems = &v
	return s
}

type ListUsersForGroupResponseBody struct {
	// Indicates whether the response is truncated.
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The marker. This parameter is returned only if the value of `IsTruncated` is `true`. If the parameter is returned, you can call this operation again and set this parameter to obtain the truncated part.````
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4B450CA1-36E8-4AA2-8461-86B42BF4CC4E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the RAM users.
	Users *ListUsersForGroupResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s ListUsersForGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersForGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupResponseBody) SetIsTruncated(v bool) *ListUsersForGroupResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListUsersForGroupResponseBody) SetMarker(v string) *ListUsersForGroupResponseBody {
	s.Marker = &v
	return s
}

func (s *ListUsersForGroupResponseBody) SetRequestId(v string) *ListUsersForGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersForGroupResponseBody) SetUsers(v *ListUsersForGroupResponseBodyUsers) *ListUsersForGroupResponseBody {
	s.Users = v
	return s
}

type ListUsersForGroupResponseBodyUsers struct {
	User []*ListUsersForGroupResponseBodyUsersUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ListUsersForGroupResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s ListUsersForGroupResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupResponseBodyUsers) SetUser(v []*ListUsersForGroupResponseBodyUsersUser) *ListUsersForGroupResponseBodyUsers {
	s.User = v
	return s
}

type ListUsersForGroupResponseBodyUsersUser struct {
	// The display name of the RAM user.
	//
	// example:
	//
	// Alice
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The time when the RAM user joined the RAM user group.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	JoinDate *string `json:"JoinDate,omitempty" xml:"JoinDate,omitempty"`
	// The name of the RAM user.
	//
	// example:
	//
	// zhangqiang
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListUsersForGroupResponseBodyUsersUser) String() string {
	return tea.Prettify(s)
}

func (s ListUsersForGroupResponseBodyUsersUser) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupResponseBodyUsersUser) SetDisplayName(v string) *ListUsersForGroupResponseBodyUsersUser {
	s.DisplayName = &v
	return s
}

func (s *ListUsersForGroupResponseBodyUsersUser) SetJoinDate(v string) *ListUsersForGroupResponseBodyUsersUser {
	s.JoinDate = &v
	return s
}

func (s *ListUsersForGroupResponseBodyUsersUser) SetUserName(v string) *ListUsersForGroupResponseBodyUsersUser {
	s.UserName = &v
	return s
}

type ListUsersForGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUsersForGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUsersForGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersForGroupResponse) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupResponse) SetHeaders(v map[string]*string) *ListUsersForGroupResponse {
	s.Headers = v
	return s
}

func (s *ListUsersForGroupResponse) SetStatusCode(v int32) *ListUsersForGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersForGroupResponse) SetBody(v *ListUsersForGroupResponseBody) *ListUsersForGroupResponse {
	s.Body = v
	return s
}

type ListVirtualMFADevicesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the MFA devices.
	VirtualMFADevices *ListVirtualMFADevicesResponseBodyVirtualMFADevices `json:"VirtualMFADevices,omitempty" xml:"VirtualMFADevices,omitempty" type:"Struct"`
}

func (s ListVirtualMFADevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualMFADevicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVirtualMFADevicesResponseBody) SetRequestId(v string) *ListVirtualMFADevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVirtualMFADevicesResponseBody) SetVirtualMFADevices(v *ListVirtualMFADevicesResponseBodyVirtualMFADevices) *ListVirtualMFADevicesResponseBody {
	s.VirtualMFADevices = v
	return s
}

type ListVirtualMFADevicesResponseBodyVirtualMFADevices struct {
	VirtualMFADevice []*ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice `json:"VirtualMFADevice,omitempty" xml:"VirtualMFADevice,omitempty" type:"Repeated"`
}

func (s ListVirtualMFADevicesResponseBodyVirtualMFADevices) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualMFADevicesResponseBodyVirtualMFADevices) GoString() string {
	return s.String()
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevices) SetVirtualMFADevice(v []*ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice) *ListVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.VirtualMFADevice = v
	return s
}

type ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice struct {
	// The time when the MFA device was enabled.
	//
	// example:
	//
	// 2015-02-18T17:22:08Z
	ActivateDate *string `json:"ActivateDate,omitempty" xml:"ActivateDate,omitempty"`
	// The serial number of the MFA device.
	//
	// example:
	//
	// acs:ram::123456789012****:mfa/device002
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The basic information about the Resource Access Management (RAM) user to which the MFA device is bound.
	User *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice) GoString() string {
	return s.String()
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice) SetActivateDate(v string) *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice {
	s.ActivateDate = &v
	return s
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice) SetSerialNumber(v string) *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice {
	s.SerialNumber = &v
	return s
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice) SetUser(v *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser) *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice {
	s.User = v
	return s
}

type ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser struct {
	// The display name of the RAM user.
	//
	// example:
	//
	// zhangq****
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The unique ID of the RAM user.
	//
	// example:
	//
	// 122748924538****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The name of the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser) GoString() string {
	return s.String()
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser) SetDisplayName(v string) *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser {
	s.DisplayName = &v
	return s
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser) SetUserId(v string) *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser {
	s.UserId = &v
	return s
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser) SetUserName(v string) *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser {
	s.UserName = &v
	return s
}

type ListVirtualMFADevicesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVirtualMFADevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVirtualMFADevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualMFADevicesResponse) GoString() string {
	return s.String()
}

func (s *ListVirtualMFADevicesResponse) SetHeaders(v map[string]*string) *ListVirtualMFADevicesResponse {
	s.Headers = v
	return s
}

func (s *ListVirtualMFADevicesResponse) SetStatusCode(v int32) *ListVirtualMFADevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVirtualMFADevicesResponse) SetBody(v *ListVirtualMFADevicesResponseBody) *ListVirtualMFADevicesResponse {
	s.Body = v
	return s
}

type RemoveUserFromGroupRequest struct {
	// The name of the RAM user group.
	//
	// example:
	//
	// Dev-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The name of the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s RemoveUserFromGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserFromGroupRequest) SetGroupName(v string) *RemoveUserFromGroupRequest {
	s.GroupName = &v
	return s
}

func (s *RemoveUserFromGroupRequest) SetUserName(v string) *RemoveUserFromGroupRequest {
	s.UserName = &v
	return s
}

type RemoveUserFromGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A07EF215-B9B3-8CB2-2899-3F9575C6E320
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveUserFromGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUserFromGroupResponseBody) SetRequestId(v string) *RemoveUserFromGroupResponseBody {
	s.RequestId = &v
	return s
}

type RemoveUserFromGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveUserFromGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveUserFromGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveUserFromGroupResponse) SetHeaders(v map[string]*string) *RemoveUserFromGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveUserFromGroupResponse) SetStatusCode(v int32) *RemoveUserFromGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveUserFromGroupResponse) SetBody(v *RemoveUserFromGroupResponseBody) *RemoveUserFromGroupResponse {
	s.Body = v
	return s
}

type SetAccountAliasRequest struct {
	// The alias of the Alibaba Cloud account.
	//
	// The alias must be 3 to 32 characters in length, and can contain lowercase letters, digits, and hyphens (-).
	//
	// >  It cannot start or end with a hyphen (-), and cannot contain consecutive hyphens (-).
	//
	// example:
	//
	// myalias
	AccountAlias *string `json:"AccountAlias,omitempty" xml:"AccountAlias,omitempty"`
}

func (s SetAccountAliasRequest) String() string {
	return tea.Prettify(s)
}

func (s SetAccountAliasRequest) GoString() string {
	return s.String()
}

func (s *SetAccountAliasRequest) SetAccountAlias(v string) *SetAccountAliasRequest {
	s.AccountAlias = &v
	return s
}

type SetAccountAliasResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetAccountAliasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetAccountAliasResponseBody) GoString() string {
	return s.String()
}

func (s *SetAccountAliasResponseBody) SetRequestId(v string) *SetAccountAliasResponseBody {
	s.RequestId = &v
	return s
}

type SetAccountAliasResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetAccountAliasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetAccountAliasResponse) String() string {
	return tea.Prettify(s)
}

func (s SetAccountAliasResponse) GoString() string {
	return s.String()
}

func (s *SetAccountAliasResponse) SetHeaders(v map[string]*string) *SetAccountAliasResponse {
	s.Headers = v
	return s
}

func (s *SetAccountAliasResponse) SetStatusCode(v int32) *SetAccountAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAccountAliasResponse) SetBody(v *SetAccountAliasResponseBody) *SetAccountAliasResponse {
	s.Body = v
	return s
}

type SetDefaultPolicyVersionRequest struct {
	// The name of the policy.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The ID of the policy version that you want to set as the default version.
	//
	// example:
	//
	// v2
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s SetDefaultPolicyVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultPolicyVersionRequest) GoString() string {
	return s.String()
}

func (s *SetDefaultPolicyVersionRequest) SetPolicyName(v string) *SetDefaultPolicyVersionRequest {
	s.PolicyName = &v
	return s
}

func (s *SetDefaultPolicyVersionRequest) SetVersionId(v string) *SetDefaultPolicyVersionRequest {
	s.VersionId = &v
	return s
}

type SetDefaultPolicyVersionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDefaultPolicyVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultPolicyVersionResponseBody) GoString() string {
	return s.String()
}

func (s *SetDefaultPolicyVersionResponseBody) SetRequestId(v string) *SetDefaultPolicyVersionResponseBody {
	s.RequestId = &v
	return s
}

type SetDefaultPolicyVersionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDefaultPolicyVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDefaultPolicyVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultPolicyVersionResponse) GoString() string {
	return s.String()
}

func (s *SetDefaultPolicyVersionResponse) SetHeaders(v map[string]*string) *SetDefaultPolicyVersionResponse {
	s.Headers = v
	return s
}

func (s *SetDefaultPolicyVersionResponse) SetStatusCode(v int32) *SetDefaultPolicyVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDefaultPolicyVersionResponse) SetBody(v *SetDefaultPolicyVersionResponseBody) *SetDefaultPolicyVersionResponse {
	s.Body = v
	return s
}

type SetPasswordPolicyRequest struct {
	// Specifies whether a password will expire. Valid values: `true` and `false`. Default value: `false`. If you leave this parameter unspecified, the default value false is used.
	//
	// 	- If you set this parameter to `true`, the Alibaba Cloud account to which the RAM users belong must reset the passwords before the RAM users can log on to the Alibaba Cloud Management Console.
	//
	// 	- If you set this parameter to `false`, the RAM users can change the passwords after the passwords expire and then log on to the Alibaba Cloud Management Console.
	//
	// example:
	//
	// false
	HardExpiry *bool `json:"HardExpiry,omitempty" xml:"HardExpiry,omitempty"`
	// The maximum number of permitted logon attempts within one hour. The number of logon attempts is reset to zero if a RAM user changes the password.
	//
	// example:
	//
	// 5
	MaxLoginAttemps *int32 `json:"MaxLoginAttemps,omitempty" xml:"MaxLoginAttemps,omitempty"`
	// The number of days for which a password is valid. If you reset a password, the password validity period restarts. Default value: 0. The default value indicates that the password never expires.
	//
	// example:
	//
	// 0
	MaxPasswordAge *int32 `json:"MaxPasswordAge,omitempty" xml:"MaxPasswordAge,omitempty"`
	// The minimum number of characters in a password.
	//
	// Valid values: 8 to 32. Default value: 8.
	//
	// example:
	//
	// 12
	MinimumPasswordLength *int32 `json:"MinimumPasswordLength,omitempty" xml:"MinimumPasswordLength,omitempty"`
	// The number of previous passwords that a RAM user is prevented from reusing. Default value: 0. The default value indicates that the RAM user can reuse previous passwords.
	//
	// example:
	//
	// 0
	PasswordReusePrevention *int32 `json:"PasswordReusePrevention,omitempty" xml:"PasswordReusePrevention,omitempty"`
	// Specifies whether a password must contain one or more lowercase letters.
	//
	// example:
	//
	// true
	RequireLowercaseCharacters *bool `json:"RequireLowercaseCharacters,omitempty" xml:"RequireLowercaseCharacters,omitempty"`
	// Specifies whether a password must contain one or more digits.
	//
	// example:
	//
	// true
	RequireNumbers *bool `json:"RequireNumbers,omitempty" xml:"RequireNumbers,omitempty"`
	// Specifies whether a password must contain one or more special characters.
	//
	// example:
	//
	// true
	RequireSymbols *bool `json:"RequireSymbols,omitempty" xml:"RequireSymbols,omitempty"`
	// Specifies whether a password must contain one or more uppercase letters.
	//
	// example:
	//
	// true
	RequireUppercaseCharacters *bool `json:"RequireUppercaseCharacters,omitempty" xml:"RequireUppercaseCharacters,omitempty"`
}

func (s SetPasswordPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s SetPasswordPolicyRequest) GoString() string {
	return s.String()
}

func (s *SetPasswordPolicyRequest) SetHardExpiry(v bool) *SetPasswordPolicyRequest {
	s.HardExpiry = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetMaxLoginAttemps(v int32) *SetPasswordPolicyRequest {
	s.MaxLoginAttemps = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetMaxPasswordAge(v int32) *SetPasswordPolicyRequest {
	s.MaxPasswordAge = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetMinimumPasswordLength(v int32) *SetPasswordPolicyRequest {
	s.MinimumPasswordLength = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetPasswordReusePrevention(v int32) *SetPasswordPolicyRequest {
	s.PasswordReusePrevention = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetRequireLowercaseCharacters(v bool) *SetPasswordPolicyRequest {
	s.RequireLowercaseCharacters = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetRequireNumbers(v bool) *SetPasswordPolicyRequest {
	s.RequireNumbers = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetRequireSymbols(v bool) *SetPasswordPolicyRequest {
	s.RequireSymbols = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetRequireUppercaseCharacters(v bool) *SetPasswordPolicyRequest {
	s.RequireUppercaseCharacters = &v
	return s
}

type SetPasswordPolicyResponseBody struct {
	// The password policy.
	PasswordPolicy *SetPasswordPolicyResponseBodyPasswordPolicy `json:"PasswordPolicy,omitempty" xml:"PasswordPolicy,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetPasswordPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetPasswordPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *SetPasswordPolicyResponseBody) SetPasswordPolicy(v *SetPasswordPolicyResponseBodyPasswordPolicy) *SetPasswordPolicyResponseBody {
	s.PasswordPolicy = v
	return s
}

func (s *SetPasswordPolicyResponseBody) SetRequestId(v string) *SetPasswordPolicyResponseBody {
	s.RequestId = &v
	return s
}

type SetPasswordPolicyResponseBodyPasswordPolicy struct {
	// Indicates whether a password expires. Valid values: `true` and `false`. Default value: `false`. If the parameter is unspecified, the default value false is returned.
	//
	// 	- If this parameter is set to `true`, the Alibaba Cloud account to which the RAM users belong must reset the password before the RAM users can log on to the Alibaba Cloud Management Console.
	//
	// 	- If this parameter is set to `false`, the RAM users can change the passwords after the passwords expire and then log on to the Alibaba Cloud Management Console.
	//
	// example:
	//
	// false
	HardExpiry *bool `json:"HardExpiry,omitempty" xml:"HardExpiry,omitempty"`
	// The maximum number of permitted logon attempts within one hour. The number of logon attempts is reset to zero if a RAM user changes the password.
	//
	// example:
	//
	// 5
	MaxLoginAttemps *int32 `json:"MaxLoginAttemps,omitempty" xml:"MaxLoginAttemps,omitempty"`
	// The number of days for which a password is valid. If you reset a password, the password validity period restarts. Default value: 0. The default value indicates that the password never expires.
	//
	// example:
	//
	// 0
	MaxPasswordAge *int32 `json:"MaxPasswordAge,omitempty" xml:"MaxPasswordAge,omitempty"`
	// The minimum number of characters in a password.
	//
	// example:
	//
	// 12
	MinimumPasswordLength *int32 `json:"MinimumPasswordLength,omitempty" xml:"MinimumPasswordLength,omitempty"`
	// The number of previous passwords that a RAM user is prevented from reusing. Default value: 0. The default value indicates that the RAM user can reuse previous passwords.
	//
	// example:
	//
	// 0
	PasswordReusePrevention *int32 `json:"PasswordReusePrevention,omitempty" xml:"PasswordReusePrevention,omitempty"`
	// Indicates whether a password must contain one or more lowercase letters.
	//
	// example:
	//
	// true
	RequireLowercaseCharacters *bool `json:"RequireLowercaseCharacters,omitempty" xml:"RequireLowercaseCharacters,omitempty"`
	// Indicates whether a password must contain one or more digits.
	//
	// example:
	//
	// true
	RequireNumbers *bool `json:"RequireNumbers,omitempty" xml:"RequireNumbers,omitempty"`
	// Indicates whether a password must contain one or more special characters.
	//
	// example:
	//
	// true
	RequireSymbols *bool `json:"RequireSymbols,omitempty" xml:"RequireSymbols,omitempty"`
	// Indicates whether a password must contain one or more uppercase letters.
	//
	// example:
	//
	// true
	RequireUppercaseCharacters *bool `json:"RequireUppercaseCharacters,omitempty" xml:"RequireUppercaseCharacters,omitempty"`
}

func (s SetPasswordPolicyResponseBodyPasswordPolicy) String() string {
	return tea.Prettify(s)
}

func (s SetPasswordPolicyResponseBodyPasswordPolicy) GoString() string {
	return s.String()
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetHardExpiry(v bool) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.HardExpiry = &v
	return s
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetMaxLoginAttemps(v int32) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.MaxLoginAttemps = &v
	return s
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetMaxPasswordAge(v int32) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.MaxPasswordAge = &v
	return s
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetMinimumPasswordLength(v int32) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.MinimumPasswordLength = &v
	return s
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetPasswordReusePrevention(v int32) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.PasswordReusePrevention = &v
	return s
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetRequireLowercaseCharacters(v bool) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.RequireLowercaseCharacters = &v
	return s
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetRequireNumbers(v bool) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.RequireNumbers = &v
	return s
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetRequireSymbols(v bool) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.RequireSymbols = &v
	return s
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetRequireUppercaseCharacters(v bool) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.RequireUppercaseCharacters = &v
	return s
}

type SetPasswordPolicyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetPasswordPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetPasswordPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s SetPasswordPolicyResponse) GoString() string {
	return s.String()
}

func (s *SetPasswordPolicyResponse) SetHeaders(v map[string]*string) *SetPasswordPolicyResponse {
	s.Headers = v
	return s
}

func (s *SetPasswordPolicyResponse) SetStatusCode(v int32) *SetPasswordPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *SetPasswordPolicyResponse) SetBody(v *SetPasswordPolicyResponseBody) *SetPasswordPolicyResponse {
	s.Body = v
	return s
}

type SetSecurityPreferenceRequest struct {
	// Specifies whether RAM users can change their passwords. Valid values:
	//
	// 	- true (default)
	//
	// 	- false
	//
	// example:
	//
	// true
	AllowUserToChangePassword *bool `json:"AllowUserToChangePassword,omitempty" xml:"AllowUserToChangePassword,omitempty"`
	// Specifies whether RAM users can manage their AccessKey pairs. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	AllowUserToManageAccessKeys *bool `json:"AllowUserToManageAccessKeys,omitempty" xml:"AllowUserToManageAccessKeys,omitempty"`
	// Specifies whether RAM users can manage their MFA devices. Valid values:
	//
	// 	- true (default)
	//
	// 	- false
	//
	// example:
	//
	// true
	AllowUserToManageMFADevices *bool `json:"AllowUserToManageMFADevices,omitempty" xml:"AllowUserToManageMFADevices,omitempty"`
	// Specifies whether RAM users can manage their public keys. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// >  This parameter is valid only for the Japan site.
	//
	// example:
	//
	// false
	AllowUserToManagePublicKeys *bool `json:"AllowUserToManagePublicKeys,omitempty" xml:"AllowUserToManagePublicKeys,omitempty"`
	// Specifies whether to remember the multi-factor authentication (MFA) devices of Resource Access Management (RAM) users for seven days. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// true
	EnableSaveMFATicket *bool `json:"EnableSaveMFATicket,omitempty" xml:"EnableSaveMFATicket,omitempty"`
	// The subnet mask that specifies the IP addresses from which you can log on to the Alibaba Cloud Management Console. This parameter takes effect on password-based logon and single sign-on (SSO). However, this parameter does not take effect on API calls that are authenticated by using AccessKey pairs.
	//
	// 	- If you specify a subnet mask, RAM users can use only the IP addresses in the subnet mask to log on to the Alibaba Cloud Management Console.
	//
	// 	- If you do not specify a subnet mask, RAM users can use all IP addresses to log on to the Alibaba Cloud Management Console.
	//
	// If you need to specify multiple subnet masks, separate the subnet masks with semicolons (;). Example: 192.168.0.0/16;10.0.0.0/8.
	//
	// You can specify up to 40 subnet masks. The total length of the subnet masks can be a maximum of 512 characters.
	//
	// example:
	//
	// 10.0.0.0/8
	LoginNetworkMasks *string `json:"LoginNetworkMasks,omitempty" xml:"LoginNetworkMasks,omitempty"`
	// The validity period of the logon session of RAM users.
	//
	// Valid values: 1 to 24. Default value: 6. Unit: hours.
	//
	// example:
	//
	// 6
	LoginSessionDuration *int32 `json:"LoginSessionDuration,omitempty" xml:"LoginSessionDuration,omitempty"`
}

func (s SetSecurityPreferenceRequest) String() string {
	return tea.Prettify(s)
}

func (s SetSecurityPreferenceRequest) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceRequest) SetAllowUserToChangePassword(v bool) *SetSecurityPreferenceRequest {
	s.AllowUserToChangePassword = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetAllowUserToManageAccessKeys(v bool) *SetSecurityPreferenceRequest {
	s.AllowUserToManageAccessKeys = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetAllowUserToManageMFADevices(v bool) *SetSecurityPreferenceRequest {
	s.AllowUserToManageMFADevices = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetAllowUserToManagePublicKeys(v bool) *SetSecurityPreferenceRequest {
	s.AllowUserToManagePublicKeys = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetEnableSaveMFATicket(v bool) *SetSecurityPreferenceRequest {
	s.EnableSaveMFATicket = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetLoginNetworkMasks(v string) *SetSecurityPreferenceRequest {
	s.LoginNetworkMasks = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetLoginSessionDuration(v int32) *SetSecurityPreferenceRequest {
	s.LoginSessionDuration = &v
	return s
}

type SetSecurityPreferenceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A978915D-F279-4CA0-A89B-9A71219FFB3E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The security preferences.
	SecurityPreference *SetSecurityPreferenceResponseBodySecurityPreference `json:"SecurityPreference,omitempty" xml:"SecurityPreference,omitempty" type:"Struct"`
}

func (s SetSecurityPreferenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetSecurityPreferenceResponseBody) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseBody) SetRequestId(v string) *SetSecurityPreferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetSecurityPreferenceResponseBody) SetSecurityPreference(v *SetSecurityPreferenceResponseBodySecurityPreference) *SetSecurityPreferenceResponseBody {
	s.SecurityPreference = v
	return s
}

type SetSecurityPreferenceResponseBodySecurityPreference struct {
	// The AccessKey pair preference.
	AccessKeyPreference *SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference `json:"AccessKeyPreference,omitempty" xml:"AccessKeyPreference,omitempty" type:"Struct"`
	// The logon preference.
	LoginProfilePreference *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference `json:"LoginProfilePreference,omitempty" xml:"LoginProfilePreference,omitempty" type:"Struct"`
	// The MFA preference.
	MFAPreference *SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference `json:"MFAPreference,omitempty" xml:"MFAPreference,omitempty" type:"Struct"`
	// The public key preference.
	//
	// >  This parameter is valid only for the Japan site.
	PublicKeyPreference *SetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference `json:"PublicKeyPreference,omitempty" xml:"PublicKeyPreference,omitempty" type:"Struct"`
}

func (s SetSecurityPreferenceResponseBodySecurityPreference) String() string {
	return tea.Prettify(s)
}

func (s SetSecurityPreferenceResponseBodySecurityPreference) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) SetAccessKeyPreference(v *SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) *SetSecurityPreferenceResponseBodySecurityPreference {
	s.AccessKeyPreference = v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) SetLoginProfilePreference(v *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) *SetSecurityPreferenceResponseBodySecurityPreference {
	s.LoginProfilePreference = v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) SetMFAPreference(v *SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) *SetSecurityPreferenceResponseBodySecurityPreference {
	s.MFAPreference = v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) SetPublicKeyPreference(v *SetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference) *SetSecurityPreferenceResponseBodySecurityPreference {
	s.PublicKeyPreference = v
	return s
}

type SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference struct {
	// Indicates whether RAM users can manage their AccessKey pairs.
	//
	// example:
	//
	// false
	AllowUserToManageAccessKeys *bool `json:"AllowUserToManageAccessKeys,omitempty" xml:"AllowUserToManageAccessKeys,omitempty"`
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) String() string {
	return tea.Prettify(s)
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) SetAllowUserToManageAccessKeys(v bool) *SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference {
	s.AllowUserToManageAccessKeys = &v
	return s
}

type SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference struct {
	// Indicates whether RAM users can change their passwords.
	//
	// example:
	//
	// true
	AllowUserToChangePassword *bool `json:"AllowUserToChangePassword,omitempty" xml:"AllowUserToChangePassword,omitempty"`
	// Indicates whether the MFA devices of RAM users are remembered.
	//
	// example:
	//
	// false
	EnableSaveMFATicket *bool `json:"EnableSaveMFATicket,omitempty" xml:"EnableSaveMFATicket,omitempty"`
	// The subnet mask.
	//
	// example:
	//
	// 10.0.0.0/8
	LoginNetworkMasks *string `json:"LoginNetworkMasks,omitempty" xml:"LoginNetworkMasks,omitempty"`
	// The validity period of the logon session of RAM users.
	//
	// example:
	//
	// 6
	LoginSessionDuration *int32 `json:"LoginSessionDuration,omitempty" xml:"LoginSessionDuration,omitempty"`
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) String() string {
	return tea.Prettify(s)
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetAllowUserToChangePassword(v bool) *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.AllowUserToChangePassword = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetEnableSaveMFATicket(v bool) *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.EnableSaveMFATicket = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetLoginNetworkMasks(v string) *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.LoginNetworkMasks = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetLoginSessionDuration(v int32) *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.LoginSessionDuration = &v
	return s
}

type SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference struct {
	// Indicates whether RAM users can manage their MFA devices.
	//
	// example:
	//
	// false
	AllowUserToManageMFADevices *bool `json:"AllowUserToManageMFADevices,omitempty" xml:"AllowUserToManageMFADevices,omitempty"`
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) String() string {
	return tea.Prettify(s)
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) SetAllowUserToManageMFADevices(v bool) *SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference {
	s.AllowUserToManageMFADevices = &v
	return s
}

type SetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference struct {
	// Indicates whether RAM users can manage their public keys.
	//
	// example:
	//
	// false
	AllowUserToManagePublicKeys *bool `json:"AllowUserToManagePublicKeys,omitempty" xml:"AllowUserToManagePublicKeys,omitempty"`
}

func (s SetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference) String() string {
	return tea.Prettify(s)
}

func (s SetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference) SetAllowUserToManagePublicKeys(v bool) *SetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference {
	s.AllowUserToManagePublicKeys = &v
	return s
}

type SetSecurityPreferenceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetSecurityPreferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetSecurityPreferenceResponse) String() string {
	return tea.Prettify(s)
}

func (s SetSecurityPreferenceResponse) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponse) SetHeaders(v map[string]*string) *SetSecurityPreferenceResponse {
	s.Headers = v
	return s
}

func (s *SetSecurityPreferenceResponse) SetStatusCode(v int32) *SetSecurityPreferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *SetSecurityPreferenceResponse) SetBody(v *SetSecurityPreferenceResponseBody) *SetSecurityPreferenceResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The names of the resources. You can specify up to 50 resource names.
	ResourceNames []*string `json:"ResourceNames,omitempty" xml:"ResourceNames,omitempty" type:"Repeated"`
	// The resource type.
	//
	// Enumerated values:
	//
	// 	- role
	//
	// 	- policy
	//
	// example:
	//
	// role
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags. You can specify up to 20 tags.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetResourceNames(v []*string) *TagResourcesRequest {
	s.ResourceNames = v
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
	// The key of the tag. The tag key can be up to 128 characters in length.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. The tag value can be up to 256 characters in length.
	//
	// example:
	//
	// v1
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

type TagResourcesShrinkRequest struct {
	// The names of the resources. You can specify up to 50 resource names.
	ResourceNamesShrink *string `json:"ResourceNames,omitempty" xml:"ResourceNames,omitempty"`
	// The resource type.
	//
	// Enumerated values:
	//
	// 	- role
	//
	// 	- policy
	//
	// example:
	//
	// role
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags. You can specify up to 20 tags.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s TagResourcesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesShrinkRequest) SetResourceNamesShrink(v string) *TagResourcesShrinkRequest {
	s.ResourceNamesShrink = &v
	return s
}

func (s *TagResourcesShrinkRequest) SetResourceType(v string) *TagResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesShrinkRequest) SetTagShrink(v string) *TagResourcesShrinkRequest {
	s.TagShrink = &v
	return s
}

type TagResourcesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 697852FB-50D7-44D9-9774-530C31EAC572
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

type UnbindMFADeviceRequest struct {
	// Specifies the username.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UnbindMFADeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *UnbindMFADeviceRequest) SetUserName(v string) *UnbindMFADeviceRequest {
	s.UserName = &v
	return s
}

type UnbindMFADeviceResponseBody struct {
	// The information about the MFA device.
	MFADevice *UnbindMFADeviceResponseBodyMFADevice `json:"MFADevice,omitempty" xml:"MFADevice,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindMFADeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindMFADeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindMFADeviceResponseBody) SetMFADevice(v *UnbindMFADeviceResponseBodyMFADevice) *UnbindMFADeviceResponseBody {
	s.MFADevice = v
	return s
}

func (s *UnbindMFADeviceResponseBody) SetRequestId(v string) *UnbindMFADeviceResponseBody {
	s.RequestId = &v
	return s
}

type UnbindMFADeviceResponseBodyMFADevice struct {
	// The serial number of the MFA device.
	//
	// example:
	//
	// acs:ram::123456789012****:mfa/device002
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s UnbindMFADeviceResponseBodyMFADevice) String() string {
	return tea.Prettify(s)
}

func (s UnbindMFADeviceResponseBodyMFADevice) GoString() string {
	return s.String()
}

func (s *UnbindMFADeviceResponseBodyMFADevice) SetSerialNumber(v string) *UnbindMFADeviceResponseBodyMFADevice {
	s.SerialNumber = &v
	return s
}

type UnbindMFADeviceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindMFADeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindMFADeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *UnbindMFADeviceResponse) SetHeaders(v map[string]*string) *UnbindMFADeviceResponse {
	s.Headers = v
	return s
}

func (s *UnbindMFADeviceResponse) SetStatusCode(v int32) *UnbindMFADeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindMFADeviceResponse) SetBody(v *UnbindMFADeviceResponseBody) *UnbindMFADeviceResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all tags from the resources.
	//
	// Enumerated values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The names of the resources. You can specify up to 50 resource names.
	ResourceNames []*string `json:"ResourceNames,omitempty" xml:"ResourceNames,omitempty" type:"Repeated"`
	// The resource type.
	//
	// Enumerated values:
	//
	// 	- role
	//
	// 	- policy
	//
	// example:
	//
	// role
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The keys of the tags. You can specify up to 20 tag keys.
	TagKeys []*string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
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

func (s *UntagResourcesRequest) SetResourceNames(v []*string) *UntagResourcesRequest {
	s.ResourceNames = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKeys(v []*string) *UntagResourcesRequest {
	s.TagKeys = v
	return s
}

type UntagResourcesShrinkRequest struct {
	// Specifies whether to remove all tags from the resources.
	//
	// Enumerated values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The names of the resources. You can specify up to 50 resource names.
	ResourceNamesShrink *string `json:"ResourceNames,omitempty" xml:"ResourceNames,omitempty"`
	// The resource type.
	//
	// Enumerated values:
	//
	// 	- role
	//
	// 	- policy
	//
	// example:
	//
	// role
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The keys of the tags. You can specify up to 20 tag keys.
	TagKeysShrink *string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty"`
}

func (s UntagResourcesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesShrinkRequest) SetAll(v bool) *UntagResourcesShrinkRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesShrinkRequest) SetResourceNamesShrink(v string) *UntagResourcesShrinkRequest {
	s.ResourceNamesShrink = &v
	return s
}

func (s *UntagResourcesShrinkRequest) SetResourceType(v string) *UntagResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesShrinkRequest) SetTagKeysShrink(v string) *UntagResourcesShrinkRequest {
	s.TagKeysShrink = &v
	return s
}

type UntagResourcesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
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

type UpdateAccessKeyRequest struct {
	// The status of the AccessKey pair. Valid values: `Active` and `Inactive`.
	//
	// example:
	//
	// Inactive
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The AccessKey ID in the AccessKey pair whose status you want to change.``
	//
	// example:
	//
	// 0wNEpMMlzy7s****
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// The name of the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UpdateAccessKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccessKeyRequest) GoString() string {
	return s.String()
}

func (s *UpdateAccessKeyRequest) SetStatus(v string) *UpdateAccessKeyRequest {
	s.Status = &v
	return s
}

func (s *UpdateAccessKeyRequest) SetUserAccessKeyId(v string) *UpdateAccessKeyRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *UpdateAccessKeyRequest) SetUserName(v string) *UpdateAccessKeyRequest {
	s.UserName = &v
	return s
}

type UpdateAccessKeyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAccessKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccessKeyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAccessKeyResponseBody) SetRequestId(v string) *UpdateAccessKeyResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAccessKeyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAccessKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAccessKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccessKeyResponse) GoString() string {
	return s.String()
}

func (s *UpdateAccessKeyResponse) SetHeaders(v map[string]*string) *UpdateAccessKeyResponse {
	s.Headers = v
	return s
}

func (s *UpdateAccessKeyResponse) SetStatusCode(v int32) *UpdateAccessKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAccessKeyResponse) SetBody(v *UpdateAccessKeyResponseBody) *UpdateAccessKeyResponse {
	s.Body = v
	return s
}

type UpdateGroupRequest struct {
	// The name of the RAM user group.
	//
	// example:
	//
	// Dev-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The new description of the RAM user group.
	//
	// The new description must be 1 to 128 characters in length.
	//
	// example:
	//
	// NewDev-Team
	NewComments *string `json:"NewComments,omitempty" xml:"NewComments,omitempty"`
	// The new name of the RAM user group.
	//
	// The name must be 1 to 64 characters in length and can contain letters, digits, periods (.), hyphens (-), and underscores (_).
	//
	// example:
	//
	// NewDev-Team
	NewGroupName *string `json:"NewGroupName,omitempty" xml:"NewGroupName,omitempty"`
}

func (s UpdateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupRequest) SetGroupName(v string) *UpdateGroupRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateGroupRequest) SetNewComments(v string) *UpdateGroupRequest {
	s.NewComments = &v
	return s
}

func (s *UpdateGroupRequest) SetNewGroupName(v string) *UpdateGroupRequest {
	s.NewGroupName = &v
	return s
}

type UpdateGroupResponseBody struct {
	// The information about the RAM user group.
	Group *UpdateGroupResponseBodyGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// EC6647CC-0A36-EC7A-BA72-CC81BF3DE182
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupResponseBody) SetGroup(v *UpdateGroupResponseBodyGroup) *UpdateGroupResponseBody {
	s.Group = v
	return s
}

func (s *UpdateGroupResponseBody) SetRequestId(v string) *UpdateGroupResponseBody {
	s.RequestId = &v
	return s
}

type UpdateGroupResponseBodyGroup struct {
	// The description.
	//
	// example:
	//
	// NewDev-Team
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The ID of the RAM user group.
	//
	// example:
	//
	// g-FpMEHiMysofp****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the RAM user group.
	//
	// example:
	//
	// NewDev-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s UpdateGroupResponseBodyGroup) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupResponseBodyGroup) GoString() string {
	return s.String()
}

func (s *UpdateGroupResponseBodyGroup) SetComments(v string) *UpdateGroupResponseBodyGroup {
	s.Comments = &v
	return s
}

func (s *UpdateGroupResponseBodyGroup) SetCreateDate(v string) *UpdateGroupResponseBodyGroup {
	s.CreateDate = &v
	return s
}

func (s *UpdateGroupResponseBodyGroup) SetGroupId(v string) *UpdateGroupResponseBodyGroup {
	s.GroupId = &v
	return s
}

func (s *UpdateGroupResponseBodyGroup) SetGroupName(v string) *UpdateGroupResponseBodyGroup {
	s.GroupName = &v
	return s
}

func (s *UpdateGroupResponseBodyGroup) SetUpdateDate(v string) *UpdateGroupResponseBodyGroup {
	s.UpdateDate = &v
	return s
}

type UpdateGroupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupResponse) SetHeaders(v map[string]*string) *UpdateGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupResponse) SetStatusCode(v int32) *UpdateGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGroupResponse) SetBody(v *UpdateGroupResponseBody) *UpdateGroupResponse {
	s.Body = v
	return s
}

type UpdateLoginProfileRequest struct {
	// Specifies whether a multi-factor authentication (MFA) device must be bound to the RAM user upon logon.
	//
	// example:
	//
	// true
	MFABindRequired *bool `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty"`
	// The logon password of the RAM user. The password must meet the password strength requirements.
	//
	// example:
	//
	// mypassword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether the RAM user has to change the password upon logon.
	//
	// example:
	//
	// true
	PasswordResetRequired *bool `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty"`
	// The name of the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UpdateLoginProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoginProfileRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoginProfileRequest) SetMFABindRequired(v bool) *UpdateLoginProfileRequest {
	s.MFABindRequired = &v
	return s
}

func (s *UpdateLoginProfileRequest) SetPassword(v string) *UpdateLoginProfileRequest {
	s.Password = &v
	return s
}

func (s *UpdateLoginProfileRequest) SetPasswordResetRequired(v bool) *UpdateLoginProfileRequest {
	s.PasswordResetRequired = &v
	return s
}

func (s *UpdateLoginProfileRequest) SetUserName(v string) *UpdateLoginProfileRequest {
	s.UserName = &v
	return s
}

type UpdateLoginProfileResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLoginProfileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoginProfileResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLoginProfileResponseBody) SetRequestId(v string) *UpdateLoginProfileResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLoginProfileResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLoginProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLoginProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoginProfileResponse) GoString() string {
	return s.String()
}

func (s *UpdateLoginProfileResponse) SetHeaders(v map[string]*string) *UpdateLoginProfileResponse {
	s.Headers = v
	return s
}

func (s *UpdateLoginProfileResponse) SetStatusCode(v int32) *UpdateLoginProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLoginProfileResponse) SetBody(v *UpdateLoginProfileResponseBody) *UpdateLoginProfileResponse {
	s.Body = v
	return s
}

type UpdatePolicyDescriptionRequest struct {
	// The description of the policy.
	//
	// The description must be 1 to 1,024 characters in length.
	//
	// example:
	//
	// This is a test policy.
	NewDescription *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// TestPolicy
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s UpdatePolicyDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePolicyDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdatePolicyDescriptionRequest) SetNewDescription(v string) *UpdatePolicyDescriptionRequest {
	s.NewDescription = &v
	return s
}

func (s *UpdatePolicyDescriptionRequest) SetPolicyName(v string) *UpdatePolicyDescriptionRequest {
	s.PolicyName = &v
	return s
}

type UpdatePolicyDescriptionResponseBody struct {
	// The information about the policy.
	Policy *UpdatePolicyDescriptionResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7572DEBD-0ECE-518E-8682-D8CB82F8FE8E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePolicyDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePolicyDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePolicyDescriptionResponseBody) SetPolicy(v *UpdatePolicyDescriptionResponseBodyPolicy) *UpdatePolicyDescriptionResponseBody {
	s.Policy = v
	return s
}

func (s *UpdatePolicyDescriptionResponseBody) SetRequestId(v string) *UpdatePolicyDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePolicyDescriptionResponseBodyPolicy struct {
	// The time when the policy was created.
	//
	// example:
	//
	// 2022-02-28T07:04:15Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The version of the policy. Default value: v1.
	//
	// example:
	//
	// v1
	DefaultVersion *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// The description of the policy.
	//
	// example:
	//
	// This is a test policy.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// TestPolicy
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy. Valid values:
	//
	// 	- Custom
	//
	// 	- System
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The time when the policy was modified.
	//
	// example:
	//
	// 2022-02-28T07:05:37Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s UpdatePolicyDescriptionResponseBodyPolicy) String() string {
	return tea.Prettify(s)
}

func (s UpdatePolicyDescriptionResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *UpdatePolicyDescriptionResponseBodyPolicy) SetCreateDate(v string) *UpdatePolicyDescriptionResponseBodyPolicy {
	s.CreateDate = &v
	return s
}

func (s *UpdatePolicyDescriptionResponseBodyPolicy) SetDefaultVersion(v string) *UpdatePolicyDescriptionResponseBodyPolicy {
	s.DefaultVersion = &v
	return s
}

func (s *UpdatePolicyDescriptionResponseBodyPolicy) SetDescription(v string) *UpdatePolicyDescriptionResponseBodyPolicy {
	s.Description = &v
	return s
}

func (s *UpdatePolicyDescriptionResponseBodyPolicy) SetPolicyName(v string) *UpdatePolicyDescriptionResponseBodyPolicy {
	s.PolicyName = &v
	return s
}

func (s *UpdatePolicyDescriptionResponseBodyPolicy) SetPolicyType(v string) *UpdatePolicyDescriptionResponseBodyPolicy {
	s.PolicyType = &v
	return s
}

func (s *UpdatePolicyDescriptionResponseBodyPolicy) SetUpdateDate(v string) *UpdatePolicyDescriptionResponseBodyPolicy {
	s.UpdateDate = &v
	return s
}

type UpdatePolicyDescriptionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePolicyDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePolicyDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePolicyDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdatePolicyDescriptionResponse) SetHeaders(v map[string]*string) *UpdatePolicyDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdatePolicyDescriptionResponse) SetStatusCode(v int32) *UpdatePolicyDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePolicyDescriptionResponse) SetBody(v *UpdatePolicyDescriptionResponseBody) *UpdatePolicyDescriptionResponse {
	s.Body = v
	return s
}

type UpdateRoleRequest struct {
	// The trust policy that specifies the trusted entity to assume the RAM role.
	//
	// example:
	//
	// { "Statement": [ { "Action": "sts:AssumeRole", "Effect": "Allow", "Principal": { "RAM": "acs:ram::12345678901234****:root" } } ], "Version": "1" }
	NewAssumeRolePolicyDocument *string `json:"NewAssumeRolePolicyDocument,omitempty" xml:"NewAssumeRolePolicyDocument,omitempty"`
	// The new description of the RAM role.
	//
	// The description must be 1 to 1,024 characters in length.
	//
	// example:
	//
	// ECS administrator
	NewDescription *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	// The maximum session time of the RAM role.
	//
	// Valid values: 3600 to 43200. Unit: seconds. Default value: 3600.
	//
	// If you do not specify this parameter, the default value is used.
	//
	// example:
	//
	// 3600
	NewMaxSessionDuration *int64 `json:"NewMaxSessionDuration,omitempty" xml:"NewMaxSessionDuration,omitempty"`
	// The name of the RAM role.
	//
	// The name must be 1 to 64 characters in length, and can contain letters, digits, periods (.), and hyphens (-).
	//
	// example:
	//
	// ECSAdmin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s UpdateRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRoleRequest) SetNewAssumeRolePolicyDocument(v string) *UpdateRoleRequest {
	s.NewAssumeRolePolicyDocument = &v
	return s
}

func (s *UpdateRoleRequest) SetNewDescription(v string) *UpdateRoleRequest {
	s.NewDescription = &v
	return s
}

func (s *UpdateRoleRequest) SetNewMaxSessionDuration(v int64) *UpdateRoleRequest {
	s.NewMaxSessionDuration = &v
	return s
}

func (s *UpdateRoleRequest) SetRoleName(v string) *UpdateRoleRequest {
	s.RoleName = &v
	return s
}

type UpdateRoleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the RAM role.
	Role *UpdateRoleResponseBodyRole `json:"Role,omitempty" xml:"Role,omitempty" type:"Struct"`
}

func (s UpdateRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRoleResponseBody) SetRequestId(v string) *UpdateRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRoleResponseBody) SetRole(v *UpdateRoleResponseBodyRole) *UpdateRoleResponseBody {
	s.Role = v
	return s
}

type UpdateRoleResponseBodyRole struct {
	// The Alibaba Cloud Resource Name (ARN) of the RAM role.
	//
	// example:
	//
	// acs:ram::123456789012****:role/ECSAdmin
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The trust policy that specifies the trusted entity to assume the RAM role.
	//
	// example:
	//
	// { "Statement": [ { "Action": "sts:AssumeRole", "Effect": "Allow", "Principal": { "RAM": "acs:ram::123456789012****:root" } } ], "Version": "1" }
	AssumeRolePolicyDocument *string `json:"AssumeRolePolicyDocument,omitempty" xml:"AssumeRolePolicyDocument,omitempty"`
	// The time when the RAM role was created.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the RAM role.
	//
	// example:
	//
	// ECS administrator
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum session time of the RAM role.
	//
	// example:
	//
	// 3600
	MaxSessionDuration *int64 `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty"`
	// The ID of the RAM role.
	//
	// example:
	//
	// 901234567890****
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// The name of the RAM role.
	//
	// example:
	//
	// ECSAdmin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The time when the description of the RAM role was changed.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s UpdateRoleResponseBodyRole) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleResponseBodyRole) GoString() string {
	return s.String()
}

func (s *UpdateRoleResponseBodyRole) SetArn(v string) *UpdateRoleResponseBodyRole {
	s.Arn = &v
	return s
}

func (s *UpdateRoleResponseBodyRole) SetAssumeRolePolicyDocument(v string) *UpdateRoleResponseBodyRole {
	s.AssumeRolePolicyDocument = &v
	return s
}

func (s *UpdateRoleResponseBodyRole) SetCreateDate(v string) *UpdateRoleResponseBodyRole {
	s.CreateDate = &v
	return s
}

func (s *UpdateRoleResponseBodyRole) SetDescription(v string) *UpdateRoleResponseBodyRole {
	s.Description = &v
	return s
}

func (s *UpdateRoleResponseBodyRole) SetMaxSessionDuration(v int64) *UpdateRoleResponseBodyRole {
	s.MaxSessionDuration = &v
	return s
}

func (s *UpdateRoleResponseBodyRole) SetRoleId(v string) *UpdateRoleResponseBodyRole {
	s.RoleId = &v
	return s
}

func (s *UpdateRoleResponseBodyRole) SetRoleName(v string) *UpdateRoleResponseBodyRole {
	s.RoleName = &v
	return s
}

func (s *UpdateRoleResponseBodyRole) SetUpdateDate(v string) *UpdateRoleResponseBodyRole {
	s.UpdateDate = &v
	return s
}

type UpdateRoleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleResponse) GoString() string {
	return s.String()
}

func (s *UpdateRoleResponse) SetHeaders(v map[string]*string) *UpdateRoleResponse {
	s.Headers = v
	return s
}

func (s *UpdateRoleResponse) SetStatusCode(v int32) *UpdateRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRoleResponse) SetBody(v *UpdateRoleResponseBody) *UpdateRoleResponse {
	s.Body = v
	return s
}

type UpdateUserRequest struct {
	// The new description of the RAM user.
	//
	// The description must be 1 to 128 characters in length.
	//
	// example:
	//
	// This is a cloud computing engineer.
	NewComments *string `json:"NewComments,omitempty" xml:"NewComments,omitempty"`
	// The new display name of the RAM user.
	//
	// The name must be 1 to 128 characters in length.
	//
	// example:
	//
	// xiaoq****
	NewDisplayName *string `json:"NewDisplayName,omitempty" xml:"NewDisplayName,omitempty"`
	// The new email address of the RAM user.
	//
	// >  This parameter applies only to the China site (aliyun.com).
	//
	// example:
	//
	// xiaoq****@example.com
	NewEmail *string `json:"NewEmail,omitempty" xml:"NewEmail,omitempty"`
	// The new mobile phone number of the RAM user.
	//
	// Format: \\<Country code>-\\<Mobile phone number>.
	//
	// >  This parameter applies only to the China site (aliyun.com).
	//
	// example:
	//
	// 86-1860000****
	NewMobilePhone *string `json:"NewMobilePhone,omitempty" xml:"NewMobilePhone,omitempty"`
	// The new username of the RAM user.
	//
	// The username must be 1 to 64 characters in length, and can contain letters, digits, periods (.), hyphens (-), and underscores (_).
	//
	// example:
	//
	// xiaoq****
	NewUserName *string `json:"NewUserName,omitempty" xml:"NewUserName,omitempty"`
	// The username of the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UpdateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) SetNewComments(v string) *UpdateUserRequest {
	s.NewComments = &v
	return s
}

func (s *UpdateUserRequest) SetNewDisplayName(v string) *UpdateUserRequest {
	s.NewDisplayName = &v
	return s
}

func (s *UpdateUserRequest) SetNewEmail(v string) *UpdateUserRequest {
	s.NewEmail = &v
	return s
}

func (s *UpdateUserRequest) SetNewMobilePhone(v string) *UpdateUserRequest {
	s.NewMobilePhone = &v
	return s
}

func (s *UpdateUserRequest) SetNewUserName(v string) *UpdateUserRequest {
	s.NewUserName = &v
	return s
}

func (s *UpdateUserRequest) SetUserName(v string) *UpdateUserRequest {
	s.UserName = &v
	return s
}

type UpdateUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the RAM user.
	User *UpdateUserResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s UpdateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserResponseBody) SetRequestId(v string) *UpdateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserResponseBody) SetUser(v *UpdateUserResponseBodyUser) *UpdateUserResponseBody {
	s.User = v
	return s
}

type UpdateUserResponseBodyUser struct {
	// The description of the RAM user.
	//
	// example:
	//
	// This is a cloud computing engineer.
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The point in time when the RAM user was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The display name of the RAM user.
	//
	// example:
	//
	// xiaoq****
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the RAM user.
	//
	// >  This parameter can be returned only on the China site (aliyun.com).
	//
	// example:
	//
	// xiaoq****@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The mobile phone number of the RAM user.
	//
	// >  This parameter can be returned only on the China site (aliyun.com).
	//
	// example:
	//
	// 86-1860000****
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	// The point in time when the information about the RAM user was last modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-02-11T03:15:21Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	// The ID of the RAM user.
	//
	// example:
	//
	// 122748924538****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username of the RAM user.
	//
	// example:
	//
	// xiaoq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UpdateUserResponseBodyUser) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponseBodyUser) GoString() string {
	return s.String()
}

func (s *UpdateUserResponseBodyUser) SetComments(v string) *UpdateUserResponseBodyUser {
	s.Comments = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetCreateDate(v string) *UpdateUserResponseBodyUser {
	s.CreateDate = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetDisplayName(v string) *UpdateUserResponseBodyUser {
	s.DisplayName = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetEmail(v string) *UpdateUserResponseBodyUser {
	s.Email = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetMobilePhone(v string) *UpdateUserResponseBodyUser {
	s.MobilePhone = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetUpdateDate(v string) *UpdateUserResponseBodyUser {
	s.UpdateDate = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetUserId(v string) *UpdateUserResponseBodyUser {
	s.UserId = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetUserName(v string) *UpdateUserResponseBodyUser {
	s.UserName = &v
	return s
}

type UpdateUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserResponse) SetHeaders(v map[string]*string) *UpdateUserResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserResponse) SetStatusCode(v int32) *UpdateUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserResponse) SetBody(v *UpdateUserResponseBody) *UpdateUserResponse {
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
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ram"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// Summary:
//
// Adds a Resource Access Management (RAM) user to a RAM user group.
//
// @param request - AddUserToGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserToGroupResponse
func (client *Client) AddUserToGroupWithOptions(request *AddUserToGroupRequest, runtime *util.RuntimeOptions) (_result *AddUserToGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddUserToGroup"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddUserToGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a Resource Access Management (RAM) user to a RAM user group.
//
// @param request - AddUserToGroupRequest
//
// @return AddUserToGroupResponse
func (client *Client) AddUserToGroup(request *AddUserToGroupRequest) (_result *AddUserToGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUserToGroupResponse{}
	_body, _err := client.AddUserToGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Attaches a policy to a Resource Access Management (RAM) user group.
//
// @param request - AttachPolicyToGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachPolicyToGroupResponse
func (client *Client) AttachPolicyToGroupWithOptions(request *AttachPolicyToGroupRequest, runtime *util.RuntimeOptions) (_result *AttachPolicyToGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyType)) {
		query["PolicyType"] = request.PolicyType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachPolicyToGroup"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachPolicyToGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches a policy to a Resource Access Management (RAM) user group.
//
// @param request - AttachPolicyToGroupRequest
//
// @return AttachPolicyToGroupResponse
func (client *Client) AttachPolicyToGroup(request *AttachPolicyToGroupRequest) (_result *AttachPolicyToGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachPolicyToGroupResponse{}
	_body, _err := client.AttachPolicyToGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Attaches a policy to a Resource Access Management (RAM) role.
//
// @param request - AttachPolicyToRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachPolicyToRoleResponse
func (client *Client) AttachPolicyToRoleWithOptions(request *AttachPolicyToRoleRequest, runtime *util.RuntimeOptions) (_result *AttachPolicyToRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyType)) {
		query["PolicyType"] = request.PolicyType
	}

	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		query["RoleName"] = request.RoleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachPolicyToRole"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachPolicyToRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches a policy to a Resource Access Management (RAM) role.
//
// @param request - AttachPolicyToRoleRequest
//
// @return AttachPolicyToRoleResponse
func (client *Client) AttachPolicyToRole(request *AttachPolicyToRoleRequest) (_result *AttachPolicyToRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachPolicyToRoleResponse{}
	_body, _err := client.AttachPolicyToRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Attaches a policy to a Resource Access Management (RAM) user.
//
// @param request - AttachPolicyToUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachPolicyToUserResponse
func (client *Client) AttachPolicyToUserWithOptions(request *AttachPolicyToUserRequest, runtime *util.RuntimeOptions) (_result *AttachPolicyToUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyType)) {
		query["PolicyType"] = request.PolicyType
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachPolicyToUser"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachPolicyToUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches a policy to a Resource Access Management (RAM) user.
//
// @param request - AttachPolicyToUserRequest
//
// @return AttachPolicyToUserResponse
func (client *Client) AttachPolicyToUser(request *AttachPolicyToUserRequest) (_result *AttachPolicyToUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachPolicyToUserResponse{}
	_body, _err := client.AttachPolicyToUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds a multi-factor authentication (MFA) device to a Resource Access Management (RAM) user.
//
// @param request - BindMFADeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindMFADeviceResponse
func (client *Client) BindMFADeviceWithOptions(request *BindMFADeviceRequest, runtime *util.RuntimeOptions) (_result *BindMFADeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthenticationCode1)) {
		query["AuthenticationCode1"] = request.AuthenticationCode1
	}

	if !tea.BoolValue(util.IsUnset(request.AuthenticationCode2)) {
		query["AuthenticationCode2"] = request.AuthenticationCode2
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNumber)) {
		query["SerialNumber"] = request.SerialNumber
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindMFADevice"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindMFADeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a multi-factor authentication (MFA) device to a Resource Access Management (RAM) user.
//
// @param request - BindMFADeviceRequest
//
// @return BindMFADeviceResponse
func (client *Client) BindMFADevice(request *BindMFADeviceRequest) (_result *BindMFADeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindMFADeviceResponse{}
	_body, _err := client.BindMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the password that is used to log on to the console for a Resource Access Management (RAM) user.
//
// Description:
//
// >  This operation is available only for RAM users. Before you call this operation, make sure that `AllowUserToChangePassword` in [SetSecurityPreference](https://help.aliyun.com/document_detail/43765.html) is set to `True`. The value True indicates that RAM users can manage their passwords.
//
// @param request - ChangePasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangePasswordResponse
func (client *Client) ChangePasswordWithOptions(request *ChangePasswordRequest, runtime *util.RuntimeOptions) (_result *ChangePasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewPassword)) {
		query["NewPassword"] = request.NewPassword
	}

	if !tea.BoolValue(util.IsUnset(request.OldPassword)) {
		query["OldPassword"] = request.OldPassword
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangePassword"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangePasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the password that is used to log on to the console for a Resource Access Management (RAM) user.
//
// Description:
//
// >  This operation is available only for RAM users. Before you call this operation, make sure that `AllowUserToChangePassword` in [SetSecurityPreference](https://help.aliyun.com/document_detail/43765.html) is set to `True`. The value True indicates that RAM users can manage their passwords.
//
// @param request - ChangePasswordRequest
//
// @return ChangePasswordResponse
func (client *Client) ChangePassword(request *ChangePasswordRequest) (_result *ChangePasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangePasswordResponse{}
	_body, _err := client.ChangePasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the alias of an Alibaba Cloud account.
//
// @param request - ClearAccountAliasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClearAccountAliasResponse
func (client *Client) ClearAccountAliasWithOptions(runtime *util.RuntimeOptions) (_result *ClearAccountAliasResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ClearAccountAlias"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ClearAccountAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the alias of an Alibaba Cloud account.
//
// @return ClearAccountAliasResponse
func (client *Client) ClearAccountAlias() (_result *ClearAccountAliasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ClearAccountAliasResponse{}
	_body, _err := client.ClearAccountAliasWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an AccessKey pair for a Resource Access Management (RAM) user.
//
// @param request - CreateAccessKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccessKeyResponse
func (client *Client) CreateAccessKeyWithOptions(request *CreateAccessKeyRequest, runtime *util.RuntimeOptions) (_result *CreateAccessKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccessKey"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAccessKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an AccessKey pair for a Resource Access Management (RAM) user.
//
// @param request - CreateAccessKeyRequest
//
// @return CreateAccessKeyResponse
func (client *Client) CreateAccessKey(request *CreateAccessKeyRequest) (_result *CreateAccessKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAccessKeyResponse{}
	_body, _err := client.CreateAccessKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a RAM user group.
//
// @param request - CreateGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupResponse
func (client *Client) CreateGroupWithOptions(request *CreateGroupRequest, runtime *util.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Comments)) {
		query["Comments"] = request.Comments
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGroup"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a RAM user group.
//
// @param request - CreateGroupRequest
//
// @return CreateGroupResponse
func (client *Client) CreateGroup(request *CreateGroupRequest) (_result *CreateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGroupResponse{}
	_body, _err := client.CreateGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables console logon for a Resource Access Management (RAM) user.
//
// @param request - CreateLoginProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLoginProfileResponse
func (client *Client) CreateLoginProfileWithOptions(request *CreateLoginProfileRequest, runtime *util.RuntimeOptions) (_result *CreateLoginProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MFABindRequired)) {
		query["MFABindRequired"] = request.MFABindRequired
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordResetRequired)) {
		query["PasswordResetRequired"] = request.PasswordResetRequired
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLoginProfile"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLoginProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables console logon for a Resource Access Management (RAM) user.
//
// @param request - CreateLoginProfileRequest
//
// @return CreateLoginProfileResponse
func (client *Client) CreateLoginProfile(request *CreateLoginProfileRequest) (_result *CreateLoginProfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLoginProfileResponse{}
	_body, _err := client.CreateLoginProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom policy.
//
// Description:
//
// For more information about policies, see [Policy overview](https://help.aliyun.com/document_detail/93732.html).
//
// This topic provides an example on how to create a custom policy to query Elastic Compute Service (ECS) instances in a specific region.
//
// @param tmpReq - CreatePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolicyResponse
func (client *Client) CreatePolicyWithOptions(tmpReq *CreatePolicyRequest, runtime *util.RuntimeOptions) (_result *CreatePolicyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreatePolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tag)) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, tea.String("Tag"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyDocument)) {
		query["PolicyDocument"] = request.PolicyDocument
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.TagShrink)) {
		query["Tag"] = request.TagShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePolicy"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom policy.
//
// Description:
//
// For more information about policies, see [Policy overview](https://help.aliyun.com/document_detail/93732.html).
//
// This topic provides an example on how to create a custom policy to query Elastic Compute Service (ECS) instances in a specific region.
//
// @param request - CreatePolicyRequest
//
// @return CreatePolicyResponse
func (client *Client) CreatePolicy(request *CreatePolicyRequest) (_result *CreatePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePolicyResponse{}
	_body, _err := client.CreatePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a version for a policy.
//
// @param request - CreatePolicyVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolicyVersionResponse
func (client *Client) CreatePolicyVersionWithOptions(request *CreatePolicyVersionRequest, runtime *util.RuntimeOptions) (_result *CreatePolicyVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyDocument)) {
		query["PolicyDocument"] = request.PolicyDocument
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.RotateStrategy)) {
		query["RotateStrategy"] = request.RotateStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.SetAsDefault)) {
		query["SetAsDefault"] = request.SetAsDefault
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePolicyVersion"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePolicyVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a version for a policy.
//
// @param request - CreatePolicyVersionRequest
//
// @return CreatePolicyVersionResponse
func (client *Client) CreatePolicyVersion(request *CreatePolicyVersionRequest) (_result *CreatePolicyVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePolicyVersionResponse{}
	_body, _err := client.CreatePolicyVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Resource Access Management (RAM) role.
//
// Description:
//
// ### [](#)Operation description
//
// For more information about RAM roles, see [Overview of RAM roles](https://help.aliyun.com/document_detail/93689.html).
//
// @param tmpReq - CreateRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRoleResponse
func (client *Client) CreateRoleWithOptions(tmpReq *CreateRoleRequest, runtime *util.RuntimeOptions) (_result *CreateRoleResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateRoleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tag)) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, tea.String("Tag"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssumeRolePolicyDocument)) {
		query["AssumeRolePolicyDocument"] = request.AssumeRolePolicyDocument
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.MaxSessionDuration)) {
		query["MaxSessionDuration"] = request.MaxSessionDuration
	}

	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		query["RoleName"] = request.RoleName
	}

	if !tea.BoolValue(util.IsUnset(request.TagShrink)) {
		query["Tag"] = request.TagShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRole"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Resource Access Management (RAM) role.
//
// Description:
//
// ### [](#)Operation description
//
// For more information about RAM roles, see [Overview of RAM roles](https://help.aliyun.com/document_detail/93689.html).
//
// @param request - CreateRoleRequest
//
// @return CreateRoleResponse
func (client *Client) CreateRole(request *CreateRoleRequest) (_result *CreateRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRoleResponse{}
	_body, _err := client.CreateRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Resource Access Management (RAM) user.
//
// Description:
//
// This topic provides an example on how to create a RAM user named `alice`.
//
// @param request - CreateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserResponse
func (client *Client) CreateUserWithOptions(request *CreateUserRequest, runtime *util.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Comments)) {
		query["Comments"] = request.Comments
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.MobilePhone)) {
		query["MobilePhone"] = request.MobilePhone
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUser"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Resource Access Management (RAM) user.
//
// Description:
//
// This topic provides an example on how to create a RAM user named `alice`.
//
// @param request - CreateUserRequest
//
// @return CreateUserResponse
func (client *Client) CreateUser(request *CreateUserRequest) (_result *CreateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUserResponse{}
	_body, _err := client.CreateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a multi-factor authentication (MFA) device.
//
// @param request - CreateVirtualMFADeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVirtualMFADeviceResponse
func (client *Client) CreateVirtualMFADeviceWithOptions(request *CreateVirtualMFADeviceRequest, runtime *util.RuntimeOptions) (_result *CreateVirtualMFADeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VirtualMFADeviceName)) {
		query["VirtualMFADeviceName"] = request.VirtualMFADeviceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVirtualMFADevice"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVirtualMFADeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a multi-factor authentication (MFA) device.
//
// @param request - CreateVirtualMFADeviceRequest
//
// @return CreateVirtualMFADeviceResponse
func (client *Client) CreateVirtualMFADevice(request *CreateVirtualMFADeviceRequest) (_result *CreateVirtualMFADeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVirtualMFADeviceResponse{}
	_body, _err := client.CreateVirtualMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Decodes the diagnostic information in the response that contains an access denied error. The error is caused by no RAM permissions.
//
// @param request - DecodeDiagnosticMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DecodeDiagnosticMessageResponse
func (client *Client) DecodeDiagnosticMessageWithOptions(request *DecodeDiagnosticMessageRequest, runtime *util.RuntimeOptions) (_result *DecodeDiagnosticMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncodedDiagnosticMessage)) {
		query["EncodedDiagnosticMessage"] = request.EncodedDiagnosticMessage
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DecodeDiagnosticMessage"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DecodeDiagnosticMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Decodes the diagnostic information in the response that contains an access denied error. The error is caused by no RAM permissions.
//
// @param request - DecodeDiagnosticMessageRequest
//
// @return DecodeDiagnosticMessageResponse
func (client *Client) DecodeDiagnosticMessage(request *DecodeDiagnosticMessageRequest) (_result *DecodeDiagnosticMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DecodeDiagnosticMessageResponse{}
	_body, _err := client.DecodeDiagnosticMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an AccessKey pair of a Resource Access Management (RAM) user.
//
// @param request - DeleteAccessKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccessKeyResponse
func (client *Client) DeleteAccessKeyWithOptions(request *DeleteAccessKeyRequest, runtime *util.RuntimeOptions) (_result *DeleteAccessKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserAccessKeyId)) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccessKey"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAccessKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an AccessKey pair of a Resource Access Management (RAM) user.
//
// @param request - DeleteAccessKeyRequest
//
// @return DeleteAccessKeyResponse
func (client *Client) DeleteAccessKey(request *DeleteAccessKeyRequest) (_result *DeleteAccessKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAccessKeyResponse{}
	_body, _err := client.DeleteAccessKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Resource Access Management (RAM) user group.
//
// Description:
//
// Before you delete a RAM user group, make sure that no policies are attached to the group and no RAM users are included in the group.
//
// @param request - DeleteGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGroupResponse
func (client *Client) DeleteGroupWithOptions(request *DeleteGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGroup"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Resource Access Management (RAM) user group.
//
// Description:
//
// Before you delete a RAM user group, make sure that no policies are attached to the group and no RAM users are included in the group.
//
// @param request - DeleteGroupRequest
//
// @return DeleteGroupResponse
func (client *Client) DeleteGroup(request *DeleteGroupRequest) (_result *DeleteGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGroupResponse{}
	_body, _err := client.DeleteGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables console logon for a Resource Access Management (RAM) user.
//
// @param request - DeleteLoginProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLoginProfileResponse
func (client *Client) DeleteLoginProfileWithOptions(request *DeleteLoginProfileRequest, runtime *util.RuntimeOptions) (_result *DeleteLoginProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLoginProfile"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLoginProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables console logon for a Resource Access Management (RAM) user.
//
// @param request - DeleteLoginProfileRequest
//
// @return DeleteLoginProfileResponse
func (client *Client) DeleteLoginProfile(request *DeleteLoginProfileRequest) (_result *DeleteLoginProfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLoginProfileResponse{}
	_body, _err := client.DeleteLoginProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a policy.
//
// @param request - DeletePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicyResponse
func (client *Client) DeletePolicyWithOptions(request *DeletePolicyRequest, runtime *util.RuntimeOptions) (_result *DeletePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CascadingDelete)) {
		query["CascadingDelete"] = request.CascadingDelete
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePolicy"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a policy.
//
// @param request - DeletePolicyRequest
//
// @return DeletePolicyResponse
func (client *Client) DeletePolicy(request *DeletePolicyRequest) (_result *DeletePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePolicyResponse{}
	_body, _err := client.DeletePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a policy version.
//
// @param request - DeletePolicyVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicyVersionResponse
func (client *Client) DeletePolicyVersionWithOptions(request *DeletePolicyVersionRequest, runtime *util.RuntimeOptions) (_result *DeletePolicyVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePolicyVersion"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePolicyVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a policy version.
//
// @param request - DeletePolicyVersionRequest
//
// @return DeletePolicyVersionResponse
func (client *Client) DeletePolicyVersion(request *DeletePolicyVersionRequest) (_result *DeletePolicyVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePolicyVersionResponse{}
	_body, _err := client.DeletePolicyVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Resource Access Management (RAM) role.
//
// @param request - DeleteRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRoleResponse
func (client *Client) DeleteRoleWithOptions(request *DeleteRoleRequest, runtime *util.RuntimeOptions) (_result *DeleteRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		query["RoleName"] = request.RoleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRole"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Resource Access Management (RAM) role.
//
// @param request - DeleteRoleRequest
//
// @return DeleteRoleResponse
func (client *Client) DeleteRole(request *DeleteRoleRequest) (_result *DeleteRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRoleResponse{}
	_body, _err := client.DeleteRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Resource Access Management (RAM) user.
//
// Description:
//
// Before you delete a RAM user, make sure that no policies are attached to the RAM user and that the RAM user does not belong to any groups.
//
// @param request - DeleteUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserResponse
func (client *Client) DeleteUserWithOptions(request *DeleteUserRequest, runtime *util.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUser"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Resource Access Management (RAM) user.
//
// Description:
//
// Before you delete a RAM user, make sure that no policies are attached to the RAM user and that the RAM user does not belong to any groups.
//
// @param request - DeleteUserRequest
//
// @return DeleteUserResponse
func (client *Client) DeleteUser(request *DeleteUserRequest) (_result *DeleteUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserResponse{}
	_body, _err := client.DeleteUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a multi-factor authentication (MFA) device.
//
// @param request - DeleteVirtualMFADeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVirtualMFADeviceResponse
func (client *Client) DeleteVirtualMFADeviceWithOptions(request *DeleteVirtualMFADeviceRequest, runtime *util.RuntimeOptions) (_result *DeleteVirtualMFADeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SerialNumber)) {
		query["SerialNumber"] = request.SerialNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVirtualMFADevice"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVirtualMFADeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a multi-factor authentication (MFA) device.
//
// @param request - DeleteVirtualMFADeviceRequest
//
// @return DeleteVirtualMFADeviceResponse
func (client *Client) DeleteVirtualMFADevice(request *DeleteVirtualMFADeviceRequest) (_result *DeleteVirtualMFADeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVirtualMFADeviceResponse{}
	_body, _err := client.DeleteVirtualMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detaches a policy from a Resource Access Management (RAM) user group.
//
// @param request - DetachPolicyFromGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachPolicyFromGroupResponse
func (client *Client) DetachPolicyFromGroupWithOptions(request *DetachPolicyFromGroupRequest, runtime *util.RuntimeOptions) (_result *DetachPolicyFromGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyType)) {
		query["PolicyType"] = request.PolicyType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachPolicyFromGroup"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachPolicyFromGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detaches a policy from a Resource Access Management (RAM) user group.
//
// @param request - DetachPolicyFromGroupRequest
//
// @return DetachPolicyFromGroupResponse
func (client *Client) DetachPolicyFromGroup(request *DetachPolicyFromGroupRequest) (_result *DetachPolicyFromGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachPolicyFromGroupResponse{}
	_body, _err := client.DetachPolicyFromGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detaches a policy from a Resource Access Management (RAM) role.
//
// @param request - DetachPolicyFromRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachPolicyFromRoleResponse
func (client *Client) DetachPolicyFromRoleWithOptions(request *DetachPolicyFromRoleRequest, runtime *util.RuntimeOptions) (_result *DetachPolicyFromRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyType)) {
		query["PolicyType"] = request.PolicyType
	}

	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		query["RoleName"] = request.RoleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachPolicyFromRole"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachPolicyFromRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detaches a policy from a Resource Access Management (RAM) role.
//
// @param request - DetachPolicyFromRoleRequest
//
// @return DetachPolicyFromRoleResponse
func (client *Client) DetachPolicyFromRole(request *DetachPolicyFromRoleRequest) (_result *DetachPolicyFromRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachPolicyFromRoleResponse{}
	_body, _err := client.DetachPolicyFromRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detaches a policy from a Resource Access Management (RAM) user.
//
// @param request - DetachPolicyFromUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachPolicyFromUserResponse
func (client *Client) DetachPolicyFromUserWithOptions(request *DetachPolicyFromUserRequest, runtime *util.RuntimeOptions) (_result *DetachPolicyFromUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyType)) {
		query["PolicyType"] = request.PolicyType
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachPolicyFromUser"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachPolicyFromUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detaches a policy from a Resource Access Management (RAM) user.
//
// @param request - DetachPolicyFromUserRequest
//
// @return DetachPolicyFromUserResponse
func (client *Client) DetachPolicyFromUser(request *DetachPolicyFromUserRequest) (_result *DetachPolicyFromUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachPolicyFromUserResponse{}
	_body, _err := client.DetachPolicyFromUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetAccessKeyLastUsedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessKeyLastUsedResponse
func (client *Client) GetAccessKeyLastUsedWithOptions(request *GetAccessKeyLastUsedRequest, runtime *util.RuntimeOptions) (_result *GetAccessKeyLastUsedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserAccessKeyId)) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAccessKeyLastUsed"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccessKeyLastUsedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetAccessKeyLastUsedRequest
//
// @return GetAccessKeyLastUsedResponse
func (client *Client) GetAccessKeyLastUsed(request *GetAccessKeyLastUsedRequest) (_result *GetAccessKeyLastUsedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccessKeyLastUsedResponse{}
	_body, _err := client.GetAccessKeyLastUsedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the alias of an Alibaba Cloud account.
//
// @param request - GetAccountAliasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountAliasResponse
func (client *Client) GetAccountAliasWithOptions(runtime *util.RuntimeOptions) (_result *GetAccountAliasResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetAccountAlias"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccountAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the alias of an Alibaba Cloud account.
//
// @return GetAccountAliasResponse
func (client *Client) GetAccountAlias() (_result *GetAccountAliasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccountAliasResponse{}
	_body, _err := client.GetAccountAliasWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a Resource Access Management (RAM) user group.
//
// @param request - GetGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGroupResponse
func (client *Client) GetGroupWithOptions(request *GetGroupRequest, runtime *util.RuntimeOptions) (_result *GetGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGroup"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a Resource Access Management (RAM) user group.
//
// @param request - GetGroupRequest
//
// @return GetGroupResponse
func (client *Client) GetGroup(request *GetGroupRequest) (_result *GetGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGroupResponse{}
	_body, _err := client.GetGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the logon configurations of a Resource Access Management (RAM) user.
//
// @param request - GetLoginProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLoginProfileResponse
func (client *Client) GetLoginProfileWithOptions(request *GetLoginProfileRequest, runtime *util.RuntimeOptions) (_result *GetLoginProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLoginProfile"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLoginProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the logon configurations of a Resource Access Management (RAM) user.
//
// @param request - GetLoginProfileRequest
//
// @return GetLoginProfileResponse
func (client *Client) GetLoginProfile(request *GetLoginProfileRequest) (_result *GetLoginProfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLoginProfileResponse{}
	_body, _err := client.GetLoginProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the password policy of Resource Access Management (RAM) users, including the password strength.
//
// @param request - GetPasswordPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPasswordPolicyResponse
func (client *Client) GetPasswordPolicyWithOptions(runtime *util.RuntimeOptions) (_result *GetPasswordPolicyResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetPasswordPolicy"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPasswordPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the password policy of Resource Access Management (RAM) users, including the password strength.
//
// @return GetPasswordPolicyResponse
func (client *Client) GetPasswordPolicy() (_result *GetPasswordPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPasswordPolicyResponse{}
	_body, _err := client.GetPasswordPolicyWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a policy.
//
// @param request - GetPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPolicyResponse
func (client *Client) GetPolicyWithOptions(request *GetPolicyRequest, runtime *util.RuntimeOptions) (_result *GetPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyType)) {
		query["PolicyType"] = request.PolicyType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPolicy"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a policy.
//
// @param request - GetPolicyRequest
//
// @return GetPolicyResponse
func (client *Client) GetPolicy(request *GetPolicyRequest) (_result *GetPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPolicyResponse{}
	_body, _err := client.GetPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a policy version.
//
// @param request - GetPolicyVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPolicyVersionResponse
func (client *Client) GetPolicyVersionWithOptions(request *GetPolicyVersionRequest, runtime *util.RuntimeOptions) (_result *GetPolicyVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyType)) {
		query["PolicyType"] = request.PolicyType
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPolicyVersion"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPolicyVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a policy version.
//
// @param request - GetPolicyVersionRequest
//
// @return GetPolicyVersionResponse
func (client *Client) GetPolicyVersion(request *GetPolicyVersionRequest) (_result *GetPolicyVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPolicyVersionResponse{}
	_body, _err := client.GetPolicyVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a Resource Access Management (RAM) role.
//
// @param request - GetRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRoleResponse
func (client *Client) GetRoleWithOptions(request *GetRoleRequest, runtime *util.RuntimeOptions) (_result *GetRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		query["RoleName"] = request.RoleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRole"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a Resource Access Management (RAM) role.
//
// @param request - GetRoleRequest
//
// @return GetRoleResponse
func (client *Client) GetRole(request *GetRoleRequest) (_result *GetRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRoleResponse{}
	_body, _err := client.GetRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the security preferences.
//
// @param request - GetSecurityPreferenceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecurityPreferenceResponse
func (client *Client) GetSecurityPreferenceWithOptions(runtime *util.RuntimeOptions) (_result *GetSecurityPreferenceResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetSecurityPreference"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSecurityPreferenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the security preferences.
//
// @return GetSecurityPreferenceResponse
func (client *Client) GetSecurityPreference() (_result *GetSecurityPreferenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSecurityPreferenceResponse{}
	_body, _err := client.GetSecurityPreferenceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a Resource Access Management (RAM) user.
//
// Description:
//
// This topic provides an example on how to query information about the RAM user `alice`.
//
// @param request - GetUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserResponse
func (client *Client) GetUserWithOptions(request *GetUserRequest, runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUser"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a Resource Access Management (RAM) user.
//
// Description:
//
// This topic provides an example on how to query information about the RAM user `alice`.
//
// @param request - GetUserRequest
//
// @return GetUserResponse
func (client *Client) GetUser(request *GetUserRequest) (_result *GetUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the multi-factor authentication (MFA) device that is bound to a Resource Access Management (RAM) user.
//
// @param request - GetUserMFAInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserMFAInfoResponse
func (client *Client) GetUserMFAInfoWithOptions(request *GetUserMFAInfoRequest, runtime *util.RuntimeOptions) (_result *GetUserMFAInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserMFAInfo"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserMFAInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the multi-factor authentication (MFA) device that is bound to a Resource Access Management (RAM) user.
//
// @param request - GetUserMFAInfoRequest
//
// @return GetUserMFAInfoResponse
func (client *Client) GetUserMFAInfo(request *GetUserMFAInfoRequest) (_result *GetUserMFAInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserMFAInfoResponse{}
	_body, _err := client.GetUserMFAInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all AccessKey pairs that belong to a Resource Access Management (RAM) user.
//
// @param request - ListAccessKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccessKeysResponse
func (client *Client) ListAccessKeysWithOptions(request *ListAccessKeysRequest, runtime *util.RuntimeOptions) (_result *ListAccessKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAccessKeys"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAccessKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all AccessKey pairs that belong to a Resource Access Management (RAM) user.
//
// @param request - ListAccessKeysRequest
//
// @return ListAccessKeysResponse
func (client *Client) ListAccessKeys(request *ListAccessKeysRequest) (_result *ListAccessKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAccessKeysResponse{}
	_body, _err := client.ListAccessKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the entities to which a policy is attached.
//
// @param request - ListEntitiesForPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEntitiesForPolicyResponse
func (client *Client) ListEntitiesForPolicyWithOptions(request *ListEntitiesForPolicyRequest, runtime *util.RuntimeOptions) (_result *ListEntitiesForPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyType)) {
		query["PolicyType"] = request.PolicyType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEntitiesForPolicy"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEntitiesForPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the entities to which a policy is attached.
//
// @param request - ListEntitiesForPolicyRequest
//
// @return ListEntitiesForPolicyResponse
func (client *Client) ListEntitiesForPolicy(request *ListEntitiesForPolicyRequest) (_result *ListEntitiesForPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEntitiesForPolicyResponse{}
	_body, _err := client.ListEntitiesForPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Resource Access Management (RAM) user groups.
//
// @param request - ListGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupsResponse
func (client *Client) ListGroupsWithOptions(request *ListGroupsRequest, runtime *util.RuntimeOptions) (_result *ListGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["Marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxItems)) {
		query["MaxItems"] = request.MaxItems
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGroups"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Resource Access Management (RAM) user groups.
//
// @param request - ListGroupsRequest
//
// @return ListGroupsResponse
func (client *Client) ListGroups(request *ListGroupsRequest) (_result *ListGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGroupsResponse{}
	_body, _err := client.ListGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Resource Access Management (RAM) user groups to which a RAM user belongs.
//
// Description:
//
// ### [](#)
//
// This topic provides an example on how to query the RAM user groups to which the RAM user `Alice` belongs. The response shows that `Alice` belongs to the RAM user group named `Dev-Team`.
//
// @param request - ListGroupsForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupsForUserResponse
func (client *Client) ListGroupsForUserWithOptions(request *ListGroupsForUserRequest, runtime *util.RuntimeOptions) (_result *ListGroupsForUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGroupsForUser"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGroupsForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Resource Access Management (RAM) user groups to which a RAM user belongs.
//
// Description:
//
// ### [](#)
//
// This topic provides an example on how to query the RAM user groups to which the RAM user `Alice` belongs. The response shows that `Alice` belongs to the RAM user group named `Dev-Team`.
//
// @param request - ListGroupsForUserRequest
//
// @return ListGroupsForUserResponse
func (client *Client) ListGroupsForUser(request *ListGroupsForUserRequest) (_result *ListGroupsForUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGroupsForUserResponse{}
	_body, _err := client.ListGroupsForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of policies.
//
// @param tmpReq - ListPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPoliciesResponse
func (client *Client) ListPoliciesWithOptions(tmpReq *ListPoliciesRequest, runtime *util.RuntimeOptions) (_result *ListPoliciesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListPoliciesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tag)) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, tea.String("Tag"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["Marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxItems)) {
		query["MaxItems"] = request.MaxItems
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyType)) {
		query["PolicyType"] = request.PolicyType
	}

	if !tea.BoolValue(util.IsUnset(request.TagShrink)) {
		query["Tag"] = request.TagShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPolicies"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of policies.
//
// @param request - ListPoliciesRequest
//
// @return ListPoliciesResponse
func (client *Client) ListPolicies(request *ListPoliciesRequest) (_result *ListPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPoliciesResponse{}
	_body, _err := client.ListPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the policies that are attached to a Resource Access Management (RAM) user group.
//
// @param request - ListPoliciesForGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPoliciesForGroupResponse
func (client *Client) ListPoliciesForGroupWithOptions(request *ListPoliciesForGroupRequest, runtime *util.RuntimeOptions) (_result *ListPoliciesForGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPoliciesForGroup"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPoliciesForGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the policies that are attached to a Resource Access Management (RAM) user group.
//
// @param request - ListPoliciesForGroupRequest
//
// @return ListPoliciesForGroupResponse
func (client *Client) ListPoliciesForGroup(request *ListPoliciesForGroupRequest) (_result *ListPoliciesForGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPoliciesForGroupResponse{}
	_body, _err := client.ListPoliciesForGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the policies that are attached to a Resource Access Management (RAM) role.
//
// @param request - ListPoliciesForRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPoliciesForRoleResponse
func (client *Client) ListPoliciesForRoleWithOptions(request *ListPoliciesForRoleRequest, runtime *util.RuntimeOptions) (_result *ListPoliciesForRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		query["RoleName"] = request.RoleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPoliciesForRole"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPoliciesForRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the policies that are attached to a Resource Access Management (RAM) role.
//
// @param request - ListPoliciesForRoleRequest
//
// @return ListPoliciesForRoleResponse
func (client *Client) ListPoliciesForRole(request *ListPoliciesForRoleRequest) (_result *ListPoliciesForRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPoliciesForRoleResponse{}
	_body, _err := client.ListPoliciesForRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the policies that are attached to a RAM user.
//
// Description:
//
// > You can call this operation to query only the policies that are attached to Alibaba Cloud accounts. You cannot query the policies that are attached to resource groups.
//
// @param request - ListPoliciesForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPoliciesForUserResponse
func (client *Client) ListPoliciesForUserWithOptions(request *ListPoliciesForUserRequest, runtime *util.RuntimeOptions) (_result *ListPoliciesForUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPoliciesForUser"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPoliciesForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the policies that are attached to a RAM user.
//
// Description:
//
// > You can call this operation to query only the policies that are attached to Alibaba Cloud accounts. You cannot query the policies that are attached to resource groups.
//
// @param request - ListPoliciesForUserRequest
//
// @return ListPoliciesForUserResponse
func (client *Client) ListPoliciesForUser(request *ListPoliciesForUserRequest) (_result *ListPoliciesForUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPoliciesForUserResponse{}
	_body, _err := client.ListPoliciesForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the versions of a policy.
//
// @param request - ListPolicyVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPolicyVersionsResponse
func (client *Client) ListPolicyVersionsWithOptions(request *ListPolicyVersionsRequest, runtime *util.RuntimeOptions) (_result *ListPolicyVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyType)) {
		query["PolicyType"] = request.PolicyType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPolicyVersions"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPolicyVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the versions of a policy.
//
// @param request - ListPolicyVersionsRequest
//
// @return ListPolicyVersionsResponse
func (client *Client) ListPolicyVersions(request *ListPolicyVersionsRequest) (_result *ListPolicyVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPolicyVersionsResponse{}
	_body, _err := client.ListPolicyVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all Resource Access Management (RAM) roles.
//
// @param tmpReq - ListRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRolesResponse
func (client *Client) ListRolesWithOptions(tmpReq *ListRolesRequest, runtime *util.RuntimeOptions) (_result *ListRolesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListRolesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tag)) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, tea.String("Tag"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["Marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxItems)) {
		query["MaxItems"] = request.MaxItems
	}

	if !tea.BoolValue(util.IsUnset(request.TagShrink)) {
		query["Tag"] = request.TagShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRoles"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all Resource Access Management (RAM) roles.
//
// @param request - ListRolesRequest
//
// @return ListRolesResponse
func (client *Client) ListRoles(request *ListRolesRequest) (_result *ListRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRolesResponse{}
	_body, _err := client.ListRolesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags that are added to resources.
//
// @param tmpReq - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(tmpReq *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListTagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceNames)) {
		request.ResourceNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceNames, tea.String("ResourceNames"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tag)) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, tea.String("Tag"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceNamesShrink)) {
		query["ResourceNames"] = request.ResourceNamesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagShrink)) {
		query["Tag"] = request.TagShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2015-05-01"),
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

// Summary:
//
// Queries the tags that are added to resources.
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

// Summary:
//
// Queries the information about all RAM users.
//
// @param request - ListUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *util.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["Marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxItems)) {
		query["MaxItems"] = request.MaxItems
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsers"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about all RAM users.
//
// @param request - ListUsersRequest
//
// @return ListUsersResponse
func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Resource Access Management (RAM) users in a RAM user group.
//
// @param request - ListUsersForGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersForGroupResponse
func (client *Client) ListUsersForGroupWithOptions(request *ListUsersForGroupRequest, runtime *util.RuntimeOptions) (_result *ListUsersForGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["Marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxItems)) {
		query["MaxItems"] = request.MaxItems
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsersForGroup"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUsersForGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Resource Access Management (RAM) users in a RAM user group.
//
// @param request - ListUsersForGroupRequest
//
// @return ListUsersForGroupResponse
func (client *Client) ListUsersForGroup(request *ListUsersForGroupRequest) (_result *ListUsersForGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUsersForGroupResponse{}
	_body, _err := client.ListUsersForGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries multi-factor authentication (MFA) devices.
//
// @param request - ListVirtualMFADevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVirtualMFADevicesResponse
func (client *Client) ListVirtualMFADevicesWithOptions(runtime *util.RuntimeOptions) (_result *ListVirtualMFADevicesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListVirtualMFADevices"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVirtualMFADevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries multi-factor authentication (MFA) devices.
//
// @return ListVirtualMFADevicesResponse
func (client *Client) ListVirtualMFADevices() (_result *ListVirtualMFADevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVirtualMFADevicesResponse{}
	_body, _err := client.ListVirtualMFADevicesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a Resource Access Management (RAM) user from a RAM user group.
//
// @param request - RemoveUserFromGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUserFromGroupResponse
func (client *Client) RemoveUserFromGroupWithOptions(request *RemoveUserFromGroupRequest, runtime *util.RuntimeOptions) (_result *RemoveUserFromGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveUserFromGroup"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveUserFromGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a Resource Access Management (RAM) user from a RAM user group.
//
// @param request - RemoveUserFromGroupRequest
//
// @return RemoveUserFromGroupResponse
func (client *Client) RemoveUserFromGroup(request *RemoveUserFromGroupRequest) (_result *RemoveUserFromGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveUserFromGroupResponse{}
	_body, _err := client.RemoveUserFromGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures an alias for an Alibaba Cloud account.
//
// @param request - SetAccountAliasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetAccountAliasResponse
func (client *Client) SetAccountAliasWithOptions(request *SetAccountAliasRequest, runtime *util.RuntimeOptions) (_result *SetAccountAliasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountAlias)) {
		query["AccountAlias"] = request.AccountAlias
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetAccountAlias"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetAccountAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures an alias for an Alibaba Cloud account.
//
// @param request - SetAccountAliasRequest
//
// @return SetAccountAliasResponse
func (client *Client) SetAccountAlias(request *SetAccountAliasRequest) (_result *SetAccountAliasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetAccountAliasResponse{}
	_body, _err := client.SetAccountAliasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Specifies a version for a policy as the default version.
//
// @param request - SetDefaultPolicyVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDefaultPolicyVersionResponse
func (client *Client) SetDefaultPolicyVersionWithOptions(request *SetDefaultPolicyVersionRequest, runtime *util.RuntimeOptions) (_result *SetDefaultPolicyVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDefaultPolicyVersion"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDefaultPolicyVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Specifies a version for a policy as the default version.
//
// @param request - SetDefaultPolicyVersionRequest
//
// @return SetDefaultPolicyVersionResponse
func (client *Client) SetDefaultPolicyVersion(request *SetDefaultPolicyVersionRequest) (_result *SetDefaultPolicyVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDefaultPolicyVersionResponse{}
	_body, _err := client.SetDefaultPolicyVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the password policy for Resource Access Management (RAM) users, including the password strength.
//
// @param request - SetPasswordPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPasswordPolicyResponse
func (client *Client) SetPasswordPolicyWithOptions(request *SetPasswordPolicyRequest, runtime *util.RuntimeOptions) (_result *SetPasswordPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HardExpiry)) {
		query["HardExpiry"] = request.HardExpiry
	}

	if !tea.BoolValue(util.IsUnset(request.MaxLoginAttemps)) {
		query["MaxLoginAttemps"] = request.MaxLoginAttemps
	}

	if !tea.BoolValue(util.IsUnset(request.MaxPasswordAge)) {
		query["MaxPasswordAge"] = request.MaxPasswordAge
	}

	if !tea.BoolValue(util.IsUnset(request.MinimumPasswordLength)) {
		query["MinimumPasswordLength"] = request.MinimumPasswordLength
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordReusePrevention)) {
		query["PasswordReusePrevention"] = request.PasswordReusePrevention
	}

	if !tea.BoolValue(util.IsUnset(request.RequireLowercaseCharacters)) {
		query["RequireLowercaseCharacters"] = request.RequireLowercaseCharacters
	}

	if !tea.BoolValue(util.IsUnset(request.RequireNumbers)) {
		query["RequireNumbers"] = request.RequireNumbers
	}

	if !tea.BoolValue(util.IsUnset(request.RequireSymbols)) {
		query["RequireSymbols"] = request.RequireSymbols
	}

	if !tea.BoolValue(util.IsUnset(request.RequireUppercaseCharacters)) {
		query["RequireUppercaseCharacters"] = request.RequireUppercaseCharacters
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetPasswordPolicy"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetPasswordPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the password policy for Resource Access Management (RAM) users, including the password strength.
//
// @param request - SetPasswordPolicyRequest
//
// @return SetPasswordPolicyResponse
func (client *Client) SetPasswordPolicy(request *SetPasswordPolicyRequest) (_result *SetPasswordPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetPasswordPolicyResponse{}
	_body, _err := client.SetPasswordPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the security preferences.
//
// @param request - SetSecurityPreferenceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetSecurityPreferenceResponse
func (client *Client) SetSecurityPreferenceWithOptions(request *SetSecurityPreferenceRequest, runtime *util.RuntimeOptions) (_result *SetSecurityPreferenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowUserToChangePassword)) {
		query["AllowUserToChangePassword"] = request.AllowUserToChangePassword
	}

	if !tea.BoolValue(util.IsUnset(request.AllowUserToManageAccessKeys)) {
		query["AllowUserToManageAccessKeys"] = request.AllowUserToManageAccessKeys
	}

	if !tea.BoolValue(util.IsUnset(request.AllowUserToManageMFADevices)) {
		query["AllowUserToManageMFADevices"] = request.AllowUserToManageMFADevices
	}

	if !tea.BoolValue(util.IsUnset(request.AllowUserToManagePublicKeys)) {
		query["AllowUserToManagePublicKeys"] = request.AllowUserToManagePublicKeys
	}

	if !tea.BoolValue(util.IsUnset(request.EnableSaveMFATicket)) {
		query["EnableSaveMFATicket"] = request.EnableSaveMFATicket
	}

	if !tea.BoolValue(util.IsUnset(request.LoginNetworkMasks)) {
		query["LoginNetworkMasks"] = request.LoginNetworkMasks
	}

	if !tea.BoolValue(util.IsUnset(request.LoginSessionDuration)) {
		query["LoginSessionDuration"] = request.LoginSessionDuration
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetSecurityPreference"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetSecurityPreferenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the security preferences.
//
// @param request - SetSecurityPreferenceRequest
//
// @return SetSecurityPreferenceResponse
func (client *Client) SetSecurityPreference(request *SetSecurityPreferenceRequest) (_result *SetSecurityPreferenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetSecurityPreferenceResponse{}
	_body, _err := client.SetSecurityPreferenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds tags to resources.
//
// @param tmpReq - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(tmpReq *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &TagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceNames)) {
		request.ResourceNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceNames, tea.String("ResourceNames"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tag)) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, tea.String("Tag"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceNamesShrink)) {
		query["ResourceNames"] = request.ResourceNamesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagShrink)) {
		query["Tag"] = request.TagShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2015-05-01"),
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

// Summary:
//
// Adds tags to resources.
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
// Unbinds a multi-factor authentication (MFA) device from a Resource Access Management (RAM) user.
//
// @param request - UnbindMFADeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindMFADeviceResponse
func (client *Client) UnbindMFADeviceWithOptions(request *UnbindMFADeviceRequest, runtime *util.RuntimeOptions) (_result *UnbindMFADeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindMFADevice"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindMFADeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds a multi-factor authentication (MFA) device from a Resource Access Management (RAM) user.
//
// @param request - UnbindMFADeviceRequest
//
// @return UnbindMFADeviceResponse
func (client *Client) UnbindMFADevice(request *UnbindMFADeviceRequest) (_result *UnbindMFADeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindMFADeviceResponse{}
	_body, _err := client.UnbindMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes tags from resources.
//
// @param tmpReq - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(tmpReq *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UntagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceNames)) {
		request.ResourceNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceNames, tea.String("ResourceNames"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TagKeys)) {
		request.TagKeysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagKeys, tea.String("TagKeys"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceNamesShrink)) {
		query["ResourceNames"] = request.ResourceNamesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKeysShrink)) {
		query["TagKeys"] = request.TagKeysShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2015-05-01"),
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

// Summary:
//
// Removes tags from resources.
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

// Summary:
//
// Changes the status of an AccessKey pair that belongs to a Resource Access Management (RAM) user.
//
// @param request - UpdateAccessKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAccessKeyResponse
func (client *Client) UpdateAccessKeyWithOptions(request *UpdateAccessKeyRequest, runtime *util.RuntimeOptions) (_result *UpdateAccessKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UserAccessKeyId)) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAccessKey"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAccessKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the status of an AccessKey pair that belongs to a Resource Access Management (RAM) user.
//
// @param request - UpdateAccessKeyRequest
//
// @return UpdateAccessKeyResponse
func (client *Client) UpdateAccessKey(request *UpdateAccessKeyRequest) (_result *UpdateAccessKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAccessKeyResponse{}
	_body, _err := client.UpdateAccessKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a Resource Access Management (RAM) user group.
//
// @param request - UpdateGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGroupResponse
func (client *Client) UpdateGroupWithOptions(request *UpdateGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.NewComments)) {
		query["NewComments"] = request.NewComments
	}

	if !tea.BoolValue(util.IsUnset(request.NewGroupName)) {
		query["NewGroupName"] = request.NewGroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGroup"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a Resource Access Management (RAM) user group.
//
// @param request - UpdateGroupRequest
//
// @return UpdateGroupResponse
func (client *Client) UpdateGroup(request *UpdateGroupRequest) (_result *UpdateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGroupResponse{}
	_body, _err := client.UpdateGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the logon configurations of a Resource Access Management (RAM) user.
//
// @param request - UpdateLoginProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLoginProfileResponse
func (client *Client) UpdateLoginProfileWithOptions(request *UpdateLoginProfileRequest, runtime *util.RuntimeOptions) (_result *UpdateLoginProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MFABindRequired)) {
		query["MFABindRequired"] = request.MFABindRequired
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordResetRequired)) {
		query["PasswordResetRequired"] = request.PasswordResetRequired
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLoginProfile"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLoginProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the logon configurations of a Resource Access Management (RAM) user.
//
// @param request - UpdateLoginProfileRequest
//
// @return UpdateLoginProfileResponse
func (client *Client) UpdateLoginProfile(request *UpdateLoginProfileRequest) (_result *UpdateLoginProfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLoginProfileResponse{}
	_body, _err := client.UpdateLoginProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the description of a custom policy.
//
// Description:
//
// ### [](#)
//
// You can call this operation to modify only the description of a custom policy. You cannot modify the description of a system policy.
//
// @param request - UpdatePolicyDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePolicyDescriptionResponse
func (client *Client) UpdatePolicyDescriptionWithOptions(request *UpdatePolicyDescriptionRequest, runtime *util.RuntimeOptions) (_result *UpdatePolicyDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewDescription)) {
		query["NewDescription"] = request.NewDescription
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePolicyDescription"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePolicyDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of a custom policy.
//
// Description:
//
// ### [](#)
//
// You can call this operation to modify only the description of a custom policy. You cannot modify the description of a system policy.
//
// @param request - UpdatePolicyDescriptionRequest
//
// @return UpdatePolicyDescriptionResponse
func (client *Client) UpdatePolicyDescription(request *UpdatePolicyDescriptionRequest) (_result *UpdatePolicyDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePolicyDescriptionResponse{}
	_body, _err := client.UpdatePolicyDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies information about a Resource Access Management (RAM) role.
//
// Description:
//
// This topic provides an example on how to change the description of `ECSAdmin` to `ECS administrator`.
//
// @param request - UpdateRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRoleResponse
func (client *Client) UpdateRoleWithOptions(request *UpdateRoleRequest, runtime *util.RuntimeOptions) (_result *UpdateRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewAssumeRolePolicyDocument)) {
		query["NewAssumeRolePolicyDocument"] = request.NewAssumeRolePolicyDocument
	}

	if !tea.BoolValue(util.IsUnset(request.NewDescription)) {
		query["NewDescription"] = request.NewDescription
	}

	if !tea.BoolValue(util.IsUnset(request.NewMaxSessionDuration)) {
		query["NewMaxSessionDuration"] = request.NewMaxSessionDuration
	}

	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		query["RoleName"] = request.RoleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRole"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies information about a Resource Access Management (RAM) role.
//
// Description:
//
// This topic provides an example on how to change the description of `ECSAdmin` to `ECS administrator`.
//
// @param request - UpdateRoleRequest
//
// @return UpdateRoleResponse
func (client *Client) UpdateRole(request *UpdateRoleRequest) (_result *UpdateRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRoleResponse{}
	_body, _err := client.UpdateRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies information about a Resource Access Management (RAM) user.
//
// Description:
//
// This topic provides an example on how to change the name of a RAM user from `zhangq****` to `xiaoq****`.
//
// @param request - UpdateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserResponse
func (client *Client) UpdateUserWithOptions(request *UpdateUserRequest, runtime *util.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewComments)) {
		query["NewComments"] = request.NewComments
	}

	if !tea.BoolValue(util.IsUnset(request.NewDisplayName)) {
		query["NewDisplayName"] = request.NewDisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.NewEmail)) {
		query["NewEmail"] = request.NewEmail
	}

	if !tea.BoolValue(util.IsUnset(request.NewMobilePhone)) {
		query["NewMobilePhone"] = request.NewMobilePhone
	}

	if !tea.BoolValue(util.IsUnset(request.NewUserName)) {
		query["NewUserName"] = request.NewUserName
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUser"),
		Version:     tea.String("2015-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies information about a Resource Access Management (RAM) user.
//
// Description:
//
// This topic provides an example on how to change the name of a RAM user from `zhangq****` to `xiaoq****`.
//
// @param request - UpdateUserRequest
//
// @return UpdateUserResponse
func (client *Client) UpdateUser(request *UpdateUserRequest) (_result *UpdateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserResponse{}
	_body, _err := client.UpdateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
