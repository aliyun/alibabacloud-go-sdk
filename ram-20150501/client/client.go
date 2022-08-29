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

type AddUserToGroupRequest struct {
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	UserName  *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddUserToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	GroupName  *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachPolicyToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	RoleName   *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachPolicyToRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	UserName   *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachPolicyToUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AuthenticationCode1 *string `json:"AuthenticationCode1,omitempty" xml:"AuthenticationCode1,omitempty"`
	AuthenticationCode2 *string `json:"AuthenticationCode2,omitempty" xml:"AuthenticationCode2,omitempty"`
	SerialNumber        *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	UserName            *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BindMFADeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	NewPassword *string `json:"NewPassword,omitempty" xml:"NewPassword,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ChangePasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ClearAccountAliasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AccessKey *CreateAccessKeyResponseBodyAccessKey `json:"AccessKey,omitempty" xml:"AccessKey,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	CreateDate      *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAccessKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Comments  *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
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
	Group     *CreateGroupResponseBodyGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Comments   *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName  *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	MFABindRequired       *bool   `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty"`
	Password              *string `json:"Password,omitempty" xml:"Password,omitempty"`
	PasswordResetRequired *bool   `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty"`
	UserName              *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	LoginProfile *CreateLoginProfileResponseBodyLoginProfile `json:"LoginProfile,omitempty" xml:"LoginProfile,omitempty" type:"Struct"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateDate            *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	MFABindRequired       *bool   `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty"`
	PasswordResetRequired *bool   `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty"`
	UserName              *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateLoginProfileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	PolicyName     *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
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

type CreatePolicyResponseBody struct {
	Policy    *CreatePolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateDate     *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DefaultVersion *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PolicyName     *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType     *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreatePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	PolicyName     *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	RotateStrategy *string `json:"RotateStrategy,omitempty" xml:"RotateStrategy,omitempty"`
	SetAsDefault   *bool   `json:"SetAsDefault,omitempty" xml:"SetAsDefault,omitempty"`
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
	PolicyVersion *CreatePolicyVersionResponseBodyPolicyVersion `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty" type:"Struct"`
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateDate       *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	IsDefaultVersion *bool   `json:"IsDefaultVersion,omitempty" xml:"IsDefaultVersion,omitempty"`
	PolicyDocument   *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	VersionId        *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreatePolicyVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AssumeRolePolicyDocument *string `json:"AssumeRolePolicyDocument,omitempty" xml:"AssumeRolePolicyDocument,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	MaxSessionDuration       *int64  `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty"`
	RoleName                 *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
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

type CreateRoleResponseBody struct {
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Role      *CreateRoleResponseBodyRole `json:"Role,omitempty" xml:"Role,omitempty" type:"Struct"`
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
	Arn                      *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	AssumeRolePolicyDocument *string `json:"AssumeRolePolicyDocument,omitempty" xml:"AssumeRolePolicyDocument,omitempty"`
	CreateDate               *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	MaxSessionDuration       *int64  `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty"`
	RoleId                   *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	RoleName                 *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Comments    *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	User      *CreateUserResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
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
	Comments    *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	CreateDate  *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Base32StringSeed *string `json:"Base32StringSeed,omitempty" xml:"Base32StringSeed,omitempty"`
	QRCodePNG        *string `json:"QRCodePNG,omitempty" xml:"QRCodePNG,omitempty"`
	SerialNumber     *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVirtualMFADeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DeleteAccessKeyRequest struct {
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	UserName        *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAccessKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteLoginProfileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s DeletePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyRequest) GoString() string {
	return s.String()
}

func (s *DeletePolicyRequest) SetPolicyName(v string) *DeletePolicyRequest {
	s.PolicyName = &v
	return s
}

type DeletePolicyResponseBody struct {
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeletePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	VersionId  *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeletePolicyVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVirtualMFADeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	GroupName  *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachPolicyFromGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	RoleName   *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachPolicyFromRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	UserName   *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachPolicyFromUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	UserName        *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	AccessKeyLastUsed *GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed `json:"AccessKeyLastUsed,omitempty" xml:"AccessKeyLastUsed,omitempty" type:"Struct"`
	RequestId         *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAccessKeyLastUsedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AccountAlias *string `json:"AccountAlias,omitempty" xml:"AccountAlias,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAccountAliasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Group     *GetGroupResponseBodyGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Comments   *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName  *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	LoginProfile *GetLoginProfileResponseBodyLoginProfile `json:"LoginProfile,omitempty" xml:"LoginProfile,omitempty" type:"Struct"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateDate            *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	MFABindRequired       *bool   `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty"`
	PasswordResetRequired *bool   `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty"`
	UserName              *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetLoginProfileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PasswordPolicy *GetPasswordPolicyResponseBodyPasswordPolicy `json:"PasswordPolicy,omitempty" xml:"PasswordPolicy,omitempty" type:"Struct"`
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	HardExpiry                 *bool  `json:"HardExpiry,omitempty" xml:"HardExpiry,omitempty"`
	MaxLoginAttemps            *int32 `json:"MaxLoginAttemps,omitempty" xml:"MaxLoginAttemps,omitempty"`
	MaxPasswordAge             *int32 `json:"MaxPasswordAge,omitempty" xml:"MaxPasswordAge,omitempty"`
	MinimumPasswordLength      *int32 `json:"MinimumPasswordLength,omitempty" xml:"MinimumPasswordLength,omitempty"`
	PasswordReusePrevention    *int32 `json:"PasswordReusePrevention,omitempty" xml:"PasswordReusePrevention,omitempty"`
	RequireLowercaseCharacters *bool  `json:"RequireLowercaseCharacters,omitempty" xml:"RequireLowercaseCharacters,omitempty"`
	RequireNumbers             *bool  `json:"RequireNumbers,omitempty" xml:"RequireNumbers,omitempty"`
	RequireSymbols             *bool  `json:"RequireSymbols,omitempty" xml:"RequireSymbols,omitempty"`
	RequireUppercaseCharacters *bool  `json:"RequireUppercaseCharacters,omitempty" xml:"RequireUppercaseCharacters,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPasswordPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
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
	DefaultPolicyVersion *GetPolicyResponseBodyDefaultPolicyVersion `json:"DefaultPolicyVersion,omitempty" xml:"DefaultPolicyVersion,omitempty" type:"Struct"`
	Policy               *GetPolicyResponseBodyPolicy               `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	RequestId            *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateDate       *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	IsDefaultVersion *bool   `json:"IsDefaultVersion,omitempty" xml:"IsDefaultVersion,omitempty"`
	PolicyDocument   *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	VersionId        *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
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
	AttachmentCount *int32  `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty"`
	CreateDate      *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DefaultVersion  *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PolicyDocument  *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	PolicyName      *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType      *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	UpdateDate      *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	VersionId  *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
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
	PolicyVersion *GetPolicyVersionResponseBodyPolicyVersion `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty" type:"Struct"`
	RequestId     *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateDate       *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	IsDefaultVersion *bool   `json:"IsDefaultVersion,omitempty" xml:"IsDefaultVersion,omitempty"`
	PolicyDocument   *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	VersionId        *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPolicyVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Role      *GetRoleResponseBodyRole `json:"Role,omitempty" xml:"Role,omitempty" type:"Struct"`
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
	Arn                      *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	AssumeRolePolicyDocument *string `json:"AssumeRolePolicyDocument,omitempty" xml:"AssumeRolePolicyDocument,omitempty"`
	CreateDate               *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	MaxSessionDuration       *int64  `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty"`
	RoleId                   *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	RoleName                 *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	UpdateDate               *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
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
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId          *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccessKeyPreference    *GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference    `json:"AccessKeyPreference,omitempty" xml:"AccessKeyPreference,omitempty" type:"Struct"`
	LoginProfilePreference *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference `json:"LoginProfilePreference,omitempty" xml:"LoginProfilePreference,omitempty" type:"Struct"`
	MFAPreference          *GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference          `json:"MFAPreference,omitempty" xml:"MFAPreference,omitempty" type:"Struct"`
	PublicKeyPreference    *GetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference    `json:"PublicKeyPreference,omitempty" xml:"PublicKeyPreference,omitempty" type:"Struct"`
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
	AllowUserToChangePassword *bool   `json:"AllowUserToChangePassword,omitempty" xml:"AllowUserToChangePassword,omitempty"`
	EnableSaveMFATicket       *bool   `json:"EnableSaveMFATicket,omitempty" xml:"EnableSaveMFATicket,omitempty"`
	LoginNetworkMasks         *string `json:"LoginNetworkMasks,omitempty" xml:"LoginNetworkMasks,omitempty"`
	LoginSessionDuration      *int32  `json:"LoginSessionDuration,omitempty" xml:"LoginSessionDuration,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSecurityPreferenceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	User      *GetUserResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
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
	Comments      *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	CreateDate    *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DisplayName   *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email         *string `json:"Email,omitempty" xml:"Email,omitempty"`
	LastLoginDate *string `json:"LastLoginDate,omitempty" xml:"LastLoginDate,omitempty"`
	MobilePhone   *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	UpdateDate    *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName      *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	MFADevice *GetUserMFAInfoResponseBodyMFADevice `json:"MFADevice,omitempty" xml:"MFADevice,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserMFAInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AccessKeys *ListAccessKeysResponseBodyAccessKeys `json:"AccessKeys,omitempty" xml:"AccessKeys,omitempty" type:"Struct"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	CreateDate  *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAccessKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
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
	Groups    *ListEntitiesForPolicyResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Struct"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Roles     *ListEntitiesForPolicyResponseBodyRoles  `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Struct"`
	Users     *ListEntitiesForPolicyResponseBodyUsers  `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
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
	AttachDate *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty"`
	Comments   *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	GroupName  *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
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
	Arn         *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	AttachDate  *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RoleId      *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	RoleName    *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
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
	AttachDate  *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListEntitiesForPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Marker   *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	MaxItems *int32  `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
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
	Groups      *ListGroupsResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Struct"`
	IsTruncated *bool                         `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	Marker      *string                       `json:"Marker,omitempty" xml:"Marker,omitempty"`
	RequestId   *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Comments   *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName  *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Groups    *ListGroupsForUserResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Comments  *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	JoinDate  *string `json:"JoinDate,omitempty" xml:"JoinDate,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListGroupsForUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Marker     *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	MaxItems   *int32  `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
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

type ListPoliciesResponseBody struct {
	IsTruncated *bool                             `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	Marker      *string                           `json:"Marker,omitempty" xml:"Marker,omitempty"`
	Policies    *ListPoliciesResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Struct"`
	RequestId   *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AttachmentCount *int32  `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty"`
	CreateDate      *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DefaultVersion  *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PolicyName      *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType      *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	UpdateDate      *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
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

func (s *ListPoliciesResponseBodyPoliciesPolicy) SetUpdateDate(v string) *ListPoliciesResponseBodyPoliciesPolicy {
	s.UpdateDate = &v
	return s
}

type ListPoliciesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Policies  *ListPoliciesForGroupResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AttachDate     *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty"`
	DefaultVersion *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PolicyName     *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType     *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPoliciesForGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Policies  *ListPoliciesForRoleResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Struct"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AttachDate     *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty"`
	DefaultVersion *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PolicyName     *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType     *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPoliciesForRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Policies  *ListPoliciesForUserResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Struct"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AttachDate     *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty"`
	DefaultVersion *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PolicyName     *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType     *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPoliciesForUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
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
	PolicyVersions *ListPolicyVersionsResponseBodyPolicyVersions `json:"PolicyVersions,omitempty" xml:"PolicyVersions,omitempty" type:"Struct"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateDate       *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	IsDefaultVersion *bool   `json:"IsDefaultVersion,omitempty" xml:"IsDefaultVersion,omitempty"`
	PolicyDocument   *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	VersionId        *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPolicyVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Marker   *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	MaxItems *int32  `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
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

type ListRolesResponseBody struct {
	IsTruncated *bool                       `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	Marker      *string                     `json:"Marker,omitempty" xml:"Marker,omitempty"`
	RequestId   *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Roles       *ListRolesResponseBodyRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Struct"`
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
	Arn                *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	CreateDate         *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	MaxSessionDuration *int64  `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty"`
	RoleId             *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	RoleName           *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	UpdateDate         *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
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

func (s *ListRolesResponseBodyRolesRole) SetUpdateDate(v string) *ListRolesResponseBodyRolesRole {
	s.UpdateDate = &v
	return s
}

type ListRolesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListUsersRequest struct {
	Marker   *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	MaxItems *int32  `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
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
	IsTruncated *bool                       `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	Marker      *string                     `json:"Marker,omitempty" xml:"Marker,omitempty"`
	RequestId   *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Users       *ListUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
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
	Comments    *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	CreateDate  *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	UpdateDate  *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Marker    *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	MaxItems  *int32  `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
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
	IsTruncated *bool                               `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	Marker      *string                             `json:"Marker,omitempty" xml:"Marker,omitempty"`
	RequestId   *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Users       *ListUsersForGroupResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
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
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	JoinDate    *string `json:"JoinDate,omitempty" xml:"JoinDate,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUsersForGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId         *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ActivateDate *string                                                                 `json:"ActivateDate,omitempty" xml:"ActivateDate,omitempty"`
	SerialNumber *string                                                                 `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	User         *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
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
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListVirtualMFADevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	UserName  *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveUserFromGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetAccountAliasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	VersionId  *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetDefaultPolicyVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	HardExpiry                 *bool  `json:"HardExpiry,omitempty" xml:"HardExpiry,omitempty"`
	MaxLoginAttemps            *int32 `json:"MaxLoginAttemps,omitempty" xml:"MaxLoginAttemps,omitempty"`
	MaxPasswordAge             *int32 `json:"MaxPasswordAge,omitempty" xml:"MaxPasswordAge,omitempty"`
	MinimumPasswordLength      *int32 `json:"MinimumPasswordLength,omitempty" xml:"MinimumPasswordLength,omitempty"`
	PasswordReusePrevention    *int32 `json:"PasswordReusePrevention,omitempty" xml:"PasswordReusePrevention,omitempty"`
	RequireLowercaseCharacters *bool  `json:"RequireLowercaseCharacters,omitempty" xml:"RequireLowercaseCharacters,omitempty"`
	RequireNumbers             *bool  `json:"RequireNumbers,omitempty" xml:"RequireNumbers,omitempty"`
	RequireSymbols             *bool  `json:"RequireSymbols,omitempty" xml:"RequireSymbols,omitempty"`
	RequireUppercaseCharacters *bool  `json:"RequireUppercaseCharacters,omitempty" xml:"RequireUppercaseCharacters,omitempty"`
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
	PasswordPolicy *SetPasswordPolicyResponseBodyPasswordPolicy `json:"PasswordPolicy,omitempty" xml:"PasswordPolicy,omitempty" type:"Struct"`
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	HardExpiry                 *bool  `json:"HardExpiry,omitempty" xml:"HardExpiry,omitempty"`
	MaxLoginAttemps            *int32 `json:"MaxLoginAttemps,omitempty" xml:"MaxLoginAttemps,omitempty"`
	MaxPasswordAge             *int32 `json:"MaxPasswordAge,omitempty" xml:"MaxPasswordAge,omitempty"`
	MinimumPasswordLength      *int32 `json:"MinimumPasswordLength,omitempty" xml:"MinimumPasswordLength,omitempty"`
	PasswordReusePrevention    *int32 `json:"PasswordReusePrevention,omitempty" xml:"PasswordReusePrevention,omitempty"`
	RequireLowercaseCharacters *bool  `json:"RequireLowercaseCharacters,omitempty" xml:"RequireLowercaseCharacters,omitempty"`
	RequireNumbers             *bool  `json:"RequireNumbers,omitempty" xml:"RequireNumbers,omitempty"`
	RequireSymbols             *bool  `json:"RequireSymbols,omitempty" xml:"RequireSymbols,omitempty"`
	RequireUppercaseCharacters *bool  `json:"RequireUppercaseCharacters,omitempty" xml:"RequireUppercaseCharacters,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetPasswordPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AllowUserToChangePassword   *bool   `json:"AllowUserToChangePassword,omitempty" xml:"AllowUserToChangePassword,omitempty"`
	AllowUserToManageAccessKeys *bool   `json:"AllowUserToManageAccessKeys,omitempty" xml:"AllowUserToManageAccessKeys,omitempty"`
	AllowUserToManageMFADevices *bool   `json:"AllowUserToManageMFADevices,omitempty" xml:"AllowUserToManageMFADevices,omitempty"`
	AllowUserToManagePublicKeys *bool   `json:"AllowUserToManagePublicKeys,omitempty" xml:"AllowUserToManagePublicKeys,omitempty"`
	EnableSaveMFATicket         *bool   `json:"EnableSaveMFATicket,omitempty" xml:"EnableSaveMFATicket,omitempty"`
	LoginNetworkMasks           *string `json:"LoginNetworkMasks,omitempty" xml:"LoginNetworkMasks,omitempty"`
	LoginSessionDuration        *int32  `json:"LoginSessionDuration,omitempty" xml:"LoginSessionDuration,omitempty"`
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
	RequestId          *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccessKeyPreference    *SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference    `json:"AccessKeyPreference,omitempty" xml:"AccessKeyPreference,omitempty" type:"Struct"`
	LoginProfilePreference *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference `json:"LoginProfilePreference,omitempty" xml:"LoginProfilePreference,omitempty" type:"Struct"`
	MFAPreference          *SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference          `json:"MFAPreference,omitempty" xml:"MFAPreference,omitempty" type:"Struct"`
	PublicKeyPreference    *SetSecurityPreferenceResponseBodySecurityPreferencePublicKeyPreference    `json:"PublicKeyPreference,omitempty" xml:"PublicKeyPreference,omitempty" type:"Struct"`
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
	AllowUserToChangePassword *bool   `json:"AllowUserToChangePassword,omitempty" xml:"AllowUserToChangePassword,omitempty"`
	EnableSaveMFATicket       *bool   `json:"EnableSaveMFATicket,omitempty" xml:"EnableSaveMFATicket,omitempty"`
	LoginNetworkMasks         *string `json:"LoginNetworkMasks,omitempty" xml:"LoginNetworkMasks,omitempty"`
	LoginSessionDuration      *int32  `json:"LoginSessionDuration,omitempty" xml:"LoginSessionDuration,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetSecurityPreferenceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type UnbindMFADeviceRequest struct {
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
	MFADevice *UnbindMFADeviceResponseBodyMFADevice `json:"MFADevice,omitempty" xml:"MFADevice,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnbindMFADeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type UpdateAccessKeyRequest struct {
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	UserName        *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAccessKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	GroupName    *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	NewComments  *string `json:"NewComments,omitempty" xml:"NewComments,omitempty"`
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
	Group     *UpdateGroupResponseBodyGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Comments   *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName  *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	MFABindRequired       *bool   `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty"`
	Password              *string `json:"Password,omitempty" xml:"Password,omitempty"`
	PasswordResetRequired *bool   `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty"`
	UserName              *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateLoginProfileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	NewDescription *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	PolicyName     *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
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
	Policy    *UpdatePolicyDescriptionResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateDate     *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DefaultVersion *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PolicyName     *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType     *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	UpdateDate     *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdatePolicyDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	NewAssumeRolePolicyDocument *string `json:"NewAssumeRolePolicyDocument,omitempty" xml:"NewAssumeRolePolicyDocument,omitempty"`
	NewDescription              *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	NewMaxSessionDuration       *int64  `json:"NewMaxSessionDuration,omitempty" xml:"NewMaxSessionDuration,omitempty"`
	RoleName                    *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
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
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Role      *UpdateRoleResponseBodyRole `json:"Role,omitempty" xml:"Role,omitempty" type:"Struct"`
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
	Arn                      *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	AssumeRolePolicyDocument *string `json:"AssumeRolePolicyDocument,omitempty" xml:"AssumeRolePolicyDocument,omitempty"`
	CreateDate               *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	MaxSessionDuration       *int64  `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty"`
	RoleId                   *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	RoleName                 *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	UpdateDate               *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	NewComments    *string `json:"NewComments,omitempty" xml:"NewComments,omitempty"`
	NewDisplayName *string `json:"NewDisplayName,omitempty" xml:"NewDisplayName,omitempty"`
	NewEmail       *string `json:"NewEmail,omitempty" xml:"NewEmail,omitempty"`
	NewMobilePhone *string `json:"NewMobilePhone,omitempty" xml:"NewMobilePhone,omitempty"`
	NewUserName    *string `json:"NewUserName,omitempty" xml:"NewUserName,omitempty"`
	UserName       *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	User      *UpdateUserResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
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
	Comments    *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	CreateDate  *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	UpdateDate  *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (client *Client) CreatePolicyWithOptions(request *CreatePolicyRequest, runtime *util.RuntimeOptions) (_result *CreatePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
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

func (client *Client) CreateRoleWithOptions(request *CreateRoleRequest, runtime *util.RuntimeOptions) (_result *CreateRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
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

func (client *Client) DeletePolicyWithOptions(request *DeletePolicyRequest, runtime *util.RuntimeOptions) (_result *DeletePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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

func (client *Client) ListPoliciesWithOptions(request *ListPoliciesRequest, runtime *util.RuntimeOptions) (_result *ListPoliciesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.PolicyType)) {
		query["PolicyType"] = request.PolicyType
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

func (client *Client) ListRolesWithOptions(request *ListRolesRequest, runtime *util.RuntimeOptions) (_result *ListRolesResponse, _err error) {
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
