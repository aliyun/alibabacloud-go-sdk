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

type AcceptHandshakeRequest struct {
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
	Handshake *AcceptHandshakeResponseBodyHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpireTime          *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	HandshakeId         *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
	MasterAccountId     *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	MasterAccountName   *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Note                *string `json:"Note,omitempty" xml:"Note,omitempty"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetEntity        *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty"`
	TargetType          *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AcceptHandshakeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PolicyName      *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType      *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	PrincipalName   *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	PrincipalType   *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AccountId         *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	SecureMobilePhone *string `json:"SecureMobilePhone,omitempty" xml:"SecureMobilePhone,omitempty"`
	VerificationCode  *string `json:"VerificationCode,omitempty" xml:"VerificationCode,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BindSecureMobilePhoneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CancelCreateCloudAccountRequest struct {
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelCreateCloudAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Handshake *CancelHandshakeResponseBodyHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpireTime          *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	HandshakeId         *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
	MasterAccountId     *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	MasterAccountName   *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Note                *string `json:"Note,omitempty" xml:"Note,omitempty"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetEntity        *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty"`
	TargetType          *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelHandshakeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelPromoteResourceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CreateCloudAccountRequest struct {
	DisplayName    *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email          *string `json:"Email,omitempty" xml:"Email,omitempty"`
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
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
	Account   *CreateCloudAccountResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccountId           *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AccountName         *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DisplayName         *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	FolderId            *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	JoinMethod          *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	RecordId            *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCloudAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EffectScope    *string `json:"EffectScope,omitempty" xml:"EffectScope,omitempty"`
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	PolicyName     *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
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
	ControlPolicy *CreateControlPolicyResponseBodyControlPolicy `json:"ControlPolicy,omitempty" xml:"ControlPolicy,omitempty" type:"Struct"`
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AttachmentCount *string `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty"`
	CreateDate      *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EffectScope     *string `json:"EffectScope,omitempty" xml:"EffectScope,omitempty"`
	PolicyId        *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	PolicyName      *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType      *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	UpdateDate      *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	FolderName     *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
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
	Folder    *CreateFolderResponseBodyFolder `json:"Folder,omitempty" xml:"Folder,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FolderId       *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	FolderName     *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFolderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CreateResourceAccountRequest struct {
	AccountNamePrefix *string                            `json:"AccountNamePrefix,omitempty" xml:"AccountNamePrefix,omitempty"`
	DisplayName       *string                            `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	ParentFolderId    *string                            `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	PayerAccountId    *string                            `json:"PayerAccountId,omitempty" xml:"PayerAccountId,omitempty"`
	Tag               []*CreateResourceAccountRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *CreateResourceAccountRequest) SetTag(v []*CreateResourceAccountRequestTag) *CreateResourceAccountRequest {
	s.Tag = v
	return s
}

type CreateResourceAccountRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Account   *CreateResourceAccountResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccountId           *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AccountName         *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DisplayName         *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	FolderId            *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	JoinMethod          *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	JoinTime            *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateResourceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

type CreateResourceGroupResponseBody struct {
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccountId      *string                                                     `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	CreateDate     *string                                                     `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DisplayName    *string                                                     `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Id             *string                                                     `json:"Id,omitempty" xml:"Id,omitempty"`
	Name           *string                                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionStatuses *CreateResourceGroupResponseBodyResourceGroupRegionStatuses `json:"RegionStatuses,omitempty" xml:"RegionStatuses,omitempty" type:"Struct"`
	Status         *string                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
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
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RolePrincipalName        *string `json:"RolePrincipalName,omitempty" xml:"RolePrincipalName,omitempty"`
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

type CreateServiceLinkedRoleRequest struct {
	CustomSuffix *string `json:"CustomSuffix,omitempty" xml:"CustomSuffix,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ServiceName  *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
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
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Role      *CreateServiceLinkedRoleResponseBodyRole `json:"Role,omitempty" xml:"Role,omitempty" type:"Struct"`
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
	Arn                      *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	AssumeRolePolicyDocument *string `json:"AssumeRolePolicyDocument,omitempty" xml:"AssumeRolePolicyDocument,omitempty"`
	CreateDate               *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IsServiceLinkedRole      *bool   `json:"IsServiceLinkedRole,omitempty" xml:"IsServiceLinkedRole,omitempty"`
	RoleId                   *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	RoleName                 *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	RolePrincipalName        *string `json:"RolePrincipalName,omitempty" xml:"RolePrincipalName,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Handshake *DeclineHandshakeResponseBodyHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpireTime          *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	HandshakeId         *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
	MasterAccountId     *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	MasterAccountName   *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Note                *string `json:"Note,omitempty" xml:"Note,omitempty"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetEntity        *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty"`
	TargetType          *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeclineHandshakeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DeleteControlPolicyRequest struct {
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteFolderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DeleteResourceGroupRequest struct {
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
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccountId      *string                                                     `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	CreateDate     *string                                                     `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DisplayName    *string                                                     `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Id             *string                                                     `json:"Id,omitempty" xml:"Id,omitempty"`
	Name           *string                                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionStatuses *DeleteResourceGroupResponseBodyResourceGroupRegionStatuses `json:"RegionStatuses,omitempty" xml:"RegionStatuses,omitempty" type:"Struct"`
	Status         *string                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
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
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DeleteServiceLinkedRoleRequest struct {
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
	DeletionTaskId *string `json:"DeletionTaskId,omitempty" xml:"DeletionTaskId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AccountId        *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
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
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeregisterDelegatedAdministratorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DestroyResourceDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PolicyName      *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType      *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	PrincipalName   *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	PrincipalType   *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DisableControlPolicyResponseBody struct {
	EnablementStatus *string `json:"EnablementStatus,omitempty" xml:"EnablementStatus,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type EnableControlPolicyResponseBody struct {
	EnablementStatus *string `json:"EnablementStatus,omitempty" xml:"EnablementStatus,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	EnableMode          *string `json:"EnableMode,omitempty" xml:"EnableMode,omitempty"`
	MAName              *string `json:"MAName,omitempty" xml:"MAName,omitempty"`
	MASecureMobilePhone *string `json:"MASecureMobilePhone,omitempty" xml:"MASecureMobilePhone,omitempty"`
	VerificationCode    *string `json:"VerificationCode,omitempty" xml:"VerificationCode,omitempty"`
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
	RequestId         *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	MasterAccountId     *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	MasterAccountName   *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	RootFolderId        *string `json:"RootFolderId,omitempty" xml:"RootFolderId,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableResourceDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AccountId   *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	IncludeTags *bool   `json:"IncludeTags,omitempty" xml:"IncludeTags,omitempty"`
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
	Account   *GetAccountResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccountId             *string                              `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AccountName           *string                              `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DisplayName           *string                              `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	FolderId              *string                              `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	IdentityInformation   *string                              `json:"IdentityInformation,omitempty" xml:"IdentityInformation,omitempty"`
	JoinMethod            *string                              `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	JoinTime              *string                              `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	Location              *string                              `json:"Location,omitempty" xml:"Location,omitempty"`
	ModifyTime            *string                              `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	ResourceDirectoryId   *string                              `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	ResourceDirectoryPath *string                              `json:"ResourceDirectoryPath,omitempty" xml:"ResourceDirectoryPath,omitempty"`
	Status                *string                              `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                  []*GetAccountResponseBodyAccountTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Type                  *string                              `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type GetControlPolicyRequest struct {
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
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
	ControlPolicy *GetControlPolicyResponseBodyControlPolicy `json:"ControlPolicy,omitempty" xml:"ControlPolicy,omitempty" type:"Struct"`
	RequestId     *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AttachmentCount *string `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty"`
	CreateDate      *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EffectScope     *string `json:"EffectScope,omitempty" xml:"EffectScope,omitempty"`
	PolicyDocument  *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	PolicyId        *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	PolicyName      *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType      *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	UpdateDate      *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	EnablementStatus *string `json:"EnablementStatus,omitempty" xml:"EnablementStatus,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetControlPolicyEnablementStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Folder    *GetFolderResponseBodyFolder `json:"Folder,omitempty" xml:"Folder,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime            *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FolderId              *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	FolderName            *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
	ParentFolderId        *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFolderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Handshake *GetHandshakeResponseBodyHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime             *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpireTime             *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	HandshakeId            *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
	InvitedAccountRealName *string `json:"InvitedAccountRealName,omitempty" xml:"InvitedAccountRealName,omitempty"`
	MasterAccountId        *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	MasterAccountName      *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	MasterAccountRealName  *string `json:"MasterAccountRealName,omitempty" xml:"MasterAccountRealName,omitempty"`
	ModifyTime             *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Note                   *string `json:"Note,omitempty" xml:"Note,omitempty"`
	ResourceDirectoryId    *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetEntity           *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty"`
	TargetType             *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetHandshakeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PayerAccountId   *string `json:"PayerAccountId,omitempty" xml:"PayerAccountId,omitempty"`
	PayerAccountName *string `json:"PayerAccountName,omitempty" xml:"PayerAccountName,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPayerForAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Language   *string `json:"Language,omitempty" xml:"Language,omitempty"`
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
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
	Policy    *GetPolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type GetResourceDirectoryResponseBody struct {
	RequestId         *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ControlPolicyStatus  *string `json:"ControlPolicyStatus,omitempty" xml:"ControlPolicyStatus,omitempty"`
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	IdentityInformation  *string `json:"IdentityInformation,omitempty" xml:"IdentityInformation,omitempty"`
	MasterAccountId      *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	MasterAccountName    *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	MemberDeletionStatus *string `json:"MemberDeletionStatus,omitempty" xml:"MemberDeletionStatus,omitempty"`
	ResourceDirectoryId  *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	RootFolderId         *string `json:"RootFolderId,omitempty" xml:"RootFolderId,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetResourceDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	IncludeTags     *bool   `json:"IncludeTags,omitempty" xml:"IncludeTags,omitempty"`
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
	RequestId     *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccountId      *string                                                  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	CreateDate     *string                                                  `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DisplayName    *string                                                  `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Id             *string                                                  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name           *string                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionStatuses *GetResourceGroupResponseBodyResourceGroupRegionStatuses `json:"RegionStatuses,omitempty" xml:"RegionStatuses,omitempty" type:"Struct"`
	Status         *string                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags           *GetResourceGroupResponseBodyResourceGroupTags           `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
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
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type GetResourceGroupListAclModeResponseBody struct {
	Mode      *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetResourceGroupListAclModeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupListAclModeResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceGroupListAclModeResponseBody) SetMode(v string) *GetResourceGroupListAclModeResponseBody {
	s.Mode = &v
	return s
}

func (s *GetResourceGroupListAclModeResponseBody) SetRequestId(v string) *GetResourceGroupListAclModeResponseBody {
	s.RequestId = &v
	return s
}

type GetResourceGroupListAclModeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetResourceGroupListAclModeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResourceGroupListAclModeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupListAclModeResponse) GoString() string {
	return s.String()
}

func (s *GetResourceGroupListAclModeResponse) SetHeaders(v map[string]*string) *GetResourceGroupListAclModeResponse {
	s.Headers = v
	return s
}

func (s *GetResourceGroupListAclModeResponse) SetStatusCode(v int32) *GetResourceGroupListAclModeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceGroupListAclModeResponse) SetBody(v *GetResourceGroupListAclModeResponseBody) *GetResourceGroupListAclModeResponse {
	s.Body = v
	return s
}

type GetRoleRequest struct {
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
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
	Arn                      *string                                    `json:"Arn,omitempty" xml:"Arn,omitempty"`
	AssumeRolePolicyDocument *string                                    `json:"AssumeRolePolicyDocument,omitempty" xml:"AssumeRolePolicyDocument,omitempty"`
	CreateDate               *string                                    `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	Description              *string                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	IsServiceLinkedRole      *bool                                      `json:"IsServiceLinkedRole,omitempty" xml:"IsServiceLinkedRole,omitempty"`
	LatestDeletionTask       *GetRoleResponseBodyRoleLatestDeletionTask `json:"LatestDeletionTask,omitempty" xml:"LatestDeletionTask,omitempty" type:"Struct"`
	MaxSessionDuration       *int64                                     `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty"`
	RoleId                   *string                                    `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	RoleName                 *string                                    `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	RolePrincipalName        *string                                    `json:"RolePrincipalName,omitempty" xml:"RolePrincipalName,omitempty"`
	UpdateDate               *string                                    `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
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
	CreateDate     *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
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

type GetServiceLinkedRoleDeletionStatusRequest struct {
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
	Reason    *GetServiceLinkedRoleDeletionStatusResponseBodyReason `json:"Reason,omitempty" xml:"Reason,omitempty" type:"Struct"`
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string                                               `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Message    *string                                                         `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Region    *string                                                                           `json:"Region,omitempty" xml:"Region,omitempty"`
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
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetServiceLinkedRoleDeletionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId         *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	MasterAccountId     *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	MasterAccountName   *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	RootFolderId        *string `json:"RootFolderId,omitempty" xml:"RootFolderId,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InitResourceDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Note         *string                                       `json:"Note,omitempty" xml:"Note,omitempty"`
	Tag          []*InviteAccountToResourceDirectoryRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	TargetEntity *string                                       `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty"`
	TargetType   *string                                       `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Handshake *InviteAccountToResourceDirectoryResponseBodyHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" type:"Struct"`
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpireTime          *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	HandshakeId         *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
	MasterAccountId     *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	MasterAccountName   *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Note                *string `json:"Note,omitempty" xml:"Note,omitempty"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetEntity        *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty"`
	TargetType          *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
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
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InviteAccountToResourceDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	IncludeTags *bool                     `json:"IncludeTags,omitempty" xml:"IncludeTags,omitempty"`
	PageNumber  *int32                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Tag         []*ListAccountsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Accounts   *ListAccountsResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
	PageNumber *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	AccountId             *string                                      `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	DisplayName           *string                                      `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	FolderId              *string                                      `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	JoinMethod            *string                                      `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	JoinTime              *string                                      `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	ModifyTime            *string                                      `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	ResourceDirectoryId   *string                                      `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	ResourceDirectoryPath *string                                      `json:"ResourceDirectoryPath,omitempty" xml:"ResourceDirectoryPath,omitempty"`
	Status                *string                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                  *ListAccountsResponseBodyAccountsAccountTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	Type                  *string                                      `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	IncludeTags    *bool                              `json:"IncludeTags,omitempty" xml:"IncludeTags,omitempty"`
	PageNumber     *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ParentFolderId *string                            `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	QueryKeyword   *string                            `json:"QueryKeyword,omitempty" xml:"QueryKeyword,omitempty"`
	Tag            []*ListAccountsForParentRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Accounts   *ListAccountsForParentResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
	PageNumber *int32                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	AccountId           *string                                               `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	DisplayName         *string                                               `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	FolderId            *string                                               `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	JoinMethod          *string                                               `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	JoinTime            *string                                               `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	ModifyTime          *string                                               `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	ResourceDirectoryId *string                                               `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	Status              *string                                               `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                *ListAccountsForParentResponseBodyAccountsAccountTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	Type                *string                                               `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAccountsForParentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Folders   *ListAncestorsResponseBodyFolders `json:"Folders,omitempty" xml:"Folders,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FolderId   *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAncestorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListControlPoliciesRequest struct {
	Language   *string `json:"Language,omitempty" xml:"Language,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	ControlPolicies *ListControlPoliciesResponseBodyControlPolicies `json:"ControlPolicies,omitempty" xml:"ControlPolicies,omitempty" type:"Struct"`
	PageNumber      *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId       *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount      *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	AttachmentCount *string `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty"`
	CreateDate      *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EffectScope     *string `json:"EffectScope,omitempty" xml:"EffectScope,omitempty"`
	PolicyId        *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	PolicyName      *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType      *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	UpdateDate      *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListControlPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
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
	ControlPolicyAttachments *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachments `json:"ControlPolicyAttachments,omitempty" xml:"ControlPolicyAttachments,omitempty" type:"Struct"`
	RequestId                *string                                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AttachDate  *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EffectScope *string `json:"EffectScope,omitempty" xml:"EffectScope,omitempty"`
	PolicyId    *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	PolicyName  *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType  *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
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
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListControlPolicyAttachmentsForTargetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PageNumber       *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	Accounts   *ListDelegatedAdministratorsResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
	PageNumber *int64                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	AccountId             *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	DelegationEnabledTime *string `json:"DelegationEnabledTime,omitempty" xml:"DelegationEnabledTime,omitempty"`
	DisplayName           *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	JoinMethod            *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	ServicePrincipal      *string `json:"ServicePrincipal,omitempty" xml:"ServicePrincipal,omitempty"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDelegatedAdministratorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DelegatedServices *ListDelegatedServicesForAccountResponseBodyDelegatedServices `json:"DelegatedServices,omitempty" xml:"DelegatedServices,omitempty" type:"Struct"`
	RequestId         *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	DelegationEnabledTime *string `json:"DelegationEnabledTime,omitempty" xml:"DelegationEnabledTime,omitempty"`
	ServicePrincipal      *string `json:"ServicePrincipal,omitempty" xml:"ServicePrincipal,omitempty"`
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

type ListDelegatedServicesForAccountResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDelegatedServicesForAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	QueryKeyword   *string `json:"QueryKeyword,omitempty" xml:"QueryKeyword,omitempty"`
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
	Folders    *ListFoldersForParentResponseBodyFolders `json:"Folders,omitempty" xml:"Folders,omitempty" type:"Struct"`
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FolderId   *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListFoldersForParentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	Handshakes *ListHandshakesForAccountResponseBodyHandshakes `json:"Handshakes,omitempty" xml:"Handshakes,omitempty" type:"Struct"`
	PageNumber *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpireTime          *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	HandshakeId         *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
	MasterAccountId     *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	MasterAccountName   *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Note                *string `json:"Note,omitempty" xml:"Note,omitempty"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetEntity        *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty"`
	TargetType          *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListHandshakesForAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	Handshakes *ListHandshakesForResourceDirectoryResponseBodyHandshakes `json:"Handshakes,omitempty" xml:"Handshakes,omitempty" type:"Struct"`
	PageNumber *int32                                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpireTime          *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	HandshakeId         *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
	MasterAccountId     *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	MasterAccountName   *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Note                *string `json:"Note,omitempty" xml:"Note,omitempty"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetEntity        *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty"`
	TargetType          *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
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
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListHandshakesForResourceDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Language   *string `json:"Language,omitempty" xml:"Language,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	PageNumber *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Policies   *ListPoliciesResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Struct"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

type ListPolicyAttachmentsRequest struct {
	Language        *string `json:"Language,omitempty" xml:"Language,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PolicyName      *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType      *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	PrincipalName   *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	PrincipalType   *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
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
	PageNumber        *int32                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PolicyAttachments *ListPolicyAttachmentsResponseBodyPolicyAttachments `json:"PolicyAttachments,omitempty" xml:"PolicyAttachments,omitempty" type:"Struct"`
	RequestId         *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount        *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	AttachDate      *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PolicyName      *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType      *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	PrincipalName   *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	PrincipalType   *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPolicyAttachmentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListResourceGroupsRequest struct {
	DisplayName      *string                         `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	IncludeTags      *bool                           `json:"IncludeTags,omitempty" xml:"IncludeTags,omitempty"`
	Name             *string                         `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNumber       *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceGroupId  *string                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceGroupIds []*string                       `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
	Status           *string                         `json:"Status,omitempty" xml:"Status,omitempty"`
	Tag              []*ListResourceGroupsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	PageNumber     *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroups *ListResourceGroupsResponseBodyResourceGroups `json:"ResourceGroups,omitempty" xml:"ResourceGroups,omitempty" type:"Struct"`
	TotalCount     *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	AccountId   *string                                                        `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	CreateDate  *string                                                        `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DisplayName *string                                                        `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Id          *string                                                        `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string                                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	Status      *string                                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags        *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
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
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListResourceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Service         *string `json:"Service,omitempty" xml:"Service,omitempty"`
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

func (s *ListResourcesRequest) SetService(v string) *ListResourcesRequest {
	s.Service = &v
	return s
}

type ListResourcesResponseBody struct {
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources  *ListResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	CreateDate      *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Service         *string `json:"Service,omitempty" xml:"Service,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Language   *string `json:"Language,omitempty" xml:"Language,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	PageNumber *int32                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Roles      *ListRolesResponseBodyRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Struct"`
	TotalCount *int32                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	Arn                 *string                                           `json:"Arn,omitempty" xml:"Arn,omitempty"`
	CreateDate          *string                                           `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	Description         *string                                           `json:"Description,omitempty" xml:"Description,omitempty"`
	IsServiceLinkedRole *bool                                             `json:"IsServiceLinkedRole,omitempty" xml:"IsServiceLinkedRole,omitempty"`
	LatestDeletionTask  *ListRolesResponseBodyRolesRoleLatestDeletionTask `json:"LatestDeletionTask,omitempty" xml:"LatestDeletionTask,omitempty" type:"Struct"`
	MaxSessionDuration  *int64                                            `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty"`
	RoleId              *string                                           `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	RoleName            *string                                           `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	RolePrincipalName   *string                                           `json:"RolePrincipalName,omitempty" xml:"RolePrincipalName,omitempty"`
	UpdateDate          *string                                           `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
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
	CreateDate     *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
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

type ListTagKeysRequest struct {
	KeyFilter    *string `json:"KeyFilter,omitempty" xml:"KeyFilter,omitempty"`
	MaxResults   *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
	NextToken *string                        `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tags      []*ListTagKeysResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	MaxResults   *int32                        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	NextToken    *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	MaxResults   *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	ValueFilter  *string `json:"ValueFilter,omitempty" xml:"ValueFilter,omitempty"`
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
	NextToken *string                          `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tags      []*ListTagValuesResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PolicyId   *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
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
	PageNumber        *int32                                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32                                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId         *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TargetAttachments *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachments `json:"TargetAttachments,omitempty" xml:"TargetAttachments,omitempty" type:"Struct"`
	TotalCount        *int32                                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	AttachDate *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty"`
	TargetId   *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
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
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTargetAttachmentsForControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AdminAccountId *string `json:"AdminAccountId,omitempty" xml:"AdminAccountId,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	EnabledServicePrincipals *ListTrustedServiceStatusResponseBodyEnabledServicePrincipals `json:"EnabledServicePrincipals,omitempty" xml:"EnabledServicePrincipals,omitempty" type:"Struct"`
	PageNumber               *int32                                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                 *int32                                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId                *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount               *int32                                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	EnableTime       *string `json:"EnableTime,omitempty" xml:"EnableTime,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTrustedServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AccountId           *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MoveAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ResourceGroupId *string                          `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Resources       []*MoveResourcesRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
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
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Service      *string `json:"Service,omitempty" xml:"Service,omitempty"`
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
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg     *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Service      *string `json:"Service,omitempty" xml:"Service,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MoveResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	Email     *string `json:"Email,omitempty" xml:"Email,omitempty"`
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
	Account   *PromoteResourceAccountResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccountId           *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AccountName         *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DisplayName         *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	FolderId            *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	JoinMethod          *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	JoinTime            *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	RecordId            *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PromoteResourceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AccountId        *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RegisterDelegatedAdministratorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveCloudAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Account   *ResendCreateCloudAccountEmailResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccountId           *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AccountName         *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DisplayName         *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	FolderId            *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	JoinMethod          *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	JoinTime            *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	RecordId            *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResendCreateCloudAccountEmailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Account   *ResendPromoteResourceAccountEmailResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccountId           *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AccountName         *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DisplayName         *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	FolderId            *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	JoinMethod          *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	JoinTime            *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	RecordId            *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResendPromoteResourceAccountEmailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type SendVerificationCodeForBindSecureMobilePhoneRequest struct {
	AccountId         *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
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
	ExpirationDate *string `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendVerificationCodeForBindSecureMobilePhoneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendVerificationCodeForEnableRDResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type SetMemberDeletionPermissionRequest struct {
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
	ManagementAccountId  *string `json:"ManagementAccountId,omitempty" xml:"ManagementAccountId,omitempty"`
	MemberDeletionStatus *string `json:"MemberDeletionStatus,omitempty" xml:"MemberDeletionStatus,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceDirectoryId  *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetMemberDeletionPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ResourceId   []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AccountId      *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	NewAccountType *string `json:"NewAccountType,omitempty" xml:"NewAccountType,omitempty"`
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
	Account   *UpdateAccountResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccountId           *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AccountName         *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DisplayName         *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	FolderId            *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	JoinMethod          *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	JoinTime            *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type UpdateControlPolicyRequest struct {
	NewDescription    *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	NewPolicyDocument *string `json:"NewPolicyDocument,omitempty" xml:"NewPolicyDocument,omitempty"`
	NewPolicyName     *string `json:"NewPolicyName,omitempty" xml:"NewPolicyName,omitempty"`
	PolicyId          *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
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
	ControlPolicy *UpdateControlPolicyResponseBodyControlPolicy `json:"ControlPolicy,omitempty" xml:"ControlPolicy,omitempty" type:"Struct"`
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AttachmentCount *string `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty"`
	CreateDate      *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EffectScope     *string `json:"EffectScope,omitempty" xml:"EffectScope,omitempty"`
	PolicyId        *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	PolicyName      *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType      *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	UpdateDate      *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	FolderId      *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
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
	Folder    *UpdateFolderResponseBodyFolder `json:"Folder,omitempty" xml:"Folder,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FolderId       *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	FolderName     *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateFolderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	NewDisplayName  *string `json:"NewDisplayName,omitempty" xml:"NewDisplayName,omitempty"`
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
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccountId   *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	CreateDate  *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RolePrincipalName        *string `json:"RolePrincipalName,omitempty" xml:"RolePrincipalName,omitempty"`
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

func (s *UpdateRoleResponseBodyRole) SetRolePrincipalName(v string) *UpdateRoleResponseBodyRole {
	s.RolePrincipalName = &v
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
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":        tea.String("resourcemanager.ap-northeast-1.aliyuncs.com"),
		"ap-south-1":            tea.String("resourcemanager.ap-south-1.aliyuncs.com"),
		"ap-southeast-1":        tea.String("resourcemanager.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        tea.String("resourcemanager.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-3":        tea.String("resourcemanager.ap-southeast-3.aliyuncs.com"),
		"ap-southeast-5":        tea.String("resourcemanager.ap-southeast-5.aliyuncs.com"),
		"cn-beijing":            tea.String("resourcemanager.cn-beijing.aliyuncs.com"),
		"cn-chengdu":            tea.String("resourcemanager.cn-chengdu.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("resourcemanager.cn-hangzhou-finance.aliyuncs.com"),
		"cn-hongkong":           tea.String("resourcemanager.cn-hongkong.aliyuncs.com"),
		"cn-huhehaote":          tea.String("resourcemanager.cn-huhehaote.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("resourcemanager.cn-north-2-gov-1.aliyuncs.com"),
		"cn-qingdao":            tea.String("resourcemanager.cn-qingdao.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("resourcemanager.cn-shanghai-finance-1.aliyuncs.com"),
		"cn-shenzhen":           tea.String("resourcemanager.cn-shenzhen.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("resourcemanager.cn-shenzhen-finance-1.aliyuncs.com"),
		"cn-wulanchabu":         tea.String("resourcemanager.cn-wulanchabu.aliyuncs.com"),
		"cn-zhangjiakou":        tea.String("resourcemanager.cn-zhangjiakou.aliyuncs.com"),
		"eu-central-1":          tea.String("resourcemanager.eu-central-1.aliyuncs.com"),
		"eu-west-1":             tea.String("resourcemanager.eu-west-1.aliyuncs.com"),
		"me-east-1":             tea.String("resourcemanager.me-east-1.aliyuncs.com"),
		"us-east-1":             tea.String("resourcemanager.us-east-1.aliyuncs.com"),
		"us-west-1":             tea.String("resourcemanager.us-west-1.aliyuncs.com"),
	}
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

func (client *Client) GetResourceGroupListAclModeWithOptions(runtime *util.RuntimeOptions) (_result *GetResourceGroupListAclModeResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetResourceGroupListAclMode"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceGroupListAclModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResourceGroupListAclMode() (_result *GetResourceGroupListAclModeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceGroupListAclModeResponse{}
	_body, _err := client.GetResourceGroupListAclModeWithOptions(runtime)
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
