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

type AcceptHandshakeRequest struct {
	// The ID of the invitation.
	//
	// You can call the [ListHandshakesForAccount](~~160006~~) operation to obtain the ID.
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
}

func (s AcceptHandshakeRequest) String() string {
	return tea.Prettify(s)
}

func (s AcceptHandshakeRequest) GoString() string {
	return s.String()
}

func (s *AcceptHandshakeRequest) SetHandshakeId(v string) *AcceptHandshakeRequest {
	s.HandshakeId = &v
	return s
}

type AcceptHandshakeResponseBody struct {
	// The information of the invitation.
	Handshake *AcceptHandshakeResponseBodyHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AcceptHandshakeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AcceptHandshakeResponseBody) GoString() string {
	return s.String()
}

func (s *AcceptHandshakeResponseBody) SetHandshake(v *AcceptHandshakeResponseBodyHandshake) *AcceptHandshakeResponseBody {
	s.Handshake = v
	return s
}

func (s *AcceptHandshakeResponseBody) SetRequestId(v string) *AcceptHandshakeResponseBody {
	s.RequestId = &v
	return s
}

type AcceptHandshakeResponseBodyHandshake struct {
	// The time when the invitation was created. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the invitation expires. The time is displayed in UTC.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the invitation.
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
	// The ID of the management account of the resource directory.
	MasterAccountId *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	// The name of the management account of the resource directory.
	MasterAccountName *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	// The time when the invitation was modified. The time is displayed in UTC.
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The comment on the invitation.
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The ID of the resource directory.
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the invitation. Valid values:
	//
	// *   Pending: The invitation is waiting for confirmation.
	// *   Accepted: The invitation is accepted.
	// *   Cancelled: The invitation is canceled.
	// *   Declined: The invitation is rejected.
	// *   Expired: The invitation expires.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID or logon email address of the invited Alibaba Cloud account.
	TargetEntity *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty"`
	// The type of the invited Alibaba Cloud account. Valid values:
	//
	// *   Account: indicates the ID of the Alibaba Cloud account.
	// *   Email: indicates the logon email address of the Alibaba Cloud account.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s AcceptHandshakeResponseBodyHandshake) String() string {
	return tea.Prettify(s)
}

func (s AcceptHandshakeResponseBodyHandshake) GoString() string {
	return s.String()
}

func (s *AcceptHandshakeResponseBodyHandshake) SetCreateTime(v string) *AcceptHandshakeResponseBodyHandshake {
	s.CreateTime = &v
	return s
}

func (s *AcceptHandshakeResponseBodyHandshake) SetExpireTime(v string) *AcceptHandshakeResponseBodyHandshake {
	s.ExpireTime = &v
	return s
}

func (s *AcceptHandshakeResponseBodyHandshake) SetHandshakeId(v string) *AcceptHandshakeResponseBodyHandshake {
	s.HandshakeId = &v
	return s
}

func (s *AcceptHandshakeResponseBodyHandshake) SetMasterAccountId(v string) *AcceptHandshakeResponseBodyHandshake {
	s.MasterAccountId = &v
	return s
}

func (s *AcceptHandshakeResponseBodyHandshake) SetMasterAccountName(v string) *AcceptHandshakeResponseBodyHandshake {
	s.MasterAccountName = &v
	return s
}

func (s *AcceptHandshakeResponseBodyHandshake) SetModifyTime(v string) *AcceptHandshakeResponseBodyHandshake {
	s.ModifyTime = &v
	return s
}

func (s *AcceptHandshakeResponseBodyHandshake) SetNote(v string) *AcceptHandshakeResponseBodyHandshake {
	s.Note = &v
	return s
}

func (s *AcceptHandshakeResponseBodyHandshake) SetResourceDirectoryId(v string) *AcceptHandshakeResponseBodyHandshake {
	s.ResourceDirectoryId = &v
	return s
}

func (s *AcceptHandshakeResponseBodyHandshake) SetStatus(v string) *AcceptHandshakeResponseBodyHandshake {
	s.Status = &v
	return s
}

func (s *AcceptHandshakeResponseBodyHandshake) SetTargetEntity(v string) *AcceptHandshakeResponseBodyHandshake {
	s.TargetEntity = &v
	return s
}

func (s *AcceptHandshakeResponseBodyHandshake) SetTargetType(v string) *AcceptHandshakeResponseBodyHandshake {
	s.TargetType = &v
	return s
}

type AcceptHandshakeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AcceptHandshakeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AcceptHandshakeResponse) String() string {
	return tea.Prettify(s)
}

func (s AcceptHandshakeResponse) GoString() string {
	return s.String()
}

func (s *AcceptHandshakeResponse) SetHeaders(v map[string]*string) *AcceptHandshakeResponse {
	s.Headers = v
	return s
}

func (s *AcceptHandshakeResponse) SetStatusCode(v int32) *AcceptHandshakeResponse {
	s.StatusCode = &v
	return s
}

func (s *AcceptHandshakeResponse) SetBody(v *AcceptHandshakeResponseBody) *AcceptHandshakeResponse {
	s.Body = v
	return s
}

type AttachControlPolicyRequest struct {
	// The ID of the access control policy.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The ID of the object to which you want to attach the access control policy. Access control policies can be attached to the following objects:
	//
	// *   Root folder
	// *   Subfolders of the Root folder
	// *   Members
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
}

func (s AttachControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *AttachControlPolicyRequest) SetPolicyId(v string) *AttachControlPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *AttachControlPolicyRequest) SetTargetId(v string) *AttachControlPolicyRequest {
	s.TargetId = &v
	return s
}

type AttachControlPolicyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *AttachControlPolicyResponseBody) SetRequestId(v string) *AttachControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type AttachControlPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *AttachControlPolicyResponse) SetHeaders(v map[string]*string) *AttachControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *AttachControlPolicyResponse) SetStatusCode(v int32) *AttachControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachControlPolicyResponse) SetBody(v *AttachControlPolicyResponseBody) *AttachControlPolicyResponse {
	s.Body = v
	return s
}

type AttachPolicyRequest struct {
	// The name of the policy.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy. Valid values:
	//
	// *   Custom: custom policy
	// *   System: system policy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The name of the object to which you want to attach the policy.
	//
	// *   If you want to attach the policy to a RAM user, specify the name in the \<UserName>@\<AccountAlias>.onaliyun.com format. \<UserName> indicates the name of the RAM user, and \<AccountAlias> indicates the alias of the Alibaba Cloud account to which the RAM user belongs.
	// *   If you want to attach the policy to a RAM user group, specify the name in the \<GroupName>@group.\<AccountAlias>.onaliyun.com format. \<GroupName> indicates the name of the RAM user group, and \<AccountAlias> indicates the alias of the Alibaba Cloud account to which the RAM user group belongs.
	// *   If you want to attach the policy to a RAM role, specify the name in the \<RoleName>@role.\<AccountAlias>.onaliyun.com format. \<RoleName> indicates the name of the RAM role, and \<AccountAlias> indicates the alias of the Alibaba Cloud account to which the RAM role belongs.
	//
	// >  The alias of an Alibaba Cloud account is a part of the default domain name. You can call the [GetDefaultDomain](~~186720~~) operation to obtain the alias of an Alibaba Cloud account.
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The type of the object to which you want to attach the policy. Valid values:
	//
	// *   IMSUser: RAM user
	// *   IMSGroup: RAM user group
	// *   ServiceRole: RAM role
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The effective scope of the policy. You can set this parameter to one of the following items:
	//
	// *   ID of a resource group: indicates that the policy takes effect for the resources in the resource group.
	// *   ID of the Alibaba Cloud account to which the authorized object belongs: indicates that the policy takes effect for the resources within the Alibaba Cloud account.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s AttachPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachPolicyRequest) GoString() string {
	return s.String()
}

func (s *AttachPolicyRequest) SetPolicyName(v string) *AttachPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *AttachPolicyRequest) SetPolicyType(v string) *AttachPolicyRequest {
	s.PolicyType = &v
	return s
}

func (s *AttachPolicyRequest) SetPrincipalName(v string) *AttachPolicyRequest {
	s.PrincipalName = &v
	return s
}

func (s *AttachPolicyRequest) SetPrincipalType(v string) *AttachPolicyRequest {
	s.PrincipalType = &v
	return s
}

func (s *AttachPolicyRequest) SetResourceGroupId(v string) *AttachPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

type AttachPolicyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *AttachPolicyResponseBody) SetRequestId(v string) *AttachPolicyResponseBody {
	s.RequestId = &v
	return s
}

type AttachPolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachPolicyResponse) GoString() string {
	return s.String()
}

func (s *AttachPolicyResponse) SetHeaders(v map[string]*string) *AttachPolicyResponse {
	s.Headers = v
	return s
}

func (s *AttachPolicyResponse) SetStatusCode(v int32) *AttachPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachPolicyResponse) SetBody(v *AttachPolicyResponseBody) *AttachPolicyResponse {
	s.Body = v
	return s
}

type BindSecureMobilePhoneRequest struct {
	// The Alibaba Cloud account ID of the member.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The mobile phone number that you want to bind to the member for security purposes.
	//
	// The mobile phone number you specify must be the same as the mobile phone number that you specify when you call the [SendVerificationCodeForBindSecureMobilePhone](~~372556~~) operation to obtain a verification code.
	//
	// Specify the mobile phone number in the \<Country code>-\<Mobile phone number> format.
	//
	// >  Mobile phone numbers in the `86-<Mobile phone number>` format in the Chinese mainland are not supported.
	SecureMobilePhone *string `json:"SecureMobilePhone,omitempty" xml:"SecureMobilePhone,omitempty"`
	// The verification code.
	//
	// You can call the [SendVerificationCodeForBindSecureMobilePhone](~~372556~~) operation to obtain the verification code.
	VerificationCode *string `json:"VerificationCode,omitempty" xml:"VerificationCode,omitempty"`
}

func (s BindSecureMobilePhoneRequest) String() string {
	return tea.Prettify(s)
}

func (s BindSecureMobilePhoneRequest) GoString() string {
	return s.String()
}

func (s *BindSecureMobilePhoneRequest) SetAccountId(v string) *BindSecureMobilePhoneRequest {
	s.AccountId = &v
	return s
}

func (s *BindSecureMobilePhoneRequest) SetSecureMobilePhone(v string) *BindSecureMobilePhoneRequest {
	s.SecureMobilePhone = &v
	return s
}

func (s *BindSecureMobilePhoneRequest) SetVerificationCode(v string) *BindSecureMobilePhoneRequest {
	s.VerificationCode = &v
	return s
}

type BindSecureMobilePhoneResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindSecureMobilePhoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindSecureMobilePhoneResponseBody) GoString() string {
	return s.String()
}

func (s *BindSecureMobilePhoneResponseBody) SetRequestId(v string) *BindSecureMobilePhoneResponseBody {
	s.RequestId = &v
	return s
}

type BindSecureMobilePhoneResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindSecureMobilePhoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindSecureMobilePhoneResponse) String() string {
	return tea.Prettify(s)
}

func (s BindSecureMobilePhoneResponse) GoString() string {
	return s.String()
}

func (s *BindSecureMobilePhoneResponse) SetHeaders(v map[string]*string) *BindSecureMobilePhoneResponse {
	s.Headers = v
	return s
}

func (s *BindSecureMobilePhoneResponse) SetStatusCode(v int32) *BindSecureMobilePhoneResponse {
	s.StatusCode = &v
	return s
}

func (s *BindSecureMobilePhoneResponse) SetBody(v *BindSecureMobilePhoneResponseBody) *BindSecureMobilePhoneResponse {
	s.Body = v
	return s
}

type CancelChangeAccountEmailRequest struct {
	// The Alibaba Cloud account ID of the member.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s CancelChangeAccountEmailRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelChangeAccountEmailRequest) GoString() string {
	return s.String()
}

func (s *CancelChangeAccountEmailRequest) SetAccountId(v string) *CancelChangeAccountEmailRequest {
	s.AccountId = &v
	return s
}

type CancelChangeAccountEmailResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelChangeAccountEmailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelChangeAccountEmailResponseBody) GoString() string {
	return s.String()
}

func (s *CancelChangeAccountEmailResponseBody) SetRequestId(v string) *CancelChangeAccountEmailResponseBody {
	s.RequestId = &v
	return s
}

type CancelChangeAccountEmailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelChangeAccountEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelChangeAccountEmailResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelChangeAccountEmailResponse) GoString() string {
	return s.String()
}

func (s *CancelChangeAccountEmailResponse) SetHeaders(v map[string]*string) *CancelChangeAccountEmailResponse {
	s.Headers = v
	return s
}

func (s *CancelChangeAccountEmailResponse) SetStatusCode(v int32) *CancelChangeAccountEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelChangeAccountEmailResponse) SetBody(v *CancelChangeAccountEmailResponseBody) *CancelChangeAccountEmailResponse {
	s.Body = v
	return s
}

type CancelCreateCloudAccountRequest struct {
	// The account record ID.
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s CancelCreateCloudAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelCreateCloudAccountRequest) GoString() string {
	return s.String()
}

func (s *CancelCreateCloudAccountRequest) SetRecordId(v string) *CancelCreateCloudAccountRequest {
	s.RecordId = &v
	return s
}

type CancelCreateCloudAccountResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelCreateCloudAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelCreateCloudAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CancelCreateCloudAccountResponseBody) SetRequestId(v string) *CancelCreateCloudAccountResponseBody {
	s.RequestId = &v
	return s
}

type CancelCreateCloudAccountResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelCreateCloudAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelCreateCloudAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelCreateCloudAccountResponse) GoString() string {
	return s.String()
}

func (s *CancelCreateCloudAccountResponse) SetHeaders(v map[string]*string) *CancelCreateCloudAccountResponse {
	s.Headers = v
	return s
}

func (s *CancelCreateCloudAccountResponse) SetStatusCode(v int32) *CancelCreateCloudAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelCreateCloudAccountResponse) SetBody(v *CancelCreateCloudAccountResponseBody) *CancelCreateCloudAccountResponse {
	s.Body = v
	return s
}

type CancelHandshakeRequest struct {
	// The ID of the invitation.
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
}

func (s CancelHandshakeRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelHandshakeRequest) GoString() string {
	return s.String()
}

func (s *CancelHandshakeRequest) SetHandshakeId(v string) *CancelHandshakeRequest {
	s.HandshakeId = &v
	return s
}

type CancelHandshakeResponseBody struct {
	// The information of the invitation.
	Handshake *CancelHandshakeResponseBodyHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelHandshakeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelHandshakeResponseBody) GoString() string {
	return s.String()
}

func (s *CancelHandshakeResponseBody) SetHandshake(v *CancelHandshakeResponseBodyHandshake) *CancelHandshakeResponseBody {
	s.Handshake = v
	return s
}

func (s *CancelHandshakeResponseBody) SetRequestId(v string) *CancelHandshakeResponseBody {
	s.RequestId = &v
	return s
}

type CancelHandshakeResponseBodyHandshake struct {
	// The time when the invitation was created. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the invitation expires. The time is displayed in UTC.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the invitation.
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
	// The ID of the management account of the resource directory.
	MasterAccountId *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	// The name of the management account of the resource directory.
	MasterAccountName *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	// The time when the invitation was modified. The time is displayed in UTC.
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The comment on the invitation.
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The ID of the resource directory.
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the invitation. Valid values:
	//
	// *   Pending: The invitation is waiting for confirmation.
	// *   Accepted: The invitation is accepted.
	// *   Cancelled: The invitation is canceled.
	// *   Declined: The invitation is rejected.
	// *   Expired: The invitation expires.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID or logon email address of the invited account.
	TargetEntity *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty"`
	// The type of the invited account. Valid values:
	//
	// *   Account: indicates the ID of the account.
	// *   Email: indicates the logon email address of the account.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s CancelHandshakeResponseBodyHandshake) String() string {
	return tea.Prettify(s)
}

func (s CancelHandshakeResponseBodyHandshake) GoString() string {
	return s.String()
}

func (s *CancelHandshakeResponseBodyHandshake) SetCreateTime(v string) *CancelHandshakeResponseBodyHandshake {
	s.CreateTime = &v
	return s
}

func (s *CancelHandshakeResponseBodyHandshake) SetExpireTime(v string) *CancelHandshakeResponseBodyHandshake {
	s.ExpireTime = &v
	return s
}

func (s *CancelHandshakeResponseBodyHandshake) SetHandshakeId(v string) *CancelHandshakeResponseBodyHandshake {
	s.HandshakeId = &v
	return s
}

func (s *CancelHandshakeResponseBodyHandshake) SetMasterAccountId(v string) *CancelHandshakeResponseBodyHandshake {
	s.MasterAccountId = &v
	return s
}

func (s *CancelHandshakeResponseBodyHandshake) SetMasterAccountName(v string) *CancelHandshakeResponseBodyHandshake {
	s.MasterAccountName = &v
	return s
}

func (s *CancelHandshakeResponseBodyHandshake) SetModifyTime(v string) *CancelHandshakeResponseBodyHandshake {
	s.ModifyTime = &v
	return s
}

func (s *CancelHandshakeResponseBodyHandshake) SetNote(v string) *CancelHandshakeResponseBodyHandshake {
	s.Note = &v
	return s
}

func (s *CancelHandshakeResponseBodyHandshake) SetResourceDirectoryId(v string) *CancelHandshakeResponseBodyHandshake {
	s.ResourceDirectoryId = &v
	return s
}

func (s *CancelHandshakeResponseBodyHandshake) SetStatus(v string) *CancelHandshakeResponseBodyHandshake {
	s.Status = &v
	return s
}

func (s *CancelHandshakeResponseBodyHandshake) SetTargetEntity(v string) *CancelHandshakeResponseBodyHandshake {
	s.TargetEntity = &v
	return s
}

func (s *CancelHandshakeResponseBodyHandshake) SetTargetType(v string) *CancelHandshakeResponseBodyHandshake {
	s.TargetType = &v
	return s
}

type CancelHandshakeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelHandshakeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelHandshakeResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelHandshakeResponse) GoString() string {
	return s.String()
}

func (s *CancelHandshakeResponse) SetHeaders(v map[string]*string) *CancelHandshakeResponse {
	s.Headers = v
	return s
}

func (s *CancelHandshakeResponse) SetStatusCode(v int32) *CancelHandshakeResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelHandshakeResponse) SetBody(v *CancelHandshakeResponseBody) *CancelHandshakeResponse {
	s.Body = v
	return s
}

type CancelPromoteResourceAccountRequest struct {
	// The account record ID.
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s CancelPromoteResourceAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelPromoteResourceAccountRequest) GoString() string {
	return s.String()
}

func (s *CancelPromoteResourceAccountRequest) SetRecordId(v string) *CancelPromoteResourceAccountRequest {
	s.RecordId = &v
	return s
}

type CancelPromoteResourceAccountResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelPromoteResourceAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelPromoteResourceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CancelPromoteResourceAccountResponseBody) SetRequestId(v string) *CancelPromoteResourceAccountResponseBody {
	s.RequestId = &v
	return s
}

type CancelPromoteResourceAccountResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelPromoteResourceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelPromoteResourceAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelPromoteResourceAccountResponse) GoString() string {
	return s.String()
}

func (s *CancelPromoteResourceAccountResponse) SetHeaders(v map[string]*string) *CancelPromoteResourceAccountResponse {
	s.Headers = v
	return s
}

func (s *CancelPromoteResourceAccountResponse) SetStatusCode(v int32) *CancelPromoteResourceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelPromoteResourceAccountResponse) SetBody(v *CancelPromoteResourceAccountResponseBody) *CancelPromoteResourceAccountResponse {
	s.Body = v
	return s
}

type ChangeAccountEmailRequest struct {
	// The Alibaba Cloud account ID of the member.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The email address to be bound to the member.
	//
	// >  The system automatically sends a verification email to the email address. After the verification is passed, the email address takes effect, and the system changes both the logon email address and secure email address of the member.
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
}

func (s ChangeAccountEmailRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeAccountEmailRequest) GoString() string {
	return s.String()
}

func (s *ChangeAccountEmailRequest) SetAccountId(v string) *ChangeAccountEmailRequest {
	s.AccountId = &v
	return s
}

func (s *ChangeAccountEmailRequest) SetEmail(v string) *ChangeAccountEmailRequest {
	s.Email = &v
	return s
}

type ChangeAccountEmailResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeAccountEmailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeAccountEmailResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeAccountEmailResponseBody) SetRequestId(v string) *ChangeAccountEmailResponseBody {
	s.RequestId = &v
	return s
}

type ChangeAccountEmailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeAccountEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeAccountEmailResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeAccountEmailResponse) GoString() string {
	return s.String()
}

func (s *ChangeAccountEmailResponse) SetHeaders(v map[string]*string) *ChangeAccountEmailResponse {
	s.Headers = v
	return s
}

func (s *ChangeAccountEmailResponse) SetStatusCode(v int32) *ChangeAccountEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeAccountEmailResponse) SetBody(v *ChangeAccountEmailResponseBody) *ChangeAccountEmailResponse {
	s.Body = v
	return s
}

type CheckAccountDeleteRequest struct {
	// The ID of the member that you want to delete.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s CheckAccountDeleteRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckAccountDeleteRequest) GoString() string {
	return s.String()
}

func (s *CheckAccountDeleteRequest) SetAccountId(v string) *CheckAccountDeleteRequest {
	s.AccountId = &v
	return s
}

type CheckAccountDeleteResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckAccountDeleteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckAccountDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAccountDeleteResponseBody) SetRequestId(v string) *CheckAccountDeleteResponseBody {
	s.RequestId = &v
	return s
}

type CheckAccountDeleteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckAccountDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckAccountDeleteResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckAccountDeleteResponse) GoString() string {
	return s.String()
}

func (s *CheckAccountDeleteResponse) SetHeaders(v map[string]*string) *CheckAccountDeleteResponse {
	s.Headers = v
	return s
}

func (s *CheckAccountDeleteResponse) SetStatusCode(v int32) *CheckAccountDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAccountDeleteResponse) SetBody(v *CheckAccountDeleteResponseBody) *CheckAccountDeleteResponse {
	s.Body = v
	return s
}

type CreateCloudAccountRequest struct {
	// The display name of the member account.
	//
	// The name must be 2 to 50 characters in length and can contain letters, digits, underscores (\_), periods (.), and hyphens (-).
	//
	// The name must be unique in the current resource directory.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address used to log on to the cloud account.
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The ID of the parent folder.
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	// The ID of the settlement account. If you do not specify this parameter, the current account is used for settlement.
	PayerAccountId *string `json:"PayerAccountId,omitempty" xml:"PayerAccountId,omitempty"`
}

func (s CreateCloudAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCloudAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudAccountRequest) SetDisplayName(v string) *CreateCloudAccountRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateCloudAccountRequest) SetEmail(v string) *CreateCloudAccountRequest {
	s.Email = &v
	return s
}

func (s *CreateCloudAccountRequest) SetParentFolderId(v string) *CreateCloudAccountRequest {
	s.ParentFolderId = &v
	return s
}

func (s *CreateCloudAccountRequest) SetPayerAccountId(v string) *CreateCloudAccountRequest {
	s.PayerAccountId = &v
	return s
}

type CreateCloudAccountResponseBody struct {
	// The information of the member account.
	Account *CreateCloudAccountResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCloudAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCloudAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCloudAccountResponseBody) SetAccount(v *CreateCloudAccountResponseBodyAccount) *CreateCloudAccountResponseBody {
	s.Account = v
	return s
}

func (s *CreateCloudAccountResponseBody) SetRequestId(v string) *CreateCloudAccountResponseBody {
	s.RequestId = &v
	return s
}

type CreateCloudAccountResponseBodyAccount struct {
	// The ID of the member account.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the member account.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The display name of the member account.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the folder.
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The way in which the member account joined the resource directory. Valid values:
	//
	// *   invited: The member account is invited to join the resource directory.
	// *   created: The member account is directly created in the resource directory.
	JoinMethod *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	// The time when the member account was modified.
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The account record ID.
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The ID of the resource directory.
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the member account. Valid values:
	//
	// *   CreateSuccess: The member account is created.
	// *   CreateVerifying: The creation of the member account is under confirmation.
	// *   CreateFailed: The member account failed to be created.
	// *   CreateExpired: The creation of the member account expired.
	// *   CreateCancelled: The creation of the member account is canceled.
	// *   PromoteVerifying: The upgrade of the member account is under confirmation.
	// *   PromoteFailed: The member account failed to be upgraded.
	// *   PromoteExpired: The upgrade of the member account expired.
	// *   PromoteCancelled: The upgrade of the member account is canceled.
	// *   PromoteSuccess: The member account is upgraded.
	// *   InviteSuccess: The owner of the member account accepted the invitation.
	// *   Removed: The member account is removed from the resource directory.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the member account. The value CloudAccount indicates that the member account is a cloud account.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateCloudAccountResponseBodyAccount) String() string {
	return tea.Prettify(s)
}

func (s CreateCloudAccountResponseBodyAccount) GoString() string {
	return s.String()
}

func (s *CreateCloudAccountResponseBodyAccount) SetAccountId(v string) *CreateCloudAccountResponseBodyAccount {
	s.AccountId = &v
	return s
}

func (s *CreateCloudAccountResponseBodyAccount) SetAccountName(v string) *CreateCloudAccountResponseBodyAccount {
	s.AccountName = &v
	return s
}

func (s *CreateCloudAccountResponseBodyAccount) SetDisplayName(v string) *CreateCloudAccountResponseBodyAccount {
	s.DisplayName = &v
	return s
}

func (s *CreateCloudAccountResponseBodyAccount) SetFolderId(v string) *CreateCloudAccountResponseBodyAccount {
	s.FolderId = &v
	return s
}

func (s *CreateCloudAccountResponseBodyAccount) SetJoinMethod(v string) *CreateCloudAccountResponseBodyAccount {
	s.JoinMethod = &v
	return s
}

func (s *CreateCloudAccountResponseBodyAccount) SetModifyTime(v string) *CreateCloudAccountResponseBodyAccount {
	s.ModifyTime = &v
	return s
}

func (s *CreateCloudAccountResponseBodyAccount) SetRecordId(v string) *CreateCloudAccountResponseBodyAccount {
	s.RecordId = &v
	return s
}

func (s *CreateCloudAccountResponseBodyAccount) SetResourceDirectoryId(v string) *CreateCloudAccountResponseBodyAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *CreateCloudAccountResponseBodyAccount) SetStatus(v string) *CreateCloudAccountResponseBodyAccount {
	s.Status = &v
	return s
}

func (s *CreateCloudAccountResponseBodyAccount) SetType(v string) *CreateCloudAccountResponseBodyAccount {
	s.Type = &v
	return s
}

type CreateCloudAccountResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCloudAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCloudAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCloudAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateCloudAccountResponse) SetHeaders(v map[string]*string) *CreateCloudAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateCloudAccountResponse) SetStatusCode(v int32) *CreateCloudAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCloudAccountResponse) SetBody(v *CreateCloudAccountResponseBody) *CreateCloudAccountResponse {
	s.Body = v
	return s
}

type CreateControlPolicyRequest struct {
	// The description of the access control policy.
	//
	// The description must be 1 to 1,024 characters in length. The description can contain letters, digits, underscores (\_), and hyphens (-) and must start with a letter.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The effective scope of the access control policy.
	//
	// The value RAM indicates that the access control policy takes effect only for RAM users and RAM roles.
	EffectScope *string `json:"EffectScope,omitempty" xml:"EffectScope,omitempty"`
	// The document of the access control policy.
	//
	// The document can be a maximum of 4,096 characters in length.
	//
	// For more information about the languages of access control policies, see [Languages of access control policies](~~179096~~).
	//
	// For more information about the examples of access control policies, see [Examples of custom access control policies](~~181474~~).
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The name of the access control policy.
	//
	// The name must be 1 to 128 characters in length. The name can contain letters, digits, and hyphens (-) and must start with a letter.
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s CreateControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateControlPolicyRequest) SetDescription(v string) *CreateControlPolicyRequest {
	s.Description = &v
	return s
}

func (s *CreateControlPolicyRequest) SetEffectScope(v string) *CreateControlPolicyRequest {
	s.EffectScope = &v
	return s
}

func (s *CreateControlPolicyRequest) SetPolicyDocument(v string) *CreateControlPolicyRequest {
	s.PolicyDocument = &v
	return s
}

func (s *CreateControlPolicyRequest) SetPolicyName(v string) *CreateControlPolicyRequest {
	s.PolicyName = &v
	return s
}

type CreateControlPolicyResponseBody struct {
	// The details of the access control policy.
	ControlPolicy *CreateControlPolicyResponseBodyControlPolicy `json:"ControlPolicy,omitempty" xml:"ControlPolicy,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateControlPolicyResponseBody) SetControlPolicy(v *CreateControlPolicyResponseBodyControlPolicy) *CreateControlPolicyResponseBody {
	s.ControlPolicy = v
	return s
}

func (s *CreateControlPolicyResponseBody) SetRequestId(v string) *CreateControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type CreateControlPolicyResponseBodyControlPolicy struct {
	// The number of times that the access control policy is referenced.
	AttachmentCount *string `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty"`
	// The time when the access control policy was created.
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the access control policy.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The effective scope of the access control policy.
	//
	// The value RAM indicates that the access control policy takes effect only for RAM users and RAM roles.
	EffectScope *string `json:"EffectScope,omitempty" xml:"EffectScope,omitempty"`
	// The ID of the access control policy.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the access control policy.
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the access control policy. Valid values:
	//
	// *   System: system access control policy
	// *   Custom: custom access control policy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The time when the access control policy was updated.
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s CreateControlPolicyResponseBodyControlPolicy) String() string {
	return tea.Prettify(s)
}

func (s CreateControlPolicyResponseBodyControlPolicy) GoString() string {
	return s.String()
}

func (s *CreateControlPolicyResponseBodyControlPolicy) SetAttachmentCount(v string) *CreateControlPolicyResponseBodyControlPolicy {
	s.AttachmentCount = &v
	return s
}

func (s *CreateControlPolicyResponseBodyControlPolicy) SetCreateDate(v string) *CreateControlPolicyResponseBodyControlPolicy {
	s.CreateDate = &v
	return s
}

func (s *CreateControlPolicyResponseBodyControlPolicy) SetDescription(v string) *CreateControlPolicyResponseBodyControlPolicy {
	s.Description = &v
	return s
}

func (s *CreateControlPolicyResponseBodyControlPolicy) SetEffectScope(v string) *CreateControlPolicyResponseBodyControlPolicy {
	s.EffectScope = &v
	return s
}

func (s *CreateControlPolicyResponseBodyControlPolicy) SetPolicyId(v string) *CreateControlPolicyResponseBodyControlPolicy {
	s.PolicyId = &v
	return s
}

func (s *CreateControlPolicyResponseBodyControlPolicy) SetPolicyName(v string) *CreateControlPolicyResponseBodyControlPolicy {
	s.PolicyName = &v
	return s
}

func (s *CreateControlPolicyResponseBodyControlPolicy) SetPolicyType(v string) *CreateControlPolicyResponseBodyControlPolicy {
	s.PolicyType = &v
	return s
}

func (s *CreateControlPolicyResponseBodyControlPolicy) SetUpdateDate(v string) *CreateControlPolicyResponseBodyControlPolicy {
	s.UpdateDate = &v
	return s
}

type CreateControlPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateControlPolicyResponse) SetHeaders(v map[string]*string) *CreateControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateControlPolicyResponse) SetStatusCode(v int32) *CreateControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateControlPolicyResponse) SetBody(v *CreateControlPolicyResponseBody) *CreateControlPolicyResponse {
	s.Body = v
	return s
}

type CreateFolderRequest struct {
	// The name of the folder.
	//
	// The name must be 1 to 24 characters in length and can contain letters, digits, underscores (\_), periods (.),and hyphens (-).
	FolderName *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
	// The ID of the parent folder.
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
}

func (s CreateFolderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFolderRequest) GoString() string {
	return s.String()
}

func (s *CreateFolderRequest) SetFolderName(v string) *CreateFolderRequest {
	s.FolderName = &v
	return s
}

func (s *CreateFolderRequest) SetParentFolderId(v string) *CreateFolderRequest {
	s.ParentFolderId = &v
	return s
}

type CreateFolderResponseBody struct {
	// The information of the folder.
	Folder *CreateFolderResponseBodyFolder `json:"Folder,omitempty" xml:"Folder,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFolderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFolderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFolderResponseBody) SetFolder(v *CreateFolderResponseBodyFolder) *CreateFolderResponseBody {
	s.Folder = v
	return s
}

func (s *CreateFolderResponseBody) SetRequestId(v string) *CreateFolderResponseBody {
	s.RequestId = &v
	return s
}

type CreateFolderResponseBodyFolder struct {
	// The time when the folder was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the folder.
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The name of the folder.
	FolderName *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
	// The ID of the parent folder.
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
}

func (s CreateFolderResponseBodyFolder) String() string {
	return tea.Prettify(s)
}

func (s CreateFolderResponseBodyFolder) GoString() string {
	return s.String()
}

func (s *CreateFolderResponseBodyFolder) SetCreateTime(v string) *CreateFolderResponseBodyFolder {
	s.CreateTime = &v
	return s
}

func (s *CreateFolderResponseBodyFolder) SetFolderId(v string) *CreateFolderResponseBodyFolder {
	s.FolderId = &v
	return s
}

func (s *CreateFolderResponseBodyFolder) SetFolderName(v string) *CreateFolderResponseBodyFolder {
	s.FolderName = &v
	return s
}

func (s *CreateFolderResponseBodyFolder) SetParentFolderId(v string) *CreateFolderResponseBodyFolder {
	s.ParentFolderId = &v
	return s
}

type CreateFolderResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFolderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFolderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFolderResponse) GoString() string {
	return s.String()
}

func (s *CreateFolderResponse) SetHeaders(v map[string]*string) *CreateFolderResponse {
	s.Headers = v
	return s
}

func (s *CreateFolderResponse) SetStatusCode(v int32) *CreateFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFolderResponse) SetBody(v *CreateFolderResponseBody) *CreateFolderResponse {
	s.Body = v
	return s
}

type CreatePolicyRequest struct {
	// The description of the policy.
	//
	// The description must be 1 to 1,024 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The document of the policy.
	//
	// The document must be 1 to 2,048 characters in length.
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The name of the policy.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
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
	// The information of the policy.
	Policy *CreatePolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// The ID of the request.
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
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The version number of the policy. Default value: v1.
	DefaultVersion *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// The description of the policy.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the policy.
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy. Valid values:
	//
	// *   Custom: custom policy
	// *   System: system policy
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
	// The document of the policy.
	//
	// The document must be 1 to 2,048 characters in length.
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The name of the policy.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// Specifies whether to set the policy version as the default version. Valid values:
	//
	// *   false: The policy version is not set as the default version.
	// *   true: The policy version is set as the default version.
	//
	// Default value: false.
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

func (s *CreatePolicyVersionRequest) SetSetAsDefault(v bool) *CreatePolicyVersionRequest {
	s.SetAsDefault = &v
	return s
}

type CreatePolicyVersionResponseBody struct {
	// The information of the policy version.
	PolicyVersion *CreatePolicyVersionResponseBodyPolicyVersion `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty" type:"Struct"`
	// The ID of the request.
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
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// Indicates whether the policy version is the default version.
	IsDefaultVersion *bool `json:"IsDefaultVersion,omitempty" xml:"IsDefaultVersion,omitempty"`
	// The ID of the policy version.
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

type CreateResourceAccountRequest struct {
	// The prefix for the Alibaba Cloud account name of the member. If you leave this parameter empty, the system randomly generates a prefix.
	//
	// The prefix must be 2 to 37 characters in length.
	//
	// The prefix can contain letters, digits, and special characters but cannot contain consecutive special characters. The prefix must start with a letter or digit and end with a letter or digit. Valid special characters include underscores (`_`), periods (.), and hyphens (`-`).
	//
	// The complete Alibaba Cloud account name of a member in a resource directory is in the \<AccountNamePrefix>@\<ResourceDirectoryId>.aliyunid.com format, such as `alice@rd-3G****.aliyunid.com`.
	//
	// Each name must be unique in the resource directory.
	AccountNamePrefix *string `json:"AccountNamePrefix,omitempty" xml:"AccountNamePrefix,omitempty"`
	// The display name of the member.
	//
	// The name must be 2 to 50 characters in length.
	//
	// The name can contain letters, digits, underscores (\_), periods (.), hyphens (-), and spaces.
	//
	// The name must be unique in the resource directory.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the parent folder.
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	// The ID of the billing account. If you leave this parameter empty, the member is used as its own billing account.
	PayerAccountId *string `json:"PayerAccountId,omitempty" xml:"PayerAccountId,omitempty"`
	// The identity type of the member. Valid values:
	//
	// *   resell: The member is an account for a reseller. This is the default value. A relationship is automatically established between the member and the reseller. The management account of the resource directory must be used as the billing account of the member.
	// *   non_resell: The member is not an account for a reseller. The member is an account that is not associated with a reseller. You can directly use the account to purchase Alibaba Cloud resources. The member is used as its own billing account.
	//
	// >  This parameter is available only for resellers at the international site (alibabacloud.com).
	ResellAccountType *string `json:"ResellAccountType,omitempty" xml:"ResellAccountType,omitempty"`
	// The tag key and value.
	Tag []*CreateResourceAccountRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateResourceAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceAccountRequest) SetAccountNamePrefix(v string) *CreateResourceAccountRequest {
	s.AccountNamePrefix = &v
	return s
}

func (s *CreateResourceAccountRequest) SetDisplayName(v string) *CreateResourceAccountRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateResourceAccountRequest) SetParentFolderId(v string) *CreateResourceAccountRequest {
	s.ParentFolderId = &v
	return s
}

func (s *CreateResourceAccountRequest) SetPayerAccountId(v string) *CreateResourceAccountRequest {
	s.PayerAccountId = &v
	return s
}

func (s *CreateResourceAccountRequest) SetResellAccountType(v string) *CreateResourceAccountRequest {
	s.ResellAccountType = &v
	return s
}

func (s *CreateResourceAccountRequest) SetTag(v []*CreateResourceAccountRequestTag) *CreateResourceAccountRequest {
	s.Tag = v
	return s
}

type CreateResourceAccountRequestTag struct {
	// A tag key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// A tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateResourceAccountRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceAccountRequestTag) GoString() string {
	return s.String()
}

func (s *CreateResourceAccountRequestTag) SetKey(v string) *CreateResourceAccountRequestTag {
	s.Key = &v
	return s
}

func (s *CreateResourceAccountRequestTag) SetValue(v string) *CreateResourceAccountRequestTag {
	s.Value = &v
	return s
}

type CreateResourceAccountResponseBody struct {
	// The information of the member.
	Account *CreateResourceAccountResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateResourceAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceAccountResponseBody) SetAccount(v *CreateResourceAccountResponseBodyAccount) *CreateResourceAccountResponseBody {
	s.Account = v
	return s
}

func (s *CreateResourceAccountResponseBody) SetRequestId(v string) *CreateResourceAccountResponseBody {
	s.RequestId = &v
	return s
}

type CreateResourceAccountResponseBodyAccount struct {
	// The Alibaba Cloud account ID of the member.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The Alibaba Cloud account name of the member.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The display name of the member.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the folder.
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The way in which the member joins the resource directory. Valid values:
	//
	// *   invited: The member is invited to join the resource directory.
	// *   created: The member is directly created in the resource directory.
	JoinMethod *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	// The time when the member joined the resource directory. The time is displayed in UTC.
	JoinTime *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// The time when the member was modified. The time is displayed in UTC.
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The ID of the resource directory.
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the member. The value CreateSuccess indicates that the member is created.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the member. The value ResourceAccount indicates that the member is a resource account.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateResourceAccountResponseBodyAccount) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceAccountResponseBodyAccount) GoString() string {
	return s.String()
}

func (s *CreateResourceAccountResponseBodyAccount) SetAccountId(v string) *CreateResourceAccountResponseBodyAccount {
	s.AccountId = &v
	return s
}

func (s *CreateResourceAccountResponseBodyAccount) SetAccountName(v string) *CreateResourceAccountResponseBodyAccount {
	s.AccountName = &v
	return s
}

func (s *CreateResourceAccountResponseBodyAccount) SetDisplayName(v string) *CreateResourceAccountResponseBodyAccount {
	s.DisplayName = &v
	return s
}

func (s *CreateResourceAccountResponseBodyAccount) SetFolderId(v string) *CreateResourceAccountResponseBodyAccount {
	s.FolderId = &v
	return s
}

func (s *CreateResourceAccountResponseBodyAccount) SetJoinMethod(v string) *CreateResourceAccountResponseBodyAccount {
	s.JoinMethod = &v
	return s
}

func (s *CreateResourceAccountResponseBodyAccount) SetJoinTime(v string) *CreateResourceAccountResponseBodyAccount {
	s.JoinTime = &v
	return s
}

func (s *CreateResourceAccountResponseBodyAccount) SetModifyTime(v string) *CreateResourceAccountResponseBodyAccount {
	s.ModifyTime = &v
	return s
}

func (s *CreateResourceAccountResponseBodyAccount) SetResourceDirectoryId(v string) *CreateResourceAccountResponseBodyAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *CreateResourceAccountResponseBodyAccount) SetStatus(v string) *CreateResourceAccountResponseBodyAccount {
	s.Status = &v
	return s
}

func (s *CreateResourceAccountResponseBodyAccount) SetType(v string) *CreateResourceAccountResponseBodyAccount {
	s.Type = &v
	return s
}

type CreateResourceAccountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceAccountResponse) SetHeaders(v map[string]*string) *CreateResourceAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceAccountResponse) SetStatusCode(v int32) *CreateResourceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceAccountResponse) SetBody(v *CreateResourceAccountResponseBody) *CreateResourceAccountResponse {
	s.Body = v
	return s
}

type CreateResourceGroupRequest struct {
	// The display name of the resource group.
	//
	// The name must be 1 to 50 characters in length.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The unique identifier of the resource group.
	//
	// The identifier must be 3 to 50 characters in length and can contain letters, digits, and hyphens (-). The identifier must start with a letter.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The tags.
	Tag []*CreateResourceGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupRequest) SetDisplayName(v string) *CreateResourceGroupRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateResourceGroupRequest) SetName(v string) *CreateResourceGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateResourceGroupRequest) SetTag(v []*CreateResourceGroupRequestTag) *CreateResourceGroupRequest {
	s.Tag = v
	return s
}

type CreateResourceGroupRequestTag struct {
	// The key of the tag.
	//
	// The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateResourceGroupRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceGroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupRequestTag) SetKey(v string) *CreateResourceGroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateResourceGroupRequestTag) SetValue(v string) *CreateResourceGroupRequestTag {
	s.Value = &v
	return s
}

type CreateResourceGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the resource group.
	ResourceGroup *CreateResourceGroupResponseBodyResourceGroup `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty" type:"Struct"`
}

func (s CreateResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupResponseBody) SetRequestId(v string) *CreateResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceGroupResponseBody) SetResourceGroup(v *CreateResourceGroupResponseBodyResourceGroup) *CreateResourceGroupResponseBody {
	s.ResourceGroup = v
	return s
}

type CreateResourceGroupResponseBodyResourceGroup struct {
	// The ID of the Alibaba Cloud account to which the resource group belongs.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The time when the resource group was created. The time is displayed in UTC.
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The display name of the resource group.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the resource group.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The unique identifier of the resource group.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the resource group in all regions.
	RegionStatuses *CreateResourceGroupResponseBodyResourceGroupRegionStatuses `json:"RegionStatuses,omitempty" xml:"RegionStatuses,omitempty" type:"Struct"`
	// The status of the resource group. Valid values:
	//
	// *   Creating: The resource group is being created.
	// *   OK: The resource group is created.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateResourceGroupResponseBodyResourceGroup) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceGroupResponseBodyResourceGroup) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupResponseBodyResourceGroup) SetAccountId(v string) *CreateResourceGroupResponseBodyResourceGroup {
	s.AccountId = &v
	return s
}

func (s *CreateResourceGroupResponseBodyResourceGroup) SetCreateDate(v string) *CreateResourceGroupResponseBodyResourceGroup {
	s.CreateDate = &v
	return s
}

func (s *CreateResourceGroupResponseBodyResourceGroup) SetDisplayName(v string) *CreateResourceGroupResponseBodyResourceGroup {
	s.DisplayName = &v
	return s
}

func (s *CreateResourceGroupResponseBodyResourceGroup) SetId(v string) *CreateResourceGroupResponseBodyResourceGroup {
	s.Id = &v
	return s
}

func (s *CreateResourceGroupResponseBodyResourceGroup) SetName(v string) *CreateResourceGroupResponseBodyResourceGroup {
	s.Name = &v
	return s
}

func (s *CreateResourceGroupResponseBodyResourceGroup) SetRegionStatuses(v *CreateResourceGroupResponseBodyResourceGroupRegionStatuses) *CreateResourceGroupResponseBodyResourceGroup {
	s.RegionStatuses = v
	return s
}

func (s *CreateResourceGroupResponseBodyResourceGroup) SetStatus(v string) *CreateResourceGroupResponseBodyResourceGroup {
	s.Status = &v
	return s
}

type CreateResourceGroupResponseBodyResourceGroupRegionStatuses struct {
	RegionStatus []*CreateResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus `json:"RegionStatus,omitempty" xml:"RegionStatus,omitempty" type:"Repeated"`
}

func (s CreateResourceGroupResponseBodyResourceGroupRegionStatuses) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceGroupResponseBodyResourceGroupRegionStatuses) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupResponseBodyResourceGroupRegionStatuses) SetRegionStatus(v []*CreateResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) *CreateResourceGroupResponseBodyResourceGroupRegionStatuses {
	s.RegionStatus = v
	return s
}

type CreateResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus struct {
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the resource group. Valid values:
	//
	// *   Creating: The resource group is being created.
	// *   OK: The resource group is created.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) SetRegionId(v string) *CreateResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus {
	s.RegionId = &v
	return s
}

func (s *CreateResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) SetStatus(v string) *CreateResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus {
	s.Status = &v
	return s
}

type CreateResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupResponse) SetHeaders(v map[string]*string) *CreateResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceGroupResponse) SetStatusCode(v int32) *CreateResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceGroupResponse) SetBody(v *CreateResourceGroupResponseBody) *CreateResourceGroupResponse {
	s.Body = v
	return s
}

type CreateRoleRequest struct {
	// The document of the policy that specifies one or more trusted entities to assume the RAM role. The trusted entities can be Alibaba Cloud accounts, Alibaba Cloud services, or identity providers (IdPs).
	//
	// >  RAM users cannot assume the RAM roles of trusted Alibaba Cloud services.
	AssumeRolePolicyDocument *string `json:"AssumeRolePolicyDocument,omitempty" xml:"AssumeRolePolicyDocument,omitempty"`
	// The description of the RAM role.
	//
	// The description must be 1 to 1,024 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum session duration of the RAM role.
	//
	// Unit: seconds. Valid values: 3600 to 43200. Default value: 3600.
	//
	// If you do not specify this parameter, the default value is used.
	MaxSessionDuration *int64 `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty"`
	// The name of the RAM role.
	//
	// The name must be 1 to 64 characters in length and can contain letters, digits, periods (.), and hyphens (-).
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
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
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the RAM role.
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
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The document of the policy that specifies the trusted entity to assume the RAM role.
	AssumeRolePolicyDocument *string `json:"AssumeRolePolicyDocument,omitempty" xml:"AssumeRolePolicyDocument,omitempty"`
	// The time when the RAM role was created.
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the RAM role.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum session duration of the RAM role.
	MaxSessionDuration *int64 `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty"`
	// The ID of the RAM role.
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// The name of the RAM role.
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The name of the RAM role after authorization.
	RolePrincipalName *string `json:"RolePrincipalName,omitempty" xml:"RolePrincipalName,omitempty"`
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

func (s *CreateRoleResponseBodyRole) SetRolePrincipalName(v string) *CreateRoleResponseBodyRole {
	s.RolePrincipalName = &v
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

type CreateServiceLinkedRoleRequest struct {
	// The suffix of the role name.
	//
	// The role name (including its suffix) must be 1 to 64 characters in length and can contain letters, digits, periods (.), and hyphens (-).
	//
	// For example, if the suffix is `Example`, the role name is `ServiceLinkedRoleName_Example`.
	CustomSuffix *string `json:"CustomSuffix,omitempty" xml:"CustomSuffix,omitempty"`
	// The description of the service-linked role.
	//
	// You must configure this parameter for service-linked roles that support custom suffixes. Otherwise, the preset value is used and cannot be modified.
	//
	// The description must be 1 to 1,024 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the service.
	//
	// For more information about the service name, see [Alibaba Cloud services that support service-linked roles](~~461722~~).
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s CreateServiceLinkedRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleRequest) SetCustomSuffix(v string) *CreateServiceLinkedRoleRequest {
	s.CustomSuffix = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) SetDescription(v string) *CreateServiceLinkedRoleRequest {
	s.Description = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) SetServiceName(v string) *CreateServiceLinkedRoleRequest {
	s.ServiceName = &v
	return s
}

type CreateServiceLinkedRoleResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the role.
	Role *CreateServiceLinkedRoleResponseBodyRole `json:"Role,omitempty" xml:"Role,omitempty" type:"Struct"`
}

func (s CreateServiceLinkedRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponseBody) SetRequestId(v string) *CreateServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBody) SetRole(v *CreateServiceLinkedRoleResponseBodyRole) *CreateServiceLinkedRoleResponseBody {
	s.Role = v
	return s
}

type CreateServiceLinkedRoleResponseBodyRole struct {
	// The Alibaba Cloud Resource Name (ARN) of the role.
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The document of the trust policy for the role.
	AssumeRolePolicyDocument *string `json:"AssumeRolePolicyDocument,omitempty" xml:"AssumeRolePolicyDocument,omitempty"`
	// The time when the role was created. The time is displayed in UTC.
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the role.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the role is a service-linked role. Valid values:
	//
	// *   true: The role is a service-linked role.
	// *   false: The role is not a service-linked role.
	IsServiceLinkedRole *bool `json:"IsServiceLinkedRole,omitempty" xml:"IsServiceLinkedRole,omitempty"`
	// The ID of the role.
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// The name of the role.
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The role name that uses a domain name as the suffix.
	RolePrincipalName *string `json:"RolePrincipalName,omitempty" xml:"RolePrincipalName,omitempty"`
}

func (s CreateServiceLinkedRoleResponseBodyRole) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleResponseBodyRole) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponseBodyRole) SetArn(v string) *CreateServiceLinkedRoleResponseBodyRole {
	s.Arn = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBodyRole) SetAssumeRolePolicyDocument(v string) *CreateServiceLinkedRoleResponseBodyRole {
	s.AssumeRolePolicyDocument = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBodyRole) SetCreateDate(v string) *CreateServiceLinkedRoleResponseBodyRole {
	s.CreateDate = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBodyRole) SetDescription(v string) *CreateServiceLinkedRoleResponseBodyRole {
	s.Description = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBodyRole) SetIsServiceLinkedRole(v bool) *CreateServiceLinkedRoleResponseBodyRole {
	s.IsServiceLinkedRole = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBodyRole) SetRoleId(v string) *CreateServiceLinkedRoleResponseBodyRole {
	s.RoleId = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBodyRole) SetRoleName(v string) *CreateServiceLinkedRoleResponseBodyRole {
	s.RoleName = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBodyRole) SetRolePrincipalName(v string) *CreateServiceLinkedRoleResponseBodyRole {
	s.RolePrincipalName = &v
	return s
}

type CreateServiceLinkedRoleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceLinkedRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponse) SetHeaders(v map[string]*string) *CreateServiceLinkedRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceLinkedRoleResponse) SetStatusCode(v int32) *CreateServiceLinkedRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceLinkedRoleResponse) SetBody(v *CreateServiceLinkedRoleResponseBody) *CreateServiceLinkedRoleResponse {
	s.Body = v
	return s
}

type DeclineHandshakeRequest struct {
	// The ID of the invitation.
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
}

func (s DeclineHandshakeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeclineHandshakeRequest) GoString() string {
	return s.String()
}

func (s *DeclineHandshakeRequest) SetHandshakeId(v string) *DeclineHandshakeRequest {
	s.HandshakeId = &v
	return s
}

type DeclineHandshakeResponseBody struct {
	// The information of the invitation.
	Handshake *DeclineHandshakeResponseBodyHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeclineHandshakeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeclineHandshakeResponseBody) GoString() string {
	return s.String()
}

func (s *DeclineHandshakeResponseBody) SetHandshake(v *DeclineHandshakeResponseBodyHandshake) *DeclineHandshakeResponseBody {
	s.Handshake = v
	return s
}

func (s *DeclineHandshakeResponseBody) SetRequestId(v string) *DeclineHandshakeResponseBody {
	s.RequestId = &v
	return s
}

type DeclineHandshakeResponseBodyHandshake struct {
	// The time when the invitation was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the invitation expires.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the invitation.
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
	// The ID of the enterprise management account of the resource directory.
	MasterAccountId *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	// The name of the enterprise management account of the resource directory.
	MasterAccountName *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	// The time when the invitation was modified.
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The invitation note.
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The ID of the resource directory.
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the invitation. Valid values:
	//
	// *   Pending: The invitation is waiting for confirmation.
	// *   Accepted: The invitation is accepted.
	// *   Cancelled: The invitation is canceled.
	// *   Declined: The invitation is rejected.
	// *   Expired: The invitation expired.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID or logon email address of the invited account.
	TargetEntity *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty"`
	// The type of the invited account. Valid values:
	//
	// *   Account: indicates the ID of the account.
	// *   Email: indicates the logon email address of the account.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s DeclineHandshakeResponseBodyHandshake) String() string {
	return tea.Prettify(s)
}

func (s DeclineHandshakeResponseBodyHandshake) GoString() string {
	return s.String()
}

func (s *DeclineHandshakeResponseBodyHandshake) SetCreateTime(v string) *DeclineHandshakeResponseBodyHandshake {
	s.CreateTime = &v
	return s
}

func (s *DeclineHandshakeResponseBodyHandshake) SetExpireTime(v string) *DeclineHandshakeResponseBodyHandshake {
	s.ExpireTime = &v
	return s
}

func (s *DeclineHandshakeResponseBodyHandshake) SetHandshakeId(v string) *DeclineHandshakeResponseBodyHandshake {
	s.HandshakeId = &v
	return s
}

func (s *DeclineHandshakeResponseBodyHandshake) SetMasterAccountId(v string) *DeclineHandshakeResponseBodyHandshake {
	s.MasterAccountId = &v
	return s
}

func (s *DeclineHandshakeResponseBodyHandshake) SetMasterAccountName(v string) *DeclineHandshakeResponseBodyHandshake {
	s.MasterAccountName = &v
	return s
}

func (s *DeclineHandshakeResponseBodyHandshake) SetModifyTime(v string) *DeclineHandshakeResponseBodyHandshake {
	s.ModifyTime = &v
	return s
}

func (s *DeclineHandshakeResponseBodyHandshake) SetNote(v string) *DeclineHandshakeResponseBodyHandshake {
	s.Note = &v
	return s
}

func (s *DeclineHandshakeResponseBodyHandshake) SetResourceDirectoryId(v string) *DeclineHandshakeResponseBodyHandshake {
	s.ResourceDirectoryId = &v
	return s
}

func (s *DeclineHandshakeResponseBodyHandshake) SetStatus(v string) *DeclineHandshakeResponseBodyHandshake {
	s.Status = &v
	return s
}

func (s *DeclineHandshakeResponseBodyHandshake) SetTargetEntity(v string) *DeclineHandshakeResponseBodyHandshake {
	s.TargetEntity = &v
	return s
}

func (s *DeclineHandshakeResponseBodyHandshake) SetTargetType(v string) *DeclineHandshakeResponseBodyHandshake {
	s.TargetType = &v
	return s
}

type DeclineHandshakeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeclineHandshakeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeclineHandshakeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeclineHandshakeResponse) GoString() string {
	return s.String()
}

func (s *DeclineHandshakeResponse) SetHeaders(v map[string]*string) *DeclineHandshakeResponse {
	s.Headers = v
	return s
}

func (s *DeclineHandshakeResponse) SetStatusCode(v int32) *DeclineHandshakeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeclineHandshakeResponse) SetBody(v *DeclineHandshakeResponseBody) *DeclineHandshakeResponse {
	s.Body = v
	return s
}

type DeleteAccountRequest struct {
	AbandonableCheckId []*string `json:"AbandonableCheckId,omitempty" xml:"AbandonableCheckId,omitempty" type:"Repeated"`
	// The type of the deletion. Valid values:
	//
	// *   0: direct deletion. If the member does not have pay-as-you-go resources that are purchased within the previous 30 days, the system directly deletes the member.
	// *   1: deletion with a silence period. If the member has pay-as-you-go resources that are purchased within the previous 30 days, the member enters a silence period of 45 days. The system starts to delete the member until the silence period ends. For more information about the silence period, see [What is the silence period for member deletion?](~~446079~~)
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s DeleteAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccountRequest) SetAbandonableCheckId(v []*string) *DeleteAccountRequest {
	s.AbandonableCheckId = v
	return s
}

func (s *DeleteAccountRequest) SetAccountId(v string) *DeleteAccountRequest {
	s.AccountId = &v
	return s
}

type DeleteAccountShrinkRequest struct {
	AbandonableCheckIdShrink *string `json:"AbandonableCheckId,omitempty" xml:"AbandonableCheckId,omitempty"`
	// The type of the deletion. Valid values:
	//
	// *   0: direct deletion. If the member does not have pay-as-you-go resources that are purchased within the previous 30 days, the system directly deletes the member.
	// *   1: deletion with a silence period. If the member has pay-as-you-go resources that are purchased within the previous 30 days, the member enters a silence period of 45 days. The system starts to delete the member until the silence period ends. For more information about the silence period, see [What is the silence period for member deletion?](~~446079~~)
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s DeleteAccountShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccountShrinkRequest) SetAbandonableCheckIdShrink(v string) *DeleteAccountShrinkRequest {
	s.AbandonableCheckIdShrink = &v
	return s
}

func (s *DeleteAccountShrinkRequest) SetAccountId(v string) *DeleteAccountShrinkRequest {
	s.AccountId = &v
	return s
}

type DeleteAccountResponseBody struct {
	DeletionType *string `json:"DeletionType,omitempty" xml:"DeletionType,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccountResponseBody) SetDeletionType(v string) *DeleteAccountResponseBody {
	s.DeletionType = &v
	return s
}

func (s *DeleteAccountResponseBody) SetRequestId(v string) *DeleteAccountResponseBody {
	s.RequestId = &v
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

type DeleteControlPolicyRequest struct {
	// The ID of the control policy.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s DeleteControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteControlPolicyRequest) SetPolicyId(v string) *DeleteControlPolicyRequest {
	s.PolicyId = &v
	return s
}

type DeleteControlPolicyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteControlPolicyResponseBody) SetRequestId(v string) *DeleteControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteControlPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteControlPolicyResponse) SetHeaders(v map[string]*string) *DeleteControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteControlPolicyResponse) SetStatusCode(v int32) *DeleteControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteControlPolicyResponse) SetBody(v *DeleteControlPolicyResponseBody) *DeleteControlPolicyResponse {
	s.Body = v
	return s
}

type DeleteFolderRequest struct {
	// The ID of the folder.
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
}

func (s DeleteFolderRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFolderRequest) GoString() string {
	return s.String()
}

func (s *DeleteFolderRequest) SetFolderId(v string) *DeleteFolderRequest {
	s.FolderId = &v
	return s
}

type DeleteFolderResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFolderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFolderResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFolderResponseBody) SetRequestId(v string) *DeleteFolderResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFolderResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFolderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFolderResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFolderResponse) GoString() string {
	return s.String()
}

func (s *DeleteFolderResponse) SetHeaders(v map[string]*string) *DeleteFolderResponse {
	s.Headers = v
	return s
}

func (s *DeleteFolderResponse) SetStatusCode(v int32) *DeleteFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFolderResponse) SetBody(v *DeleteFolderResponseBody) *DeleteFolderResponse {
	s.Body = v
	return s
}

type DeletePolicyRequest struct {
	// The name of the policy.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
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
	// The ID of the request.
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
	// The name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The ID of the policy version.
	//
	// You can call the [ListPolicyVersions](~~159982~~) operation to query the ID.
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
	// The ID of the request.
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

type DeleteResourceGroupRequest struct {
	// The ID of the resource group.
	//
	// You can call the [ListResourceGroups](~~158855~~) operation to obtain the ID.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeleteResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupRequest) SetResourceGroupId(v string) *DeleteResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

type DeleteResourceGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the resource group.
	ResourceGroup *DeleteResourceGroupResponseBodyResourceGroup `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty" type:"Struct"`
}

func (s DeleteResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupResponseBody) SetRequestId(v string) *DeleteResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceGroupResponseBody) SetResourceGroup(v *DeleteResourceGroupResponseBodyResourceGroup) *DeleteResourceGroupResponseBody {
	s.ResourceGroup = v
	return s
}

type DeleteResourceGroupResponseBodyResourceGroup struct {
	// The ID of the Alibaba Cloud account to which the resource group belongs.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The time when the resource group was created. The time is displayed in UTC.
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The display name of the resource group.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the resource group.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The unique identifier of the resource group.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the resource group in all regions.
	RegionStatuses *DeleteResourceGroupResponseBodyResourceGroupRegionStatuses `json:"RegionStatuses,omitempty" xml:"RegionStatuses,omitempty" type:"Struct"`
	// The status of the resource group. Valid values:
	//
	// *   Creating: The resource group is being created.
	// *   OK: The resource group is created.
	// *   PendingDelete: The resource group is waiting to be deleted.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteResourceGroupResponseBodyResourceGroup) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceGroupResponseBodyResourceGroup) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupResponseBodyResourceGroup) SetAccountId(v string) *DeleteResourceGroupResponseBodyResourceGroup {
	s.AccountId = &v
	return s
}

func (s *DeleteResourceGroupResponseBodyResourceGroup) SetCreateDate(v string) *DeleteResourceGroupResponseBodyResourceGroup {
	s.CreateDate = &v
	return s
}

func (s *DeleteResourceGroupResponseBodyResourceGroup) SetDisplayName(v string) *DeleteResourceGroupResponseBodyResourceGroup {
	s.DisplayName = &v
	return s
}

func (s *DeleteResourceGroupResponseBodyResourceGroup) SetId(v string) *DeleteResourceGroupResponseBodyResourceGroup {
	s.Id = &v
	return s
}

func (s *DeleteResourceGroupResponseBodyResourceGroup) SetName(v string) *DeleteResourceGroupResponseBodyResourceGroup {
	s.Name = &v
	return s
}

func (s *DeleteResourceGroupResponseBodyResourceGroup) SetRegionStatuses(v *DeleteResourceGroupResponseBodyResourceGroupRegionStatuses) *DeleteResourceGroupResponseBodyResourceGroup {
	s.RegionStatuses = v
	return s
}

func (s *DeleteResourceGroupResponseBodyResourceGroup) SetStatus(v string) *DeleteResourceGroupResponseBodyResourceGroup {
	s.Status = &v
	return s
}

type DeleteResourceGroupResponseBodyResourceGroupRegionStatuses struct {
	RegionStatus []*DeleteResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus `json:"RegionStatus,omitempty" xml:"RegionStatus,omitempty" type:"Repeated"`
}

func (s DeleteResourceGroupResponseBodyResourceGroupRegionStatuses) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceGroupResponseBodyResourceGroupRegionStatuses) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupResponseBodyResourceGroupRegionStatuses) SetRegionStatus(v []*DeleteResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) *DeleteResourceGroupResponseBodyResourceGroupRegionStatuses {
	s.RegionStatus = v
	return s
}

type DeleteResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus struct {
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the resource group. Valid values:
	//
	// *   Creating: The resource group is being created.
	// *   OK: The resource group is created.
	// *   PendingDelete: The resource group is waiting to be deleted.
	// *   Deleting: The resource group is being deleted.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) SetRegionId(v string) *DeleteResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus {
	s.RegionId = &v
	return s
}

func (s *DeleteResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) SetStatus(v string) *DeleteResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus {
	s.Status = &v
	return s
}

type DeleteResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupResponse) SetHeaders(v map[string]*string) *DeleteResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceGroupResponse) SetStatusCode(v int32) *DeleteResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceGroupResponse) SetBody(v *DeleteResourceGroupResponseBody) *DeleteResourceGroupResponse {
	s.Body = v
	return s
}

type DeleteRoleRequest struct {
	// The name of the RAM role.
	//
	// The name must be 1 to 64 characters in length and can contain letters, digits, periods (.), and hyphens (-).
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
	// The ID of the request.
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

type DeleteServiceLinkedRoleRequest struct {
	// The name of the role.
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s DeleteServiceLinkedRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceLinkedRoleRequest) SetRoleName(v string) *DeleteServiceLinkedRoleRequest {
	s.RoleName = &v
	return s
}

type DeleteServiceLinkedRoleResponseBody struct {
	// The ID of the deletion task.
	DeletionTaskId *string `json:"DeletionTaskId,omitempty" xml:"DeletionTaskId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceLinkedRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceLinkedRoleResponseBody) SetDeletionTaskId(v string) *DeleteServiceLinkedRoleResponseBody {
	s.DeletionTaskId = &v
	return s
}

func (s *DeleteServiceLinkedRoleResponseBody) SetRequestId(v string) *DeleteServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceLinkedRoleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceLinkedRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceLinkedRoleResponse) SetHeaders(v map[string]*string) *DeleteServiceLinkedRoleResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceLinkedRoleResponse) SetStatusCode(v int32) *DeleteServiceLinkedRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceLinkedRoleResponse) SetBody(v *DeleteServiceLinkedRoleResponseBody) *DeleteServiceLinkedRoleResponse {
	s.Body = v
	return s
}

type DeregisterDelegatedAdministratorRequest struct {
	// The ID of the member in the resource directory.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The identifier of the trusted service.
	//
	// For more information, see the `Trusted service identifier` column in [Supported trusted services](~~208133~~).
	ServicePrincipal *string `json:"ServicePrincipal,omitempty" xml:"ServicePrincipal,omitempty"`
}

func (s DeregisterDelegatedAdministratorRequest) String() string {
	return tea.Prettify(s)
}

func (s DeregisterDelegatedAdministratorRequest) GoString() string {
	return s.String()
}

func (s *DeregisterDelegatedAdministratorRequest) SetAccountId(v string) *DeregisterDelegatedAdministratorRequest {
	s.AccountId = &v
	return s
}

func (s *DeregisterDelegatedAdministratorRequest) SetServicePrincipal(v string) *DeregisterDelegatedAdministratorRequest {
	s.ServicePrincipal = &v
	return s
}

type DeregisterDelegatedAdministratorResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeregisterDelegatedAdministratorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeregisterDelegatedAdministratorResponseBody) GoString() string {
	return s.String()
}

func (s *DeregisterDelegatedAdministratorResponseBody) SetRequestId(v string) *DeregisterDelegatedAdministratorResponseBody {
	s.RequestId = &v
	return s
}

type DeregisterDelegatedAdministratorResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeregisterDelegatedAdministratorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeregisterDelegatedAdministratorResponse) String() string {
	return tea.Prettify(s)
}

func (s DeregisterDelegatedAdministratorResponse) GoString() string {
	return s.String()
}

func (s *DeregisterDelegatedAdministratorResponse) SetHeaders(v map[string]*string) *DeregisterDelegatedAdministratorResponse {
	s.Headers = v
	return s
}

func (s *DeregisterDelegatedAdministratorResponse) SetStatusCode(v int32) *DeregisterDelegatedAdministratorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeregisterDelegatedAdministratorResponse) SetBody(v *DeregisterDelegatedAdministratorResponseBody) *DeregisterDelegatedAdministratorResponse {
	s.Body = v
	return s
}

type DestroyResourceDirectoryResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DestroyResourceDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DestroyResourceDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *DestroyResourceDirectoryResponseBody) SetRequestId(v string) *DestroyResourceDirectoryResponseBody {
	s.RequestId = &v
	return s
}

type DestroyResourceDirectoryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DestroyResourceDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DestroyResourceDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DestroyResourceDirectoryResponse) GoString() string {
	return s.String()
}

func (s *DestroyResourceDirectoryResponse) SetHeaders(v map[string]*string) *DestroyResourceDirectoryResponse {
	s.Headers = v
	return s
}

func (s *DestroyResourceDirectoryResponse) SetStatusCode(v int32) *DestroyResourceDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DestroyResourceDirectoryResponse) SetBody(v *DestroyResourceDirectoryResponseBody) *DestroyResourceDirectoryResponse {
	s.Body = v
	return s
}

type DetachControlPolicyRequest struct {
	// The ID of the access control policy.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The ID of the object from which you want to detach the access control policy. Access control policies can be attached to the following objects:
	//
	// *   Root folder
	// *   Subfolders of the Root folder
	// *   Members
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
}

func (s DetachControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *DetachControlPolicyRequest) SetPolicyId(v string) *DetachControlPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *DetachControlPolicyRequest) SetTargetId(v string) *DetachControlPolicyRequest {
	s.TargetId = &v
	return s
}

type DetachControlPolicyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DetachControlPolicyResponseBody) SetRequestId(v string) *DetachControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DetachControlPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *DetachControlPolicyResponse) SetHeaders(v map[string]*string) *DetachControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *DetachControlPolicyResponse) SetStatusCode(v int32) *DetachControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachControlPolicyResponse) SetBody(v *DetachControlPolicyResponseBody) *DetachControlPolicyResponse {
	s.Body = v
	return s
}

type DetachPolicyRequest struct {
	// The name of the policy.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy. Valid values:
	//
	// *   Custom: custom policy
	// *   System: system policy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The name of the object to which the policy is attached.
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The type of the object to which the policy is attached. Valid values:
	//
	// *   IMSUser: RAM user
	// *   IMSGroup: RAM user group
	// *   ServiceRole: RAM role
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The ID of the resource group or the ID of the Alibaba Cloud account to which the resource group belongs.
	//
	// This parameter specifies the resource group or Alibaba Cloud account for which you want to revoke permissions.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DetachPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachPolicyRequest) GoString() string {
	return s.String()
}

func (s *DetachPolicyRequest) SetPolicyName(v string) *DetachPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *DetachPolicyRequest) SetPolicyType(v string) *DetachPolicyRequest {
	s.PolicyType = &v
	return s
}

func (s *DetachPolicyRequest) SetPrincipalName(v string) *DetachPolicyRequest {
	s.PrincipalName = &v
	return s
}

func (s *DetachPolicyRequest) SetPrincipalType(v string) *DetachPolicyRequest {
	s.PrincipalType = &v
	return s
}

func (s *DetachPolicyRequest) SetResourceGroupId(v string) *DetachPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

type DetachPolicyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DetachPolicyResponseBody) SetRequestId(v string) *DetachPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DetachPolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachPolicyResponse) GoString() string {
	return s.String()
}

func (s *DetachPolicyResponse) SetHeaders(v map[string]*string) *DetachPolicyResponse {
	s.Headers = v
	return s
}

func (s *DetachPolicyResponse) SetStatusCode(v int32) *DetachPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachPolicyResponse) SetBody(v *DetachPolicyResponseBody) *DetachPolicyResponse {
	s.Body = v
	return s
}

type DisableAssociatedTransferResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableAssociatedTransferResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableAssociatedTransferResponseBody) GoString() string {
	return s.String()
}

func (s *DisableAssociatedTransferResponseBody) SetRequestId(v string) *DisableAssociatedTransferResponseBody {
	s.RequestId = &v
	return s
}

type DisableAssociatedTransferResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableAssociatedTransferResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableAssociatedTransferResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableAssociatedTransferResponse) GoString() string {
	return s.String()
}

func (s *DisableAssociatedTransferResponse) SetHeaders(v map[string]*string) *DisableAssociatedTransferResponse {
	s.Headers = v
	return s
}

func (s *DisableAssociatedTransferResponse) SetStatusCode(v int32) *DisableAssociatedTransferResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableAssociatedTransferResponse) SetBody(v *DisableAssociatedTransferResponseBody) *DisableAssociatedTransferResponse {
	s.Body = v
	return s
}

type DisableControlPolicyResponseBody struct {
	// The status of the Control Policy feature. Valid values:
	//
	// *   Enabled: The Control Policy feature is enabled.
	// *   PendingEnable: The Control Policy feature is being enabled.
	// *   Disabled: The Control Policy feature is disabled.
	// *   PendingDisable: The Control Policy feature is being disabled.
	EnablementStatus *string `json:"EnablementStatus,omitempty" xml:"EnablementStatus,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DisableControlPolicyResponseBody) SetEnablementStatus(v string) *DisableControlPolicyResponseBody {
	s.EnablementStatus = &v
	return s
}

func (s *DisableControlPolicyResponseBody) SetRequestId(v string) *DisableControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DisableControlPolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *DisableControlPolicyResponse) SetHeaders(v map[string]*string) *DisableControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *DisableControlPolicyResponse) SetStatusCode(v int32) *DisableControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableControlPolicyResponse) SetBody(v *DisableControlPolicyResponseBody) *DisableControlPolicyResponse {
	s.Body = v
	return s
}

type EnableAssociatedTransferResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableAssociatedTransferResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableAssociatedTransferResponseBody) GoString() string {
	return s.String()
}

func (s *EnableAssociatedTransferResponseBody) SetRequestId(v string) *EnableAssociatedTransferResponseBody {
	s.RequestId = &v
	return s
}

type EnableAssociatedTransferResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableAssociatedTransferResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableAssociatedTransferResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableAssociatedTransferResponse) GoString() string {
	return s.String()
}

func (s *EnableAssociatedTransferResponse) SetHeaders(v map[string]*string) *EnableAssociatedTransferResponse {
	s.Headers = v
	return s
}

func (s *EnableAssociatedTransferResponse) SetStatusCode(v int32) *EnableAssociatedTransferResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableAssociatedTransferResponse) SetBody(v *EnableAssociatedTransferResponseBody) *EnableAssociatedTransferResponse {
	s.Body = v
	return s
}

type EnableControlPolicyResponseBody struct {
	// The status of the Control Policy feature. Valid values:
	//
	// *   Enabled: The Control Policy feature is enabled.
	// *   PendingEnable: The Control Policy feature is being enabled.
	// *   Disabled: The Control Policy feature is disabled.
	// *   PendingDisable: The Control Policy feature is being disabled.
	EnablementStatus *string `json:"EnablementStatus,omitempty" xml:"EnablementStatus,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *EnableControlPolicyResponseBody) SetEnablementStatus(v string) *EnableControlPolicyResponseBody {
	s.EnablementStatus = &v
	return s
}

func (s *EnableControlPolicyResponseBody) SetRequestId(v string) *EnableControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type EnableControlPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *EnableControlPolicyResponse) SetHeaders(v map[string]*string) *EnableControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *EnableControlPolicyResponse) SetStatusCode(v int32) *EnableControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableControlPolicyResponse) SetBody(v *EnableControlPolicyResponseBody) *EnableControlPolicyResponse {
	s.Body = v
	return s
}

type EnableResourceDirectoryRequest struct {
	// The mode in which you enable a resource directory. Valid values:
	//
	// *   CurrentAccount: indicates that the current account is used to enable a resource directory.
	// *   NewManagementAccount: indicates that a newly created account is used to enable a resource directory. If you select this mode, you must configure the `MAName`, `MASecureMobilePhone`, and `VerificationCode` parameters.
	EnableMode *string `json:"EnableMode,omitempty" xml:"EnableMode,omitempty"`
	// The name of the newly created account.
	//
	// Specify the name in the `<Prefix>@rdadmin.aliyunid.com` format. The prefix can contain letters, digits, and special characters but cannot contain consecutive special characters. The prefix must start with a letter or digit and end with a letter or digit. Valid special characters include underscores (\_), periods (.), and hyphens (-). The prefix must be 2 to 50 characters in length.
	MAName *string `json:"MAName,omitempty" xml:"MAName,omitempty"`
	// The mobile phone number that is bound to the newly created account.
	//
	// If you leave this parameter empty, the mobile phone number that is bound to the current account is used. The mobile phone number you specify must be the same as the mobile phone number that you specify when you call the [SendVerificationCodeForEnableRD](~~364248~~) operation to obtain a verification code.
	//
	// Specify the mobile phone number in the `<Country code>-<Mobile phone number>` format.
	//
	// >  Mobile phone numbers in the `86-<Mobile phone number>` format in the Chinese mainland are not supported.
	MASecureMobilePhone *string `json:"MASecureMobilePhone,omitempty" xml:"MASecureMobilePhone,omitempty"`
	// The verification code.
	//
	// You can call the [SendVerificationCodeForEnableRD](~~364248~~) operation to obtain the verification code.
	VerificationCode *string `json:"VerificationCode,omitempty" xml:"VerificationCode,omitempty"`
}

func (s EnableResourceDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableResourceDirectoryRequest) GoString() string {
	return s.String()
}

func (s *EnableResourceDirectoryRequest) SetEnableMode(v string) *EnableResourceDirectoryRequest {
	s.EnableMode = &v
	return s
}

func (s *EnableResourceDirectoryRequest) SetMAName(v string) *EnableResourceDirectoryRequest {
	s.MAName = &v
	return s
}

func (s *EnableResourceDirectoryRequest) SetMASecureMobilePhone(v string) *EnableResourceDirectoryRequest {
	s.MASecureMobilePhone = &v
	return s
}

func (s *EnableResourceDirectoryRequest) SetVerificationCode(v string) *EnableResourceDirectoryRequest {
	s.VerificationCode = &v
	return s
}

type EnableResourceDirectoryResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the resource directory.
	ResourceDirectory *EnableResourceDirectoryResponseBodyResourceDirectory `json:"ResourceDirectory,omitempty" xml:"ResourceDirectory,omitempty" type:"Struct"`
}

func (s EnableResourceDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableResourceDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *EnableResourceDirectoryResponseBody) SetRequestId(v string) *EnableResourceDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableResourceDirectoryResponseBody) SetResourceDirectory(v *EnableResourceDirectoryResponseBodyResourceDirectory) *EnableResourceDirectoryResponseBody {
	s.ResourceDirectory = v
	return s
}

type EnableResourceDirectoryResponseBodyResourceDirectory struct {
	// The time when the resource directory was enabled.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the management account.
	MasterAccountId *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	// The name of the management account.
	MasterAccountName *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	// The ID of the resource directory.
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The ID of the Root folder.
	RootFolderId *string `json:"RootFolderId,omitempty" xml:"RootFolderId,omitempty"`
}

func (s EnableResourceDirectoryResponseBodyResourceDirectory) String() string {
	return tea.Prettify(s)
}

func (s EnableResourceDirectoryResponseBodyResourceDirectory) GoString() string {
	return s.String()
}

func (s *EnableResourceDirectoryResponseBodyResourceDirectory) SetCreateTime(v string) *EnableResourceDirectoryResponseBodyResourceDirectory {
	s.CreateTime = &v
	return s
}

func (s *EnableResourceDirectoryResponseBodyResourceDirectory) SetMasterAccountId(v string) *EnableResourceDirectoryResponseBodyResourceDirectory {
	s.MasterAccountId = &v
	return s
}

func (s *EnableResourceDirectoryResponseBodyResourceDirectory) SetMasterAccountName(v string) *EnableResourceDirectoryResponseBodyResourceDirectory {
	s.MasterAccountName = &v
	return s
}

func (s *EnableResourceDirectoryResponseBodyResourceDirectory) SetResourceDirectoryId(v string) *EnableResourceDirectoryResponseBodyResourceDirectory {
	s.ResourceDirectoryId = &v
	return s
}

func (s *EnableResourceDirectoryResponseBodyResourceDirectory) SetRootFolderId(v string) *EnableResourceDirectoryResponseBodyResourceDirectory {
	s.RootFolderId = &v
	return s
}

type EnableResourceDirectoryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableResourceDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableResourceDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableResourceDirectoryResponse) GoString() string {
	return s.String()
}

func (s *EnableResourceDirectoryResponse) SetHeaders(v map[string]*string) *EnableResourceDirectoryResponse {
	s.Headers = v
	return s
}

func (s *EnableResourceDirectoryResponse) SetStatusCode(v int32) *EnableResourceDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableResourceDirectoryResponse) SetBody(v *EnableResourceDirectoryResponseBody) *EnableResourceDirectoryResponse {
	s.Body = v
	return s
}

type GetAccountRequest struct {
	// The Alibaba Cloud account ID of the member.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Specifies whether to return the information of tags. Valid values:
	//
	// *   false (default value)
	// *   true
	IncludeTags *bool `json:"IncludeTags,omitempty" xml:"IncludeTags,omitempty"`
}

func (s GetAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccountRequest) GoString() string {
	return s.String()
}

func (s *GetAccountRequest) SetAccountId(v string) *GetAccountRequest {
	s.AccountId = &v
	return s
}

func (s *GetAccountRequest) SetIncludeTags(v bool) *GetAccountRequest {
	s.IncludeTags = &v
	return s
}

type GetAccountResponseBody struct {
	// The information of the member.
	Account *GetAccountResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccountResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountResponseBody) SetAccount(v *GetAccountResponseBodyAccount) *GetAccountResponseBody {
	s.Account = v
	return s
}

func (s *GetAccountResponseBody) SetRequestId(v string) *GetAccountResponseBody {
	s.RequestId = &v
	return s
}

type GetAccountResponseBodyAccount struct {
	// The Alibaba Cloud account ID of the member.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The Alibaba Cloud account name of the member.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The display name of the member.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The status of the modification for the email address bound to the member. Valid values:
	//
	// *   WAIT_MODIFY: in progress
	// *   CANCELLED: canceled
	// *   EXPIRED: expired
	//
	// If the value of this parameter is empty, no modification is performed for the email address.
	EmailStatus *string `json:"EmailStatus,omitempty" xml:"EmailStatus,omitempty"`
	// The ID of the folder.
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The real-name verification information.
	IdentityInformation *string `json:"IdentityInformation,omitempty" xml:"IdentityInformation,omitempty"`
	// The way in which the member joins the resource directory. Valid values:
	//
	// *   invited: The member is invited to join the resource directory.
	// *   created: The member is directly created in the resource directory.
	JoinMethod *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	// The time when the member joined the resource directory.
	JoinTime *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// The location of the member in the resource directory.
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The time when the member was modified.
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The ID of the resource directory.
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The path of the member in the resource directory.
	ResourceDirectoryPath *string `json:"ResourceDirectoryPath,omitempty" xml:"ResourceDirectoryPath,omitempty"`
	// The status of the member. Valid values:
	//
	// *   CreateSuccess: The member is created.
	// *   PromoteVerifying: The upgrade of the member is being confirmed.
	// *   PromoteFailed: The upgrade of the member fails.
	// *   PromoteExpired: The upgrade of the member expires.
	// *   PromoteCancelled: The upgrade of the member is canceled.
	// *   PromoteSuccess: The member is upgraded.
	// *   InviteSuccess: The member accepts the invitation.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags that are added to the member.
	Tags []*GetAccountResponseBodyAccountTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of the member. Valid values:
	//
	// *   CloudAccount: cloud account
	// *   ResourceAccount: resource account
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAccountResponseBodyAccount) String() string {
	return tea.Prettify(s)
}

func (s GetAccountResponseBodyAccount) GoString() string {
	return s.String()
}

func (s *GetAccountResponseBodyAccount) SetAccountId(v string) *GetAccountResponseBodyAccount {
	s.AccountId = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetAccountName(v string) *GetAccountResponseBodyAccount {
	s.AccountName = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetDisplayName(v string) *GetAccountResponseBodyAccount {
	s.DisplayName = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetEmailStatus(v string) *GetAccountResponseBodyAccount {
	s.EmailStatus = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetFolderId(v string) *GetAccountResponseBodyAccount {
	s.FolderId = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetIdentityInformation(v string) *GetAccountResponseBodyAccount {
	s.IdentityInformation = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetJoinMethod(v string) *GetAccountResponseBodyAccount {
	s.JoinMethod = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetJoinTime(v string) *GetAccountResponseBodyAccount {
	s.JoinTime = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetLocation(v string) *GetAccountResponseBodyAccount {
	s.Location = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetModifyTime(v string) *GetAccountResponseBodyAccount {
	s.ModifyTime = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetResourceDirectoryId(v string) *GetAccountResponseBodyAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetResourceDirectoryPath(v string) *GetAccountResponseBodyAccount {
	s.ResourceDirectoryPath = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetStatus(v string) *GetAccountResponseBodyAccount {
	s.Status = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetTags(v []*GetAccountResponseBodyAccountTags) *GetAccountResponseBodyAccount {
	s.Tags = v
	return s
}

func (s *GetAccountResponseBodyAccount) SetType(v string) *GetAccountResponseBodyAccount {
	s.Type = &v
	return s
}

type GetAccountResponseBodyAccountTags struct {
	// A tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// A tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetAccountResponseBodyAccountTags) String() string {
	return tea.Prettify(s)
}

func (s GetAccountResponseBodyAccountTags) GoString() string {
	return s.String()
}

func (s *GetAccountResponseBodyAccountTags) SetKey(v string) *GetAccountResponseBodyAccountTags {
	s.Key = &v
	return s
}

func (s *GetAccountResponseBodyAccountTags) SetValue(v string) *GetAccountResponseBodyAccountTags {
	s.Value = &v
	return s
}

type GetAccountResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccountResponse) GoString() string {
	return s.String()
}

func (s *GetAccountResponse) SetHeaders(v map[string]*string) *GetAccountResponse {
	s.Headers = v
	return s
}

func (s *GetAccountResponse) SetStatusCode(v int32) *GetAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountResponse) SetBody(v *GetAccountResponseBody) *GetAccountResponse {
	s.Body = v
	return s
}

type GetAccountDeletionCheckResultRequest struct {
	// The ID of the member that you want to delete.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s GetAccountDeletionCheckResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccountDeletionCheckResultRequest) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionCheckResultRequest) SetAccountId(v string) *GetAccountDeletionCheckResultRequest {
	s.AccountId = &v
	return s
}

type GetAccountDeletionCheckResultResponseBody struct {
	// The result of the deletion check for the member.
	AccountDeletionCheckResultInfo *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo `json:"AccountDeletionCheckResultInfo,omitempty" xml:"AccountDeletionCheckResultInfo,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccountDeletionCheckResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccountDeletionCheckResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionCheckResultResponseBody) SetAccountDeletionCheckResultInfo(v *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo) *GetAccountDeletionCheckResultResponseBody {
	s.AccountDeletionCheckResultInfo = v
	return s
}

func (s *GetAccountDeletionCheckResultResponseBody) SetRequestId(v string) *GetAccountDeletionCheckResultResponseBody {
	s.RequestId = &v
	return s
}

type GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo struct {
	// The check items that you can choose to ignore for the member deletion.
	//
	// >  This parameter may be returned if the value of AllowDelete is true.
	AbandonableChecks []*GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks `json:"AbandonableChecks,omitempty" xml:"AbandonableChecks,omitempty" type:"Repeated"`
	// Indicates whether the member can be deleted. Valid values:
	//
	// *   true: The member can be deleted.
	// *   false: The member cannot be deleted.
	AllowDelete *string `json:"AllowDelete,omitempty" xml:"AllowDelete,omitempty"`
	// The reasons why the member cannot be deleted.
	//
	// >  This parameter is returned only if the value of AllowDelete is false.
	NotAllowReason []*GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason `json:"NotAllowReason,omitempty" xml:"NotAllowReason,omitempty" type:"Repeated"`
	// The status of the check. Valid values:
	//
	// *   PreCheckComplete: The check is complete.
	// *   PreChecking: The check is in progress.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo) SetAbandonableChecks(v []*GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo {
	s.AbandonableChecks = v
	return s
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo) SetAllowDelete(v string) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo {
	s.AllowDelete = &v
	return s
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo) SetNotAllowReason(v []*GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo {
	s.NotAllowReason = v
	return s
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo) SetStatus(v string) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo {
	s.Status = &v
	return s
}

type GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks struct {
	// The ID of the check item.
	CheckId *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The name of the cloud service to which the check item belongs.
	CheckName *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	// The description of the check item.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks) String() string {
	return tea.Prettify(s)
}

func (s GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks) SetCheckId(v string) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks {
	s.CheckId = &v
	return s
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks) SetCheckName(v string) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks {
	s.CheckName = &v
	return s
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks) SetDescription(v string) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks {
	s.Description = &v
	return s
}

type GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason struct {
	// The ID of the check item.
	CheckId *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The name of the cloud service to which the check item belongs.
	CheckName *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	// The description of the check item.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason) String() string {
	return tea.Prettify(s)
}

func (s GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason) SetCheckId(v string) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason {
	s.CheckId = &v
	return s
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason) SetCheckName(v string) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason {
	s.CheckName = &v
	return s
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason) SetDescription(v string) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason {
	s.Description = &v
	return s
}

type GetAccountDeletionCheckResultResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountDeletionCheckResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountDeletionCheckResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccountDeletionCheckResultResponse) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionCheckResultResponse) SetHeaders(v map[string]*string) *GetAccountDeletionCheckResultResponse {
	s.Headers = v
	return s
}

func (s *GetAccountDeletionCheckResultResponse) SetStatusCode(v int32) *GetAccountDeletionCheckResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountDeletionCheckResultResponse) SetBody(v *GetAccountDeletionCheckResultResponseBody) *GetAccountDeletionCheckResultResponse {
	s.Body = v
	return s
}

type GetAccountDeletionStatusRequest struct {
	// The Alibaba Cloud account ID of the member.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s GetAccountDeletionStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccountDeletionStatusRequest) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionStatusRequest) SetAccountId(v string) *GetAccountDeletionStatusRequest {
	s.AccountId = &v
	return s
}

type GetAccountDeletionStatusResponseBody struct {
	// The deletion status of the member.
	RdAccountDeletionStatus *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus `json:"RdAccountDeletionStatus,omitempty" xml:"RdAccountDeletionStatus,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccountDeletionStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccountDeletionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionStatusResponseBody) SetRdAccountDeletionStatus(v *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) *GetAccountDeletionStatusResponseBody {
	s.RdAccountDeletionStatus = v
	return s
}

func (s *GetAccountDeletionStatusResponseBody) SetRequestId(v string) *GetAccountDeletionStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus struct {
	// The Alibaba Cloud account ID of the member.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The start time of the deletion.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The end time of the deletion.
	DeletionTime *string `json:"DeletionTime,omitempty" xml:"DeletionTime,omitempty"`
	// The type of the deletion. Valid values:
	//
	// *   0: direct deletion. If the member does not have pay-as-you-go resources that are purchased within the previous 30 days, the system directly deletes the member.
	// *   1: deletion with a silence period. If the member has pay-as-you-go resources that are purchased within the previous 30 days, the member enters a silence period. The system starts to delete the member until the silence period ends. For more information about the silence period, see [What is the silence period for member deletion?](~~446079~~)
	DeletionType *string `json:"DeletionType,omitempty" xml:"DeletionType,omitempty"`
	// The reasons why the member fails to be deleted.
	FailReasonList []*GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList `json:"FailReasonList,omitempty" xml:"FailReasonList,omitempty" type:"Repeated"`
	// The status. Valid values:
	//
	// *   Success: The member is deleted.
	// *   Checking: A deletion check is being performed for the member.
	// *   Deleting: The member is being deleted.
	// *   CheckFailed: The deletion check for the member fails.
	// *   DeleteFailed: The member fails to be deleted.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) String() string {
	return tea.Prettify(s)
}

func (s GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) SetAccountId(v string) *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus {
	s.AccountId = &v
	return s
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) SetCreateTime(v string) *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus {
	s.CreateTime = &v
	return s
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) SetDeletionTime(v string) *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus {
	s.DeletionTime = &v
	return s
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) SetDeletionType(v string) *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus {
	s.DeletionType = &v
	return s
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) SetFailReasonList(v []*GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList) *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus {
	s.FailReasonList = v
	return s
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) SetStatus(v string) *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus {
	s.Status = &v
	return s
}

type GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList struct {
	// The description of the check item.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the cloud service to which the check item belongs.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList) String() string {
	return tea.Prettify(s)
}

func (s GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList) SetDescription(v string) *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList {
	s.Description = &v
	return s
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList) SetName(v string) *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList {
	s.Name = &v
	return s
}

type GetAccountDeletionStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountDeletionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountDeletionStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccountDeletionStatusResponse) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionStatusResponse) SetHeaders(v map[string]*string) *GetAccountDeletionStatusResponse {
	s.Headers = v
	return s
}

func (s *GetAccountDeletionStatusResponse) SetStatusCode(v int32) *GetAccountDeletionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountDeletionStatusResponse) SetBody(v *GetAccountDeletionStatusResponseBody) *GetAccountDeletionStatusResponse {
	s.Body = v
	return s
}

type GetControlPolicyRequest struct {
	// The language in which you want to return the description of the access control policy. Valid values:
	//
	// *   zh-CN (default value): Chinese
	// *   en: English
	// *   ja: Japanese
	//
	// >  This parameter is valid only for system access control policies.
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The ID of the access control policy.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s GetControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetControlPolicyRequest) SetLanguage(v string) *GetControlPolicyRequest {
	s.Language = &v
	return s
}

func (s *GetControlPolicyRequest) SetPolicyId(v string) *GetControlPolicyRequest {
	s.PolicyId = &v
	return s
}

type GetControlPolicyResponseBody struct {
	// The details of the access control policy.
	ControlPolicy *GetControlPolicyResponseBodyControlPolicy `json:"ControlPolicy,omitempty" xml:"ControlPolicy,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetControlPolicyResponseBody) SetControlPolicy(v *GetControlPolicyResponseBodyControlPolicy) *GetControlPolicyResponseBody {
	s.ControlPolicy = v
	return s
}

func (s *GetControlPolicyResponseBody) SetRequestId(v string) *GetControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type GetControlPolicyResponseBodyControlPolicy struct {
	// The number of times that the access control policy is referenced.
	AttachmentCount *string `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty"`
	// The time when the access control policy was created.
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the access control policy.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The effective scope of the access control policy. Valid values:
	//
	// *   All: The access control policy is in effect for Alibaba Cloud accounts, RAM users, and RAM roles.
	// *   RAM: The access control policy is in effect only for RAM users and RAM roles.
	EffectScope *string `json:"EffectScope,omitempty" xml:"EffectScope,omitempty"`
	// The document of the access control policy.
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The ID of the access control policy.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the access control policy.
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the access control policy. Valid values:
	//
	// *   System: system access control policy
	// *   Custom: custom access control policy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The time when the access control policy was updated.
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s GetControlPolicyResponseBodyControlPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetControlPolicyResponseBodyControlPolicy) GoString() string {
	return s.String()
}

func (s *GetControlPolicyResponseBodyControlPolicy) SetAttachmentCount(v string) *GetControlPolicyResponseBodyControlPolicy {
	s.AttachmentCount = &v
	return s
}

func (s *GetControlPolicyResponseBodyControlPolicy) SetCreateDate(v string) *GetControlPolicyResponseBodyControlPolicy {
	s.CreateDate = &v
	return s
}

func (s *GetControlPolicyResponseBodyControlPolicy) SetDescription(v string) *GetControlPolicyResponseBodyControlPolicy {
	s.Description = &v
	return s
}

func (s *GetControlPolicyResponseBodyControlPolicy) SetEffectScope(v string) *GetControlPolicyResponseBodyControlPolicy {
	s.EffectScope = &v
	return s
}

func (s *GetControlPolicyResponseBodyControlPolicy) SetPolicyDocument(v string) *GetControlPolicyResponseBodyControlPolicy {
	s.PolicyDocument = &v
	return s
}

func (s *GetControlPolicyResponseBodyControlPolicy) SetPolicyId(v string) *GetControlPolicyResponseBodyControlPolicy {
	s.PolicyId = &v
	return s
}

func (s *GetControlPolicyResponseBodyControlPolicy) SetPolicyName(v string) *GetControlPolicyResponseBodyControlPolicy {
	s.PolicyName = &v
	return s
}

func (s *GetControlPolicyResponseBodyControlPolicy) SetPolicyType(v string) *GetControlPolicyResponseBodyControlPolicy {
	s.PolicyType = &v
	return s
}

func (s *GetControlPolicyResponseBodyControlPolicy) SetUpdateDate(v string) *GetControlPolicyResponseBodyControlPolicy {
	s.UpdateDate = &v
	return s
}

type GetControlPolicyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetControlPolicyResponse) SetHeaders(v map[string]*string) *GetControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetControlPolicyResponse) SetStatusCode(v int32) *GetControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetControlPolicyResponse) SetBody(v *GetControlPolicyResponseBody) *GetControlPolicyResponse {
	s.Body = v
	return s
}

type GetControlPolicyEnablementStatusResponseBody struct {
	// The status of the Control Policy feature. Valid values:
	//
	// *   Enabled: The Control Policy feature is enabled.
	// *   PendingEnable: The Control Policy feature is being enabled.
	// *   Disabled: The Control Policy feature is disabled.
	// *   PendingDisable: The Control Policy feature is being disabled.
	EnablementStatus *string `json:"EnablementStatus,omitempty" xml:"EnablementStatus,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetControlPolicyEnablementStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetControlPolicyEnablementStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetControlPolicyEnablementStatusResponseBody) SetEnablementStatus(v string) *GetControlPolicyEnablementStatusResponseBody {
	s.EnablementStatus = &v
	return s
}

func (s *GetControlPolicyEnablementStatusResponseBody) SetRequestId(v string) *GetControlPolicyEnablementStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetControlPolicyEnablementStatusResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetControlPolicyEnablementStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetControlPolicyEnablementStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetControlPolicyEnablementStatusResponse) GoString() string {
	return s.String()
}

func (s *GetControlPolicyEnablementStatusResponse) SetHeaders(v map[string]*string) *GetControlPolicyEnablementStatusResponse {
	s.Headers = v
	return s
}

func (s *GetControlPolicyEnablementStatusResponse) SetStatusCode(v int32) *GetControlPolicyEnablementStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetControlPolicyEnablementStatusResponse) SetBody(v *GetControlPolicyEnablementStatusResponseBody) *GetControlPolicyEnablementStatusResponse {
	s.Body = v
	return s
}

type GetFolderRequest struct {
	// The ID of the folder.
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
}

func (s GetFolderRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFolderRequest) GoString() string {
	return s.String()
}

func (s *GetFolderRequest) SetFolderId(v string) *GetFolderRequest {
	s.FolderId = &v
	return s
}

type GetFolderResponseBody struct {
	// The information of the folder.
	Folder *GetFolderResponseBodyFolder `json:"Folder,omitempty" xml:"Folder,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFolderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFolderResponseBody) GoString() string {
	return s.String()
}

func (s *GetFolderResponseBody) SetFolder(v *GetFolderResponseBodyFolder) *GetFolderResponseBody {
	s.Folder = v
	return s
}

func (s *GetFolderResponseBody) SetRequestId(v string) *GetFolderResponseBody {
	s.RequestId = &v
	return s
}

type GetFolderResponseBodyFolder struct {
	// The time when the folder was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the folder.
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The name of the folder.
	FolderName *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
	// The ID of the parent folder.
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	// The path of the folder in the resource directory.
	ResourceDirectoryPath *string `json:"ResourceDirectoryPath,omitempty" xml:"ResourceDirectoryPath,omitempty"`
}

func (s GetFolderResponseBodyFolder) String() string {
	return tea.Prettify(s)
}

func (s GetFolderResponseBodyFolder) GoString() string {
	return s.String()
}

func (s *GetFolderResponseBodyFolder) SetCreateTime(v string) *GetFolderResponseBodyFolder {
	s.CreateTime = &v
	return s
}

func (s *GetFolderResponseBodyFolder) SetFolderId(v string) *GetFolderResponseBodyFolder {
	s.FolderId = &v
	return s
}

func (s *GetFolderResponseBodyFolder) SetFolderName(v string) *GetFolderResponseBodyFolder {
	s.FolderName = &v
	return s
}

func (s *GetFolderResponseBodyFolder) SetParentFolderId(v string) *GetFolderResponseBodyFolder {
	s.ParentFolderId = &v
	return s
}

func (s *GetFolderResponseBodyFolder) SetResourceDirectoryPath(v string) *GetFolderResponseBodyFolder {
	s.ResourceDirectoryPath = &v
	return s
}

type GetFolderResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFolderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFolderResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFolderResponse) GoString() string {
	return s.String()
}

func (s *GetFolderResponse) SetHeaders(v map[string]*string) *GetFolderResponse {
	s.Headers = v
	return s
}

func (s *GetFolderResponse) SetStatusCode(v int32) *GetFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFolderResponse) SetBody(v *GetFolderResponseBody) *GetFolderResponse {
	s.Body = v
	return s
}

type GetHandshakeRequest struct {
	// The ID of the invitation.
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
}

func (s GetHandshakeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHandshakeRequest) GoString() string {
	return s.String()
}

func (s *GetHandshakeRequest) SetHandshakeId(v string) *GetHandshakeRequest {
	s.HandshakeId = &v
	return s
}

type GetHandshakeResponseBody struct {
	// The information of the invitation.
	Handshake *GetHandshakeResponseBodyHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetHandshakeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHandshakeResponseBody) GoString() string {
	return s.String()
}

func (s *GetHandshakeResponseBody) SetHandshake(v *GetHandshakeResponseBodyHandshake) *GetHandshakeResponseBody {
	s.Handshake = v
	return s
}

func (s *GetHandshakeResponseBody) SetRequestId(v string) *GetHandshakeResponseBody {
	s.RequestId = &v
	return s
}

type GetHandshakeResponseBodyHandshake struct {
	// The time when the invitation was created. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the invitation expires. The time is displayed in UTC.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the invitation.
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
	// The real-name verification information of the invitee.
	//
	// >  This parameter is available only when an invitee calls this operation.
	InvitedAccountRealName *string `json:"InvitedAccountRealName,omitempty" xml:"InvitedAccountRealName,omitempty"`
	// The ID of the management account of the resource directory.
	MasterAccountId *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	// The name of the management account of the resource directory.
	MasterAccountName *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	// The real-name verification information of the management account of the resource directory.
	//
	// >  This parameter is available only when an invitee calls this operation.
	MasterAccountRealName *string `json:"MasterAccountRealName,omitempty" xml:"MasterAccountRealName,omitempty"`
	// The time when the invitation was modified. The time is displayed in UTC.
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The comment on the invitation.
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The ID of the resource directory.
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the invitation. Valid values:
	//
	// *   Pending: The invitation is waiting for confirmation.
	// *   Accepted: The invitation is accepted.
	// *   Cancelled: The invitation is canceled.
	// *   Declined: The invitation is rejected.
	// *   Expired: The invitation expires.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID or logon email address of the invited account.
	TargetEntity *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty"`
	// The type of the invited account. Valid values:
	//
	// *   Account: indicates the ID of the account.
	// *   Email: indicates the logon email address of the account.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s GetHandshakeResponseBodyHandshake) String() string {
	return tea.Prettify(s)
}

func (s GetHandshakeResponseBodyHandshake) GoString() string {
	return s.String()
}

func (s *GetHandshakeResponseBodyHandshake) SetCreateTime(v string) *GetHandshakeResponseBodyHandshake {
	s.CreateTime = &v
	return s
}

func (s *GetHandshakeResponseBodyHandshake) SetExpireTime(v string) *GetHandshakeResponseBodyHandshake {
	s.ExpireTime = &v
	return s
}

func (s *GetHandshakeResponseBodyHandshake) SetHandshakeId(v string) *GetHandshakeResponseBodyHandshake {
	s.HandshakeId = &v
	return s
}

func (s *GetHandshakeResponseBodyHandshake) SetInvitedAccountRealName(v string) *GetHandshakeResponseBodyHandshake {
	s.InvitedAccountRealName = &v
	return s
}

func (s *GetHandshakeResponseBodyHandshake) SetMasterAccountId(v string) *GetHandshakeResponseBodyHandshake {
	s.MasterAccountId = &v
	return s
}

func (s *GetHandshakeResponseBodyHandshake) SetMasterAccountName(v string) *GetHandshakeResponseBodyHandshake {
	s.MasterAccountName = &v
	return s
}

func (s *GetHandshakeResponseBodyHandshake) SetMasterAccountRealName(v string) *GetHandshakeResponseBodyHandshake {
	s.MasterAccountRealName = &v
	return s
}

func (s *GetHandshakeResponseBodyHandshake) SetModifyTime(v string) *GetHandshakeResponseBodyHandshake {
	s.ModifyTime = &v
	return s
}

func (s *GetHandshakeResponseBodyHandshake) SetNote(v string) *GetHandshakeResponseBodyHandshake {
	s.Note = &v
	return s
}

func (s *GetHandshakeResponseBodyHandshake) SetResourceDirectoryId(v string) *GetHandshakeResponseBodyHandshake {
	s.ResourceDirectoryId = &v
	return s
}

func (s *GetHandshakeResponseBodyHandshake) SetStatus(v string) *GetHandshakeResponseBodyHandshake {
	s.Status = &v
	return s
}

func (s *GetHandshakeResponseBodyHandshake) SetTargetEntity(v string) *GetHandshakeResponseBodyHandshake {
	s.TargetEntity = &v
	return s
}

func (s *GetHandshakeResponseBodyHandshake) SetTargetType(v string) *GetHandshakeResponseBodyHandshake {
	s.TargetType = &v
	return s
}

type GetHandshakeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHandshakeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHandshakeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHandshakeResponse) GoString() string {
	return s.String()
}

func (s *GetHandshakeResponse) SetHeaders(v map[string]*string) *GetHandshakeResponse {
	s.Headers = v
	return s
}

func (s *GetHandshakeResponse) SetStatusCode(v int32) *GetHandshakeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHandshakeResponse) SetBody(v *GetHandshakeResponseBody) *GetHandshakeResponse {
	s.Body = v
	return s
}

type GetPayerForAccountRequest struct {
	// The ID of the account.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s GetPayerForAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPayerForAccountRequest) GoString() string {
	return s.String()
}

func (s *GetPayerForAccountRequest) SetAccountId(v string) *GetPayerForAccountRequest {
	s.AccountId = &v
	return s
}

type GetPayerForAccountResponseBody struct {
	// The ID of the settlement account.
	PayerAccountId *string `json:"PayerAccountId,omitempty" xml:"PayerAccountId,omitempty"`
	// The name of the settlement account.
	PayerAccountName *string `json:"PayerAccountName,omitempty" xml:"PayerAccountName,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPayerForAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPayerForAccountResponseBody) GoString() string {
	return s.String()
}

func (s *GetPayerForAccountResponseBody) SetPayerAccountId(v string) *GetPayerForAccountResponseBody {
	s.PayerAccountId = &v
	return s
}

func (s *GetPayerForAccountResponseBody) SetPayerAccountName(v string) *GetPayerForAccountResponseBody {
	s.PayerAccountName = &v
	return s
}

func (s *GetPayerForAccountResponseBody) SetRequestId(v string) *GetPayerForAccountResponseBody {
	s.RequestId = &v
	return s
}

type GetPayerForAccountResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPayerForAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPayerForAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPayerForAccountResponse) GoString() string {
	return s.String()
}

func (s *GetPayerForAccountResponse) SetHeaders(v map[string]*string) *GetPayerForAccountResponse {
	s.Headers = v
	return s
}

func (s *GetPayerForAccountResponse) SetStatusCode(v int32) *GetPayerForAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPayerForAccountResponse) SetBody(v *GetPayerForAccountResponseBody) *GetPayerForAccountResponse {
	s.Body = v
	return s
}

type GetPolicyRequest struct {
	// The language that is used to return the description of the system policy. Valid values:
	//
	// *   en: English
	// *   zh-CN: Chinese
	// *   ja: Japanese
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The name of the policy.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy. Valid values:
	//
	// *   Custom: custom policy
	// *   System: system policy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s GetPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetPolicyRequest) SetLanguage(v string) *GetPolicyRequest {
	s.Language = &v
	return s
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
	// The information of the policy.
	Policy *GetPolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetPolicyResponseBody) SetPolicy(v *GetPolicyResponseBodyPolicy) *GetPolicyResponseBody {
	s.Policy = v
	return s
}

func (s *GetPolicyResponseBody) SetRequestId(v string) *GetPolicyResponseBody {
	s.RequestId = &v
	return s
}

type GetPolicyResponseBodyPolicy struct {
	// The number of times the policy is referenced.
	AttachmentCount *int32 `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty"`
	// The time when the policy was created.
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The default version of the policy.
	DefaultVersion *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// The description of the policy.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The document of the policy.
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The name of the policy.
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy. Valid values:
	//
	// *   Custom: custom policy
	// *   System: system policy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The time when the policy was updated.
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
	// The name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy. Valid values:
	//
	// *   Custom: custom policy
	// *   System: system policy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The ID of the policy version.
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
	// The information of the policy version.
	PolicyVersion *GetPolicyVersionResponseBodyPolicyVersion `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty" type:"Struct"`
	// The ID of the request.
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
	// The time when the policy version was created.
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// Indicates whether the policy version is the default version.
	IsDefaultVersion *bool `json:"IsDefaultVersion,omitempty" xml:"IsDefaultVersion,omitempty"`
	// The document of the policy.
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The ID of the policy version.
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

type GetResourceDirectoryResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the resource directory.
	ResourceDirectory *GetResourceDirectoryResponseBodyResourceDirectory `json:"ResourceDirectory,omitempty" xml:"ResourceDirectory,omitempty" type:"Struct"`
}

func (s GetResourceDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceDirectoryResponseBody) SetRequestId(v string) *GetResourceDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceDirectoryResponseBody) SetResourceDirectory(v *GetResourceDirectoryResponseBodyResourceDirectory) *GetResourceDirectoryResponseBody {
	s.ResourceDirectory = v
	return s
}

type GetResourceDirectoryResponseBodyResourceDirectory struct {
	// The status of the Control Policy feature. Valid values:
	//
	// *   Enabled: The feature is enabled.
	// *   PendingEnable: The feature is being enabled.
	// *   Disabled: The feature is disabled.
	// *   PendingDisable: The feature is being disabled.
	ControlPolicyStatus *string `json:"ControlPolicyStatus,omitempty" xml:"ControlPolicyStatus,omitempty"`
	// The time when the resource directory was enabled.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The real-name verification information.
	IdentityInformation *string `json:"IdentityInformation,omitempty" xml:"IdentityInformation,omitempty"`
	// The ID of the management account.
	MasterAccountId *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	// The name of the management account.
	MasterAccountName *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	// The status of the member deletion feature. Valid values:
	//
	// *   Enabled: The feature is enabled. You can call the [DeleteAccount](~~311546~~) operation to delete members of the resource account type.
	// *   Disabled: The feature is disabled. You cannot delete members of the resource account type.
	MemberDeletionStatus *string `json:"MemberDeletionStatus,omitempty" xml:"MemberDeletionStatus,omitempty"`
	// The ID of the resource directory.
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The ID of the Root folder.
	RootFolderId *string `json:"RootFolderId,omitempty" xml:"RootFolderId,omitempty"`
}

func (s GetResourceDirectoryResponseBodyResourceDirectory) String() string {
	return tea.Prettify(s)
}

func (s GetResourceDirectoryResponseBodyResourceDirectory) GoString() string {
	return s.String()
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) SetControlPolicyStatus(v string) *GetResourceDirectoryResponseBodyResourceDirectory {
	s.ControlPolicyStatus = &v
	return s
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) SetCreateTime(v string) *GetResourceDirectoryResponseBodyResourceDirectory {
	s.CreateTime = &v
	return s
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) SetIdentityInformation(v string) *GetResourceDirectoryResponseBodyResourceDirectory {
	s.IdentityInformation = &v
	return s
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) SetMasterAccountId(v string) *GetResourceDirectoryResponseBodyResourceDirectory {
	s.MasterAccountId = &v
	return s
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) SetMasterAccountName(v string) *GetResourceDirectoryResponseBodyResourceDirectory {
	s.MasterAccountName = &v
	return s
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) SetMemberDeletionStatus(v string) *GetResourceDirectoryResponseBodyResourceDirectory {
	s.MemberDeletionStatus = &v
	return s
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) SetResourceDirectoryId(v string) *GetResourceDirectoryResponseBodyResourceDirectory {
	s.ResourceDirectoryId = &v
	return s
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) SetRootFolderId(v string) *GetResourceDirectoryResponseBodyResourceDirectory {
	s.RootFolderId = &v
	return s
}

type GetResourceDirectoryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceDirectoryResponse) GoString() string {
	return s.String()
}

func (s *GetResourceDirectoryResponse) SetHeaders(v map[string]*string) *GetResourceDirectoryResponse {
	s.Headers = v
	return s
}

func (s *GetResourceDirectoryResponse) SetStatusCode(v int32) *GetResourceDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceDirectoryResponse) SetBody(v *GetResourceDirectoryResponseBody) *GetResourceDirectoryResponse {
	s.Body = v
	return s
}

type GetResourceGroupRequest struct {
	// The ID of the request.
	IncludeTags *bool `json:"IncludeTags,omitempty" xml:"IncludeTags,omitempty"`
	// Specifies whether to return the information of tags. Valid values:
	//
	// *   false (default value)
	// *   true
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s GetResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *GetResourceGroupRequest) SetIncludeTags(v bool) *GetResourceGroupRequest {
	s.IncludeTags = &v
	return s
}

func (s *GetResourceGroupRequest) SetResourceGroupId(v string) *GetResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

type GetResourceGroupResponseBody struct {
	// The information of the resource group.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The display name of the resource group.
	ResourceGroup *GetResourceGroupResponseBodyResourceGroup `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty" type:"Struct"`
}

func (s GetResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponseBody) SetRequestId(v string) *GetResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceGroupResponseBody) SetResourceGroup(v *GetResourceGroupResponseBodyResourceGroup) *GetResourceGroupResponseBody {
	s.ResourceGroup = v
	return s
}

type GetResourceGroupResponseBodyResourceGroup struct {
	// The identifier of the resource group.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the resource group.
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The status of the resource group. Valid values:
	//
	// *   Creating: The resource group is being created.
	// *   OK: The resource group is created.
	// *   PendingDelete: The resource group is waiting to be deleted.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The tags that are added to the resource group.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the resource group was created. The time is displayed in UTC.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the resource group. Valid values:
	//
	// *   Creating: The resource group is being created.
	// *   OK: The resource group is created.
	// *   PendingDelete: The resource group is waiting to be deleted.
	// *   Deleting: The resource group is being deleted.
	RegionStatuses *GetResourceGroupResponseBodyResourceGroupRegionStatuses `json:"RegionStatuses,omitempty" xml:"RegionStatuses,omitempty" type:"Struct"`
	// The status of the resource group in all regions.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag key.
	Tags *GetResourceGroupResponseBodyResourceGroupTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s GetResourceGroupResponseBodyResourceGroup) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupResponseBodyResourceGroup) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetAccountId(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.AccountId = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetCreateDate(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.CreateDate = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetDisplayName(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.DisplayName = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetId(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.Id = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetName(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.Name = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetRegionStatuses(v *GetResourceGroupResponseBodyResourceGroupRegionStatuses) *GetResourceGroupResponseBodyResourceGroup {
	s.RegionStatuses = v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetStatus(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.Status = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetTags(v *GetResourceGroupResponseBodyResourceGroupTags) *GetResourceGroupResponseBodyResourceGroup {
	s.Tags = v
	return s
}

type GetResourceGroupResponseBodyResourceGroupRegionStatuses struct {
	RegionStatus []*GetResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus `json:"RegionStatus,omitempty" xml:"RegionStatus,omitempty" type:"Repeated"`
}

func (s GetResourceGroupResponseBodyResourceGroupRegionStatuses) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupResponseBodyResourceGroupRegionStatuses) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponseBodyResourceGroupRegionStatuses) SetRegionStatus(v []*GetResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) *GetResourceGroupResponseBodyResourceGroupRegionStatuses {
	s.RegionStatus = v
	return s
}

type GetResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus struct {
	// The ID of the Alibaba Cloud account to which the resource group belongs.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The region ID.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) SetRegionId(v string) *GetResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus {
	s.RegionId = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) SetStatus(v string) *GetResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus {
	s.Status = &v
	return s
}

type GetResourceGroupResponseBodyResourceGroupTags struct {
	Tag []*GetResourceGroupResponseBodyResourceGroupTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s GetResourceGroupResponseBodyResourceGroupTags) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupResponseBodyResourceGroupTags) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponseBodyResourceGroupTags) SetTag(v []*GetResourceGroupResponseBodyResourceGroupTagsTag) *GetResourceGroupResponseBodyResourceGroupTags {
	s.Tag = v
	return s
}

type GetResourceGroupResponseBodyResourceGroupTagsTag struct {
	// The tag value.
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetResourceGroupResponseBodyResourceGroupTagsTag) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupResponseBodyResourceGroupTagsTag) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponseBodyResourceGroupTagsTag) SetTagKey(v string) *GetResourceGroupResponseBodyResourceGroupTagsTag {
	s.TagKey = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroupTagsTag) SetTagValue(v string) *GetResourceGroupResponseBodyResourceGroupTagsTag {
	s.TagValue = &v
	return s
}

type GetResourceGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponse) SetHeaders(v map[string]*string) *GetResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *GetResourceGroupResponse) SetStatusCode(v int32) *GetResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceGroupResponse) SetBody(v *GetResourceGroupResponseBody) *GetResourceGroupResponse {
	s.Body = v
	return s
}

type GetRoleRequest struct {
	// The language that is used to return the description of the RAM role. Valid values:
	//
	// *   en: English
	// *   zh-CN: Chinese
	// *   ja: Japanese
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The name of the RAM role.
	//
	// The name must be 1 to 64 characters in length and can contain letters, digits, periods (.), and hyphens (-).
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s GetRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRoleRequest) GoString() string {
	return s.String()
}

func (s *GetRoleRequest) SetLanguage(v string) *GetRoleRequest {
	s.Language = &v
	return s
}

func (s *GetRoleRequest) SetRoleName(v string) *GetRoleRequest {
	s.RoleName = &v
	return s
}

type GetRoleResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the RAM role.
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
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The document of the policy that specifies the trusted entity to assume the RAM role.
	AssumeRolePolicyDocument *string `json:"AssumeRolePolicyDocument,omitempty" xml:"AssumeRolePolicyDocument,omitempty"`
	// The time when the RAM role was created.
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the RAM role.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the RAM role is a service linked role.
	IsServiceLinkedRole *bool `json:"IsServiceLinkedRole,omitempty" xml:"IsServiceLinkedRole,omitempty"`
	// The information of the most recent deletion task.
	LatestDeletionTask *GetRoleResponseBodyRoleLatestDeletionTask `json:"LatestDeletionTask,omitempty" xml:"LatestDeletionTask,omitempty" type:"Struct"`
	// The maximum session duration of the RAM role.
	MaxSessionDuration *int64 `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty"`
	// The ID of the RAM role.
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// The name of the RAM role.
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The name of the RAM role after authorization.
	RolePrincipalName *string `json:"RolePrincipalName,omitempty" xml:"RolePrincipalName,omitempty"`
	// The time when the RAM role was updated.
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

func (s *GetRoleResponseBodyRole) SetIsServiceLinkedRole(v bool) *GetRoleResponseBodyRole {
	s.IsServiceLinkedRole = &v
	return s
}

func (s *GetRoleResponseBodyRole) SetLatestDeletionTask(v *GetRoleResponseBodyRoleLatestDeletionTask) *GetRoleResponseBodyRole {
	s.LatestDeletionTask = v
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

func (s *GetRoleResponseBodyRole) SetRolePrincipalName(v string) *GetRoleResponseBodyRole {
	s.RolePrincipalName = &v
	return s
}

func (s *GetRoleResponseBodyRole) SetUpdateDate(v string) *GetRoleResponseBodyRole {
	s.UpdateDate = &v
	return s
}

type GetRoleResponseBodyRoleLatestDeletionTask struct {
	// The time when the deletion task was created.
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The ID of the deletion task.
	DeletionTaskId *string `json:"DeletionTaskId,omitempty" xml:"DeletionTaskId,omitempty"`
}

func (s GetRoleResponseBodyRoleLatestDeletionTask) String() string {
	return tea.Prettify(s)
}

func (s GetRoleResponseBodyRoleLatestDeletionTask) GoString() string {
	return s.String()
}

func (s *GetRoleResponseBodyRoleLatestDeletionTask) SetCreateDate(v string) *GetRoleResponseBodyRoleLatestDeletionTask {
	s.CreateDate = &v
	return s
}

func (s *GetRoleResponseBodyRoleLatestDeletionTask) SetDeletionTaskId(v string) *GetRoleResponseBodyRoleLatestDeletionTask {
	s.DeletionTaskId = &v
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

type GetServiceLinkedRoleDeletionStatusRequest struct {
	// The ID of the deletion task.
	DeletionTaskId *string `json:"DeletionTaskId,omitempty" xml:"DeletionTaskId,omitempty"`
}

func (s GetServiceLinkedRoleDeletionStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceLinkedRoleDeletionStatusRequest) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleDeletionStatusRequest) SetDeletionTaskId(v string) *GetServiceLinkedRoleDeletionStatusRequest {
	s.DeletionTaskId = &v
	return s
}

type GetServiceLinkedRoleDeletionStatusResponseBody struct {
	// The reason why the deletion task failed.
	Reason *GetServiceLinkedRoleDeletionStatusResponseBodyReason `json:"Reason,omitempty" xml:"Reason,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the task.
	//
	// - SUCCEEDED
	// - IN_PROGRESS
	// - FAILED
	// - NOT_STARTED
	// - INTERNAL_ERROR
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetServiceLinkedRoleDeletionStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceLinkedRoleDeletionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBody) SetReason(v *GetServiceLinkedRoleDeletionStatusResponseBodyReason) *GetServiceLinkedRoleDeletionStatusResponseBody {
	s.Reason = v
	return s
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBody) SetRequestId(v string) *GetServiceLinkedRoleDeletionStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBody) SetStatus(v string) *GetServiceLinkedRoleDeletionStatusResponseBody {
	s.Status = &v
	return s
}

type GetServiceLinkedRoleDeletionStatusResponseBodyReason struct {
	// Failure information.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Use resource information of the service linked role.
	RoleUsages *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsages `json:"RoleUsages,omitempty" xml:"RoleUsages,omitempty" type:"Struct"`
}

func (s GetServiceLinkedRoleDeletionStatusResponseBodyReason) String() string {
	return tea.Prettify(s)
}

func (s GetServiceLinkedRoleDeletionStatusResponseBodyReason) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBodyReason) SetMessage(v string) *GetServiceLinkedRoleDeletionStatusResponseBodyReason {
	s.Message = &v
	return s
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBodyReason) SetRoleUsages(v *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsages) *GetServiceLinkedRoleDeletionStatusResponseBodyReason {
	s.RoleUsages = v
	return s
}

type GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsages struct {
	RoleUsage []*GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsage `json:"RoleUsage,omitempty" xml:"RoleUsage,omitempty" type:"Repeated"`
}

func (s GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsages) String() string {
	return tea.Prettify(s)
}

func (s GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsages) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsages) SetRoleUsage(v []*GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsage) *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsages {
	s.RoleUsage = v
	return s
}

type GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsage struct {
	// The IDs of the regions in which the resources are to be queried.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The returned resources.
	Resources *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsageResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
}

func (s GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsage) String() string {
	return tea.Prettify(s)
}

func (s GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsage) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsage) SetRegion(v string) *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsage {
	s.Region = &v
	return s
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsage) SetResources(v *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsageResources) *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsage {
	s.Resources = v
	return s
}

type GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsageResources struct {
	Resource []*string `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsageResources) String() string {
	return tea.Prettify(s)
}

func (s GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsageResources) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsageResources) SetResource(v []*string) *GetServiceLinkedRoleDeletionStatusResponseBodyReasonRoleUsagesRoleUsageResources {
	s.Resource = v
	return s
}

type GetServiceLinkedRoleDeletionStatusResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceLinkedRoleDeletionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceLinkedRoleDeletionStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceLinkedRoleDeletionStatusResponse) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleDeletionStatusResponse) SetHeaders(v map[string]*string) *GetServiceLinkedRoleDeletionStatusResponse {
	s.Headers = v
	return s
}

func (s *GetServiceLinkedRoleDeletionStatusResponse) SetStatusCode(v int32) *GetServiceLinkedRoleDeletionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceLinkedRoleDeletionStatusResponse) SetBody(v *GetServiceLinkedRoleDeletionStatusResponseBody) *GetServiceLinkedRoleDeletionStatusResponse {
	s.Body = v
	return s
}

type InitResourceDirectoryResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the resource directory.
	ResourceDirectory *InitResourceDirectoryResponseBodyResourceDirectory `json:"ResourceDirectory,omitempty" xml:"ResourceDirectory,omitempty" type:"Struct"`
}

func (s InitResourceDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitResourceDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *InitResourceDirectoryResponseBody) SetRequestId(v string) *InitResourceDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitResourceDirectoryResponseBody) SetResourceDirectory(v *InitResourceDirectoryResponseBodyResourceDirectory) *InitResourceDirectoryResponseBody {
	s.ResourceDirectory = v
	return s
}

type InitResourceDirectoryResponseBodyResourceDirectory struct {
	// The time when the resource directory was enabled.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the enterprise management account.
	MasterAccountId *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	// The name of the enterprise management account.
	MasterAccountName *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	// The ID of the resource directory.
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The ID of the root folder.
	RootFolderId *string `json:"RootFolderId,omitempty" xml:"RootFolderId,omitempty"`
}

func (s InitResourceDirectoryResponseBodyResourceDirectory) String() string {
	return tea.Prettify(s)
}

func (s InitResourceDirectoryResponseBodyResourceDirectory) GoString() string {
	return s.String()
}

func (s *InitResourceDirectoryResponseBodyResourceDirectory) SetCreateTime(v string) *InitResourceDirectoryResponseBodyResourceDirectory {
	s.CreateTime = &v
	return s
}

func (s *InitResourceDirectoryResponseBodyResourceDirectory) SetMasterAccountId(v string) *InitResourceDirectoryResponseBodyResourceDirectory {
	s.MasterAccountId = &v
	return s
}

func (s *InitResourceDirectoryResponseBodyResourceDirectory) SetMasterAccountName(v string) *InitResourceDirectoryResponseBodyResourceDirectory {
	s.MasterAccountName = &v
	return s
}

func (s *InitResourceDirectoryResponseBodyResourceDirectory) SetResourceDirectoryId(v string) *InitResourceDirectoryResponseBodyResourceDirectory {
	s.ResourceDirectoryId = &v
	return s
}

func (s *InitResourceDirectoryResponseBodyResourceDirectory) SetRootFolderId(v string) *InitResourceDirectoryResponseBodyResourceDirectory {
	s.RootFolderId = &v
	return s
}

type InitResourceDirectoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitResourceDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitResourceDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s InitResourceDirectoryResponse) GoString() string {
	return s.String()
}

func (s *InitResourceDirectoryResponse) SetHeaders(v map[string]*string) *InitResourceDirectoryResponse {
	s.Headers = v
	return s
}

func (s *InitResourceDirectoryResponse) SetStatusCode(v int32) *InitResourceDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *InitResourceDirectoryResponse) SetBody(v *InitResourceDirectoryResponseBody) *InitResourceDirectoryResponse {
	s.Body = v
	return s
}

type InviteAccountToResourceDirectoryRequest struct {
	// The comment on the invitation.
	//
	// The comment can be up to 1,024 characters in length.
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The tag key and value.
	Tag []*InviteAccountToResourceDirectoryRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID or logon email address of the account that you want to invite.
	TargetEntity *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty"`
	// The type of the account. Valid values:
	//
	// *   Account: indicates the ID of the account.
	// *   Email: indicates the logon email address of the account.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s InviteAccountToResourceDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s InviteAccountToResourceDirectoryRequest) GoString() string {
	return s.String()
}

func (s *InviteAccountToResourceDirectoryRequest) SetNote(v string) *InviteAccountToResourceDirectoryRequest {
	s.Note = &v
	return s
}

func (s *InviteAccountToResourceDirectoryRequest) SetTag(v []*InviteAccountToResourceDirectoryRequestTag) *InviteAccountToResourceDirectoryRequest {
	s.Tag = v
	return s
}

func (s *InviteAccountToResourceDirectoryRequest) SetTargetEntity(v string) *InviteAccountToResourceDirectoryRequest {
	s.TargetEntity = &v
	return s
}

func (s *InviteAccountToResourceDirectoryRequest) SetTargetType(v string) *InviteAccountToResourceDirectoryRequest {
	s.TargetType = &v
	return s
}

type InviteAccountToResourceDirectoryRequestTag struct {
	// A tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// A tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s InviteAccountToResourceDirectoryRequestTag) String() string {
	return tea.Prettify(s)
}

func (s InviteAccountToResourceDirectoryRequestTag) GoString() string {
	return s.String()
}

func (s *InviteAccountToResourceDirectoryRequestTag) SetKey(v string) *InviteAccountToResourceDirectoryRequestTag {
	s.Key = &v
	return s
}

func (s *InviteAccountToResourceDirectoryRequestTag) SetValue(v string) *InviteAccountToResourceDirectoryRequestTag {
	s.Value = &v
	return s
}

type InviteAccountToResourceDirectoryResponseBody struct {
	// The information of the invitation.
	Handshake *InviteAccountToResourceDirectoryResponseBodyHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InviteAccountToResourceDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InviteAccountToResourceDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *InviteAccountToResourceDirectoryResponseBody) SetHandshake(v *InviteAccountToResourceDirectoryResponseBodyHandshake) *InviteAccountToResourceDirectoryResponseBody {
	s.Handshake = v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseBody) SetRequestId(v string) *InviteAccountToResourceDirectoryResponseBody {
	s.RequestId = &v
	return s
}

type InviteAccountToResourceDirectoryResponseBodyHandshake struct {
	// The time when the invitation was created. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the invitation expires. The time is displayed in UTC.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the invitation.
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
	// The ID of the management account of the resource directory.
	MasterAccountId *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	// The name of the management account of the resource directory.
	MasterAccountName *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	// The time when the invitation was modified. The time is displayed in UTC.
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The comment on the invitation.
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The ID of the resource directory.
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the invitation.
	//
	// *   Pending: The invitation is waiting for confirmation.
	// *   Accepted: The invitation is accepted.
	// *   Cancelled: The invitation is canceled.
	// *   Declined: The invitation is rejected.
	// *   Expired: The invitation expires.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID or logon email address of the invited account.
	TargetEntity *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty"`
	// The type of the invited account. Valid values:
	//
	// *   Account: indicates the ID of the account.
	// *   Email: indicates the logon email address of the account.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s InviteAccountToResourceDirectoryResponseBodyHandshake) String() string {
	return tea.Prettify(s)
}

func (s InviteAccountToResourceDirectoryResponseBodyHandshake) GoString() string {
	return s.String()
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) SetCreateTime(v string) *InviteAccountToResourceDirectoryResponseBodyHandshake {
	s.CreateTime = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) SetExpireTime(v string) *InviteAccountToResourceDirectoryResponseBodyHandshake {
	s.ExpireTime = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) SetHandshakeId(v string) *InviteAccountToResourceDirectoryResponseBodyHandshake {
	s.HandshakeId = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) SetMasterAccountId(v string) *InviteAccountToResourceDirectoryResponseBodyHandshake {
	s.MasterAccountId = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) SetMasterAccountName(v string) *InviteAccountToResourceDirectoryResponseBodyHandshake {
	s.MasterAccountName = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) SetModifyTime(v string) *InviteAccountToResourceDirectoryResponseBodyHandshake {
	s.ModifyTime = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) SetNote(v string) *InviteAccountToResourceDirectoryResponseBodyHandshake {
	s.Note = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) SetResourceDirectoryId(v string) *InviteAccountToResourceDirectoryResponseBodyHandshake {
	s.ResourceDirectoryId = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) SetStatus(v string) *InviteAccountToResourceDirectoryResponseBodyHandshake {
	s.Status = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) SetTargetEntity(v string) *InviteAccountToResourceDirectoryResponseBodyHandshake {
	s.TargetEntity = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) SetTargetType(v string) *InviteAccountToResourceDirectoryResponseBodyHandshake {
	s.TargetType = &v
	return s
}

type InviteAccountToResourceDirectoryResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InviteAccountToResourceDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InviteAccountToResourceDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s InviteAccountToResourceDirectoryResponse) GoString() string {
	return s.String()
}

func (s *InviteAccountToResourceDirectoryResponse) SetHeaders(v map[string]*string) *InviteAccountToResourceDirectoryResponse {
	s.Headers = v
	return s
}

func (s *InviteAccountToResourceDirectoryResponse) SetStatusCode(v int32) *InviteAccountToResourceDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponse) SetBody(v *InviteAccountToResourceDirectoryResponseBody) *InviteAccountToResourceDirectoryResponse {
	s.Body = v
	return s
}

type ListAccountsRequest struct {
	// Specifies whether to return the information of tags. Valid values:
	//
	// *   false (default value)
	// *   true
	IncludeTags *bool `json:"IncludeTags,omitempty" xml:"IncludeTags,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The tag key and value.
	Tag []*ListAccountsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListAccountsRequest) SetIncludeTags(v bool) *ListAccountsRequest {
	s.IncludeTags = &v
	return s
}

func (s *ListAccountsRequest) SetPageNumber(v int32) *ListAccountsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAccountsRequest) SetPageSize(v int32) *ListAccountsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAccountsRequest) SetTag(v []*ListAccountsRequestTag) *ListAccountsRequest {
	s.Tag = v
	return s
}

type ListAccountsRequestTag struct {
	// A tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// A tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAccountsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsRequestTag) GoString() string {
	return s.String()
}

func (s *ListAccountsRequestTag) SetKey(v string) *ListAccountsRequestTag {
	s.Key = &v
	return s
}

func (s *ListAccountsRequestTag) SetValue(v string) *ListAccountsRequestTag {
	s.Value = &v
	return s
}

type ListAccountsResponseBody struct {
	// The members returned.
	Accounts *ListAccountsResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccountsResponseBody) SetAccounts(v *ListAccountsResponseBodyAccounts) *ListAccountsResponseBody {
	s.Accounts = v
	return s
}

func (s *ListAccountsResponseBody) SetPageNumber(v int32) *ListAccountsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAccountsResponseBody) SetPageSize(v int32) *ListAccountsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAccountsResponseBody) SetRequestId(v string) *ListAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccountsResponseBody) SetTotalCount(v int32) *ListAccountsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAccountsResponseBodyAccounts struct {
	Account []*ListAccountsResponseBodyAccountsAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Repeated"`
}

func (s ListAccountsResponseBodyAccounts) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsResponseBodyAccounts) GoString() string {
	return s.String()
}

func (s *ListAccountsResponseBodyAccounts) SetAccount(v []*ListAccountsResponseBodyAccountsAccount) *ListAccountsResponseBodyAccounts {
	s.Account = v
	return s
}

type ListAccountsResponseBodyAccountsAccount struct {
	// The Alibaba Cloud account ID of the member.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The display name of the member.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the folder.
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The way in which the member joins the resource directory. Valid values:
	//
	// *   invited: The member is invited to join the resource directory.
	// *   created: The member is directly created in the resource directory.
	JoinMethod *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	// The time when the member joined the resource directory. The time is displayed in UTC.
	JoinTime *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// The time when the member was modified. The time is displayed in UTC.
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The ID of the resource directory.
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The path of the member in the resource directory.
	ResourceDirectoryPath *string `json:"ResourceDirectoryPath,omitempty" xml:"ResourceDirectoryPath,omitempty"`
	// The status of the member. Valid values:
	//
	// *   CreateSuccess: The member is created.
	// *   PromoteVerifying: The upgrade of the member is being confirmed.
	// *   PromoteFailed: The upgrade of the member fails.
	// *   PromoteExpired: The upgrade of the member expires.
	// *   PromoteCancelled: The upgrade of the member is canceled.
	// *   PromoteSuccess: The member is upgraded.
	// *   InviteSuccess: The member accepts the invitation.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags that are added to the member.
	Tags *ListAccountsResponseBodyAccountsAccountTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The type of the member. Valid values:
	//
	// *   CloudAccount: cloud account
	// *   ResourceAccount: resource account
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAccountsResponseBodyAccountsAccount) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsResponseBodyAccountsAccount) GoString() string {
	return s.String()
}

func (s *ListAccountsResponseBodyAccountsAccount) SetAccountId(v string) *ListAccountsResponseBodyAccountsAccount {
	s.AccountId = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetDisplayName(v string) *ListAccountsResponseBodyAccountsAccount {
	s.DisplayName = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetFolderId(v string) *ListAccountsResponseBodyAccountsAccount {
	s.FolderId = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetJoinMethod(v string) *ListAccountsResponseBodyAccountsAccount {
	s.JoinMethod = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetJoinTime(v string) *ListAccountsResponseBodyAccountsAccount {
	s.JoinTime = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetModifyTime(v string) *ListAccountsResponseBodyAccountsAccount {
	s.ModifyTime = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetResourceDirectoryId(v string) *ListAccountsResponseBodyAccountsAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetResourceDirectoryPath(v string) *ListAccountsResponseBodyAccountsAccount {
	s.ResourceDirectoryPath = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetStatus(v string) *ListAccountsResponseBodyAccountsAccount {
	s.Status = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetTags(v *ListAccountsResponseBodyAccountsAccountTags) *ListAccountsResponseBodyAccountsAccount {
	s.Tags = v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetType(v string) *ListAccountsResponseBodyAccountsAccount {
	s.Type = &v
	return s
}

type ListAccountsResponseBodyAccountsAccountTags struct {
	Tag []*ListAccountsResponseBodyAccountsAccountTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListAccountsResponseBodyAccountsAccountTags) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsResponseBodyAccountsAccountTags) GoString() string {
	return s.String()
}

func (s *ListAccountsResponseBodyAccountsAccountTags) SetTag(v []*ListAccountsResponseBodyAccountsAccountTagsTag) *ListAccountsResponseBodyAccountsAccountTags {
	s.Tag = v
	return s
}

type ListAccountsResponseBodyAccountsAccountTagsTag struct {
	// A tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// A tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAccountsResponseBodyAccountsAccountTagsTag) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsResponseBodyAccountsAccountTagsTag) GoString() string {
	return s.String()
}

func (s *ListAccountsResponseBodyAccountsAccountTagsTag) SetKey(v string) *ListAccountsResponseBodyAccountsAccountTagsTag {
	s.Key = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccountTagsTag) SetValue(v string) *ListAccountsResponseBodyAccountsAccountTagsTag {
	s.Value = &v
	return s
}

type ListAccountsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsResponse) GoString() string {
	return s.String()
}

func (s *ListAccountsResponse) SetHeaders(v map[string]*string) *ListAccountsResponse {
	s.Headers = v
	return s
}

func (s *ListAccountsResponse) SetStatusCode(v int32) *ListAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccountsResponse) SetBody(v *ListAccountsResponseBody) *ListAccountsResponse {
	s.Body = v
	return s
}

type ListAccountsForParentRequest struct {
	// Specifies whether to return the information of tags. Valid values:
	//
	// false (default value)
	//
	// true
	IncludeTags *bool `json:"IncludeTags,omitempty" xml:"IncludeTags,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the folder.
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	// The keyword used for the query, such as the display name of a member.
	//
	// Fuzzy match is supported.
	QueryKeyword *string `json:"QueryKeyword,omitempty" xml:"QueryKeyword,omitempty"`
	// The tag key and value.
	Tag []*ListAccountsForParentRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListAccountsForParentRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsForParentRequest) GoString() string {
	return s.String()
}

func (s *ListAccountsForParentRequest) SetIncludeTags(v bool) *ListAccountsForParentRequest {
	s.IncludeTags = &v
	return s
}

func (s *ListAccountsForParentRequest) SetPageNumber(v int32) *ListAccountsForParentRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAccountsForParentRequest) SetPageSize(v int32) *ListAccountsForParentRequest {
	s.PageSize = &v
	return s
}

func (s *ListAccountsForParentRequest) SetParentFolderId(v string) *ListAccountsForParentRequest {
	s.ParentFolderId = &v
	return s
}

func (s *ListAccountsForParentRequest) SetQueryKeyword(v string) *ListAccountsForParentRequest {
	s.QueryKeyword = &v
	return s
}

func (s *ListAccountsForParentRequest) SetTag(v []*ListAccountsForParentRequestTag) *ListAccountsForParentRequest {
	s.Tag = v
	return s
}

type ListAccountsForParentRequestTag struct {
	// A tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// A tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAccountsForParentRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsForParentRequestTag) GoString() string {
	return s.String()
}

func (s *ListAccountsForParentRequestTag) SetKey(v string) *ListAccountsForParentRequestTag {
	s.Key = &v
	return s
}

func (s *ListAccountsForParentRequestTag) SetValue(v string) *ListAccountsForParentRequestTag {
	s.Value = &v
	return s
}

type ListAccountsForParentResponseBody struct {
	// The information of the members.
	Accounts *ListAccountsForParentResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAccountsForParentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsForParentResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccountsForParentResponseBody) SetAccounts(v *ListAccountsForParentResponseBodyAccounts) *ListAccountsForParentResponseBody {
	s.Accounts = v
	return s
}

func (s *ListAccountsForParentResponseBody) SetPageNumber(v int32) *ListAccountsForParentResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAccountsForParentResponseBody) SetPageSize(v int32) *ListAccountsForParentResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAccountsForParentResponseBody) SetRequestId(v string) *ListAccountsForParentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccountsForParentResponseBody) SetTotalCount(v int32) *ListAccountsForParentResponseBody {
	s.TotalCount = &v
	return s
}

type ListAccountsForParentResponseBodyAccounts struct {
	Account []*ListAccountsForParentResponseBodyAccountsAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Repeated"`
}

func (s ListAccountsForParentResponseBodyAccounts) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsForParentResponseBodyAccounts) GoString() string {
	return s.String()
}

func (s *ListAccountsForParentResponseBodyAccounts) SetAccount(v []*ListAccountsForParentResponseBodyAccountsAccount) *ListAccountsForParentResponseBodyAccounts {
	s.Account = v
	return s
}

type ListAccountsForParentResponseBodyAccountsAccount struct {
	// The Alibaba Cloud account ID of the member.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The display name of the member.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the folder.
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The way in which the member joins the resource directory.
	//
	// *   invited: The member is invited to join the resource directory.
	// *   created: The member is directly created in the resource directory.
	JoinMethod *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	// The time when the member joined the resource directory. The time is displayed in UTC.
	JoinTime *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// The time when the member was modified. The time is displayed in UTC.
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The ID of the resource directory.
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the member. Valid values:
	//
	// *   CreateSuccess: The member is created.
	// *   PromoteVerifying: The upgrade of the member is being confirmed.
	// *   PromoteFailed: The upgrade of the member fails.
	// *   PromoteExpired: The upgrade of the member expires.
	// *   PromoteCancelled: The upgrade of the member is canceled.
	// *   PromoteSuccess: The member is upgraded.
	// *   InviteSuccess: The member accepts the invitation.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags that are added to the member.
	Tags *ListAccountsForParentResponseBodyAccountsAccountTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The type of the member.
	//
	// *   CloudAccount: cloud account
	// *   ResourceAccount: resource account
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAccountsForParentResponseBodyAccountsAccount) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsForParentResponseBodyAccountsAccount) GoString() string {
	return s.String()
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetAccountId(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.AccountId = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetDisplayName(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.DisplayName = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetFolderId(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.FolderId = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetJoinMethod(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.JoinMethod = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetJoinTime(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.JoinTime = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetModifyTime(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.ModifyTime = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetResourceDirectoryId(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetStatus(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.Status = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetTags(v *ListAccountsForParentResponseBodyAccountsAccountTags) *ListAccountsForParentResponseBodyAccountsAccount {
	s.Tags = v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetType(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.Type = &v
	return s
}

type ListAccountsForParentResponseBodyAccountsAccountTags struct {
	Tag []*ListAccountsForParentResponseBodyAccountsAccountTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListAccountsForParentResponseBodyAccountsAccountTags) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsForParentResponseBodyAccountsAccountTags) GoString() string {
	return s.String()
}

func (s *ListAccountsForParentResponseBodyAccountsAccountTags) SetTag(v []*ListAccountsForParentResponseBodyAccountsAccountTagsTag) *ListAccountsForParentResponseBodyAccountsAccountTags {
	s.Tag = v
	return s
}

type ListAccountsForParentResponseBodyAccountsAccountTagsTag struct {
	// A tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// A tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAccountsForParentResponseBodyAccountsAccountTagsTag) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsForParentResponseBodyAccountsAccountTagsTag) GoString() string {
	return s.String()
}

func (s *ListAccountsForParentResponseBodyAccountsAccountTagsTag) SetKey(v string) *ListAccountsForParentResponseBodyAccountsAccountTagsTag {
	s.Key = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccountTagsTag) SetValue(v string) *ListAccountsForParentResponseBodyAccountsAccountTagsTag {
	s.Value = &v
	return s
}

type ListAccountsForParentResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccountsForParentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccountsForParentResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsForParentResponse) GoString() string {
	return s.String()
}

func (s *ListAccountsForParentResponse) SetHeaders(v map[string]*string) *ListAccountsForParentResponse {
	s.Headers = v
	return s
}

func (s *ListAccountsForParentResponse) SetStatusCode(v int32) *ListAccountsForParentResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccountsForParentResponse) SetBody(v *ListAccountsForParentResponseBody) *ListAccountsForParentResponse {
	s.Body = v
	return s
}

type ListAncestorsRequest struct {
	// The ID of the child folder.
	ChildId *string `json:"ChildId,omitempty" xml:"ChildId,omitempty"`
}

func (s ListAncestorsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAncestorsRequest) GoString() string {
	return s.String()
}

func (s *ListAncestorsRequest) SetChildId(v string) *ListAncestorsRequest {
	s.ChildId = &v
	return s
}

type ListAncestorsResponseBody struct {
	// The information of the folders.
	Folders *ListAncestorsResponseBodyFolders `json:"Folders,omitempty" xml:"Folders,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAncestorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAncestorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAncestorsResponseBody) SetFolders(v *ListAncestorsResponseBodyFolders) *ListAncestorsResponseBody {
	s.Folders = v
	return s
}

func (s *ListAncestorsResponseBody) SetRequestId(v string) *ListAncestorsResponseBody {
	s.RequestId = &v
	return s
}

type ListAncestorsResponseBodyFolders struct {
	Folder []*ListAncestorsResponseBodyFoldersFolder `json:"Folder,omitempty" xml:"Folder,omitempty" type:"Repeated"`
}

func (s ListAncestorsResponseBodyFolders) String() string {
	return tea.Prettify(s)
}

func (s ListAncestorsResponseBodyFolders) GoString() string {
	return s.String()
}

func (s *ListAncestorsResponseBodyFolders) SetFolder(v []*ListAncestorsResponseBodyFoldersFolder) *ListAncestorsResponseBodyFolders {
	s.Folder = v
	return s
}

type ListAncestorsResponseBodyFoldersFolder struct {
	// The time when the folder was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the folder.
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The name of the folder.
	FolderName *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
}

func (s ListAncestorsResponseBodyFoldersFolder) String() string {
	return tea.Prettify(s)
}

func (s ListAncestorsResponseBodyFoldersFolder) GoString() string {
	return s.String()
}

func (s *ListAncestorsResponseBodyFoldersFolder) SetCreateTime(v string) *ListAncestorsResponseBodyFoldersFolder {
	s.CreateTime = &v
	return s
}

func (s *ListAncestorsResponseBodyFoldersFolder) SetFolderId(v string) *ListAncestorsResponseBodyFoldersFolder {
	s.FolderId = &v
	return s
}

func (s *ListAncestorsResponseBodyFoldersFolder) SetFolderName(v string) *ListAncestorsResponseBodyFoldersFolder {
	s.FolderName = &v
	return s
}

type ListAncestorsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAncestorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAncestorsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAncestorsResponse) GoString() string {
	return s.String()
}

func (s *ListAncestorsResponse) SetHeaders(v map[string]*string) *ListAncestorsResponse {
	s.Headers = v
	return s
}

func (s *ListAncestorsResponse) SetStatusCode(v int32) *ListAncestorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAncestorsResponse) SetBody(v *ListAncestorsResponseBody) *ListAncestorsResponse {
	s.Body = v
	return s
}

type ListAssociatedTransferSettingResponseBody struct {
	// The settings of the Transfer Associated Resources feature.
	AssociatedTransferSetting *ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting `json:"AssociatedTransferSetting,omitempty" xml:"AssociatedTransferSetting,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAssociatedTransferSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAssociatedTransferSettingResponseBody) GoString() string {
	return s.String()
}

func (s *ListAssociatedTransferSettingResponseBody) SetAssociatedTransferSetting(v *ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting) *ListAssociatedTransferSettingResponseBody {
	s.AssociatedTransferSetting = v
	return s
}

func (s *ListAssociatedTransferSettingResponseBody) SetRequestId(v string) *ListAssociatedTransferSettingResponseBody {
	s.RequestId = &v
	return s
}

type ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting struct {
	// The settings of the Transfer Associated Resources feature.
	AccountId                       *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	EnableExistingResourcesTransfer *string `json:"EnableExistingResourcesTransfer,omitempty" xml:"EnableExistingResourcesTransfer,omitempty"`
	// The settings of the transfer rules.
	RuleSettings []*ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings `json:"RuleSettings,omitempty" xml:"RuleSettings,omitempty" type:"Repeated"`
	// The status of the Transfer Associated Resources feature. Valid values:
	//
	// - Enable: enabled
	// - Disable: disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting) String() string {
	return tea.Prettify(s)
}

func (s ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting) GoString() string {
	return s.String()
}

func (s *ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting) SetAccountId(v string) *ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting {
	s.AccountId = &v
	return s
}

func (s *ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting) SetEnableExistingResourcesTransfer(v string) *ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting {
	s.EnableExistingResourcesTransfer = &v
	return s
}

func (s *ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting) SetRuleSettings(v []*ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings) *ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting {
	s.RuleSettings = v
	return s
}

func (s *ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting) SetStatus(v string) *ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting {
	s.Status = &v
	return s
}

type ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings struct {
	// The type of the associated resource.
	AssociatedResourceType *string `json:"AssociatedResourceType,omitempty" xml:"AssociatedResourceType,omitempty"`
	// The service code of the associated resource.
	AssociatedService *string `json:"AssociatedService,omitempty" xml:"AssociatedService,omitempty"`
	// The type of the primary resource.
	MasterResourceType *string `json:"MasterResourceType,omitempty" xml:"MasterResourceType,omitempty"`
	// The service code of the primary resource.
	MasterService *string `json:"MasterService,omitempty" xml:"MasterService,omitempty"`
	// The status of the Transfer Associated Resources feature. Valid values:
	//
	// - Enable: enabled
	// - Disable: disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings) String() string {
	return tea.Prettify(s)
}

func (s ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings) GoString() string {
	return s.String()
}

func (s *ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings) SetAssociatedResourceType(v string) *ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings {
	s.AssociatedResourceType = &v
	return s
}

func (s *ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings) SetAssociatedService(v string) *ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings {
	s.AssociatedService = &v
	return s
}

func (s *ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings) SetMasterResourceType(v string) *ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings {
	s.MasterResourceType = &v
	return s
}

func (s *ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings) SetMasterService(v string) *ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings {
	s.MasterService = &v
	return s
}

func (s *ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings) SetStatus(v string) *ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings {
	s.Status = &v
	return s
}

type ListAssociatedTransferSettingResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAssociatedTransferSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAssociatedTransferSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAssociatedTransferSettingResponse) GoString() string {
	return s.String()
}

func (s *ListAssociatedTransferSettingResponse) SetHeaders(v map[string]*string) *ListAssociatedTransferSettingResponse {
	s.Headers = v
	return s
}

func (s *ListAssociatedTransferSettingResponse) SetStatusCode(v int32) *ListAssociatedTransferSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAssociatedTransferSettingResponse) SetBody(v *ListAssociatedTransferSettingResponseBody) *ListAssociatedTransferSettingResponse {
	s.Body = v
	return s
}

type ListControlPoliciesRequest struct {
	// The language in which you want to return the descriptions of the access control policies. Valid values:
	//
	// - zh-CN (default value): Chinese
	// - en: English
	// - ja: Japanese
	//
	// >  This parameter is valid only for system access control policies.
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The number of the page to return.
	//
	// Page start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the access control policy. Valid values:
	//
	// - System: system access control policy
	// - Custom: custom access control policy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListControlPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListControlPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListControlPoliciesRequest) SetLanguage(v string) *ListControlPoliciesRequest {
	s.Language = &v
	return s
}

func (s *ListControlPoliciesRequest) SetPageNumber(v int32) *ListControlPoliciesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListControlPoliciesRequest) SetPageSize(v int32) *ListControlPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *ListControlPoliciesRequest) SetPolicyType(v string) *ListControlPoliciesRequest {
	s.PolicyType = &v
	return s
}

type ListControlPoliciesResponseBody struct {
	// The access control policies.
	ControlPolicies *ListControlPoliciesResponseBodyControlPolicies `json:"ControlPolicies,omitempty" xml:"ControlPolicies,omitempty" type:"Struct"`
	// The number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of access control policies.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListControlPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListControlPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListControlPoliciesResponseBody) SetControlPolicies(v *ListControlPoliciesResponseBodyControlPolicies) *ListControlPoliciesResponseBody {
	s.ControlPolicies = v
	return s
}

func (s *ListControlPoliciesResponseBody) SetPageNumber(v int32) *ListControlPoliciesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListControlPoliciesResponseBody) SetPageSize(v int32) *ListControlPoliciesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListControlPoliciesResponseBody) SetRequestId(v string) *ListControlPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListControlPoliciesResponseBody) SetTotalCount(v int32) *ListControlPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

type ListControlPoliciesResponseBodyControlPolicies struct {
	ControlPolicy []*ListControlPoliciesResponseBodyControlPoliciesControlPolicy `json:"ControlPolicy,omitempty" xml:"ControlPolicy,omitempty" type:"Repeated"`
}

func (s ListControlPoliciesResponseBodyControlPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListControlPoliciesResponseBodyControlPolicies) GoString() string {
	return s.String()
}

func (s *ListControlPoliciesResponseBodyControlPolicies) SetControlPolicy(v []*ListControlPoliciesResponseBodyControlPoliciesControlPolicy) *ListControlPoliciesResponseBodyControlPolicies {
	s.ControlPolicy = v
	return s
}

type ListControlPoliciesResponseBodyControlPoliciesControlPolicy struct {
	// The number of times that the access control policy is referenced.
	AttachmentCount *string `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty"`
	// The time when the access control policy was created.
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the access control policy.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The effective scope of the access control policy. Valid values:
	//
	// - All: The access control policy is in effect for Alibaba Cloud accounts, RAM users, and RAM roles.
	// - RAM: The access control policy is in effect only for RAM users and RAM roles.
	EffectScope *string `json:"EffectScope,omitempty" xml:"EffectScope,omitempty"`
	// The ID of the access control policy.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the access control policy.
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the access control policy. Valid values:
	//
	// - System: system access control policy
	// - Custom: custom access control policy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The time when the access control policy was updated.
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListControlPoliciesResponseBodyControlPoliciesControlPolicy) String() string {
	return tea.Prettify(s)
}

func (s ListControlPoliciesResponseBodyControlPoliciesControlPolicy) GoString() string {
	return s.String()
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) SetAttachmentCount(v string) *ListControlPoliciesResponseBodyControlPoliciesControlPolicy {
	s.AttachmentCount = &v
	return s
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) SetCreateDate(v string) *ListControlPoliciesResponseBodyControlPoliciesControlPolicy {
	s.CreateDate = &v
	return s
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) SetDescription(v string) *ListControlPoliciesResponseBodyControlPoliciesControlPolicy {
	s.Description = &v
	return s
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) SetEffectScope(v string) *ListControlPoliciesResponseBodyControlPoliciesControlPolicy {
	s.EffectScope = &v
	return s
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) SetPolicyId(v string) *ListControlPoliciesResponseBodyControlPoliciesControlPolicy {
	s.PolicyId = &v
	return s
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) SetPolicyName(v string) *ListControlPoliciesResponseBodyControlPoliciesControlPolicy {
	s.PolicyName = &v
	return s
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) SetPolicyType(v string) *ListControlPoliciesResponseBodyControlPoliciesControlPolicy {
	s.PolicyType = &v
	return s
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) SetUpdateDate(v string) *ListControlPoliciesResponseBodyControlPoliciesControlPolicy {
	s.UpdateDate = &v
	return s
}

type ListControlPoliciesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListControlPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListControlPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListControlPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListControlPoliciesResponse) SetHeaders(v map[string]*string) *ListControlPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListControlPoliciesResponse) SetStatusCode(v int32) *ListControlPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListControlPoliciesResponse) SetBody(v *ListControlPoliciesResponseBody) *ListControlPoliciesResponse {
	s.Body = v
	return s
}

type ListControlPolicyAttachmentsForTargetRequest struct {
	// The language in which you want to return the descriptions of the access control policies. Valid values:
	//
	// *   zh-CN (default value): Chinese
	// *   en: English
	// *   ja: Japanese
	//
	// >  This parameter is valid only for system access control policies.
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The ID of the object whose access control policies you want to query. Access control policies can be attached to the following objects:
	//
	// *   Root folder
	// *   Subfolders of the Root folder
	// *   Members
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
}

func (s ListControlPolicyAttachmentsForTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s ListControlPolicyAttachmentsForTargetRequest) GoString() string {
	return s.String()
}

func (s *ListControlPolicyAttachmentsForTargetRequest) SetLanguage(v string) *ListControlPolicyAttachmentsForTargetRequest {
	s.Language = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetRequest) SetTargetId(v string) *ListControlPolicyAttachmentsForTargetRequest {
	s.TargetId = &v
	return s
}

type ListControlPolicyAttachmentsForTargetResponseBody struct {
	// The attached access control policies.
	ControlPolicyAttachments *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachments `json:"ControlPolicyAttachments,omitempty" xml:"ControlPolicyAttachments,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListControlPolicyAttachmentsForTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListControlPolicyAttachmentsForTargetResponseBody) GoString() string {
	return s.String()
}

func (s *ListControlPolicyAttachmentsForTargetResponseBody) SetControlPolicyAttachments(v *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachments) *ListControlPolicyAttachmentsForTargetResponseBody {
	s.ControlPolicyAttachments = v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponseBody) SetRequestId(v string) *ListControlPolicyAttachmentsForTargetResponseBody {
	s.RequestId = &v
	return s
}

type ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachments struct {
	ControlPolicyAttachment []*ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment `json:"ControlPolicyAttachment,omitempty" xml:"ControlPolicyAttachment,omitempty" type:"Repeated"`
}

func (s ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachments) String() string {
	return tea.Prettify(s)
}

func (s ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachments) GoString() string {
	return s.String()
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachments) SetControlPolicyAttachment(v []*ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachments {
	s.ControlPolicyAttachment = v
	return s
}

type ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment struct {
	// The time when the access control policy was attached.
	AttachDate *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty"`
	// The description of the access control policy.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The effective scope of the access control policy. Valid values:
	//
	// *   All: The access control policy is in effect for Alibaba Cloud accounts, RAM users, and RAM roles.
	// *   RAM: The access control policy is in effect only for RAM users and RAM roles.
	EffectScope *string `json:"EffectScope,omitempty" xml:"EffectScope,omitempty"`
	// The ID of the access control policy.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the access control policy.
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the access control policy. Valid values:
	//
	// *   System: system access control policy
	// *   Custom: custom access control policy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) String() string {
	return tea.Prettify(s)
}

func (s ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) GoString() string {
	return s.String()
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) SetAttachDate(v string) *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment {
	s.AttachDate = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) SetDescription(v string) *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment {
	s.Description = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) SetEffectScope(v string) *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment {
	s.EffectScope = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) SetPolicyId(v string) *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment {
	s.PolicyId = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) SetPolicyName(v string) *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment {
	s.PolicyName = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) SetPolicyType(v string) *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment {
	s.PolicyType = &v
	return s
}

type ListControlPolicyAttachmentsForTargetResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListControlPolicyAttachmentsForTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListControlPolicyAttachmentsForTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s ListControlPolicyAttachmentsForTargetResponse) GoString() string {
	return s.String()
}

func (s *ListControlPolicyAttachmentsForTargetResponse) SetHeaders(v map[string]*string) *ListControlPolicyAttachmentsForTargetResponse {
	s.Headers = v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponse) SetStatusCode(v int32) *ListControlPolicyAttachmentsForTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponse) SetBody(v *ListControlPolicyAttachmentsForTargetResponseBody) *ListControlPolicyAttachmentsForTargetResponse {
	s.Body = v
	return s
}

type ListDelegatedAdministratorsRequest struct {
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The identifier of the trusted service.
	//
	// For more information, see the `Trusted service identifier` column in [Supported trusted services](~~208133~~).
	ServicePrincipal *string `json:"ServicePrincipal,omitempty" xml:"ServicePrincipal,omitempty"`
}

func (s ListDelegatedAdministratorsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDelegatedAdministratorsRequest) GoString() string {
	return s.String()
}

func (s *ListDelegatedAdministratorsRequest) SetPageNumber(v int64) *ListDelegatedAdministratorsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDelegatedAdministratorsRequest) SetPageSize(v int64) *ListDelegatedAdministratorsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDelegatedAdministratorsRequest) SetServicePrincipal(v string) *ListDelegatedAdministratorsRequest {
	s.ServicePrincipal = &v
	return s
}

type ListDelegatedAdministratorsResponseBody struct {
	// The information of the delegated administrator accounts.
	Accounts *ListDelegatedAdministratorsResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
	// The page number of the returned page.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDelegatedAdministratorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDelegatedAdministratorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDelegatedAdministratorsResponseBody) SetAccounts(v *ListDelegatedAdministratorsResponseBodyAccounts) *ListDelegatedAdministratorsResponseBody {
	s.Accounts = v
	return s
}

func (s *ListDelegatedAdministratorsResponseBody) SetPageNumber(v int64) *ListDelegatedAdministratorsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDelegatedAdministratorsResponseBody) SetPageSize(v int64) *ListDelegatedAdministratorsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDelegatedAdministratorsResponseBody) SetRequestId(v string) *ListDelegatedAdministratorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDelegatedAdministratorsResponseBody) SetTotalCount(v int64) *ListDelegatedAdministratorsResponseBody {
	s.TotalCount = &v
	return s
}

type ListDelegatedAdministratorsResponseBodyAccounts struct {
	Account []*ListDelegatedAdministratorsResponseBodyAccountsAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Repeated"`
}

func (s ListDelegatedAdministratorsResponseBodyAccounts) String() string {
	return tea.Prettify(s)
}

func (s ListDelegatedAdministratorsResponseBodyAccounts) GoString() string {
	return s.String()
}

func (s *ListDelegatedAdministratorsResponseBodyAccounts) SetAccount(v []*ListDelegatedAdministratorsResponseBodyAccountsAccount) *ListDelegatedAdministratorsResponseBodyAccounts {
	s.Account = v
	return s
}

type ListDelegatedAdministratorsResponseBodyAccountsAccount struct {
	// The ID of the member.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The time when the member was specified as a delegated administrator account.
	DelegationEnabledTime *string `json:"DelegationEnabledTime,omitempty" xml:"DelegationEnabledTime,omitempty"`
	// The display name of the member.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The way in which the member joins the resource directory. Valid values:
	//
	// *   invited: The member is invited to join the resource directory.
	// *   created: The member is directly created in the resource directory.
	JoinMethod *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	// The identifier of the trusted service.
	ServicePrincipal *string `json:"ServicePrincipal,omitempty" xml:"ServicePrincipal,omitempty"`
}

func (s ListDelegatedAdministratorsResponseBodyAccountsAccount) String() string {
	return tea.Prettify(s)
}

func (s ListDelegatedAdministratorsResponseBodyAccountsAccount) GoString() string {
	return s.String()
}

func (s *ListDelegatedAdministratorsResponseBodyAccountsAccount) SetAccountId(v string) *ListDelegatedAdministratorsResponseBodyAccountsAccount {
	s.AccountId = &v
	return s
}

func (s *ListDelegatedAdministratorsResponseBodyAccountsAccount) SetDelegationEnabledTime(v string) *ListDelegatedAdministratorsResponseBodyAccountsAccount {
	s.DelegationEnabledTime = &v
	return s
}

func (s *ListDelegatedAdministratorsResponseBodyAccountsAccount) SetDisplayName(v string) *ListDelegatedAdministratorsResponseBodyAccountsAccount {
	s.DisplayName = &v
	return s
}

func (s *ListDelegatedAdministratorsResponseBodyAccountsAccount) SetJoinMethod(v string) *ListDelegatedAdministratorsResponseBodyAccountsAccount {
	s.JoinMethod = &v
	return s
}

func (s *ListDelegatedAdministratorsResponseBodyAccountsAccount) SetServicePrincipal(v string) *ListDelegatedAdministratorsResponseBodyAccountsAccount {
	s.ServicePrincipal = &v
	return s
}

type ListDelegatedAdministratorsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDelegatedAdministratorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDelegatedAdministratorsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDelegatedAdministratorsResponse) GoString() string {
	return s.String()
}

func (s *ListDelegatedAdministratorsResponse) SetHeaders(v map[string]*string) *ListDelegatedAdministratorsResponse {
	s.Headers = v
	return s
}

func (s *ListDelegatedAdministratorsResponse) SetStatusCode(v int32) *ListDelegatedAdministratorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDelegatedAdministratorsResponse) SetBody(v *ListDelegatedAdministratorsResponseBody) *ListDelegatedAdministratorsResponse {
	s.Body = v
	return s
}

type ListDelegatedServicesForAccountRequest struct {
	// The ID of the member.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s ListDelegatedServicesForAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDelegatedServicesForAccountRequest) GoString() string {
	return s.String()
}

func (s *ListDelegatedServicesForAccountRequest) SetAccountId(v string) *ListDelegatedServicesForAccountRequest {
	s.AccountId = &v
	return s
}

type ListDelegatedServicesForAccountResponseBody struct {
	// The trusted services.
	//
	// >  If the value of this parameter is empty, the member is not specified as a delegated administrator account.
	DelegatedServices *ListDelegatedServicesForAccountResponseBodyDelegatedServices `json:"DelegatedServices,omitempty" xml:"DelegatedServices,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDelegatedServicesForAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDelegatedServicesForAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ListDelegatedServicesForAccountResponseBody) SetDelegatedServices(v *ListDelegatedServicesForAccountResponseBodyDelegatedServices) *ListDelegatedServicesForAccountResponseBody {
	s.DelegatedServices = v
	return s
}

func (s *ListDelegatedServicesForAccountResponseBody) SetRequestId(v string) *ListDelegatedServicesForAccountResponseBody {
	s.RequestId = &v
	return s
}

type ListDelegatedServicesForAccountResponseBodyDelegatedServices struct {
	DelegatedService []*ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService `json:"DelegatedService,omitempty" xml:"DelegatedService,omitempty" type:"Repeated"`
}

func (s ListDelegatedServicesForAccountResponseBodyDelegatedServices) String() string {
	return tea.Prettify(s)
}

func (s ListDelegatedServicesForAccountResponseBodyDelegatedServices) GoString() string {
	return s.String()
}

func (s *ListDelegatedServicesForAccountResponseBodyDelegatedServices) SetDelegatedService(v []*ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService) *ListDelegatedServicesForAccountResponseBodyDelegatedServices {
	s.DelegatedService = v
	return s
}

type ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService struct {
	// The time when the member was specified as a delegated administrator account of the trusted service.
	DelegationEnabledTime *string `json:"DelegationEnabledTime,omitempty" xml:"DelegationEnabledTime,omitempty"`
	// The identification of the trusted service.
	ServicePrincipal *string `json:"ServicePrincipal,omitempty" xml:"ServicePrincipal,omitempty"`
	// The status of the trusted service. Valid values:
	//
	// *   ENABLED: enabled
	// *   DISABLED: disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService) String() string {
	return tea.Prettify(s)
}

func (s ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService) GoString() string {
	return s.String()
}

func (s *ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService) SetDelegationEnabledTime(v string) *ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService {
	s.DelegationEnabledTime = &v
	return s
}

func (s *ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService) SetServicePrincipal(v string) *ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService {
	s.ServicePrincipal = &v
	return s
}

func (s *ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService) SetStatus(v string) *ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService {
	s.Status = &v
	return s
}

type ListDelegatedServicesForAccountResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDelegatedServicesForAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDelegatedServicesForAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDelegatedServicesForAccountResponse) GoString() string {
	return s.String()
}

func (s *ListDelegatedServicesForAccountResponse) SetHeaders(v map[string]*string) *ListDelegatedServicesForAccountResponse {
	s.Headers = v
	return s
}

func (s *ListDelegatedServicesForAccountResponse) SetStatusCode(v int32) *ListDelegatedServicesForAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDelegatedServicesForAccountResponse) SetBody(v *ListDelegatedServicesForAccountResponseBody) *ListDelegatedServicesForAccountResponse {
	s.Body = v
	return s
}

type ListFoldersForParentRequest struct {
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the parent folder.
	//
	// If you leave this parameter empty, the information of the first-level subfolders of the Root folder is queried.
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	// The keyword used for the query, such as a folder name.
	//
	// Fuzzy match is supported.
	QueryKeyword *string `json:"QueryKeyword,omitempty" xml:"QueryKeyword,omitempty"`
}

func (s ListFoldersForParentRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFoldersForParentRequest) GoString() string {
	return s.String()
}

func (s *ListFoldersForParentRequest) SetPageNumber(v int32) *ListFoldersForParentRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFoldersForParentRequest) SetPageSize(v int32) *ListFoldersForParentRequest {
	s.PageSize = &v
	return s
}

func (s *ListFoldersForParentRequest) SetParentFolderId(v string) *ListFoldersForParentRequest {
	s.ParentFolderId = &v
	return s
}

func (s *ListFoldersForParentRequest) SetQueryKeyword(v string) *ListFoldersForParentRequest {
	s.QueryKeyword = &v
	return s
}

type ListFoldersForParentResponseBody struct {
	// The information of the folders.
	Folders *ListFoldersForParentResponseBodyFolders `json:"Folders,omitempty" xml:"Folders,omitempty" type:"Struct"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFoldersForParentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFoldersForParentResponseBody) GoString() string {
	return s.String()
}

func (s *ListFoldersForParentResponseBody) SetFolders(v *ListFoldersForParentResponseBodyFolders) *ListFoldersForParentResponseBody {
	s.Folders = v
	return s
}

func (s *ListFoldersForParentResponseBody) SetPageNumber(v int32) *ListFoldersForParentResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFoldersForParentResponseBody) SetPageSize(v int32) *ListFoldersForParentResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFoldersForParentResponseBody) SetRequestId(v string) *ListFoldersForParentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFoldersForParentResponseBody) SetTotalCount(v int32) *ListFoldersForParentResponseBody {
	s.TotalCount = &v
	return s
}

type ListFoldersForParentResponseBodyFolders struct {
	Folder []*ListFoldersForParentResponseBodyFoldersFolder `json:"Folder,omitempty" xml:"Folder,omitempty" type:"Repeated"`
}

func (s ListFoldersForParentResponseBodyFolders) String() string {
	return tea.Prettify(s)
}

func (s ListFoldersForParentResponseBodyFolders) GoString() string {
	return s.String()
}

func (s *ListFoldersForParentResponseBodyFolders) SetFolder(v []*ListFoldersForParentResponseBodyFoldersFolder) *ListFoldersForParentResponseBodyFolders {
	s.Folder = v
	return s
}

type ListFoldersForParentResponseBodyFoldersFolder struct {
	// The time when the folder was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the folder.
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The name of the folder.
	FolderName *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
}

func (s ListFoldersForParentResponseBodyFoldersFolder) String() string {
	return tea.Prettify(s)
}

func (s ListFoldersForParentResponseBodyFoldersFolder) GoString() string {
	return s.String()
}

func (s *ListFoldersForParentResponseBodyFoldersFolder) SetCreateTime(v string) *ListFoldersForParentResponseBodyFoldersFolder {
	s.CreateTime = &v
	return s
}

func (s *ListFoldersForParentResponseBodyFoldersFolder) SetFolderId(v string) *ListFoldersForParentResponseBodyFoldersFolder {
	s.FolderId = &v
	return s
}

func (s *ListFoldersForParentResponseBodyFoldersFolder) SetFolderName(v string) *ListFoldersForParentResponseBodyFoldersFolder {
	s.FolderName = &v
	return s
}

type ListFoldersForParentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFoldersForParentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFoldersForParentResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFoldersForParentResponse) GoString() string {
	return s.String()
}

func (s *ListFoldersForParentResponse) SetHeaders(v map[string]*string) *ListFoldersForParentResponse {
	s.Headers = v
	return s
}

func (s *ListFoldersForParentResponse) SetStatusCode(v int32) *ListFoldersForParentResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFoldersForParentResponse) SetBody(v *ListFoldersForParentResponseBody) *ListFoldersForParentResponse {
	s.Body = v
	return s
}

type ListHandshakesForAccountRequest struct {
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListHandshakesForAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHandshakesForAccountRequest) GoString() string {
	return s.String()
}

func (s *ListHandshakesForAccountRequest) SetPageNumber(v int32) *ListHandshakesForAccountRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHandshakesForAccountRequest) SetPageSize(v int32) *ListHandshakesForAccountRequest {
	s.PageSize = &v
	return s
}

type ListHandshakesForAccountResponseBody struct {
	// The information of the invitations.
	Handshakes *ListHandshakesForAccountResponseBodyHandshakes `json:"Handshakes,omitempty" xml:"Handshakes,omitempty" type:"Struct"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of invitations.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHandshakesForAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHandshakesForAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ListHandshakesForAccountResponseBody) SetHandshakes(v *ListHandshakesForAccountResponseBodyHandshakes) *ListHandshakesForAccountResponseBody {
	s.Handshakes = v
	return s
}

func (s *ListHandshakesForAccountResponseBody) SetPageNumber(v int32) *ListHandshakesForAccountResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListHandshakesForAccountResponseBody) SetPageSize(v int32) *ListHandshakesForAccountResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListHandshakesForAccountResponseBody) SetRequestId(v string) *ListHandshakesForAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHandshakesForAccountResponseBody) SetTotalCount(v int32) *ListHandshakesForAccountResponseBody {
	s.TotalCount = &v
	return s
}

type ListHandshakesForAccountResponseBodyHandshakes struct {
	Handshake []*ListHandshakesForAccountResponseBodyHandshakesHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" type:"Repeated"`
}

func (s ListHandshakesForAccountResponseBodyHandshakes) String() string {
	return tea.Prettify(s)
}

func (s ListHandshakesForAccountResponseBodyHandshakes) GoString() string {
	return s.String()
}

func (s *ListHandshakesForAccountResponseBodyHandshakes) SetHandshake(v []*ListHandshakesForAccountResponseBodyHandshakesHandshake) *ListHandshakesForAccountResponseBodyHandshakes {
	s.Handshake = v
	return s
}

type ListHandshakesForAccountResponseBodyHandshakesHandshake struct {
	// The time when the invitation was created. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the invitation expires. The time is displayed in UTC.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the invitation.
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
	// The ID of the management account of the resource directory.
	MasterAccountId *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	// The name of the management account of the resource directory.
	MasterAccountName *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	// The time when the invitation was modified. The time is displayed in UTC.
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The comment on the invitation.
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The ID of the resource directory.
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the invitation. Valid values:
	//
	// *   Pending: The invitation is waiting for confirmation.
	// *   Accepted: The invitation is accepted.
	// *   Cancelled: The invitation is canceled.
	// *   Declined: The invitation is rejected.
	// *   Expired: The invitation expires.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID or logon email address of the invited Alibaba Cloud account.
	TargetEntity *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty"`
	// The type of the invited Alibaba Cloud account. Valid values:
	//
	// *   Account: indicates the ID of the account.
	// *   Email: indicates the logon email address of the account.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListHandshakesForAccountResponseBodyHandshakesHandshake) String() string {
	return tea.Prettify(s)
}

func (s ListHandshakesForAccountResponseBodyHandshakesHandshake) GoString() string {
	return s.String()
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) SetCreateTime(v string) *ListHandshakesForAccountResponseBodyHandshakesHandshake {
	s.CreateTime = &v
	return s
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) SetExpireTime(v string) *ListHandshakesForAccountResponseBodyHandshakesHandshake {
	s.ExpireTime = &v
	return s
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) SetHandshakeId(v string) *ListHandshakesForAccountResponseBodyHandshakesHandshake {
	s.HandshakeId = &v
	return s
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) SetMasterAccountId(v string) *ListHandshakesForAccountResponseBodyHandshakesHandshake {
	s.MasterAccountId = &v
	return s
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) SetMasterAccountName(v string) *ListHandshakesForAccountResponseBodyHandshakesHandshake {
	s.MasterAccountName = &v
	return s
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) SetModifyTime(v string) *ListHandshakesForAccountResponseBodyHandshakesHandshake {
	s.ModifyTime = &v
	return s
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) SetNote(v string) *ListHandshakesForAccountResponseBodyHandshakesHandshake {
	s.Note = &v
	return s
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) SetResourceDirectoryId(v string) *ListHandshakesForAccountResponseBodyHandshakesHandshake {
	s.ResourceDirectoryId = &v
	return s
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) SetStatus(v string) *ListHandshakesForAccountResponseBodyHandshakesHandshake {
	s.Status = &v
	return s
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) SetTargetEntity(v string) *ListHandshakesForAccountResponseBodyHandshakesHandshake {
	s.TargetEntity = &v
	return s
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) SetTargetType(v string) *ListHandshakesForAccountResponseBodyHandshakesHandshake {
	s.TargetType = &v
	return s
}

type ListHandshakesForAccountResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHandshakesForAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHandshakesForAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHandshakesForAccountResponse) GoString() string {
	return s.String()
}

func (s *ListHandshakesForAccountResponse) SetHeaders(v map[string]*string) *ListHandshakesForAccountResponse {
	s.Headers = v
	return s
}

func (s *ListHandshakesForAccountResponse) SetStatusCode(v int32) *ListHandshakesForAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHandshakesForAccountResponse) SetBody(v *ListHandshakesForAccountResponseBody) *ListHandshakesForAccountResponse {
	s.Body = v
	return s
}

type ListHandshakesForResourceDirectoryRequest struct {
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListHandshakesForResourceDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHandshakesForResourceDirectoryRequest) GoString() string {
	return s.String()
}

func (s *ListHandshakesForResourceDirectoryRequest) SetPageNumber(v int32) *ListHandshakesForResourceDirectoryRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryRequest) SetPageSize(v int32) *ListHandshakesForResourceDirectoryRequest {
	s.PageSize = &v
	return s
}

type ListHandshakesForResourceDirectoryResponseBody struct {
	// The information of the invitations.
	Handshakes *ListHandshakesForResourceDirectoryResponseBodyHandshakes `json:"Handshakes,omitempty" xml:"Handshakes,omitempty" type:"Struct"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHandshakesForResourceDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHandshakesForResourceDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListHandshakesForResourceDirectoryResponseBody) SetHandshakes(v *ListHandshakesForResourceDirectoryResponseBodyHandshakes) *ListHandshakesForResourceDirectoryResponseBody {
	s.Handshakes = v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseBody) SetPageNumber(v int32) *ListHandshakesForResourceDirectoryResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseBody) SetPageSize(v int32) *ListHandshakesForResourceDirectoryResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseBody) SetRequestId(v string) *ListHandshakesForResourceDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseBody) SetTotalCount(v int32) *ListHandshakesForResourceDirectoryResponseBody {
	s.TotalCount = &v
	return s
}

type ListHandshakesForResourceDirectoryResponseBodyHandshakes struct {
	Handshake []*ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" type:"Repeated"`
}

func (s ListHandshakesForResourceDirectoryResponseBodyHandshakes) String() string {
	return tea.Prettify(s)
}

func (s ListHandshakesForResourceDirectoryResponseBodyHandshakes) GoString() string {
	return s.String()
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakes) SetHandshake(v []*ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) *ListHandshakesForResourceDirectoryResponseBodyHandshakes {
	s.Handshake = v
	return s
}

type ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake struct {
	// The time when the invitation was created. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the invitation expires. The time is displayed in UTC.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the invitation.
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
	// The ID of the management account of the resource directory.
	MasterAccountId *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	// The name of the management account of the resource directory.
	MasterAccountName *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	// The time when the invitation was modified. The time is displayed in UTC.
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The comment on the invitation.
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The ID of the resource directory.
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the invitation. Valid values:
	//
	// *   Pending: The invitation is waiting for confirmation.
	// *   Accepted: The invitation is accepted.
	// *   Cancelled: The invitation is canceled.
	// *   Declined: The invitation is rejected.
	// *   Expired: The invitation expires.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID or logon email address of the invited account.
	TargetEntity *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty"`
	// The type of the invited account. Valid values:
	//
	// *   Account: indicates the ID of the account.
	// *   Email: indicates the logon email address of the account.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) String() string {
	return tea.Prettify(s)
}

func (s ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) GoString() string {
	return s.String()
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) SetCreateTime(v string) *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake {
	s.CreateTime = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) SetExpireTime(v string) *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake {
	s.ExpireTime = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) SetHandshakeId(v string) *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake {
	s.HandshakeId = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) SetMasterAccountId(v string) *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake {
	s.MasterAccountId = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) SetMasterAccountName(v string) *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake {
	s.MasterAccountName = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) SetModifyTime(v string) *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake {
	s.ModifyTime = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) SetNote(v string) *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake {
	s.Note = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) SetResourceDirectoryId(v string) *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake {
	s.ResourceDirectoryId = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) SetStatus(v string) *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake {
	s.Status = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) SetTargetEntity(v string) *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake {
	s.TargetEntity = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) SetTargetType(v string) *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake {
	s.TargetType = &v
	return s
}

type ListHandshakesForResourceDirectoryResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHandshakesForResourceDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHandshakesForResourceDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHandshakesForResourceDirectoryResponse) GoString() string {
	return s.String()
}

func (s *ListHandshakesForResourceDirectoryResponse) SetHeaders(v map[string]*string) *ListHandshakesForResourceDirectoryResponse {
	s.Headers = v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponse) SetStatusCode(v int32) *ListHandshakesForResourceDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponse) SetBody(v *ListHandshakesForResourceDirectoryResponseBody) *ListHandshakesForResourceDirectoryResponse {
	s.Body = v
	return s
}

type ListPoliciesRequest struct {
	// The language that is used to return the description of the system policy. Valid values:
	//
	// *   en: English
	// *   zh-CN: Chinese
	// *   ja: Japanese
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the policy. If you do not specify this parameter, the system lists all types of policies. Valid values:
	//
	// *   Custom: custom policy
	// *   System: system policy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListPoliciesRequest) SetLanguage(v string) *ListPoliciesRequest {
	s.Language = &v
	return s
}

func (s *ListPoliciesRequest) SetPageNumber(v int32) *ListPoliciesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPoliciesRequest) SetPageSize(v int32) *ListPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *ListPoliciesRequest) SetPolicyType(v string) *ListPoliciesRequest {
	s.PolicyType = &v
	return s
}

type ListPoliciesResponseBody struct {
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The information of the policies.
	Policies *ListPoliciesResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBody) SetPageNumber(v int32) *ListPoliciesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPoliciesResponseBody) SetPageSize(v int32) *ListPoliciesResponseBody {
	s.PageSize = &v
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

func (s *ListPoliciesResponseBody) SetTotalCount(v int32) *ListPoliciesResponseBody {
	s.TotalCount = &v
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
	// The number of times the policy is referenced.
	AttachmentCount *int32 `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty"`
	// The time when the policy was created.
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The default version of the policy.
	DefaultVersion *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// The description of the policy.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the policy.
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy. Valid values:
	//
	// *   Custom: custom policy
	// *   System: system policy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The time when the policy was updated.
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

func (s *ListPoliciesResponseBodyPoliciesPolicy) SetUpdateDate(v string) *ListPoliciesResponseBodyPoliciesPolicy {
	s.UpdateDate = &v
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

type ListPolicyAttachmentsRequest struct {
	// The language that is used to return the description of the system policy. Valid values:
	//
	// *   en: English
	// *   zh-CN: Chinese
	// *   ja: Japanese
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the policy.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy. If you do not specify this parameter, the system lists all types of policies. Valid values:
	//
	// *   Custom: custom policy
	// *   System: system policy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The name of the object to which the policy is attached.
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The type of the object to which the policy is attached. If you do not specify this parameter, the system lists all types of objects. Valid values:
	//
	// *   IMSUser: RAM user
	// *   IMSGroup: RAM user group
	// *   ServiceRole: RAM role
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The ID of the resource group or the ID of the Alibaba Cloud account to which the resource group belongs. If you do not specify this parameter, the system lists all policy attachment records under the current account.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListPolicyAttachmentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyAttachmentsRequest) GoString() string {
	return s.String()
}

func (s *ListPolicyAttachmentsRequest) SetLanguage(v string) *ListPolicyAttachmentsRequest {
	s.Language = &v
	return s
}

func (s *ListPolicyAttachmentsRequest) SetPageNumber(v int32) *ListPolicyAttachmentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPolicyAttachmentsRequest) SetPageSize(v int32) *ListPolicyAttachmentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPolicyAttachmentsRequest) SetPolicyName(v string) *ListPolicyAttachmentsRequest {
	s.PolicyName = &v
	return s
}

func (s *ListPolicyAttachmentsRequest) SetPolicyType(v string) *ListPolicyAttachmentsRequest {
	s.PolicyType = &v
	return s
}

func (s *ListPolicyAttachmentsRequest) SetPrincipalName(v string) *ListPolicyAttachmentsRequest {
	s.PrincipalName = &v
	return s
}

func (s *ListPolicyAttachmentsRequest) SetPrincipalType(v string) *ListPolicyAttachmentsRequest {
	s.PrincipalType = &v
	return s
}

func (s *ListPolicyAttachmentsRequest) SetResourceGroupId(v string) *ListPolicyAttachmentsRequest {
	s.ResourceGroupId = &v
	return s
}

type ListPolicyAttachmentsResponseBody struct {
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The policy attachment records.
	PolicyAttachments *ListPolicyAttachmentsResponseBodyPolicyAttachments `json:"PolicyAttachments,omitempty" xml:"PolicyAttachments,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPolicyAttachmentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyAttachmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPolicyAttachmentsResponseBody) SetPageNumber(v int32) *ListPolicyAttachmentsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPolicyAttachmentsResponseBody) SetPageSize(v int32) *ListPolicyAttachmentsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPolicyAttachmentsResponseBody) SetPolicyAttachments(v *ListPolicyAttachmentsResponseBodyPolicyAttachments) *ListPolicyAttachmentsResponseBody {
	s.PolicyAttachments = v
	return s
}

func (s *ListPolicyAttachmentsResponseBody) SetRequestId(v string) *ListPolicyAttachmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPolicyAttachmentsResponseBody) SetTotalCount(v int32) *ListPolicyAttachmentsResponseBody {
	s.TotalCount = &v
	return s
}

type ListPolicyAttachmentsResponseBodyPolicyAttachments struct {
	PolicyAttachment []*ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment `json:"PolicyAttachment,omitempty" xml:"PolicyAttachment,omitempty" type:"Repeated"`
}

func (s ListPolicyAttachmentsResponseBodyPolicyAttachments) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyAttachmentsResponseBodyPolicyAttachments) GoString() string {
	return s.String()
}

func (s *ListPolicyAttachmentsResponseBodyPolicyAttachments) SetPolicyAttachment(v []*ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment) *ListPolicyAttachmentsResponseBodyPolicyAttachments {
	s.PolicyAttachment = v
	return s
}

type ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment struct {
	// The time when the policy was attached.
	AttachDate *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty"`
	// The description of the policy.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the policy.
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy. Valid values:
	//
	// *   Custom: custom policy
	// *   System: system policy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The name of the object to which the policy is attached.
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The type of the object to which the policy is attached. Valid values:
	//
	// *   IMSUser: RAM user
	// *   IMSGroup: RAM user group
	// *   ServiceRole: RAM role
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The ID of the resource group or the ID of the Alibaba Cloud account to which the resource group belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment) GoString() string {
	return s.String()
}

func (s *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment) SetAttachDate(v string) *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment {
	s.AttachDate = &v
	return s
}

func (s *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment) SetDescription(v string) *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment {
	s.Description = &v
	return s
}

func (s *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment) SetPolicyName(v string) *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment {
	s.PolicyName = &v
	return s
}

func (s *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment) SetPolicyType(v string) *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment {
	s.PolicyType = &v
	return s
}

func (s *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment) SetPrincipalName(v string) *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment {
	s.PrincipalName = &v
	return s
}

func (s *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment) SetPrincipalType(v string) *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment {
	s.PrincipalType = &v
	return s
}

func (s *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment) SetResourceGroupId(v string) *ListPolicyAttachmentsResponseBodyPolicyAttachmentsPolicyAttachment {
	s.ResourceGroupId = &v
	return s
}

type ListPolicyAttachmentsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPolicyAttachmentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPolicyAttachmentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyAttachmentsResponse) GoString() string {
	return s.String()
}

func (s *ListPolicyAttachmentsResponse) SetHeaders(v map[string]*string) *ListPolicyAttachmentsResponse {
	s.Headers = v
	return s
}

func (s *ListPolicyAttachmentsResponse) SetStatusCode(v int32) *ListPolicyAttachmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPolicyAttachmentsResponse) SetBody(v *ListPolicyAttachmentsResponseBody) *ListPolicyAttachmentsResponse {
	s.Body = v
	return s
}

type ListPolicyVersionsRequest struct {
	// The name of the policy.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy. Valid values:
	//
	// *   Custom: custom policy
	// *   System: system policy
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
	// The information of the policy versions.
	PolicyVersions *ListPolicyVersionsResponseBodyPolicyVersions `json:"PolicyVersions,omitempty" xml:"PolicyVersions,omitempty" type:"Struct"`
	// The ID of the request.
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
	// The time when the policy version was created.
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// Indicates whether the policy version is the default version.
	IsDefaultVersion *bool `json:"IsDefaultVersion,omitempty" xml:"IsDefaultVersion,omitempty"`
	// The ID of the policy version.
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

type ListResourceGroupsRequest struct {
	// The display name of the resource group. This parameter specifies a filter condition for the query. Fuzzy match is supported.
	//
	// The display name can be a maximum of 50 characters in length.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Specifies whether to return the information of tags. Valid values:
	//
	// *   false (default value)
	// *   true
	//
	// >  If you configure the Tag parameter, the system returns the information of tags regardless of the setting of the `IncludeTags` parameter.
	IncludeTags *bool `json:"IncludeTags,omitempty" xml:"IncludeTags,omitempty"`
	// The identifier of the resource group. This parameter specifies a filter condition for the query. Fuzzy match is supported.
	//
	// The identifier can be a maximum of 50 characters in length and can contain letters, digits, and hyphens (-).
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group. This parameter specifies a filter condition for the query.
	//
	// The ID can be a maximum of 18 characters in length and must start with `rg-`.
	//
	// >  This parameter is incorporated into the `ResourceGroupIds` parameter. If you configure both the `ResourceGroupId` and `ResourceGroupIds` parameters, the value of the `ResourceGroupIds` parameter prevails.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The IDs of the resource groups. This parameter specifies a filter condition for the query.
	//
	// You can specify a maximum of 100 resource group IDs.
	//
	// >  If you configure both the `ResourceGroupId` and `ResourceGroupIds` parameters, the value of the `ResourceGroupIds` parameter prevails.
	ResourceGroupIds []*string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
	// The status of the resource group. This parameter specifies a filter condition for the query. Valid values:
	//
	// *   Creating: The resource group is being created.
	// *   OK: The resource group is created.
	// *   PendingDelete: The resource group is waiting to be deleted.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag. This parameter specifies a filter condition for the query.
	Tag []*ListResourceGroupsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListResourceGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsRequest) SetDisplayName(v string) *ListResourceGroupsRequest {
	s.DisplayName = &v
	return s
}

func (s *ListResourceGroupsRequest) SetIncludeTags(v bool) *ListResourceGroupsRequest {
	s.IncludeTags = &v
	return s
}

func (s *ListResourceGroupsRequest) SetName(v string) *ListResourceGroupsRequest {
	s.Name = &v
	return s
}

func (s *ListResourceGroupsRequest) SetPageNumber(v int32) *ListResourceGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceGroupsRequest) SetPageSize(v int32) *ListResourceGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourceGroupsRequest) SetResourceGroupId(v string) *ListResourceGroupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListResourceGroupsRequest) SetResourceGroupIds(v []*string) *ListResourceGroupsRequest {
	s.ResourceGroupIds = v
	return s
}

func (s *ListResourceGroupsRequest) SetStatus(v string) *ListResourceGroupsRequest {
	s.Status = &v
	return s
}

func (s *ListResourceGroupsRequest) SetTag(v []*ListResourceGroupsRequestTag) *ListResourceGroupsRequest {
	s.Tag = v
	return s
}

type ListResourceGroupsRequestTag struct {
	// The tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListResourceGroupsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsRequestTag) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsRequestTag) SetKey(v string) *ListResourceGroupsRequestTag {
	s.Key = &v
	return s
}

func (s *ListResourceGroupsRequestTag) SetValue(v string) *ListResourceGroupsRequestTag {
	s.Value = &v
	return s
}

type ListResourceGroupsResponseBody struct {
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the resource groups.
	ResourceGroups *ListResourceGroupsResponseBodyResourceGroups `json:"ResourceGroups,omitempty" xml:"ResourceGroups,omitempty" type:"Struct"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourceGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBody) SetPageNumber(v int32) *ListResourceGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResourceGroupsResponseBody) SetPageSize(v int32) *ListResourceGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListResourceGroupsResponseBody) SetRequestId(v string) *ListResourceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceGroupsResponseBody) SetResourceGroups(v *ListResourceGroupsResponseBodyResourceGroups) *ListResourceGroupsResponseBody {
	s.ResourceGroups = v
	return s
}

func (s *ListResourceGroupsResponseBody) SetTotalCount(v int32) *ListResourceGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListResourceGroupsResponseBodyResourceGroups struct {
	ResourceGroup []*ListResourceGroupsResponseBodyResourceGroupsResourceGroup `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty" type:"Repeated"`
}

func (s ListResourceGroupsResponseBodyResourceGroups) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsResponseBodyResourceGroups) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBodyResourceGroups) SetResourceGroup(v []*ListResourceGroupsResponseBodyResourceGroupsResourceGroup) *ListResourceGroupsResponseBodyResourceGroups {
	s.ResourceGroup = v
	return s
}

type ListResourceGroupsResponseBodyResourceGroupsResourceGroup struct {
	// The ID of the Alibaba Cloud account to which the resource group belongs.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The time when the resource group was created. The time is displayed in UTC.
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The display name of the resource group.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the resource group.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The identifier of the resource group.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the resource group. Valid values:
	//
	// *   Creating: The resource group is being created.
	// *   OK: The resource group is created.
	// *   PendingDelete: The resource group is waiting to be deleted.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags that are added to the resource group.
	Tags *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s ListResourceGroupsResponseBodyResourceGroupsResourceGroup) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsResponseBodyResourceGroupsResourceGroup) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroup) SetAccountId(v string) *ListResourceGroupsResponseBodyResourceGroupsResourceGroup {
	s.AccountId = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroup) SetCreateDate(v string) *ListResourceGroupsResponseBodyResourceGroupsResourceGroup {
	s.CreateDate = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroup) SetDisplayName(v string) *ListResourceGroupsResponseBodyResourceGroupsResourceGroup {
	s.DisplayName = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroup) SetId(v string) *ListResourceGroupsResponseBodyResourceGroupsResourceGroup {
	s.Id = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroup) SetName(v string) *ListResourceGroupsResponseBodyResourceGroupsResourceGroup {
	s.Name = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroup) SetStatus(v string) *ListResourceGroupsResponseBodyResourceGroupsResourceGroup {
	s.Status = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroup) SetTags(v *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTags) *ListResourceGroupsResponseBodyResourceGroupsResourceGroup {
	s.Tags = v
	return s
}

type ListResourceGroupsResponseBodyResourceGroupsResourceGroupTags struct {
	Tag []*ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListResourceGroupsResponseBodyResourceGroupsResourceGroupTags) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsResponseBodyResourceGroupsResourceGroupTags) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTags) SetTag(v []*ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag) *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTags {
	s.Tag = v
	return s
}

type ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag struct {
	// The tag key.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag) SetTagKey(v string) *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag {
	s.TagKey = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag) SetTagValue(v string) *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag {
	s.TagValue = &v
	return s
}

type ListResourceGroupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponse) SetHeaders(v map[string]*string) *ListResourceGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceGroupsResponse) SetStatusCode(v int32) *ListResourceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceGroupsResponse) SetBody(v *ListResourceGroupsResponseBody) *ListResourceGroupsResponse {
	s.Body = v
	return s
}

type ListResourcesRequest struct {
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// For more information about the supported resource types, see the **Resource type** column in [Services that work with Resource Group](~~94479~~).
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The resource types. A maximum of 50 resource types are supported.
	//
	// >  If you configure `ResourceTypes`, you must configure both `Service` and `ResourceType`. Otherwise, the configured Service or ResourceType does not take effect.
	ResourceTypes []*ListResourcesRequestResourceTypes `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
	// The ID of the Alibaba Cloud service.
	//
	// You can obtain the ID from the **Service code** column in [Services that work with Resource Group](~~94479~~).
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s ListResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListResourcesRequest) SetPageNumber(v int32) *ListResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesRequest) SetPageSize(v int32) *ListResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourcesRequest) SetRegion(v string) *ListResourcesRequest {
	s.Region = &v
	return s
}

func (s *ListResourcesRequest) SetResourceGroupId(v string) *ListResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListResourcesRequest) SetResourceId(v string) *ListResourcesRequest {
	s.ResourceId = &v
	return s
}

func (s *ListResourcesRequest) SetResourceType(v string) *ListResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListResourcesRequest) SetResourceTypes(v []*ListResourcesRequestResourceTypes) *ListResourcesRequest {
	s.ResourceTypes = v
	return s
}

func (s *ListResourcesRequest) SetService(v string) *ListResourcesRequest {
	s.Service = &v
	return s
}

type ListResourcesRequestResourceTypes struct {
	// The resource type.
	//
	// Valid values of N: 1 to 50.
	//
	// For more information about the supported resource types, see the **Resource type** column in [Services that work with Resource Group](~~94479~~).
	//
	// >  You must configure both `Service` and `ResourceType` in `ResourceTypes`. Otherwise, the two parameters do not take effect.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the Alibaba Cloud service.
	//
	// Valid values of N: 1 to 50.
	//
	// You can obtain the ID from the **Service code** column in [Services that work with Resource Group](~~94479~~).
	//
	// >  You must configure both `Service` and `ResourceType` in `ResourceTypes`. Otherwise, the two parameters do not take effect.
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s ListResourcesRequestResourceTypes) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesRequestResourceTypes) GoString() string {
	return s.String()
}

func (s *ListResourcesRequestResourceTypes) SetResourceType(v string) *ListResourcesRequestResourceTypes {
	s.ResourceType = &v
	return s
}

func (s *ListResourcesRequestResourceTypes) SetService(v string) *ListResourcesRequestResourceTypes {
	s.Service = &v
	return s
}

type ListResourcesResponseBody struct {
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the resources.
	Resources *ListResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBody) SetPageNumber(v int32) *ListResourcesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesResponseBody) SetPageSize(v int32) *ListResourcesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListResourcesResponseBody) SetRequestId(v string) *ListResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourcesResponseBody) SetResources(v *ListResourcesResponseBodyResources) *ListResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *ListResourcesResponseBody) SetTotalCount(v int32) *ListResourcesResponseBody {
	s.TotalCount = &v
	return s
}

type ListResourcesResponseBodyResources struct {
	Resource []*ListResourcesResponseBodyResourcesResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s ListResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyResources) SetResource(v []*ListResourcesResponseBodyResourcesResource) *ListResourcesResponseBodyResources {
	s.Resource = v
	return s
}

type ListResourcesResponseBodyResourcesResource struct {
	// The time when the resource was created. The time is displayed in UTC.
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the Alibaba Cloud service.
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s ListResourcesResponseBodyResourcesResource) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyResourcesResource) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyResourcesResource) SetCreateDate(v string) *ListResourcesResponseBodyResourcesResource {
	s.CreateDate = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesResource) SetRegionId(v string) *ListResourcesResponseBodyResourcesResource {
	s.RegionId = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesResource) SetResourceGroupId(v string) *ListResourcesResponseBodyResourcesResource {
	s.ResourceGroupId = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesResource) SetResourceId(v string) *ListResourcesResponseBodyResourcesResource {
	s.ResourceId = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesResource) SetResourceType(v string) *ListResourcesResponseBodyResourcesResource {
	s.ResourceType = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesResource) SetService(v string) *ListResourcesResponseBodyResourcesResource {
	s.Service = &v
	return s
}

type ListResourcesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListResourcesResponse) SetHeaders(v map[string]*string) *ListResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListResourcesResponse) SetStatusCode(v int32) *ListResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourcesResponse) SetBody(v *ListResourcesResponseBody) *ListResourcesResponse {
	s.Body = v
	return s
}

type ListRolesRequest struct {
	// The language that is used to return the descriptions of the RAM roles. Valid values:
	//
	// *   en: English
	// *   zh-CN: Chinese
	// *   ja: Japanese
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListRolesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRolesRequest) GoString() string {
	return s.String()
}

func (s *ListRolesRequest) SetLanguage(v string) *ListRolesRequest {
	s.Language = &v
	return s
}

func (s *ListRolesRequest) SetPageNumber(v int32) *ListRolesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRolesRequest) SetPageSize(v int32) *ListRolesRequest {
	s.PageSize = &v
	return s
}

type ListRolesResponseBody struct {
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the RAM roles.
	Roles *ListRolesResponseBodyRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Struct"`
	// The total number of RAM roles.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBody) SetPageNumber(v int32) *ListRolesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRolesResponseBody) SetPageSize(v int32) *ListRolesResponseBody {
	s.PageSize = &v
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

func (s *ListRolesResponseBody) SetTotalCount(v int32) *ListRolesResponseBody {
	s.TotalCount = &v
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
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The time when the RAM role was created.
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the RAM role.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the RAM role is a service linked role.
	IsServiceLinkedRole *bool `json:"IsServiceLinkedRole,omitempty" xml:"IsServiceLinkedRole,omitempty"`
	// The information of the most recent deletion task.
	LatestDeletionTask *ListRolesResponseBodyRolesRoleLatestDeletionTask `json:"LatestDeletionTask,omitempty" xml:"LatestDeletionTask,omitempty" type:"Struct"`
	// The maximum session duration of the RAM role.
	MaxSessionDuration *int64 `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty"`
	// The ID of the RAM role.
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// The name of the RAM role.
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The name of the RAM role after authorization.
	RolePrincipalName *string `json:"RolePrincipalName,omitempty" xml:"RolePrincipalName,omitempty"`
	// The time when the RAM role was updated.
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

func (s *ListRolesResponseBodyRolesRole) SetIsServiceLinkedRole(v bool) *ListRolesResponseBodyRolesRole {
	s.IsServiceLinkedRole = &v
	return s
}

func (s *ListRolesResponseBodyRolesRole) SetLatestDeletionTask(v *ListRolesResponseBodyRolesRoleLatestDeletionTask) *ListRolesResponseBodyRolesRole {
	s.LatestDeletionTask = v
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

func (s *ListRolesResponseBodyRolesRole) SetRolePrincipalName(v string) *ListRolesResponseBodyRolesRole {
	s.RolePrincipalName = &v
	return s
}

func (s *ListRolesResponseBodyRolesRole) SetUpdateDate(v string) *ListRolesResponseBodyRolesRole {
	s.UpdateDate = &v
	return s
}

type ListRolesResponseBodyRolesRoleLatestDeletionTask struct {
	// The time when the deletion task was created.
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The ID of the deletion task.
	DeletionTaskId *string `json:"DeletionTaskId,omitempty" xml:"DeletionTaskId,omitempty"`
}

func (s ListRolesResponseBodyRolesRoleLatestDeletionTask) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBodyRolesRoleLatestDeletionTask) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyRolesRoleLatestDeletionTask) SetCreateDate(v string) *ListRolesResponseBodyRolesRoleLatestDeletionTask {
	s.CreateDate = &v
	return s
}

func (s *ListRolesResponseBodyRolesRoleLatestDeletionTask) SetDeletionTaskId(v string) *ListRolesResponseBodyRolesRoleLatestDeletionTask {
	s.DeletionTaskId = &v
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

type ListTagKeysRequest struct {
	// The tag key for a fuzzy query.
	KeyFilter *string `json:"KeyFilter,omitempty" xml:"KeyFilter,omitempty"`
	// The maximum number of entries to return for a single request.
	//
	// Valid values: 1 to 100. Default value: 10.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to start the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource type.
	//
	// The value Account indicates the members of the resource directory.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetKeyFilter(v string) *ListTagKeysRequest {
	s.KeyFilter = &v
	return s
}

func (s *ListTagKeysRequest) SetMaxResults(v int32) *ListTagKeysRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagKeysRequest) SetNextToken(v string) *ListTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

type ListTagKeysResponseBody struct {
	// Indicates whether the next query is required.
	//
	// *   If the value of this parameter is empty (`"NextToken": ""`), all results are returned, and the next query is not required.
	// *   If the value of this parameter is not empty, the next query is required, and the value is the token used to start the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the tag keys.
	Tags []*ListTagKeysResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetNextToken(v string) *ListTagKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagKeysResponseBody) SetTags(v []*ListTagKeysResponseBodyTags) *ListTagKeysResponseBody {
	s.Tags = v
	return s
}

type ListTagKeysResponseBodyTags struct {
	// The tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s ListTagKeysResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBodyTags) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBodyTags) SetKey(v string) *ListTagKeysResponseBodyTags {
	s.Key = &v
	return s
}

type ListTagKeysResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponse) SetHeaders(v map[string]*string) *ListTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListTagKeysResponse) SetStatusCode(v int32) *ListTagKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to start the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource ID.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the objects whose tags you want to query. This parameter specifies a filter condition for the query. Valid values:
	//
	// *   ResourceGroup: resource group. This is the default value.
	// *   Account: member.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key and value.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetMaxResults(v int32) *ListTagResourcesRequest {
	s.MaxResults = &v
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

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// A tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// A tag value.
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
	// Indicates whether the next query is required.
	//
	// *   If the value of this parameter is empty (`"NextToken": ""`), all results are returned, and the next query is not required.
	// *   If the value of this parameter is not empty, the next query is required, and the value is the token used to start the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tags.
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
	// The ID of the resource group or member.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the object whose tags are queried. Valid values:
	//
	// *   resourcegroup: resource group
	// *   Account: member
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
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

type ListTagValuesRequest struct {
	// The maximum number of entries to return for a single request.
	//
	// Valid values: 1 to 100. Default value: 10.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to start the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource type.
	//
	// The value Account indicates the members of the resource directory.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key. This parameter specifies a filter condition for the query.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value for a fuzzy query.
	ValueFilter *string `json:"ValueFilter,omitempty" xml:"ValueFilter,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) SetMaxResults(v int32) *ListTagValuesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagValuesRequest) SetNextToken(v string) *ListTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceType(v string) *ListTagValuesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagValuesRequest) SetTagKey(v string) *ListTagValuesRequest {
	s.TagKey = &v
	return s
}

func (s *ListTagValuesRequest) SetValueFilter(v string) *ListTagValuesRequest {
	s.ValueFilter = &v
	return s
}

type ListTagValuesResponseBody struct {
	// Indicates whether the next query is required.
	//
	// *   If the value of this parameter is empty (`"NextToken": ""`), all results are returned, and the next query is not required.
	// *   If the value of this parameter is not empty, the next query is required, and the value is the token used to start the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the tag values.
	Tags []*ListTagValuesResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListTagValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBody) SetNextToken(v string) *ListTagValuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesResponseBody) SetRequestId(v string) *ListTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagValuesResponseBody) SetTags(v []*ListTagValuesResponseBodyTags) *ListTagValuesResponseBody {
	s.Tags = v
	return s
}

type ListTagValuesResponseBodyTags struct {
	// The tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagValuesResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponseBodyTags) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBodyTags) SetValue(v string) *ListTagValuesResponseBodyTags {
	s.Value = &v
	return s
}

type ListTagValuesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponse) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponse) SetHeaders(v map[string]*string) *ListTagValuesResponse {
	s.Headers = v
	return s
}

func (s *ListTagValuesResponse) SetStatusCode(v int32) *ListTagValuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagValuesResponse) SetBody(v *ListTagValuesResponseBody) *ListTagValuesResponse {
	s.Body = v
	return s
}

type ListTargetAttachmentsForControlPolicyRequest struct {
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the control policy.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s ListTargetAttachmentsForControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTargetAttachmentsForControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *ListTargetAttachmentsForControlPolicyRequest) SetPageNumber(v int32) *ListTargetAttachmentsForControlPolicyRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyRequest) SetPageSize(v int32) *ListTargetAttachmentsForControlPolicyRequest {
	s.PageSize = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyRequest) SetPolicyId(v string) *ListTargetAttachmentsForControlPolicyRequest {
	s.PolicyId = &v
	return s
}

type ListTargetAttachmentsForControlPolicyResponseBody struct {
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of objects to which the control policy is attached.
	TargetAttachments *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachments `json:"TargetAttachments,omitempty" xml:"TargetAttachments,omitempty" type:"Struct"`
	// The total number of objects to which the control policy is attached.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTargetAttachmentsForControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTargetAttachmentsForControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ListTargetAttachmentsForControlPolicyResponseBody) SetPageNumber(v int32) *ListTargetAttachmentsForControlPolicyResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponseBody) SetPageSize(v int32) *ListTargetAttachmentsForControlPolicyResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponseBody) SetRequestId(v string) *ListTargetAttachmentsForControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponseBody) SetTargetAttachments(v *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachments) *ListTargetAttachmentsForControlPolicyResponseBody {
	s.TargetAttachments = v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponseBody) SetTotalCount(v int32) *ListTargetAttachmentsForControlPolicyResponseBody {
	s.TotalCount = &v
	return s
}

type ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachments struct {
	TargetAttachment []*ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment `json:"TargetAttachment,omitempty" xml:"TargetAttachment,omitempty" type:"Repeated"`
}

func (s ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachments) String() string {
	return tea.Prettify(s)
}

func (s ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachments) GoString() string {
	return s.String()
}

func (s *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachments) SetTargetAttachment(v []*ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment) *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachments {
	s.TargetAttachment = v
	return s
}

type ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment struct {
	// The time when the control policy was attached to the object.
	AttachDate *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty"`
	// The ID of the object.
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the object.
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The type of the object. Valid values:
	//
	// *   Root: Root folder
	// *   Folder: child folder of the Root folder
	// *   Account: member account
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment) String() string {
	return tea.Prettify(s)
}

func (s ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment) GoString() string {
	return s.String()
}

func (s *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment) SetAttachDate(v string) *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment {
	s.AttachDate = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment) SetTargetId(v string) *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment {
	s.TargetId = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment) SetTargetName(v string) *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment {
	s.TargetName = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment) SetTargetType(v string) *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment {
	s.TargetType = &v
	return s
}

type ListTargetAttachmentsForControlPolicyResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTargetAttachmentsForControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTargetAttachmentsForControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTargetAttachmentsForControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *ListTargetAttachmentsForControlPolicyResponse) SetHeaders(v map[string]*string) *ListTargetAttachmentsForControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponse) SetStatusCode(v int32) *ListTargetAttachmentsForControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponse) SetBody(v *ListTargetAttachmentsForControlPolicyResponseBody) *ListTargetAttachmentsForControlPolicyResponse {
	s.Body = v
	return s
}

type ListTrustedServiceStatusRequest struct {
	// The ID of the enterprise management account or delegated administrator account.
	//
	// *   If you set this parameter to the ID of an enterprise management account, the trusted services that are enabled within the enterprise management account are queried. The default value of this parameter is the ID of an enterprise management account.
	// *   If you set this parameter to the ID of a delegated administrator account, the trusted services that are enabled within the delegated administrator account are queried.
	//
	// For more information about trusted services and delegated administrator accounts, see [Overview of trusted services](~~208133~~) and [Delegated administrator accounts](~~208117~~).
	AdminAccountId *string `json:"AdminAccountId,omitempty" xml:"AdminAccountId,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListTrustedServiceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTrustedServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *ListTrustedServiceStatusRequest) SetAdminAccountId(v string) *ListTrustedServiceStatusRequest {
	s.AdminAccountId = &v
	return s
}

func (s *ListTrustedServiceStatusRequest) SetPageNumber(v int32) *ListTrustedServiceStatusRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTrustedServiceStatusRequest) SetPageSize(v int32) *ListTrustedServiceStatusRequest {
	s.PageSize = &v
	return s
}

type ListTrustedServiceStatusResponseBody struct {
	// The trusted services that are enabled.
	EnabledServicePrincipals *ListTrustedServiceStatusResponseBodyEnabledServicePrincipals `json:"EnabledServicePrincipals,omitempty" xml:"EnabledServicePrincipals,omitempty" type:"Struct"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTrustedServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTrustedServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrustedServiceStatusResponseBody) SetEnabledServicePrincipals(v *ListTrustedServiceStatusResponseBodyEnabledServicePrincipals) *ListTrustedServiceStatusResponseBody {
	s.EnabledServicePrincipals = v
	return s
}

func (s *ListTrustedServiceStatusResponseBody) SetPageNumber(v int32) *ListTrustedServiceStatusResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTrustedServiceStatusResponseBody) SetPageSize(v int32) *ListTrustedServiceStatusResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTrustedServiceStatusResponseBody) SetRequestId(v string) *ListTrustedServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTrustedServiceStatusResponseBody) SetTotalCount(v int32) *ListTrustedServiceStatusResponseBody {
	s.TotalCount = &v
	return s
}

type ListTrustedServiceStatusResponseBodyEnabledServicePrincipals struct {
	EnabledServicePrincipal []*ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal `json:"EnabledServicePrincipal,omitempty" xml:"EnabledServicePrincipal,omitempty" type:"Repeated"`
}

func (s ListTrustedServiceStatusResponseBodyEnabledServicePrincipals) String() string {
	return tea.Prettify(s)
}

func (s ListTrustedServiceStatusResponseBodyEnabledServicePrincipals) GoString() string {
	return s.String()
}

func (s *ListTrustedServiceStatusResponseBodyEnabledServicePrincipals) SetEnabledServicePrincipal(v []*ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal) *ListTrustedServiceStatusResponseBodyEnabledServicePrincipals {
	s.EnabledServicePrincipal = v
	return s
}

type ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal struct {
	// The time when the trusted service was enabled.
	EnableTime *string `json:"EnableTime,omitempty" xml:"EnableTime,omitempty"`
	// The identification of the trusted service.
	ServicePrincipal *string `json:"ServicePrincipal,omitempty" xml:"ServicePrincipal,omitempty"`
}

func (s ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal) String() string {
	return tea.Prettify(s)
}

func (s ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal) GoString() string {
	return s.String()
}

func (s *ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal) SetEnableTime(v string) *ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal {
	s.EnableTime = &v
	return s
}

func (s *ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal) SetServicePrincipal(v string) *ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal {
	s.ServicePrincipal = &v
	return s
}

type ListTrustedServiceStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrustedServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrustedServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTrustedServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *ListTrustedServiceStatusResponse) SetHeaders(v map[string]*string) *ListTrustedServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *ListTrustedServiceStatusResponse) SetStatusCode(v int32) *ListTrustedServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrustedServiceStatusResponse) SetBody(v *ListTrustedServiceStatusResponseBody) *ListTrustedServiceStatusResponse {
	s.Body = v
	return s
}

type MoveAccountRequest struct {
	// The ID of the account you want to move.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the destination folder.
	DestinationFolderId *string `json:"DestinationFolderId,omitempty" xml:"DestinationFolderId,omitempty"`
}

func (s MoveAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveAccountRequest) GoString() string {
	return s.String()
}

func (s *MoveAccountRequest) SetAccountId(v string) *MoveAccountRequest {
	s.AccountId = &v
	return s
}

func (s *MoveAccountRequest) SetDestinationFolderId(v string) *MoveAccountRequest {
	s.DestinationFolderId = &v
	return s
}

type MoveAccountResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveAccountResponseBody) GoString() string {
	return s.String()
}

func (s *MoveAccountResponseBody) SetRequestId(v string) *MoveAccountResponseBody {
	s.RequestId = &v
	return s
}

type MoveAccountResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveAccountResponse) GoString() string {
	return s.String()
}

func (s *MoveAccountResponse) SetHeaders(v map[string]*string) *MoveAccountResponse {
	s.Headers = v
	return s
}

func (s *MoveAccountResponse) SetStatusCode(v int32) *MoveAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveAccountResponse) SetBody(v *MoveAccountResponseBody) *MoveAccountResponse {
	s.Body = v
	return s
}

type MoveResourcesRequest struct {
	// The ID of the resource group to which you want to move the resources.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resources that you want to move.
	//
	// >  You can move a maximum of 10 resources at a time. If you want to move more than 10 resources, move them in batches.
	Resources []*MoveResourcesRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s MoveResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveResourcesRequest) GoString() string {
	return s.String()
}

func (s *MoveResourcesRequest) SetResourceGroupId(v string) *MoveResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *MoveResourcesRequest) SetResources(v []*MoveResourcesRequestResources) *MoveResourcesRequest {
	s.Resources = v
	return s
}

type MoveResourcesRequestResources struct {
	// The region ID of the resource.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the Alibaba Cloud service to which the resource belongs.
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s MoveResourcesRequestResources) String() string {
	return tea.Prettify(s)
}

func (s MoveResourcesRequestResources) GoString() string {
	return s.String()
}

func (s *MoveResourcesRequestResources) SetRegionId(v string) *MoveResourcesRequestResources {
	s.RegionId = &v
	return s
}

func (s *MoveResourcesRequestResources) SetResourceId(v string) *MoveResourcesRequestResources {
	s.ResourceId = &v
	return s
}

func (s *MoveResourcesRequestResources) SetResourceType(v string) *MoveResourcesRequestResources {
	s.ResourceType = &v
	return s
}

func (s *MoveResourcesRequestResources) SetService(v string) *MoveResourcesRequestResources {
	s.Service = &v
	return s
}

type MoveResourcesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned results.
	Responses []*MoveResourcesResponseBodyResponses `json:"Responses,omitempty" xml:"Responses,omitempty" type:"Repeated"`
}

func (s MoveResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *MoveResourcesResponseBody) SetRequestId(v string) *MoveResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveResourcesResponseBody) SetResponses(v []*MoveResourcesResponseBodyResponses) *MoveResourcesResponseBody {
	s.Responses = v
	return s
}

type MoveResourcesResponseBodyResponses struct {
	// The error code returned.
	//
	// >  This parameter is returned if the resource failed to be moved.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// >  This parameter is returned if the resource failed to be moved.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The region ID of the resource.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the Alibaba Cloud service.
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The status of the move task. Valid values:
	//
	// *   SUCCESS
	// *   FAIL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s MoveResourcesResponseBodyResponses) String() string {
	return tea.Prettify(s)
}

func (s MoveResourcesResponseBodyResponses) GoString() string {
	return s.String()
}

func (s *MoveResourcesResponseBodyResponses) SetErrorCode(v string) *MoveResourcesResponseBodyResponses {
	s.ErrorCode = &v
	return s
}

func (s *MoveResourcesResponseBodyResponses) SetErrorMsg(v string) *MoveResourcesResponseBodyResponses {
	s.ErrorMsg = &v
	return s
}

func (s *MoveResourcesResponseBodyResponses) SetRegionId(v string) *MoveResourcesResponseBodyResponses {
	s.RegionId = &v
	return s
}

func (s *MoveResourcesResponseBodyResponses) SetRequestId(v string) *MoveResourcesResponseBodyResponses {
	s.RequestId = &v
	return s
}

func (s *MoveResourcesResponseBodyResponses) SetResourceId(v string) *MoveResourcesResponseBodyResponses {
	s.ResourceId = &v
	return s
}

func (s *MoveResourcesResponseBodyResponses) SetResourceType(v string) *MoveResourcesResponseBodyResponses {
	s.ResourceType = &v
	return s
}

func (s *MoveResourcesResponseBodyResponses) SetService(v string) *MoveResourcesResponseBodyResponses {
	s.Service = &v
	return s
}

func (s *MoveResourcesResponseBodyResponses) SetStatus(v string) *MoveResourcesResponseBodyResponses {
	s.Status = &v
	return s
}

type MoveResourcesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveResourcesResponse) GoString() string {
	return s.String()
}

func (s *MoveResourcesResponse) SetHeaders(v map[string]*string) *MoveResourcesResponse {
	s.Headers = v
	return s
}

func (s *MoveResourcesResponse) SetStatusCode(v int32) *MoveResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveResourcesResponse) SetBody(v *MoveResourcesResponseBody) *MoveResourcesResponse {
	s.Body = v
	return s
}

type PromoteResourceAccountRequest struct {
	// The ID of the resource account you want to upgrade.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The email address used to log on to the cloud account after the upgrade.
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
}

func (s PromoteResourceAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s PromoteResourceAccountRequest) GoString() string {
	return s.String()
}

func (s *PromoteResourceAccountRequest) SetAccountId(v string) *PromoteResourceAccountRequest {
	s.AccountId = &v
	return s
}

func (s *PromoteResourceAccountRequest) SetEmail(v string) *PromoteResourceAccountRequest {
	s.Email = &v
	return s
}

type PromoteResourceAccountResponseBody struct {
	// The information of the member account.
	Account *PromoteResourceAccountResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PromoteResourceAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PromoteResourceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *PromoteResourceAccountResponseBody) SetAccount(v *PromoteResourceAccountResponseBodyAccount) *PromoteResourceAccountResponseBody {
	s.Account = v
	return s
}

func (s *PromoteResourceAccountResponseBody) SetRequestId(v string) *PromoteResourceAccountResponseBody {
	s.RequestId = &v
	return s
}

type PromoteResourceAccountResponseBodyAccount struct {
	// The ID of the member account.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the member account.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The display name of the member account.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the folder.
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The way in which the member account joined the resource directory. Valid values:
	//
	// *   invited: The member account is invited to join the resource directory.
	// *   created: The member account is directly created in the resource directory.
	JoinMethod *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	// The time when the member account joined the resource directory.
	JoinTime *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// The time when the member account was modified.
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The account record ID.
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The ID of the resource directory.
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the member account. Valid values:
	//
	// *   CreateSuccess: The member account is created.
	// *   CreateVerifying: The creation of the member account is under confirmation.
	// *   CreateFailed: The member account failed to be created.
	// *   CreateExpired: The creation of the member account expired.
	// *   CreateCancelled: The creation of the member account is canceled.
	// *   PromoteVerifying: The upgrade of the member account is under confirmation.
	// *   PromoteFailed: The member account failed to be upgraded.
	// *   PromoteExpired: The upgrade of the member account expired.
	// *   PromoteCancelled: The upgrade of the member account is canceled.
	// *   PromoteSuccess: The member account is upgraded.
	// *   InviteSuccess: The owner of the member account accepted the invitation.
	// *   Removed: The member account is removed from the resource directory.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the member account. Valid values:
	//
	// *   CloudAccount
	// *   ResourceAccount
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PromoteResourceAccountResponseBodyAccount) String() string {
	return tea.Prettify(s)
}

func (s PromoteResourceAccountResponseBodyAccount) GoString() string {
	return s.String()
}

func (s *PromoteResourceAccountResponseBodyAccount) SetAccountId(v string) *PromoteResourceAccountResponseBodyAccount {
	s.AccountId = &v
	return s
}

func (s *PromoteResourceAccountResponseBodyAccount) SetAccountName(v string) *PromoteResourceAccountResponseBodyAccount {
	s.AccountName = &v
	return s
}

func (s *PromoteResourceAccountResponseBodyAccount) SetDisplayName(v string) *PromoteResourceAccountResponseBodyAccount {
	s.DisplayName = &v
	return s
}

func (s *PromoteResourceAccountResponseBodyAccount) SetFolderId(v string) *PromoteResourceAccountResponseBodyAccount {
	s.FolderId = &v
	return s
}

func (s *PromoteResourceAccountResponseBodyAccount) SetJoinMethod(v string) *PromoteResourceAccountResponseBodyAccount {
	s.JoinMethod = &v
	return s
}

func (s *PromoteResourceAccountResponseBodyAccount) SetJoinTime(v string) *PromoteResourceAccountResponseBodyAccount {
	s.JoinTime = &v
	return s
}

func (s *PromoteResourceAccountResponseBodyAccount) SetModifyTime(v string) *PromoteResourceAccountResponseBodyAccount {
	s.ModifyTime = &v
	return s
}

func (s *PromoteResourceAccountResponseBodyAccount) SetRecordId(v string) *PromoteResourceAccountResponseBodyAccount {
	s.RecordId = &v
	return s
}

func (s *PromoteResourceAccountResponseBodyAccount) SetResourceDirectoryId(v string) *PromoteResourceAccountResponseBodyAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *PromoteResourceAccountResponseBodyAccount) SetStatus(v string) *PromoteResourceAccountResponseBodyAccount {
	s.Status = &v
	return s
}

func (s *PromoteResourceAccountResponseBodyAccount) SetType(v string) *PromoteResourceAccountResponseBodyAccount {
	s.Type = &v
	return s
}

type PromoteResourceAccountResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PromoteResourceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PromoteResourceAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s PromoteResourceAccountResponse) GoString() string {
	return s.String()
}

func (s *PromoteResourceAccountResponse) SetHeaders(v map[string]*string) *PromoteResourceAccountResponse {
	s.Headers = v
	return s
}

func (s *PromoteResourceAccountResponse) SetStatusCode(v int32) *PromoteResourceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *PromoteResourceAccountResponse) SetBody(v *PromoteResourceAccountResponseBody) *PromoteResourceAccountResponse {
	s.Body = v
	return s
}

type RegisterDelegatedAdministratorRequest struct {
	// The ID of the member in the resource directory.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The identifier of the trusted service.
	//
	// For more information, see the `Trusted service identifier` column in [Supported trusted services](~~208133~~).
	ServicePrincipal *string `json:"ServicePrincipal,omitempty" xml:"ServicePrincipal,omitempty"`
}

func (s RegisterDelegatedAdministratorRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterDelegatedAdministratorRequest) GoString() string {
	return s.String()
}

func (s *RegisterDelegatedAdministratorRequest) SetAccountId(v string) *RegisterDelegatedAdministratorRequest {
	s.AccountId = &v
	return s
}

func (s *RegisterDelegatedAdministratorRequest) SetServicePrincipal(v string) *RegisterDelegatedAdministratorRequest {
	s.ServicePrincipal = &v
	return s
}

type RegisterDelegatedAdministratorResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterDelegatedAdministratorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterDelegatedAdministratorResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterDelegatedAdministratorResponseBody) SetRequestId(v string) *RegisterDelegatedAdministratorResponseBody {
	s.RequestId = &v
	return s
}

type RegisterDelegatedAdministratorResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterDelegatedAdministratorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterDelegatedAdministratorResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterDelegatedAdministratorResponse) GoString() string {
	return s.String()
}

func (s *RegisterDelegatedAdministratorResponse) SetHeaders(v map[string]*string) *RegisterDelegatedAdministratorResponse {
	s.Headers = v
	return s
}

func (s *RegisterDelegatedAdministratorResponse) SetStatusCode(v int32) *RegisterDelegatedAdministratorResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterDelegatedAdministratorResponse) SetBody(v *RegisterDelegatedAdministratorResponseBody) *RegisterDelegatedAdministratorResponse {
	s.Body = v
	return s
}

type RemoveCloudAccountRequest struct {
	// The ID of the member.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s RemoveCloudAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveCloudAccountRequest) GoString() string {
	return s.String()
}

func (s *RemoveCloudAccountRequest) SetAccountId(v string) *RemoveCloudAccountRequest {
	s.AccountId = &v
	return s
}

type RemoveCloudAccountResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveCloudAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveCloudAccountResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveCloudAccountResponseBody) SetRequestId(v string) *RemoveCloudAccountResponseBody {
	s.RequestId = &v
	return s
}

type RemoveCloudAccountResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveCloudAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveCloudAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveCloudAccountResponse) GoString() string {
	return s.String()
}

func (s *RemoveCloudAccountResponse) SetHeaders(v map[string]*string) *RemoveCloudAccountResponse {
	s.Headers = v
	return s
}

func (s *RemoveCloudAccountResponse) SetStatusCode(v int32) *RemoveCloudAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveCloudAccountResponse) SetBody(v *RemoveCloudAccountResponseBody) *RemoveCloudAccountResponse {
	s.Body = v
	return s
}

type ResendCreateCloudAccountEmailRequest struct {
	// The account record ID.
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s ResendCreateCloudAccountEmailRequest) String() string {
	return tea.Prettify(s)
}

func (s ResendCreateCloudAccountEmailRequest) GoString() string {
	return s.String()
}

func (s *ResendCreateCloudAccountEmailRequest) SetRecordId(v string) *ResendCreateCloudAccountEmailRequest {
	s.RecordId = &v
	return s
}

type ResendCreateCloudAccountEmailResponseBody struct {
	// The information of the member account.
	Account *ResendCreateCloudAccountEmailResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResendCreateCloudAccountEmailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResendCreateCloudAccountEmailResponseBody) GoString() string {
	return s.String()
}

func (s *ResendCreateCloudAccountEmailResponseBody) SetAccount(v *ResendCreateCloudAccountEmailResponseBodyAccount) *ResendCreateCloudAccountEmailResponseBody {
	s.Account = v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseBody) SetRequestId(v string) *ResendCreateCloudAccountEmailResponseBody {
	s.RequestId = &v
	return s
}

type ResendCreateCloudAccountEmailResponseBodyAccount struct {
	// The ID of the account.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the account.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The display name of the member account.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the folder.
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The way in which the member account joined the resource directory. Valid values:
	//
	// *   invited: The member account is invited to join the resource directory.
	// *   created: The member account is directly created in the resource directory.
	JoinMethod *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	// The time when the member account joined the resource directory.
	JoinTime *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// The time when the member account was modified.
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The account record ID.
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The ID of the resource directory.
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the member account. Valid values:
	//
	// *   CreateSuccess: The member account is created.
	// *   CreateVerifying: The creation of the member account is under confirmation.
	// *   CreateFailed: The member account failed to be created.
	// *   CreateExpired: The creation of the member account expired.
	// *   CreateCancelled: The creation of the member account is canceled.
	// *   PromoteVerifying: The upgrade of the member account is under confirmation.
	// *   PromoteFailed: The member account failed to be upgraded.
	// *   PromoteExpired: The upgrade of the member account expired.
	// *   PromoteCancelled: The upgrade of the member account is canceled.
	// *   PromoteSuccess: The member account is upgraded.
	// *   InviteSuccess: The owner of the member account accepted the invitation.
	// *   Removed: The member account is removed from the resource directory.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the member account. The value CloudAccount indicates that the member account is a cloud account.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ResendCreateCloudAccountEmailResponseBodyAccount) String() string {
	return tea.Prettify(s)
}

func (s ResendCreateCloudAccountEmailResponseBodyAccount) GoString() string {
	return s.String()
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) SetAccountId(v string) *ResendCreateCloudAccountEmailResponseBodyAccount {
	s.AccountId = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) SetAccountName(v string) *ResendCreateCloudAccountEmailResponseBodyAccount {
	s.AccountName = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) SetDisplayName(v string) *ResendCreateCloudAccountEmailResponseBodyAccount {
	s.DisplayName = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) SetFolderId(v string) *ResendCreateCloudAccountEmailResponseBodyAccount {
	s.FolderId = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) SetJoinMethod(v string) *ResendCreateCloudAccountEmailResponseBodyAccount {
	s.JoinMethod = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) SetJoinTime(v string) *ResendCreateCloudAccountEmailResponseBodyAccount {
	s.JoinTime = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) SetModifyTime(v string) *ResendCreateCloudAccountEmailResponseBodyAccount {
	s.ModifyTime = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) SetRecordId(v string) *ResendCreateCloudAccountEmailResponseBodyAccount {
	s.RecordId = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) SetResourceDirectoryId(v string) *ResendCreateCloudAccountEmailResponseBodyAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) SetStatus(v string) *ResendCreateCloudAccountEmailResponseBodyAccount {
	s.Status = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) SetType(v string) *ResendCreateCloudAccountEmailResponseBodyAccount {
	s.Type = &v
	return s
}

type ResendCreateCloudAccountEmailResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResendCreateCloudAccountEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResendCreateCloudAccountEmailResponse) String() string {
	return tea.Prettify(s)
}

func (s ResendCreateCloudAccountEmailResponse) GoString() string {
	return s.String()
}

func (s *ResendCreateCloudAccountEmailResponse) SetHeaders(v map[string]*string) *ResendCreateCloudAccountEmailResponse {
	s.Headers = v
	return s
}

func (s *ResendCreateCloudAccountEmailResponse) SetStatusCode(v int32) *ResendCreateCloudAccountEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponse) SetBody(v *ResendCreateCloudAccountEmailResponseBody) *ResendCreateCloudAccountEmailResponse {
	s.Body = v
	return s
}

type ResendPromoteResourceAccountEmailRequest struct {
	// The account record ID.
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s ResendPromoteResourceAccountEmailRequest) String() string {
	return tea.Prettify(s)
}

func (s ResendPromoteResourceAccountEmailRequest) GoString() string {
	return s.String()
}

func (s *ResendPromoteResourceAccountEmailRequest) SetRecordId(v string) *ResendPromoteResourceAccountEmailRequest {
	s.RecordId = &v
	return s
}

type ResendPromoteResourceAccountEmailResponseBody struct {
	// The information of the member account.
	Account *ResendPromoteResourceAccountEmailResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResendPromoteResourceAccountEmailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResendPromoteResourceAccountEmailResponseBody) GoString() string {
	return s.String()
}

func (s *ResendPromoteResourceAccountEmailResponseBody) SetAccount(v *ResendPromoteResourceAccountEmailResponseBodyAccount) *ResendPromoteResourceAccountEmailResponseBody {
	s.Account = v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseBody) SetRequestId(v string) *ResendPromoteResourceAccountEmailResponseBody {
	s.RequestId = &v
	return s
}

type ResendPromoteResourceAccountEmailResponseBodyAccount struct {
	// The ID of the account.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the account.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The display name of the member account.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the folder.
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The way in which the member account joined the resource directory. Valid values:
	//
	// *   invited: The member account is invited to join the resource directory.
	// *   created: The member account is directly created in the resource directory.
	JoinMethod *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	// The time when the member account joined the resource directory.
	JoinTime *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// The time when the member account was modified.
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The account record ID.
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The ID of the resource directory.
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the member account. Valid values:
	//
	// *   CreateSuccess: The member account is created.
	// *   CreateVerifying: The creation of the member account is under confirmation.
	// *   CreateFailed: The member account failed to be created.
	// *   CreateExpired: The creation of the member account expired.
	// *   CreateCancelled: The creation of the member account is canceled.
	// *   PromoteVerifying: The upgrade of the member account is under confirmation.
	// *   PromoteFailed: The member account failed to be upgraded.
	// *   PromoteExpired: The upgrade of the member account expired.
	// *   PromoteCancelled: The upgrade of the member account is canceled.
	// *   PromoteSuccess: The member account is upgraded.
	// *   InviteSuccess: The owner of the member account accepted the invitation.
	// *   Removed: The member account is removed from the resource directory.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the member account. Valid values:
	//
	// *   CloudAccount: cloud account
	// *   ResourceAccount: resource account
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ResendPromoteResourceAccountEmailResponseBodyAccount) String() string {
	return tea.Prettify(s)
}

func (s ResendPromoteResourceAccountEmailResponseBodyAccount) GoString() string {
	return s.String()
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) SetAccountId(v string) *ResendPromoteResourceAccountEmailResponseBodyAccount {
	s.AccountId = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) SetAccountName(v string) *ResendPromoteResourceAccountEmailResponseBodyAccount {
	s.AccountName = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) SetDisplayName(v string) *ResendPromoteResourceAccountEmailResponseBodyAccount {
	s.DisplayName = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) SetFolderId(v string) *ResendPromoteResourceAccountEmailResponseBodyAccount {
	s.FolderId = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) SetJoinMethod(v string) *ResendPromoteResourceAccountEmailResponseBodyAccount {
	s.JoinMethod = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) SetJoinTime(v string) *ResendPromoteResourceAccountEmailResponseBodyAccount {
	s.JoinTime = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) SetModifyTime(v string) *ResendPromoteResourceAccountEmailResponseBodyAccount {
	s.ModifyTime = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) SetRecordId(v string) *ResendPromoteResourceAccountEmailResponseBodyAccount {
	s.RecordId = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) SetResourceDirectoryId(v string) *ResendPromoteResourceAccountEmailResponseBodyAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) SetStatus(v string) *ResendPromoteResourceAccountEmailResponseBodyAccount {
	s.Status = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) SetType(v string) *ResendPromoteResourceAccountEmailResponseBodyAccount {
	s.Type = &v
	return s
}

type ResendPromoteResourceAccountEmailResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResendPromoteResourceAccountEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResendPromoteResourceAccountEmailResponse) String() string {
	return tea.Prettify(s)
}

func (s ResendPromoteResourceAccountEmailResponse) GoString() string {
	return s.String()
}

func (s *ResendPromoteResourceAccountEmailResponse) SetHeaders(v map[string]*string) *ResendPromoteResourceAccountEmailResponse {
	s.Headers = v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponse) SetStatusCode(v int32) *ResendPromoteResourceAccountEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponse) SetBody(v *ResendPromoteResourceAccountEmailResponseBody) *ResendPromoteResourceAccountEmailResponse {
	s.Body = v
	return s
}

type RetryChangeAccountEmailRequest struct {
	// The Alibaba Cloud account ID of the member.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s RetryChangeAccountEmailRequest) String() string {
	return tea.Prettify(s)
}

func (s RetryChangeAccountEmailRequest) GoString() string {
	return s.String()
}

func (s *RetryChangeAccountEmailRequest) SetAccountId(v string) *RetryChangeAccountEmailRequest {
	s.AccountId = &v
	return s
}

type RetryChangeAccountEmailResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RetryChangeAccountEmailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetryChangeAccountEmailResponseBody) GoString() string {
	return s.String()
}

func (s *RetryChangeAccountEmailResponseBody) SetRequestId(v string) *RetryChangeAccountEmailResponseBody {
	s.RequestId = &v
	return s
}

type RetryChangeAccountEmailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryChangeAccountEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryChangeAccountEmailResponse) String() string {
	return tea.Prettify(s)
}

func (s RetryChangeAccountEmailResponse) GoString() string {
	return s.String()
}

func (s *RetryChangeAccountEmailResponse) SetHeaders(v map[string]*string) *RetryChangeAccountEmailResponse {
	s.Headers = v
	return s
}

func (s *RetryChangeAccountEmailResponse) SetStatusCode(v int32) *RetryChangeAccountEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryChangeAccountEmailResponse) SetBody(v *RetryChangeAccountEmailResponseBody) *RetryChangeAccountEmailResponse {
	s.Body = v
	return s
}

type SendVerificationCodeForBindSecureMobilePhoneRequest struct {
	// The ID of the resource account.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The mobile phone number that you want to bind to the resource account.
	//
	// Specify the mobile phone number in the \<Country code>-\<Mobile phone number> format.
	//
	// >  Mobile phone numbers in the `86-<Mobile phone number>` format in the Chinese mainland are not supported.
	SecureMobilePhone *string `json:"SecureMobilePhone,omitempty" xml:"SecureMobilePhone,omitempty"`
}

func (s SendVerificationCodeForBindSecureMobilePhoneRequest) String() string {
	return tea.Prettify(s)
}

func (s SendVerificationCodeForBindSecureMobilePhoneRequest) GoString() string {
	return s.String()
}

func (s *SendVerificationCodeForBindSecureMobilePhoneRequest) SetAccountId(v string) *SendVerificationCodeForBindSecureMobilePhoneRequest {
	s.AccountId = &v
	return s
}

func (s *SendVerificationCodeForBindSecureMobilePhoneRequest) SetSecureMobilePhone(v string) *SendVerificationCodeForBindSecureMobilePhoneRequest {
	s.SecureMobilePhone = &v
	return s
}

type SendVerificationCodeForBindSecureMobilePhoneResponseBody struct {
	// The time when the verification code expires.
	ExpirationDate *string `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendVerificationCodeForBindSecureMobilePhoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendVerificationCodeForBindSecureMobilePhoneResponseBody) GoString() string {
	return s.String()
}

func (s *SendVerificationCodeForBindSecureMobilePhoneResponseBody) SetExpirationDate(v string) *SendVerificationCodeForBindSecureMobilePhoneResponseBody {
	s.ExpirationDate = &v
	return s
}

func (s *SendVerificationCodeForBindSecureMobilePhoneResponseBody) SetRequestId(v string) *SendVerificationCodeForBindSecureMobilePhoneResponseBody {
	s.RequestId = &v
	return s
}

type SendVerificationCodeForBindSecureMobilePhoneResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendVerificationCodeForBindSecureMobilePhoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendVerificationCodeForBindSecureMobilePhoneResponse) String() string {
	return tea.Prettify(s)
}

func (s SendVerificationCodeForBindSecureMobilePhoneResponse) GoString() string {
	return s.String()
}

func (s *SendVerificationCodeForBindSecureMobilePhoneResponse) SetHeaders(v map[string]*string) *SendVerificationCodeForBindSecureMobilePhoneResponse {
	s.Headers = v
	return s
}

func (s *SendVerificationCodeForBindSecureMobilePhoneResponse) SetStatusCode(v int32) *SendVerificationCodeForBindSecureMobilePhoneResponse {
	s.StatusCode = &v
	return s
}

func (s *SendVerificationCodeForBindSecureMobilePhoneResponse) SetBody(v *SendVerificationCodeForBindSecureMobilePhoneResponseBody) *SendVerificationCodeForBindSecureMobilePhoneResponse {
	s.Body = v
	return s
}

type SendVerificationCodeForEnableRDRequest struct {
	// The mobile phone number that is bound to the newly created account. If you leave this parameter empty, the mobile phone number that is bound to the current account is used.
	//
	// Specify the mobile phone number in the `<Country code>-<Mobile phone number>` format.
	//
	// >  Mobile phone numbers in the `86-<Mobile phone number>` format in the Chinese mainland are not supported.
	SecureMobilePhone *string `json:"SecureMobilePhone,omitempty" xml:"SecureMobilePhone,omitempty"`
}

func (s SendVerificationCodeForEnableRDRequest) String() string {
	return tea.Prettify(s)
}

func (s SendVerificationCodeForEnableRDRequest) GoString() string {
	return s.String()
}

func (s *SendVerificationCodeForEnableRDRequest) SetSecureMobilePhone(v string) *SendVerificationCodeForEnableRDRequest {
	s.SecureMobilePhone = &v
	return s
}

type SendVerificationCodeForEnableRDResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendVerificationCodeForEnableRDResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendVerificationCodeForEnableRDResponseBody) GoString() string {
	return s.String()
}

func (s *SendVerificationCodeForEnableRDResponseBody) SetRequestId(v string) *SendVerificationCodeForEnableRDResponseBody {
	s.RequestId = &v
	return s
}

type SendVerificationCodeForEnableRDResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendVerificationCodeForEnableRDResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendVerificationCodeForEnableRDResponse) String() string {
	return tea.Prettify(s)
}

func (s SendVerificationCodeForEnableRDResponse) GoString() string {
	return s.String()
}

func (s *SendVerificationCodeForEnableRDResponse) SetHeaders(v map[string]*string) *SendVerificationCodeForEnableRDResponse {
	s.Headers = v
	return s
}

func (s *SendVerificationCodeForEnableRDResponse) SetStatusCode(v int32) *SendVerificationCodeForEnableRDResponse {
	s.StatusCode = &v
	return s
}

func (s *SendVerificationCodeForEnableRDResponse) SetBody(v *SendVerificationCodeForEnableRDResponseBody) *SendVerificationCodeForEnableRDResponse {
	s.Body = v
	return s
}

type SetDefaultPolicyVersionRequest struct {
	// The name of the policy.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The ID of the policy version.
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
	// The ID of the request.
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

type SetMemberDeletionPermissionRequest struct {
	// Specifies whether to enable the member deletion feature. Valid values:
	//
	// *   Enabled: enables the member deletion feature
	// *   Disabled: disables the member deletion feature
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetMemberDeletionPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s SetMemberDeletionPermissionRequest) GoString() string {
	return s.String()
}

func (s *SetMemberDeletionPermissionRequest) SetStatus(v string) *SetMemberDeletionPermissionRequest {
	s.Status = &v
	return s
}

type SetMemberDeletionPermissionResponseBody struct {
	// The ID of the management account of the resource directory.
	ManagementAccountId *string `json:"ManagementAccountId,omitempty" xml:"ManagementAccountId,omitempty"`
	// The status of the member deletion feature. Valid values:
	//
	// *   Enabled: The feature is enabled.
	// *   Disabled: The feature is disabled.
	MemberDeletionStatus *string `json:"MemberDeletionStatus,omitempty" xml:"MemberDeletionStatus,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource directory.
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
}

func (s SetMemberDeletionPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetMemberDeletionPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *SetMemberDeletionPermissionResponseBody) SetManagementAccountId(v string) *SetMemberDeletionPermissionResponseBody {
	s.ManagementAccountId = &v
	return s
}

func (s *SetMemberDeletionPermissionResponseBody) SetMemberDeletionStatus(v string) *SetMemberDeletionPermissionResponseBody {
	s.MemberDeletionStatus = &v
	return s
}

func (s *SetMemberDeletionPermissionResponseBody) SetRequestId(v string) *SetMemberDeletionPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetMemberDeletionPermissionResponseBody) SetResourceDirectoryId(v string) *SetMemberDeletionPermissionResponseBody {
	s.ResourceDirectoryId = &v
	return s
}

type SetMemberDeletionPermissionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetMemberDeletionPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetMemberDeletionPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s SetMemberDeletionPermissionResponse) GoString() string {
	return s.String()
}

func (s *SetMemberDeletionPermissionResponse) SetHeaders(v map[string]*string) *SetMemberDeletionPermissionResponse {
	s.Headers = v
	return s
}

func (s *SetMemberDeletionPermissionResponse) SetStatusCode(v int32) *SetMemberDeletionPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetMemberDeletionPermissionResponse) SetBody(v *SetMemberDeletionPermissionResponseBody) *SetMemberDeletionPermissionResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The ID of a resource group or member.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the objects to which you want to add tags. Valid values:
	//
	// *   ResourceGroup : resource group. This is the default value.
	// *   Account: member.
	//
	// >  This parameter is required if you add tags to members in a resource directory.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
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
	// A tag key.
	//
	// A tag key can be a maximum of 128 characters in length. It cannot contain `http://` or `https://` and cannot start with `acs:` or `aliyun`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// A tag value.
	//
	// A tag value can be a maximum of 128 characters in length. It cannot contain `http://` or `https://` and cannot start with `acs:`.
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
	// The ID of the request.
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
	// Specifies whether to remove all tags from the specified resource groups or members. Valid values:
	//
	// *   false (default value)
	// *   true
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The ID of a resource group or member.
	//
	// You can specify a maximum of 50 IDs.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the objects from which you want to remove tags. Valid values:
	//
	// *   ResourceGroup: resource group. This is the default value.
	// *   Account: member.
	//
	// >  This parameter is required if you remove tags from members in a resource directory.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// A tag key.
	//
	// You can specify a maximum of 20 tag keys.
	//
	// >  If you set the `All` parameter to `true`, you do not need to configure this parameter.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
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
	// The ID of the request.
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

type UpdateAccountRequest struct {
	// The ID of the Alibaba Cloud account that corresponds to the member.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The new type of the member. Valid values:
	//
	// *   ResourceAccount: resource account
	// *   CloudAccount: cloud account
	//
	// >  You can configure either the `NewDisplayName` or `NewAccountType` parameter.
	NewAccountType *string `json:"NewAccountType,omitempty" xml:"NewAccountType,omitempty"`
	// The new display name of the member.
	//
	// >  You can configure either the `NewDisplayName` or `NewAccountType` parameter.
	NewDisplayName *string `json:"NewDisplayName,omitempty" xml:"NewDisplayName,omitempty"`
}

func (s UpdateAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccountRequest) GoString() string {
	return s.String()
}

func (s *UpdateAccountRequest) SetAccountId(v string) *UpdateAccountRequest {
	s.AccountId = &v
	return s
}

func (s *UpdateAccountRequest) SetNewAccountType(v string) *UpdateAccountRequest {
	s.NewAccountType = &v
	return s
}

func (s *UpdateAccountRequest) SetNewDisplayName(v string) *UpdateAccountRequest {
	s.NewDisplayName = &v
	return s
}

type UpdateAccountResponseBody struct {
	// The information of the member.
	Account *UpdateAccountResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccountResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAccountResponseBody) SetAccount(v *UpdateAccountResponseBodyAccount) *UpdateAccountResponseBody {
	s.Account = v
	return s
}

func (s *UpdateAccountResponseBody) SetRequestId(v string) *UpdateAccountResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAccountResponseBodyAccount struct {
	// The ID of the Alibaba Cloud account that corresponds to the member.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the Alibaba Cloud account that corresponds to the member.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The display name of the member.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the folder.
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The way in which the member joins the resource directory. Valid values:
	//
	// *   invited: The member is invited to join the resource directory.
	// *   created: The member is directly created in the resource directory.
	JoinMethod *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	// The time when the member joined the resource directory. The time is displayed in UTC.
	JoinTime *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// The time when the member was modified. The time is displayed in UTC.
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The ID of the resource directory.
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the member. Valid values:
	//
	// *   CreateSuccess: The member is created.
	// *   InviteSuccess: The member accepts the invitation.
	// *   Removed: The member is removed.
	// *   SwitchSuccess: The type of the member is switched.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the member. Valid values:
	//
	// *   CloudAccount: cloud account
	// *   ResourceAccount: resource account
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateAccountResponseBodyAccount) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccountResponseBodyAccount) GoString() string {
	return s.String()
}

func (s *UpdateAccountResponseBodyAccount) SetAccountId(v string) *UpdateAccountResponseBodyAccount {
	s.AccountId = &v
	return s
}

func (s *UpdateAccountResponseBodyAccount) SetAccountName(v string) *UpdateAccountResponseBodyAccount {
	s.AccountName = &v
	return s
}

func (s *UpdateAccountResponseBodyAccount) SetDisplayName(v string) *UpdateAccountResponseBodyAccount {
	s.DisplayName = &v
	return s
}

func (s *UpdateAccountResponseBodyAccount) SetFolderId(v string) *UpdateAccountResponseBodyAccount {
	s.FolderId = &v
	return s
}

func (s *UpdateAccountResponseBodyAccount) SetJoinMethod(v string) *UpdateAccountResponseBodyAccount {
	s.JoinMethod = &v
	return s
}

func (s *UpdateAccountResponseBodyAccount) SetJoinTime(v string) *UpdateAccountResponseBodyAccount {
	s.JoinTime = &v
	return s
}

func (s *UpdateAccountResponseBodyAccount) SetModifyTime(v string) *UpdateAccountResponseBodyAccount {
	s.ModifyTime = &v
	return s
}

func (s *UpdateAccountResponseBodyAccount) SetResourceDirectoryId(v string) *UpdateAccountResponseBodyAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *UpdateAccountResponseBodyAccount) SetStatus(v string) *UpdateAccountResponseBodyAccount {
	s.Status = &v
	return s
}

func (s *UpdateAccountResponseBodyAccount) SetType(v string) *UpdateAccountResponseBodyAccount {
	s.Type = &v
	return s
}

type UpdateAccountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccountResponse) GoString() string {
	return s.String()
}

func (s *UpdateAccountResponse) SetHeaders(v map[string]*string) *UpdateAccountResponse {
	s.Headers = v
	return s
}

func (s *UpdateAccountResponse) SetStatusCode(v int32) *UpdateAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAccountResponse) SetBody(v *UpdateAccountResponseBody) *UpdateAccountResponse {
	s.Body = v
	return s
}

type UpdateAssociatedTransferSettingRequest struct {
	EnableExistingResourcesTransfer *string `json:"EnableExistingResourcesTransfer,omitempty" xml:"EnableExistingResourcesTransfer,omitempty"`
	// The settings of the transfer rules.
	RuleSettings []*UpdateAssociatedTransferSettingRequestRuleSettings `json:"RuleSettings,omitempty" xml:"RuleSettings,omitempty" type:"Repeated"`
}

func (s UpdateAssociatedTransferSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAssociatedTransferSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateAssociatedTransferSettingRequest) SetEnableExistingResourcesTransfer(v string) *UpdateAssociatedTransferSettingRequest {
	s.EnableExistingResourcesTransfer = &v
	return s
}

func (s *UpdateAssociatedTransferSettingRequest) SetRuleSettings(v []*UpdateAssociatedTransferSettingRequestRuleSettings) *UpdateAssociatedTransferSettingRequest {
	s.RuleSettings = v
	return s
}

type UpdateAssociatedTransferSettingRequestRuleSettings struct {
	// The type of the associated resource.
	//
	// You can obtain the resource type from the **Resource type** column in [Services that work with Resource Group](~~94479~~).
	AssociatedResourceType *string `json:"AssociatedResourceType,omitempty" xml:"AssociatedResourceType,omitempty"`
	// The service code of the associated resource.
	//
	// You can obtain the service code from the **Service code** column in [Services that work with Resource Group](~~94479~~).
	AssociatedService *string `json:"AssociatedService,omitempty" xml:"AssociatedService,omitempty"`
	// The type of the primary resource.
	//
	// You can obtain the resource type from the **Resource type** column in [Services that work with Resource Group](~~94479~~).
	MasterResourceType *string `json:"MasterResourceType,omitempty" xml:"MasterResourceType,omitempty"`
	// The service code of the primary resource.
	//
	// You can obtain the service code from the **Service code** column in [Services that work with Resource Group](~~94479~~).
	MasterService *string `json:"MasterService,omitempty" xml:"MasterService,omitempty"`
	// The status of the Transfer Associated Resources feature. Valid values:
	//
	// - Enable: enabled
	// - Disable: disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateAssociatedTransferSettingRequestRuleSettings) String() string {
	return tea.Prettify(s)
}

func (s UpdateAssociatedTransferSettingRequestRuleSettings) GoString() string {
	return s.String()
}

func (s *UpdateAssociatedTransferSettingRequestRuleSettings) SetAssociatedResourceType(v string) *UpdateAssociatedTransferSettingRequestRuleSettings {
	s.AssociatedResourceType = &v
	return s
}

func (s *UpdateAssociatedTransferSettingRequestRuleSettings) SetAssociatedService(v string) *UpdateAssociatedTransferSettingRequestRuleSettings {
	s.AssociatedService = &v
	return s
}

func (s *UpdateAssociatedTransferSettingRequestRuleSettings) SetMasterResourceType(v string) *UpdateAssociatedTransferSettingRequestRuleSettings {
	s.MasterResourceType = &v
	return s
}

func (s *UpdateAssociatedTransferSettingRequestRuleSettings) SetMasterService(v string) *UpdateAssociatedTransferSettingRequestRuleSettings {
	s.MasterService = &v
	return s
}

func (s *UpdateAssociatedTransferSettingRequestRuleSettings) SetStatus(v string) *UpdateAssociatedTransferSettingRequestRuleSettings {
	s.Status = &v
	return s
}

type UpdateAssociatedTransferSettingResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAssociatedTransferSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAssociatedTransferSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAssociatedTransferSettingResponseBody) SetRequestId(v string) *UpdateAssociatedTransferSettingResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAssociatedTransferSettingResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAssociatedTransferSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAssociatedTransferSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAssociatedTransferSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateAssociatedTransferSettingResponse) SetHeaders(v map[string]*string) *UpdateAssociatedTransferSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateAssociatedTransferSettingResponse) SetStatusCode(v int32) *UpdateAssociatedTransferSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAssociatedTransferSettingResponse) SetBody(v *UpdateAssociatedTransferSettingResponseBody) *UpdateAssociatedTransferSettingResponse {
	s.Body = v
	return s
}

type UpdateControlPolicyRequest struct {
	// The new description of the access control policy.
	//
	// The description must be 1 to 1,024 characters in length. The description can contain letters, digits, underscores (\_), and hyphens (-) and must start with a letter.
	NewDescription *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	// The new document of the access control policy.
	//
	// The document can be a maximum of 4,096 characters in length.
	//
	// For more information about the languages of access control policies, see [Languages of access control policies](~~179096~~).
	//
	// For more information about the examples of access control policies, see [Examples of custom access control policies](~~181474~~).
	NewPolicyDocument *string `json:"NewPolicyDocument,omitempty" xml:"NewPolicyDocument,omitempty"`
	// The new name of the access control policy.
	//
	// The name must be 1 to 128 characters in length. The name can contain letters, digits, and hyphens (-) and must start with a letter.
	NewPolicyName *string `json:"NewPolicyName,omitempty" xml:"NewPolicyName,omitempty"`
	// The ID of the access control policy.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s UpdateControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateControlPolicyRequest) SetNewDescription(v string) *UpdateControlPolicyRequest {
	s.NewDescription = &v
	return s
}

func (s *UpdateControlPolicyRequest) SetNewPolicyDocument(v string) *UpdateControlPolicyRequest {
	s.NewPolicyDocument = &v
	return s
}

func (s *UpdateControlPolicyRequest) SetNewPolicyName(v string) *UpdateControlPolicyRequest {
	s.NewPolicyName = &v
	return s
}

func (s *UpdateControlPolicyRequest) SetPolicyId(v string) *UpdateControlPolicyRequest {
	s.PolicyId = &v
	return s
}

type UpdateControlPolicyResponseBody struct {
	// The details of the access control policy.
	ControlPolicy *UpdateControlPolicyResponseBodyControlPolicy `json:"ControlPolicy,omitempty" xml:"ControlPolicy,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateControlPolicyResponseBody) SetControlPolicy(v *UpdateControlPolicyResponseBodyControlPolicy) *UpdateControlPolicyResponseBody {
	s.ControlPolicy = v
	return s
}

func (s *UpdateControlPolicyResponseBody) SetRequestId(v string) *UpdateControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type UpdateControlPolicyResponseBodyControlPolicy struct {
	// The number of times that the access control policy is referenced.
	AttachmentCount *string `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty"`
	// The time when the access control policy was created.
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the access control policy.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The effective scope of the access control policy. Valid values:
	//
	// *   All: The access control policy is in effect for Alibaba Cloud accounts, RAM users, and RAM roles.
	// *   RAM: The access control policy is in effect only for RAM users and RAM roles.
	EffectScope *string `json:"EffectScope,omitempty" xml:"EffectScope,omitempty"`
	// The ID of the access control policy.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the access control policy.
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the access control policy. Valid values:
	//
	// *   System: system access control policy
	// *   Custom: custom access control policy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The time when the access control policy was updated.
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s UpdateControlPolicyResponseBodyControlPolicy) String() string {
	return tea.Prettify(s)
}

func (s UpdateControlPolicyResponseBodyControlPolicy) GoString() string {
	return s.String()
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) SetAttachmentCount(v string) *UpdateControlPolicyResponseBodyControlPolicy {
	s.AttachmentCount = &v
	return s
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) SetCreateDate(v string) *UpdateControlPolicyResponseBodyControlPolicy {
	s.CreateDate = &v
	return s
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) SetDescription(v string) *UpdateControlPolicyResponseBodyControlPolicy {
	s.Description = &v
	return s
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) SetEffectScope(v string) *UpdateControlPolicyResponseBodyControlPolicy {
	s.EffectScope = &v
	return s
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) SetPolicyId(v string) *UpdateControlPolicyResponseBodyControlPolicy {
	s.PolicyId = &v
	return s
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) SetPolicyName(v string) *UpdateControlPolicyResponseBodyControlPolicy {
	s.PolicyName = &v
	return s
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) SetPolicyType(v string) *UpdateControlPolicyResponseBodyControlPolicy {
	s.PolicyType = &v
	return s
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) SetUpdateDate(v string) *UpdateControlPolicyResponseBodyControlPolicy {
	s.UpdateDate = &v
	return s
}

type UpdateControlPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateControlPolicyResponse) SetHeaders(v map[string]*string) *UpdateControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateControlPolicyResponse) SetStatusCode(v int32) *UpdateControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateControlPolicyResponse) SetBody(v *UpdateControlPolicyResponseBody) *UpdateControlPolicyResponse {
	s.Body = v
	return s
}

type UpdateFolderRequest struct {
	// The ID of the folder.
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The new name of the folder.
	//
	// The name must be 1 to 24 characters in length and can contain letters, digits, underscores (\_), periods (.), and hyphens (-).
	NewFolderName *string `json:"NewFolderName,omitempty" xml:"NewFolderName,omitempty"`
}

func (s UpdateFolderRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFolderRequest) GoString() string {
	return s.String()
}

func (s *UpdateFolderRequest) SetFolderId(v string) *UpdateFolderRequest {
	s.FolderId = &v
	return s
}

func (s *UpdateFolderRequest) SetNewFolderName(v string) *UpdateFolderRequest {
	s.NewFolderName = &v
	return s
}

type UpdateFolderResponseBody struct {
	// The information of the folder.
	Folder *UpdateFolderResponseBodyFolder `json:"Folder,omitempty" xml:"Folder,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFolderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFolderResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFolderResponseBody) SetFolder(v *UpdateFolderResponseBodyFolder) *UpdateFolderResponseBody {
	s.Folder = v
	return s
}

func (s *UpdateFolderResponseBody) SetRequestId(v string) *UpdateFolderResponseBody {
	s.RequestId = &v
	return s
}

type UpdateFolderResponseBodyFolder struct {
	// The time when the folder was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the folder.
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The name of the folder.
	FolderName *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
	// The ID of the parent folder.
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
}

func (s UpdateFolderResponseBodyFolder) String() string {
	return tea.Prettify(s)
}

func (s UpdateFolderResponseBodyFolder) GoString() string {
	return s.String()
}

func (s *UpdateFolderResponseBodyFolder) SetCreateTime(v string) *UpdateFolderResponseBodyFolder {
	s.CreateTime = &v
	return s
}

func (s *UpdateFolderResponseBodyFolder) SetFolderId(v string) *UpdateFolderResponseBodyFolder {
	s.FolderId = &v
	return s
}

func (s *UpdateFolderResponseBodyFolder) SetFolderName(v string) *UpdateFolderResponseBodyFolder {
	s.FolderName = &v
	return s
}

func (s *UpdateFolderResponseBodyFolder) SetParentFolderId(v string) *UpdateFolderResponseBodyFolder {
	s.ParentFolderId = &v
	return s
}

type UpdateFolderResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFolderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFolderResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFolderResponse) GoString() string {
	return s.String()
}

func (s *UpdateFolderResponse) SetHeaders(v map[string]*string) *UpdateFolderResponse {
	s.Headers = v
	return s
}

func (s *UpdateFolderResponse) SetStatusCode(v int32) *UpdateFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFolderResponse) SetBody(v *UpdateFolderResponseBody) *UpdateFolderResponse {
	s.Body = v
	return s
}

type UpdateResourceGroupRequest struct {
	// The display name of the resource group.
	//
	// The name must be 1 to 50 characters in length.
	NewDisplayName *string `json:"NewDisplayName,omitempty" xml:"NewDisplayName,omitempty"`
	// The ID of the resource group.
	//
	// You can call the [ListResourceGroups](~~158855~~) operation to obtain the ID.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s UpdateResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupRequest) SetNewDisplayName(v string) *UpdateResourceGroupRequest {
	s.NewDisplayName = &v
	return s
}

func (s *UpdateResourceGroupRequest) SetResourceGroupId(v string) *UpdateResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

type UpdateResourceGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the resource group.
	ResourceGroup *UpdateResourceGroupResponseBodyResourceGroup `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty" type:"Struct"`
}

func (s UpdateResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupResponseBody) SetRequestId(v string) *UpdateResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceGroupResponseBody) SetResourceGroup(v *UpdateResourceGroupResponseBodyResourceGroup) *UpdateResourceGroupResponseBody {
	s.ResourceGroup = v
	return s
}

type UpdateResourceGroupResponseBodyResourceGroup struct {
	// The ID of the Alibaba Cloud account to which the resource group belongs.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The time when the resource group was created. The time is displayed in UTC.
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The display name of the resource group.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the resource group.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The unique identifier of the resource group.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateResourceGroupResponseBodyResourceGroup) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceGroupResponseBodyResourceGroup) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupResponseBodyResourceGroup) SetAccountId(v string) *UpdateResourceGroupResponseBodyResourceGroup {
	s.AccountId = &v
	return s
}

func (s *UpdateResourceGroupResponseBodyResourceGroup) SetCreateDate(v string) *UpdateResourceGroupResponseBodyResourceGroup {
	s.CreateDate = &v
	return s
}

func (s *UpdateResourceGroupResponseBodyResourceGroup) SetDisplayName(v string) *UpdateResourceGroupResponseBodyResourceGroup {
	s.DisplayName = &v
	return s
}

func (s *UpdateResourceGroupResponseBodyResourceGroup) SetId(v string) *UpdateResourceGroupResponseBodyResourceGroup {
	s.Id = &v
	return s
}

func (s *UpdateResourceGroupResponseBodyResourceGroup) SetName(v string) *UpdateResourceGroupResponseBodyResourceGroup {
	s.Name = &v
	return s
}

type UpdateResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupResponse) SetHeaders(v map[string]*string) *UpdateResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceGroupResponse) SetStatusCode(v int32) *UpdateResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceGroupResponse) SetBody(v *UpdateResourceGroupResponseBody) *UpdateResourceGroupResponse {
	s.Body = v
	return s
}

type UpdateRoleRequest struct {
	// The document of the policy that specifies the trusted entity to assume the RAM role.
	NewAssumeRolePolicyDocument *string `json:"NewAssumeRolePolicyDocument,omitempty" xml:"NewAssumeRolePolicyDocument,omitempty"`
	// The description of the RAM role.
	//
	// The description must be 1 to 1,024 characters in length.
	NewDescription *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	// The maximum session duration of the RAM role.
	//
	// Unit: seconds. Valid values: 3600 to 43200. Default value: 3600.
	//
	// If you do not specify this parameter, the default value is used.
	NewMaxSessionDuration *int64 `json:"NewMaxSessionDuration,omitempty" xml:"NewMaxSessionDuration,omitempty"`
	// The name of the RAM role.
	//
	// The name must be 1 to 64 characters in length and can contain letters, digits, periods (.),and hyphens (-).
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
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the RAM role.
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
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The document of the policy that specifies the trusted entity to assume the RAM role.
	AssumeRolePolicyDocument *string `json:"AssumeRolePolicyDocument,omitempty" xml:"AssumeRolePolicyDocument,omitempty"`
	// The time when the RAM role was created.
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the RAM role.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum session duration of the RAM role.
	MaxSessionDuration *int64 `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty"`
	// The ID of the RAM role.
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// The name of the RAM role.
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The name of the RAM role after authorization.
	RolePrincipalName *string `json:"RolePrincipalName,omitempty" xml:"RolePrincipalName,omitempty"`
	// The time when the RAM role was updated.
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

func (s *UpdateRoleResponseBodyRole) SetRolePrincipalName(v string) *UpdateRoleResponseBodyRole {
	s.RolePrincipalName = &v
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("resourcemanager"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
 * After an invited Alibaba Cloud account joins a resource directory, it becomes a member of the resource directory. By default, the name of the invited Alibaba Cloud account is used as the display name of the account in the resource directory.
 * This topic provides an example on how to call the API operation to accept the invitation `h-Ih8IuPfvV0t0****` that is initiated to invite the Alibaba Cloud account `177242285274****` to join the resource directory `rd-3G****`.
 *
 * @param request AcceptHandshakeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AcceptHandshakeResponse
 */
func (client *Client) AcceptHandshakeWithOptions(request *AcceptHandshakeRequest, runtime *util.RuntimeOptions) (_result *AcceptHandshakeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HandshakeId)) {
		query["HandshakeId"] = request.HandshakeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AcceptHandshake"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AcceptHandshakeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * After an invited Alibaba Cloud account joins a resource directory, it becomes a member of the resource directory. By default, the name of the invited Alibaba Cloud account is used as the display name of the account in the resource directory.
 * This topic provides an example on how to call the API operation to accept the invitation `h-Ih8IuPfvV0t0****` that is initiated to invite the Alibaba Cloud account `177242285274****` to join the resource directory `rd-3G****`.
 *
 * @param request AcceptHandshakeRequest
 * @return AcceptHandshakeResponse
 */
func (client *Client) AcceptHandshake(request *AcceptHandshakeRequest) (_result *AcceptHandshakeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AcceptHandshakeResponse{}
	_body, _err := client.AcceptHandshakeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * After you attach an access control policy, the operations performed on resources by using members are limited by the policy. Make sure that the attached policy meets your expectations. Otherwise, your business may be affected.
 * By default, the system access control policy FullAliyunAccess is attached to each folder and member.
 * The access control policy that is attached to a folder also applies to all its subfolders and all members in the subfolders.
 * A maximum of 10 access control policies can be attached to a folder or member.
 * This topic provides an example on how to call the API operation to attach the custom access control policy `cp-jExXAqIYkwHN****` to the folder `fd-ZDNPiT****`.
 *
 * @param request AttachControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AttachControlPolicyResponse
 */
func (client *Client) AttachControlPolicyWithOptions(request *AttachControlPolicyRequest, runtime *util.RuntimeOptions) (_result *AttachControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["TargetId"] = request.TargetId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachControlPolicy"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * After you attach an access control policy, the operations performed on resources by using members are limited by the policy. Make sure that the attached policy meets your expectations. Otherwise, your business may be affected.
 * By default, the system access control policy FullAliyunAccess is attached to each folder and member.
 * The access control policy that is attached to a folder also applies to all its subfolders and all members in the subfolders.
 * A maximum of 10 access control policies can be attached to a folder or member.
 * This topic provides an example on how to call the API operation to attach the custom access control policy `cp-jExXAqIYkwHN****` to the folder `fd-ZDNPiT****`.
 *
 * @param request AttachControlPolicyRequest
 * @return AttachControlPolicyResponse
 */
func (client *Client) AttachControlPolicy(request *AttachControlPolicyRequest) (_result *AttachControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachControlPolicyResponse{}
	_body, _err := client.AttachControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this example, the policy `AdministratorAccess` is attached to the RAM user `alice@demo.onaliyun.com` and takes effect only for resources in the `rg-9gLOoK****` resource group.
 *
 * @param request AttachPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AttachPolicyResponse
 */
func (client *Client) AttachPolicyWithOptions(request *AttachPolicyRequest, runtime *util.RuntimeOptions) (_result *AttachPolicyResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.PrincipalName)) {
		query["PrincipalName"] = request.PrincipalName
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalType)) {
		query["PrincipalType"] = request.PrincipalType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachPolicy"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, the policy `AdministratorAccess` is attached to the RAM user `alice@demo.onaliyun.com` and takes effect only for resources in the `rg-9gLOoK****` resource group.
 *
 * @param request AttachPolicyRequest
 * @return AttachPolicyResponse
 */
func (client *Client) AttachPolicy(request *AttachPolicyRequest) (_result *AttachPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachPolicyResponse{}
	_body, _err := client.AttachPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation only to bind a mobile phone number to a member of the resource account type. You cannot call this API operation to change the mobile phone number that is bound to a member of the resource account type.
 * To ensure that the system can record the operators of management operations, you must use a RAM user or RAM role to which the AliyunResourceDirectoryFullAccess policy is attached within the management account of your resource directory to call this API operation.
 * This topic provides an example on how to call the API operation to bind a mobile phone number to the member `138660628348****` for security purposes.
 *
 * @param request BindSecureMobilePhoneRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return BindSecureMobilePhoneResponse
 */
func (client *Client) BindSecureMobilePhoneWithOptions(request *BindSecureMobilePhoneRequest, runtime *util.RuntimeOptions) (_result *BindSecureMobilePhoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.SecureMobilePhone)) {
		query["SecureMobilePhone"] = request.SecureMobilePhone
	}

	if !tea.BoolValue(util.IsUnset(request.VerificationCode)) {
		query["VerificationCode"] = request.VerificationCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindSecureMobilePhone"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindSecureMobilePhoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation only to bind a mobile phone number to a member of the resource account type. You cannot call this API operation to change the mobile phone number that is bound to a member of the resource account type.
 * To ensure that the system can record the operators of management operations, you must use a RAM user or RAM role to which the AliyunResourceDirectoryFullAccess policy is attached within the management account of your resource directory to call this API operation.
 * This topic provides an example on how to call the API operation to bind a mobile phone number to the member `138660628348****` for security purposes.
 *
 * @param request BindSecureMobilePhoneRequest
 * @return BindSecureMobilePhoneResponse
 */
func (client *Client) BindSecureMobilePhone(request *BindSecureMobilePhoneRequest) (_result *BindSecureMobilePhoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindSecureMobilePhoneResponse{}
	_body, _err := client.BindSecureMobilePhoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelChangeAccountEmailWithOptions(request *CancelChangeAccountEmailRequest, runtime *util.RuntimeOptions) (_result *CancelChangeAccountEmailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelChangeAccountEmail"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelChangeAccountEmailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelChangeAccountEmail(request *CancelChangeAccountEmailRequest) (_result *CancelChangeAccountEmailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelChangeAccountEmailResponse{}
	_body, _err := client.CancelChangeAccountEmailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelCreateCloudAccountWithOptions(request *CancelCreateCloudAccountRequest, runtime *util.RuntimeOptions) (_result *CancelCreateCloudAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RecordId)) {
		query["RecordId"] = request.RecordId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelCreateCloudAccount"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelCreateCloudAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelCreateCloudAccount(request *CancelCreateCloudAccountRequest) (_result *CancelCreateCloudAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelCreateCloudAccountResponse{}
	_body, _err := client.CancelCreateCloudAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to cancel the invitation whose ID is `h-ycm4rp****`.
 *
 * @param request CancelHandshakeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CancelHandshakeResponse
 */
func (client *Client) CancelHandshakeWithOptions(request *CancelHandshakeRequest, runtime *util.RuntimeOptions) (_result *CancelHandshakeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HandshakeId)) {
		query["HandshakeId"] = request.HandshakeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelHandshake"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelHandshakeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to cancel the invitation whose ID is `h-ycm4rp****`.
 *
 * @param request CancelHandshakeRequest
 * @return CancelHandshakeResponse
 */
func (client *Client) CancelHandshake(request *CancelHandshakeRequest) (_result *CancelHandshakeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelHandshakeResponse{}
	_body, _err := client.CancelHandshakeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelPromoteResourceAccountWithOptions(request *CancelPromoteResourceAccountRequest, runtime *util.RuntimeOptions) (_result *CancelPromoteResourceAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RecordId)) {
		query["RecordId"] = request.RecordId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelPromoteResourceAccount"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelPromoteResourceAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelPromoteResourceAccount(request *CancelPromoteResourceAccountRequest) (_result *CancelPromoteResourceAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelPromoteResourceAccountResponse{}
	_body, _err := client.CancelPromoteResourceAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangeAccountEmailWithOptions(request *ChangeAccountEmailRequest, runtime *util.RuntimeOptions) (_result *ChangeAccountEmailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeAccountEmail"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeAccountEmailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeAccountEmail(request *ChangeAccountEmailRequest) (_result *ChangeAccountEmailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeAccountEmailResponse{}
	_body, _err := client.ChangeAccountEmailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Before you delete a member, you must call this API operation to check whether the member can be deleted.
 * This topic provides an example on how to call the API operation to perform a deletion check on the member whose ID is `179855839641****`.
 *
 * @param request CheckAccountDeleteRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CheckAccountDeleteResponse
 */
func (client *Client) CheckAccountDeleteWithOptions(request *CheckAccountDeleteRequest, runtime *util.RuntimeOptions) (_result *CheckAccountDeleteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckAccountDelete"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckAccountDeleteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Before you delete a member, you must call this API operation to check whether the member can be deleted.
 * This topic provides an example on how to call the API operation to perform a deletion check on the member whose ID is `179855839641****`.
 *
 * @param request CheckAccountDeleteRequest
 * @return CheckAccountDeleteResponse
 */
func (client *Client) CheckAccountDelete(request *CheckAccountDeleteRequest) (_result *CheckAccountDeleteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckAccountDeleteResponse{}
	_body, _err := client.CheckAccountDeleteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * A resource directory supports two types of member accounts: resource accounts and cloud accounts.
 * *   Resource account (recommended): A resource account is only used as a resource container and fully depends on a resource directory. Such member accounts are secure and easy-to-create. For more information about how to create a resource account, see [CreateResourceAccount](~~159392~~).
 * >  A resource account can be upgraded to a cloud account. For more information, see [PromoteResourceAccount](~~159395~~) .
 * *   Cloud account: A cloud account has all the features of an Alibaba Cloud account, including root permissions.
 *
 * @param request CreateCloudAccountRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateCloudAccountResponse
 */
func (client *Client) CreateCloudAccountWithOptions(request *CreateCloudAccountRequest, runtime *util.RuntimeOptions) (_result *CreateCloudAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.ParentFolderId)) {
		query["ParentFolderId"] = request.ParentFolderId
	}

	if !tea.BoolValue(util.IsUnset(request.PayerAccountId)) {
		query["PayerAccountId"] = request.PayerAccountId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCloudAccount"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCloudAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * A resource directory supports two types of member accounts: resource accounts and cloud accounts.
 * *   Resource account (recommended): A resource account is only used as a resource container and fully depends on a resource directory. Such member accounts are secure and easy-to-create. For more information about how to create a resource account, see [CreateResourceAccount](~~159392~~).
 * >  A resource account can be upgraded to a cloud account. For more information, see [PromoteResourceAccount](~~159395~~) .
 * *   Cloud account: A cloud account has all the features of an Alibaba Cloud account, including root permissions.
 *
 * @param request CreateCloudAccountRequest
 * @return CreateCloudAccountResponse
 */
func (client *Client) CreateCloudAccount(request *CreateCloudAccountRequest) (_result *CreateCloudAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCloudAccountResponse{}
	_body, _err := client.CreateCloudAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to create a custom access control policy named `ExampleControlPolicy`. This access control policy is used to prohibit modifications to the ResourceDirectoryAccountAccessRole role and the permissions of the role.
 *
 * @param request CreateControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateControlPolicyResponse
 */
func (client *Client) CreateControlPolicyWithOptions(request *CreateControlPolicyRequest, runtime *util.RuntimeOptions) (_result *CreateControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EffectScope)) {
		query["EffectScope"] = request.EffectScope
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
		Action:      tea.String("CreateControlPolicy"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to create a custom access control policy named `ExampleControlPolicy`. This access control policy is used to prohibit modifications to the ResourceDirectoryAccountAccessRole role and the permissions of the role.
 *
 * @param request CreateControlPolicyRequest
 * @return CreateControlPolicyResponse
 */
func (client *Client) CreateControlPolicy(request *CreateControlPolicyRequest) (_result *CreateControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateControlPolicyResponse{}
	_body, _err := client.CreateControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * >  A maximum of five levels of folders can be created under the root folder.
 * In this example, a folder named `rdFolder` is created under the root folder.
 *
 * @param request CreateFolderRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateFolderResponse
 */
func (client *Client) CreateFolderWithOptions(request *CreateFolderRequest, runtime *util.RuntimeOptions) (_result *CreateFolderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FolderName)) {
		query["FolderName"] = request.FolderName
	}

	if !tea.BoolValue(util.IsUnset(request.ParentFolderId)) {
		query["ParentFolderId"] = request.ParentFolderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFolder"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFolderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * >  A maximum of five levels of folders can be created under the root folder.
 * In this example, a folder named `rdFolder` is created under the root folder.
 *
 * @param request CreateFolderRequest
 * @return CreateFolderResponse
 */
func (client *Client) CreateFolder(request *CreateFolderRequest) (_result *CreateFolderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFolderResponse{}
	_body, _err := client.CreateFolderWithOptions(request, runtime)
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
		Version:     tea.String("2020-03-31"),
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

	if !tea.BoolValue(util.IsUnset(request.SetAsDefault)) {
		query["SetAsDefault"] = request.SetAsDefault
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePolicyVersion"),
		Version:     tea.String("2020-03-31"),
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

/**
 * A member serves as a container for resources and is also an organizational unit in a resource directory. A member indicates a project or application. The resources of different members are isolated.
 * This topic provides an example on how to call the API operation to create a member in the `fd-r23M55****` folder. The display name of the member is `Dev`, and the prefix for the Alibaba Cloud account name of the member is `alice`.
 *
 * @param request CreateResourceAccountRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateResourceAccountResponse
 */
func (client *Client) CreateResourceAccountWithOptions(request *CreateResourceAccountRequest, runtime *util.RuntimeOptions) (_result *CreateResourceAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountNamePrefix)) {
		query["AccountNamePrefix"] = request.AccountNamePrefix
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.ParentFolderId)) {
		query["ParentFolderId"] = request.ParentFolderId
	}

	if !tea.BoolValue(util.IsUnset(request.PayerAccountId)) {
		query["PayerAccountId"] = request.PayerAccountId
	}

	if !tea.BoolValue(util.IsUnset(request.ResellAccountType)) {
		query["ResellAccountType"] = request.ResellAccountType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateResourceAccount"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateResourceAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * A member serves as a container for resources and is also an organizational unit in a resource directory. A member indicates a project or application. The resources of different members are isolated.
 * This topic provides an example on how to call the API operation to create a member in the `fd-r23M55****` folder. The display name of the member is `Dev`, and the prefix for the Alibaba Cloud account name of the member is `alice`.
 *
 * @param request CreateResourceAccountRequest
 * @return CreateResourceAccountResponse
 */
func (client *Client) CreateResourceAccount(request *CreateResourceAccountRequest) (_result *CreateResourceAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateResourceAccountResponse{}
	_body, _err := client.CreateResourceAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * >  A maximum of 30 resource groups can be created within an Alibaba Cloud account.
 *
 * @param request CreateResourceGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateResourceGroupResponse
 */
func (client *Client) CreateResourceGroupWithOptions(request *CreateResourceGroupRequest, runtime *util.RuntimeOptions) (_result *CreateResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateResourceGroup"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * >  A maximum of 30 resource groups can be created within an Alibaba Cloud account.
 *
 * @param request CreateResourceGroupRequest
 * @return CreateResourceGroupResponse
 */
func (client *Client) CreateResourceGroup(request *CreateResourceGroupRequest) (_result *CreateResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateResourceGroupResponse{}
	_body, _err := client.CreateResourceGroupWithOptions(request, runtime)
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
		Version:     tea.String("2020-03-31"),
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

func (client *Client) CreateServiceLinkedRoleWithOptions(request *CreateServiceLinkedRoleRequest, runtime *util.RuntimeOptions) (_result *CreateServiceLinkedRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomSuffix)) {
		query["CustomSuffix"] = request.CustomSuffix
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceLinkedRole"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateServiceLinkedRole(request *CreateServiceLinkedRoleRequest) (_result *CreateServiceLinkedRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.CreateServiceLinkedRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeclineHandshakeWithOptions(request *DeclineHandshakeRequest, runtime *util.RuntimeOptions) (_result *DeclineHandshakeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HandshakeId)) {
		query["HandshakeId"] = request.HandshakeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeclineHandshake"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeclineHandshakeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeclineHandshake(request *DeclineHandshakeRequest) (_result *DeclineHandshakeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeclineHandshakeResponse{}
	_body, _err := client.DeclineHandshakeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The ID of the member that you want to delete.
 *
 * @param tmpReq DeleteAccountRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteAccountResponse
 */
func (client *Client) DeleteAccountWithOptions(tmpReq *DeleteAccountRequest, runtime *util.RuntimeOptions) (_result *DeleteAccountResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteAccountShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AbandonableCheckId)) {
		request.AbandonableCheckIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AbandonableCheckId, tea.String("AbandonableCheckId"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AbandonableCheckIdShrink)) {
		query["AbandonableCheckId"] = request.AbandonableCheckIdShrink
	}

	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccount"),
		Version:     tea.String("2020-03-31"),
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

/**
 * The ID of the member that you want to delete.
 *
 * @param request DeleteAccountRequest
 * @return DeleteAccountResponse
 */
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

/**
 * If you want to delete a custom control policy that is attached to folders or member accounts, you must call the [DetachControlPolicy](~~208331~~) operation to detach the policy before you delete it.
 * In this example, the custom control policy `cp-SImPt8GCEwiq****` is deleted.
 *
 * @param request DeleteControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteControlPolicyResponse
 */
func (client *Client) DeleteControlPolicyWithOptions(request *DeleteControlPolicyRequest, runtime *util.RuntimeOptions) (_result *DeleteControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteControlPolicy"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If you want to delete a custom control policy that is attached to folders or member accounts, you must call the [DetachControlPolicy](~~208331~~) operation to detach the policy before you delete it.
 * In this example, the custom control policy `cp-SImPt8GCEwiq****` is deleted.
 *
 * @param request DeleteControlPolicyRequest
 * @return DeleteControlPolicyResponse
 */
func (client *Client) DeleteControlPolicy(request *DeleteControlPolicyRequest) (_result *DeleteControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteControlPolicyResponse{}
	_body, _err := client.DeleteControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * >  Before you delete a folder, make sure that the folder does not contain any member accounts or child folders.
 *
 * @param request DeleteFolderRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteFolderResponse
 */
func (client *Client) DeleteFolderWithOptions(request *DeleteFolderRequest, runtime *util.RuntimeOptions) (_result *DeleteFolderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FolderId)) {
		query["FolderId"] = request.FolderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFolder"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFolderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * >  Before you delete a folder, make sure that the folder does not contain any member accounts or child folders.
 *
 * @param request DeleteFolderRequest
 * @return DeleteFolderResponse
 */
func (client *Client) DeleteFolder(request *DeleteFolderRequest) (_result *DeleteFolderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFolderResponse{}
	_body, _err := client.DeleteFolderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * >
 * *   Before you delete a policy, you must delete all non-default versions of the policy. For more information about how to delete a policy version, see [DeletePolicyVersion](~~159041~~).
 * *   Before you delete a policy, make sure that the policy is not referenced. This means that the policy is not attached to RAM users, RAM user groups, or RAM roles. For more information about how to detach a policy, see [DetachPolicy](~~159168~~).
 *
 * @param request DeletePolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeletePolicyResponse
 */
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
		Version:     tea.String("2020-03-31"),
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

/**
 * >
 * *   Before you delete a policy, you must delete all non-default versions of the policy. For more information about how to delete a policy version, see [DeletePolicyVersion](~~159041~~).
 * *   Before you delete a policy, make sure that the policy is not referenced. This means that the policy is not attached to RAM users, RAM user groups, or RAM roles. For more information about how to detach a policy, see [DetachPolicy](~~159168~~).
 *
 * @param request DeletePolicyRequest
 * @return DeletePolicyResponse
 */
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

/**
 * >  The default version of a permission policy cannot be deleted.
 *
 * @param request DeletePolicyVersionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeletePolicyVersionResponse
 */
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
		Version:     tea.String("2020-03-31"),
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

/**
 * >  The default version of a permission policy cannot be deleted.
 *
 * @param request DeletePolicyVersionRequest
 * @return DeletePolicyVersionResponse
 */
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

/**
 * >  Before you delete a resource group, you must delete all the resources in it.
 * In this example, the resource group whose ID is `rg-9gLOoK****` is deleted.
 *
 * @param request DeleteResourceGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteResourceGroupResponse
 */
func (client *Client) DeleteResourceGroupWithOptions(request *DeleteResourceGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteResourceGroup"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * >  Before you delete a resource group, you must delete all the resources in it.
 * In this example, the resource group whose ID is `rg-9gLOoK****` is deleted.
 *
 * @param request DeleteResourceGroupRequest
 * @return DeleteResourceGroupResponse
 */
func (client *Client) DeleteResourceGroup(request *DeleteResourceGroupRequest) (_result *DeleteResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteResourceGroupResponse{}
	_body, _err := client.DeleteResourceGroupWithOptions(request, runtime)
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
		Version:     tea.String("2020-03-31"),
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

func (client *Client) DeleteServiceLinkedRoleWithOptions(request *DeleteServiceLinkedRoleRequest, runtime *util.RuntimeOptions) (_result *DeleteServiceLinkedRoleResponse, _err error) {
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
		Action:      tea.String("DeleteServiceLinkedRole"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceLinkedRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteServiceLinkedRole(request *DeleteServiceLinkedRoleRequest) (_result *DeleteServiceLinkedRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteServiceLinkedRoleResponse{}
	_body, _err := client.DeleteServiceLinkedRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * >  If the delegated administrator account that you want to remove has historical management tasks in the related trusted service, the trusted service may be affected after the delegated administrator account is removed. Therefore, proceed with caution.
 * This topic provides an example on how to call the API operation to remove the delegated administrator account `181761095690****` for Cloud Firewall.
 *
 * @param request DeregisterDelegatedAdministratorRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeregisterDelegatedAdministratorResponse
 */
func (client *Client) DeregisterDelegatedAdministratorWithOptions(request *DeregisterDelegatedAdministratorRequest, runtime *util.RuntimeOptions) (_result *DeregisterDelegatedAdministratorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.ServicePrincipal)) {
		query["ServicePrincipal"] = request.ServicePrincipal
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeregisterDelegatedAdministrator"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeregisterDelegatedAdministratorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * >  If the delegated administrator account that you want to remove has historical management tasks in the related trusted service, the trusted service may be affected after the delegated administrator account is removed. Therefore, proceed with caution.
 * This topic provides an example on how to call the API operation to remove the delegated administrator account `181761095690****` for Cloud Firewall.
 *
 * @param request DeregisterDelegatedAdministratorRequest
 * @return DeregisterDelegatedAdministratorResponse
 */
func (client *Client) DeregisterDelegatedAdministrator(request *DeregisterDelegatedAdministratorRequest) (_result *DeregisterDelegatedAdministratorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeregisterDelegatedAdministratorResponse{}
	_body, _err := client.DeregisterDelegatedAdministratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Before you disable a resource directory, make sure that the following requirements are met:
 * *   All member accounts must be removed from the resource directory. For more information about how to remove a member account, see [RemoveCloudAccount](~~159431~~).
 * *   All folders except the root folder must be deleted from the resource directory. For more information about how to delete a folder, see [DeleteFolder](~~159432~~).
 *
 * @param request DestroyResourceDirectoryRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DestroyResourceDirectoryResponse
 */
func (client *Client) DestroyResourceDirectoryWithOptions(runtime *util.RuntimeOptions) (_result *DestroyResourceDirectoryResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DestroyResourceDirectory"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DestroyResourceDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Before you disable a resource directory, make sure that the following requirements are met:
 * *   All member accounts must be removed from the resource directory. For more information about how to remove a member account, see [RemoveCloudAccount](~~159431~~).
 * *   All folders except the root folder must be deleted from the resource directory. For more information about how to delete a folder, see [DeleteFolder](~~159432~~).
 *
 * @return DestroyResourceDirectoryResponse
 */
func (client *Client) DestroyResourceDirectory() (_result *DestroyResourceDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DestroyResourceDirectoryResponse{}
	_body, _err := client.DestroyResourceDirectoryWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * After you detach an access control policy, the operations performed on resources by using members are not limited by the policy. Make sure that the detached policy meets your expectations. Otherwise, your business may be affected.
 * Both the system and custom access control policies can be detached. If an object has only one access control policy attached, the policy cannot be detached.
 * This topic provides an example on how to call the API operation to detach the custom control policy `cp-jExXAqIYkwHN****` from the folder `fd-ZDNPiT****`.
 *
 * @param request DetachControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DetachControlPolicyResponse
 */
func (client *Client) DetachControlPolicyWithOptions(request *DetachControlPolicyRequest, runtime *util.RuntimeOptions) (_result *DetachControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["TargetId"] = request.TargetId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachControlPolicy"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * After you detach an access control policy, the operations performed on resources by using members are not limited by the policy. Make sure that the detached policy meets your expectations. Otherwise, your business may be affected.
 * Both the system and custom access control policies can be detached. If an object has only one access control policy attached, the policy cannot be detached.
 * This topic provides an example on how to call the API operation to detach the custom control policy `cp-jExXAqIYkwHN****` from the folder `fd-ZDNPiT****`.
 *
 * @param request DetachControlPolicyRequest
 * @return DetachControlPolicyResponse
 */
func (client *Client) DetachControlPolicy(request *DetachControlPolicyRequest) (_result *DetachControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachControlPolicyResponse{}
	_body, _err := client.DetachControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachPolicyWithOptions(request *DetachPolicyRequest, runtime *util.RuntimeOptions) (_result *DetachPolicyResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.PrincipalName)) {
		query["PrincipalName"] = request.PrincipalName
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalType)) {
		query["PrincipalType"] = request.PrincipalType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachPolicy"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachPolicy(request *DetachPolicyRequest) (_result *DetachPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachPolicyResponse{}
	_body, _err := client.DetachPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableAssociatedTransferWithOptions(runtime *util.RuntimeOptions) (_result *DisableAssociatedTransferResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DisableAssociatedTransfer"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableAssociatedTransferResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableAssociatedTransfer() (_result *DisableAssociatedTransferResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableAssociatedTransferResponse{}
	_body, _err := client.DisableAssociatedTransferWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * After you disable the Control Policy feature, the system automatically detaches all control policies that are attached to folders and member accounts. The system does not delete these control policies, but you cannot attach them to folders or member accounts again.
 * >  If you disable the Control Policy feature, the permissions of all folders and member accounts in a resource directory are affected. You must proceed with caution.
 *
 * @param request DisableControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DisableControlPolicyResponse
 */
func (client *Client) DisableControlPolicyWithOptions(runtime *util.RuntimeOptions) (_result *DisableControlPolicyResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DisableControlPolicy"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * After you disable the Control Policy feature, the system automatically detaches all control policies that are attached to folders and member accounts. The system does not delete these control policies, but you cannot attach them to folders or member accounts again.
 * >  If you disable the Control Policy feature, the permissions of all folders and member accounts in a resource directory are affected. You must proceed with caution.
 *
 * @return DisableControlPolicyResponse
 */
func (client *Client) DisableControlPolicy() (_result *DisableControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableControlPolicyResponse{}
	_body, _err := client.DisableControlPolicyWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableAssociatedTransferWithOptions(runtime *util.RuntimeOptions) (_result *EnableAssociatedTransferResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("EnableAssociatedTransfer"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableAssociatedTransferResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableAssociatedTransfer() (_result *EnableAssociatedTransferResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableAssociatedTransferResponse{}
	_body, _err := client.EnableAssociatedTransferWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The Control Policy feature allows you to manage the permission boundaries of the folders or member accounts in a resource directory in a centralized manner. This feature is implemented based on the resource directory. You can use this feature to develop common or dedicated rules for access control. The Control Policy feature does not grant permissions but only defines permission boundaries. A member account in a resource directory can be used to access resources only after it is granted the required permissions by using the Resource Access Management (RAM) service. For more information, see [Overview of the Control Policy feature](~~178671~~).
 *
 * @param request EnableControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return EnableControlPolicyResponse
 */
func (client *Client) EnableControlPolicyWithOptions(runtime *util.RuntimeOptions) (_result *EnableControlPolicyResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("EnableControlPolicy"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The Control Policy feature allows you to manage the permission boundaries of the folders or member accounts in a resource directory in a centralized manner. This feature is implemented based on the resource directory. You can use this feature to develop common or dedicated rules for access control. The Control Policy feature does not grant permissions but only defines permission boundaries. A member account in a resource directory can be used to access resources only after it is granted the required permissions by using the Resource Access Management (RAM) service. For more information, see [Overview of the Control Policy feature](~~178671~~).
 *
 * @return EnableControlPolicyResponse
 */
func (client *Client) EnableControlPolicy() (_result *EnableControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableControlPolicyResponse{}
	_body, _err := client.EnableControlPolicyWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can use the current account or a newly created account to enable a resource directory. For more information, see [Enable a resource directory](~~111215~~).
 * In this example, the current account is used to enable a resource directory.
 *
 * @param request EnableResourceDirectoryRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return EnableResourceDirectoryResponse
 */
func (client *Client) EnableResourceDirectoryWithOptions(request *EnableResourceDirectoryRequest, runtime *util.RuntimeOptions) (_result *EnableResourceDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnableMode)) {
		query["EnableMode"] = request.EnableMode
	}

	if !tea.BoolValue(util.IsUnset(request.MAName)) {
		query["MAName"] = request.MAName
	}

	if !tea.BoolValue(util.IsUnset(request.MASecureMobilePhone)) {
		query["MASecureMobilePhone"] = request.MASecureMobilePhone
	}

	if !tea.BoolValue(util.IsUnset(request.VerificationCode)) {
		query["VerificationCode"] = request.VerificationCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableResourceDirectory"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableResourceDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can use the current account or a newly created account to enable a resource directory. For more information, see [Enable a resource directory](~~111215~~).
 * In this example, the current account is used to enable a resource directory.
 *
 * @param request EnableResourceDirectoryRequest
 * @return EnableResourceDirectoryResponse
 */
func (client *Client) EnableResourceDirectory(request *EnableResourceDirectoryRequest) (_result *EnableResourceDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableResourceDirectoryResponse{}
	_body, _err := client.EnableResourceDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the information of the member whose Alibaba Cloud account ID is `181761095690****`.
 *
 * @param request GetAccountRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetAccountResponse
 */
func (client *Client) GetAccountWithOptions(request *GetAccountRequest, runtime *util.RuntimeOptions) (_result *GetAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeTags)) {
		query["IncludeTags"] = request.IncludeTags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAccount"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the information of the member whose Alibaba Cloud account ID is `181761095690****`.
 *
 * @param request GetAccountRequest
 * @return GetAccountResponse
 */
func (client *Client) GetAccount(request *GetAccountRequest) (_result *GetAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccountResponse{}
	_body, _err := client.GetAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * After you call the [CheckAccountDelete](~~448542~~) operation to perform a member deletion check, you can call the GetAccountDeletionCheckResult operation to query the check result. If the check result shows that the member meets deletion requirements, you can delete the member. Otherwise, you need to first modify the items that do not meet requirements.
 * This topic provides an example on how to call the API operation to query the result of the deletion check for the member whose ID is `179855839641****`. The response shows that the member does not meet deletion requirements.
 *
 * @param request GetAccountDeletionCheckResultRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetAccountDeletionCheckResultResponse
 */
func (client *Client) GetAccountDeletionCheckResultWithOptions(request *GetAccountDeletionCheckResultRequest, runtime *util.RuntimeOptions) (_result *GetAccountDeletionCheckResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAccountDeletionCheckResult"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccountDeletionCheckResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * After you call the [CheckAccountDelete](~~448542~~) operation to perform a member deletion check, you can call the GetAccountDeletionCheckResult operation to query the check result. If the check result shows that the member meets deletion requirements, you can delete the member. Otherwise, you need to first modify the items that do not meet requirements.
 * This topic provides an example on how to call the API operation to query the result of the deletion check for the member whose ID is `179855839641****`. The response shows that the member does not meet deletion requirements.
 *
 * @param request GetAccountDeletionCheckResultRequest
 * @return GetAccountDeletionCheckResultResponse
 */
func (client *Client) GetAccountDeletionCheckResult(request *GetAccountDeletionCheckResultRequest) (_result *GetAccountDeletionCheckResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccountDeletionCheckResultResponse{}
	_body, _err := client.GetAccountDeletionCheckResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the deletion status of the member whose Alibaba Cloud account ID is `169946124551****`. The response shows that the member is deleted.
 *
 * @param request GetAccountDeletionStatusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetAccountDeletionStatusResponse
 */
func (client *Client) GetAccountDeletionStatusWithOptions(request *GetAccountDeletionStatusRequest, runtime *util.RuntimeOptions) (_result *GetAccountDeletionStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAccountDeletionStatus"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccountDeletionStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the deletion status of the member whose Alibaba Cloud account ID is `169946124551****`. The response shows that the member is deleted.
 *
 * @param request GetAccountDeletionStatusRequest
 * @return GetAccountDeletionStatusResponse
 */
func (client *Client) GetAccountDeletionStatus(request *GetAccountDeletionStatusRequest) (_result *GetAccountDeletionStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccountDeletionStatusResponse{}
	_body, _err := client.GetAccountDeletionStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the details of the access control policy whose ID is `cp-SImPt8GCEwiq****`.
 *
 * @param request GetControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetControlPolicyResponse
 */
func (client *Client) GetControlPolicyWithOptions(request *GetControlPolicyRequest, runtime *util.RuntimeOptions) (_result *GetControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetControlPolicy"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the details of the access control policy whose ID is `cp-SImPt8GCEwiq****`.
 *
 * @param request GetControlPolicyRequest
 * @return GetControlPolicyResponse
 */
func (client *Client) GetControlPolicy(request *GetControlPolicyRequest) (_result *GetControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetControlPolicyResponse{}
	_body, _err := client.GetControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetControlPolicyEnablementStatusWithOptions(runtime *util.RuntimeOptions) (_result *GetControlPolicyEnablementStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetControlPolicyEnablementStatus"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetControlPolicyEnablementStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetControlPolicyEnablementStatus() (_result *GetControlPolicyEnablementStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetControlPolicyEnablementStatusResponse{}
	_body, _err := client.GetControlPolicyEnablementStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this example, the information of the folder `fd-Jyl5U7****` is queried.
 *
 * @param request GetFolderRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetFolderResponse
 */
func (client *Client) GetFolderWithOptions(request *GetFolderRequest, runtime *util.RuntimeOptions) (_result *GetFolderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FolderId)) {
		query["FolderId"] = request.FolderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFolder"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFolderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, the information of the folder `fd-Jyl5U7****` is queried.
 *
 * @param request GetFolderRequest
 * @return GetFolderResponse
 */
func (client *Client) GetFolder(request *GetFolderRequest) (_result *GetFolderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFolderResponse{}
	_body, _err := client.GetFolderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this example, the information of the invitation whose ID is `h-ycm4rp****` is queried.
 *
 * @param request GetHandshakeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetHandshakeResponse
 */
func (client *Client) GetHandshakeWithOptions(request *GetHandshakeRequest, runtime *util.RuntimeOptions) (_result *GetHandshakeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HandshakeId)) {
		query["HandshakeId"] = request.HandshakeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHandshake"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHandshakeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, the information of the invitation whose ID is `h-ycm4rp****` is queried.
 *
 * @param request GetHandshakeRequest
 * @return GetHandshakeResponse
 */
func (client *Client) GetHandshake(request *GetHandshakeRequest) (_result *GetHandshakeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHandshakeResponse{}
	_body, _err := client.GetHandshakeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPayerForAccountWithOptions(request *GetPayerForAccountRequest, runtime *util.RuntimeOptions) (_result *GetPayerForAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPayerForAccount"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPayerForAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPayerForAccount(request *GetPayerForAccountRequest) (_result *GetPayerForAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPayerForAccountResponse{}
	_body, _err := client.GetPayerForAccountWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
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
		Action:      tea.String("GetPolicy"),
		Version:     tea.String("2020-03-31"),
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
		Version:     tea.String("2020-03-31"),
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

/**
 * This topic provides an example on how to use a management account to call the API operation to query the information of the resource directory that is enabled by using the management account.
 *
 * @param request GetResourceDirectoryRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetResourceDirectoryResponse
 */
func (client *Client) GetResourceDirectoryWithOptions(runtime *util.RuntimeOptions) (_result *GetResourceDirectoryResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetResourceDirectory"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to use a management account to call the API operation to query the information of the resource directory that is enabled by using the management account.
 *
 * @return GetResourceDirectoryResponse
 */
func (client *Client) GetResourceDirectory() (_result *GetResourceDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceDirectoryResponse{}
	_body, _err := client.GetResourceDirectoryWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * For more information about common request parameters, see [Common parameters](~~159973~~).
 *
 * @param request GetResourceGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetResourceGroupResponse
 */
func (client *Client) GetResourceGroupWithOptions(request *GetResourceGroupRequest, runtime *util.RuntimeOptions) (_result *GetResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncludeTags)) {
		query["IncludeTags"] = request.IncludeTags
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceGroup"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * For more information about common request parameters, see [Common parameters](~~159973~~).
 *
 * @param request GetResourceGroupRequest
 * @return GetResourceGroupResponse
 */
func (client *Client) GetResourceGroup(request *GetResourceGroupRequest) (_result *GetResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceGroupResponse{}
	_body, _err := client.GetResourceGroupWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		query["RoleName"] = request.RoleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRole"),
		Version:     tea.String("2020-03-31"),
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

func (client *Client) GetServiceLinkedRoleDeletionStatusWithOptions(request *GetServiceLinkedRoleDeletionStatusRequest, runtime *util.RuntimeOptions) (_result *GetServiceLinkedRoleDeletionStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeletionTaskId)) {
		query["DeletionTaskId"] = request.DeletionTaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceLinkedRoleDeletionStatus"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceLinkedRoleDeletionStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetServiceLinkedRoleDeletionStatus(request *GetServiceLinkedRoleDeletionStatusRequest) (_result *GetServiceLinkedRoleDeletionStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServiceLinkedRoleDeletionStatusResponse{}
	_body, _err := client.GetServiceLinkedRoleDeletionStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * >
 * *   An account can be used to enable a resource directory only after it passes enterprise real-name verification. An account that only passed individual real-name verification cannot be used to enable a resource directory.
 * *   We recommend that you only use the enterprise management account as the administrator of the resource directory. Do not use the enterprise management account to purchase cloud resources.
 *
 * @param request InitResourceDirectoryRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return InitResourceDirectoryResponse
 */
func (client *Client) InitResourceDirectoryWithOptions(runtime *util.RuntimeOptions) (_result *InitResourceDirectoryResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("InitResourceDirectory"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InitResourceDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * >
 * *   An account can be used to enable a resource directory only after it passes enterprise real-name verification. An account that only passed individual real-name verification cannot be used to enable a resource directory.
 * *   We recommend that you only use the enterprise management account as the administrator of the resource directory. Do not use the enterprise management account to purchase cloud resources.
 *
 * @return InitResourceDirectoryResponse
 */
func (client *Client) InitResourceDirectory() (_result *InitResourceDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitResourceDirectoryResponse{}
	_body, _err := client.InitResourceDirectoryWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to invite the account `someone@example.com` to join a resource directory.
 *
 * @param request InviteAccountToResourceDirectoryRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return InviteAccountToResourceDirectoryResponse
 */
func (client *Client) InviteAccountToResourceDirectoryWithOptions(request *InviteAccountToResourceDirectoryRequest, runtime *util.RuntimeOptions) (_result *InviteAccountToResourceDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Note)) {
		query["Note"] = request.Note
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TargetEntity)) {
		query["TargetEntity"] = request.TargetEntity
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InviteAccountToResourceDirectory"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InviteAccountToResourceDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to invite the account `someone@example.com` to join a resource directory.
 *
 * @param request InviteAccountToResourceDirectoryRequest
 * @return InviteAccountToResourceDirectoryResponse
 */
func (client *Client) InviteAccountToResourceDirectory(request *InviteAccountToResourceDirectoryRequest) (_result *InviteAccountToResourceDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InviteAccountToResourceDirectoryResponse{}
	_body, _err := client.InviteAccountToResourceDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can use only the management account of a resource directory or a delegated administrator account of a trusted service to call this operation.
 *
 * @param request ListAccountsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListAccountsResponse
 */
func (client *Client) ListAccountsWithOptions(request *ListAccountsRequest, runtime *util.RuntimeOptions) (_result *ListAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncludeTags)) {
		query["IncludeTags"] = request.IncludeTags
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAccounts"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can use only the management account of a resource directory or a delegated administrator account of a trusted service to call this operation.
 *
 * @param request ListAccountsRequest
 * @return ListAccountsResponse
 */
func (client *Client) ListAccounts(request *ListAccountsRequest) (_result *ListAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAccountsResponse{}
	_body, _err := client.ListAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAccountsForParentWithOptions(request *ListAccountsForParentRequest, runtime *util.RuntimeOptions) (_result *ListAccountsForParentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncludeTags)) {
		query["IncludeTags"] = request.IncludeTags
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ParentFolderId)) {
		query["ParentFolderId"] = request.ParentFolderId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryKeyword)) {
		query["QueryKeyword"] = request.QueryKeyword
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAccountsForParent"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAccountsForParentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAccountsForParent(request *ListAccountsForParentRequest) (_result *ListAccountsForParentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAccountsForParentResponse{}
	_body, _err := client.ListAccountsForParentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAncestorsWithOptions(request *ListAncestorsRequest, runtime *util.RuntimeOptions) (_result *ListAncestorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChildId)) {
		query["ChildId"] = request.ChildId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAncestors"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAncestorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAncestors(request *ListAncestorsRequest) (_result *ListAncestorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAncestorsResponse{}
	_body, _err := client.ListAncestorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAssociatedTransferSettingWithOptions(runtime *util.RuntimeOptions) (_result *ListAssociatedTransferSettingResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListAssociatedTransferSetting"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAssociatedTransferSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAssociatedTransferSetting() (_result *ListAssociatedTransferSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAssociatedTransferSettingResponse{}
	_body, _err := client.ListAssociatedTransferSettingWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the system access control policies within a resource directory. The response shows that the resource directory has only one system access control policy. The policy is named `FullAliyunAccess`.
 *
 * @param request ListControlPoliciesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListControlPoliciesResponse
 */
func (client *Client) ListControlPoliciesWithOptions(request *ListControlPoliciesRequest, runtime *util.RuntimeOptions) (_result *ListControlPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyType)) {
		query["PolicyType"] = request.PolicyType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListControlPolicies"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListControlPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the system access control policies within a resource directory. The response shows that the resource directory has only one system access control policy. The policy is named `FullAliyunAccess`.
 *
 * @param request ListControlPoliciesRequest
 * @return ListControlPoliciesResponse
 */
func (client *Client) ListControlPolicies(request *ListControlPoliciesRequest) (_result *ListControlPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListControlPoliciesResponse{}
	_body, _err := client.ListControlPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the access control policies that are attached to the folder `fd-ZDNPiT****`.
 *
 * @param request ListControlPolicyAttachmentsForTargetRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListControlPolicyAttachmentsForTargetResponse
 */
func (client *Client) ListControlPolicyAttachmentsForTargetWithOptions(request *ListControlPolicyAttachmentsForTargetRequest, runtime *util.RuntimeOptions) (_result *ListControlPolicyAttachmentsForTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["TargetId"] = request.TargetId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListControlPolicyAttachmentsForTarget"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListControlPolicyAttachmentsForTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the access control policies that are attached to the folder `fd-ZDNPiT****`.
 *
 * @param request ListControlPolicyAttachmentsForTargetRequest
 * @return ListControlPolicyAttachmentsForTargetResponse
 */
func (client *Client) ListControlPolicyAttachmentsForTarget(request *ListControlPolicyAttachmentsForTargetRequest) (_result *ListControlPolicyAttachmentsForTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListControlPolicyAttachmentsForTargetResponse{}
	_body, _err := client.ListControlPolicyAttachmentsForTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query all delegated administrator accounts in a resource directory. The response shows that two delegated administrator accounts for Cloud Firewall exist in the resource directory.
 *
 * @param request ListDelegatedAdministratorsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListDelegatedAdministratorsResponse
 */
func (client *Client) ListDelegatedAdministratorsWithOptions(request *ListDelegatedAdministratorsRequest, runtime *util.RuntimeOptions) (_result *ListDelegatedAdministratorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ServicePrincipal)) {
		query["ServicePrincipal"] = request.ServicePrincipal
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDelegatedAdministrators"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDelegatedAdministratorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query all delegated administrator accounts in a resource directory. The response shows that two delegated administrator accounts for Cloud Firewall exist in the resource directory.
 *
 * @param request ListDelegatedAdministratorsRequest
 * @return ListDelegatedAdministratorsResponse
 */
func (client *Client) ListDelegatedAdministrators(request *ListDelegatedAdministratorsRequest) (_result *ListDelegatedAdministratorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDelegatedAdministratorsResponse{}
	_body, _err := client.ListDelegatedAdministratorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the trusted services for which the member `138660628348****` is specified as a delegated administrator account. The response shows that the member is specified as a delegated administrator account of Cloud Firewall.
 *
 * @param request ListDelegatedServicesForAccountRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListDelegatedServicesForAccountResponse
 */
func (client *Client) ListDelegatedServicesForAccountWithOptions(request *ListDelegatedServicesForAccountRequest, runtime *util.RuntimeOptions) (_result *ListDelegatedServicesForAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDelegatedServicesForAccount"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDelegatedServicesForAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the trusted services for which the member `138660628348****` is specified as a delegated administrator account. The response shows that the member is specified as a delegated administrator account of Cloud Firewall.
 *
 * @param request ListDelegatedServicesForAccountRequest
 * @return ListDelegatedServicesForAccountResponse
 */
func (client *Client) ListDelegatedServicesForAccount(request *ListDelegatedServicesForAccountRequest) (_result *ListDelegatedServicesForAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDelegatedServicesForAccountResponse{}
	_body, _err := client.ListDelegatedServicesForAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * >  You can view the information of only the first-level subfolders of a folder.
 *
 * @param request ListFoldersForParentRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListFoldersForParentResponse
 */
func (client *Client) ListFoldersForParentWithOptions(request *ListFoldersForParentRequest, runtime *util.RuntimeOptions) (_result *ListFoldersForParentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ParentFolderId)) {
		query["ParentFolderId"] = request.ParentFolderId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryKeyword)) {
		query["QueryKeyword"] = request.QueryKeyword
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFoldersForParent"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFoldersForParentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * >  You can view the information of only the first-level subfolders of a folder.
 *
 * @param request ListFoldersForParentRequest
 * @return ListFoldersForParentResponse
 */
func (client *Client) ListFoldersForParent(request *ListFoldersForParentRequest) (_result *ListFoldersForParentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFoldersForParentResponse{}
	_body, _err := client.ListFoldersForParentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the invitations that are associated with the management account `172841235500****`. The response shows that two invitations are associated with the management account.
 *
 * @param request ListHandshakesForAccountRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListHandshakesForAccountResponse
 */
func (client *Client) ListHandshakesForAccountWithOptions(request *ListHandshakesForAccountRequest, runtime *util.RuntimeOptions) (_result *ListHandshakesForAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHandshakesForAccount"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHandshakesForAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the invitations that are associated with the management account `172841235500****`. The response shows that two invitations are associated with the management account.
 *
 * @param request ListHandshakesForAccountRequest
 * @return ListHandshakesForAccountResponse
 */
func (client *Client) ListHandshakesForAccount(request *ListHandshakesForAccountRequest) (_result *ListHandshakesForAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHandshakesForAccountResponse{}
	_body, _err := client.ListHandshakesForAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHandshakesForResourceDirectoryWithOptions(request *ListHandshakesForResourceDirectoryRequest, runtime *util.RuntimeOptions) (_result *ListHandshakesForResourceDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHandshakesForResourceDirectory"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHandshakesForResourceDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHandshakesForResourceDirectory(request *ListHandshakesForResourceDirectoryRequest) (_result *ListHandshakesForResourceDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHandshakesForResourceDirectoryResponse{}
	_body, _err := client.ListHandshakesForResourceDirectoryWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyType)) {
		query["PolicyType"] = request.PolicyType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPolicies"),
		Version:     tea.String("2020-03-31"),
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

/**
 * You can view the following information:
 * *   Policy attachment records under an Alibaba Cloud account or a resource group
 * *   Policies attached to RAM users, RAM user groups, or RAM roles
 * *   RAM users, RAM user groups, or RAM roles to which policies are attached under an Alibaba Cloud account or a resource group
 *
 * @param request ListPolicyAttachmentsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListPolicyAttachmentsResponse
 */
func (client *Client) ListPolicyAttachmentsWithOptions(request *ListPolicyAttachmentsRequest, runtime *util.RuntimeOptions) (_result *ListPolicyAttachmentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyType)) {
		query["PolicyType"] = request.PolicyType
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalName)) {
		query["PrincipalName"] = request.PrincipalName
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalType)) {
		query["PrincipalType"] = request.PrincipalType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPolicyAttachments"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPolicyAttachmentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can view the following information:
 * *   Policy attachment records under an Alibaba Cloud account or a resource group
 * *   Policies attached to RAM users, RAM user groups, or RAM roles
 * *   RAM users, RAM user groups, or RAM roles to which policies are attached under an Alibaba Cloud account or a resource group
 *
 * @param request ListPolicyAttachmentsRequest
 * @return ListPolicyAttachmentsResponse
 */
func (client *Client) ListPolicyAttachments(request *ListPolicyAttachmentsRequest) (_result *ListPolicyAttachmentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPolicyAttachmentsResponse{}
	_body, _err := client.ListPolicyAttachmentsWithOptions(request, runtime)
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
		Version:     tea.String("2020-03-31"),
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

/**
 * You can call this API operation to query all resource groups within the current account. You can also call this API operation to query a specific resource group based on the status, ID, identifier, or display name of the resource group.
 * This topic provides an example on how to call the API operation to query the basic information about the resource groups `rg-1hSBH2****` and `rg-9gLOoK****` within the current account.
 *
 * @param request ListResourceGroupsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListResourceGroupsResponse
 */
func (client *Client) ListResourceGroupsWithOptions(request *ListResourceGroupsRequest, runtime *util.RuntimeOptions) (_result *ListResourceGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeTags)) {
		query["IncludeTags"] = request.IncludeTags
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupIds)) {
		query["ResourceGroupIds"] = request.ResourceGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceGroups"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to query all resource groups within the current account. You can also call this API operation to query a specific resource group based on the status, ID, identifier, or display name of the resource group.
 * This topic provides an example on how to call the API operation to query the basic information about the resource groups `rg-1hSBH2****` and `rg-9gLOoK****` within the current account.
 *
 * @param request ListResourceGroupsRequest
 * @return ListResourceGroupsResponse
 */
func (client *Client) ListResourceGroups(request *ListResourceGroupsRequest) (_result *ListResourceGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourceGroupsResponse{}
	_body, _err := client.ListResourceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * >  You can use a RAM role that is not associated with a session policy to call this API operation.
 * This topic provides an example on how to call the API operation to query resources that can be accessed by the current account in resource groups. The response shows that the current account can access only the Elastic Compute Service (ECS) instance `i-23v38****` in the resource group `rg-uPJpP****`.
 *
 * @param request ListResourcesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListResourcesResponse
 */
func (client *Client) ListResourcesWithOptions(request *ListResourcesRequest, runtime *util.RuntimeOptions) (_result *ListResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceTypes)) {
		query["ResourceTypes"] = request.ResourceTypes
	}

	if !tea.BoolValue(util.IsUnset(request.Service)) {
		query["Service"] = request.Service
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResources"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * >  You can use a RAM role that is not associated with a session policy to call this API operation.
 * This topic provides an example on how to call the API operation to query resources that can be accessed by the current account in resource groups. The response shows that the current account can access only the Elastic Compute Service (ECS) instance `i-23v38****` in the resource group `rg-uPJpP****`.
 *
 * @param request ListResourcesRequest
 * @return ListResourcesResponse
 */
func (client *Client) ListResources(request *ListResourcesRequest) (_result *ListResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourcesResponse{}
	_body, _err := client.ListResourcesWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRoles"),
		Version:     tea.String("2020-03-31"),
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

/**
 * This topic provides an example on how to call the API operation to query tag keys. The response shows that the custom tag key team exists.
 *
 * @param request ListTagKeysRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListTagKeysResponse
 */
func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyFilter)) {
		query["KeyFilter"] = request.KeyFilter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagKeys"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query tag keys. The response shows that the custom tag key team exists.
 *
 * @param request ListTagKeysRequest
 * @return ListTagKeysResponse
 */
func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the tags that are added to the resource group with an ID of `rg-aekz6bre2uq****`. The response shows that only the `k1:v1` tag is added to the resource group.
 *
 * @param request ListTagResourcesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListTagResourcesResponse
 */
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
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
		Version:     tea.String("2020-03-31"),
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

/**
 * This topic provides an example on how to call the API operation to query the tags that are added to the resource group with an ID of `rg-aekz6bre2uq****`. The response shows that only the `k1:v1` tag is added to the resource group.
 *
 * @param request ListTagResourcesRequest
 * @return ListTagResourcesResponse
 */
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

/**
 * This topic provides an example on how to call the API operation to query the tag values of the tag key k1. The response shows that the tag value of the tag key k1 is v1.
 *
 * @param request ListTagValuesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListTagValuesResponse
 */
func (client *Client) ListTagValuesWithOptions(request *ListTagValuesRequest, runtime *util.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	if !tea.BoolValue(util.IsUnset(request.ValueFilter)) {
		query["ValueFilter"] = request.ValueFilter
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagValues"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagValuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the tag values of the tag key k1. The response shows that the tag value of the tag key k1 is v1.
 *
 * @param request ListTagValuesRequest
 * @return ListTagValuesResponse
 */
func (client *Client) ListTagValues(request *ListTagValuesRequest) (_result *ListTagValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagValuesResponse{}
	_body, _err := client.ListTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this example, the folders or member accounts to which the control policy `cp-jExXAqIYkwHN****` is attached are queried. The returned result shows that the control policy is attached to the folder `fd-ZDNPiT****`.
 *
 * @param request ListTargetAttachmentsForControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListTargetAttachmentsForControlPolicyResponse
 */
func (client *Client) ListTargetAttachmentsForControlPolicyWithOptions(request *ListTargetAttachmentsForControlPolicyRequest, runtime *util.RuntimeOptions) (_result *ListTargetAttachmentsForControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTargetAttachmentsForControlPolicy"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTargetAttachmentsForControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, the folders or member accounts to which the control policy `cp-jExXAqIYkwHN****` is attached are queried. The returned result shows that the control policy is attached to the folder `fd-ZDNPiT****`.
 *
 * @param request ListTargetAttachmentsForControlPolicyRequest
 * @return ListTargetAttachmentsForControlPolicyResponse
 */
func (client *Client) ListTargetAttachmentsForControlPolicy(request *ListTargetAttachmentsForControlPolicyRequest) (_result *ListTargetAttachmentsForControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTargetAttachmentsForControlPolicyResponse{}
	_body, _err := client.ListTargetAttachmentsForControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * >  Only an enterprise management account or delegated administrator account can be used to call this operation.
 * In this example, the trusted services that are enabled within an enterprise management account are queried. The returned result shows that the trusted services Cloud Config and ActionTrail are enabled within the enterprise management account.
 *
 * @param request ListTrustedServiceStatusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListTrustedServiceStatusResponse
 */
func (client *Client) ListTrustedServiceStatusWithOptions(request *ListTrustedServiceStatusRequest, runtime *util.RuntimeOptions) (_result *ListTrustedServiceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdminAccountId)) {
		query["AdminAccountId"] = request.AdminAccountId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTrustedServiceStatus"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTrustedServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * >  Only an enterprise management account or delegated administrator account can be used to call this operation.
 * In this example, the trusted services that are enabled within an enterprise management account are queried. The returned result shows that the trusted services Cloud Config and ActionTrail are enabled within the enterprise management account.
 *
 * @param request ListTrustedServiceStatusRequest
 * @return ListTrustedServiceStatusResponse
 */
func (client *Client) ListTrustedServiceStatus(request *ListTrustedServiceStatusRequest) (_result *ListTrustedServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTrustedServiceStatusResponse{}
	_body, _err := client.ListTrustedServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MoveAccountWithOptions(request *MoveAccountRequest, runtime *util.RuntimeOptions) (_result *MoveAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationFolderId)) {
		query["DestinationFolderId"] = request.DestinationFolderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveAccount"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MoveAccount(request *MoveAccountRequest) (_result *MoveAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveAccountResponse{}
	_body, _err := client.MoveAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * For more information about Alibaba Cloud services whose resources can be moved between resource groups, see the **Supported by the API** column in [Alibaba Cloud services that support resource groups](~~94479~~).
 * In this example, two virtual private clouds (VPCs) `vpc-bp1sig0mjktx5ewx1****` and `vpc-bp1visxm225pv49dz****` that reside in different regions and belong to different resource groups are moved to the resource group `rg-aekzmeobk5w****`.
 *
 * @param request MoveResourcesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return MoveResourcesResponse
 */
func (client *Client) MoveResourcesWithOptions(request *MoveResourcesRequest, runtime *util.RuntimeOptions) (_result *MoveResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Resources)) {
		query["Resources"] = request.Resources
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveResources"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * For more information about Alibaba Cloud services whose resources can be moved between resource groups, see the **Supported by the API** column in [Alibaba Cloud services that support resource groups](~~94479~~).
 * In this example, two virtual private clouds (VPCs) `vpc-bp1sig0mjktx5ewx1****` and `vpc-bp1visxm225pv49dz****` that reside in different regions and belong to different resource groups are moved to the resource group `rg-aekzmeobk5w****`.
 *
 * @param request MoveResourcesRequest
 * @return MoveResourcesResponse
 */
func (client *Client) MoveResources(request *MoveResourcesRequest) (_result *MoveResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveResourcesResponse{}
	_body, _err := client.MoveResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PromoteResourceAccountWithOptions(request *PromoteResourceAccountRequest, runtime *util.RuntimeOptions) (_result *PromoteResourceAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PromoteResourceAccount"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PromoteResourceAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PromoteResourceAccount(request *PromoteResourceAccountRequest) (_result *PromoteResourceAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PromoteResourceAccountResponse{}
	_body, _err := client.PromoteResourceAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The delegated administrator account can be used to access the information of the resource directory and view the structure and members of the resource directory. The delegated administrator account can also be used to perform service-related management operations in the trusted service on behalf of the management account of the resource directory.
 * When you call this operation, you must take note of the following limits:
 * *   Only some trusted services support delegated administrator accounts. For more information, see [Supported trusted services](~~208133~~).
 * *   Only the management account of a resource directory or an authorized RAM user or RAM role of the management account can be used to call this operation.
 * *   The number of delegated administrator accounts that are allowed for a trusted service is defined by the trusted service.
 * This topic provides an example on how to call the API operation to specify the member `181761095690****` as a delegated administrator account of Cloud Firewall.
 *
 * @param request RegisterDelegatedAdministratorRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RegisterDelegatedAdministratorResponse
 */
func (client *Client) RegisterDelegatedAdministratorWithOptions(request *RegisterDelegatedAdministratorRequest, runtime *util.RuntimeOptions) (_result *RegisterDelegatedAdministratorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.ServicePrincipal)) {
		query["ServicePrincipal"] = request.ServicePrincipal
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RegisterDelegatedAdministrator"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RegisterDelegatedAdministratorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The delegated administrator account can be used to access the information of the resource directory and view the structure and members of the resource directory. The delegated administrator account can also be used to perform service-related management operations in the trusted service on behalf of the management account of the resource directory.
 * When you call this operation, you must take note of the following limits:
 * *   Only some trusted services support delegated administrator accounts. For more information, see [Supported trusted services](~~208133~~).
 * *   Only the management account of a resource directory or an authorized RAM user or RAM role of the management account can be used to call this operation.
 * *   The number of delegated administrator accounts that are allowed for a trusted service is defined by the trusted service.
 * This topic provides an example on how to call the API operation to specify the member `181761095690****` as a delegated administrator account of Cloud Firewall.
 *
 * @param request RegisterDelegatedAdministratorRequest
 * @return RegisterDelegatedAdministratorResponse
 */
func (client *Client) RegisterDelegatedAdministrator(request *RegisterDelegatedAdministratorRequest) (_result *RegisterDelegatedAdministratorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterDelegatedAdministratorResponse{}
	_body, _err := client.RegisterDelegatedAdministratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to remove the member `177242285274****` from a resource directory.
 *
 * @param request RemoveCloudAccountRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RemoveCloudAccountResponse
 */
func (client *Client) RemoveCloudAccountWithOptions(request *RemoveCloudAccountRequest, runtime *util.RuntimeOptions) (_result *RemoveCloudAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveCloudAccount"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveCloudAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to remove the member `177242285274****` from a resource directory.
 *
 * @param request RemoveCloudAccountRequest
 * @return RemoveCloudAccountResponse
 */
func (client *Client) RemoveCloudAccount(request *RemoveCloudAccountRequest) (_result *RemoveCloudAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveCloudAccountResponse{}
	_body, _err := client.RemoveCloudAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResendCreateCloudAccountEmailWithOptions(request *ResendCreateCloudAccountEmailRequest, runtime *util.RuntimeOptions) (_result *ResendCreateCloudAccountEmailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RecordId)) {
		query["RecordId"] = request.RecordId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResendCreateCloudAccountEmail"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResendCreateCloudAccountEmailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResendCreateCloudAccountEmail(request *ResendCreateCloudAccountEmailRequest) (_result *ResendCreateCloudAccountEmailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResendCreateCloudAccountEmailResponse{}
	_body, _err := client.ResendCreateCloudAccountEmailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResendPromoteResourceAccountEmailWithOptions(request *ResendPromoteResourceAccountEmailRequest, runtime *util.RuntimeOptions) (_result *ResendPromoteResourceAccountEmailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RecordId)) {
		query["RecordId"] = request.RecordId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResendPromoteResourceAccountEmail"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResendPromoteResourceAccountEmailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResendPromoteResourceAccountEmail(request *ResendPromoteResourceAccountEmailRequest) (_result *ResendPromoteResourceAccountEmailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResendPromoteResourceAccountEmailResponse{}
	_body, _err := client.ResendPromoteResourceAccountEmailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RetryChangeAccountEmailWithOptions(request *RetryChangeAccountEmailRequest, runtime *util.RuntimeOptions) (_result *RetryChangeAccountEmailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RetryChangeAccountEmail"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RetryChangeAccountEmailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RetryChangeAccountEmail(request *RetryChangeAccountEmailRequest) (_result *RetryChangeAccountEmailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RetryChangeAccountEmailResponse{}
	_body, _err := client.RetryChangeAccountEmailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * To ensure that the system can record the operators of management operations, you must use a RAM user or RAM role to which the AliyunResourceDirectoryFullAccess policy is attached within the management account of your resource directory to call this operation.
 * In this example, a verification code is sent to the mobile phone number that you want to bind to the resource account `138660628348****`.
 *
 * @param request SendVerificationCodeForBindSecureMobilePhoneRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SendVerificationCodeForBindSecureMobilePhoneResponse
 */
func (client *Client) SendVerificationCodeForBindSecureMobilePhoneWithOptions(request *SendVerificationCodeForBindSecureMobilePhoneRequest, runtime *util.RuntimeOptions) (_result *SendVerificationCodeForBindSecureMobilePhoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.SecureMobilePhone)) {
		query["SecureMobilePhone"] = request.SecureMobilePhone
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendVerificationCodeForBindSecureMobilePhone"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendVerificationCodeForBindSecureMobilePhoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * To ensure that the system can record the operators of management operations, you must use a RAM user or RAM role to which the AliyunResourceDirectoryFullAccess policy is attached within the management account of your resource directory to call this operation.
 * In this example, a verification code is sent to the mobile phone number that you want to bind to the resource account `138660628348****`.
 *
 * @param request SendVerificationCodeForBindSecureMobilePhoneRequest
 * @return SendVerificationCodeForBindSecureMobilePhoneResponse
 */
func (client *Client) SendVerificationCodeForBindSecureMobilePhone(request *SendVerificationCodeForBindSecureMobilePhoneRequest) (_result *SendVerificationCodeForBindSecureMobilePhoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendVerificationCodeForBindSecureMobilePhoneResponse{}
	_body, _err := client.SendVerificationCodeForBindSecureMobilePhoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Each Alibaba Cloud account can be used to send a maximum of 100 verification codes per day.
 *
 * @param request SendVerificationCodeForEnableRDRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SendVerificationCodeForEnableRDResponse
 */
func (client *Client) SendVerificationCodeForEnableRDWithOptions(request *SendVerificationCodeForEnableRDRequest, runtime *util.RuntimeOptions) (_result *SendVerificationCodeForEnableRDResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SecureMobilePhone)) {
		query["SecureMobilePhone"] = request.SecureMobilePhone
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendVerificationCodeForEnableRD"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendVerificationCodeForEnableRDResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Each Alibaba Cloud account can be used to send a maximum of 100 verification codes per day.
 *
 * @param request SendVerificationCodeForEnableRDRequest
 * @return SendVerificationCodeForEnableRDResponse
 */
func (client *Client) SendVerificationCodeForEnableRD(request *SendVerificationCodeForEnableRDRequest) (_result *SendVerificationCodeForEnableRDResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendVerificationCodeForEnableRDResponse{}
	_body, _err := client.SendVerificationCodeForEnableRDWithOptions(request, runtime)
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
		Version:     tea.String("2020-03-31"),
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

/**
 * Members of the resource account type can be deleted only after the member deletion feature is enabled.
 *
 * @param request SetMemberDeletionPermissionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SetMemberDeletionPermissionResponse
 */
func (client *Client) SetMemberDeletionPermissionWithOptions(request *SetMemberDeletionPermissionRequest, runtime *util.RuntimeOptions) (_result *SetMemberDeletionPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetMemberDeletionPermission"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetMemberDeletionPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Members of the resource account type can be deleted only after the member deletion feature is enabled.
 *
 * @param request SetMemberDeletionPermissionRequest
 * @return SetMemberDeletionPermissionResponse
 */
func (client *Client) SetMemberDeletionPermission(request *SetMemberDeletionPermissionRequest) (_result *SetMemberDeletionPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetMemberDeletionPermissionResponse{}
	_body, _err := client.SetMemberDeletionPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to add the tag `k1:v1` to the resource group with an ID of `rg-aekz6bre2uq****`.
 *
 * @param request TagResourcesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return TagResourcesResponse
 */
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Version:     tea.String("2020-03-31"),
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

/**
 * This topic provides an example on how to call the API operation to add the tag `k1:v1` to the resource group with an ID of `rg-aekz6bre2uq****`.
 *
 * @param request TagResourcesRequest
 * @return TagResourcesResponse
 */
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

/**
 * This topic provides an example on how to call the API operation to remove the tag whose tag key is `k1` from the resource group whose ID is `rg-aek2dpwyrfr****`.
 *
 * @param request UntagResourcesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UntagResourcesResponse
 */
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
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
		Version:     tea.String("2020-03-31"),
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

/**
 * This topic provides an example on how to call the API operation to remove the tag whose tag key is `k1` from the resource group whose ID is `rg-aek2dpwyrfr****`.
 *
 * @param request UntagResourcesRequest
 * @return UntagResourcesResponse
 */
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

/**
 * *   To ensure that the system can record the operators of management operations, you must use a RAM user or RAM role to which the AliyunResourceDirectoryFullAccess policy is attached within the management account of your resource directory to call this operation.
 * *   Before you switch the type of a member from resource account to cloud account, make sure that specific conditions are met. For more information about the conditions, see [Switch a resource account to a cloud account](~~111233~~).
 * *   Before you switch the type of a member from cloud account to resource account, make sure that specific conditions are met. For more information about the conditions, see [Switch a cloud account to a resource account](~~209980~~).
 * This example provides an example on how to call the API operation to change the display name of the member `12323344****` to `admin`.
 *
 * @param request UpdateAccountRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateAccountResponse
 */
func (client *Client) UpdateAccountWithOptions(request *UpdateAccountRequest, runtime *util.RuntimeOptions) (_result *UpdateAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.NewAccountType)) {
		query["NewAccountType"] = request.NewAccountType
	}

	if !tea.BoolValue(util.IsUnset(request.NewDisplayName)) {
		query["NewDisplayName"] = request.NewDisplayName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAccount"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   To ensure that the system can record the operators of management operations, you must use a RAM user or RAM role to which the AliyunResourceDirectoryFullAccess policy is attached within the management account of your resource directory to call this operation.
 * *   Before you switch the type of a member from resource account to cloud account, make sure that specific conditions are met. For more information about the conditions, see [Switch a resource account to a cloud account](~~111233~~).
 * *   Before you switch the type of a member from cloud account to resource account, make sure that specific conditions are met. For more information about the conditions, see [Switch a cloud account to a resource account](~~209980~~).
 * This example provides an example on how to call the API operation to change the display name of the member `12323344****` to `admin`.
 *
 * @param request UpdateAccountRequest
 * @return UpdateAccountResponse
 */
func (client *Client) UpdateAccount(request *UpdateAccountRequest) (_result *UpdateAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAccountResponse{}
	_body, _err := client.UpdateAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * For information about the resources that support the Transfer Associated Resources feature, see [Use the Transfer Associated Resources feature](~~2639129~~).
 *
 * @param request UpdateAssociatedTransferSettingRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateAssociatedTransferSettingResponse
 */
func (client *Client) UpdateAssociatedTransferSettingWithOptions(request *UpdateAssociatedTransferSettingRequest, runtime *util.RuntimeOptions) (_result *UpdateAssociatedTransferSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnableExistingResourcesTransfer)) {
		query["EnableExistingResourcesTransfer"] = request.EnableExistingResourcesTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.RuleSettings)) {
		query["RuleSettings"] = request.RuleSettings
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAssociatedTransferSetting"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAssociatedTransferSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * For information about the resources that support the Transfer Associated Resources feature, see [Use the Transfer Associated Resources feature](~~2639129~~).
 *
 * @param request UpdateAssociatedTransferSettingRequest
 * @return UpdateAssociatedTransferSettingResponse
 */
func (client *Client) UpdateAssociatedTransferSetting(request *UpdateAssociatedTransferSettingRequest) (_result *UpdateAssociatedTransferSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAssociatedTransferSettingResponse{}
	_body, _err := client.UpdateAssociatedTransferSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this example, the name of the access control policy whose ID is `cp-jExXAqIYkwHN****` is changed to `NewControlPolicy`.
 *
 * @param request UpdateControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateControlPolicyResponse
 */
func (client *Client) UpdateControlPolicyWithOptions(request *UpdateControlPolicyRequest, runtime *util.RuntimeOptions) (_result *UpdateControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewDescription)) {
		query["NewDescription"] = request.NewDescription
	}

	if !tea.BoolValue(util.IsUnset(request.NewPolicyDocument)) {
		query["NewPolicyDocument"] = request.NewPolicyDocument
	}

	if !tea.BoolValue(util.IsUnset(request.NewPolicyName)) {
		query["NewPolicyName"] = request.NewPolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateControlPolicy"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, the name of the access control policy whose ID is `cp-jExXAqIYkwHN****` is changed to `NewControlPolicy`.
 *
 * @param request UpdateControlPolicyRequest
 * @return UpdateControlPolicyResponse
 */
func (client *Client) UpdateControlPolicy(request *UpdateControlPolicyRequest) (_result *UpdateControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateControlPolicyResponse{}
	_body, _err := client.UpdateControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFolderWithOptions(request *UpdateFolderRequest, runtime *util.RuntimeOptions) (_result *UpdateFolderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FolderId)) {
		query["FolderId"] = request.FolderId
	}

	if !tea.BoolValue(util.IsUnset(request.NewFolderName)) {
		query["NewFolderName"] = request.NewFolderName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFolder"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFolderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFolder(request *UpdateFolderRequest) (_result *UpdateFolderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFolderResponse{}
	_body, _err := client.UpdateFolderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this example, the display name of the resource group `rg-9gLOoK****` is changed to `project`.
 *
 * @param request UpdateResourceGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateResourceGroupResponse
 */
func (client *Client) UpdateResourceGroupWithOptions(request *UpdateResourceGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewDisplayName)) {
		query["NewDisplayName"] = request.NewDisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateResourceGroup"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, the display name of the resource group `rg-9gLOoK****` is changed to `project`.
 *
 * @param request UpdateResourceGroupRequest
 * @return UpdateResourceGroupResponse
 */
func (client *Client) UpdateResourceGroup(request *UpdateResourceGroupRequest) (_result *UpdateResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateResourceGroupResponse{}
	_body, _err := client.UpdateResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this example, the description of the RAM role `ECSAdmin` is updated to `ECS administrator`.
 *
 * @param request UpdateRoleRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateRoleResponse
 */
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
		Version:     tea.String("2020-03-31"),
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

/**
 * In this example, the description of the RAM role `ECSAdmin` is updated to `ECS administrator`.
 *
 * @param request UpdateRoleRequest
 * @return UpdateRoleResponse
 */
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
