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

type AddExternalSAMLIdPCertificateRequest struct {
	DirectoryId     *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
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
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddExternalSAMLIdPCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddExternalSAMLIdPCertificateResponse) SetBody(v *AddExternalSAMLIdPCertificateResponseBody) *AddExternalSAMLIdPCertificateResponse {
	s.Body = v
	return s
}

type AddPermissionPolicyToAccessConfigurationRequest struct {
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	DirectoryId           *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	InlinePolicyDocument  *string `json:"InlinePolicyDocument,omitempty" xml:"InlinePolicyDocument,omitempty"`
	PermissionPolicyName  *string `json:"PermissionPolicyName,omitempty" xml:"PermissionPolicyName,omitempty"`
	PermissionPolicyType  *string `json:"PermissionPolicyType,omitempty" xml:"PermissionPolicyType,omitempty"`
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
	Headers map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddPermissionPolicyToAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddPermissionPolicyToAccessConfigurationResponse) SetBody(v *AddPermissionPolicyToAccessConfigurationResponseBody) *AddPermissionPolicyToAccessConfigurationResponse {
	s.Body = v
	return s
}

type AddUserToGroupRequest struct {
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddUserToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddUserToGroupResponse) SetBody(v *AddUserToGroupResponseBody) *AddUserToGroupResponse {
	s.Body = v
	return s
}

type ClearExternalSAMLIdentityProviderRequest struct {
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
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ClearExternalSAMLIdentityProviderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ClearExternalSAMLIdentityProviderResponse) SetBody(v *ClearExternalSAMLIdentityProviderResponseBody) *ClearExternalSAMLIdentityProviderResponse {
	s.Body = v
	return s
}

type CreateAccessAssignmentRequest struct {
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	DirectoryId           *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	PrincipalId           *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	PrincipalType         *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	TargetId              *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetType            *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
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
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Task      *CreateAccessAssignmentResponseBodyTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
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
	AccessConfigurationId   *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	AccessConfigurationName *string `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	PrincipalId             *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	PrincipalName           *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	PrincipalType           *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetId                *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetName              *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	TargetPath              *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	TargetPathName          *string `json:"TargetPathName,omitempty" xml:"TargetPathName,omitempty"`
	TargetType              *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	TaskId                  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType                *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAccessAssignmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateAccessAssignmentResponse) SetBody(v *CreateAccessAssignmentResponseBody) *CreateAccessAssignmentResponse {
	s.Body = v
	return s
}

type CreateAccessConfigurationRequest struct {
	AccessConfigurationName *string `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DirectoryId             *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	RelayState              *string `json:"RelayState,omitempty" xml:"RelayState,omitempty"`
	SessionDuration         *int32  `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
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
	AccessConfiguration *CreateAccessConfigurationResponseBodyAccessConfiguration `json:"AccessConfiguration,omitempty" xml:"AccessConfiguration,omitempty" type:"Struct"`
	RequestId           *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccessConfigurationId   *string   `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	AccessConfigurationName *string   `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	CreateTime              *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description             *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	RelayState              *string   `json:"RelayState,omitempty" xml:"RelayState,omitempty"`
	SessionDuration         *int32    `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	StatusNotifications     []*string `json:"StatusNotifications,omitempty" xml:"StatusNotifications,omitempty" type:"Repeated"`
	UpdateTime              *string   `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateAccessConfigurationResponse) SetBody(v *CreateAccessConfigurationResponseBody) *CreateAccessConfigurationResponse {
	s.Body = v
	return s
}

type CreateDirectoryRequest struct {
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
	Directory *CreateDirectoryResponseBodyDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DirectoryId   *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	DirectoryName *string `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateDirectoryResponse) SetBody(v *CreateDirectoryResponseBody) *CreateDirectoryResponse {
	s.Body = v
	return s
}

type CreateGroupRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	GroupName   *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
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
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateGroupResponse) SetBody(v *CreateGroupResponseBody) *CreateGroupResponse {
	s.Body = v
	return s
}

type CreateSCIMServerCredentialRequest struct {
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
	RequestId            *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CredentialId     *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
	CredentialSecret *string `json:"CredentialSecret,omitempty" xml:"CredentialSecret,omitempty"`
	CredentialType   *string `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	DirectoryId      *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	ExpireTime       *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSCIMServerCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateSCIMServerCredentialResponse) SetBody(v *CreateSCIMServerCredentialResponseBody) *CreateSCIMServerCredentialResponse {
	s.Body = v
	return s
}

type CreateUserRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	FirstName   *string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	LastName    *string `json:"LastName,omitempty" xml:"LastName,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName   *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email         *string `json:"Email,omitempty" xml:"Email,omitempty"`
	FirstName     *string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	LastName      *string `json:"LastName,omitempty" xml:"LastName,omitempty"`
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName      *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateUserResponse) SetBody(v *CreateUserResponseBody) *CreateUserResponse {
	s.Body = v
	return s
}

type DeleteAccessAssignmentRequest struct {
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	DeprovisionStrategy   *string `json:"DeprovisionStrategy,omitempty" xml:"DeprovisionStrategy,omitempty"`
	DirectoryId           *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	PrincipalId           *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	PrincipalType         *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	TargetId              *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetType            *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
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
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Task      *DeleteAccessAssignmentResponseBodyTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
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
	AccessConfigurationId   *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	AccessConfigurationName *string `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	PrincipalId             *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	PrincipalName           *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	PrincipalType           *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetId                *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetName              *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	TargetPath              *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	TargetPathName          *string `json:"TargetPathName,omitempty" xml:"TargetPathName,omitempty"`
	TargetType              *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	TaskId                  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType                *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAccessAssignmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteAccessAssignmentResponse) SetBody(v *DeleteAccessAssignmentResponseBody) *DeleteAccessAssignmentResponse {
	s.Body = v
	return s
}

type DeleteAccessConfigurationRequest struct {
	AccessConfigurationId         *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	DirectoryId                   *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	ForceRemovePermissionPolicies *bool   `json:"ForceRemovePermissionPolicies,omitempty" xml:"ForceRemovePermissionPolicies,omitempty"`
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteAccessConfigurationResponse) SetBody(v *DeleteAccessConfigurationResponseBody) *DeleteAccessConfigurationResponse {
	s.Body = v
	return s
}

type DeleteDirectoryRequest struct {
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteDirectoryResponse) SetBody(v *DeleteDirectoryResponseBody) *DeleteDirectoryResponse {
	s.Body = v
	return s
}

type DeleteGroupRequest struct {
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteGroupResponse) SetBody(v *DeleteGroupResponseBody) *DeleteGroupResponse {
	s.Body = v
	return s
}

type DeleteMFADeviceForUserRequest struct {
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	MFADeviceId *string `json:"MFADeviceId,omitempty" xml:"MFADeviceId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteMFADeviceForUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteMFADeviceForUserResponse) SetBody(v *DeleteMFADeviceForUserResponseBody) *DeleteMFADeviceForUserResponse {
	s.Body = v
	return s
}

type DeleteSCIMServerCredentialRequest struct {
	CredentialId *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
	DirectoryId  *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
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
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSCIMServerCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteSCIMServerCredentialResponse) SetBody(v *DeleteSCIMServerCredentialResponseBody) *DeleteSCIMServerCredentialResponse {
	s.Body = v
	return s
}

type DeleteUserRequest struct {
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteUserResponse) SetBody(v *DeleteUserResponseBody) *DeleteUserResponse {
	s.Body = v
	return s
}

type DeprovisionAccessConfigurationRequest struct {
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	DirectoryId           *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	TargetId              *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetType            *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
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
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks     []*DeprovisionAccessConfigurationResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
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
	AccessConfigurationId   *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	AccessConfigurationName *string `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetId                *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetName              *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	TargetPath              *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	TargetPathName          *string `json:"TargetPathName,omitempty" xml:"TargetPathName,omitempty"`
	TargetType              *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	TaskId                  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType                *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
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
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeprovisionAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeprovisionAccessConfigurationResponse) SetBody(v *DeprovisionAccessConfigurationResponseBody) *DeprovisionAccessConfigurationResponse {
	s.Body = v
	return s
}

type DisableServiceResponseBody struct {
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DisableServiceResponse) SetBody(v *DisableServiceResponseBody) *DisableServiceResponse {
	s.Body = v
	return s
}

type EnableServiceResponseBody struct {
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *EnableServiceResponse) SetBody(v *EnableServiceResponseBody) *EnableServiceResponse {
	s.Body = v
	return s
}

type GetAccessConfigurationRequest struct {
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	DirectoryId           *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
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
	AccessConfiguration *GetAccessConfigurationResponseBodyAccessConfiguration `json:"AccessConfiguration,omitempty" xml:"AccessConfiguration,omitempty" type:"Struct"`
	RequestId           *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccessConfigurationId   *string   `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	AccessConfigurationName *string   `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	CreateTime              *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description             *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	RelayState              *string   `json:"RelayState,omitempty" xml:"RelayState,omitempty"`
	SessionDuration         *int32    `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	StatusNotifications     []*string `json:"StatusNotifications,omitempty" xml:"StatusNotifications,omitempty" type:"Repeated"`
	UpdateTime              *string   `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetAccessConfigurationResponse) SetBody(v *GetAccessConfigurationResponseBody) *GetAccessConfigurationResponse {
	s.Body = v
	return s
}

type GetDirectoryRequest struct {
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
	Directory *GetDirectoryResponseBodyDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DirectoryId   *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	DirectoryName *string `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetDirectoryResponse) SetBody(v *GetDirectoryResponseBody) *GetDirectoryResponse {
	s.Body = v
	return s
}

type GetDirectorySAMLServiceProviderInfoRequest struct {
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
	RequestId           *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AcsUrl                  *string `json:"AcsUrl,omitempty" xml:"AcsUrl,omitempty"`
	DirectoryId             *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EncodedMetadataDocument *string `json:"EncodedMetadataDocument,omitempty" xml:"EncodedMetadataDocument,omitempty"`
	EntityId                *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
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

type GetDirectorySAMLServiceProviderInfoResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDirectorySAMLServiceProviderInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetDirectorySAMLServiceProviderInfoResponse) SetBody(v *GetDirectorySAMLServiceProviderInfoResponseBody) *GetDirectorySAMLServiceProviderInfoResponse {
	s.Body = v
	return s
}

type GetDirectoryStatisticsRequest struct {
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
	DirectoryStatistics *GetDirectoryStatisticsResponseBodyDirectoryStatistics `json:"DirectoryStatistics,omitempty" xml:"DirectoryStatistics,omitempty" type:"Struct"`
	RequestId           *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccessAssignmentCount     *int32  `json:"AccessAssignmentCount,omitempty" xml:"AccessAssignmentCount,omitempty"`
	AccessConfigurationCount  *int32  `json:"AccessConfigurationCount,omitempty" xml:"AccessConfigurationCount,omitempty"`
	AccessConfigurationQuota  *int32  `json:"AccessConfigurationQuota,omitempty" xml:"AccessConfigurationQuota,omitempty"`
	DirectoryId               *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	DirectoryName             *string `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	GroupCount                *int32  `json:"GroupCount,omitempty" xml:"GroupCount,omitempty"`
	GroupQuota                *int32  `json:"GroupQuota,omitempty" xml:"GroupQuota,omitempty"`
	InProgressTaskCount       *int32  `json:"InProgressTaskCount,omitempty" xml:"InProgressTaskCount,omitempty"`
	Region                    *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SCIMServerCredentialCount *int32  `json:"SCIMServerCredentialCount,omitempty" xml:"SCIMServerCredentialCount,omitempty"`
	SCIMSyncEnabled           *bool   `json:"SCIMSyncEnabled,omitempty" xml:"SCIMSyncEnabled,omitempty"`
	SSOEnabled                *bool   `json:"SSOEnabled,omitempty" xml:"SSOEnabled,omitempty"`
	UserCount                 *int32  `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
	UserQuota                 *int32  `json:"UserQuota,omitempty" xml:"UserQuota,omitempty"`
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

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetUserCount(v int32) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.UserCount = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetUserQuota(v int32) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.UserQuota = &v
	return s
}

type GetDirectoryStatisticsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDirectoryStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetDirectoryStatisticsResponse) SetBody(v *GetDirectoryStatisticsResponseBody) *GetDirectoryStatisticsResponse {
	s.Body = v
	return s
}

type GetExternalSAMLIdentityProviderRequest struct {
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
	RequestId                         *string                                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CertificateIds          []*string `json:"CertificateIds,omitempty" xml:"CertificateIds,omitempty" type:"Repeated"`
	CreateTime              *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DirectoryId             *string   `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EncodedMetadataDocument *string   `json:"EncodedMetadataDocument,omitempty" xml:"EncodedMetadataDocument,omitempty"`
	EntityId                *string   `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	LoginUrl                *string   `json:"LoginUrl,omitempty" xml:"LoginUrl,omitempty"`
	SSOStatus               *string   `json:"SSOStatus,omitempty" xml:"SSOStatus,omitempty"`
	UpdateTime              *string   `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	WantRequestSigned       *bool     `json:"WantRequestSigned,omitempty" xml:"WantRequestSigned,omitempty"`
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
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetExternalSAMLIdentityProviderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetExternalSAMLIdentityProviderResponse) SetBody(v *GetExternalSAMLIdentityProviderResponseBody) *GetExternalSAMLIdentityProviderResponse {
	s.Body = v
	return s
}

type GetGroupRequest struct {
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
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
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetGroupResponse) SetBody(v *GetGroupResponseBody) *GetGroupResponse {
	s.Body = v
	return s
}

type GetMFAAuthenticationStatusRequest struct {
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
	MFAAuthenticationStatus *string `json:"MFAAuthenticationStatus,omitempty" xml:"MFAAuthenticationStatus,omitempty"`
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMFAAuthenticationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetMFAAuthenticationStatusResponse) SetBody(v *GetMFAAuthenticationStatusResponseBody) *GetMFAAuthenticationStatusResponse {
	s.Body = v
	return s
}

type GetSCIMSynchronizationStatusRequest struct {
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
	RequestId                 *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSCIMSynchronizationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetSCIMSynchronizationStatusResponse) SetBody(v *GetSCIMSynchronizationStatusResponseBody) *GetSCIMSynchronizationStatusResponse {
	s.Body = v
	return s
}

type GetServiceStatusResponseBody struct {
	RequestId     *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccountId               *string   `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	PrerequisiteCheckResult *string   `json:"PrerequisiteCheckResult,omitempty" xml:"PrerequisiteCheckResult,omitempty"`
	RegionsInUse            []*string `json:"RegionsInUse,omitempty" xml:"RegionsInUse,omitempty" type:"Repeated"`
	Status                  *string   `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetServiceStatusResponse) SetBody(v *GetServiceStatusResponseBody) *GetServiceStatusResponse {
	s.Body = v
	return s
}

type GetTaskRequest struct {
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Task      *GetTaskResponseBodyTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
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
	AccessConfigurationId   *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	AccessConfigurationName *string `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	EndTime                 *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FailureReason           *string `json:"FailureReason,omitempty" xml:"FailureReason,omitempty"`
	PrincipalId             *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	PrincipalName           *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	PrincipalType           *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	StartTime               *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetId                *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetName              *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	TargetPath              *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	TargetPathName          *string `json:"TargetPathName,omitempty" xml:"TargetPathName,omitempty"`
	TargetType              *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	TaskId                  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType                *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
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
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetTaskResponse) SetBody(v *GetTaskResponseBody) *GetTaskResponse {
	s.Body = v
	return s
}

type GetTaskStatusRequest struct {
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FailureReason *string `json:"FailureReason,omitempty" xml:"FailureReason,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId        *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType      *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetTaskStatusResponse) SetBody(v *GetTaskStatusResponseBody) *GetTaskStatusResponse {
	s.Body = v
	return s
}

type GetUserRequest struct {
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName   *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email         *string `json:"Email,omitempty" xml:"Email,omitempty"`
	FirstName     *string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	LastName      *string `json:"LastName,omitempty" xml:"LastName,omitempty"`
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName      *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetUserResponse) SetBody(v *GetUserResponseBody) *GetUserResponse {
	s.Body = v
	return s
}

type ListAccessAssignmentsRequest struct {
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	DirectoryId           *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	MaxResults            *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken             *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PrincipalId           *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	PrincipalType         *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	TargetId              *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetType            *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
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
	AccessAssignments []*ListAccessAssignmentsResponseBodyAccessAssignments `json:"AccessAssignments,omitempty" xml:"AccessAssignments,omitempty" type:"Repeated"`
	IsTruncated       *bool                                                 `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	MaxResults        *int32                                                `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken         *string                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId         *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCounts       *int32                                                `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
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
	AccessConfigurationId   *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	AccessConfigurationName *string `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	CreateTime              *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	PrincipalId             *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	PrincipalName           *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	PrincipalType           *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	TargetId                *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetName              *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	TargetPath              *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	TargetPathName          *string `json:"TargetPathName,omitempty" xml:"TargetPathName,omitempty"`
	TargetType              *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAccessAssignmentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListAccessAssignmentsResponse) SetBody(v *ListAccessAssignmentsResponseBody) *ListAccessAssignmentsResponse {
	s.Body = v
	return s
}

type ListAccessConfigurationProvisioningsRequest struct {
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	DirectoryId           *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	MaxResults            *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken             *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ProvisioningStatus    *string `json:"ProvisioningStatus,omitempty" xml:"ProvisioningStatus,omitempty"`
	TargetId              *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetType            *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
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
	AccessConfigurationProvisionings []*ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings `json:"AccessConfigurationProvisionings,omitempty" xml:"AccessConfigurationProvisionings,omitempty" type:"Repeated"`
	IsTruncated                      *bool                                                                               `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	MaxResults                       *int32                                                                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                        *string                                                                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId                        *string                                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCounts                      *int32                                                                              `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
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
	AccessConfigurationId   *string   `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	AccessConfigurationName *string   `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	CreateTime              *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	RAMPolicyNames          []*string `json:"RAMPolicyNames,omitempty" xml:"RAMPolicyNames,omitempty" type:"Repeated"`
	RAMRoleName             *string   `json:"RAMRoleName,omitempty" xml:"RAMRoleName,omitempty"`
	SAMLProviderName        *string   `json:"SAMLProviderName,omitempty" xml:"SAMLProviderName,omitempty"`
	Status                  *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetId                *string   `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetName              *string   `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	TargetPath              *string   `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	TargetPathName          *string   `json:"TargetPathName,omitempty" xml:"TargetPathName,omitempty"`
	TargetType              *string   `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	UpdateTime              *string   `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAccessConfigurationProvisioningsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListAccessConfigurationProvisioningsResponse) SetBody(v *ListAccessConfigurationProvisioningsResponseBody) *ListAccessConfigurationProvisioningsResponse {
	s.Body = v
	return s
}

type ListAccessConfigurationsRequest struct {
	DirectoryId         *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	Filter              *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	MaxResults          *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken           *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
	AccessConfigurations []*ListAccessConfigurationsResponseBodyAccessConfigurations `json:"AccessConfigurations,omitempty" xml:"AccessConfigurations,omitempty" type:"Repeated"`
	IsTruncated          *bool                                                       `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	MaxResults           *int32                                                      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken            *string                                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId            *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCounts          *int32                                                      `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
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
	AccessConfigurationId   *string   `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	AccessConfigurationName *string   `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	CreateTime              *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description             *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	RelayState              *string   `json:"RelayState,omitempty" xml:"RelayState,omitempty"`
	SessionDuration         *int32    `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	StatusNotifications     []*string `json:"StatusNotifications,omitempty" xml:"StatusNotifications,omitempty" type:"Repeated"`
	UpdateTime              *string   `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAccessConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListAccessConfigurationsResponse) SetBody(v *ListAccessConfigurationsResponseBody) *ListAccessConfigurationsResponse {
	s.Body = v
	return s
}

type ListDirectoriesResponseBody struct {
	Directories []*ListDirectoriesResponseBodyDirectories `json:"Directories,omitempty" xml:"Directories,omitempty" type:"Repeated"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCounts *int32                                    `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
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
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DirectoryId   *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	DirectoryName *string `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListDirectoriesResponse) SetBody(v *ListDirectoriesResponseBody) *ListDirectoriesResponse {
	s.Body = v
	return s
}

type ListExternalSAMLIdPCertificatesRequest struct {
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
	RequestId           *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SAMLIdPCertificates []*ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates `json:"SAMLIdPCertificates,omitempty" xml:"SAMLIdPCertificates,omitempty" type:"Repeated"`
	TotalCounts         *int32                                                            `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
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
	CertificateId      *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	Issuer             *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	NotAfter           *string `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	NotBefore          *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	PublicKey          *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	SerialNumber       *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	SignatureAlgorithm *string `json:"SignatureAlgorithm,omitempty" xml:"SignatureAlgorithm,omitempty"`
	Subject            *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	Version            *int32  `json:"Version,omitempty" xml:"Version,omitempty"`
	X509Certificate    *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
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
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListExternalSAMLIdPCertificatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListExternalSAMLIdPCertificatesResponse) SetBody(v *ListExternalSAMLIdPCertificatesResponseBody) *ListExternalSAMLIdPCertificatesResponse {
	s.Body = v
	return s
}

type ListGroupMembersRequest struct {
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	MaxResults  *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
	GroupMembers []*ListGroupMembersResponseBodyGroupMembers `json:"GroupMembers,omitempty" xml:"GroupMembers,omitempty" type:"Repeated"`
	IsTruncated  *bool                                       `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	MaxResults   *int32                                      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCounts  *int32                                      `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
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
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName   *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email         *string `json:"Email,omitempty" xml:"Email,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	JoinTime      *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName      *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListGroupMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListGroupMembersResponse) SetBody(v *ListGroupMembersResponseBody) *ListGroupMembersResponse {
	s.Body = v
	return s
}

type ListGroupsRequest struct {
	DirectoryId   *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	Filter        *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	MaxResults    *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
	Groups      []*ListGroupsResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	IsTruncated *bool                           `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	MaxResults  *int32                          `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId   *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCounts *int32                          `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
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
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListGroupsResponse) SetBody(v *ListGroupsResponseBody) *ListGroupsResponse {
	s.Body = v
	return s
}

type ListJoinedGroupsForUserRequest struct {
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	MaxResults  *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	IsTruncated  *bool                                              `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	JoinedGroups []*ListJoinedGroupsForUserResponseBodyJoinedGroups `json:"JoinedGroups,omitempty" xml:"JoinedGroups,omitempty" type:"Repeated"`
	MaxResults   *int32                                             `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string                                            `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCounts  *int32                                             `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
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
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	JoinTime      *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListJoinedGroupsForUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListJoinedGroupsForUserResponse) SetBody(v *ListJoinedGroupsForUserResponseBody) *ListJoinedGroupsForUserResponse {
	s.Body = v
	return s
}

type ListMFADevicesForUserRequest struct {
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	MFADevices  []*ListMFADevicesForUserResponseBodyMFADevices `json:"MFADevices,omitempty" xml:"MFADevices,omitempty" type:"Repeated"`
	RequestId   *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCounts *int32                                         `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
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
	DeviceId      *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceType    *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListMFADevicesForUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListMFADevicesForUserResponse) SetBody(v *ListMFADevicesForUserResponseBody) *ListMFADevicesForUserResponse {
	s.Body = v
	return s
}

type ListPermissionPoliciesInAccessConfigurationRequest struct {
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	DirectoryId           *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	PermissionPolicyType  *string `json:"PermissionPolicyType,omitempty" xml:"PermissionPolicyType,omitempty"`
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
	PermissionPolicies []*ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies `json:"PermissionPolicies,omitempty" xml:"PermissionPolicies,omitempty" type:"Repeated"`
	RequestId          *string                                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCounts        *int32                                                                       `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
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
	AddTime                  *string `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	PermissionPolicyDocument *string `json:"PermissionPolicyDocument,omitempty" xml:"PermissionPolicyDocument,omitempty"`
	PermissionPolicyName     *string `json:"PermissionPolicyName,omitempty" xml:"PermissionPolicyName,omitempty"`
	PermissionPolicyType     *string `json:"PermissionPolicyType,omitempty" xml:"PermissionPolicyType,omitempty"`
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
	Headers map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPermissionPoliciesInAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListPermissionPoliciesInAccessConfigurationResponse) SetBody(v *ListPermissionPoliciesInAccessConfigurationResponseBody) *ListPermissionPoliciesInAccessConfigurationResponse {
	s.Body = v
	return s
}

type ListSCIMServerCredentialsRequest struct {
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
	RequestId             *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SCIMServerCredentials []*ListSCIMServerCredentialsResponseBodySCIMServerCredentials `json:"SCIMServerCredentials,omitempty" xml:"SCIMServerCredentials,omitempty" type:"Repeated"`
	TotalCounts           *int32                                                        `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
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
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CredentialId   *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
	CredentialType *string `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	DirectoryId    *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	ExpireTime     *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSCIMServerCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListSCIMServerCredentialsResponse) SetBody(v *ListSCIMServerCredentialsResponseBody) *ListSCIMServerCredentialsResponse {
	s.Body = v
	return s
}

type ListTasksRequest struct {
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	DirectoryId           *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	Filter                *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	MaxResults            *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken             *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PrincipalId           *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	PrincipalType         *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetId              *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetType            *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	TaskType              *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
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
	IsTruncated *bool                         `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	MaxResults  *int32                        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId   *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks       []*ListTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	TotalCounts *int32                        `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
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
	AccessConfigurationId   *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	AccessConfigurationName *string `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	EndTime                 *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FailureReason           *string `json:"FailureReason,omitempty" xml:"FailureReason,omitempty"`
	PrincipalId             *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	PrincipalName           *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	PrincipalType           *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	StartTime               *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetId                *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetName              *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	TargetPath              *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	TargetPathName          *string `json:"TargetPathName,omitempty" xml:"TargetPathName,omitempty"`
	TargetType              *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	TaskId                  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType                *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
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
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTasksResponse) SetBody(v *ListTasksResponseBody) *ListTasksResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	DirectoryId   *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	Filter        *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	MaxResults    *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	IsTruncated *bool                         `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	MaxResults  *int32                        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId   *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCounts *int32                        `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
	Users       []*ListUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
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
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName   *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email         *string `json:"Email,omitempty" xml:"Email,omitempty"`
	FirstName     *string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	LastName      *string `json:"LastName,omitempty" xml:"LastName,omitempty"`
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName      *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListUsersResponse) SetBody(v *ListUsersResponseBody) *ListUsersResponse {
	s.Body = v
	return s
}

type ProvisionAccessConfigurationRequest struct {
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	DirectoryId           *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	TargetId              *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetType            *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
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
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks     []*ProvisionAccessConfigurationResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
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
	AccessConfigurationId   *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	AccessConfigurationName *string `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetId                *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetName              *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	TargetPath              *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	TargetPathName          *string `json:"TargetPathName,omitempty" xml:"TargetPathName,omitempty"`
	TargetType              *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	TaskId                  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType                *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
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
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ProvisionAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ProvisionAccessConfigurationResponse) SetBody(v *ProvisionAccessConfigurationResponseBody) *ProvisionAccessConfigurationResponse {
	s.Body = v
	return s
}

type RemoveExternalSAMLIdPCertificateRequest struct {
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	DirectoryId   *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
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
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveExternalSAMLIdPCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RemoveExternalSAMLIdPCertificateResponse) SetBody(v *RemoveExternalSAMLIdPCertificateResponseBody) *RemoveExternalSAMLIdPCertificateResponse {
	s.Body = v
	return s
}

type RemovePermissionPolicyFromAccessConfigurationRequest struct {
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	DirectoryId           *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	PermissionPolicyName  *string `json:"PermissionPolicyName,omitempty" xml:"PermissionPolicyName,omitempty"`
	PermissionPolicyType  *string `json:"PermissionPolicyType,omitempty" xml:"PermissionPolicyType,omitempty"`
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
	Headers map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemovePermissionPolicyFromAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RemovePermissionPolicyFromAccessConfigurationResponse) SetBody(v *RemovePermissionPolicyFromAccessConfigurationResponseBody) *RemovePermissionPolicyFromAccessConfigurationResponse {
	s.Body = v
	return s
}

type RemoveUserFromGroupRequest struct {
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveUserFromGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RemoveUserFromGroupResponse) SetBody(v *RemoveUserFromGroupResponseBody) *RemoveUserFromGroupResponse {
	s.Body = v
	return s
}

type ResetUserPasswordRequest struct {
	DirectoryId                      *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	GenerateRandomPassword           *bool   `json:"GenerateRandomPassword,omitempty" xml:"GenerateRandomPassword,omitempty"`
	Password                         *string `json:"Password,omitempty" xml:"Password,omitempty"`
	RequirePasswordResetForNextLogin *bool   `json:"RequirePasswordResetForNextLogin,omitempty" xml:"RequirePasswordResetForNextLogin,omitempty"`
	UserId                           *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	NewPassword *string `json:"NewPassword,omitempty" xml:"NewPassword,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetUserPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ResetUserPasswordResponse) SetBody(v *ResetUserPasswordResponseBody) *ResetUserPasswordResponse {
	s.Body = v
	return s
}

type SetExternalSAMLIdentityProviderRequest struct {
	DirectoryId             *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EncodedMetadataDocument *string `json:"EncodedMetadataDocument,omitempty" xml:"EncodedMetadataDocument,omitempty"`
	EntityId                *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	LoginUrl                *string `json:"LoginUrl,omitempty" xml:"LoginUrl,omitempty"`
	SSOStatus               *string `json:"SSOStatus,omitempty" xml:"SSOStatus,omitempty"`
	WantRequestSigned       *bool   `json:"WantRequestSigned,omitempty" xml:"WantRequestSigned,omitempty"`
	X509Certificate         *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
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
	RequestId                         *string                                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CertificateIds          []*string `json:"CertificateIds,omitempty" xml:"CertificateIds,omitempty" type:"Repeated"`
	CreateTime              *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DirectoryId             *string   `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EncodedMetadataDocument *string   `json:"EncodedMetadataDocument,omitempty" xml:"EncodedMetadataDocument,omitempty"`
	EntityId                *string   `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	LoginUrl                *string   `json:"LoginUrl,omitempty" xml:"LoginUrl,omitempty"`
	SSOStatus               *string   `json:"SSOStatus,omitempty" xml:"SSOStatus,omitempty"`
	UpdateTime              *string   `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	WantRequestSigned       *bool     `json:"WantRequestSigned,omitempty" xml:"WantRequestSigned,omitempty"`
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
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetExternalSAMLIdentityProviderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SetExternalSAMLIdentityProviderResponse) SetBody(v *SetExternalSAMLIdentityProviderResponseBody) *SetExternalSAMLIdentityProviderResponse {
	s.Body = v
	return s
}

type SetMFAAuthenticationStatusRequest struct {
	DirectoryId             *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
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
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetMFAAuthenticationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SetMFAAuthenticationStatusResponse) SetBody(v *SetMFAAuthenticationStatusResponseBody) *SetMFAAuthenticationStatusResponse {
	s.Body = v
	return s
}

type SetSCIMSynchronizationStatusRequest struct {
	DirectoryId               *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
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
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetSCIMSynchronizationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SetSCIMSynchronizationStatusResponse) SetBody(v *SetSCIMSynchronizationStatusResponseBody) *SetSCIMSynchronizationStatusResponse {
	s.Body = v
	return s
}

type UpdateAccessConfigurationRequest struct {
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	DirectoryId           *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	NewDescription        *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	NewRelayState         *string `json:"NewRelayState,omitempty" xml:"NewRelayState,omitempty"`
	NewSessionDuration    *int32  `json:"NewSessionDuration,omitempty" xml:"NewSessionDuration,omitempty"`
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
	AccessConfiguration *UpdateAccessConfigurationResponseBodyAccessConfiguration `json:"AccessConfiguration,omitempty" xml:"AccessConfiguration,omitempty" type:"Struct"`
	RequestId           *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccessConfigurationId   *string   `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	AccessConfigurationName *string   `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	CreateTime              *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description             *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	RelayState              *string   `json:"RelayState,omitempty" xml:"RelayState,omitempty"`
	SessionDuration         *int32    `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	StatusNotifications     []*string `json:"StatusNotifications,omitempty" xml:"StatusNotifications,omitempty" type:"Repeated"`
	UpdateTime              *string   `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateAccessConfigurationResponse) SetBody(v *UpdateAccessConfigurationResponseBody) *UpdateAccessConfigurationResponse {
	s.Body = v
	return s
}

type UpdateDirectoryRequest struct {
	DirectoryId      *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
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
	Directory *UpdateDirectoryResponseBodyDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DirectoryId   *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	DirectoryName *string `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateDirectoryResponse) SetBody(v *UpdateDirectoryResponseBody) *UpdateDirectoryResponse {
	s.Body = v
	return s
}

type UpdateGroupRequest struct {
	DirectoryId    *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	NewDescription *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	NewGroupName   *string `json:"NewGroupName,omitempty" xml:"NewGroupName,omitempty"`
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
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateGroupResponse) SetBody(v *UpdateGroupResponseBody) *UpdateGroupResponse {
	s.Body = v
	return s
}

type UpdateInlinePolicyForAccessConfigurationRequest struct {
	AccessConfigurationId   *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	DirectoryId             *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	InlinePolicyName        *string `json:"InlinePolicyName,omitempty" xml:"InlinePolicyName,omitempty"`
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
	Headers map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateInlinePolicyForAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateInlinePolicyForAccessConfigurationResponse) SetBody(v *UpdateInlinePolicyForAccessConfigurationResponseBody) *UpdateInlinePolicyForAccessConfigurationResponse {
	s.Body = v
	return s
}

type UpdateSCIMServerCredentialStatusRequest struct {
	CredentialId *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
	DirectoryId  *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	NewStatus    *string `json:"NewStatus,omitempty" xml:"NewStatus,omitempty"`
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
	RequestId            *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CredentialId   *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
	CredentialType *string `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	DirectoryId    *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	ExpireTime     *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSCIMServerCredentialStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateSCIMServerCredentialStatusResponse) SetBody(v *UpdateSCIMServerCredentialStatusResponseBody) *UpdateSCIMServerCredentialStatusResponse {
	s.Body = v
	return s
}

type UpdateUserRequest struct {
	DirectoryId    *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	NewDescription *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	NewDisplayName *string `json:"NewDisplayName,omitempty" xml:"NewDisplayName,omitempty"`
	NewEmail       *string `json:"NewEmail,omitempty" xml:"NewEmail,omitempty"`
	NewFirstName   *string `json:"NewFirstName,omitempty" xml:"NewFirstName,omitempty"`
	NewLastName    *string `json:"NewLastName,omitempty" xml:"NewLastName,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName   *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email         *string `json:"Email,omitempty" xml:"Email,omitempty"`
	FirstName     *string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	LastName      *string `json:"LastName,omitempty" xml:"LastName,omitempty"`
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName      *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateUserResponse) SetBody(v *UpdateUserResponseBody) *UpdateUserResponse {
	s.Body = v
	return s
}

type UpdateUserStatusRequest struct {
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	NewStatus   *string `json:"NewStatus,omitempty" xml:"NewStatus,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateUserStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (client *Client) AddExternalSAMLIdPCertificateWithOptions(request *AddExternalSAMLIdPCertificateRequest, runtime *util.RuntimeOptions) (_result *AddExternalSAMLIdPCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	query["X509Certificate"] = request.X509Certificate
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("AddExternalSAMLIdPCertificate"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) AddPermissionPolicyToAccessConfigurationWithOptions(request *AddPermissionPolicyToAccessConfigurationRequest, runtime *util.RuntimeOptions) (_result *AddPermissionPolicyToAccessConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccessConfigurationId"] = request.AccessConfigurationId
	query["DirectoryId"] = request.DirectoryId
	query["InlinePolicyDocument"] = request.InlinePolicyDocument
	query["PermissionPolicyName"] = request.PermissionPolicyName
	query["PermissionPolicyType"] = request.PermissionPolicyType
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("AddPermissionPolicyToAccessConfiguration"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) AddUserToGroupWithOptions(request *AddUserToGroupRequest, runtime *util.RuntimeOptions) (_result *AddUserToGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	query["GroupId"] = request.GroupId
	query["UserId"] = request.UserId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("AddUserToGroup"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) ClearExternalSAMLIdentityProviderWithOptions(request *ClearExternalSAMLIdentityProviderRequest, runtime *util.RuntimeOptions) (_result *ClearExternalSAMLIdentityProviderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ClearExternalSAMLIdentityProvider"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) CreateAccessAssignmentWithOptions(request *CreateAccessAssignmentRequest, runtime *util.RuntimeOptions) (_result *CreateAccessAssignmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccessConfigurationId"] = request.AccessConfigurationId
	query["DirectoryId"] = request.DirectoryId
	query["PrincipalId"] = request.PrincipalId
	query["PrincipalType"] = request.PrincipalType
	query["TargetId"] = request.TargetId
	query["TargetType"] = request.TargetType
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccessAssignment"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) CreateAccessConfigurationWithOptions(request *CreateAccessConfigurationRequest, runtime *util.RuntimeOptions) (_result *CreateAccessConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccessConfigurationName"] = request.AccessConfigurationName
	query["Description"] = request.Description
	query["DirectoryId"] = request.DirectoryId
	query["RelayState"] = request.RelayState
	query["SessionDuration"] = request.SessionDuration
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccessConfiguration"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) CreateDirectoryWithOptions(request *CreateDirectoryRequest, runtime *util.RuntimeOptions) (_result *CreateDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryName"] = request.DirectoryName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDirectory"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) CreateGroupWithOptions(request *CreateGroupRequest, runtime *util.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Description"] = request.Description
	query["DirectoryId"] = request.DirectoryId
	query["GroupName"] = request.GroupName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGroup"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) CreateSCIMServerCredentialWithOptions(request *CreateSCIMServerCredentialRequest, runtime *util.RuntimeOptions) (_result *CreateSCIMServerCredentialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSCIMServerCredential"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) CreateUserWithOptions(request *CreateUserRequest, runtime *util.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Description"] = request.Description
	query["DirectoryId"] = request.DirectoryId
	query["DisplayName"] = request.DisplayName
	query["Email"] = request.Email
	query["FirstName"] = request.FirstName
	query["LastName"] = request.LastName
	query["Status"] = request.Status
	query["UserName"] = request.UserName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUser"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DeleteAccessAssignmentWithOptions(request *DeleteAccessAssignmentRequest, runtime *util.RuntimeOptions) (_result *DeleteAccessAssignmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccessConfigurationId"] = request.AccessConfigurationId
	query["DeprovisionStrategy"] = request.DeprovisionStrategy
	query["DirectoryId"] = request.DirectoryId
	query["PrincipalId"] = request.PrincipalId
	query["PrincipalType"] = request.PrincipalType
	query["TargetId"] = request.TargetId
	query["TargetType"] = request.TargetType
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccessAssignment"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DeleteAccessConfigurationWithOptions(request *DeleteAccessConfigurationRequest, runtime *util.RuntimeOptions) (_result *DeleteAccessConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccessConfigurationId"] = request.AccessConfigurationId
	query["DirectoryId"] = request.DirectoryId
	query["ForceRemovePermissionPolicies"] = request.ForceRemovePermissionPolicies
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccessConfiguration"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DeleteDirectoryWithOptions(request *DeleteDirectoryRequest, runtime *util.RuntimeOptions) (_result *DeleteDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDirectory"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DeleteGroupWithOptions(request *DeleteGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	query["GroupId"] = request.GroupId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGroup"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DeleteMFADeviceForUserWithOptions(request *DeleteMFADeviceForUserRequest, runtime *util.RuntimeOptions) (_result *DeleteMFADeviceForUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	query["MFADeviceId"] = request.MFADeviceId
	query["UserId"] = request.UserId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMFADeviceForUser"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DeleteSCIMServerCredentialWithOptions(request *DeleteSCIMServerCredentialRequest, runtime *util.RuntimeOptions) (_result *DeleteSCIMServerCredentialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CredentialId"] = request.CredentialId
	query["DirectoryId"] = request.DirectoryId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSCIMServerCredential"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DeleteUserWithOptions(request *DeleteUserRequest, runtime *util.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	query["UserId"] = request.UserId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUser"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DeprovisionAccessConfigurationWithOptions(request *DeprovisionAccessConfigurationRequest, runtime *util.RuntimeOptions) (_result *DeprovisionAccessConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccessConfigurationId"] = request.AccessConfigurationId
	query["DirectoryId"] = request.DirectoryId
	query["TargetId"] = request.TargetId
	query["TargetType"] = request.TargetType
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeprovisionAccessConfiguration"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
		ReqBodyType: tea.String("json"),
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
		ReqBodyType: tea.String("json"),
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

func (client *Client) GetAccessConfigurationWithOptions(request *GetAccessConfigurationRequest, runtime *util.RuntimeOptions) (_result *GetAccessConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccessConfigurationId"] = request.AccessConfigurationId
	query["DirectoryId"] = request.DirectoryId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAccessConfiguration"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) GetDirectoryWithOptions(request *GetDirectoryRequest, runtime *util.RuntimeOptions) (_result *GetDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDirectory"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) GetDirectorySAMLServiceProviderInfoWithOptions(request *GetDirectorySAMLServiceProviderInfoRequest, runtime *util.RuntimeOptions) (_result *GetDirectorySAMLServiceProviderInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDirectorySAMLServiceProviderInfo"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) GetDirectoryStatisticsWithOptions(request *GetDirectoryStatisticsRequest, runtime *util.RuntimeOptions) (_result *GetDirectoryStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDirectoryStatistics"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) GetExternalSAMLIdentityProviderWithOptions(request *GetExternalSAMLIdentityProviderRequest, runtime *util.RuntimeOptions) (_result *GetExternalSAMLIdentityProviderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("GetExternalSAMLIdentityProvider"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) GetGroupWithOptions(request *GetGroupRequest, runtime *util.RuntimeOptions) (_result *GetGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	query["GroupId"] = request.GroupId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGroup"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) GetMFAAuthenticationStatusWithOptions(request *GetMFAAuthenticationStatusRequest, runtime *util.RuntimeOptions) (_result *GetMFAAuthenticationStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMFAAuthenticationStatus"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) GetSCIMSynchronizationStatusWithOptions(request *GetSCIMSynchronizationStatusRequest, runtime *util.RuntimeOptions) (_result *GetSCIMSynchronizationStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSCIMSynchronizationStatus"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
		ReqBodyType: tea.String("json"),
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

func (client *Client) GetTaskWithOptions(request *GetTaskRequest, runtime *util.RuntimeOptions) (_result *GetTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	query["TaskId"] = request.TaskId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTask"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) GetTaskStatusWithOptions(request *GetTaskStatusRequest, runtime *util.RuntimeOptions) (_result *GetTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	query["TaskId"] = request.TaskId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTaskStatus"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) GetUserWithOptions(request *GetUserRequest, runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	query["UserId"] = request.UserId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUser"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) ListAccessAssignmentsWithOptions(request *ListAccessAssignmentsRequest, runtime *util.RuntimeOptions) (_result *ListAccessAssignmentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccessConfigurationId"] = request.AccessConfigurationId
	query["DirectoryId"] = request.DirectoryId
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["PrincipalId"] = request.PrincipalId
	query["PrincipalType"] = request.PrincipalType
	query["TargetId"] = request.TargetId
	query["TargetType"] = request.TargetType
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAccessAssignments"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) ListAccessConfigurationProvisioningsWithOptions(request *ListAccessConfigurationProvisioningsRequest, runtime *util.RuntimeOptions) (_result *ListAccessConfigurationProvisioningsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccessConfigurationId"] = request.AccessConfigurationId
	query["DirectoryId"] = request.DirectoryId
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["ProvisioningStatus"] = request.ProvisioningStatus
	query["TargetId"] = request.TargetId
	query["TargetType"] = request.TargetType
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAccessConfigurationProvisionings"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) ListAccessConfigurationsWithOptions(request *ListAccessConfigurationsRequest, runtime *util.RuntimeOptions) (_result *ListAccessConfigurationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	query["Filter"] = request.Filter
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["StatusNotifications"] = request.StatusNotifications
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAccessConfigurations"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
		ReqBodyType: tea.String("json"),
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

func (client *Client) ListExternalSAMLIdPCertificatesWithOptions(request *ListExternalSAMLIdPCertificatesRequest, runtime *util.RuntimeOptions) (_result *ListExternalSAMLIdPCertificatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ListExternalSAMLIdPCertificates"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) ListGroupMembersWithOptions(request *ListGroupMembersRequest, runtime *util.RuntimeOptions) (_result *ListGroupMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	query["GroupId"] = request.GroupId
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGroupMembers"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) ListGroupsWithOptions(request *ListGroupsRequest, runtime *util.RuntimeOptions) (_result *ListGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	query["Filter"] = request.Filter
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["ProvisionType"] = request.ProvisionType
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGroups"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) ListJoinedGroupsForUserWithOptions(request *ListJoinedGroupsForUserRequest, runtime *util.RuntimeOptions) (_result *ListJoinedGroupsForUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["UserId"] = request.UserId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ListJoinedGroupsForUser"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) ListMFADevicesForUserWithOptions(request *ListMFADevicesForUserRequest, runtime *util.RuntimeOptions) (_result *ListMFADevicesForUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	query["UserId"] = request.UserId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMFADevicesForUser"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) ListPermissionPoliciesInAccessConfigurationWithOptions(request *ListPermissionPoliciesInAccessConfigurationRequest, runtime *util.RuntimeOptions) (_result *ListPermissionPoliciesInAccessConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccessConfigurationId"] = request.AccessConfigurationId
	query["DirectoryId"] = request.DirectoryId
	query["PermissionPolicyType"] = request.PermissionPolicyType
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPermissionPoliciesInAccessConfiguration"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) ListSCIMServerCredentialsWithOptions(request *ListSCIMServerCredentialsRequest, runtime *util.RuntimeOptions) (_result *ListSCIMServerCredentialsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSCIMServerCredentials"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) ListTasksWithOptions(request *ListTasksRequest, runtime *util.RuntimeOptions) (_result *ListTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccessConfigurationId"] = request.AccessConfigurationId
	query["DirectoryId"] = request.DirectoryId
	query["Filter"] = request.Filter
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["PrincipalId"] = request.PrincipalId
	query["PrincipalType"] = request.PrincipalType
	query["Status"] = request.Status
	query["TargetId"] = request.TargetId
	query["TargetType"] = request.TargetType
	query["TaskType"] = request.TaskType
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTasks"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *util.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	query["Filter"] = request.Filter
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["ProvisionType"] = request.ProvisionType
	query["Status"] = request.Status
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsers"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) ProvisionAccessConfigurationWithOptions(request *ProvisionAccessConfigurationRequest, runtime *util.RuntimeOptions) (_result *ProvisionAccessConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccessConfigurationId"] = request.AccessConfigurationId
	query["DirectoryId"] = request.DirectoryId
	query["TargetId"] = request.TargetId
	query["TargetType"] = request.TargetType
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ProvisionAccessConfiguration"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) RemoveExternalSAMLIdPCertificateWithOptions(request *RemoveExternalSAMLIdPCertificateRequest, runtime *util.RuntimeOptions) (_result *RemoveExternalSAMLIdPCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CertificateId"] = request.CertificateId
	query["DirectoryId"] = request.DirectoryId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveExternalSAMLIdPCertificate"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) RemovePermissionPolicyFromAccessConfigurationWithOptions(request *RemovePermissionPolicyFromAccessConfigurationRequest, runtime *util.RuntimeOptions) (_result *RemovePermissionPolicyFromAccessConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccessConfigurationId"] = request.AccessConfigurationId
	query["DirectoryId"] = request.DirectoryId
	query["PermissionPolicyName"] = request.PermissionPolicyName
	query["PermissionPolicyType"] = request.PermissionPolicyType
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("RemovePermissionPolicyFromAccessConfiguration"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) RemoveUserFromGroupWithOptions(request *RemoveUserFromGroupRequest, runtime *util.RuntimeOptions) (_result *RemoveUserFromGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	query["GroupId"] = request.GroupId
	query["UserId"] = request.UserId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveUserFromGroup"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) ResetUserPasswordWithOptions(request *ResetUserPasswordRequest, runtime *util.RuntimeOptions) (_result *ResetUserPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	query["GenerateRandomPassword"] = request.GenerateRandomPassword
	query["Password"] = request.Password
	query["RequirePasswordResetForNextLogin"] = request.RequirePasswordResetForNextLogin
	query["UserId"] = request.UserId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetUserPassword"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) SetExternalSAMLIdentityProviderWithOptions(request *SetExternalSAMLIdentityProviderRequest, runtime *util.RuntimeOptions) (_result *SetExternalSAMLIdentityProviderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	query["EncodedMetadataDocument"] = request.EncodedMetadataDocument
	query["EntityId"] = request.EntityId
	query["LoginUrl"] = request.LoginUrl
	query["SSOStatus"] = request.SSOStatus
	query["WantRequestSigned"] = request.WantRequestSigned
	query["X509Certificate"] = request.X509Certificate
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("SetExternalSAMLIdentityProvider"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) SetMFAAuthenticationStatusWithOptions(request *SetMFAAuthenticationStatusRequest, runtime *util.RuntimeOptions) (_result *SetMFAAuthenticationStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	query["MFAAuthenticationStatus"] = request.MFAAuthenticationStatus
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("SetMFAAuthenticationStatus"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) SetSCIMSynchronizationStatusWithOptions(request *SetSCIMSynchronizationStatusRequest, runtime *util.RuntimeOptions) (_result *SetSCIMSynchronizationStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	query["SCIMSynchronizationStatus"] = request.SCIMSynchronizationStatus
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("SetSCIMSynchronizationStatus"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) UpdateAccessConfigurationWithOptions(request *UpdateAccessConfigurationRequest, runtime *util.RuntimeOptions) (_result *UpdateAccessConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccessConfigurationId"] = request.AccessConfigurationId
	query["DirectoryId"] = request.DirectoryId
	query["NewDescription"] = request.NewDescription
	query["NewRelayState"] = request.NewRelayState
	query["NewSessionDuration"] = request.NewSessionDuration
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAccessConfiguration"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) UpdateDirectoryWithOptions(request *UpdateDirectoryRequest, runtime *util.RuntimeOptions) (_result *UpdateDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	query["NewDirectoryName"] = request.NewDirectoryName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDirectory"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) UpdateGroupWithOptions(request *UpdateGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	query["GroupId"] = request.GroupId
	query["NewDescription"] = request.NewDescription
	query["NewGroupName"] = request.NewGroupName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGroup"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) UpdateInlinePolicyForAccessConfigurationWithOptions(request *UpdateInlinePolicyForAccessConfigurationRequest, runtime *util.RuntimeOptions) (_result *UpdateInlinePolicyForAccessConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccessConfigurationId"] = request.AccessConfigurationId
	query["DirectoryId"] = request.DirectoryId
	query["InlinePolicyName"] = request.InlinePolicyName
	query["NewInlinePolicyDocument"] = request.NewInlinePolicyDocument
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInlinePolicyForAccessConfiguration"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) UpdateSCIMServerCredentialStatusWithOptions(request *UpdateSCIMServerCredentialStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateSCIMServerCredentialStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CredentialId"] = request.CredentialId
	query["DirectoryId"] = request.DirectoryId
	query["NewStatus"] = request.NewStatus
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSCIMServerCredentialStatus"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) UpdateUserWithOptions(request *UpdateUserRequest, runtime *util.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	query["NewDescription"] = request.NewDescription
	query["NewDisplayName"] = request.NewDisplayName
	query["NewEmail"] = request.NewEmail
	query["NewFirstName"] = request.NewFirstName
	query["NewLastName"] = request.NewLastName
	query["UserId"] = request.UserId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUser"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) UpdateUserStatusWithOptions(request *UpdateUserStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateUserStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DirectoryId"] = request.DirectoryId
	query["NewStatus"] = request.NewStatus
	query["UserId"] = request.UserId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUserStatus"),
		Version:     tea.String("2021-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
