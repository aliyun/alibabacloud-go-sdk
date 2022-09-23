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

type AttachPolicyRequest struct {
	AccountId       *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
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

func (s *AttachPolicyRequest) SetAccountId(v string) *AttachPolicyRequest {
	s.AccountId = &v
	return s
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
	DisplayName               *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email                     *string `json:"Email,omitempty" xml:"Email,omitempty"`
	EnableConsolidatedBilling *bool   `json:"EnableConsolidatedBilling,omitempty" xml:"EnableConsolidatedBilling,omitempty"`
	ParentFolderId            *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	PayerAccountId            *string `json:"PayerAccountId,omitempty" xml:"PayerAccountId,omitempty"`
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

func (s *CreateCloudAccountRequest) SetEnableConsolidatedBilling(v bool) *CreateCloudAccountRequest {
	s.EnableConsolidatedBilling = &v
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
	JoinTime            *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
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

func (s *CreateCloudAccountResponseBodyAccount) SetJoinTime(v string) *CreateCloudAccountResponseBodyAccount {
	s.JoinTime = &v
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

type CreateFolderRequest struct {
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
}

func (s CreateFolderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFolderRequest) GoString() string {
	return s.String()
}

func (s *CreateFolderRequest) SetName(v string) *CreateFolderRequest {
	s.Name = &v
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
	CreateDate     *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	FolderId       *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
}

func (s CreateFolderResponseBodyFolder) String() string {
	return tea.Prettify(s)
}

func (s CreateFolderResponseBodyFolder) GoString() string {
	return s.String()
}

func (s *CreateFolderResponseBodyFolder) SetCreateDate(v string) *CreateFolderResponseBodyFolder {
	s.CreateDate = &v
	return s
}

func (s *CreateFolderResponseBodyFolder) SetFolderId(v string) *CreateFolderResponseBodyFolder {
	s.FolderId = &v
	return s
}

func (s *CreateFolderResponseBodyFolder) SetName(v string) *CreateFolderResponseBodyFolder {
	s.Name = &v
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
	DisplayName               *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	EnableConsolidatedBilling *bool   `json:"EnableConsolidatedBilling,omitempty" xml:"EnableConsolidatedBilling,omitempty"`
	ParentFolderId            *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	PayerAccountId            *string `json:"PayerAccountId,omitempty" xml:"PayerAccountId,omitempty"`
}

func (s CreateResourceAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceAccountRequest) SetDisplayName(v string) *CreateResourceAccountRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateResourceAccountRequest) SetEnableConsolidatedBilling(v bool) *CreateResourceAccountRequest {
	s.EnableConsolidatedBilling = &v
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
	AccountId   *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupRequest) SetAccountId(v string) *CreateResourceGroupRequest {
	s.AccountId = &v
	return s
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

type DeleteInvalidCloudAccountRecordRequest struct {
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s DeleteInvalidCloudAccountRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInvalidCloudAccountRecordRequest) GoString() string {
	return s.String()
}

func (s *DeleteInvalidCloudAccountRecordRequest) SetRecordId(v string) *DeleteInvalidCloudAccountRecordRequest {
	s.RecordId = &v
	return s
}

type DeleteInvalidCloudAccountRecordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInvalidCloudAccountRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInvalidCloudAccountRecordResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInvalidCloudAccountRecordResponseBody) SetRequestId(v string) *DeleteInvalidCloudAccountRecordResponseBody {
	s.RequestId = &v
	return s
}

type DeleteInvalidCloudAccountRecordResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteInvalidCloudAccountRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteInvalidCloudAccountRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInvalidCloudAccountRecordResponse) GoString() string {
	return s.String()
}

func (s *DeleteInvalidCloudAccountRecordResponse) SetHeaders(v map[string]*string) *DeleteInvalidCloudAccountRecordResponse {
	s.Headers = v
	return s
}

func (s *DeleteInvalidCloudAccountRecordResponse) SetStatusCode(v int32) *DeleteInvalidCloudAccountRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInvalidCloudAccountRecordResponse) SetBody(v *DeleteInvalidCloudAccountRecordResponseBody) *DeleteInvalidCloudAccountRecordResponse {
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
	AccountId       *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeleteResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupRequest) SetAccountId(v string) *DeleteResourceGroupRequest {
	s.AccountId = &v
	return s
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

type DestoryResourceDirectoryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DestoryResourceDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DestoryResourceDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *DestoryResourceDirectoryResponseBody) SetRequestId(v string) *DestoryResourceDirectoryResponseBody {
	s.RequestId = &v
	return s
}

type DestoryResourceDirectoryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DestoryResourceDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DestoryResourceDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DestoryResourceDirectoryResponse) GoString() string {
	return s.String()
}

func (s *DestoryResourceDirectoryResponse) SetHeaders(v map[string]*string) *DestoryResourceDirectoryResponse {
	s.Headers = v
	return s
}

func (s *DestoryResourceDirectoryResponse) SetStatusCode(v int32) *DestoryResourceDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DestoryResourceDirectoryResponse) SetBody(v *DestoryResourceDirectoryResponseBody) *DestoryResourceDirectoryResponse {
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

type DetachPolicyRequest struct {
	AccountId       *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
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

func (s *DetachPolicyRequest) SetAccountId(v string) *DetachPolicyRequest {
	s.AccountId = &v
	return s
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

type GetAccountSummaryRequest struct {
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s GetAccountSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccountSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetAccountSummaryRequest) SetAccountId(v string) *GetAccountSummaryRequest {
	s.AccountId = &v
	return s
}

type GetAccountSummaryResponseBody struct {
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SummaryMap *GetAccountSummaryResponseBodySummaryMap `json:"SummaryMap,omitempty" xml:"SummaryMap,omitempty" type:"Struct"`
}

func (s GetAccountSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccountSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountSummaryResponseBody) SetRequestId(v string) *GetAccountSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccountSummaryResponseBody) SetSummaryMap(v *GetAccountSummaryResponseBodySummaryMap) *GetAccountSummaryResponseBody {
	s.SummaryMap = v
	return s
}

type GetAccountSummaryResponseBodySummaryMap struct {
	AttachedPoliciesPerGroupQuota       *int32 `json:"AttachedPoliciesPerGroupQuota,omitempty" xml:"AttachedPoliciesPerGroupQuota,omitempty"`
	AttachedPoliciesPerRoleQuota        *int32 `json:"AttachedPoliciesPerRoleQuota,omitempty" xml:"AttachedPoliciesPerRoleQuota,omitempty"`
	AttachedPoliciesPerUserQuota        *int32 `json:"AttachedPoliciesPerUserQuota,omitempty" xml:"AttachedPoliciesPerUserQuota,omitempty"`
	AttachedSystemPoliciesPerGroupQuota *int32 `json:"AttachedSystemPoliciesPerGroupQuota,omitempty" xml:"AttachedSystemPoliciesPerGroupQuota,omitempty"`
	AttachedSystemPoliciesPerRoleQuota  *int32 `json:"AttachedSystemPoliciesPerRoleQuota,omitempty" xml:"AttachedSystemPoliciesPerRoleQuota,omitempty"`
	AttachedSystemPoliciesPerUserQuota  *int32 `json:"AttachedSystemPoliciesPerUserQuota,omitempty" xml:"AttachedSystemPoliciesPerUserQuota,omitempty"`
	Policies                            *int32 `json:"Policies,omitempty" xml:"Policies,omitempty"`
	PoliciesQuota                       *int32 `json:"PoliciesQuota,omitempty" xml:"PoliciesQuota,omitempty"`
	PolicySizeQuota                     *int32 `json:"PolicySizeQuota,omitempty" xml:"PolicySizeQuota,omitempty"`
	ResourceGroups                      *int32 `json:"ResourceGroups,omitempty" xml:"ResourceGroups,omitempty"`
	ResourceGroupsQuota                 *int32 `json:"ResourceGroupsQuota,omitempty" xml:"ResourceGroupsQuota,omitempty"`
	Roles                               *int32 `json:"Roles,omitempty" xml:"Roles,omitempty"`
	RolesQuota                          *int32 `json:"RolesQuota,omitempty" xml:"RolesQuota,omitempty"`
	VersionsPerPolicyQuota              *int32 `json:"VersionsPerPolicyQuota,omitempty" xml:"VersionsPerPolicyQuota,omitempty"`
}

func (s GetAccountSummaryResponseBodySummaryMap) String() string {
	return tea.Prettify(s)
}

func (s GetAccountSummaryResponseBodySummaryMap) GoString() string {
	return s.String()
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetAttachedPoliciesPerGroupQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.AttachedPoliciesPerGroupQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetAttachedPoliciesPerRoleQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.AttachedPoliciesPerRoleQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetAttachedPoliciesPerUserQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.AttachedPoliciesPerUserQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetAttachedSystemPoliciesPerGroupQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.AttachedSystemPoliciesPerGroupQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetAttachedSystemPoliciesPerRoleQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.AttachedSystemPoliciesPerRoleQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetAttachedSystemPoliciesPerUserQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.AttachedSystemPoliciesPerUserQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetPolicies(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.Policies = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetPoliciesQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.PoliciesQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetPolicySizeQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.PolicySizeQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetResourceGroups(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.ResourceGroups = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetResourceGroupsQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.ResourceGroupsQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetRoles(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.Roles = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetRolesQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.RolesQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetVersionsPerPolicyQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.VersionsPerPolicyQuota = &v
	return s
}

type GetAccountSummaryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAccountSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAccountSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccountSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetAccountSummaryResponse) SetHeaders(v map[string]*string) *GetAccountSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetAccountSummaryResponse) SetStatusCode(v int32) *GetAccountSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountSummaryResponse) SetBody(v *GetAccountSummaryResponseBody) *GetAccountSummaryResponse {
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
	CreateDate     *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	FolderId       *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
}

func (s GetFolderResponseBodyFolder) String() string {
	return tea.Prettify(s)
}

func (s GetFolderResponseBodyFolder) GoString() string {
	return s.String()
}

func (s *GetFolderResponseBodyFolder) SetCreateDate(v string) *GetFolderResponseBodyFolder {
	s.CreateDate = &v
	return s
}

func (s *GetFolderResponseBodyFolder) SetFolderId(v string) *GetFolderResponseBodyFolder {
	s.FolderId = &v
	return s
}

func (s *GetFolderResponseBodyFolder) SetName(v string) *GetFolderResponseBodyFolder {
	s.Name = &v
	return s
}

func (s *GetFolderResponseBodyFolder) SetParentFolderId(v string) *GetFolderResponseBodyFolder {
	s.ParentFolderId = &v
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
	CreateDate          *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	MasterAccountId     *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	MasterAccountName   *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	RootFolderId        *string `json:"RootFolderId,omitempty" xml:"RootFolderId,omitempty"`
}

func (s GetResourceDirectoryResponseBodyResourceDirectory) String() string {
	return tea.Prettify(s)
}

func (s GetResourceDirectoryResponseBodyResourceDirectory) GoString() string {
	return s.String()
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) SetCreateDate(v string) *GetResourceDirectoryResponseBodyResourceDirectory {
	s.CreateDate = &v
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

type GetResourceDirectoryAccountRequest struct {
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s GetResourceDirectoryAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceDirectoryAccountRequest) GoString() string {
	return s.String()
}

func (s *GetResourceDirectoryAccountRequest) SetAccountId(v string) *GetResourceDirectoryAccountRequest {
	s.AccountId = &v
	return s
}

type GetResourceDirectoryAccountResponseBody struct {
	Account   *GetResourceDirectoryAccountResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetResourceDirectoryAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceDirectoryAccountResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceDirectoryAccountResponseBody) SetAccount(v *GetResourceDirectoryAccountResponseBodyAccount) *GetResourceDirectoryAccountResponseBody {
	s.Account = v
	return s
}

func (s *GetResourceDirectoryAccountResponseBody) SetRequestId(v string) *GetResourceDirectoryAccountResponseBody {
	s.RequestId = &v
	return s
}

type GetResourceDirectoryAccountResponseBodyAccount struct {
	AccountId           *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AccountName         *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DisplayName         *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	FolderId            *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	IdentityInformation *string `json:"IdentityInformation,omitempty" xml:"IdentityInformation,omitempty"`
	JoinMethod          *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	JoinTime            *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetResourceDirectoryAccountResponseBodyAccount) String() string {
	return tea.Prettify(s)
}

func (s GetResourceDirectoryAccountResponseBodyAccount) GoString() string {
	return s.String()
}

func (s *GetResourceDirectoryAccountResponseBodyAccount) SetAccountId(v string) *GetResourceDirectoryAccountResponseBodyAccount {
	s.AccountId = &v
	return s
}

func (s *GetResourceDirectoryAccountResponseBodyAccount) SetAccountName(v string) *GetResourceDirectoryAccountResponseBodyAccount {
	s.AccountName = &v
	return s
}

func (s *GetResourceDirectoryAccountResponseBodyAccount) SetDisplayName(v string) *GetResourceDirectoryAccountResponseBodyAccount {
	s.DisplayName = &v
	return s
}

func (s *GetResourceDirectoryAccountResponseBodyAccount) SetFolderId(v string) *GetResourceDirectoryAccountResponseBodyAccount {
	s.FolderId = &v
	return s
}

func (s *GetResourceDirectoryAccountResponseBodyAccount) SetIdentityInformation(v string) *GetResourceDirectoryAccountResponseBodyAccount {
	s.IdentityInformation = &v
	return s
}

func (s *GetResourceDirectoryAccountResponseBodyAccount) SetJoinMethod(v string) *GetResourceDirectoryAccountResponseBodyAccount {
	s.JoinMethod = &v
	return s
}

func (s *GetResourceDirectoryAccountResponseBodyAccount) SetJoinTime(v string) *GetResourceDirectoryAccountResponseBodyAccount {
	s.JoinTime = &v
	return s
}

func (s *GetResourceDirectoryAccountResponseBodyAccount) SetModifyTime(v string) *GetResourceDirectoryAccountResponseBodyAccount {
	s.ModifyTime = &v
	return s
}

func (s *GetResourceDirectoryAccountResponseBodyAccount) SetResourceDirectoryId(v string) *GetResourceDirectoryAccountResponseBodyAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *GetResourceDirectoryAccountResponseBodyAccount) SetStatus(v string) *GetResourceDirectoryAccountResponseBodyAccount {
	s.Status = &v
	return s
}

func (s *GetResourceDirectoryAccountResponseBodyAccount) SetType(v string) *GetResourceDirectoryAccountResponseBodyAccount {
	s.Type = &v
	return s
}

type GetResourceDirectoryAccountResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetResourceDirectoryAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResourceDirectoryAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceDirectoryAccountResponse) GoString() string {
	return s.String()
}

func (s *GetResourceDirectoryAccountResponse) SetHeaders(v map[string]*string) *GetResourceDirectoryAccountResponse {
	s.Headers = v
	return s
}

func (s *GetResourceDirectoryAccountResponse) SetStatusCode(v int32) *GetResourceDirectoryAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceDirectoryAccountResponse) SetBody(v *GetResourceDirectoryAccountResponseBody) *GetResourceDirectoryAccountResponse {
	s.Body = v
	return s
}

type GetResourceGroupRequest struct {
	AccountId       *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s GetResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *GetResourceGroupRequest) SetAccountId(v string) *GetResourceGroupRequest {
	s.AccountId = &v
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

type GetServiceLinkedRoleTemplateRequest struct {
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s GetServiceLinkedRoleTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceLinkedRoleTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleTemplateRequest) SetServiceName(v string) *GetServiceLinkedRoleTemplateRequest {
	s.ServiceName = &v
	return s
}

type GetServiceLinkedRoleTemplateResponseBody struct {
	RequestId                 *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceLinkedRoleTemplate *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate `json:"ServiceLinkedRoleTemplate,omitempty" xml:"ServiceLinkedRoleTemplate,omitempty" type:"Struct"`
}

func (s GetServiceLinkedRoleTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceLinkedRoleTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleTemplateResponseBody) SetRequestId(v string) *GetServiceLinkedRoleTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceLinkedRoleTemplateResponseBody) SetServiceLinkedRoleTemplate(v *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate) *GetServiceLinkedRoleTemplateResponseBody {
	s.ServiceLinkedRoleTemplate = v
	return s
}

type GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate struct {
	MultipleRolesAllowed *bool                                                                              `json:"MultipleRolesAllowed,omitempty" xml:"MultipleRolesAllowed,omitempty"`
	RoleDescriptions     *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptions `json:"RoleDescriptions,omitempty" xml:"RoleDescriptions,omitempty" type:"Struct"`
	RoleNamePrefix       *string                                                                            `json:"RoleNamePrefix,omitempty" xml:"RoleNamePrefix,omitempty"`
	ServiceName          *string                                                                            `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	SystemPolicyName     *string                                                                            `json:"SystemPolicyName,omitempty" xml:"SystemPolicyName,omitempty"`
}

func (s GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate) String() string {
	return tea.Prettify(s)
}

func (s GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate) SetMultipleRolesAllowed(v bool) *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate {
	s.MultipleRolesAllowed = &v
	return s
}

func (s *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate) SetRoleDescriptions(v *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptions) *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate {
	s.RoleDescriptions = v
	return s
}

func (s *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate) SetRoleNamePrefix(v string) *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate {
	s.RoleNamePrefix = &v
	return s
}

func (s *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate) SetServiceName(v string) *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate {
	s.ServiceName = &v
	return s
}

func (s *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate) SetSystemPolicyName(v string) *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate {
	s.SystemPolicyName = &v
	return s
}

type GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptions struct {
	RoleDescription []*GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptionsRoleDescription `json:"RoleDescription,omitempty" xml:"RoleDescription,omitempty" type:"Repeated"`
}

func (s GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptions) String() string {
	return tea.Prettify(s)
}

func (s GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptions) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptions) SetRoleDescription(v []*GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptionsRoleDescription) *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptions {
	s.RoleDescription = v
	return s
}

type GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptionsRoleDescription struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Language    *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptionsRoleDescription) String() string {
	return tea.Prettify(s)
}

func (s GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptionsRoleDescription) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptionsRoleDescription) SetDescription(v string) *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptionsRoleDescription {
	s.Description = &v
	return s
}

func (s *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptionsRoleDescription) SetLanguage(v string) *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptionsRoleDescription {
	s.Language = &v
	return s
}

type GetServiceLinkedRoleTemplateResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetServiceLinkedRoleTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetServiceLinkedRoleTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceLinkedRoleTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleTemplateResponse) SetHeaders(v map[string]*string) *GetServiceLinkedRoleTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetServiceLinkedRoleTemplateResponse) SetStatusCode(v int32) *GetServiceLinkedRoleTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceLinkedRoleTemplateResponse) SetBody(v *GetServiceLinkedRoleTemplateResponseBody) *GetServiceLinkedRoleTemplateResponse {
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
	CreateDate          *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
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

func (s *InitResourceDirectoryResponseBodyResourceDirectory) SetCreateDate(v string) *InitResourceDirectoryResponseBodyResourceDirectory {
	s.CreateDate = &v
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

type ListAccountRecordsForParentRequest struct {
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	QueryKeyword   *string `json:"QueryKeyword,omitempty" xml:"QueryKeyword,omitempty"`
}

func (s ListAccountRecordsForParentRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccountRecordsForParentRequest) GoString() string {
	return s.String()
}

func (s *ListAccountRecordsForParentRequest) SetPageNumber(v int32) *ListAccountRecordsForParentRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAccountRecordsForParentRequest) SetPageSize(v int32) *ListAccountRecordsForParentRequest {
	s.PageSize = &v
	return s
}

func (s *ListAccountRecordsForParentRequest) SetParentFolderId(v string) *ListAccountRecordsForParentRequest {
	s.ParentFolderId = &v
	return s
}

func (s *ListAccountRecordsForParentRequest) SetQueryKeyword(v string) *ListAccountRecordsForParentRequest {
	s.QueryKeyword = &v
	return s
}

type ListAccountRecordsForParentResponseBody struct {
	PageNumber *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Records    *ListAccountRecordsForParentResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Struct"`
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAccountRecordsForParentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAccountRecordsForParentResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccountRecordsForParentResponseBody) SetPageNumber(v int32) *ListAccountRecordsForParentResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAccountRecordsForParentResponseBody) SetPageSize(v int32) *ListAccountRecordsForParentResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAccountRecordsForParentResponseBody) SetRecords(v *ListAccountRecordsForParentResponseBodyRecords) *ListAccountRecordsForParentResponseBody {
	s.Records = v
	return s
}

func (s *ListAccountRecordsForParentResponseBody) SetRequestId(v string) *ListAccountRecordsForParentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccountRecordsForParentResponseBody) SetTotalCount(v int32) *ListAccountRecordsForParentResponseBody {
	s.TotalCount = &v
	return s
}

type ListAccountRecordsForParentResponseBodyRecords struct {
	Record []*ListAccountRecordsForParentResponseBodyRecordsRecord `json:"Record,omitempty" xml:"Record,omitempty" type:"Repeated"`
}

func (s ListAccountRecordsForParentResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s ListAccountRecordsForParentResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListAccountRecordsForParentResponseBodyRecords) SetRecord(v []*ListAccountRecordsForParentResponseBodyRecordsRecord) *ListAccountRecordsForParentResponseBodyRecords {
	s.Record = v
	return s
}

type ListAccountRecordsForParentResponseBodyRecordsRecord struct {
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

func (s ListAccountRecordsForParentResponseBodyRecordsRecord) String() string {
	return tea.Prettify(s)
}

func (s ListAccountRecordsForParentResponseBodyRecordsRecord) GoString() string {
	return s.String()
}

func (s *ListAccountRecordsForParentResponseBodyRecordsRecord) SetAccountId(v string) *ListAccountRecordsForParentResponseBodyRecordsRecord {
	s.AccountId = &v
	return s
}

func (s *ListAccountRecordsForParentResponseBodyRecordsRecord) SetAccountName(v string) *ListAccountRecordsForParentResponseBodyRecordsRecord {
	s.AccountName = &v
	return s
}

func (s *ListAccountRecordsForParentResponseBodyRecordsRecord) SetDisplayName(v string) *ListAccountRecordsForParentResponseBodyRecordsRecord {
	s.DisplayName = &v
	return s
}

func (s *ListAccountRecordsForParentResponseBodyRecordsRecord) SetFolderId(v string) *ListAccountRecordsForParentResponseBodyRecordsRecord {
	s.FolderId = &v
	return s
}

func (s *ListAccountRecordsForParentResponseBodyRecordsRecord) SetJoinMethod(v string) *ListAccountRecordsForParentResponseBodyRecordsRecord {
	s.JoinMethod = &v
	return s
}

func (s *ListAccountRecordsForParentResponseBodyRecordsRecord) SetJoinTime(v string) *ListAccountRecordsForParentResponseBodyRecordsRecord {
	s.JoinTime = &v
	return s
}

func (s *ListAccountRecordsForParentResponseBodyRecordsRecord) SetModifyTime(v string) *ListAccountRecordsForParentResponseBodyRecordsRecord {
	s.ModifyTime = &v
	return s
}

func (s *ListAccountRecordsForParentResponseBodyRecordsRecord) SetRecordId(v string) *ListAccountRecordsForParentResponseBodyRecordsRecord {
	s.RecordId = &v
	return s
}

func (s *ListAccountRecordsForParentResponseBodyRecordsRecord) SetResourceDirectoryId(v string) *ListAccountRecordsForParentResponseBodyRecordsRecord {
	s.ResourceDirectoryId = &v
	return s
}

func (s *ListAccountRecordsForParentResponseBodyRecordsRecord) SetStatus(v string) *ListAccountRecordsForParentResponseBodyRecordsRecord {
	s.Status = &v
	return s
}

func (s *ListAccountRecordsForParentResponseBodyRecordsRecord) SetType(v string) *ListAccountRecordsForParentResponseBodyRecordsRecord {
	s.Type = &v
	return s
}

type ListAccountRecordsForParentResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAccountRecordsForParentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAccountRecordsForParentResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccountRecordsForParentResponse) GoString() string {
	return s.String()
}

func (s *ListAccountRecordsForParentResponse) SetHeaders(v map[string]*string) *ListAccountRecordsForParentResponse {
	s.Headers = v
	return s
}

func (s *ListAccountRecordsForParentResponse) SetStatusCode(v int32) *ListAccountRecordsForParentResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccountRecordsForParentResponse) SetBody(v *ListAccountRecordsForParentResponseBody) *ListAccountRecordsForParentResponse {
	s.Body = v
	return s
}

type ListAccountsRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListAccountsRequest) SetPageNumber(v int32) *ListAccountsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAccountsRequest) SetPageSize(v int32) *ListAccountsRequest {
	s.PageSize = &v
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
	AccountId           *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	DisplayName         *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	FolderId            *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	JoinMethod          *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	JoinTime            *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
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

func (s *ListAccountsResponseBodyAccountsAccount) SetStatus(v string) *ListAccountsResponseBodyAccountsAccount {
	s.Status = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetType(v string) *ListAccountsResponseBodyAccountsAccount {
	s.Type = &v
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
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	QueryKeyword   *string `json:"QueryKeyword,omitempty" xml:"QueryKeyword,omitempty"`
}

func (s ListAccountsForParentRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsForParentRequest) GoString() string {
	return s.String()
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
	AccountId           *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	DisplayName         *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	FolderId            *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	JoinMethod          *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	JoinTime            *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
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

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetType(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.Type = &v
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
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	FolderId   *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListAncestorsResponseBodyFoldersFolder) String() string {
	return tea.Prettify(s)
}

func (s ListAncestorsResponseBodyFoldersFolder) GoString() string {
	return s.String()
}

func (s *ListAncestorsResponseBodyFoldersFolder) SetCreateDate(v string) *ListAncestorsResponseBodyFoldersFolder {
	s.CreateDate = &v
	return s
}

func (s *ListAncestorsResponseBodyFoldersFolder) SetFolderId(v string) *ListAncestorsResponseBodyFoldersFolder {
	s.FolderId = &v
	return s
}

func (s *ListAncestorsResponseBodyFoldersFolder) SetName(v string) *ListAncestorsResponseBodyFoldersFolder {
	s.Name = &v
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
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	FolderId   *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListFoldersForParentResponseBodyFoldersFolder) String() string {
	return tea.Prettify(s)
}

func (s ListFoldersForParentResponseBodyFoldersFolder) GoString() string {
	return s.String()
}

func (s *ListFoldersForParentResponseBodyFoldersFolder) SetCreateDate(v string) *ListFoldersForParentResponseBodyFoldersFolder {
	s.CreateDate = &v
	return s
}

func (s *ListFoldersForParentResponseBodyFoldersFolder) SetFolderId(v string) *ListFoldersForParentResponseBodyFoldersFolder {
	s.FolderId = &v
	return s
}

func (s *ListFoldersForParentResponseBodyFoldersFolder) SetName(v string) *ListFoldersForParentResponseBodyFoldersFolder {
	s.Name = &v
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

type ListParentsRequest struct {
	ChildId *string `json:"ChildId,omitempty" xml:"ChildId,omitempty"`
}

func (s ListParentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListParentsRequest) GoString() string {
	return s.String()
}

func (s *ListParentsRequest) SetChildId(v string) *ListParentsRequest {
	s.ChildId = &v
	return s
}

type ListParentsResponseBody struct {
	Folders   *ListParentsResponseBodyFolders `json:"Folders,omitempty" xml:"Folders,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListParentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListParentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListParentsResponseBody) SetFolders(v *ListParentsResponseBodyFolders) *ListParentsResponseBody {
	s.Folders = v
	return s
}

func (s *ListParentsResponseBody) SetRequestId(v string) *ListParentsResponseBody {
	s.RequestId = &v
	return s
}

type ListParentsResponseBodyFolders struct {
	Folder []*ListParentsResponseBodyFoldersFolder `json:"Folder,omitempty" xml:"Folder,omitempty" type:"Repeated"`
}

func (s ListParentsResponseBodyFolders) String() string {
	return tea.Prettify(s)
}

func (s ListParentsResponseBodyFolders) GoString() string {
	return s.String()
}

func (s *ListParentsResponseBodyFolders) SetFolder(v []*ListParentsResponseBodyFoldersFolder) *ListParentsResponseBodyFolders {
	s.Folder = v
	return s
}

type ListParentsResponseBodyFoldersFolder struct {
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	FolderId   *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListParentsResponseBodyFoldersFolder) String() string {
	return tea.Prettify(s)
}

func (s ListParentsResponseBodyFoldersFolder) GoString() string {
	return s.String()
}

func (s *ListParentsResponseBodyFoldersFolder) SetCreateDate(v string) *ListParentsResponseBodyFoldersFolder {
	s.CreateDate = &v
	return s
}

func (s *ListParentsResponseBodyFoldersFolder) SetFolderId(v string) *ListParentsResponseBodyFoldersFolder {
	s.FolderId = &v
	return s
}

func (s *ListParentsResponseBodyFoldersFolder) SetName(v string) *ListParentsResponseBodyFoldersFolder {
	s.Name = &v
	return s
}

type ListParentsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListParentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListParentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListParentsResponse) GoString() string {
	return s.String()
}

func (s *ListParentsResponse) SetHeaders(v map[string]*string) *ListParentsResponse {
	s.Headers = v
	return s
}

func (s *ListParentsResponse) SetStatusCode(v int32) *ListParentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListParentsResponse) SetBody(v *ListParentsResponseBody) *ListParentsResponse {
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
	AccountId       *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
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

func (s *ListPolicyAttachmentsRequest) SetAccountId(v string) *ListPolicyAttachmentsRequest {
	s.AccountId = &v
	return s
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
	AccountId  *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListResourceGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsRequest) SetAccountId(v string) *ListResourceGroupsRequest {
	s.AccountId = &v
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

func (s *ListResourceGroupsRequest) SetStatus(v string) *ListResourceGroupsRequest {
	s.Status = &v
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
	AccountId   *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	CreateDate  *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

type ListRolesForServiceRequest struct {
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Service  *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s ListRolesForServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRolesForServiceRequest) GoString() string {
	return s.String()
}

func (s *ListRolesForServiceRequest) SetLanguage(v string) *ListRolesForServiceRequest {
	s.Language = &v
	return s
}

func (s *ListRolesForServiceRequest) SetService(v string) *ListRolesForServiceRequest {
	s.Service = &v
	return s
}

type ListRolesForServiceResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Roles     *ListRolesForServiceResponseBodyRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Struct"`
}

func (s ListRolesForServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRolesForServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ListRolesForServiceResponseBody) SetRequestId(v string) *ListRolesForServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRolesForServiceResponseBody) SetRoles(v *ListRolesForServiceResponseBodyRoles) *ListRolesForServiceResponseBody {
	s.Roles = v
	return s
}

type ListRolesForServiceResponseBodyRoles struct {
	Role []*ListRolesForServiceResponseBodyRolesRole `json:"Role,omitempty" xml:"Role,omitempty" type:"Repeated"`
}

func (s ListRolesForServiceResponseBodyRoles) String() string {
	return tea.Prettify(s)
}

func (s ListRolesForServiceResponseBodyRoles) GoString() string {
	return s.String()
}

func (s *ListRolesForServiceResponseBodyRoles) SetRole(v []*ListRolesForServiceResponseBodyRolesRole) *ListRolesForServiceResponseBodyRoles {
	s.Role = v
	return s
}

type ListRolesForServiceResponseBodyRolesRole struct {
	Arn                 *string                                                     `json:"Arn,omitempty" xml:"Arn,omitempty"`
	CreateDate          *string                                                     `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	Description         *string                                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	IsServiceLinkedRole *bool                                                       `json:"IsServiceLinkedRole,omitempty" xml:"IsServiceLinkedRole,omitempty"`
	LatestDeletionTask  *ListRolesForServiceResponseBodyRolesRoleLatestDeletionTask `json:"LatestDeletionTask,omitempty" xml:"LatestDeletionTask,omitempty" type:"Struct"`
	MaxSessionDuration  *int64                                                      `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty"`
	RoleId              *string                                                     `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	RoleName            *string                                                     `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	RolePrincipalName   *string                                                     `json:"RolePrincipalName,omitempty" xml:"RolePrincipalName,omitempty"`
	UpdateDate          *string                                                     `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListRolesForServiceResponseBodyRolesRole) String() string {
	return tea.Prettify(s)
}

func (s ListRolesForServiceResponseBodyRolesRole) GoString() string {
	return s.String()
}

func (s *ListRolesForServiceResponseBodyRolesRole) SetArn(v string) *ListRolesForServiceResponseBodyRolesRole {
	s.Arn = &v
	return s
}

func (s *ListRolesForServiceResponseBodyRolesRole) SetCreateDate(v string) *ListRolesForServiceResponseBodyRolesRole {
	s.CreateDate = &v
	return s
}

func (s *ListRolesForServiceResponseBodyRolesRole) SetDescription(v string) *ListRolesForServiceResponseBodyRolesRole {
	s.Description = &v
	return s
}

func (s *ListRolesForServiceResponseBodyRolesRole) SetIsServiceLinkedRole(v bool) *ListRolesForServiceResponseBodyRolesRole {
	s.IsServiceLinkedRole = &v
	return s
}

func (s *ListRolesForServiceResponseBodyRolesRole) SetLatestDeletionTask(v *ListRolesForServiceResponseBodyRolesRoleLatestDeletionTask) *ListRolesForServiceResponseBodyRolesRole {
	s.LatestDeletionTask = v
	return s
}

func (s *ListRolesForServiceResponseBodyRolesRole) SetMaxSessionDuration(v int64) *ListRolesForServiceResponseBodyRolesRole {
	s.MaxSessionDuration = &v
	return s
}

func (s *ListRolesForServiceResponseBodyRolesRole) SetRoleId(v string) *ListRolesForServiceResponseBodyRolesRole {
	s.RoleId = &v
	return s
}

func (s *ListRolesForServiceResponseBodyRolesRole) SetRoleName(v string) *ListRolesForServiceResponseBodyRolesRole {
	s.RoleName = &v
	return s
}

func (s *ListRolesForServiceResponseBodyRolesRole) SetRolePrincipalName(v string) *ListRolesForServiceResponseBodyRolesRole {
	s.RolePrincipalName = &v
	return s
}

func (s *ListRolesForServiceResponseBodyRolesRole) SetUpdateDate(v string) *ListRolesForServiceResponseBodyRolesRole {
	s.UpdateDate = &v
	return s
}

type ListRolesForServiceResponseBodyRolesRoleLatestDeletionTask struct {
	CreateDate     *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DeletionTaskId *string `json:"DeletionTaskId,omitempty" xml:"DeletionTaskId,omitempty"`
}

func (s ListRolesForServiceResponseBodyRolesRoleLatestDeletionTask) String() string {
	return tea.Prettify(s)
}

func (s ListRolesForServiceResponseBodyRolesRoleLatestDeletionTask) GoString() string {
	return s.String()
}

func (s *ListRolesForServiceResponseBodyRolesRoleLatestDeletionTask) SetCreateDate(v string) *ListRolesForServiceResponseBodyRolesRoleLatestDeletionTask {
	s.CreateDate = &v
	return s
}

func (s *ListRolesForServiceResponseBodyRolesRoleLatestDeletionTask) SetDeletionTaskId(v string) *ListRolesForServiceResponseBodyRolesRoleLatestDeletionTask {
	s.DeletionTaskId = &v
	return s
}

type ListRolesForServiceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRolesForServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRolesForServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRolesForServiceResponse) GoString() string {
	return s.String()
}

func (s *ListRolesForServiceResponse) SetHeaders(v map[string]*string) *ListRolesForServiceResponse {
	s.Headers = v
	return s
}

func (s *ListRolesForServiceResponse) SetStatusCode(v int32) *ListRolesForServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRolesForServiceResponse) SetBody(v *ListRolesForServiceResponseBody) *ListRolesForServiceResponse {
	s.Body = v
	return s
}

type ListTrustedServiceStatusRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListTrustedServiceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTrustedServiceStatusRequest) GoString() string {
	return s.String()
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
	RecordId  *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PromoteResourceAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PromoteResourceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *PromoteResourceAccountResponseBody) SetRecordId(v string) *PromoteResourceAccountResponseBody {
	s.RecordId = &v
	return s
}

func (s *PromoteResourceAccountResponseBody) SetRequestId(v string) *PromoteResourceAccountResponseBody {
	s.RequestId = &v
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

type QueryResourceRequest struct {
	AccountId       *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Service         *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s QueryResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryResourceRequest) GoString() string {
	return s.String()
}

func (s *QueryResourceRequest) SetAccountId(v string) *QueryResourceRequest {
	s.AccountId = &v
	return s
}

func (s *QueryResourceRequest) SetPageNumber(v int32) *QueryResourceRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryResourceRequest) SetPageSize(v int32) *QueryResourceRequest {
	s.PageSize = &v
	return s
}

func (s *QueryResourceRequest) SetRegion(v string) *QueryResourceRequest {
	s.Region = &v
	return s
}

func (s *QueryResourceRequest) SetResourceGroupId(v string) *QueryResourceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *QueryResourceRequest) SetResourceId(v string) *QueryResourceRequest {
	s.ResourceId = &v
	return s
}

func (s *QueryResourceRequest) SetResourceType(v string) *QueryResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *QueryResourceRequest) SetService(v string) *QueryResourceRequest {
	s.Service = &v
	return s
}

type QueryResourceResponseBody struct {
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources  *QueryResourceResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryResourceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryResourceResponseBody) SetPageNumber(v int32) *QueryResourceResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryResourceResponseBody) SetPageSize(v int32) *QueryResourceResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryResourceResponseBody) SetRequestId(v string) *QueryResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryResourceResponseBody) SetResources(v *QueryResourceResponseBodyResources) *QueryResourceResponseBody {
	s.Resources = v
	return s
}

func (s *QueryResourceResponseBody) SetTotalCount(v int32) *QueryResourceResponseBody {
	s.TotalCount = &v
	return s
}

type QueryResourceResponseBodyResources struct {
	Resource []*QueryResourceResponseBodyResourcesResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s QueryResourceResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s QueryResourceResponseBodyResources) GoString() string {
	return s.String()
}

func (s *QueryResourceResponseBodyResources) SetResource(v []*QueryResourceResponseBodyResourcesResource) *QueryResourceResponseBodyResources {
	s.Resource = v
	return s
}

type QueryResourceResponseBodyResourcesResource struct {
	CreateDate      *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Service         *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s QueryResourceResponseBodyResourcesResource) String() string {
	return tea.Prettify(s)
}

func (s QueryResourceResponseBodyResourcesResource) GoString() string {
	return s.String()
}

func (s *QueryResourceResponseBodyResourcesResource) SetCreateDate(v string) *QueryResourceResponseBodyResourcesResource {
	s.CreateDate = &v
	return s
}

func (s *QueryResourceResponseBodyResourcesResource) SetRegionId(v string) *QueryResourceResponseBodyResourcesResource {
	s.RegionId = &v
	return s
}

func (s *QueryResourceResponseBodyResourcesResource) SetResourceGroupId(v string) *QueryResourceResponseBodyResourcesResource {
	s.ResourceGroupId = &v
	return s
}

func (s *QueryResourceResponseBodyResourcesResource) SetResourceId(v string) *QueryResourceResponseBodyResourcesResource {
	s.ResourceId = &v
	return s
}

func (s *QueryResourceResponseBodyResourcesResource) SetResourceType(v string) *QueryResourceResponseBodyResourcesResource {
	s.ResourceType = &v
	return s
}

func (s *QueryResourceResponseBodyResourcesResource) SetService(v string) *QueryResourceResponseBodyResourcesResource {
	s.Service = &v
	return s
}

type QueryResourceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryResourceResponse) GoString() string {
	return s.String()
}

func (s *QueryResourceResponse) SetHeaders(v map[string]*string) *QueryResourceResponse {
	s.Headers = v
	return s
}

func (s *QueryResourceResponse) SetStatusCode(v int32) *QueryResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryResourceResponse) SetBody(v *QueryResourceResponseBody) *QueryResourceResponse {
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
	RecordId  *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResendCreateCloudAccountEmailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResendCreateCloudAccountEmailResponseBody) GoString() string {
	return s.String()
}

func (s *ResendCreateCloudAccountEmailResponseBody) SetRecordId(v string) *ResendCreateCloudAccountEmailResponseBody {
	s.RecordId = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseBody) SetRequestId(v string) *ResendCreateCloudAccountEmailResponseBody {
	s.RequestId = &v
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
	RecordId  *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResendPromoteResourceAccountEmailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResendPromoteResourceAccountEmailResponseBody) GoString() string {
	return s.String()
}

func (s *ResendPromoteResourceAccountEmailResponseBody) SetRecordId(v string) *ResendPromoteResourceAccountEmailResponseBody {
	s.RecordId = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseBody) SetRequestId(v string) *ResendPromoteResourceAccountEmailResponseBody {
	s.RequestId = &v
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

type UpdateFolderRequest struct {
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

func (s *UpdateFolderRequest) SetName(v string) *UpdateFolderRequest {
	s.Name = &v
	return s
}

type UpdateFolderResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFolderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFolderResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFolderResponseBody) SetRequestId(v string) *UpdateFolderResponseBody {
	s.RequestId = &v
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
	AccountId       *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	NewDisplayName  *string `json:"NewDisplayName,omitempty" xml:"NewDisplayName,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s UpdateResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupRequest) SetAccountId(v string) *UpdateResourceGroupRequest {
	s.AccountId = &v
	return s
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

func (client *Client) AttachPolicyWithOptions(request *AttachPolicyRequest, runtime *util.RuntimeOptions) (_result *AttachPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
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
		Action:      tea.String("AttachPolicy"),
		Version:     tea.String("2016-11-11"),
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
		Version:     tea.String("2016-11-11"),
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
		Version:     tea.String("2016-11-11"),
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

	if !tea.BoolValue(util.IsUnset(request.EnableConsolidatedBilling)) {
		query["EnableConsolidatedBilling"] = request.EnableConsolidatedBilling
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
		Version:     tea.String("2016-11-11"),
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

func (client *Client) CreateFolderWithOptions(request *CreateFolderRequest, runtime *util.RuntimeOptions) (_result *CreateFolderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ParentFolderId)) {
		query["ParentFolderId"] = request.ParentFolderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFolder"),
		Version:     tea.String("2016-11-11"),
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
		Version:     tea.String("2016-11-11"),
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
		Version:     tea.String("2016-11-11"),
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
	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.EnableConsolidatedBilling)) {
		query["EnableConsolidatedBilling"] = request.EnableConsolidatedBilling
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
		Action:      tea.String("CreateResourceAccount"),
		Version:     tea.String("2016-11-11"),
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
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

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
		Version:     tea.String("2016-11-11"),
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
		Version:     tea.String("2016-11-11"),
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
		Version:     tea.String("2016-11-11"),
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
		Version:     tea.String("2016-11-11"),
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

func (client *Client) DeleteInvalidCloudAccountRecordWithOptions(request *DeleteInvalidCloudAccountRecordRequest, runtime *util.RuntimeOptions) (_result *DeleteInvalidCloudAccountRecordResponse, _err error) {
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
		Action:      tea.String("DeleteInvalidCloudAccountRecord"),
		Version:     tea.String("2016-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInvalidCloudAccountRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteInvalidCloudAccountRecord(request *DeleteInvalidCloudAccountRecordRequest) (_result *DeleteInvalidCloudAccountRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteInvalidCloudAccountRecordResponse{}
	_body, _err := client.DeleteInvalidCloudAccountRecordWithOptions(request, runtime)
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
		Version:     tea.String("2016-11-11"),
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
		Version:     tea.String("2016-11-11"),
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
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteResourceGroup"),
		Version:     tea.String("2016-11-11"),
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
		Version:     tea.String("2016-11-11"),
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
		Version:     tea.String("2016-11-11"),
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

func (client *Client) DestoryResourceDirectoryWithOptions(runtime *util.RuntimeOptions) (_result *DestoryResourceDirectoryResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DestoryResourceDirectory"),
		Version:     tea.String("2016-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DestoryResourceDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DestoryResourceDirectory() (_result *DestoryResourceDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DestoryResourceDirectoryResponse{}
	_body, _err := client.DestoryResourceDirectoryWithOptions(runtime)
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
		Version:     tea.String("2016-11-11"),
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

func (client *Client) DetachPolicyWithOptions(request *DetachPolicyRequest, runtime *util.RuntimeOptions) (_result *DetachPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
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
		Action:      tea.String("DetachPolicy"),
		Version:     tea.String("2016-11-11"),
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

func (client *Client) GetAccountSummaryWithOptions(request *GetAccountSummaryRequest, runtime *util.RuntimeOptions) (_result *GetAccountSummaryResponse, _err error) {
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
		Action:      tea.String("GetAccountSummary"),
		Version:     tea.String("2016-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccountSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAccountSummary(request *GetAccountSummaryRequest) (_result *GetAccountSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccountSummaryResponse{}
	_body, _err := client.GetAccountSummaryWithOptions(request, runtime)
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
		Version:     tea.String("2016-11-11"),
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
		Version:     tea.String("2016-11-11"),
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
		Version:     tea.String("2016-11-11"),
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
		Version:     tea.String("2016-11-11"),
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

func (client *Client) GetResourceDirectoryAccountWithOptions(request *GetResourceDirectoryAccountRequest, runtime *util.RuntimeOptions) (_result *GetResourceDirectoryAccountResponse, _err error) {
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
		Action:      tea.String("GetResourceDirectoryAccount"),
		Version:     tea.String("2016-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceDirectoryAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResourceDirectoryAccount(request *GetResourceDirectoryAccountRequest) (_result *GetResourceDirectoryAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceDirectoryAccountResponse{}
	_body, _err := client.GetResourceDirectoryAccountWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceGroup"),
		Version:     tea.String("2016-11-11"),
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
		Version:     tea.String("2016-11-11"),
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
		Version:     tea.String("2016-11-11"),
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

func (client *Client) GetServiceLinkedRoleTemplateWithOptions(request *GetServiceLinkedRoleTemplateRequest, runtime *util.RuntimeOptions) (_result *GetServiceLinkedRoleTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceLinkedRoleTemplate"),
		Version:     tea.String("2016-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceLinkedRoleTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetServiceLinkedRoleTemplate(request *GetServiceLinkedRoleTemplateRequest) (_result *GetServiceLinkedRoleTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServiceLinkedRoleTemplateResponse{}
	_body, _err := client.GetServiceLinkedRoleTemplateWithOptions(request, runtime)
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
		Version:     tea.String("2016-11-11"),
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

func (client *Client) ListAccountRecordsForParentWithOptions(request *ListAccountRecordsForParentRequest, runtime *util.RuntimeOptions) (_result *ListAccountRecordsForParentResponse, _err error) {
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
		Action:      tea.String("ListAccountRecordsForParent"),
		Version:     tea.String("2016-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAccountRecordsForParentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAccountRecordsForParent(request *ListAccountRecordsForParentRequest) (_result *ListAccountRecordsForParentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAccountRecordsForParentResponse{}
	_body, _err := client.ListAccountRecordsForParentWithOptions(request, runtime)
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
		Action:      tea.String("ListAccounts"),
		Version:     tea.String("2016-11-11"),
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
		Action:      tea.String("ListAccountsForParent"),
		Version:     tea.String("2016-11-11"),
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
		Version:     tea.String("2016-11-11"),
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
		Version:     tea.String("2016-11-11"),
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

func (client *Client) ListParentsWithOptions(request *ListParentsRequest, runtime *util.RuntimeOptions) (_result *ListParentsResponse, _err error) {
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
		Action:      tea.String("ListParents"),
		Version:     tea.String("2016-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListParentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListParents(request *ListParentsRequest) (_result *ListParentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListParentsResponse{}
	_body, _err := client.ListParentsWithOptions(request, runtime)
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
		Version:     tea.String("2016-11-11"),
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
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

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
		Version:     tea.String("2016-11-11"),
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
		Version:     tea.String("2016-11-11"),
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
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceGroups"),
		Version:     tea.String("2016-11-11"),
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
		Version:     tea.String("2016-11-11"),
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

func (client *Client) ListRolesForServiceWithOptions(request *ListRolesForServiceRequest, runtime *util.RuntimeOptions) (_result *ListRolesForServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.Service)) {
		query["Service"] = request.Service
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRolesForService"),
		Version:     tea.String("2016-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRolesForServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRolesForService(request *ListRolesForServiceRequest) (_result *ListRolesForServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRolesForServiceResponse{}
	_body, _err := client.ListRolesForServiceWithOptions(request, runtime)
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
		Version:     tea.String("2016-11-11"),
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
		Version:     tea.String("2016-11-11"),
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
		Version:     tea.String("2016-11-11"),
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

func (client *Client) QueryResourceWithOptions(request *QueryResourceRequest, runtime *util.RuntimeOptions) (_result *QueryResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

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
		Action:      tea.String("QueryResource"),
		Version:     tea.String("2016-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryResource(request *QueryResourceRequest) (_result *QueryResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryResourceResponse{}
	_body, _err := client.QueryResourceWithOptions(request, runtime)
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
		Version:     tea.String("2016-11-11"),
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
		Version:     tea.String("2016-11-11"),
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
		Version:     tea.String("2016-11-11"),
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
		Version:     tea.String("2016-11-11"),
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

func (client *Client) UpdateFolderWithOptions(request *UpdateFolderRequest, runtime *util.RuntimeOptions) (_result *UpdateFolderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FolderId)) {
		query["FolderId"] = request.FolderId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFolder"),
		Version:     tea.String("2016-11-11"),
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
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

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
		Version:     tea.String("2016-11-11"),
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
		Version:     tea.String("2016-11-11"),
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
