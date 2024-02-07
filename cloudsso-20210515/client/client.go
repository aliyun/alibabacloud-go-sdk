// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddExternalSAMLIdPCertificateRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The X.509 certificate in the PEM format.
	//
	// The certificate is provided by the SAML IdP.
	X509Certificate *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
}

func (s AddExternalSAMLIdPCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s AddExternalSAMLIdPCertificateRequest) GoString() string {
	return s.String()
}

func (s *AddExternalSAMLIdPCertificateRequest) SetDirectoryId(v string) *AddExternalSAMLIdPCertificateRequest {
	s.DirectoryId = &v
	return s
}

func (s *AddExternalSAMLIdPCertificateRequest) SetX509Certificate(v string) *AddExternalSAMLIdPCertificateRequest {
	s.X509Certificate = &v
	return s
}

type AddExternalSAMLIdPCertificateResponseBody struct {
	// The ID of the SAML signing certificate.
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddExternalSAMLIdPCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddExternalSAMLIdPCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *AddExternalSAMLIdPCertificateResponseBody) SetCertificateId(v string) *AddExternalSAMLIdPCertificateResponseBody {
	s.CertificateId = &v
	return s
}

func (s *AddExternalSAMLIdPCertificateResponseBody) SetRequestId(v string) *AddExternalSAMLIdPCertificateResponseBody {
	s.RequestId = &v
	return s
}

type AddExternalSAMLIdPCertificateResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddExternalSAMLIdPCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddExternalSAMLIdPCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s AddExternalSAMLIdPCertificateResponse) GoString() string {
	return s.String()
}

func (s *AddExternalSAMLIdPCertificateResponse) SetHeaders(v map[string]*string) *AddExternalSAMLIdPCertificateResponse {
	s.Headers = v
	return s
}

func (s *AddExternalSAMLIdPCertificateResponse) SetStatusCode(v int32) *AddExternalSAMLIdPCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *AddExternalSAMLIdPCertificateResponse) SetBody(v *AddExternalSAMLIdPCertificateResponseBody) *AddExternalSAMLIdPCertificateResponse {
	s.Body = v
	return s
}

type AddPermissionPolicyToAccessConfigurationRequest struct {
	// The ID of the access configuration.
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The configurations of the inline policy.
	//
	// The value can be up to 4,096 characters in length.
	//
	// If you set `PermissionPolicyType` to `Inline`, you must specify this parameter. For more information about the syntax and structure of RAM policies, see [Policy syntax and structure](~~93739~~).
	InlinePolicyDocument *string `json:"InlinePolicyDocument,omitempty" xml:"InlinePolicyDocument,omitempty"`
	// The name of the policy.
	//
	// *   If you set `PermissionPolicyType` to `System`, you must set this parameter to the name of the system policy. You can obtain the name of the system policy from RAM.
	// *   If you set `PermissionPolicyType` to `Inline`, you must set this parameter to the name of the inline policy. A custom value is supported.
	PermissionPolicyName *string `json:"PermissionPolicyName,omitempty" xml:"PermissionPolicyName,omitempty"`
	// The type of the policy. Valid values:
	//
	// *   System: system policy. Resource Access Management (RAM) system policies are reused.
	// *   Inline: inline policy. Inline policies are created based on the RAM policy syntax and structure.
	PermissionPolicyType *string `json:"PermissionPolicyType,omitempty" xml:"PermissionPolicyType,omitempty"`
}

func (s AddPermissionPolicyToAccessConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s AddPermissionPolicyToAccessConfigurationRequest) GoString() string {
	return s.String()
}

func (s *AddPermissionPolicyToAccessConfigurationRequest) SetAccessConfigurationId(v string) *AddPermissionPolicyToAccessConfigurationRequest {
	s.AccessConfigurationId = &v
	return s
}

func (s *AddPermissionPolicyToAccessConfigurationRequest) SetDirectoryId(v string) *AddPermissionPolicyToAccessConfigurationRequest {
	s.DirectoryId = &v
	return s
}

func (s *AddPermissionPolicyToAccessConfigurationRequest) SetInlinePolicyDocument(v string) *AddPermissionPolicyToAccessConfigurationRequest {
	s.InlinePolicyDocument = &v
	return s
}

func (s *AddPermissionPolicyToAccessConfigurationRequest) SetPermissionPolicyName(v string) *AddPermissionPolicyToAccessConfigurationRequest {
	s.PermissionPolicyName = &v
	return s
}

func (s *AddPermissionPolicyToAccessConfigurationRequest) SetPermissionPolicyType(v string) *AddPermissionPolicyToAccessConfigurationRequest {
	s.PermissionPolicyType = &v
	return s
}

type AddPermissionPolicyToAccessConfigurationResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddPermissionPolicyToAccessConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddPermissionPolicyToAccessConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *AddPermissionPolicyToAccessConfigurationResponseBody) SetRequestId(v string) *AddPermissionPolicyToAccessConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type AddPermissionPolicyToAccessConfigurationResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddPermissionPolicyToAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddPermissionPolicyToAccessConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s AddPermissionPolicyToAccessConfigurationResponse) GoString() string {
	return s.String()
}

func (s *AddPermissionPolicyToAccessConfigurationResponse) SetHeaders(v map[string]*string) *AddPermissionPolicyToAccessConfigurationResponse {
	s.Headers = v
	return s
}

func (s *AddPermissionPolicyToAccessConfigurationResponse) SetStatusCode(v int32) *AddPermissionPolicyToAccessConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *AddPermissionPolicyToAccessConfigurationResponse) SetBody(v *AddPermissionPolicyToAccessConfigurationResponseBody) *AddPermissionPolicyToAccessConfigurationResponse {
	s.Body = v
	return s
}

type AddUserToGroupRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the group.
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddUserToGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUserToGroupRequest) GoString() string {
	return s.String()
}

func (s *AddUserToGroupRequest) SetDirectoryId(v string) *AddUserToGroupRequest {
	s.DirectoryId = &v
	return s
}

func (s *AddUserToGroupRequest) SetGroupId(v string) *AddUserToGroupRequest {
	s.GroupId = &v
	return s
}

func (s *AddUserToGroupRequest) SetUserId(v string) *AddUserToGroupRequest {
	s.UserId = &v
	return s
}

type AddUserToGroupResponseBody struct {
	// The ID of the request.
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

type ClearExternalSAMLIdentityProviderRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s ClearExternalSAMLIdentityProviderRequest) String() string {
	return tea.Prettify(s)
}

func (s ClearExternalSAMLIdentityProviderRequest) GoString() string {
	return s.String()
}

func (s *ClearExternalSAMLIdentityProviderRequest) SetDirectoryId(v string) *ClearExternalSAMLIdentityProviderRequest {
	s.DirectoryId = &v
	return s
}

type ClearExternalSAMLIdentityProviderResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClearExternalSAMLIdentityProviderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClearExternalSAMLIdentityProviderResponseBody) GoString() string {
	return s.String()
}

func (s *ClearExternalSAMLIdentityProviderResponseBody) SetRequestId(v string) *ClearExternalSAMLIdentityProviderResponseBody {
	s.RequestId = &v
	return s
}

type ClearExternalSAMLIdentityProviderResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearExternalSAMLIdentityProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClearExternalSAMLIdentityProviderResponse) String() string {
	return tea.Prettify(s)
}

func (s ClearExternalSAMLIdentityProviderResponse) GoString() string {
	return s.String()
}

func (s *ClearExternalSAMLIdentityProviderResponse) SetHeaders(v map[string]*string) *ClearExternalSAMLIdentityProviderResponse {
	s.Headers = v
	return s
}

func (s *ClearExternalSAMLIdentityProviderResponse) SetStatusCode(v int32) *ClearExternalSAMLIdentityProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearExternalSAMLIdentityProviderResponse) SetBody(v *ClearExternalSAMLIdentityProviderResponseBody) *ClearExternalSAMLIdentityProviderResponse {
	s.Body = v
	return s
}

type CreateAccessAssignmentRequest struct {
	// The ID of the access configuration.
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the CloudSSO identity.
	//
	// *   If you set `PrincipalType` to `User`, set `PrincipalId` to the ID of the CloudSSO user.
	// *   If you set `PrincipalType` to `Group`, set `PrincipalId` to the ID of the CloudSSO group.
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The type of the CloudSSO identity. Valid values:
	//
	// *   User
	// *   Group
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The ID of the task object.
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the task object. Set the value to RD-Account, which specifies the accounts in the resource directory.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s CreateAccessAssignmentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessAssignmentRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessAssignmentRequest) SetAccessConfigurationId(v string) *CreateAccessAssignmentRequest {
	s.AccessConfigurationId = &v
	return s
}

func (s *CreateAccessAssignmentRequest) SetDirectoryId(v string) *CreateAccessAssignmentRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateAccessAssignmentRequest) SetPrincipalId(v string) *CreateAccessAssignmentRequest {
	s.PrincipalId = &v
	return s
}

func (s *CreateAccessAssignmentRequest) SetPrincipalType(v string) *CreateAccessAssignmentRequest {
	s.PrincipalType = &v
	return s
}

func (s *CreateAccessAssignmentRequest) SetTargetId(v string) *CreateAccessAssignmentRequest {
	s.TargetId = &v
	return s
}

func (s *CreateAccessAssignmentRequest) SetTargetType(v string) *CreateAccessAssignmentRequest {
	s.TargetType = &v
	return s
}

type CreateAccessAssignmentResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried task.
	Task *CreateAccessAssignmentResponseBodyTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
}

func (s CreateAccessAssignmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessAssignmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccessAssignmentResponseBody) SetRequestId(v string) *CreateAccessAssignmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccessAssignmentResponseBody) SetTask(v *CreateAccessAssignmentResponseBodyTask) *CreateAccessAssignmentResponseBody {
	s.Task = v
	return s
}

type CreateAccessAssignmentResponseBodyTask struct {
	// The ID of the access configuration.
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The name of the access configuration.
	AccessConfigurationName *string `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	// The ID of the CloudSSO identity.
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The name of the CloudSSO identity.
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The type of the CloudSSO identity. Valid values:
	//
	// *   User
	// *   Group
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The task status. Valid values:
	//
	// *   InProgress: The task is running.
	// *   Success: The task is successful.
	// *   Failed: The task failed.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the task object.
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the task object.
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The path ID of the task object in the resource directory.
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// The path name of the task object in the resource directory.
	TargetPathName *string `json:"TargetPathName,omitempty" xml:"TargetPathName,omitempty"`
	// The type of the task object. The value is fixed as RD-Account, which indicates the accounts in the resource directory.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The ID of the job.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type. The value is fixed as CreateAccessAssignment, which indicates that access permissions on an account in your resource directory are assigned.
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s CreateAccessAssignmentResponseBodyTask) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessAssignmentResponseBodyTask) GoString() string {
	return s.String()
}

func (s *CreateAccessAssignmentResponseBodyTask) SetAccessConfigurationId(v string) *CreateAccessAssignmentResponseBodyTask {
	s.AccessConfigurationId = &v
	return s
}

func (s *CreateAccessAssignmentResponseBodyTask) SetAccessConfigurationName(v string) *CreateAccessAssignmentResponseBodyTask {
	s.AccessConfigurationName = &v
	return s
}

func (s *CreateAccessAssignmentResponseBodyTask) SetPrincipalId(v string) *CreateAccessAssignmentResponseBodyTask {
	s.PrincipalId = &v
	return s
}

func (s *CreateAccessAssignmentResponseBodyTask) SetPrincipalName(v string) *CreateAccessAssignmentResponseBodyTask {
	s.PrincipalName = &v
	return s
}

func (s *CreateAccessAssignmentResponseBodyTask) SetPrincipalType(v string) *CreateAccessAssignmentResponseBodyTask {
	s.PrincipalType = &v
	return s
}

func (s *CreateAccessAssignmentResponseBodyTask) SetStatus(v string) *CreateAccessAssignmentResponseBodyTask {
	s.Status = &v
	return s
}

func (s *CreateAccessAssignmentResponseBodyTask) SetTargetId(v string) *CreateAccessAssignmentResponseBodyTask {
	s.TargetId = &v
	return s
}

func (s *CreateAccessAssignmentResponseBodyTask) SetTargetName(v string) *CreateAccessAssignmentResponseBodyTask {
	s.TargetName = &v
	return s
}

func (s *CreateAccessAssignmentResponseBodyTask) SetTargetPath(v string) *CreateAccessAssignmentResponseBodyTask {
	s.TargetPath = &v
	return s
}

func (s *CreateAccessAssignmentResponseBodyTask) SetTargetPathName(v string) *CreateAccessAssignmentResponseBodyTask {
	s.TargetPathName = &v
	return s
}

func (s *CreateAccessAssignmentResponseBodyTask) SetTargetType(v string) *CreateAccessAssignmentResponseBodyTask {
	s.TargetType = &v
	return s
}

func (s *CreateAccessAssignmentResponseBodyTask) SetTaskId(v string) *CreateAccessAssignmentResponseBodyTask {
	s.TaskId = &v
	return s
}

func (s *CreateAccessAssignmentResponseBodyTask) SetTaskType(v string) *CreateAccessAssignmentResponseBodyTask {
	s.TaskType = &v
	return s
}

type CreateAccessAssignmentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAccessAssignmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAccessAssignmentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessAssignmentResponse) GoString() string {
	return s.String()
}

func (s *CreateAccessAssignmentResponse) SetHeaders(v map[string]*string) *CreateAccessAssignmentResponse {
	s.Headers = v
	return s
}

func (s *CreateAccessAssignmentResponse) SetStatusCode(v int32) *CreateAccessAssignmentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccessAssignmentResponse) SetBody(v *CreateAccessAssignmentResponseBody) *CreateAccessAssignmentResponse {
	s.Body = v
	return s
}

type CreateAccessConfigurationRequest struct {
	// The name of the access configuration.
	//
	// The name can contain letters, digits, and hyphens (-).
	//
	// The name can be up to 32 characters in length.
	AccessConfigurationName *string `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	// The description of the access configuration.
	//
	// The description can be up to 1,024 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The initial web page that is displayed after a CloudSSO user accesses an account in your resource directory by using the access configuration.
	//
	// The web page must be a page of the Alibaba Cloud Management Console. By default, this parameter is empty, which indicates that the initial web page is the homepage of the Alibaba Cloud Management Console.
	RelayState *string `json:"RelayState,omitempty" xml:"RelayState,omitempty"`
	// The duration of a session in which a CloudSSO user accesses an account in your resource directory by using the access configuration.
	//
	// Unit: seconds.
	//
	// Valid values: 900 to 43200. The value 900 indicates 15 minutes. The value 43200 indicates 12 hours.
	//
	// Default value: 3600. The value indicates 1 hour.
	SessionDuration *int32 `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
}

func (s CreateAccessConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessConfigurationRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessConfigurationRequest) SetAccessConfigurationName(v string) *CreateAccessConfigurationRequest {
	s.AccessConfigurationName = &v
	return s
}

func (s *CreateAccessConfigurationRequest) SetDescription(v string) *CreateAccessConfigurationRequest {
	s.Description = &v
	return s
}

func (s *CreateAccessConfigurationRequest) SetDirectoryId(v string) *CreateAccessConfigurationRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateAccessConfigurationRequest) SetRelayState(v string) *CreateAccessConfigurationRequest {
	s.RelayState = &v
	return s
}

func (s *CreateAccessConfigurationRequest) SetSessionDuration(v int32) *CreateAccessConfigurationRequest {
	s.SessionDuration = &v
	return s
}

type CreateAccessConfigurationResponseBody struct {
	// The information about the access configuration.
	AccessConfiguration *CreateAccessConfigurationResponseBodyAccessConfiguration `json:"AccessConfiguration,omitempty" xml:"AccessConfiguration,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccessConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccessConfigurationResponseBody) SetAccessConfiguration(v *CreateAccessConfigurationResponseBodyAccessConfiguration) *CreateAccessConfigurationResponseBody {
	s.AccessConfiguration = v
	return s
}

func (s *CreateAccessConfigurationResponseBody) SetRequestId(v string) *CreateAccessConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type CreateAccessConfigurationResponseBodyAccessConfiguration struct {
	// The ID of the access configuration.
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The name of the access configuration.
	AccessConfigurationName *string `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	// The time when the access configuration was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the access configuration.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The initial web page that is displayed after a CloudSSO user accesses an account in your resource directory by using the access configuration.
	RelayState *string `json:"RelayState,omitempty" xml:"RelayState,omitempty"`
	// The duration of a session in which a CloudSSO user accesses an account in your resource directory by using the access configuration.
	//
	// Unit: seconds.
	SessionDuration *int32 `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	// The status notification.
	StatusNotifications []*string `json:"StatusNotifications,omitempty" xml:"StatusNotifications,omitempty" type:"Repeated"`
	// The time when the information about the access configuration was modified.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s CreateAccessConfigurationResponseBodyAccessConfiguration) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessConfigurationResponseBodyAccessConfiguration) GoString() string {
	return s.String()
}

func (s *CreateAccessConfigurationResponseBodyAccessConfiguration) SetAccessConfigurationId(v string) *CreateAccessConfigurationResponseBodyAccessConfiguration {
	s.AccessConfigurationId = &v
	return s
}

func (s *CreateAccessConfigurationResponseBodyAccessConfiguration) SetAccessConfigurationName(v string) *CreateAccessConfigurationResponseBodyAccessConfiguration {
	s.AccessConfigurationName = &v
	return s
}

func (s *CreateAccessConfigurationResponseBodyAccessConfiguration) SetCreateTime(v string) *CreateAccessConfigurationResponseBodyAccessConfiguration {
	s.CreateTime = &v
	return s
}

func (s *CreateAccessConfigurationResponseBodyAccessConfiguration) SetDescription(v string) *CreateAccessConfigurationResponseBodyAccessConfiguration {
	s.Description = &v
	return s
}

func (s *CreateAccessConfigurationResponseBodyAccessConfiguration) SetRelayState(v string) *CreateAccessConfigurationResponseBodyAccessConfiguration {
	s.RelayState = &v
	return s
}

func (s *CreateAccessConfigurationResponseBodyAccessConfiguration) SetSessionDuration(v int32) *CreateAccessConfigurationResponseBodyAccessConfiguration {
	s.SessionDuration = &v
	return s
}

func (s *CreateAccessConfigurationResponseBodyAccessConfiguration) SetStatusNotifications(v []*string) *CreateAccessConfigurationResponseBodyAccessConfiguration {
	s.StatusNotifications = v
	return s
}

func (s *CreateAccessConfigurationResponseBodyAccessConfiguration) SetUpdateTime(v string) *CreateAccessConfigurationResponseBodyAccessConfiguration {
	s.UpdateTime = &v
	return s
}

type CreateAccessConfigurationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAccessConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessConfigurationResponse) GoString() string {
	return s.String()
}

func (s *CreateAccessConfigurationResponse) SetHeaders(v map[string]*string) *CreateAccessConfigurationResponse {
	s.Headers = v
	return s
}

func (s *CreateAccessConfigurationResponse) SetStatusCode(v int32) *CreateAccessConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccessConfigurationResponse) SetBody(v *CreateAccessConfigurationResponseBody) *CreateAccessConfigurationResponse {
	s.Body = v
	return s
}

type CreateDirectoryRequest struct {
	// The name of the directory. The name must be globally unique.
	//
	// The name can contain lowercase letters, digits, or hyphens (-). The name cannot start or end with a hyphen (-) and cannot contain two consecutive hyphens (-). The name cannot start with d-.
	//
	// The name must be 2 to 64 characters in length.
	//
	// >  If you do not specify this parameter, the value of this parameter is automatically generated by the system.
	DirectoryName *string `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
}

func (s CreateDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDirectoryRequest) GoString() string {
	return s.String()
}

func (s *CreateDirectoryRequest) SetDirectoryName(v string) *CreateDirectoryRequest {
	s.DirectoryName = &v
	return s
}

type CreateDirectoryResponseBody struct {
	// The information about the directory.
	Directory *CreateDirectoryResponseBodyDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDirectoryResponseBody) SetDirectory(v *CreateDirectoryResponseBodyDirectory) *CreateDirectoryResponseBody {
	s.Directory = v
	return s
}

func (s *CreateDirectoryResponseBody) SetRequestId(v string) *CreateDirectoryResponseBody {
	s.RequestId = &v
	return s
}

type CreateDirectoryResponseBodyDirectory struct {
	// The time when the directory was created. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The name of the directory.
	DirectoryName *string `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	// The region ID of the directory.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The time when the directory was modified. The time is displayed in UTC.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s CreateDirectoryResponseBodyDirectory) String() string {
	return tea.Prettify(s)
}

func (s CreateDirectoryResponseBodyDirectory) GoString() string {
	return s.String()
}

func (s *CreateDirectoryResponseBodyDirectory) SetCreateTime(v string) *CreateDirectoryResponseBodyDirectory {
	s.CreateTime = &v
	return s
}

func (s *CreateDirectoryResponseBodyDirectory) SetDirectoryId(v string) *CreateDirectoryResponseBodyDirectory {
	s.DirectoryId = &v
	return s
}

func (s *CreateDirectoryResponseBodyDirectory) SetDirectoryName(v string) *CreateDirectoryResponseBodyDirectory {
	s.DirectoryName = &v
	return s
}

func (s *CreateDirectoryResponseBodyDirectory) SetRegion(v string) *CreateDirectoryResponseBodyDirectory {
	s.Region = &v
	return s
}

func (s *CreateDirectoryResponseBodyDirectory) SetUpdateTime(v string) *CreateDirectoryResponseBodyDirectory {
	s.UpdateTime = &v
	return s
}

type CreateDirectoryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDirectoryResponse) GoString() string {
	return s.String()
}

func (s *CreateDirectoryResponse) SetHeaders(v map[string]*string) *CreateDirectoryResponse {
	s.Headers = v
	return s
}

func (s *CreateDirectoryResponse) SetStatusCode(v int32) *CreateDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDirectoryResponse) SetBody(v *CreateDirectoryResponseBody) *CreateDirectoryResponse {
	s.Body = v
	return s
}

type CreateGroupRequest struct {
	// The description of the group.
	//
	// The description can be up to 1,024 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The name of the group.
	//
	// The name can contain letters, digits, underscores (\_), hyphens (-), and periods (.).
	//
	// The name can be up to 128 characters in length.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s CreateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupRequest) SetDescription(v string) *CreateGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupRequest) SetDirectoryId(v string) *CreateGroupRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupRequest) SetGroupName(v string) *CreateGroupRequest {
	s.GroupName = &v
	return s
}

type CreateGroupResponseBody struct {
	// The information about the group.
	Group *CreateGroupResponseBodyGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	// The ID of the request.
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
	// The time when the group was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the group.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the group.
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the group. The value is fixed as Manual, which indicates that the group is manually created.
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	// The time when the information about the group was modified.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s CreateGroupResponseBodyGroup) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponseBodyGroup) GoString() string {
	return s.String()
}

func (s *CreateGroupResponseBodyGroup) SetCreateTime(v string) *CreateGroupResponseBodyGroup {
	s.CreateTime = &v
	return s
}

func (s *CreateGroupResponseBodyGroup) SetDescription(v string) *CreateGroupResponseBodyGroup {
	s.Description = &v
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

func (s *CreateGroupResponseBodyGroup) SetProvisionType(v string) *CreateGroupResponseBodyGroup {
	s.ProvisionType = &v
	return s
}

func (s *CreateGroupResponseBodyGroup) SetUpdateTime(v string) *CreateGroupResponseBodyGroup {
	s.UpdateTime = &v
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

type CreateSCIMServerCredentialRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s CreateSCIMServerCredentialRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSCIMServerCredentialRequest) GoString() string {
	return s.String()
}

func (s *CreateSCIMServerCredentialRequest) SetDirectoryId(v string) *CreateSCIMServerCredentialRequest {
	s.DirectoryId = &v
	return s
}

type CreateSCIMServerCredentialResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the SCIM credential.
	SCIMServerCredential *CreateSCIMServerCredentialResponseBodySCIMServerCredential `json:"SCIMServerCredential,omitempty" xml:"SCIMServerCredential,omitempty" type:"Struct"`
}

func (s CreateSCIMServerCredentialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSCIMServerCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSCIMServerCredentialResponseBody) SetRequestId(v string) *CreateSCIMServerCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSCIMServerCredentialResponseBody) SetSCIMServerCredential(v *CreateSCIMServerCredentialResponseBodySCIMServerCredential) *CreateSCIMServerCredentialResponseBody {
	s.SCIMServerCredential = v
	return s
}

type CreateSCIMServerCredentialResponseBodySCIMServerCredential struct {
	// The time when the SCIM credential was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the SCIM credential.
	CredentialId *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
	// The SCIM credential.
	//
	// >  The SCIM credential is returned only when it is created. After the SCIM credential is created, you cannot query it. Keep the SCIM 
	CredentialSecret *string `json:"CredentialSecret,omitempty" xml:"CredentialSecret,omitempty"`
	// The type of the SCIM credential.
	CredentialType *string `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The time when the SCIM credential expires.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The status of the SCIM credential. The value is fixed as Enabled, which indicates that the SCIM credential is enabled.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateSCIMServerCredentialResponseBodySCIMServerCredential) String() string {
	return tea.Prettify(s)
}

func (s CreateSCIMServerCredentialResponseBodySCIMServerCredential) GoString() string {
	return s.String()
}

func (s *CreateSCIMServerCredentialResponseBodySCIMServerCredential) SetCreateTime(v string) *CreateSCIMServerCredentialResponseBodySCIMServerCredential {
	s.CreateTime = &v
	return s
}

func (s *CreateSCIMServerCredentialResponseBodySCIMServerCredential) SetCredentialId(v string) *CreateSCIMServerCredentialResponseBodySCIMServerCredential {
	s.CredentialId = &v
	return s
}

func (s *CreateSCIMServerCredentialResponseBodySCIMServerCredential) SetCredentialSecret(v string) *CreateSCIMServerCredentialResponseBodySCIMServerCredential {
	s.CredentialSecret = &v
	return s
}

func (s *CreateSCIMServerCredentialResponseBodySCIMServerCredential) SetCredentialType(v string) *CreateSCIMServerCredentialResponseBodySCIMServerCredential {
	s.CredentialType = &v
	return s
}

func (s *CreateSCIMServerCredentialResponseBodySCIMServerCredential) SetDirectoryId(v string) *CreateSCIMServerCredentialResponseBodySCIMServerCredential {
	s.DirectoryId = &v
	return s
}

func (s *CreateSCIMServerCredentialResponseBodySCIMServerCredential) SetExpireTime(v string) *CreateSCIMServerCredentialResponseBodySCIMServerCredential {
	s.ExpireTime = &v
	return s
}

func (s *CreateSCIMServerCredentialResponseBodySCIMServerCredential) SetStatus(v string) *CreateSCIMServerCredentialResponseBodySCIMServerCredential {
	s.Status = &v
	return s
}

type CreateSCIMServerCredentialResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSCIMServerCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSCIMServerCredentialResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSCIMServerCredentialResponse) GoString() string {
	return s.String()
}

func (s *CreateSCIMServerCredentialResponse) SetHeaders(v map[string]*string) *CreateSCIMServerCredentialResponse {
	s.Headers = v
	return s
}

func (s *CreateSCIMServerCredentialResponse) SetStatusCode(v int32) *CreateSCIMServerCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSCIMServerCredentialResponse) SetBody(v *CreateSCIMServerCredentialResponseBody) *CreateSCIMServerCredentialResponse {
	s.Body = v
	return s
}

type CreateUserRequest struct {
	// The description of the user.
	//
	// The description can be up to 1,024 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The display name of the user.
	//
	// The name can be up to 256 characters in length.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the user. The email address must be unique within the directory.
	//
	// The email address can be up to 128 characters in length.
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The first name of the user.
	//
	// The name can be up to 64 characters in length.
	FirstName *string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	// The last name of the user.
	//
	// The name can be up to 64 characters in length.
	LastName *string `json:"LastName,omitempty" xml:"LastName,omitempty"`
	// The status of the user. Valid values:
	//
	// *   Enabled: The logon of the user is enabled. This is the default value.
	// *   Disabled: The logon of the user is disabled.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the user. The name must be unique within the directory. The name cannot be changed.
	//
	// The name can contain numbers, letters, and the following special characters: `@_-.`
	//
	// The name can be up to 64 characters in length.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) SetDescription(v string) *CreateUserRequest {
	s.Description = &v
	return s
}

func (s *CreateUserRequest) SetDirectoryId(v string) *CreateUserRequest {
	s.DirectoryId = &v
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

func (s *CreateUserRequest) SetFirstName(v string) *CreateUserRequest {
	s.FirstName = &v
	return s
}

func (s *CreateUserRequest) SetLastName(v string) *CreateUserRequest {
	s.LastName = &v
	return s
}

func (s *CreateUserRequest) SetStatus(v string) *CreateUserRequest {
	s.Status = &v
	return s
}

func (s *CreateUserRequest) SetUserName(v string) *CreateUserRequest {
	s.UserName = &v
	return s
}

type CreateUserResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the user.
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
	// The time when the user was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the user.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the user.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the user.
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The first name of the user.
	FirstName *string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	// The last name of the user.
	LastName *string `json:"LastName,omitempty" xml:"LastName,omitempty"`
	// The type of the user. Valid values:
	//
	// *   Manual: The user is manually created.
	// *   Synchronized: The user is synchronized from an external identity provider (IdP).
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	// The status of the user. Valid values:
	//
	// *   Enabled: The logon of the user is enabled.
	// *   Disabled: The logon of the user is disabled.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the user was modified.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The name of the user.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateUserResponseBodyUser) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponseBodyUser) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBodyUser) SetCreateTime(v string) *CreateUserResponseBodyUser {
	s.CreateTime = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetDescription(v string) *CreateUserResponseBodyUser {
	s.Description = &v
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

func (s *CreateUserResponseBodyUser) SetFirstName(v string) *CreateUserResponseBodyUser {
	s.FirstName = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetLastName(v string) *CreateUserResponseBodyUser {
	s.LastName = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetProvisionType(v string) *CreateUserResponseBodyUser {
	s.ProvisionType = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetStatus(v string) *CreateUserResponseBodyUser {
	s.Status = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetUpdateTime(v string) *CreateUserResponseBodyUser {
	s.UpdateTime = &v
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

type CreateUserProvisioningRequest struct {
	DeletionStrategy    *string `json:"DeletionStrategy,omitempty" xml:"DeletionStrategy,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DirectoryId         *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	DuplicationStrategy *string `json:"DuplicationStrategy,omitempty" xml:"DuplicationStrategy,omitempty"`
	PrincipalId         *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	PrincipalType       *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	TargetId            *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetType          *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s CreateUserProvisioningRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserProvisioningRequest) GoString() string {
	return s.String()
}

func (s *CreateUserProvisioningRequest) SetDeletionStrategy(v string) *CreateUserProvisioningRequest {
	s.DeletionStrategy = &v
	return s
}

func (s *CreateUserProvisioningRequest) SetDescription(v string) *CreateUserProvisioningRequest {
	s.Description = &v
	return s
}

func (s *CreateUserProvisioningRequest) SetDirectoryId(v string) *CreateUserProvisioningRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateUserProvisioningRequest) SetDuplicationStrategy(v string) *CreateUserProvisioningRequest {
	s.DuplicationStrategy = &v
	return s
}

func (s *CreateUserProvisioningRequest) SetPrincipalId(v string) *CreateUserProvisioningRequest {
	s.PrincipalId = &v
	return s
}

func (s *CreateUserProvisioningRequest) SetPrincipalType(v string) *CreateUserProvisioningRequest {
	s.PrincipalType = &v
	return s
}

func (s *CreateUserProvisioningRequest) SetTargetId(v string) *CreateUserProvisioningRequest {
	s.TargetId = &v
	return s
}

func (s *CreateUserProvisioningRequest) SetTargetType(v string) *CreateUserProvisioningRequest {
	s.TargetType = &v
	return s
}

type CreateUserProvisioningResponseBody struct {
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserProvisioning *CreateUserProvisioningResponseBodyUserProvisioning `json:"UserProvisioning,omitempty" xml:"UserProvisioning,omitempty" type:"Struct"`
}

func (s CreateUserProvisioningResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserProvisioningResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserProvisioningResponseBody) SetRequestId(v string) *CreateUserProvisioningResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserProvisioningResponseBody) SetUserProvisioning(v *CreateUserProvisioningResponseBodyUserProvisioning) *CreateUserProvisioningResponseBody {
	s.UserProvisioning = v
	return s
}

type CreateUserProvisioningResponseBodyUserProvisioning struct {
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeletionStrategy    *string `json:"DeletionStrategy,omitempty" xml:"DeletionStrategy,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DirectoryId         *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	DuplicationStrategy *string `json:"DuplicationStrategy,omitempty" xml:"DuplicationStrategy,omitempty"`
	OwnerPk             *string `json:"OwnerPk,omitempty" xml:"OwnerPk,omitempty"`
	PrincipalId         *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	PrincipalName       *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	PrincipalType       *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetId            *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetName          *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	TargetPath          *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	TargetType          *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	UpdateTime          *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserProvisioningId  *string `json:"UserProvisioningId,omitempty" xml:"UserProvisioningId,omitempty"`
}

func (s CreateUserProvisioningResponseBodyUserProvisioning) String() string {
	return tea.Prettify(s)
}

func (s CreateUserProvisioningResponseBodyUserProvisioning) GoString() string {
	return s.String()
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetCreateTime(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.CreateTime = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetDeletionStrategy(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.DeletionStrategy = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetDescription(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.Description = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetDirectoryId(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.DirectoryId = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetDuplicationStrategy(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.DuplicationStrategy = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetOwnerPk(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.OwnerPk = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetPrincipalId(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.PrincipalId = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetPrincipalName(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.PrincipalName = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetPrincipalType(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.PrincipalType = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetStatus(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.Status = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetTargetId(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.TargetId = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetTargetName(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.TargetName = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetTargetPath(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.TargetPath = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetTargetType(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.TargetType = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetUpdateTime(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.UpdateTime = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetUserProvisioningId(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.UserProvisioningId = &v
	return s
}

type CreateUserProvisioningResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserProvisioningResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserProvisioningResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserProvisioningResponse) GoString() string {
	return s.String()
}

func (s *CreateUserProvisioningResponse) SetHeaders(v map[string]*string) *CreateUserProvisioningResponse {
	s.Headers = v
	return s
}

func (s *CreateUserProvisioningResponse) SetStatusCode(v int32) *CreateUserProvisioningResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserProvisioningResponse) SetBody(v *CreateUserProvisioningResponseBody) *CreateUserProvisioningResponse {
	s.Body = v
	return s
}

type DeleteAccessAssignmentRequest struct {
	// The ID of the access configuration.
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// Specifies whether to de-provision the access configuration when you remove the access permissions from the CloudSSO identity. The access configuration is used to assign the access permissions, and the identity is the only one that uses the access configuration and is associated with the account. Valid values:
	//
	// *   DeprovisionForLastAccessAssignmentOnAccount: de-provisions the access configuration.
	// *   None: does not de-provision the access configuration. This is the default value.
	DeprovisionStrategy *string `json:"DeprovisionStrategy,omitempty" xml:"DeprovisionStrategy,omitempty"`
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the CloudSSO identity.
	//
	// *   If you set `PrincipalType` to `User`, set `PrincipalId` to the ID of the CloudSSO user.
	// *   If you set `PrincipalType` to `Group`, set `PrincipalId` to the ID of the CloudSSO group.
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The type of the CloudSSO identity. Valid values:
	//
	// *   User
	// *   Group
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The ID of the task object.
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the task object. Set the value to RD-Account, which specifies the accounts in the resource directory.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s DeleteAccessAssignmentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessAssignmentRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessAssignmentRequest) SetAccessConfigurationId(v string) *DeleteAccessAssignmentRequest {
	s.AccessConfigurationId = &v
	return s
}

func (s *DeleteAccessAssignmentRequest) SetDeprovisionStrategy(v string) *DeleteAccessAssignmentRequest {
	s.DeprovisionStrategy = &v
	return s
}

func (s *DeleteAccessAssignmentRequest) SetDirectoryId(v string) *DeleteAccessAssignmentRequest {
	s.DirectoryId = &v
	return s
}

func (s *DeleteAccessAssignmentRequest) SetPrincipalId(v string) *DeleteAccessAssignmentRequest {
	s.PrincipalId = &v
	return s
}

func (s *DeleteAccessAssignmentRequest) SetPrincipalType(v string) *DeleteAccessAssignmentRequest {
	s.PrincipalType = &v
	return s
}

func (s *DeleteAccessAssignmentRequest) SetTargetId(v string) *DeleteAccessAssignmentRequest {
	s.TargetId = &v
	return s
}

func (s *DeleteAccessAssignmentRequest) SetTargetType(v string) *DeleteAccessAssignmentRequest {
	s.TargetType = &v
	return s
}

type DeleteAccessAssignmentResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task information.
	Task *DeleteAccessAssignmentResponseBodyTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
}

func (s DeleteAccessAssignmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessAssignmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccessAssignmentResponseBody) SetRequestId(v string) *DeleteAccessAssignmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAccessAssignmentResponseBody) SetTask(v *DeleteAccessAssignmentResponseBodyTask) *DeleteAccessAssignmentResponseBody {
	s.Task = v
	return s
}

type DeleteAccessAssignmentResponseBodyTask struct {
	// The ID of the access configuration.
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The name of the access configuration.
	AccessConfigurationName *string `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	// The ID of the CloudSSO identity.
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The name of the CloudSSO identity.
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The type of the CloudSSO identity. Valid values:
	//
	// *   User
	// *   Group
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The task status. Valid values:
	//
	// *   InProgress: The task is running.
	// *   Success: The task is successful.
	// *   Failed: The task failed.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the task object.
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the task object.
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The path ID of the task object in the resource directory.
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// The path name of the task object in the resource directory.
	TargetPathName *string `json:"TargetPathName,omitempty" xml:"TargetPathName,omitempty"`
	// The type of the task object. The value is fixed as RD-Account, which indicates the accounts in the resource directory.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The task ID.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type. The value is fixed as DeleteAccessAssignment, which indicates that access permissions on an account in your resource directory are removed.
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DeleteAccessAssignmentResponseBodyTask) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessAssignmentResponseBodyTask) GoString() string {
	return s.String()
}

func (s *DeleteAccessAssignmentResponseBodyTask) SetAccessConfigurationId(v string) *DeleteAccessAssignmentResponseBodyTask {
	s.AccessConfigurationId = &v
	return s
}

func (s *DeleteAccessAssignmentResponseBodyTask) SetAccessConfigurationName(v string) *DeleteAccessAssignmentResponseBodyTask {
	s.AccessConfigurationName = &v
	return s
}

func (s *DeleteAccessAssignmentResponseBodyTask) SetPrincipalId(v string) *DeleteAccessAssignmentResponseBodyTask {
	s.PrincipalId = &v
	return s
}

func (s *DeleteAccessAssignmentResponseBodyTask) SetPrincipalName(v string) *DeleteAccessAssignmentResponseBodyTask {
	s.PrincipalName = &v
	return s
}

func (s *DeleteAccessAssignmentResponseBodyTask) SetPrincipalType(v string) *DeleteAccessAssignmentResponseBodyTask {
	s.PrincipalType = &v
	return s
}

func (s *DeleteAccessAssignmentResponseBodyTask) SetStatus(v string) *DeleteAccessAssignmentResponseBodyTask {
	s.Status = &v
	return s
}

func (s *DeleteAccessAssignmentResponseBodyTask) SetTargetId(v string) *DeleteAccessAssignmentResponseBodyTask {
	s.TargetId = &v
	return s
}

func (s *DeleteAccessAssignmentResponseBodyTask) SetTargetName(v string) *DeleteAccessAssignmentResponseBodyTask {
	s.TargetName = &v
	return s
}

func (s *DeleteAccessAssignmentResponseBodyTask) SetTargetPath(v string) *DeleteAccessAssignmentResponseBodyTask {
	s.TargetPath = &v
	return s
}

func (s *DeleteAccessAssignmentResponseBodyTask) SetTargetPathName(v string) *DeleteAccessAssignmentResponseBodyTask {
	s.TargetPathName = &v
	return s
}

func (s *DeleteAccessAssignmentResponseBodyTask) SetTargetType(v string) *DeleteAccessAssignmentResponseBodyTask {
	s.TargetType = &v
	return s
}

func (s *DeleteAccessAssignmentResponseBodyTask) SetTaskId(v string) *DeleteAccessAssignmentResponseBodyTask {
	s.TaskId = &v
	return s
}

func (s *DeleteAccessAssignmentResponseBodyTask) SetTaskType(v string) *DeleteAccessAssignmentResponseBodyTask {
	s.TaskType = &v
	return s
}

type DeleteAccessAssignmentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAccessAssignmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAccessAssignmentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessAssignmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessAssignmentResponse) SetHeaders(v map[string]*string) *DeleteAccessAssignmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessAssignmentResponse) SetStatusCode(v int32) *DeleteAccessAssignmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessAssignmentResponse) SetBody(v *DeleteAccessAssignmentResponseBody) *DeleteAccessAssignmentResponse {
	s.Body = v
	return s
}

type DeleteAccessConfigurationRequest struct {
	// The ID of the access configuration.
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// Specifies whether to forcibly remove system policies and inline policies. Valid values:
	//
	// *   true: When you delete the access configuration, the associated system policies and inline policies are forcibly removed.
	// *   false: When you delete the access configuration, the associated system policies and inline policies are not forcibly removed. This is the default value. If these policies exist in the access configuration, the deletion fails. Before you delete the access configuration, you must remove the system policies and inline policies. For more information, see [RemovePermissionPolicyFromAccessConfiguration](~~336904~~).
	ForceRemovePermissionPolicies *bool `json:"ForceRemovePermissionPolicies,omitempty" xml:"ForceRemovePermissionPolicies,omitempty"`
}

func (s DeleteAccessConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessConfigurationRequest) SetAccessConfigurationId(v string) *DeleteAccessConfigurationRequest {
	s.AccessConfigurationId = &v
	return s
}

func (s *DeleteAccessConfigurationRequest) SetDirectoryId(v string) *DeleteAccessConfigurationRequest {
	s.DirectoryId = &v
	return s
}

func (s *DeleteAccessConfigurationRequest) SetForceRemovePermissionPolicies(v bool) *DeleteAccessConfigurationRequest {
	s.ForceRemovePermissionPolicies = &v
	return s
}

type DeleteAccessConfigurationResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccessConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccessConfigurationResponseBody) SetRequestId(v string) *DeleteAccessConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAccessConfigurationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAccessConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessConfigurationResponse) SetHeaders(v map[string]*string) *DeleteAccessConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessConfigurationResponse) SetStatusCode(v int32) *DeleteAccessConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessConfigurationResponse) SetBody(v *DeleteAccessConfigurationResponseBody) *DeleteAccessConfigurationResponse {
	s.Body = v
	return s
}

type DeleteDirectoryRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s DeleteDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirectoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteDirectoryRequest) SetDirectoryId(v string) *DeleteDirectoryRequest {
	s.DirectoryId = &v
	return s
}

type DeleteDirectoryResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDirectoryResponseBody) SetRequestId(v string) *DeleteDirectoryResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDirectoryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirectoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteDirectoryResponse) SetHeaders(v map[string]*string) *DeleteDirectoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteDirectoryResponse) SetStatusCode(v int32) *DeleteDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDirectoryResponse) SetBody(v *DeleteDirectoryResponseBody) *DeleteDirectoryResponse {
	s.Body = v
	return s
}

type DeleteGroupRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the group.
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s DeleteGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupRequest) SetDirectoryId(v string) *DeleteGroupRequest {
	s.DirectoryId = &v
	return s
}

func (s *DeleteGroupRequest) SetGroupId(v string) *DeleteGroupRequest {
	s.GroupId = &v
	return s
}

type DeleteGroupResponseBody struct {
	// The ID of the request.
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

type DeleteMFADeviceForUserRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the MFA device.
	//
	// You can call the [ListMFADevicesForUser](~~333531~~) operation to query the IDs of MFA devices.
	MFADeviceId *string `json:"MFADeviceId,omitempty" xml:"MFADeviceId,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteMFADeviceForUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMFADeviceForUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteMFADeviceForUserRequest) SetDirectoryId(v string) *DeleteMFADeviceForUserRequest {
	s.DirectoryId = &v
	return s
}

func (s *DeleteMFADeviceForUserRequest) SetMFADeviceId(v string) *DeleteMFADeviceForUserRequest {
	s.MFADeviceId = &v
	return s
}

func (s *DeleteMFADeviceForUserRequest) SetUserId(v string) *DeleteMFADeviceForUserRequest {
	s.UserId = &v
	return s
}

type DeleteMFADeviceForUserResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMFADeviceForUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMFADeviceForUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMFADeviceForUserResponseBody) SetRequestId(v string) *DeleteMFADeviceForUserResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMFADeviceForUserResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMFADeviceForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMFADeviceForUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMFADeviceForUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteMFADeviceForUserResponse) SetHeaders(v map[string]*string) *DeleteMFADeviceForUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteMFADeviceForUserResponse) SetStatusCode(v int32) *DeleteMFADeviceForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMFADeviceForUserResponse) SetBody(v *DeleteMFADeviceForUserResponseBody) *DeleteMFADeviceForUserResponse {
	s.Body = v
	return s
}

type DeleteSCIMServerCredentialRequest struct {
	// The ID of the SCIM credential.
	CredentialId *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s DeleteSCIMServerCredentialRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSCIMServerCredentialRequest) GoString() string {
	return s.String()
}

func (s *DeleteSCIMServerCredentialRequest) SetCredentialId(v string) *DeleteSCIMServerCredentialRequest {
	s.CredentialId = &v
	return s
}

func (s *DeleteSCIMServerCredentialRequest) SetDirectoryId(v string) *DeleteSCIMServerCredentialRequest {
	s.DirectoryId = &v
	return s
}

type DeleteSCIMServerCredentialResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSCIMServerCredentialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSCIMServerCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSCIMServerCredentialResponseBody) SetRequestId(v string) *DeleteSCIMServerCredentialResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSCIMServerCredentialResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSCIMServerCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSCIMServerCredentialResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSCIMServerCredentialResponse) GoString() string {
	return s.String()
}

func (s *DeleteSCIMServerCredentialResponse) SetHeaders(v map[string]*string) *DeleteSCIMServerCredentialResponse {
	s.Headers = v
	return s
}

func (s *DeleteSCIMServerCredentialResponse) SetStatusCode(v int32) *DeleteSCIMServerCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSCIMServerCredentialResponse) SetBody(v *DeleteSCIMServerCredentialResponseBody) *DeleteSCIMServerCredentialResponse {
	s.Body = v
	return s
}

type DeleteUserRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserRequest) SetDirectoryId(v string) *DeleteUserRequest {
	s.DirectoryId = &v
	return s
}

func (s *DeleteUserRequest) SetUserId(v string) *DeleteUserRequest {
	s.UserId = &v
	return s
}

type DeleteUserResponseBody struct {
	// The ID of the request.
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

type DeleteUserProvisioningRequest struct {
	DeletionStrategy   *string `json:"DeletionStrategy,omitempty" xml:"DeletionStrategy,omitempty"`
	DirectoryId        *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	UserProvisioningId *string `json:"UserProvisioningId,omitempty" xml:"UserProvisioningId,omitempty"`
}

func (s DeleteUserProvisioningRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserProvisioningRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserProvisioningRequest) SetDeletionStrategy(v string) *DeleteUserProvisioningRequest {
	s.DeletionStrategy = &v
	return s
}

func (s *DeleteUserProvisioningRequest) SetDirectoryId(v string) *DeleteUserProvisioningRequest {
	s.DirectoryId = &v
	return s
}

func (s *DeleteUserProvisioningRequest) SetUserProvisioningId(v string) *DeleteUserProvisioningRequest {
	s.UserProvisioningId = &v
	return s
}

type DeleteUserProvisioningResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserProvisioningResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserProvisioningResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserProvisioningResponseBody) SetRequestId(v string) *DeleteUserProvisioningResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserProvisioningResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserProvisioningResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserProvisioningResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserProvisioningResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserProvisioningResponse) SetHeaders(v map[string]*string) *DeleteUserProvisioningResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserProvisioningResponse) SetStatusCode(v int32) *DeleteUserProvisioningResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserProvisioningResponse) SetBody(v *DeleteUserProvisioningResponseBody) *DeleteUserProvisioningResponse {
	s.Body = v
	return s
}

type DeleteUserProvisioningEventRequest struct {
	DirectoryId        *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EventId            *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	UserProvisioningId *string `json:"UserProvisioningId,omitempty" xml:"UserProvisioningId,omitempty"`
}

func (s DeleteUserProvisioningEventRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserProvisioningEventRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserProvisioningEventRequest) SetDirectoryId(v string) *DeleteUserProvisioningEventRequest {
	s.DirectoryId = &v
	return s
}

func (s *DeleteUserProvisioningEventRequest) SetEventId(v string) *DeleteUserProvisioningEventRequest {
	s.EventId = &v
	return s
}

func (s *DeleteUserProvisioningEventRequest) SetUserProvisioningId(v string) *DeleteUserProvisioningEventRequest {
	s.UserProvisioningId = &v
	return s
}

type DeleteUserProvisioningEventResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserProvisioningEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserProvisioningEventResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserProvisioningEventResponseBody) SetRequestId(v string) *DeleteUserProvisioningEventResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserProvisioningEventResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserProvisioningEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserProvisioningEventResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserProvisioningEventResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserProvisioningEventResponse) SetHeaders(v map[string]*string) *DeleteUserProvisioningEventResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserProvisioningEventResponse) SetStatusCode(v int32) *DeleteUserProvisioningEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserProvisioningEventResponse) SetBody(v *DeleteUserProvisioningEventResponseBody) *DeleteUserProvisioningEventResponse {
	s.Body = v
	return s
}

type DeprovisionAccessConfigurationRequest struct {
	// The ID of the access configuration.
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The directory ID.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the task object.
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the task object. Set the value to RD-Account, which specifies the accounts in the resource directory.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s DeprovisionAccessConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeprovisionAccessConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DeprovisionAccessConfigurationRequest) SetAccessConfigurationId(v string) *DeprovisionAccessConfigurationRequest {
	s.AccessConfigurationId = &v
	return s
}

func (s *DeprovisionAccessConfigurationRequest) SetDirectoryId(v string) *DeprovisionAccessConfigurationRequest {
	s.DirectoryId = &v
	return s
}

func (s *DeprovisionAccessConfigurationRequest) SetTargetId(v string) *DeprovisionAccessConfigurationRequest {
	s.TargetId = &v
	return s
}

func (s *DeprovisionAccessConfigurationRequest) SetTargetType(v string) *DeprovisionAccessConfigurationRequest {
	s.TargetType = &v
	return s
}

type DeprovisionAccessConfigurationResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task information.
	Tasks []*DeprovisionAccessConfigurationResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s DeprovisionAccessConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeprovisionAccessConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DeprovisionAccessConfigurationResponseBody) SetRequestId(v string) *DeprovisionAccessConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeprovisionAccessConfigurationResponseBody) SetTasks(v []*DeprovisionAccessConfigurationResponseBodyTasks) *DeprovisionAccessConfigurationResponseBody {
	s.Tasks = v
	return s
}

type DeprovisionAccessConfigurationResponseBodyTasks struct {
	// The ID of the access configuration.
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The name of the access configuration.
	AccessConfigurationName *string `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	// The task status. Valid values:
	//
	// *   InProgress: The task is running.
	// *   Success: The task is successful.
	// *   Failed: The task failed.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the task object.
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the task object.
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The path ID of the task object in the resource directory.
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// The path name of the task object in the resource directory.
	TargetPathName *string `json:"TargetPathName,omitempty" xml:"TargetPathName,omitempty"`
	// The type of the task object. The value is fixed as RD-Account, which indicates the accounts in the resource directory.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The task ID.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type. The value is fixed as DeprovisionAccessConfiguration, which indicates that the access configuration is de-provisioned.
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DeprovisionAccessConfigurationResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s DeprovisionAccessConfigurationResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) SetAccessConfigurationId(v string) *DeprovisionAccessConfigurationResponseBodyTasks {
	s.AccessConfigurationId = &v
	return s
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) SetAccessConfigurationName(v string) *DeprovisionAccessConfigurationResponseBodyTasks {
	s.AccessConfigurationName = &v
	return s
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) SetStatus(v string) *DeprovisionAccessConfigurationResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) SetTargetId(v string) *DeprovisionAccessConfigurationResponseBodyTasks {
	s.TargetId = &v
	return s
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) SetTargetName(v string) *DeprovisionAccessConfigurationResponseBodyTasks {
	s.TargetName = &v
	return s
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) SetTargetPath(v string) *DeprovisionAccessConfigurationResponseBodyTasks {
	s.TargetPath = &v
	return s
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) SetTargetPathName(v string) *DeprovisionAccessConfigurationResponseBodyTasks {
	s.TargetPathName = &v
	return s
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) SetTargetType(v string) *DeprovisionAccessConfigurationResponseBodyTasks {
	s.TargetType = &v
	return s
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) SetTaskId(v string) *DeprovisionAccessConfigurationResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) SetTaskType(v string) *DeprovisionAccessConfigurationResponseBodyTasks {
	s.TaskType = &v
	return s
}

type DeprovisionAccessConfigurationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeprovisionAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeprovisionAccessConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeprovisionAccessConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DeprovisionAccessConfigurationResponse) SetHeaders(v map[string]*string) *DeprovisionAccessConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DeprovisionAccessConfigurationResponse) SetStatusCode(v int32) *DeprovisionAccessConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeprovisionAccessConfigurationResponse) SetBody(v *DeprovisionAccessConfigurationResponseBody) *DeprovisionAccessConfigurationResponse {
	s.Body = v
	return s
}

type DisableServiceResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DisableServiceResponseBody) SetRequestId(v string) *DisableServiceResponseBody {
	s.RequestId = &v
	return s
}

type DisableServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableServiceResponse) GoString() string {
	return s.String()
}

func (s *DisableServiceResponse) SetHeaders(v map[string]*string) *DisableServiceResponse {
	s.Headers = v
	return s
}

func (s *DisableServiceResponse) SetStatusCode(v int32) *DisableServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableServiceResponse) SetBody(v *DisableServiceResponseBody) *DisableServiceResponse {
	s.Body = v
	return s
}

type EnableServiceResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableServiceResponseBody) GoString() string {
	return s.String()
}

func (s *EnableServiceResponseBody) SetRequestId(v string) *EnableServiceResponseBody {
	s.RequestId = &v
	return s
}

type EnableServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableServiceResponse) GoString() string {
	return s.String()
}

func (s *EnableServiceResponse) SetHeaders(v map[string]*string) *EnableServiceResponse {
	s.Headers = v
	return s
}

func (s *EnableServiceResponse) SetStatusCode(v int32) *EnableServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableServiceResponse) SetBody(v *EnableServiceResponseBody) *EnableServiceResponse {
	s.Body = v
	return s
}

type GetAccessConfigurationRequest struct {
	// The ID of the access configuration.
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s GetAccessConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccessConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetAccessConfigurationRequest) SetAccessConfigurationId(v string) *GetAccessConfigurationRequest {
	s.AccessConfigurationId = &v
	return s
}

func (s *GetAccessConfigurationRequest) SetDirectoryId(v string) *GetAccessConfigurationRequest {
	s.DirectoryId = &v
	return s
}

type GetAccessConfigurationResponseBody struct {
	// The information about the access configuration.
	AccessConfiguration *GetAccessConfigurationResponseBodyAccessConfiguration `json:"AccessConfiguration,omitempty" xml:"AccessConfiguration,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccessConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccessConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessConfigurationResponseBody) SetAccessConfiguration(v *GetAccessConfigurationResponseBodyAccessConfiguration) *GetAccessConfigurationResponseBody {
	s.AccessConfiguration = v
	return s
}

func (s *GetAccessConfigurationResponseBody) SetRequestId(v string) *GetAccessConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type GetAccessConfigurationResponseBodyAccessConfiguration struct {
	// The ID of the access configuration.
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The name of the access configuration.
	AccessConfigurationName *string `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	// The time when the access configuration was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the access configuration.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The initial web page that is displayed after a CloudSSO user accesses an account in your resource directory by using the access configuration.
	RelayState *string `json:"RelayState,omitempty" xml:"RelayState,omitempty"`
	// The duration of a session in which a CloudSSO user accesses an account in your resource directory by using the access configuration.
	//
	// Unit: seconds.
	SessionDuration *int32 `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	// The status notification.
	StatusNotifications []*string `json:"StatusNotifications,omitempty" xml:"StatusNotifications,omitempty" type:"Repeated"`
	// The time when the information about the access configuration was modified.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetAccessConfigurationResponseBodyAccessConfiguration) String() string {
	return tea.Prettify(s)
}

func (s GetAccessConfigurationResponseBodyAccessConfiguration) GoString() string {
	return s.String()
}

func (s *GetAccessConfigurationResponseBodyAccessConfiguration) SetAccessConfigurationId(v string) *GetAccessConfigurationResponseBodyAccessConfiguration {
	s.AccessConfigurationId = &v
	return s
}

func (s *GetAccessConfigurationResponseBodyAccessConfiguration) SetAccessConfigurationName(v string) *GetAccessConfigurationResponseBodyAccessConfiguration {
	s.AccessConfigurationName = &v
	return s
}

func (s *GetAccessConfigurationResponseBodyAccessConfiguration) SetCreateTime(v string) *GetAccessConfigurationResponseBodyAccessConfiguration {
	s.CreateTime = &v
	return s
}

func (s *GetAccessConfigurationResponseBodyAccessConfiguration) SetDescription(v string) *GetAccessConfigurationResponseBodyAccessConfiguration {
	s.Description = &v
	return s
}

func (s *GetAccessConfigurationResponseBodyAccessConfiguration) SetRelayState(v string) *GetAccessConfigurationResponseBodyAccessConfiguration {
	s.RelayState = &v
	return s
}

func (s *GetAccessConfigurationResponseBodyAccessConfiguration) SetSessionDuration(v int32) *GetAccessConfigurationResponseBodyAccessConfiguration {
	s.SessionDuration = &v
	return s
}

func (s *GetAccessConfigurationResponseBodyAccessConfiguration) SetStatusNotifications(v []*string) *GetAccessConfigurationResponseBodyAccessConfiguration {
	s.StatusNotifications = v
	return s
}

func (s *GetAccessConfigurationResponseBodyAccessConfiguration) SetUpdateTime(v string) *GetAccessConfigurationResponseBodyAccessConfiguration {
	s.UpdateTime = &v
	return s
}

type GetAccessConfigurationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccessConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccessConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetAccessConfigurationResponse) SetHeaders(v map[string]*string) *GetAccessConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetAccessConfigurationResponse) SetStatusCode(v int32) *GetAccessConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessConfigurationResponse) SetBody(v *GetAccessConfigurationResponseBody) *GetAccessConfigurationResponse {
	s.Body = v
	return s
}

type GetDirectoryRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s GetDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDirectoryRequest) GoString() string {
	return s.String()
}

func (s *GetDirectoryRequest) SetDirectoryId(v string) *GetDirectoryRequest {
	s.DirectoryId = &v
	return s
}

type GetDirectoryResponseBody struct {
	// The information about the directory.
	Directory *GetDirectoryResponseBodyDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetDirectoryResponseBody) SetDirectory(v *GetDirectoryResponseBodyDirectory) *GetDirectoryResponseBody {
	s.Directory = v
	return s
}

func (s *GetDirectoryResponseBody) SetRequestId(v string) *GetDirectoryResponseBody {
	s.RequestId = &v
	return s
}

type GetDirectoryResponseBodyDirectory struct {
	// The time when the directory was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The name of the directory.
	DirectoryName *string `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	// The region ID of the directory.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The time when the directory was modified.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetDirectoryResponseBodyDirectory) String() string {
	return tea.Prettify(s)
}

func (s GetDirectoryResponseBodyDirectory) GoString() string {
	return s.String()
}

func (s *GetDirectoryResponseBodyDirectory) SetCreateTime(v string) *GetDirectoryResponseBodyDirectory {
	s.CreateTime = &v
	return s
}

func (s *GetDirectoryResponseBodyDirectory) SetDirectoryId(v string) *GetDirectoryResponseBodyDirectory {
	s.DirectoryId = &v
	return s
}

func (s *GetDirectoryResponseBodyDirectory) SetDirectoryName(v string) *GetDirectoryResponseBodyDirectory {
	s.DirectoryName = &v
	return s
}

func (s *GetDirectoryResponseBodyDirectory) SetRegion(v string) *GetDirectoryResponseBodyDirectory {
	s.Region = &v
	return s
}

func (s *GetDirectoryResponseBodyDirectory) SetUpdateTime(v string) *GetDirectoryResponseBodyDirectory {
	s.UpdateTime = &v
	return s
}

type GetDirectoryResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDirectoryResponse) GoString() string {
	return s.String()
}

func (s *GetDirectoryResponse) SetHeaders(v map[string]*string) *GetDirectoryResponse {
	s.Headers = v
	return s
}

func (s *GetDirectoryResponse) SetStatusCode(v int32) *GetDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDirectoryResponse) SetBody(v *GetDirectoryResponseBody) *GetDirectoryResponse {
	s.Body = v
	return s
}

type GetDirectorySAMLServiceProviderInfoRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s GetDirectorySAMLServiceProviderInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDirectorySAMLServiceProviderInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDirectorySAMLServiceProviderInfoRequest) SetDirectoryId(v string) *GetDirectorySAMLServiceProviderInfoRequest {
	s.DirectoryId = &v
	return s
}

type GetDirectorySAMLServiceProviderInfoResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the SP.
	SAMLServiceProvider *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider `json:"SAMLServiceProvider,omitempty" xml:"SAMLServiceProvider,omitempty" type:"Struct"`
}

func (s GetDirectorySAMLServiceProviderInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDirectorySAMLServiceProviderInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDirectorySAMLServiceProviderInfoResponseBody) SetRequestId(v string) *GetDirectorySAMLServiceProviderInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDirectorySAMLServiceProviderInfoResponseBody) SetSAMLServiceProvider(v *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider) *GetDirectorySAMLServiceProviderInfoResponseBody {
	s.SAMLServiceProvider = v
	return s
}

type GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider struct {
	// The Assertion Consumer Service (ACS) URL of the SP.
	AcsUrl          *string `json:"AcsUrl,omitempty" xml:"AcsUrl,omitempty"`
	AuthnSignAlgo   *string `json:"AuthnSignAlgo,omitempty" xml:"AuthnSignAlgo,omitempty"`
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The metadata file of the SP. The value of this parameter is Base64-encoded.
	EncodedMetadataDocument *string `json:"EncodedMetadataDocument,omitempty" xml:"EncodedMetadataDocument,omitempty"`
	// The entity ID of the SP.
	EntityId                  *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	SupportEncryptedAssertion *bool   `json:"SupportEncryptedAssertion,omitempty" xml:"SupportEncryptedAssertion,omitempty"`
}

func (s GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider) String() string {
	return tea.Prettify(s)
}

func (s GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider) GoString() string {
	return s.String()
}

func (s *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider) SetAcsUrl(v string) *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider {
	s.AcsUrl = &v
	return s
}

func (s *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider) SetAuthnSignAlgo(v string) *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider {
	s.AuthnSignAlgo = &v
	return s
}

func (s *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider) SetCertificateType(v string) *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider {
	s.CertificateType = &v
	return s
}

func (s *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider) SetDirectoryId(v string) *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider {
	s.DirectoryId = &v
	return s
}

func (s *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider) SetEncodedMetadataDocument(v string) *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider {
	s.EncodedMetadataDocument = &v
	return s
}

func (s *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider) SetEntityId(v string) *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider {
	s.EntityId = &v
	return s
}

func (s *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider) SetSupportEncryptedAssertion(v bool) *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider {
	s.SupportEncryptedAssertion = &v
	return s
}

type GetDirectorySAMLServiceProviderInfoResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDirectorySAMLServiceProviderInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDirectorySAMLServiceProviderInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDirectorySAMLServiceProviderInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDirectorySAMLServiceProviderInfoResponse) SetHeaders(v map[string]*string) *GetDirectorySAMLServiceProviderInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDirectorySAMLServiceProviderInfoResponse) SetStatusCode(v int32) *GetDirectorySAMLServiceProviderInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDirectorySAMLServiceProviderInfoResponse) SetBody(v *GetDirectorySAMLServiceProviderInfoResponseBody) *GetDirectorySAMLServiceProviderInfoResponse {
	s.Body = v
	return s
}

type GetDirectoryStatisticsRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s GetDirectoryStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDirectoryStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetDirectoryStatisticsRequest) SetDirectoryId(v string) *GetDirectoryStatisticsRequest {
	s.DirectoryId = &v
	return s
}

type GetDirectoryStatisticsResponseBody struct {
	// The statistics of the directory.
	DirectoryStatistics *GetDirectoryStatisticsResponseBodyDirectoryStatistics `json:"DirectoryStatistics,omitempty" xml:"DirectoryStatistics,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDirectoryStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDirectoryStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDirectoryStatisticsResponseBody) SetDirectoryStatistics(v *GetDirectoryStatisticsResponseBodyDirectoryStatistics) *GetDirectoryStatisticsResponseBody {
	s.DirectoryStatistics = v
	return s
}

func (s *GetDirectoryStatisticsResponseBody) SetRequestId(v string) *GetDirectoryStatisticsResponseBody {
	s.RequestId = &v
	return s
}

type GetDirectoryStatisticsResponseBodyDirectoryStatistics struct {
	// The number of access permissions that are assigned.
	AccessAssignmentCount *int32 `json:"AccessAssignmentCount,omitempty" xml:"AccessAssignmentCount,omitempty"`
	// The number of access configurations.
	AccessConfigurationCount *int32 `json:"AccessConfigurationCount,omitempty" xml:"AccessConfigurationCount,omitempty"`
	// The quota for access configurations.
	AccessConfigurationQuota *int32 `json:"AccessConfigurationQuota,omitempty" xml:"AccessConfigurationQuota,omitempty"`
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The name of the directory.
	DirectoryName *string `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	// The number of user groups.
	GroupCount *int32 `json:"GroupCount,omitempty" xml:"GroupCount,omitempty"`
	// The quota for user groups.
	GroupQuota *int32 `json:"GroupQuota,omitempty" xml:"GroupQuota,omitempty"`
	// The number of tasks that are being performed.
	InProgressTaskCount *int32 `json:"InProgressTaskCount,omitempty" xml:"InProgressTaskCount,omitempty"`
	// The region ID of the directory.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The number of SCIM credentials.
	SCIMServerCredentialCount *int32 `json:"SCIMServerCredentialCount,omitempty" xml:"SCIMServerCredentialCount,omitempty"`
	// Indicates whether SCIM synchronization is enabled. Valid values:
	//
	// *   true
	// *   false
	SCIMSyncEnabled *bool `json:"SCIMSyncEnabled,omitempty" xml:"SCIMSyncEnabled,omitempty"`
	// Indicates whether SSO is enabled. Valid values:
	//
	// *   true
	// *   false
	SSOEnabled *bool `json:"SSOEnabled,omitempty" xml:"SSOEnabled,omitempty"`
	// The quota for system policies that can be configured for an access configuration.
	SystemPolicyPerAccessConfigurationQuota *int32 `json:"SystemPolicyPerAccessConfigurationQuota,omitempty" xml:"SystemPolicyPerAccessConfigurationQuota,omitempty"`
	// The number of users.
	UserCount *int32 `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
	// The quota for users.
	UserQuota *int32 `json:"UserQuota,omitempty" xml:"UserQuota,omitempty"`
}

func (s GetDirectoryStatisticsResponseBodyDirectoryStatistics) String() string {
	return tea.Prettify(s)
}

func (s GetDirectoryStatisticsResponseBodyDirectoryStatistics) GoString() string {
	return s.String()
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetAccessAssignmentCount(v int32) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.AccessAssignmentCount = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetAccessConfigurationCount(v int32) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.AccessConfigurationCount = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetAccessConfigurationQuota(v int32) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.AccessConfigurationQuota = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetDirectoryId(v string) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.DirectoryId = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetDirectoryName(v string) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.DirectoryName = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetGroupCount(v int32) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.GroupCount = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetGroupQuota(v int32) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.GroupQuota = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetInProgressTaskCount(v int32) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.InProgressTaskCount = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetRegion(v string) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.Region = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetSCIMServerCredentialCount(v int32) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.SCIMServerCredentialCount = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetSCIMSyncEnabled(v bool) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.SCIMSyncEnabled = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetSSOEnabled(v bool) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.SSOEnabled = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetSystemPolicyPerAccessConfigurationQuota(v int32) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.SystemPolicyPerAccessConfigurationQuota = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetUserCount(v int32) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.UserCount = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetUserQuota(v int32) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.UserQuota = &v
	return s
}

type GetDirectoryStatisticsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDirectoryStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDirectoryStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDirectoryStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetDirectoryStatisticsResponse) SetHeaders(v map[string]*string) *GetDirectoryStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetDirectoryStatisticsResponse) SetStatusCode(v int32) *GetDirectoryStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDirectoryStatisticsResponse) SetBody(v *GetDirectoryStatisticsResponseBody) *GetDirectoryStatisticsResponse {
	s.Body = v
	return s
}

type GetExternalSAMLIdentityProviderRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s GetExternalSAMLIdentityProviderRequest) String() string {
	return tea.Prettify(s)
}

func (s GetExternalSAMLIdentityProviderRequest) GoString() string {
	return s.String()
}

func (s *GetExternalSAMLIdentityProviderRequest) SetDirectoryId(v string) *GetExternalSAMLIdentityProviderRequest {
	s.DirectoryId = &v
	return s
}

type GetExternalSAMLIdentityProviderResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configurations of the IdP.
	SAMLIdentityProviderConfiguration *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration `json:"SAMLIdentityProviderConfiguration,omitempty" xml:"SAMLIdentityProviderConfiguration,omitempty" type:"Struct"`
}

func (s GetExternalSAMLIdentityProviderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExternalSAMLIdentityProviderResponseBody) GoString() string {
	return s.String()
}

func (s *GetExternalSAMLIdentityProviderResponseBody) SetRequestId(v string) *GetExternalSAMLIdentityProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExternalSAMLIdentityProviderResponseBody) SetSAMLIdentityProviderConfiguration(v *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) *GetExternalSAMLIdentityProviderResponseBody {
	s.SAMLIdentityProviderConfiguration = v
	return s
}

type GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration struct {
	// The ID of the SAML signing certificate.
	CertificateIds []*string `json:"CertificateIds,omitempty" xml:"CertificateIds,omitempty" type:"Repeated"`
	// The time when the IdP was configured for the first time.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The metadata file of the IdP. The value of this parameter is Base64-encoded.
	EncodedMetadataDocument *string `json:"EncodedMetadataDocument,omitempty" xml:"EncodedMetadataDocument,omitempty"`
	// The entity ID of the IdP.
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The logon URL of the IdP.
	LoginUrl *string `json:"LoginUrl,omitempty" xml:"LoginUrl,omitempty"`
	// The status of SSO logon. Valid values:
	//
	// *   Enabled
	// *   Disabled
	SSOStatus *string `json:"SSOStatus,omitempty" xml:"SSOStatus,omitempty"`
	// The time when the IdP configurations were last modified.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// Indicates whether CloudSSO needs to sign SAML requests. The requests are sent when users log on to the CloudSSO user portal to initiate SAML-based SSO. Valid values:
	//
	// *   true: yes
	// *   false: no (default)
	WantRequestSigned *bool `json:"WantRequestSigned,omitempty" xml:"WantRequestSigned,omitempty"`
}

func (s GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) String() string {
	return tea.Prettify(s)
}

func (s GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) GoString() string {
	return s.String()
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetCertificateIds(v []*string) *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.CertificateIds = v
	return s
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetCreateTime(v string) *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.CreateTime = &v
	return s
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetDirectoryId(v string) *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.DirectoryId = &v
	return s
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetEncodedMetadataDocument(v string) *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.EncodedMetadataDocument = &v
	return s
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetEntityId(v string) *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.EntityId = &v
	return s
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetLoginUrl(v string) *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.LoginUrl = &v
	return s
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetSSOStatus(v string) *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.SSOStatus = &v
	return s
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetUpdateTime(v string) *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.UpdateTime = &v
	return s
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetWantRequestSigned(v bool) *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.WantRequestSigned = &v
	return s
}

type GetExternalSAMLIdentityProviderResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExternalSAMLIdentityProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExternalSAMLIdentityProviderResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExternalSAMLIdentityProviderResponse) GoString() string {
	return s.String()
}

func (s *GetExternalSAMLIdentityProviderResponse) SetHeaders(v map[string]*string) *GetExternalSAMLIdentityProviderResponse {
	s.Headers = v
	return s
}

func (s *GetExternalSAMLIdentityProviderResponse) SetStatusCode(v int32) *GetExternalSAMLIdentityProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExternalSAMLIdentityProviderResponse) SetBody(v *GetExternalSAMLIdentityProviderResponseBody) *GetExternalSAMLIdentityProviderResponse {
	s.Body = v
	return s
}

type GetGroupRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the group.
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s GetGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGroupRequest) GoString() string {
	return s.String()
}

func (s *GetGroupRequest) SetDirectoryId(v string) *GetGroupRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetGroupRequest) SetGroupId(v string) *GetGroupRequest {
	s.GroupId = &v
	return s
}

type GetGroupResponseBody struct {
	// The information about the group.
	Group *GetGroupResponseBodyGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	// The ID of the request.
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
	// The time when the group was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the group.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the group.
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the group. Valid values:
	//
	// *   Manual: The group is manually created.
	// *   Synchronized: The group is synchronized from an external identity provider (IdP).
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	// The time when the information about the group was modified.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetGroupResponseBodyGroup) String() string {
	return tea.Prettify(s)
}

func (s GetGroupResponseBodyGroup) GoString() string {
	return s.String()
}

func (s *GetGroupResponseBodyGroup) SetCreateTime(v string) *GetGroupResponseBodyGroup {
	s.CreateTime = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetDescription(v string) *GetGroupResponseBodyGroup {
	s.Description = &v
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

func (s *GetGroupResponseBodyGroup) SetProvisionType(v string) *GetGroupResponseBodyGroup {
	s.ProvisionType = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetUpdateTime(v string) *GetGroupResponseBodyGroup {
	s.UpdateTime = &v
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

type GetLoginPreferenceRequest struct {
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s GetLoginPreferenceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLoginPreferenceRequest) GoString() string {
	return s.String()
}

func (s *GetLoginPreferenceRequest) SetDirectoryId(v string) *GetLoginPreferenceRequest {
	s.DirectoryId = &v
	return s
}

type GetLoginPreferenceResponseBody struct {
	LoginPreference *GetLoginPreferenceResponseBodyLoginPreference `json:"LoginPreference,omitempty" xml:"LoginPreference,omitempty" type:"Struct"`
	RequestId       *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLoginPreferenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLoginPreferenceResponseBody) GoString() string {
	return s.String()
}

func (s *GetLoginPreferenceResponseBody) SetLoginPreference(v *GetLoginPreferenceResponseBodyLoginPreference) *GetLoginPreferenceResponseBody {
	s.LoginPreference = v
	return s
}

func (s *GetLoginPreferenceResponseBody) SetRequestId(v string) *GetLoginPreferenceResponseBody {
	s.RequestId = &v
	return s
}

type GetLoginPreferenceResponseBodyLoginPreference struct {
	LoginNetworkMasks *string `json:"LoginNetworkMasks,omitempty" xml:"LoginNetworkMasks,omitempty"`
}

func (s GetLoginPreferenceResponseBodyLoginPreference) String() string {
	return tea.Prettify(s)
}

func (s GetLoginPreferenceResponseBodyLoginPreference) GoString() string {
	return s.String()
}

func (s *GetLoginPreferenceResponseBodyLoginPreference) SetLoginNetworkMasks(v string) *GetLoginPreferenceResponseBodyLoginPreference {
	s.LoginNetworkMasks = &v
	return s
}

type GetLoginPreferenceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLoginPreferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLoginPreferenceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLoginPreferenceResponse) GoString() string {
	return s.String()
}

func (s *GetLoginPreferenceResponse) SetHeaders(v map[string]*string) *GetLoginPreferenceResponse {
	s.Headers = v
	return s
}

func (s *GetLoginPreferenceResponse) SetStatusCode(v int32) *GetLoginPreferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLoginPreferenceResponse) SetBody(v *GetLoginPreferenceResponseBody) *GetLoginPreferenceResponse {
	s.Body = v
	return s
}

type GetMFAAuthenticationSettingInfoRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s GetMFAAuthenticationSettingInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMFAAuthenticationSettingInfoRequest) GoString() string {
	return s.String()
}

func (s *GetMFAAuthenticationSettingInfoRequest) SetDirectoryId(v string) *GetMFAAuthenticationSettingInfoRequest {
	s.DirectoryId = &v
	return s
}

type GetMFAAuthenticationSettingInfoResponseBody struct {
	// The MFA setting of all users.
	MFAAuthenticationSettingInfo *GetMFAAuthenticationSettingInfoResponseBodyMFAAuthenticationSettingInfo `json:"MFAAuthenticationSettingInfo,omitempty" xml:"MFAAuthenticationSettingInfo,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMFAAuthenticationSettingInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMFAAuthenticationSettingInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetMFAAuthenticationSettingInfoResponseBody) SetMFAAuthenticationSettingInfo(v *GetMFAAuthenticationSettingInfoResponseBodyMFAAuthenticationSettingInfo) *GetMFAAuthenticationSettingInfoResponseBody {
	s.MFAAuthenticationSettingInfo = v
	return s
}

func (s *GetMFAAuthenticationSettingInfoResponseBody) SetRequestId(v string) *GetMFAAuthenticationSettingInfoResponseBody {
	s.RequestId = &v
	return s
}

type GetMFAAuthenticationSettingInfoResponseBodyMFAAuthenticationSettingInfo struct {
	// The MFA policy of all users. Valid values:
	//
	// *   Enabled: MFA is enabled for all users.
	// *   Byuser: User-specific settings are applied. For more information about how to configure MFA for a single user, see [UpdateUserMFAAuthenticationSettings](~~450135~~).
	// *   Disabled: MFA is disabled for all users.
	// *   OnlyRiskyLogin: MFA is required only for unusual logons.
	MfaAuthenticationAdvanceSettings *string `json:"MfaAuthenticationAdvanceSettings,omitempty" xml:"MfaAuthenticationAdvanceSettings,omitempty"`
	// The MFA policy for unusual logons. Valid values:
	//
	// *   Autonomous: MFA is prompted for users who initiated unusual logons. However, the users are allowed to skip MFA. If an MFA device is bound to a user who initiated an unusual logon, the user must pass MFA.
	// *   EnforceVerify: MFA is required. If no MFA devices are bound to a user who initiated an unusual logon, the user must bind an MFA device. If an MFA device is already bound to a user who initiated an unusual logon, the user must pass MFA.
	//
	// > This parameter is displayed only when Byuser or OnlyRiskyLogin is returned for the MfaAuthenticationAdvanceSettings parameter.
	OperationForRiskLogin *string `json:"OperationForRiskLogin,omitempty" xml:"OperationForRiskLogin,omitempty"`
}

func (s GetMFAAuthenticationSettingInfoResponseBodyMFAAuthenticationSettingInfo) String() string {
	return tea.Prettify(s)
}

func (s GetMFAAuthenticationSettingInfoResponseBodyMFAAuthenticationSettingInfo) GoString() string {
	return s.String()
}

func (s *GetMFAAuthenticationSettingInfoResponseBodyMFAAuthenticationSettingInfo) SetMfaAuthenticationAdvanceSettings(v string) *GetMFAAuthenticationSettingInfoResponseBodyMFAAuthenticationSettingInfo {
	s.MfaAuthenticationAdvanceSettings = &v
	return s
}

func (s *GetMFAAuthenticationSettingInfoResponseBodyMFAAuthenticationSettingInfo) SetOperationForRiskLogin(v string) *GetMFAAuthenticationSettingInfoResponseBodyMFAAuthenticationSettingInfo {
	s.OperationForRiskLogin = &v
	return s
}

type GetMFAAuthenticationSettingInfoResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMFAAuthenticationSettingInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMFAAuthenticationSettingInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMFAAuthenticationSettingInfoResponse) GoString() string {
	return s.String()
}

func (s *GetMFAAuthenticationSettingInfoResponse) SetHeaders(v map[string]*string) *GetMFAAuthenticationSettingInfoResponse {
	s.Headers = v
	return s
}

func (s *GetMFAAuthenticationSettingInfoResponse) SetStatusCode(v int32) *GetMFAAuthenticationSettingInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMFAAuthenticationSettingInfoResponse) SetBody(v *GetMFAAuthenticationSettingInfoResponseBody) *GetMFAAuthenticationSettingInfoResponse {
	s.Body = v
	return s
}

type GetMFAAuthenticationSettingsRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s GetMFAAuthenticationSettingsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMFAAuthenticationSettingsRequest) GoString() string {
	return s.String()
}

func (s *GetMFAAuthenticationSettingsRequest) SetDirectoryId(v string) *GetMFAAuthenticationSettingsRequest {
	s.DirectoryId = &v
	return s
}

type GetMFAAuthenticationSettingsResponseBody struct {
	// Indicates whether MFA is enabled for all users. Valid values:
	//
	// *   Enabled: MFA is enabled for all users.
	// *   Byuser: User-specific settings are applied.
	// *   Disabled: MFA is disabled for all users.
	MFAAuthenticationAdvanceSettings *string `json:"MFAAuthenticationAdvanceSettings,omitempty" xml:"MFAAuthenticationAdvanceSettings,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMFAAuthenticationSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMFAAuthenticationSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *GetMFAAuthenticationSettingsResponseBody) SetMFAAuthenticationAdvanceSettings(v string) *GetMFAAuthenticationSettingsResponseBody {
	s.MFAAuthenticationAdvanceSettings = &v
	return s
}

func (s *GetMFAAuthenticationSettingsResponseBody) SetRequestId(v string) *GetMFAAuthenticationSettingsResponseBody {
	s.RequestId = &v
	return s
}

type GetMFAAuthenticationSettingsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMFAAuthenticationSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMFAAuthenticationSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMFAAuthenticationSettingsResponse) GoString() string {
	return s.String()
}

func (s *GetMFAAuthenticationSettingsResponse) SetHeaders(v map[string]*string) *GetMFAAuthenticationSettingsResponse {
	s.Headers = v
	return s
}

func (s *GetMFAAuthenticationSettingsResponse) SetStatusCode(v int32) *GetMFAAuthenticationSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMFAAuthenticationSettingsResponse) SetBody(v *GetMFAAuthenticationSettingsResponseBody) *GetMFAAuthenticationSettingsResponse {
	s.Body = v
	return s
}

type GetMFAAuthenticationStatusRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s GetMFAAuthenticationStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMFAAuthenticationStatusRequest) GoString() string {
	return s.String()
}

func (s *GetMFAAuthenticationStatusRequest) SetDirectoryId(v string) *GetMFAAuthenticationStatusRequest {
	s.DirectoryId = &v
	return s
}

type GetMFAAuthenticationStatusResponseBody struct {
	// The status of MFA. Valid values:
	//
	// *   Enabled
	// *   Disabled
	MFAAuthenticationStatus *string `json:"MFAAuthenticationStatus,omitempty" xml:"MFAAuthenticationStatus,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMFAAuthenticationStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMFAAuthenticationStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetMFAAuthenticationStatusResponseBody) SetMFAAuthenticationStatus(v string) *GetMFAAuthenticationStatusResponseBody {
	s.MFAAuthenticationStatus = &v
	return s
}

func (s *GetMFAAuthenticationStatusResponseBody) SetRequestId(v string) *GetMFAAuthenticationStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetMFAAuthenticationStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMFAAuthenticationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMFAAuthenticationStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMFAAuthenticationStatusResponse) GoString() string {
	return s.String()
}

func (s *GetMFAAuthenticationStatusResponse) SetHeaders(v map[string]*string) *GetMFAAuthenticationStatusResponse {
	s.Headers = v
	return s
}

func (s *GetMFAAuthenticationStatusResponse) SetStatusCode(v int32) *GetMFAAuthenticationStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMFAAuthenticationStatusResponse) SetBody(v *GetMFAAuthenticationStatusResponseBody) *GetMFAAuthenticationStatusResponse {
	s.Body = v
	return s
}

type GetPasswordPolicyRequest struct {
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s GetPasswordPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPasswordPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetPasswordPolicyRequest) SetDirectoryId(v string) *GetPasswordPolicyRequest {
	s.DirectoryId = &v
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
	HardExpire                 *bool  `json:"HardExpire,omitempty" xml:"HardExpire,omitempty"`
	MaxLoginAttempts           *int32 `json:"MaxLoginAttempts,omitempty" xml:"MaxLoginAttempts,omitempty"`
	MaxPasswordAge             *int32 `json:"MaxPasswordAge,omitempty" xml:"MaxPasswordAge,omitempty"`
	MaxPasswordLength          *int32 `json:"MaxPasswordLength,omitempty" xml:"MaxPasswordLength,omitempty"`
	MinPasswordDifferentChars  *int32 `json:"MinPasswordDifferentChars,omitempty" xml:"MinPasswordDifferentChars,omitempty"`
	MinPasswordLength          *int32 `json:"MinPasswordLength,omitempty" xml:"MinPasswordLength,omitempty"`
	PasswordNotContainUsername *bool  `json:"PasswordNotContainUsername,omitempty" xml:"PasswordNotContainUsername,omitempty"`
	PasswordReusePrevention    *int32 `json:"PasswordReusePrevention,omitempty" xml:"PasswordReusePrevention,omitempty"`
	RequireLowerCaseChars      *bool  `json:"RequireLowerCaseChars,omitempty" xml:"RequireLowerCaseChars,omitempty"`
	RequireNumbers             *bool  `json:"RequireNumbers,omitempty" xml:"RequireNumbers,omitempty"`
	RequireSymbols             *bool  `json:"RequireSymbols,omitempty" xml:"RequireSymbols,omitempty"`
	RequireUpperCaseChars      *bool  `json:"RequireUpperCaseChars,omitempty" xml:"RequireUpperCaseChars,omitempty"`
}

func (s GetPasswordPolicyResponseBodyPasswordPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetPasswordPolicyResponseBodyPasswordPolicy) GoString() string {
	return s.String()
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetHardExpire(v bool) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.HardExpire = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetMaxLoginAttempts(v int32) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.MaxLoginAttempts = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetMaxPasswordAge(v int32) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.MaxPasswordAge = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetMaxPasswordLength(v int32) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.MaxPasswordLength = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetMinPasswordDifferentChars(v int32) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.MinPasswordDifferentChars = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetMinPasswordLength(v int32) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.MinPasswordLength = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetPasswordNotContainUsername(v bool) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.PasswordNotContainUsername = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetPasswordReusePrevention(v int32) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.PasswordReusePrevention = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetRequireLowerCaseChars(v bool) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.RequireLowerCaseChars = &v
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

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetRequireUpperCaseChars(v bool) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.RequireUpperCaseChars = &v
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

type GetSCIMSynchronizationStatusRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s GetSCIMSynchronizationStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSCIMSynchronizationStatusRequest) GoString() string {
	return s.String()
}

func (s *GetSCIMSynchronizationStatusRequest) SetDirectoryId(v string) *GetSCIMSynchronizationStatusRequest {
	s.DirectoryId = &v
	return s
}

type GetSCIMSynchronizationStatusResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of SCIM synchronization. Valid values:
	//
	// *   Enabled
	// *   Disabled
	SCIMSynchronizationStatus *string `json:"SCIMSynchronizationStatus,omitempty" xml:"SCIMSynchronizationStatus,omitempty"`
}

func (s GetSCIMSynchronizationStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSCIMSynchronizationStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetSCIMSynchronizationStatusResponseBody) SetRequestId(v string) *GetSCIMSynchronizationStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSCIMSynchronizationStatusResponseBody) SetSCIMSynchronizationStatus(v string) *GetSCIMSynchronizationStatusResponseBody {
	s.SCIMSynchronizationStatus = &v
	return s
}

type GetSCIMSynchronizationStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSCIMSynchronizationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSCIMSynchronizationStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSCIMSynchronizationStatusResponse) GoString() string {
	return s.String()
}

func (s *GetSCIMSynchronizationStatusResponse) SetHeaders(v map[string]*string) *GetSCIMSynchronizationStatusResponse {
	s.Headers = v
	return s
}

func (s *GetSCIMSynchronizationStatusResponse) SetStatusCode(v int32) *GetSCIMSynchronizationStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSCIMSynchronizationStatusResponse) SetBody(v *GetSCIMSynchronizationStatusResponseBody) *GetSCIMSynchronizationStatusResponse {
	s.Body = v
	return s
}

type GetServiceStatusResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status information of CloudSSO.
	ServiceStatus *GetServiceStatusResponseBodyServiceStatus `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty" type:"Struct"`
}

func (s GetServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceStatusResponseBody) SetRequestId(v string) *GetServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceStatusResponseBody) SetServiceStatus(v *GetServiceStatusResponseBodyServiceStatus) *GetServiceStatusResponseBody {
	s.ServiceStatus = v
	return s
}

type GetServiceStatusResponseBodyServiceStatus struct {
	// The ID of your Alibaba Cloud account.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Indicates whether you have permissions to enable CloudSSO. Valid values:
	//
	// *   Success: You have permissions to enable CloudSSO.
	// *   Failed: You do not have permissions to enable CloudSSO.
	//
	// >  The value of this parameter is returned only if the value of `Status` is `Disabled`.
	PrerequisiteCheckResult *string `json:"PrerequisiteCheckResult,omitempty" xml:"PrerequisiteCheckResult,omitempty"`
	// The ID of the region.
	RegionsInUse []*string `json:"RegionsInUse,omitempty" xml:"RegionsInUse,omitempty" type:"Repeated"`
	// Indicates whether CloudSSO is enabled. Valid values:
	//
	// *   Enabled
	// *   Disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetServiceStatusResponseBodyServiceStatus) String() string {
	return tea.Prettify(s)
}

func (s GetServiceStatusResponseBodyServiceStatus) GoString() string {
	return s.String()
}

func (s *GetServiceStatusResponseBodyServiceStatus) SetAccountId(v string) *GetServiceStatusResponseBodyServiceStatus {
	s.AccountId = &v
	return s
}

func (s *GetServiceStatusResponseBodyServiceStatus) SetPrerequisiteCheckResult(v string) *GetServiceStatusResponseBodyServiceStatus {
	s.PrerequisiteCheckResult = &v
	return s
}

func (s *GetServiceStatusResponseBodyServiceStatus) SetRegionsInUse(v []*string) *GetServiceStatusResponseBodyServiceStatus {
	s.RegionsInUse = v
	return s
}

func (s *GetServiceStatusResponseBodyServiceStatus) SetStatus(v string) *GetServiceStatusResponseBodyServiceStatus {
	s.Status = &v
	return s
}

type GetServiceStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetServiceStatusResponse) SetHeaders(v map[string]*string) *GetServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetServiceStatusResponse) SetStatusCode(v int32) *GetServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceStatusResponse) SetBody(v *GetServiceStatusResponseBody) *GetServiceStatusResponse {
	s.Body = v
	return s
}

type GetTaskRequest struct {
	// The directory ID.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The task ID.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskRequest) GoString() string {
	return s.String()
}

func (s *GetTaskRequest) SetDirectoryId(v string) *GetTaskRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetTaskRequest) SetTaskId(v string) *GetTaskRequest {
	s.TaskId = &v
	return s
}

type GetTaskResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task information.
	Task *GetTaskResponseBodyTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
}

func (s GetTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBody) SetRequestId(v string) *GetTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskResponseBody) SetTask(v *GetTaskResponseBodyTask) *GetTaskResponseBody {
	s.Task = v
	return s
}

type GetTaskResponseBodyTask struct {
	// The ID of the access configuration.
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The name of the access configuration.
	AccessConfigurationName *string `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	// The end time of the task.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The cause of the task failure.
	//
	// >  This parameter is returned only when the value of `Status` is `Failed`.
	FailureReason *string `json:"FailureReason,omitempty" xml:"FailureReason,omitempty"`
	// The ID of the CloudSSO identity.
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The name of the CloudSSO identity.
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The type of the CloudSSO identity. Valid values:
	//
	// *   User
	// *   Group
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The start time of the task.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task status. Valid values:
	//
	// *   InProgress: The task is running.
	// *   Success: The task is successful.
	// *   Failed: The task failed.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the task object.
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the task object.
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The path ID of the task object in the resource directory.
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// The path name of the task object in the resource directory.
	TargetPathName *string `json:"TargetPathName,omitempty" xml:"TargetPathName,omitempty"`
	// The type of the task object. The value is fixed as RD-Account, which indicates the accounts in the resource directory.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The task ID.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type. Valid values:
	//
	// *   ProvisionAccessConfiguration: An access configuration is provisioned.
	// *   DeprovisionAccessConfiguration: An access configuration is de-provisioned.
	// *   CreateAccessAssignment: Access permissions on an account in the resource directory are assigned.
	// *   DeleteAccessAssignment: Access permissions on an account in the resource directory are removed.
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetTaskResponseBodyTask) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponseBodyTask) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTask) SetAccessConfigurationId(v string) *GetTaskResponseBodyTask {
	s.AccessConfigurationId = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetAccessConfigurationName(v string) *GetTaskResponseBodyTask {
	s.AccessConfigurationName = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetEndTime(v string) *GetTaskResponseBodyTask {
	s.EndTime = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetFailureReason(v string) *GetTaskResponseBodyTask {
	s.FailureReason = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetPrincipalId(v string) *GetTaskResponseBodyTask {
	s.PrincipalId = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetPrincipalName(v string) *GetTaskResponseBodyTask {
	s.PrincipalName = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetPrincipalType(v string) *GetTaskResponseBodyTask {
	s.PrincipalType = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetStartTime(v string) *GetTaskResponseBodyTask {
	s.StartTime = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetStatus(v string) *GetTaskResponseBodyTask {
	s.Status = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTargetId(v string) *GetTaskResponseBodyTask {
	s.TargetId = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTargetName(v string) *GetTaskResponseBodyTask {
	s.TargetName = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTargetPath(v string) *GetTaskResponseBodyTask {
	s.TargetPath = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTargetPathName(v string) *GetTaskResponseBodyTask {
	s.TargetPathName = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTargetType(v string) *GetTaskResponseBodyTask {
	s.TargetType = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTaskId(v string) *GetTaskResponseBodyTask {
	s.TaskId = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTaskType(v string) *GetTaskResponseBodyTask {
	s.TaskType = &v
	return s
}

type GetTaskResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTaskResponse) SetHeaders(v map[string]*string) *GetTaskResponse {
	s.Headers = v
	return s
}

func (s *GetTaskResponse) SetStatusCode(v int32) *GetTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskResponse) SetBody(v *GetTaskResponseBody) *GetTaskResponse {
	s.Body = v
	return s
}

type GetTaskStatusRequest struct {
	// The directory ID.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The task ID.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetTaskStatusRequest) SetDirectoryId(v string) *GetTaskStatusRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetTaskStatusRequest) SetTaskId(v string) *GetTaskStatusRequest {
	s.TaskId = &v
	return s
}

type GetTaskStatusResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status information about the task.
	TaskStatus *GetTaskStatusResponseBodyTaskStatus `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty" type:"Struct"`
}

func (s GetTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskStatusResponseBody) SetRequestId(v string) *GetTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetTaskStatus(v *GetTaskStatusResponseBodyTaskStatus) *GetTaskStatusResponseBody {
	s.TaskStatus = v
	return s
}

type GetTaskStatusResponseBodyTaskStatus struct {
	// The end time of the task.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The cause of the task failure.
	//
	// >  This parameter is returned only when the value of `Status` is `Failed`.
	FailureReason *string `json:"FailureReason,omitempty" xml:"FailureReason,omitempty"`
	// The start time of the task.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task status. Valid values:
	//
	// *   InProgress: The task is running.
	// *   Success: The task is successful.
	// *   Failed: The task failed.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type. Valid values:
	//
	// *   ProvisionAccessConfiguration: An access configuration is provisioned.
	// *   DeprovisionAccessConfiguration: An access configuration is de-provisioned.
	// *   CreateAccessAssignment: Access permissions on an account in the resource directory are assigned.
	// *   DeleteAccessAssignment: Access permissions on an account in the resource directory are removed.
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetTaskStatusResponseBodyTaskStatus) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusResponseBodyTaskStatus) GoString() string {
	return s.String()
}

func (s *GetTaskStatusResponseBodyTaskStatus) SetEndTime(v string) *GetTaskStatusResponseBodyTaskStatus {
	s.EndTime = &v
	return s
}

func (s *GetTaskStatusResponseBodyTaskStatus) SetFailureReason(v string) *GetTaskStatusResponseBodyTaskStatus {
	s.FailureReason = &v
	return s
}

func (s *GetTaskStatusResponseBodyTaskStatus) SetStartTime(v string) *GetTaskStatusResponseBodyTaskStatus {
	s.StartTime = &v
	return s
}

func (s *GetTaskStatusResponseBodyTaskStatus) SetStatus(v string) *GetTaskStatusResponseBodyTaskStatus {
	s.Status = &v
	return s
}

func (s *GetTaskStatusResponseBodyTaskStatus) SetTaskId(v string) *GetTaskStatusResponseBodyTaskStatus {
	s.TaskId = &v
	return s
}

func (s *GetTaskStatusResponseBodyTaskStatus) SetTaskType(v string) *GetTaskStatusResponseBodyTaskStatus {
	s.TaskType = &v
	return s
}

type GetTaskStatusResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetTaskStatusResponse) SetHeaders(v map[string]*string) *GetTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetTaskStatusResponse) SetStatusCode(v int32) *GetTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskStatusResponse) SetBody(v *GetTaskStatusResponseBody) *GetTaskStatusResponse {
	s.Body = v
	return s
}

type GetUserRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) SetDirectoryId(v string) *GetUserRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetUserRequest) SetUserId(v string) *GetUserRequest {
	s.UserId = &v
	return s
}

type GetUserResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the user.
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
	// The time when the user was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the user.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the user.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the user.
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The first name of the user.
	FirstName *string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	// The last name of the user.
	LastName *string `json:"LastName,omitempty" xml:"LastName,omitempty"`
	// The type of the user. Valid values:
	//
	// *   Manual: The user is manually created.
	// *   Synchronized: The user is synchronized from an external identity provider (IdP).
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	// The status of the user. Valid values:
	//
	// *   Enabled: The logon of the user is enabled.
	// *   Disabled: The logon of the user is disabled.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the information about the user was modified.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The name of the user.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetUserResponseBodyUser) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyUser) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUser) SetCreateTime(v string) *GetUserResponseBodyUser {
	s.CreateTime = &v
	return s
}

func (s *GetUserResponseBodyUser) SetDescription(v string) *GetUserResponseBodyUser {
	s.Description = &v
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

func (s *GetUserResponseBodyUser) SetFirstName(v string) *GetUserResponseBodyUser {
	s.FirstName = &v
	return s
}

func (s *GetUserResponseBodyUser) SetLastName(v string) *GetUserResponseBodyUser {
	s.LastName = &v
	return s
}

func (s *GetUserResponseBodyUser) SetProvisionType(v string) *GetUserResponseBodyUser {
	s.ProvisionType = &v
	return s
}

func (s *GetUserResponseBodyUser) SetStatus(v string) *GetUserResponseBodyUser {
	s.Status = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUpdateTime(v string) *GetUserResponseBodyUser {
	s.UpdateTime = &v
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

type GetUserMFAAuthenticationSettingsRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetUserMFAAuthenticationSettingsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserMFAAuthenticationSettingsRequest) GoString() string {
	return s.String()
}

func (s *GetUserMFAAuthenticationSettingsRequest) SetDirectoryId(v string) *GetUserMFAAuthenticationSettingsRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetUserMFAAuthenticationSettingsRequest) SetUserId(v string) *GetUserMFAAuthenticationSettingsRequest {
	s.UserId = &v
	return s
}

type GetUserMFAAuthenticationSettingsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether MFA is enabled for the user. Valid values:
	//
	// *   Enabled: MFA is enabled for the user.
	// *   Disabled: MFA is disabled for the user.
	UserMFAAuthenticationSettings *string `json:"UserMFAAuthenticationSettings,omitempty" xml:"UserMFAAuthenticationSettings,omitempty"`
}

func (s GetUserMFAAuthenticationSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserMFAAuthenticationSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserMFAAuthenticationSettingsResponseBody) SetRequestId(v string) *GetUserMFAAuthenticationSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserMFAAuthenticationSettingsResponseBody) SetUserMFAAuthenticationSettings(v string) *GetUserMFAAuthenticationSettingsResponseBody {
	s.UserMFAAuthenticationSettings = &v
	return s
}

type GetUserMFAAuthenticationSettingsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserMFAAuthenticationSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserMFAAuthenticationSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserMFAAuthenticationSettingsResponse) GoString() string {
	return s.String()
}

func (s *GetUserMFAAuthenticationSettingsResponse) SetHeaders(v map[string]*string) *GetUserMFAAuthenticationSettingsResponse {
	s.Headers = v
	return s
}

func (s *GetUserMFAAuthenticationSettingsResponse) SetStatusCode(v int32) *GetUserMFAAuthenticationSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserMFAAuthenticationSettingsResponse) SetBody(v *GetUserMFAAuthenticationSettingsResponseBody) *GetUserMFAAuthenticationSettingsResponse {
	s.Body = v
	return s
}

type GetUserProvisioningRequest struct {
	DirectoryId        *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	UserProvisioningId *string `json:"UserProvisioningId,omitempty" xml:"UserProvisioningId,omitempty"`
}

func (s GetUserProvisioningRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserProvisioningRequest) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningRequest) SetDirectoryId(v string) *GetUserProvisioningRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetUserProvisioningRequest) SetUserProvisioningId(v string) *GetUserProvisioningRequest {
	s.UserProvisioningId = &v
	return s
}

type GetUserProvisioningResponseBody struct {
	RequestId        *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserProvisioning *GetUserProvisioningResponseBodyUserProvisioning `json:"UserProvisioning,omitempty" xml:"UserProvisioning,omitempty" type:"Struct"`
}

func (s GetUserProvisioningResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserProvisioningResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningResponseBody) SetRequestId(v string) *GetUserProvisioningResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserProvisioningResponseBody) SetUserProvisioning(v *GetUserProvisioningResponseBodyUserProvisioning) *GetUserProvisioningResponseBody {
	s.UserProvisioning = v
	return s
}

type GetUserProvisioningResponseBodyUserProvisioning struct {
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeletionStrategy    *string `json:"DeletionStrategy,omitempty" xml:"DeletionStrategy,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DirectoryId         *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	DuplicationStrategy *string `json:"DuplicationStrategy,omitempty" xml:"DuplicationStrategy,omitempty"`
	OwnerPk             *string `json:"OwnerPk,omitempty" xml:"OwnerPk,omitempty"`
	PrincipalId         *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	PrincipalName       *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	PrincipalType       *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetId            *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetName          *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	TargetPath          *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	TargetType          *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	UpdateTime          *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserProvisioningId  *string `json:"UserProvisioningId,omitempty" xml:"UserProvisioningId,omitempty"`
}

func (s GetUserProvisioningResponseBodyUserProvisioning) String() string {
	return tea.Prettify(s)
}

func (s GetUserProvisioningResponseBodyUserProvisioning) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetCreateTime(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.CreateTime = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetDeletionStrategy(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.DeletionStrategy = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetDescription(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.Description = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetDirectoryId(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.DirectoryId = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetDuplicationStrategy(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.DuplicationStrategy = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetOwnerPk(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.OwnerPk = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetPrincipalId(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.PrincipalId = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetPrincipalName(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.PrincipalName = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetPrincipalType(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.PrincipalType = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetStatus(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.Status = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetTargetId(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.TargetId = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetTargetName(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.TargetName = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetTargetPath(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.TargetPath = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetTargetType(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.TargetType = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetUpdateTime(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.UpdateTime = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetUserProvisioningId(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.UserProvisioningId = &v
	return s
}

type GetUserProvisioningResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserProvisioningResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserProvisioningResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserProvisioningResponse) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningResponse) SetHeaders(v map[string]*string) *GetUserProvisioningResponse {
	s.Headers = v
	return s
}

func (s *GetUserProvisioningResponse) SetStatusCode(v int32) *GetUserProvisioningResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserProvisioningResponse) SetBody(v *GetUserProvisioningResponseBody) *GetUserProvisioningResponse {
	s.Body = v
	return s
}

type GetUserProvisioningConfigurationRequest struct {
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s GetUserProvisioningConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserProvisioningConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningConfigurationRequest) SetDirectoryId(v string) *GetUserProvisioningConfigurationRequest {
	s.DirectoryId = &v
	return s
}

type GetUserProvisioningConfigurationResponseBody struct {
	RequestId                     *string                                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserProvisioningConfiguration *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration `json:"UserProvisioningConfiguration,omitempty" xml:"UserProvisioningConfiguration,omitempty" type:"Struct"`
}

func (s GetUserProvisioningConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserProvisioningConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningConfigurationResponseBody) SetRequestId(v string) *GetUserProvisioningConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserProvisioningConfigurationResponseBody) SetUserProvisioningConfiguration(v *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) *GetUserProvisioningConfigurationResponseBody {
	s.UserProvisioningConfiguration = v
	return s
}

type GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration struct {
	CreateTime         *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DefaultLandingPage *string `json:"DefaultLandingPage,omitempty" xml:"DefaultLandingPage,omitempty"`
	DirectoryId        *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	SessionDuration    *int32  `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	UpdateTime         *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) String() string {
	return tea.Prettify(s)
}

func (s GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) SetCreateTime(v string) *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration {
	s.CreateTime = &v
	return s
}

func (s *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) SetDefaultLandingPage(v string) *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration {
	s.DefaultLandingPage = &v
	return s
}

func (s *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) SetDirectoryId(v string) *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration {
	s.DirectoryId = &v
	return s
}

func (s *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) SetSessionDuration(v int32) *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration {
	s.SessionDuration = &v
	return s
}

func (s *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) SetUpdateTime(v string) *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration {
	s.UpdateTime = &v
	return s
}

type GetUserProvisioningConfigurationResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserProvisioningConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserProvisioningConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserProvisioningConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningConfigurationResponse) SetHeaders(v map[string]*string) *GetUserProvisioningConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetUserProvisioningConfigurationResponse) SetStatusCode(v int32) *GetUserProvisioningConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserProvisioningConfigurationResponse) SetBody(v *GetUserProvisioningConfigurationResponseBody) *GetUserProvisioningConfigurationResponse {
	s.Body = v
	return s
}

type GetUserProvisioningEventRequest struct {
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EventId     *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
}

func (s GetUserProvisioningEventRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserProvisioningEventRequest) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningEventRequest) SetDirectoryId(v string) *GetUserProvisioningEventRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetUserProvisioningEventRequest) SetEventId(v string) *GetUserProvisioningEventRequest {
	s.EventId = &v
	return s
}

type GetUserProvisioningEventResponseBody struct {
	RequestId             *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserProvisioningEvent *GetUserProvisioningEventResponseBodyUserProvisioningEvent `json:"UserProvisioningEvent,omitempty" xml:"UserProvisioningEvent,omitempty" type:"Struct"`
}

func (s GetUserProvisioningEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserProvisioningEventResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningEventResponseBody) SetRequestId(v string) *GetUserProvisioningEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserProvisioningEventResponseBody) SetUserProvisioningEvent(v *GetUserProvisioningEventResponseBodyUserProvisioningEvent) *GetUserProvisioningEventResponseBody {
	s.UserProvisioningEvent = v
	return s
}

type GetUserProvisioningEventResponseBodyUserProvisioningEvent struct {
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeletionStrategy    *string `json:"DeletionStrategy,omitempty" xml:"DeletionStrategy,omitempty"`
	DirectoryId         *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	DuplicationStrategy *string `json:"DuplicationStrategy,omitempty" xml:"DuplicationStrategy,omitempty"`
	ErrorCount          *int64  `json:"ErrorCount,omitempty" xml:"ErrorCount,omitempty"`
	ErrorInfo           *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	EventId             *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	LatestAsyncTime     *string `json:"LatestAsyncTime,omitempty" xml:"LatestAsyncTime,omitempty"`
	PrincipalId         *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	PrincipalName       *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	PrincipalType       *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	SourceType          *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	TargetId            *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetName          *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	TargetPath          *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	TargetType          *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	UpdateTime          *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserProvisioningId  *string `json:"UserProvisioningId,omitempty" xml:"UserProvisioningId,omitempty"`
}

func (s GetUserProvisioningEventResponseBodyUserProvisioningEvent) String() string {
	return tea.Prettify(s)
}

func (s GetUserProvisioningEventResponseBodyUserProvisioningEvent) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetCreateTime(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.CreateTime = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetDeletionStrategy(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.DeletionStrategy = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetDirectoryId(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.DirectoryId = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetDuplicationStrategy(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.DuplicationStrategy = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetErrorCount(v int64) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.ErrorCount = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetErrorInfo(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.ErrorInfo = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetEventId(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.EventId = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetLatestAsyncTime(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.LatestAsyncTime = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetPrincipalId(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.PrincipalId = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetPrincipalName(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.PrincipalName = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetPrincipalType(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.PrincipalType = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetSourceType(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.SourceType = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetTargetId(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.TargetId = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetTargetName(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.TargetName = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetTargetPath(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.TargetPath = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetTargetType(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.TargetType = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetUpdateTime(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.UpdateTime = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetUserProvisioningId(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.UserProvisioningId = &v
	return s
}

type GetUserProvisioningEventResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserProvisioningEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserProvisioningEventResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserProvisioningEventResponse) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningEventResponse) SetHeaders(v map[string]*string) *GetUserProvisioningEventResponse {
	s.Headers = v
	return s
}

func (s *GetUserProvisioningEventResponse) SetStatusCode(v int32) *GetUserProvisioningEventResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserProvisioningEventResponse) SetBody(v *GetUserProvisioningEventResponseBody) *GetUserProvisioningEventResponse {
	s.Body = v
	return s
}

type GetUserProvisioningRdAccountStatisticsRequest struct {
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	RdMemberId  *string `json:"RdMemberId,omitempty" xml:"RdMemberId,omitempty"`
}

func (s GetUserProvisioningRdAccountStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserProvisioningRdAccountStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningRdAccountStatisticsRequest) SetDirectoryId(v string) *GetUserProvisioningRdAccountStatisticsRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetUserProvisioningRdAccountStatisticsRequest) SetRdMemberId(v string) *GetUserProvisioningRdAccountStatisticsRequest {
	s.RdMemberId = &v
	return s
}

type GetUserProvisioningRdAccountStatisticsResponseBody struct {
	RequestId                  *string                                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserProvisioningStatistics *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics `json:"UserProvisioningStatistics,omitempty" xml:"UserProvisioningStatistics,omitempty" type:"Struct"`
}

func (s GetUserProvisioningRdAccountStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserProvisioningRdAccountStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningRdAccountStatisticsResponseBody) SetRequestId(v string) *GetUserProvisioningRdAccountStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserProvisioningRdAccountStatisticsResponseBody) SetUserProvisioningStatistics(v *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics) *GetUserProvisioningRdAccountStatisticsResponseBody {
	s.UserProvisioningStatistics = v
	return s
}

type GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics struct {
	DirectoryId      *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EntityId         *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	FailedEventCount *int64  `json:"FailedEventCount,omitempty" xml:"FailedEventCount,omitempty"`
	LatestAsyncTime  *string `json:"LatestAsyncTime,omitempty" xml:"LatestAsyncTime,omitempty"`
	OwnerPk          *string `json:"OwnerPk,omitempty" xml:"OwnerPk,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics) String() string {
	return tea.Prettify(s)
}

func (s GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics) SetDirectoryId(v string) *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics {
	s.DirectoryId = &v
	return s
}

func (s *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics) SetEntityId(v string) *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics {
	s.EntityId = &v
	return s
}

func (s *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics) SetFailedEventCount(v int64) *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics {
	s.FailedEventCount = &v
	return s
}

func (s *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics) SetLatestAsyncTime(v string) *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics {
	s.LatestAsyncTime = &v
	return s
}

func (s *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics) SetOwnerPk(v string) *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics {
	s.OwnerPk = &v
	return s
}

func (s *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics) SetType(v string) *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics {
	s.Type = &v
	return s
}

type GetUserProvisioningRdAccountStatisticsResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserProvisioningRdAccountStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserProvisioningRdAccountStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserProvisioningRdAccountStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningRdAccountStatisticsResponse) SetHeaders(v map[string]*string) *GetUserProvisioningRdAccountStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetUserProvisioningRdAccountStatisticsResponse) SetStatusCode(v int32) *GetUserProvisioningRdAccountStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserProvisioningRdAccountStatisticsResponse) SetBody(v *GetUserProvisioningRdAccountStatisticsResponseBody) *GetUserProvisioningRdAccountStatisticsResponse {
	s.Body = v
	return s
}

type GetUserProvisioningStatisticsRequest struct {
	DirectoryId        *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	UserProvisioningId *string `json:"UserProvisioningId,omitempty" xml:"UserProvisioningId,omitempty"`
}

func (s GetUserProvisioningStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserProvisioningStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningStatisticsRequest) SetDirectoryId(v string) *GetUserProvisioningStatisticsRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetUserProvisioningStatisticsRequest) SetUserProvisioningId(v string) *GetUserProvisioningStatisticsRequest {
	s.UserProvisioningId = &v
	return s
}

type GetUserProvisioningStatisticsResponseBody struct {
	RequestId                  *string                                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserProvisioningStatistics *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics `json:"UserProvisioningStatistics,omitempty" xml:"UserProvisioningStatistics,omitempty" type:"Struct"`
}

func (s GetUserProvisioningStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserProvisioningStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningStatisticsResponseBody) SetRequestId(v string) *GetUserProvisioningStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserProvisioningStatisticsResponseBody) SetUserProvisioningStatistics(v *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics) *GetUserProvisioningStatisticsResponseBody {
	s.UserProvisioningStatistics = v
	return s
}

type GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics struct {
	DirectoryId      *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EntityId         *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	FailedEventCount *int64  `json:"FailedEventCount,omitempty" xml:"FailedEventCount,omitempty"`
	LatestAsyncTime  *string `json:"LatestAsyncTime,omitempty" xml:"LatestAsyncTime,omitempty"`
	OwnerPk          *string `json:"OwnerPk,omitempty" xml:"OwnerPk,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics) String() string {
	return tea.Prettify(s)
}

func (s GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics) SetDirectoryId(v string) *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics {
	s.DirectoryId = &v
	return s
}

func (s *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics) SetEntityId(v string) *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics {
	s.EntityId = &v
	return s
}

func (s *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics) SetFailedEventCount(v int64) *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics {
	s.FailedEventCount = &v
	return s
}

func (s *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics) SetLatestAsyncTime(v string) *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics {
	s.LatestAsyncTime = &v
	return s
}

func (s *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics) SetOwnerPk(v string) *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics {
	s.OwnerPk = &v
	return s
}

func (s *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics) SetType(v string) *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics {
	s.Type = &v
	return s
}

type GetUserProvisioningStatisticsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserProvisioningStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserProvisioningStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserProvisioningStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningStatisticsResponse) SetHeaders(v map[string]*string) *GetUserProvisioningStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetUserProvisioningStatisticsResponse) SetStatusCode(v int32) *GetUserProvisioningStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserProvisioningStatisticsResponse) SetBody(v *GetUserProvisioningStatisticsResponseBody) *GetUserProvisioningStatisticsResponse {
	s.Body = v
	return s
}

type ListAccessAssignmentsRequest struct {
	// The ID of the access configuration. The ID can be used to filter access permissions.
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The directory ID.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 20.
	//
	// Default value: 10.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If this is your first time to call this operation, you do not need to specify the `NextToken` parameter.
	//
	// When you call this operation for the first time, if the total number of entries to return exceeds the value of `MaxResults`, the entries are truncated. Only the entries that match the value of `MaxResults` are returned, and the excess entries are not returned. In this case, the value of the response parameter `IsTruncated` is `true`, and `NextToken` is returned. In the next call, you can use the value of `NextToken` and maintain the settings of the other request parameters to query the excess entries. You can repeat the call until the value of `IsTruncated` becomes `false`. This way, all entries are returned.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the CloudSSO identity. The ID can be used to filter access permissions.
	//
	// *   If you set `PrincipalType` to User, set `PrincipalId` to the ID of the CloudSSO user.
	// *   If you set `PrincipalType` to Group, set `PrincipalId` to the ID of the CloudSSO group.
	//
	// >  You can use the type to filter access permissions only if you specify both PrincipalId and `PrincipalType`.``
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The type of the CloudSSO identity. The type can be used to filter access permissions. Valid values:
	//
	// *   User
	// *   Group
	//
	// >  You can use the type to filter access permissions only if you specify both PrincipalId and `PrincipalType`.``
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The ID of the task object. The ID can be used to filter access permissions.
	//
	// >  You can use the type to filter access permissions only if you specify both `TargetId` and `TargetType`.
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the task object. The type can be used to filter access permissions.
	//
	// Set the value to RD-Account, which specifies the accounts in the resource directory.
	//
	// >  You can use the type to filter access permissions only if you specify both `TargetId` and `TargetType`.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListAccessAssignmentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccessAssignmentsRequest) GoString() string {
	return s.String()
}

func (s *ListAccessAssignmentsRequest) SetAccessConfigurationId(v string) *ListAccessAssignmentsRequest {
	s.AccessConfigurationId = &v
	return s
}

func (s *ListAccessAssignmentsRequest) SetDirectoryId(v string) *ListAccessAssignmentsRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListAccessAssignmentsRequest) SetMaxResults(v int32) *ListAccessAssignmentsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAccessAssignmentsRequest) SetNextToken(v string) *ListAccessAssignmentsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAccessAssignmentsRequest) SetPrincipalId(v string) *ListAccessAssignmentsRequest {
	s.PrincipalId = &v
	return s
}

func (s *ListAccessAssignmentsRequest) SetPrincipalType(v string) *ListAccessAssignmentsRequest {
	s.PrincipalType = &v
	return s
}

func (s *ListAccessAssignmentsRequest) SetTargetId(v string) *ListAccessAssignmentsRequest {
	s.TargetId = &v
	return s
}

func (s *ListAccessAssignmentsRequest) SetTargetType(v string) *ListAccessAssignmentsRequest {
	s.TargetType = &v
	return s
}

type ListAccessAssignmentsResponseBody struct {
	// The access permissions that are assigned.
	AccessAssignments []*ListAccessAssignmentsResponseBodyAccessAssignments `json:"AccessAssignments,omitempty" xml:"AccessAssignments,omitempty" type:"Repeated"`
	// Indicates whether the queried entries are truncated. Valid values:
	//
	// *   true
	// *   false
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The maximum number of entries returned per page.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// >  This parameter is returned only when the value of IsTruncated is `true`.``
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
}

func (s ListAccessAssignmentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAccessAssignmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccessAssignmentsResponseBody) SetAccessAssignments(v []*ListAccessAssignmentsResponseBodyAccessAssignments) *ListAccessAssignmentsResponseBody {
	s.AccessAssignments = v
	return s
}

func (s *ListAccessAssignmentsResponseBody) SetIsTruncated(v bool) *ListAccessAssignmentsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListAccessAssignmentsResponseBody) SetMaxResults(v int32) *ListAccessAssignmentsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAccessAssignmentsResponseBody) SetNextToken(v string) *ListAccessAssignmentsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAccessAssignmentsResponseBody) SetRequestId(v string) *ListAccessAssignmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccessAssignmentsResponseBody) SetTotalCounts(v int32) *ListAccessAssignmentsResponseBody {
	s.TotalCounts = &v
	return s
}

type ListAccessAssignmentsResponseBodyAccessAssignments struct {
	// The ID of the access configuration.
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The name of the access configuration.
	AccessConfigurationName *string `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	// The time when the access permissions were assigned.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the CloudSSO identity.
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The name of the CloudSSO identity.
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The type of the CloudSSO identity. Valid values:
	//
	// *   User
	// *   Group
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The ID of the task object.
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the task object.
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The path ID of the task object in the resource directory.
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// The path name of the task object in the resource directory.
	TargetPathName *string `json:"TargetPathName,omitempty" xml:"TargetPathName,omitempty"`
	// The type of the task object.
	//
	// The value is fixed as RD-Account, which indicates the accounts in the resource directory.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListAccessAssignmentsResponseBodyAccessAssignments) String() string {
	return tea.Prettify(s)
}

func (s ListAccessAssignmentsResponseBodyAccessAssignments) GoString() string {
	return s.String()
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) SetAccessConfigurationId(v string) *ListAccessAssignmentsResponseBodyAccessAssignments {
	s.AccessConfigurationId = &v
	return s
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) SetAccessConfigurationName(v string) *ListAccessAssignmentsResponseBodyAccessAssignments {
	s.AccessConfigurationName = &v
	return s
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) SetCreateTime(v string) *ListAccessAssignmentsResponseBodyAccessAssignments {
	s.CreateTime = &v
	return s
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) SetPrincipalId(v string) *ListAccessAssignmentsResponseBodyAccessAssignments {
	s.PrincipalId = &v
	return s
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) SetPrincipalName(v string) *ListAccessAssignmentsResponseBodyAccessAssignments {
	s.PrincipalName = &v
	return s
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) SetPrincipalType(v string) *ListAccessAssignmentsResponseBodyAccessAssignments {
	s.PrincipalType = &v
	return s
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) SetTargetId(v string) *ListAccessAssignmentsResponseBodyAccessAssignments {
	s.TargetId = &v
	return s
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) SetTargetName(v string) *ListAccessAssignmentsResponseBodyAccessAssignments {
	s.TargetName = &v
	return s
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) SetTargetPath(v string) *ListAccessAssignmentsResponseBodyAccessAssignments {
	s.TargetPath = &v
	return s
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) SetTargetPathName(v string) *ListAccessAssignmentsResponseBodyAccessAssignments {
	s.TargetPathName = &v
	return s
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) SetTargetType(v string) *ListAccessAssignmentsResponseBodyAccessAssignments {
	s.TargetType = &v
	return s
}

type ListAccessAssignmentsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccessAssignmentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccessAssignmentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccessAssignmentsResponse) GoString() string {
	return s.String()
}

func (s *ListAccessAssignmentsResponse) SetHeaders(v map[string]*string) *ListAccessAssignmentsResponse {
	s.Headers = v
	return s
}

func (s *ListAccessAssignmentsResponse) SetStatusCode(v int32) *ListAccessAssignmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccessAssignmentsResponse) SetBody(v *ListAccessAssignmentsResponseBody) *ListAccessAssignmentsResponse {
	s.Body = v
	return s
}

type ListAccessConfigurationProvisioningsRequest struct {
	// The ID of the access configuration. The ID can be used to filter access permissions.
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 20.
	//
	// Default value: 10.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. If this is your first time to call this operation, you do not need to specify the `NextToken` parameter.
	//
	// When you call this operation for the first time, if the total number of entries to return exceeds the value of `MaxResults`, the entries are truncated. Only the entries that match the value of `MaxResults` are returned, and the excess entries are not returned. In this case, the value of the response parameter `IsTruncated` is `true`, and `NextToken` is returned. In the next call, you can use the value of `NextToken` and maintain the settings of the other request parameters to query the excess entries. You can repeat the call until the value of `IsTruncated` becomes `false`. This way, all entries are returned.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The status of the access configuration. The value can be used to filter accounts. Valid values:
	//
	// *   Provisioned: The access configuration is provisioned.
	// *   ReprovisionRequired: The access configuration needs to be re-provisioned.
	// *   DeprovisionFailed: The access configuration failed to be provisioned.
	ProvisioningStatus *string `json:"ProvisioningStatus,omitempty" xml:"ProvisioningStatus,omitempty"`
	// The ID of the task object. The ID can be used to filter access permissions.
	//
	// >  You can use the type to filter access permissions only if you specify both `TargetId` and `TargetType`.
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the task object. The type can be used to filter access permissions.
	//
	// Set the value to RD-Account, which specifies the accounts in the resource directory.
	//
	// >  You can use the type to filter access permissions only if you specify both `TargetId` and `TargetType`.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListAccessConfigurationProvisioningsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccessConfigurationProvisioningsRequest) GoString() string {
	return s.String()
}

func (s *ListAccessConfigurationProvisioningsRequest) SetAccessConfigurationId(v string) *ListAccessConfigurationProvisioningsRequest {
	s.AccessConfigurationId = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsRequest) SetDirectoryId(v string) *ListAccessConfigurationProvisioningsRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsRequest) SetMaxResults(v int32) *ListAccessConfigurationProvisioningsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsRequest) SetNextToken(v string) *ListAccessConfigurationProvisioningsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsRequest) SetProvisioningStatus(v string) *ListAccessConfigurationProvisioningsRequest {
	s.ProvisioningStatus = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsRequest) SetTargetId(v string) *ListAccessConfigurationProvisioningsRequest {
	s.TargetId = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsRequest) SetTargetType(v string) *ListAccessConfigurationProvisioningsRequest {
	s.TargetType = &v
	return s
}

type ListAccessConfigurationProvisioningsResponseBody struct {
	// The accounts for which the access configuration is provisioned.
	AccessConfigurationProvisionings []*ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings `json:"AccessConfigurationProvisionings,omitempty" xml:"AccessConfigurationProvisionings,omitempty" type:"Repeated"`
	// Indicates whether the queried entries are truncated. Valid values:
	//
	// *   true
	// *   false
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The maximum number of entries returned per page.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// >  This parameter is returned only when the value of `IsTruncated` is `true`.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
}

func (s ListAccessConfigurationProvisioningsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAccessConfigurationProvisioningsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccessConfigurationProvisioningsResponseBody) SetAccessConfigurationProvisionings(v []*ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) *ListAccessConfigurationProvisioningsResponseBody {
	s.AccessConfigurationProvisionings = v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBody) SetIsTruncated(v bool) *ListAccessConfigurationProvisioningsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBody) SetMaxResults(v int32) *ListAccessConfigurationProvisioningsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBody) SetNextToken(v string) *ListAccessConfigurationProvisioningsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBody) SetRequestId(v string) *ListAccessConfigurationProvisioningsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBody) SetTotalCounts(v int32) *ListAccessConfigurationProvisioningsResponseBody {
	s.TotalCounts = &v
	return s
}

type ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings struct {
	// The ID of the access configuration.
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The name of the access configuration.
	AccessConfigurationName *string `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	// The first time when the access configuration was provisioned.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The name of the custom policy that is created for an account in the resource directory.
	RAMPolicyNames []*string `json:"RAMPolicyNames,omitempty" xml:"RAMPolicyNames,omitempty" type:"Repeated"`
	// The name of the RAM role that is created for an account in the resource directory.
	RAMRoleName *string `json:"RAMRoleName,omitempty" xml:"RAMRoleName,omitempty"`
	// The name of the Security Assertion Markup Language (SAML) identity provider (IdP) that is created within an account in the resource directory.
	SAMLProviderName *string `json:"SAMLProviderName,omitempty" xml:"SAMLProviderName,omitempty"`
	// The status of the access configuration. Valid values:
	//
	// *   Provisioned: The access configuration is provisioned.
	// *   ReprovisionRequired: The access configuration needs to be re-provisioned.
	// *   DeprovisionFailed: The access configuration failed to be provisioned.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the task object.
	//
	// If the value of TargetType is `RD-Account`, the value of this parameter is the UID of an account in the resource directory.
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the task object.
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The path ID of the task object in the resource directory.
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// The path name of the task object in the resource directory.
	TargetPathName *string `json:"TargetPathName,omitempty" xml:"TargetPathName,omitempty"`
	// The type of the task object.
	//
	// Set the value to RD-Account, which specifies the accounts in the resource directory.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The last time when the access configuration was provisioned.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) String() string {
	return tea.Prettify(s)
}

func (s ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) GoString() string {
	return s.String()
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) SetAccessConfigurationId(v string) *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings {
	s.AccessConfigurationId = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) SetAccessConfigurationName(v string) *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings {
	s.AccessConfigurationName = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) SetCreateTime(v string) *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings {
	s.CreateTime = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) SetRAMPolicyNames(v []*string) *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings {
	s.RAMPolicyNames = v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) SetRAMRoleName(v string) *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings {
	s.RAMRoleName = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) SetSAMLProviderName(v string) *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings {
	s.SAMLProviderName = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) SetStatus(v string) *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings {
	s.Status = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) SetTargetId(v string) *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings {
	s.TargetId = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) SetTargetName(v string) *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings {
	s.TargetName = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) SetTargetPath(v string) *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings {
	s.TargetPath = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) SetTargetPathName(v string) *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings {
	s.TargetPathName = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) SetTargetType(v string) *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings {
	s.TargetType = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) SetUpdateTime(v string) *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings {
	s.UpdateTime = &v
	return s
}

type ListAccessConfigurationProvisioningsResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccessConfigurationProvisioningsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccessConfigurationProvisioningsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccessConfigurationProvisioningsResponse) GoString() string {
	return s.String()
}

func (s *ListAccessConfigurationProvisioningsResponse) SetHeaders(v map[string]*string) *ListAccessConfigurationProvisioningsResponse {
	s.Headers = v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponse) SetStatusCode(v int32) *ListAccessConfigurationProvisioningsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponse) SetBody(v *ListAccessConfigurationProvisioningsResponseBody) *ListAccessConfigurationProvisioningsResponse {
	s.Body = v
	return s
}

type ListAccessConfigurationsRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The filter condition.
	//
	// Specify the value in the \<Attribute> \<Operator> \<Value> format. The value is not case sensitive. You can set \<Attribute> only to AccessConfigurationName and \<Operator> only to eq or sw. The value eq indicates Equals. The value sw indicates Starts With.
	//
	// For example, if you set Filter to AccessConfigurationName sw test, the operation queries all access configurations whose names start with test. If you set Filter to AccessConfigurationName eq TestAccessConfiguration, the operation queries the access configuration whose name is TestAccessConfiguration.
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token to return for the next page. If this is your first time to call this operation, you do not need to specify the `NextToken` parameter.
	//
	// When you call this operation for the first time, if the total number of entries to return exceeds the value of `MaxResults`, the entries are truncated. Only the entries that match the value of `MaxResults` are returned, and the excess entries are not returned. In this case, the value of the response parameter `IsTruncated` is `true`, and `NextToken` is returned. In the next call, you can use the value of `NextToken` and maintain the settings of the other request parameters to query the excess entries. You can repeat the call until the value of `IsTruncated` becomes `false`. This way, all entries are returned.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The status notification. The status notification can be used to filter access configurations.
	//
	// Set the value to ReprovisionRequired, which indicates that the operation queries all access configurations that need to be re-provisioned.
	StatusNotifications *string `json:"StatusNotifications,omitempty" xml:"StatusNotifications,omitempty"`
}

func (s ListAccessConfigurationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccessConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *ListAccessConfigurationsRequest) SetDirectoryId(v string) *ListAccessConfigurationsRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListAccessConfigurationsRequest) SetFilter(v string) *ListAccessConfigurationsRequest {
	s.Filter = &v
	return s
}

func (s *ListAccessConfigurationsRequest) SetMaxResults(v int32) *ListAccessConfigurationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAccessConfigurationsRequest) SetNextToken(v string) *ListAccessConfigurationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAccessConfigurationsRequest) SetStatusNotifications(v string) *ListAccessConfigurationsRequest {
	s.StatusNotifications = &v
	return s
}

type ListAccessConfigurationsResponseBody struct {
	// The access configurations.
	AccessConfigurations []*ListAccessConfigurationsResponseBodyAccessConfigurations `json:"AccessConfigurations,omitempty" xml:"AccessConfigurations,omitempty" type:"Repeated"`
	// Indicates whether the queried entries are truncated. Valid values:
	//
	// *   true: The queried entries are truncated.
	// *   false: The queried entries are not truncated.
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The number of entries returned per page.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is returned for the next page.
	//
	// >  This parameter is returned only when the `IsTruncated` parameter is set to `true`.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
}

func (s ListAccessConfigurationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAccessConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccessConfigurationsResponseBody) SetAccessConfigurations(v []*ListAccessConfigurationsResponseBodyAccessConfigurations) *ListAccessConfigurationsResponseBody {
	s.AccessConfigurations = v
	return s
}

func (s *ListAccessConfigurationsResponseBody) SetIsTruncated(v bool) *ListAccessConfigurationsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListAccessConfigurationsResponseBody) SetMaxResults(v int32) *ListAccessConfigurationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAccessConfigurationsResponseBody) SetNextToken(v string) *ListAccessConfigurationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAccessConfigurationsResponseBody) SetRequestId(v string) *ListAccessConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccessConfigurationsResponseBody) SetTotalCounts(v int32) *ListAccessConfigurationsResponseBody {
	s.TotalCounts = &v
	return s
}

type ListAccessConfigurationsResponseBodyAccessConfigurations struct {
	// The ID of the access configuration.
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The name of the access configuration.
	AccessConfigurationName *string `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	// The time when the access configuration was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the access configuration.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The initial web page that is displayed after a CloudSSO user accesses an account in your resource directory by using the access configuration.
	RelayState *string `json:"RelayState,omitempty" xml:"RelayState,omitempty"`
	// The duration of a session in which a CloudSSO user accesses an account in your resource directory by using the access configuration.
	//
	// Unit: seconds.
	SessionDuration *int32 `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	// The status notification.
	StatusNotifications []*string `json:"StatusNotifications,omitempty" xml:"StatusNotifications,omitempty" type:"Repeated"`
	// The time when the information about the access configuration was modified.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListAccessConfigurationsResponseBodyAccessConfigurations) String() string {
	return tea.Prettify(s)
}

func (s ListAccessConfigurationsResponseBodyAccessConfigurations) GoString() string {
	return s.String()
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurations) SetAccessConfigurationId(v string) *ListAccessConfigurationsResponseBodyAccessConfigurations {
	s.AccessConfigurationId = &v
	return s
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurations) SetAccessConfigurationName(v string) *ListAccessConfigurationsResponseBodyAccessConfigurations {
	s.AccessConfigurationName = &v
	return s
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurations) SetCreateTime(v string) *ListAccessConfigurationsResponseBodyAccessConfigurations {
	s.CreateTime = &v
	return s
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurations) SetDescription(v string) *ListAccessConfigurationsResponseBodyAccessConfigurations {
	s.Description = &v
	return s
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurations) SetRelayState(v string) *ListAccessConfigurationsResponseBodyAccessConfigurations {
	s.RelayState = &v
	return s
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurations) SetSessionDuration(v int32) *ListAccessConfigurationsResponseBodyAccessConfigurations {
	s.SessionDuration = &v
	return s
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurations) SetStatusNotifications(v []*string) *ListAccessConfigurationsResponseBodyAccessConfigurations {
	s.StatusNotifications = v
	return s
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurations) SetUpdateTime(v string) *ListAccessConfigurationsResponseBodyAccessConfigurations {
	s.UpdateTime = &v
	return s
}

type ListAccessConfigurationsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccessConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccessConfigurationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccessConfigurationsResponse) GoString() string {
	return s.String()
}

func (s *ListAccessConfigurationsResponse) SetHeaders(v map[string]*string) *ListAccessConfigurationsResponse {
	s.Headers = v
	return s
}

func (s *ListAccessConfigurationsResponse) SetStatusCode(v int32) *ListAccessConfigurationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccessConfigurationsResponse) SetBody(v *ListAccessConfigurationsResponseBody) *ListAccessConfigurationsResponse {
	s.Body = v
	return s
}

type ListDirectoriesResponseBody struct {
	// The directories.
	Directories []*ListDirectoriesResponseBodyDirectories `json:"Directories,omitempty" xml:"Directories,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of directories.
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
}

func (s ListDirectoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDirectoriesResponseBody) SetDirectories(v []*ListDirectoriesResponseBodyDirectories) *ListDirectoriesResponseBody {
	s.Directories = v
	return s
}

func (s *ListDirectoriesResponseBody) SetRequestId(v string) *ListDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDirectoriesResponseBody) SetTotalCounts(v int32) *ListDirectoriesResponseBody {
	s.TotalCounts = &v
	return s
}

type ListDirectoriesResponseBodyDirectories struct {
	// The time when the directory was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The name of the directory.
	DirectoryName *string `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	// The region ID of the directory.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The time when the directory was modified.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListDirectoriesResponseBodyDirectories) String() string {
	return tea.Prettify(s)
}

func (s ListDirectoriesResponseBodyDirectories) GoString() string {
	return s.String()
}

func (s *ListDirectoriesResponseBodyDirectories) SetCreateTime(v string) *ListDirectoriesResponseBodyDirectories {
	s.CreateTime = &v
	return s
}

func (s *ListDirectoriesResponseBodyDirectories) SetDirectoryId(v string) *ListDirectoriesResponseBodyDirectories {
	s.DirectoryId = &v
	return s
}

func (s *ListDirectoriesResponseBodyDirectories) SetDirectoryName(v string) *ListDirectoriesResponseBodyDirectories {
	s.DirectoryName = &v
	return s
}

func (s *ListDirectoriesResponseBodyDirectories) SetRegion(v string) *ListDirectoriesResponseBodyDirectories {
	s.Region = &v
	return s
}

func (s *ListDirectoriesResponseBodyDirectories) SetUpdateTime(v string) *ListDirectoriesResponseBodyDirectories {
	s.UpdateTime = &v
	return s
}

type ListDirectoriesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDirectoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *ListDirectoriesResponse) SetHeaders(v map[string]*string) *ListDirectoriesResponse {
	s.Headers = v
	return s
}

func (s *ListDirectoriesResponse) SetStatusCode(v int32) *ListDirectoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDirectoriesResponse) SetBody(v *ListDirectoriesResponseBody) *ListDirectoriesResponse {
	s.Body = v
	return s
}

type ListExternalSAMLIdPCertificatesRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s ListExternalSAMLIdPCertificatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExternalSAMLIdPCertificatesRequest) GoString() string {
	return s.String()
}

func (s *ListExternalSAMLIdPCertificatesRequest) SetDirectoryId(v string) *ListExternalSAMLIdPCertificatesRequest {
	s.DirectoryId = &v
	return s
}

type ListExternalSAMLIdPCertificatesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The SAML signing certificates.
	SAMLIdPCertificates []*ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates `json:"SAMLIdPCertificates,omitempty" xml:"SAMLIdPCertificates,omitempty" type:"Repeated"`
	// The total number of entries returned.
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
}

func (s ListExternalSAMLIdPCertificatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExternalSAMLIdPCertificatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListExternalSAMLIdPCertificatesResponseBody) SetRequestId(v string) *ListExternalSAMLIdPCertificatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExternalSAMLIdPCertificatesResponseBody) SetSAMLIdPCertificates(v []*ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) *ListExternalSAMLIdPCertificatesResponseBody {
	s.SAMLIdPCertificates = v
	return s
}

func (s *ListExternalSAMLIdPCertificatesResponseBody) SetTotalCounts(v int32) *ListExternalSAMLIdPCertificatesResponseBody {
	s.TotalCounts = &v
	return s
}

type ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates struct {
	// The ID of the certificate.
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The issuer of the certificate.
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The time when the certificate expires.
	NotAfter *string `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// The time when the certificate was created.
	NotBefore *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// The public key of the certificate. The value of this parameter is in the PEM format and is Base64-encoded.
	PublicKey *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	// The serial number of the certificate.
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The signature algorithm of the certificate.
	SignatureAlgorithm *string `json:"SignatureAlgorithm,omitempty" xml:"SignatureAlgorithm,omitempty"`
	// The subject of the certificate.
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// The version of the certificate.
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
	// The X.509 certificate in the PEM format.
	X509Certificate *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
}

func (s ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) String() string {
	return tea.Prettify(s)
}

func (s ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) GoString() string {
	return s.String()
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) SetCertificateId(v string) *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates {
	s.CertificateId = &v
	return s
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) SetIssuer(v string) *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates {
	s.Issuer = &v
	return s
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) SetNotAfter(v string) *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates {
	s.NotAfter = &v
	return s
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) SetNotBefore(v string) *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates {
	s.NotBefore = &v
	return s
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) SetPublicKey(v string) *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates {
	s.PublicKey = &v
	return s
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) SetSerialNumber(v string) *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates {
	s.SerialNumber = &v
	return s
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) SetSignatureAlgorithm(v string) *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates {
	s.SignatureAlgorithm = &v
	return s
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) SetSubject(v string) *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates {
	s.Subject = &v
	return s
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) SetVersion(v int32) *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates {
	s.Version = &v
	return s
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) SetX509Certificate(v string) *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates {
	s.X509Certificate = &v
	return s
}

type ListExternalSAMLIdPCertificatesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExternalSAMLIdPCertificatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExternalSAMLIdPCertificatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExternalSAMLIdPCertificatesResponse) GoString() string {
	return s.String()
}

func (s *ListExternalSAMLIdPCertificatesResponse) SetHeaders(v map[string]*string) *ListExternalSAMLIdPCertificatesResponse {
	s.Headers = v
	return s
}

func (s *ListExternalSAMLIdPCertificatesResponse) SetStatusCode(v int32) *ListExternalSAMLIdPCertificatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExternalSAMLIdPCertificatesResponse) SetBody(v *ListExternalSAMLIdPCertificatesResponseBody) *ListExternalSAMLIdPCertificatesResponse {
	s.Body = v
	return s
}

type ListGroupMembersRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the group.
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token to return for the next page. If this is your first time to call this operation, you do not need to specify `NextToken` .
	//
	// When you call this operation for the first time, if the total number of entries to return exceeds the value of `MaxResults`, the entries are truncated. Only the entries that match the value of `MaxResults` are returned, and the excess entries are not returned. In this case, the value of the response parameter `IsTruncated` is `true`, and `NextToken` is returned. In the next call, you can use the value of `NextToken` and maintain the settings of the other request parameters to query the excess entries. You can repeat the call until the value of `IsTruncated` becomes `false`. This way, all entries are returned.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListGroupMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupMembersRequest) GoString() string {
	return s.String()
}

func (s *ListGroupMembersRequest) SetDirectoryId(v string) *ListGroupMembersRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListGroupMembersRequest) SetGroupId(v string) *ListGroupMembersRequest {
	s.GroupId = &v
	return s
}

func (s *ListGroupMembersRequest) SetMaxResults(v int32) *ListGroupMembersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListGroupMembersRequest) SetNextToken(v string) *ListGroupMembersRequest {
	s.NextToken = &v
	return s
}

type ListGroupMembersResponseBody struct {
	// The users in the group.
	GroupMembers []*ListGroupMembersResponseBodyGroupMembers `json:"GroupMembers,omitempty" xml:"GroupMembers,omitempty" type:"Repeated"`
	// Indicates whether the queried entries are truncated. Valid values:
	//
	// *   true: The queried entries are truncated.
	// *   false: The queried entries are not truncated.
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The number of entries returned per page.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is returned for the next page.
	//
	// >  This parameter is returned only when the value of `IsTruncated` is `true`.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
}

func (s ListGroupMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupMembersResponseBody) SetGroupMembers(v []*ListGroupMembersResponseBodyGroupMembers) *ListGroupMembersResponseBody {
	s.GroupMembers = v
	return s
}

func (s *ListGroupMembersResponseBody) SetIsTruncated(v bool) *ListGroupMembersResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListGroupMembersResponseBody) SetMaxResults(v int32) *ListGroupMembersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListGroupMembersResponseBody) SetNextToken(v string) *ListGroupMembersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListGroupMembersResponseBody) SetRequestId(v string) *ListGroupMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupMembersResponseBody) SetTotalCounts(v int32) *ListGroupMembersResponseBody {
	s.TotalCounts = &v
	return s
}

type ListGroupMembersResponseBodyGroupMembers struct {
	// The description of the user.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the user.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the user.
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The ID of the group.
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The time when the user was added to the user group.
	JoinTime *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// The type of the user. Valid values:
	//
	// *   Manual: The user is manually created.
	// *   Synchronized: The user is synchronized from an external identity provider (IdP).
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	// The status of the user. Valid values:
	//
	// *   Enabled: The logon of the user is enabled.
	// *   Disabled: The logon of the user is disabled.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The name of the user.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListGroupMembersResponseBodyGroupMembers) String() string {
	return tea.Prettify(s)
}

func (s ListGroupMembersResponseBodyGroupMembers) GoString() string {
	return s.String()
}

func (s *ListGroupMembersResponseBodyGroupMembers) SetDescription(v string) *ListGroupMembersResponseBodyGroupMembers {
	s.Description = &v
	return s
}

func (s *ListGroupMembersResponseBodyGroupMembers) SetDisplayName(v string) *ListGroupMembersResponseBodyGroupMembers {
	s.DisplayName = &v
	return s
}

func (s *ListGroupMembersResponseBodyGroupMembers) SetEmail(v string) *ListGroupMembersResponseBodyGroupMembers {
	s.Email = &v
	return s
}

func (s *ListGroupMembersResponseBodyGroupMembers) SetGroupId(v string) *ListGroupMembersResponseBodyGroupMembers {
	s.GroupId = &v
	return s
}

func (s *ListGroupMembersResponseBodyGroupMembers) SetJoinTime(v string) *ListGroupMembersResponseBodyGroupMembers {
	s.JoinTime = &v
	return s
}

func (s *ListGroupMembersResponseBodyGroupMembers) SetProvisionType(v string) *ListGroupMembersResponseBodyGroupMembers {
	s.ProvisionType = &v
	return s
}

func (s *ListGroupMembersResponseBodyGroupMembers) SetStatus(v string) *ListGroupMembersResponseBodyGroupMembers {
	s.Status = &v
	return s
}

func (s *ListGroupMembersResponseBodyGroupMembers) SetUserId(v string) *ListGroupMembersResponseBodyGroupMembers {
	s.UserId = &v
	return s
}

func (s *ListGroupMembersResponseBodyGroupMembers) SetUserName(v string) *ListGroupMembersResponseBodyGroupMembers {
	s.UserName = &v
	return s
}

type ListGroupMembersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupMembersResponse) GoString() string {
	return s.String()
}

func (s *ListGroupMembersResponse) SetHeaders(v map[string]*string) *ListGroupMembersResponse {
	s.Headers = v
	return s
}

func (s *ListGroupMembersResponse) SetStatusCode(v int32) *ListGroupMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupMembersResponse) SetBody(v *ListGroupMembersResponseBody) *ListGroupMembersResponse {
	s.Body = v
	return s
}

type ListGroupsRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The filter condition.
	//
	// Specify the value in the `<Attribute> <Operator> <Value>` format. The value is not case sensitive. You can set `<Attribute>` only to `GroupName` and `<Operator>` only to `eq` or `sw`. The value eq indicates Equals. The value sw indicates Starts With.
	//
	// For example, if you set Filter to GroupName sw test, the operation queries the groups whose names start with test. If you set Filter to GroupName eq testgroup, the operation queries the group whose name is testgroup.
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token to return for the next page. If this is your first time to call this operation, you do not need to specify `NextToken`.
	//
	// When you call this operation for the first time, if the total number of entries to return exceeds the value of `MaxResults`, the entries are truncated. Only the entries that match the value of `MaxResults` are returned, and the excess entries are not returned. In this case, the value of the response parameter `IsTruncated` is `true`, and `NextToken` is returned. In the next call, you can use the value of `NextToken` and maintain the settings of the other request parameters to query the excess entries. You can repeat the call until the value of `IsTruncated` becomes `false`. This way, all entries are returned.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The type of the group. The type can be used to filter groups. Valid values:
	//
	// *   Manual: The group is manually created.
	// *   Synchronized: The group is synchronized from an external identity provider (IdP).
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
}

func (s ListGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsRequest) SetDirectoryId(v string) *ListGroupsRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListGroupsRequest) SetFilter(v string) *ListGroupsRequest {
	s.Filter = &v
	return s
}

func (s *ListGroupsRequest) SetMaxResults(v int32) *ListGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListGroupsRequest) SetNextToken(v string) *ListGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListGroupsRequest) SetProvisionType(v string) *ListGroupsRequest {
	s.ProvisionType = &v
	return s
}

type ListGroupsResponseBody struct {
	// The groups.
	Groups []*ListGroupsResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// Indicates whether the queried entries are truncated. Valid values:
	//
	// *   true: The queried entries are truncated.
	// *   false: The queried entries are not truncated.
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The number of entries returned per page.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is returned for the next page.
	//
	// >  This parameter is returned only when the `IsTruncated` parameter is set to `true`.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
}

func (s ListGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBody) SetGroups(v []*ListGroupsResponseBodyGroups) *ListGroupsResponseBody {
	s.Groups = v
	return s
}

func (s *ListGroupsResponseBody) SetIsTruncated(v bool) *ListGroupsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListGroupsResponseBody) SetMaxResults(v int32) *ListGroupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListGroupsResponseBody) SetNextToken(v string) *ListGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListGroupsResponseBody) SetRequestId(v string) *ListGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupsResponseBody) SetTotalCounts(v int32) *ListGroupsResponseBody {
	s.TotalCounts = &v
	return s
}

type ListGroupsResponseBodyGroups struct {
	// The time when the group was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the group.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the group.
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the group. Valid values:
	//
	// *   Manual: The group is manually created.
	// *   Synchronized: The group is synchronized from an external IdP.
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	// The time when the information about the group was modified.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListGroupsResponseBodyGroups) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBodyGroups) SetCreateTime(v string) *ListGroupsResponseBodyGroups {
	s.CreateTime = &v
	return s
}

func (s *ListGroupsResponseBodyGroups) SetDescription(v string) *ListGroupsResponseBodyGroups {
	s.Description = &v
	return s
}

func (s *ListGroupsResponseBodyGroups) SetGroupId(v string) *ListGroupsResponseBodyGroups {
	s.GroupId = &v
	return s
}

func (s *ListGroupsResponseBodyGroups) SetGroupName(v string) *ListGroupsResponseBodyGroups {
	s.GroupName = &v
	return s
}

func (s *ListGroupsResponseBodyGroups) SetProvisionType(v string) *ListGroupsResponseBodyGroups {
	s.ProvisionType = &v
	return s
}

func (s *ListGroupsResponseBodyGroups) SetUpdateTime(v string) *ListGroupsResponseBodyGroups {
	s.UpdateTime = &v
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

type ListJoinedGroupsForUserRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token to return for the next page. If this is your first time to call this operation, you do not need to specify `NextToken` .
	//
	// When you call this operation for the first time, if the total number of entries to return exceeds the value of `MaxResults`, the entries are truncated. Only the entries that match the value of `MaxResults` are returned, and the excess entries are not returned. In this case, the value of the response parameter `IsTruncated` is `true`, and `NextToken` is returned. In the next call, you can use the value of `NextToken` and maintain the settings of the other request parameters to query the excess entries. You can repeat the call until the value of `IsTruncated` becomes `false`. This way, all entries are returned.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListJoinedGroupsForUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJoinedGroupsForUserRequest) GoString() string {
	return s.String()
}

func (s *ListJoinedGroupsForUserRequest) SetDirectoryId(v string) *ListJoinedGroupsForUserRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListJoinedGroupsForUserRequest) SetMaxResults(v int32) *ListJoinedGroupsForUserRequest {
	s.MaxResults = &v
	return s
}

func (s *ListJoinedGroupsForUserRequest) SetNextToken(v string) *ListJoinedGroupsForUserRequest {
	s.NextToken = &v
	return s
}

func (s *ListJoinedGroupsForUserRequest) SetUserId(v string) *ListJoinedGroupsForUserRequest {
	s.UserId = &v
	return s
}

type ListJoinedGroupsForUserResponseBody struct {
	// Indicates whether the queried entries are truncated. Valid values:
	//
	// *   true: The queried entries are truncated.
	// *   false: The queried entries are not truncated.
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The groups to which the user is added.
	JoinedGroups []*ListJoinedGroupsForUserResponseBodyJoinedGroups `json:"JoinedGroups,omitempty" xml:"JoinedGroups,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is returned for the next page.
	//
	// >  This parameter is returned only when the value of `IsTruncated` is `true`.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
}

func (s ListJoinedGroupsForUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJoinedGroupsForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListJoinedGroupsForUserResponseBody) SetIsTruncated(v bool) *ListJoinedGroupsForUserResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListJoinedGroupsForUserResponseBody) SetJoinedGroups(v []*ListJoinedGroupsForUserResponseBodyJoinedGroups) *ListJoinedGroupsForUserResponseBody {
	s.JoinedGroups = v
	return s
}

func (s *ListJoinedGroupsForUserResponseBody) SetMaxResults(v int32) *ListJoinedGroupsForUserResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListJoinedGroupsForUserResponseBody) SetNextToken(v string) *ListJoinedGroupsForUserResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListJoinedGroupsForUserResponseBody) SetRequestId(v string) *ListJoinedGroupsForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJoinedGroupsForUserResponseBody) SetTotalCounts(v int32) *ListJoinedGroupsForUserResponseBody {
	s.TotalCounts = &v
	return s
}

type ListJoinedGroupsForUserResponseBodyJoinedGroups struct {
	// The description of the group.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the group.
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The time when the user was added to the user group.
	JoinTime *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// The type of the group. Valid values:
	//
	// *   Manual: The group is manually created.
	// *   Synchronized: The user is synchronized from an external identity provider (IdP).
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListJoinedGroupsForUserResponseBodyJoinedGroups) String() string {
	return tea.Prettify(s)
}

func (s ListJoinedGroupsForUserResponseBodyJoinedGroups) GoString() string {
	return s.String()
}

func (s *ListJoinedGroupsForUserResponseBodyJoinedGroups) SetDescription(v string) *ListJoinedGroupsForUserResponseBodyJoinedGroups {
	s.Description = &v
	return s
}

func (s *ListJoinedGroupsForUserResponseBodyJoinedGroups) SetGroupId(v string) *ListJoinedGroupsForUserResponseBodyJoinedGroups {
	s.GroupId = &v
	return s
}

func (s *ListJoinedGroupsForUserResponseBodyJoinedGroups) SetGroupName(v string) *ListJoinedGroupsForUserResponseBodyJoinedGroups {
	s.GroupName = &v
	return s
}

func (s *ListJoinedGroupsForUserResponseBodyJoinedGroups) SetJoinTime(v string) *ListJoinedGroupsForUserResponseBodyJoinedGroups {
	s.JoinTime = &v
	return s
}

func (s *ListJoinedGroupsForUserResponseBodyJoinedGroups) SetProvisionType(v string) *ListJoinedGroupsForUserResponseBodyJoinedGroups {
	s.ProvisionType = &v
	return s
}

func (s *ListJoinedGroupsForUserResponseBodyJoinedGroups) SetUserId(v string) *ListJoinedGroupsForUserResponseBodyJoinedGroups {
	s.UserId = &v
	return s
}

type ListJoinedGroupsForUserResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJoinedGroupsForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJoinedGroupsForUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJoinedGroupsForUserResponse) GoString() string {
	return s.String()
}

func (s *ListJoinedGroupsForUserResponse) SetHeaders(v map[string]*string) *ListJoinedGroupsForUserResponse {
	s.Headers = v
	return s
}

func (s *ListJoinedGroupsForUserResponse) SetStatusCode(v int32) *ListJoinedGroupsForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJoinedGroupsForUserResponse) SetBody(v *ListJoinedGroupsForUserResponseBody) *ListJoinedGroupsForUserResponse {
	s.Body = v
	return s
}

type ListMFADevicesForUserRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListMFADevicesForUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMFADevicesForUserRequest) GoString() string {
	return s.String()
}

func (s *ListMFADevicesForUserRequest) SetDirectoryId(v string) *ListMFADevicesForUserRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListMFADevicesForUserRequest) SetUserId(v string) *ListMFADevicesForUserRequest {
	s.UserId = &v
	return s
}

type ListMFADevicesForUserResponseBody struct {
	// The MFA devices.
	MFADevices []*ListMFADevicesForUserResponseBodyMFADevices `json:"MFADevices,omitempty" xml:"MFADevices,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of MFA devices.
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
}

func (s ListMFADevicesForUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMFADevicesForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListMFADevicesForUserResponseBody) SetMFADevices(v []*ListMFADevicesForUserResponseBodyMFADevices) *ListMFADevicesForUserResponseBody {
	s.MFADevices = v
	return s
}

func (s *ListMFADevicesForUserResponseBody) SetRequestId(v string) *ListMFADevicesForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMFADevicesForUserResponseBody) SetTotalCounts(v int32) *ListMFADevicesForUserResponseBody {
	s.TotalCounts = &v
	return s
}

type ListMFADevicesForUserResponseBodyMFADevices struct {
	// The ID of the MFA device.
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// The name of the MFA device.
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The type of the MFA device. The value is fixed as TOTP, which indicates a virtual MFA device. Virtual MFA devices are based on the Time-based One-time Password (TOTP) algorithm.
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The time when the MFA device was enabled.
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListMFADevicesForUserResponseBodyMFADevices) String() string {
	return tea.Prettify(s)
}

func (s ListMFADevicesForUserResponseBodyMFADevices) GoString() string {
	return s.String()
}

func (s *ListMFADevicesForUserResponseBodyMFADevices) SetDeviceId(v string) *ListMFADevicesForUserResponseBodyMFADevices {
	s.DeviceId = &v
	return s
}

func (s *ListMFADevicesForUserResponseBodyMFADevices) SetDeviceName(v string) *ListMFADevicesForUserResponseBodyMFADevices {
	s.DeviceName = &v
	return s
}

func (s *ListMFADevicesForUserResponseBodyMFADevices) SetDeviceType(v string) *ListMFADevicesForUserResponseBodyMFADevices {
	s.DeviceType = &v
	return s
}

func (s *ListMFADevicesForUserResponseBodyMFADevices) SetEffectiveTime(v string) *ListMFADevicesForUserResponseBodyMFADevices {
	s.EffectiveTime = &v
	return s
}

func (s *ListMFADevicesForUserResponseBodyMFADevices) SetUserId(v string) *ListMFADevicesForUserResponseBodyMFADevices {
	s.UserId = &v
	return s
}

type ListMFADevicesForUserResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMFADevicesForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMFADevicesForUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMFADevicesForUserResponse) GoString() string {
	return s.String()
}

func (s *ListMFADevicesForUserResponse) SetHeaders(v map[string]*string) *ListMFADevicesForUserResponse {
	s.Headers = v
	return s
}

func (s *ListMFADevicesForUserResponse) SetStatusCode(v int32) *ListMFADevicesForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMFADevicesForUserResponse) SetBody(v *ListMFADevicesForUserResponseBody) *ListMFADevicesForUserResponse {
	s.Body = v
	return s
}

type ListPermissionPoliciesInAccessConfigurationRequest struct {
	// The ID of the access configuration.
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The type of the policy. The type can be used to filter policies. Valid values:
	//
	// *   System: system policy
	// *   Inline: inline policy
	//
	// If you do not specify this parameter, all types of policies are queried.
	PermissionPolicyType *string `json:"PermissionPolicyType,omitempty" xml:"PermissionPolicyType,omitempty"`
}

func (s ListPermissionPoliciesInAccessConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionPoliciesInAccessConfigurationRequest) GoString() string {
	return s.String()
}

func (s *ListPermissionPoliciesInAccessConfigurationRequest) SetAccessConfigurationId(v string) *ListPermissionPoliciesInAccessConfigurationRequest {
	s.AccessConfigurationId = &v
	return s
}

func (s *ListPermissionPoliciesInAccessConfigurationRequest) SetDirectoryId(v string) *ListPermissionPoliciesInAccessConfigurationRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListPermissionPoliciesInAccessConfigurationRequest) SetPermissionPolicyType(v string) *ListPermissionPoliciesInAccessConfigurationRequest {
	s.PermissionPolicyType = &v
	return s
}

type ListPermissionPoliciesInAccessConfigurationResponseBody struct {
	// The policies.
	PermissionPolicies []*ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies `json:"PermissionPolicies,omitempty" xml:"PermissionPolicies,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of policies.
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
}

func (s ListPermissionPoliciesInAccessConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionPoliciesInAccessConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *ListPermissionPoliciesInAccessConfigurationResponseBody) SetPermissionPolicies(v []*ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies) *ListPermissionPoliciesInAccessConfigurationResponseBody {
	s.PermissionPolicies = v
	return s
}

func (s *ListPermissionPoliciesInAccessConfigurationResponseBody) SetRequestId(v string) *ListPermissionPoliciesInAccessConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPermissionPoliciesInAccessConfigurationResponseBody) SetTotalCounts(v int32) *ListPermissionPoliciesInAccessConfigurationResponseBody {
	s.TotalCounts = &v
	return s
}

type ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies struct {
	// The time when the policy was created for the access configuration.
	AddTime *string `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	// The configurations of the inline policy.
	//
	// >  This parameter is returned only when the value of the PermissionPolicyType parameter is Inline.
	PermissionPolicyDocument *string `json:"PermissionPolicyDocument,omitempty" xml:"PermissionPolicyDocument,omitempty"`
	// The name of the policy.
	PermissionPolicyName *string `json:"PermissionPolicyName,omitempty" xml:"PermissionPolicyName,omitempty"`
	// The type of the policy.
	PermissionPolicyType *string `json:"PermissionPolicyType,omitempty" xml:"PermissionPolicyType,omitempty"`
}

func (s ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies) GoString() string {
	return s.String()
}

func (s *ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies) SetAddTime(v string) *ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies {
	s.AddTime = &v
	return s
}

func (s *ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies) SetPermissionPolicyDocument(v string) *ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies {
	s.PermissionPolicyDocument = &v
	return s
}

func (s *ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies) SetPermissionPolicyName(v string) *ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies {
	s.PermissionPolicyName = &v
	return s
}

func (s *ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies) SetPermissionPolicyType(v string) *ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies {
	s.PermissionPolicyType = &v
	return s
}

type ListPermissionPoliciesInAccessConfigurationResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPermissionPoliciesInAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPermissionPoliciesInAccessConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionPoliciesInAccessConfigurationResponse) GoString() string {
	return s.String()
}

func (s *ListPermissionPoliciesInAccessConfigurationResponse) SetHeaders(v map[string]*string) *ListPermissionPoliciesInAccessConfigurationResponse {
	s.Headers = v
	return s
}

func (s *ListPermissionPoliciesInAccessConfigurationResponse) SetStatusCode(v int32) *ListPermissionPoliciesInAccessConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPermissionPoliciesInAccessConfigurationResponse) SetBody(v *ListPermissionPoliciesInAccessConfigurationResponseBody) *ListPermissionPoliciesInAccessConfigurationResponse {
	s.Body = v
	return s
}

type ListSCIMServerCredentialsRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s ListSCIMServerCredentialsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSCIMServerCredentialsRequest) GoString() string {
	return s.String()
}

func (s *ListSCIMServerCredentialsRequest) SetDirectoryId(v string) *ListSCIMServerCredentialsRequest {
	s.DirectoryId = &v
	return s
}

type ListSCIMServerCredentialsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The SCIM credentials.
	SCIMServerCredentials []*ListSCIMServerCredentialsResponseBodySCIMServerCredentials `json:"SCIMServerCredentials,omitempty" xml:"SCIMServerCredentials,omitempty" type:"Repeated"`
	// The total number of entries returned.
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
}

func (s ListSCIMServerCredentialsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSCIMServerCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSCIMServerCredentialsResponseBody) SetRequestId(v string) *ListSCIMServerCredentialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSCIMServerCredentialsResponseBody) SetSCIMServerCredentials(v []*ListSCIMServerCredentialsResponseBodySCIMServerCredentials) *ListSCIMServerCredentialsResponseBody {
	s.SCIMServerCredentials = v
	return s
}

func (s *ListSCIMServerCredentialsResponseBody) SetTotalCounts(v int32) *ListSCIMServerCredentialsResponseBody {
	s.TotalCounts = &v
	return s
}

type ListSCIMServerCredentialsResponseBodySCIMServerCredentials struct {
	// The time when the SCIM credential was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the SCIM credential.
	CredentialId *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
	// The type of the SCIM credential.
	CredentialType *string `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The time when the SCIM credential expires.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The status of the SCIM credential. Valid values:
	//
	// *   Enabled: The SCIM credential is enabled.
	// *   Disabled: The SCIM credential is disabled.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListSCIMServerCredentialsResponseBodySCIMServerCredentials) String() string {
	return tea.Prettify(s)
}

func (s ListSCIMServerCredentialsResponseBodySCIMServerCredentials) GoString() string {
	return s.String()
}

func (s *ListSCIMServerCredentialsResponseBodySCIMServerCredentials) SetCreateTime(v string) *ListSCIMServerCredentialsResponseBodySCIMServerCredentials {
	s.CreateTime = &v
	return s
}

func (s *ListSCIMServerCredentialsResponseBodySCIMServerCredentials) SetCredentialId(v string) *ListSCIMServerCredentialsResponseBodySCIMServerCredentials {
	s.CredentialId = &v
	return s
}

func (s *ListSCIMServerCredentialsResponseBodySCIMServerCredentials) SetCredentialType(v string) *ListSCIMServerCredentialsResponseBodySCIMServerCredentials {
	s.CredentialType = &v
	return s
}

func (s *ListSCIMServerCredentialsResponseBodySCIMServerCredentials) SetDirectoryId(v string) *ListSCIMServerCredentialsResponseBodySCIMServerCredentials {
	s.DirectoryId = &v
	return s
}

func (s *ListSCIMServerCredentialsResponseBodySCIMServerCredentials) SetExpireTime(v string) *ListSCIMServerCredentialsResponseBodySCIMServerCredentials {
	s.ExpireTime = &v
	return s
}

func (s *ListSCIMServerCredentialsResponseBodySCIMServerCredentials) SetStatus(v string) *ListSCIMServerCredentialsResponseBodySCIMServerCredentials {
	s.Status = &v
	return s
}

type ListSCIMServerCredentialsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSCIMServerCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSCIMServerCredentialsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSCIMServerCredentialsResponse) GoString() string {
	return s.String()
}

func (s *ListSCIMServerCredentialsResponse) SetHeaders(v map[string]*string) *ListSCIMServerCredentialsResponse {
	s.Headers = v
	return s
}

func (s *ListSCIMServerCredentialsResponse) SetStatusCode(v int32) *ListSCIMServerCredentialsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSCIMServerCredentialsResponse) SetBody(v *ListSCIMServerCredentialsResponseBody) *ListSCIMServerCredentialsResponse {
	s.Body = v
	return s
}

type ListTasksRequest struct {
	// The ID of the access configuration. The ID can be used to filter access permissions.
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The filter condition.
	//
	// The condition is not case-sensitive. The condition must be in the StartTime ge YYYY-MM-DDTHH:mm:SSZ format. You must set YYYY-MM-DDTHH:mm:SSZ to a value that is no more than 7 days from the current time. ge indicates Greater Than or Equals.
	//
	// For example, if you set the Filter parameter to StartTime ge 2021-03-15T01:12:23Z, the operation queries the tasks from 2021-03-15T01:12:23 GMT.
	//
	// >  If you do not specify this parameter, the operation queries the tasks within the previous 24 hours by default.
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 20.
	//
	// Default value: 10.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If this is your first time to call this operation, you do not need to specify the `NextToken` parameter.
	//
	// When you call this operation for the first time, if the total number of entries to return exceeds the value of `MaxResults`, the entries are truncated. Only the entries that match the value of `MaxResults` are returned, and the excess entries are not returned. In this case, the value of the response parameter `IsTruncated` is `true`, and `NextToken` is returned. In the next call, you can use the value of `NextToken` and maintain the settings of the other request parameters to query the excess entries. You can repeat the call until the value of `IsTruncated` becomes `false`. This way, all entries are returned.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the CloudSSO identity. The ID can be used to filter access permissions.
	//
	// *   If you set `PrincipalType` to `User`, set `PrincipalId` to the ID of the CloudSSO user.
	// *   If you set `PrincipalType` to `Group`, set `PrincipalId` to the ID of the CloudSSO group.
	//
	// >  You can use the type to filter access permissions only if you specify both `PrincipalId` and `PrincipalType`.
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The type of the CloudSSO identity. The type can be used to filter access permissions. Valid values:
	//
	// *   User
	// *   Group
	//
	// >  You can use the type to filter access permissions only if you specify both `PrincipalId` and `PrincipalType`.
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The ID of the task. The ID can be used to filter tasks. Valid values:
	//
	// *   InProgress: The task is running.
	// *   Success: The task is successful.
	// *   Failed: The task failed.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the task object. The ID can be used to filter access permissions.
	//
	// >  You can use the type to filter access permissions only if you specify both `TargetId` and `TargetType`.
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the task object. The type can be used to filter access permissions.
	//
	// Set the value to RD-Account, which specifies the accounts in the resource directory.
	//
	// >  You can use the type to filter access permissions only if you specify both `TargetId` and `TargetType`.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The type of the task. The type can be used to filter tasks. Valid values:
	//
	// *   ProvisionAccessConfiguration: An access configuration is provisioned.
	// *   DeprovisionAccessConfiguration: An access configuration is de-provisioned.
	// *   CreateAccessAssignment: Access permissions on an account in the resource directory are assigned.
	// *   DeleteAccessAssignment: Access permissions on an account in the resource directory are removed.
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ListTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTasksRequest) SetAccessConfigurationId(v string) *ListTasksRequest {
	s.AccessConfigurationId = &v
	return s
}

func (s *ListTasksRequest) SetDirectoryId(v string) *ListTasksRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListTasksRequest) SetFilter(v string) *ListTasksRequest {
	s.Filter = &v
	return s
}

func (s *ListTasksRequest) SetMaxResults(v int32) *ListTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTasksRequest) SetNextToken(v string) *ListTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListTasksRequest) SetPrincipalId(v string) *ListTasksRequest {
	s.PrincipalId = &v
	return s
}

func (s *ListTasksRequest) SetPrincipalType(v string) *ListTasksRequest {
	s.PrincipalType = &v
	return s
}

func (s *ListTasksRequest) SetStatus(v string) *ListTasksRequest {
	s.Status = &v
	return s
}

func (s *ListTasksRequest) SetTargetId(v string) *ListTasksRequest {
	s.TargetId = &v
	return s
}

func (s *ListTasksRequest) SetTargetType(v string) *ListTasksRequest {
	s.TargetType = &v
	return s
}

func (s *ListTasksRequest) SetTaskType(v string) *ListTasksRequest {
	s.TaskType = &v
	return s
}

type ListTasksResponseBody struct {
	// Indicates whether the queried entries are truncated. Valid values:
	//
	// *   true
	// *   false
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The maximum number of entries returned per page.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// >  This parameter is returned only when the value of `IsTruncated` is `true`.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tasks.
	Tasks []*ListTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// The total number of entries returned.
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
}

func (s ListTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBody) SetIsTruncated(v bool) *ListTasksResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListTasksResponseBody) SetMaxResults(v int32) *ListTasksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTasksResponseBody) SetNextToken(v string) *ListTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTasksResponseBody) SetRequestId(v string) *ListTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTasksResponseBody) SetTasks(v []*ListTasksResponseBodyTasks) *ListTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *ListTasksResponseBody) SetTotalCounts(v int32) *ListTasksResponseBody {
	s.TotalCounts = &v
	return s
}

type ListTasksResponseBodyTasks struct {
	// The ID of the access configuration.
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The name of the access configuration.
	AccessConfigurationName *string `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	// The end time of the task.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The cause of the task failure.
	//
	// >  This parameter is returned only when the value of `Status` is `Failed`.
	FailureReason *string `json:"FailureReason,omitempty" xml:"FailureReason,omitempty"`
	// The ID of the CloudSSO identity.
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The name of the CloudSSO identity.
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The type of the CloudSSO identity. Valid values:
	//
	// *   User
	// *   Group
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The start time of the task.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task status. Valid values:
	//
	// *   InProgress: The task is running.
	// *   Success: The task is successful.
	// *   Failed: The task failed.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the task object.
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the task object.
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The path ID of the task object in the resource directory.
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// The path name of the task object in the resource directory.
	TargetPathName *string `json:"TargetPathName,omitempty" xml:"TargetPathName,omitempty"`
	// The type of the task object.
	//
	// The value is fixed as RD-Account, which indicates the accounts in the resource directory.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The ID of the job.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type. Valid values:
	//
	// *   ProvisionAccessConfiguration: An access configuration is provisioned.
	// *   DeprovisionAccessConfiguration: An access configuration is de-provisioned.
	// *   CreateAccessAssignment: Access permissions on an account in the resource directory are assigned.
	// *   DeleteAccessAssignment: Access permissions on an account in the resource directory are removed.
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ListTasksResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyTasks) SetAccessConfigurationId(v string) *ListTasksResponseBodyTasks {
	s.AccessConfigurationId = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetAccessConfigurationName(v string) *ListTasksResponseBodyTasks {
	s.AccessConfigurationName = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetEndTime(v string) *ListTasksResponseBodyTasks {
	s.EndTime = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetFailureReason(v string) *ListTasksResponseBodyTasks {
	s.FailureReason = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetPrincipalId(v string) *ListTasksResponseBodyTasks {
	s.PrincipalId = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetPrincipalName(v string) *ListTasksResponseBodyTasks {
	s.PrincipalName = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetPrincipalType(v string) *ListTasksResponseBodyTasks {
	s.PrincipalType = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetStartTime(v string) *ListTasksResponseBodyTasks {
	s.StartTime = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetStatus(v string) *ListTasksResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetTargetId(v string) *ListTasksResponseBodyTasks {
	s.TargetId = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetTargetName(v string) *ListTasksResponseBodyTasks {
	s.TargetName = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetTargetPath(v string) *ListTasksResponseBodyTasks {
	s.TargetPath = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetTargetPathName(v string) *ListTasksResponseBodyTasks {
	s.TargetPathName = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetTargetType(v string) *ListTasksResponseBodyTasks {
	s.TargetType = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetTaskId(v string) *ListTasksResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetTaskType(v string) *ListTasksResponseBodyTasks {
	s.TaskType = &v
	return s
}

type ListTasksResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponse) GoString() string {
	return s.String()
}

func (s *ListTasksResponse) SetHeaders(v map[string]*string) *ListTasksResponse {
	s.Headers = v
	return s
}

func (s *ListTasksResponse) SetStatusCode(v int32) *ListTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTasksResponse) SetBody(v *ListTasksResponseBody) *ListTasksResponse {
	s.Body = v
	return s
}

type ListUserProvisioningEventsRequest struct {
	DirectoryId        *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	MaxResults         *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken          *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	UserProvisioningId *string `json:"UserProvisioningId,omitempty" xml:"UserProvisioningId,omitempty"`
}

func (s ListUserProvisioningEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserProvisioningEventsRequest) GoString() string {
	return s.String()
}

func (s *ListUserProvisioningEventsRequest) SetDirectoryId(v string) *ListUserProvisioningEventsRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListUserProvisioningEventsRequest) SetMaxResults(v int32) *ListUserProvisioningEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUserProvisioningEventsRequest) SetNextToken(v string) *ListUserProvisioningEventsRequest {
	s.NextToken = &v
	return s
}

func (s *ListUserProvisioningEventsRequest) SetUserProvisioningId(v string) *ListUserProvisioningEventsRequest {
	s.UserProvisioningId = &v
	return s
}

type ListUserProvisioningEventsResponseBody struct {
	IsTruncated            *bool                                                           `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	MaxResults             *int32                                                          `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken              *string                                                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId              *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCounts            *int32                                                          `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
	UserProvisioningEvents []*ListUserProvisioningEventsResponseBodyUserProvisioningEvents `json:"UserProvisioningEvents,omitempty" xml:"UserProvisioningEvents,omitempty" type:"Repeated"`
}

func (s ListUserProvisioningEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserProvisioningEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserProvisioningEventsResponseBody) SetIsTruncated(v bool) *ListUserProvisioningEventsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBody) SetMaxResults(v int32) *ListUserProvisioningEventsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBody) SetNextToken(v string) *ListUserProvisioningEventsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBody) SetRequestId(v string) *ListUserProvisioningEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBody) SetTotalCounts(v int32) *ListUserProvisioningEventsResponseBody {
	s.TotalCounts = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBody) SetUserProvisioningEvents(v []*ListUserProvisioningEventsResponseBodyUserProvisioningEvents) *ListUserProvisioningEventsResponseBody {
	s.UserProvisioningEvents = v
	return s
}

type ListUserProvisioningEventsResponseBodyUserProvisioningEvents struct {
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeletionStrategy    *string `json:"DeletionStrategy,omitempty" xml:"DeletionStrategy,omitempty"`
	DirectoryId         *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	DuplicationStrategy *string `json:"DuplicationStrategy,omitempty" xml:"DuplicationStrategy,omitempty"`
	ErrorCount          *int64  `json:"ErrorCount,omitempty" xml:"ErrorCount,omitempty"`
	ErrorInfo           *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	EventId             *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	LatestAsyncTime     *string `json:"LatestAsyncTime,omitempty" xml:"LatestAsyncTime,omitempty"`
	PrincipalId         *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	PrincipalName       *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	PrincipalType       *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	SourceType          *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	TargetId            *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetName          *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	TargetPath          *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	TargetType          *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	UpdateTime          *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserProvisioningId  *string `json:"UserProvisioningId,omitempty" xml:"UserProvisioningId,omitempty"`
}

func (s ListUserProvisioningEventsResponseBodyUserProvisioningEvents) String() string {
	return tea.Prettify(s)
}

func (s ListUserProvisioningEventsResponseBodyUserProvisioningEvents) GoString() string {
	return s.String()
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetCreateTime(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.CreateTime = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetDeletionStrategy(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.DeletionStrategy = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetDirectoryId(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.DirectoryId = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetDuplicationStrategy(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.DuplicationStrategy = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetErrorCount(v int64) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.ErrorCount = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetErrorInfo(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.ErrorInfo = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetEventId(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.EventId = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetLatestAsyncTime(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.LatestAsyncTime = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetPrincipalId(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.PrincipalId = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetPrincipalName(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.PrincipalName = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetPrincipalType(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.PrincipalType = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetSourceType(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.SourceType = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetTargetId(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.TargetId = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetTargetName(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.TargetName = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetTargetPath(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.TargetPath = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetTargetType(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.TargetType = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetUpdateTime(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.UpdateTime = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetUserProvisioningId(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.UserProvisioningId = &v
	return s
}

type ListUserProvisioningEventsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserProvisioningEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserProvisioningEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserProvisioningEventsResponse) GoString() string {
	return s.String()
}

func (s *ListUserProvisioningEventsResponse) SetHeaders(v map[string]*string) *ListUserProvisioningEventsResponse {
	s.Headers = v
	return s
}

func (s *ListUserProvisioningEventsResponse) SetStatusCode(v int32) *ListUserProvisioningEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserProvisioningEventsResponse) SetBody(v *ListUserProvisioningEventsResponseBody) *ListUserProvisioningEventsResponse {
	s.Body = v
	return s
}

type ListUserProvisioningsRequest struct {
	DirectoryId   *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	MaxResults    *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PrincipalId   *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	TargetId      *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetType    *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListUserProvisioningsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserProvisioningsRequest) GoString() string {
	return s.String()
}

func (s *ListUserProvisioningsRequest) SetDirectoryId(v string) *ListUserProvisioningsRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListUserProvisioningsRequest) SetMaxResults(v int32) *ListUserProvisioningsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUserProvisioningsRequest) SetNextToken(v string) *ListUserProvisioningsRequest {
	s.NextToken = &v
	return s
}

func (s *ListUserProvisioningsRequest) SetPrincipalId(v string) *ListUserProvisioningsRequest {
	s.PrincipalId = &v
	return s
}

func (s *ListUserProvisioningsRequest) SetPrincipalType(v string) *ListUserProvisioningsRequest {
	s.PrincipalType = &v
	return s
}

func (s *ListUserProvisioningsRequest) SetTargetId(v string) *ListUserProvisioningsRequest {
	s.TargetId = &v
	return s
}

func (s *ListUserProvisioningsRequest) SetTargetType(v string) *ListUserProvisioningsRequest {
	s.TargetType = &v
	return s
}

type ListUserProvisioningsResponseBody struct {
	IsTruncated       *bool                                                 `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	MaxResults        *int32                                                `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken         *string                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId         *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCounts       *int32                                                `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
	UserProvisionings []*ListUserProvisioningsResponseBodyUserProvisionings `json:"UserProvisionings,omitempty" xml:"UserProvisionings,omitempty" type:"Repeated"`
}

func (s ListUserProvisioningsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserProvisioningsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserProvisioningsResponseBody) SetIsTruncated(v bool) *ListUserProvisioningsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListUserProvisioningsResponseBody) SetMaxResults(v int32) *ListUserProvisioningsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListUserProvisioningsResponseBody) SetNextToken(v string) *ListUserProvisioningsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUserProvisioningsResponseBody) SetRequestId(v string) *ListUserProvisioningsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserProvisioningsResponseBody) SetTotalCounts(v int32) *ListUserProvisioningsResponseBody {
	s.TotalCounts = &v
	return s
}

func (s *ListUserProvisioningsResponseBody) SetUserProvisionings(v []*ListUserProvisioningsResponseBodyUserProvisionings) *ListUserProvisioningsResponseBody {
	s.UserProvisionings = v
	return s
}

type ListUserProvisioningsResponseBodyUserProvisionings struct {
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeletionStrategy    *string `json:"DeletionStrategy,omitempty" xml:"DeletionStrategy,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DirectoryId         *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	DuplicationStrategy *string `json:"DuplicationStrategy,omitempty" xml:"DuplicationStrategy,omitempty"`
	OwnerPk             *string `json:"OwnerPk,omitempty" xml:"OwnerPk,omitempty"`
	PrincipalId         *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	PrincipalName       *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	PrincipalType       *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetId            *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetName          *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	TargetPath          *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	TargetType          *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	UpdateTime          *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserProvisioningId  *string `json:"UserProvisioningId,omitempty" xml:"UserProvisioningId,omitempty"`
}

func (s ListUserProvisioningsResponseBodyUserProvisionings) String() string {
	return tea.Prettify(s)
}

func (s ListUserProvisioningsResponseBodyUserProvisionings) GoString() string {
	return s.String()
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetCreateTime(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.CreateTime = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetDeletionStrategy(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.DeletionStrategy = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetDescription(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.Description = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetDirectoryId(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.DirectoryId = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetDuplicationStrategy(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.DuplicationStrategy = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetOwnerPk(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.OwnerPk = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetPrincipalId(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.PrincipalId = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetPrincipalName(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.PrincipalName = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetPrincipalType(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.PrincipalType = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetStatus(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.Status = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetTargetId(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.TargetId = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetTargetName(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.TargetName = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetTargetPath(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.TargetPath = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetTargetType(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.TargetType = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetUpdateTime(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.UpdateTime = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetUserProvisioningId(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.UserProvisioningId = &v
	return s
}

type ListUserProvisioningsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserProvisioningsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserProvisioningsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserProvisioningsResponse) GoString() string {
	return s.String()
}

func (s *ListUserProvisioningsResponse) SetHeaders(v map[string]*string) *ListUserProvisioningsResponse {
	s.Headers = v
	return s
}

func (s *ListUserProvisioningsResponse) SetStatusCode(v int32) *ListUserProvisioningsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserProvisioningsResponse) SetBody(v *ListUserProvisioningsResponseBody) *ListUserProvisioningsResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The filter condition.
	//
	// Specify the value in the `<Attribute> <Operator> <Value>` format. The value is not case-sensitive. You can set `<Attribute>` only to `UserName` and `Operator` only to `eq` or `sw`. The value eq indicates Equals, and the value sw indicates Starts With.
	//
	// For example, if you set Filter to UserName sw test, the operation queries the users whose names start with test. If you set Filter to UserName eq testuser, the operation queries the user whose name is `testuser`.
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token to return for the next page. If this is your first time to call this operation, you do not need to specify `NextToken`.
	//
	// When you call this operation for the first time, if the total number of entries to return exceeds the value of `MaxResults`, the entries are truncated. Only the entries that match the value of `MaxResults` are returned, and the excess entries are not returned. In this case, the value of the response parameter `IsTruncated` is `true`, and `NextToken` is returned. In the next call, you can use the value of `NextToken` and maintain the settings of the other request parameters to query the excess entries. You can repeat the call until the value of `IsTruncated` becomes `false`. This way, all entries are returned.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The type of the user. The type can be used to filter users. Valid values:
	//
	// *   Manual: The user is manually created.
	// *   Synchronized: The user is synchronized from an external IdP.
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	// The status of the user. The status can be used to filter users. Valid values:
	//
	// *   Enabled: The logon of the user is enabled.
	// *   Disabled: The logon of the user is disabled.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetDirectoryId(v string) *ListUsersRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListUsersRequest) SetFilter(v string) *ListUsersRequest {
	s.Filter = &v
	return s
}

func (s *ListUsersRequest) SetMaxResults(v int32) *ListUsersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUsersRequest) SetNextToken(v string) *ListUsersRequest {
	s.NextToken = &v
	return s
}

func (s *ListUsersRequest) SetProvisionType(v string) *ListUsersRequest {
	s.ProvisionType = &v
	return s
}

func (s *ListUsersRequest) SetStatus(v string) *ListUsersRequest {
	s.Status = &v
	return s
}

type ListUsersResponseBody struct {
	// Indicates whether the queried entries are truncated. Valid values:
	//
	// *   true: The queried entries are truncated.
	// *   false: The queried entries are not truncated.
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The number of entries returned per page.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is returned for the next page.
	//
	// >  This parameter is returned only when the value of `IsTruncated` is `true`.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
	// The users.
	Users []*ListUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
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

func (s *ListUsersResponseBody) SetMaxResults(v int32) *ListUsersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListUsersResponseBody) SetNextToken(v string) *ListUsersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetTotalCounts(v int32) *ListUsersResponseBody {
	s.TotalCounts = &v
	return s
}

func (s *ListUsersResponseBody) SetUsers(v []*ListUsersResponseBodyUsers) *ListUsersResponseBody {
	s.Users = v
	return s
}

type ListUsersResponseBodyUsers struct {
	// The time when the user was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the user.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the user.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the user.
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The first name of the user.
	FirstName *string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	// The last name of the user.
	LastName *string `json:"LastName,omitempty" xml:"LastName,omitempty"`
	// The type of the user. Valid values:
	//
	// *   Manual: The user is manually created.
	// *   Synchronized: The user is synchronized from an external IdP.
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	// The status of the user. Valid values:
	//
	// *   Enabled: The logon of the user is enabled.
	// *   Disabled: The logon of the user is disabled.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the information about the user was modified.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The name of the user.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListUsersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsers) SetCreateTime(v string) *ListUsersResponseBodyUsers {
	s.CreateTime = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetDescription(v string) *ListUsersResponseBodyUsers {
	s.Description = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetDisplayName(v string) *ListUsersResponseBodyUsers {
	s.DisplayName = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetEmail(v string) *ListUsersResponseBodyUsers {
	s.Email = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetFirstName(v string) *ListUsersResponseBodyUsers {
	s.FirstName = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetLastName(v string) *ListUsersResponseBodyUsers {
	s.LastName = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetProvisionType(v string) *ListUsersResponseBodyUsers {
	s.ProvisionType = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetStatus(v string) *ListUsersResponseBodyUsers {
	s.Status = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUpdateTime(v string) *ListUsersResponseBodyUsers {
	s.UpdateTime = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUserId(v string) *ListUsersResponseBodyUsers {
	s.UserId = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUserName(v string) *ListUsersResponseBodyUsers {
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

type ProvisionAccessConfigurationRequest struct {
	// The ID of the access configuration.
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The directory ID.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the task object.
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the task object. Set the value to RD-Account, which specifies the accounts in the resource directory.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ProvisionAccessConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s ProvisionAccessConfigurationRequest) GoString() string {
	return s.String()
}

func (s *ProvisionAccessConfigurationRequest) SetAccessConfigurationId(v string) *ProvisionAccessConfigurationRequest {
	s.AccessConfigurationId = &v
	return s
}

func (s *ProvisionAccessConfigurationRequest) SetDirectoryId(v string) *ProvisionAccessConfigurationRequest {
	s.DirectoryId = &v
	return s
}

func (s *ProvisionAccessConfigurationRequest) SetTargetId(v string) *ProvisionAccessConfigurationRequest {
	s.TargetId = &v
	return s
}

func (s *ProvisionAccessConfigurationRequest) SetTargetType(v string) *ProvisionAccessConfigurationRequest {
	s.TargetType = &v
	return s
}

type ProvisionAccessConfigurationResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task information.
	Tasks []*ProvisionAccessConfigurationResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s ProvisionAccessConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ProvisionAccessConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *ProvisionAccessConfigurationResponseBody) SetRequestId(v string) *ProvisionAccessConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ProvisionAccessConfigurationResponseBody) SetTasks(v []*ProvisionAccessConfigurationResponseBodyTasks) *ProvisionAccessConfigurationResponseBody {
	s.Tasks = v
	return s
}

type ProvisionAccessConfigurationResponseBodyTasks struct {
	// The ID of the access configuration.
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The name of the access configuration.
	AccessConfigurationName *string `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	// The task status. Valid values:
	//
	// *   InProgress: The task is running.
	// *   Success: The task is successful.
	// *   Failed: The task failed.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the task object.
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the task object.
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The path ID of the task object in the resource directory.
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// The path name of the task object in the resource directory.
	TargetPathName *string `json:"TargetPathName,omitempty" xml:"TargetPathName,omitempty"`
	// The type of the task object. The value is fixed as RD-Account, which indicates the accounts in the resource directory.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The task ID.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type. The value is fixed as ProvisionAccessConfiguration, which indicates that an access configuration is provisioned.
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ProvisionAccessConfigurationResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s ProvisionAccessConfigurationResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) SetAccessConfigurationId(v string) *ProvisionAccessConfigurationResponseBodyTasks {
	s.AccessConfigurationId = &v
	return s
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) SetAccessConfigurationName(v string) *ProvisionAccessConfigurationResponseBodyTasks {
	s.AccessConfigurationName = &v
	return s
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) SetStatus(v string) *ProvisionAccessConfigurationResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) SetTargetId(v string) *ProvisionAccessConfigurationResponseBodyTasks {
	s.TargetId = &v
	return s
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) SetTargetName(v string) *ProvisionAccessConfigurationResponseBodyTasks {
	s.TargetName = &v
	return s
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) SetTargetPath(v string) *ProvisionAccessConfigurationResponseBodyTasks {
	s.TargetPath = &v
	return s
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) SetTargetPathName(v string) *ProvisionAccessConfigurationResponseBodyTasks {
	s.TargetPathName = &v
	return s
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) SetTargetType(v string) *ProvisionAccessConfigurationResponseBodyTasks {
	s.TargetType = &v
	return s
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) SetTaskId(v string) *ProvisionAccessConfigurationResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) SetTaskType(v string) *ProvisionAccessConfigurationResponseBodyTasks {
	s.TaskType = &v
	return s
}

type ProvisionAccessConfigurationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProvisionAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ProvisionAccessConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s ProvisionAccessConfigurationResponse) GoString() string {
	return s.String()
}

func (s *ProvisionAccessConfigurationResponse) SetHeaders(v map[string]*string) *ProvisionAccessConfigurationResponse {
	s.Headers = v
	return s
}

func (s *ProvisionAccessConfigurationResponse) SetStatusCode(v int32) *ProvisionAccessConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *ProvisionAccessConfigurationResponse) SetBody(v *ProvisionAccessConfigurationResponseBody) *ProvisionAccessConfigurationResponse {
	s.Body = v
	return s
}

type RemoveExternalSAMLIdPCertificateRequest struct {
	// The ID of the certificate.
	//
	// You can call the [ListExternalSAMLIdPCertificates](~~341629~~) operation to query the IDs of certificates.
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s RemoveExternalSAMLIdPCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveExternalSAMLIdPCertificateRequest) GoString() string {
	return s.String()
}

func (s *RemoveExternalSAMLIdPCertificateRequest) SetCertificateId(v string) *RemoveExternalSAMLIdPCertificateRequest {
	s.CertificateId = &v
	return s
}

func (s *RemoveExternalSAMLIdPCertificateRequest) SetDirectoryId(v string) *RemoveExternalSAMLIdPCertificateRequest {
	s.DirectoryId = &v
	return s
}

type RemoveExternalSAMLIdPCertificateResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveExternalSAMLIdPCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveExternalSAMLIdPCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveExternalSAMLIdPCertificateResponseBody) SetRequestId(v string) *RemoveExternalSAMLIdPCertificateResponseBody {
	s.RequestId = &v
	return s
}

type RemoveExternalSAMLIdPCertificateResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveExternalSAMLIdPCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveExternalSAMLIdPCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveExternalSAMLIdPCertificateResponse) GoString() string {
	return s.String()
}

func (s *RemoveExternalSAMLIdPCertificateResponse) SetHeaders(v map[string]*string) *RemoveExternalSAMLIdPCertificateResponse {
	s.Headers = v
	return s
}

func (s *RemoveExternalSAMLIdPCertificateResponse) SetStatusCode(v int32) *RemoveExternalSAMLIdPCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveExternalSAMLIdPCertificateResponse) SetBody(v *RemoveExternalSAMLIdPCertificateResponseBody) *RemoveExternalSAMLIdPCertificateResponse {
	s.Body = v
	return s
}

type RemovePermissionPolicyFromAccessConfigurationRequest struct {
	// The ID of the access configuration.
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The name of the policy.
	PermissionPolicyName *string `json:"PermissionPolicyName,omitempty" xml:"PermissionPolicyName,omitempty"`
	// The type of the policy. Valid values:
	//
	// *   System: system policy
	// *   Inline: inline policy
	PermissionPolicyType *string `json:"PermissionPolicyType,omitempty" xml:"PermissionPolicyType,omitempty"`
}

func (s RemovePermissionPolicyFromAccessConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s RemovePermissionPolicyFromAccessConfigurationRequest) GoString() string {
	return s.String()
}

func (s *RemovePermissionPolicyFromAccessConfigurationRequest) SetAccessConfigurationId(v string) *RemovePermissionPolicyFromAccessConfigurationRequest {
	s.AccessConfigurationId = &v
	return s
}

func (s *RemovePermissionPolicyFromAccessConfigurationRequest) SetDirectoryId(v string) *RemovePermissionPolicyFromAccessConfigurationRequest {
	s.DirectoryId = &v
	return s
}

func (s *RemovePermissionPolicyFromAccessConfigurationRequest) SetPermissionPolicyName(v string) *RemovePermissionPolicyFromAccessConfigurationRequest {
	s.PermissionPolicyName = &v
	return s
}

func (s *RemovePermissionPolicyFromAccessConfigurationRequest) SetPermissionPolicyType(v string) *RemovePermissionPolicyFromAccessConfigurationRequest {
	s.PermissionPolicyType = &v
	return s
}

type RemovePermissionPolicyFromAccessConfigurationResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemovePermissionPolicyFromAccessConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemovePermissionPolicyFromAccessConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *RemovePermissionPolicyFromAccessConfigurationResponseBody) SetRequestId(v string) *RemovePermissionPolicyFromAccessConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type RemovePermissionPolicyFromAccessConfigurationResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemovePermissionPolicyFromAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemovePermissionPolicyFromAccessConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s RemovePermissionPolicyFromAccessConfigurationResponse) GoString() string {
	return s.String()
}

func (s *RemovePermissionPolicyFromAccessConfigurationResponse) SetHeaders(v map[string]*string) *RemovePermissionPolicyFromAccessConfigurationResponse {
	s.Headers = v
	return s
}

func (s *RemovePermissionPolicyFromAccessConfigurationResponse) SetStatusCode(v int32) *RemovePermissionPolicyFromAccessConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *RemovePermissionPolicyFromAccessConfigurationResponse) SetBody(v *RemovePermissionPolicyFromAccessConfigurationResponseBody) *RemovePermissionPolicyFromAccessConfigurationResponse {
	s.Body = v
	return s
}

type RemoveUserFromGroupRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the group.
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RemoveUserFromGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserFromGroupRequest) SetDirectoryId(v string) *RemoveUserFromGroupRequest {
	s.DirectoryId = &v
	return s
}

func (s *RemoveUserFromGroupRequest) SetGroupId(v string) *RemoveUserFromGroupRequest {
	s.GroupId = &v
	return s
}

func (s *RemoveUserFromGroupRequest) SetUserId(v string) *RemoveUserFromGroupRequest {
	s.UserId = &v
	return s
}

type RemoveUserFromGroupResponseBody struct {
	// The ID of the request.
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

type ResetUserPasswordRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// Specifies whether to enable the system to automatically generate a new password. Valid values:
	//
	// *   True: The new password is automatically generated by the system.
	// *   False: The new password must be manually specified. This is the default value.
	GenerateRandomPassword *bool `json:"GenerateRandomPassword,omitempty" xml:"GenerateRandomPassword,omitempty"`
	// The new password.
	//
	// The password must contain the following types of characters: uppercase letters, lowercase letters, digits, and special characters.
	//
	// The password must be 8 to 32 characters in length.
	//
	// >  If you set `GenerateRandomPassword` to `False`, you must specify `Password` .
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether password reset is required upon the next logon. Valid values:
	//
	// *   True: Password reset is required upon the next logon.
	// *   False: Password reset is not required upon the next logon. This is the default value.
	RequirePasswordResetForNextLogin *bool `json:"RequirePasswordResetForNextLogin,omitempty" xml:"RequirePasswordResetForNextLogin,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ResetUserPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordRequest) SetDirectoryId(v string) *ResetUserPasswordRequest {
	s.DirectoryId = &v
	return s
}

func (s *ResetUserPasswordRequest) SetGenerateRandomPassword(v bool) *ResetUserPasswordRequest {
	s.GenerateRandomPassword = &v
	return s
}

func (s *ResetUserPasswordRequest) SetPassword(v string) *ResetUserPasswordRequest {
	s.Password = &v
	return s
}

func (s *ResetUserPasswordRequest) SetRequirePasswordResetForNextLogin(v bool) *ResetUserPasswordRequest {
	s.RequirePasswordResetForNextLogin = &v
	return s
}

func (s *ResetUserPasswordRequest) SetUserId(v string) *ResetUserPasswordRequest {
	s.UserId = &v
	return s
}

type ResetUserPasswordResponseBody struct {
	// The new password.
	//
	// >  This parameter is returned only when the new password is automatically generated by the system.
	NewPassword *string `json:"NewPassword,omitempty" xml:"NewPassword,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetUserPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetUserPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordResponseBody) SetNewPassword(v string) *ResetUserPasswordResponseBody {
	s.NewPassword = &v
	return s
}

func (s *ResetUserPasswordResponseBody) SetRequestId(v string) *ResetUserPasswordResponseBody {
	s.RequestId = &v
	return s
}

type ResetUserPasswordResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetUserPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetUserPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetUserPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordResponse) SetHeaders(v map[string]*string) *ResetUserPasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetUserPasswordResponse) SetStatusCode(v int32) *ResetUserPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetUserPasswordResponse) SetBody(v *ResetUserPasswordResponseBody) *ResetUserPasswordResponse {
	s.Body = v
	return s
}

type RetryUserProvisioningEventRequest struct {
	DirectoryId         *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	DuplicationStrategy *string `json:"DuplicationStrategy,omitempty" xml:"DuplicationStrategy,omitempty"`
	EventId             *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
}

func (s RetryUserProvisioningEventRequest) String() string {
	return tea.Prettify(s)
}

func (s RetryUserProvisioningEventRequest) GoString() string {
	return s.String()
}

func (s *RetryUserProvisioningEventRequest) SetDirectoryId(v string) *RetryUserProvisioningEventRequest {
	s.DirectoryId = &v
	return s
}

func (s *RetryUserProvisioningEventRequest) SetDuplicationStrategy(v string) *RetryUserProvisioningEventRequest {
	s.DuplicationStrategy = &v
	return s
}

func (s *RetryUserProvisioningEventRequest) SetEventId(v string) *RetryUserProvisioningEventRequest {
	s.EventId = &v
	return s
}

type RetryUserProvisioningEventResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RetryUserProvisioningEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetryUserProvisioningEventResponseBody) GoString() string {
	return s.String()
}

func (s *RetryUserProvisioningEventResponseBody) SetRequestId(v string) *RetryUserProvisioningEventResponseBody {
	s.RequestId = &v
	return s
}

type RetryUserProvisioningEventResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryUserProvisioningEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryUserProvisioningEventResponse) String() string {
	return tea.Prettify(s)
}

func (s RetryUserProvisioningEventResponse) GoString() string {
	return s.String()
}

func (s *RetryUserProvisioningEventResponse) SetHeaders(v map[string]*string) *RetryUserProvisioningEventResponse {
	s.Headers = v
	return s
}

func (s *RetryUserProvisioningEventResponse) SetStatusCode(v int32) *RetryUserProvisioningEventResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryUserProvisioningEventResponse) SetBody(v *RetryUserProvisioningEventResponseBody) *RetryUserProvisioningEventResponse {
	s.Body = v
	return s
}

type SetExternalSAMLIdentityProviderRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The metadata file of the IdP. The value of this parameter is Base64-encoded.
	//
	// The file is provided by the IdP that supports SAML 2.0.
	EncodedMetadataDocument *string `json:"EncodedMetadataDocument,omitempty" xml:"EncodedMetadataDocument,omitempty"`
	// The entity ID of the IdP.
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The logon URL of the IdP.
	LoginUrl *string `json:"LoginUrl,omitempty" xml:"LoginUrl,omitempty"`
	// The status of SSO logon. Valid values:
	//
	// *   Enabled
	// *   Disabled (default)
	SSOStatus *string `json:"SSOStatus,omitempty" xml:"SSOStatus,omitempty"`
	// Specifies whether CloudSSO needs to sign SAML requests. The requests are sent when users log on to the CloudSSO user portal to initiate SAML-based SSO. Valid values:
	//
	// *   true: yes
	// *   false: no (default)
	WantRequestSigned *bool `json:"WantRequestSigned,omitempty" xml:"WantRequestSigned,omitempty"`
	// The X.509 certificate in the PEM format. If you specify this parameter, all existing certificates are replaced.
	X509Certificate *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
}

func (s SetExternalSAMLIdentityProviderRequest) String() string {
	return tea.Prettify(s)
}

func (s SetExternalSAMLIdentityProviderRequest) GoString() string {
	return s.String()
}

func (s *SetExternalSAMLIdentityProviderRequest) SetDirectoryId(v string) *SetExternalSAMLIdentityProviderRequest {
	s.DirectoryId = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderRequest) SetEncodedMetadataDocument(v string) *SetExternalSAMLIdentityProviderRequest {
	s.EncodedMetadataDocument = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderRequest) SetEntityId(v string) *SetExternalSAMLIdentityProviderRequest {
	s.EntityId = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderRequest) SetLoginUrl(v string) *SetExternalSAMLIdentityProviderRequest {
	s.LoginUrl = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderRequest) SetSSOStatus(v string) *SetExternalSAMLIdentityProviderRequest {
	s.SSOStatus = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderRequest) SetWantRequestSigned(v bool) *SetExternalSAMLIdentityProviderRequest {
	s.WantRequestSigned = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderRequest) SetX509Certificate(v string) *SetExternalSAMLIdentityProviderRequest {
	s.X509Certificate = &v
	return s
}

type SetExternalSAMLIdentityProviderResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configurations of the IdP.
	SAMLIdentityProviderConfiguration *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration `json:"SAMLIdentityProviderConfiguration,omitempty" xml:"SAMLIdentityProviderConfiguration,omitempty" type:"Struct"`
}

func (s SetExternalSAMLIdentityProviderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetExternalSAMLIdentityProviderResponseBody) GoString() string {
	return s.String()
}

func (s *SetExternalSAMLIdentityProviderResponseBody) SetRequestId(v string) *SetExternalSAMLIdentityProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderResponseBody) SetSAMLIdentityProviderConfiguration(v *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) *SetExternalSAMLIdentityProviderResponseBody {
	s.SAMLIdentityProviderConfiguration = v
	return s
}

type SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration struct {
	// The ID of the SAML signing certificate.
	CertificateIds []*string `json:"CertificateIds,omitempty" xml:"CertificateIds,omitempty" type:"Repeated"`
	// The time when the IdP was configured for the first time.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The metadata file of the IdP. The value of this parameter is Base64-encoded.
	EncodedMetadataDocument *string `json:"EncodedMetadataDocument,omitempty" xml:"EncodedMetadataDocument,omitempty"`
	// The entity ID of the IdP.
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The logon URL of the IdP.
	LoginUrl *string `json:"LoginUrl,omitempty" xml:"LoginUrl,omitempty"`
	// The status of SSO logon. Valid values:
	//
	// *   Enabled
	// *   Disabled
	SSOStatus *string `json:"SSOStatus,omitempty" xml:"SSOStatus,omitempty"`
	// The time when the IdP configurations were last modified.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// Indicates whether CloudSSO needs to sign SAML requests. The requests are sent when users log on to the CloudSSO user portal to initiate SAML-based SSO. Valid values:
	//
	// *   true: yes
	// *   false: no (default)
	WantRequestSigned *bool `json:"WantRequestSigned,omitempty" xml:"WantRequestSigned,omitempty"`
}

func (s SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) String() string {
	return tea.Prettify(s)
}

func (s SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) GoString() string {
	return s.String()
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetCertificateIds(v []*string) *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.CertificateIds = v
	return s
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetCreateTime(v string) *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.CreateTime = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetDirectoryId(v string) *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.DirectoryId = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetEncodedMetadataDocument(v string) *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.EncodedMetadataDocument = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetEntityId(v string) *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.EntityId = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetLoginUrl(v string) *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.LoginUrl = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetSSOStatus(v string) *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.SSOStatus = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetUpdateTime(v string) *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.UpdateTime = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetWantRequestSigned(v bool) *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.WantRequestSigned = &v
	return s
}

type SetExternalSAMLIdentityProviderResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetExternalSAMLIdentityProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetExternalSAMLIdentityProviderResponse) String() string {
	return tea.Prettify(s)
}

func (s SetExternalSAMLIdentityProviderResponse) GoString() string {
	return s.String()
}

func (s *SetExternalSAMLIdentityProviderResponse) SetHeaders(v map[string]*string) *SetExternalSAMLIdentityProviderResponse {
	s.Headers = v
	return s
}

func (s *SetExternalSAMLIdentityProviderResponse) SetStatusCode(v int32) *SetExternalSAMLIdentityProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderResponse) SetBody(v *SetExternalSAMLIdentityProviderResponseBody) *SetExternalSAMLIdentityProviderResponse {
	s.Body = v
	return s
}

type SetLoginPreferenceRequest struct {
	DirectoryId       *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	LoginNetworkMasks *string `json:"LoginNetworkMasks,omitempty" xml:"LoginNetworkMasks,omitempty"`
}

func (s SetLoginPreferenceRequest) String() string {
	return tea.Prettify(s)
}

func (s SetLoginPreferenceRequest) GoString() string {
	return s.String()
}

func (s *SetLoginPreferenceRequest) SetDirectoryId(v string) *SetLoginPreferenceRequest {
	s.DirectoryId = &v
	return s
}

func (s *SetLoginPreferenceRequest) SetLoginNetworkMasks(v string) *SetLoginPreferenceRequest {
	s.LoginNetworkMasks = &v
	return s
}

type SetLoginPreferenceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLoginPreferenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetLoginPreferenceResponseBody) GoString() string {
	return s.String()
}

func (s *SetLoginPreferenceResponseBody) SetRequestId(v string) *SetLoginPreferenceResponseBody {
	s.RequestId = &v
	return s
}

type SetLoginPreferenceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetLoginPreferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetLoginPreferenceResponse) String() string {
	return tea.Prettify(s)
}

func (s SetLoginPreferenceResponse) GoString() string {
	return s.String()
}

func (s *SetLoginPreferenceResponse) SetHeaders(v map[string]*string) *SetLoginPreferenceResponse {
	s.Headers = v
	return s
}

func (s *SetLoginPreferenceResponse) SetStatusCode(v int32) *SetLoginPreferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLoginPreferenceResponse) SetBody(v *SetLoginPreferenceResponseBody) *SetLoginPreferenceResponse {
	s.Body = v
	return s
}

type SetMFAAuthenticationStatusRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The status of MFA. Valid values:
	//
	// *   Enabled
	// *   Disabled
	MFAAuthenticationStatus *string `json:"MFAAuthenticationStatus,omitempty" xml:"MFAAuthenticationStatus,omitempty"`
}

func (s SetMFAAuthenticationStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetMFAAuthenticationStatusRequest) GoString() string {
	return s.String()
}

func (s *SetMFAAuthenticationStatusRequest) SetDirectoryId(v string) *SetMFAAuthenticationStatusRequest {
	s.DirectoryId = &v
	return s
}

func (s *SetMFAAuthenticationStatusRequest) SetMFAAuthenticationStatus(v string) *SetMFAAuthenticationStatusRequest {
	s.MFAAuthenticationStatus = &v
	return s
}

type SetMFAAuthenticationStatusResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetMFAAuthenticationStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetMFAAuthenticationStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetMFAAuthenticationStatusResponseBody) SetRequestId(v string) *SetMFAAuthenticationStatusResponseBody {
	s.RequestId = &v
	return s
}

type SetMFAAuthenticationStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetMFAAuthenticationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetMFAAuthenticationStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetMFAAuthenticationStatusResponse) GoString() string {
	return s.String()
}

func (s *SetMFAAuthenticationStatusResponse) SetHeaders(v map[string]*string) *SetMFAAuthenticationStatusResponse {
	s.Headers = v
	return s
}

func (s *SetMFAAuthenticationStatusResponse) SetStatusCode(v int32) *SetMFAAuthenticationStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetMFAAuthenticationStatusResponse) SetBody(v *SetMFAAuthenticationStatusResponseBody) *SetMFAAuthenticationStatusResponse {
	s.Body = v
	return s
}

type SetPasswordPolicyRequest struct {
	DirectoryId                *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	MaxLoginAttempts           *int32  `json:"MaxLoginAttempts,omitempty" xml:"MaxLoginAttempts,omitempty"`
	MaxPasswordAge             *int32  `json:"MaxPasswordAge,omitempty" xml:"MaxPasswordAge,omitempty"`
	MinPasswordDifferentChars  *int32  `json:"MinPasswordDifferentChars,omitempty" xml:"MinPasswordDifferentChars,omitempty"`
	MinPasswordLength          *int32  `json:"MinPasswordLength,omitempty" xml:"MinPasswordLength,omitempty"`
	PasswordNotContainUsername *bool   `json:"PasswordNotContainUsername,omitempty" xml:"PasswordNotContainUsername,omitempty"`
	PasswordReusePrevention    *int32  `json:"PasswordReusePrevention,omitempty" xml:"PasswordReusePrevention,omitempty"`
}

func (s SetPasswordPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s SetPasswordPolicyRequest) GoString() string {
	return s.String()
}

func (s *SetPasswordPolicyRequest) SetDirectoryId(v string) *SetPasswordPolicyRequest {
	s.DirectoryId = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetMaxLoginAttempts(v int32) *SetPasswordPolicyRequest {
	s.MaxLoginAttempts = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetMaxPasswordAge(v int32) *SetPasswordPolicyRequest {
	s.MaxPasswordAge = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetMinPasswordDifferentChars(v int32) *SetPasswordPolicyRequest {
	s.MinPasswordDifferentChars = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetMinPasswordLength(v int32) *SetPasswordPolicyRequest {
	s.MinPasswordLength = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetPasswordNotContainUsername(v bool) *SetPasswordPolicyRequest {
	s.PasswordNotContainUsername = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetPasswordReusePrevention(v int32) *SetPasswordPolicyRequest {
	s.PasswordReusePrevention = &v
	return s
}

type SetPasswordPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetPasswordPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetPasswordPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *SetPasswordPolicyResponseBody) SetRequestId(v string) *SetPasswordPolicyResponseBody {
	s.RequestId = &v
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

type SetSCIMSynchronizationStatusRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The status of SCIM synchronization. Valid values:
	//
	// *   Enabled
	// *   Disabled
	SCIMSynchronizationStatus *string `json:"SCIMSynchronizationStatus,omitempty" xml:"SCIMSynchronizationStatus,omitempty"`
}

func (s SetSCIMSynchronizationStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetSCIMSynchronizationStatusRequest) GoString() string {
	return s.String()
}

func (s *SetSCIMSynchronizationStatusRequest) SetDirectoryId(v string) *SetSCIMSynchronizationStatusRequest {
	s.DirectoryId = &v
	return s
}

func (s *SetSCIMSynchronizationStatusRequest) SetSCIMSynchronizationStatus(v string) *SetSCIMSynchronizationStatusRequest {
	s.SCIMSynchronizationStatus = &v
	return s
}

type SetSCIMSynchronizationStatusResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetSCIMSynchronizationStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetSCIMSynchronizationStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetSCIMSynchronizationStatusResponseBody) SetRequestId(v string) *SetSCIMSynchronizationStatusResponseBody {
	s.RequestId = &v
	return s
}

type SetSCIMSynchronizationStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetSCIMSynchronizationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetSCIMSynchronizationStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetSCIMSynchronizationStatusResponse) GoString() string {
	return s.String()
}

func (s *SetSCIMSynchronizationStatusResponse) SetHeaders(v map[string]*string) *SetSCIMSynchronizationStatusResponse {
	s.Headers = v
	return s
}

func (s *SetSCIMSynchronizationStatusResponse) SetStatusCode(v int32) *SetSCIMSynchronizationStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetSCIMSynchronizationStatusResponse) SetBody(v *SetSCIMSynchronizationStatusResponseBody) *SetSCIMSynchronizationStatusResponse {
	s.Body = v
	return s
}

type UpdateAccessConfigurationRequest struct {
	// The ID of the access configuration.
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The new description of the access configuration.
	//
	// The description can be up to 1,024 characters in length.
	NewDescription *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	// The new initial web page that is displayed after a CloudSSO user accesses an account in your resource directory by using the access configuration.
	//
	// The web page must be a page of the Alibaba Cloud Management Console.
	NewRelayState *string `json:"NewRelayState,omitempty" xml:"NewRelayState,omitempty"`
	// The new duration of a session in which a CloudSSO user accesses an account in your resource directory by using the access configuration.
	//
	// Unit: seconds.
	//
	// Valid values: 900 to 43200. The value 900 indicates 15 minutes. The value 43200 indicates 12 hours.
	NewSessionDuration *int32 `json:"NewSessionDuration,omitempty" xml:"NewSessionDuration,omitempty"`
}

func (s UpdateAccessConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccessConfigurationRequest) GoString() string {
	return s.String()
}

func (s *UpdateAccessConfigurationRequest) SetAccessConfigurationId(v string) *UpdateAccessConfigurationRequest {
	s.AccessConfigurationId = &v
	return s
}

func (s *UpdateAccessConfigurationRequest) SetDirectoryId(v string) *UpdateAccessConfigurationRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateAccessConfigurationRequest) SetNewDescription(v string) *UpdateAccessConfigurationRequest {
	s.NewDescription = &v
	return s
}

func (s *UpdateAccessConfigurationRequest) SetNewRelayState(v string) *UpdateAccessConfigurationRequest {
	s.NewRelayState = &v
	return s
}

func (s *UpdateAccessConfigurationRequest) SetNewSessionDuration(v int32) *UpdateAccessConfigurationRequest {
	s.NewSessionDuration = &v
	return s
}

type UpdateAccessConfigurationResponseBody struct {
	// The information about the access configuration.
	AccessConfiguration *UpdateAccessConfigurationResponseBodyAccessConfiguration `json:"AccessConfiguration,omitempty" xml:"AccessConfiguration,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAccessConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccessConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAccessConfigurationResponseBody) SetAccessConfiguration(v *UpdateAccessConfigurationResponseBodyAccessConfiguration) *UpdateAccessConfigurationResponseBody {
	s.AccessConfiguration = v
	return s
}

func (s *UpdateAccessConfigurationResponseBody) SetRequestId(v string) *UpdateAccessConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAccessConfigurationResponseBodyAccessConfiguration struct {
	// The ID of the access configuration.
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The name of the access configuration.
	AccessConfigurationName *string `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	// The time when the access configuration was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the access configuration.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The initial web page that is displayed after a CloudSSO user accesses an account in your resource directory by using the access configuration.
	RelayState *string `json:"RelayState,omitempty" xml:"RelayState,omitempty"`
	// The duration of a session in which a CloudSSO user accesses an account in your resource directory by using the access configuration.
	//
	// Unit: seconds.
	SessionDuration *int32 `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	// The status notification.
	StatusNotifications []*string `json:"StatusNotifications,omitempty" xml:"StatusNotifications,omitempty" type:"Repeated"`
	// The time when the information about the access configuration was modified.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s UpdateAccessConfigurationResponseBodyAccessConfiguration) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccessConfigurationResponseBodyAccessConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateAccessConfigurationResponseBodyAccessConfiguration) SetAccessConfigurationId(v string) *UpdateAccessConfigurationResponseBodyAccessConfiguration {
	s.AccessConfigurationId = &v
	return s
}

func (s *UpdateAccessConfigurationResponseBodyAccessConfiguration) SetAccessConfigurationName(v string) *UpdateAccessConfigurationResponseBodyAccessConfiguration {
	s.AccessConfigurationName = &v
	return s
}

func (s *UpdateAccessConfigurationResponseBodyAccessConfiguration) SetCreateTime(v string) *UpdateAccessConfigurationResponseBodyAccessConfiguration {
	s.CreateTime = &v
	return s
}

func (s *UpdateAccessConfigurationResponseBodyAccessConfiguration) SetDescription(v string) *UpdateAccessConfigurationResponseBodyAccessConfiguration {
	s.Description = &v
	return s
}

func (s *UpdateAccessConfigurationResponseBodyAccessConfiguration) SetRelayState(v string) *UpdateAccessConfigurationResponseBodyAccessConfiguration {
	s.RelayState = &v
	return s
}

func (s *UpdateAccessConfigurationResponseBodyAccessConfiguration) SetSessionDuration(v int32) *UpdateAccessConfigurationResponseBodyAccessConfiguration {
	s.SessionDuration = &v
	return s
}

func (s *UpdateAccessConfigurationResponseBodyAccessConfiguration) SetStatusNotifications(v []*string) *UpdateAccessConfigurationResponseBodyAccessConfiguration {
	s.StatusNotifications = v
	return s
}

func (s *UpdateAccessConfigurationResponseBodyAccessConfiguration) SetUpdateTime(v string) *UpdateAccessConfigurationResponseBodyAccessConfiguration {
	s.UpdateTime = &v
	return s
}

type UpdateAccessConfigurationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAccessConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccessConfigurationResponse) GoString() string {
	return s.String()
}

func (s *UpdateAccessConfigurationResponse) SetHeaders(v map[string]*string) *UpdateAccessConfigurationResponse {
	s.Headers = v
	return s
}

func (s *UpdateAccessConfigurationResponse) SetStatusCode(v int32) *UpdateAccessConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAccessConfigurationResponse) SetBody(v *UpdateAccessConfigurationResponseBody) *UpdateAccessConfigurationResponse {
	s.Body = v
	return s
}

type UpdateDirectoryRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The new name of the directory. The name must be globally unique.
	//
	// The name can contain lowercase letters, digits, and hyphens (-). The name cannot start or end with a hyphen (-) and cannot have two consecutive hyphens (-). If you want to start the new name of the directory starts with `d-`, you must set this parameter to the ID of the directory.
	//
	// The name must be 2 to 64 characters in length.
	NewDirectoryName *string `json:"NewDirectoryName,omitempty" xml:"NewDirectoryName,omitempty"`
}

func (s UpdateDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDirectoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateDirectoryRequest) SetDirectoryId(v string) *UpdateDirectoryRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateDirectoryRequest) SetNewDirectoryName(v string) *UpdateDirectoryRequest {
	s.NewDirectoryName = &v
	return s
}

type UpdateDirectoryResponseBody struct {
	// The information about the directory.
	Directory *UpdateDirectoryResponseBodyDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDirectoryResponseBody) SetDirectory(v *UpdateDirectoryResponseBodyDirectory) *UpdateDirectoryResponseBody {
	s.Directory = v
	return s
}

func (s *UpdateDirectoryResponseBody) SetRequestId(v string) *UpdateDirectoryResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDirectoryResponseBodyDirectory struct {
	// The time when the directory was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The name of the directory.
	DirectoryName *string `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	// The region ID of the directory.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The time when the directory was modified.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s UpdateDirectoryResponseBodyDirectory) String() string {
	return tea.Prettify(s)
}

func (s UpdateDirectoryResponseBodyDirectory) GoString() string {
	return s.String()
}

func (s *UpdateDirectoryResponseBodyDirectory) SetCreateTime(v string) *UpdateDirectoryResponseBodyDirectory {
	s.CreateTime = &v
	return s
}

func (s *UpdateDirectoryResponseBodyDirectory) SetDirectoryId(v string) *UpdateDirectoryResponseBodyDirectory {
	s.DirectoryId = &v
	return s
}

func (s *UpdateDirectoryResponseBodyDirectory) SetDirectoryName(v string) *UpdateDirectoryResponseBodyDirectory {
	s.DirectoryName = &v
	return s
}

func (s *UpdateDirectoryResponseBodyDirectory) SetRegion(v string) *UpdateDirectoryResponseBodyDirectory {
	s.Region = &v
	return s
}

func (s *UpdateDirectoryResponseBodyDirectory) SetUpdateTime(v string) *UpdateDirectoryResponseBodyDirectory {
	s.UpdateTime = &v
	return s
}

type UpdateDirectoryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDirectoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateDirectoryResponse) SetHeaders(v map[string]*string) *UpdateDirectoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateDirectoryResponse) SetStatusCode(v int32) *UpdateDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDirectoryResponse) SetBody(v *UpdateDirectoryResponseBody) *UpdateDirectoryResponse {
	s.Body = v
	return s
}

type UpdateGroupRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the group.
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The new description of the group.
	NewDescription *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	// The new name of the group.
	NewGroupName *string `json:"NewGroupName,omitempty" xml:"NewGroupName,omitempty"`
}

func (s UpdateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupRequest) SetDirectoryId(v string) *UpdateGroupRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateGroupRequest) SetGroupId(v string) *UpdateGroupRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateGroupRequest) SetNewDescription(v string) *UpdateGroupRequest {
	s.NewDescription = &v
	return s
}

func (s *UpdateGroupRequest) SetNewGroupName(v string) *UpdateGroupRequest {
	s.NewGroupName = &v
	return s
}

type UpdateGroupResponseBody struct {
	// The information about the group.
	Group *UpdateGroupResponseBodyGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	// The ID of the request.
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
	// The time when the group was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the group.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the group.
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the group. Valid values:
	//
	// *   Manual: The group is manually created.
	// *   Synchronized: The user is synchronized from an external identity provider (IdP).
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	// The time when the information about the group was modified.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s UpdateGroupResponseBodyGroup) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupResponseBodyGroup) GoString() string {
	return s.String()
}

func (s *UpdateGroupResponseBodyGroup) SetCreateTime(v string) *UpdateGroupResponseBodyGroup {
	s.CreateTime = &v
	return s
}

func (s *UpdateGroupResponseBodyGroup) SetDescription(v string) *UpdateGroupResponseBodyGroup {
	s.Description = &v
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

func (s *UpdateGroupResponseBodyGroup) SetProvisionType(v string) *UpdateGroupResponseBodyGroup {
	s.ProvisionType = &v
	return s
}

func (s *UpdateGroupResponseBodyGroup) SetUpdateTime(v string) *UpdateGroupResponseBodyGroup {
	s.UpdateTime = &v
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

type UpdateInlinePolicyForAccessConfigurationRequest struct {
	// The ID of the access configuration.
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The name of the inline policy.
	InlinePolicyName *string `json:"InlinePolicyName,omitempty" xml:"InlinePolicyName,omitempty"`
	// The new configurations of the inline policy.
	//
	// The value can be up to 4,096 characters in length.
	//
	// For more information about the syntax and structure of RAM policies, see [Policy syntax and structure](~~93739~~).
	NewInlinePolicyDocument *string `json:"NewInlinePolicyDocument,omitempty" xml:"NewInlinePolicyDocument,omitempty"`
}

func (s UpdateInlinePolicyForAccessConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInlinePolicyForAccessConfigurationRequest) GoString() string {
	return s.String()
}

func (s *UpdateInlinePolicyForAccessConfigurationRequest) SetAccessConfigurationId(v string) *UpdateInlinePolicyForAccessConfigurationRequest {
	s.AccessConfigurationId = &v
	return s
}

func (s *UpdateInlinePolicyForAccessConfigurationRequest) SetDirectoryId(v string) *UpdateInlinePolicyForAccessConfigurationRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateInlinePolicyForAccessConfigurationRequest) SetInlinePolicyName(v string) *UpdateInlinePolicyForAccessConfigurationRequest {
	s.InlinePolicyName = &v
	return s
}

func (s *UpdateInlinePolicyForAccessConfigurationRequest) SetNewInlinePolicyDocument(v string) *UpdateInlinePolicyForAccessConfigurationRequest {
	s.NewInlinePolicyDocument = &v
	return s
}

type UpdateInlinePolicyForAccessConfigurationResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInlinePolicyForAccessConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInlinePolicyForAccessConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInlinePolicyForAccessConfigurationResponseBody) SetRequestId(v string) *UpdateInlinePolicyForAccessConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type UpdateInlinePolicyForAccessConfigurationResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInlinePolicyForAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInlinePolicyForAccessConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInlinePolicyForAccessConfigurationResponse) GoString() string {
	return s.String()
}

func (s *UpdateInlinePolicyForAccessConfigurationResponse) SetHeaders(v map[string]*string) *UpdateInlinePolicyForAccessConfigurationResponse {
	s.Headers = v
	return s
}

func (s *UpdateInlinePolicyForAccessConfigurationResponse) SetStatusCode(v int32) *UpdateInlinePolicyForAccessConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInlinePolicyForAccessConfigurationResponse) SetBody(v *UpdateInlinePolicyForAccessConfigurationResponseBody) *UpdateInlinePolicyForAccessConfigurationResponse {
	s.Body = v
	return s
}

type UpdateMFAAuthenticationSettingsRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// Specifies whether to enable MFA for all users. Valid value:
	//
	// - Enabled: enables MFA for all users.
	// - Byuser: uses user-specific settings. For more information about how to configure MFA for a single user, see [UpdateUserMFAAuthenticationSettings](~~450135~~).
	// - Disabled: disables MFA for all users.
	// - OnlyRiskyLogin: MFA is required only for unusual logons.
	MFAAuthenticationSettings *string `json:"MFAAuthenticationSettings,omitempty" xml:"MFAAuthenticationSettings,omitempty"`
	// Specifies whether MFA is required for users who initiated unusual logons. Valid value:
	//
	// - Autonomous: MFA is prompted for users who initiated unusual logons. However, the users are allowed to skip MFA. If an MFA device is bound to a user who initiated an unusual logon, the user must pass MFA.
	//
	// - EnforceVerify: MFA is required. If no MFA devices are bound to a user who initiated an unusual logon, the user must bind an MFA device. If an MFA device is already bound to a user who initiated an unusual logon, the user must pass MFA.
	OperationForRiskLogin *string `json:"OperationForRiskLogin,omitempty" xml:"OperationForRiskLogin,omitempty"`
}

func (s UpdateMFAAuthenticationSettingsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMFAAuthenticationSettingsRequest) GoString() string {
	return s.String()
}

func (s *UpdateMFAAuthenticationSettingsRequest) SetDirectoryId(v string) *UpdateMFAAuthenticationSettingsRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateMFAAuthenticationSettingsRequest) SetMFAAuthenticationSettings(v string) *UpdateMFAAuthenticationSettingsRequest {
	s.MFAAuthenticationSettings = &v
	return s
}

func (s *UpdateMFAAuthenticationSettingsRequest) SetOperationForRiskLogin(v string) *UpdateMFAAuthenticationSettingsRequest {
	s.OperationForRiskLogin = &v
	return s
}

type UpdateMFAAuthenticationSettingsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMFAAuthenticationSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMFAAuthenticationSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMFAAuthenticationSettingsResponseBody) SetRequestId(v string) *UpdateMFAAuthenticationSettingsResponseBody {
	s.RequestId = &v
	return s
}

type UpdateMFAAuthenticationSettingsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMFAAuthenticationSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMFAAuthenticationSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMFAAuthenticationSettingsResponse) GoString() string {
	return s.String()
}

func (s *UpdateMFAAuthenticationSettingsResponse) SetHeaders(v map[string]*string) *UpdateMFAAuthenticationSettingsResponse {
	s.Headers = v
	return s
}

func (s *UpdateMFAAuthenticationSettingsResponse) SetStatusCode(v int32) *UpdateMFAAuthenticationSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMFAAuthenticationSettingsResponse) SetBody(v *UpdateMFAAuthenticationSettingsResponseBody) *UpdateMFAAuthenticationSettingsResponse {
	s.Body = v
	return s
}

type UpdateSCIMServerCredentialStatusRequest struct {
	// The ID of the SCIM credential.
	CredentialId *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The new status of the SCIM credential. Valid values:
	//
	// *   Enabled: The SCIM credential is enabled.
	// *   Disabled: The SCIM credential is disabled.
	NewStatus *string `json:"NewStatus,omitempty" xml:"NewStatus,omitempty"`
}

func (s UpdateSCIMServerCredentialStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSCIMServerCredentialStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateSCIMServerCredentialStatusRequest) SetCredentialId(v string) *UpdateSCIMServerCredentialStatusRequest {
	s.CredentialId = &v
	return s
}

func (s *UpdateSCIMServerCredentialStatusRequest) SetDirectoryId(v string) *UpdateSCIMServerCredentialStatusRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateSCIMServerCredentialStatusRequest) SetNewStatus(v string) *UpdateSCIMServerCredentialStatusRequest {
	s.NewStatus = &v
	return s
}

type UpdateSCIMServerCredentialStatusResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the SCIM credential.
	SCIMServerCredential *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential `json:"SCIMServerCredential,omitempty" xml:"SCIMServerCredential,omitempty" type:"Struct"`
}

func (s UpdateSCIMServerCredentialStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSCIMServerCredentialStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSCIMServerCredentialStatusResponseBody) SetRequestId(v string) *UpdateSCIMServerCredentialStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSCIMServerCredentialStatusResponseBody) SetSCIMServerCredential(v *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential) *UpdateSCIMServerCredentialStatusResponseBody {
	s.SCIMServerCredential = v
	return s
}

type UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential struct {
	// The time when the SCIM credential was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the SCIM credential.
	CredentialId *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
	// The type of the SCIM credential.
	CredentialType *string `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The time when the SCIM credential expires.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The status of the SCIM credential. Valid values:
	//
	// *   Enabled: The SCIM credential is enabled.
	// *   Disabled: The SCIM credential is disabled.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential) String() string {
	return tea.Prettify(s)
}

func (s UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential) GoString() string {
	return s.String()
}

func (s *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential) SetCreateTime(v string) *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential {
	s.CreateTime = &v
	return s
}

func (s *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential) SetCredentialId(v string) *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential {
	s.CredentialId = &v
	return s
}

func (s *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential) SetCredentialType(v string) *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential {
	s.CredentialType = &v
	return s
}

func (s *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential) SetDirectoryId(v string) *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential {
	s.DirectoryId = &v
	return s
}

func (s *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential) SetExpireTime(v string) *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential {
	s.ExpireTime = &v
	return s
}

func (s *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential) SetStatus(v string) *UpdateSCIMServerCredentialStatusResponseBodySCIMServerCredential {
	s.Status = &v
	return s
}

type UpdateSCIMServerCredentialStatusResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSCIMServerCredentialStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSCIMServerCredentialStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSCIMServerCredentialStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateSCIMServerCredentialStatusResponse) SetHeaders(v map[string]*string) *UpdateSCIMServerCredentialStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateSCIMServerCredentialStatusResponse) SetStatusCode(v int32) *UpdateSCIMServerCredentialStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSCIMServerCredentialStatusResponse) SetBody(v *UpdateSCIMServerCredentialStatusResponseBody) *UpdateSCIMServerCredentialStatusResponse {
	s.Body = v
	return s
}

type UpdateUserRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The new description of the user.
	NewDescription *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	// The new display name of the user.
	NewDisplayName *string `json:"NewDisplayName,omitempty" xml:"NewDisplayName,omitempty"`
	// The new email address of the user.
	NewEmail *string `json:"NewEmail,omitempty" xml:"NewEmail,omitempty"`
	// The new first name of the user.
	NewFirstName *string `json:"NewFirstName,omitempty" xml:"NewFirstName,omitempty"`
	// The new last name of the user.
	NewLastName *string `json:"NewLastName,omitempty" xml:"NewLastName,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) SetDirectoryId(v string) *UpdateUserRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateUserRequest) SetNewDescription(v string) *UpdateUserRequest {
	s.NewDescription = &v
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

func (s *UpdateUserRequest) SetNewFirstName(v string) *UpdateUserRequest {
	s.NewFirstName = &v
	return s
}

func (s *UpdateUserRequest) SetNewLastName(v string) *UpdateUserRequest {
	s.NewLastName = &v
	return s
}

func (s *UpdateUserRequest) SetUserId(v string) *UpdateUserRequest {
	s.UserId = &v
	return s
}

type UpdateUserResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the user.
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
	// The time when the user was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the user.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the user.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the user.
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The first name of the user.
	FirstName *string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	// The last name of the user.
	LastName *string `json:"LastName,omitempty" xml:"LastName,omitempty"`
	// The type of the user. Valid values:
	//
	// *   Manual: The user is manually created.
	// *   Synchronized: The user is synchronized from an external identity provider (IdP).
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	// The status of the user. Valid values:
	//
	// *   Enabled: The logon of the user is enabled.
	// *   Disabled: The logon of the user is disabled.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the information about the user was modified.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The name of the user.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UpdateUserResponseBodyUser) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponseBodyUser) GoString() string {
	return s.String()
}

func (s *UpdateUserResponseBodyUser) SetCreateTime(v string) *UpdateUserResponseBodyUser {
	s.CreateTime = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetDescription(v string) *UpdateUserResponseBodyUser {
	s.Description = &v
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

func (s *UpdateUserResponseBodyUser) SetFirstName(v string) *UpdateUserResponseBodyUser {
	s.FirstName = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetLastName(v string) *UpdateUserResponseBodyUser {
	s.LastName = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetProvisionType(v string) *UpdateUserResponseBodyUser {
	s.ProvisionType = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetStatus(v string) *UpdateUserResponseBodyUser {
	s.Status = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetUpdateTime(v string) *UpdateUserResponseBodyUser {
	s.UpdateTime = &v
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

type UpdateUserMFAAuthenticationSettingsRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// Specifies whether to enable MFA for the user. Valid values:
	//
	// *   Enabled: enables MFA for the user.
	// *   Disabled: disables MFA for the user.
	UserMFAAuthenticationSettings *string `json:"UserMFAAuthenticationSettings,omitempty" xml:"UserMFAAuthenticationSettings,omitempty"`
}

func (s UpdateUserMFAAuthenticationSettingsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserMFAAuthenticationSettingsRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserMFAAuthenticationSettingsRequest) SetDirectoryId(v string) *UpdateUserMFAAuthenticationSettingsRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateUserMFAAuthenticationSettingsRequest) SetUserId(v string) *UpdateUserMFAAuthenticationSettingsRequest {
	s.UserId = &v
	return s
}

func (s *UpdateUserMFAAuthenticationSettingsRequest) SetUserMFAAuthenticationSettings(v string) *UpdateUserMFAAuthenticationSettingsRequest {
	s.UserMFAAuthenticationSettings = &v
	return s
}

type UpdateUserMFAAuthenticationSettingsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUserMFAAuthenticationSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserMFAAuthenticationSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserMFAAuthenticationSettingsResponseBody) SetRequestId(v string) *UpdateUserMFAAuthenticationSettingsResponseBody {
	s.RequestId = &v
	return s
}

type UpdateUserMFAAuthenticationSettingsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserMFAAuthenticationSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserMFAAuthenticationSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserMFAAuthenticationSettingsResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserMFAAuthenticationSettingsResponse) SetHeaders(v map[string]*string) *UpdateUserMFAAuthenticationSettingsResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserMFAAuthenticationSettingsResponse) SetStatusCode(v int32) *UpdateUserMFAAuthenticationSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserMFAAuthenticationSettingsResponse) SetBody(v *UpdateUserMFAAuthenticationSettingsResponseBody) *UpdateUserMFAAuthenticationSettingsResponse {
	s.Body = v
	return s
}

type UpdateUserProvisioningRequest struct {
	DirectoryId            *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	NewDeletionStrategy    *string `json:"NewDeletionStrategy,omitempty" xml:"NewDeletionStrategy,omitempty"`
	NewDescription         *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	NewDuplicationStrategy *string `json:"NewDuplicationStrategy,omitempty" xml:"NewDuplicationStrategy,omitempty"`
	UserProvisioningId     *string `json:"UserProvisioningId,omitempty" xml:"UserProvisioningId,omitempty"`
}

func (s UpdateUserProvisioningRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserProvisioningRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserProvisioningRequest) SetDirectoryId(v string) *UpdateUserProvisioningRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateUserProvisioningRequest) SetNewDeletionStrategy(v string) *UpdateUserProvisioningRequest {
	s.NewDeletionStrategy = &v
	return s
}

func (s *UpdateUserProvisioningRequest) SetNewDescription(v string) *UpdateUserProvisioningRequest {
	s.NewDescription = &v
	return s
}

func (s *UpdateUserProvisioningRequest) SetNewDuplicationStrategy(v string) *UpdateUserProvisioningRequest {
	s.NewDuplicationStrategy = &v
	return s
}

func (s *UpdateUserProvisioningRequest) SetUserProvisioningId(v string) *UpdateUserProvisioningRequest {
	s.UserProvisioningId = &v
	return s
}

type UpdateUserProvisioningResponseBody struct {
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserProvisioning *UpdateUserProvisioningResponseBodyUserProvisioning `json:"UserProvisioning,omitempty" xml:"UserProvisioning,omitempty" type:"Struct"`
}

func (s UpdateUserProvisioningResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserProvisioningResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserProvisioningResponseBody) SetRequestId(v string) *UpdateUserProvisioningResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserProvisioningResponseBody) SetUserProvisioning(v *UpdateUserProvisioningResponseBodyUserProvisioning) *UpdateUserProvisioningResponseBody {
	s.UserProvisioning = v
	return s
}

type UpdateUserProvisioningResponseBodyUserProvisioning struct {
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeletionStrategy    *string `json:"DeletionStrategy,omitempty" xml:"DeletionStrategy,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DirectoryId         *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	DuplicationStrategy *string `json:"DuplicationStrategy,omitempty" xml:"DuplicationStrategy,omitempty"`
	OwnerPk             *string `json:"OwnerPk,omitempty" xml:"OwnerPk,omitempty"`
	PrincipalId         *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	PrincipalName       *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	PrincipalType       *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetId            *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetName          *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	TargetPath          *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	TargetType          *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	UpdateTime          *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserProvisioningId  *string `json:"UserProvisioningId,omitempty" xml:"UserProvisioningId,omitempty"`
}

func (s UpdateUserProvisioningResponseBodyUserProvisioning) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserProvisioningResponseBodyUserProvisioning) GoString() string {
	return s.String()
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetCreateTime(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.CreateTime = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetDeletionStrategy(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.DeletionStrategy = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetDescription(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.Description = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetDirectoryId(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.DirectoryId = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetDuplicationStrategy(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.DuplicationStrategy = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetOwnerPk(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.OwnerPk = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetPrincipalId(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.PrincipalId = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetPrincipalName(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.PrincipalName = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetPrincipalType(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.PrincipalType = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetStatus(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.Status = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetTargetId(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.TargetId = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetTargetName(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.TargetName = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetTargetPath(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.TargetPath = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetTargetType(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.TargetType = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetUpdateTime(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.UpdateTime = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetUserProvisioningId(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.UserProvisioningId = &v
	return s
}

type UpdateUserProvisioningResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserProvisioningResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserProvisioningResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserProvisioningResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserProvisioningResponse) SetHeaders(v map[string]*string) *UpdateUserProvisioningResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserProvisioningResponse) SetStatusCode(v int32) *UpdateUserProvisioningResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserProvisioningResponse) SetBody(v *UpdateUserProvisioningResponseBody) *UpdateUserProvisioningResponse {
	s.Body = v
	return s
}

type UpdateUserProvisioningConfigurationRequest struct {
	DirectoryId           *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	NewDefaultLandingPage *string `json:"NewDefaultLandingPage,omitempty" xml:"NewDefaultLandingPage,omitempty"`
	NewSessionDuration    *int32  `json:"NewSessionDuration,omitempty" xml:"NewSessionDuration,omitempty"`
}

func (s UpdateUserProvisioningConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserProvisioningConfigurationRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserProvisioningConfigurationRequest) SetDirectoryId(v string) *UpdateUserProvisioningConfigurationRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateUserProvisioningConfigurationRequest) SetNewDefaultLandingPage(v string) *UpdateUserProvisioningConfigurationRequest {
	s.NewDefaultLandingPage = &v
	return s
}

func (s *UpdateUserProvisioningConfigurationRequest) SetNewSessionDuration(v int32) *UpdateUserProvisioningConfigurationRequest {
	s.NewSessionDuration = &v
	return s
}

type UpdateUserProvisioningConfigurationResponseBody struct {
	RequestId                     *string                                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserProvisioningConfiguration *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration `json:"UserProvisioningConfiguration,omitempty" xml:"UserProvisioningConfiguration,omitempty" type:"Struct"`
}

func (s UpdateUserProvisioningConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserProvisioningConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserProvisioningConfigurationResponseBody) SetRequestId(v string) *UpdateUserProvisioningConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserProvisioningConfigurationResponseBody) SetUserProvisioningConfiguration(v *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) *UpdateUserProvisioningConfigurationResponseBody {
	s.UserProvisioningConfiguration = v
	return s
}

type UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration struct {
	CreateTime         *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DefaultLandingPage *string `json:"DefaultLandingPage,omitempty" xml:"DefaultLandingPage,omitempty"`
	DirectoryId        *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	SessionDuration    *int32  `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	UpdateTime         *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) SetCreateTime(v string) *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration {
	s.CreateTime = &v
	return s
}

func (s *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) SetDefaultLandingPage(v string) *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration {
	s.DefaultLandingPage = &v
	return s
}

func (s *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) SetDirectoryId(v string) *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration {
	s.DirectoryId = &v
	return s
}

func (s *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) SetSessionDuration(v int32) *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration {
	s.SessionDuration = &v
	return s
}

func (s *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) SetUpdateTime(v string) *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration {
	s.UpdateTime = &v
	return s
}

type UpdateUserProvisioningConfigurationResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserProvisioningConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserProvisioningConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserProvisioningConfigurationResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserProvisioningConfigurationResponse) SetHeaders(v map[string]*string) *UpdateUserProvisioningConfigurationResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserProvisioningConfigurationResponse) SetStatusCode(v int32) *UpdateUserProvisioningConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserProvisioningConfigurationResponse) SetBody(v *UpdateUserProvisioningConfigurationResponseBody) *UpdateUserProvisioningConfigurationResponse {
	s.Body = v
	return s
}

type UpdateUserStatusRequest struct {
	// The ID of the directory.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The new status of the user. Valid values:
	//
	// *   Enabled: The logon of the user is enabled.
	// *   Disabled: The logon of the user is disabled.
	NewStatus *string `json:"NewStatus,omitempty" xml:"NewStatus,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateUserStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserStatusRequest) SetDirectoryId(v string) *UpdateUserStatusRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateUserStatusRequest) SetNewStatus(v string) *UpdateUserStatusRequest {
	s.NewStatus = &v
	return s
}

func (s *UpdateUserStatusRequest) SetUserId(v string) *UpdateUserStatusRequest {
	s.UserId = &v
	return s
}

type UpdateUserStatusResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUserStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserStatusResponseBody) SetRequestId(v string) *UpdateUserStatusResponseBody {
	s.RequestId = &v
	return s
}

type UpdateUserStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserStatusResponse) SetHeaders(v map[string]*string) *UpdateUserStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserStatusResponse) SetStatusCode(v int32) *UpdateUserStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserStatusResponse) SetBody(v *UpdateUserStatusResponseBody) *UpdateUserStatusResponse {
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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cloudsso"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

/**
 * You can add up to two SAML signing certificates.
 * This topic provides an example on how to add a SAML signing certificate to the directory `d-00fc2p61****`.
 *
 * @param request AddExternalSAMLIdPCertificateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AddExternalSAMLIdPCertificateResponse
 */
func (client *Client) AddExternalSAMLIdPCertificateWithOptions(request *AddExternalSAMLIdPCertificateRequest, runtime *util.RuntimeOptions) (_result *AddExternalSAMLIdPCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.X509Certificate)) {
		query["X509Certificate"] = request.X509Certificate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddExternalSAMLIdPCertificate"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddExternalSAMLIdPCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can add up to two SAML signing certificates.
 * This topic provides an example on how to add a SAML signing certificate to the directory `d-00fc2p61****`.
 *
 * @param request AddExternalSAMLIdPCertificateRequest
 * @return AddExternalSAMLIdPCertificateResponse
 */
func (client *Client) AddExternalSAMLIdPCertificate(request *AddExternalSAMLIdPCertificateRequest) (_result *AddExternalSAMLIdPCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddExternalSAMLIdPCertificateResponse{}
	_body, _err := client.AddExternalSAMLIdPCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to add the system policy `AliyunECSFullAccess` to the access configuration `ac-00jhtfl8thteu6uj****`.
 *
 * @param request AddPermissionPolicyToAccessConfigurationRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AddPermissionPolicyToAccessConfigurationResponse
 */
func (client *Client) AddPermissionPolicyToAccessConfigurationWithOptions(request *AddPermissionPolicyToAccessConfigurationRequest, runtime *util.RuntimeOptions) (_result *AddPermissionPolicyToAccessConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessConfigurationId)) {
		query["AccessConfigurationId"] = request.AccessConfigurationId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.InlinePolicyDocument)) {
		query["InlinePolicyDocument"] = request.InlinePolicyDocument
	}

	if !tea.BoolValue(util.IsUnset(request.PermissionPolicyName)) {
		query["PermissionPolicyName"] = request.PermissionPolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.PermissionPolicyType)) {
		query["PermissionPolicyType"] = request.PermissionPolicyType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddPermissionPolicyToAccessConfiguration"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddPermissionPolicyToAccessConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to add the system policy `AliyunECSFullAccess` to the access configuration `ac-00jhtfl8thteu6uj****`.
 *
 * @param request AddPermissionPolicyToAccessConfigurationRequest
 * @return AddPermissionPolicyToAccessConfigurationResponse
 */
func (client *Client) AddPermissionPolicyToAccessConfiguration(request *AddPermissionPolicyToAccessConfigurationRequest) (_result *AddPermissionPolicyToAccessConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddPermissionPolicyToAccessConfigurationResponse{}
	_body, _err := client.AddPermissionPolicyToAccessConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * If System for Cross-domain Identity Management (SCIM) synchronization is enabled, you cannot add a user to a group that is synchronized by using SCIM.
 * This topic provides an example of how to add the user `u-00q8wbq42wiltcrk****` to the group `g-00jqzghi2n3o5hkh****`.
 *
 * @param request AddUserToGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AddUserToGroupResponse
 */
func (client *Client) AddUserToGroupWithOptions(request *AddUserToGroupRequest, runtime *util.RuntimeOptions) (_result *AddUserToGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddUserToGroup"),
		Version:     tea.String("2021-05-15"),
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

/**
 * If System for Cross-domain Identity Management (SCIM) synchronization is enabled, you cannot add a user to a group that is synchronized by using SCIM.
 * This topic provides an example of how to add the user `u-00q8wbq42wiltcrk****` to the group `g-00jqzghi2n3o5hkh****`.
 *
 * @param request AddUserToGroupRequest
 * @return AddUserToGroupResponse
 */
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

/**
 * If single sign-on (SSO) logon is disabled, you can clear the configurations of a SAML IdP. If SSO logon is enabled, you cannot clear the configurations.
 * This topic provides an example on how to clear the configurations of the SAML IdP within the directory `d-00fc2p61****`.
 *
 * @param request ClearExternalSAMLIdentityProviderRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ClearExternalSAMLIdentityProviderResponse
 */
func (client *Client) ClearExternalSAMLIdentityProviderWithOptions(request *ClearExternalSAMLIdentityProviderRequest, runtime *util.RuntimeOptions) (_result *ClearExternalSAMLIdentityProviderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ClearExternalSAMLIdentityProvider"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ClearExternalSAMLIdentityProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If single sign-on (SSO) logon is disabled, you can clear the configurations of a SAML IdP. If SSO logon is enabled, you cannot clear the configurations.
 * This topic provides an example on how to clear the configurations of the SAML IdP within the directory `d-00fc2p61****`.
 *
 * @param request ClearExternalSAMLIdentityProviderRequest
 * @return ClearExternalSAMLIdentityProviderResponse
 */
func (client *Client) ClearExternalSAMLIdentityProvider(request *ClearExternalSAMLIdentityProviderRequest) (_result *ClearExternalSAMLIdentityProviderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ClearExternalSAMLIdentityProviderResponse{}
	_body, _err := client.ClearExternalSAMLIdentityProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * When you call this operation, an asynchronous task is created. You can call the [GetTask](~~340670~~) operation to query the progress of the task based on the value of the `TaskId` response parameter.
 * For more information about how to assign permissions on an account in your resource directory, see [Overview of multi-account authorization](~~266726~~).
 * This topic provides an example on how to assign access permissions on the account `114240524784****` in your resource directory to the CloudSSO user `u-00q8wbq42wiltcrk****` by using the access configuration `ac-00jhtfl8thteu6uj****`. After the call is successful, the CloudSSO user can access resources within the account in the resource directory.
 *
 * @param request CreateAccessAssignmentRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateAccessAssignmentResponse
 */
func (client *Client) CreateAccessAssignmentWithOptions(request *CreateAccessAssignmentRequest, runtime *util.RuntimeOptions) (_result *CreateAccessAssignmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessConfigurationId)) {
		query["AccessConfigurationId"] = request.AccessConfigurationId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalId)) {
		query["PrincipalId"] = request.PrincipalId
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalType)) {
		query["PrincipalType"] = request.PrincipalType
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["TargetId"] = request.TargetId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccessAssignment"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAccessAssignmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * When you call this operation, an asynchronous task is created. You can call the [GetTask](~~340670~~) operation to query the progress of the task based on the value of the `TaskId` response parameter.
 * For more information about how to assign permissions on an account in your resource directory, see [Overview of multi-account authorization](~~266726~~).
 * This topic provides an example on how to assign access permissions on the account `114240524784****` in your resource directory to the CloudSSO user `u-00q8wbq42wiltcrk****` by using the access configuration `ac-00jhtfl8thteu6uj****`. After the call is successful, the CloudSSO user can access resources within the account in the resource directory.
 *
 * @param request CreateAccessAssignmentRequest
 * @return CreateAccessAssignmentResponse
 */
func (client *Client) CreateAccessAssignment(request *CreateAccessAssignmentRequest) (_result *CreateAccessAssignmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAccessAssignmentResponse{}
	_body, _err := client.CreateAccessAssignmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * For more information about access configurations, see [Overview of access configurations](~~266737~~).
 * This topic provides an example on how to create an access configuration named `ECS-Admin`.
 *
 * @param request CreateAccessConfigurationRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateAccessConfigurationResponse
 */
func (client *Client) CreateAccessConfigurationWithOptions(request *CreateAccessConfigurationRequest, runtime *util.RuntimeOptions) (_result *CreateAccessConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessConfigurationName)) {
		query["AccessConfigurationName"] = request.AccessConfigurationName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.RelayState)) {
		query["RelayState"] = request.RelayState
	}

	if !tea.BoolValue(util.IsUnset(request.SessionDuration)) {
		query["SessionDuration"] = request.SessionDuration
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccessConfiguration"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAccessConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * For more information about access configurations, see [Overview of access configurations](~~266737~~).
 * This topic provides an example on how to create an access configuration named `ECS-Admin`.
 *
 * @param request CreateAccessConfigurationRequest
 * @return CreateAccessConfigurationResponse
 */
func (client *Client) CreateAccessConfiguration(request *CreateAccessConfigurationRequest) (_result *CreateAccessConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAccessConfigurationResponse{}
	_body, _err := client.CreateAccessConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * A directory is a CloudSSO instance. Before you can use CloudSSO, you must create a directory. The directory is used to manage all CloudSSO resources.
 * To create a directory, you must select a region. Alibaba Cloud stores data in the directory only in the region that you select. However, you can deploy Alibaba Cloud resources including Elastic Compute Service (ECS) instances and ApsaraDB RDS instances in other regions. You can also use your cloud account for logons and access the Alibaba Cloud resources in other regions. You can select a region to create a directory based on your security compliance requirements and the geographic location of specific users. If you do not have strict security compliance requirements, we recommend that you select a region that is the closest to the geographical location of the specific users. This way, access to cloud resources is accelerated. You can create the CloudSSO directory in the China (Shanghai), China (Hong Kong), US (Silicon Valley), or Germany (Frankfurt) region.
 * This topic provides an example on how to create a directory named `example` in the China (Shanghai) region.
 * ## Limits
 * - You can create only one directory for a management account.
 * - If you want to change the region of a directory, you must delete the directory and then create a directory in a different region.
 *
 * @param request CreateDirectoryRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateDirectoryResponse
 */
func (client *Client) CreateDirectoryWithOptions(request *CreateDirectoryRequest, runtime *util.RuntimeOptions) (_result *CreateDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryName)) {
		query["DirectoryName"] = request.DirectoryName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDirectory"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * A directory is a CloudSSO instance. Before you can use CloudSSO, you must create a directory. The directory is used to manage all CloudSSO resources.
 * To create a directory, you must select a region. Alibaba Cloud stores data in the directory only in the region that you select. However, you can deploy Alibaba Cloud resources including Elastic Compute Service (ECS) instances and ApsaraDB RDS instances in other regions. You can also use your cloud account for logons and access the Alibaba Cloud resources in other regions. You can select a region to create a directory based on your security compliance requirements and the geographic location of specific users. If you do not have strict security compliance requirements, we recommend that you select a region that is the closest to the geographical location of the specific users. This way, access to cloud resources is accelerated. You can create the CloudSSO directory in the China (Shanghai), China (Hong Kong), US (Silicon Valley), or Germany (Frankfurt) region.
 * This topic provides an example on how to create a directory named `example` in the China (Shanghai) region.
 * ## Limits
 * - You can create only one directory for a management account.
 * - If you want to change the region of a directory, you must delete the directory and then create a directory in a different region.
 *
 * @param request CreateDirectoryRequest
 * @return CreateDirectoryResponse
 */
func (client *Client) CreateDirectory(request *CreateDirectoryRequest) (_result *CreateDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDirectoryResponse{}
	_body, _err := client.CreateDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to create a group named `TestGroup`.
 *
 * @param request CreateGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateGroupResponse
 */
func (client *Client) CreateGroupWithOptions(request *CreateGroupRequest, runtime *util.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGroup"),
		Version:     tea.String("2021-05-15"),
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

/**
 * This topic provides an example on how to create a group named `TestGroup`.
 *
 * @param request CreateGroupRequest
 * @return CreateGroupResponse
 */
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

/**
 * SCIM credentials are required for SCIM synchronization. You can create up to two SCIM credentials.
 * This topic provides an example on how to create a SCIM credential within the directory `d-00fc2p61****`.
 *
 * @param request CreateSCIMServerCredentialRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateSCIMServerCredentialResponse
 */
func (client *Client) CreateSCIMServerCredentialWithOptions(request *CreateSCIMServerCredentialRequest, runtime *util.RuntimeOptions) (_result *CreateSCIMServerCredentialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSCIMServerCredential"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSCIMServerCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * SCIM credentials are required for SCIM synchronization. You can create up to two SCIM credentials.
 * This topic provides an example on how to create a SCIM credential within the directory `d-00fc2p61****`.
 *
 * @param request CreateSCIMServerCredentialRequest
 * @return CreateSCIMServerCredentialResponse
 */
func (client *Client) CreateSCIMServerCredential(request *CreateSCIMServerCredentialRequest) (_result *CreateSCIMServerCredentialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSCIMServerCredentialResponse{}
	_body, _err := client.CreateSCIMServerCredentialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to create a user named `Alice`.
 *
 * @param request CreateUserRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateUserResponse
 */
func (client *Client) CreateUserWithOptions(request *CreateUserRequest, runtime *util.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.FirstName)) {
		query["FirstName"] = request.FirstName
	}

	if !tea.BoolValue(util.IsUnset(request.LastName)) {
		query["LastName"] = request.LastName
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUser"),
		Version:     tea.String("2021-05-15"),
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

/**
 * This topic provides an example on how to create a user named `Alice`.
 *
 * @param request CreateUserRequest
 * @return CreateUserResponse
 */
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

func (client *Client) CreateUserProvisioningWithOptions(request *CreateUserProvisioningRequest, runtime *util.RuntimeOptions) (_result *CreateUserProvisioningResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeletionStrategy)) {
		query["DeletionStrategy"] = request.DeletionStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.DuplicationStrategy)) {
		query["DuplicationStrategy"] = request.DuplicationStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalId)) {
		query["PrincipalId"] = request.PrincipalId
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalType)) {
		query["PrincipalType"] = request.PrincipalType
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["TargetId"] = request.TargetId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUserProvisioning"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserProvisioningResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUserProvisioning(request *CreateUserProvisioningRequest) (_result *CreateUserProvisioningResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUserProvisioningResponse{}
	_body, _err := client.CreateUserProvisioningWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * When you call this operation, an asynchronous task is created. You can call the [GetTask](~~340670~~) operation to query the progress of the task based on the value of the `TaskId` response parameter.
 * This topic provides an example on how to remove the access permissions on the account `114240524784****` in the resource directory from the CloudSSO user `u-00q8wbq42wiltcrk****`. The access permissions are assigned by using the access configuration `ac-00jhtfl8thteu6uj****`.
 *
 * @param request DeleteAccessAssignmentRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteAccessAssignmentResponse
 */
func (client *Client) DeleteAccessAssignmentWithOptions(request *DeleteAccessAssignmentRequest, runtime *util.RuntimeOptions) (_result *DeleteAccessAssignmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessConfigurationId)) {
		query["AccessConfigurationId"] = request.AccessConfigurationId
	}

	if !tea.BoolValue(util.IsUnset(request.DeprovisionStrategy)) {
		query["DeprovisionStrategy"] = request.DeprovisionStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalId)) {
		query["PrincipalId"] = request.PrincipalId
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalType)) {
		query["PrincipalType"] = request.PrincipalType
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["TargetId"] = request.TargetId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccessAssignment"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAccessAssignmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * When you call this operation, an asynchronous task is created. You can call the [GetTask](~~340670~~) operation to query the progress of the task based on the value of the `TaskId` response parameter.
 * This topic provides an example on how to remove the access permissions on the account `114240524784****` in the resource directory from the CloudSSO user `u-00q8wbq42wiltcrk****`. The access permissions are assigned by using the access configuration `ac-00jhtfl8thteu6uj****`.
 *
 * @param request DeleteAccessAssignmentRequest
 * @return DeleteAccessAssignmentResponse
 */
func (client *Client) DeleteAccessAssignment(request *DeleteAccessAssignmentRequest) (_result *DeleteAccessAssignmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAccessAssignmentResponse{}
	_body, _err := client.DeleteAccessAssignmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to delete the access configuration whose ID is `ac-001j9mcm3k7335bc****`.
 * ## Prerequisites
 * The access configuration that you want to delete is de-provisioned from the accounts in your resource directory. For more information, see [DeprovisionAccessConfiguration](~~338352~~).
 *
 * @param request DeleteAccessConfigurationRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteAccessConfigurationResponse
 */
func (client *Client) DeleteAccessConfigurationWithOptions(request *DeleteAccessConfigurationRequest, runtime *util.RuntimeOptions) (_result *DeleteAccessConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessConfigurationId)) {
		query["AccessConfigurationId"] = request.AccessConfigurationId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.ForceRemovePermissionPolicies)) {
		query["ForceRemovePermissionPolicies"] = request.ForceRemovePermissionPolicies
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccessConfiguration"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAccessConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to delete the access configuration whose ID is `ac-001j9mcm3k7335bc****`.
 * ## Prerequisites
 * The access configuration that you want to delete is de-provisioned from the accounts in your resource directory. For more information, see [DeprovisionAccessConfiguration](~~338352~~).
 *
 * @param request DeleteAccessConfigurationRequest
 * @return DeleteAccessConfigurationResponse
 */
func (client *Client) DeleteAccessConfiguration(request *DeleteAccessConfigurationRequest) (_result *DeleteAccessConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAccessConfigurationResponse{}
	_body, _err := client.DeleteAccessConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to delete a directory whose ID is `d-00fc2p61****`.
 * ## Prerequisites
 * No resources are contained in the directory that you want to delete.
 * *   Access permissions on the accounts in your resource directory are removed from all users and groups. For more information, see [DeleteAccessAssignment](~~338350~~).
 * *   Users are deleted. For more information, see [DeleteUser](~~341671~~).
 * *   Groups are deleted. For more information, see [DeleteGroup](~~341821~~).
 * *   Access configurations are deleted. For more information, see [DeleteAccessConfiguration](~~336907~~).
 * *   System for Cross-domain Identity Management (SCIM) credentials are deleted. For more information, see [DeleteSCIMServerCredential](~~341842~~).
 * *   SSO logon configurations are deleted. For more information, see [ClearExternalSAMLIdentityProvider](~~341573~~).
 *
 * @param request DeleteDirectoryRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteDirectoryResponse
 */
func (client *Client) DeleteDirectoryWithOptions(request *DeleteDirectoryRequest, runtime *util.RuntimeOptions) (_result *DeleteDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDirectory"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to delete a directory whose ID is `d-00fc2p61****`.
 * ## Prerequisites
 * No resources are contained in the directory that you want to delete.
 * *   Access permissions on the accounts in your resource directory are removed from all users and groups. For more information, see [DeleteAccessAssignment](~~338350~~).
 * *   Users are deleted. For more information, see [DeleteUser](~~341671~~).
 * *   Groups are deleted. For more information, see [DeleteGroup](~~341821~~).
 * *   Access configurations are deleted. For more information, see [DeleteAccessConfiguration](~~336907~~).
 * *   System for Cross-domain Identity Management (SCIM) credentials are deleted. For more information, see [DeleteSCIMServerCredential](~~341842~~).
 * *   SSO logon configurations are deleted. For more information, see [ClearExternalSAMLIdentityProvider](~~341573~~).
 *
 * @param request DeleteDirectoryRequest
 * @return DeleteDirectoryResponse
 */
func (client *Client) DeleteDirectory(request *DeleteDirectoryRequest) (_result *DeleteDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDirectoryResponse{}
	_body, _err := client.DeleteDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * If System for Cross-domain Identity Management (SCIM) synchronization is enabled, you cannot delete a group that is synchronized by using SCIM.
 * ## Prerequisites
 * The group that you want to delete is not associated with the following resources. If the group is associated with the resources, the deletion fails.
 * *   Users: You must remove users from the group. For more information, see [RemoveUserFromGroup](~~335116~~).
 * *   Access permissions: You must remove the access permissions on the accounts in your resource directory from the group. For more information, see [DeleteAccessAssignment](~~338350~~).
 *
 * @param request DeleteGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteGroupResponse
 */
func (client *Client) DeleteGroupWithOptions(request *DeleteGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGroup"),
		Version:     tea.String("2021-05-15"),
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

/**
 * If System for Cross-domain Identity Management (SCIM) synchronization is enabled, you cannot delete a group that is synchronized by using SCIM.
 * ## Prerequisites
 * The group that you want to delete is not associated with the following resources. If the group is associated with the resources, the deletion fails.
 * *   Users: You must remove users from the group. For more information, see [RemoveUserFromGroup](~~335116~~).
 * *   Access permissions: You must remove the access permissions on the accounts in your resource directory from the group. For more information, see [DeleteAccessAssignment](~~338350~~).
 *
 * @param request DeleteGroupRequest
 * @return DeleteGroupResponse
 */
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

/**
 * This topic provides an example on how to unbind the MFA device `mfa-00ujhet8pycljj7j****` from the user `u-00q8wbq42wiltcrk****`.
 *
 * @param request DeleteMFADeviceForUserRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteMFADeviceForUserResponse
 */
func (client *Client) DeleteMFADeviceForUserWithOptions(request *DeleteMFADeviceForUserRequest, runtime *util.RuntimeOptions) (_result *DeleteMFADeviceForUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.MFADeviceId)) {
		query["MFADeviceId"] = request.MFADeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMFADeviceForUser"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMFADeviceForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to unbind the MFA device `mfa-00ujhet8pycljj7j****` from the user `u-00q8wbq42wiltcrk****`.
 *
 * @param request DeleteMFADeviceForUserRequest
 * @return DeleteMFADeviceForUserResponse
 */
func (client *Client) DeleteMFADeviceForUser(request *DeleteMFADeviceForUserRequest) (_result *DeleteMFADeviceForUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMFADeviceForUserResponse{}
	_body, _err := client.DeleteMFADeviceForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * After a SCIM credential is deleted, the synchronization task that uses the SCIM credential fails.
 * This topic provides an example on how to delete the SCIM credential whose ID is `scimcred-004whl0kvfwcypbi****`.
 *
 * @param request DeleteSCIMServerCredentialRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteSCIMServerCredentialResponse
 */
func (client *Client) DeleteSCIMServerCredentialWithOptions(request *DeleteSCIMServerCredentialRequest, runtime *util.RuntimeOptions) (_result *DeleteSCIMServerCredentialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CredentialId)) {
		query["CredentialId"] = request.CredentialId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSCIMServerCredential"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSCIMServerCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * After a SCIM credential is deleted, the synchronization task that uses the SCIM credential fails.
 * This topic provides an example on how to delete the SCIM credential whose ID is `scimcred-004whl0kvfwcypbi****`.
 *
 * @param request DeleteSCIMServerCredentialRequest
 * @return DeleteSCIMServerCredentialResponse
 */
func (client *Client) DeleteSCIMServerCredential(request *DeleteSCIMServerCredentialRequest) (_result *DeleteSCIMServerCredentialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSCIMServerCredentialResponse{}
	_body, _err := client.DeleteSCIMServerCredentialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * If System for Cross-domain Identity Management (SCIM) synchronization is enabled, you cannot delete a user that is synchronized by using SCIM.
 * ## Prerequisites
 * The user that you want to delete is not associated with the following resources. If the user is associated with the resources, the deletion fails.
 * *   Multi-factor authentication (MFA) devices: You must unbind the MFA devices from the user. For more information, see [DeleteMFADeviceForUser](~~341675~~).
 * *   Access permissions: You must remove the access permissions on the accounts in your resource directory from the user. For more information, see [DeleteAccessAssignment](~~338350~~).
 * *   Groups: You must remove the user from groups. For more information, see [RemoveUserFromGroup](~~335116~~).
 *
 * @param request DeleteUserRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteUserResponse
 */
func (client *Client) DeleteUserWithOptions(request *DeleteUserRequest, runtime *util.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUser"),
		Version:     tea.String("2021-05-15"),
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

/**
 * If System for Cross-domain Identity Management (SCIM) synchronization is enabled, you cannot delete a user that is synchronized by using SCIM.
 * ## Prerequisites
 * The user that you want to delete is not associated with the following resources. If the user is associated with the resources, the deletion fails.
 * *   Multi-factor authentication (MFA) devices: You must unbind the MFA devices from the user. For more information, see [DeleteMFADeviceForUser](~~341675~~).
 * *   Access permissions: You must remove the access permissions on the accounts in your resource directory from the user. For more information, see [DeleteAccessAssignment](~~338350~~).
 * *   Groups: You must remove the user from groups. For more information, see [RemoveUserFromGroup](~~335116~~).
 *
 * @param request DeleteUserRequest
 * @return DeleteUserResponse
 */
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

func (client *Client) DeleteUserProvisioningWithOptions(request *DeleteUserProvisioningRequest, runtime *util.RuntimeOptions) (_result *DeleteUserProvisioningResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeletionStrategy)) {
		query["DeletionStrategy"] = request.DeletionStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.UserProvisioningId)) {
		query["UserProvisioningId"] = request.UserProvisioningId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUserProvisioning"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserProvisioningResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUserProvisioning(request *DeleteUserProvisioningRequest) (_result *DeleteUserProvisioningResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserProvisioningResponse{}
	_body, _err := client.DeleteUserProvisioningWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserProvisioningEventWithOptions(request *DeleteUserProvisioningEventRequest, runtime *util.RuntimeOptions) (_result *DeleteUserProvisioningEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.EventId)) {
		query["EventId"] = request.EventId
	}

	if !tea.BoolValue(util.IsUnset(request.UserProvisioningId)) {
		query["UserProvisioningId"] = request.UserProvisioningId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUserProvisioningEvent"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserProvisioningEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUserProvisioningEvent(request *DeleteUserProvisioningEventRequest) (_result *DeleteUserProvisioningEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserProvisioningEventResponse{}
	_body, _err := client.DeleteUserProvisioningEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * When you call this operation, an asynchronous task is automatically created. You can call the [GetTask](~~340670~~) operation to query the progress of the task based on the value of the `TaskId` response parameter.
 * This topic provides an example on how to de-provision the access configuration `ac-00jhtfl8thteu6uj****` from the account `114240524784****` in your resource directory.
 *
 * @param request DeprovisionAccessConfigurationRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeprovisionAccessConfigurationResponse
 */
func (client *Client) DeprovisionAccessConfigurationWithOptions(request *DeprovisionAccessConfigurationRequest, runtime *util.RuntimeOptions) (_result *DeprovisionAccessConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessConfigurationId)) {
		query["AccessConfigurationId"] = request.AccessConfigurationId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["TargetId"] = request.TargetId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeprovisionAccessConfiguration"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeprovisionAccessConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * When you call this operation, an asynchronous task is automatically created. You can call the [GetTask](~~340670~~) operation to query the progress of the task based on the value of the `TaskId` response parameter.
 * This topic provides an example on how to de-provision the access configuration `ac-00jhtfl8thteu6uj****` from the account `114240524784****` in your resource directory.
 *
 * @param request DeprovisionAccessConfigurationRequest
 * @return DeprovisionAccessConfigurationResponse
 */
func (client *Client) DeprovisionAccessConfiguration(request *DeprovisionAccessConfigurationRequest) (_result *DeprovisionAccessConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeprovisionAccessConfigurationResponse{}
	_body, _err := client.DeprovisionAccessConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * If your CloudSSO has no directory, you can disable CloudSSO based on your business requirements. After you disable CloudSSO, you can enable it at any time.
 *
 * @param request DisableServiceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DisableServiceResponse
 */
func (client *Client) DisableServiceWithOptions(runtime *util.RuntimeOptions) (_result *DisableServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DisableService"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If your CloudSSO has no directory, you can disable CloudSSO based on your business requirements. After you disable CloudSSO, you can enable it at any time.
 *
 * @return DisableServiceResponse
 */
func (client *Client) DisableService() (_result *DisableServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableServiceResponse{}
	_body, _err := client.DisableServiceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation only if your account belongs to the management account that is used to enable a resource directory and has permissions to enable CloudSSO. For more information, see [Enable CloudSSO](~~262819~~).
 * If you call this operation, you agree to the [Alibaba Cloud International Website Product Terms of Service](https://www.alibabacloud.com/help/doc-detail/42416.htm).
 *
 * @param request EnableServiceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return EnableServiceResponse
 */
func (client *Client) EnableServiceWithOptions(runtime *util.RuntimeOptions) (_result *EnableServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("EnableService"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation only if your account belongs to the management account that is used to enable a resource directory and has permissions to enable CloudSSO. For more information, see [Enable CloudSSO](~~262819~~).
 * If you call this operation, you agree to the [Alibaba Cloud International Website Product Terms of Service](https://www.alibabacloud.com/help/doc-detail/42416.htm).
 *
 * @return EnableServiceResponse
 */
func (client *Client) EnableService() (_result *EnableServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableServiceResponse{}
	_body, _err := client.EnableServiceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to query the information about the access configuration whose ID is `ac-00ccule7tadaijxc****`.
 *
 * @param request GetAccessConfigurationRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetAccessConfigurationResponse
 */
func (client *Client) GetAccessConfigurationWithOptions(request *GetAccessConfigurationRequest, runtime *util.RuntimeOptions) (_result *GetAccessConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessConfigurationId)) {
		query["AccessConfigurationId"] = request.AccessConfigurationId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAccessConfiguration"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccessConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to query the information about the access configuration whose ID is `ac-00ccule7tadaijxc****`.
 *
 * @param request GetAccessConfigurationRequest
 * @return GetAccessConfigurationResponse
 */
func (client *Client) GetAccessConfiguration(request *GetAccessConfigurationRequest) (_result *GetAccessConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccessConfigurationResponse{}
	_body, _err := client.GetAccessConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to query information about the directory whose ID is `d-00fc2p61****`.
 *
 * @param request GetDirectoryRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetDirectoryResponse
 */
func (client *Client) GetDirectoryWithOptions(request *GetDirectoryRequest, runtime *util.RuntimeOptions) (_result *GetDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDirectory"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to query information about the directory whose ID is `d-00fc2p61****`.
 *
 * @param request GetDirectoryRequest
 * @return GetDirectoryResponse
 */
func (client *Client) GetDirectory(request *GetDirectoryRequest) (_result *GetDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDirectoryResponse{}
	_body, _err := client.GetDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * During SAML 2.0-based single sign-on (SSO) logon, CloudSSO is an SP, and the identity management system of an enterprise is an identity provider (IdP).
 * This topic provides an example on how to query the information about the SP within the directory `d-00fc2p61****`.
 *
 * @param request GetDirectorySAMLServiceProviderInfoRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetDirectorySAMLServiceProviderInfoResponse
 */
func (client *Client) GetDirectorySAMLServiceProviderInfoWithOptions(request *GetDirectorySAMLServiceProviderInfoRequest, runtime *util.RuntimeOptions) (_result *GetDirectorySAMLServiceProviderInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDirectorySAMLServiceProviderInfo"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDirectorySAMLServiceProviderInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * During SAML 2.0-based single sign-on (SSO) logon, CloudSSO is an SP, and the identity management system of an enterprise is an identity provider (IdP).
 * This topic provides an example on how to query the information about the SP within the directory `d-00fc2p61****`.
 *
 * @param request GetDirectorySAMLServiceProviderInfoRequest
 * @return GetDirectorySAMLServiceProviderInfoResponse
 */
func (client *Client) GetDirectorySAMLServiceProviderInfo(request *GetDirectorySAMLServiceProviderInfoRequest) (_result *GetDirectorySAMLServiceProviderInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDirectorySAMLServiceProviderInfoResponse{}
	_body, _err := client.GetDirectorySAMLServiceProviderInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### [](#)
 * This topic provides an example on how to query the statistics of a directory whose ID is `d-00fc2p61****`. The statistics include the number of users, quota for users, number of groups, quota for groups, number of access configurations, quota for access configurations, quota for system policies that can be configured for an access configuration, number of access permissions that are assigned, number of System for Cross-domain Identity Management (SCIM) credentials, number of asynchronous tasks, status of single sign-on (SSO), and status of SCIM synchronization.
 * ### [](#qps)Limit
 * You can call this operation up to 100 times per second per account. This operation is globally limited to 100 times per second across all accounts. If the number of the calls per second exceeds a limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limits when you call this operation.
 *
 * @param request GetDirectoryStatisticsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetDirectoryStatisticsResponse
 */
func (client *Client) GetDirectoryStatisticsWithOptions(request *GetDirectoryStatisticsRequest, runtime *util.RuntimeOptions) (_result *GetDirectoryStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDirectoryStatistics"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDirectoryStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### [](#)
 * This topic provides an example on how to query the statistics of a directory whose ID is `d-00fc2p61****`. The statistics include the number of users, quota for users, number of groups, quota for groups, number of access configurations, quota for access configurations, quota for system policies that can be configured for an access configuration, number of access permissions that are assigned, number of System for Cross-domain Identity Management (SCIM) credentials, number of asynchronous tasks, status of single sign-on (SSO), and status of SCIM synchronization.
 * ### [](#qps)Limit
 * You can call this operation up to 100 times per second per account. This operation is globally limited to 100 times per second across all accounts. If the number of the calls per second exceeds a limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limits when you call this operation.
 *
 * @param request GetDirectoryStatisticsRequest
 * @return GetDirectoryStatisticsResponse
 */
func (client *Client) GetDirectoryStatistics(request *GetDirectoryStatisticsRequest) (_result *GetDirectoryStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDirectoryStatisticsResponse{}
	_body, _err := client.GetDirectoryStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to query the configurations of the SAML IdP within the directory `d-00fc2p61****`.
 *
 * @param request GetExternalSAMLIdentityProviderRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetExternalSAMLIdentityProviderResponse
 */
func (client *Client) GetExternalSAMLIdentityProviderWithOptions(request *GetExternalSAMLIdentityProviderRequest, runtime *util.RuntimeOptions) (_result *GetExternalSAMLIdentityProviderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetExternalSAMLIdentityProvider"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetExternalSAMLIdentityProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to query the configurations of the SAML IdP within the directory `d-00fc2p61****`.
 *
 * @param request GetExternalSAMLIdentityProviderRequest
 * @return GetExternalSAMLIdentityProviderResponse
 */
func (client *Client) GetExternalSAMLIdentityProvider(request *GetExternalSAMLIdentityProviderRequest) (_result *GetExternalSAMLIdentityProviderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetExternalSAMLIdentityProviderResponse{}
	_body, _err := client.GetExternalSAMLIdentityProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to query the information about the group `g-00jqzghi2n3o5hkh****` in the directory `d-00fc2p61****`.
 *
 * @param request GetGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetGroupResponse
 */
func (client *Client) GetGroupWithOptions(request *GetGroupRequest, runtime *util.RuntimeOptions) (_result *GetGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGroup"),
		Version:     tea.String("2021-05-15"),
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

/**
 * This topic provides an example on how to query the information about the group `g-00jqzghi2n3o5hkh****` in the directory `d-00fc2p61****`.
 *
 * @param request GetGroupRequest
 * @return GetGroupResponse
 */
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

func (client *Client) GetLoginPreferenceWithOptions(request *GetLoginPreferenceRequest, runtime *util.RuntimeOptions) (_result *GetLoginPreferenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLoginPreference"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLoginPreferenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLoginPreference(request *GetLoginPreferenceRequest) (_result *GetLoginPreferenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLoginPreferenceResponse{}
	_body, _err := client.GetLoginPreferenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * If you enable username-password logon for CloudSSO users, you can also configure MFA for the users.
 * This topic provides an example on how to query the MFA setting of all CloudSSO users that belong to the directory named `00q8wbq42wiltcrk****`.
 *
 * @param request GetMFAAuthenticationSettingInfoRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetMFAAuthenticationSettingInfoResponse
 */
func (client *Client) GetMFAAuthenticationSettingInfoWithOptions(request *GetMFAAuthenticationSettingInfoRequest, runtime *util.RuntimeOptions) (_result *GetMFAAuthenticationSettingInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMFAAuthenticationSettingInfo"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMFAAuthenticationSettingInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If you enable username-password logon for CloudSSO users, you can also configure MFA for the users.
 * This topic provides an example on how to query the MFA setting of all CloudSSO users that belong to the directory named `00q8wbq42wiltcrk****`.
 *
 * @param request GetMFAAuthenticationSettingInfoRequest
 * @return GetMFAAuthenticationSettingInfoResponse
 */
func (client *Client) GetMFAAuthenticationSettingInfo(request *GetMFAAuthenticationSettingInfoRequest) (_result *GetMFAAuthenticationSettingInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMFAAuthenticationSettingInfoResponse{}
	_body, _err := client.GetMFAAuthenticationSettingInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * > This operation is no longer maintained and updated. You can call the [GetMFAAuthenticationSettingInfo](~~611286~~) operation to query more detailed information.
 * This topic provides an example on how to query the MFA setting of the users that belong to the directory named `d-00fc2p61****`. The returned result shows that MFA is enabled for all the users.
 *
 * @param request GetMFAAuthenticationSettingsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetMFAAuthenticationSettingsResponse
 */
func (client *Client) GetMFAAuthenticationSettingsWithOptions(request *GetMFAAuthenticationSettingsRequest, runtime *util.RuntimeOptions) (_result *GetMFAAuthenticationSettingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMFAAuthenticationSettings"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMFAAuthenticationSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * > This operation is no longer maintained and updated. You can call the [GetMFAAuthenticationSettingInfo](~~611286~~) operation to query more detailed information.
 * This topic provides an example on how to query the MFA setting of the users that belong to the directory named `d-00fc2p61****`. The returned result shows that MFA is enabled for all the users.
 *
 * @param request GetMFAAuthenticationSettingsRequest
 * @return GetMFAAuthenticationSettingsResponse
 */
func (client *Client) GetMFAAuthenticationSettings(request *GetMFAAuthenticationSettingsRequest) (_result *GetMFAAuthenticationSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMFAAuthenticationSettingsResponse{}
	_body, _err := client.GetMFAAuthenticationSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to check whether MFA is enabled for users in the directory whose ID is `00fc2p61****`. The returned result shows that MFA is in the Enabled state.
 *
 * @param request GetMFAAuthenticationStatusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetMFAAuthenticationStatusResponse
 */
func (client *Client) GetMFAAuthenticationStatusWithOptions(request *GetMFAAuthenticationStatusRequest, runtime *util.RuntimeOptions) (_result *GetMFAAuthenticationStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMFAAuthenticationStatus"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMFAAuthenticationStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to check whether MFA is enabled for users in the directory whose ID is `00fc2p61****`. The returned result shows that MFA is in the Enabled state.
 *
 * @param request GetMFAAuthenticationStatusRequest
 * @return GetMFAAuthenticationStatusResponse
 */
func (client *Client) GetMFAAuthenticationStatus(request *GetMFAAuthenticationStatusRequest) (_result *GetMFAAuthenticationStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMFAAuthenticationStatusResponse{}
	_body, _err := client.GetMFAAuthenticationStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPasswordPolicyWithOptions(request *GetPasswordPolicyRequest, runtime *util.RuntimeOptions) (_result *GetPasswordPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPasswordPolicy"),
		Version:     tea.String("2021-05-15"),
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

func (client *Client) GetPasswordPolicy(request *GetPasswordPolicyRequest) (_result *GetPasswordPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPasswordPolicyResponse{}
	_body, _err := client.GetPasswordPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to query the status of SCIM synchronization within the directory `d-00fc2p61****`. The returned result shows that SCIM synchronization is in the Enabled state.
 *
 * @param request GetSCIMSynchronizationStatusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetSCIMSynchronizationStatusResponse
 */
func (client *Client) GetSCIMSynchronizationStatusWithOptions(request *GetSCIMSynchronizationStatusRequest, runtime *util.RuntimeOptions) (_result *GetSCIMSynchronizationStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSCIMSynchronizationStatus"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSCIMSynchronizationStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to query the status of SCIM synchronization within the directory `d-00fc2p61****`. The returned result shows that SCIM synchronization is in the Enabled state.
 *
 * @param request GetSCIMSynchronizationStatusRequest
 * @return GetSCIMSynchronizationStatusResponse
 */
func (client *Client) GetSCIMSynchronizationStatus(request *GetSCIMSynchronizationStatusRequest) (_result *GetSCIMSynchronizationStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSCIMSynchronizationStatusResponse{}
	_body, _err := client.GetSCIMSynchronizationStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetServiceStatusWithOptions(runtime *util.RuntimeOptions) (_result *GetServiceStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetServiceStatus"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetServiceStatus() (_result *GetServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServiceStatusResponse{}
	_body, _err := client.GetServiceStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to query the information about the task whose ID is `t-shfqw1u1edszvxw5****`.
 *
 * @param request GetTaskRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetTaskResponse
 */
func (client *Client) GetTaskWithOptions(request *GetTaskRequest, runtime *util.RuntimeOptions) (_result *GetTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTask"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to query the information about the task whose ID is `t-shfqw1u1edszvxw5****`.
 *
 * @param request GetTaskRequest
 * @return GetTaskResponse
 */
func (client *Client) GetTask(request *GetTaskRequest) (_result *GetTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTaskResponse{}
	_body, _err := client.GetTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the GetTaskStatus operation to query the status of an asynchronous task. If you want to query more information about an asynchronous task, call the [GetTask](~~340670~~) operation.
 * This topic provides an example on how to query the information about the task whose ID is `t-shfqw1u1edszvxw5****`.
 *
 * @param request GetTaskStatusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetTaskStatusResponse
 */
func (client *Client) GetTaskStatusWithOptions(request *GetTaskStatusRequest, runtime *util.RuntimeOptions) (_result *GetTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTaskStatus"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the GetTaskStatus operation to query the status of an asynchronous task. If you want to query more information about an asynchronous task, call the [GetTask](~~340670~~) operation.
 * This topic provides an example on how to query the information about the task whose ID is `t-shfqw1u1edszvxw5****`.
 *
 * @param request GetTaskStatusRequest
 * @return GetTaskStatusResponse
 */
func (client *Client) GetTaskStatus(request *GetTaskStatusRequest) (_result *GetTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTaskStatusResponse{}
	_body, _err := client.GetTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to query information about the user whose ID is `u-00q8wbq42wiltcrk****` in the `d-00fc2p61****` directory.
 *
 * @param request GetUserRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetUserResponse
 */
func (client *Client) GetUserWithOptions(request *GetUserRequest, runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUser"),
		Version:     tea.String("2021-05-15"),
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

/**
 * This topic provides an example on how to query information about the user whose ID is `u-00q8wbq42wiltcrk****` in the `d-00fc2p61****` directory.
 *
 * @param request GetUserRequest
 * @return GetUserResponse
 */
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

/**
 * This topic provides an example on how to query the MFA setting of the user named `u-00q8wbq42wiltcrk****`. The returned result shows that MFA is enabled for the user.
 *
 * @param request GetUserMFAAuthenticationSettingsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetUserMFAAuthenticationSettingsResponse
 */
func (client *Client) GetUserMFAAuthenticationSettingsWithOptions(request *GetUserMFAAuthenticationSettingsRequest, runtime *util.RuntimeOptions) (_result *GetUserMFAAuthenticationSettingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserMFAAuthenticationSettings"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserMFAAuthenticationSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to query the MFA setting of the user named `u-00q8wbq42wiltcrk****`. The returned result shows that MFA is enabled for the user.
 *
 * @param request GetUserMFAAuthenticationSettingsRequest
 * @return GetUserMFAAuthenticationSettingsResponse
 */
func (client *Client) GetUserMFAAuthenticationSettings(request *GetUserMFAAuthenticationSettingsRequest) (_result *GetUserMFAAuthenticationSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserMFAAuthenticationSettingsResponse{}
	_body, _err := client.GetUserMFAAuthenticationSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserProvisioningWithOptions(request *GetUserProvisioningRequest, runtime *util.RuntimeOptions) (_result *GetUserProvisioningResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.UserProvisioningId)) {
		query["UserProvisioningId"] = request.UserProvisioningId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserProvisioning"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserProvisioningResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserProvisioning(request *GetUserProvisioningRequest) (_result *GetUserProvisioningResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserProvisioningResponse{}
	_body, _err := client.GetUserProvisioningWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserProvisioningConfigurationWithOptions(request *GetUserProvisioningConfigurationRequest, runtime *util.RuntimeOptions) (_result *GetUserProvisioningConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserProvisioningConfiguration"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserProvisioningConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserProvisioningConfiguration(request *GetUserProvisioningConfigurationRequest) (_result *GetUserProvisioningConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserProvisioningConfigurationResponse{}
	_body, _err := client.GetUserProvisioningConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserProvisioningEventWithOptions(request *GetUserProvisioningEventRequest, runtime *util.RuntimeOptions) (_result *GetUserProvisioningEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.EventId)) {
		query["EventId"] = request.EventId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserProvisioningEvent"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserProvisioningEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserProvisioningEvent(request *GetUserProvisioningEventRequest) (_result *GetUserProvisioningEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserProvisioningEventResponse{}
	_body, _err := client.GetUserProvisioningEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserProvisioningRdAccountStatisticsWithOptions(request *GetUserProvisioningRdAccountStatisticsRequest, runtime *util.RuntimeOptions) (_result *GetUserProvisioningRdAccountStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.RdMemberId)) {
		query["RdMemberId"] = request.RdMemberId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserProvisioningRdAccountStatistics"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserProvisioningRdAccountStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserProvisioningRdAccountStatistics(request *GetUserProvisioningRdAccountStatisticsRequest) (_result *GetUserProvisioningRdAccountStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserProvisioningRdAccountStatisticsResponse{}
	_body, _err := client.GetUserProvisioningRdAccountStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserProvisioningStatisticsWithOptions(request *GetUserProvisioningStatisticsRequest, runtime *util.RuntimeOptions) (_result *GetUserProvisioningStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.UserProvisioningId)) {
		query["UserProvisioningId"] = request.UserProvisioningId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserProvisioningStatistics"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserProvisioningStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserProvisioningStatistics(request *GetUserProvisioningStatisticsRequest) (_result *GetUserProvisioningStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserProvisioningStatisticsResponse{}
	_body, _err := client.GetUserProvisioningStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to query the assigned access permissions on the account `114240524784****` in your resource directory. The returned result shows that access permissions on the account in your resource directory is assigned to one user.
 *
 * @param request ListAccessAssignmentsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListAccessAssignmentsResponse
 */
func (client *Client) ListAccessAssignmentsWithOptions(request *ListAccessAssignmentsRequest, runtime *util.RuntimeOptions) (_result *ListAccessAssignmentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessConfigurationId)) {
		query["AccessConfigurationId"] = request.AccessConfigurationId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalId)) {
		query["PrincipalId"] = request.PrincipalId
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalType)) {
		query["PrincipalType"] = request.PrincipalType
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["TargetId"] = request.TargetId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAccessAssignments"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAccessAssignmentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to query the assigned access permissions on the account `114240524784****` in your resource directory. The returned result shows that access permissions on the account in your resource directory is assigned to one user.
 *
 * @param request ListAccessAssignmentsRequest
 * @return ListAccessAssignmentsResponse
 */
func (client *Client) ListAccessAssignments(request *ListAccessAssignmentsRequest) (_result *ListAccessAssignmentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAccessAssignmentsResponse{}
	_body, _err := client.ListAccessAssignmentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to query the accounts for which the access permission `ac-00ccule7tadaijxc****` is provisioned. The returned result shows that the access configuration is provisioned for two accounts in your resource directory.
 *
 * @param request ListAccessConfigurationProvisioningsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListAccessConfigurationProvisioningsResponse
 */
func (client *Client) ListAccessConfigurationProvisioningsWithOptions(request *ListAccessConfigurationProvisioningsRequest, runtime *util.RuntimeOptions) (_result *ListAccessConfigurationProvisioningsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessConfigurationId)) {
		query["AccessConfigurationId"] = request.AccessConfigurationId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProvisioningStatus)) {
		query["ProvisioningStatus"] = request.ProvisioningStatus
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["TargetId"] = request.TargetId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAccessConfigurationProvisionings"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAccessConfigurationProvisioningsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to query the accounts for which the access permission `ac-00ccule7tadaijxc****` is provisioned. The returned result shows that the access configuration is provisioned for two accounts in your resource directory.
 *
 * @param request ListAccessConfigurationProvisioningsRequest
 * @return ListAccessConfigurationProvisioningsResponse
 */
func (client *Client) ListAccessConfigurationProvisionings(request *ListAccessConfigurationProvisioningsRequest) (_result *ListAccessConfigurationProvisioningsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAccessConfigurationProvisioningsResponse{}
	_body, _err := client.ListAccessConfigurationProvisioningsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to query the access configurations within the directory `d-00fc2p61****`. The returned result shows that the directory contains the `VPC-Admin` and `ECS-Admin` access configurations.
 *
 * @param request ListAccessConfigurationsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListAccessConfigurationsResponse
 */
func (client *Client) ListAccessConfigurationsWithOptions(request *ListAccessConfigurationsRequest, runtime *util.RuntimeOptions) (_result *ListAccessConfigurationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.StatusNotifications)) {
		query["StatusNotifications"] = request.StatusNotifications
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAccessConfigurations"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAccessConfigurationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to query the access configurations within the directory `d-00fc2p61****`. The returned result shows that the directory contains the `VPC-Admin` and `ECS-Admin` access configurations.
 *
 * @param request ListAccessConfigurationsRequest
 * @return ListAccessConfigurationsResponse
 */
func (client *Client) ListAccessConfigurations(request *ListAccessConfigurationsRequest) (_result *ListAccessConfigurationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAccessConfigurationsResponse{}
	_body, _err := client.ListAccessConfigurationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to query the directories within your Alibaba Cloud account. The returned result shows that only one directory with the ID `d-00fc2p61****` is created within your Alibaba Cloud account.
 *
 * @param request ListDirectoriesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListDirectoriesResponse
 */
func (client *Client) ListDirectoriesWithOptions(runtime *util.RuntimeOptions) (_result *ListDirectoriesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListDirectories"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDirectoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to query the directories within your Alibaba Cloud account. The returned result shows that only one directory with the ID `d-00fc2p61****` is created within your Alibaba Cloud account.
 *
 * @return ListDirectoriesResponse
 */
func (client *Client) ListDirectories() (_result *ListDirectoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDirectoriesResponse{}
	_body, _err := client.ListDirectoriesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to query the SAML signing certificates within the directory `d-00fc2p61****`. The returned result shows that the directory contains one SAML signing certificate.
 *
 * @param request ListExternalSAMLIdPCertificatesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListExternalSAMLIdPCertificatesResponse
 */
func (client *Client) ListExternalSAMLIdPCertificatesWithOptions(request *ListExternalSAMLIdPCertificatesRequest, runtime *util.RuntimeOptions) (_result *ListExternalSAMLIdPCertificatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListExternalSAMLIdPCertificates"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListExternalSAMLIdPCertificatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to query the SAML signing certificates within the directory `d-00fc2p61****`. The returned result shows that the directory contains one SAML signing certificate.
 *
 * @param request ListExternalSAMLIdPCertificatesRequest
 * @return ListExternalSAMLIdPCertificatesResponse
 */
func (client *Client) ListExternalSAMLIdPCertificates(request *ListExternalSAMLIdPCertificatesRequest) (_result *ListExternalSAMLIdPCertificatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListExternalSAMLIdPCertificatesResponse{}
	_body, _err := client.ListExternalSAMLIdPCertificatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to query the users in the group `g-00jqzghi2n3o5hkh****`. The returned result shows that the group contains the user `Alice` and the user `user1`.
 *
 * @param request ListGroupMembersRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListGroupMembersResponse
 */
func (client *Client) ListGroupMembersWithOptions(request *ListGroupMembersRequest, runtime *util.RuntimeOptions) (_result *ListGroupMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGroupMembers"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGroupMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to query the users in the group `g-00jqzghi2n3o5hkh****`. The returned result shows that the group contains the user `Alice` and the user `user1`.
 *
 * @param request ListGroupMembersRequest
 * @return ListGroupMembersResponse
 */
func (client *Client) ListGroupMembers(request *ListGroupMembersRequest) (_result *ListGroupMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGroupMembersResponse{}
	_body, _err := client.ListGroupMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to query the groups in the directory `d-00fc2p61****`. The returned result shows that the directory contains three groups. The groups `group1` and `group2` are synchronized from an external identity provider (IdP). The group `TestGroup` is manually created in CloudSSO.
 *
 * @param request ListGroupsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListGroupsResponse
 */
func (client *Client) ListGroupsWithOptions(request *ListGroupsRequest, runtime *util.RuntimeOptions) (_result *ListGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProvisionType)) {
		query["ProvisionType"] = request.ProvisionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGroups"),
		Version:     tea.String("2021-05-15"),
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

/**
 * This topic provides an example on how to query the groups in the directory `d-00fc2p61****`. The returned result shows that the directory contains three groups. The groups `group1` and `group2` are synchronized from an external identity provider (IdP). The group `TestGroup` is manually created in CloudSSO.
 *
 * @param request ListGroupsRequest
 * @return ListGroupsResponse
 */
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

/**
 * This topic provides an example on how to query the groups to which the user `u-00q8wbq42wiltcrk****` is added. The returned result shows that the user is added to both the `TestGroup` and the `group1` groups.
 *
 * @param request ListJoinedGroupsForUserRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListJoinedGroupsForUserResponse
 */
func (client *Client) ListJoinedGroupsForUserWithOptions(request *ListJoinedGroupsForUserRequest, runtime *util.RuntimeOptions) (_result *ListJoinedGroupsForUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListJoinedGroupsForUser"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListJoinedGroupsForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to query the groups to which the user `u-00q8wbq42wiltcrk****` is added. The returned result shows that the user is added to both the `TestGroup` and the `group1` groups.
 *
 * @param request ListJoinedGroupsForUserRequest
 * @return ListJoinedGroupsForUserResponse
 */
func (client *Client) ListJoinedGroupsForUser(request *ListJoinedGroupsForUserRequest) (_result *ListJoinedGroupsForUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListJoinedGroupsForUserResponse{}
	_body, _err := client.ListJoinedGroupsForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to query the MFA devices that are bound to the user `u-00q8wbq42wiltcrk****`. The returned result shows that the MFA device named `Alice-MFA1` is bound to the user.
 *
 * @param request ListMFADevicesForUserRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListMFADevicesForUserResponse
 */
func (client *Client) ListMFADevicesForUserWithOptions(request *ListMFADevicesForUserRequest, runtime *util.RuntimeOptions) (_result *ListMFADevicesForUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMFADevicesForUser"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMFADevicesForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to query the MFA devices that are bound to the user `u-00q8wbq42wiltcrk****`. The returned result shows that the MFA device named `Alice-MFA1` is bound to the user.
 *
 * @param request ListMFADevicesForUserRequest
 * @return ListMFADevicesForUserResponse
 */
func (client *Client) ListMFADevicesForUser(request *ListMFADevicesForUserRequest) (_result *ListMFADevicesForUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMFADevicesForUserResponse{}
	_body, _err := client.ListMFADevicesForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to query the policies that are created for the access configuration `ac-00jhtfl8thteu6uj****`. The returned result shows that the access configuration contains one system policy and one inline policy.
 *
 * @param request ListPermissionPoliciesInAccessConfigurationRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListPermissionPoliciesInAccessConfigurationResponse
 */
func (client *Client) ListPermissionPoliciesInAccessConfigurationWithOptions(request *ListPermissionPoliciesInAccessConfigurationRequest, runtime *util.RuntimeOptions) (_result *ListPermissionPoliciesInAccessConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessConfigurationId)) {
		query["AccessConfigurationId"] = request.AccessConfigurationId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.PermissionPolicyType)) {
		query["PermissionPolicyType"] = request.PermissionPolicyType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPermissionPoliciesInAccessConfiguration"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPermissionPoliciesInAccessConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to query the policies that are created for the access configuration `ac-00jhtfl8thteu6uj****`. The returned result shows that the access configuration contains one system policy and one inline policy.
 *
 * @param request ListPermissionPoliciesInAccessConfigurationRequest
 * @return ListPermissionPoliciesInAccessConfigurationResponse
 */
func (client *Client) ListPermissionPoliciesInAccessConfiguration(request *ListPermissionPoliciesInAccessConfigurationRequest) (_result *ListPermissionPoliciesInAccessConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPermissionPoliciesInAccessConfigurationResponse{}
	_body, _err := client.ListPermissionPoliciesInAccessConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to query the SCIM credentials within the `d-00fc2p61****` directory.
 *
 * @param request ListSCIMServerCredentialsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListSCIMServerCredentialsResponse
 */
func (client *Client) ListSCIMServerCredentialsWithOptions(request *ListSCIMServerCredentialsRequest, runtime *util.RuntimeOptions) (_result *ListSCIMServerCredentialsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSCIMServerCredentials"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSCIMServerCredentialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to query the SCIM credentials within the `d-00fc2p61****` directory.
 *
 * @param request ListSCIMServerCredentialsRequest
 * @return ListSCIMServerCredentialsResponse
 */
func (client *Client) ListSCIMServerCredentials(request *ListSCIMServerCredentialsRequest) (_result *ListSCIMServerCredentialsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSCIMServerCredentialsResponse{}
	_body, _err := client.ListSCIMServerCredentialsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * By default, this operation queries the tasks within the previous 24 hours. This operation allows you to query the tasks within a maximum of 7 days. You can specify the start time of the query by using `Filter`.
 * This topic provides an example on how to query the tasks within the previous 24 hours.
 *
 * @param request ListTasksRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListTasksResponse
 */
func (client *Client) ListTasksWithOptions(request *ListTasksRequest, runtime *util.RuntimeOptions) (_result *ListTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessConfigurationId)) {
		query["AccessConfigurationId"] = request.AccessConfigurationId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalId)) {
		query["PrincipalId"] = request.PrincipalId
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalType)) {
		query["PrincipalType"] = request.PrincipalType
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["TargetId"] = request.TargetId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTasks"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * By default, this operation queries the tasks within the previous 24 hours. This operation allows you to query the tasks within a maximum of 7 days. You can specify the start time of the query by using `Filter`.
 * This topic provides an example on how to query the tasks within the previous 24 hours.
 *
 * @param request ListTasksRequest
 * @return ListTasksResponse
 */
func (client *Client) ListTasks(request *ListTasksRequest) (_result *ListTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTasksResponse{}
	_body, _err := client.ListTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserProvisioningEventsWithOptions(request *ListUserProvisioningEventsRequest, runtime *util.RuntimeOptions) (_result *ListUserProvisioningEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.UserProvisioningId)) {
		query["UserProvisioningId"] = request.UserProvisioningId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserProvisioningEvents"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserProvisioningEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserProvisioningEvents(request *ListUserProvisioningEventsRequest) (_result *ListUserProvisioningEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserProvisioningEventsResponse{}
	_body, _err := client.ListUserProvisioningEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserProvisioningsWithOptions(request *ListUserProvisioningsRequest, runtime *util.RuntimeOptions) (_result *ListUserProvisioningsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalId)) {
		query["PrincipalId"] = request.PrincipalId
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalType)) {
		query["PrincipalType"] = request.PrincipalType
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["TargetId"] = request.TargetId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserProvisionings"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserProvisioningsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserProvisionings(request *ListUserProvisioningsRequest) (_result *ListUserProvisioningsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserProvisioningsResponse{}
	_body, _err := client.ListUserProvisioningsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to query users in the `d-00fc2p61****` directory. The returned result shows that the directory contains two users. The user `AliceLee` is synchronized from an external identity provider (IdP). The user `user1` is manually created within CloudSSO.
 *
 * @param request ListUsersRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListUsersResponse
 */
func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *util.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProvisionType)) {
		query["ProvisionType"] = request.ProvisionType
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsers"),
		Version:     tea.String("2021-05-15"),
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

/**
 * This topic provides an example on how to query users in the `d-00fc2p61****` directory. The returned result shows that the directory contains two users. The user `AliceLee` is synchronized from an external identity provider (IdP). The user `user1` is manually created within CloudSSO.
 *
 * @param request ListUsersRequest
 * @return ListUsersResponse
 */
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

/**
 * When you call this operation, an asynchronous task is automatically created. You can call the [GetTask](~~340670~~) operation to query the progress of the task based on the value of the `TaskId` response parameter.
 * This topic provides an example on how to provision the access configuration `ac-00jhtfl8thteu6uj****` for the account `114240524784****` in your resource directory.
 *
 * @param request ProvisionAccessConfigurationRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ProvisionAccessConfigurationResponse
 */
func (client *Client) ProvisionAccessConfigurationWithOptions(request *ProvisionAccessConfigurationRequest, runtime *util.RuntimeOptions) (_result *ProvisionAccessConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessConfigurationId)) {
		query["AccessConfigurationId"] = request.AccessConfigurationId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["TargetId"] = request.TargetId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ProvisionAccessConfiguration"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ProvisionAccessConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * When you call this operation, an asynchronous task is automatically created. You can call the [GetTask](~~340670~~) operation to query the progress of the task based on the value of the `TaskId` response parameter.
 * This topic provides an example on how to provision the access configuration `ac-00jhtfl8thteu6uj****` for the account `114240524784****` in your resource directory.
 *
 * @param request ProvisionAccessConfigurationRequest
 * @return ProvisionAccessConfigurationResponse
 */
func (client *Client) ProvisionAccessConfiguration(request *ProvisionAccessConfigurationRequest) (_result *ProvisionAccessConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ProvisionAccessConfigurationResponse{}
	_body, _err := client.ProvisionAccessConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to remove the SAML signing certificate whose ID is `idp-c-00dt9gnl7fmjaw9c****`.
 *
 * @param request RemoveExternalSAMLIdPCertificateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RemoveExternalSAMLIdPCertificateResponse
 */
func (client *Client) RemoveExternalSAMLIdPCertificateWithOptions(request *RemoveExternalSAMLIdPCertificateRequest, runtime *util.RuntimeOptions) (_result *RemoveExternalSAMLIdPCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertificateId)) {
		query["CertificateId"] = request.CertificateId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveExternalSAMLIdPCertificate"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveExternalSAMLIdPCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to remove the SAML signing certificate whose ID is `idp-c-00dt9gnl7fmjaw9c****`.
 *
 * @param request RemoveExternalSAMLIdPCertificateRequest
 * @return RemoveExternalSAMLIdPCertificateResponse
 */
func (client *Client) RemoveExternalSAMLIdPCertificate(request *RemoveExternalSAMLIdPCertificateRequest) (_result *RemoveExternalSAMLIdPCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveExternalSAMLIdPCertificateResponse{}
	_body, _err := client.RemoveExternalSAMLIdPCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * After you remove an inline policy from an access configuration, the policy cannot be restored.
 * This topic provides an example on how to remove the system policy `AliyunECSFullAccess` from the access configuration `ac-00jhtfl8thteu6uj****`.
 *
 * @param request RemovePermissionPolicyFromAccessConfigurationRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RemovePermissionPolicyFromAccessConfigurationResponse
 */
func (client *Client) RemovePermissionPolicyFromAccessConfigurationWithOptions(request *RemovePermissionPolicyFromAccessConfigurationRequest, runtime *util.RuntimeOptions) (_result *RemovePermissionPolicyFromAccessConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessConfigurationId)) {
		query["AccessConfigurationId"] = request.AccessConfigurationId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.PermissionPolicyName)) {
		query["PermissionPolicyName"] = request.PermissionPolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.PermissionPolicyType)) {
		query["PermissionPolicyType"] = request.PermissionPolicyType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemovePermissionPolicyFromAccessConfiguration"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemovePermissionPolicyFromAccessConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * After you remove an inline policy from an access configuration, the policy cannot be restored.
 * This topic provides an example on how to remove the system policy `AliyunECSFullAccess` from the access configuration `ac-00jhtfl8thteu6uj****`.
 *
 * @param request RemovePermissionPolicyFromAccessConfigurationRequest
 * @return RemovePermissionPolicyFromAccessConfigurationResponse
 */
func (client *Client) RemovePermissionPolicyFromAccessConfiguration(request *RemovePermissionPolicyFromAccessConfigurationRequest) (_result *RemovePermissionPolicyFromAccessConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemovePermissionPolicyFromAccessConfigurationResponse{}
	_body, _err := client.RemovePermissionPolicyFromAccessConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * If System for Cross-domain Identity Management (SCIM) synchronization is enabled, you cannot remove a user from a group that is synchronized by using SCIM.
 * This topic provides an example on how to remove the user `u-00q8wbq42wiltcrk****` from the group `g-00jqzghi2n3o5hkh****`.
 *
 * @param request RemoveUserFromGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RemoveUserFromGroupResponse
 */
func (client *Client) RemoveUserFromGroupWithOptions(request *RemoveUserFromGroupRequest, runtime *util.RuntimeOptions) (_result *RemoveUserFromGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveUserFromGroup"),
		Version:     tea.String("2021-05-15"),
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

/**
 * If System for Cross-domain Identity Management (SCIM) synchronization is enabled, you cannot remove a user from a group that is synchronized by using SCIM.
 * This topic provides an example on how to remove the user `u-00q8wbq42wiltcrk****` from the group `g-00jqzghi2n3o5hkh****`.
 *
 * @param request RemoveUserFromGroupRequest
 * @return RemoveUserFromGroupResponse
 */
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

/**
 * If you forget your password or your password expires or is at risk, you must contact a CloudSSO administrator to reset your password.
 * >  After you enable SSO logon, your password cannot be reset.
 * This topic provides an example on how to reset the password of the user `u-00q8wbq42wiltcrk****`. The new password is automatically generated by the system.
 *
 * @param request ResetUserPasswordRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ResetUserPasswordResponse
 */
func (client *Client) ResetUserPasswordWithOptions(request *ResetUserPasswordRequest, runtime *util.RuntimeOptions) (_result *ResetUserPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.GenerateRandomPassword)) {
		query["GenerateRandomPassword"] = request.GenerateRandomPassword
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.RequirePasswordResetForNextLogin)) {
		query["RequirePasswordResetForNextLogin"] = request.RequirePasswordResetForNextLogin
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetUserPassword"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetUserPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If you forget your password or your password expires or is at risk, you must contact a CloudSSO administrator to reset your password.
 * >  After you enable SSO logon, your password cannot be reset.
 * This topic provides an example on how to reset the password of the user `u-00q8wbq42wiltcrk****`. The new password is automatically generated by the system.
 *
 * @param request ResetUserPasswordRequest
 * @return ResetUserPasswordResponse
 */
func (client *Client) ResetUserPassword(request *ResetUserPasswordRequest) (_result *ResetUserPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetUserPasswordResponse{}
	_body, _err := client.ResetUserPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RetryUserProvisioningEventWithOptions(request *RetryUserProvisioningEventRequest, runtime *util.RuntimeOptions) (_result *RetryUserProvisioningEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.DuplicationStrategy)) {
		query["DuplicationStrategy"] = request.DuplicationStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.EventId)) {
		query["EventId"] = request.EventId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RetryUserProvisioningEvent"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RetryUserProvisioningEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RetryUserProvisioningEvent(request *RetryUserProvisioningEventRequest) (_result *RetryUserProvisioningEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RetryUserProvisioningEventResponse{}
	_body, _err := client.RetryUserProvisioningEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * During SAML 2.0-based single sign-on (SSO) logon, CloudSSO is an SP, and the identity management system of an enterprise is an IdP.
 * You can use one of the following methods to configure a SAML IdP. You can obtain the required metadata file or parameter values from your IdP.
 * *   Use the metadata file. You can specify the `EncodedMetadataDocument` parameter to upload the metadata file.
 * *   Manually configure the IdP. You can manually specify the following parameters for your IdP: `EntityId`, `LoginUrl`, `WantRequestSigned`, and `X509Certificate`.
 * If you have configured a SAML IdP, the existing configurations are replaced after you call this operation.
 * *   If the IdP is configured by using the metadata file, all existing configurations are replaced with new configurations.
 * *   If the IdP is manually configured, the original parameter values that are different from the new parameter values are replaced.
 * >  If SSO logon is enabled, new configurations immediately take effect. Take note of the impacts on the production environment.
 * This topic provides an example on how to configure an IdP by using the metadata file within the directory `d-00fc2p61****`.
 *
 * @param request SetExternalSAMLIdentityProviderRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SetExternalSAMLIdentityProviderResponse
 */
func (client *Client) SetExternalSAMLIdentityProviderWithOptions(request *SetExternalSAMLIdentityProviderRequest, runtime *util.RuntimeOptions) (_result *SetExternalSAMLIdentityProviderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.EncodedMetadataDocument)) {
		query["EncodedMetadataDocument"] = request.EncodedMetadataDocument
	}

	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		query["EntityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginUrl)) {
		query["LoginUrl"] = request.LoginUrl
	}

	if !tea.BoolValue(util.IsUnset(request.SSOStatus)) {
		query["SSOStatus"] = request.SSOStatus
	}

	if !tea.BoolValue(util.IsUnset(request.WantRequestSigned)) {
		query["WantRequestSigned"] = request.WantRequestSigned
	}

	if !tea.BoolValue(util.IsUnset(request.X509Certificate)) {
		query["X509Certificate"] = request.X509Certificate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetExternalSAMLIdentityProvider"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetExternalSAMLIdentityProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * During SAML 2.0-based single sign-on (SSO) logon, CloudSSO is an SP, and the identity management system of an enterprise is an IdP.
 * You can use one of the following methods to configure a SAML IdP. You can obtain the required metadata file or parameter values from your IdP.
 * *   Use the metadata file. You can specify the `EncodedMetadataDocument` parameter to upload the metadata file.
 * *   Manually configure the IdP. You can manually specify the following parameters for your IdP: `EntityId`, `LoginUrl`, `WantRequestSigned`, and `X509Certificate`.
 * If you have configured a SAML IdP, the existing configurations are replaced after you call this operation.
 * *   If the IdP is configured by using the metadata file, all existing configurations are replaced with new configurations.
 * *   If the IdP is manually configured, the original parameter values that are different from the new parameter values are replaced.
 * >  If SSO logon is enabled, new configurations immediately take effect. Take note of the impacts on the production environment.
 * This topic provides an example on how to configure an IdP by using the metadata file within the directory `d-00fc2p61****`.
 *
 * @param request SetExternalSAMLIdentityProviderRequest
 * @return SetExternalSAMLIdentityProviderResponse
 */
func (client *Client) SetExternalSAMLIdentityProvider(request *SetExternalSAMLIdentityProviderRequest) (_result *SetExternalSAMLIdentityProviderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetExternalSAMLIdentityProviderResponse{}
	_body, _err := client.SetExternalSAMLIdentityProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetLoginPreferenceWithOptions(request *SetLoginPreferenceRequest, runtime *util.RuntimeOptions) (_result *SetLoginPreferenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginNetworkMasks)) {
		query["LoginNetworkMasks"] = request.LoginNetworkMasks
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetLoginPreference"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetLoginPreferenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetLoginPreference(request *SetLoginPreferenceRequest) (_result *SetLoginPreferenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetLoginPreferenceResponse{}
	_body, _err := client.SetLoginPreferenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * If a CloudSSO administrator enables username-password logon for users, CloudSSO automatically enables MFA to improve security. The administrator can call this operation to enable or disable MFA based on the business requirements.
 * This topic provides an example on how to enable MFA for users.
 *
 * @param request SetMFAAuthenticationStatusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SetMFAAuthenticationStatusResponse
 */
func (client *Client) SetMFAAuthenticationStatusWithOptions(request *SetMFAAuthenticationStatusRequest, runtime *util.RuntimeOptions) (_result *SetMFAAuthenticationStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.MFAAuthenticationStatus)) {
		query["MFAAuthenticationStatus"] = request.MFAAuthenticationStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetMFAAuthenticationStatus"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetMFAAuthenticationStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If a CloudSSO administrator enables username-password logon for users, CloudSSO automatically enables MFA to improve security. The administrator can call this operation to enable or disable MFA based on the business requirements.
 * This topic provides an example on how to enable MFA for users.
 *
 * @param request SetMFAAuthenticationStatusRequest
 * @return SetMFAAuthenticationStatusResponse
 */
func (client *Client) SetMFAAuthenticationStatus(request *SetMFAAuthenticationStatusRequest) (_result *SetMFAAuthenticationStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetMFAAuthenticationStatusResponse{}
	_body, _err := client.SetMFAAuthenticationStatusWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxLoginAttempts)) {
		query["MaxLoginAttempts"] = request.MaxLoginAttempts
	}

	if !tea.BoolValue(util.IsUnset(request.MaxPasswordAge)) {
		query["MaxPasswordAge"] = request.MaxPasswordAge
	}

	if !tea.BoolValue(util.IsUnset(request.MinPasswordDifferentChars)) {
		query["MinPasswordDifferentChars"] = request.MinPasswordDifferentChars
	}

	if !tea.BoolValue(util.IsUnset(request.MinPasswordLength)) {
		query["MinPasswordLength"] = request.MinPasswordLength
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordNotContainUsername)) {
		query["PasswordNotContainUsername"] = request.PasswordNotContainUsername
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordReusePrevention)) {
		query["PasswordReusePrevention"] = request.PasswordReusePrevention
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetPasswordPolicy"),
		Version:     tea.String("2021-05-15"),
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

/**
 * You can synchronize users or groups from an external identity provider (IdP) that supports SCIM 2.0 to CloudSSO only after SCIM synchronization is enabled. If you disable SCIM synchronization, you can no longer synchronize users or groups to CloudSSO. The following list describes the impacts after SCIM synchronization is enabled or disabled:
 * *   After you enable SCIM synchronization, you cannot modify or delete the users or groups that are synchronized to CloudSSO by using SCIM. In addition, you cannot add users to or remove users from the groups. However, you can change the passwords of the users and enable or disable the logon of the users.
 * *   After you disable SCIM synchronization, you can modify and delete the users and groups that are synchronized to CloudSSO by using SCIM. You can also add users to or remove users from the groups.
 * This topic provides an example on how to enable SCIM synchronization within the directory `d-00fc2p61****`.
 *
 * @param request SetSCIMSynchronizationStatusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SetSCIMSynchronizationStatusResponse
 */
func (client *Client) SetSCIMSynchronizationStatusWithOptions(request *SetSCIMSynchronizationStatusRequest, runtime *util.RuntimeOptions) (_result *SetSCIMSynchronizationStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.SCIMSynchronizationStatus)) {
		query["SCIMSynchronizationStatus"] = request.SCIMSynchronizationStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetSCIMSynchronizationStatus"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetSCIMSynchronizationStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can synchronize users or groups from an external identity provider (IdP) that supports SCIM 2.0 to CloudSSO only after SCIM synchronization is enabled. If you disable SCIM synchronization, you can no longer synchronize users or groups to CloudSSO. The following list describes the impacts after SCIM synchronization is enabled or disabled:
 * *   After you enable SCIM synchronization, you cannot modify or delete the users or groups that are synchronized to CloudSSO by using SCIM. In addition, you cannot add users to or remove users from the groups. However, you can change the passwords of the users and enable or disable the logon of the users.
 * *   After you disable SCIM synchronization, you can modify and delete the users and groups that are synchronized to CloudSSO by using SCIM. You can also add users to or remove users from the groups.
 * This topic provides an example on how to enable SCIM synchronization within the directory `d-00fc2p61****`.
 *
 * @param request SetSCIMSynchronizationStatusRequest
 * @return SetSCIMSynchronizationStatusResponse
 */
func (client *Client) SetSCIMSynchronizationStatus(request *SetSCIMSynchronizationStatusRequest) (_result *SetSCIMSynchronizationStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetSCIMSynchronizationStatusResponse{}
	_body, _err := client.SetSCIMSynchronizationStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can modify the `Description`, `SessionDuration`, and `RelayState` parameters for an access configuration.
 * This topic provides an example on how to change the initial web page in the access configuration `ac-00jhtfl8thteu6uj****` to `https://cloudsso.console.aliyun.com`.
 *
 * @param request UpdateAccessConfigurationRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateAccessConfigurationResponse
 */
func (client *Client) UpdateAccessConfigurationWithOptions(request *UpdateAccessConfigurationRequest, runtime *util.RuntimeOptions) (_result *UpdateAccessConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessConfigurationId)) {
		query["AccessConfigurationId"] = request.AccessConfigurationId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.NewDescription)) {
		query["NewDescription"] = request.NewDescription
	}

	if !tea.BoolValue(util.IsUnset(request.NewRelayState)) {
		query["NewRelayState"] = request.NewRelayState
	}

	if !tea.BoolValue(util.IsUnset(request.NewSessionDuration)) {
		query["NewSessionDuration"] = request.NewSessionDuration
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAccessConfiguration"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAccessConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can modify the `Description`, `SessionDuration`, and `RelayState` parameters for an access configuration.
 * This topic provides an example on how to change the initial web page in the access configuration `ac-00jhtfl8thteu6uj****` to `https://cloudsso.console.aliyun.com`.
 *
 * @param request UpdateAccessConfigurationRequest
 * @return UpdateAccessConfigurationResponse
 */
func (client *Client) UpdateAccessConfiguration(request *UpdateAccessConfigurationRequest) (_result *UpdateAccessConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAccessConfigurationResponse{}
	_body, _err := client.UpdateAccessConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * After you change the name of a directory, the URL that is used to log on to the Cloud SSO user portal is changed. You must notify the Cloud SSO users of the correct URL.
 * This topic provides an example on how to change the name of a directory to `new-example`.
 *
 * @param request UpdateDirectoryRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateDirectoryResponse
 */
func (client *Client) UpdateDirectoryWithOptions(request *UpdateDirectoryRequest, runtime *util.RuntimeOptions) (_result *UpdateDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.NewDirectoryName)) {
		query["NewDirectoryName"] = request.NewDirectoryName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDirectory"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * After you change the name of a directory, the URL that is used to log on to the Cloud SSO user portal is changed. You must notify the Cloud SSO users of the correct URL.
 * This topic provides an example on how to change the name of a directory to `new-example`.
 *
 * @param request UpdateDirectoryRequest
 * @return UpdateDirectoryResponse
 */
func (client *Client) UpdateDirectory(request *UpdateDirectoryRequest) (_result *UpdateDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDirectoryResponse{}
	_body, _err := client.UpdateDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can modify `GroupName` and `Description` for a group.
 * >  If System for Cross-domain Identity Management (SCIM) synchronization is enabled, you cannot modify the information about a group that is synchronized by using SCIM.
 * This topic provides an example on how to change the name of the group `g-00jqzghi2n3o5hkh****` to `NewTestGroup`.
 *
 * @param request UpdateGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateGroupResponse
 */
func (client *Client) UpdateGroupWithOptions(request *UpdateGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.NewDescription)) {
		query["NewDescription"] = request.NewDescription
	}

	if !tea.BoolValue(util.IsUnset(request.NewGroupName)) {
		query["NewGroupName"] = request.NewGroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGroup"),
		Version:     tea.String("2021-05-15"),
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

/**
 * You can modify `GroupName` and `Description` for a group.
 * >  If System for Cross-domain Identity Management (SCIM) synchronization is enabled, you cannot modify the information about a group that is synchronized by using SCIM.
 * This topic provides an example on how to change the name of the group `g-00jqzghi2n3o5hkh****` to `NewTestGroup`.
 *
 * @param request UpdateGroupRequest
 * @return UpdateGroupResponse
 */
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

/**
 * This topic provides an example on how to modify an inline policy that is created for the access configuration `ac-00jhtfl8thteu6uj****`.
 *
 * @param request UpdateInlinePolicyForAccessConfigurationRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateInlinePolicyForAccessConfigurationResponse
 */
func (client *Client) UpdateInlinePolicyForAccessConfigurationWithOptions(request *UpdateInlinePolicyForAccessConfigurationRequest, runtime *util.RuntimeOptions) (_result *UpdateInlinePolicyForAccessConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessConfigurationId)) {
		query["AccessConfigurationId"] = request.AccessConfigurationId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.InlinePolicyName)) {
		query["InlinePolicyName"] = request.InlinePolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.NewInlinePolicyDocument)) {
		query["NewInlinePolicyDocument"] = request.NewInlinePolicyDocument
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInlinePolicyForAccessConfiguration"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInlinePolicyForAccessConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to modify an inline policy that is created for the access configuration `ac-00jhtfl8thteu6uj****`.
 *
 * @param request UpdateInlinePolicyForAccessConfigurationRequest
 * @return UpdateInlinePolicyForAccessConfigurationResponse
 */
func (client *Client) UpdateInlinePolicyForAccessConfiguration(request *UpdateInlinePolicyForAccessConfigurationRequest) (_result *UpdateInlinePolicyForAccessConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateInlinePolicyForAccessConfigurationResponse{}
	_body, _err := client.UpdateInlinePolicyForAccessConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * If you enable username-password logon for CloudSSO users, you can also configure MFA for the users.
 * This topic provides an example on how to enable MFA for all CloudSSO users that belong to the directory named `d-00fc2p61****`.
 *
 * @param request UpdateMFAAuthenticationSettingsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateMFAAuthenticationSettingsResponse
 */
func (client *Client) UpdateMFAAuthenticationSettingsWithOptions(request *UpdateMFAAuthenticationSettingsRequest, runtime *util.RuntimeOptions) (_result *UpdateMFAAuthenticationSettingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.MFAAuthenticationSettings)) {
		query["MFAAuthenticationSettings"] = request.MFAAuthenticationSettings
	}

	if !tea.BoolValue(util.IsUnset(request.OperationForRiskLogin)) {
		query["OperationForRiskLogin"] = request.OperationForRiskLogin
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateMFAAuthenticationSettings"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMFAAuthenticationSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If you enable username-password logon for CloudSSO users, you can also configure MFA for the users.
 * This topic provides an example on how to enable MFA for all CloudSSO users that belong to the directory named `d-00fc2p61****`.
 *
 * @param request UpdateMFAAuthenticationSettingsRequest
 * @return UpdateMFAAuthenticationSettingsResponse
 */
func (client *Client) UpdateMFAAuthenticationSettings(request *UpdateMFAAuthenticationSettingsRequest) (_result *UpdateMFAAuthenticationSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateMFAAuthenticationSettingsResponse{}
	_body, _err := client.UpdateMFAAuthenticationSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to disable the SCIM credential whose ID is `scimcred-004whl0kvfwcypbi****`. After the SCIM credential is disabled, the synchronization task that uses the SCIM credential fails. You can call this operation again to enable the SCIM credential.
 *
 * @param request UpdateSCIMServerCredentialStatusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateSCIMServerCredentialStatusResponse
 */
func (client *Client) UpdateSCIMServerCredentialStatusWithOptions(request *UpdateSCIMServerCredentialStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateSCIMServerCredentialStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CredentialId)) {
		query["CredentialId"] = request.CredentialId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.NewStatus)) {
		query["NewStatus"] = request.NewStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSCIMServerCredentialStatus"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSCIMServerCredentialStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to disable the SCIM credential whose ID is `scimcred-004whl0kvfwcypbi****`. After the SCIM credential is disabled, the synchronization task that uses the SCIM credential fails. You can call this operation again to enable the SCIM credential.
 *
 * @param request UpdateSCIMServerCredentialStatusRequest
 * @return UpdateSCIMServerCredentialStatusResponse
 */
func (client *Client) UpdateSCIMServerCredentialStatus(request *UpdateSCIMServerCredentialStatusRequest) (_result *UpdateSCIMServerCredentialStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSCIMServerCredentialStatusResponse{}
	_body, _err := client.UpdateSCIMServerCredentialStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can modify `FirstName`, `LastName`, `DisplayName`, `Email`, and `Description` for a user. You cannot modify `UserName` for a user.
 * >  If System for Cross-domain Identity Management (SCIM) synchronization is enabled, you cannot modify the information about a user that is synchronized by using SCIM.
 * This topic provides an example on how to change the email address of the user whose ID is `u-00q8wbq42wiltcrk****` to `AliceLee@example.com`.
 *
 * @param request UpdateUserRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateUserResponse
 */
func (client *Client) UpdateUserWithOptions(request *UpdateUserRequest, runtime *util.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.NewDescription)) {
		query["NewDescription"] = request.NewDescription
	}

	if !tea.BoolValue(util.IsUnset(request.NewDisplayName)) {
		query["NewDisplayName"] = request.NewDisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.NewEmail)) {
		query["NewEmail"] = request.NewEmail
	}

	if !tea.BoolValue(util.IsUnset(request.NewFirstName)) {
		query["NewFirstName"] = request.NewFirstName
	}

	if !tea.BoolValue(util.IsUnset(request.NewLastName)) {
		query["NewLastName"] = request.NewLastName
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUser"),
		Version:     tea.String("2021-05-15"),
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

/**
 * You can modify `FirstName`, `LastName`, `DisplayName`, `Email`, and `Description` for a user. You cannot modify `UserName` for a user.
 * >  If System for Cross-domain Identity Management (SCIM) synchronization is enabled, you cannot modify the information about a user that is synchronized by using SCIM.
 * This topic provides an example on how to change the email address of the user whose ID is `u-00q8wbq42wiltcrk****` to `AliceLee@example.com`.
 *
 * @param request UpdateUserRequest
 * @return UpdateUserResponse
 */
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

/**
 * If you call the [UpdateMFAAuthenticationSettings](~~450134~~) operation to set the MFAAuthenticationSettings parameter to `Byuser`, user-specific settings are applied. Then, you must call the UpdateUserMFAAuthenticationSettings operation to configure MFA for each user.
 * By default, the MFAAuthenticationSettings parameter is set to `Enabled` for a new user.
 * This topic provides an example on how to enable MFA for the user named `u-00q8wbq42wiltcrk****`.
 *
 * @param request UpdateUserMFAAuthenticationSettingsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateUserMFAAuthenticationSettingsResponse
 */
func (client *Client) UpdateUserMFAAuthenticationSettingsWithOptions(request *UpdateUserMFAAuthenticationSettingsRequest, runtime *util.RuntimeOptions) (_result *UpdateUserMFAAuthenticationSettingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserMFAAuthenticationSettings)) {
		query["UserMFAAuthenticationSettings"] = request.UserMFAAuthenticationSettings
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUserMFAAuthenticationSettings"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserMFAAuthenticationSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If you call the [UpdateMFAAuthenticationSettings](~~450134~~) operation to set the MFAAuthenticationSettings parameter to `Byuser`, user-specific settings are applied. Then, you must call the UpdateUserMFAAuthenticationSettings operation to configure MFA for each user.
 * By default, the MFAAuthenticationSettings parameter is set to `Enabled` for a new user.
 * This topic provides an example on how to enable MFA for the user named `u-00q8wbq42wiltcrk****`.
 *
 * @param request UpdateUserMFAAuthenticationSettingsRequest
 * @return UpdateUserMFAAuthenticationSettingsResponse
 */
func (client *Client) UpdateUserMFAAuthenticationSettings(request *UpdateUserMFAAuthenticationSettingsRequest) (_result *UpdateUserMFAAuthenticationSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserMFAAuthenticationSettingsResponse{}
	_body, _err := client.UpdateUserMFAAuthenticationSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateUserProvisioningWithOptions(request *UpdateUserProvisioningRequest, runtime *util.RuntimeOptions) (_result *UpdateUserProvisioningResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.NewDeletionStrategy)) {
		query["NewDeletionStrategy"] = request.NewDeletionStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.NewDescription)) {
		query["NewDescription"] = request.NewDescription
	}

	if !tea.BoolValue(util.IsUnset(request.NewDuplicationStrategy)) {
		query["NewDuplicationStrategy"] = request.NewDuplicationStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.UserProvisioningId)) {
		query["UserProvisioningId"] = request.UserProvisioningId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUserProvisioning"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserProvisioningResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateUserProvisioning(request *UpdateUserProvisioningRequest) (_result *UpdateUserProvisioningResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserProvisioningResponse{}
	_body, _err := client.UpdateUserProvisioningWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateUserProvisioningConfigurationWithOptions(request *UpdateUserProvisioningConfigurationRequest, runtime *util.RuntimeOptions) (_result *UpdateUserProvisioningConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.NewDefaultLandingPage)) {
		query["NewDefaultLandingPage"] = request.NewDefaultLandingPage
	}

	if !tea.BoolValue(util.IsUnset(request.NewSessionDuration)) {
		query["NewSessionDuration"] = request.NewSessionDuration
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUserProvisioningConfiguration"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserProvisioningConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateUserProvisioningConfiguration(request *UpdateUserProvisioningConfigurationRequest) (_result *UpdateUserProvisioningConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserProvisioningConfigurationResponse{}
	_body, _err := client.UpdateUserProvisioningConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to change the status of the user whose ID is `u-00q8wbq42wiltcrk****` to Disabled. Users in the Disabled state cannot log on to the CloudSSO user portal.
 *
 * @param request UpdateUserStatusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateUserStatusResponse
 */
func (client *Client) UpdateUserStatusWithOptions(request *UpdateUserStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateUserStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.NewStatus)) {
		query["NewStatus"] = request.NewStatus
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUserStatus"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to change the status of the user whose ID is `u-00q8wbq42wiltcrk****` to Disabled. Users in the Disabled state cannot log on to the CloudSSO user portal.
 *
 * @param request UpdateUserStatusRequest
 * @return UpdateUserStatusResponse
 */
func (client *Client) UpdateUserStatus(request *UpdateUserStatusRequest) (_result *UpdateUserStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserStatusResponse{}
	_body, _err := client.UpdateUserStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
